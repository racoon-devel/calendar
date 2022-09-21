package model

import "gorm.io/gorm"

// User описывает данные пользователя, записанные при регистрации
type User struct {
	gorm.Model
	Login        string `gorm:"unique, not null"`
	PasswordHash string `gorm:"not null"`
	Name         string `gorm:"not null"`
	Surname      string `gorm:"not null"`
}
