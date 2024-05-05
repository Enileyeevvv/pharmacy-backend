package database

import (
	"backend/app/models"
)

func IsUserSignedUpByLogin(login string) bool {
	var count int64 = 0
	var user models.User
	GetDB().First(&user, models.User{Login: login}).Count(&count)
	return count == 1 && user.ID > 0
}
