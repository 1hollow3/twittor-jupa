package routers

import (
	"encoding/json"
	"github.com/1hollow3/twittor-jupa/db"
	"github.com/1hollow3/twittor-jupa/models"
	"net/http"
)

func ChangeProfile(writer http.ResponseWriter, reader *http.Request) {

	var t models.User

	err := json.NewDecoder(reader.Body).Decode(&t)
	if err != nil {
		http.Error(writer, "Incorrect data "+err.Error(), 400)
		return
	}

	var status bool
	status, err = db.ModifyUser(t, IDUser)
	if err != nil {
		http.Error(writer, "An error has happened at the modification of the register, try again "+err.Error(), 500)
	}
	if status == false {
		http.Error(writer, "It couldn't change the register of the user", 500)
	}
	writer.WriteHeader(http.StatusCreated)
}
