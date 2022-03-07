package routers

import (
	"github.com/1hollow3/twittor-jupa/db"
	"net/http"
)

func RemoveTweet(w http.ResponseWriter, r *http.Request)  {
	ID := r.URL.Query().Get("id")
	if len(ID)<1 {
		http.Error(w, "ID parameter must be send", http.StatusBadRequest)
		return
	}

	err := db.DeleteTweet(ID, IDUser)

	if err != nil {
		http.Error(w, "Error while deleting tweet from db "+ err.Error(), http.StatusBadRequest)
		return
	}

	w.WriteHeader(http.StatusOK)

}
