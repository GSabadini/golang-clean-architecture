package http

import (
	"net/http"
)

type (
	// Client is the http wrapper for the application
	Client struct {
		req *Request
	}
)

// NewClient returns a configured Client
func NewClient(r *Request) *Client {
	return &Client{r}
}

// Get executes a GET http request
func (c *Client) Get(url string) (*http.Response, error) {
	return c.req.Do(http.MethodGet, url, "application/json", nil)
}
