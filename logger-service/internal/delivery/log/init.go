package log_delivery

import (
	"log-service/internal/data"
	"log-service/internal/validation"
)

type logDelivery struct {
	validation *validation.Validate
	models     *data.Models
}

func NewLogDelivery(validation *validation.Validate, models *data.Models) LogDelivery {
	return &logDelivery{
		validation: validation,
		models:     models,
	}
}
