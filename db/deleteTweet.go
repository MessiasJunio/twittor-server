package db

import (
	"context"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

//DeleteTweet delete a tweet
func DeleteTweet(ID string, UserID string) error {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*15)
	defer cancel()

	db := ConnectDB().Database("twittor-react")
	col := db.Collection("tweet")

	objID, _ := primitive.ObjectIDFromHex(ID)
	condition := bson.M{
		"_id":    objID,
		"userid": UserID,
	}

	_, err := col.DeleteOne(ctx, condition)
	return err
}
