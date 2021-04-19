package routers

import (
	"io"
	"net/http"
	"os"
	"strings"

	"github.com/MessiasJunio/twittor/db"
	"github.com/MessiasJunio/twittor/models"
)

func UploadBanner(w http.ResponseWriter, r *http.Request) {
	file, handler, err := r.FormFile("banner")
	var extension = strings.Split(handler.Filename, ".")[1]
	var avatarFile = "uploads/banners/" + UserID + "." + extension
	f, err := os.OpenFile(avatarFile, os.O_WRONLY|os.O_CREATE, 0666)
	if err != nil {
		http.Error(w, "Error while uploading the image "+err.Error(), http.StatusBadRequest)
		return
	}

	_, err = io.Copy(f, file)
	if err != nil {
		http.Error(w, "Error while copying the image "+err.Error(), http.StatusBadRequest)
		return
	}

	var user models.User
	var status bool

	user.Banner = UserID + "." + extension
	status, err = db.UpdateRegistry(user, UserID)
	if err != nil || status == false {
		http.Error(w, "Error while saving the banner in the Database"+err.Error(), http.StatusAccepted)
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
}
