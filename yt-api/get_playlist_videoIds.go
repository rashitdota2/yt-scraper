package yt_api

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"

	"github.com/rashitdota2/yt-scraper/models"
)

func (y *YtClient) GetPlaylistVideoIds(ctx context.Context, playlistId string) ([]string, error) {
	req, err := http.NewRequestWithContext(ctx, http.MethodGet,
		"https://www.googleapis.com/youtube/v3/playlistItems", nil)
	if err != nil {
		return nil, err
	}

	q := req.URL.Query()
	q.Add("part", "contentDetails")
	q.Add("playlistId", playlistId)
	q.Add("maxResults", "10")
	q.Add("fields", "items(contentDetails/videoId)")
	req.URL.RawQuery = q.Encode()

	req.Header.Add("Accept", "application/json")

	resp, err := y.client.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("unexpected status code: %d", resp.StatusCode)
	}

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	var respBody struct {
		Items []models.PlaylistInfo `json:"items"`
	}

	if err = json.Unmarshal(body, &respBody); err != nil {
		return nil, err
	}

	if len(respBody.Items) == 0 {
		return nil, fmt.Errorf("no videos found for playlistId: %s", playlistId)
	}

	var ids = make([]string, len(respBody.Items))

	for i, item := range respBody.Items {
		ids[i] = item.ContentDetails.VideoId
	}

	return ids, nil
}
