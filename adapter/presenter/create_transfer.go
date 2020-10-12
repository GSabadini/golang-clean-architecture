package presenter

import (
	"github.com/GSabadini/go-challenge/domain/entity"
	"github.com/GSabadini/go-challenge/usecase"
)

type createTransferPresenter struct{}

// NewCreateTransferPresenter creates new createTransferPresenter
func NewCreateTransferPresenter() usecase.CreateTransferPresenter {
	return createTransferPresenter{}
}

// Output returns the transfer creation response
func (c createTransferPresenter) Output(t entity.Transfer) usecase.CreateTransferOutput {
	return usecase.CreateTransferOutput{
		ID:        t.ID().Value(),
		PayerID:   t.Payer().Value(),
		PayeeID:   t.Payee().Value(),
		Value:     t.Value().Amount().Value(),
		CreatedAt: t.CreatedAt().String(),
	}
}
