package db

import (
	"context"
	"time"

	"github.com/MessiasJunio/twittor/models"
	"go.mongodb.org/mongo-driver/bson"
)

// CheckIfUserExists receives a email and check if user exists in Database
func CheckIfUserExists(email string) (models.User, bool, string) {
	ctx, cancel := context.WithTimeout(context.Background(), 15*time.Second)
	defer cancel()

	db := ConnectDB().Database("twittor")
	col := db.Collection("users")

	condition := bson.M{"email": email}

	var result models.User

	err := col.FindOne(ctx, condition).Decode(&result)
	ID := result.ID.Hex()
	if err != nil {
		return result, false, ID
	}
	return result, true, ID
}
