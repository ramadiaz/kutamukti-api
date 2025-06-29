package dto

type ItemTypes string

const (
	Image ItemTypes = "image"
	Video ItemTypes = "video"
)

type Gallery struct {
	Type        ItemTypes `json:"type" validate:"required"`
	Title       string    `json:"title" validate:"required"`
	Description string    `json:"description" validate:"required"`
	Slug        string    `json:"slug" validate:"required"`
}
