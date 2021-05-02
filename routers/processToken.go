package routers

import (
	"errors"
	"strings"

	"github.com/MessiasJunio/twittor-server/db"
	"github.com/MessiasJunio/twittor-server/models"
	jwt "github.com/dgrijalva/jwt-go"
)

// Email value of Email used in EndPoints
var Email string

//UserID is the ID returned of the model, that uses all the EnPoints
var UserID string

// ProcessToken process token of extracted values
func ProcessToken(tk string) (*models.Claim, bool, string, error) {
	myKey := []byte("MastersOfDevelopment")
	claims := &models.Claim{}

	splitToken := strings.Split(tk, "Bearer")
	if len(splitToken) != 2 {
		return claims, false, string(""), errors.New("Invalid token format")
	}

	tk = strings.TrimSpace(splitToken[1])

	tkn, err := jwt.ParseWithClaims(tk, claims, func(token *jwt.Token) (interface{}, error) {
		return myKey, nil
	})

	if err == nil {
		_, found, _ := db.CheckIfUserExists(claims.Email)
		if found == true {
			Email = claims.Email
			UserID = claims.ID.Hex()
		}
		return claims, found, UserID, nil
	}
	if !tkn.Valid {
		return claims, false, string(""), errors.New("Invalid token")
	}

	return claims, false, string(""), err
}
