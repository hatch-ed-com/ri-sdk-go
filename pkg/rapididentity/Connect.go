package rapididentity

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
)

// BUG(Identity Automation): Downloading a compressed file is not possible with GetConnectFileContent. If compression is needed use GetConnectFileContentZip

// Input for retrieving Connect actions.
type GetConnectActionsInput struct {
	// The Connect project to filter by.
	// If empty, all projects will be searched.
	// For identifying the <Main> project use the
	// const variable MainProject. For bulitin
	// actions use $builtin
	Project string `json:"project" jsonschema:"The Connect project to filter by. If empty, all projects will be searched. For identifying the <Main> project use the const variable MainProject. For builtin actions use $builtin"`

	// Whether to return full action details
	// or just metadata.
	MetaDataOnly bool `json:"metaDataOnly" jsonschema:"Whether to return full action details or just metadata."`
}

// Output for retrieving Connect actions
type GetConnectActionsOutput struct {
	// Query type name. For example "all".
	Name string `json:"name" jsonschema:"Query type name. For example \"all\"."`

	// The list of actions.
	ActionDefs ActionDefList `json:"actionDefs" jsonschema:"The list of actions."`
}

// Input for retrieving file content from the
// Connect files module and logs
type GetConnectFileContentInput struct {
	// The path to the file to retrieve.
	// This can also be used to retrieve job and run logs
	// by setting the path to log/job or log/run.
	// This member is required
	Path string `json:"path" jsonschema:"The path to the file to retrieve. This can also be used to retrieve job and run logs by setting the path to log/job or log/run. This member is required"`

	// The connect project name that the directory or file resides
	// The default is the .Main project
	Project string `json:"project" jsonschema:"The connect project name that the directory or file resides The default is the .Main project"`

	// Determines whether to decompress file on the
	// RapidIdentity Server. At this time, there
	// is a bug in retrieving a compressed
	// file
	Decompress bool `json:"decompress" jsonschema:"Determines whether to decompress file on the RapidIdentity Server. At this time, there is a bug in retrieving a compressed file"`

	// The format of the result
	// The default is text/plain
	ResponseType string `json:"responseType" jsonschema:"The format of the result The default is text/plain"`
}

// Input for retrieving multiple files zipped from the
// Connect files module and logs
type GetConnectFileContentZipInput struct {
	// An array of paths to the files to retrieve.
	// This can also be used to retrieve job and run logs
	// by setting the path to log/job or log/run.
	PathList StringList `json:"pathList" jsonschema:"An array of paths to the files to retrieve. This can also be used to retrieve job and run logs by setting the path to log/job or log/run. This member is required"`

	// The connect project name that the directory or file resides
	// The default is the .Main project
	Project string `json:"project" jsonschema:"The connect project name that the directory or file resides The default is the .Main project"`
}

// Input for retrieving metadata on files from the
// Connect files module and logs.
type GetConnectFilesInput struct {
	// The path to the directory or file metadata to retrieve.
	// This can also be used to retrieve job and run logs
	// by setting the path to log/job or log/run.
	// This member is required
	Path string `json:"path" jsonschema:"The path to the directory or file metadata to retrieve. This can also be used to retrieve job and run logs by setting the path to log/job or log/run. This member is required"`

	// The connect project name that the directory or file resides
	// The default is the .Main project
	Project string `json:"project" jsonschema:"The connect project name that the directory or file resides The default is the .Main project"`

	// The format of the result
	// The default is application/json
	ResponseType string `json:"responseType" jsonschema:"The format of the result The default is application/json"`
}

// Represents the metadata of a file
type FileEntry struct {
	// The path to the directory or file
	Path string `json:"path" jsonschema:"The path to the directory or file"`

	// The size of the file in bytes
	Size int `json:"size" jsonschema:"The size of the file in bytes"`

	// The unix timestamp in milliseconds of when the
	// file or directory was modified
	Timestamp int64 `json:"timestamp" jsonschema:"The unix timestamp in milliseconds of when the file or directory was modified"`

	// The Connect project where the file or directory resides
	Project string `json:"project" jsonschema:"The Connect project where the file or directory resides"`

	// Whether or not the file or directory is readable
	Readable bool `json:"readable" jsonschema:"Whether or not the file or directory is readable"`

	// Whether or not the file or directory is writable
	Writable bool `json:"writable" jsonschema:"Whether or not the file or directory is writable"`
}

