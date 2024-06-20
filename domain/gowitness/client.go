package gowitness

import (
	"bytes"
	"encoding/json"
	"io"
	"net/http"
)

type (
	Client interface {
		AddUrl(string) error
		Search(string)
	}
	clientImpl struct {
		baseURL string
		client  http.Client
	}
)

// AddUrl implements Client.
func (c *clientImpl) AddUrl(url string) error {
	postBody, _ := json.Marshal(map[string]string{
		"url":     url,
		"oneshot": "false",
	})
	response, err := http.Post(c.baseURL+"/screenshot", "application/json", bytes.NewBuffer(postBody))
	if err != nil {
		return err
	}
	if response.Body != nil {
		defer response.Body.Close()
	}
	_, readErr := io.ReadAll(response.Body)
	if readErr != nil {
		return readErr
	}
	return nil
}

// Search implements Client.
func (c *clientImpl) Search(string) {
	panic("gowitness.Search is unimplemented")
}

func NewClient(client http.Client, baseURL string) Client {
	return &clientImpl{client: client, baseURL: baseURL}
}
