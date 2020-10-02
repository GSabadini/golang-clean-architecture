package http

import (
	"errors"
	"fmt"
	"net/http"
	"testing"
	"time"
)

var (
	errFail = errors.New("fail")
)

func TestRetry(t *testing.T) {
	t.Parallel()

	tests := []struct {
		scenario string
		function func(*testing.T)
	}{
		{
			scenario: "do retry",
			function: testDoRetry,
		},
		{
			scenario: "do retry with three attempts",
			function: testDoRetryThreeAttempts,
		},
		{
			scenario: "do retry with fail",
			function: testDoRetryWithFail,
		},
	}

	for _, test := range tests {
		t.Run(test.scenario, func(t *testing.T) {
			test.function(t)
		})
	}
}

func testDoRetry(t *testing.T) {
	r := Retry{
		attempts:    2,
		sleep:       time.Millisecond,
		statusCodes: nil,
	}

	attemptsCount := 0
	fn := func() (*http.Response, error) {
		attemptsCount++
		return &http.Response{}, nil
	}

	_, err := r.Do(fn)
	if err != nil {
		t.Errorf("retry.Do returned wrong err value: got %v want %v", err, nil)
	}

	if attemptsCount != 1 {
		t.Errorf("attemptsCount returned wrong count value: got %v want %v", attemptsCount, 1)
	}
}

func testDoRetryThreeAttempts(t *testing.T) {
	r := Retry{
		attempts:    3,
		sleep:       time.Millisecond,
		statusCodes: []int{http.StatusInternalServerError},
	}

	attemptsCount := 0
	fn := func() (*http.Response, error) {
		attemptsCount++
		res := &http.Response{
			StatusCode: http.StatusInternalServerError,
		}

		if attemptsCount == 3 {
			res.StatusCode = http.StatusOK
		}

		for _, httpCode := range r.statusCodes {
			if res.StatusCode == httpCode {
				return nil, fmt.Errorf("failed to request: %v ", http.StatusText(httpCode))
			}
		}

		return res, nil
	}

	_, err := r.Do(fn)
	if err != nil {
		t.Errorf("retry.Do returned wrong err value: got %v want %v", err, nil)
	}

	if attemptsCount != 3 {
		t.Errorf("attemptsCount returned wrong count value: got %v want %v", attemptsCount, 3)
	}
}

func testDoRetryWithFail(t *testing.T) {
	r := Retry{
		attempts:    2,
		sleep:       time.Millisecond,
		statusCodes: nil,
	}

	attemptsCount := 0
	fn := func() (*http.Response, error) {
		attemptsCount++
		return &http.Response{}, errFail
	}

	_, err := r.Do(fn)
	if err == nil {
		t.Errorf("retry.Do returned wrong err value: got %v want %v", err, nil)
	}

	if attemptsCount != 2 {
		t.Errorf("attemptsCount returned wrong count value: got %v want %v", attemptsCount, 2)
	}
}
