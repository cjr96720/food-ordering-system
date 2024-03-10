package models

import "gorm.io/gorm"

type OrderItem struct {
	gorm.Model
	OrderId   int
	FoodId    int
	Quantity  int
	Subtotal  float64
	CreatedAt int
	UpdatedAt int
}
