package entity

import "errors"

var (
	ErrInvalidTypeDocument = errors.New("invalid type document")
)

const (
	CPF  TypeDocument = "CPF"
	CNPJ TypeDocument = "CNPJ"
)

type TypeDocument string

func (t TypeDocument) String() string {
	return string(t)
}

type Document struct {
	//kind   TypeDocument
	Type   TypeDocument
	Number string
}

func NewDocument(t TypeDocument, number string) (Document, error) {
	switch t {
	case CPF:
		return Document{Type: CPF, Number: number}, nil
	case CNPJ:
		return Document{Type: CNPJ, Number: number}, nil
	}

	return Document{}, ErrInvalidTypeDocument
}
