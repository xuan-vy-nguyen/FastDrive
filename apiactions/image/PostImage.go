package image

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/xuan-vy-nguyen/SE_Project01/dbactions"

	"github.com/xuan-vy-nguyen/SE_Project01/datastruct"
)

// PostImage is used for testing
func PostImage(w http.ResponseWriter, r *http.Request) {
	jwtStr := r.Header["Access-Token"][0]
	filename := r.Header["File-Name"][0]
	var message string

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

	// get image data
	file, _, err2 := r.FormFile("Image")
	if err2 != nil {
		w.WriteHeader(http.StatusBadRequest)
		message = err2.Error()
		return
	}

	// Read the file into memory
	data, err2 := ioutil.ReadAll(file)
	if err2 != nil {
		w.WriteHeader(http.StatusBadRequest)
		message = err2.Error()
		return
	}

	// insert to mongodb
	err3 := dbactions.AddOneImageDB(data, UserInfor.Mail, filename)
	if err3 {
		w.WriteHeader(http.StatusInternalServerError)
		message = "Cannot insert image to mongodb"
		return
	}

	// return ok
	w.WriteHeader(http.StatusCreated)
	message = "created"
}
