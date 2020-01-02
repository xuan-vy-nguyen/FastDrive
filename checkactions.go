package main

import (
	"github.com/xuan-vy-nguyen/SE_Project01/database"
	"regexp"
)

func checkingLogin(p database.LoginAccount) (database.JWTRespone, int) {
	var reponseJson database.JWTRespone
	// check if user pass wrong
	switch (checkAccInSignUpDB(p)){
		case 0:	// wrong user - pass
			return reponseJson, 0
		case 1:	// server bug
			return reponseJson, 1
		case 4:	// do not have this account
			return reponseJson, 4
			// case 2 is ok
	}

	// check if user has login in another place
	if checkAccInLoginDB(p) {
		return reponseJson, 3
	}
	token, err := createJWT(p)
	if err {
		return reponseJson, 1	// server bug
	}
	reponseJson = database.JWTRespone{
		AccessToken: token,
		RefreshToken: "",
	}
	return reponseJson, 2	// return ok
}

func checkingSignUp(infor database.SignUpAccount)(string) {
	reMail := regexp.MustCompile("^[a-zA-Z0-9.!#$%&'*+/=?^_`{|}~-]+@[a-zA-Z0-9](?:[a-zA-Z0-9-]{0,61}[a-zA-Z0-9])?(?:\\.[a-zA-Z0-9](?:[a-zA-Z0-9-]{0,61}[a-zA-Z0-9])?)*$")
	rePhone := regexp.MustCompile(`^(?:(?:\(?(?:00|\+)([1-4]\d\d|[1-9]\d?)\)?)?[\-\.\ \\\/]?)?((?:\(?\d{1,}\)?[\-\.\ \\\/]?){0,})(?:[\-\.\ \\\/]?(?:#|ext\.?|extension|x)[\-\.\ \\\/]?(\d+))?$`)
	// check infor is validated
	if !reMail.MatchString(infor.Mail) {
		return "email is wrong"
	}
	if len(infor.Pass) < 7 {
		return "password is atleast 7 characters"
	}
	if !rePhone.MatchString(infor.PhoneNumber){
		return "phone number is wrong"
	}
	// check exist account in DB
	if existInSignUpDB(infor.Mail){
		return "email is used by another user"
	}
	// insert to mongoDB and return err, if err = "" => no err
	err := addSignUpDB(infor)
	return err
}
