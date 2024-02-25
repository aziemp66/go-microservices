package rabbitmq

import (
	"context"
	"log"

	"github.com/rabbitmq/amqp091-go"
)

func handleRequeue(ctx context.Context, c *amqp091.Channel, d amqp091.Delivery) {
	requeueCount := d.Headers["x-requeue-count"].(int)
	if requeueCount < 3 {
		c.PublishWithContext(
			ctx,
			"logger",
			d.RoutingKey,
			false,
			false,
			amqp091.Publishing{
				ContentType: d.ContentType,
				Headers: amqp091.Table{
					"x-requeue-count": requeueCount + 1,
				},
				Body: d.Body,
			},
		)
		d.Ack(false)
		return
	}
	err := c.ExchangeDeclare(
		"logger",
		"topic",
		false,
		false,
		false,
		false,
		nil,
	)
	if err != nil {
		log.Fatalf("Failed to declare exchange : %s", err.Error())
	}

	_, err = c.QueueDeclare(
		"log_dlq",
		false,
		false,
		false,
		false,
		nil,
	)
	if err != nil {
		log.Fatalf("Failed to declare queue : %s", err.Error())
	}

	err = c.PublishWithContext(ctx, "logger", d.RoutingKey, false, false, amqp091.Publishing{
		ContentType: d.ContentType,
		Headers: amqp091.Table{
			"x-requeue-count": requeueCount + 1,
		},
		Body: d.Body,
	})
	if err != nil {
		log.Fatalf("Failed to publish message : %s", err.Error())
	}
	d.Ack(false)
}
