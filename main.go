package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/gorilla/mux"
	"github.com/xuan-vy-nguyen/SE_Project01/apiactions"
	"github.com/xuan-vy-nguyen/SE_Project01/apiactions/account"
	"github.com/xuan-vy-nguyen/SE_Project01/apiactions/image"
)

func main() {
	x2()
	// Log PORT
	fmt.Println("PORT: ", os.Getenv("PORT"))
	// create API
	r := mux.NewRouter()
	api := r.PathPrefix("/api/v1").Subrouter()
	// "The Gate"
	api.HandleFunc("/logout", apiactions.LogOutGet).Methods(http.MethodGet)
	api.HandleFunc("/login", apiactions.LoginPost).Methods(http.MethodPost)
	api.HandleFunc("/signup", apiactions.SignUpPost).Methods(http.MethodPost)
	// Account's Behaviors
	api.HandleFunc("/account", account.GetAcountGet).Methods(http.MethodGet)
	api.HandleFunc("/account", account.UpdateAcountPut).Methods(http.MethodPut)
	api.HandleFunc("/account/password", account.ComparePasswordGet).Methods(http.MethodGet)
	// Image's Behaviors
	api.HandleFunc("/image", image.GetImage).Methods(http.MethodGet)
	api.HandleFunc("/image", image.PostImage).Methods(http.MethodPost)
	api.HandleFunc("/image/enhancement", image.EnhancementImage).Methods(http.MethodPost)
	api.HandleFunc("/image", image.DeleteImage).Methods(http.MethodDelete)
	api.HandleFunc("/image/list", image.GetListImages).Methods(http.MethodGet)
	// api.HandleFunc("", put).Methods(http.MethodPut)
	// api.HandleFunc("", delete).Methods(http.MethodDelete)
	// api.HandleFunc("", notFound)
	port := os.Getenv("PORT")
	print(port)
	if port != "" {
		log.Fatal(http.ListenAndServe(":"+port, r))
	} else {
		log.Fatal(http.ListenAndServe(":80", r))
	}
}
