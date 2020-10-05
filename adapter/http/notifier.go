package http

import (
	"encoding/json"
	"os"

	"github.com/GSabadini/go-challenge/domain/entity"
	"github.com/pkg/errors"
)

const (
	enviado = "Enviado"
)

var (
	errFailedToNotify = errors.New("failed to notify")
)

type (
	Notifier struct {
		client HTTPGetter
	}

	NotifierResponse struct {
		Message string
	}
)

func NewNotifier(client HTTPGetter) Notifier {
	return Notifier{
		client: client,
	}
}

func (n Notifier) Notify(t entity.Transfer) error {
	r, err := n.client.Get(os.Getenv("NOTIFY_URI"))
	//r, err := a.client.Get("https://run.mocky.io/v3/ed736a57-0c29-4433-92af-42228052e5ae")
	if err != nil {
		return errors.Wrap(err, errFailedToNotify.Error())
	}

	b := &NotifierResponse{}
	err = json.NewDecoder(r.Body).Decode(&b)
	if err != nil {
		return errors.Wrap(err, errFailedToNotify.Error())
	}

	if b.Message != enviado {
		return errFailedToNotify
	}

	return nil
}
