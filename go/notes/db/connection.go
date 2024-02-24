package db

import (
	"log"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

var DB *gorm.DB

func DBConnection() {
	var err error
	DB, err = gorm.Open(sqlite.Open("notes.sqlite"), &gorm.Config{})

	if err != nil {
		log.Fatal(err)
	}
}
