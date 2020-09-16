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

	payer := entity.NewUserFactory(
		"0db298eb-c8e7-4829-84b7-c1036b4f0791",
		"Gabriel Facina",
		email,
		"passw",
		vo.Document{Type: vo.RG, Number: "102476239"},
		vo.Money{Currency: vo.BRL, Value: 100},
		entity.Custom,
	)

	payee := entity.NewUserFactory(
		"0db298eb-c8e7-4829-84b7-c1036b4f0792",
		"Gabriel Facina",
		email,
		"passw",
		vo.Document{Type: vo.CNPJ, Number: "6239532017000000"},
		vo.Money{Currency: vo.BRL, Value: 100},
		entity.Merchant,
	)

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
			PayerID:        payer.ID,
			PayeeID:   payee.ID,
			Value:     	vo.Money{
				Currency: vo.BRL,
				Value:    100,
			},
			CreatedAt: time.Time{},
		})
	if err != nil {
		fmt.Println(err)
	}

	payerR, _ := userRepo.FindByID(context.TODO(), payer.ID)
	fmt.Println(payerR, "payer \n\n")

	payeeR, _ := userRepo.FindByID(context.TODO(), payee.ID)
	fmt.Println(payeeR, "payee \n\n")

	fmt.Println(transfer, "transfer")
}
