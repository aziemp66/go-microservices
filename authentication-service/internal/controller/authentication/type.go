package authentication_controller

import (
	data "authentication/internal/models"

	"github.com/gin-gonic/gin"
)

type controller struct {
	Models data.Models
}

func AuthenticationController(router *gin.RouterGroup, models data.Models) {
	c := controller{
		Models: models,
	}

	router.POST("/authenticate", c.Authenticate)
}
