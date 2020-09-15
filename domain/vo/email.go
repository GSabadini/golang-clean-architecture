package vo

import (
	"errors"
	"regexp"
)

var (
	// ErrInvalidEmail return invalid Email
	ErrInvalidEmail = errors.New("invalid email")

	rxEmail = regexp.MustCompile("^[a-zA-Z0-9.!#$%&'*+\\/=?^_`{|}~-]+@[a-zA-Z0-9](?:[a-zA-Z0-9-]{0,61}[a-zA-Z0-9])?(?:\\.[a-zA-Z0-9](?:[a-zA-Z0-9-]{0,61}[a-zA-Z0-9])?)*$")
)

// Email structure
type Email struct {
	value string
}

// NewEmail create new Email
func NewEmail(value string) (Email, error) {
	var email = Email{value: value}

	if !email.validate() {
		return Email{}, ErrInvalidEmail
	}

	return email, nil
}

func (e Email) validate() bool {
	return rxEmail.MatchString(e.value)
}

// Value return value Email
func (e Email) Value() string {
	return e.value
}

// String returns string representation of the Email
func (e Email) String() string {
	return string(e.value)
}

// Equals checks that two Email are the same
func (e Email) Equals(value Value) bool {
	o, ok := value.(Email)
	return ok && e.value == o.value
}
