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
			w.Write([]byte(`{"message": "mail or|and password is wrong"}`))
			return
		case 1:
			w.WriteHeader(http.StatusInternalServerError)
			w.Write([]byte(`{"message": "server has something wrong"}`))
			return
		default:
			errDB := addLoginDB(p.Mail, JsonToken.AccessToken)
			if errDB{
				w.WriteHeader(http.StatusInternalServerError)
				w.Write([]byte(`{"message": "server has something wrong"}`))
				return
			}
	}
	// convert jsonToken to  and send
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(JsonToken)
}


func signUpPost(w http.ResponseWriter, r *http.Request) {
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
			w.Write([]byte(`{"message": "mail or/and password is wrong"}`))
			return
		case 1:
			w.WriteHeader(http.StatusInternalServerError)
			w.Write([]byte(`{"message": "server has something wrong"}`))
			return
	}
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(JsonToken) 
}