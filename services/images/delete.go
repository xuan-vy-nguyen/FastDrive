package images

import (
	"encoding/json"
	"fmt"
	"net/http"

	dbactions "github.com/xuan-vy-nguyen/SE_Project01/controllers"
	datastruct "github.com/xuan-vy-nguyen/SE_Project01/models"
)

// DeleteImage is used for delete image
func DeleteImage(w http.ResponseWriter, r *http.Request) {
	fmt.Println("DeleteImage")

	// check access token
	jwtStr := r.Header["Accesstoken"][0]
	filename := r.Header["File-Name"][0]
	_, err := dbactions.GetOneLoginDB(jwtStr)
	if err == true {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	// return values
	var message string

	// prepare for return
	w.Header().Set("Content-Type", "application/form-data")

	defer func() {
		responser := datastruct.MessageRespone{
			Message: message,
			Body:    nil,
		}
		json.NewEncoder(w).Encode(responser)
		fmt.Println("")
	}()

	// find userMail
	UserInfor, err := dbactions.GetOneLoginDB(jwtStr)
	if err == true {
		w.WriteHeader(http.StatusBadRequest)
		message = "your access-token is wrong"
		return
	}

	// get list of user's images
	listNameImg, errListNameImg := dbactions.GetAllNameUserImage(UserInfor.Mail)
	if errListNameImg {
		w.WriteHeader(http.StatusInternalServerError)
		message = "StatusInternalServerError"
		return
	}

	// check name of new image
	i := 0
	for i < len(listNameImg) {
		if filename == listNameImg[i] {
			// delete Image and error
			if err = dbactions.DeleteOneImageDB(filename, UserInfor.Mail); err == true {
				w.WriteHeader(http.StatusInternalServerError)
				message = "StatusInternalServerError"
				return
			}
			// return true
			w.WriteHeader(http.StatusOK)
			message = "deleted"
			return
		}
		i++
	}
	w.WriteHeader(http.StatusNotFound)
	message = "This image is not in Database"
}
