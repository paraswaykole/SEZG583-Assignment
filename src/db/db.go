package db

import (
	"book-catalogue-service/src/constants"
	"fmt"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var db *gorm.DB

func InitDB() {
	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=disable TimeZone=Asia/Kolkata",
		constants.DB_HOST, constants.DB_USER, constants.DB_PASSWORD, constants.DB_NAME, constants.DB_PORT)
	var err error
	db, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		panic(err)
	}
}

func Get() *gorm.DB {
	return db
}
