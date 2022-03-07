package routers

import (
	"github.com/1hollow3/twittor-jupa/db"
	"io"
	"net/http"
	"os"
)

func GetBanner(w http.ResponseWriter, r *http.Request)  {

	ID := r.URL.Query().Get("id")
	if len(ID)<1{
		http.Error(w, "ID parameter is requested", http.StatusBadRequest)
		return
	}

	profile, err := db.SearchProfile(ID)
	if err != nil {
		http.Error(w, "User not found", http.StatusBadRequest)
		return
	}

	file, err := os.Open("uploads/banners/" + profile.Banner)
	if err != nil {
		http.Error(w, "Image not found", http.StatusBadRequest)
		return
	}

	_, err = io.Copy(w, file)
	if err != nil {
		http.Error(w, "Error copying the image", http.StatusInternalServerError)
		return
	}
}
