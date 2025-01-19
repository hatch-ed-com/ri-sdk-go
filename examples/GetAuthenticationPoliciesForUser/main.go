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
	defer client.Close()

	input := rapididentity.GetAuthenticationPoliciesForUserInput{
		ShowAuthenticationPolicies: true,
		ShowClaims:                 true,
		User: rapididentity.GetAuthenticationPoliciesForUserPayload{
			Username: os.Getenv("RI_USER"),
		},
	}

	ctx := context.Background()
	output, err := client.GetAuthenticationPoliciesForUser(ctx, input)
	if err != nil {
		riError, ok := err.(rapididentity.RapidIdentityError)
		if ok {
			log.Fatalf("Request URL: %s, Status Code: %d, Message: %s", riError.ReqUrl, riError.Code, riError.Message)
		}
		log.Fatal(err)
	}

	for _, method := range output.AuthenticationPolicies[0].Methods {
		if method.GetBaseAuthenticationMethodInfo().Type == "webAuthn" {
			webAuthn := method.(rapididentity.WebAuthnMethod)
			fmt.Printf("Allow Change Deferral: %t", webAuthn.AllowChallengeDeferral)
		}
	}
}
