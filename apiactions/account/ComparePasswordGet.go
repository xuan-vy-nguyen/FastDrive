package account

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/xuan-vy-nguyen/SE_Project01/datastruct"
	"github.com/xuan-vy-nguyen/SE_Project01/dbactions"
)

func ComparePasswordGet(w http.ResponseWriter, r *http.Request) {
	fmt.Println("ComparePasswordGet")

	message := ""
	jwtStr := r.Header["Access-Token"][0]

	w.Header().Set("Content-Type", "application/json")
	defer func() {
		responser := datastruct.MessageRespone{
			Message: message,
			Body:    nil,
		}
		json.NewEncoder(w).Encode(responser)
		fmt.Println("")
	}()

	// get password in body
	var p struct {
		Password string `json:"pass"`
	}
	errGetReq := json.NewDecoder(r.Body).Decode(&p)
	if errGetReq != nil {
		w.WriteHeader(http.StatusBadRequest)
		message = errGetReq.Error()
		return
	}

	// check information and return bugs
	acc, err := dbactions.GetOneLoginDB(jwtStr)
	if err == true {
		w.WriteHeader(http.StatusBadRequest)
		message = "your access-token is wrong"
		return
	}

	// if no bug -> return OK
	userAccount, err := dbactions.GetOneSignUpDB(acc.Mail)
	if err == true {
		w.WriteHeader(http.StatusInternalServerError)
		message = "server has something not true"
		return
	}

	// if request.pass != userAccount.Pass -> 400
	if userAccount.Pass != p.Password {
		w.WriteHeader(http.StatusBadRequest)
		message = "your password is wrong"
		return
	}
	// if no bug -> return ok
	w.WriteHeader(http.StatusOK)
	message = "your password is OK"
}
