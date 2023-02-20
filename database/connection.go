package database

import (
	"pro/models"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB

func Connect() {
	connection, err := gorm.Open(mysql.Open("root:@/project"), &gorm.Config{})

	if err != nil {
		panic("Could not connect to database")
	}
	connection.AutoMigrate(&models.User{})

	DB = connection
}
