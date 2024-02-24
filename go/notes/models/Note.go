package models

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Note struct {
	gorm.Model
	Id          uuid.UUID `json:"id" gorm:"primaryKey;default:uuid_generate_v4()"`
	Title       string    `json:"title" gorm:"type:varchar(50)"`
	Description string    `json:"description" gorm:"type:text"`
	TypeNote    TypeNote  `json:"type_note" gorm:"foreignKey:Code"`
}
