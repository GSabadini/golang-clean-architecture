package middleware

import (
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
)

func TestCorrelationID_Execute(t *testing.T) {
	tests := []struct {
		name    string
		headers map[string]string
		want    string
	}{
		{
			name:    "Define correlation id",
			headers: map[string]string{"X-Correlation-Id": "f9882930-1914-47d7-8b58-18bff092e081"},
			want:    "f9882930-1914-47d7-8b58-18bff092e081",
		},
		{
			name: "Auto generated correlation id",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			req, err := http.NewRequest(http.MethodGet, "/middleware", nil)
			if err != nil {
				t.Fatal(err)
			}

			if tt.headers != nil {
				req.Header.Set("X-Correlation-Id", tt.headers["X-Correlation-Id"])
			}

			testHandler := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {})

			rr := httptest.NewRecorder()

			handler := NewCorrelationID().Execute(testHandler)
			handler.ServeHTTP(rr, req)

			gotHeader := rr.Header().Get("X-Correlation-Id")
			if gotHeader == "" {
				t.Errorf("[TestCase '%s'] Header X-Correlation-Id undefined", tt.name)
			}

			if (tt.want != "") && (!strings.EqualFold(gotHeader, tt.want)) {
				t.Errorf(
					"[TestCase '%s'] Got header: '%v' | Want header: '%v'",
					tt.name,
					gotHeader,
					tt.want,
				)
			}
		})
	}
}
