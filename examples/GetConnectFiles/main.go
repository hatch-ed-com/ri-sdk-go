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

	input := rapididentity.GetConnectFilesInput{
		Path:    "/",
		Project: "sec_mgr",
	}

	output, err := client.GetConnectFiles(input)
	if err != nil {
		fmt.Println(err)
	}

	fmt.Printf("%+v\n", output)
}
