package vo

import (
	"github.com/pkg/errors"
)

var (
	ErrInvalidDocument = errors.New("invalid document")
)

// Document structure
type Document struct {
	typeDoc TypeDocument
	value   string
}

// NewDocument create new Document
func NewDocument(t string, value string) (Document, error) {
	var doc = Document{
		typeDoc: TypeDocument(t),
		value:   value,
	}

	if err := doc.validate(); err != nil {
		return Document{}, errors.Wrap(ErrInvalidDocument, err.Error())
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
		cpf, err := NewCNPJ(d.value)
		if err != nil {
			return err
		}
		d.value = cpf.String()

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

// NewDocumentTest create new Document for testing
func NewDocumentTest(t TypeDocument, value string) Document {
	return Document{
		typeDoc: t,
		value:   value,
	}
}
