package presenter

import (
	"github.com/GSabadini/go-challenge/domain/entity"
	"github.com/GSabadini/go-challenge/usecase"
)

type CreateTransferPresenter struct{}

func (c CreateTransferPresenter) Output(t entity.Transfer) usecase.CreateTransferOutput {
	return usecase.CreateTransferOutput{
		ID:        string(t.ID()),
		PayerID:   string(t.Payer()),
		PayeeID:   string(t.Payee()),
		Value:     t.Value().Amount(),
		CreatedAt: t.CreatedAt().String(),
	}
}
