package middlew

import (
	"net/http"

	"github.com/1hollow3/twittor-jupa/routers"
)

// ValidateJWT is a Middleware for JWT validation
func ValidateJWT(next http.HandlerFunc) http.HandlerFunc {
	return func(writer http.ResponseWriter, request *http.Request) {
		_, _, _, err := routers.ProcessToken(request.Header.Get("Authorization"))
		if err != nil {
			http.Error(writer, "Error on Token! "+err.Error(), http.StatusBadRequest)
			return
		}
		next.ServeHTTP(writer, request)
	}
}
