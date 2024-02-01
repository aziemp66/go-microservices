package log_delivery

import (
	"log-service/internal/data"
	"log-service/internal/validation"
)

type LogDelivery struct {
	validation *validation.Validate
	models     *data.Models
}

func NewLogDelivery(validation *validation.Validate, models *data.Models) *LogDelivery {
	return &LogDelivery{
		validation: validation,
		models:     models,
	}
}
