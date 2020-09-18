package http

import (
	"fmt"
	"net/http"
)

type Notifier struct {
	client http.Client
}

func (n Notifier) Notify() error {
	fmt.Println("Notificado")
	return nil
}
