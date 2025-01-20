package rapididentity

import (
	"context"
	"encoding/json"
	"fmt"
)

// Input for retrieving a RapidIdentity user
// by DN or idautoID.
type GetUserByIdInput struct {
	// The DN or idautoID of the user
	// to retrieve.
	Id string
}

// Retrieve a RapidIdentity user by DN or idautoID.
//
//meta:operation GET /admin/ldap/users/{dnOrId}
func (c *Client) GetUserById(ctx context.Context, params GetUserByIdInput) (*User, error) {
	var output User

	url := fmt.Sprintf("%s/admin/ldap/users/%s", c.baseEndpoint, params.Id)
	req, err := c.GenerateRequest(ctx, "GET", url, nil)
	if err != nil {
		return nil, err
	}

	res, err := c.httpClient.Do(req)
	if err != nil {
		return nil, err
	}
	resBody, err := c.ReceiveResponse(res)
	if err != nil {
		return nil, err
	}

	err = json.Unmarshal(resBody, &output)
	if err != nil {
		return nil, err
	}

	return &output, nil
}
