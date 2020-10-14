package http

import (
	"fmt"
	"math/rand"
	"net"
	"net/http"
	"time"
)

var (
	defaultSleep = 500 * time.Millisecond
)

type (
	// Retry is mechanism the application retry.
	Retry struct {
		attempts    int
		sleep       time.Duration
		statusCodes []int
		rt          http.RoundTripper
	}

	// Func is the function to be executed and eventually retried.
	Func func() error
)

// NewRetry returns a new configured CircuitBreaker with retry.
func NewRetry(attempts int, statusCode []int, sleep time.Duration) *Retry {
	rt := &http.Transport{
		DialContext: (&net.Dialer{
			Timeout:   10 * time.Second,
			KeepAlive: 90 * time.Second,
			DualStack: true,
		}).DialContext,
	}

	return &Retry{
		attempts:    attempts,
		sleep:       sleep,
		statusCodes: statusCode,
		rt:          rt,
	}
}

// RoundTrip decorates RoundTrip with a retry.
func (r *Retry) RoundTrip(req *http.Request) (*http.Response, error) {
	var res *http.Response
	var err error

	fn := func() (*http.Response, error) {
		res, err := r.rt.RoundTrip(req)
		if err != nil {
			return res, err
		}

		for _, statusCode := range r.statusCodes {
			if res.StatusCode == statusCode {
				return nil, fmt.Errorf("failed to request: %v ", http.StatusText(statusCode))
			}
		}

		return res, err
	}

	err = retry(func() error {
		var err error
		res, err = fn()
		return err
	}, r.attempts, r.sleep)

	return res, err
}

// retry runs the passed function until the number of attempts is reached.
// Whenever Func returns err it will sleep and Func will be executed again in a recursive fashion.
// The sleep value is slightly modified on every retry (exponential backoff) to prevent the thundering herd problem (https://en.wikipedia.org/wiki/Thundering_herd_problem).
// If no value is given to sleep it will defaults to 500ms.
func retry(fn Func, attempts int, sleep time.Duration) error {
	if sleep == 0 {
		sleep = defaultSleep
	}

	if err := fn(); err != nil {
		attempts--
		if attempts <= 0 {
			return err
		}

		// preventing thundering herd problem (https://en.wikipedia.org/wiki/Thundering_herd_problem)
		sleep += (time.Duration(rand.Int63n(int64(sleep)))) / 2
		time.Sleep(sleep)

		return retry(fn, attempts, 2*sleep)
	}

	return nil
}
