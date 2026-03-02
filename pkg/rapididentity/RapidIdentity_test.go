package rapididentity

import (
	"bytes"
	"context"
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

func TestDoCustomRequest(t *testing.T) {
	t.Parallel()
	mux := http.NewServeMux()
	server := httptest.NewServer(mux)
	mux.HandleFunc(baseUrlPath+"/util/regex/v2/validate", func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, "POST")
		testHeader(t, r, "Content-Type", "application/json")
		var regExp struct {
			Value string `json:"value"`
		}
		reqBody, err := io.ReadAll(r.Body)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			fmt.Fprintln(w, err)
			return
		}
		err = json.Unmarshal(reqBody, &regExp)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			fmt.Fprintln(w, err)
			return
		}
		w.WriteHeader(http.StatusOK)
		fmt.Fprintf(w,
			`{ 
					"result": true,
					"message": "%s"
				}`,
			regExp.Value)
	})

	baseUrl, _ := url.Parse(server.URL)
	client, _ := New(Options{
		HTTPClient:      &http.Client{},
		BaseUrl:         baseUrl,
		ServiceIdentity: mockServiceIdentity,
	})
	defer func(client *Client) {
		err := client.Close()
		if err != nil {
			t.Errorf("error closing client: %s", err)
		}
	}(client)

	input := struct {
		Value string `json:"value"`
	}{
		Value: "a.+",
	}

	payload, err := json.Marshal(input)
	if err != nil {
		t.Errorf("error marshalling request body: %s", err)
	}

	payloadBuf := bytes.NewBuffer(payload)

	ctx := context.Background()
	res, err := client.DoCustomRequest(ctx, "POST", "util/regex/v2/validate", payloadBuf)
	if err != nil {
		t.Errorf("error making custom request: %s", err)
	}

	defer func(res *http.Response) {
		err := res.Body.Close()
		if err != nil {
			t.Errorf("error closing response body: %s", err)
		}
	}(res)

	resBody, err := io.ReadAll(res.Body)
	if err != nil {
		t.Errorf("error reading response body: %s", err)
	}

	var endpointResponse struct {
		Result  bool   `json:"result"`
		Message string `json:"message"`
	}

	err = json.Unmarshal(resBody, &endpointResponse)
	if err != nil {
		t.Errorf("error unmarshalling response body: %s", err)
	}

	got := endpointResponse.Message
	want := input.Value

	if got != want {
		t.Errorf("response message: got %s, want %s", got, want)
	}

	t.Cleanup(server.Close)
}

