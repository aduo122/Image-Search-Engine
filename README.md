# Image-Search
Project Structure

Backend: Go, Redis, REST
- Use Claifai API (https://clarifai.com/developer/guide/) to analyze image and get tags
- Create Revert Index, using tag to search image url
- Store the key-value pair to redis using "Zadd" function, sort the value by ranking Value

Frontend: Node.js, Express.js, Handlebars
- Use "zrevrange" to retrive top 10 images of the tag
