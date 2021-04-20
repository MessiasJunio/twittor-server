package db

import (
	"context"
	"time"

	"github.com/MessiasJunio/twittor/models"
	"go.mongodb.org/mongo-driver/bson"
)

func ReadRatio(t models.Ratio) (bool, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 15*time.Second)
	defer cancel()

	db := ConnectDB().Database("twittor-react")
	col := db.Collection("ratio")

	condition := bson.M{
		"userid":      t.UserID,
		"userratioid": t.UserRatioID,
	}

	var result models.Ratio
	err := col.FindOne(ctx, condition).Decode(&result)
	if err != nil {
		return false, err
	}

	return true, nil
}
