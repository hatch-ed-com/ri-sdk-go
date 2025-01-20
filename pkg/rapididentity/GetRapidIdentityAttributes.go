package rapididentity

import (
	"context"
	"encoding/json"
	"fmt"
)

// Retrieves RapidIdentity LDAP attributes.
//
//meta:operation GET /admin/ldap/schema/attributes
func (c *Client) GetRapidIdentityAttributes(ctx context.Context) ([]string, error) {
	var output []string

	url := fmt.Sprintf("%s/admin/ldap/schema/attributes", c.baseEndpoint)
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

	return output, nil
}
