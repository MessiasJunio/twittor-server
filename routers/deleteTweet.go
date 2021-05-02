package routers

import (
	"net/http"

	"github.com/MessiasJunio/twittor-server/db"
)

func DeleteTweet(w http.ResponseWriter, r *http.Request) {
	ID := r.URL.Query().Get("id")
	if len(ID) < 1 {
		http.Error(w, "Should send parameter id", http.StatusBadRequest)
		return
	}
	err := db.DeleteTweet(ID, UserID)
	if err != nil {
		http.Error(w, "An error occurred while trying to use Tweet "+err.Error(), http.StatusBadRequest)
		return
	}

	w.Header().Set("Content-type", "application/json")
	w.WriteHeader(http.StatusCreated)
}
