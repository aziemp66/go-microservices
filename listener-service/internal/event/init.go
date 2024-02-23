// Package event : amqp event
package event

import (
	amqp "github.com/rabbitmq/amqp091-go"
)

// DeclareExchange : Declare DeclareExchange
func DeclareExchange(ch *amqp.Channel) error {
	return ch.ExchangeDeclare(
		"logs_topic", // name
		"topic",      //type
		true,         //durable?
		false,        //auto-delete
		false,        //internal
		false,        //no-wait
		nil,          //arguments
	)
}

func DeclareRandomQueue(ch *amqp.Channel) (amqp.Queue, error) {
	return ch.QueueDeclare(
		"",    //name
		false, //durable
		false, //auto-delete
		true,  //exclusive
		false, //no-wait
		nil,   //arguments
	)
}
