package main

import (
	"github.com/xuan-vy-nguyen/SE_Project01/database"
	"regexp"
	"time"
	"strings"
)

func checkingLogin(p database.LoginAccount) (database.JWTRespone, database.SignUpAccount, int) {
	var reponseJson database.JWTRespone
	var userInformation database.SignUpAccount
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
	reponseJson = database.JWTRespone{
		AccessToken: token,
		RefreshToken: "",
	}
	return reponseJson, userInformation, 2	// return ok
}

func checkingSignUp(infor database.SignUpAccount)(string) {
	reMail := regexp.MustCompile("^[a-zA-Z0-9.!#$%&'*+/=?^_`{|}~-]+@[a-zA-Z0-9](?:[a-zA-Z0-9-]{0,61}[a-zA-Z0-9])?(?:\\.[a-zA-Z0-9](?:[a-zA-Z0-9-]{0,61}[a-zA-Z0-9])?)*$")
	rePhone := regexp.MustCompile(`^(?:(?:\(?(?:00|\+)([1-4]\d\d|[1-9]\d?)\)?)?[\-\.\ \\\/]?)?((?:\(?\d{1,}\)?[\-\.\ \\\/]?){0,})(?:[\-\.\ \\\/]?(?:#|ext\.?|extension|x)[\-\.\ \\\/]?(\d+))?$`)
	reDate := regexp.MustCompile("((19|20)\\d\\d)/(0?[1-9]|1[012])/(0?[1-9]|[12][0-9]|3[01])")
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
	if !reDate.MatchString(infor.BirthDay) || (strings.Split(infor.BirthDay, ".")[0] >= strings.Split(time.Now().Format("2006/01/02"), ".")[0]) {
		return "birthday is wrong"
	}

	// check exist account in DB
	if existInSignUpDB(infor.Mail){
		return "email is used by another user"
	}
	// insert to mongoDB and return err, if err = "" => no err
	
	return ""
}

func checkingUpdateSignUp(infor database.SignUpAccount)(string) {
	reMail := regexp.MustCompile("^[a-zA-Z0-9.!#$%&'*+/=?^_`{|}~-]+@[a-zA-Z0-9](?:[a-zA-Z0-9-]{0,61}[a-zA-Z0-9])?(?:\\.[a-zA-Z0-9](?:[a-zA-Z0-9-]{0,61}[a-zA-Z0-9])?)*$")
	rePhone := regexp.MustCompile(`^(?:(?:\(?(?:00|\+)([1-4]\d\d|[1-9]\d?)\)?)?[\-\.\ \\\/]?)?((?:\(?\d{1,}\)?[\-\.\ \\\/]?){0,})(?:[\-\.\ \\\/]?(?:#|ext\.?|extension|x)[\-\.\ \\\/]?(\d+))?$`)
	reDate := regexp.MustCompile("((19|20)\\d\\d)/(0?[1-9]|1[012])/(0?[1-9]|[12][0-9]|3[01])")
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
	if !reDate.MatchString(infor.BirthDay) || (strings.Split(infor.BirthDay, ".")[0] >= strings.Split(time.Now().Format("2006/01/02"), ".")[0]) {
		return "birthday is wrong"
	}

	// check exist account in DB
	if existInSignUpDB(infor.Mail){
		return "email is used by another user"
	}
	// insert to mongoDB and return err, if err = "" => no err
	
	return ""
}