package database

import (
	"fmt"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/schema"

	"food-ordering-system/bootstrap"
	"food-ordering-system/internal/models"
)

func NewPostgresDSN(env *bootstrap.Env) string {
	return fmt.Sprintf(
		"user=%s password=%s dbname=%s host=database port=%s sslmode=disable",
		env.PostgresUser,
		env.PostgresPassword,
		env.PostgresDB,
		env.PostgresPort,
	)
}

func NewPostgresConnection(env *bootstrap.Env) *gorm.DB {
	dsn := NewPostgresDSN(env)
	postgresConfig := postgres.Open(dsn)
	gormConfig := gorm.Config{NamingStrategy: schema.NamingStrategy{SingularTable: false}}

	db, err := gorm.Open(postgresConfig, &gormConfig)
	if err != nil {
		panic(err)
	}

	db.Table("users").AutoMigrate(&models.User{})
	db.Table("food").AutoMigrate(&models.Food{})
	db.Table("orders").AutoMigrate(&models.Order{})
	db.Table("order_items").AutoMigrate(&models.OrderItem{})

	return db
}
