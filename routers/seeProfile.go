package routers

import (
	"encoding/json"
	"net/http"

	"github.com/1hollow3/twittor-jupa/db"
)

// SeeProfile is to extract the profile values
func SeeProfile(w http.ResponseWriter, r *http.Request)  {

	ID := r.URL.Query().Get("id")
	if len(ID) < 1{
		http.Error(w, "ID is a required value", http.StatusBadRequest)
		return
	}

	profile, err := db.SearchProfile(ID)
	if err != nil {
		http.Error(w, "An error happened in the search of the register " + err.Error(), 400)
		return
	}

	w.Header().Set("context-type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(profile )
}