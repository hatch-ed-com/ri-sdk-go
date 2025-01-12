package rapididentity

import (
	"encoding/json"
	"fmt"
	"time"
)

type GetBootstrapInfoOutput struct {
	// RapidIdentity license information.
	LicenseInfo LicenseInfo `json:"licenseInfo"`

	// RapidIdentity version information.
	VersionInfo VersionInfo `json:"versionInfo"`

	// RapidIdentity session information from the invoking user.
	SessionInfo SessionInfo `json:"sessionInfo"`

	// RapidIdentity module information.
	ModuleInfo ModuleInfo `json:"moduleInfo"`

	// RapidIdentity Portal UI appearance information.
	UIInfo UIInfo `json:"uiInfo"`

	// The default RapidIdentity module page that a user lands after login.
	DefaultLandingModule string `json:"defaultLandingModule"`

	// Whether the RapidIdentity tenant is restarted after an upgrade is applied.
	DisableUpgradesRestarts bool `json:"disableUpgradesRestarts"`

	// Whether the RapidIdentity ProxyAs feature can be enabled for non admins.
	AllowProxy bool `json:"allowProxy"`

	// Whether the tenant is a RapidIdentity Cloud tenant.
	Idaas bool `json:"idaas"`

	// The RapidIdentity Tenant ID.
	TenantId string `json:"tenantId"`

	// Whether notifications are enabled for the tenant.
	NotificationsEnabled bool `json:"notificationsEnabled"`

	// Whether Global Search is enabled.
	GlobalSearchEnabled bool `json:"globalSearchEnabled"`

	// Whether the LDAP directory is within the RapidIdentity cloud tenant and not external.
	IsRICloudLdap bool `json:"isRICloudLdap"`

	// Information on the RapidIdentity features applied.
	Features FeatureInfo `json:"features"`

	// The url for retrieving studio library apps.
	DepotProxy DepotProxyInfo `json:"depotProxy"`

	// Whether the RapidIdentity tenant has personas enabled
	HasPersonas bool `json:"hasPersonas"`

	// ShieldID information for the RapidIdentity tenant
	ShieldIdInfo ShieldIdInfo `json:"shieldIdInfo"`

	// Whether the tenant is an ID Hub enabled tenant.
	IdHub bool `json:"idHub"`

	// Whether the new Google Social login is enabled.
	IsIdAutoGoogleEnabled bool `json:"isIdAutoGoogleEnabled"`

	// Whether the new Apple Social login is enabled.
	IsIdAutoAppleEnabled bool `json:"isIdAutoAppleEnabled"`
}

type LicenseInfo struct {
	// The license type such as subscription.
	Type string `json:"type"`

	// The display name for the person or entity
	// that the license is owned by.
	Licensee string `json:"licensee"`

	// The unique ID for the license.
	LicenseeId string `json:"licenseeId"`

	// The date the license expires.
	ExpirationDate string `json:"expirationDate"`

	// The cluster count.
	ClusterCount int `json:"clusterCount"`

	// The number of licensed users.
	LicensedUserCount int `json:"licensedUserCount"`

	// The modules provided by the license (Legacy)
	Modules []string `json:"modules"`
}

type VersionInfo struct {
	// The RapidIdentity version number.
	Version string `json:"version"`

	// The timestamp for the RapidIdentity version build.
	BuildTimestamp time.Time `json:"buildTimestamp"`
}

type ModuleInfo struct {
	// RapidIdentity Applications module information.
	Applications ApplicationModuleInfo `json:"applications"`

	// RapidIdentity Dashboard module information.
	Dashboard DashboardModuleInfo `json:"dashboard"`

	// RapidIdentity SSO Portal (Personas) information.
	DashboardV2 DashboardV2ModuleInfo `json:"dashboard_V2"`

	// RapidIdentity Files module information.
	Files FileModuleInfo `json:"files"`

	// RapidIdentity People module information.
	Profiles ProfileModuleInfo `json:"profiles"`

	// RapidIdentity Reports module information.
	Reporting ReportingModuleInfo `json:"reporting"`

	// RapidIdentity Groups module information.
	Roles RolesModuleInfo `json:"roles"`

	// RapidIdentity Sponsorship module information.
	Sponsorship SponsorshipModuleInfo `json:"sponsorship"`

	// RapidIdentity Requests module information.
	Workflow WorkflowModuleInfo `json:"workflow"`

	// RapidIdentiy Admin information.
	Admin AdminModuleInfo `json:"admin"`

	// RapidIdentity Connect module information.
	Connect ConnectModuleInfo `json:"connect"`

	// RapidIdentity Studio module information.
	Studio StudioModuleInfo `json:"studio"`

	// RapidIdentity folders module information.
	Folders FoldersModuleInfo `json:"folders"`

	// RapidIdentity Insights module information.
	Insights InsightsModuleInfo `json:"insights"`

	// RapidIdentity configuration module information.
	Configuration ConfigurationModuleInfo `json:"configuration"`

	// RapidIdentity IdHub information.
	IdHub IdHubModuleInfo `json:"idHub"`
}

