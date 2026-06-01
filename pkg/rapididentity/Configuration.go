package rapididentity

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"time"
)

// A RapidIdentity user session
type Session struct {
	// The session object.
	Session SessionInfo `json:"session"`

	// Whether a password update is required with the session.
	PasswordUpdateRequired bool `json:"passwordUpdateRequired"`

	// Number of logins remaining before user is locked out.
	GraceLoginsRemaining int `json:"graceLoginsRemaining"`
}
type StringList []string

func (sl StringList) MarshalJSON() ([]byte, error) {
	if sl == nil {
		return []byte("[]"), nil
	}
	return json.Marshal([]string(sl))
}

type SessionInfo struct {
	// The session ID.
	Id string `json:"id"`

	// The session token.
	Token string `json:"token"`

	// The session user information. If using a proxy session
	// this will be the proxied user.
	User User `json:"user"`

	// If using a proxy session this will be the user
	// who invoked the proxy.
	RealUser User `json:"realUser"`

	// The RapidIdentity roles associated with the user
	// This does not contain the groups the user is a member of.
	Roles StringList `json:"roles"`

	// When the session was created.
	Created time.Time `json:"created"`

	// The Client IP Address used to create the session.
	CreatedClientIp string `json:"createdClientIp"`

	// The Host IP address used to create the session.
	CreatedHostIp string `json:"createdHostIp"`

	// The time the session was last used.
	LastUsed time.Time `json:"lastUsed"`

	// The Client IP Address that was last used with the session.
	LastUsedClientIp string `json:"lastUsedClientIp"`

	// The Host IP Address that was last used with the session.
	LastUsedHostIp string `json:"lastUsedHostIp"`

	// When the session was invalidated.
	Invalidated time.Time `json:"invalidated"`

	// Proxy data associated with the session.
	ProxyData ProxyData `json:"proxyData"`
}

type ProxyData struct {
	// The RapidIdentity roles of the user who initiated the proxy
	Permissions StringList `json:"permissions"`
}

type AuthenticationPolicyList []AuthenticationPolicy

func (apl AuthenticationPolicyList) MarshalJSON() ([]byte, error) {
	if apl == nil {
		return []byte("[]"), nil
	}
	return json.Marshal([]AuthenticationPolicy(apl))
}

type AuthenticationPolicyCriteriaList []AuthenticationPolicyCriteria

func (apcl AuthenticationPolicyCriteriaList) MarshalJSON() ([]byte, error) {
	if apcl == nil {
		return []byte("[]"), nil
	}
	return json.Marshal([]AuthenticationPolicyCriteria(apcl))
}

type AuthenticationPolicyMethodList []AuthenticationPolicyMethod

func (apml AuthenticationPolicyMethodList) MarshalJSON() ([]byte, error) {
	if apml == nil {
		return []byte("[]"), nil
	}
	return json.Marshal([]AuthenticationPolicyMethod(apml))
}

// Params for GetAuthenticationPoliciesForUser method.
type GetAuthenticationPoliciesForUserInput struct {
	// Whether to provide authentication policies in response.
	// If this is false, user information is shown only.
	ShowAuthenticationPolicies bool `json:"showAuthenticationPolicies" jsonschema:"Whether to provide authentication policies in response. If this is false, user information is shown only."`

	// Whether to provide a claim for the user
	// in the form of a json web token (jwt).
	ShowClaims bool `json:"showClaims" jsonschema:"Whether to provide a claim for the user in the form of a json web token (jwt)."`

	// The fields to show in the authenticationPolicies
	// response. By default all fields are shown.
	AuthenticationPolicyFieldsToShow StringList `json:"authenticationPolicyFieldsToShow" jsonschema:"The fields to show in the authenticationPolicies response. By default all fields are shown."`

	// The user to get authentication policies.
	User GetAuthenticationPoliciesForUserPayload `json:"user" jsonschema:"The user to get authentication policies."`
}

// Request payload for retrieving authentication policies
// for a user.
type GetAuthenticationPoliciesForUserPayload struct {
	// Username of the user. This can be any value within
	// the idautoPersonUsernameMV attribute within RapidIdentity.
	Username string `json:"username" jsonschema:"Username of the user. This can be any value within the idautoPersonUsernameMV attribute within RapidIdentity."`
}

// Output for the GetAuthenticationPoliciesForUser method.
type GetAuthenticationPoliciesForUserOutput struct {
	// User information for the username provided in the request.
	User User `json:"user" jsonschema:"User information for the username provided in the request."`

	// Authentication policies for the username provided in the request.
	AuthenticationPolicies AuthenticationPolicyList `json:"authenticationPolicies" jsonschema:"Authentication policies for the username provided in the request."`
}

// RapidIdentity authentication policy information.
type AuthenticationPolicy struct {
	// Unique id for the authentication policy.
	Id string `json:"id" jsonschema:"Unique id for the authentication policy."`

	// The version of the authentication policy.
	Version int `json:"version" jsonschema:"The version of the authentication policy."`

	// The name of the authentication policy.
	Name string `json:"name" jsonschema:"The name of the authentication policy."`

	// Whether the authentication policy is enabled.
	Enabled bool `json:"enabled" jsonschema:"Whether the authentication policy is enabled."`

	// The criteria for the authentication policy.
	// This utilize an interface as the criteria array
	// returns several different objects. Due to this
	// several Criteria structs were created such as
	// DaysOfTheWeekCriteria that implement the interface.
	Criteria AuthenticationPolicyCriteriaList `json:"criteria" jsonschema:"The criteria for the authentication policy. This utilize an interface as the criteria array returns several different objects. Due to this several Criteria structs were created such as DaysOfTheWeekCriteria that implement the interface."`

	// The methods for the authentication policy.
	// This utilize an interface as the methods array
	// returns several different objects. Due to this
	// several Method structs were created such as
	// DuoMethod that implement the interface.
	Methods AuthenticationPolicyMethodList `json:"methods" jsonschema:"The methods for the authentication policy. This utilize an interface as the methods array returns several different objects. Due to this several Method structs were created such as DuoMethod that implement the interface."`

	// Whether the authentication policy can be initated with a QR Code.
	InsecureQRIdEnabled bool `json:"insecureQRIdEnabled" jsonschema:"Whether the authentication policy can be initated with a QR Code."`

	// Whether the authentication policy should always fail.
	AlwaysFail bool `json:"alwaysFail" jsonschema:"Whether the authentication policy should always fail."`

	// Whether the authentication policy is a forgot password policy.
	IsResetPasswordPolicy bool `json:"isResetPasswordPolicy" jsonschema:"Whether the authentication policy is a forgot password policy."`
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
	Type string `json:"type" jsonschema:"The type of method or criteria."`

	// Whether the method or criteria is enabled.
	Enabled bool `json:"enabled" jsonschema:"Whether the method or criteria is enabled."`
}

