package http

import (
	"net/http"
)

type (
	// HTTPClient is the http wrapper for the application
	HTTPClient interface {
		HTTPGetter
	}

	// HTTPGetter holds fields and dependencies for executing an http GET request
	HTTPGetter interface {
		// Get executes a GET http request
		Get(url string) (*http.Response, error)
	}
)

type (
	stubHTTPGetter struct {
		res *http.Response
		err error
	}
)

func (h stubHTTPGetter) Get(_ string) (*http.Response, error) {
	return h.res, h.err
}
