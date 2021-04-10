package handlers

import (
	"log"
	"net/http"
	"os"

	"github.com/MessiasJunio/twittor/middlewares"
	"github.com/MessiasJunio/twittor/routers"
	"github.com/gorilla/mux"
	"github.com/rs/cors"
)

// Managers set port, handler and listening server
func Managers() {
	router := mux.NewRouter()

	router.HandleFunc("/register", middlewares.CheckDB(routers.Register)).Methods("POST")
	router.HandleFunc("/login", middlewares.CheckDB(routers.Login)).Methods("POST")

	PORT := os.Getenv("PORT")

	if PORT == "" {
		PORT = "8080"
	}
	handler := cors.AllowAll().Handler(router)

	log.Fatal(http.ListenAndServe(":"+PORT, handler))
}
