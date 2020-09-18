package entity

import "errors"

var (
	ErrInvalidTypeDocument = errors.New("invalid type document")
)

type TypeDocument string

type Document struct {
	Type   TypeDocument
	Number string
}

const (
	CPF  TypeDocument = "CPF"
	CNPJ TypeDocument = "CNPJ"
)
