package models

import (
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB

func ConnectDatabase() {
	database, err := gorm.Open(mysql.Open("root:root@tcp(mysql.default.svc.cluster.local:3306)/gordon"))
	if err != nil {
		panic(err)
	}

	database.AutoMigrate(&User{})
	database.AutoMigrate(&Todo{})

	DB = database
}
