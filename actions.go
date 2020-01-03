package main

import (
	"encoding/json"
	"net/http"
	"fmt"
	"time"
	"github.com/xuan-vy-nguyen/SE_Project01/database"
)

func loginPost(w http.ResponseWriter, r *http.Request) {
	fmt.Println("loginPost")

	var p database.LoginAccount

	err := json.NewDecoder(r.Body).Decode(&p)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	var message string
	JsonToken, UserInformation, errr := checkingLogin(p)
	
	w.Header().Set("Content-Type", "application/json")
	defer func() {
		UserInformation.Pass = ""
		type bodyStruct struct {
			Tokens	interface{} `json:"tokens"`
			Users	interface{} `json:"user"`
		}
		responser := database.MessageRespone{
			Message: message,
			Body: bodyStruct{Tokens:JsonToken, Users : UserInformation},
		}
		json.NewEncoder(w).Encode(responser)
		fmt.Println("")
	}()

	switch(errr){
		case 0:
			w.WriteHeader(http.StatusBadRequest)
			message = "password is wrong"
			return
		case 1:
			w.WriteHeader(http.StatusInternalServerError)
			message = "server has something wrong"
			return
		case 3:
			w.WriteHeader(http.StatusBadRequest)
			message = "account is logging in another place"
			return
		case 4:
			w.WriteHeader(http.StatusBadRequest)
			message = "email is wrong"
			return
		default:
			errDB := addOneLoginDB(p.Mail, JsonToken.AccessToken)
			if errDB {	// if have a bug when add acc to LoginDB
				w.WriteHeader(http.StatusInternalServerError)
				message = "server has something wrong"
			} else {
				w.WriteHeader(http.StatusOK)
				message = "OK"
			}
			return
	}
}

func signUpPost(w http.ResponseWriter, r *http.Request) {
	fmt.Println("signUpPost")

	var p database.SignUpAccount
	var message string

	w.Header().Set("Content-Type", "application/json")
	defer func() {
		responser := database.MessageRespone{
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
		if errrStr := checkingSignUp(p); errrStr != "" {
			w.WriteHeader(http.StatusBadRequest)
			message = errrStr
			return 
		} 
		// update information
		p.CreateAt = time.Now().Format("2006/01/02")
		p.IsDeleted = false
		p.IsActive = false
		if errrStr := addOneSignUpDB(p); errrStr != "" {
			w.WriteHeader(http.StatusBadRequest)
			message = errrStr
			return 
		}
		w.WriteHeader(http.StatusCreated)
		message = "created"
		return
	}
}

func logOutGet(w http.ResponseWriter, r *http.Request) {
	fmt.Println("logOutGet")

	var message string
	jwtStr := r.Header["Access-Token"][0]
	w.Header().Set("Content-Type", "application/json")
	defer func() {
		responser := database.MessageRespone{
			Message: message,
			Body: nil,
		}
		json.NewEncoder(w).Encode(responser)
		fmt.Println("")
	}()

	// check information and return bugs 
	if (!isInLoginDB(jwtStr)) {	
		w.WriteHeader(http.StatusBadRequest)
		message = "your access-token is wrong"
		return 
	}

	// remove in login DB
	if err:= deleteOneLoginDB(jwtStr); err == true {
		w.WriteHeader(http.StatusInternalServerError)
		message = "Internal Server Error"
		return
	}

	// if no bug -> return OK
	w.WriteHeader(http.StatusOK)
	message = "OK"
}

func getRandomGet(w http.ResponseWriter, r *http.Request){
	fmt.Println("getRandomGet")

	message, body := "", ""
	jwtStr := r.Header["Access-Token"][0]
	w.Header().Set("Content-Type", "application/json")
	defer func() {
		responser := database.MessageRespone{
			Message: message,
			Body: body,
		}
		json.NewEncoder(w).Encode(responser)
		fmt.Println("")
	}()

	// check information and return bugs 
	if (!isInLoginDB(jwtStr)) {	
		w.WriteHeader(http.StatusBadRequest)
		message = "your access-token is wrong"
		return
	}
	// if no bug -> return OK
	w.WriteHeader(http.StatusOK)
	message = "OK"
	body = "this is random-getting"
}

func getAcountGet(w http.ResponseWriter, r *http.Request){
	fmt.Println("getAccountGet")

	message:= ""
	var body database.SignUpAccount
	jwtStr := r.Header["Access-Token"][0]
	w.Header().Set("Content-Type", "application/json")
	defer func() {
		responser := database.MessageRespone{
			Message: message,
			Body: body,
		}
		json.NewEncoder(w).Encode(responser)
		fmt.Println("")
	}()

	// check information and return bugs 
	acc, err := getOneLoginDB(jwtStr)
	if err == true {	
		w.WriteHeader(http.StatusBadRequest)
		message = "your access-token is wrong"
		return
	}

	// if no bug -> return OK
	if body, err = getOneSignUpDB(acc.Mail); err == true {
		w.WriteHeader(http.StatusInternalServerError)
		message = "server has something not true"
		return
	}
	w.WriteHeader(http.StatusOK)
	message = "OK"
}
func updateAcountPut(w http.ResponseWriter, r *http.Request){
	fmt.Println("updateAcountPut")

	var p database.SignUpAccount
	var message string
	jwtStr := r.Header["Access-Token"][0]

	w.Header().Set("Content-Type", "application/json")
	defer func() {
		responser := database.MessageRespone{
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
		newLogin := database.LoginDB{
			Mail: p.Mail, 
			Token: jwtStr,
		}
		// get old mail
		oldLogin, err := getOneLoginDB(jwtStr);
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
		if err := updateOneLoginDB(oldLogin.Mail, newLogin); err == true {
			w.WriteHeader(http.StatusInternalServerError)
			message = "something wrong"
			return
		}
		// Update SignUpDB
		if err := updateOneSignUpDB(oldLogin.Mail, p); err == true {
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