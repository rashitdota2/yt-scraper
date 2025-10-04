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

func (y *YtClient) GetVideoInfo(ctx context.Context, id string) (models.VideoInfo, error) {
	req, err := http.NewRequestWithContext(ctx, http.MethodGet,
		"https://www.googleapis.com/youtube/v3/videos", nil)
	if err != nil {
		return models.VideoInfo{}, err
	}

	q := req.URL.Query()
	q.Add("part", "contentDetails,snippet")
	q.Add("fields", "items(snippet(title,description,thumbnails/maxres/url,tags,categoryId),contentDetails/duration)")
	q.Add("id", id)
	q.Add("key", y.apiKey)
	req.URL.RawQuery = q.Encode()

	req.Header.Add("Accept", "application/json")

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return models.VideoInfo{}, err
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return models.VideoInfo{}, err
	}

	if resp.StatusCode != http.StatusOK {
		return models.VideoInfo{}, fmt.Errorf("invalid status code: %v,\n body: %s", resp.Status,
			string(body))
	}

	var respBody struct {
		Items []models.VideoInfo `json:"items"`
	}

	if err = json.Unmarshal(body, &respBody); err != nil {
		return models.VideoInfo{}, err
	}

	if len(respBody.Items) == 0 {
		return models.VideoInfo{}, errors.New("no videos found")
	}

	return respBody.Items[0], nil
}
