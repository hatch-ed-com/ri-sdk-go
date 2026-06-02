package main

import (
	"context"
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

	client, err := rapididentity.New(options)
	if err != nil {
		riError, ok := err.(rapididentity.RapidIdentityError)
		if ok {
			log.Fatalf("Request URL: %s, Status Code: %d, Message: %s", riError.ReqUrl, riError.Code, riError.Message)
		}
		log.Fatal(err)
	}

	input := rapididentity.SaveConnectActionInput{
		Action: rapididentity.ActionDef{
			Id:           "A5500922-4C5D-43B8-B407-93B42DEA96E2",
			Version:      0,
			Name:         "TestRISDK",
			Description:  "This is a test from RI SDK",
			ReturnsValue: false,
			Actions: rapididentity.ConnectActionList{
				{
					Args: rapididentity.ArgDefList{
						{
							Description: "The message to be logged",
							Name:        "message",
							Type:        "string",
							Value:       "\"This is a test from the RI SDK\"",
						},
						{
							Description: "The log level (default: INFO)",
							Name:        "level",
							Type:        "string",
							Value:       "\"WARN\"",
						},
					},
					Id:   "7F3F77A5-737E-4036-A75D-8DD39A336ED1",
					Name: "log",
				},
			},
		},
	}

	ctx := context.Background()
	output, err := client.SaveConnectAction(ctx, input)
	if err != nil {
		riError, ok := err.(rapididentity.RapidIdentityError)
		if ok {
			log.Fatalf("Request URL: %s, Status Code: %d, Message: %s", riError.ReqUrl, riError.Code, riError.Message)
		}
		log.Fatal(err)
	}

	fmt.Printf("%+v\n", output.Action)

}
