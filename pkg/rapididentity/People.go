package rapididentity

import (
	"bytes"
	"cmp"
	"context"
	"encoding/json"
	"fmt"
)

// Input for getting user delegations.
type GetDelegationsForUserInput struct {
	// The idautoID of the user to retrieve
	// delegations.
	UserId string `json:"userId" jsonschema:"The idautoID of the user to retrieve delegations."`
}

// Output of delegations and profiles for a user.
type GetDelegationsForUserOutput struct {
	// Full list of delegations and profiles for a user.
	// Only provides delegations and profiles that are acessible
	// by the invoking session.
	AggregatedDelegation AggregatedDelegation `json:"aggregatedDelegation" jsonschema:"Full list of delegations and profiles for a user. Only provides delegations and profiles that are acessible by the invoking session."`
}

type DelegationProfileList []DelegationProfile

func (dpl DelegationProfileList) MarshalJSON() ([]byte, error) {
	if dpl == nil {
		return []byte("[]"), nil
	}
	return json.Marshal([]DelegationProfile(dpl))
}

type AggregatedDelegation struct {
	// idautoID of the user.
	Id string `json:"id" jsonschema:"idautoID of the user."`

	// Base user information.
	User User `json:"user" jsonschema:"Base user information."`

	// Helpdesk questions for the user.
	HelpdeskQuestions StringList `json:"helpdeskQuestions" jsonschema:"Helpdesk questions for the user."`

	// Profiles and delegations for the user
	// that the invoking session has access.
	DelegationProfiles DelegationProfileList `json:"delegationProfiles" jsonschema:"Profiles and delegations for the user that the invoking session has access."`
}

type User struct {
	// idautoID of the user.
	Id string `json:"id" jsonschema:"idautoID of the user."`

	// Distinguished Name of the user.
	Dn string `json:"dn" jsonschema:"Distinguished Name of the user."`

	// Username of the user.
	Username string `json:"username" jsonschema:"Username of the user."`

	// First name of the user.
	FirstName string `json:"firstName" jsonschema:"First name of the user."`

	// Last name of the user.
	LastName string `json:"lastName" jsonschema:"Last name of the user."`

	// Email address of the user.
	Email string `json:"email" jsonschema:"Email address of the user."`

	// Email address of the user.
	Distinguisher string `json:"distinguisher" jsonschema:"Email address of the user."`

	// Image url of the user.
	ImageUrl string `json:"imageUrl" jsonschema:"Image url of the user."`

	// Mobile numbers for the user.
	MobileNumbers StringList `json:"mobileNumbers" jsonschema:"Mobile numbers for the user."`

	// Alternate email address for the user
	AlternateEmail string `json:"alternateEmail" jsonschema:"Alternate email address for the user"`
}

type DelegationProfile struct {
	// Delegations for the user. Only provides
	// delegations that the invoking session has
	// access.
	Delegation Delegation `json:"delegation" jsonschema:"Delegations for the user. Only provides delegations that the invoking session has access."`

	// Profiles for the user. Only provides
	// profiles that the invoking session has
	// access.
	Profile Profile `json:"profile" jsonschema:"Profiles for the user. Only provides profiles that the invoking session has access."`
}

type DelegationAttributeList []DelegationAttribute

func (dal DelegationAttributeList) MarshalJSON() ([]byte, error) {
	if dal == nil {
		return []byte("[]"), nil
	}
	return json.Marshal([]DelegationAttribute(dal))
}

type ActionList []Action

func (al ActionList) MarshalJSON() ([]byte, error) {
	if al == nil {
		return []byte("[]"), nil
	}
	return json.Marshal([]Action(al))
}

type Delegation struct {
	// ID of the delegation.
	Id string `json:"id" jsonschema:"ID of the delegation."`

	// Name of the delegation.
	Name string `json:"name" jsonschema:"Name of the delegation."`

	// Tyoe of delegation. Will be MY or CUSTOM.
	Type string `json:"type" jsonschema:"Tyoe of delegation. Will be MY or CUSTOM."`

	// The Base OU in RapidIdentity LDAP to start search
	// for LDAP Objects. Source meaning who has access
	// to view this delegation.
	SourceBaseDN string `json:"sourceBaseDN" jsonschema:"The Base OU in RapidIdentity LDAP to start search for LDAP Objects. Source meaning who has access to view this delegation."`

	// The Base OU in RapidIdentity LDAP to start search
	// for LDAP Objects. Target meaning what LDAP Objects
	// that show up in the delegation view.
	TargetBaseDN string `json:"targetBaseDN" jsonschema:"The Base OU in RapidIdentity LDAP to start search for LDAP Objects. Target meaning what LDAP Objects that show up in the delegation view."`

	// Attributes associated with the delegations.
	Attributes DelegationAttributeList `json:"attributes" jsonschema:"Attributes associated with the delegations."`

	// The image in the profile layout.
	LayoutImage string `json:"layoutImage" jsonschema:"The image in the profile layout."`

	// The attribute to user for the 1 position of the
	// profile layout.
	Layout1 string `json:"layout1" jsonschema:"The attribute to user for the 1 position of the profile layout."`

	// The attribute to user for the 2 position of the
	// profile layout.
	Layout2 string `json:"layout2" jsonschema:"The attribute to user for the 2 position of the profile layout."`

	// The attribute to user for the 3 position of the
	// profile layout.
	Layout3 string `json:"layout3" jsonschema:"The attribute to user for the 3 position of the profile layout."`

	// The actions associated with the delegation.
	Actions ActionList `json:"actions" jsonschema:"The actions associated with the delegation."`
}

