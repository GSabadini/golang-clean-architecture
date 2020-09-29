package vo

import (
	"errors"
	"strings"
)

var (
	ErrInvalidTypeUser = errors.New("invalid type user")
)

const (
	CUSTOM   TypeUser = "CUSTOM"
	MERCHANT TypeUser = "MERCHANT"
)

type TypeUser string

func (t TypeUser) String() string {
	return string(t)
}

func (t TypeUser) ToUpper() TypeUser {
	return TypeUser(strings.ToUpper(string(t)))
}
