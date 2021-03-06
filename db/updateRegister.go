package db

import (
	"context"
	"time"

	"github.com/MessiasJunio/twittor-server/models"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func UpdateRegistry(u models.User, ID string) (bool, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 15*time.Second)
	defer cancel()

	db := ConnectDB().Database("twittor-react")
	col := db.Collection("users")

	registry := make(map[string]interface{})

	if len(u.Name) > 0 {
		registry["name"] = u.Name
	}

	if len(u.LastName) > 0 {
		registry["lastName"] = u.LastName
	}

	registry["birthDate"] = u.BirthDate

	if len(u.Avatar) > 0 {
		registry["avatar"] = u.Avatar
	}

	if len(u.Banner) > 0 {
		registry["banner"] = u.Banner
	}

	if len(u.Biography) > 0 {
		registry["biography"] = u.Biography
	}

	if len(u.Location) > 0 {
		registry["location"] = u.Location
	}

	if len(u.WebSite) > 0 {
		registry["website"] = u.WebSite
	}

	updateString := bson.M{
		"$set": registry,
	}

	objID, _ := primitive.ObjectIDFromHex(ID)
	filter := bson.M{"_id": bson.M{"$eq": objID}}

	_, err := col.UpdateOne(ctx, filter, updateString)
	if err != nil {
		return false, err
	}

	return true, nil
}
