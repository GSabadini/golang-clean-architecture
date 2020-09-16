package entity

import (
	"errors"
	"strings"
)

var (
	ErrInvalidTypeUser = errors.New("invalid type user")
)

type TypeUser string

func (t TypeUser) toUpper() TypeUser {
	return TypeUser(strings.ToUpper(string(t)))
}

const (
	Custom   TypeUser = "CUSTOM"
	Merchant TypeUser = "MERCHANT"
)
