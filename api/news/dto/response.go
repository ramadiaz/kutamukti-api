package dto

import "time"

type Response struct {
	Status  int         `json:"status"`
	Message string      `json:"message"`
	Body    interface{} `json:"body,omitempty"`
}

type NewsResponse struct {
	UUID         string       `json:"uuid"`
	Title        string       `json:"title"`
	Content      string       `json:"content"`
	RawText      string       `json:"raw_text"`
	Slug         string       `json:"slug"`
	ThumbnailURL string       `json:"thumbnail_url"`
	Images       []NewsImages `json:"images"`

	CreatedAt time.Time `json:"created_at"`

	User NewsAuthor `json:"author"`
}

type NewsAuthor struct {
	UUID string `json:"uuid"`
	Name string `json:"name"`
}
