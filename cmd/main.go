package main

import (
	"github.com/gin-gonic/gin"
	log "github.com/sirupsen/logrus"

	"food-ordering-system/api/routers"
	"food-ordering-system/bootstrap"
	"food-ordering-system/internal/database"
	_ "food-ordering-system/internal/logger"
)

func main() {
	env := bootstrap.NewEnv()
	db := database.NewPostgresConnection(env)
	gin := gin.Default()

	routers.Setup(db, gin)

	log.Info("Running server on port 8080")
	gin.Run("0.0.0.0:8080")
}
