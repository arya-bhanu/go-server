package auth

import (
	"os"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"github.com/joho/godotenv"

	"go-server/core/model"
)

var secretJwt *string

func init() {
	err := godotenv.Load()
	secret := os.Getenv("JWT_SECRET")
	if err == nil {
		if secret == "" {
			secretJwt = nil
		} else {
			secretJwt = &secret
		}
	}

}

func LoginAuth(data model.JwtData) (*model.LoginAuth, error) {
	token, err := generateJWTToken(data)
	if err != nil {
		return nil, err
	}
	return &model.LoginAuth{
		JwtAccessToken:  token,
		JwtRefreshToken: "refresh_token_here",
	}, nil

}

func generateJWTToken(data model.JwtData) (string, error) {
	expires := jwt.NewNumericDate(time.Now().Add(1 * time.Hour))
	if secretJwt != nil {
		token := jwt.NewWithClaims(jwt.SigningMethodHS256, model.JwtClaims{
			RegisteredClaims: jwt.RegisteredClaims{
				ExpiresAt: expires,
				IssuedAt:  jwt.NewNumericDate(time.Now()),
				NotBefore: jwt.NewNumericDate(time.Now()),
			},
			Data: data,
		})

		tokenString, err := token.SignedString([]byte(*secretJwt))
		if err != nil {
			return "", err
		}
		return tokenString, nil
	}
	return "", nil
}
