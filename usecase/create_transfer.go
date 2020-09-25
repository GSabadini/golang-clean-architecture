package usecase

import (
	"context"
	"errors"
	"fmt"
	"time"

	"github.com/GSabadini/go-challenge/domain/entity"
	"github.com/GSabadini/go-challenge/domain/vo"
)

type Authorizer interface {
	Authorized(entity.Transfer) (bool, error)
}

type Notifier interface {
	Notify(entity.Transfer) error
}

//Output port
type CreateTransferPresenter interface {
	Output(entity.Transfer) CreateTransferOutput
}

//Input port
type CreateTransferUseCase interface {
	Execute(context.Context, CreateTransferInput) (CreateTransferOutput, error)
}

//Input data
type CreateTransferInput struct {
	ID        vo.Uuid
	PayerID   vo.Uuid
	PayeeID   vo.Uuid
	Value     vo.Money
	CreatedAt time.Time
}

//Output data
type CreateTransferOutput struct {
	ID        string `json:"id"`
	PayerID   string `json:"payer"`
	PayeeID   string `json:"payee"`
	Value     int64  `json:"value"`
	CreatedAt string `json:"created_at"`
}

type CreateTransferInteractor struct {
	createTransferRepo   entity.CreateTransferRepository
	updateUserWalletRepo entity.UpdateUserWalletRepository
	findUserByIDRepo     entity.FindUserByIDRepository
	pre                  CreateTransferPresenter
	externalAuthorizer   Authorizer
	notifier             Notifier
}

func NewCreateTransferInteractor(
	createTransferRepo entity.CreateTransferRepository,
	updateUserWalletRepo entity.UpdateUserWalletRepository,
	findUserByIDRepo entity.FindUserByIDRepository,
	pre CreateTransferPresenter,
	externalAuthorizer Authorizer,
	notifier Notifier,
) CreateTransferInteractor {
	return CreateTransferInteractor{
		createTransferRepo:   createTransferRepo,
		updateUserWalletRepo: updateUserWalletRepo,
		findUserByIDRepo:     findUserByIDRepo,
		pre:                  pre,
		externalAuthorizer:   externalAuthorizer,
		notifier:             notifier,
	}
}

func (c CreateTransferInteractor) Execute(ctx context.Context, i CreateTransferInput) (CreateTransferOutput, error) {
	if err := c.process(ctx, i.PayerID, i.PayeeID, i.Value); err != nil {
		return c.pre.Output(entity.Transfer{}), err
	}

	uuid, err := vo.NewUuid(entity.NewUUID())
	if err != nil {
		return c.pre.Output(entity.Transfer{}), err
	}

	transfer, err := c.createTransferRepo.Create(ctx, entity.NewTransfer(
		uuid,
		i.PayerID,
		i.PayeeID,
		i.Value,
		time.Now(),
	))
	if err != nil {
		return c.pre.Output(entity.Transfer{}), err
	}

	ok, err := c.externalAuthorizer.Authorized(transfer)
	if err != nil || !ok {
		//c.updateUserWalletRepo.Rollback()
		return c.pre.Output(entity.Transfer{}), err
	}

	err = c.notifier.Notify(transfer)
	if err != nil {
		return c.pre.Output(entity.Transfer{}), err
	}

	return c.pre.Output(transfer), nil
}

func (c CreateTransferInteractor) process(ctx context.Context, payerID vo.Uuid, payeeID vo.Uuid, value vo.Money) error {
	payer, err := c.findUserByIDRepo.FindByID(ctx, payerID)
	if err != nil {
		return err
	}

	if !payer.CanTransfer() {
		return errors.New("!authorized")
	}

	payee, err := c.findUserByIDRepo.FindByID(ctx, payeeID)
	if err != nil {
		return err
	}

	err = payer.Withdraw(value)
	if err != nil {
		return err
	}

	payee.Deposit(value)

	/**
	Start Transaction
	*/

	//c.updateUserWalletRepo.InitTransaction()
	err = c.updateUserWalletRepo.UpdateWallet(ctx, payerID, payer.Wallet().Money())
	fmt.Println(payer.Wallet().Money())
	if err != nil {
		return err
	}

	err = c.updateUserWalletRepo.UpdateWallet(ctx, payeeID, payee.Wallet().Money())
	if err != nil {
		//c.updateUserWalletRepo.Rollback()
		return err
	}

	//c.updateUserWalletRepo.Commit()
	/**
	End Transaction
	*/

	return nil
}
