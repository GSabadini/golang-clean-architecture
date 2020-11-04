package http

import (
	"context"
	"encoding/json"
	"github.com/GSabadini/golang-clean-architecture/adapter/queue"
	"os"

	"github.com/GSabadini/golang-clean-architecture/adapter/logger"
	"github.com/GSabadini/golang-clean-architecture/domain/entity"
	"github.com/GSabadini/golang-clean-architecture/usecase"
	"github.com/pkg/errors"
)

const (
	enviado = "Enviado"
)

var errFailedToNotify = errors.New("failed to notify")

type (
	notifier struct {
		client    HTTPGetter
		publisher queue.Producer
		log       logger.Logger
		logKey    string
	}

	notifierResponse struct {
		Message string
	}
)

// NewNotifier creates new notifier with its dependencies
func NewNotifier(c HTTPGetter, p queue.Producer, l logger.Logger) usecase.Notifier {
	return notifier{
		client:    c,
		publisher: p,
		log:       l,
		logKey:    "send_notify",
	}
}

// Notify send a notification
func (n notifier) Notify(_ context.Context, _ entity.Transfer) {
	res, err := n.client.Get(os.Getenv("NOTIFY_URI"))
	if err != nil {
		n.log.WithFields(logger.Fields{
			"key":   n.logKey,
			"error": err.Error(),
		}).Errorf("failed to client")

		n.publish(err)
		return
	}

	b := &notifierResponse{}
	err = json.NewDecoder(res.Body).Decode(&b)
	if err != nil {
		n.log.WithFields(logger.Fields{
			"key":   n.logKey,
			"error": err.Error(),
		}).Errorf("failed to marshal message")

		n.publish(err)
		return
	}

	if b.Message != enviado {
		n.publish(errFailedToNotify)
		return
	}

	n.log.WithFields(logger.Fields{
		"key":         n.logKey,
		"http_status": res.StatusCode,
	}).Infof("success to notify")
}

func (n notifier) publish(err error) {
	message, err := json.Marshal(map[string]string{
		"uri":   os.Getenv("NOTIFY_URI"),
		"error": err.Error(),
	})
	if err != nil {
		n.log.WithFields(logger.Fields{
			"key":   n.logKey,
			"error": err.Error(),
		}).Errorf("failed to marshal message")
		return
	}

	if err := n.publisher.Publish(message); err != nil {
		n.log.WithFields(logger.Fields{
			"key":   n.logKey,
			"error": err.Error(),
		}).Errorf("failed to publish to the queue")
		return
	}

	n.log.WithFields(logger.Fields{
		"key": n.logKey,
	}).Infof("success to publish to the queue")
}
