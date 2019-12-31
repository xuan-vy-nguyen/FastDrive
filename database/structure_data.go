package database

import (
	"github.com/dgrijalva/jwt-go"
)
type MessageRespone struct {
	Message string `json:"message"`
	Body interface{} `json:"body"`
}
// LoginAccount type of login data
type LoginAccount struct {
	Mail string `json:"mail"`
	Pass string `json:"pass"`
}

// SignUpAccount type of signup
type SignUpAccount struct {
	Mail string `json:"mail"`
	Pass string `json:"pass"`
	Name string `json:"name"`
	PhoneNumber string `json:"phoneNumber"`
}

// use for JWT
type Claims struct {
	Mail string `json:"mail"`
	jwt.StandardClaims
}

type JWTRespone struct {
	AccessToken string `json:"accesstoken"`
	RefreshToken string `json:"refreshtoken"`
}

// use for DB
type LoginDB struct {
	Mail string `json:"mail"`
	Token string `json:"token"`
}