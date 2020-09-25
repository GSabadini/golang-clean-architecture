package vo

import (
	"errors"
	"regexp"
)

var (
	// ErrInvalidCNPJ return invalid CNPJ
	ErrInvalidCNPJ = errors.New("invalid cnpj")

	rxCNPJ = regexp.MustCompile(`^\d{2}\.?\d{3}\.?\d{3}\/?(:?\d{3}[1-9]|\d{2}[1-9]\d|\d[1-9]\d{2}|[1-9]\d{3})-?\d{2}$`)
)