type ProfileAttributeList []ProfileAttribute

func (pal ProfileAttributeList) MarshalJSON() ([]byte, error) {
	if pal == nil {
		return []byte("[]"), nil
	}
	return json.Marshal([]ProfileAttribute(pal))
}

type Profile struct {
	// The idautoID of the user.
	Id string `json:"id" jsonschema:"The idautoID of the user."`

	// The attributes and values associated with
	// the user's profiles.
	Attributes ProfileAttributeList `json:"attributes" jsonschema:"The attributes and values associated with the user's profiles."`
}

type DelegationAttribute struct {
	// The GalItem associated with the
	// delegation attribute.
	GalItem GalItem `json:"galItem" jsonschema:"The GalItem associated with the delegation attribute."`

	// The name of the attribute for the delegation.
	Name string `json:"name" jsonschema:"The name of the attribute for the delegation."`

	// Whether the delegation attribute is editable.
	Editable bool `json:"editable" jsonschema:"Whether the delegation attribute is editable."`

	// Whether the delegation attribute is in the table view
	// of a delegation view.
	ShowInList bool `json:"showInList" jsonschema:"Whether the delegation attribute is in the table view of a delegation view."`

	// Whether the delegation attribute is in the details
	// of a delegation view.
	ShowInDetails bool `json:"showInDetails" jsonschema:"Whether the delegation attribute is in the details of a delegation view."`

	// Whether or not the attribute is required in the delegation
	// view.
	Required bool `json:"required" jsonschema:"Whether or not the attribute is required in the delegation view."`
}

type GalItem struct {
	// The ID of the GalItem attribute.
	Id string `json:"id" jsonschema:"The ID of the GalItem attribute."`

	// The friendly name for the GalItem attribute.
	FriendlyName string `json:"friendlyName" jsonschema:"The friendly name for the GalItem attribute."`

	// Whether the GalItem attribute accepts multiple
	// values.
	AllowMultiValue bool `json:"allowMultiValue" jsonschema:"Whether the GalItem attribute accepts multiple values."`

	// The type of the GalItem attribute.
	Type string `json:"type" jsonschema:"The type of the GalItem attribute."`
}

type Action struct {
	// The ID of the action.
	Id string `json:"id" jsonschema:"The ID of the action."`

	// The action friendly name.
	Name string `json:"name" jsonschema:"The action friendly name."`

	// The action description.
	Description string `json:"description" jsonschema:"The action description."`
}

type ProfileAttribute struct {
	// The ID of the GalItem
	Id string `json:"id" jsonschema:"The ID of the GalItem"`

	// The GalItem friendly name.
	Name string `json:"name" jsonschema:"The GalItem friendly name."`

	// The value(s) of the attribute for the user
	Values StringList `json:"values" jsonschema:"The value(s) of the attribute for the user"`
}

// Input for retrieving a RapidIdentity user
// by DN or idautoID.
type GetUserByIdInput struct {
	// The DN or idautoID of the user
	// to retrieve.
	Id string `json:"id" jsonschema:"The DN or idautoID of the user to retrieve."`
}

type UserList []User

func (ul UserList) MarshalJSON() ([]byte, error) {
	if ul == nil {
		return []byte("[]"), nil
	}
	return json.Marshal([]User(ul))
}

// Input for getting users.
type RunUserQueryInput struct {
	// The type of search to initiate.
	// The default is "advanced".
	SearchType string `json:"searchType" jsonschema:"The type of search to initiate. The default is \"advanced\"."`

	// The maximum amount of users to return.
	// The default is 1000.
	Limit int `json:"limit" jsonschema:"The maximum amount of users to return. The default is 1000."`

	// The delegation ids of the delegations that
	// will be searched.
	DelegationIds StringList `json:"delegationIds" jsonschema:"The delegation ids of the delegations that will be searched."`

	// The user query to run.
	Query AuditReportQuery `json:"query" jsonschema:"The user query to run."`
}

// Gets all associated delegations and profiles for the user
// based on their idautoID. This method will only return
// delegations and profiles that the invoking session has access.
//
//meta:operation GET /profiles/aggregated/for/{userId}
func (c *Client) GetDelegationsForUser(ctx context.Context, params GetDelegationsForUserInput) (*GetDelegationsForUserOutput, error) {
	var output GetDelegationsForUserOutput

	url := fmt.Sprintf("%s/profiles/aggregated/for/%s", c.baseEndpoint, params.UserId)
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
		return nil, RapidIdentityError{
			Method:  req.Method,
			ReqUrl:  req.URL,
			Message: string(resBody),
			Reason:  err.Error(),
			Code:    res.StatusCode,
		}
	}

	return &output, nil
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

// Run a user query.
//
//meta:operation POST /users
func (c *Client) RunUserQuery(ctx context.Context, params RunUserQueryInput) (UserList, error) {
	var output UserList
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
