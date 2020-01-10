package main

import (
	"log"
	"net/http"
	"os"
	"github.com/gorilla/mux"
	"github.com/xuan-vy-nguyen/SE_Project01/apiactions"
)


func main() {
	r := mux.NewRouter()
	api := r.PathPrefix("/api/v1").Subrouter()
	api.HandleFunc("/logout", apiactions.LogOutGet).Methods(http.MethodGet)
	api.HandleFunc("/random", apiactions.GetRandomGet).Methods(http.MethodGet)
	api.HandleFunc("/login", apiactions.LoginPost).Methods(http.MethodPost)
	api.HandleFunc("/signup", apiactions.SignUpPost).Methods(http.MethodPost)
	api.HandleFunc("/account", apiactions.GetAcountGet).Methods(http.MethodGet)
	api.HandleFunc("/account", apiactions.UpdateAcountPut).Methods(http.MethodPut)
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
