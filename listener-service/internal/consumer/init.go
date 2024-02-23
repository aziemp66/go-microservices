package consumer

import (
	amqp "github.com/rabbitmq/amqp091-go"
)

func NewConsumer(conn *amqp.Connection, queueName string) (*Consumer, error) {
	consumer := &Consumer{
		conn: conn,
	}

	err := consumer.setup()
	if err != nil {
		return nil, err
	}

	return consumer, nil
}
