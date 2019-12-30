package actions

import (
	"github.com/xuan-vy-nguyen/SE_Project01/database"
)

func checkingLogin(p database.LoginAccount) bool {
	if p.Mail != "xuanvy99" && p.Pass != "12345678" {
		return false
	}
	return true
}
