package main

import (
	"context"
	"encoding/json"
	"fmt"
	"time"

	"github.com/GSabadini/go-challenge/adapter/http"
	"github.com/GSabadini/go-challenge/adapter/presenter"
	"github.com/GSabadini/go-challenge/domain/vo"
	"github.com/GSabadini/go-challenge/infrastructure/db"
	"github.com/GSabadini/go-challenge/usecase"
)

func main() {
	email, err := vo.NewEmail("gfacina@hotmail.com")
	if err != nil {
		panic(err)
	}

	uuid, err := vo.NewUuid(vo.CreateUuid())
	if err != nil {
		//fmt.Println(err)
		panic(err)
	}
	fmt.Print(uuid)

	payer := usecase.NewCreateUserInput(
		vo.NewFullName("Gabriel Facina"),
		vo.NewDocumentTest("CPF", "1231231231"),
		email,
		vo.NewPassword("passw"),
		vo.NewWallet(vo.NewMoneyBRL(vo.NewAmountTest(100))),
		vo.CUSTOM,
	)
	if err != nil {
		fmt.Println(err)
	}

	payee := usecase.NewCreateUserInput(
		vo.NewFullName("Gabriel Facina"),
		vo.NewDocumentTest("CPF", "1231231231"),
		email,
		vo.NewPassword("passw"),
		vo.NewWallet(vo.NewMoneyBRL(vo.NewAmountTest(100))),
		vo.MERCHANT,
	)
	if err != nil {
		fmt.Println(err)
	}

	createUserRepo := &db.UserInMen{}
	createUserUC := usecase.NewCreateUserInteractor(createUserRepo, presenter.CreateUserPresenter{})
	u1, _ := createUserUC.Execute(
		context.TODO(),
		payer,
	)

	b2, err := json.Marshal(u1)
	if err != nil {
		panic(err)
	}
	fmt.Println(string(b2), "b222222222")

	u2, _ := createUserUC.Execute(
		context.TODO(),
		payee,
	)

	payerID, err := vo.NewUuid(u1.ID)
	if err != nil {
		fmt.Println(err)
	}

	payeeID, err := vo.NewUuid(u2.ID)
	if err != nil {
		fmt.Println(err)
	}

	createTransferRepo := &db.TransferInMen{}
	createTransfer := usecase.NewCreateTransferInteractor(
		createTransferRepo,
		createUserRepo,
		createUserRepo,
		presenter.CreateTransferPresenter{},
		http.Authorizer{},
		http.Notifier{},
	)

	transfer, err := createTransfer.Execute(
		context.TODO(),
		usecase.CreateTransferInput{
			ID:        uuid,
			PayerID:   payerID,
			PayeeID:   payeeID,
			Value:     vo.NewMoneyBRL(vo.NewAmountTest(100)),
			CreatedAt: time.Time{},
		})
	if err != nil {
		fmt.Println(err)
	}

	payerR, _ := createUserRepo.FindByID(context.TODO(), payerID)
	fmt.Println(" \n\npayer")
	fmt.Printf("%+v: ", payerR.Wallet())

	payeeR, _ := createUserRepo.FindByID(context.TODO(), payeeID)
	fmt.Println(" \n\npayee")
	fmt.Printf("%+v: ", payeeR.Wallet())

	fmt.Println("\n\ntransfer")
	fmt.Printf("%+v: ", transfer)
	b, _ := json.Marshal(transfer)
	fmt.Println(string(b))
}
