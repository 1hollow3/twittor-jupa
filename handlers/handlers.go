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

	PORT := os.Getenv("PORT")
	if PORT == "" {
		PORT = "8080"
	}
	handler := cors.AllowAll().Handler(router)
	log.Fatal(http.ListenAndServe(":"+PORT, handler))
}
