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
	Captcha     string `json:"captcha" binding:"required"`
}

type RecaptchaResponse struct {
	Success     bool     `json:"success"`
	ChallengeTS string   `json:"challenge_ts"`
	Hostname    string   `json:"hostname"`
	ErrorCodes  []string `json:"error-codes"`
}
