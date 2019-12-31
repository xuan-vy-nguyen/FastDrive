package database

import (
	"github.com/dgrijalva/jwt-go"
)

// LoginAccount type of login data
type LoginAccount struct {
	Mail string `json:"mail"`
	Pass string `json:"pass"`
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