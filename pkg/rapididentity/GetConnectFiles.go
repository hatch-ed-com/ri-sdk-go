package rapididentity

import (
	"context"
	"encoding/json"
	"fmt"
)

// Input for retrieving metadata on files from the
// Connect files module and logs.
type GetConnectFilesInput struct {
	// The path to the directory or file metadata to retrieve.
	// This can also be used to retrieve job and run logs
	// by setting the path to log/job or log/run.
	// This member is required
	Path string

	// The connect project name that the directory or file resides
	// The default is the .Main project
	Project string

	// The format of the result
	// The default is application/json
	ResponseType string
}

// Represents the metadata of a file
type FileEntry struct {
	// The path to the directory or file
	Path string `json:"path"`

	// The size of the file in bytes
	Size int `json:"size"`

	// The unix timestamp in milliseconds of when the
	// file or directory was modified
	Timestamp int64 `json:"timestamp"`

	// The Connect project where the file or directory resides
	Project string `json:"project"`

	// Whether or not the file or directory is readable
	Readable bool `json:"readable"`

	// Whether or not the file or directory is writable
	Writable bool `json:"writable"`
}

// Output for retrieving Connect files metadata
type GetConnectFilesOutput struct {
	FileEntry

	// If the path is a directory this will contain
	// the list of all files in the directory.
	// Only goes one level deep
	FileEntries []FileEntry `json:"fileEntries"`
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
