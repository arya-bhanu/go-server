package model

type UserRegisterRequest struct {
	Name     string `json:"name"`
	Email    string `json:"email"`
	Password string `json:"password"`
	Username string `json:"username"`
}

type UserLoginRequest struct {
	Identifier string `json:"identifier,omitempty"`
	Password   string `json:"password"`
}
