package rapididentity

import (
	"context"
	"encoding/json"
	"fmt"
)

// Output for retrieving Connect projects
type GetConnectProjectsOutput struct {
	// List of Connect projects.
	Projects []ConnectProject `json:"projects"`
}

type ConnectProject struct {
	// The name of the project.
	Name string `json:"name"`

	// The unique ID of the project.
	Id string `json:"id"`

	// The description of the project.
	Description string `json:"description"`

	// The group DN with administrator privileges
	// for the project.
	AdminGroupDN string `json:"adminGroupDN"`

	// The group DN with operator privileges
	// for the project.
	OperatorGroupDN string `json:"operatorGroupDN"`

	// The group DN with auditor privileges
	// for the project.
	AuditorGroupDN string `json:"auditorGroupDN"`

	// The RESTPoint configuration for the project.
	RestPoints RestPointConfig `json:"restPoints"`

	// The number of times the project has been updated.
	ChangeCount int `json:"changeCount"`

	// The timestamp of the last update to the project.
	ModifiedMs int64 `json:"modifiedMs"`

	// The DN of the user who made the last update.
	ModifiedBy string `json:"modifiedBy"`

	// The display name of the user who made the last update.
	ModifiedByName string `json:"modifiedByName"`
}

type RestPointConfig struct {
	// Whether RESTPoints are enabled.
	Disabled bool `json:"disabled"`

	// The default authentication used for the RESTPoints.
	AuthSpec AuthSpecConfig `json:"authSpec"`

	// The RESTPoints information.
	RestPoints []RestPoint `json:"restPoints"`
}

type AuthSpecConfig struct {
	// No authentication is used for the RESTPoint.
	Anonymous bool `json:"anonymous"`

	// OAuth1 authentication utilizing the OAuth1 consumer
	// keys in the Connect module.
	Oauth1 bool `json:"oauth1"`

	// Basic authentication utilizing username and password
	// of a RapidIdentity user with a minimum of Connect operator
	// privileges.
	Basic bool `json:"basic"`

	// Basic authentication utilizing OAuth1 consumer keys
	// in the Connect module.
	BasicWithOAuthKeys bool `json:"basicWithOAuthKeys"`
}

type RestPoint struct {
	// The unique ID of the RESTPoint.
	Id string `json:"id"`

	// The description of the RESTPoint.
	Description string `json:"description"`

	// The HTTP method of the RESTPoint.
	Method string `json:"method"`

	// Whether the RESTPoint is enabled.
	Disabled bool `json:"disabled"`

	// The HTTP path for the RESTPoint.
	Path string `json:"path"`

	// The Content-Type produced by the RESTPoint.
	Produces string `json:"produces"`

	// The action set called by the RESTPoint.
	ActionSet string `json:"actionSet"`

	// The arguments passed into the action set.
	// input parameters.
	ArgMap []RestPointArgMap `json:"argMap"`
}

type RestPointArgMap struct {
	// The HTTP source type. For example:
	// METHOD, QUERY_PARAM, etc...
	SourceType string `json:"sourceType"`

	// The type of the source.
	DestType string `json:"destType"`

	// The key of the source. For example:
	// if QUERY_PARAM this would be the name
	// of the query param.
	DestKey string `json:"destKey"`
}

// Retrieves a list of all Connect projects
//
//meta:operation GET /admin/connect/projects
func (c *Client) GetConnectProjects(ctx context.Context) (*GetConnectProjectsOutput, error) {
	var output GetConnectProjectsOutput

	url := fmt.Sprintf("%s/admin/connect/projects", c.baseEndpoint)

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
