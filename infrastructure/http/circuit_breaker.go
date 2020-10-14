package http

import (
	"fmt"
	"net"
	"net/http"
	"time"
)

type (
	// Breaker is the http circuit breaker.
	Breaker interface {
		// Execute runs the given request if the circuit breaker is closed or half-open states.
		// An error is instantly returned when the circuit breaker is tripped.
		Execute(func() (interface{}, error)) (interface{}, error)
	}

	// CircuitBreaker is the application http transport.
	CircuitBreaker struct {
		rt      http.RoundTripper
		breaker Breaker
	}
)

// NewCircuitBreaker returns a new configured CircuitBreaker with circuit breaker.
func NewCircuitBreaker(cb Breaker) *CircuitBreaker {
	rt := &http.Transport{
		DialContext: (&net.Dialer{
			Timeout:   10 * time.Second,
			KeepAlive: 90 * time.Second,
			DualStack: true,
		}).DialContext,
	}

	return &CircuitBreaker{
		rt:      rt,
		breaker: cb,
	}
}

// RoundTrip decorates rt.RoundTrip with a circuit breaker.
// An error is returned if the circuit breaker rejects the request.
func (t *CircuitBreaker) RoundTrip(r *http.Request) (*http.Response, error) {
	res, err := t.breaker.Execute(func() (interface{}, error) {
		res, err := t.rt.RoundTrip(r)
		if err != nil {
			return nil, err
		}

		if res != nil && res.StatusCode >= http.StatusInternalServerError {
			return res, fmt.Errorf("http response error: %v", res.StatusCode)
		}

		return res, err
	})

	if err != nil {
		return nil, err
	}

	return res.(*http.Response), err
}
