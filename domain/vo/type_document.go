package vo

import (
	"github.com/pkg/errors"
	"strings"
)

const (
	// Document types
	CPF  TypeDocument = "CPF"
	CNPJ TypeDocument = "CNPJ"
)

var (
	ErrInvalidTypeDocument = errors.New("invalid type document")
)

// TypeDocument define document types
type TypeDocument string

func (t TypeDocument) toUpper() TypeDocument {
	return TypeDocument(strings.ToUpper(string(t)))
}

func (t TypeDocument) String() string {
	return string(t)
}
