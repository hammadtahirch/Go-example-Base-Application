package middlewares

import (
	"net/http"
)

// LoggingMiddleware ... This function helps to log the stuff
func LoggingMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// todo: write code here.
		//end
		next.ServeHTTP(w, r)
	})
}
