var express = require('express');
var bodyParser = require('body-parser');
var cors = require('cors');
var path = require('path');
const { title } = require('process');

var app = express();

app.use(cors());
app.use(bodyParser.json());

app.set('views', path.join(__dirname, 'views'));
app.set('view engine', 'ejs');

app.get('/', function(req, res) {
    res.render('index', {
        title: 'Hello World'
    });
});

app.listen(7000, function() {
  console.log('Server started on port 7000');
});