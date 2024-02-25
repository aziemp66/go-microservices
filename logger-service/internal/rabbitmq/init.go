package rabbitmq

import (
	"log-service/internal/data"
	"log-service/internal/validation"

	amqp "github.com/rabbitmq/amqp091-go"
)

func NewConsumer(conn string, models *data.Models, validate *validation.Validate) (*consumer, error) {
	connection, err := amqp.Dial(conn)
	if err != nil {
		return nil, err
	}

	Channel, err := connection.Channel()
	if err != nil {
		return nil, err
	}

	return &consumer{
		Channel:    Channel,
		Connection: connection,
		models:     models,
		validate:   validate,
	}, nil
}