// Determines if authentication policy is satisfied based on
// the day of the week
type DaysOfWeekCriteria struct {
	BaseAuthenticationInfo
	// Authentication policy is enabled on Sundays
	Sunday bool `json:"sunday" jsonschema:"Authentication policy is enabled on Sundays"`

	// Authentication policy is enabled on Mondays
	Monday bool `json:"monday" jsonschema:"Authentication policy is enabled on Mondays"`

	// Authentication policy is enabled on Tuesdays
	Tuesday bool `json:"tuesday" jsonschema:"Authentication policy is enabled on Tuesdays"`

	// Authentication policy is enabled on Wednesdays
	Wednesday bool `json:"wednesday" jsonschema:"Authentication policy is enabled on Wednesdays"`

	// Authentication policy is enabled on Thursdays
	Thursday bool `json:"thursday" jsonschema:"Authentication policy is enabled on Thursdays"`

	// Authentication policy is enabled on Fridays
	Friday bool `json:"friday" jsonschema:"Authentication policy is enabled on Fridays"`

	// Authentication policy is enabled on Saturdays
	Saturday bool `json:"saturday" jsonschema:"Authentication policy is enabled on Saturdays"`
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
	Negate bool `json:"negate" jsonschema:"If Negate is set to true then any user with WebAuthn devices registered will not satisfy the authentication policy"`
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
	LdapFilter string `json:"ldapFilter" jsonschema:"The LDAP filter to determine what users apply to policy."`

	// Whether to match the LDAP Admin account.
	MatchNonLdapAdmin bool `json:"matchNonLdapAdmin" jsonschema:"Whether to match the LDAP Admin account."`
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
	Subnets StringList `json:"subnets" jsonschema:"List of subnets to evaluate."`

	// Allow for insecure http.
	EnableHttpHeaderProcessing bool `json:"enableHttpHeaderProcessing" jsonschema:"Allow for insecure http."`

	// Whether to allow or deny the subnets listed.
	Negate bool `json:"negate" jsonschema:"Whether to allow or deny the subnets listed."`
}

func (snc SourceNetworkCriteria) GetBaseAuthenticationCriteriaInfo() BaseAuthenticationInfo {
	return snc.BaseAuthenticationInfo
}

type RoleAuthValueList []RoleAuthValue

func (ravl RoleAuthValueList) MarshalJSON() ([]byte, error) {
	if ravl == nil {
		return []byte("[]"), nil
	}
	return json.Marshal([]RoleAuthValue(ravl))
}

// Determines if authentication policy is satisfied based on
// user being within certain roles.
type RoleCriteria struct {
	BaseAuthenticationInfo

	// List of roles to evaluate.
	Roles RoleAuthValueList `json:"roles" jsonschema:"List of roles to evaluate."`

	// Whether to apply policy to everyone.
	ApplyToEveryone bool `json:"applyToEveryone" jsonschema:"Whether to apply policy to everyone."`

	// Whether to allow or deny users in the roles associated with
	// this policy
	InverseMatch bool `json:"inverseMatch" jsonschema:"Whether to allow or deny users in the roles associated with this policy"`
}

func (rc RoleCriteria) GetBaseAuthenticationCriteriaInfo() BaseAuthenticationInfo {
	return rc.BaseAuthenticationInfo
}

// Role information for RoleCriteria.
type RoleAuthValue struct {
	// The idautoID of the role in RapidIdentity.
	Id string `json:"id" jsonschema:"The idautoID of the role in RapidIdentity."`

	// The name of the role in RapidIdentity.
	Name string `json:"name" jsonschema:"The name of the role in RapidIdentity."`
}

// Determines whether authentication policy is satisfied
// based on the time of the day.
type TimeOfDayCriteria struct {
	BaseAuthenticationInfo

	// The start time to evaluate.
	Start TimeOfDayCriteriaSyntax `json:"start" jsonschema:"The start time to evaluate."`

	// The end time to evaluate.
	End TimeOfDayCriteriaSyntax `json:"end" jsonschema:"The end time to evaluate."`

	// The time zone to evaluate.
	TimeZone string `json:"timeZone" jsonschema:"The time zone to evaluate."`
}

func (tdc TimeOfDayCriteria) GetBaseAuthenticationCriteriaInfo() BaseAuthenticationInfo {
	return tdc.BaseAuthenticationInfo
}

// Time of day value for the TimeOfDayCriteria.
type TimeOfDayCriteriaSyntax struct {
	// Hour from 0 - 23.
	Hour int `json:"hour" jsonschema:"Hour from 0 - 23."`

	// Minute from 0 - 60.
	Minute int `json:"minute" jsonschema:"Minute from 0 - 60."`
}

