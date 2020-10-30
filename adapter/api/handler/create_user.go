package handler

import (
	"encoding/json"
	"net/http"
	"time"

	"github.com/GSabadini/go-challenge/adapter/api/response"
	"github.com/GSabadini/go-challenge/adapter/logger"
	"github.com/GSabadini/go-challenge/domain/vo"
	"github.com/GSabadini/go-challenge/usecase"
	"github.com/google/uuid"
)

type (
	// Request data
	CreateUserRequest struct {
		FullName string
		Email    string
		Password string
		Document CreateUserDocumentRequest
		Wallet   CreateUserWalletRequest
		Type     string
	}

	// Request data
	CreateUserDocumentRequest struct {
		Type  string
		Value string
	}

	// Request data
	CreateUserWalletRequest struct {
		Currency string
		Amount   int64
	}

	// CreateUserHandler defines the dependencies of the HTTP handler for the use case
	CreateUserHandler struct {
		uc     usecase.CreateUserUseCase
		log    logger.Logger
		logKey string
	}
)

// NewCreateUserHandler creates new CreateUserHandler with its dependencies
func NewCreateUserHandler(uc usecase.CreateUserUseCase, l logger.Logger) CreateUserHandler {
	return CreateUserHandler{
		uc:     uc,
		log:    l,
		logKey: "create_user",
	}
}

// Handle handles http request
func (c CreateUserHandler) Handle(w http.ResponseWriter, r *http.Request) {
	var reqData CreateUserRequest
	if err := json.NewDecoder(r.Body).Decode(&reqData); err != nil {
		c.log.WithFields(logger.Fields{
			"key":         c.logKey,
			"error":       err.Error(),
			"http_status": http.StatusBadRequest,
		}).Errorf("failed to marshal message")

		response.NewError(err, http.StatusBadRequest).Send(w)
		return
	}
	defer r.Body.Close()

	input, err := NewCreateUserInput(reqData)
	if err != nil {
		c.log.WithFields(logger.Fields{
			"key":         c.logKey,
			"error":       err.Error(),
			"http_status": http.StatusBadRequest,
		}).Errorf("failed to data")

		response.NewError(err, http.StatusBadRequest).Send(w)
		return
	}

	output, err := c.uc.Execute(r.Context(), input)
	if err != nil {
		c.log.WithFields(logger.Fields{
			"key":         c.logKey,
			"error":       err.Error(),
			"http_status": http.StatusInternalServerError,
		}).Errorf("error when creating a new user")

		response.NewError(err, http.StatusInternalServerError).Send(w)
		return
	}

	c.log.WithFields(logger.Fields{
		"key":         c.logKey,
		"http_status": http.StatusCreated,
	}).Infof("success creating user")

	response.NewSuccess(output, http.StatusCreated).Send(w)
}

func NewCreateUserInput(i CreateUserRequest) (usecase.CreateUserInput, error) {
	id, err := vo.NewUuid(uuid.New().String())
	doc, err := vo.NewDocument(vo.TypeDocument(i.Document.Type), i.Document.Value)
	email, err := vo.NewEmail(i.Email)
	amount, err := vo.NewAmount(i.Wallet.Amount)
	wallet := vo.NewWallet(vo.NewMoneyBRL(amount))
	typeUser, err := vo.NewTypeUser(i.Type)
	if err != nil {
		return usecase.CreateUserInput{}, err
	}

	return usecase.CreateUserInput{
		ID:        id,
		FullName:  vo.NewFullName(i.FullName),
		Document:  doc,
		Email:     email,
		Password:  vo.NewPassword(i.Password),
		Wallet:    wallet,
		Type:      typeUser,
		CreatedAt: time.Now(),
	}, nil
}
