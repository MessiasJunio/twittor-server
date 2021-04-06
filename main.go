package main

import (
	"log"

	"github.com/MessiasJunio/twittor/db"
	"github.com/MessiasJunio/twittor/handlers"
	"github.com/joho/godotenv"
)

func init() {
	loadEnv()
}

func main() {
	if db.CheckConnection() == 0 {
		log.Fatal("No Database connection")
	}

	handlers.Managers()
}

func loadEnv() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}
}
