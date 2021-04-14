package routers

import (
	"encoding/json"
	"net/http"

	"github.com/MessiasJunio/twittor/db"
	"github.com/MessiasJunio/twittor/models"
)

//Register a function that creates a user database register
func Register(w http.ResponseWriter, r *http.Request) {
	var t models.User
	err := json.NewDecoder(r.Body).Decode(&t)

	if err != nil {
		http.Error(w, "Data Received Error "+err.Error(), 400)
		return
	}

	if len(t.Email) == 0 {
		http.Error(w, "User Email Required", 400)
	}

	if len(t.Password) < 6 {
		http.Error(w, "Password must have at least 6 characters", 400)
	}

	_, finded, _ := db.CheckIfUserExists(t.Email)
	if finded == true {
		http.Error(w, "User email already exists", 400)
		return
	}

	_, status, err := db.InsertRegister(t)
	if err != nil {
		http.Error(w, "Error when try to insert a new user"+err.Error(), 400)
		return
	}

	if status == false {
		http.Error(w, "Can not insert user"+err.Error(), 400)
		return
	}

	w.WriteHeader(http.StatusCreated)
}
