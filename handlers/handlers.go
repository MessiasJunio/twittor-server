package handlers

import (
	"log"
	"net/http"
	"os"

	"github.com/gorilla/middleWares"
	"github.com/gorilla/mux"
	"github.com/gorilla/routers"
	"github.com/rs/cors"
)

// Managers set port, handler and listening server
func Managers() {
	router := mux.NewRouter()

	router.HandleFunc("/register", middleWares.CheckDB(routers.Register)).Methods("POST")

	PORT := os.Getenv("PORT")

	if PORT == "" {
		PORT = "8080"
	}
	handler := cors.AllowAll().Handler(router)

	log.Fatal(http.ListenAndServe(":"+PORT, handler))
}
