// Package rapididentity provides an easy to use SDK
// for the Identity Automation RapidIdentity REST API
package rapididentity

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"time"
)

const (
	Version          = "v1.0.0"
	defaultUserAgent = "ri-sdk-go" + "/" + Version
)

// Configurable options for the RapidIdentity
// Client.
type Options struct {
	// The http client to use to make requests
	// to the RapidIdentity REST API
	HTTPClient *http.Client

	// The service identity key to use for authorization.
	// See https://help.rapididentity.com/docs/service-identities-in-rapididentity
	// for setting up Service Identities in RapidIdentity Connect
	ServiceIdentity string

	// The RapidIdentity user to
	// create a session. If using Service Identity
	// This can be left empty or nil.
	RapidIdentityUser *RapidIdentityUser

	// The rapididentity base host url.
	// For example https://portal.us001-rapididentity.com.
	// Do NOT add a trailing slash
	BaseUrl *url.URL

	// The user agent to used in requests.
	// The default is the ri-sdk-go user agent
	UserAgent string
}

// RapidIdentity username and password for
// generating a user session
type RapidIdentityUser struct {
	// RapidIdentity user username.
	Username string `json:"username"`

	// RapidIdentity user password.
	Password string `json:"password"`
}

// A RapidIdentity user session
type Session struct {
	// The session object.
	Session SessionInfo `json:"session"`

	// Whether a password update is required with the session.
	PasswordUpdateRequired bool `json:"passwordUpdateRequired"`

	// Number of logins remaining before user is locked out.
	GraceLoginsRemaining int `json:"graceLoginsRemaining"`
}
type SessionInfo struct {
	// The session ID.
	Id string `json:"id"`

	// The session token.
	Token string `json:"token"`

	// The session user information. If using a proxy session
	// this will be the proxied user.
	User User `json:"user"`

	// If using a proxy session this will be the user
	// who invoked the proxy.
	RealUser User `json:"realUser"`

	// The RapidIdentity roles associated with the user
	// This does not contain the groups the user is a member of.
	Roles []string `json:"roles"`

	// When the session was created.
	Created time.Time `json:"created"`

	// The Client IP Address used to create the session.
	CreatedClientIp string `json:"createdClientIp"`

	// The Host IP address used to create the session.
	CreatedHostIp string `json:"createdHostIp"`

	// The time the session was last used.
	LastUsed time.Time `json:"lastUsed"`

	// The Client IP Address that was last used with the session.
	LastUsedClientIp string `json:"lastUsedClientIp"`

	// The Host IP Address that was last used with the session.
	LastUsedHostIp string `json:"lastUsedHostIp"`

	// When the session was invalidated.
	Invalidated time.Time `json:"invalidated"`

	// Proxy data associated with the session.
	ProxyData ProxyData `json:"proxyData"`
}

type ProxyData struct {
	// The RapidIdentity roles of the user who initiated the proxy
	Permissions []string `json:"permissions"`
}

// Client to make RapidIdentity REST API Calls.
type Client struct {
	httpClient         *http.Client
	serviceIdentityKey string
	session            *Session
	userAgent          string
	baseEndpoint       string
}

// Generates a base RapidIdentity API request that
// includes authorization and other reused headers.
func (c *Client) GenerateRequest(method string, url string, body io.Reader) (*http.Request, error) {
	req, err := http.NewRequest(method, url, body)
	if err != nil {
		return nil, err
	}

	token := c.serviceIdentityKey
	if c.session != nil {
		token = c.session.Session.Token
	}
	req.Header.Add("Authorization", "Bearer "+token)
	req.Header.Add("UserAgent", c.userAgent)
	req.Header.Add("Accept", "application/json")

	return req, nil
}

