package random

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/xuan-vy-nguyen/SE_Project01/datastruct"
	"github.com/xuan-vy-nguyen/SE_Project01/dbactions"
)

func GetRandomGet(w http.ResponseWriter, r *http.Request) {
	fmt.Println("getRandomGet")

	message, body := "", ""
	jwtStr := r.Header["Access-Token"][0]
	w.Header().Set("Content-Type", "application/json")
	defer func() {
		responser := datastruct.MessageRespone{
			Message: message,
			Body:    body,
		}
		json.NewEncoder(w).Encode(responser)
		fmt.Println("")
	}()

	// check information and return bugs
	if !dbactions.IsInLoginDB(jwtStr) {
		w.WriteHeader(http.StatusBadRequest)
		message = "your access-token is wrong"
		return
	}
	// if no bug -> return OK
	w.WriteHeader(http.StatusOK)
	message = "OK"
	body = "this is random-getting"
}
