package middlew

import (
	"net/http"

	"github.com/1hollow3/twittor-jupa/db"
)

func CheckDB(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if db.CheckConnection() == 0 {
			http.Error(w, "Lost connection with database", 500)
			return
		}
		next.ServeHTTP(w, r)
	}
}