type ApplicationModuleInfo struct {
	ModuleLicenseInfo

	// My tab visibility information.
	MyTabInfo TabInfo `json:"myTabInfo"`

	// Team tab visibility information.
	TeamTabInfo TabInfo `json:"teamTabInfo"`

	// Other tab visibility information.
	OtherTabInfo TabInfo `json:"otherTabInfo"`

	// Preferences information.
	Preferences PreferenceInfo `json:"preferences"`
}

type ModuleLicenseInfo struct {
	// Whether RapidIdentity module is licensed.
	Licensed bool `json:"licensed"`

	// Whether RapidIdentity module is visible for the invoking user.
	Visible bool `json:"visible"`
}

type TabInfo struct {
	// Whether tab is visible for the invoking user.
	Visible bool `json:"visible"`
}

type TabActionInfo struct {
	TabInfo
	// Available actions within the specified module tab.
	Actions []string `json:"actions"`
}

type TabAdminInfo struct {
	// Whether config is visible for the invoking user.
	ConfigTabVisible bool `json:"configTabVisible"`

	// Whether admin tab is visible for the invoking user.
	AdminTabVisible bool `json:"adminTabVisible"`
}

type PreferenceInfo struct {
	// Whether to start on favorites in the applications module.
	StartAtFavorites bool `json:"startAtFavorites"`
}

type DashboardModuleInfo struct {
	ModuleLicenseInfo
	// Visibility information for the My Activity Tab
	MyActivityTab TabInfo `json:"myActivityTab"`

	// Visibility information for the Team Activity Tab
	TeamActivityTab TabInfo `json:"teamActivityTab"`

	// Visibility information for the Other Activity Tab
	OtherActivityTab TabInfo `json:"otherActivityTab"`

	// Visibility information for the My Entitlements Tab
	MyEntitlementsTab TabInfo `json:"myEntitlementsTab"`

	// Visibility information for the Team Entitlements Tab
	TeamEntitlementsTab TabInfo `json:"teamEntitlementsTab"`

	// Visibility information for the Other Entitlements Tab
	OtherEntitlementsTab TabInfo `json:"otherEntitlementsTab"`

	// Visibility information for the Executive Tab
	ExecutiveTab TabInfo `json:"executiveTab"`
}

type DashboardV2ModuleInfo struct {
	ModuleLicenseInfo
}

type FileModuleInfo struct {
	ModuleLicenseInfo
	// The maximum file size, in MB, that can be uploaded.
	MaxUploadSize float32 `json:"maxUploadSize"`

	// Whether SSL upload is enabled.
	EnableSSLUpload bool `json:"enableSSLUpload"`

	// Whether public access is enabled.
	EnableMakePublic bool `json:"enableMakePublic"`
}

type ProfileModuleInfo struct {
	ModuleLicenseInfo

	// Whether challenge questions are enabled (Legacy).
	ChallengeQuestionsEnabled bool `json:"challengeQuestionsEnabled"`

	// Whether the invoking user must update their challenge questions.
	MustUpdateChallengeQuestions bool `json:"mustUpdateChallengeQuestions"`

	// The invalid challenge set message (Legacy).
	InvalidChallengeSetMessage string `json:"invalidChallengeSetMessage"`

	// Whether show all is enabled.
	EnableShowAll bool `json:"enableShowAll"`

	// Whether must change password options are supported.
	MustChangePasswordOptionsSupported bool `json:"mustChangePasswordOptionsSupported"`

	// Whether an unlock requires a password reset.
	UnlockRequiresPasswordReset bool `json:"unlockRequiresPasswordReset"`

	// Whether a delegation attribute is required.
	DelegationAttrRequired bool `json:"delegationAttrRequired"`
}

type ReportingModuleInfo struct {
	ModuleLicenseInfo

	// The max results returned from a Reports module audit report query.
	AuditReportMax int `json:"auditReportMax"`
}

