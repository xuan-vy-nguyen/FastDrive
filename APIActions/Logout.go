package APIActions

import (
	"encoding/json"
	"net/http"
	"fmt"
	"time"
	"github.com/xuan-vy-nguyen/SE_Project01/DataStruct"
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