// Use DUO push or one time code to authenticate.
type DuoMethod struct {
	BaseAuthenticationInfo

	// The configuration ID for the Duo Configuration.
	ConfigId string `json:"configId" jsonschema:"The configuration ID for the Duo Configuration."`

	// Removes additional click to enter DUO prompt.
	AutoProcess bool `json:"autoProcess" jsonschema:"Removes additional click to enter DUO prompt."`
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
	TrustedIdp FederationAuthValue `json:"trustedIdp" jsonschema:"The Trusted IDP configuration to reference"`

	// The url to redirect to once authentication at trusted IDP
	// has been completed.
	PostAuthRedirectUrl string `json:"postAuthRedirectUrl" jsonschema:"The url to redirect to once authentication at trusted IDP has been completed."`

	// Whether to expose atrributes in SAML response.
	ExposeAttributes bool `json:"exposeAttributes" jsonschema:"Whether to expose atrributes in SAML response."`

	// Whether to forward the username to the IDP.
	ForwardUsernameEnabled bool `json:"forwardUsernameEnabled" jsonschema:"Whether to forward the username to the IDP."`

	// The username attribute to forward to the IDP.
	ForwardUsernameAttribute string `json:"forwardUsernameAttribute" jsonschema:"The username attribute to forward to the IDP."`

	// The SAML NAME ID Format of the forward attribute.
	ForwardUsernameNameIDFormat string `json:"forwardUsernameNameIDFormat" jsonschema:"The SAML NAME ID Format of the forward attribute."`
}

func (fm FederationMethod) GetBaseAuthenticationMethodInfo() BaseAuthenticationInfo {
	return fm.BaseAuthenticationInfo
}

// The federation values from the FederationMethod.
type FederationAuthValue struct {

	// The unique ID of the trusted IDP configuration.
	Id string `json:"id" jsonschema:"The unique ID of the trusted IDP configuration."`

	// THe name of the trusted IDP configuration.
	Name string `json:"name" jsonschema:"THe name of the trusted IDP configuration."`
}

// Use FIDO, or device login to authenticate.
type WebAuthnMethod struct {
	BaseAuthenticationInfo

	// Allow login to be remembered for 30 days.
	AllowChallengeDeferral bool `json:"allowChallengeDeferral" jsonschema:"Allow login to be remembered for 30 days."`
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
	ExpirationWarningEnabled bool `json:"expirationWarningEnabled" jsonschema:"Whether to display expiration warning when password is close to expiring."`

	// How many days prior to password expiration to show expiration warning.
	ExpirationWarningDays int `json:"expirationWarningDays" jsonschema:"How many days prior to password expiration to show expiration warning."`

	// How long since the password has been changed.
	CurrentPasswordAgeDays int `json:"currentPasswordAgeDays" jsonschema:"How long since the password has been changed."`

	// The maximum number of days a password can be used before it must be changed.
	PasswordMaximumAgeDays int `json:"passwordMaximumAgeDays" jsonschema:"The maximum number of days a password can be used before it must be changed."`

	// Whether the user must change their password on login.
	MustChange bool `json:"mustChange" jsonschema:"Whether the user must change their password on login."`
}

func (pm PasswordMethod) GetBaseAuthenticationMethodInfo() BaseAuthenticationInfo {
	return pm.BaseAuthenticationInfo
}

// Utilize a pool of images to authenticate.
type PictographMethod struct {
	BaseAuthenticationInfo

	// Number of images to display to user to choose from.
	NumToChallenge int `json:"numToChallenge" jsonschema:"Number of images to display to user to choose from."`

	// The number of images a user must select to authenticate.
	NumToChoose int `json:"numToChoose" jsonschema:"The number of images a user must select to authenticate."`

	// Whether to use the default pool of images or custom images.
	UseDefaultImagePool bool `json:"useDefaultImagePool" jsonschema:"Whether to use the default pool of images or custom images."`

	// The imageIds to use if custom image pool is used.
	ImageIds StringList `json:"imageIds" jsonschema:"The imageIds to use if custom image pool is used."`
}

func (pm PictographMethod) GetBaseAuthenticationMethodInfo() BaseAuthenticationInfo {
	return pm.BaseAuthenticationInfo
}

// Utilize the RapidIdentity Mobile app to authenticate.
// via push or one time code.
type PingMeMethod struct {
	BaseAuthenticationInfo

	// Whether to use cloud based pingMe (Legacy).
	NativePingMe bool `json:"nativePingMe" jsonschema:"Whether to use cloud based pingMe (Legacy)."`

	// A friendly description for the service.
	ServiceDescription string `json:"serviceDescription" jsonschema:"A friendly description for the service."`
}

func (pmm PingMeMethod) GetBaseAuthenticationMethodInfo() BaseAuthenticationInfo {
	return pmm.BaseAuthenticationInfo
}

// Utilize challenge questions and answers to authenticate.
type RapidPortalChallengeMethod struct {
	BaseAuthenticationInfo

	// The RapidIdentity Portal base URL.
	RapidPortalBaseUrl string `json:"rapidPortalBaseUrl" jsonschema:"The RapidIdentity Portal base URL."`
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
	Apple SocialProviderAppleInfo `json:"apple" jsonschema:"Apple provider information."`

	// Google provider information.
	GooglePlus SocialProviderGoogleInfo `json:"googlePlus" jsonschema:"Google provider information."`
}

func (sm SocialMethod) GetBaseAuthenticationMethodInfo() BaseAuthenticationInfo {
	return sm.BaseAuthenticationInfo
}

// Apple social provider information.
type SocialProviderAppleInfo struct {
	// Whether Apple is enabled to be used.
	Enabled bool `json:"enabled" jsonschema:"Whether Apple is enabled to be used."`

	// The private key associated with the Apple social provider.
	PrivateKey string `json:"privateKey" jsonschema:"The private key associated with the Apple social provider."`
}

type SocialProviderGoogleInfo struct {
	// Whether Google is enabled to be used
	Enabled bool `json:"enabled" jsonschema:"Whether Google is enabled to be used"`

	// The client secret associated with the Google social provider.
	ClientSecret string `json:"clientSecret" jsonschema:"The client secret associated with the Google social provider."`
}