type RolesModuleInfo struct {
	ModuleLicenseInfo
	// Whether to enable selection of group types.
	EnableGroupTypeSelection bool `json:"enableGroupTypeSelection"`

	// Available group types to choose form if group type selection is enabled.
	PossibleGroupTypes []string `json:"possibleGroupTypes"`

	// Allowed group types to choose from if group type selection is enabled.
	AllowedGroupTypes []string `json:"allowedGroupTypes"`

	// Whether to preload all groups in Groups module.
	PreloadGroups bool `json:"preloadGroups"`

	// Whether to show the distinguished name of the RapidIdentity group.
	ShowDN bool `json:"showDN"`

	// Visibility and actions for the My Groups tab.
	MyTabInfo TabActionInfo `json:"myTabInfo"`

	// Visibility and actions for the Team Groups tab.
	TeamTabInfo TabActionInfo `json:"teamTabInfo"`

	// Visibility and actions for the Other Groups tab.
	OtherTabInfo TabActionInfo `json:"otherTabInfo"`

	// Any custom attributes added to the groups module.
	CustomAttributes []DelegationAttribute `json:"customAttributes"`

	// Whether wildcard (*) searches are allowed.
	EnableWildcardSearch bool `json:"enableWildcardSearch"`
}

type SponsorshipModuleInfo struct {
	ModuleLicenseInfo

	// Visibility and actions for the My Sponsored Accounts tab.
	MyTabInfo TabActionInfo `json:"myTabInfo"`

	// Visibility and actions for the Team Sponsored Accounts tab.
	TeamTabInfo TabActionInfo `json:"teamTabInfo"`

	// Visibility and actions for the Other Sponsored Accounts tab.
	OtherTabInfo TabActionInfo `json:"otherTabInfo"`

	// The max expiration date you can select when creating or certifying a sponsored account.
	MaxExpirationDays int `json:"maxExpirationDays"`

	// Whether an email address is required for a sponsored account.
	EmailAddressRequired bool `json:"emailAddressRequired"`

	// Whether an expiration date is required for a sponsored account.
	ExpirationRequired bool `json:"expirationRequired"`

	// Whether to preload all sponsors.
	PreloadSponsors bool `json:"preloadSponsors"`

	// Whether to load all sponsored accounts upon accessing the Sponsorship module.
	PreloadSponsoredAccounts bool `json:"preloadSponsoredAccounts"`

	// Custom attributes added to the sponsorship module.
	CustomAttributes []DelegationAttribute `json:"customAttributes"`
}

type WorkflowModuleInfo struct {
	ModuleLicenseInfo

	// Whether SSL uploads are required.
	EnableSSLUpload bool `json:"enableSSLUpload"`

	// Visibility for My Dashboard tab.
	MyDashboardTabInfo TabInfo `json:"myDashboardTabInfo"`

	// Visibility for Team Dashboard tab.
	TeamDashboardTabInfo TabInfo `json:"teamDashboardTabInfo"`

	// Visibility for Other Dashboard tab.
	OtherDashboardTabInfo TabInfo `json:"otherDashboardTabInfo"`

	// Visibility for My Entitlements tab.
	MyRequestsTabInfo TabInfo `json:"myRequestsTabInfo"`

	// Visibility for Team Entitlements tab.
	TeamRequestsTabInfo TabInfo `json:"teamRequestsTabInfo"`

	// Visibility for Other Entitlements tab.
	OtherRequestsTabInfo TabInfo `json:"otherRequestsTabInfo"`

	// Visibility information for My Task Approvals.
	MyApprovalsTabInfo TabInfo `json:"myApprovalsTabInfo"`

	// Visibility information for Team Task Approvals.
	TeamApprovalsTabInfo TabInfo `json:"teamApprovalsTabInfo"`

	// Visibility information for Other Task Approvals.
	OtherApprovalsTabInfo TabInfo `json:"otherApprovalsTabInfo"`

	// Visibility information for Certification tasks.
	UpcomingCertificationsTabInfo TabInfo `json:"upcomingCertificationsTabInfo"`

	// Visibility information for searching Certification tasks.
	SearchCertificationsTabInfo TabInfo `json:"searchCertificationsTabInfo"`

	// Visibility information for task manager tab..
	TaskManagerTabInfo TabInfo `json:"taskManagerTabInfo"`
}

type AdminModuleInfo struct {
	// Whether configuration is visible for invoking user.
	Visible bool

	// Portal admin tab info.
	Portal TabAdminInfo `json:"portal"`

	// Application admin tab info.
	Applications TabAdminInfo `json:"applications"`

	// Dashboard admin tab info.
	Dashboard TabAdminInfo `json:"dashboard"`

	// Files admin tab info.
	Files TabAdminInfo `json:"files"`

	// People admin tab info.
	Profiles TabAdminInfo `json:"profiles"`

	// Reports admin tab info.
	Reporting TabAdminInfo `json:"reporting"`

	// Groups admin tab info.
	Roles TabAdminInfo `json:"roles"`

	// Sponsorship admin tab info.
	Sponsorship TabAdminInfo `json:"sponsorship"`

	// Workflow admin tab info.
	Workflow TabAdminInfo `json:"workflow"`
}

