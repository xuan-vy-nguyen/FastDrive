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

func SignUpPost(w http.ResponseWriter, r *http.Request) {
	fmt.Println("signUpPost")

	var p DataStruct.SignUpAccount
	var message string

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

func CheckingSignUp(infor DataStructSignUpAccount)(string) {
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
