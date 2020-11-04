package vo

import (
	"github.com/pkg/errors"
)

const (
	// Document types
	CPF  TypeDocument = "CPF"
	CNPJ TypeDocument = "CNPJ"
)

var (
	ErrInvalidTypeDocument = errors.New("invalid type document")
)

type (
	// TypeDocument define document types
	TypeDocument string
)

// String returns string representation of the TypeDocument
func (t TypeDocument) String() string {
	return string(t)
}

var (
	ErrInvalidDocument = errors.New("invalid document")
)

// Document structure
type Document struct {
	typeDoc TypeDocument
	value   string
}

// NewDocument create new Document
func NewDocument(typeDoc TypeDocument, value string) (Document, error) {
	var doc = Document{
		typeDoc: typeDoc,
		value:   value,
	}

	if err := doc.validate(); err != nil {
		return Document{}, err
	}

	return doc, nil
}

func (d *Document) validate() error {
	switch d.typeDoc {
	case CPF:
		cpf, err := NewCPF(d.value)
		if err != nil {
			return err
		}
		d.value = cpf.String()

		return nil
	case CNPJ:
		cnpj, err := NewCNPJ(d.value)
		if err != nil {
			return err
		}
		d.value = cnpj.String()

		return nil
	}

	return ErrInvalidTypeDocument
}

// Value return value Document
func (d Document) Value() string {
	return d.value
}

// Type return type Document
func (d Document) Type() TypeDocument {
	return d.typeDoc
}

// Equals checks that two Document are the same
func (d Document) Equals(value Value) bool {
	o, ok := value.(Document)
	return ok && d.typeDoc == o.typeDoc && d.value == o.value
}

// NewDocumentTest create new Document for testing
func NewDocumentTest(t TypeDocument, value string) Document {
	return Document{
		typeDoc: t,
		value:   value,
	}
}
