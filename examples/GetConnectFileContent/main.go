package main

import (
	"fmt"
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

	input := rapididentity.GetConnectFileContentInput{
		Path:         "log/run/RESTPointAPIGateway/2024-09-09/2024-09-09-16_18_24.798.html.gz",
		Project:      "sec_mgr",
		Decompress:   true,
		ResponseType: "text/html",
	}

	output, err := client.GetConnectFileContent(input)
	if err != nil {
		riError, ok := err.(rapididentity.RapidIdentityError)
		if ok {
			log.Fatalf("Request URL: %s, Status Code: %d, Message: %s", riError.ReqUrl, riError.Code, riError.Message)
		}
		log.Fatal(err)
	}

	fmt.Printf("%s\n", output)
}
