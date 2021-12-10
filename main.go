package main

import (
	"log"

	"github.com/1hollow3/twittor-jupa/db"
	"github.com/1hollow3/twittor-jupa/handlers"
)

func main() {
	if db.CheckConnection() == 0 {
		log.Fatal("No Connection to database")
		return
	}
	handlers.Handlers()
}
