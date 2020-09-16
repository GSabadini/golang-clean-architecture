package vo

import (
	"errors"
	"regexp"
)

var (
	// ErrInvalidUuid return invalid Uuid
	ErrInvalidUuid = errors.New("invalid uuid")

	rxUuidClear = regexp.MustCompile(`/[^0-9a-fA-F]/`)
	rxUuid      = regexp.MustCompile(`/^[0-9a-fA-F]{32}$/`)
)

type Uuid string

// Uuid structure
//type Uuid struct {
//	value string
//}

//// NewUuid create new Uuid
//func NewUuid(value string) (Uuid, error) {
//	var uuid = Uuid{value: value}
//
//	if !uuid.validate() {
//		return Uuid{}, ErrInvalidUuid
//	}
//
//	return uuid, nil
//}
//
//func (e Uuid) validate() bool {
//	return rxUuid.MatchString(e.value)
//}
//
//func (e Uuid) clear() bool {
//	//strings.Replace(rxUuidClear.Match(e.value), "", 1)
//	return rxUuid.MatchString(e.value)
//}
//
//// Value return value Uuid
//func (e Uuid) Value() string {
//	return e.value
//}
//
//// String returns string representation of the Uuid
//func (e Uuid) String() string {
//	return string(e.value)
//}
//
//// Equals checks that two Uuid are the same
//func (e Uuid) Equals(value Value) bool {
//	o, ok := value.(Uuid)
//	return ok && e.value == o.value
//}
