package apiactions

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/xuan-vy-nguyen/SE_Project01/datastruct"
	"github.com/xuan-vy-nguyen/SE_Project01/dbactions"
)

func LogOutGet(w http.ResponseWriter, r *http.Request) {
	fmt.Println("logOutGet")

	var message string
	jwtStr := r.Header["AccessToken"][0]
	w.Header().Set("Content-Type", "application/json")
	defer func() {
		responser := datastruct.MessageRespone{
			Message: message,
			Body:    nil,
		}
		json.NewEncoder(w).Encode(responser)
		fmt.Println("")
	}()

	// check information and return bugs
	if !dbactions.IsInLoginDB(jwtStr) {
		w.WriteHeader(http.StatusBadRequest)
		message = "your access-token is wrong"
		return
	}

	// remove in login DB
	if err := dbactions.DeleteOneLoginDB(jwtStr); err == true {
		w.WriteHeader(http.StatusInternalServerError)
		message = "Internal Server Error"
		return
	}

	// if no bug -> return OK
	w.WriteHeader(http.StatusOK)
	message = "OK"
}
