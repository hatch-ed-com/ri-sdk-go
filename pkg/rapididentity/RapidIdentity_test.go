package rapididentity

import (
	"net/http"
	"net/http/httptest"
	"net/url"
	"testing"
)

const (
	baseUrlPath         = "/api/rest"
	mockServiceIdentity = "service_identity_key"
)

func setup(t *testing.T) (client *Client, mux *http.ServeMux) {
	t.Helper()
	mux = http.NewServeMux()
	server := httptest.NewServer(mux)
	baseUrl, _ := url.Parse(server.URL)
	client = New(Options{
		HTTPClient:      &http.Client{},
		ServiceIdentity: mockServiceIdentity,
		BaseUrl:         baseUrl,
	})

	t.Cleanup(server.Close)

	return client, mux
}

func testMethod(t *testing.T, r *http.Request, want string) {
	t.Helper()
	got := r.Method
	if got != want {
		t.Errorf("request method: %s, want %s", got, want)
	}
}

func testHeader(t *testing.T, r *http.Request, header string, want string) {
	t.Helper()
	if got := r.Header.Get(header); got != want {
		t.Errorf("request header %s value: %s, want %s", header, got, want)
	}
}

func testQueryParam(t *testing.T, r *http.Request, param string, want string) {
	t.Helper()
	if got := r.URL.Query().Get(param); got != want {
		t.Errorf("request query param %s value: %s, want %s", param, got, want)
	}
}
