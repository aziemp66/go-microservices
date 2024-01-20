package handlers

import (
	"broker/internal/helpers"
	"net/http"
)

func Broker(w http.ResponseWriter, r *http.Request) {
	payload := helpers.JsonResponse{
		Error:   false,
		Message: "Hit The Broker",
	}

	helpers.WriteJSON(w, http.StatusOK, payload)
}
