package routers

import (
	"encoding/json"
	"github.com/1hollow3/twittor-jupa/db"
	"net/http"
	"strconv"
)

func ReadRelatedTweets(w http.ResponseWriter, r *http.Request) {

	if len(r.URL.Query().Get("page")) < 1 {
		http.Error(w, "page parameter is requested", 400)
		return
	}

	page, err := strconv.Atoi(r.URL.Query().Get("page"))
	if err != nil || page < 1{
		http.Error(w, "page parameter must be a value greater than 0", 400)
		return
	}

	result, right:= db.ReadTweetsFollowers(IDUser, page)

	if !right {
		http.Error(w, "Error reading tweets", 500)
		return
	}


	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(200)
	json.NewEncoder(w).Encode(result)

}
