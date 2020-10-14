package queue

import (
	"os"

	"github.com/streadway/amqp"
)

// RabbitMQHandler defines the RabbitMQ handler
type RabbitMQHandler struct {
	conn    *amqp.Connection
	queue   amqp.Queue
	channel *amqp.Channel
}

// NewRabbitMQHandler creates new RabbitMQHandler
func NewRabbitMQHandler() (*RabbitMQHandler, error) {
	conn, err := amqp.Dial(os.Getenv("RABBITMQ_URI"))
	if err != nil {
		return &RabbitMQHandler{}, err
	}

	channel, err := conn.Channel()
	if err != nil {
		return &RabbitMQHandler{}, err
	}

	queue, err := channel.QueueDeclare(
		"notify",
		false,
		false,
		false,
		false,
		nil,
	)
	if err != nil {
		return &RabbitMQHandler{}, err
	}

	return &RabbitMQHandler{
		conn:    conn,
		queue:   queue,
		channel: channel,
	}, nil
}

// Conn returns the conn property
func (r RabbitMQHandler) Conn() *amqp.Connection {
	return r.conn
}

// Queue returns the queue property
func (r RabbitMQHandler) Queue() amqp.Queue {
	return r.queue
}

// Channel returns the channel property
func (r RabbitMQHandler) Channel() *amqp.Channel {
	return r.channel
}
