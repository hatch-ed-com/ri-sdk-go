package rapididentity

import (
	"bytes"
	"encoding/json"
	"fmt"
)

// Params for GetAuthenticationPoliciesForUser method.
type GetAuthenticationPoliciesForUserInput struct {
	// Whether to provide authentication policies in response.
	// If this is fault, user information is shown only.
	ShowAuthenticationPolicies bool

	// Whether to provide a claim for the user
	// in the form of a json web token (jwt).
	ShowClaims bool

	// The fields to show in the authenticationPolicies
	// response. By default all fields are shown.
	AuthenticationPolicyFieldsToShow []string

	// The user to get authentication policies.
	User GetAuthenticationPoliciesForUserPayload
}

// Request payload for retrieving authentication policies
// for a user.
type GetAuthenticationPoliciesForUserPayload struct {
	// Username of the user. This can be any value within
	// the idautoPersonUsernameMV attribute within RapidIdentity.
	Username string `json:"username"`
}

// Output for the GetAuthenticationPoliciesForUser method.
type GetAuthenticationPoliciesForUserOutput struct {
	// User information for the username provided in the request.
	User User `json:"user"`

	// Authentication policies for the username provided in the request.
	AuthenticationPolicies []AuthenticationPolicy `json:"authenticationPolicies"`
}

// RapidIdentity authentication policy information.
type AuthenticationPolicy struct {
	// Unique id for the authentication policy.
	Id string `json:"id"`

	// The version of the authentication policy.
	Version int `json:"version"`

	// The name of the authentication policy.
	Name string `json:"name"`

	// Whether the authentication policy is enabled.
	Enabled bool `json:"enabled"`

	// The criteria for the authentication policy.
	// This utilize an interface as the criteria array
	// returns several different objects. Due to this
	// several Criteria structs were created such as
	// DaysOfTheWeekCriteria that implement the interface.
	Criteria []AuthenticationPolicyCriteria `json:"criteria"`

	// The methods for the authentication policy.
	// This utilize an interface as the methods array
	// returns several different objects. Due to this
	// several Method structs were created such as
	// DuoMethod that implement the interface.
	Methods []AuthenticationPolicyMethod `json:"methods"`

	// Whether the authentication policy can be initated with a QR Code.
	InsecureQRIdEnabled bool `json:"insecureQRIdEnabled"`

	// Whether the authentication policy should always fail.
	AlwaysFail bool `json:"alwaysFail"`

	// Whether the authentication policy is a forgot password policy.
	IsResetPasswordPolicy bool `json:"isResetPasswordPolicy"`
}

// Interface for all policy criteria to implement.
type AuthenticationPolicyCriteria interface {
	GetBaseAuthenticationCriteriaInfo() BaseAuthenticationInfo
}

// Interface for all policy methods to implement.
type AuthenticationPolicyMethod interface {
	GetBaseAuthenticationMethodInfo() BaseAuthenticationInfo
}

