package mocks

import (
	"io"
	"net/http"
)

var (
	GetBuildRequestFunc func(method, url string, body io.Reader) (*http.Request, error)
)

type MockHttp struct {
	BuildRequestFunc func(method, url string, body io.Reader) (*http.Request, error)
}

func (m *MockHttp) BuildRequest(method, url string, body io.Reader) (*http.Request, error) {
	return GetBuildRequestFunc(method, url, body)
}
