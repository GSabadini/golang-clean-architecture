package vo

import (
	"errors"
	"strconv"
)

var (
	errInvalidAmount = errors.New("invalid amount")
)

// Amount structure
type Amount struct {
	value int64
}

// NewAmount create new Amount
func NewAmount(value int64) (Amount, error) {
	var a = Amount{value: value}

	if !a.validate() {
		return Amount{}, errInvalidAmount
	}

	return a, nil
}

func (a Amount) validate() bool {
	return a.value >= 0
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

// NewAmountTest create new Amount for testing
func NewAmountTest(value int64) Amount {
	return Amount{value: value}
}
