package rapididentity

import (
	"bytes"
	"cmp"
	"context"
	"encoding/json"
	"fmt"
)

// Input for getting users.
type RunUserQueryInput struct {
	// The type of search to initiate.
	// The default is "advanced".
	SearchType string

	// The maximum amount of users to return.
	// The default is 1000.
	Limit int

	// The delegation ids of the delegations that
	// will be searched.
	DelegationIds []string

	// The user query to run.
	Query AuditReportQuery
}

// Run a user query.
//
//meta:operation POST /users
func (c *Client) RunUserQuery(ctx context.Context, params RunUserQueryInput) ([]User, error) {
	var output []User
	limit := cmp.Or(params.Limit, 1000)
	searchType := cmp.Or(params.SearchType, "advanced")

	url := fmt.Sprintf("%s/users?search=%s&limit=%d", c.baseEndpoint, searchType, limit)
	query, err := json.Marshal(params.Query)
	if err != nil {
		return nil, err
	}
	for _, field := range params.DelegationIds {
		url = fmt.Sprintf("%s&did=%s", url, field)
	}
	requestBody := bytes.NewBuffer(query)
	req, err := c.GenerateRequest(ctx, "POST", url, requestBody)
	if err != nil {
		return nil, err
	}
	req.Header.Add("Content-Type", "application/json")

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
		return nil, RapidIdentityError{
			Method:  req.Method,
			ReqUrl:  req.URL,
			Message: string(resBody),
			Reason:  err.Error(),
			Code:    res.StatusCode,
		}
	}

	return output, nil
}
