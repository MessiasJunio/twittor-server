package routers

import (
	"encoding/json"
	"net/http"

	"github.com/MessiasJunio/twittor-server/db"
	"github.com/MessiasJunio/twittor-server/models"
)

func ReadRatio(w http.ResponseWriter, r *http.Request) {
	ID := r.URL.Query().Get("id")

	var t models.Ratio
	t.UserID = UserID
	t.UserRatioID = ID

	var resp models.StatusRatio
	status, err := db.ReadRatio(t)
	if err != nil || status == false {
		resp.Status = false
	} else {
		resp.Status = true
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(resp)
}
