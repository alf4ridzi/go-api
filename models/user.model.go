package models

import (
	"gorm.io/gorm"
)

type userRole string

const (
	ADMIN userRole = "admin"
	USER  userRole = "user"
)

type User struct {
	gorm.Model
	Username string   `gorm:"not null;unique" json:"username"`
	Name     string   `gorm:"type:varchar(100);not null" json:"name"`
	Email    *string  `gorm:"unique;type:varchar(255)" json:"email"`
	Role     userRole `gorm:"type:varchar(10);not null;default:'user'" json:"role"`
	Password string   `gorm:"not null" json:"-"`
}
