package models

import "gorm.io/gorm"

type Food struct {
	gorm.Model
	FoodName    string
	Category    string
	Description string
	Price       float32
	Available   bool
	CreatedAt   int
	UpdatedAt   int
}
