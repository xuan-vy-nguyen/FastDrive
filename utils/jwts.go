package utils

import (
	"time"

	"github.com/dgrijalva/jwt-go"
	datastruct "github.com/xuan-vy-nguyen/SE_Project01/models"
)

var jwtKey = []byte("my_secret_key")

// CreateJWT is ok
func CreateJWT(p datastruct.LoginAccount) (string, bool) { // return tokenstring, err
	expirationTime := time.Now().Add(100000 * time.Hour)
	// Create the JWT claims, which includes the username and expiry time
	claims := &datastruct.Claims{
		Mail: p.Mail,
		StandardClaims: jwt.StandardClaims{
			// In JWT, the expiry time is expressed as unix milliseconds
			ExpiresAt: expirationTime.Unix(),
		},
	}

	// Declare the token with the algorithm used for signing, and the claims
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	// Create the JWT string
	tokenString, err := token.SignedString(jwtKey)
	if err != nil {
		// If there is an error in creating the JWT return an internal server error
		return "", true
	}
	return tokenString, false
}
