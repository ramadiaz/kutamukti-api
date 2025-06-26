package dto

type Complaint struct {
	Title       string `json:"title" bind:"required"`
	Description string `json:"description" bind:"required"`
}
