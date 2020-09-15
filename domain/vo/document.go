package vo

type TypeDocument string

const (
	RG   TypeDocument = "RG"
	CNPJ TypeDocument = "CNPJ"
)

type Document struct {
	Type   TypeDocument
	Number string
}
