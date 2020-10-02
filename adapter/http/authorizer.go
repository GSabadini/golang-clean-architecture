package http

import (
	"encoding/json"
	"fmt"
	"github.com/GSabadini/go-challenge/domain/entity"
	"github.com/pkg/errors"
	"os"
)

const (
	autorizado = "Autorizado"
)

var (
	errAuthorizationDenied = errors.New("authorization denied")
)

type (
	Authorizer struct {
		client HTTPGetter
	}

	AuthorizerResponse struct {
		Message string
	}
)

func NewAuthorizer(client HTTPGetter) Authorizer {
	return Authorizer{
		client: client,
	}
}

func (a Authorizer) Authorized(_ entity.Transfer) (bool, error) {
	r, err := a.client.Get(os.Getenv("AUTHORIZER_URI"))
	//r, err := a.client.Get("https://run.mocky.io/v3/ed736a57-0c29-4433-92af-42228052e5ae")
	if err != nil {
		return false, errors.Wrap(err, errAuthorizationDenied.Error())
	}

	b := &AuthorizerResponse{}
	err = json.NewDecoder(r.Body).Decode(&b)
	if err != nil {
		return false, errors.Wrap(err, errAuthorizationDenied.Error())
	}

	if b.Message != autorizado {
		return false, errAuthorizationDenied
	}

	fmt.Println("Autorizado")

	return true, nil
}
