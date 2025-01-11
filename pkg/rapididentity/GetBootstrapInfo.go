package rapididentity

import (
	"encoding/json"
	"fmt"
	"time"
)

type GetBootstrapInfoOutput struct {
	LicenseInfo             LicenseInfo  `json:"licenseInfo"`
	VersionInfo             VersionInfo  `json:"versionInfo"`
	SessionInfo             SessionInfo  `json:"sessionInfo"`
	ModuleInfo              ModuleInfo   `json:"moduleInfo"`
	UIInfo                  UIInfo       `json:"uiInfo"`
	DefaultLandingModule    string       `json:"defaultLandingModule"`
	DisableUpgradesRestarts bool         `json:"disableUpgradesRestarts"`
	AllowProxy              bool         `json:"allowProxy"`
	Idaas                   bool         `json:"idaas"`
	TenantId                string       `json:"tenantId"`
	NotificationEnabled     bool         `json:"notificationEnabled"`
	GlobalSearchEnabled     bool         `json:"globalSearchEnabled"`
	IsRICloudLdap           bool         `json:"isRICloudLdap"`
	Features                FeatureInfo  `json:"features"`
	HasPersonas             bool         `json:"hasPersonas"`
	ShieldIdInfo            ShieldIdInfo `json:"shieldIdInfo"`
	IdHub                   bool         `json:"idHub"`
	IsIdAutoGoogleEnabled   bool         `json:"isIdAutoGoogleEnabled"`
	IsIdAutoAppleEnabled    bool         `json:"isIdAutoAppleEnabled"`
}

type LicenseInfo struct {
	Type              string   `json:"type"`
	Licensee          string   `json:"licensee"`
	LicenseeId        string   `json:"licenseeId"`
	ExpirationDate    string   `json:"expirationDate"`
	ClusterCount      int      `json:"clusterCount"`
	LicensedUserCount int      `json:"licensedUserCount"`
	Modules           []string `json:"modules"`
}

type VersionInfo struct {
	Version        string    `json:"version"`
	BuildTimestamp time.Time `json:"buildTimestamp"`
}

type ModuleInfo struct {
	Applications  ApplicationModuleInfo   `json:"applications"`
	Dashboard     ModuleLicenseInfo       `json:"dashboard"`
	DashboardV2   ModuleLicenseInfo       `json:"dashboard_V2"`
	Files         FileModuleInfo          `json:"files"`
	Profiles      ProfileModuleInfo       `json:"profiles"`
	Reporting     ModuleLicenseInfo       `json:"reporting"`
	Roles         ModuleLicenseInfo       `json:"roles"`
	Sponsorship   ModuleLicenseInfo       `json:"sponsorship"`
	Workflow      WorkflowModuleInfo      `json:"workflow"`
	Admin         ModuleLicenseInfo       `json:"admin"`
	Connect       ModuleLicenseInfo       `json:"connect"`
	Studio        StudioModuleInfo        `json:"studio"`
	Folders       ModuleLicenseInfo       `json:"folders"`
	Insights      InsightsModuleInfo      `json:"insights"`
	Configuration ConfigurationModuleInfo `json:"configuration"`
	IdHub         ModuleLicenseInfo       `json:"idHub"`
}

type ApplicationModuleInfo struct {
	ModuleLicenseInfo
	MyTabInfo    TabInfo        `json:"myTabInfo"`
	TeamTabInfo  TabInfo        `json:"teamTabInfo"`
	OtherTabInfo TabInfo        `json:"otherTabInfo"`
	Preferences  PreferenceInfo `json:"preferences"`
}

type ModuleLicenseInfo struct {
	Licensed bool `json:"licensed"`
	Visible  bool `json:"visible"`
}

type TabInfo struct {
	Visible bool `json:"visible"`
}

type PreferenceInfo struct {
	StartAtFavorites bool `json:"startAtFavorites"`
}

type FileModuleInfo struct {
	ModuleLicenseInfo
	MaxUploadSize    float32 `json:"maxUploadSize"`
	EnableSSLUpload  bool    `json:"enableSSLUpload"`
	EnableMakePublic bool    `json:"enableMakePublic"`
}