// Custom UnmarshJSON method due to Criteria and Methods fields on
// GetAuthenticationPoliciesForUserOutput.
func (ap *AuthenticationPolicy) UnmarshalJSON(b []byte) error {
	var AuthenticationBaseData map[string]interface{}
	err := json.Unmarshal(b, &AuthenticationBaseData)
	if err != nil {
		return err
	}

	for apk, apv := range AuthenticationBaseData {
		switch apk {
		case "id":
			ap.Id = apv.(string)
		case "version":
			ap.Version = int(apv.(float64))
		case "name":
			ap.Name = apv.(string)
		case "enabled":
			ap.Enabled = apv.(bool)
		case "insecureQRIdEnabled":
			ap.InsecureQRIdEnabled = apv.(bool)
		case "alwaysFail":
			ap.AlwaysFail = apv.(bool)
		case "isResetPasswordPolicy":
			ap.IsResetPasswordPolicy = apv.(bool)
		case "criteria":
			var criteriaDataList []map[string]interface{}
			criteriaData, err := json.Marshal(apv)
			if err != nil {
				return err
			}

			err = json.Unmarshal(criteriaData, &criteriaDataList)
			if err != nil {
				return err
			}

			for _, cd := range criteriaDataList {
				typeValue := cd["type"].(string)
				criteria, err := json.Marshal(cd)
				if err != nil {
					return err
				}

				switch typeValue {
				case "dayOfWeek":
					var dayOfWeekCriteria DaysOfWeekCriteria
					err = json.Unmarshal(criteria, &dayOfWeekCriteria)
					if err != nil {
						return err
					}
					ap.Criteria = append(ap.Criteria, dayOfWeekCriteria)
				case "webAuthn":
					var webAuthnCriteria WebAuthnCriteria
					err = json.Unmarshal(criteria, &webAuthnCriteria)
					if err != nil {
						return err
					}
					ap.Criteria = append(ap.Criteria, webAuthnCriteria)
				case "kerberos":
					var kerberosCriteria KerberosCriteria
					err = json.Unmarshal(criteria, &kerberosCriteria)
					if err != nil {
						return err
					}
					ap.Criteria = append(ap.Criteria, kerberosCriteria)
				case "ldapFilter":
					var ldapFilterCriteria LdapFilterCriteria
					err = json.Unmarshal(criteria, &ldapFilterCriteria)
					if err != nil {
						return err
					}
					ap.Criteria = append(ap.Criteria, ldapFilterCriteria)
				case "qrCode":
					var qrCodeCriteria QrCodeCriteria
					err = json.Unmarshal(criteria, &qrCodeCriteria)
					if err != nil {
						return err
					}
					ap.Criteria = append(ap.Criteria, qrCodeCriteria)
				case "sourceNetwork":
					var sourceNetworkCriteria SourceNetworkCriteria
					err = json.Unmarshal(criteria, &sourceNetworkCriteria)
					if err != nil {
						return err
					}
					ap.Criteria = append(ap.Criteria, sourceNetworkCriteria)
				case "role":
					var roleCriteria RoleCriteria
					err = json.Unmarshal(criteria, &roleCriteria)
					if err != nil {
						return err
					}
					ap.Criteria = append(ap.Criteria, roleCriteria)
				case "timeOfDay":
					var timeOfDayCriteria TimeOfDayCriteria
					err = json.Unmarshal(criteria, &timeOfDayCriteria)
					if err != nil {
						return err
					}
					ap.Criteria = append(ap.Criteria, timeOfDayCriteria)
				}
			}
		case "methods":
			var methodDataList []map[string]interface{}
			methodData, err := json.Marshal(apv)
			if err != nil {
				return err
			}

			err = json.Unmarshal(methodData, &methodDataList)
			if err != nil {
				return err
			}

			for _, cd := range methodDataList {
				typeValue := cd["type"].(string)
				method, err := json.Marshal(cd)
				if err != nil {
					return err
				}

				switch typeValue {
				case "duo":
					var duoMethod DuoMethod
					err = json.Unmarshal(method, &duoMethod)
					if err != nil {
						return err
					}
					ap.Methods = append(ap.Methods, duoMethod)
				case "email":
					var emailMethod EmailMethod
					err = json.Unmarshal(method, &emailMethod)
					if err != nil {
						return err
					}
					ap.Methods = append(ap.Methods, emailMethod)
				case "federation":
					var federationMethod FederationMethod
					err = json.Unmarshal(method, &federationMethod)
					if err != nil {
						return err
					}
					ap.Methods = append(ap.Methods, federationMethod)
				case "webAuthn":
					var webAuthnMethod WebAuthnMethod
					err = json.Unmarshal(method, &webAuthnMethod)
					if err != nil {
						return err
					}
					ap.Methods = append(ap.Methods, webAuthnMethod)
				case "kerberos":
					var kerberosMethod KerberosMethod
					err = json.Unmarshal(method, &kerberosMethod)
					if err != nil {
						return err
					}
					ap.Methods = append(ap.Methods, kerberosMethod)
				case "password":
					var passwordMethod PasswordMethod
					err = json.Unmarshal(method, &passwordMethod)
					if err != nil {
						return err
					}
					ap.Methods = append(ap.Methods, passwordMethod)
				case "pictograph":
					var pictographMethod PictographMethod
					err = json.Unmarshal(method, &pictographMethod)
					if err != nil {
						return err
					}
					ap.Methods = append(ap.Methods, pictographMethod)
				case "pingMe":
					var pingMeMethod PingMeMethod
					err = json.Unmarshal(method, &pingMeMethod)
					if err != nil {
						return err
					}
					ap.Methods = append(ap.Methods, pingMeMethod)
				case "rapidPortalChallenge":
					var rapidPortalChallengeMethod RapidPortalChallengeMethod
					err = json.Unmarshal(method, &rapidPortalChallengeMethod)
					if err != nil {
						return err
					}
					ap.Methods = append(ap.Methods, rapidPortalChallengeMethod)
				case "qrCode":
					var qrCodeMethod QrCodeMethod
					err = json.Unmarshal(method, &qrCodeMethod)
					if err != nil {
						return err
					}
					ap.Methods = append(ap.Methods, qrCodeMethod)
				case "sms":
					var smsMethod SmsMethod
					err = json.Unmarshal(method, &smsMethod)
					if err != nil {
						return err
					}
					ap.Methods = append(ap.Methods, smsMethod)
				case "social":
					var socialMethod SocialMethod
					err = json.Unmarshal(method, &socialMethod)
					if err != nil {
						return err
					}
					ap.Methods = append(ap.Methods, socialMethod)
				case "totp":
					var totpMethod TotpMethod
					err = json.Unmarshal(method, &totpMethod)
					if err != nil {
						return err
					}
					ap.Methods = append(ap.Methods, totpMethod)
				case "userAgreement":
					var userAgreementMethod UserAgreementMethod
					err = json.Unmarshal(method, &userAgreementMethod)
					if err != nil {
						return err
					}
					ap.Methods = append(ap.Methods, userAgreementMethod)
				}
			}
		}
	}
	return nil
}

