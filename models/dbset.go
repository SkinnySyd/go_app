package models

import (
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB

func ConnectDatabase() {
	database, err := gorm.Open(mysql.Open("root:root@tcp(10.106.117.167:3306)/gordon"))
	if err != nil {
		panic(err)
	}

	database.AutoMigrate(&User{})

	DB = database
}