// Authenticate with one time code to TOTP application.
type TotpMethod struct {
	BaseAuthenticationInfo

	// The window size for the registration QR Code.
	TotpWindowSize int `json:"totpWindowSize" jsonschema:"The window size for the registration QR Code."`

	// Only challenge user with one time code every 30 days.
	AllowChallengeDeferral bool `json:"allowChallengeDeferral" jsonschema:"Only challenge user with one time code every 30 days."`

	// The issuer name for the TOTP code.
	IssuerName string `json:"issuerName" jsonschema:"The issuer name for the TOTP code."`

	// The setup instructions to display during registration.
	SetupInstructions string `json:"setupInstructions" jsonschema:"The setup instructions to display during registration."`
}

func (tm TotpMethod) GetBaseAuthenticationMethodInfo() BaseAuthenticationInfo {
	return tm.BaseAuthenticationInfo
}

// User agreement to display on login.
type UserAgreementMethod struct {
	BaseAuthenticationInfo

	// The unique id that references the User Agreement.
	UserAgreementId string `json:"userAgreementId" jsonschema:"The unique id that references the User Agreement."`

	// Whether to show user agreement on every login or just one time.
	ShowuserAgreementEveryTime bool `json:"showuserAgreementEveryTime" jsonschema:"Whether to show user agreement on every login or just one time."`
}

func (uam UserAgreementMethod) GetBaseAuthenticationMethodInfo() BaseAuthenticationInfo {
	return uam.BaseAuthenticationInfo
}

type GetBootstrapInfoOutput struct {
	// RapidIdentity license information.
	LicenseInfo LicenseInfo `json:"licenseInfo" jsonschema:"RapidIdentity license information."`

	// RapidIdentity version information.
	VersionInfo VersionInfo `json:"versionInfo" jsonschema:"RapidIdentity version information."`

	// RapidIdentity session information from the invoking user.
	SessionInfo SessionInfo `json:"sessionInfo" jsonschema:"RapidIdentity session information from the invoking user."`

	// RapidIdentity module information.
	ModuleInfo ModuleInfo `json:"moduleInfo" jsonschema:"RapidIdentity module information."`

	// RapidIdentity Portal UI appearance information.
	UIInfo UIInfo `json:"uiInfo" jsonschema:"RapidIdentity Portal UI appearance information."`

	// The default RapidIdentity module page that a user lands after login.
	DefaultLandingModule string `json:"defaultLandingModule" jsonschema:"The default RapidIdentity module page that a user lands after login."`

	// Whether the RapidIdentity tenant is restarted after an upgrade is applied.
	DisableUpgradesRestarts bool `json:"disableUpgradesRestarts" jsonschema:"Whether the RapidIdentity tenant is restarted after an upgrade is applied."`

	// Whether the RapidIdentity ProxyAs feature can be enabled for non admins.
	AllowProxy bool `json:"allowProxy" jsonschema:"Whether the RapidIdentity ProxyAs feature can be enabled for non admins."`

	// Whether the tenant is a RapidIdentity Cloud tenant.
	Idaas bool `json:"idaas" jsonschema:"Whether the tenant is a RapidIdentity Cloud tenant."`

	// The RapidIdentity Tenant ID.
	TenantId string `json:"tenantId" jsonschema:"The RapidIdentity Tenant ID."`

	// Whether notifications are enabled for the tenant.
	NotificationsEnabled bool `json:"notificationsEnabled" jsonschema:"Whether notifications are enabled for the tenant."`

	// Whether Global Search is enabled.
	GlobalSearchEnabled bool `json:"globalSearchEnabled" jsonschema:"Whether Global Search is enabled."`

	// Whether the LDAP directory is within the RapidIdentity cloud tenant and not external.
	IsRICloudLdap bool `json:"isRICloudLdap" jsonschema:"Whether the LDAP directory is within the RapidIdentity cloud tenant and not external."`

	// Information on the RapidIdentity features applied.
	Features FeatureInfo `json:"features" jsonschema:"Information on the RapidIdentity features applied."`

	// The url for retrieving studio library apps.
	DepotProxy DepotProxyInfo `json:"depotProxy" jsonschema:"The url for retrieving studio library apps."`

	// Whether the RapidIdentity tenant has personas enabled
	HasPersonas bool `json:"hasPersonas" jsonschema:"Whether the RapidIdentity tenant has personas enabled"`

	// ShieldID information for the RapidIdentity tenant
	ShieldIdInfo ShieldIdInfo `json:"shieldIdInfo" jsonschema:"ShieldID information for the RapidIdentity tenant"`

	// Whether the tenant is an ID Hub enabled tenant.
	IdHub bool `json:"idHub" jsonschema:"Whether the tenant is an ID Hub enabled tenant."`

	// Whether the new Google Social login is enabled.
	IsIdAutoGoogleEnabled bool `json:"isIdAutoGoogleEnabled" jsonschema:"Whether the new Google Social login is enabled."`

	// Whether the new Apple Social login is enabled.
	IsIdAutoAppleEnabled bool `json:"isIdAutoAppleEnabled" jsonschema:"Whether the new Apple Social login is enabled."`
}

type LicenseInfo struct {
	// The license type such as subscription.
	Type string `json:"type" jsonschema:"The license type such as subscription."`

	// The display name for the person or entity
	// that the license is owned by.
	Licensee string `json:"licensee" jsonschema:"The display name for the person or entity that the license is owned by."`

	// The unique ID for the license.
	LicenseeId string `json:"licenseeId" jsonschema:"The unique ID for the license."`

	// The date the license expires.
	ExpirationDate string `json:"expirationDate" jsonschema:"The date the license expires."`

	// The cluster count.
	ClusterCount int `json:"clusterCount" jsonschema:"The cluster count."`

	// The number of licensed users.
	LicensedUserCount int `json:"licensedUserCount" jsonschema:"The number of licensed users."`

	// The modules provided by the license (Legacy)
	Modules StringList `json:"modules" jsonschema:"The modules provided by the license (Legacy)"`
}

