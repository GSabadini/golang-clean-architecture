package http

import (
	"context"
	"encoding/json"
	"github.com/GSabadini/go-challenge/adapter/queue"
	"os"

	"github.com/GSabadini/go-challenge/adapter/logger"
	"github.com/GSabadini/go-challenge/domain/entity"
	"github.com/GSabadini/go-challenge/usecase"
	"github.com/pkg/errors"
)

const (
	enviado = "Enviado"
	logKey  = "send_notify"
)

var errFailedToNotify = errors.New("failed to notify")

type (
	notifier struct {
		client    HTTPGetter
		publisher queue.Producer
		log       logger.Logger
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
	}
}

// Notify send a notification
func (n notifier) Notify(_ context.Context, _ entity.Transfer) {
	res, err := n.client.Get(os.Getenv("NOTIFY_URI"))
	if err != nil {
		n.publish(err)
		return
	}

	b := &notifierResponse{}
	err = json.NewDecoder(res.Body).Decode(&b)
	if err != nil {
		n.publish(err)
		return
	}

	if b.Message != enviado {
		n.publish(errFailedToNotify)
		return
	}

	n.log.WithFields(logger.Fields{
		"key":         logKey,
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
			"key":   logKey,
			"error": err.Error(),
		}).Errorf("failed to marshal message")
		return
	}

	if err := n.publisher.Publish(message); err != nil {
		n.log.WithFields(logger.Fields{
			"key":   logKey,
			"error": err.Error(),
		}).Errorf("failed to publish to the queue")
		return
	}

	n.log.WithFields(logger.Fields{
		"key": logKey,
	}).Infof("success to publish to the queue")
}
