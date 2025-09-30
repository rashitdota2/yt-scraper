package yt_api

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"

	"github.com/rashitdota2/yt-scraper/models"
)

func (y *YtClient) GetChanInfo(ctx context.Context, username string) (models.ChanInfo, error) {
	req, err := http.NewRequestWithContext(ctx, http.MethodGet,
		"https://www.googleapis.com/youtube/v3/channels", nil)
	if err != nil {
		return models.ChanInfo{}, err
	}

	q := req.URL.Query()
	q.Add("part", "snippet,contentDetails")
	q.Add("fields", "items(snippet(title,description,thumbnails/high/url),contentDetails/relatedPlaylists/uploads)")
	q.Add("forHandle", username)
	q.Set("key", y.apiKey)
	req.URL.RawQuery = q.Encode()

	resp, err := y.client.Do(req)
	if err != nil {
		return models.ChanInfo{}, err
	}
	defer resp.Body.Close()

	var respBody struct {
		Items []models.ChanInfo `json:"items"`
	}

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return models.ChanInfo{}, err
	}

	if resp.StatusCode != http.StatusOK {
		return models.ChanInfo{}, fmt.Errorf("invalid status code: %v,\n body: %s", resp.Status,
			string(body))
	}

	if err := json.Unmarshal(body, &respBody); err != nil {
		return models.ChanInfo{}, err
	}

	if len(respBody.Items) < 1 {
		return models.ChanInfo{}, errors.New("chan not found")
	}

	return respBody.Items[0], nil
}
