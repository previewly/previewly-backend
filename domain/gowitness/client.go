package gowitness

import (
	"bytes"
	"encoding/json"
	"io"
	"net/http"
	"wsw/backend/lib/utils"
)

type (
	Client interface {
		AddUrl(string) error
		Search(string) (string, error)
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
	result, readErr := io.ReadAll(response.Body)
	if readErr != nil {
		return readErr
	}
	utils.D(string(result))
	return nil
}

// Search implements Client.
func (c *clientImpl) Search(url string) (string, error) {
	response, err := http.Get(c.baseURL + "/search?q=" + url)
	if err != nil {
		return "", err
	}
	if response.Body != nil {
		defer response.Body.Close()
	}
	_, readErr := io.ReadAll(response.Body)
	if readErr != nil {
		return "", readErr
	}
	// utils.D(string(result))
	return "", nil
}

func NewClient(client http.Client, baseURL string) Client {
	return &clientImpl{client: client, baseURL: baseURL}
}
