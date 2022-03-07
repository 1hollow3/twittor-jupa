package routers

import (
	"encoding/json"
	"github.com/1hollow3/twittor-jupa/db"
	"github.com/1hollow3/twittor-jupa/models"
	"net/http"
)

func QueryRelation(w http.ResponseWriter, r *http.Request) {

	ID := r.URL.Query().Get("id")
	if len(ID) < 1 {
		http.Error(w, "id parameter is requested", 400)
		return
	}

	var t models.Relation
	t.RelatedUserID = ID
	t.UserID = IDUser

	var response models.ResponseQueryRelation

	status, err := db.ConsultRelation(t)
	if !status || err != nil {
		response.Status = false
	}else {
		response.Status = true
	}

	w.Header().Set("Content-type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(response)

}