type ConnectModuleInfo struct {
	// Whether the Connect module is visible for the invoking user.
	Visible bool `json:"visible"`
}

type StudioModuleInfo struct {
	ModuleLicenseInfo
	// Whether the invoking user is a Studio operator.
	IsOperator bool `json:"isOperator"`

	// Whether the invoking user is a Studio administrator.
	IsAdmin bool `json:"isAdmin"`
}

type FoldersModuleInfo struct {
	// Whether folder users are visible for the invoking user.
	UsersVisible bool `json:"usersVisible"`

	// Whether folder groups are visible for the invoking user.
	GroupsVisible bool `json:"groupsVisible"`

	// Whether the invoking user is a Folders operator.
	IsOperator bool `json:"isOperator"`

	// Whether the invoking user is a Folders administrator.
	IsAdmin bool `json:"isAdmin"`

	// Whether the folers shcema is up to date.
	SchemaUpToDate bool `json:"schemaUpToDate"`
}

type InsightsModuleInfo struct {
	ModuleLicenseInfo

	// Whether the invoking user is an Insights manager.
	IsManager bool `json:"isManager"`

	// Whether the invoking user is an Insights viewer.
	IsViewer bool `json:"isViewer"`
}

type ConfigurationModuleInfo struct {
	ModuleLicenseInfo

	// Visibility information for the audit tab within configuration.
	AuditTabInfo TabInfo `json:"auditTabInfo"`
}

type IdHubModuleInfo struct {
	// Whether ID Hub is visible for the invoking user.
	Visible bool `json:"visible"`

	// LCS External Auth Client Id.
	LcsExternalAuthClientId string `json:"lcsExternalAuthClientId"`

	// The LCS Domain.
	LcsDomain string `json:"lcsDomain"`

	// The catalog domain.
	CatalogDomain string `json:"catalogDomain"`
}

type UIInfo struct {
	// The portal logo url.
	LogoUrl string `json:"logoUrl"`

	// The portal background gradient color 1.
	BackgroundGradient1 string `json:"backgroundGradient1"`

	// The portal background gradient color 2.
	BackgroundGradient2 string `json:"backgroundGradient2"`

	// The portal wide logo url.
	WideLogoURL string `json:"wideLogoURL"`

	// The portal narrow logo url.
	NarrowLogoURL string `json:"narrowLogoURL"`

	// The portal favicon url.
	FaviconURL string `json:"faviconURL"`

	// The portal brand color 1.
	BrandColorOne string `json:"brandColorOne"`

	// The portal brand color 2.
	BrandColorTwo string `json:"brandColorTwo"`
}

type FeatureInfo struct {
	// Whether login configs is enabled.
	LoginConfigs bool `json:"loginConfigs"`

	// Pendo information for RapidIdentity tenant.
	Pendo PendoInfo `json:"pendo"`

	// ChurnZero information for RapidIdentity tenant.
	ChurnZero ChurnZeroInfo `json:"churnZero"`

	// Whether personas is enabled.
	SsoPortal bool `json:"ssoPortal"`

	// Whether universal authentication director is available.
	ThirdPartyPortal bool `json:"thirdPartyPortal"`

	// Whether SafeID is available.
	SafeId bool `json:"safeId"`

	// Whether ShieldID is available.
	ShieldId bool `json:"shieldId"`

	// Whether ID Hub is available.
	IdHub bool `json:"idHub"`

	// Whether Password Vault is available.
	PasswordVault bool `json:"passwordVault"`

	// Whether ProxyAs is available for non admins.
	ProxyAs bool `json:"proxyAs"`
}

type PendoInfo struct {
	// Unique Pendo ID for the invoking user.
	Id string `json:"id"`

	// UserType of the invoking user.
	UserType string `json:"userType"`
}

type ChurnZeroInfo struct {
	// Unique Churn Zero ID for the invoking user.
	Id string `json:"id"`

	// User type of the invoking user.
	UserType string `json:"userType"`

	// User role of the invoking user.
	UserRole string `json:"userRole"`
}

type DepotProxyInfo struct {
	// The Studio library url.
	VettedStudioAppsUrl string `json:"vettedStudioAppsUrl"`
}

type ShieldIdInfo struct {
	// API Domain
	ApiDomain string `json:"apiDomain"`

	// CLient ID
	ClientId string `json:"clientId"`

	// Host ID
	HostId string `json:"hostId"`
}

// Retrieves RapidIdentity tenant and user access information for the invoking user.
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
