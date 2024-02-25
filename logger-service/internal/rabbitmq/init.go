package rabbitmq

import (
	"log-service/internal/data"
	"log-service/internal/validation"
	"time"

	amqp "github.com/rabbitmq/amqp091-go"
)

func NewConsumer(conn string, models *data.Models, validate *validation.Validate) (*consumer, error) {
	var connection *amqp.Connection

	var count int = 1
	for {
		conn, err := amqp.Dial(conn)
		if err != nil {
			if count <= 5 {
				time.Sleep(3 * time.Second)
				count++
				continue
			}
			return nil, err
		}
		connection = conn
		break
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
