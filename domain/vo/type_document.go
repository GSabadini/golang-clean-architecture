package vo

import "strings"

const (
	// Document types
	CPF  TypeDocument = "CPF"
	CNPJ TypeDocument = "CNPJ"
)

// TypeDocument define document types
type TypeDocument string

func (t TypeDocument) toUpper() TypeDocument {
	return TypeDocument(strings.ToUpper(string(t)))
}

func (t TypeDocument) String() string {
	return string(t)
}
