package auth

import (
	"context"
	"fmt"
	"os"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"github.com/google/uuid"
	"github.com/joho/godotenv"

	"go-server/core/database"
	"go-server/core/model"
	"go-server/utils"

)

var secretJwt *string
var ctx = context.Background()

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
	// 1 hour JWT Token
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

func generateRefreshToken(data model.JwtData) (string, error) {

}

func generateDeviceUUID() string {
	return uuid.New().String()
}

func storeRefreshToken(userId string, deviceId string, ttl time.Duration) (string, error) {
	redisCl := database.Redis

	key := fmt.Sprintf("refresh:%s:%s", userId, deviceId)

	oldToken, _ := redisCl.Get(ctx, key).Result()

	if oldToken != "" {
		redisCl.Del(ctx, oldToken)
	}

	secureToken, err := utils.GenerateSecureToken(32)
	if err != nil {
		return "", err
	}
	// set with userId & deviceId as a key
	err = redisCl.Set(ctx, key, secureToken, ttl).Err()
	if err != nil {
		redisCl.Del(ctx, key)
		return "", err
	}

	// set with secureToken as a key
	reverseKey := fmt.Sprintf("token:%s", secureToken)
	err = redisCl.Set(ctx, reverseKey, key, ttl).Err()
	if err != nil {
		redisCl.Del(ctx, reverseKey)
		return "", err
	}
	return secureToken, nil
}
