package model

type UserRegisterRequest struct {
	Name     string `json:"name"`
	Email    string `json:"email"`
	Password string `json:"password"`
}

type UserLoginRequest struct {
	Identifier string `json:"identifier,omitempty"`
	Password   string `json:"password"`
}
