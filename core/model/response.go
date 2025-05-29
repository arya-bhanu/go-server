package model

import "github.com/golang-jwt/jwt/v5"

type ResponseType = int
type Response struct {
	Status  int    `json:"status"`
	Message string `json:"message"`
	Data    any    `json:"data,omitempty"`
}

type JwtClaims struct {
	Data JwtData `json:"data"`
	jwt.RegisteredClaims
}

type JwtData struct {
	Name     string `json:"name"`
	Id       string `json:"id"`
	Email    string `json:"email"`
	Role     string `json:"role"`
	Username string `json:"username"`
}

type LoginAuth struct {
	JwtAccessToken  string `json:"jwt_access_token"`
	JwtRefreshToken string `json:"jwt_refresh_token"`
}

func GenerateResponse(status uint, message string, data any) (*Response, int) {
	return &Response{
		Status:  int(status),
		Message: message,
		Data:    data,
	}, int(status)
}
