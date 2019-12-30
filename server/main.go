package server

import (
	_ "SE_Project01/server/restapi"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

// func get(w http.ResponseWriter, r *http.Request) {
// 	w.Header().Set("Content-Type", "application/json")
// 	w.WriteHeader(http.StatusOK)
// 	w.Write([]byte(`{"message": "get called"}`))
// }

// func post(w http.ResponseWriter, r *http.Request) {
// 	w.Header().Set("Content-Type", "application/json")
// 	w.WriteHeader(http.StatusCreated)
// 	w.Write([]byte(`{"message": "post called"}`))

// // 	var p Person

// // 	// Try to decode the request body into the struct. If there is an error,
// // 	// respond to the client with the error message and a 400 status code.
// // 	err := json.NewDecoder(r.Body).Decode(&p)
// // 	if err != nil {
// // 		http.Error(w, err.Error(), http.StatusBadRequest)
// // 		return
// // 	}

// // 	// Do something with the Person struct...
// // 	fmt.Fprintf(w, "Person: %+v", p)
// // }

// func put(w http.ResponseWriter, r *http.Request) {
// 	w.Header().Set("Content-Type", "application/json")
// 	w.WriteHeader(http.StatusAccepted)
// 	w.Write([]byte(`{"message": "put called"}`))
// }

// func delete(w http.ResponseWriter, r *http.Request) {
// 	w.Header().Set("Content-Type", "application/json")
// 	w.WriteHeader(http.StatusOK)
// 	w.Write([]byte(`{"message": "delete called"}`))
// }

// func notFound(w http.ResponseWriter, r *http.Request) {
// 	w.Header().Set("Content-Type", "application/json")
// 	w.WriteHeader(http.StatusNotFound)
// 	w.Write([]byte(`{"message": "not found"}`))
// }

func main() {
	r := mux.NewRouter()
	api := r.PathPrefix("/api/v1").Subrouter()
	// api.HandleFunc("", get).Methods(http.MethodGet)
	api.HandleFunc("login", restapi.loginPost).Methods(http.MethodPost)
	// api.HandleFunc("", put).Methods(http.MethodPut)
	// api.HandleFunc("", delete).Methods(http.MethodDelete)
	// api.HandleFunc("", notFound)
	log.Fatal(http.ListenAndServe(":8080", r))
}
