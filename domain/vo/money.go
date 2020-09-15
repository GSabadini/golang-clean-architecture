package vo

type Currency string

const (
	BRL Currency = "BRL"
)

type Money struct {
	Currency Currency
	Value    int64
}
