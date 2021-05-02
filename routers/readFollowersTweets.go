package routers

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/MessiasJunio/twittor-server/db"
)

func ReadFollowersTweets(w http.ResponseWriter, r *http.Request) {

	if len(r.URL.Query().Get("page")) < 1 {
		http.Error(w, "You must send the page parameter", http.StatusBadRequest)
		return
	}
	page, err := strconv.Atoi(r.URL.Query().Get("page"))
	if err != nil {
		http.Error(w, "You must send the page parameter as an integer greater than 0", http.StatusBadRequest)
		return
	}

	response, correct := db.ReadFollowersTweets(UserID, page)
	if correct == false {
		http.Error(w, "Error Reading tweets", http.StatusBadRequest)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(response)
}