type VersionInfo struct {
	// The RapidIdentity version number.
	Version string `json:"version" jsonschema:"The RapidIdentity version number."`

	// The timestamp for the RapidIdentity version build.
	BuildTimestamp time.Time `json:"buildTimestamp" jsonschema:"The timestamp for the RapidIdentity version build."`
}

type ModuleInfo struct {
	// RapidIdentity Applications module information.
	Applications ApplicationModuleInfo `json:"applications" jsonschema:"RapidIdentity Applications module information."`

	// RapidIdentity Dashboard module information.
	Dashboard DashboardModuleInfo `json:"dashboard" jsonschema:"RapidIdentity Dashboard module information."`

	// RapidIdentity SSO Portal (Personas) information.
	DashboardV2 DashboardV2ModuleInfo `json:"dashboard_V2" jsonschema:"RapidIdentity SSO Portal (Personas) information."`

	// RapidIdentity Files module information.
	Files FileModuleInfo `json:"files" jsonschema:"RapidIdentity Files module information."`

	// RapidIdentity People module information.
	Profiles ProfileModuleInfo `json:"profiles" jsonschema:"RapidIdentity People module information."`

	// RapidIdentity Reports module information.
	Reporting ReportingModuleInfo `json:"reporting" jsonschema:"RapidIdentity Reports module information."`

	// RapidIdentity Groups module information.
	Roles RolesModuleInfo `json:"roles" jsonschema:"RapidIdentity Groups module information."`

	// RapidIdentity Sponsorship module information.
	Sponsorship SponsorshipModuleInfo `json:"sponsorship" jsonschema:"RapidIdentity Sponsorship module information."`

	// RapidIdentity Requests module information.
	Workflow WorkflowModuleInfo `json:"workflow" jsonschema:"RapidIdentity Requests module information."`

	// RapidIdentiy Admin information.
	Admin AdminModuleInfo `json:"admin" jsonschema:"RapidIdentiy Admin information."`

	// RapidIdentity Connect module information.
	Connect ConnectModuleInfo `json:"connect" jsonschema:"RapidIdentity Connect module information."`

	// RapidIdentity Studio module information.
	Studio StudioModuleInfo `json:"studio" jsonschema:"RapidIdentity Studio module information."`

	// RapidIdentity folders module information.
	Folders FoldersModuleInfo `json:"folders" jsonschema:"RapidIdentity folders module information."`

	// RapidIdentity Insights module information.
	Insights InsightsModuleInfo `json:"insights" jsonschema:"RapidIdentity Insights module information."`

	// RapidIdentity configuration module information.
	Configuration ConfigurationModuleInfo `json:"configuration" jsonschema:"RapidIdentity configuration module information."`

	// RapidIdentity IdHub information.
	IdHub IdHubModuleInfo `json:"idHub" jsonschema:"RapidIdentity IdHub information."`
}

type ApplicationModuleInfo struct {
	ModuleLicenseInfo

	// My tab visibility information.
	MyTabInfo TabInfo `json:"myTabInfo" jsonschema:"My tab visibility information."`

	// Team tab visibility information.
	TeamTabInfo TabInfo `json:"teamTabInfo" jsonschema:"Team tab visibility information."`

	// Other tab visibility information.
	OtherTabInfo TabInfo `json:"otherTabInfo" jsonschema:"Other tab visibility information."`

	// Preferences information.
	Preferences PreferenceInfo `json:"preferences" jsonschema:"Preferences information."`
}

type ModuleLicenseInfo struct {
	// Whether RapidIdentity module is licensed.
	Licensed bool `json:"licensed" jsonschema:"Whether RapidIdentity module is licensed."`

	// Whether RapidIdentity module is visible for the invoking user.
	Visible bool `json:"visible" jsonschema:"Whether RapidIdentity module is visible for the invoking user."`
}

type TabInfo struct {
	// Whether tab is visible for the invoking user.
	Visible bool `json:"visible" jsonschema:"Whether tab is visible for the invoking user."`
}

type TabActionInfo struct {
	TabInfo
	// Available actions within the specified module tab.
	Actions StringList `json:"actions" jsonschema:"Available actions within the specified module tab."`
}

type TabAdminInfo struct {
	// Whether config is visible for the invoking user.
	ConfigTabVisible bool `json:"configTabVisible" jsonschema:"Whether config is visible for the invoking user."`

	// Whether admin tab is visible for the invoking user.
	AdminTabVisible bool `json:"adminTabVisible" jsonschema:"Whether admin tab is visible for the invoking user."`
}

type PreferenceInfo struct {
	// Whether to start on favorites in the applications module.
	StartAtFavorites bool `json:"startAtFavorites" jsonschema:"Whether to start on favorites in the applications module."`
}

type DashboardModuleInfo struct {
	ModuleLicenseInfo
	// Visibility information for the My Activity Tab
	MyActivityTab TabInfo `json:"myActivityTab" jsonschema:"Visibility information for the My Activity Tab"`

	// Visibility information for the Team Activity Tab
	TeamActivityTab TabInfo `json:"teamActivityTab" jsonschema:"Visibility information for the Team Activity Tab"`

	// Visibility information for the Other Activity Tab
	OtherActivityTab TabInfo `json:"otherActivityTab" jsonschema:"Visibility information for the Other Activity Tab"`

	// Visibility information for the My Entitlements Tab
	MyEntitlementsTab TabInfo `json:"myEntitlementsTab" jsonschema:"Visibility information for the My Entitlements Tab"`

	// Visibility information for the Team Entitlements Tab
	TeamEntitlementsTab TabInfo `json:"teamEntitlementsTab" jsonschema:"Visibility information for the Team Entitlements Tab"`

	// Visibility information for the Other Entitlements Tab
	OtherEntitlementsTab TabInfo `json:"otherEntitlementsTab" jsonschema:"Visibility information for the Other Entitlements Tab"`

	// Visibility information for the Executive Tab
	ExecutiveTab TabInfo `json:"executiveTab" jsonschema:"Visibility information for the Executive Tab"`
}

