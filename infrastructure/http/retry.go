package http

import (
	"math/rand"
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
	}

	// Func is the function to be executed and eventually retried.
	Func func() error

	// HTTPFunc is the function to be executed and eventually retried.
	// The only difference from Func is that it expects an *http.Response on the first returning argument.
	HTTPFunc func() (*http.Response, error)
)

// Do wraps Func and returns *http.Response and error as returning arguments.
func (r Retry) Do(fn HTTPFunc) (*http.Response, error) {
	var res *http.Response

	err := retry(func() error {
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
