package middleware

import (
	"context"
	"net/http"
	"strings"

	"github.com/golang-jwt/jwt/v5"

	"go-server/utils/auth"
)

func AuthGuardJWTMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		authHeader := r.Header.Get("Authorization")
		if !strings.Contains(authHeader, "Bearer") {
			http.Error(w, "Authorization Bearer Not Provided", http.StatusBadRequest)
			return
		}
		authToken := strings.TrimPrefix(authHeader, "Bearer ")
		token, err := auth.ValidateToken(authToken)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}
		claims := token.Claims.(jwt.MapClaims)
		data := claims["data"].(map[string]any)
		ctx := context.WithValue(r.Context(), "data", data)
		next.ServeHTTP(w, r.WithContext(ctx))
	})
}
