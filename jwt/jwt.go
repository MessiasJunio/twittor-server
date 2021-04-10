package jwt

import (
	"time"

	"github.com/MessiasJunio/twittor/models"
	"github.com/dgrijalva/jwt-go"
)

func GenerateJWT(t models.User) (string, error) {
	myKey := []byte("MastersDevelopers")

	payload := jwt.MapClaims{

		"email":     t.Email,
		"name":      t.Name,
		"surname":   t.Surname,
		"birthDate": t.BirthDate,
		"biography": t.Biography,
		"location":  t.Location,
		"webSite":   t.WebSite,
		"_id":       t.ID.Hex(),
		"exp":       time.Now().Add(time.Hour * 24).Unix(),
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, payload)
	tokenStr, err := token.SignedString(myKey)
	if err != nil {
		return tokenStr, err
	}
	return tokenStr, nil
}