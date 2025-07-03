package dto

type ComplaintStatus string

const (
	Open    ComplaintStatus = "open"
	Process ComplaintStatus = "process"
	Closed  ComplaintStatus = "closed"
)

func (s ComplaintStatus) String() string {
	return string(s)
}

type Complaint struct {
	Title       string `json:"title" validate:"required"`
	Description string `json:"description" validate:"required"`
}
