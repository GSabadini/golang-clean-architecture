package http

import (
	"context"
	"fmt"
	"io"
	"net/http"
	"time"
)

var (
	reqTimeout = time.Second * 5
)

type (
	// RequestOption is the request options.
	RequestOption func(*Request)

	// Request is the application http request.
	Request struct {
		client *http.Client
	}
)

// WithRoundTripper receives the http.RoundTripper implementation.
func WithRoundTripper(rt http.RoundTripper) RequestOption {
	return func(r *Request) {
		r.client.Transport = rt
	}
}

// NewRequest returns a new configured Request.
func NewRequest(opts ...RequestOption) *Request {
	r := &Request{client: new(http.Client)}
	for _, o := range opts {
		o(r)
	}
	return r
}

// Do is a convenient method for executing http requests.
func (r *Request) Do(method, url, contentType string, body io.Reader) (*http.Response, error) {
	req, err := http.NewRequest(method, url, body)
	if err != nil {
		return nil, fmt.Errorf("failed to create request %v: ", err)
	}

	req.Header.Set("Content-Type", contentType)

	ctx, cancel := context.WithTimeout(req.Context(), reqTimeout)
	defer cancel()

	req = req.WithContext(ctx)

	return r.client.Do(req)
}
