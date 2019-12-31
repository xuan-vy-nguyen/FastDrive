package main

import (
	"github.com/xuan-vy-nguyen/SE_Project01/database"
)

func checkingLogin(p database.LoginAccount) (database.JWTRespone, int) {
	var reponseJson database.JWTRespone
	if p.Mail != "xuanvy99" || p.Pass != "12345678" {
		return reponseJson, 0
	}
	token, err := createJWT(p)
	if err {
		return reponseJson, 1
	}
	reponseJson = database.JWTRespone{
		AccessToken: token,
		RefreshToken: "",
	}
	return reponseJson, 2
}

