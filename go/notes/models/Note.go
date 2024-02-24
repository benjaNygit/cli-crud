package models

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Note struct {
	gorm.Model
	Id          uuid.UUID `gorm:"primaryKey"`
	Title       string    `gorm:"type:varchar(50)"`
	Description string
	TypeNote    TypeNote `gorm:"foreignKey:Code"`
}
