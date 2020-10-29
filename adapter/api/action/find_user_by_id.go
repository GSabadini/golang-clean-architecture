package action

import (
	"errors"
	"net/http"

	"github.com/GSabadini/go-challenge/adapter/api/response"
	"github.com/GSabadini/go-challenge/adapter/logger"
	"github.com/GSabadini/go-challenge/domain/vo"
	"github.com/GSabadini/go-challenge/usecase"
)

type FindUserByIDAction struct {
	uc     usecase.FindUserByID
	log    logger.Logger
	logKey string
}

func NewFindUserByIDAction(uc usecase.FindUserByID, l logger.Logger) FindUserByIDAction {
	return FindUserByIDAction{
		uc:     uc,
		log:    l,
		logKey: "find_user_by_id",
	}
}

func (f FindUserByIDAction) Execute(w http.ResponseWriter, r *http.Request) {
	reqID := r.URL.Query().Get("user_id")
	if reqID == "" {
		err := errors.New("invalid parameter")
		f.log.WithFields(logger.Fields{
			"key":         f.logKey,
			"error":       err.Error(),
			"http_status": http.StatusBadRequest,
		}).Errorf("invalid parameter")

		response.NewError(err, http.StatusBadRequest).Send(w)
		return
	}

	ID, err := vo.NewUuid(reqID)
	if err != nil {
		err := errors.New("invalid uuid")
		f.log.WithFields(logger.Fields{
			"key":         f.logKey,
			"error":       err.Error(),
			"http_status": http.StatusBadRequest,
		}).Errorf("invalid uuid")

		response.NewError(err, http.StatusBadRequest).Send(w)
		return
	}

	output, err := f.uc.Execute(r.Context(), usecase.FindUserByIDInput{ID: ID})
	if err != nil {
		f.log.WithFields(logger.Fields{
			"key":         f.logKey,
			"error":       err.Error(),
			"http_status": http.StatusInternalServerError,
		}).Errorf("error fetching user by id")

		response.NewError(err, http.StatusInternalServerError).Send(w)
		return
	}

	f.log.WithFields(logger.Fields{
		"key":         f.logKey,
		"http_status": http.StatusOK,
	}).Infof("success when returning user by id")

	response.NewSuccess(output, http.StatusOK).Send(w)
}
