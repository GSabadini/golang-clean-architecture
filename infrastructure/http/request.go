package http

import (
	"context"
	"fmt"
	"io"
	"net/http"
	"time"
)

var (
	defaultTimeout = time.Second * 5
)

type (
	// RequestOption is the request options.
	RequestOption func(*Request)

	// Request is the application http request.
	Request struct {
		client *http.Client
	}
)

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

	if r.client.Timeout == 0 {
		r.client.Timeout = defaultTimeout
	}

	ctx, cancel := context.WithTimeout(req.Context(), r.client.Timeout)
	defer cancel()

	req = req.WithContext(ctx)

	return r.client.Do(req)
}

// WithCircuitBreaker receives the http.RoundTripper implementation.
func WithCircuitBreaker(rt http.RoundTripper) RequestOption {
	return func(r *Request) {
		r.client.Transport = rt
	}
}

// WithRetry receives the http.RoundTripper implementation.
func WithRetry(rt http.RoundTripper) RequestOption {
	return func(r *Request) {
		r.client.Transport = rt
	}
}

func WithTimeout(t time.Duration) RequestOption {
	return func(r *Request) {
		r.client.Timeout = t
	}
}