// Handles the responses provided by the
// RapidIdentity REST API
func (c *Client) ReceiveResponse(res *http.Response) ([]byte, error) {
	defer res.Body.Close()
	resBody, err := io.ReadAll(res.Body)
	if err != nil {
		return nil, RapidIdentityError{
			Method:  res.Request.Method,
			ReqUrl:  res.Request.URL,
			Message: string(resBody),
			Reason:  err.Error(),
			Code:    res.StatusCode,
		}
	}

	if res.StatusCode >= 200 && res.StatusCode < 300 {
		return resBody, nil
	} else {
		return nil, RapidIdentityError{
			Method:  res.Request.Method,
			ReqUrl:  res.Request.URL,
			Message: string(resBody),
			Reason:  string(resBody),
			Code:    res.StatusCode,
		}
	}
}

// If user session is available the session
// is revoked.
func (c *Client) Close() error {
	if c.session != nil {
		url := fmt.Sprintf("%s/sessions", c.baseEndpoint)
		req, err := http.NewRequest("DELETE", url, nil)
		if err != nil {
			return err
		}
		req.Header.Add("Authorization", "Bearer "+c.session.Session.Token)
		req.Header.Add("User-Agent", c.userAgent)

		res, err := c.httpClient.Do(req)
		if err != nil {
			return err
		}

		if res.StatusCode >= 200 && res.StatusCode < 300 {
			c.session = nil
			return nil
		}

		resBody, err := io.ReadAll(res.Body)
		if err != nil {
			return RapidIdentityError{
				Method:  req.Method,
				ReqUrl:  req.URL,
				Message: string(resBody),
				Reason:  err.Error(),
				Code:    res.StatusCode,
			}
		}
		return RapidIdentityError{
			Method:  req.Method,
			ReqUrl:  req.URL,
			Message: string(resBody),
			Reason:  "Unknown",
			Code:    res.StatusCode,
		}
	}

	return nil
}

// Creates a new RapidIdentity Client
// with the provided options
func New(options Options) (*Client, error) {
	if options.UserAgent == "" {
		options.UserAgent = defaultUserAgent
	}

	c := &Client{
		serviceIdentityKey: options.ServiceIdentity,
		httpClient:         options.HTTPClient,
		userAgent:          options.UserAgent,
		baseEndpoint:       fmt.Sprintf("%s/api/rest", options.BaseUrl),
	}

	if options.RapidIdentityUser != nil {
		url := fmt.Sprintf("%s/sessions", c.baseEndpoint)
		rapidIdentityUser, err := json.Marshal(options.RapidIdentityUser)
		if err != nil {
			return nil, err
		}
		reqBody := bytes.NewBuffer(rapidIdentityUser)
		req, err := http.NewRequest("POST", url, reqBody)
		if err != nil {
			return nil, err
		}
		req.Header.Add("Content-Type", "application/json")
		req.Header.Add("User-Agent", c.userAgent)
		res, err := options.HTTPClient.Do(req)
		if err != nil {
			return nil, err
		}
		defer res.Body.Close()
		resBody, err := io.ReadAll(res.Body)
		if err != nil {
			return nil, RapidIdentityError{
				Method:  req.Method,
				ReqUrl:  req.URL,
				Message: string(resBody),
				Reason:  err.Error(),
				Code:    res.StatusCode,
			}
		}

		if res.StatusCode >= 200 && res.StatusCode < 300 {
			var session Session
			err = json.Unmarshal(resBody, &session)
			if err != nil {
				return nil, RapidIdentityError{
					Method:  req.Method,
					ReqUrl:  req.URL,
					Message: string(resBody),
					Reason:  err.Error(),
					Code:    res.StatusCode,
				}
			}
			c.session = &session
		} else {
			return nil, RapidIdentityError{
				Method:  req.Method,
				ReqUrl:  req.URL,
				Message: string(resBody),
				Reason:  "Unknown",
				Code:    res.StatusCode,
			}
		}

	}

	return c, nil
}

// Error message to be used for additional
// information for all endpoints
type RapidIdentityError struct {
	Method  string
	ReqUrl  *url.URL
	Message string
	Reason  string
	Code    int
}

func (re RapidIdentityError) Error() string {
	return re.Message
}