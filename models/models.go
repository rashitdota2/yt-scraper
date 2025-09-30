package models

type ChanInfo struct {
	Snippet struct {
		Title       string     `json:"title"`
		Description string     `json:"description"`
		Thumbnails  Thumbnails `json:"thumbnails"`
	} `json:"snippet"`
	ContentDetails struct {
		RelatedPlaylist struct {
			Uploads string `json:"uploads"`
		} `json:"relatedPlaylists"`
	} `json:"contentDetails"`
}

type Thumbnails struct {
	High struct {
		URL string `json:"url"`
	} `json:"high"` // for chan info
	Maxres struct {
		URL string `json:"url"`
	} `json:"maxres"` // for video info
}

type VideoInfo struct {
	Snippet struct {
		Title       string     `json:"title"`
		Description string     `json:"description"`
		Thumbnails  Thumbnails `json:"thumbnails"`
		Tags        []string   `json:"tags"`
		CategoryId  string     `json:"categoryId"`
	} `json:"snippet"`
	ContentDetails struct {
		Duration string `json:"duration"`
	} `json:"contentDetails"`
}

type PlaylistInfo struct {
	ContentDetails struct {
		VideoId string `json:"videoId"`
	} `json:"contentDetails"`
}
