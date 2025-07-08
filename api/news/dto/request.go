package dto

type News struct {
	Title        string       `json:"title" validate:"required"`
	Content      string       `json:"content" validate:"required"`
	ThumbnailURL string       `json:"thumbnail_url" validate:"required"`
	Images       []NewsImages `json:"images" validate:"required"`
}

type NewsImages struct {
	ImageURL string `json:"image_url" validate:"required"`
}
