package yt_api

import (
	"net/http"
	"time"
)

type YtClient struct {
	client *http.Client
	apiKey string
}

func NewYtClient(apiKey string) *YtClient {
	return &YtClient{
		client: &http.Client{
			Timeout: time.Minute,
		},
		apiKey: apiKey,
	}
}
