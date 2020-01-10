package APIActions

import (
	"encoding/json"
	"net/http"
	"fmt"
	"time"
	"github.com/xuan-vy-nguyen/SE_Project01/DataStruct"
)

func GetRandomGet(w http.ResponseWriter, r *http.Request){
	fmt.Println("getRandomGet")

	message, body := "", ""
	jwtStr := r.Header["Access-Token"][0]
	w.Header().Set("Content-Type", "application/json")
	defer func() {
		responser := DataStruct.MessageRespone{
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
