package services

import (
	"errors"
	"fmt"

	log "github.com/sirupsen/logrus"
	"golang.org/x/crypto/bcrypt"

	"food-ordering-system/internal/domains"
	_ "food-ordering-system/internal/logger"
	"food-ordering-system/internal/models"
	"food-ordering-system/internal/repositories"
)

type UserService interface {
	CreateUser(user domains.SignupRequest) error
	GetUserByEmail(email string) (models.User, error)
	Login(user domains.LoginRequest) error
}

type UserServiceImpl struct {
	UserRepository repositories.UserRepository
}

func NewUserServiceImpl(ur repositories.UserRepository) UserService {
	return &UserServiceImpl{
		UserRepository: ur,
	}
}

func (us *UserServiceImpl) CreateUser(user domains.SignupRequest) error {
	userModel := models.User{
		Username: user.Username,
		Password: user.Password,
		Email:    user.Email,
		Phone:    user.Phone,
	}

	return us.UserRepository.Create(userModel)
}

func (us *UserServiceImpl) GetUserByEmail(email string) (models.User, error) {
	return us.UserRepository.GetByEmail(email)
}

func (us *UserServiceImpl) Login(user domains.LoginRequest) error {
	userFromDb, err := us.UserRepository.GetByEmail(user.Email)

	if err != nil {
		log.Warn(fmt.Sprintf("error: %s", err))
		return errors.New("user not found")
	}

	// varified password
	isCorrectPassword := VarifyPassword(user.Password, userFromDb.Password)
	if !isCorrectPassword {
		return errors.New("incorrect password")
	}
	return nil
}

func VarifyPassword(userPassword string, correctPassword string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(correctPassword), []byte(userPassword))

	return err == nil
}
