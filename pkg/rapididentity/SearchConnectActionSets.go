package rapididentity

import (
	"context"
	"encoding/json"
	"fmt"
)

// Input for searching action sets within
// a Connect Project
type SearchConnectActionSetsInput struct {
	// The text to search for. This can
	// also be a regex pattern if Regex is
	// set to true.
	SearchString string `json:"searchString" jsonschema:"The text to search for. This can also be a regex pattern if Regex is set to true."`

	// The Connect project to search within.
	// If empty, all projects will be searched.
	// For identifying the <Main> project use the
	// const variable MainProject.
	Project string `json:"project" jsonschema:"The Connect project to search within. If empty, all projects will be searched. For identifying the <Main> project use the const variable MainProject."`

	// Whether to match action set names.
	MatchAction bool `json:"matchAction" jsonschema:"Whether to match action set names."`

	// Whether to apply a case sensitive search.
	MatchCase bool `json:"matchCase" jsonschema:"Whether to apply a case sensitive search."`

	// Whether the search string contains regex.
	Regex bool `json:"regex" jsonschema:"Whether the search string contains regex."`
}

// Result from searching for action sets within a
// Connect project.
type SearchConnectActionSetsOutput struct {
	// The name of the search query.
	Name string `json:"name" jsonschema:"The name of the search query."`

	// The description of the search query.
	Description string `json:"description" jsonschema:"The description of the search query."`

	// The action set definition results.
	ActionDefs []ActionDef `json:"actionDefs" jsonschema:"The action set definition results."`

	// The http status code returned.
	HttpStatus int `json:"httpStatus" jsonschema:"The http status code returned."`
}

type ActionDef struct {
	// The action set ID.
	Id string `json:"id" jsonschema:"The action set ID."`

	// The action set version.
	Version int `json:"version" jsonschema:"The action set version."`

	// The project where the action set resides.
	Project string `json:"project" jsonschema:"The project where the action set resides."`

	// The name of the action set.
	Name string `json:"name" jsonschema:"The name of the action set."`

	// The category the action set is a part of.
	Category string `json:"category" jsonschema:"The category the action set is a part of."`

	// Whether the action set is built in or custom.
	BuiltIn bool `json:"builtIn" jsonschema:"Whether the action set is built in or custom."`

	// Whether the action set is a part of the community depot.
	Community bool `json:"community" jsonschema:"Whether the action set is a part of the community depot."`

	// Whether the action set returns a value.
	ReturnsValue bool `json:"returnsValue" jsonschema:"Whether the action set returns a value."`

	// The description of the action set.
	Description string `json:"description" jsonschema:"The description of the action set."`

	// Whether the action set is licensed or not.
	Unlicensed bool `json:"unlicensed" jsonschema:"Whether the action set is licensed or not."`

	// Whether the action set contains sensitive information.
	Sensitive bool `json:"sensitive" jsonschema:"Whether the action set contains sensitive information."`

	// The input parameters of the action set.
	ArgDefs []ArgDef `json:"argDefs" jsonschema:"The input parameters of the action set."`

	// The actions within the action set.
	Actions []ConnectAction `json:"actions" jsonschema:"The actions within the action set."`

	// Whether the action set is deprecated.
	Deprecated string `json:"deprecated" jsonschema:"Whether the action set is deprecated."`

	// The http status code of the return.
	HttpStatus int `json:"httpStatus" jsonschema:"The http status code of the return."`

	// The number of times the action set has been modified.
	ChangeCount int `json:"changeCount" jsonschema:"The number of times the action set has been modified."`

	// When the action set was last modified.
	ModifiedMs int64 `json:"modifiedMs" jsonschema:"When the action set was last modified."`

	// The idautoID of the user who modified the action set.
	ModifiedBy string `json:"modifiedBy" jsonschema:"The idautoID of the user who modified the action set."`

	// The display name of the user who modified the action set.
	ModifiedByName string `json:"modifiedByName" jsonschema:"The display name of the user who modified the action set."`
}

type ArgDef struct {
	// Whether the action set input parameter is optional.
	Optional bool `json:"optional" jsonschema:"Whether the action set input parameter is optional."`

	// The type of the input parameter
	Type string `json:"type" jsonschema:"The type of the input parameter"`

	// The name of the input parameter.
	Name string `json:"name" jsonschema:"The name of the input parameter."`

	// The description for the input parameter.
	Description string `json:"description" jsonschema:"The description for the input parameter."`
}

type ConnectAction struct {
	// The unique ID of the Connect action.
	Id string `json:"id" jsonschema:"The unique ID of the Connect action."`

	// The name of the Connect action.
	Name string `json:"name" jsonschema:"The name of the Connect action."`

	// Whether the action returns a value.
	OutputVar string `json:"outputVar" jsonschema:"Whether the action returns a value."`

	// Whether the action is disabled.
	Disabled bool `json:"disabled" jsonschema:"Whether the action is disabled."`

	// The project where the action resides.
	Project string `json:"project" jsonschema:"The project where the action resides."`

	// The input parameters for the action.
	Args []ConnectActionArg `json:"args" jsonschema:"The input parameters for the action."`
}

type ConnectActionArg struct {
	// The name of the input parameter.
	Name string `json:"name" jsonschema:"The name of the input parameter."`

	// The value of the input parameter.
	Value string `json:"value" jsonschema:"The value of the input parameter."`

	// The Connect actions.
	Actions []ConnectAction `json:"actions" jsonschema:"The Connect actions."`

	// The http status code returned.
	HttpStatus int `json:"httpStatus" jsonschema:"The http status code returned."`
}

// Searches for text within action sets in a project.
//
//meta:operation GET /admin/connect/search/actions
func (c *Client) SearchConnectActionSets(ctx context.Context, params SearchConnectActionSetsInput) (*SearchConnectActionSetsOutput, error) {
	var output SearchConnectActionSetsOutput

	url := fmt.Sprintf("%s/admin/connect/search/actions?searchString=%s&matchAction=%t&matchCase=%t&regex=%t", c.baseEndpoint, params.SearchString, params.MatchAction, params.MatchCase, params.Regex)
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
