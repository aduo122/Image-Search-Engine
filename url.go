package main

import (
		"fmt"
		"io/ioutil"
		"net/http"
		"net/url"
)

func main(){

	type image struct {
		Inputs []struct {
			Data struct {
				Image struct {
					URL string `json:"url"`
				} `json:"image"`
			} `json:"data"`
		} `json:"inputs"`
	}
	data := image{
		Inputs []struct{
			Data struct{
				Image struct{
					URL : "https://samples.clarifai.com/metro-north.jpg"
				},
			},
		},
	}
	payloadBytes, err := json.Marshal(data)
	if err != nil {
		// handle err
	}
	body := bytes.NewReader(payloadBytes)

	req, err := http.NewRequest("POST", "https://api.clarifai.com/v2/models/aaa03c23b3724a16a56b629203edc62c/outputs", body)
	if err != nil {
		// handle err
	}
	req.Header.Set("Authorization", "Key d4f76e005d404eb69893a5f721550d62")
	req.Header.Set("Content-Type", "application/json")

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		// handle err
	}
	defer resp.Body.Close()
}
