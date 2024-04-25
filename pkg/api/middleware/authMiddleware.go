package middleware

import (
	"fmt"
	"net/http"
	"strings"

	"github.com/dgrijalva/jwt-go"
)

var jwtKey = []byte("your_secret_key")

func AuthMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// Get the JWT string from the header
		tokenString := r.Header.Get("Authorization")

		// Check if the token is empty
		if tokenString == "" {
			http.Error(w, "Authorization header must be present", http.StatusUnauthorized)
			return
		}

		// The token string should be in the format `Bearer {token}`
		splitToken := strings.Split(tokenString, "Bearer ")
		if len(splitToken) != 2 {
			http.Error(w, "Invalid token format", http.StatusUnauthorized)
			return
		}
		tokenString = splitToken[1]

		// Parse the token
		token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
			// Make sure the token method is what we expect
			if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
				return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
			}
			return jwtKey, nil
		})

		if err != nil {
			http.Error(w, "Invalid token", http.StatusUnauthorized)
			return
		}

		if !token.Valid {
			http.Error(w, "Invalid token", http.StatusUnauthorized)
			return
		}

		// If everything is OK, call the next handler
		next.ServeHTTP(w, r)
	})
}
