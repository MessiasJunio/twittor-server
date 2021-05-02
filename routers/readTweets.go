package routers

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/MessiasJunio/twittor-server/db"
)

//ReadTweets read the tweets
func ReadTweets(w http.ResponseWriter, r *http.Request) {
	ID := r.URL.Query().Get("id")
	if len(ID) < 1 {
		http.Error(w, "Should send parameter id", http.StatusBadRequest)
		return
	}

	if len(r.URL.Query().Get("page")) < 1 {
		http.Error(w, "Should send parameter page", http.StatusBadRequest)
		return
	}

	page, err := strconv.Atoi(r.URL.Query().Get("page"))
	if err != nil {
		http.Error(w, "Should send the parameter page with a value greater than 0", http.StatusBadRequest)
		return
	}

	pag := int64(page)
	response, correct := db.ReadTweets(ID, pag)
	if correct == false {
		http.Error(w, "Error while reading the tweets", http.StatusBadRequest)
	}
	w.Header().Set("Content-type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(response)
}