type ProfileModuleInfo struct {
	ModuleLicenseInfo
	ChallengeQuestionsEnabled          bool   `json:"challengeQuestionsEnabled"`
	MustUpdateChallengeQuestions       bool   `json:"mustUpdateChallengeQuestions"`
	InvalidChallengeSetMessage         string `json:"invalidChallengeSetMessage"`
	EnableShowAll                      bool   `json:"enableShowAll"`
	MustChangePasswordOptionsSupported bool   `json:"mustChangePasswordOptionsSupported"`
	UnlockRequiresPasswordReset        bool   `json:"unlockRequiresPasswordReset"`
	DelegationAttrRequired             bool   `json:"delegationAttrRequired"`
}

type WorkflowModuleInfo struct {
	ModuleLicenseInfo
	EnableSSLUpload               bool    `json:"enableSSLUpload"`
	MyDashboardTabInfo            TabInfo `json:"myDashboardTabInfo"`
	TeamDashboardTabInfo          TabInfo `json:"teamDashboardTabInfo"`
	OtherDashboardTabInfo         TabInfo `json:"otherDashboardTabInfo"`
	MyRequestsTabInfo             TabInfo `json:"myRequestsTabInfo"`
	TeamRequestsTabInfo           TabInfo `json:"teamRequestsTabInfo"`
	OtherRequestsTabInfo          TabInfo `json:"otherRequestsTabInfo"`
	MyApprovalsTabInfo            TabInfo `json:"myApprovalsTabInfo"`
	TeamApprovalsTabInfo          TabInfo `json:"teamApprovalsTabInfo"`
	OtherApprovalsTabInfo         TabInfo `json:"otherApprovalsTabInfo"`
	UpcomingCertificationsTabInfo TabInfo `json:"upcomingCertificationsTabInfo"`
	SearchCertificationsTabInfo   TabInfo `json:"searchCertificationsTabInfo"`
	TaskManagerTabInfo            TabInfo `json:"taskManagerTabInfo"`
}

type StudioModuleInfo struct {
	ModuleLicenseInfo
	IsOperator bool `json:"isOperator"`
	IsAdmin    bool `json:"isAdmin"`
}

type InsightsModuleInfo struct {
	ModuleLicenseInfo
	IsManager bool `json:"isManager"`
	IsViewer  bool `json:"isViewer"`
}

type ConfigurationModuleInfo struct {
	ModuleLicenseInfo
	AuditTabInfo TabInfo `json:"auditTabInfo"`
}

type UIInfo struct {
	LogoUrl             string `json:"logoUrl"`
	BackgroundGradient1 string `json:"backgroundGradient1"`
	BackgroundGradient2 string `json:"backgroundGradient2"`
	WideLogoURL         string `json:"wideLogoURL"`
	NarrowLogoURL       string `json:"narrowLogoURL"`
	FaviconURL          string `json:"faviconURL"`
	BrandColorOne       string `json:"brandColorOne"`
	BrandColorTwo       string `json:"brandColorTwo"`
}

type FeatureInfo struct {
	LoginConfigs     bool `json:"loginConfigs"`
	SsoPortal        bool `json:"ssoPortal"`
	ThirdPartyPortal bool `json:"thirdPartyPortal"`
	SafeId           bool `json:"safeId"`
	ShieldId         bool `json:"shieldId"`
	IdHub            bool `json:"idHub"`
	PasswordVault    bool `json:"passwordVault"`
	ProxyAs          bool `json:"proxyAs"`
}

type ShieldIdInfo struct {
	ApiDomain string `json:"apiDomain"`
	ClientId  string `json:"clientId"`
	HostId    string `json:"hostId"`
}

func (c *Client) GetBootstrapInfo() (*GetBootstrapInfoOutput, error) {
	var output GetBootstrapInfoOutput

	url := fmt.Sprintf("%s/bootstrapInfo", c.baseEndpoint)
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
		return nil, err
	}

	return &output, nil
}