type DashboardV2ModuleInfo struct {
	ModuleLicenseInfo
}

type FileModuleInfo struct {
	ModuleLicenseInfo
	// The maximum file size, in MB, that can be uploaded.
	MaxUploadSize float32 `json:"maxUploadSize" jsonschema:"The maximum file size, in MB, that can be uploaded."`

	// Whether SSL upload is enabled.
	EnableSSLUpload bool `json:"enableSSLUpload" jsonschema:"Whether SSL upload is enabled."`

	// Whether public access is enabled.
	EnableMakePublic bool `json:"enableMakePublic" jsonschema:"Whether public access is enabled."`
}

type ProfileModuleInfo struct {
	ModuleLicenseInfo

	// Whether challenge questions are enabled (Legacy).
	ChallengeQuestionsEnabled bool `json:"challengeQuestionsEnabled" jsonschema:"Whether challenge questions are enabled (Legacy)."`

	// Whether the invoking user must update their challenge questions.
	MustUpdateChallengeQuestions bool `json:"mustUpdateChallengeQuestions" jsonschema:"Whether the invoking user must update their challenge questions."`

	// The invalid challenge set message (Legacy).
	InvalidChallengeSetMessage string `json:"invalidChallengeSetMessage" jsonschema:"The invalid challenge set message (Legacy)."`

	// Whether show all is enabled.
	EnableShowAll bool `json:"enableShowAll" jsonschema:"Whether show all is enabled."`

	// Whether must change password options are supported.
	MustChangePasswordOptionsSupported bool `json:"mustChangePasswordOptionsSupported" jsonschema:"Whether must change password options are supported."`

	// Whether an unlock requires a password reset.
	UnlockRequiresPasswordReset bool `json:"unlockRequiresPasswordReset" jsonschema:"Whether an unlock requires a password reset."`

	// Whether a delegation attribute is required.
	DelegationAttrRequired bool `json:"delegationAttrRequired" jsonschema:"Whether a delegation attribute is required."`
}

type ReportingModuleInfo struct {
	ModuleLicenseInfo

	// The max results returned from a Reports module audit report query.
	AuditReportMax int `json:"auditReportMax" jsonschema:"The max results returned from a Reports module audit report query."`
}

type RolesModuleInfo struct {
	ModuleLicenseInfo
	// Whether to enable selection of group types.
	EnableGroupTypeSelection bool `json:"enableGroupTypeSelection" jsonschema:"Whether to enable selection of group types."`

	// Available group types to choose form if group type selection is enabled.
	PossibleGroupTypes StringList `json:"possibleGroupTypes" jsonschema:"Available group types to choose form if group type selection is enabled."`

	// Allowed group types to choose from if group type selection is enabled.
	AllowedGroupTypes StringList `json:"allowedGroupTypes" jsonschema:"Allowed group types to choose from if group type selection is enabled."`

	// Whether to preload all groups in Groups module.
	PreloadGroups bool `json:"preloadGroups" jsonschema:"Whether to preload all groups in Groups module."`

	// Whether to show the distinguished name of the RapidIdentity group.
	ShowDN bool `json:"showDN" jsonschema:"Whether to show the distinguished name of the RapidIdentity group."`

	// Visibility and actions for the My Groups tab.
	MyTabInfo TabActionInfo `json:"myTabInfo" jsonschema:"Visibility and actions for the My Groups tab."`

	// Visibility and actions for the Team Groups tab.
	TeamTabInfo TabActionInfo `json:"teamTabInfo" jsonschema:"Visibility and actions for the Team Groups tab."`

	// Visibility and actions for the Other Groups tab.
	OtherTabInfo TabActionInfo `json:"otherTabInfo" jsonschema:"Visibility and actions for the Other Groups tab."`

	// Any custom attributes added to the groups module.
	CustomAttributes DelegationAttributeList `json:"customAttributes" jsonschema:"Any custom attributes added to the groups module."`

	// Whether wildcard (*) searches are allowed.
	EnableWildcardSearch bool `json:"enableWildcardSearch" jsonschema:"Whether wildcard (*) searches are allowed."`
}

type SponsorshipModuleInfo struct {
	ModuleLicenseInfo

	// Visibility and actions for the My Sponsored Accounts tab.
	MyTabInfo TabActionInfo `json:"myTabInfo" jsonschema:"Visibility and actions for the My Sponsored Accounts tab."`

	// Visibility and actions for the Team Sponsored Accounts tab.
	TeamTabInfo TabActionInfo `json:"teamTabInfo" jsonschema:"Visibility and actions for the Team Sponsored Accounts tab."`

	// Visibility and actions for the Other Sponsored Accounts tab.
	OtherTabInfo TabActionInfo `json:"otherTabInfo" jsonschema:"Visibility and actions for the Other Sponsored Accounts tab."`

	// The max expiration date you can select when creating or certifying a sponsored account.
	MaxExpirationDays int `json:"maxExpirationDays" jsonschema:"The max expiration date you can select when creating or certifying a sponsored account."`

	// Whether an email address is required for a sponsored account.
	EmailAddressRequired bool `json:"emailAddressRequired" jsonschema:"Whether an email address is required for a sponsored account."`

	// Whether an expiration date is required for a sponsored account.
	ExpirationRequired bool `json:"expirationRequired" jsonschema:"Whether an expiration date is required for a sponsored account."`

	// Whether to preload all sponsors.
	PreloadSponsors bool `json:"preloadSponsors" jsonschema:"Whether to preload all sponsors."`

	// Whether to load all sponsored accounts upon accessing the Sponsorship module.
	PreloadSponsoredAccounts bool `json:"preloadSponsoredAccounts" jsonschema:"Whether to load all sponsored accounts upon accessing the Sponsorship module."`

	// Custom attributes added to the sponsorship module.
	CustomAttributes DelegationAttributeList `json:"customAttributes" jsonschema:"Custom attributes added to the sponsorship module."`
}

