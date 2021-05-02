package routers

import (
	"net/http"

	"github.com/MessiasJunio/twittor-server/db"
	"github.com/MessiasJunio/twittor-server/models"
)

func HighRatio(w http.ResponseWriter, r *http.Request) {
	ID := r.URL.Query().Get("id")
	if len(ID) < 1 {
		http.Error(w, "Parameter ID is required", http.StatusBadRequest)
		return
	}

	var t models.Ratio
	t.UserID = UserID
	t.UserRatioID = ID

	status, err := db.InsertRatio(t)
	if err != nil {
		http.Error(w, "An error occurred while trying to insert ratio", http.StatusBadRequest)
		return
	}

	if status == false {
		http.Error(w, "Ratio can not be inserted "+err.Error(), http.StatusBadRequest)
		return
	}

	w.WriteHeader(http.StatusCreated)
}
