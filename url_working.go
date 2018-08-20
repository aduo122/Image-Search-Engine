package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"
	"time"

	"github.com/go-redis/redis"
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

// type pic_index struct {
// 	Value float64 `json:"value"`
// 	Url   string  `json:"url"`
// }
type chData struct {
	Url string `json:"url"`
	Tag []byte `json:"tag"`
}

const TAR string = "https://api.clarifai.com/v2/models/aaa03c23b3724a16a56b629203edc62c/outputs"

func main() {
	channel := make(chan *chData, 8)
	client := &http.Client{}

	//initial redis
	redisClient := redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		Password: "", // no password set
		DB:       0,  // use default DB
	})
	pong, err := redisClient.Ping().Result()
	if err != nil {
		fmt.Println("Reids not connected")
	} else {
		fmt.Println(pong, err)
	}
	// get url and send to redis
	urls := getURLs(client)
	for index, url := range urls {
		if index > 5 {
			break
		}
		go getTags(client, url, channel)
		go fetch(redisClient, channel)
	}
	time.Sleep(1 * 1e9)
}

func getURLs(client *http.Client) []string {
	// client := &http.Client{}
	target := "https://s3.amazonaws.com/clarifai-data/backend/api-take-home/images.txt"
	request, err := http.NewRequest("GET", target, nil)
	if err != nil {
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

func getTags(client *http.Client, url string, ch chan *chData) {
	// create url json
	t := `{"inputs": [{"data": {"image": { "url": "` + url + `"}}}]}`
	picUrl := []byte(t)

	// make post request, get tag struct
	req, err := http.NewRequest("POST", TAR, bytes.NewReader(picUrl))
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

	// add chData to channel
	res := new(chData)
	res.Tag = result
	res.Url = url
	fmt.Println("get tag of " + url)
	ch <- res
}

func fetch(redisClient *redis.Client, ch chan *chData) {
	temp := <-ch
	result := temp.Tag
	url := temp.Url
	res := new(pic_tag)
	json.Unmarshal(result, &res) // res: tag struct

	// save {label: [url]} to redis
	for _, scores := range res.Outputs[0].Data.Concepts {
		//initial info for the key
		err := redisClient.ZAdd(scores.Name, redis.Z{
			Score:  scores.Value,
			Member: url,
		}).Err()
		if err != nil {
			panic(err)
		}
		fmt.Println(scores.Name, scores.Value)
	}
}
