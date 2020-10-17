package action

import (
	"context"
	"fmt"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
	"time"

	"github.com/GSabadini/go-challenge/adapter/logger"
	"github.com/GSabadini/go-challenge/adapter/presenter"
	"github.com/GSabadini/go-challenge/domain/entity"
	"github.com/GSabadini/go-challenge/domain/vo"
	infralogger "github.com/GSabadini/go-challenge/infrastructure/logger"
	"github.com/GSabadini/go-challenge/usecase"
)

type stubFindUserByIDRepo struct {
	result entity.User
	err    error
}

func (f stubFindUserByIDRepo) FindByID(_ context.Context, _ vo.Uuid) (entity.User, error) {
	return f.result, f.err
}

func TestFindUserByIDAction_Execute(t *testing.T) {
	type fields struct {
		uc  usecase.FindUserByID
		log logger.Logger
	}
	type args struct {
		w  http.ResponseWriter
		r  *http.Request
		ID string
	}
	tests := []struct {
		name               string
		fields             fields
		args               args
		expectedBody       string
		expectedStatusCode int
	}{
		{
			name: "Find user by id success",
			fields: fields{
				uc: usecase.NewFindUserByIDInteractor(
					stubFindUserByIDRepo{
						result: entity.NewCustomUser(
							vo.NewUuidStaticTest(),
							vo.NewFullName("Custom user"),
							vo.NewEmailTest("test@testing.com"),
							vo.NewPassword("passw"),
							vo.NewDocumentTest(vo.CPF, "07091054954"),
							vo.NewWallet(vo.NewMoneyBRL(vo.NewAmountTest(100))),
							time.Time{},
						),
						err: nil,
					},
					presenter.NewFindUserByIDPresenter()),
				log: infralogger.Dummy{},
			},
			args: args{
				ID: vo.NewUuidStaticTest().Value(),
			},
			expectedBody:       `{"id":"0db298eb-c8e7-4829-84b7-c1036b4f0791","fullname":"Custom user","email":"test@testing.com","document":{"type":"CPF","value":"07091054954"},"wallet":{"currency":"BRL","amount":100},"roles":{"can_transfer":true},"type":"CUSTOM","created_at":"0001-01-01 00:00:00 +0000 UTC"}`,
			expectedStatusCode: http.StatusOK,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			uri := fmt.Sprintf("/users/%s", tt.args.ID)
			req, _ := http.NewRequest(http.MethodGet, uri, nil)

			q := req.URL.Query()
			q.Add("user_id", tt.args.ID)
			req.URL.RawQuery = q.Encode()

			var (
				w      = httptest.NewRecorder()
				action = NewFindUserByIDAction(tt.fields.uc, tt.fields.log)
			)

			action.Execute(w, req)

			if w.Code != tt.expectedStatusCode {
				t.Errorf(
					"[TestCase '%s'] O handler retornou um HTTP status code inesperado: retornado '%v' esperado '%v'",
					tt.name,
					w.Code,
					tt.expectedStatusCode,
				)
			}

			var result = strings.TrimSpace(w.Body.String())
			if !strings.EqualFold(result, tt.expectedBody) {
				t.Errorf(
					"[TestCase '%s'] Result: '%v' | Expected: '%v'",
					tt.name,
					result,
					tt.expectedBody,
				)
			}
		})
	}
}
