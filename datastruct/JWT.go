package datastruct

import (
	"github.com/dgrijalva/jwt-go"
)

// Claims is used for JWT
type Claims struct {
	Mail string `json:"mail"`
	jwt.StandardClaims
}

// JWTRespone is ok
type JWTRespone struct {
	AccessToken  string `json:"accesstoken"`
	RefreshToken string `json:"refreshtoken"`
}
