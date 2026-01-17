package pubsub

import (
	"encoding/json"
	"fmt"
	"log"

	amqp "github.com/rabbitmq/amqp091-go"
)

func SubscribeJSON[T any](conn *amqp.Connection, exchange, queueName, key string, queueType SimpleQueueType, handler func(T)) error {
	ch, queue, err := DeclareAndBind(conn, exchange, queueName, key, queueType)
	if err != nil {
		return fmt.Errorf("failed to check binding to queue")
	}
	amqpCH, err := ch.Consume(queue.Name, "", false, false, false, false, nil)
	if err != nil {
		return fmt.Errorf("failed to consume channel %v", err)
	}
	go func() {
		for delivery := range amqpCH {
			var msg T
			err := json.Unmarshal(delivery.Body, &msg)
			if err != nil {
				delivery.Nack(false, false)
				fmt.Printf("failde to unmarshal delivery message: %v", err)
				continue
			}
			handler(msg)
			if err := delivery.Ack(false); err != nil {
				log.Printf("failed to ack messaage: %v", err)
			}
		}
	}()
	return nil
}
