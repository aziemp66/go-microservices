package broker_controller

import (
	authentication_adapter "broker/internal/adapter/authentication"
	log_adapter "broker/internal/adapter/log"
	mail_adapter "broker/internal/adapter/mail"
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
	case "log":
		log_adapter.LogItem(ctx, req.Log)
	case "mail":
		mail_adapter.SendMail(ctx, req.Mail)
	default:
		ctx.Error(http_error.NewBadRequest("Unknown Action"))
		return
	}
}
