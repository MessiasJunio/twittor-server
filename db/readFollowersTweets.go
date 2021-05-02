package db

import (
	"context"
	"time"

	"github.com/MessiasJunio/twittor-server/models"
	"go.mongodb.org/mongo-driver/bson"
)

func ReadFollowersTweets(ID string, page int) ([]models.ReturnFollowersTweets, bool) {
	ctx, cancel := context.WithTimeout(context.Background(), 15*time.Second)
	defer cancel()

	db := ConnectDB().Database("twittor-react")
	col := db.Collection("ratio")

	skip := (page - 1) * 20

	condition := make([]bson.M, 0)
	condition = append(condition, bson.M{"$match": bson.M{"userid": ID}})

	condition = append(condition, bson.M{
		"$lookup": bson.M{
			"from":         "tweet",
			"localField":   "userratioid",
			"foreignField": "userid",
			"as":           "tweet",
		}})
	condition = append(condition, bson.M{"$unwind": "$tweet"})
	condition = append(condition, bson.M{"$sort": bson.M{"tweet.date": -1}})
	condition = append(condition, bson.M{"$skip": skip})
	condition = append(condition, bson.M{"$limit": 20})

	cursor, err := col.Aggregate(ctx, condition)
	var result []models.ReturnFollowersTweets
	err = cursor.All(ctx, &result)
	if err != nil {
		return result, false
	}
	return result, true

}
