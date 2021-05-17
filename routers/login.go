package routers

import (
	"encoding/json"
	"net/http"
	"time"

	"github.com/MessiasJunio/twittor-server/db"
	JWT "github.com/MessiasJunio/twittor-server/jwt"
	"github.com/MessiasJunio/twittor-server/models"
)

/* Login sign in account */
func Login(w http.ResponseWriter, r *http.Request) {
	w.Header().Add("content-type", "application/json")

	var t models.User

	err := json.NewDecoder(r.Body).Decode(&t)

	if err != nil {
		http.Error(w, "Invalid User or Password"+err.Error(), 400)
		return
	}

	if len(t.Email) == 0 {
		http.Error(w, "Required Email", 400)
	}

	document, exists := db.SignIn(t.Email, t.Password)

	if exists == false {
		http.Error(w, "Invalid User or Password", 400)
		return
	}

	jwtKey, err := JWT.GenerateJWT(document)
	if err != nil {
		http.Error(w, "Error trying to generate Token", 400)
	}

	resp := models.LoginResponse{
		Token: jwtKey,
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)

	json.NewEncoder(w).Encode(resp)

	expirelationnTime := time.Now().Add(24 * time.Hour)

	http.SetCookie(w, &http.Cookie{
		Name:    "token",
		Value:   jwtKey,
		Expires: expirelationnTime,
	})
}
