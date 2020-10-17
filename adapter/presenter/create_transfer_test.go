package presenter

import (
	"reflect"
	"testing"
	"time"

	"github.com/GSabadini/go-challenge/domain/entity"
	"github.com/GSabadini/go-challenge/domain/vo"
	"github.com/GSabadini/go-challenge/usecase"
)

func Test_createTransferPresenter_Output(t *testing.T) {
	type args struct {
		t entity.Transfer
	}
	tests := []struct {
		name string
		args args
		want usecase.CreateTransferOutput
	}{
		{
			name: "Create transfer output",
			args: args{
				t: entity.NewTransfer(
					vo.NewUuidStaticTest(),
					vo.NewUuidStaticTest(),
					vo.NewUuidStaticTest(),
					vo.NewMoneyBRL(vo.NewAmountTest(100)),
					time.Time{},
				),
			},
			want: usecase.CreateTransferOutput{
				ID:        "0db298eb-c8e7-4829-84b7-c1036b4f0791",
				PayerID:   "0db298eb-c8e7-4829-84b7-c1036b4f0791",
				PayeeID:   "0db298eb-c8e7-4829-84b7-c1036b4f0791",
				Value:     100,
				CreatedAt: time.Time{}.String(),
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := NewCreateTransferPresenter()
			if got := c.Output(tt.args.t); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("[TestCase '%s'] Got: '%+v' | Want: '%+v'", tt.name, got, tt.want)
			}
		})
	}
}