// Base fields for policy methods and criteria.
type BaseAuthenticationInfo struct {
	// The type of method or criteria.
	Type string `json:"type"`

	// Whether the method or criteria is enabled.
	Enabled bool `json:"enabled"`
}

// Determines if authentication policy is satisfied based on
// the day of the week
type DaysOfWeekCriteria struct {
	BaseAuthenticationInfo
	// Authentication policy is enabled on Sundays
	Sunday bool `json:"sunday"`

	// Authentication policy is enabled on Mondays
	Monday bool `json:"monday"`

	// Authentication policy is enabled on Tuesdays
	Tuesday bool `json:"tuesday"`

	// Authentication policy is enabled on Wednesdays
	Wednesday bool `json:"wednesday"`

	// Authentication policy is enabled on Thursdays
	Thursday bool `json:"thursday"`

	// Authentication policy is enabled on Fridays
	Friday bool `json:"friday"`

	// Authentication policy is enabled on Saturdays
	Saturday bool `json:"saturday"`
}

func (dwc DaysOfWeekCriteria) GetBaseAuthenticationCriteriaInfo() BaseAuthenticationInfo {
	return dwc.BaseAuthenticationInfo
}

// Determines if authentication policy is satisfied based on
// WebAuthn devices registered.
type WebAuthnCriteria struct {
	BaseAuthenticationInfo
	// If Negate is set to true then any user with
	// WebAuthn devices registered will not satisfy the
	// authentication policy
	Negate bool `json:"negate"`
}

func (wac WebAuthnCriteria) GetBaseAuthenticationCriteriaInfo() BaseAuthenticationInfo {
	return wac.BaseAuthenticationInfo
}

// Determines if authentication policy is satisfied based on
// a kerberos token being present.
type KerberosCriteria struct {
	BaseAuthenticationInfo
}

func (kc KerberosCriteria) GetBaseAuthenticationCriteriaInfo() BaseAuthenticationInfo {
	return kc.BaseAuthenticationInfo
}

