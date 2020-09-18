package http

import (
	"fmt"
	"net/http"
)

type Authorizer struct {
	client http.Client
}

func (a Authorizer) Authorized() (bool, error) {
	fmt.Println("Authorizado")
	return true, nil
}
