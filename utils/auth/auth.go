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

var secretJwtKey *string
var ctx = context.Background()

func init() {
	err := godotenv.Load()
	secret := os.Getenv("JWT_SECRET")
	if err == nil {
		if secret == "" {
			secretJwtKey = nil
		} else {
			secretJwtKey = &secret
		}
	}

}

func LoginAuth(data model.JwtData) (*model.LoginAuth, error) {
	token, err := generateJWTToken(data)
	tokenRefresh, errRefresh := generateRefreshToken(data)
	if err != nil {
		return nil, err
	}
	if errRefresh != nil {
		return nil, errRefresh
	}
	return &model.LoginAuth{
		JwtAccessToken:  token,
		JwtRefreshToken: tokenRefresh,
	}, nil

}

func RefreshToken() {}

func ValidateToken(tokenStr string) (*jwt.Token, error) {
	tokenParsed, err := jwt.Parse(tokenStr, func(t *jwt.Token) (any, error) {
		if _, ok := t.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", t.Header["alg"])
		}
		return []byte(*secretJwtKey), nil
	})
	if err != nil || !tokenParsed.Valid {
		return nil, fmt.Errorf("error: invalid token %+v", err)
	}

	claims := tokenParsed.Claims.(jwt.MapClaims)

	if exp, ok := claims["exp"].(float64); !ok || int64(exp) < time.Now().Unix() {
		fmt.Print(exp)
		fmt.Print(ok)
		return nil, fmt.Errorf("error: token expire")
	}

	return tokenParsed, nil

}

func generateJWTToken(data model.JwtData) (string, error) {
	// 1 hour JWT Token
	expires := jwt.NewNumericDate(time.Now().Add(1 * time.Hour))
	if secretJwtKey != nil {
		token := jwt.NewWithClaims(jwt.SigningMethodHS256, model.JwtClaims{
			RegisteredClaims: jwt.RegisteredClaims{
				ExpiresAt: expires,
				IssuedAt:  jwt.NewNumericDate(time.Now()),
				NotBefore: jwt.NewNumericDate(time.Now()),
			},
			Data: data,
		})

		tokenString, err := token.SignedString([]byte(*secretJwtKey))
		if err != nil {
			return "", err
		}
		return tokenString, nil
	}
	return "", nil
}

func generateRefreshToken(data model.JwtData) (string, error) {
	// 1 day token
	expireTime := 24 * time.Hour
	deviceId := generateDeviceUUID()
	refreshToken, err := storeRefreshToken(data.Id, deviceId, expireTime)
	if err != nil {
		return "", err
	}
	return refreshToken, nil
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
