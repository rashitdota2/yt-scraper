package yt_api

import (
	"context"
	"testing"
)

const testApiKey = "AIzaSyD--"

const chanUsername = "PGL_DOTA2"

func TestGetChanInfo(t *testing.T) {
	client := NewYtClient(testApiKey)

	info, err := client.GetChanInfo(context.Background(), chanUsername)
	if err != nil {
		t.Fatal(err)
	}
	t.Log("title:", info.Snippet.Title)
	t.Log("desc:", info.Snippet.Description)
	t.Log("len_of_desc", len(info.Snippet.Description))
	t.Log("thumbnails:", info.Snippet.Thumbnails.High.URL)
	t.Log("upload_playlist_id:", info.ContentDetails.RelatedPlaylist.Uploads)
}
