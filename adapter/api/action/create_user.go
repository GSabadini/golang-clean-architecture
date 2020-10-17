package action

import (
	"encoding/json"
	"net/http"

	"github.com/GSabadini/go-challenge/adapter/api/response"
	"github.com/GSabadini/go-challenge/adapter/logger"
	"github.com/GSabadini/go-challenge/usecase"
)

type CreateUserAction struct {
	uc     usecase.CreateUserUseCase
	log    logger.Logger
	logKey string
}

func NewCreateUserAction(uc usecase.CreateUserUseCase, l logger.Logger) CreateUserAction {
	return CreateUserAction{
		uc:     uc,
		log:    l,
		logKey: "create_user",
	}
}

type ()

func (c CreateUserAction) Execute(w http.ResponseWriter, r *http.Request) {
	var input usecase.CreateUserInput

	if err := json.NewDecoder(r.Body).Decode(&input); err != nil {
		c.log.WithFields(logger.Fields{
			"key":         c.logKey,
			"error":       err.Error(),
			"http_status": http.StatusBadRequest,
		}).Errorf("failed to marshal message")

		response.NewError(err, http.StatusBadRequest).Send(w)
		return
	}
	defer r.Body.Close()

	//output, err := c.uc.Execute(r.Context(), user)
	//if err != nil {
	//	c.log.WithFields(logger.Fields{
	//		"key":   c.logKey,
	//		"error": err.Error(),
	//		"http_status": http.StatusInternalServerError,
	//	}).Errorf("error when creating a new user")
	//
	//	response.NewError(err, http.StatusInternalServerError).Send(w)
	//	return
	//}

	c.log.WithFields(logger.Fields{
		"key":         c.logKey,
		"http_status": http.StatusCreated,
	}).Infof("success creating user")

	response.NewSuccess(input, http.StatusCreated).Send(w)
}
