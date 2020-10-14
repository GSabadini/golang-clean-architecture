package http

import (
	"errors"
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
			function: testRetry,
		},
		{
			scenario: "do retry with fail",
			function: testRetryWithFail,
		},
	}

	for _, test := range tests {
		t.Run(test.scenario, func(t *testing.T) {
			test.function(t)
		})
	}
}

func testRetry(t *testing.T) {
	attemptsCount := 0
	fn := func() error {
		attemptsCount++
		return nil
	}

	err := retry(fn, 2, time.Second)
	if err != nil {
		t.Errorf("retry.Do returned wrong err value: got %v want %v", err, nil)
	}

	if attemptsCount != 1 {
		t.Errorf("attemptsCount returned wrong count value: got %v want %v", attemptsCount, 1)
	}
}

func testRetryWithFail(t *testing.T) {
	errFail := errors.New("fail")
	attemptsCount := 0
	fail := func() error {
		attemptsCount++
		return errFail
	}

	err := retry(fail, 2, time.Second)
	if err == nil {
		t.Errorf("retry.Do returned wrong err value: got %v want %v", err, errFail)
	}

	if attemptsCount != 2 {
		t.Errorf("attemptsCount returned wrong count value: got %v want %v", attemptsCount, 2)
	}
}
