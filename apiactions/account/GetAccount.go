package account

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/xuan-vy-nguyen/SE_Project01/datastruct"
	"github.com/xuan-vy-nguyen/SE_Project01/dbactions"
)

func GetAcountGet(w http.ResponseWriter, r *http.Request) {
	fmt.Println("getAccountGet")

	message := ""
	var body datastruct.SignUpAccount
	jwtStr := r.Header["AccessToken"][0]

	w.Header().Set("Content-Type", "application/json")
	defer func() {
		body.Pass = ""
		responser := datastruct.MessageRespone{
			Message: message,
			Body:    body,
		}
		json.NewEncoder(w).Encode(responser)
		fmt.Println("")
	}()

	// check information and return bugs
	acc, err := dbactions.GetOneLoginDB(jwtStr)
	if err == true {
		w.WriteHeader(http.StatusBadRequest)
		message = "your access-token is wrong"
		return
	}

	// if no bug -> return OK
	if body, err = dbactions.GetOneSignUpDB(acc.Mail); err == true {
		w.WriteHeader(http.StatusInternalServerError)
		message = "server has something not true"
		return
	}

	// if no bug -> return ok
	w.WriteHeader(http.StatusOK)
	message = "OK"
}
