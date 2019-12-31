package main

import (
	"encoding/json"
	"net/http"

	"github.com/xuan-vy-nguyen/SE_Project01/database"
)

func loginPost(w http.ResponseWriter, r *http.Request) {
	var p database.LoginAccount

	err := json.NewDecoder(r.Body).Decode(&p)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	JsonToken, errr := checkingLogin(p)
	switch(errr){
		case 0:
			w.WriteHeader(http.StatusBadRequest)
			w.Write([]byte(`{"message": "password is wrong"}`))
			return
		case 1:
			w.WriteHeader(http.StatusInternalServerError)
			w.Write([]byte(`{"message": "server has something wrong"}`))
			return
		case 3:
			w.WriteHeader(http.StatusBadRequest)
			w.Write([]byte(`{"message": "account is logging in another place"}`))
			return
		case 4:
			w.WriteHeader(http.StatusBadRequest)
			w.Write([]byte(`{"message": "email is wrong"}`))
			return
		default:
			errDB := addLoginDB(p.Mail, JsonToken.AccessToken)
			if errDB {	// if have a bug when add acc to LoginDB
				w.WriteHeader(http.StatusInternalServerError)
				w.Write([]byte(`{"message": "server has something wrong"}`))
				return
			}
	}
	// and send
	responser := database.MessageRespone{
		Message: "OK",
		Body: JsonToken,
	}
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(responser)
}


func signUpPost(w http.ResponseWriter, r *http.Request) {
	var p database.SignUpAccount

	err := json.NewDecoder(r.Body).Decode(&p)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	// check information and return bugs 
	errrStr := checkingSignUp(p)
	if errrStr != "" {
		responser := database.MessageRespone{
			Message: errrStr,
			Body: nil,
		}
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(responser)
		return 
	}
	// if no bug -> return OK
	w.WriteHeader(http.StatusCreated)
	w.Write([]byte(`{"message": "created"}`))
}