package main

import (
		"fmt"
		// "io/ioutil"
		// "net/http"
		// "net/url"
		// "bytes"
		// "encoding/json"
		"github.com/mikemintang/go-curl"
)

// type ImageStruct struct {
// 	URL string
// }
// type DataStruct struct {
// 	Image ImageStruct
// }
// type InputStruct struct {
// 	Data DataStruct
// }
// type weburl struct {
// 	Inputs []InputStruct
// }
type Payload struct {
	Inputs []struct {
		Data struct {
			Image struct {
				URL string `json:"url"`
			} `json:"image"`
		} `json:"data"`
	} `json:"inputs"`
}

func main(){

	// var input InputStruct
	// input.Data.Image.URL = "https://samples.clarifai.com/metro-north.jpg"
	// // fmt.Println(input)
	// var data weburl
	// data.Inputs = append(data.Inputs, input)
	// fmt.Println(data)
	url := "https://api.clarifai.com/v2/models/aaa03c23b3724a16a56b629203edc62c/outputs"
	headers := map[string]string{
			"Authorization": "Key d4f76e005d404eb69893a5f721550d62",
			"Content-Type":  "application/json",
	}

	postData := map[string]interface{}{
	    "inputs": []map[string]interface{}{
	        "data": map[string]interface{}{
	          "image": map[string]string{
	            "url": "https://samples.clarifai.com/metro-north.jpg",
	          },
	        },
	      },
	  }

	req := curl.NewRequest()
	resp, err := req.
			SetUrl(url).
			SetHeaders(headers).
			SetPostData(postData).
			Post()

	if err != nil {
			fmt.Println(err)
	} else {
			if resp.IsOk() {
					fmt.Println(resp.Body)
			} else {
					fmt.Println(resp.Raw)
			}
	}

	// datajson := `{"inputs": [{"data": {"image": {"url": "https://samples.clarifai.com/metro-north.jpg"}}}]}`
	// var data Payload
	// json.Unmarshal([]byte(datajson), &data)
	// fmt.Println(data)
	// payloadBytes, err := json.Marshal(data)
	// // fmt.Println(payloadBytes)
	// if err != nil {
	// 	// handle err
	// }
	// body := bytes.NewReader(payloadBytes)
	//
	// req, err := http.NewRequest("POST", "https://api.clarifai.com/v2/models/aaa03c23b3724a16a56b629203edc62c/outputs", body)
	// if err != nil {
	// 	// handle err
	// }
	// req.Header.Set("Authorization", "Key d4f76e005d404eb69893a5f721550d62")
	// req.Header.Set("Content-Type", "application/json")
	// // fmt.Println(req)
	// resp, err := http.DefaultClient.Do(req)
	// if err != nil {
	// 	// handle err
	// }
	//
	// defer resp.Body.Close()
	fmt.Println(resp)
}
