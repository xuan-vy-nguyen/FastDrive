package restapi

import (
	"SE_Project01/server/database"
	"encoding/json"
	"fmt"
	"net/http"
)

// loginPost just login
func loginPost(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	w.Write([]byte(`{"message": "received"}`))

	var p database.LoginAccount

	// Try to decode the request body into the struct. If there is an error,
	// respond to the client with the error message and a 400 status code.
	err := json.NewDecoder(r.Body).Decode(&p)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	// Do something with the Person struct...
	fmt.Fprintf(w, "Account: %+v", p)
}
