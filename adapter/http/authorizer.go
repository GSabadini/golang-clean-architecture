package http

import (
	"encoding/json"
	"fmt"
	"github.com/GSabadini/go-challenge/domain/entity"
	"github.com/pkg/errors"
	"time"
)

var (
	errNotAuthorizer = errors.New("failed authorizer")
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

func (a Authorizer) Authorized(t entity.Transfer) (bool, error) {
	fmt.Println("Authorizado")
	return true, nil
}

func (a Authorizer) Authorized1() (bool, error) {
	start := time.Now()
	r, err := a.client.Get("https://run.mocky.io/v3/ed736a57-0c29-4433-92af-42228052e5ae")
	if err != nil {
		fmt.Println("end", time.Since(start).Seconds())

		return false, errors.Wrap(err, errNotAuthorizer.Error())
	}

	b := &AuthorizerResponse{}

	err = json.NewDecoder(r.Body).Decode(&b)
	if err != nil {
		//http.Error(w, err.Error(), 400)
		return false, errors.Wrap(err, errNotAuthorizer.Error())
	}

	if b.Message != "Autorizado" {
		return false, errNotAuthorizer
	}
	fmt.Println(b.Message)
	return true, nil
}
