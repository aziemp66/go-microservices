package rabbitmq

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"log-service/internal/data"
	"log-service/internal/models/request"
	"strings"

	amqp "github.com/rabbitmq/amqp091-go"
)

func (consumer *consumer) declare() error {
	err := consumer.Channel.ExchangeDeclare(
		"logger",
		"topic",
		false,
		false,
		false,
		false,
		nil,
	)
	if err != nil {
		return err
	}

	q, err := consumer.Channel.QueueDeclare(
		"log",
		false,
		false,
		false,
		false,
		nil,
	)
	if err != nil {
		return err
	}

	err = consumer.Channel.QueueBind(
		q.Name,
		"log.*.*",
		"logger",
		false,
		nil,
	)
	if err != nil {
		return err
	}

	return nil
}

func (consumer *consumer) Consume(ctx context.Context) {
	forever := make(chan bool)
	err := consumer.declare()
	if err != nil {
		log.Fatalf("Failed to declare customer : %s", err.Error())
	}

	msgs, err := consumer.Channel.ConsumeWithContext(ctx, "log", "log-consumer", false, false, false, false, nil)
	if err != nil {
		log.Fatalf("Failed to declare consumer : %s", err.Error())
	}

	fmt.Println("Successfuly Connect to rabbitMQ instance")
	fmt.Println("[*] waiting message..")

	for d := range msgs {
		keys := strings.Split(d.RoutingKey, ".")

		switch keys[len(keys)-1] {
		case "create":
			consumer.createLogEntry(ctx, d)
		}
	}

	<-forever
}

func (consumer *consumer) createLogEntry(ctx context.Context, d amqp.Delivery) {
	var createLogEntry request.CreateLogEntry
	err := json.Unmarshal(d.Body, &createLogEntry)
	if err != nil {
		handleRequeue(ctx, consumer.Channel, d)
		log.Fatalf("Failed parsing json to struct : %s", err.Error())
	}

	if err := consumer.validate.Validate(createLogEntry); err != nil {
		handleRequeue(ctx, consumer.Channel, d)
		log.Fatalf("Failed validating json : %s", err.Error())
	}

	err = consumer.models.LogEntry.Insert(data.LogEntry{
		Name: createLogEntry.Name,
		Data: createLogEntry.Data,
	})
	if err != nil {
		handleRequeue(ctx, consumer.Channel, d)
		log.Fatalf("Failed inserting data to database : %s", err.Error())
	}

	d.Ack(false)
}
