package gupshup

import (
	"io"
	"net/http"
)

type httpClient interface {
	Do(req *http.Request) (*http.Response, error)
}

type httpBuilder interface {
	BuildRequest(method, url string, body io.Reader) (*http.Request, error)
}

type httpRequest struct{}

func (r *httpRequest) BuildRequest(method, url string, body io.Reader) (*http.Request, error) {
	return http.NewRequest(method, url, body)
}