type FileEntryList []FileEntry

func (fel FileEntryList) MarshalJSON() ([]byte, error) {
	if fel == nil {
		return []byte("[]"), nil
	}
	return json.Marshal([]FileEntry(fel))
}

// Output for retrieving Connect files metadata
type GetConnectFilesOutput struct {
	FileEntry

	// If the path is a directory this will contain
	// the list of all files in the directory.
	// Only goes one level deep
	FileEntries FileEntryList `json:"fileEntries" jsonschema:"If the path is a directory this will contain the list of all files in the directory. Only goes one level deep"`
}

// Input for retrieving Connect jobs.
type GetConnectJobsInput struct {
	// The connect project name to retrieve jobs.
	// For identifying the <Main> project use the
	// const variable MainProject.
	Project string `json:"project" jsonschema:"The connect project name to retrieve jobs. For identifying the <Main> project use the const variable MainProject."`
}

type ConnectJobList []ConnectJob

func (cjl ConnectJobList) MarshalJSON() ([]byte, error) {
	if cjl == nil {
		return []byte("[]"), nil
	}
	return json.Marshal([]ConnectJob(cjl))
}

// Output for retrieving Connect jobs
type GetConnectJobsOutput struct {
	// List of Connect jobs.
	Jobs ConnectJobList `json:"jobs" jsonschema:"List of Connect jobs."`
}

type ConnectJob struct {
	// The name of the job.
	Name string `json:"name" jsonschema:"The name of the job."`

	// The unique ID of the job.
	Id string `json:"id" jsonschema:"The unique ID of the job."`

	// The version of the job.
	Version int `json:"version" jsonschema:"The version of the job."`

	// The description of the job.
	Description string `json:"description" jsonschema:"The description of the job."`

	// Whether trace logging is enabled of the job.
	TraceEnabled bool `json:"traceEnabled" jsonschema:"Whether trace logging is enabled of the job."`

	// The cron formatted job schedule.
	// The format is <second> <minute> <hour> <day of month> <month> <day of week>
	// An example for every 5 minutes would be 0 */5 * * * ?
	CronSpec string `json:"cronSpec" jsonschema:"The cron formatted job schedule. The format is <second> <minute> <hour> <day of month> <month> <day of week> An example for every 5 minutes would be 0 */5 * * * ?"`

	// The Timezone for the schedule.
	TimeZone string `json:"timeZone" jsonschema:"The Timezone for the schedule."`

	// Whether the job is disabled.
	Disabled bool `json:"disabled" jsonschema:"Whether the job is disabled."`

	// Whether the job will attached the generated log
	// when sending the job completion email.
	AttachLog bool `json:"attachLog" jsonschema:"Whether the job will attached the generated log when sending the job completion email."`

	// Whether to skip job execution if the previous
	// job is still running
	SkipOverlap bool `json:"skipOverlap" jsonschema:"Whether to skip job execution if the previous job is still running"`

	// A comma separated list of email addresses to send
	// job completion email.
	EmailRecipients string `json:"emailRecipients" jsonschema:"A comma separated list of email addresses to send job completion email."`

	// The project where the job resides.
	Project string `json:"project" jsonschema:"The project where the job resides."`

	// The number of days to keep logs for the job.
	LogRetentionDays int `json:"logRetentionDays" jsonschema:"The number of days to keep logs for the job."`

	// The Connect Action Set associated with the Job.
	Action ConnectAction `json:"action" jsonschema:"The Connect Action Set associated with the Job."`

	// How long the job should run before timing out.
	TimeoutSeconds int `json:"timeoutSeconds" jsonschema:"How long the job should run before timing out."`

	// Whether the job can be run externally.
	RunExternal bool `json:"runExternal" jsonschema:"Whether the job can be run externally."`
}

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

