package routers

import (
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"

	"food-ordering-system/api/controllers"
)

func Setup(db *gorm.DB, gin *gin.Engine) {
	baseRouter := gin.Group("")

	// setup health check router
	hc := controllers.HealthCheckController{}
	baseRouter.GET("/healthz", hc.HealthCheck)

	v1Router := gin.Group("/api/v1")
	NewUserRouter(db, v1Router)
}
