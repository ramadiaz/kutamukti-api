package dto

type Roles string

const (
	Admin Roles = "admin"
	Staff Roles = "staff"
)

func (r Roles) String() string {
	return string(r)
}

type User struct {
	Email    string `json:"email" validate:"required,email"`
	Username string `json:"username" validate:"required"`
	Name     string `json:"name" validate:"required"`
	Role     Roles  `json:"role" validate:"required,oneof=admin staff"`
}
