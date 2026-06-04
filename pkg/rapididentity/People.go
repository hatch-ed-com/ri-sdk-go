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

// Input for setting the RapidIdentity Password for the user via delegations.
type SetPasswordInput struct {
	// Whether the password change is self-service.
	IsSelfService bool `json:"isSelfService" jsonschema:"Whether the password change is self-service."`
	// The delegation ID to use for the password change.
	DelegationId string `json:"delegationId" jsonschema:"The delegation ID to use for the password change."`
	// Whether the user must update their password on next login.
	MustUpdate bool `json:"mustUpdate" jsonschema:"Whether the user must update their password on next login."`
	// The idautoIDs of the target users.
	Targets StringList `json:"targets" jsonschema:"The idautoIDs of the target users."`
	// The new password for the user.
	NewPassword string `json:"newPassword" jsonschema:"The new password for the user."`
}

// Result of a password change for a target user.
type SetPasswordResult struct {
	// The idautoID of the target user.
	Target string `json:"target" jsonschema:"The idautoID of the target user."`
	// Whether the password change was successful.
	Success bool `json:"success" jsonschema:"Whether the password change was successful."`
	// The name of the target user.
	TargetName string `json:"targetName" jsonschema:"The name of the target user."`
}

// Output of setting the RapidIdentity Password for the user via delegations.
type SetPasswordOutput []SetPasswordResult

func (spo SetPasswordOutput) MarshalJSON() ([]byte, error) {
	if spo == nil {
		return []byte("[]"), nil
	}
	return json.Marshal([]SetPasswordResult(spo))
}

// Input for retrieving password policies for specified users.
type GetPasswordPoliciesForInput struct {
	// The idautoIDs of the target users.
	UserIds StringList `json:"userIds" jsonschema:"The idautoIDs of the target users."`
	// The type of policy to retrieve. Default is "passwordPolicy".
	Type string `json:"type" jsonschema:"The type of policy to retrieve. Default is \"passwordPolicy\"."`
}

type CharSetList []CharSet

func (csl CharSetList) MarshalJSON() ([]byte, error) {
	if csl == nil {
		return []byte("[]"), nil
	}
	return json.Marshal([]CharSet(csl))
}

type PasswordPolicyAttributeList []PasswordPolicyAttribute

func (ppal PasswordPolicyAttributeList) MarshalJSON() ([]byte, error) {
	if ppal == nil {
		return []byte("[]"), nil
	}
	return json.Marshal([]PasswordPolicyAttribute(ppal))
}

type GroupAclList []GroupAcl

func (gal GroupAclList) MarshalJSON() ([]byte, error) {
	if gal == nil {
		return []byte("[]"), nil
	}
	return json.Marshal([]GroupAcl(gal))
}

