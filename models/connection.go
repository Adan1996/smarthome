package models

import (
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB

func ConnectionDatabase() {
	database, err := gorm.Open(mysql.Open("root:toor@tcp(localhost:3306)/smarthome"))
	if err != nil {
		panic("failed to connect database")
	}

	database.AutoMigrate(&Rfid{})

	DB = database
}
