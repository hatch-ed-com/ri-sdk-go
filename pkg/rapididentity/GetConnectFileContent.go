package rapididentity

import "fmt"

// BUG(Identity Automation): Downloading a compressed file is not possible with GetConnectFileContent. If compression is needed use GetConnectFileContentZip

// Input for retrieving file content from the
// Connect files module and logs
type GetConnectFileContentInput struct {
	// The path to the file to retrieve.
	// This can also be used to retrieve job and run logs
	// by setting the path to log/job or log/run.
	// This member is required
	Path string

	// The connect project name that the directory or file resides
	// The default is the .Main project
	Project string

	// Determines whether to decompress file on the
	// RapidIdentity Server. At this time, there
	// is a bug in retrieving a compressed
	// file
	Decompress bool

	// The format of the result
	// The default is text/plain
	ResponseType string
}

// Retrieves file content from a file within the Connect files
// module and logs.
//
//meta:operation GET /admin/connect/fileContent/{path}
func (c *Client) GetConnectFileContent(params GetConnectFileContentInput) ([]byte, error) {
	url := fmt.Sprintf("%s/admin/connect/fileContent/%s?project=%s&decompress=%t", c.baseEndpoint, params.Path, params.Project, params.Decompress)
	req, err := c.GenerateRequest("GET", url, nil)
	if err != nil {
		return nil, err
	}

	if params.ResponseType == "" {
		req.Header.Set("Accept", params.ResponseType)
	} else {
		req.Header.Set("Accept", "text/plain")
	}

	res, err := c.options.HTTPClient.Do(req)
	if err != nil {
		return nil, err
	}
	resBody, err := c.ReceiveResponse(res)
	if err != nil {
		return nil, err
	}

	return resBody, nil
}