type WorkflowModuleInfo struct {
	ModuleLicenseInfo

	// Whether SSL uploads are required.
	EnableSSLUpload bool `json:"enableSSLUpload" jsonschema:"Whether SSL uploads are required."`

	// Visibility for My Dashboard tab.
	MyDashboardTabInfo TabInfo `json:"myDashboardTabInfo" jsonschema:"Visibility for My Dashboard tab."`

	// Visibility for Team Dashboard tab.
	TeamDashboardTabInfo TabInfo `json:"teamDashboardTabInfo" jsonschema:"Visibility for Team Dashboard tab."`

	// Visibility for Other Dashboard tab.
	OtherDashboardTabInfo TabInfo `json:"otherDashboardTabInfo" jsonschema:"Visibility for Other Dashboard tab."`

	// Visibility for My Entitlements tab.
	MyRequestsTabInfo TabInfo `json:"myRequestsTabInfo" jsonschema:"Visibility for My Entitlements tab."`

	// Visibility for Team Entitlements tab.
	TeamRequestsTabInfo TabInfo `json:"teamRequestsTabInfo" jsonschema:"Visibility for Team Entitlements tab."`

	// Visibility for Other Entitlements tab.
	OtherRequestsTabInfo TabInfo `json:"otherRequestsTabInfo" jsonschema:"Visibility for Other Entitlements tab."`

	// Visibility information for My Task Approvals.
	MyApprovalsTabInfo TabInfo `json:"myApprovalsTabInfo" jsonschema:"Visibility information for My Task Approvals."`

	// Visibility information for Team Task Approvals.
	TeamApprovalsTabInfo TabInfo `json:"teamApprovalsTabInfo" jsonschema:"Visibility information for Team Task Approvals."`

	// Visibility information for Other Task Approvals.
	OtherApprovalsTabInfo TabInfo `json:"otherApprovalsTabInfo" jsonschema:"Visibility information for Other Task Approvals."`

	// Visibility information for Certification tasks.
	UpcomingCertificationsTabInfo TabInfo `json:"upcomingCertificationsTabInfo" jsonschema:"Visibility information for Certification tasks."`

	// Visibility information for searching Certification tasks.
	SearchCertificationsTabInfo TabInfo `json:"searchCertificationsTabInfo" jsonschema:"Visibility information for searching Certification tasks."`

	// Visibility information for task manager tab..
	TaskManagerTabInfo TabInfo `json:"taskManagerTabInfo" jsonschema:"Visibility information for task manager tab.."`
}

type AdminModuleInfo struct {
	// Whether configuration is visible for invoking user.
	Visible bool `json:"visible" jsonschema:"Whether configuration is visible for invoking user."`

	// Portal admin tab info.
	Portal TabAdminInfo `json:"portal" jsonschema:"Portal admin tab info."`

	// Application admin tab info.
	Applications TabAdminInfo `json:"applications" jsonschema:"Application admin tab info."`

	// Dashboard admin tab info.
	Dashboard TabAdminInfo `json:"dashboard" jsonschema:"Dashboard admin tab info."`

	// Files admin tab info.
	Files TabAdminInfo `json:"files" jsonschema:"Files admin tab info."`

	// People admin tab info.
	Profiles TabAdminInfo `json:"profiles" jsonschema:"People admin tab info."`

	// Reports admin tab info.
	Reporting TabAdminInfo `json:"reporting" jsonschema:"Reports admin tab info."`

	// Groups admin tab info.
	Roles TabAdminInfo `json:"roles" jsonschema:"Groups admin tab info."`

	// Sponsorship admin tab info.
	Sponsorship TabAdminInfo `json:"sponsorship" jsonschema:"Sponsorship admin tab info."`

	// Workflow admin tab info.
	Workflow TabAdminInfo `json:"workflow" jsonschema:"Workflow admin tab info."`
}

type ConnectModuleInfo struct {
	// Whether the Connect module is visible for the invoking user.
	Visible bool `json:"visible" jsonschema:"Whether the Connect module is visible for the invoking user."`
}

type StudioModuleInfo struct {
	ModuleLicenseInfo
	// Whether the invoking user is a Studio operator.
	IsOperator bool `json:"isOperator" jsonschema:"Whether the invoking user is a Studio operator."`

	// Whether the invoking user is a Studio administrator.
	IsAdmin bool `json:"isAdmin" jsonschema:"Whether the invoking user is a Studio administrator."`
}

type FoldersModuleInfo struct {
	// Whether folder users are visible for the invoking user.
	UsersVisible bool `json:"usersVisible" jsonschema:"Whether folder users are visible for the invoking user."`

	// Whether folder groups are visible for the invoking user.
	GroupsVisible bool `json:"groupsVisible" jsonschema:"Whether folder groups are visible for the invoking user."`

	// Whether the invoking user is a Folders operator.
	IsOperator bool `json:"isOperator" jsonschema:"Whether the invoking user is a Folders operator."`

	// Whether the invoking user is a Folders administrator.
	IsAdmin bool `json:"isAdmin" jsonschema:"Whether the invoking user is a Folders administrator."`

	// Whether the folers shcema is up to date.
	SchemaUpToDate bool `json:"schemaUpToDate" jsonschema:"Whether the folers shcema is up to date."`
}

type InsightsModuleInfo struct {
	ModuleLicenseInfo

	// Whether the invoking user is an Insights manager.
	IsManager bool `json:"isManager" jsonschema:"Whether the invoking user is an Insights manager."`

	// Whether the invoking user is an Insights viewer.
	IsViewer bool `json:"isViewer" jsonschema:"Whether the invoking user is an Insights viewer."`
}

type ConfigurationModuleInfo struct {
	ModuleLicenseInfo

	// Visibility information for the audit tab within configuration.
	AuditTabInfo TabInfo `json:"auditTabInfo" jsonschema:"Visibility information for the audit tab within configuration."`
}

