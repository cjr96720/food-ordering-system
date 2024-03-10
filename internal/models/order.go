package models

import (
	"time"

	"gorm.io/gorm"
)

type Order struct {
	gorm.Model
	UserID      int
	OrderDate   time.Time
	TotalAmount float64
	Status      string
	OrderItem   []OrderItem //Order consists of many OrderItems
	CreatedAt   int
	UpdatedAt   int
}
