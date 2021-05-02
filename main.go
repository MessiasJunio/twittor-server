package main

import (
	"log"

	"github.com/MessiasJunio/twittor-server/db"
	"github.com/MessiasJunio/twittor-server/handlers"
	"github.com/joho/godotenv"
)

func loadEnv() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}
}

func init() {
	loadEnv()
}

func main() {
	if db.CheckConnection() == 0 {
		log.Fatal("No Database connection")
		return
	}

	handlers.Managers()
}
