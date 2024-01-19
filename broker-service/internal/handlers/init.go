package handlers

import (
	"broker/internal/helpers"
	"encoding/json"
	"net/http"
)

func Broker(w http.ResponseWriter, r *http.Request) {
	payload := helpers.JsonResponse{
		Error:   false,
		Message: "Hit The Broker",
	}

	helpers.WriteJSON(w, http.StatusOK, payload)

	out, _ := json.MarshalIndent(payload, "", "\t")
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusAccepted)

	w.Write(out)
}
