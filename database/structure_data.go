package database

import (
	"github.com/dgrijalva/jwt-go"
)

// LoginAccount type of login data
type LoginAccount struct {
	Mail string `json:"mail"`
	Pass string `json:"pass"`
}

type Claims struct {
	Mail string `json:"mail"`
	jwt.StandardClaims
}