// Determines if authentication policy is satisfied based on
// LDAP filter provided.
type LdapFilterCriteria struct {
	BaseAuthenticationInfo
	// The LDAP filter to determine what users apply to policy.
	LdapFilter string `json:"ldapFilter"`

	// Whether to match the LDAP Admin account.
	MatchNonLdapAdmin bool `json:"matchNonLdapAdmin"`
}

func (lfc LdapFilterCriteria) GetBaseAuthenticationCriteriaInfo() BaseAuthenticationInfo {
	return lfc.BaseAuthenticationInfo
}

// Allows for QR Code to initiate authentication policy.
type QrCodeCriteria struct {
	BaseAuthenticationInfo
}

func (qcc QrCodeCriteria) GetBaseAuthenticationCriteriaInfo() BaseAuthenticationInfo {
	return qcc.BaseAuthenticationInfo
}

// Determines if authentication policy is satisfied based on
// Source IP Address.
type SourceNetworkCriteria struct {
	BaseAuthenticationInfo
	// List of subnets to evaluate.
	Subnets []string `json:"subnets"`

	// Allow for insecure http.
	EnableHttpHeaderProcessing bool `json:"enableHttpHeaderProcessing"`

	// Whether to allow or deny the subnets listed.
	Negate bool `json:"negate"`
}

func (snc SourceNetworkCriteria) GetBaseAuthenticationCriteriaInfo() BaseAuthenticationInfo {
	return snc.BaseAuthenticationInfo
}

// Determines if authentication policy is satisfied based on
// user being within certain roles.
type RoleCriteria struct {
	BaseAuthenticationInfo

	// List of roles to evaluate.
	Roles []RoleAuthValue `json:"roles"`

	// Whether to apply policy to everyone.
	ApplyToEveryone bool `json:"applyToEveryone"`

	// Whether to allow or deny users in the roles associated with
	// this policy
	InverseMatch bool `json:"inverseMatch"`
}

func (rc RoleCriteria) GetBaseAuthenticationCriteriaInfo() BaseAuthenticationInfo {
	return rc.BaseAuthenticationInfo
}

// Role information for RoleCriteria.
type RoleAuthValue struct {
	// The idautoID of the role in RapidIdentity.
	Id string `json:"id"`

	// The name of the role in RapidIdentity.
	Name string `json:"name"`
}

// Determines whether authentication policy is satisfied
// based on the time of the day.
type TimeOfDayCriteria struct {
	BaseAuthenticationInfo

	// The start time to evaluate.
	Start TimeOfDayCriteriaSyntax

	// The end time to evaluate.
	End TimeOfDayCriteriaSyntax

	// The time zone to evaluate.
	TimeZone string
}

func (tdc TimeOfDayCriteria) GetBaseAuthenticationCriteriaInfo() BaseAuthenticationInfo {
	return tdc.BaseAuthenticationInfo
}

// Time of day value for the TimeOfDayCriteria.
type TimeOfDayCriteriaSyntax struct {
	// Hour from 0 - 23.
	Hour int `json:"hour"`

	// Minute from 0 - 60.
	Minute int `json:"minute"`
}

// Use DUO push or one time code to authenticate.
type DuoMethod struct {
	BaseAuthenticationInfo

	// The configuration ID for the Duo Configuration.
	ConfigId string `json:"configId"`

	// Removes additional click to enter DUO prompt.
	AutoProcess bool `json:"autoProcess"`
}

func (dm DuoMethod) GetBaseAuthenticationMethodInfo() BaseAuthenticationInfo {
	return dm.BaseAuthenticationInfo
}

// Recieve a one time code to email to authenticate.
type EmailMethod struct {
	BaseAuthenticationInfo
}

func (em EmailMethod) GetBaseAuthenticationMethodInfo() BaseAuthenticationInfo {
	return em.BaseAuthenticationInfo
}

