package rapididentity

import (
	"context"
	"encoding/json"
	"fmt"
)

// Input for retrieving Connect actions.
type GetConnectActionsInput struct {
	// The Connect project to filter by.
	// If empty, all projects will be searched.
	// For identifying the <Main> project use the
	// const variable MainProject.
	Project string

	// Whether to return full action details
	// or just metadata.
	MetaDataOnly bool
}

// Output for retrieving Connect actions
type GetConnectActionsOutput struct {
	// Query type name. For example "all".
	Name string

	// The list of actions.
	ActionDefs []ActionDef
}

// Retrieves actions from Connect.
//
//meta:operation GET /admin/connect/actions
func (c *Client) GetConnectActions(ctx context.Context, params GetConnectActionsInput) (*GetConnectActionsOutput, error) {
	var output GetConnectActionsOutput

	url := fmt.Sprintf("%s/admin/connect/actions?metaDataOnly=%t", c.baseEndpoint, params.MetaDataOnly)
	if params.Project != "" {
		if params.Project == MainProject {
			url = fmt.Sprintf("%s&project=", url)
		} else {
			url = fmt.Sprintf("%s&project=%s", url, params.Project)
		}
	}
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