// Input for searching action sets within
// a Connect Project
type SearchConnectActionSetsInput struct {
	// The text to search for. This can
	// also be a regex pattern if Regex is
	// set to true.
	SearchString string `json:"searchString" jsonschema:"The text to search for. This can also be a regex pattern if Regex is set to true."`

	// The Connect project to search within.
	// If empty, all projects will be searched.
	// For identifying the <Main> project use the
	// const variable MainProject.
	Project string `json:"project" jsonschema:"The Connect project to search within. If empty, all projects will be searched. For identifying the <Main> project use the const variable MainProject."`

	// Whether to match action set names.
	MatchAction bool `json:"matchAction" jsonschema:"Whether to match action set names."`

	// Whether to apply a case sensitive search.
	MatchCase bool `json:"matchCase" jsonschema:"Whether to apply a case sensitive search."`

	// Whether the search string contains regex.
	Regex bool `json:"regex" jsonschema:"Whether the search string contains regex."`
}

// Result from searching for action sets within a
// Connect project.
type SearchConnectActionSetsOutput struct {
	// The name of the search query.
	Name string `json:"name" jsonschema:"The name of the search query."`

	// The description of the search query.
	Description string `json:"description" jsonschema:"The description of the search query."`

	// The action set definition results.
	ActionDefs ActionDefList `json:"actionDefs" jsonschema:"The action set definition results."`

	// The http status code returned.
	HttpStatus int `json:"httpStatus" jsonschema:"The http status code returned."`
}

type ActionDef struct {
	// The action set ID.
	Id string `json:"id" jsonschema:"The action set ID. On creation of an action this must be populated with a UUID of the format xxxxxxxx-xxxx-xxxx-xxxx-xxxxxxxxxxxx"`

	// The action set version.
	Version int `json:"version" jsonschema:"The action set version."`

	// The project where the action set resides.
	Project string `json:"project" jsonschema:"The project where the action set resides."`

	// The name of the action set.
	Name string `json:"name" jsonschema:"The name of the action set."`

	// The category the action set is a part of.
	Category string `json:"category" jsonschema:"The category the action set is a part of."`

	// Whether the action set is built in or custom.
	BuiltIn bool `json:"builtIn" jsonschema:"Whether the action set is built in or custom."`

	// Whether the action set is a part of the community depot.
	Community bool `json:"community" jsonschema:"Whether the action set is a part of the community depot."`

	// Whether the action set returns a value.
	ReturnsValue bool `json:"returnsValue" jsonschema:"Whether the action set returns a value."`

	// The description of the action set.
	Description string `json:"description" jsonschema:"The description of the action set."`

	// Whether the action set is licensed or not.
	Unlicensed bool `json:"unlicensed" jsonschema:"Whether the action set is licensed or not."`

	// Whether the action set contains sensitive information.
	Sensitive bool `json:"sensitive" jsonschema:"Whether the action set contains sensitive information."`

	// The input parameters of the action set.
	ArgDefs ArgDefList `json:"argDefs" jsonschema:"The input parameters of the action set."`

	// The actions within the action set.
	Actions ConnectActionList `json:"actions" jsonschema:"The actions within the action set."`

	// Whether the action set is deprecated.
	Deprecated string `json:"deprecated" jsonschema:"Whether the action set is deprecated."`

	// The http status code of the return.
	HttpStatus int `json:"httpStatus" jsonschema:"The http status code of the return."`

	// The number of times the action set has been modified.
	ChangeCount int `json:"changeCount" jsonschema:"The number of times the action set has been modified."`

	// When the action set was last modified.
	ModifiedMs int64 `json:"modifiedMs" jsonschema:"When the action set was last modified."`

	// The idautoID of the user who modified the action set.
	ModifiedBy string `json:"modifiedBy" jsonschema:"The idautoID of the user who modified the action set."`

	// The display name of the user who modified the action set.
	ModifiedByName string `json:"modifiedByName" jsonschema:"The display name of the user who modified the action set."`
}

type ActionDefList []ActionDef

func (adl ActionDefList) MarshalJSON() ([]byte, error) {
	if adl == nil {
		return []byte("[]"), nil
	}
	return json.Marshal([]ActionDef(adl))
}

