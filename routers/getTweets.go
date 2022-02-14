package routers

import (
	"encoding/json"
	"github.com/1hollow3/twittor-jupa/db"
	"net/http"
	"strconv"
)

func GetTweets(w http.ResponseWriter, r *http.Request) {

	ID := r.URL.Query().Get("id")
	if len(ID) < 1 {
		http.Error(w, "id parameter is requested", 400)
		return
	}

	if len(r.URL.Query().Get("page")) < 1 {
		http.Error(w, "page parameter is requested", 400)
		return
	}
	page, err := strconv.Atoi(r.URL.Query().Get("page"))
	if err != nil {
		http.Error(w, "page parameter must be a value greater than 0", 400)
		return
	}

	response, ok := db.ReadTweets(ID, int64(page))
	if ok == false {
		http.Error(w, "Error reading tweets", 500)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(200)
	json.NewEncoder(w).Encode(response)

}
