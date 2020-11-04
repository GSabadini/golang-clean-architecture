package http

import (
	"bytes"
	"context"
	"errors"
	"io/ioutil"
	"net/http"
	"testing"

	"github.com/GSabadini/go-challenge/adapter/queue"
	"github.com/GSabadini/go-challenge/domain/entity"
	"github.com/GSabadini/go-challenge/infrastructure/logger"
)

type spyProducer struct {
	invoked bool
	err     error
}

func (s *spyProducer) Publish(_ []byte) error {
	s.invoked = true

	return s.err
}

func TestNotifier_Notify(t *testing.T) {
	type fields struct {
		client   HTTPGetter
		producer queue.Producer
	}
	type args struct {
		t entity.Transfer
	}
	tests := []struct {
		name             string
		fields           fields
		args             args
		publishIsInvoked bool
		publishErr       error
	}{
		{
			name: "Test notify success",
			fields: fields{
				client: httpGetterStub{
					res: &http.Response{
						Body: ioutil.NopCloser(
							bytes.NewReader([]byte(`{"message":"Enviado"}`)),
						),
					},
					err: nil,
				},
			},
			args: args{
				t: entity.Transfer{},
			},
			publishIsInvoked: false,
			publishErr:       nil,
		},
		{
			name: "Test notify error response",
			fields: fields{
				client: httpGetterStub{
					res: &http.Response{
						Body: ioutil.NopCloser(
							bytes.NewReader([]byte(`{"message":"error"}`)),
						),
					},
					err: nil,
				},
			},
			args: args{
				t: entity.Transfer{},
			},
			publishIsInvoked: true,
			publishErr:       nil,
		},
		{
			name: "Test notify client error",
			fields: fields{
				client: httpGetterStub{
					res: &http.Response{},
					err: errors.New("failure client"),
				},
			},
			args: args{
				t: entity.Transfer{},
			},
			publishIsInvoked: true,
			publishErr:       nil,
		},
		{
			name: "Test notify publish error",
			fields: fields{
				client: httpGetterStub{
					res: &http.Response{},
					err: errors.New("failure client"),
				},
			},
			args: args{
				t: entity.Transfer{},
			},
			publishIsInvoked: true,
			publishErr:       errors.New("failed publish"),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			spyProducer := &spyProducer{
				err: tt.publishErr,
			}

			n := NewNotifier(tt.fields.client, spyProducer, logger.Dummy{})
			n.Notify(context.TODO(), tt.args.t)

			if tt.publishIsInvoked != spyProducer.invoked {
				t.Errorf("[TestCase '%s'] Got: '%v' | Want: '%v'",
					tt.name,
					tt.publishIsInvoked,
					spyProducer.invoked,
				)
				t.Errorf("Expected to call 'Search' in 'Find', but it wasn't.")
			}
		})
	}
}
