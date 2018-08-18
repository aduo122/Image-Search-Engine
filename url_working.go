package main

import (
  "bytes"
  "io/ioutil"
  "net/http"
  "fmt"
  "strings"
  "encoding/json"
  "time"
)

type pic_tag struct {
	Status struct {
		Code        int    `json:"code"`
		Description string `json:"description"`
	} `json:"status"`
	Outputs []struct {
		ID     string `json:"id"`
		Status struct {
			Code        int    `json:"code"`
			Description string `json:"description"`
		} `json:"status"`
		CreatedAt time.Time `json:"created_at"`
		Model     struct {
			ID         string    `json:"id"`
			Name       string    `json:"name"`
			CreatedAt  time.Time `json:"created_at"`
			AppID      string    `json:"app_id"`
			OutputInfo struct {
				Message string `json:"message"`
				Type    string `json:"type"`
				TypeExt string `json:"type_ext"`
			} `json:"output_info"`
			ModelVersion struct {
				ID        string    `json:"id"`
				CreatedAt time.Time `json:"created_at"`
				Status    struct {
					Code        int    `json:"code"`
					Description string `json:"description"`
				} `json:"status"`
			} `json:"model_version"`
			DisplayName string `json:"display_name"`
		} `json:"model"`
		Input struct {
			ID   string `json:"id"`
			Data struct {
				Image struct {
					URL string `json:"url"`
				} `json:"image"`
			} `json:"data"`
		} `json:"input"`
		Data struct {
			Concepts []struct {
				ID    string  `json:"id"`
				Name  string  `json:"name"`
				Value float64 `json:"value"`
				AppID string  `json:"app_id"`
			} `json:"concepts"`
		} `json:"data"`
	} `json:"outputs"`
}

func main(){
  client := &http.Client{}
  TAR := "https://api.clarifai.com/v2/models/aaa03c23b3724a16a56b629203edc62c/outputs"

  urls := getURLs()
  for index, url := range urls{
    if index > 3{
      break
    }
    fmt.Println(index, string(url))
    t := `{"inputs": [{"data": {"image": { "url": "` + string(url) + `"}}}]}`
    fmt.Println(t)
    body := []byte(t)
    req, err := http.NewRequest("POST", TAR, bytes.NewReader(body))
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
    res := new(pic_tag)
    json.Unmarshal(result, &res)
    // fmt.Println(res.Outputs[0].Data.Concepts)
    for _, scores := range res.Outputs[0].Data.Concepts{
      fmt.Println(scores.Name,scores.Value)
    }
    // fmt.Println(string(result)) // remember to turn into string before print
  }
}

func getURLs() []string {
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
  }
  urls := strings.Split(string(res), "\n")
  return urls
}
