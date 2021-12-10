package routers

import (
	"encoding/json"
	"github.com/1hollow3/twittor-jupa/db"
	"github.com/1hollow3/twittor-jupa/models"
	"net/http"
)

func Register(w http.ResponseWriter, r *http.Request) {

	var t models.User
	err := json.NewDecoder(r.Body).Decode(&t)
	if err != nil {
		http.Error(w, "Error in requested Data:"+err.Error(), 400)
		return
	}

	if len(t.Email) == 0 {
		http.Error(w, "Email is required", 400)
	}

	if len(t.Password) < 6 {
		http.Error(w, "Password must be at least 6 characters long", 400)
	}

	_, encontrado, _ := db.UserExist(t.Email)
	if encontrado == true {
		http.Error(w, "There's already an user registered with that email", 400)
		return
	}

	_, status, err := db.NewRegister(t)
	if err != nil {
		http.Error(w, "An error happened trying to create the user "+err.Error(), 400)
	}
	if status == false {
		http.Error(w, "Error saving the user in db " + err.Error(), 400 )
	}

	w.WriteHeader(http.StatusCreated)
}
