package middlewares

import (
	env "AuthApp/config/env"
	"context"
	"fmt"
	"net/http"
	"strconv"
	"strings"

	"github.com/golang-jwt/jwt/v4"
)

func JWTAuthMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		authHeader := r.Header.Get("Authorization")

		if authHeader == "" {
			http.Error(w, "Authorization header is required", http.StatusUnauthorized)
			return
		}

		if !strings.HasPrefix(authHeader, "Bearer ") {
			http.Error(w, "Authorization header must start with Bearer", http.StatusUnauthorized)
			return
		}

		token := strings.TrimPrefix(authHeader, "Bearer ")

		if token == "" {
			http.Error(w, "Token is required", http.StatusUnauthorized)
			return
		}
		claims := jwt.MapClaims{}

		_, err := jwt.ParseWithClaims(token, &claims, func(token *jwt.Token) (any, error) {
			return []byte(env.GetString("JWT_TOKEN", "TOKEN")), nil
		})

		if err != nil {
			http.Error(w, "Invalid token: "+err.Error(), http.StatusUnauthorized)
			return
		}

		userId, okId := claims["id"].(float64)

		email, okEmail := claims["Email"].(string)

		if !okId || !okEmail {
			http.Error(w, "Invalid token claims", http.StatusUnauthorized)
			return
		}

		fmt.Println("Authenticated user id: ", userId, "Email: ", email)

		ctx := context.WithValue(r.Context(), "userId", strconv.FormatFloat(userId, 'f', 0, 64))

		ctx = context.WithValue(ctx, "email", email)

		next.ServeHTTP(w, r.WithContext(ctx))
	})
}
