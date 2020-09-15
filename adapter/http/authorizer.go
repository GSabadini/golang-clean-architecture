package http

import "fmt"

type Authorizer struct{}

func (a Authorizer) Authorized() (bool, error) {
	fmt.Println("Authorizado")
	return true, nil
}
