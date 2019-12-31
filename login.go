package main

import (
	"github.com/xuan-vy-nguyen/SE_Project01/database"
)

func checkingLogin(p database.LoginAccount) (string, int) {
	if p.Mail != "xuanvy99" || p.Pass != "12345678" {
		return "", 0
	}
	token, err := createJWT(p)
	if err {
		return "", 1
	}
	return token, 2
}

