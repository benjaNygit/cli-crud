package models

import "gorm.io/gorm"

type TypeNote struct {
	gorm.Model
	Id          uint   `json:"id" gorm:"primaryKey"`
	Description string `json:"description" gorm:"type:varchar(50)"`
}
