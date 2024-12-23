// Package rapididentity provides an easy to use SDK
// for the Identity Automation RapidIdentity REST API
package rapididentity

import (
	"fmt"
	"io"
	"net/http"
)

// Configurable options for the RapidIdentity
// Client.
type Options struct {
	// The http client to use to make requests
	// to the RapidIdentity REST API
	HTTPClient http.Client

	// The service identity key to use for authorization.
	// See https://help.rapididentity.com/docs/service-identities-in-rapididentity
	// for setting up Service Identities in RapidIdentity Connect
	ServiceIdentity string

	// The rapididentity base host url.
	// For example https://portal.us001-rapididentity.com.
	// Do NOT add a trailing slash
	HostUrl string
}

// Client to make RapidIdentity REST API Calls.
type Client struct {
	options      Options
	baseEndpoint string
}

// Generates a base RapidIdentity API request that
// includes authorization and other reused headers.
func (c *Client) GenerateRequest(method string, url string, body io.Reader) (*http.Request, error) {
	req, err := http.NewRequest(method, url, body)
	if err != nil {
		return nil, err
	}

	req.Header.Add("Authorization", "Bearer "+c.options.ServiceIdentity)
	req.Header.Add("Accept", "application/json")

	return req, nil
}

// Handles the responses provided by the
// RapidIdentity REST API
func (c *Client) ReceiveResponse(res *http.Response) ([]byte, error) {
	defer res.Body.Close()
	resBody, err := io.ReadAll(res.Body)
	if err != nil {
		return nil, err
	}

	if res.StatusCode >= 200 && res.StatusCode < 300 {
		return resBody, nil
	} else {
		return nil, fmt.Errorf("status code %d. %s", res.StatusCode, resBody)
	}
}

// Creates a new RapidIdentity Client
// with the provided options
func New(options Options) *Client {
	baseEndpoint := fmt.Sprintf("%s/api/rest", options.HostUrl)
	return &Client{
		options:      options,
		baseEndpoint: baseEndpoint,
	}
}
