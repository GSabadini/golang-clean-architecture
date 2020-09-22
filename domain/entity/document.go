package entity

import "errors"

var (
	ErrInvalidTypeDocument = errors.New("invalid type document")
)

type TypeDocument string

type Document struct {
	//kind   TypeDocument
	Type   TypeDocument
	Number string
}

func NewDocument(t TypeDocument, number string) (*Document, error) {
	switch t {
	case CPF:
		return &Document{Type: CPF, Number: number}, nil
	case CNPJ:
		return &Document{Type: CNPJ, Number: number}, nil
	}

	return nil, ErrInvalidTypeDocument
}

const (
	CPF  TypeDocument = "CPF"
	CNPJ TypeDocument = "CNPJ"
)
