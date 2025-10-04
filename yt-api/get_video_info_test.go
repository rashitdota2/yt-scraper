package yt_api

import (
	"context"
	"testing"
)

const videoId = "OGtdfYFIz00"

func TestGetVideoInfo(t *testing.T) {
	client := NewYtClient(testApiKey)

	info, err := client.GetVideoInfo(context.Background(), videoId)
	if err != nil {
		t.Fatal(err)
	}
	t.Log(info)
}
