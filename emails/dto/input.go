package dto

type EmailRequest struct {
	Email   string
	Subject string
	Body    string
}

type AccountCredentials struct {
	Name        string
	Position    string
	Email       string
	Username    string
	Password    string
	LoginURL    string
	Year        string
	FacebookURL string
	TwitterURL  string
	LinkedinURL string
}
