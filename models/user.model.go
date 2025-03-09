package models

import (
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	Username string  `gorm:"not null;unique" json:"username"`
	Name     string  `gorm:"type:varchar(100);not null" json:"name"`
	Email    *string `gorm:"unique" json:"email"`
	Password string  `gorm:"not null" json:"-"`
}
