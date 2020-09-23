package http

import (
	"fmt"
	"net/http"

	"github.com/GSabadini/go-challenge/domain/entity"
)

type Authorizer struct {
	client http.Client
}

func NewAuthorizer(client http.Client) Authorizer {
	return Authorizer{
		client: client,
	}
}

func (a Authorizer) Authorized(t entity.Transfer) (bool, error) {
	fmt.Println("Authorizado")
	return true, nil
}
