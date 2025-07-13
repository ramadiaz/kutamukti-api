package dto

type UMKM struct {
	Name        string  `json:"name" validate:"required"`
	Description string  `json:"description" validate:"required"`
	Is24Hours   bool    `json:"is24_hours"`
	Location    string  `json:"location" validate:"required"`
	Langitude   float64 `json:"langitude" validate:"required"`
	Latitude    float64 `json:"latitude" validate:"required"`
	OpenAt      int     `json:"open_at" validate:"required,gte=0,lte=24"`
	CloseAt     int     `json:"close_at" validate:"required,gte=0,lte=24"`
}