package routers

import (
	"net/http"

	"github.com/MessiasJunio/twittor/db"
	"github.com/MessiasJunio/twittor/models"
)

func LowRatio(w http.ResponseWriter, r *http.Request) {
	ID := r.URL.Query().Get("id")

	var t models.Ratio
	t.UserID = UserID
	t.UserRatioID = ID

	status, err := db.DeleteRatio(t)
	if err != nil {
		http.Error(w, "An error occurred while trying to delete ratio "+err.Error(), http.StatusBadRequest)
		return
	}
	if status == false {
		http.Error(w, "The ratio could not be deleted "+err.Error(), http.StatusBadRequest)
		return
	}

	w.WriteHeader(http.StatusCreated)
}
