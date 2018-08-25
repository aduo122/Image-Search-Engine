# Image-Search
Project Structure

Backend: Go, Redis, REST
- Use Claifai API (https://clarifai.com/developer/guide/) to analyze image and get tags
- Create Revert Index, using tag to search image url
- Store the key-value pair to redis using "Zadd" function, sort the value by ranking Value

Frontend: Node.js, Express.js, Handlebars
- Use "zrevrange" to retrive top 10 images of the tag

under Image_Search_engine/app/
  docker build -t go-app .

under Image_Search_engine/node/
  docker build -t node-web .

under Image_Search_engine/
  docker-compose up

in browser, visit page
  localhost:3000
