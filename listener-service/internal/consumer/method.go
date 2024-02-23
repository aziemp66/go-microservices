package consumer

import (
	"encoding/json"
	"fmt"
	"listener/internal/event"
	"log"
)

func (c *Consumer) setup() error {
	channel, err := c.conn.Channel()
	if err != nil {
		return err
	}

	return event.DeclareExchange(channel)
}

func (c *Consumer) Listen(topics []string) error {
	ch, err := c.conn.Channel()
	if err != nil {
		return err
	}
	defer ch.Close()

	q, err := event.DeclareRandomQueue(ch)
	if err != nil {
		return err
	}

	for _, v := range topics {
		err = ch.QueueBind(
			q.Name,
			v,
			"logs_topic",
			false,
			nil,
		)
		if err != nil {
			return err
		}
	}

	messages, err := ch.Consume(q.Name, "", true, false, false, false, nil)

	forever := make(chan bool)
	go func() {
		for d := range messages {
			var payload Payload
			_ = json.Unmarshal(d.Body, &payload)

			go handlePayload(payload)
		}
	}()
	<-forever

	fmt.Printf("Waiting for messages [Exchange,Queue] [logs_topic, %s]\n", q.Name)
}

func handlePayload(p Payload) {
	switch p.Name {
	case "log", "event":
		//log whatever we got
		err := logEvent(p)
		if err != nil {
			log.Println(err)
		}
	default:
		err := logEvent(p)
		if err != nil {
			log.Println(err)
		}
	}
}

func logEvent(p Payload) error {

}
