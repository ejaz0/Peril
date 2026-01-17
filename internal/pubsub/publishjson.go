package pubsub

import (
	"context"
	"encoding/json"
	"fmt"

	amqp "github.com/rabbitmq/amqp091-go"
)

func PublishJSON[T any](ch *amqp.Channel, exchange, key string, val T) error {
	marshalVal, err := json.Marshal(val)
	if err != nil {
		return fmt.Errorf("failed to marshal into bytes")
	}
	err = ch.PublishWithContext(context.Background(), exchange, key, false, false, amqp.Publishing{
		ContentType: "application/json",
		Body:        marshalVal,
	})
	if err != nil {
		return fmt.Errorf("failed to publish the message")
	}
	return nil
}
