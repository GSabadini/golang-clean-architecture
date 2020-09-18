package main

import (
	"context"
	"fmt"
	"github.com/GSabadini/go-challenge/adapter/db/repository"
	"github.com/GSabadini/go-challenge/adapter/http"
	"github.com/GSabadini/go-challenge/domain/entity"
	"github.com/GSabadini/go-challenge/domain/vo"
	"github.com/GSabadini/go-challenge/usecase"
	"time"
)

func main() {
	email, err := vo.NewEmail("gfacina@hotmail.com")
	if err != nil {
		fmt.Println(err)
	}

	payer, err := entity.NewUser(
		"0db298eb-c8e7-4829-84b7-c1036b4f0791",
		"Gabriel Facina",
		email,
		"passw",
		entity.Document{Type: entity.CPF, Number: "102476239"},
		entity.NewWallet(vo.NewMoneyBRL(100)),
		entity.CUSTOM,
		time.Now(),
	)
	if err != nil {
		fmt.Println(err)
	}

	payee, err := entity.NewUser(
		"0db298eb-c8e7-4829-84b7-c1036b4f0792",
		"Gabriel Facina",
		email,
		"passw",
		entity.Document{Type: entity.CPF, Number: "1231231231"},
		entity.NewWallet(vo.NewMoneyBRL(100)),
		entity.MERCHANT,
		time.Now(),
	)
	if err != nil {
		fmt.Println(err)
	}

	userRepo := &repository.UserInMen{}
	_ = userRepo.Save(
		context.TODO(),
		payer,
	)
	_ = userRepo.Save(
		context.TODO(),
		payee,
	)

	transferRepo := &repository.TransferInMen{}
	createTransfer := usecase.CreateTransferInteractor{
		TransferRepo:       transferRepo,
		UserRepo:           userRepo,
		ExternalAuthorizer: http.Authorizer{},
		Notifier:           http.Notifier{},
	}

	transfer, err := createTransfer.Execute(
		context.TODO(),
		usecase.TransferInput{
			ID:        "",
			PayerID:   payer.ID(),
			PayeeID:   payee.ID(),
			Value:     vo.NewMoneyBRL(100),
			CreatedAt: time.Time{},
		})
	if err != nil {
		fmt.Println(err)
	}

	payerR, _ := userRepo.FindByID(context.TODO(), payer.ID())
	fmt.Println(" \n\npayer")
	fmt.Printf("%+v: ", payerR.Wallet())

	payeeR, _ := userRepo.FindByID(context.TODO(), payee.ID())
	fmt.Println(" \n\npayee")
	fmt.Printf("%+v: ", payeeR.Wallet())

	fmt.Println("\n\ntransfer")
	fmt.Printf("%+v: ", transfer)
}
