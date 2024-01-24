package broker_controller

import (
	authentication_adapter "broker/internal/adapter/authentication"
	http_server "broker/internal/http"
	http_error "broker/internal/http/error"
	"broker/internal/model/request"
	"net/http"

	"github.com/gin-gonic/gin"
)

func (controller *controller) Broker(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, http_server.Response{
		Message: "Hit The broker",
	})
}

func (controller *controller) HandleSubmission(ctx *gin.Context) {
	var req request.RequestPayload
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.Error(err)
		return
	}

	switch req.Action {
	case "auth":
		authentication_adapter.Authenticate(ctx, req.Auth)
	default:
		ctx.Error(http_error.NewBadRequest("Unknown Action"))
		return
	}
}
