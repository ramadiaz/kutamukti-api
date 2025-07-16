package dto

type Response struct {
	Status  int         `json:"status"`
	Message string      `json:"message"`
	Body    interface{} `json:"body,omitempty"`
}

type UMKMResponse struct {
	UUID        string  `json:"uuid" validate:"required"`
	Name        string  `json:"name" validate:"required"`
	Description string  `json:"description" validate:"required"`
	Is24Hours   bool    `json:"is24_hours" validate:"required"`
	Location    string  `json:"location" validate:"required"`
	Langitude   float64 `json:"langitude" validate:"required"`
	Latitude    float64 `json:"latitude" validate:"required"`
	OpenAt      int     `json:"open_at" validate:"required"`
	CloseAt     int     `json:"close_at" validate:"required"`
	Owner       string  `json:"owner"`
	Contact     string  `json:"contact"`

	Products []UMKMProductResponse `json:"products,omitempty"`
}

type UMKMProductResponse struct {
	UUID        string              `json:"uuid"`
	UMKMUUID    string              `json:"umkm_uuid"`
	Name        string              `json:"name"`
	Description string              `json:"description"`
	Variation   string              `json:"variation"`
	Price       int                 `json:"price"`
	Images      []UMKMImageResponse `json:"images"`

	UMKM UMKMResponse `json:"umkm,omitempty"`
}

type UMKMImageResponse struct {
	ImageURL string `json:"image_url"`
}
