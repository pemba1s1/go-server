package middleware

import (
	"fmt"
	"net/http"

	"github.com/golang-jwt/jwt"
	"github.com/pemba1s1/go-server/internal/auth"
	"github.com/pemba1s1/go-server/utils"
)

func AuthMiddleware(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		// Get the JWT token from the Authorization header
		tokenString, err := auth.GetAPIKey(r.Header)
		if err != nil {
			utils.RespondWithError(w, http.StatusUnauthorized, fmt.Sprintf("Error: %v", err))
		}
		// Parse the JWT token
		token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
			// Provide the key or secret used to sign the token
			return []byte("secret"), nil
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
	}
}