// Utilized SAML to extenral IDP to authenticate.
type FederationMethod struct {
	BaseAuthenticationInfo

	// The Trusted IDP configuration to reference
	TrustedIdp FederationAuthValue `json:"trustedIdp"`

	// The url to redirect to once authentication at trusted IDP
	// has been completed.
	PostAuthRedirectUrl string `json:"postAuthRedirectUrl"`

	// Whether to expose atrributes in SAML response.
	ExposeAttributes bool `json:"exposeAttributes"`

	// Whether to forward the username to the IDP.
	ForwardUsernameEnabled bool `json:"forwardUsernameEnabled"`

	// The username attribute to forward to the IDP.
	ForwardUsernameAttribute string `json:"forwardUsernameAttribute"`

	// The SAML NAME ID Format of the forward attribute.
	ForwardUsernameNameIDFormat string `json:"forwardUsernameNameIDFormat"`
}

func (fm FederationMethod) GetBaseAuthenticationMethodInfo() BaseAuthenticationInfo {
	return fm.BaseAuthenticationInfo
}

// The federation values from the FederationMethod.
type FederationAuthValue struct {

	// The unique ID of the trusted IDP configuration.
	Id string `json:"id"`

	// THe name of the trusted IDP configuration.
	Name string `json:"name"`
}

// Use FIDO, or device login to authenticate.
type WebAuthnMethod struct {
	BaseAuthenticationInfo

	// Allow login to be remembered for 30 days.
	AllowChallengeDeferral bool `json:"allowChallengeDeferral"`
}

func (wam WebAuthnMethod) GetBaseAuthenticationMethodInfo() BaseAuthenticationInfo {
	return wam.BaseAuthenticationInfo
}

// Authenticate with device login on Active Directory joined device.
type KerberosMethod struct {
	BaseAuthenticationInfo
}

func (km KerberosMethod) GetBaseAuthenticationMethodInfo() BaseAuthenticationInfo {
	return km.BaseAuthenticationInfo
}

// Use password to authenticate.
type PasswordMethod struct {
	BaseAuthenticationInfo

	// Whether to display expiration warning when password is close to expiring.
	ExpirationWarningEnabled bool `json:"expirationWarningEnabled"`

	// How many days prior to password expiration to show expiration warning.
	ExpirationWarningDays int `json:"expirationWarningDays"`

	// How long since the password has been changed.
	CurrentPasswordAgeDays int `json:"currentPasswordAgeDays"`

	// The maximum number of days a password can be used before it must be changed.
	PasswordMaximumAgeDays int `json:"passwordMaximumAgeDays"`

	// Whether the user must change their password on login.
	MustChange bool `json:"mustChange"`
}

func (pm PasswordMethod) GetBaseAuthenticationMethodInfo() BaseAuthenticationInfo {
	return pm.BaseAuthenticationInfo
}

// Utilize a pool of images to authenticate.
type PictographMethod struct {
	BaseAuthenticationInfo

	// Number of images to display to user to choose from.
	NumToChallenge int `json:"numToChallenge"`

	// The number of images a user must select to authenticate.
	NumToChoose int `json:"numToChoose"`

	// Whether to use the default pool of images or custom images.
	UseDefaultImagePool bool `json:"useDefaultImagePool"`

	// The imageIds to use if custom image pool is used.
	ImageIds []string `json:"imageIds"`
}

func (pm PictographMethod) GetBaseAuthenticationMethodInfo() BaseAuthenticationInfo {
	return pm.BaseAuthenticationInfo
}

// Utilize the RapidIdentity Mobile app to authenticate.
// via push or one time code.
type PingMeMethod struct {
	BaseAuthenticationInfo

	// Whether to use cloud based pingMe (Legacy).
	NativePingMe bool `json:"nativePingMe"`

	// A friendly description for the service.
	ServiceDescription string `json:"serviceDescription"`
}

func (pmm PingMeMethod) GetBaseAuthenticationMethodInfo() BaseAuthenticationInfo {
	return pmm.BaseAuthenticationInfo
}

