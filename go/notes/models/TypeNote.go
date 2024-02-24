package models

import "gorm.io/gorm"

type TypeNote struct {
	gorm.Model
	Code        uint   `gorm:"primaryKey"`
	Description string `gorm:"unique"`
}
