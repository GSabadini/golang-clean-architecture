package http

import (
	"fmt"
	"github.com/GSabadini/go-challenge/domain/entity"
)

type Authorizer struct {
	client HTTPGetter
}

func NewAuthorizer(client HTTPGetter) Authorizer {
	return Authorizer{
		client: client,
	}
}

func (a Authorizer) Authorized(t entity.Transfer) (bool, error) {
	fmt.Println("Authorizado")
	return true, nil
}
