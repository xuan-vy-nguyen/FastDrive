package main

import (
	"log"
	"net/http"
	"os"
	"github.com/gorilla/mux"
)

func main() {
	r := mux.NewRouter()
	api := r.PathPrefix("/api/v1").Subrouter()
	api.HandleFunc("/logout", logOutGet).Methods(http.MethodGet)
	api.HandleFunc("/get", getRandomGet).Methods(http.MethodGet)
	api.HandleFunc("/login", loginPost).Methods(http.MethodPost)
	api.HandleFunc("/signup", signUpPost).Methods(http.MethodPost)
	// api.HandleFunc("", put).Methods(http.MethodPut)
	// api.HandleFunc("", delete).Methods(http.MethodDelete)
	// api.HandleFunc("", notFound)
	port := os.Getenv("PORT")
	print(port)
	if (port != "") {
		log.Fatal(http.ListenAndServe(":" + port, r))
	} else {
		log.Fatal(http.ListenAndServe(":8080", r))
	} 
}
