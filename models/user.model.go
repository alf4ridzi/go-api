package models

import (
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	Username string  `gorm:"not null;unique"`
	Name     string  `gorm:"type:varchar(100);not null"`
	Email    *string `gorm:"unique"`
	Password string  `gorm:"not null"`
}
