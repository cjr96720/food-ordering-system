package repositories

import (
	"errors"

	"gorm.io/gorm"

	_ "food-ordering-system/internal/logger"
	"food-ordering-system/internal/models"
)

type UserRepository interface {
	Create(user models.User) error
	GetByEmail(email string) (models.User, error)
}

type UserRepositoryImpl struct {
	db *gorm.DB
}

func NewUserRepositoryImpl(db *gorm.DB) UserRepository {
	return &UserRepositoryImpl{db: db}
}

func (ur *UserRepositoryImpl) Create(user models.User) error {
	result := ur.db.Create(&user)

	return result.Error
}

func (ur *UserRepositoryImpl) GetByEmail(email string) (models.User, error) {
	var user models.User

	result := ur.db.Model(&models.User{}).Where("email = ?", email).First(&user)

	if result.Error != nil {
		return user, errors.New("user not found")
	} else {
		return user, nil
	}
}
