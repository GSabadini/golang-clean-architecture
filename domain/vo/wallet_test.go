package vo

import (
	"reflect"
	"testing"
)

func TestWallet_Add(t *testing.T) {
	type fields struct {
		money Money
	}
	type args struct {
		amount Amount
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   Money
	}{
		{
			name: "Test add value in money",
			fields: fields{
				money: Money{
					currency: Currency{
						value: BRL,
					},
					amount: Amount{
						value: 100,
					},
				},
			},
			args: args{
				amount: Amount{
					value: 100,
				},
			},
			want: Money{
				currency: Currency{
					value: BRL,
				},
				amount: Amount{
					value: 200,
				},
			},
		},
		{
			name: "Test add value in money",
			fields: fields{
				money: Money{
					currency: Currency{
						value: BRL,
					},
					amount: Amount{
						value: 999,
					},
				},
			},
			args: args{
				amount: Amount{
					value: 250,
				},
			},
			want: Money{
				currency: Currency{
					value: BRL,
				},
				amount: Amount{
					value: 1249,
				},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			w := NewWallet(tt.fields.money)
			if got := w.Add(tt.args.amount); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("[TestCase '%s'] Got: '%v' | Want: '%v'", tt.name, got, tt.want)
			}
		})
	}
}

func TestWallet_Sub(t *testing.T) {
	type fields struct {
		money Money
	}
	type args struct {
		amount Amount
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   Money
	}{
		{
			name: "Test sub value in money",
			fields: fields{
				money: Money{
					currency: Currency{
						value: BRL,
					},
					amount: Amount{
						value: 200,
					},
				},
			},
			args: args{
				amount: Amount{
					value: 100,
				},
			},
			want: Money{
				currency: Currency{
					value: BRL,
				},
				amount: Amount{
					value: 100,
				},
			},
		},
		{
			name: "Test sub value in money",
			fields: fields{
				money: Money{
					currency: Currency{
						value: BRL,
					},
					amount: Amount{
						value: 100,
					},
				},
			},
			args: args{
				amount: Amount{
					value: 100,
				},
			},
			want: Money{
				currency: Currency{
					value: BRL,
				},
				amount: Amount{
					value: 0,
				},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			w := NewWallet(tt.fields.money)
			if got := w.Sub(tt.args.amount); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("[TestCase '%s'] Got: '%v' | Want: '%v'", tt.name, got, tt.want)
			}
		})
	}
}

// @todo
func TestWallet_Equals(t *testing.T) {
	type fields struct {
		money Money
	}
	type args struct {
		value Value
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   bool
	}{
		//{
		//	name: "Wallet value equals",
		//	fields: fields{
		//		money: Money{
		//			currency: Currency{
		//				value: BRL,
		//			},
		//			amount: Amount{
		//				value: 100,
		//			},
		//		},
		//	},
		//	args: args{
		//		value: Money{
		//			currency: Currency{
		//				value: BRL,
		//			},
		//			amount: Amount{
		//				value: 100,
		//			},
		//		},
		//	},
		//	want: true,
		//},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			w := NewWallet(tt.fields.money)
			if got := w.Equals(tt.args.value); got != tt.want {
				t.Errorf("[TestCase '%s'] Got: '%v' | Want: '%v'", tt.name, got, tt.want)
			}
		})
	}
}
