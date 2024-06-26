package gowitness

import (
	"bytes"
	"encoding/json"
	"io"
	"net/http"
)

type (
	Client interface {
		AddURL(string) (int, error)
		Details(int) (DetailsURL, error)
	}
	DetailsURL     struct{}
	addURLResponse struct{ ID int }

	clientImpl struct {
		baseURL string
		client  http.Client
	}
)

// Details implements Client.
func (c *clientImpl) Details(int) (DetailsURL, error) {
	panic("Details unimplemented")
}

// AddURL implements Client.
func (c *clientImpl) AddURL(url string) (int, error) {
	postBody, _ := json.Marshal(map[string]string{
		"url":        url,
		"oneshot":    "false",
		"foreground": "true",
	})
	response, err := http.Post(c.baseURL+"/screenshot", "application/json", bytes.NewBuffer(postBody))
	if err != nil {
		return 0, err
	}
	if response.Body != nil {
		defer response.Body.Close()
	}
	result, readErr := io.ReadAll(response.Body)
	if readErr != nil {
		return 0, readErr
	}
	responseType := addURLResponse{}
	jsonErr := json.Unmarshal(result, &responseType)
	if jsonErr != nil {
		return 0, nil
	}
	return responseType.ID, nil
}

func NewClient(client http.Client, baseURL string) Client {
	return &clientImpl{client: client, baseURL: baseURL}
}
