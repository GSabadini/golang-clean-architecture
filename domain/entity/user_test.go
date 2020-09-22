package entity

import (
	"github.com/GSabadini/go-challenge/domain/vo"
	"reflect"
	"testing"
	"time"
)

func TestNewUser(t *testing.T) {
	type args struct {
		ID        vo.Uuid
		fullName  vo.FullName
		email     vo.Email
		password  vo.Password
		document  Document
		wallet    *Wallet
		typeUser  TypeUser
		createdAt time.Time
	}
	tests := []struct {
		name    string
		args    args
		want    User
		wantErr error
	}{
		{
			name: "Test create custom user",
			args: args{
				ID:       "0db298eb-c8e7-4829-84b7-c1036b4f0791",
				fullName: "Test testing",
				email:    vo.Email{},
				password: "123",
				document: Document{
					Type:   CPF,
					Number: "07010965836",
				},
				wallet:    &Wallet{money: vo.NewMoneyBRL(100)},
				typeUser:  CUSTOM,
				createdAt: time.Time{},
			},
			want: User{
				id:       "0db298eb-c8e7-4829-84b7-c1036b4f0791",
				fullName: "Test testing",
				email:    vo.Email{},
				password: "123",
				document: Document{
					Type:   CPF,
					Number: "07010965836",
				},
				roles:     Roles{canTransfer: true},
				wallet:    &Wallet{money: vo.NewMoneyBRL(100)},
				typeUser:  CUSTOM,
				createdAt: time.Time{},
			},
		},
		{
			name: "Test create merchant user",
			args: args{
				ID:       "0db298eb-c8e7-4829-84b7-c1036b4f0791",
				fullName: "Test testing",
				email:    vo.Email{},
				password: "123",
				document: Document{
					Type:   CNPJ,
					Number: "07010965836",
				},
				wallet:    &Wallet{money: vo.NewMoneyBRL(100)},
				typeUser:  MERCHANT,
				createdAt: time.Time{},
			},
			want: User{
				id:       "0db298eb-c8e7-4829-84b7-c1036b4f0791",
				fullName: "Test testing",
				email:    vo.Email{},
				password: "123",
				document: Document{
					Type:   CNPJ,
					Number: "07010965836",
				},
				roles:     Roles{canTransfer: false},
				wallet:    &Wallet{money: vo.NewMoneyBRL(100)},
				typeUser:  MERCHANT,
				createdAt: time.Time{},
			},
		},
		{
			name: "Test create invalid user",
			args: args{
				ID:       "0db298eb-c8e7-4829-84b7-c1036b4f0791",
				fullName: "Test testing",
				email:    vo.Email{},
				password: "123",
				document: Document{
					Type:   CNPJ,
					Number: "07010965836",
				},
				wallet:    &Wallet{money: vo.NewMoneyBRL(100)},
				typeUser:  "INVALID",
				createdAt: time.Time{},
			},
			want:    User{},
			wantErr: ErrInvalidTypeUser,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := NewUser(
				tt.args.ID,
				tt.args.fullName,
				tt.args.email,
				tt.args.password,
				tt.args.document,
				tt.args.wallet,
				tt.args.typeUser,
				tt.args.createdAt,
			)
			if (err != nil) && (tt.wantErr != err) {
				t.Errorf("[TestCase '%s'] Err: '%v' | WantErr: '%v'", tt.name, err, tt.wantErr)
				return
			}

			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("[TestCase '%s'] Got: '%+v' | Want: '%+v'", tt.name, got, tt.want)
			}
		})
	}
}

func TestTypeUser_toUpper(t *testing.T) {
	tests := []struct {
		name string
		t    TypeUser
		want TypeUser
	}{
		{
			name: "Test upper custom type",
			t:    "cUstOm",
			want: CUSTOM,
		},
		{
			name: "Test upper merchant type",
			t:    "merchant",
			want: MERCHANT,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.t.toUpper(); got != tt.want {
				t.Errorf("[TestCase '%s'] Got: '%+v' | Want: '%+v'", tt.name, got, tt.want)
			}
		})
	}
}

func TestUser_CanTransfer(t *testing.T) {
	type args struct {
		id        vo.Uuid
		fullName  vo.FullName
		email     vo.Email
		password  vo.Password
		document  Document
		wallet    *Wallet
		typeUser  TypeUser
		roles     Roles
		createdAt time.Time
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "Test whether custom type user can transfer",
			args: args{
				id:        "0db298eb-c8e7-4829-84b7-c1036b4f0791",
				fullName:  "Test testing",
				email:     vo.Email{},
				password:  "123",
				document:  Document{},
				wallet:    nil,
				typeUser:  CUSTOM,
				roles:     Roles{},
				createdAt: time.Time{},
			},
			want: true,
		},
		{
			name: "Test whether merchant type user can transfer",
			args: args{
				id:        "0db298eb-c8e7-4829-84b7-c1036b4f0791",
				fullName:  "Test testing",
				email:     vo.Email{},
				password:  "123",
				document:  Document{},
				wallet:    nil,
				typeUser:  MERCHANT,
				roles:     Roles{},
				createdAt: time.Time{},
			},
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := NewUser(
				tt.args.id,
				tt.args.fullName,
				tt.args.email,
				tt.args.password,
				tt.args.document,
				tt.args.wallet,
				tt.args.typeUser,
				tt.args.createdAt,
			)
			if err != nil {
				t.Errorf("[TestCase '%s'] Err: '%v", tt.name, err)
				return
			}

			if got := got.CanTransfer(); got != tt.want {
				t.Errorf("[TestCase '%s'] Got: '%+v' | Want: '%+v'", tt.name, got, tt.want)
			}
		})
	}
}

