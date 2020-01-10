package APIActions

import (
	"encoding/json"
	"net/http"
	"fmt"
	"time"
	"github.com/xuan-vy-nguyen/SE_Project01/DataStruct"
)

func updateAcountPut(w http.ResponseWriter, r *http.Request){
	fmt.Println("updateAcountPut")

	var p DataStruct.SignUpAccount
	var message string
	jwtStr := r.Header["Access-Token"][0]

	w.Header().Set("Content-Type", "application/json")
	defer func() {
		responser := DataStruct.MessageRespone{
			Message: message,
			Body: nil,
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
		newLogin := DataStruct.LoginDB{
			Mail: p.Mail, 
			Token: jwtStr,
		}
		// get old mail
		oldLogin, err := GetOneLoginDB(jwtStr);
		if err == true {
			w.WriteHeader(http.StatusInternalServerError)
			message = "something wrong"
			return
		}
		// checking 
		if errrStr := checkingSignUp(p); errrStr != ""  {
			if (errrStr == "email is used by another user" && p.Mail == oldLogin.Mail){
				fmt.Println("No conflict")
			} else {
				w.WriteHeader(http.StatusBadRequest)
				message = errrStr
				return 
			}
		} 
		// Update LoginDB
		if err := UpdateOneLoginDB(oldLogin.Mail, newLogin); err == true {
			w.WriteHeader(http.StatusInternalServerError)
			message = "something wrong"
			return
		}
		// Update SignUpDB
		if err := UpdateOneSignUpDB(oldLogin.Mail, p); err == true {
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