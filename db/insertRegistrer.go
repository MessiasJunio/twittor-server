package db

import (
	"context"
	"time"

	"github.com/MessiasJunio/twittor-server/models"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

// InsertRegister insert user data
func InsertRegister(u models.User) (string, bool, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 15*time.Second)

	defer cancel()

	db := ConnectDB().Database("twittor-react")

	col := db.Collection("users")

	u.Password, _ = EncryptPassword(u.Password)

	result, err := col.InsertOne(ctx, u)
	if err != nil {
		return "", false, err
	}

	ObjID, _ := result.InsertedID.(primitive.ObjectID)

	return ObjID.String(), true, nil
}
