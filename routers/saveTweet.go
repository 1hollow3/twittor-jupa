package routers

import (
	"encoding/json"
	"github.com/1hollow3/twittor-jupa/db"
	"github.com/1hollow3/twittor-jupa/models"
	"net/http"
	"time"
)

func SaveTweet(w http.ResponseWriter, r *http.Request) {

	var message models.Tweet
	err := json.NewDecoder(r.Body).Decode(&message)
	if err != nil {
		http.Error(w, "Error in getting the message "+err.Error(), 400)
		return
	}

	register := models.RecordTweet{
		UserID:  IDUser,
		Date:    time.Now(),
		Message: message.Message,
	}

	_, status, err := db.InsertTweet(register)
	if err != nil {
		http.Error(w, "Error recording the message"+err.Error(), 400)
		return
	}

	if !status {
		http.Error(w, "The message couldn't be recorded "+err.Error(), 400)
		return
	}
	w.WriteHeader(http.StatusCreated)

}
