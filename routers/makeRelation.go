package routers

import (
	"github.com/1hollow3/twittor-jupa/db"
	"github.com/1hollow3/twittor-jupa/models"
	"net/http"
)

func MakeRelation(w http.ResponseWriter, r *http.Request) {

	ID := r.URL.Query().Get("id")
	if len(ID) < 1 {
		http.Error(w, "id parameter is requested", 400)
		return
	}

	var t models.Relation
	t.UserID = IDUser
	t.RelatedUserID = ID

	status, err := db.InsertRelation(t)
	if err != nil {
		http.Error(w, "Error saving the relation" + err.Error(), 500)
		return
	}
	if !status {
		http.Error(w, "The relation couldn't be saved", 500)
		return
	}

	w.WriteHeader(http.StatusCreated)

}