type ArgDef struct {
	// Whether the action set input parameter is optional.
	Optional bool `json:"optional" jsonschema:"Whether the action set input parameter is optional."`

	// The type of the input parameter
	Type string `json:"type" jsonschema:"The type of the input parameter"`

	// The name of the input parameter.
	Name string `json:"name" jsonschema:"The name of the input parameter."`

	// The description for the input parameter.
	Description string `json:"description" jsonschema:"The description for the input parameter."`

	// The value of the input parameter.
	Value string `json:"value" jsonschema:"The value of the input parameter"`
}

type ArgDefList []ArgDef

func (arg ArgDefList) MarshalJSON() ([]byte, error) {
	if arg == nil {
		return []byte("[]"), nil
	}
	return json.Marshal([]ArgDef(arg))
}

type ConnectAction struct {
	// The unique ID of the Connect action.
	Id string `json:"id" jsonschema:"The unique ID of the Connect action."`

	// The name of the Connect action.
	Name string `json:"name" jsonschema:"The name of the Connect action."`

	// Whether the action returns a value.
	OutputVar string `json:"outputVar" jsonschema:"Whether the action returns a value."`

	// Whether the action is disabled.
	Disabled bool `json:"disabled" jsonschema:"Whether the action is disabled."`

	// The project where the action resides.
	Project string `json:"project" jsonschema:"The project where the action resides."`

	// The input parameters for the action.
	Args ArgDefList `json:"args" jsonschema:"The input parameters for the action."`
}

type ConnectActionList []ConnectAction

func (cal ConnectActionList) MarshalJSON() ([]byte, error) {
	if cal == nil {
		return []byte("[]"), nil
	}
	return json.Marshal([]ConnectAction(cal))
}

type GetConnectActionByIdInput struct {
	// The Connect action ID, or name. The
	// name is in the format <project>.<name>
	Id string `json:"id" jsonschema:"The unique Connect action ID or name."`

	// Whether to return full action details
	// or just metadata.
	MetaDataOnly bool `json:"metaDataOnly" jsonschema:"Whether to return full action details or just metadata."`
}

type GetConnectActionByIdOutput struct {
	Action ActionDef `json:"action" jsonschema:"The action that was retrieved based on the name of id. The version that is returned is the value of the version when updating the action. If there is a conflict, query the action and ensure the version number is correct"`
}

type SaveConnectActionInput struct {
	Action ActionDef `json:"action" jsonschema:"The action to save or update. When updating an existing action set the version number must be the one provided to you in a Connect action query."`
}

type SaveConnectActionOutput struct {
	Action ActionDef `json:"action" jsonschema:"The action that was created or updated. The version that is returned is the value of the version when updating the action. If there is a conflict, query the action and ensure the version number is correct"`
}

