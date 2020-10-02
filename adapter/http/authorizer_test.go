package http

import (
	"bytes"
	"github.com/GSabadini/go-challenge/domain/entity"
	"io/ioutil"
	"net/http"
	"testing"
)

func TestAuthorizer_Authorized(t *testing.T) {
	type fields struct {
		client HTTPGetter
	}
	type args struct {
		transfer entity.Transfer
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    bool
		wantErr bool
	}{
		{
			name: "Test authorized success",
			fields: fields{
				client: HTTPGetterStub{
					res: &http.Response{
						Body: ioutil.NopCloser(
							bytes.NewReader([]byte(`{"message":"Autorizado"}`)),
						),
					},
					err: nil,
				},
			},
			args: args{
				transfer: entity.Transfer{},
			},
			want:    true,
			wantErr: false,
		},
		{
			name: "Test authorized error",
			fields: fields{
				client: HTTPGetterStub{
					res: &http.Response{},
					err: errAuthorizationDenied,
				},
			},
			args: args{
				transfer: entity.Transfer{},
			},
			want:    false,
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			a := NewAuthorizer(tt.fields.client)
			got, err := a.Authorized(tt.args.transfer)
			if (err != nil) != tt.wantErr {
				t.Errorf("[TestCase '%s'] Err: '%v' | WantErr: '%v'", tt.name, err, tt.wantErr)
				return
			}

			if got != tt.want {
				t.Errorf("[TestCase '%s'] Got: '%v' | Want: '%v'", tt.name, got, tt.want)
			}

		})
	}
}