func TestUser_Deposit(t *testing.T) {
	type argsUser struct {
		id        vo.Uuid
		fullName  vo.FullName
		email     vo.Email
		password  vo.Password
		document  Document
		wallet    *Wallet
		typeUser  TypeUser
		roles     Roles
		createdAt time.Time
	}
	type args struct {
		money vo.Money
	}
	tests := []struct {
		name     string
		argsUser argsUser
		args     args
		want     int64
	}{
		{
			name: "Test deposit 100",
			argsUser: argsUser{
				id:        "0db298eb-c8e7-4829-84b7-c1036b4f0791",
				fullName:  "Test testing",
				email:     vo.Email{},
				password:  "123",
				document:  Document{},
				wallet:    &Wallet{money: vo.NewMoneyBRL(100)},
				typeUser:  CUSTOM,
				roles:     Roles{},
				createdAt: time.Time{},
			},
			args: args{
				money: vo.NewMoneyBRL(100),
			},
			want: 200,
		},
		{
			name: "Test deposit 1000",
			argsUser: argsUser{
				id:        "0db298eb-c8e7-4829-84b7-c1036b4f0791",
				fullName:  "Test testing",
				email:     vo.Email{},
				password:  "123",
				document:  Document{},
				wallet:    &Wallet{money: vo.NewMoneyBRL(100)},
				typeUser:  CUSTOM,
				roles:     Roles{},
				createdAt: time.Time{},
			},
			args: args{
				money: vo.NewMoneyBRL(1000),
			},
			want: 1100,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := NewUser(
				tt.argsUser.id,
				tt.argsUser.fullName,
				tt.argsUser.email,
				tt.argsUser.password,
				tt.argsUser.document,
				tt.argsUser.wallet,
				tt.argsUser.typeUser,
				tt.argsUser.createdAt,
			)
			if err != nil {
				t.Errorf("[TestCase '%s'] Err: '%v", tt.name, err)
				return
			}

			got.Deposit(tt.args.money)

			if got.Wallet().Money().Amount() != tt.want {
				t.Errorf("[TestCase '%s'] Got: '%+v' | Want: '%+v'", tt.name, got.Wallet().Money().Amount(), tt.want)
			}
		})
	}
}

func TestUser_Withdraw(t *testing.T) {
	type argsUser struct {
		id        vo.Uuid
		fullName  vo.FullName
		email     vo.Email
		password  vo.Password
		document  Document
		wallet    *Wallet
		typeUser  TypeUser
		roles     Roles
		createdAt time.Time
	}
	type args struct {
		money vo.Money
	}
	tests := []struct {
		name     string
		argsUser argsUser
		args     args
		want     int64
		wantErr  error
	}{
		{
			name: "Test withdraw 100",
			argsUser: argsUser{
				id:        "0db298eb-c8e7-4829-84b7-c1036b4f0791",
				fullName:  "Test testing",
				email:     vo.Email{},
				password:  "123",
				document:  Document{},
				wallet:    &Wallet{money: vo.NewMoneyBRL(100)},
				typeUser:  CUSTOM,
				roles:     Roles{},
				createdAt: time.Time{},
			},
			args: args{
				money: vo.NewMoneyBRL(100),
			},
			want: 0,
		},
		{
			name: "Test withdraw 50",
			argsUser: argsUser{
				id:        "0db298eb-c8e7-4829-84b7-c1036b4f0791",
				fullName:  "Test testing",
				email:     vo.Email{},
				password:  "123",
				document:  Document{},
				wallet:    &Wallet{money: vo.NewMoneyBRL(100)},
				typeUser:  CUSTOM,
				roles:     Roles{},
				createdAt: time.Time{},
			},
			args: args{
				money: vo.NewMoneyBRL(50),
			},
			want: 50,
		},
		{
			name: "Test withdraw insufficient balance",
			argsUser: argsUser{
				id:        "0db298eb-c8e7-4829-84b7-c1036b4f0791",
				fullName:  "Test testing",
				email:     vo.Email{},
				password:  "123",
				document:  Document{},
				wallet:    &Wallet{money: vo.NewMoneyBRL(100)},
				typeUser:  CUSTOM,
				roles:     Roles{},
				createdAt: time.Time{},
			},
			args: args{
				money: vo.NewMoneyBRL(1000),
			},
			wantErr: ErrInsufficientBalance,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := NewUser(
				tt.argsUser.id,
				tt.argsUser.fullName,
				tt.argsUser.email,
				tt.argsUser.password,
				tt.argsUser.document,
				tt.argsUser.wallet,
				tt.argsUser.typeUser,
				tt.argsUser.createdAt,
			)
			if err != nil {
				t.Errorf("[TestCase '%s'] Err: '%v", tt.name, err)
				return
			}

			err = got.Withdraw(tt.args.money)
			if (err != nil) && (tt.wantErr != err) {
				t.Errorf("[TestCase '%s'] Err: '%v' | WantErr: '%v'", tt.name, err, tt.wantErr)
				return
			}

			if (err == nil) && (got.Wallet().Money().Amount() != tt.want) {
				t.Errorf("[TestCase '%s'] Got: '%+v' | Want: '%+v'", tt.name, got.Wallet().Money().Amount(), tt.want)
			}
		})
	}
}
