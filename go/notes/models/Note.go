package models

import (
	"gorm.io/gorm"
)

type Note struct {
	gorm.Model
	Id          uint     `json:"id" gorm:"primaryKey"`
	Title       string   `json:"title" gorm:"type:varchar(50)"`
	Description string   `json:"description" gorm:"type:text"`
	TypeNote    TypeNote `json:"type_note" gorm:"foreignKey:Id"`
}
