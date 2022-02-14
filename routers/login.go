package routers

import (
	"encoding/json"
	"net/http"
	"time"

	"github.com/1hollow3/twittor-jupa/db"
	"github.com/1hollow3/twittor-jupa/jwt"
	"github.com/1hollow3/twittor-jupa/models"
)

func Login(w http.ResponseWriter, r *http.Request) {
	w.Header().Add("content-type", "application/json")

	var t models.User

	err := json.NewDecoder(r.Body).Decode(&t)
	if err != nil {
		http.Error(w, "Username and/or password are invalid "+err.Error(), 400)
		return
	}
	if len(t.Email) == 0 {
		http.Error(w, "Email required "+err.Error(), 400)
		return
	}

	document, exist := db.Login(t.Email, t.Password)

	if exist == false {
		http.Error(w, "Username and/or password are invalid", 400)
		return
	}

	jwtKey, err := jwt.GenerateJWT(document)
	if err != nil {
		http.Error(w, "Error on the generation of Token " + err.Error(), 500)
		return
	}

	resp := models.LoginResponse {
		Token : jwtKey,
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(resp)

	expirationTime := time.Now().Add(24 * time.Hour)
	http.SetCookie(w, &http.Cookie{
		Name: "token",
		Value: jwtKey,
		Expires: expirationTime	})
}
