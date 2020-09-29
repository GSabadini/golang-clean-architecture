package vo

// FullName structure
type FullName struct {
	value string
}

// NewFullName create new FullName
func NewFullName(value string) FullName {
	return FullName{
		value: value,
	}
}

// Value return value FullName
func (p FullName) Value() string {
	return p.value
}

// Equals checks that two FullName are the same
func (p FullName) Equals(value Value) bool {
	o, ok := value.(FullName)
	return ok && p.value == o.value
}
