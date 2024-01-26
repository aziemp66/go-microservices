package authentication_adapter

import (
	http_server "broker/internal/http"
	http_error "broker/internal/http/error"
	"broker/internal/model/request"
	"bytes"
	"encoding/json"
	"errors"
	"net/http"

	"github.com/gin-gonic/gin"
)

func Authenticate(ctx *gin.Context, a request.AuthPayload) {
	jsonData, _ := json.MarshalIndent(a, "", "\t")

	req, err := http.NewRequest("POST", "http://authentication-service/authenticate", bytes.NewBuffer(jsonData))
	if err != nil {
		ctx.Error(http_error.NewBadRequest(err.Error()))
		return
	}

	client := &http.Client{}
	response, err := client.Do(req)
	if err != nil {
		ctx.Error(http_error.NewBadRequest(err.Error()))
		return
	}
	defer response.Body.Close()

	var jsonFromService http_server.Response
	err = json.NewDecoder(response.Body).Decode(&jsonFromService)
	if err != nil {
		ctx.Error(http_error.NewBadRequest(err.Error()))
		return
	}

	if response.StatusCode == http.StatusUnauthorized {
		ctx.Error(http_error.NewUnauthorized(jsonFromService.Message))
		return
	} else if response.StatusCode == http.StatusBadRequest {
		ctx.Error(http_error.NewBadRequest(jsonFromService.Message))
		return
	} else if response.StatusCode != http.StatusOK {
		ctx.Error(errors.New(jsonFromService.Message))
		return
	}

	ctx.JSON(http.StatusOK, http_server.Response{
		Message: jsonFromService.Message,
		Value:   jsonFromService.Value,
	})
}
