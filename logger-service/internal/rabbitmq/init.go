package rabbitmq

import (
	"log-service/internal/data"
	"log-service/internal/validation"

	amqp "github.com/rabbitmq/amqp091-go"
)

func NewConsumer(conn string, models *data.Models, validate *validation.Validate) (*consumer, error) {
	connection, err := amqp.Dial("amqp://admin:password@rabbitmq:5672/")
	if err != nil {
		return nil, err
	}

	channel, err := connection.Channel()
	if err != nil {
		return nil, err
	}

	return &consumer{
		channel:    channel,
		connection: connection,
		models:     models,
	}, nil
}