// Retrieves actions from Connect.
//
//meta:operation GET /admin/connect/actions
func (c *Client) GetConnectActions(ctx context.Context, params GetConnectActionsInput) (*GetConnectActionsOutput, error) {
	var output GetConnectActionsOutput

	url := fmt.Sprintf("%s/admin/connect/actions?metaDataOnly=%t", c.baseEndpoint, params.MetaDataOnly)
	if params.Project != "" {
		if params.Project == MainProject {
			url = fmt.Sprintf("%s&project=", url)
		} else {
			url = fmt.Sprintf("%s&project=%s", url, params.Project)
		}
	}
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

// Retrieves a Connect action by name or ID.
//
//meta:operation GET /admin/connect/actions/{nameOrId}
func (c *Client) GetConnectActionById(ctx context.Context, params GetConnectActionByIdInput) (*GetConnectActionByIdOutput, error) {
	var output ActionDef

	url := fmt.Sprintf("%s/admin/connect/actions/%s?metaDataOnly=%t", c.baseEndpoint, params.Id, params.MetaDataOnly)
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

	return &GetConnectActionByIdOutput{
		Action: output,
	}, nil
}

// Retrieves file content from a file within the Connect files
// module and logs.
//
//meta:operation GET /admin/connect/fileContent/{path}
func (c *Client) GetConnectFileContent(ctx context.Context, params GetConnectFileContentInput) ([]byte, error) {
	url := fmt.Sprintf("%s/admin/connect/fileContent/%s?project=%s&decompress=%t", c.baseEndpoint, params.Path, params.Project, params.Decompress)
	req, err := c.GenerateRequest(ctx, "GET", url, nil)
	if err != nil {
		return nil, err
	}

	if params.ResponseType == "" {
		req.Header.Set("Accept", "text/plain")
	} else {
		req.Header.Set("Accept", params.ResponseType)
	}

	res, err := c.httpClient.Do(req)
	if err != nil {
		return nil, err
	}
	resBody, err := c.ReceiveResponse(res)
	if err != nil {
		return nil, err
	}

	return resBody, nil
}

// Retrieves multiple files zipped from the Connect files
// module and logs.
//
//meta:operation GET /admin/connect/fileContentZip
func (c *Client) GetConnectFileContentZip(ctx context.Context, params GetConnectFileContentZipInput) ([]byte, error) {
	url := fmt.Sprintf("%s/admin/connect/fileContentZip?project=%s", c.baseEndpoint, params.Project)
	for _, path := range params.PathList {
		url = fmt.Sprintf("%s&path=%s", url, path)
	}
	req, err := c.GenerateRequest(ctx, "GET", url, nil)
	if err != nil {
		return nil, err
	}

	req.Header.Set("Accept", "application/zip")

	res, err := c.httpClient.Do(req)
	if err != nil {
		return nil, err
	}
	resBody, err := c.ReceiveResponse(res)
	if err != nil {
		return nil, err
	}

	return resBody, nil
}

// Retrieves metadata for files within the Connect files
// module and logs. This does NOT retrieve the file contents
// only the metadata as shown in the GetConnectFilesOutput
//
//meta:operation GET /admin/connect/files/{path}
func (c *Client) GetConnectFiles(ctx context.Context, params GetConnectFilesInput) (*GetConnectFilesOutput, error) {
	var output GetConnectFilesOutput

	url := fmt.Sprintf("%s/admin/connect/files/%s?project=%s", c.baseEndpoint, params.Path, params.Project)
	req, err := c.GenerateRequest(ctx, "GET", url, nil)
	if err != nil {
		return nil, err
	}

	if params.ResponseType != "" {
		req.Header.Set("Accept", params.ResponseType)
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

// Retrieves Connect Jobs for all projects or specified project.
//
//meta:operation GET /admin/connect/jobs
func (c *Client) GetConnectJobs(ctx context.Context, params GetConnectJobsInput) (*GetConnectJobsOutput, error) {
	var output GetConnectJobsOutput

	url := fmt.Sprintf("%s/admin/connect/jobs", c.baseEndpoint)
	if params.Project != "" {
		if params.Project == MainProject {
			url = fmt.Sprintf("%s?project=", url)
		} else {
			url = fmt.Sprintf("%s?project=%s", url, params.Project)
		}
	}

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

// Searches for text within action sets in a project.
//
//meta:operation GET /admin/connect/search/actions
func (c *Client) SearchConnectActionSets(ctx context.Context, params SearchConnectActionSetsInput) (*SearchConnectActionSetsOutput, error) {
	var output SearchConnectActionSetsOutput

	url := fmt.Sprintf("%s/admin/connect/search/actions?searchString=%s&matchAction=%t&matchCase=%t&regex=%t", c.baseEndpoint, params.SearchString, params.MatchAction, params.MatchCase, params.Regex)
	if params.Project != "" {
		if params.Project == MainProject {
			url = fmt.Sprintf("%s&project=", url)
		} else {
			url = fmt.Sprintf("%s&project=%s", url, params.Project)
		}
	}
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

// Create or update a Connect Action Set
//
//meta:operation POST /admin/connect/actions
func (c *Client) SaveConnectAction(ctx context.Context, params SaveConnectActionInput) (*SaveConnectActionOutput, error) {
	url := fmt.Sprintf("%s/admin/connect/actions", c.baseEndpoint)
	action, err := json.Marshal(params.Action)
	if err != nil {
		return nil, err
	}
	requestBody := bytes.NewBuffer(action)
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
	var output ActionDef
	err = json.Unmarshal(resBody, &output)
	if err != nil {
		return nil, err
	}

	return &SaveConnectActionOutput{
		Action: output,
	}, nil
}
