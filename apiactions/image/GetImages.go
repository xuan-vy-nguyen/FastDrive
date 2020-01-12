package image

import (
	"fmt"
	"net/http"

	"github.com/xuan-vy-nguyen/SE_Project01/datastruct"
	"github.com/xuan-vy-nguyen/SE_Project01/dbactions"
)

// GetImage is used for testing
func GetImage(w http.ResponseWriter, r *http.Request) {
	fmt.Println("GetImage")

	// check access token
	jwtStr := r.Header["Access-Token"][0]
	filename := r.Header["File-Name"][0]
	// checking jwt
	_, err := dbactions.GetOneLoginDB(jwtStr)
	if err == true {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	var data datastruct.ImageDB

	// prepare for return
	w.Header().Set("Content-Type", "multipart/form-data")
	defer func() {
		w.Write(data.Image)
		fmt.Println("")
	}()

	// find userMail
	UserInfor, err := dbactions.GetOneLoginDB(jwtStr)
	if err == true {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	// get list of user's images
	listNameImg, errListNameImg := dbactions.GetAllNameUserImage(UserInfor.Mail)
	if errListNameImg {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	// check name of new image
	i := 0
	for i < len(listNameImg) {
		if filename == listNameImg[i] {
			// get Image and error
			data, err = dbactions.GetOneImageDB(filename, UserInfor.Mail)
			if err == true {
				w.WriteHeader(http.StatusInternalServerError)
				return
			}
			// return true
			w.WriteHeader(http.StatusOK)
			return
		}
		i++
	}
	w.WriteHeader(http.StatusNotFound)
}
