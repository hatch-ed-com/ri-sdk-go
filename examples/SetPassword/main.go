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
		HTTPClient: &http.Client{},
		BaseUrl:    baseUrl,
		RapidIdentityUser: &rapididentity.RapidIdentityUser{
			Username: os.Getenv("RI_USER"),
			Password: os.Getenv("RI_PWD"),
		},
	}

	client, err := rapididentity.New(options)
	if err != nil {
		riError, ok := err.(rapididentity.RapidIdentityError)
		if ok {
			log.Fatalf("Method: %s, Request URL: %s, Status Code: %d, Message: %s, Reason: %s",
				riError.Method,
				riError.ReqUrl,
				riError.Code,
				riError.Message,
				riError.Reason)
		}
		log.Fatal(err)
	}
	defer client.Close()

	input := rapididentity.SetPasswordInput{
		IsSelfService: false,
		DelegationId:  "c6643ce0-1c9b-4280-8c2c-aa72b5689dce",
		MustUpdate:    false,
		Targets:       []string{"996e6333-0731-4d1d-9def-0d0d9afae523"},
		NewPassword:   "SuperSecret#123",
	}

	ctx := context.Background()
	output, err := client.SetPassword(ctx, input)
	if err != nil {
		riError, ok := err.(rapididentity.RapidIdentityError)
		if ok {
			log.Fatalf("Method: %s, Request URL: %s, Status Code: %d, Message: %s, Reason: %s",
				riError.Method,
				riError.ReqUrl,
				riError.Code,
				riError.Message,
				riError.Reason)
		}
		log.Fatal(err)
	}

	fmt.Printf("%+v\n", output)
}
