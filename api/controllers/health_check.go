package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"

	"food-ordering-system/internal/domains"
)

type HealthCheckController struct{}

func NewHealthCheckController() *HealthCheckController {
	return &HealthCheckController{}
}

// @Summary		Health Check
// @Success		200	{object}	response.HealthCheckResponse{}
// @Router		/healthz [get]
func (*HealthCheckController) HealthCheck(c *gin.Context) {
	response := domains.HealthCheckResponse{Messaage: "OK"}

	c.JSON(http.StatusOK, response)
}
