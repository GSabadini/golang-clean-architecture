package presenter

import (
	"github.com/GSabadini/go-challenge/domain/entity"
	"github.com/GSabadini/go-challenge/usecase"
)

type CreateTransferPresenter struct{}

func (c CreateTransferPresenter) Output(t entity.Transfer) usecase.CreateTransferOutput {
	return usecase.CreateTransferOutput{
		ID:        t.ID(),
		PayerID:   t.Payer(),
		PayeeID:   t.Payee(),
		Value:     t.Value(),
		CreatedAt: t.CreatedAt(),
	}
}
