package http

import (
	"context"
	"encoding/json"
	"os"

	"github.com/GSabadini/go-challenge/domain/entity"
	"github.com/GSabadini/go-challenge/usecase"
	"github.com/pkg/errors"
)

const autorizado = "Autorizado"

var errAuthorizationDenied = errors.New("authorization denied")

type (
	authorizer struct {
		client HTTPGetter
	}

	authorizerResponse struct {
		Message string
	}
)

// NewAuthorizer creates new authorizer with its dependencies
func NewAuthorizer(client HTTPGetter) usecase.Authorizer {
	return authorizer{
		client: client,
	}
}

// Authorized
func (a authorizer) Authorized(_ context.Context, _ entity.Transfer) (bool, error) {
	r, err := a.client.Get(os.Getenv("AUTHORIZER_URI"))
	if err != nil {
		return false, errors.Wrap(err, errAuthorizationDenied.Error())
	}

	b := &authorizerResponse{}
	err = json.NewDecoder(r.Body).Decode(&b)
	if err != nil {
		return false, errors.Wrap(err, errAuthorizationDenied.Error())
	}

	if b.Message != autorizado {
		return false, errAuthorizationDenied
	}

	return true, nil
}
