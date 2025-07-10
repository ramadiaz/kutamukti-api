package dto

import "time"

type Response struct {
	Status  int         `json:"status"`
	Message string      `json:"message"`
	Body    interface{} `json:"body,omitempty"`
}

type ImageGalleriesResponse struct {
	UUID      string    `json:"uuid"`
	Title     string    `json:"title"`
	CreatedAt time.Time `json:"created_at"`

	Images []ImagesResponse `json:"images"`
}

type ImagesResponse struct {
	GalleryUUID string `json:"gallery_uuid"`
	ImageURL    string `json:"image_url"`
	Description string `json:"description"`
}
