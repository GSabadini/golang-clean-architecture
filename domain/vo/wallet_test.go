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
		// TODO: Add test cases.
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
		// TODO: Add test cases.
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
		// TODO: Add test cases.
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
