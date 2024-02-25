package rabbitmq

import (
	"log-service/internal/data"
	"log-service/internal/validation"

	amqp "github.com/rabbitmq/amqp091-go"
)

type consumer struct {
	connection *amqp.Connection
	channel    *amqp.Channel
	models     *data.Models
	validate   *validation.Validate
}
