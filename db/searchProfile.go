package db

import (
	"context"
	"fmt"
	"time"

	"github.com/MessiasJunio/twittor-server/models"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

// SearchProfile search a profile in the database
func SearchProfile(ID string) (models.User, error) {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*15)
	defer cancel()

	db := ConnectDB().Database("twittor-react")
	col := db.Collection("users")

	var profile models.User

	objID, _ := primitive.ObjectIDFromHex(ID)

	condition := bson.M{
		"_id": objID,
	}

	err := col.FindOne(ctx, condition).Decode(&profile)
	profile.Password = ""

	if err != nil {
		fmt.Println("Register not found" + err.Error())
		return profile, err
	}

	return profile, nil
}
