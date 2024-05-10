package database

import (
	"github.com/Enileyeevvv/pharmacy-backend/pharmacy-service/models"
	"github.com/gofiber/fiber/v2/log"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var db *gorm.DB

func Init() *gorm.DB {
	// todo replace password
	dsn := "host=localhost user=postgres password=password dbname=pharmacy sslmode=disable"

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})

	if err != nil {
		log.Fatal(err)
	}

	db.AutoMigrate(&models.User{}, &models.MedicinalProduct{})

	return db
}

func GetDB() *gorm.DB {
	if db != nil {
		return db
	}

	db = Init()
	return db
}
