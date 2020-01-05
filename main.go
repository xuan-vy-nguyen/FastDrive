package main

import (
	"log"
	"net/http"
	"os"
	"github.com/gorilla/mux"
	"github.com/xuan-vy-nguyen.SE_Project01/test"
)

var MongoURI = //"mongodb+srv://xuanvyClone1:az1731999@cluster0-ktqay.mongodb.net/test?retryWrites=true&w=majority"
	"mongodb+srv://xuanvy99:az1731999@cluster0-mzeio.mongodb.net/test?retryWrites=true&w=majority" 
	// "mongodb://localhost:27017"
var Collection = "app"
var LoginDB = "LoginDB"
var SignDB = "SignDB"

func main() {
	test.testing()
	r := mux.NewRouter()
	api := r.PathPrefix("/api/v1").Subrouter()
	api.HandleFunc("/logout", logOutGet).Methods(http.MethodGet)
	api.HandleFunc("/random", getRandomGet).Methods(http.MethodGet)
	api.HandleFunc("/login", loginPost).Methods(http.MethodPost)
	api.HandleFunc("/signup", signUpPost).Methods(http.MethodPost)
	api.HandleFunc("/account", getAcountGet).Methods(http.MethodGet)
	api.HandleFunc("/account", updateAcountPut).Methods(http.MethodPut)
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
