package main

import (
  "bytes"
  "io/ioutil"
  "net/http"
  "fmt"
)

func main(){
  client := &http.Client{}
  tar := "https://api.clarifai.com/v2/models/aaa03c23b3724a16a56b629203edc62c/outputs"
  body := []byte(`{"inputs": [{"data": {"image": { "url": "https://samples.clarifai.com/metro-north.jpg"}}}]}`)

  req, err := http.NewRequest("POST", tar, bytes.NewReader(body))
  if err != nil {
    return
  }

  req.Header.Set("Authorization", " Key d4f76e005d404eb69893a5f721550d62")
  req.Header.Add("Content-Type", "application/json")

  resp, err := client.Do(req)
  defer resp.Body.Close()
  if err != nil {
    return
  }

  result, err := ioutil.ReadAll(resp.Body) // read body
  if err != nil {
    return
  }
  fmt.Println(string(result)) // remember to turn into string before print
}
