package database

import (
	"github.com/Enileyeevvv/pharmacy-backend/pharmacy-service/models"
)

func IsUserCreatedByLogin(login string) bool {
	var count int64 = 0
	var user models.User

	GetDB().First(
		&user,
		models.User{Login: login}).Count(&count)

	return count == 1 && user.ID > 0
}

func CreateUser(user *models.User) error {
	result := GetDB().Create(&user)

	return result.Error
}
