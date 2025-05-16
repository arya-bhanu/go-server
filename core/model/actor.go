package model

type User struct {
	Id       string  `json:"id"`
	Name     string  `json:"name"`
	Role     *string `json:"role,omitempty"`
	Email    string  `json:"email"`
	Password string  `json:"password"`
}
