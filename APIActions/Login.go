package APIAction

import (
	"encoding/json"
	"net/http"
	"fmt"
	"time"
	"github.com/xuan-vy-nguyen/SE_Project01/DataStruct"
	"regexp"
	"strings"
)

func LoginPost(w http.ResponseWriter, r *http.Request) {
	fmt.Println("loginPost")

	var p DataStruct.LoginAccount

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
			Tokens	string `json:"access-tokens"`
			Users	string `json:"username"`
			BirthDay string 	`json:"birthday"`
			CreateAt string 	`json:"createat"`
			Phone	string	`json:"phoneNumber"`
		}
		responser := DataStruct.MessageRespone{
			Message: message,
			Body: bodyStruct{
				Tokens: JsonToken.AccessToken, 
				Users: UserInformation.UserName,
				BirthDay: UserInformation.BirthDay,
				Phone: UserInformation.PhoneNumber,
				CreateAt: UserInformation.CreateAt,
			},
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

func checkingLogin(p DataStruct.LoginAccount) (DataStruct.JWTRespone, DataStruct.SignUpAccount, int) {
	var reponseJson DataStruct.JWTRespone
	var userInformation DataStruct.SignUpAccount
	// check if user pass wrong
	switch (checkAccInSignUpDB(p)){
		case 0:	// wrong user - pass
			return reponseJson, userInformation, 0
		case 1:	// server bug
			return reponseJson, userInformation, 1
		case 4:	// do not have this account
			return reponseJson, userInformation, 4
			// case 2 is ok
	}

	// check if user has login in another place
	if checkAccInLoginDB(p) {
		return reponseJson, userInformation, 3
	}
	token, err := createJWT(p)
	if err {
		return reponseJson, userInformation, 1	// server bug
	}
	userInformation, err = getOneSignUpDB(p.Mail)
	if err {
		return reponseJson, userInformation, 1	// server bug
	}
	reponseJson = DataStruct.JWTRespone{
		AccessToken: token,
		RefreshToken: "",
	}
	return reponseJson, userInformation, 2	// return ok
}