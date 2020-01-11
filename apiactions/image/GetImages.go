package image

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/xuan-vy-nguyen/SE_Project01/datastruct"
)

// GetImage is used for testing
func GetImage(w http.ResponseWriter, r *http.Request) {
	fmt.Println("GetImage")

	var p datastruct.ImageRequest
	var message string
	var body datastruct.ImageResponse

	err := json.NewDecoder(r.Body).Decode(&p)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		message = err.Error()
		return
	}

	w.Header().Set("Content-Type", "application/json")
	defer func() {
		responser := datastruct.MessageRespone{
			Message: message,
			Body:    body,
		}
		json.NewEncoder(w).Encode(responser)
		fmt.Println("")
	}()

}
