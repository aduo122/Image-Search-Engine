package main

import (
  // "bytes"
  "io/ioutil"
  "net/http"
  "fmt"
  "strings"
)

func main(){
  client := &http.Client{}
  target := "https://s3.amazonaws.com/clarifai-data/backend/api-take-home/images.txt"
  request, err := http.NewRequest("GET", target, nil)
  if err != nil{
    fmt.Println("Fatal error", err.Error())
  }
  response, err := client.Do(request)
  defer response.Body.Close()
  res, err1 := ioutil.ReadAll(response.Body)
  if err1 != nil {
      return
  }
  urls := strings.Split(string(res), "\n")
  for _, url := range urls{
    fmt.Println(index, string(url))
    // fmt.Println(index)
  }
}
