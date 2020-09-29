package vo

// Password structure
type Password struct {
	value string
}

// NewPassword create new Password
func NewPassword(value string) Password {
	return Password{
		value: value,
	}
}

// Value return value Password
func (p Password) Value() string {
	return p.value
}

// Equals checks that two Password are the same
func (p Password) Equals(value Value) bool {
	o, ok := value.(Password)
	return ok && p.value == o.value
}
