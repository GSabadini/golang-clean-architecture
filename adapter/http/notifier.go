package http

import (
	"context"
	"encoding/json"
	"os"

	"github.com/GSabadini/go-challenge/domain/entity"
	"github.com/GSabadini/go-challenge/usecase"
	"github.com/pkg/errors"
)

const enviado = "Enviado"

var errFailedToNotify = errors.New("failed to notify")

type (
	notifier struct {
		client HTTPGetter
	}

	notifierResponse struct {
		Message string
	}
)

// NewNotifier creates new notifier with its dependencies
func NewNotifier(client HTTPGetter) usecase.Notifier {
	return notifier{
		client: client,
	}
}

// Notify
func (n notifier) Notify(_ context.Context, _ entity.Transfer) error {
	r, err := n.client.Get(os.Getenv("NOTIFY_URI"))
	if err != nil {
		return errors.Wrap(err, errFailedToNotify.Error())
	}

	b := &notifierResponse{}
	err = json.NewDecoder(r.Body).Decode(&b)
	if err != nil {
		return errors.Wrap(err, errFailedToNotify.Error())
	}

	if b.Message != enviado {
		return errFailedToNotify
	}

	return nil
}
