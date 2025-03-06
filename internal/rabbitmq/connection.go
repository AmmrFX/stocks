package rabbitmq

import (
	"github.com/rabbitmq/amqp091-go"
)

func Connect() (*amqp091.Connection, error) {
	conn, err := amqp091.Dial("amqp://guest:guest@localhost:5672/")
	if err != nil {
		return nil, err
	}
	return conn, nil
}
