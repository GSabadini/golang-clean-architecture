package response

import (
	"encoding/json"
	"net/http"
)

// Error defines the structure of success for http responses
type Error struct {
	statusCode int
	Errors     []string `json:"errors"`
}

// NewError creates new Error
func NewError(err error, status int) *Error {
	return &Error{
		statusCode: status,
		Errors:     []string{err.Error()},
	}
}

func NewErrorMessage(messages []string, status int) *Error {
	return &Error{
		statusCode: status,
		Errors:     messages,
	}
}

// Send returns a response with JSON format
func (e Error) Send(w http.ResponseWriter) error {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(e.statusCode)
	return json.NewEncoder(w).Encode(e)
}
