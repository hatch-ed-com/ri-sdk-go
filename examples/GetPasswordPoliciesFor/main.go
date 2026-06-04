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

	input := rapididentity.GetPasswordPoliciesForInput{
		UserIds: []string{"019e4ba9-95aa-7a17-b15d-0e1c8ca70467"},
	}

	ctx := context.Background()
	output, err := client.GetPasswordPoliciesFor(ctx, input)
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
