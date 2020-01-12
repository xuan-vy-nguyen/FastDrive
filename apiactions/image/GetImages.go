package image

import (
	"fmt"
	"net/http"

	"github.com/xuan-vy-nguyen/SE_Project01/datastruct"
	"github.com/xuan-vy-nguyen/SE_Project01/dbactions"
)

// GetImage is used for testing
func GetImage(w http.ResponseWriter, r *http.Request) {
	fmt.Println("GetImage")

	// check access token
	jwtStr := r.Header["Access-Token"][0]
	fileName := r.Header["File-Name"][0]
	_, err := dbactions.GetOneLoginDB(jwtStr)
	if err == true {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	var data datastruct.ImageDB

	// prepare for return
	w.Header().Set("Content-Type", "application/form-data")
	defer func() {
		w.Write(data.Image)
		fmt.Println("")
	}()

	// get Image and error
	data, err2 := dbactions.GetOneImageDB(fileName, jwtStr)
	if err2 == true {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	// return true
	w.WriteHeader(http.StatusOK)
}
