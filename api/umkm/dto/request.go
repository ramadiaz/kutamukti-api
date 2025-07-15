package dto

type UMKM struct {
	Name        string  `json:"name" validate:"required"`
	Description string  `json:"description" validate:"required"`
	Is24Hours   bool    `json:"is24_hours"`
	Location    string  `json:"location" validate:"required"`
	Langitude   float64 `json:"langitude" validate:"required"`
	Latitude    float64 `json:"latitude" validate:"required"`
	OpenAt      int     `json:"open_at" validate:"gte=0,lte=24"`
	CloseAt     int     `json:"close_at" validate:"gte=0,lte=24"`
}

type UMKMProduct struct {
	UMKMUUID    string      `json:"umkm_uuid" validate:"required"`
	Name        string      `json:"name" validate:"required"`
	Description string      `json:"description" validate:"required"`
	Variation   string      `json:"variation" validate:"required"`
	Price       int         `json:"price" validate:"required,gte=0"`
	Images      []UMKMImage `json:"images"`
}

type UMKMImage struct {
	ImageURL string `json:"image_url"`
}
