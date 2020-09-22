package http

import (
	"fmt"
	"net/http"

	"github.com/GSabadini/go-challenge/domain/entity"
)

type Authorizer struct {
	client http.Client
}

func (a Authorizer) Authorized(t entity.Transfer) (bool, error) {
	fmt.Println("Authorizado")
	return true, nil
}
