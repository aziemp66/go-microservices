package authentication_controller

import (
	http_server "authentication/internal/http"
	http_error "authentication/internal/http/error"
	"fmt"

	"github.com/gin-gonic/gin"
)

func (controller *controller) Authenticate(ctx *gin.Context) {
	var req struct {
		Email    string `json:"email"`
		Password string `json:"password"`
	}

	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.Error(http_error.NewBadRequest(err.Error()))
		return
	}

	user, err := controller.Models.User.GetByEmail(req.Email)
	if err != nil {
		ctx.Error(http_error.NewUnauthorized(fmt.Sprintf("user with %s email is not found : %s", req.Email, err.Error())))
		return
	}

	valid, err := controller.Models.User.PasswordMatches(req.Email, req.Password)
	if err != nil || !valid {
		ctx.Error(http_error.NewUnauthorized(fmt.Sprintf("password is wrong : %s", err.Error())))
		return
	}

	ctx.JSON(200, http_server.Response{
		Message: fmt.Sprintf("Logged In As User : %s", user.Email),
		Value:   user,
	})
}
