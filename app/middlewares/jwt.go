package middlewares

import (
	"fmt"
	"net/http"
	"os"
	"strings"

	"github.com/dgrijalva/jwt-go"
)

type MyCustomClaims struct {
	Token string `json:"token"`
	jwt.StandardClaims
}

// ValidateJwtMiddlewear ... This function helps to log the stuff
func ValidateJwtMiddlewear(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		authToken := r.Header.Get("Authorization")
		splitToekn := strings.Split(authToken, "Bearer")
		if len(splitToekn) != 2 {
			//todo log token error
		}
		tokenSring := strings.TrimSpace(splitToekn[1])

		token, err := jwt.ParseWithClaims(tokenSring, &MyCustomClaims{}, func(token *jwt.Token) (interface{}, error) {
			return []byte([]byte(os.Getenv("JWT_SECRATE"))), nil
		})

		if claims, ok := token.Claims.(*MyCustomClaims); ok && token.Valid {
			//todo: add logic to check token xpired or not
			fmt.Printf("%v %v", claims.Token, claims.StandardClaims.ExpiresAt)
		} else {
			//todo : add some changes for catching error
			fmt.Println(err)
		}

		next.ServeHTTP(w, r)
	})
}