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
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			m := Money{
				currency: tt.fields.currency,
				amount:   tt.fields.amount,
			}
			if got := m.Add(tt.args.amount); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Add() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestMoney_Amount(t *testing.T) {
	type fields struct {
		currency Currency
		amount   Amount
	}
	tests := []struct {
		name   string
		fields fields
		want   Amount
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			m := Money{
				currency: tt.fields.currency,
				amount:   tt.fields.amount,
			}
			if got := m.Amount(); got != tt.want {
				t.Errorf("Amount() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestMoney_Currency(t *testing.T) {
	type fields struct {
		currency Currency
		amount   Amount
	}
	tests := []struct {
		name   string
		fields fields
		want   Currency
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			m := Money{
				currency: tt.fields.currency,
				amount:   tt.fields.amount,
			}
			if got := m.Currency(); got != tt.want {
				t.Errorf("Currency() = %v, want %v", got, tt.want)
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
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			m := Money{
				currency: tt.fields.currency,
				amount:   tt.fields.amount,
			}
			if got := m.Equals(tt.args.value); got != tt.want {
				t.Errorf("Equals() = %v, want %v", got, tt.want)
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
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			m := Money{
				currency: tt.fields.currency,
				amount:   tt.fields.amount,
			}
			if got := m.Sub(tt.args.amount); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Sub() = %v, want %v", got, tt.want)
			}
		})
	}
}
