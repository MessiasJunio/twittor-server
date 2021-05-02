package db

import (
	"context"
	"time"

	"github.com/MessiasJunio/twittor-server/models"
)

func InsertRatio(t models.Ratio) (bool, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 15*time.Second)
	defer cancel()

	db := ConnectDB().Database("twittor-react")
	col := db.Collection("ratio")
	_, err := col.InsertOne(ctx, t)
	if err != nil {
		return false, err
	}
	return true, nil
}