type IdHubModuleInfo struct {
	// Whether ID Hub is visible for the invoking user.
	Visible bool `json:"visible" jsonschema:"Whether ID Hub is visible for the invoking user."`

	// LCS External Auth Client Id.
	LcsExternalAuthClientId string `json:"lcsExternalAuthClientId" jsonschema:"LCS External Auth Client Id."`

	// The LCS Domain.
	LcsDomain string `json:"lcsDomain" jsonschema:"The LCS Domain."`

	// The catalog domain.
	CatalogDomain string `json:"catalogDomain" jsonschema:"The catalog domain."`
}

type UIInfo struct {
	// The portal logo url.
	LogoUrl string `json:"logoUrl" jsonschema:"The portal logo url."`

	// The portal background gradient color 1.
	BackgroundGradient1 string `json:"backgroundGradient1" jsonschema:"The portal background gradient color 1."`

	// The portal background gradient color 2.
	BackgroundGradient2 string `json:"backgroundGradient2" jsonschema:"The portal background gradient color 2."`

	// The portal wide logo url.
	WideLogoURL string `json:"wideLogoURL" jsonschema:"The portal wide logo url."`

	// The portal narrow logo url.
	NarrowLogoURL string `json:"narrowLogoURL" jsonschema:"The portal narrow logo url."`

	// The portal favicon url.
	FaviconURL string `json:"faviconURL" jsonschema:"The portal favicon url."`

	// The portal brand color 1.
	BrandColorOne string `json:"brandColorOne" jsonschema:"The portal brand color 1."`

	// The portal brand color 2.
	BrandColorTwo string `json:"brandColorTwo" jsonschema:"The portal brand color 2."`
}

type FeatureInfo struct {
	// Whether login configs is enabled.
	LoginConfigs bool `json:"loginConfigs" jsonschema:"Whether login configs is enabled."`

	// Pendo information for RapidIdentity tenant.
	Pendo PendoInfo `json:"pendo" jsonschema:"Pendo information for RapidIdentity tenant."`

	// ChurnZero information for RapidIdentity tenant.
	ChurnZero ChurnZeroInfo `json:"churnZero" jsonschema:"ChurnZero information for RapidIdentity tenant."`

	// Whether personas is enabled.
	SsoPortal bool `json:"ssoPortal" jsonschema:"Whether personas is enabled."`

	// Whether universal authentication director is available.
	ThirdPartyPortal bool `json:"thirdPartyPortal" jsonschema:"Whether universal authentication director is available."`

	// Whether SafeID is available.
	SafeId bool `json:"safeId" jsonschema:"Whether SafeID is available."`

	// Whether ShieldID is available.
	ShieldId bool `json:"shieldId" jsonschema:"Whether ShieldID is available."`

	// Whether ID Hub is available.
	IdHub bool `json:"idHub" jsonschema:"Whether ID Hub is available."`

	// Whether Password Vault is available.
	PasswordVault bool `json:"passwordVault" jsonschema:"Whether Password Vault is available."`

	// Whether ProxyAs is available for non admins.
	ProxyAs bool `json:"proxyAs" jsonschema:"Whether ProxyAs is available for non admins."`
}

type PendoInfo struct {
	// Unique Pendo ID for the invoking user.
	Id string `json:"id" jsonschema:"Unique Pendo ID for the invoking user."`

	// UserType of the invoking user.
	UserType string `json:"userType" jsonschema:"UserType of the invoking user."`
}

type ChurnZeroInfo struct {
	// Unique Churn Zero ID for the invoking user.
	Id string `json:"id" jsonschema:"Unique Churn Zero ID for the invoking user."`

	// User type of the invoking user.
	UserType string `json:"userType" jsonschema:"User type of the invoking user."`

	// User role of the invoking user.
	UserRole string `json:"userRole" jsonschema:"User role of the invoking user."`
}

type DepotProxyInfo struct {
	// The Studio library url.
	VettedStudioAppsUrl string `json:"vettedStudioAppsUrl" jsonschema:"The Studio library url."`
}

type ShieldIdInfo struct {
	// API Domain
	ApiDomain string `json:"apiDomain" jsonschema:"API Domain"`

	// CLient ID
	ClientId string `json:"clientId" jsonschema:"CLient ID"`

	// Host ID
	HostId string `json:"hostId" jsonschema:"Host ID"`
}

// Retrieves authentication policies for specified user.
//
//meta:operation POST /authn/v1/username
func (c *Client) GetAuthenticationPoliciesForUser(ctx context.Context, params GetAuthenticationPoliciesForUserInput) (*GetAuthenticationPoliciesForUserOutput, error) {
	url := fmt.Sprintf("%s/authn/v1/username?authenticationPolicies=%t&claim=%t", c.baseEndpoint, params.ShowAuthenticationPolicies, params.ShowClaims)
	for _, field := range params.AuthenticationPolicyFieldsToShow {
		url = fmt.Sprintf("%s&authenticationPolicyField=%s", url, field)
	}
	user, err := json.Marshal(params.User)
	if err != nil {
		return nil, err
	}
	requestBody := bytes.NewBuffer(user)
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
	var output GetAuthenticationPoliciesForUserOutput
	err = json.Unmarshal(resBody, &output)
	if err != nil {
		return nil, err
	}

	return &output, nil
}

// Retrieves RapidIdentity tenant and user access information for the invoking user.
//
//meta:operation GET /bootstrapInfo
func (c *Client) GetBootstrapInfo(ctx context.Context) (*GetBootstrapInfoOutput, error) {
	var output GetBootstrapInfoOutput

	url := fmt.Sprintf("%s/bootstrapInfo", c.baseEndpoint)
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

// Retrieves RapidIdentity LDAP attributes.
//
//meta:operation GET /admin/ldap/schema/attributes
func (c *Client) GetRapidIdentityAttributes(ctx context.Context) (StringList, error) {
	var output StringList

	url := fmt.Sprintf("%s/admin/ldap/schema/attributes", c.baseEndpoint)
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

	return output, nil
}
