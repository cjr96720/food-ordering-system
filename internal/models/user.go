package models

import "gorm.io/gorm"

type User struct {
	gorm.Model
	Username  string
	Password  string
	Email     string  `gorm:"unique"`
	Phone     string  `gorm:"unique;default:null"`
	Orders    []Order //User can create many Orders
	CreatedAt int
	UpdatedAt int
}
