package routers

import (
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"

	"food-ordering-system/api/controllers"
	"food-ordering-system/internal/repositories"
	"food-ordering-system/internal/services"
)

func NewUserRouter(db *gorm.DB, group *gin.RouterGroup) {
	ur := repositories.NewUserRepositoryImpl(db)
	us := services.NewUserServiceImpl(ur)
	uc := controllers.NewUserController(us)

	group.POST("/signup", uc.Signup)
	group.POST("/login", uc.Login)
}
