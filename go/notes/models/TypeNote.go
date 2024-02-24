package models

import "gorm.io/gorm"

type TypeNote struct {
	gorm.Model
	Code        uint   `json:"code" gorm:"primaryKey"`
	Description string `json:"description" gorm:"unique"`
}
