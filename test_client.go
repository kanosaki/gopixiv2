package pixiv

import (
	"fmt"
	"io"
	"net/http"
)

type MockClient struct {
	Delegate *http.Client
}

func (m *MockClient) Get(base, query string) (*http.Response, error) {
	u := fmt.Sprintf("http://localhost/%s", base)
	if query != "" {
		u += "?" + query
	}
	return m.Delegate.Get(u)
}

func (m *MockClient) Post(base, query string, body io.Reader) (*http.Response, error) {
	u := fmt.Sprintf("http://localhost/%s", base)
	if query != "" {
		u += "?" + query
	}
	return m.Delegate.Post(u, "", body)
}
