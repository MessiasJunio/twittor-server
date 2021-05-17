package db

import (
	"context"
	"fmt"
	"log"
	"os"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

// ConnectDB allows connect in the DB
func ConnectDB() *mongo.Client {
	var URI = fmt.Sprintf("mongodb+srv://twittor:" + os.Getenv("DB_PASSWORD") + "@twittor.wwrlx.mongodb.net/myFirstDatabase?retryWrites=true&w=majority")

	var clientOptions = options.Client().ApplyURI(URI)

	client, err := mongo.Connect(context.TODO(), clientOptions)

	if err != nil {
		log.Fatal(err.Error())
		return client
	}

	err = client.Ping(context.TODO(), nil)

	if err != nil {
		log.Fatal(err.Error())
		return client
	}

	log.Println("Database Connection Success")
	return client
}

// CheckConnection test ping in the DB
func CheckConnection() int {
	err := ConnectDB().Ping(context.TODO(), nil)

	if err != nil {
		return 0
	}
	return 1
}

func CreateDefaultContext() (context.Context, context.CancelFunc) {
	ctx, cancel := context.WithTimeout(context.Background(), 15*time.Second)
	return ctx, cancel
}

func GetCollection(collection string) *mongo.Collection {
	database := ConnectDB().Database("twittor-react")
	colllection := database.Collection(collection)
	return colllection
}