// Password Policy details.
type PasswordPolicy struct {
	// The ID of the password policy.
	Id string `json:"id" jsonschema:"The ID of the password policy."`
	// The version of the password policy.
	Version int `json:"version" jsonschema:"The version of the password policy."`
	// The name of the password policy.
	Name string `json:"name" jsonschema:"The name of the password policy."`
	// The description of the password policy.
	Description string `json:"description" jsonschema:"The description of the password policy."`
	// The priority of the password policy.
	Priority int `json:"priority" jsonschema:"The priority of the password policy."`
	// Whether the password policy is enabled.
	Enabled bool `json:"enabled" jsonschema:"Whether the password policy is enabled."`
	// Whether this is the default password policy.
	Default bool `json:"default" jsonschema:"Whether this is the default password policy."`
	// Whether group ACLs are enabled.
	GroupAclsEnabled bool `json:"groupAclsEnabled" jsonschema:"Whether group ACLs are enabled."`
	// The list of group ACLs.
	GroupAcls GroupAclList `json:"groupAcls" jsonschema:"The list of group ACLs."`
	// Whether filter ACL is enabled.
	FilterAclEnabled bool `json:"filterAclEnabled" jsonschema:"Whether filter ACL is enabled."`
	// The filter ACL expression.
	FilterAcl string `json:"filterAcl" jsonschema:"The filter ACL expression."`
	// The password reset attribute.
	PasswordResetAttribute PasswordPolicyAttribute `json:"passwordResetAttribute" jsonschema:"The password reset attribute."`
	// The minimum length of the password.
	MinLength int `json:"minLength" jsonschema:"The minimum length of the password."`
	// The maximum length of the password.
	MaxLength int `json:"maxLength" jsonschema:"The maximum length of the password."`
	// The character sets required for the password policy.
	CharSets CharSetList `json:"charSets" jsonschema:"The character sets required for the password policy."`
	// The number of required character sets.
	RequiredCharSets int `json:"requiredCharSets" jsonschema:"The number of required character sets."`
	// Whether random passwords are allowed.
	AllowRandomPassword bool `json:"allowRandomPassword" jsonschema:"Whether random passwords are allowed."`
	// The regex pattern for allowed characters.
	AllowedCharacterRegex string `json:"allowedCharacterRegex" jsonschema:"The regex pattern for allowed characters."`
	// The list of attributes to match against the password.
	MatchingAttributes PasswordPolicyAttributeList `json:"matchingAttributes" jsonschema:"The list of attributes to match against the password."`
	// Whether matching attributes are case sensitive.
	MatchingAttributesCaseSensitive bool `json:"matchingAttributesCaseSensitive" jsonschema:"Whether matching attributes are case sensitive."`
	// Whether matching attributes must match the entire string.
	MatchingAttributesMatchEntire bool `json:"matchingAttributesMatchEntire" jsonschema:"Whether matching attributes must match the entire string."`
	// The list of blacklisted passwords.
	BlackListed StringList `json:"blackListed" jsonschema:"The list of blacklisted passwords."`
	// Whether the blacklist is case sensitive.
	BlackListCaseSensitive bool `json:"blackListCaseSensitive" jsonschema:"Whether the blacklist is case sensitive."`
	// Whether the blacklist must match the entire string.
	BlackListMatchEntire bool `json:"blackListMatchEntire" jsonschema:"Whether the blacklist must match the entire string."`
	// The list of blacklisted regexes.
	BlackListRegexes StringList `json:"blackListRegexes" jsonschema:"The list of blacklisted regexes."`
	// Whether users must change their password by default.
	DefaultForceUserPasswordChange bool `json:"defaultForceUserPasswordChange" jsonschema:"Whether users must change their password by default."`
	// Whether to hide the force password change option.
	HideForceUserPasswordChange bool `json:"hideForceUserPasswordChange" jsonschema:"Whether to hide the force password change option."`
	// Whether to enforce AD complexity attributes.
	EnforceADComplexityAttributes bool `json:"enforceADComplexityAttributes" jsonschema:"Whether to enforce AD complexity attributes."`
	// Whether to enforce password history on admin reset.
	EnforcePasswordHistoryAdminReset bool `json:"enforcePasswordHistoryAdminReset" jsonschema:"Whether to enforce password history on admin reset."`
	// Whether the policy can be overridden by a delegated policy.
	DelegatedPolicyOverride bool `json:"delegatedPolicyOverride" jsonschema:"Whether the policy can be overridden by a delegated policy."`
	// The screening configuration for the password policy.
	ScreeningConfig ScreeningConfig `json:"screeningConfig" jsonschema:"The screening configuration for the password policy."`
	// The number of previous passwords to remember.
	PasswordRememberedCount int `json:"passwordRememberedCount" jsonschema:"The number of previous passwords to remember."`
	// The maximum age of a password in days.
	PasswordMaximumAgeDays int `json:"passwordMaximumAgeDays" jsonschema:"The maximum age of a password in days."`
	// The number of days before expiration to start warning the user.
	ExpirationWarningDays int `json:"expirationWarningDays" jsonschema:"The number of days before expiration to start warning the user."`
	// The maximum number of failed login attempts allowed.
	MaxFailedAttempts int `json:"maxFailedAttempts" jsonschema:"The maximum number of failed login attempts allowed."`
	// The window in minutes for failed attempts.
	FailedAttemptsWindow int `json:"failedAttemptsWindow" jsonschema:"The window in minutes for failed attempts."`
	// The duration of account lockout in minutes.
	AccountLockoutDuration int `json:"accountLockoutDuration" jsonschema:"The duration of account lockout in minutes."`
}

