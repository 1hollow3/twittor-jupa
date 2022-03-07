package routers

import (
	"github.com/1hollow3/twittor-jupa/db"
	"github.com/1hollow3/twittor-jupa/models"
	"net/http"
)

func RemoveRelation(w http.ResponseWriter, r *http.Request) {

	ID := r.URL.Query().Get("id")
	if len(ID) < 1 {
		http.Error(w, "ID parameter must be send", http.StatusBadRequest)
		return
	}

	var t models.Relation
	t.UserID = IDUser
	t.RelatedUserID = ID

	err := db.DeleteRelation(t)
	if err != nil {
		http.Error(w, "Error deleting the relation", 500)
	}

	w.WriteHeader(http.StatusOK)

}
