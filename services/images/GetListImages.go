package image

import (
	"encoding/json"
	"fmt"
	"net/http"

	datastruct "github.com/xuan-vy-nguyen/SE_Project01/models"
	dbactions "github.com/xuan-vy-nguyen/SE_Project01/controllers"
)

// GetListImages is used for testing
func GetListImages(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Get All Name Image")

	// check access token
	jwtStr := r.Header["Accesstoken"][0]

	// return values
	var body []string
	var message string

	// prepare for return
	w.Header().Set("Content-Type", "application/form-data")

	defer func() {
		responser := datastruct.MessageRespone{
			Message: message,
			Body:    body,
		}
		json.NewEncoder(w).Encode(responser)
		fmt.Println("")
	}()

	// checking JWT
	_, err := dbactions.GetOneLoginDB(jwtStr)
	if err == true {
		w.WriteHeader(http.StatusBadRequest)
		message = "your access-token is wrong"
		return
	}

	// find userMail
	UserInfor, err := dbactions.GetOneLoginDB(jwtStr)
	if err == true {
		w.WriteHeader(http.StatusInternalServerError)
		message = "StatusInternalServerError"
		return
	}

	// get list of user's images
	body, errListNameImg := dbactions.GetAllNameUserImage(UserInfor.Mail)
	if errListNameImg {
		w.WriteHeader(http.StatusInternalServerError)
		message = "StatusInternalServerError"
		return
	}

	// return ok
	w.WriteHeader(http.StatusOK)
	message = "OK"
}
