package usecase

import (
	"context"
	"errors"
	"time"

	"github.com/GSabadini/go-challenge/domain/entity"
	"github.com/GSabadini/go-challenge/domain/vo"
)

type Authorizer interface {
	Authorized() (bool, error)
}

type Notifier interface {
	//Notify(entity.Transfer) error
	Notify() error
}

type TransferInput struct {
	ID        vo.Uuid   `json:"id"`
	PayerID   vo.Uuid   `json:"payer"`
	PayeeID   vo.Uuid   `json:"payee"`
	Value     vo.Money  `json:"value"`
	CreatedAt time.Time `json:"created_at"`
}

type CreateTransferUseCase interface {
	Execute(context.Context, TransferInput) (entity.Transfer, error)
}

type CreateTransferInteractor struct {
	TransferRepo entity.TransferRepository
	UserRepo     entity.UserRepository

	ExternalAuthorizer Authorizer
	Notifier           Notifier
}

func (c CreateTransferInteractor) Execute(ctx context.Context, i TransferInput) (entity.Transfer, error) {
	if err := c.process(ctx, i.PayerID, i.PayeeID, i.Value); err != nil {
		return entity.Transfer{}, err
	}

	transfer := entity.NewTransfer(
		"vo.Uuid{}",
		i.PayerID,
		i.PayeeID,
		i.Value,
		time.Now(),
	)

	err := c.TransferRepo.Save(ctx, transfer)
	if err != nil {
		return entity.Transfer{}, errors.New("")
	}

	return transfer, nil
}

func (c CreateTransferInteractor) process(ctx context.Context, payerID vo.Uuid, payeeID vo.Uuid, value vo.Money) error {
	payer, err := c.UserRepo.FindByID(ctx, payerID)
	if err != nil {
		return errors.New("")
	}

	if !payer.CanTransfer() {
		return errors.New("!authorized")
	}

	payee, err := c.UserRepo.FindByID(ctx, payeeID)
	if err != nil {
		return errors.New("")
	}

	err = payer.Withdraw(value)
	if err != nil {
		return errors.New("")
	}

	payee.Deposit(value)

	/**
	Start Transaction
	*/

	//c.UserRepo.InitTransaction()
	err = c.UserRepo.UpdateWallet(ctx, payerID, payer.Wallet().Money())
	if err != nil {
		return errors.New("")
	}

	err = c.UserRepo.UpdateWallet(ctx, payeeID, payee.Wallet().Money())
	if err != nil {
		//c.UserRepo.Rollback()
		return errors.New("")
	}

	ok, err := c.ExternalAuthorizer.Authorized()
	if err != nil || !ok {
		//c.UserRepo.Rollback()
		return errors.New("")
	}

	//c.UserRepo.Commit()
	/**
	End Transaction
	*/

	err = c.Notifier.Notify()
	if err != nil {
		return errors.New("")
	}

	return nil
}
