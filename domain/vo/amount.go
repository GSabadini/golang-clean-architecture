package vo

import (
	"errors"
	"strconv"
)

var (
	errInvalidAmount = errors.New("invalid Amount")
)

type Amount struct {
	value int64
}

func NewAmount(value int64) (Amount, error) {
	if value < 0 {
		return Amount{}, errInvalidAmount
	}

	return Amount{value: value}, nil
}

// Value return value Amount
func (a Amount) Value() int64 {
	return a.value
}

// String returns string representation of the Amount
func (a Amount) String() string {
	return strconv.FormatInt(a.value, 10)
}

// Equals checks that two Amount are the same
func (a Amount) Equals(value Value) bool {
	o, ok := value.(Amount)
	return ok && a.value == o.value
}

func NewAmountTest(value int64) Amount {
	return Amount{value: value}
}
