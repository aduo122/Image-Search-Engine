package main

import (
	"bytes"
	// "encoding/json"
	"fmt"
	"io/ioutil"
	// "log"
	"net/http"
	// "strings"
	// "time"

	// "github.com/go-redis/redis"
)

const TAR string = "https://api.clarifai.com/v2/models/aaa03c23b3724a16a56b629203edc62c/outputs"


func main(){
	ch := make(chan []byte, 10)

	client := &http.Client{}
	url := "https://farm7.staticflickr.com/5769/21094803716_da3cea21b8_o.jpg"
	go getTags(client, url, ch)
	// fmt.Println(res)
}

func getTags(client *http.Client, url string, ch chan []byte) {
	// create url json
	fmt.Println("getting here")
	t := `{"inputs": [{"data": {"image": { "url": "` + url + `"}}}]}`
	picUrl := []byte(t)

	// make post requirement， get tag struct
	req, err := http.NewRequest("POST", TAR, bytes.NewReader(picUrl))
	if err != nil {
	}
	req.Header.Set("Authorization", " Key d4f76e005d404eb69893a5f721550d62")
	req.Header.Add("Content-Type", "application/json")
	resp, err := client.Do(req)
	defer resp.Body.Close()
	if err != nil {
	}
	result, err := ioutil.ReadAll(resp.Body) // read body
	if err != nil {
	}
	ch <- result
}
