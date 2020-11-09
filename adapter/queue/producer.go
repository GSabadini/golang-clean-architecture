package queue

import (
	"github.com/GSabadini/golang-clean-architecture/adapter/logger"
	"github.com/streadway/amqp"
)

type producer struct {
	channel   *amqp.Channel
	queueName string
	log       logger.Logger
	logKey    string
}

// NewProducer creates new producer with its dependencies
func NewProducer(ch *amqp.Channel, qn string, l logger.Logger) Producer {
	return producer{
		channel:   ch,
		queueName: qn,
		log:       l,
		logKey:    "queue_producer",
	}
}

// Publish sends a Publishing from the client to an exchange on the server
func (p producer) Publish(message []byte) error {
	if err := p.channel.Publish(
		"",
		p.queueName,
		false,
		false,
		amqp.Publishing{
			Headers:     amqp.Table{},
			ContentType: "text/plain",
			Body:        message,
		}); err != nil {
		p.log.WithFields(logger.Fields{
			"key":   p.logKey,
			"error": err.Error(),
		}).Errorf("failed to publish message: %s", message)

		return err
	}

	p.log.WithFields(logger.Fields{
		"key": p.logKey,
	}).Infof("new message publish: %s", message)

	return nil
}
