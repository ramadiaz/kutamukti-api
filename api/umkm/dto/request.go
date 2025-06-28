package dto

import "time"

type UMKM struct {
	Name        string    `json:"name" validate:"required"`
	Description string    `json:"description" validate:"required"`
	Is24Hours   bool      `json:"is24_hours"`
	Location    string    `json:"location" validate:"required"`
	Langitude   float64   `json:"langitude" validate:"required"`
	Latitude    float64   `json:"latitude" validate:"required"`
	OpenAt      time.Time `json:"open_at" validate:"required"`
	CloseAt     time.Time `json:"close_at" validate:"required"`
}
