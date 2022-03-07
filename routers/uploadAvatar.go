package routers

import (
	"github.com/1hollow3/twittor-jupa/db"
	"github.com/1hollow3/twittor-jupa/models"
	"io"
	"net/http"
	"os"
	"strings"
)

func UploadAvatar(w http.ResponseWriter, r *http.Request)  {

	file, handler, err := r.FormFile("avatar")
	var extension = strings.Split(handler.Filename,".")[1]
	var fileRoute = "uploads/avatars/" + IDUser + "." + extension
	var f *os.File

	f, err = os.OpenFile(fileRoute, os.O_WRONLY|os.O_CREATE, 0666)
	if err != nil {
		http.Error(w, "error uploading the avatar " + err.Error(), 500)
		return
	}

	_, err = io.Copy(f, file)
	if err != nil {
		http.Error(w, "error copying the avatar " + err.Error(), 500)
		return
	}

	var user models.User
	var status bool

	user.Avatar = IDUser + "." + extension
	status, err = db.ModifyUser(user, IDUser)
	if err != nil || status == false{
		http.Error(w, "error recording the avatar " + err.Error(), 500)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusAccepted)
}