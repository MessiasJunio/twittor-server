package main

import (
	"log"

	"github.com/MessiasJunio/twittor/db"
	"github.com/MessiasJunio/twittor/handlers"
)

func main() {
	if db.CheckConnection() == 0 {
		log.Fatal("No Database connection")
	}

	handlers.Managers()
}
