package main

import (
	"log"
	"net/http"
	"os"
	"github.com/gorilla/mux"
	"github.com/xuan-vy-nguyen/SE_Project01/APIActions"
)


func main() {
	r := mux.NewRouter()
	api := r.PathPrefix("/api/v1").Subrouter()
	api.HandleFunc("/logout", APIActions.LogOutGet).Methods(http.MethodGet)
	api.HandleFunc("/random", APIActions.GetRandomGet).Methods(http.MethodGet)
	api.HandleFunc("/login", APIActions.LoginPost).Methods(http.MethodPost)
	api.HandleFunc("/signup", APIActions.SignUpPost).Methods(http.MethodPost)
	api.HandleFunc("/account", APIActions.GetAcountGet).Methods(http.MethodGet)
	api.HandleFunc("/account", APIActions.UpdateAcountPut).Methods(http.MethodPut)
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
