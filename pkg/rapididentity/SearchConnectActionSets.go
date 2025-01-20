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
	SearchString string `json:"searchString"`

	// The Connect project to search within.
	// If empty, all projects will be searched.
	// For identifying the <Main> project use the
	// const variable MainProject.
	Project string `json:"project"`

	// Whether to match action set names.
	MatchAction bool `json:"matchAction"`

	// Whether to apply a case sensitive search.
	MatchCase bool `json:"matchCase"`

	// Whether the search string contains regex.
	Regex bool `json:"regex"`
}

// Result from searching for action sets within a
// Connect project.
type SearchConnectActionSetsOutput struct {
	// The name of the search query.
	Name string `json:"name"`

	// The description of the search query.
	Description string `json:"description"`

	// The action set definition results.
	ActionDefs []ActionDef `json:"actionDefs"`

	// The http status code returned.
	HttpStatus int `json:"httpStatus"`
}

type ActionDef struct {
	// The action set ID.
	Id string `json:"id"`

	// The action set version.
	Version int `json:"version"`

	// The project where the action set resides.
	Project string `json:"project"`

	// The name of the action set.
	Name string `json:"name"`

	// The category the action set is a part of.
	Category string `json:"category"`

	// Whether the action set is built in or custom.
	BuiltIn bool `json:"builtIn"`

	// Whether the action set is a part of the community depot.
	Community bool `json:"community"`

	// Whether the action set returns a value.
	ReturnsValue bool `json:"returnsValue"`

	// The description of the action set.
	Description string `json:"description"`

	// Whether the action set is licensed or not.
	Unlicensed bool `json:"unlicensed"`

	// Whether the action set contains sensitive information.
	Sensitive bool `json:"sensitive"`

	// The input parameters of the action set.
	ArgDefs []ArgDef `json:"argDefs"`

	// The actions within the action set.
	Actions []ConnectAction `json:"actions"`

	// Whether the action set is deprecated.
	Deprecated string `json:"deprecated"`

	// The http status code of the return.
	HttpStatus int `json:"httpStatus"`

	// The number of times the action set has been modified.
	ChangeCount int `json:"changeCount"`

	// When the action set was last modified.
	ModifiedMs int64 `json:"modifiedMs"`

	// The idautoID of the user who modified the action set.
	ModifiedBy string `json:"modifiedBy"`

	// The display name of the user who modified the action set.
	ModifiedByName string `json:"modifiedByName"`
}

type ArgDef struct {
	// Whether the action set input parameter is optional.
	Optional bool `json:"optional"`

	// The type of the input parameter
	Type string `json:"type"`

	// The name of the input parameter.
	Name string `json:"name"`

	// The description for the input parameter.
	Description string `json:"description"`
}

type ConnectAction struct {
	// The unique ID of the Connect action.
	Id string `json:"id"`

	// The name of the Connect action.
	Name string `json:"name"`

	// Whether the action returns a value.
	OutputVar string `json:"outputVar"`

	// Whether the action is disabled.
	Disabled bool `json:"disabled"`

	// The project where the action resides.
	Project string `json:"project"`

	// The input parameters for the action.
	Args []ConnectActionArg `json:"args"`
}

type ConnectActionArg struct {
	// The name of the input parameter.
	Name string `json:"name"`

	// The value of the input parameter.
	Value string `json:"value"`

	// The Connect actions.
	Actions []ConnectAction `json:"actions"`

	// The http status code returned.
	HttpStatus int `json:"httpStatus"`
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
