package db

import (
	"context"
	"time"

	"github.com/MessiasJunio/twittor-server/models"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func ReadTweets(ID string, page int64) ([]*models.ReturnTweet, bool) {
	ctx, cancel := context.WithTimeout(context.Background(), 15*time.Second)
	defer cancel()

	db := ConnectDB().Database("twittor-react")
	col := db.Collection("tweet")

	var result []*models.ReturnTweet

	condition := bson.M{
		"userid": ID,
	}

	optionsMongo := options.Find()
	optionsMongo.SetLimit(20)
	optionsMongo.SetSort(bson.D{{Key: "date", Value: -1}})
	optionsMongo.SetSkip((page - 1) * 20)

	cursor, err := col.Find(ctx, condition, optionsMongo)
	if err != nil {
		return result, false
	}
	for cursor.Next(context.TODO()) {
		var registry models.ReturnTweet
		err := cursor.Decode(&registry)
		if err != nil {
			return result, false
		}
		result = append(result, &registry)
	}

	return result, true
}
