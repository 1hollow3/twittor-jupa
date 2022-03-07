package routers

import (
	"encoding/json"
	"github.com/1hollow3/twittor-jupa/db"
	"net/http"
	"strconv"
)

func UserList(w http.ResponseWriter, r *http.Request) {

	typeUser := r.URL.Query().Get("type")
	page := r.URL.Query().Get("page")
	search := r.URL.Query().Get("search")

	pagTemp, err := strconv.Atoi(page)


	pag := int64(pagTemp)
	if err != nil || pag < 1 {
		http.Error(w, "Page parameter must be send with a value greater than 0", 400)
		return
	}

	result, status := db.ReadAllUsers(IDUser, pag, search, typeUser)
	if status == false {
		http.Error(w, "Error reading the users", 500)
		return
	}
	w.Header().Set("Content-type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(result)
}
