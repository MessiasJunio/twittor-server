package routers

import (
	"encoding/json"
	"net/http"
	"time"

	"github.com/MessiasJunio/twittor-server/db"
	"github.com/MessiasJunio/twittor-server/models"
)

//SaveTweet allows save tweet in the database
func SaveTweet(w http.ResponseWriter, r *http.Request) {
	var message models.Tweet
	err := json.NewDecoder(r.Body).Decode(&message)

	registry := models.SaveTweet{
		UserID:  UserID,
		Message: message.Message,
		Date:    time.Now(),
	}

	_, status, err := db.InsertTweet(registry)
	if err != nil {
		http.Error(w, "An error occurred while trying to insert the record, please try again"+err.Error(), 400)
		return
	}

	if status == false {
		http.Error(w, "The Tweet could not be inserted", 400)
		return
	}

	w.WriteHeader(http.StatusCreated)
}
