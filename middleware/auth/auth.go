package middleware

import (
	"fmt"
	"net/http"
	"os"

	"github.com/golang-jwt/jwt"
	"github.com/pemba1s1/go-server/internal/auth"
	"github.com/pemba1s1/go-server/utils"
)

func AuthMiddleware() func(http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			// Get the JWT token from the Authorization header
			tokenString, err := auth.GetAPIKey(r.Header)
			if err != nil {
				utils.RespondWithError(w, http.StatusUnauthorized, fmt.Sprintf("Error: %v", err))
			}
			// Parse the JWT token
			token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
				// Provide the key or secret used to sign the token
				return []byte(os.Getenv("SECRET_KEY")), nil
			})

			if err != nil || !token.Valid {
				// Return an unauthorized error if the token is invalid
				w.WriteHeader(http.StatusUnauthorized)
				w.Write([]byte("Unauthorized"))
				return
			}
			claims := token.Claims.(jwt.MapClaims)
			issuer := claims["iss"].(string)
			fmt.Println("Issuer:", issuer)
			// Call the next handler if the token is valid
			next.ServeHTTP(w, r)
		})

	}
}
