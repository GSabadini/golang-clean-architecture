package vo

import (
	"errors"
	"strings"
)

var (
	ErrInvalidTypeDocument = errors.New("invalid type document")
)

const (
	CPF  TypeDocument = "CPF"
	CNPJ TypeDocument = "CNPJ"
)

type TypeDocument string

func (t TypeDocument) toUpper() TypeDocument {
	return TypeDocument(strings.ToUpper(string(t)))
}

func (t TypeDocument) String() string {
	return string(t)
}

type Document struct {
	typeDoc TypeDocument
	value   string
}

func NewDocument(t string, value string) (Document, error) {
	switch t {
	case "CPF":
		return Document{typeDoc: CPF, value: value}, nil
	case "CNPJ":
		return Document{typeDoc: CNPJ, value: value}, nil
	}

	return Document{}, ErrInvalidTypeDocument
}

func (d Document) Value() string {
	return d.value
}

func (d Document) Type() TypeDocument {
	return d.typeDoc
}

func NewDocumentTest(t TypeDocument, value string) Document {
	return Document{
		typeDoc: TypeDocument(t),
		value:   value,
	}
}
