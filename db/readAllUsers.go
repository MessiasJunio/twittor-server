package db

import (
	"context"
	"time"

	"github.com/MessiasJunio/twittor/models"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func ReadAllUsers(ID string, page int64, search string, userType string) ([]*models.User, bool) {
	ctx, cancel := context.WithTimeout(context.Background(), 15*time.Second)
	defer cancel()

	db := ConnectDB().Database("twittor-react")
	col := db.Collection("users")

	var results []*models.User

	findOptions := options.Find()
	findOptions.SetSkip((page - 1) * 20)
	findOptions.SetLimit(20)

	query := bson.M{
		"name": bson.M{"$regex": `(?i)` + search},
	}

	cur, err := col.Find(ctx, query, findOptions)
	if err != nil {
		return results, false
	}

	var found, insert bool

	for cur.Next(ctx) {
		var s models.User
		err := cur.Decode(&s)
		if err != nil {
			return results, false
		}

		var r models.Ratio
		r.UserID = ID
		r.UserRatioID = s.ID.Hex()

		insert = false

		found, err = ReadRatio(r)
		if userType == "new" && found == false {
			insert = true
		}

		if userType == "follow" && found == true {
			insert = true
		}

		if r.UserRatioID == ID {
			insert = false
		}

		if insert == true {
			s.Password = ""
			s.Biography = ""
			s.WebSite = ""
			s.Location = ""
			s.Banner = ""
			s.Email = ""

			results = append(results, &s)
		}
	}

	err = cur.Err()
	if err != nil {
		return results, false
	}

	cur.Close(ctx)
	return results, true
}
