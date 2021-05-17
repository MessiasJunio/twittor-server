package db

import (
	"github.com/MessiasJunio/twittor-server/models"
	"go.mongodb.org/mongo-driver/bson"
)

func ReadRelation(t models.Relation) (bool, error) {
	ctx, cancel := CreateDefaultContext()
	defer cancel()

	collection := GetCollection("relation")

	condition := bson.M{
		"userid":         t.UserID,
		"userrelationid": t.UserRelationID,
	}

	var result models.Relation
	err := collection.FindOne(ctx, condition).Decode(&result)
	if err != nil {
		return false, err
	}

	return true, nil
}
