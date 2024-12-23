package main

import (
	"fmt"
	"net/http"

	"github.com/hatch-ed-com/ri-sdk-go/pkg/rapididentity"
)

func main() {
	options := rapididentity.Options{
		HTTPClient:      http.Client{},
		HostUrl:         "https://portal.us001-rapididentity.com",
		ServiceIdentity: "service_identity_key",
	}

	client := rapididentity.New(options)

	input := rapididentity.GetConnectFileContentInput{
		Path:         "log/jobs/GitCloneProject/2024-12-23/2024-12-23-00_31_12.870.html.gz",
		Project:      "gitops",
		Decompress:   true,
		ResponseType: "text/html",
	}

	output, err := client.GetConnectFileContent(input)
	if err != nil {
		fmt.Println(err)
	}

	fmt.Printf("%s\n", output)
}
