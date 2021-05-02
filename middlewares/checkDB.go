package middlewares

import (
	"net/http"

	"github.com/MessiasJunio/twittor-server/db"
)

func CheckDB(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if db.CheckConnection() == 0 {
			http.Error(w, "Database Connection was Lost", 500)
		}
		next.ServeHTTP(w, r)
	}
}
