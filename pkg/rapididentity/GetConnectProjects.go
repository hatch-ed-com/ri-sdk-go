package rapididentity

import (
	"context"
	"encoding/json"
	"fmt"
)

type ConnectProjectList []ConnectProject

func (cpl ConnectProjectList) MarshalJSON() ([]byte, error) {
	if cpl == nil {
		return []byte("[]"), nil
	}
	return json.Marshal([]ConnectProject(cpl))
}

// Output for retrieving Connect projects
type GetConnectProjectsOutput struct {
	// List of Connect projects.
	Projects ConnectProjectList `json:"projects" jsonschema:"List of Connect projects."`
}

type ConnectProject struct {
	// The name of the project.
	Name string `json:"name" jsonschema:"The name of the project."`

	// The unique ID of the project.
	Id string `json:"id" jsonschema:"The unique ID of the project."`

	// The description of the project.
	Description string `json:"description" jsonschema:"The description of the project."`

	// The group DN with administrator privileges
	// for the project.
	AdminGroupDN string `json:"adminGroupDN" jsonschema:"The group DN with administrator privileges for the project."`

	// The group DN with operator privileges
	// for the project.
	OperatorGroupDN string `json:"operatorGroupDN" jsonschema:"The group DN with operator privileges for the project."`

	// The group DN with auditor privileges
	// for the project.
	AuditorGroupDN string `json:"auditorGroupDN" jsonschema:"The group DN with auditor privileges for the project."`

	// The RESTPoint configuration for the project.
	RestPoints RestPointConfig `json:"restPoints" jsonschema:"The RESTPoint configuration for the project."`

	// The number of times the project has been updated.
	ChangeCount int `json:"changeCount" jsonschema:"The number of times the project has been updated."`

	// The timestamp of the last update to the project.
	ModifiedMs int64 `json:"modifiedMs" jsonschema:"The timestamp of the last update to the project."`

	// The DN of the user who made the last update.
	ModifiedBy string `json:"modifiedBy" jsonschema:"The DN of the user who made the last update."`

	// The display name of the user who made the last update.
	ModifiedByName string `json:"modifiedByName" jsonschema:"The display name of the user who made the last update."`
}

type RestPointList []RestPoint

func (rpl RestPointList) MarshalJSON() ([]byte, error) {
	if rpl == nil {
		return []byte("[]"), nil
	}
	return json.Marshal([]RestPoint(rpl))
}

type RestPointConfig struct {
	// Whether RESTPoints are enabled.
	Disabled bool `json:"disabled" jsonschema:"Whether RESTPoints are enabled."`

	// The default authentication used for the RESTPoints.
	AuthSpec AuthSpecConfig `json:"authSpec" jsonschema:"The default authentication used for the RESTPoints."`

	// The RESTPoints information.
	RestPoints RestPointList `json:"restPoints" jsonschema:"The RESTPoints information."`
}

type AuthSpecConfig struct {
	// No authentication is used for the RESTPoint.
	Anonymous bool `json:"anonymous" jsonschema:"No authentication is used for the RESTPoint."`

	// OAuth1 authentication utilizing the OAuth1 consumer
	// keys in the Connect module.
	Oauth1 bool `json:"oauth1" jsonschema:"OAuth1 authentication utilizing the OAuth1 consumer keys in the Connect module."`

	// Basic authentication utilizing username and password
	// of a RapidIdentity user with a minimum of Connect operator
	// privileges.
	Basic bool `json:"basic" jsonschema:"Basic authentication utilizing username and password of a RapidIdentity user with a minimum of Connect operator privileges."`

	// Basic authentication utilizing OAuth1 consumer keys
	// in the Connect module.
	BasicWithOAuthKeys bool `json:"basicWithOAuthKeys" jsonschema:"Basic authentication utilizing OAuth1 consumer keys in the Connect module."`
}

type RestPointArgMapList []RestPointArgMap

func (rpaml RestPointArgMapList) MarshalJSON() ([]byte, error) {
	if rpaml == nil {
		return []byte("[]"), nil
	}
	return json.Marshal([]RestPointArgMap(rpaml))
}

type RestPoint struct {
	// The unique ID of the RESTPoint.
	Id string `json:"id" jsonschema:"The unique ID of the RESTPoint."`

	// The description of the RESTPoint.
	Description string `json:"description" jsonschema:"The description of the RESTPoint."`

	// The HTTP method of the RESTPoint.
	Method string `json:"method" jsonschema:"The HTTP method of the RESTPoint."`

	// Whether the RESTPoint is enabled.
	Disabled bool `json:"disabled" jsonschema:"Whether the RESTPoint is enabled."`

	// The HTTP path for the RESTPoint.
	Path string `json:"path" jsonschema:"The HTTP path for the RESTPoint."`

	// The Content-Type produced by the RESTPoint.
	Produces string `json:"produces" jsonschema:"The Content-Type produced by the RESTPoint."`

	// The action set called by the RESTPoint.
	ActionSet string `json:"actionSet" jsonschema:"The action set called by the RESTPoint."`

	// The arguments passed into the action set.
	// input parameters.
	ArgMap RestPointArgMapList `json:"argMap" jsonschema:"The arguments passed into the action set. input parameters."`
}

type RestPointArgMap struct {
	// The HTTP source type. For example:
	// METHOD, QUERY_PARAM, etc...
	SourceType string `json:"sourceType" jsonschema:"The HTTP source type. For example: METHOD, QUERY_PARAM, etc..."`

	// The type of the source.
	DestType string `json:"destType" jsonschema:"The type of the source."`

	// The key of the source. For example:
	// if QUERY_PARAM this would be the name
	// of the query param.
	DestKey string `json:"destKey" jsonschema:"The key of the source. For example: if QUERY_PARAM this would be the name of the query param."`
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
