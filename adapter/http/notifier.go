package http

import (
	"fmt"
	"github.com/GSabadini/go-challenge/domain/entity"
)

type Notifier struct {
	client HTTPGetter
}

func NewNotifier(client HTTPGetter) Notifier {
	return Notifier{
		client: client,
	}
}

func (n Notifier) Notify(t entity.Transfer) error {
	fmt.Println("Notificado")
	return nil
}
