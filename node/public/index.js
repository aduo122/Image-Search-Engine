var argv = require('minimist')(process.argv.slice(2));
var port = argv['port'];
var redis_host = argv['redis_host'];
var redis_port = argv['redis_port'];
var redis_channel = argv['redis_channel'];

// Setup dependency instances.
var express = require('express');
var app = express();
var server = require('http').createServer(app);
// var io = require('socket.io')(server);

// Setup redis client.
var redis = require('redis');
console.log('Creating a redis client');
var redisClient = redis.createClient(redis_port, redis_host);
console.log('Subscribing to redis topic: %s', redis_channel);

function searchFunc() {
	var input, filter, ul, li, a, i;
	input = document.getElementById('myinput');
	filter = input.value.toUpperCase();
	// get ul from redis
	ul = document.getElementById('wrapper');
	li = ul.getElementByTagName('li');

	for(i=0; i<li.length; i++){
		a = li[i].getElementByTagName('a')[0];
		if (a.innerHTML.toUpperCase().indexOf(filter) > -1){
			li[i].style.display = "";
		}

		else{
			li[i].style.display = 'none';
		}
	}
}
