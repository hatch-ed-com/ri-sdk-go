package main

import (
	"fmt"
	"net/http"
	"os"

	"github.com/hatch-ed-com/ri-sdk-go/pkg/rapididentity"
)

func main() {
	options := rapididentity.Options{
		HTTPClient:      http.Client{},
		HostUrl:         "https://portal.us001-rapididentity.com",
		ServiceIdentity: "service_identity_key",
	}

	client := rapididentity.New(options)

	input := rapididentity.GetConnectFileContentZip{
		PathList: []string{
			"log/run/RESTPointAPIGateway/2024-09-09/2024-09-09-16_18_24.798.html.gz",
			"log/run/RESTPointAPIGateway/2024-09-09/2024-09-09-16_25_40.972.html.gz",
			"log/run/RESTPointAPIGateway/2024-09-09/2024-09-09-16_32_53.006.html.gz",
			"log/run/RESTPointAPIGateway/2024-09-09/2024-09-09-16_32_53.042.html.gz",
		},
		Project: "sec_mgr",
	}

	output, err := client.GetConnectFileContentZip(input)
	if err != nil {
		fmt.Println(err)
	}

	f, _ := os.Create("myFiles.zip")
	defer f.Close()
	f.Write(output)
}
