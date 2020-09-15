package vo

import "fmt"

// Value object, includes method String() from fmt.Stringer
type Value interface {
	fmt.Stringer
	Equals(value Value) bool
}
