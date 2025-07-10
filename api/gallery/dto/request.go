package dto

type ImageGalleries struct {
	Title  string   `json:"title" validate:"required"`
	Images []Images `json:"images" validate:"required"`
}

type Images struct {
	ImageURL    string `json:"image_url" validate:"required"`
	Description string `json:"description" validate:"required"`
}


