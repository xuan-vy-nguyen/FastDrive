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
	token, errr := checkingLogin(p)
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
	w.Write([]byte(token))
}


func signUpPost(w http.ResponseWriter, r *http.Request) {
	var p database.LoginAccount

	err := json.NewDecoder(r.Body).Decode(&p)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	token, errr := checkingLogin(p)
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
	w.Write([]byte(token))
}