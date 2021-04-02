package database

import (
	"github.com/codergab/gojwt/models"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DBConn *gorm.DB

func Connect() {
	connection, err := gorm.Open(mysql.Open("root:@/go_basic_jwt"), &gorm.Config{})

	if err != nil {
		panic("Cannot Connect to DB")
	}

	DBConn = connection

	//Auto Migrate Users
	connection.AutoMigrate(&models.User{})
}
