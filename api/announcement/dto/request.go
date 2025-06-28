package dto

type Announcement struct {
	Title       string `json:"title" validate:"required"`
	Description string `json:"description" validate:"required"`
}
