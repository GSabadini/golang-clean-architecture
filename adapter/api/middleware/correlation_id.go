package middleware

import (
	"context"
	"net/http"

	"github.com/google/uuid"
)

type CorrelationID struct{}

func NewCorrelationID() *CorrelationID {
	return &CorrelationID{}
}

func (c CorrelationID) Execute(w http.ResponseWriter, r *http.Request, next http.HandlerFunc) {
	ctx := r.Context()
	id := r.Header.Get("X-Correlation-Id")
	if id == "" {
		id = uuid.New().String()
	}

	ctx = context.WithValue(ctx, "correlation_id", id)
	r = r.WithContext(ctx)

	w.Header().Set("X-Correlation-Id", id)
	next.ServeHTTP(w, r)
}
