package account

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/xuan-vy-nguyen/SE_Project01/apiactions"
	"github.com/xuan-vy-nguyen/SE_Project01/datastruct"
	"github.com/xuan-vy-nguyen/SE_Project01/dbactions"
)

func UpdateAcountPut(w http.ResponseWriter, r *http.Request) {
	fmt.Println("updateAcountPut")

	var p datastruct.SignUpAccount
	var message string
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

	err := json.NewDecoder(r.Body).Decode(&p)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		message = err.Error()
		return
	} else {
		// Update LoginDB
		newLogin := datastruct.LoginDB{
			Mail:  p.Mail,
			Token: jwtStr,
		}
		// get old mail
		oldLogin, err := dbactions.GetOneLoginDB(jwtStr)
		if err == true {
			w.WriteHeader(http.StatusInternalServerError)
			message = "something wrong"
			return
		}
		// checking
		if errrStr := apiactions.CheckingSignUp(p); errrStr != "" {
			if errrStr == "email is used by another user" && p.Mail == oldLogin.Mail {
				fmt.Println("No conflict")
			} else {
				w.WriteHeader(http.StatusBadRequest)
				message = errrStr
				return
			}
		}
		// Update LoginDB
		if err := dbactions.UpdateOneLoginDB(oldLogin.Mail, newLogin); err == true {
			w.WriteHeader(http.StatusInternalServerError)
			message = "something wrong"
			return
		}
		// Update SignUpDB
		if err := dbactions.UpdateOneSignUpDB(oldLogin.Mail, p); err == true {
			w.WriteHeader(http.StatusInternalServerError)
			message = "something wrong"
			return
		}
		// return OK
		w.WriteHeader(http.StatusCreated)
		message = "created"
		return
	}
}
