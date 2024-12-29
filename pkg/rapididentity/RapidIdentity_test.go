package rapididentity

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/http/httptest"
	"net/url"
	"testing"
)

const (
	baseUrlPath         = "/api/rest"
	mockServiceIdentity = "service_identity_key"
	mockUsername        = "rapididentity@example.com"
	mockPassword        = "donottellanyone"
)

func setup(t *testing.T) (*Client, *http.ServeMux) {
	t.Helper()
	mux := http.NewServeMux()
	server := httptest.NewServer(mux)
	baseUrl, _ := url.Parse(server.URL)
	client, _ := New(Options{
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

func TestNew(t *testing.T) {
	t.Parallel()
	mux := http.NewServeMux()
	server := httptest.NewServer(mux)
	mux.HandleFunc(baseUrlPath+"/sessions", func(w http.ResponseWriter, r *http.Request) {
		if r.Method == "POST" {
			testMethod(t, r, "POST")
			testHeader(t, r, "Content-Type", "application/json")
			var rapidIdentityUser RapidIdentityUser
			reqBody, err := io.ReadAll(r.Body)
			if err != nil {
				w.WriteHeader(http.StatusInternalServerError)
				fmt.Fprintln(w, err)
				return
			}
			err = json.Unmarshal(reqBody, &rapidIdentityUser)
			if err != nil {
				w.WriteHeader(http.StatusInternalServerError)
				fmt.Fprintln(w, err)
				return
			}
			w.WriteHeader(http.StatusOK)
			fmt.Fprintf(w,
				`{ 
					"session": {
						"id": "1234",
						"token": "4321",
						"user": {
							"email":"%s"
						}
					}
				}`,
				rapidIdentityUser.Username)
		} else {
			testMethod(t, r, "DELETE")
			testHeader(t, r, "Authorization", "Bearer 4321")
			w.WriteHeader(http.StatusNoContent)
		}
	})
	baseUrl, _ := url.Parse(server.URL)
	client, _ := New(Options{
		HTTPClient: &http.Client{},
		RapidIdentityUser: &RapidIdentityUser{
			Username: mockUsername,
			Password: mockPassword,
		},
		BaseUrl: baseUrl,
	})
	if client.session.Session.User.Email != mockUsername {
		t.Errorf("session payload username: got %s, want %s", client.session.Session.User.Email, mockUsername)
	}
	defer client.Close()
	t.Cleanup(server.Close)
}
