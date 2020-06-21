package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/gorilla/mux"
	apiactions "github.com/xuan-vy-nguyen/SE_Project01/services"
	account "github.com/xuan-vy-nguyen/SE_Project01/services/account"
	image "github.com/xuan-vy-nguyen/SE_Project01/services/image"
)

func main() {
	// Log PORT
	fmt.Println("PORT: ", os.Getenv("PORT"))
	// create API
	r := mux.NewRouter()
	api := r.PathPrefix("/api/v2").Subrouter()
	// "The Gate"
	api.HandleFunc("/logout", apiactions.LogOutGet).Methods(http.MethodGet)
	api.HandleFunc("/login", apiactions.LoginPost).Methods(http.MethodPost)
	api.HandleFunc("/signup", apiactions.SignUpPost).Methods(http.MethodPost)
	// Account's Behaviors
	api.HandleFunc("/accounts", account.GetAcountGet).Methods(http.MethodGet)
	api.HandleFunc("/accounts", account.UpdateAcountPut).Methods(http.MethodPut)
	api.HandleFunc("/accounts/password", account.ComparePasswordGet).Methods(http.MethodGet)
	// Image's Behaviors
	api.HandleFunc("/images", image.GetImage).Methods(http.MethodGet)
	api.HandleFunc("/images", image.PostImage).Methods(http.MethodPost)
	api.HandleFunc("/images/enhancements", image.EnhancementImage).Methods(http.MethodPost)
	api.HandleFunc("/images", image.DeleteImage).Methods(http.MethodDelete)
	api.HandleFunc("/images/lists", image.GetListImages).Methods(http.MethodGet)
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
