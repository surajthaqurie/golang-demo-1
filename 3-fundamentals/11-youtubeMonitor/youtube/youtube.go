package youtube

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"
	"time"

	"github.com/joho/godotenv"
)

type Response struct {
	Kind  string `json:"kind"`
	Items []Item `json:"items"`
}

type Item struct {
	Kind  string `json:"kind"`
	Id    string `json:"id"`
	Stats Stats  `json:"statistics"`
}

type Stats struct {
	Views       string `json:"viewCount"`
	Subscribers string `json:"subscriberCount"`
	Videos      string `json:"videoCount"`
}

func GetSubscriber() (Item, error) {
	err := godotenv.Load()
	if err != nil {
		return Item{}, fmt.Errorf("failed to load .env file: %w", err)
	}

	req, err := http.NewRequest("GET", "https://www.googleapis.com/youtube/v3/channels", nil)
	if err != nil {
		fmt.Println(err)
		return Item{}, err
	}

	q := req.URL.Query()
	q.Add("key", os.Getenv("YOUTUBE_API_KEY"))
	q.Add("id", os.Getenv("YOUTUBE_CHANNEL_ID"))
	q.Add("part", "statistics")
	req.URL.RawQuery = q.Encode()

	client := &http.Client{
		Timeout: 10 * time.Second, // Set a timeout for the HTTP request
	}

	resp, err := client.Do(req)
	if err != nil {
		fmt.Println(err)
		return Item{}, err
	}
	defer resp.Body.Close()

	fmt.Println("Response status: ", resp.Status)

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		fmt.Println(err)
		return Item{}, err
	}

	var response Response
	err = json.Unmarshal(body, &response)
	if err != nil {
		fmt.Println(err)
		return Item{}, err
	}

	return response.Items[0], nil
}