type GroupAcl struct {
	// The ID of the group.
	Id string `json:"id" jsonschema:"The ID of the group."`
	// The name of the group.
	Name string `json:"name" jsonschema:"The name of the group."`
	// The description of the group.
	Description string `json:"description" jsonschema:"The description of the group."`
}

type PasswordPolicyAttribute struct {
	// The ID of the attribute.
	Id string `json:"id" jsonschema:"The ID of the attribute."`
	// The friendly name of the attribute.
	FriendlyName string `json:"friendlyName" jsonschema:"The friendly name of the attribute."`
	// Whether the attribute is searchable.
	Searchable bool `json:"searchable" jsonschema:"Whether the attribute is searchable."`
	// Whether the attribute is multi-valued (Legacy).
	MultiValued bool `json:"multiValued" jsonschema:"Whether the attribute is multi-valued (Legacy)."`
	// Whether the attribute allows multiple values.
	AllowMultiValue bool `json:"allowMultiValue" jsonschema:"Whether the attribute allows multiple values."`
	// The type of the attribute.
	Type string `json:"type" jsonschema:"The type of the attribute."`
}

type ScreeningConfig struct {
	// Whether screening is enabled.
	Enabled bool `json:"enabled" jsonschema:"Whether screening is enabled."`
	// The error message to display when screening fails.
	ErrorMessage string `json:"errorMessage" jsonschema:"The error message to display when screening fails."`
	// The configuration for the screening type.
	TypeConfig TypeConfig `json:"typeConfig" jsonschema:"The configuration for the screening type."`
}

type TypeConfig struct {
	// The type of screening.
	Type string `json:"@type" jsonschema:"The type of screening."`
}

type CharSet struct {
	// The ID of the character set.
	Id string `json:"id" jsonschema:"The ID of the character set."`
	// The type of the character set.
	Type string `json:"type" jsonschema:"The type of the character set."`
	// The minimum number of characters from this set.
	Min int `json:"min" jsonschema:"The minimum number of characters from this set."`
	// The maximum number of characters from this set.
	Max int `json:"max" jsonschema:"The maximum number of characters from this set."`
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

// Sets the RapidIdentity Password for the user via delegations.
//
//meta:operation POST /profiles/actions/password
func (c *Client) SetPassword(ctx context.Context, params SetPasswordInput) (SetPasswordOutput, error) {
	var output SetPasswordOutput

	url := fmt.Sprintf("%s/profiles/actions/password", c.baseEndpoint)
	body, err := json.Marshal(params)
	if err != nil {
		return nil, err
	}
	requestBody := bytes.NewBuffer(body)
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

// Retrieves the password policy for specified users.
//
//meta:operation POST /profiles/passwordPolicies/for
func (c *Client) GetPasswordPoliciesFor(ctx context.Context, params GetPasswordPoliciesForInput) (*PasswordPolicy, error) {
	var output PasswordPolicy
	params.Type = cmp.Or(params.Type, "passwordPolicy")

	url := fmt.Sprintf("%s/profiles/passwordPolicies/for", c.baseEndpoint)
	body, err := json.Marshal(params)
	if err != nil {
		return nil, err
	}
	requestBody := bytes.NewBuffer(body)
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

	return &output, nil
}
