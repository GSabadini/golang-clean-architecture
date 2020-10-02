package http

import (
	"context"
	"fmt"
	"io"
	"net/http"
	"time"
)

var (
	defaulTimeout = time.Second * 5
)

type (
	// RequestOption is the request options.
	RequestOption func(*Request)

	// Request is the application http request.
	Request struct {
		client  *http.Client
		retry   Retry
		timeout time.Duration
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

	if r.timeout == 0 {
		r.timeout = defaulTimeout
	}

	ctx, cancel := context.WithTimeout(req.Context(), r.timeout)
	defer cancel()

	req = req.WithContext(ctx)

	if r.retry.attempts > 0 {
		fn := func() (*http.Response, error) {
			res, err := r.client.Do(req)
			if err != nil {
				return res, err
			}

			for _, statusCode := range r.retry.statusCodes {
				if res.StatusCode == statusCode {
					return nil, fmt.Errorf("failed to request: %v ", http.StatusText(statusCode))
				}
			}

			return res, err
		}

		return r.retry.Do(fn)
	}

	return r.client.Do(req)
}

// WithRoundTripper receives the http.RoundTripper implementation.
func WithRoundTripper(rt http.RoundTripper) RequestOption {
	return func(r *Request) {
		r.client.Transport = rt
	}
}

// WithRetry receives the http.RoundTripper implementation.
func WithRetry(attempts int, sleep time.Duration, statusCode []int) RequestOption {
	return func(r *Request) {
		r.retry = Retry{
			attempts:    attempts,
			sleep:       sleep,
			statusCodes: statusCode,
		}
	}
}

func WithTimeout(t time.Duration) RequestOption {
	return func(r *Request) {
		r.timeout = t
	}
}
