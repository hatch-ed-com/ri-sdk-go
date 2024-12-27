package main

import (
	"log"
	"net/http"
	"net/url"
	"os"

	"github.com/hatch-ed-com/ri-sdk-go/pkg/rapididentity"
)

func main() {
	baseUrl, err := url.Parse(os.Getenv("RI_URL"))
	if err != nil {
		log.Fatal(err)
	}
	options := rapididentity.Options{
		HTTPClient:      &http.Client{},
		BaseUrl:         baseUrl,
		ServiceIdentity: os.Getenv("RI_KEY"),
	}

	client := rapididentity.New(options)

	input := rapididentity.GetConnectFileContentZipInput{
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
		riError, ok := err.(rapididentity.RapidIdentityError)
		if ok {
			log.Fatalf("Request URL: %s, Status Code: %d, Message: %s", riError.ReqUrl, riError.Code, riError.Message)
		}
		log.Fatal(err)
	}

	f, _ := os.Create("myFiles.zip")
	defer f.Close()
	f.Write(output)
}
