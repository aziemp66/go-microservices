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

	err := helpers.WriteJSON(w, http.StatusOK, payload)
	if err != nil {
		helpers.ErrorJSON(w, r, err, http.StatusBadRequest)
	}
}
