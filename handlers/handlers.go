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
	router.HandleFunc("/viewprofile", middlewares.CheckDB(middlewares.ValidateJWT(routers.ViewProfile))).Methods("GET")
	router.HandleFunc("/updateProfile", middlewares.CheckDB(middlewares.ValidateJWT(routers.UpdateProfile))).Methods("PUT")
	router.HandleFunc("/tweet", middlewares.CheckDB(middlewares.ValidateJWT(routers.SaveTweet))).Methods("POST")
	router.HandleFunc("/readTweets", middlewares.CheckDB(middlewares.ValidateJWT(routers.ReadTweets))).Methods("GET")
	router.HandleFunc("/deleteTweet", middlewares.CheckDB(middlewares.ValidateJWT(routers.DeleteTweet))).Methods("DELETE")

	router.HandleFunc("/uploadAvatar", middlewares.CheckDB(middlewares.ValidateJWT(routers.UploadAvatar))).Methods("POST")
	router.HandleFunc("/getAvatar", middlewares.CheckDB(routers.GetAvatar)).Methods("GET")
	router.HandleFunc("/uploadBanner", middlewares.CheckDB(middlewares.ValidateJWT(routers.UploadBanner))).Methods("POST")
	router.HandleFunc("/getBanner", middlewares.CheckDB(routers.GetBanner)).Methods("GET")

	router.HandleFunc("/highRatio", middlewares.CheckDB(middlewares.ValidateJWT(routers.HighRatio))).Methods("POST")
	router.HandleFunc("/lowRatio", middlewares.CheckDB(middlewares.ValidateJWT(routers.LowRatio))).Methods("DELETE")
	router.HandleFunc("/readRatio", middlewares.CheckDB(middlewares.ValidateJWT(routers.ReadRatio))).Methods("GET")

	router.HandleFunc("/usersList", middlewares.CheckDB(middlewares.ValidateJWT(routers.UsersList))).Methods("GET")

	PORT := os.Getenv("PORT")

	if PORT == "" {
		PORT = "8080"
	}
	handler := cors.AllowAll().Handler(router)

	log.Fatal(http.ListenAndServe(":"+PORT, handler))
}
