package broker_controller

import (
	http_server "broker/internal/http"
	"net/http"

	"github.com/gin-gonic/gin"
)

func (controller *controller) Broker(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, http_server.Response{
		Message: "Hit The broker",
	})
}
