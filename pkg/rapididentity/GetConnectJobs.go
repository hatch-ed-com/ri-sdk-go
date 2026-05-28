package rapididentity

import (
	"context"
	"encoding/json"
	"fmt"
)

// Input for retrieving Connect jobs.
type GetConnectJobsInput struct {
	// The connect project name to retrieve jobs.
	// For identifying the <Main> project use the
	// const variable MainProject.
	Project string `json:"project" jsonschema:"The connect project name to retrieve jobs. For identifying the <Main> project use the const variable MainProject."`
}

// Output for retrieving Connect jobs
type GetConnectJobsOutput struct {
	// List of Connect jobs.
	Jobs []ConnectJob `json:"jobs" jsonschema:"List of Connect jobs."`
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
