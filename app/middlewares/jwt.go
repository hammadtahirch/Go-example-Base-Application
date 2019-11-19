package middlewares

import (
	"fmt"
	"net/http"
	"os"
	"strings"

	"git-lab.boldapps.net/nifty-logix/mvc/app/services"
	"git-lab.boldapps.net/nifty-logix/mvc/app/utils"
	"github.com/dgrijalva/jwt-go"
)

// CustomClaims ...
type CustomClaims struct {
	Token string `json:"token"`
	jwt.StandardClaims
}

// DependencyInjection ...
type DependencyInjection struct {
	us services.LogService
}

// ValidateJwtMiddlewear ... This function helps to log the stuff
func (di DependencyInjection) ValidateJwtMiddlewear(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

		authToken := r.Header.Get("Authorization")
		splitToekn := strings.Split(authToken, "Bearer")
		if len(splitToekn) != 2 {
			utils.RespondError(w, 401, "Invalid Token")
			return
		}
		tokenSring := strings.TrimSpace(splitToekn[1])

		token, err := jwt.ParseWithClaims(tokenSring, &CustomClaims{}, func(token *jwt.Token) (interface{}, error) {
			return []byte([]byte(os.Getenv("JWT_SECRATE"))), nil
		})

		if claims, ok := token.Claims.(*CustomClaims); ok && token.Valid {
			//todo: add logic to check token xpired or not
			fmt.Printf("%v %v", claims.Token, claims.StandardClaims.ExpiresAt)
		} else {
			//todo : add some changes for catching error
			fmt.Println(err)
		}

		next.ServeHTTP(w, r)
	})
}
