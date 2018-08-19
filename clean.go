package main

import (

	"fmt"
	// "net/http"
	// "strings"
	// "time"
	// "reflect"
	"github.com/go-redis/redis"
)

func main(){
	redisClient := redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		Password: "", // no password set
		DB:       0,  // use default DB
	})
	pong, err := redisClient.Ping().Result()
	fmt.Println(pong, err)
	redisClient.FlushAll()
}
