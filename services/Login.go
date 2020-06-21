package apiactions

import (
	"encoding/json"
	"fmt"
	"net/http"

	dbactions "github.com/xuan-vy-nguyen/SE_Project01/controllers"
	datastruct "github.com/xuan-vy-nguyen/SE_Project01/models"
	utils "github.com/xuan-vy-nguyen/SE_Project01/utils"
)

func LoginPost(w http.ResponseWriter, r *http.Request) {
	fmt.Println("loginPost")

	var p datastruct.LoginAccount
	var message string

	err := json.NewDecoder(r.Body).Decode(&p)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		message = err.Error()
		return
	}

	jsonToken, UserInformation, errr := CheckingLogin(p)

	w.Header().Set("Content-Type", "application/json")
	defer func() {
		type bodyStruct struct {
			Tokens   string `json:"Accesstokens"`
			Users    string `json:"username"`
			BirthDay string `json:"birthday"`
			CreateAt string `json:"createat"`
			Phone    string `json:"phoneNumber"`
		}
		responser := datastruct.MessageRespone{
			Message: message,
			Body: bodyStruct{
				Tokens:   jsonToken.Accesstoken,
				Users:    UserInformation.UserName,
				BirthDay: UserInformation.BirthDay,
				Phone:    UserInformation.PhoneNumber,
				CreateAt: UserInformation.CreateAt,
			},
		}
		json.NewEncoder(w).Encode(responser)
		fmt.Println("")
	}()

	switch errr {
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
		message = "mail is wrong"
		return
	default:
		errDB := dbactions.AddOneLoginDB(p.Mail, jsonToken.Accesstoken)
		if errDB { // if have a bug when add acc to LoginDB
			w.WriteHeader(http.StatusInternalServerError)
			message = "server has something wrong"
		} else {
			w.WriteHeader(http.StatusOK)
			message = "OK"
		}
		return
	}
}

func CheckingLogin(p datastruct.LoginAccount) (datastruct.JWTRespone, datastruct.SignUpAccount, int) {
	var reponseJson datastruct.JWTRespone
	var userInformation datastruct.SignUpAccount
	// check if user pass wrong
	switch dbactions.CheckAccInSignUpDB(p) {
	case 0: // wrong user - pass
		return reponseJson, userInformation, 0
	case 1: // server bug
		return reponseJson, userInformation, 1
	case 4: // do not have this account
		return reponseJson, userInformation, 4
		// case 2 is ok
	}

	// check if user has login in another place
	if dbactions.CheckAccInLoginDB(p) {
		return reponseJson, userInformation, 3
	}
	token, err := utils.CreateJWT(p)
	if err {
		return reponseJson, userInformation, 1 // server bug
	}
	userInformation, err = dbactions.GetOneSignUpDB(p.Mail)
	if err {
		return reponseJson, userInformation, 1 // server bug
	}
	reponseJson = datastruct.JWTRespone{
		Accesstoken:  token,
		RefreshToken: "",
	}
	return reponseJson, userInformation, 2 // return ok
}
