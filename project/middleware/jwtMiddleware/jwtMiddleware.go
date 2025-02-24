package jwtmiddleware

import (
	"net/http"
	"pstgSQL/project/models/user"
	"strings"

	"github.com/golang-jwt/jwt/v5"
)

var jwtKey []byte
func init() {
    jwtKey = []byte(user.GetJWTKey())
}

func JWTMiddleware(next http.Handler) http.Handler{
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		authHeader := r.Header.Get("Authorization")
		if authHeader == "" || !strings.HasPrefix(authHeader, "Bearer ") {
			http.Error(w, "Missing or invalid token", http.StatusUnauthorized)
			return
		}

		tokenStr := strings.TrimPrefix(authHeader, "Bearer ")

		token, err := jwt.Parse(tokenStr, func(t *jwt.Token) (interface{}, error) {
			return jwtKey, nil
		})
		if err != nil || !token.Valid {
			http.Error(w, "Invalid token", http.StatusUnauthorized)
		} 

		next.ServeHTTP(w, r)
	})
}