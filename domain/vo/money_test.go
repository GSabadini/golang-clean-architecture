package vo

import (
	"reflect"
	"testing"
)

func TestNewMoney(t *testing.T) {
	type args struct {
		currency Currency
		amount   Amount
	}
	tests := []struct {
		name string
		args args
		want Money
	}{
		{
			name: "Test new valid money",
			args: args{
				currency: Currency{
					value: BRL,
				},
				amount: Amount{value: 100},
			},
			want: Money{
				currency: Currency{
					value: BRL,
				},
				amount: Amount{value: 100},
			},
		},
		{
			name: "Test new valid money",
			args: args{
				currency: Currency{
					value: USD,
				},
				amount: Amount{value: 10},
			},
			want: Money{
				currency: Currency{
					value: USD,
				},
				amount: Amount{value: 10},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewMoney(tt.args.currency, tt.args.amount); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("[TestCase '%s'] Got: '%v' | Want: '%v'", tt.name, got, tt.want)
			}
		})
	}
}

func TestNewMoneyBRL(t *testing.T) {
	type args struct {
		amount Amount
	}
	tests := []struct {
		name string
		args args
		want Money
	}{
		{
			name: "Test create money with currency BRL",
			args: args{
				amount: Amount{value: 10},
			},
			want: Money{
				currency: Currency{
					value: BRL,
				},
				amount: Amount{value: 10},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewMoneyBRL(tt.args.amount); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("[TestCase '%s'] Got: '%v' | Want: '%v'", tt.name, got, tt.want)
			}
		})
	}
}

func TestMoney_Add(t *testing.T) {
	type fields struct {
		currency Currency
		amount   Amount
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
			name: "Test add money",
			fields: fields{
				currency: Currency{
					value: BRL,
				},
				amount: Amount{
					value: 100,
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
			name: "Test add money",
			fields: fields{
				currency: Currency{
					value: BRL,
				},
				amount: Amount{
					value: 10,
				},
			},
			args: args{
				amount: Amount{
					value: 10,
				},
			},
			want: Money{
				currency: Currency{
					value: BRL,
				},
				amount: Amount{
					value: 20,
				},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			m := NewMoney(tt.fields.currency, tt.fields.amount)
			if got := m.Add(tt.args.amount); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("[TestCase '%s'] Got: '%v' | Want: '%v'", tt.name, got, tt.want)
			}
		})
	}
}

func TestMoney_Sub(t *testing.T) {
	type fields struct {
		currency Currency
		amount   Amount
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
			name: "Test sub money",
			fields: fields{
				currency: Currency{
					value: BRL,
				},
				amount: Amount{
					value: 100,
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
		{
			name: "Test sub money",
			fields: fields{
				currency: Currency{
					value: BRL,
				},
				amount: Amount{
					value: 1050,
				},
			},
			args: args{
				amount: Amount{
					value: 1000,
				},
			},
			want: Money{
				currency: Currency{
					value: BRL,
				},
				amount: Amount{
					value: 50,
				},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			m := NewMoney(tt.fields.currency, tt.fields.amount)
			if got := m.Sub(tt.args.amount); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("[TestCase '%s'] Got: '%v' | Want: '%v'", tt.name, got, tt.want)
			}
		})
	}
}

func TestMoney_Equals(t *testing.T) {
	type fields struct {
		currency Currency
		amount   Amount
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
		{
			name: "Money value equals",
			fields: fields{
				currency: Currency{
					value: BRL,
				},
				amount: Amount{
					value: 0,
				},
			},
			args: args{
				value: Money{
					currency: Currency{
						value: BRL,
					},
					amount: Amount{
						value: 0,
					},
				},
			},
			want: true,
		},
		{
			name: "Money value equals",
			fields: fields{
				currency: Currency{
					value: USD,
				},
				amount: Amount{
					value: 2020,
				},
			},
			args: args{
				value: Money{
					currency: Currency{
						value: USD,
					},
					amount: Amount{
						value: 2020,
					},
				},
			},
			want: true,
		},
		{
			name: "Money value not equals",
			fields: fields{
				currency: Currency{
					value: USD,
				},
				amount: Amount{
					value: 2020,
				},
			},
			args: args{
				value: Money{
					currency: Currency{
						value: USD,
					},
					amount: Amount{
						value: 20,
					},
				},
			},
			want: false,
		},
		{
			name: "Money value not equals",
			fields: fields{
				currency: Currency{
					value: BRL,
				},
				amount: Amount{
					value: 2020,
				},
			},
			args: args{
				value: Money{
					currency: Currency{
						value: USD,
					},
					amount: Amount{
						value: 2020,
					},
				},
			},
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			m := NewMoney(tt.fields.currency, tt.fields.amount)
			if got := m.Equals(tt.args.value); got != tt.want {
				t.Errorf("[TestCase '%s'] Got: '%v' | Want: '%v'", tt.name, got, tt.want)
			}
		})
	}
}