// Utilize challenge questions and answers to authenticate.
type RapidPortalChallengeMethod struct {
	BaseAuthenticationInfo

	// The RapidIdentity Portal base URL.
	RapidPortalBaseUrl string `json:"rapidPortalBaseUrl"`
}

func (rpcm RapidPortalChallengeMethod) GetBaseAuthenticationMethodInfo() BaseAuthenticationInfo {
	return rpcm.BaseAuthenticationInfo
}

// Utilize QR Code to authenticate.
type QrCodeMethod struct {
	BaseAuthenticationInfo
}

func (qcm QrCodeMethod) GetBaseAuthenticationMethodInfo() BaseAuthenticationInfo {
	return qcm.BaseAuthenticationInfo
}

// Utilize one time code sent to mobile number to authenticate.
type SmsMethod struct {
	BaseAuthenticationInfo
}

func (sm SmsMethod) GetBaseAuthenticationMethodInfo() BaseAuthenticationInfo {
	return sm.BaseAuthenticationInfo
}

// Utilize social network provider to authenticate.
type SocialMethod struct {
	BaseAuthenticationInfo

	// Apple provider information.
	Apple SocialProviderAppleInfo `json:"apple"`

	// Google provider information.
	GooglePlus SocialProviderGoogleInfo `json:"googlePlus"`
}

func (sm SocialMethod) GetBaseAuthenticationMethodInfo() BaseAuthenticationInfo {
	return sm.BaseAuthenticationInfo
}

// Apple social provider information.
type SocialProviderAppleInfo struct {
	// Whether Apple is enabled to be used.
	Enabled bool `json:"enabled"`

	// The private key associated with the Apple social provider.
	PrivateKey string `json:"privateKey"`
}

type SocialProviderGoogleInfo struct {
	// Whether Google is enabled to be used
	Enabled bool `json:"enabled"`

	// The client secret associated with the Google social provider.
	ClientSecret string `json:"clientSecret"`
}

// Authenticate with one time code to TOTP application.
type TotpMethod struct {
	BaseAuthenticationInfo

	// The window size for the registration QR Code.
	TotpWindowSize int `json:"totpWindowSize"`

	// Only challenge user with one time code every 30 days.
	AllowChallengeDeferral bool `json:"allowChallengeDeferral"`

	// The issuer name for the TOTP code.
	IssuerName string `json:"issuerName"`

	// The setup instructions to display during registration.
	SetupInstructions string `json:"setupInstructions"`
}

func (tm TotpMethod) GetBaseAuthenticationMethodInfo() BaseAuthenticationInfo {
	return tm.BaseAuthenticationInfo
}

// User agreement to display on login.
type UserAgreementMethod struct {
	BaseAuthenticationInfo

	// The unique id that references the User Agreement.
	UserAgreementId string `json:"userAgreementId"`

	// Whether to show user agreement on every login or just one time.
	ShowuserAgreementEveryTime bool `json:"showuserAgreementEveryTime"`
}

func (uam UserAgreementMethod) GetBaseAuthenticationMethodInfo() BaseAuthenticationInfo {
	return uam.BaseAuthenticationInfo
}

// Retrieves authentication policies for specified user.
//
//meta:operation POST /authn/v1/username
func (c *Client) GetAuthenticationPoliciesForUser(params GetAuthenticationPoliciesForUserInput) (*GetAuthenticationPoliciesForUserOutput, error) {
	url := fmt.Sprintf("%s/authn/v1/username?authenticationPolicies=%t&claim=%t", c.baseEndpoint, params.ShowAuthenticationPolicies, params.ShowClaims)
	for _, field := range params.AuthenticationPolicyFieldsToShow {
		url = fmt.Sprintf("%s&authenticationPolicyField=%s", url, field)
	}
	user, err := json.Marshal(params.User)
	if err != nil {
		return nil, err
	}
	requestBody := bytes.NewBuffer(user)
	req, err := c.GenerateRequest("POST", url, requestBody)
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
	var output GetAuthenticationPoliciesForUserOutput
	err = json.Unmarshal(resBody, &output)
	if err != nil {
		return nil, err
	}

	return &output, nil
}
