package routers

import (
	"encoding/json"
	"net/http"

	"github.com/MessiasJunio/twittor-server/db"
	"github.com/MessiasJunio/twittor-server/models"
)

// UpdateProfile updates user profile
func UpdateProfile(w http.ResponseWriter, r *http.Request) {
	var t models.User

	err := json.NewDecoder(r.Body).Decode(&t)
	if err != nil {
		http.Error(w, "Incorrect Data"+err.Error(), 400)
	}

	var status bool

	status, err = db.UpdateRegistry(t, UserID)
	if err != nil {
		http.Error(w, "An error occurred while trying to modify the registry. Try again "+err.Error(), 400)
	}
	if status == false {
		http.Error(w, "Unable to modify the user's registry ", 400)
		return
	}

	w.WriteHeader(http.StatusCreated)
}
