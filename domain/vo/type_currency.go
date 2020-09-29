package vo

const (
	// Currency types
	BRL TypeCurrency = "BRL"
	USD TypeCurrency = "USD"
)

// TypeCurrency define currency types
type TypeCurrency string

func (tc TypeCurrency) String() string {
	return string(tc)
}
