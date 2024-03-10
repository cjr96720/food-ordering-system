package controllers

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	log "github.com/sirupsen/logrus"
	"golang.org/x/crypto/bcrypt"

	"food-ordering-system/internal/domains"
	_ "food-ordering-system/internal/logger"
	"food-ordering-system/internal/services"
)

type UserController struct {
	userService services.UserService
}

func NewUserController(us services.UserService) *UserController {
	return &UserController{userService: us}
}

func (uc *UserController) Signup(c *gin.Context) {
	var user domains.SignupRequest

	err := c.ShouldBindJSON(&user)
	if err != nil {
		log.Error("Error: " + err.Error())
		c.JSON(
			http.StatusBadRequest,
			domains.ErrorResponse{Error: "Please check your request body."},
		)
		return
	}

	//check if the user exists
	_, err = uc.userService.GetUserByEmail(user.Email)
	if err == nil {
		c.JSON(
			http.StatusConflict,
			domains.ErrorResponse{Error: fmt.Sprintf("User already exists with the given email [%s].", user.Email)},
		)
		return
	}

	//encrypt password
	encryptedPassword, err := bcrypt.GenerateFromPassword(
		[]byte(user.Password),
		bcrypt.DefaultCost,
	)
	if err != nil {
		log.Error("Error: " + err.Error())
		c.JSON(
			http.StatusBadRequest,
			domains.ErrorResponse{Error: "Backend Error."},
		)
		return
	}
	user.Password = string(encryptedPassword)

	err = uc.userService.CreateUser(user)
	if err != nil {
		log.Error("Error: " + err.Error())
		c.JSON(
			http.StatusBadRequest,
			domains.ErrorResponse{Error: "Backend Error."},
		)
		return
	}

	response := domains.DefaultResponse{
		Message: fmt.Sprintf("New user %s is created.", user.Username),
		Data:    nil,
	}

	c.JSON(http.StatusOK, response)
}

func (uc *UserController) Login(c *gin.Context) {
	var loginInfo domains.LoginRequest

	err := c.ShouldBindJSON(&loginInfo)
	if err != nil {
		log.Error("Error: " + err.Error())
		c.JSON(
			http.StatusBadRequest,
			domains.ErrorResponse{Error: "Please check your request body."},
		)
		return
	}

	err = uc.userService.Login(loginInfo)
	if err != nil {
		log.Error("Error: " + err.Error())
		c.JSON(
			http.StatusBadRequest,
			domains.ErrorResponse{Error: "Incorrect email or password."},
		)
		return
	}

	response := domains.DefaultResponse{
		Message: "Success",
		Data:    nil,
	}
	c.JSON(http.StatusOK, response)
}
