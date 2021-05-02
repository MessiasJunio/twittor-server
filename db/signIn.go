package db

import (
	"github.com/MessiasJunio/twittor-server/models"
	"golang.org/x/crypto/bcrypt"
)

// SignIn check login account
func SignIn(email string, password string) (models.User, bool) {
	user, found, _ := CheckIfUserExists(email)

	if found == false {
		return user, false
	}

	passwordBytes := []byte(password)
	passwordDB := []byte(user.Password)

	err := bcrypt.CompareHashAndPassword(passwordDB, passwordBytes)

	if err != nil {
		return user, false
	}

	return user, true
}
