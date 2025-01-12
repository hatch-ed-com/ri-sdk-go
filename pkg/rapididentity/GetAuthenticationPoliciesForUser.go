package rapididentity

import (
	"bytes"
	"encoding/json"
	"fmt"
)

type GetAuthenticationPoliciesForUserInput struct {
	ShowAuthenticationPolicies       bool
	ShowClaims                       bool
	AuthenticationPolicyFieldsToShow []string
	User                             GetAuthenticationPoliciesForUserPayload
}

type GetAuthenticationPoliciesForUserPayload struct {
	Username string `json:"username"`
}

type GetAuthenticationPoliciesForUserOutput struct {
	User                   User                   `json:"user"`
	AuthenticationPolicies []AuthenticationPolicy `json:"authenticationPolicies"`
}

type AuthenticationPolicy struct {
	Id                    string                         `json:"id"`
	Version               int                            `json:"version"`
	Name                  string                         `json:"name"`
	Enabled               bool                           `json:"enabled"`
	Criteria              []AuthenticationPolicyCriteria `json:"criteria"`
	Methods               []AuthenticationPolicyMethod   `json:"methods"`
	InsecureQRIdEnabled   bool                           `json:"insecureQRIdEnabled"`
	AlwaysFail            bool                           `json:"alwaysFail"`
	IsResetPasswordPolicy bool                           `json:"isResetPasswordPolicy"`
}

type AuthenticationPolicyCriteria interface {
	GetBaseAuthenticationCriteriaInfo() BaseAuthenticationInfo
}

type AuthenticationPolicyMethod interface {
	GetBaseAuthenticationMethodInfo() BaseAuthenticationInfo
}

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

type BaseAuthenticationInfo struct {
	Type    string `json:"type"`
	Enabled bool   `json:"enabled"`
}

type DaysOfWeekCriteria struct {
	BaseAuthenticationInfo
	Sunday    bool `json:"sunday"`
	Monday    bool `json:"monday"`
	Tuesday   bool `json:"tuesday"`
	Wednesday bool `json:"wednesday"`
	Thursday  bool `json:"thursday"`
	Friday    bool `json:"friday"`
	Saturday  bool `json:"saturday"`
}

func (dwc DaysOfWeekCriteria) GetBaseAuthenticationCriteriaInfo() BaseAuthenticationInfo {
	return dwc.BaseAuthenticationInfo
}

type WebAuthnCriteria struct {
	BaseAuthenticationInfo
	Negate bool `json:"negate"`
}

func (wac WebAuthnCriteria) GetBaseAuthenticationCriteriaInfo() BaseAuthenticationInfo {
	return wac.BaseAuthenticationInfo
}

type KerberosCriteria struct {
	BaseAuthenticationInfo
}

func (kc KerberosCriteria) GetBaseAuthenticationCriteriaInfo() BaseAuthenticationInfo {
	return kc.BaseAuthenticationInfo
}

type LdapFilterCriteria struct {
	BaseAuthenticationInfo
	LdapFilter        string `json:"ldapFilter"`
	MatchNonLdapAdmin bool   `json:"matchNonLdapAdmin"`
}

func (lfc LdapFilterCriteria) GetBaseAuthenticationCriteriaInfo() BaseAuthenticationInfo {
	return lfc.BaseAuthenticationInfo
}

type QrCodeCriteria struct {
	BaseAuthenticationInfo
}

func (qcc QrCodeCriteria) GetBaseAuthenticationCriteriaInfo() BaseAuthenticationInfo {
	return qcc.BaseAuthenticationInfo
}

type SourceNetworkCriteria struct {
	BaseAuthenticationInfo
	Subnets                    []string `json:"subnets"`
	EnableHttpHeaderProcessing bool     `json:"enableHttpHeaderProcessing"`
	Negate                     bool     `json:"negate"`
}

func (snc SourceNetworkCriteria) GetBaseAuthenticationCriteriaInfo() BaseAuthenticationInfo {
	return snc.BaseAuthenticationInfo
}

type RoleCriteria struct {
	BaseAuthenticationInfo
	Roles           []RoleAuthValue `json:"roles"`
	ApplyToEveryone bool            `json:"applyToEveryone"`
	InverseMatch    bool            `json:"inverseMatch"`
}

func (rc RoleCriteria) GetBaseAuthenticationCriteriaInfo() BaseAuthenticationInfo {
	return rc.BaseAuthenticationInfo
}

type RoleAuthValue struct {
	Id   string `json:"id"`
	Name string `json:"name"`
}

type TimeOfDayCriteria struct {
	BaseAuthenticationInfo
	Start    TimeOfDayCriteriaSyntax
	End      TimeOfDayCriteriaSyntax
	TimeZone string
}

func (tdc TimeOfDayCriteria) GetBaseAuthenticationCriteriaInfo() BaseAuthenticationInfo {
	return tdc.BaseAuthenticationInfo
}

type TimeOfDayCriteriaSyntax struct {
	Hour   int `json:"hour"`
	Minute int `json:"minute"`
}

type DuoMethod struct {
	BaseAuthenticationInfo
	ConfigId    string `json:"configId"`
	AutoProcess bool   `json:"autoProcess"`
}

func (dm DuoMethod) GetBaseAuthenticationMethodInfo() BaseAuthenticationInfo {
	return dm.BaseAuthenticationInfo
}

