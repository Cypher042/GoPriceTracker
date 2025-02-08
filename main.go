package main

import ("fmt"
		"log"
		"bufio"
		"os"
		"github.com/gocolly/colly"	
	)

func main() {

	scraper := colly.NewCollector()

	var URLS []string

	file, err := os.Open("urls.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()
	scanner := bufio.NewScanner(file)
1``

	for scanner.Scan() {
		URLS = append(URLS, scanner.Text())
	}
	
	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}


	// := "https://www.amazon.in/..."
	

	scraper.OnRequest(func(r *colly.Request) {
		r.Headers.Set("User-Agent", "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/114.0.0.0 Safari/537.36")
		r.Headers.Set("Accept-Language", "en-IN,en;q=0.9,hi;q=0.8")
		r.Headers.Set("X-Forwarded-For", "103.211.212.105") 
		r.Headers.Set("Cookie", "session=idk; region=IN")
		r.Headers.Set("Referer", "https://www.google.co.in/")
		// fmt.Printf("Visiting... %s\n", r.URL)
	})

	scraper.OnHTML("span.a-price", func(e *colly.HTMLElement) {
		
		if e.Index == 5{
			priceWhole := e.ChildText("span.a-price-whole")
			priceSymbol := e.ChildText("span.a-price-symbol")
			fmt.Printf("Price: %s%s\n", priceSymbol, priceWhole)
		}
	})

	scraper.OnError(func(r *colly.Response, err error) {
		log.Println("Error:", err)
	})

	for (i := 0; i < len(URLS); i++) {
		
		err := scraper.Visit(URLS[i])
		if err != nil {
			log.Fatalf("Failed to visit the website: %v", err)
		
		}
	}
}
