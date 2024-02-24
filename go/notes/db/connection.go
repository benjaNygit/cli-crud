package db

import (
	"log"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

var db *gorm.DB

func DBConnection() {
	var err error
	db, err = gorm.Open(sqlite.Open("notes.db"), &gorm.Config{})

	if err != nil {
		log.Fatal(err)
	}
}
