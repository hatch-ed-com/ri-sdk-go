package rapididentity

import (
	"encoding/json"
	"fmt"
)

// Input for getting user delegations.
type GetDelegationsForUserInput struct {
	// The idautoID of the user to retrieve
	// delegations.
	UserId string
}

// Output of delegations and profiles for a user.
type GetDelegationsForUserOutput struct {
	// Full list of delegations and profiles for a user.
	// Only provides delegations and profiles that are acessible
	// by the invoking session.
	AggregatedDelegation AggregatedDelegation `json:"aggregatedDelegation"`
}

type AggregatedDelegation struct {
	// idautoID of the user.
	Id string `json:"id"`

	// Base user information.
	User User `json:"user"`

	// Helpdesk questions for the user.
	HelpdeskQuestions []string `json:"helpdeskQuestions"`

	// Profiles and delegations for the user
	// that the invoking session has access.
	DelegationProfiles []DelegationProfile `json:"delegationProfiles"`
}

type User struct {
	// idautoID of the user.
	Id string `json:"id"`

	// Distinguished Name of the user.
	Dn string `json:"dn"`

	// Username of the user.
	Username string `json:"username"`

	// First name of the user.
	FirstName string `json:"firstName"`

	// Last name of the user.
	LastName string `json:"lastName"`

	// Email address of the user.
	Email string `json:"email"`

	// Email address of the user.
	Distinguisher string `json:"distinguisher"`

	// Image url of the user.
	ImageUrl string `json:"imageUrl"`

	// Mobile numbers for the user.
	MobileNumbers []string `json:"mobileNumbers"`

	// Alternate email address for the user
	AlternateEmail string `json:"alternateEmail"`
}

type DelegationProfile struct {
	// Delegations for the user. Only provides
	// delegations that the invoking session has
	// access.
	Delegation Delegation `json:"delegation"`

	// Profiles for the user. Only provides
	// profiles that the invoking session has
	// access.
	Profile Profile `json:"profile"`
}

type Delegation struct {
	// ID of the delegation.
	Id string `json:"id"`

	// Name of the delegation.
	Name string `json:"name"`

	// Tyoe of delegation. Will be MY or CUSTOM.
	Type string `json:"type"`

	// The Base OU in RapidIdentity LDAP to start search
	// for LDAP Objects. Source meaning who has access
	// to view this delegation.
	SourceBaseDN string `json:"sourceBaseDN"`

	// The Base OU in RapidIdentity LDAP to start search
	// for LDAP Objects. Target meaning what LDAP Objects
	// that show up in the delegation view.
	TargetBaseDN string `json:"targetBaseDN"`

	// Attributes associated with the delegations.
	Attributes []DelegationAttribute `json:"attributes"`

	// The image in the profile layout.
	LayoutImage string `json:"layoutImage"`

	// The attribute to user for the 1 position of the
	// profile layout.
	Layout1 string `json:"layout1"`

	// The attribute to user for the 2 position of the
	// profile layout.
	Layout2 string `json:"layout2"`

	// The attribute to user for the 3 position of the
	// profile layout.
	Layout3 string `json:"layout3"`

	// The actions associated with the delegation.
	Actions []Action `json:"actions"`
}

type Profile struct {
	// The idautoID of the user.
	Id string `json:"id"`

	// The attributes and values associated with
	// the user's profiles.
	Attributes []ProfileAttribute `json:"attributes"`
}

type DelegationAttribute struct {
	// The GalItem associated with the
	// delegation attribute.
	GalItem GalItem `json:"galItem"`

	// The name of the attribute for the delegation.
	Name string `json:"name"`

	// Whether the delegation attribute is editable.
	Editable bool `json:"editable"`

	// Whether the delegation attribute is in the table view
	// of a delegation view.
	ShowInList bool `json:"showInList"`

	// Whether the delegation attribute is in the details
	// of a delegation view.
	ShowInDetails bool `json:"showInDetails"`

	// Whether or not the attribute is required in the delegation
	// view.
	Required bool `json:"required"`
}

type GalItem struct {
	// The ID of the GalItem attribute.
	Id string `json:"id"`

	// The friendly name for the GalItem attribute.
	FriendlyName string `json:"friendlyName"`

	// Whether the GalItem attribute accepts multiple
	// values.
	AllowMultiValue bool `json:"allowMultiValue"`

	// The type of the GalItem attribute.
	Type string `json:"type"`
}

type Action struct {
	// The ID of the action.
	Id string `json:"id"`

	// The action friendly name.
	Name string `json:"name"`

	// The action description.
	Description string `json:"description"`
}

type ProfileAttribute struct {
	// The ID of the GalItem
	Id string `json:"id"`

	// The GalItem friendly name.
	Name string `json:"name"`

	// The value(s) of the attribute for the user
	Values []string `json:"values"`
}

// Gets all associated delegations and profiles for the user
// based on their idautoID. This method will only return
// delegations and profiles that the invoking session has access.
//
//meta:operation GET /profiles/aggregated/for/{userId}
func (c *Client) GetDelegationsForUser(params GetDelegationsForUserInput) (*GetDelegationsForUserOutput, error) {
	var output GetDelegationsForUserOutput

	url := fmt.Sprintf("%s/profiles/aggregated/for/%s", c.baseEndpoint, params.UserId)
	req, err := c.GenerateRequest("GET", url, nil)
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
