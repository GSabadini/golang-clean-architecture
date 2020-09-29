package vo

import (
	"errors"
)

var (
	ErrInvalidTypeDocument = errors.New("invalid type document")
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

	if !doc.validate() {
		return Document{}, ErrInvalidTypeDocument
	}

	return doc, nil
}

func (d Document) validate() bool {
	switch d.typeDoc {
	case CPF, CNPJ:
		return true
	}

	return false
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
