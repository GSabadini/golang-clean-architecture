package vo

import (
	"strings"

	"github.com/pkg/errors"
)

const (
	CUSTOM   TypeUser = "CUSTOM"
	MERCHANT TypeUser = "MERCHANT"
)

var (
	ErrInvalidTypeUser = errors.New("invalid type user")

	ErrNotAllowedTypeUser = errors.New("not allowed user type")
)

type (
	// TypeUser define user types
	TypeUser string
)

func NewTypeUser(value string) (TypeUser, error) {
	switch TypeUser(strings.ToUpper(value)) {
	case MERCHANT, CUSTOM:
		return TypeUser(strings.ToUpper(value)), nil
	}

	return "", ErrInvalidTypeUser
}

// String returns string representation of the TypeUser
func (t TypeUser) String() string {
	return string(t)
}

// ToUpper
func (t TypeUser) ToUpper() TypeUser {
	return TypeUser(strings.ToUpper(string(t)))
}
