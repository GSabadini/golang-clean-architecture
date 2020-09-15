package http

import "fmt"

type Notifier struct{}

func (n Notifier) Notify() error {
	fmt.Println("Notificado")
	return nil
}
