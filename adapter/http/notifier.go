package http

import (
	"fmt"
	"net/http"

	"github.com/GSabadini/go-challenge/domain/entity"
)

type Notifier struct {
	client http.Client
}

func (n Notifier) Notify(t entity.Transfer) error {
	fmt.Println("Notificado")
	return nil
}
