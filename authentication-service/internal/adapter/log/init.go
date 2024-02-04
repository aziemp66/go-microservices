package log_adapter

import (
	http_server "authentication/internal/http"
	"bytes"
	"encoding/json"
	"errors"
	"net/http"
)

func LogRequest(name string, data any) error {
	var entry struct {
		Name string `json:"name"`
		Data any    `json:"data"`
	}

	entry.Name = name
	entry.Data = data

	jsonData, _ := json.MarshalIndent(entry, "", "\t")

	logServiceURL := "http://logger-service/log"

	request, err := http.NewRequest("POST", logServiceURL, bytes.NewBuffer(jsonData))
	if err != nil {
		return err
	}
	request.Header.Set("Content-Type", "application/json")

	client := &http.Client{}

	response, err := client.Do(request)
	if err != nil {
		return err
	}

	defer response.Body.Close()
	var jsonFromService http_server.Response
	err = json.NewDecoder(response.Body).Decode(&jsonFromService)
	if err != nil {
		return err
	}
	if response.StatusCode != http.StatusCreated {
		var body http_server.Response
		json.NewDecoder(response.Body).Decode(&body)

		return errors.New(body.Message)
	}

	return nil
}