type EmailMethod struct {
	BaseAuthenticationInfo
}

func (em EmailMethod) GetBaseAuthenticationMethodInfo() BaseAuthenticationInfo {
	return em.BaseAuthenticationInfo
}

type FederationMethod struct {
	BaseAuthenticationInfo
	TrustedIdp                  FederationAuthValue `json:"trustedIdp"`
	PostAuthRedirectUrl         string              `json:"postAuthRedirectUrl"`
	ExposeAttributes            bool                `json:"exposeAttributes"`
	ForwardUsernameEnabled      bool                `json:"forwardUsernameEnabled"`
	ForwardUsernameAttribute    string              `json:"forwardUsernameAttribute"`
	ForwardUsernameNameIDFormat string              `json:"forwardUsernameNameIDFormat"`
}

func (fm FederationMethod) GetBaseAuthenticationMethodInfo() BaseAuthenticationInfo {
	return fm.BaseAuthenticationInfo
}

type FederationAuthValue struct {
	Id   string `json:"id"`
	Name string `json:"name"`
}

type WebAuthnMethod struct {
	BaseAuthenticationInfo
	AllowChallengeDeferral bool `json:"allowChallengeDeferral"`
}

func (wam WebAuthnMethod) GetBaseAuthenticationMethodInfo() BaseAuthenticationInfo {
	return wam.BaseAuthenticationInfo
}

type KerberosMethod struct {
	BaseAuthenticationInfo
}

func (km KerberosMethod) GetBaseAuthenticationMethodInfo() BaseAuthenticationInfo {
	return km.BaseAuthenticationInfo
}

type PasswordMethod struct {
	BaseAuthenticationInfo
	ExpirationWarningEnabled bool `json:"expirationWarningEnabled"`
	ExpirationWarningDays    int  `json:"expirationWarningDays"`
	CurrentPasswordAgeDays   int  `json:"currentPasswordAgeDays"`
	PasswordMaximumAgeDays   int  `json:"passwordMaximumAgeDays"`
	MustChange               bool `json:"mustChange"`
}

func (pm PasswordMethod) GetBaseAuthenticationMethodInfo() BaseAuthenticationInfo {
	return pm.BaseAuthenticationInfo
}

type PictographMethod struct {
	BaseAuthenticationInfo
	NumToChallenge      int      `json:"numToChallenge"`
	NumToChoose         int      `json:"numToChoose"`
	UseDefaultImagePool bool     `json:"useDefaultImagePool"`
	ImageIds            []string `json:"imageIds"`
}

func (pm PictographMethod) GetBaseAuthenticationMethodInfo() BaseAuthenticationInfo {
	return pm.BaseAuthenticationInfo
}

type PingMeMethod struct {
	BaseAuthenticationInfo
	NativePingMe       bool   `json:"nativePingMe"`
	ServiceDescription string `json:"serviceDescription"`
}

func (pmm PingMeMethod) GetBaseAuthenticationMethodInfo() BaseAuthenticationInfo {
	return pmm.BaseAuthenticationInfo
}

type RapidPortalChallengeMethod struct {
	BaseAuthenticationInfo
	RapidPortalBaseUrl string `json:"rapidPortalBaseUrl"`
}

func (rpcm RapidPortalChallengeMethod) GetBaseAuthenticationMethodInfo() BaseAuthenticationInfo {
	return rpcm.BaseAuthenticationInfo
}

type QrCodeMethod struct {
	BaseAuthenticationInfo
}

func (qcm QrCodeMethod) GetBaseAuthenticationMethodInfo() BaseAuthenticationInfo {
	return qcm.BaseAuthenticationInfo
}

type SmsMethod struct {
	BaseAuthenticationInfo
}

func (sm SmsMethod) GetBaseAuthenticationMethodInfo() BaseAuthenticationInfo {
	return sm.BaseAuthenticationInfo
}

type SocialMethod struct {
	BaseAuthenticationInfo
	Apple      SocialProviderAppleInfo  `json:"apple"`
	GooglePlus SocialProviderGoogleInfo `json:"googlePlus"`
}

func (sm SocialMethod) GetBaseAuthenticationMethodInfo() BaseAuthenticationInfo {
	return sm.BaseAuthenticationInfo
}

type SocialProviderAppleInfo struct {
	Enabled    bool   `json:"enabled"`
	PrivateKey string `json:"privateKey"`
}

type SocialProviderGoogleInfo struct {
	Enabled      bool   `json:"enabled"`
	ClientSecret string `json:"clientSecret"`
}

type TotpMethod struct {
	BaseAuthenticationInfo
	TotpWindowSize         int    `json:"totpWindowSize"`
	AllowChallengeDeferral bool   `json:"allowChallengeDeferral"`
	IssuerName             string `json:"issuerName"`
	SetupInstructions      string `json:"setupInstructions"`
}

func (tm TotpMethod) GetBaseAuthenticationMethodInfo() BaseAuthenticationInfo {
	return tm.BaseAuthenticationInfo
}

type UserAgreementMethod struct {
	BaseAuthenticationInfo
	UserAgreementId            string `json:"userAgreementId"`
	ShowuserAgreementEveryTime bool   `json:"showuserAgreementEveryTime"`
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
