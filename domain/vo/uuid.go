package vo

import (
	"errors"
	"regexp"
)

var (
	ErrInvalidUuid = errors.New("invalid uuid")

	rxUuid = regexp.MustCompile(`[0-9a-fA-F]{8}\-[0-9a-fA-F]{4}\-[0-9a-fA-F]{4}\-[0-9a-fA-F]{4}\-[0-9a-fA-F]{12}`)
)

//Uuid structure
type Uuid struct {
	value string
}

// NewUuid create new Uuid
func NewUuid(value string) (Uuid, error) {
	var uuid = Uuid{value: value}

	if !uuid.validate() {
		return Uuid{}, ErrInvalidUuid
	}

	return uuid, nil
}

func (e Uuid) validate() bool {
	return rxUuid.MatchString(e.value)
}

// Value return value Uuid
func (e Uuid) Value() string {
	return e.value
}

// String returns string representation of the Uuid
func (e Uuid) String() string {
	return e.value
}

// Equals checks that two Uuid are the same
func (e Uuid) Equals(value Value) bool {
	o, ok := value.(Uuid)
	return ok && e.value == o.value
}

// NewUuidStaticTest create new Uuid static
func NewUuidStaticTest() Uuid {
	return Uuid{value: "0db298eb-c8e7-4829-84b7-c1036b4f0791"}
}
