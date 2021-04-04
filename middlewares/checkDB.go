package middleWare

import (
	"net/http"

	"github.com/MessiasJunio/twittor/db"
)

func CheckDB(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if db.CheckConnection() == 0 {
			http.Error(w, "Database Connection Lost", 500)
		}
		next.ServeHTTP(w, r)
	}
}
