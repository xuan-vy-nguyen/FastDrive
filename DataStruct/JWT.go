package DataStruct

import (
	"github.com/dgrijalva/jwt-go"
)

// use for JWT
type Claims struct {
	Mail string `json:"mail"`
	jwt.StandardClaims
}

type JWTRespone struct {
	AccessToken string `json:"accesstoken"`
	RefreshToken string `json:"refreshtoken"`
}