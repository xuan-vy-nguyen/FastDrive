package main

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func main() {
	r := mux.NewRouter()
	api := r.PathPrefix("/api/v1").Subrouter()
	// api.HandleFunc("", get).Methods(http.MethodGet)
	api.HandleFunc("/login", loginPost).Methods(http.MethodPost)
	api.HandleFunc("/signup", signUpPost).Methods(http.MethodPost)
	// api.HandleFunc("", put).Methods(http.MethodPut)
	// api.HandleFunc("", delete).Methods(http.MethodDelete)
	// api.HandleFunc("", notFound)
	log.Fatal(http.ListenAndServe(":8080", r))
}
