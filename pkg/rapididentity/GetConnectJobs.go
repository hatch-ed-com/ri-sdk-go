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
	Project string
}

// Output for retrieving Connect jobs
type GetConnectJobsOutput struct {
	// List of Connect jobs.
	Jobs []ConnectJob `json:"jobs"`
}

type ConnectJob struct {
	// The name of the job.
	Name string `json:"name"`

	// The unique ID of the job.
	Id string `json:"id"`

	// The version of the job.
	Version int `json:"version"`

	// The description of the job.
	Description string `json:"description"`

	// Whether trace logging is enabled of the job.
	TraceEnabled bool `json:"traceEnabled"`

	// The cron formatted job schedule.
	// The format is <second> <minute> <hour> <day of month> <month> <day of week>
	// An example for every 5 minutes would be 0 */5 * * * ?
	CronSpec string `json:"cronSpec"`

	// The Timezone for the schedule.
	TimeZone string `json:"timeZone"`

	// Whether the job is disabled.
	Disabled bool `json:"disabled"`

	// Whether the job will attached the generated log
	// when sending the job completion email.
	AttachLog bool `json:"attachLog"`

	// Whether to skip job execution if the previous
	// job is still running
	SkipOverlap bool `json:"skipOverlap"`

	// A comma separated list of email addresses to send
	// job completion email.
	EmailRecipients string `json:"emailRecipients"`

	// The project where the job resides.
	Project string `json:"project"`

	// The number of days to keep logs for the job.
	LogRetentionDays int `json:"logRetentionDays"`

	// The Connect Action Set associated with the Job.
	Action ConnectAction `json:"action"`

	// How long the job should run before timing out.
	TimeoutSeconds int `json:"timeoutSeconds"`

	// Whether the job can be run externally.
	RunExternal bool `json:"runExternal"`
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
