package handlers

import (
	"log"
	"net/http"
	"os"

	"github.com/1hollow3/twittor-jupa/middlew"
	"github.com/1hollow3/twittor-jupa/routers"
	"github.com/gorilla/mux"
	"github.com/rs/cors"
)

func Handlers() {
	router := mux.NewRouter()

	router.HandleFunc("/register", middlew.CheckDB(routers.Register)).Methods("POST")
	router.HandleFunc("/login", middlew.CheckDB(routers.Login)).Methods("POST")
	router.HandleFunc("/seeProfile", middlew.ValidateJWT(routers.SeeProfile)).Methods("GET")
	router.HandleFunc("/changeProfile", middlew.ValidateJWT(routers.ChangeProfile)).Methods("PUT")
	router.HandleFunc("/saveTweet", middlew.ValidateJWT(routers.SaveTweet)).Methods("POST")
	router.HandleFunc("/readTweet", middlew.ValidateJWT(routers.GetTweets)).Methods("GET")

	PORT := os.Getenv("PORT")
	if PORT == "" {
		PORT = "8080"
	}
	handler := cors.AllowAll().Handler(router)
	log.Fatal(http.ListenAndServe(":"+PORT, handler))
}
