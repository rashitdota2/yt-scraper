package yt_api

import (
	"context"
	"testing"
)

const playlistId = "UUK5d3n3kfkzlArMccS0TTXA"

func TestGetPlaylistVideoIds(t *testing.T) {
	client := NewYtClient(testApiKey)

	ids, err := client.GetPlaylistVideoIds(context.Background(), playlistId)
	if err != nil {
		t.Fatal(err)
	}
	t.Log(ids)
}
