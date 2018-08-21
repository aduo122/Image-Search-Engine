const express = require('express');
const exphbs = require('express-handlebars');
const path = require('path');
const bodyParser = require('body-parser');
const redis = require('redis')

// create redis client
var redis_host = process.env.REDIS_HOST || '127.0.0.1';
var redis_port = 6379;
let client = redis.createClient(redis_port, redis_host);

client.on('connect', function(){
	console.log('Connected to Redis...')
})

//set port
const port = 3000;

// init app
const app = express();

// view engine
app.engine('handlebars', exphbs({defaultLayout:'main'}));
app.set('view engine', 'handlebars');

// body-parser
app.use(bodyParser.json());
app.use(bodyParser.urlencoded({extended:false}));

// search page
app.get('/', function(req, res, next){
	res.render('searchusers');
});

// search processing
app.post('/picture/search', function(req, res, next){
	let tag = req.body.tag.toLowerCase();
	client.zrevrange(tag, 0, 9, function(err, obj){
		if(err){
			res.render('searchusers', {
				error: 'Picture of '+ tag + ' does not exist'
			});
		}else{
			obj.id = tag;
			console.log(obj);
			res.render('details', {
				picture: obj
			});
		}
	});
});

app.listen(port, function(){
	console.log('Server started on port '+port);
});
