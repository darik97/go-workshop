package jokes

import (
	"encoding/json"
	"fmt"
	"net/http"
	"workshop/internal/api"
)

const getJokePath = "/api?format=json"

// JokeClient url for requesting joke
type JokeClient struct {
	url string
}

// NewJokeClient constructor for JokeClient
func NewJokeClient(baseURL string) *JokeClient {
	return &JokeClient{
		url: baseURL,
	}
}

// GetJoke send request and return response with joke
func (jc *JokeClient) GetJoke() (*api.JokeResponse, error) {
	urlPath := jc.url + getJokePath

	resp, err := http.Get(urlPath)
	if err != nil {
		return nil, err
	}

	var data api.JokeResponse

	err = json.NewDecoder(resp.Body).Decode(&data)
	if err != nil {
		return nil, err
	}
	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("API request error: %s", http.StatusText(resp.StatusCode))
	}

	return &data, nil
}
