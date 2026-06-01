package rapididentity

import (
	"context"
	"encoding/json"
	"fmt"
	"time"
)

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
