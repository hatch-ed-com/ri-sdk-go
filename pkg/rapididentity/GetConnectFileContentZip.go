package rapididentity

import "fmt"

// Input for retrieving multiple files zipped from the
// Connect files module and logs
type GetConnectFileContentZip struct {
	// An array of paths to the files to retrieve.
	// This can also be used to retrieve job and run logs
	// by setting the path to log/job or log/run.
	// This member is required
	PathList []string

	// The connect project name that the directory or file resides
	// The default is the .Main project
	Project string
}

// Retrieves multiple files zipped from the Connect files
// module and logs.
func (c *Client) GetConnectFileContentZip(params GetConnectFileContentZip) ([]byte, error) {
	url := fmt.Sprintf("%s/admin/connect/fileContentZip?project=%s", c.baseEndpoint, params.Project)
	for _, path := range params.PathList {
		url = fmt.Sprintf("%s&path=%s", url, path)
	}
	req, err := c.GenerateRequest("GET", url, nil)
	if err != nil {
		return nil, err
	}

	req.Header.Set("Accept", "application/zip")

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
