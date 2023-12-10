package models

import "gorm.io/gorm"

// User datos del usuario
type User struct {
	gorm.Model
	Name  string `gorm:"not null" json:"name"`
	Email string `gorm:"not null" json:"email"`
}

// Users lista de usuarios
type Users []*User
