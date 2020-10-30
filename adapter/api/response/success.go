package response

import (
	"encoding/json"
	"net/http"
)

// Success defines the structure of success for http responses
type Success struct {
	statusCode int
	result     interface{}
}

// NewSuccess creates new Success
func NewSuccess(result interface{}, status int) Success {
	return Success{
		statusCode: status,
		result:     result,
	}
}

// Send returns a response with JSON format
func (r Success) Send(w http.ResponseWriter) error {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(r.statusCode)
	return json.NewEncoder(w).Encode(r.result)
}
