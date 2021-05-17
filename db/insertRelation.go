package db

import (
	"github.com/MessiasJunio/twittor-server/models"
)

func InsertRelation(t models.Relation) (bool, error) {
	ctx, cancel := CreateDefaultContext()
	defer cancel()

	collection := GetCollection("relation")

	_, err := collection.InsertOne(ctx, t)
	if err != nil {
		return false, err
	}
	return true, nil
}