func TestDoCustomRequestWithHeaders(t *testing.T) {
	t.Parallel()

	t.Run("custom content type with body", func(t *testing.T) {
		t.Parallel()
		mux := http.NewServeMux()
		server := httptest.NewServer(mux)
		mux.HandleFunc(baseUrlPath+"/connect/actionSets/execute", func(w http.ResponseWriter, r *http.Request) {
			testMethod(t, r, "POST")
			testHeader(t, r, "Content-Type", "text/plain")
			reqBody, err := io.ReadAll(r.Body)
			if err != nil {
				w.WriteHeader(http.StatusInternalServerError)
				fmt.Fprintln(w, err)
				return
			}
			w.WriteHeader(http.StatusOK)
			fmt.Fprint(w, string(reqBody))
		})

		baseUrl, _ := url.Parse(server.URL)
		client, _ := New(Options{
			HTTPClient:      &http.Client{},
			BaseUrl:         baseUrl,
			ServiceIdentity: mockServiceIdentity,
		})
		defer func(client *Client) {
			err := client.Close()
			if err != nil {
				t.Errorf("error closing client: %s", err)
			}
		}(client)

		body := bytes.NewBufferString("plain text body")
		headers := http.Header{}
		headers.Set("Content-Type", "text/plain")

		ctx := context.Background()
		res, err := client.DoCustomRequestWithHeaders(ctx, "POST", "connect/actionSets/execute", headers, body)
		if err != nil {
			t.Errorf("error making custom request: %s", err)
		}

		defer func(res *http.Response) {
			err := res.Body.Close()
			if err != nil {
				t.Errorf("error closing response body: %s", err)
			}
		}(res)

		resBody, err := io.ReadAll(res.Body)
		if err != nil {
			t.Errorf("error reading response body: %s", err)
		}

		got := string(resBody)
		want := "plain text body"

		if got != want {
			t.Errorf("response body: got %s, want %s", got, want)
		}

		t.Cleanup(server.Close)
	})

	t.Run("default content type with body", func(t *testing.T) {
		t.Parallel()
		mux := http.NewServeMux()
		server := httptest.NewServer(mux)
		mux.HandleFunc(baseUrlPath+"/util/regex/v2/validate", func(w http.ResponseWriter, r *http.Request) {
			testMethod(t, r, "POST")
			testHeader(t, r, "Content-Type", "application/json")
			var regExp struct {
				Value string `json:"value"`
			}
			reqBody, err := io.ReadAll(r.Body)
			if err != nil {
				w.WriteHeader(http.StatusInternalServerError)
				fmt.Fprintln(w, err)
				return
			}
			err = json.Unmarshal(reqBody, &regExp)
			if err != nil {
				w.WriteHeader(http.StatusInternalServerError)
				fmt.Fprintln(w, err)
				return
			}
			w.WriteHeader(http.StatusOK)
			fmt.Fprintf(w,
				`{
					"result": true,
					"message": "%s"
				}`,
				regExp.Value)
		})

		baseUrl, _ := url.Parse(server.URL)
		client, _ := New(Options{
			HTTPClient:      &http.Client{},
			BaseUrl:         baseUrl,
			ServiceIdentity: mockServiceIdentity,
		})
		defer func(client *Client) {
			err := client.Close()
			if err != nil {
				t.Errorf("error closing client: %s", err)
			}
		}(client)

		input := struct {
			Value string `json:"value"`
		}{
			Value: "a.+",
		}

		payload, err := json.Marshal(input)
		if err != nil {
			t.Errorf("error marshalling request body: %s", err)
		}

		payloadBuf := bytes.NewBuffer(payload)

		ctx := context.Background()
		res, err := client.DoCustomRequestWithHeaders(ctx, "POST", "util/regex/v2/validate", nil, payloadBuf)
		if err != nil {
			t.Errorf("error making custom request: %s", err)
		}

		defer func(res *http.Response) {
			err := res.Body.Close()
			if err != nil {
				t.Errorf("error closing response body: %s", err)
			}
		}(res)

		resBody, err := io.ReadAll(res.Body)
		if err != nil {
			t.Errorf("error reading response body: %s", err)
		}

		var endpointResponse struct {
			Result  bool   `json:"result"`
			Message string `json:"message"`
		}

		err = json.Unmarshal(resBody, &endpointResponse)
		if err != nil {
			t.Errorf("error unmarshalling response body: %s", err)
		}

		got := endpointResponse.Message
		want := input.Value

		if got != want {
			t.Errorf("response message: got %s, want %s", got, want)
		}

		t.Cleanup(server.Close)
	})

	t.Run("custom accept header with no body", func(t *testing.T) {
		t.Parallel()
		mux := http.NewServeMux()
		server := httptest.NewServer(mux)
		mux.HandleFunc(baseUrlPath+"/admin/workflow/resources", func(w http.ResponseWriter, r *http.Request) {
			testMethod(t, r, "GET")
			testHeader(t, r, "Accept", "text/plain")
			w.WriteHeader(http.StatusOK)
			fmt.Fprint(w, "plain text response")
		})

		baseUrl, _ := url.Parse(server.URL)
		client, _ := New(Options{
			HTTPClient:      &http.Client{},
			BaseUrl:         baseUrl,
			ServiceIdentity: mockServiceIdentity,
		})
		defer func(client *Client) {
			err := client.Close()
			if err != nil {
				t.Errorf("error closing client: %s", err)
			}
		}(client)

		headers := http.Header{}
		headers.Set("Accept", "text/plain")

		ctx := context.Background()
		res, err := client.DoCustomRequestWithHeaders(ctx, "GET", "admin/workflow/resources", headers, nil)
		if err != nil {
			t.Errorf("error making custom request: %s", err)
		}

		defer func(res *http.Response) {
			err := res.Body.Close()
			if err != nil {
				t.Errorf("error closing response body: %s", err)
			}
		}(res)

		resBody, err := io.ReadAll(res.Body)
		if err != nil {
			t.Errorf("error reading response body: %s", err)
		}

		got := string(resBody)
		want := "plain text response"

		if got != want {
			t.Errorf("response body: got %s, want %s", got, want)
		}

		t.Cleanup(server.Close)
	})
}
