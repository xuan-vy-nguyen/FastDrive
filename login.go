package main

import (
	"github.com/xuan-vy-nguyen/SE_Project01/database"
	"regexp"
)

func checkingLogin(p database.LoginAccount) (database.JWTRespone, int) {
	var reponseJson database.JWTRespone
	// check if user pass wrong
	// if !checkAccInSignUpDB(p) {
	// 	return reponseJson, 0
	// }
	// // check if user has login in another place
	// if !checkAccInLoginDB(p) {
	// 	return reponseJson, 3
	// }
	token, err := createJWT(p)
	if err {
		return reponseJson, 1
	}
	reponseJson = database.JWTRespone{
		AccessToken: token,
		RefreshToken: "",
	}
	return reponseJson, 2
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
		return "email is used by an user"
	}
	// insert to mongoDB and return err, if err = "" => no err
	err := addSignUpDB(infor)
	return err
}
