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

	input := rapididentity.GetDelegationsForUserInput{
		UserId: "08b5f0ec-d56a-4712-ada5-c86074ab11db",
	}

	output, err := client.GetDelegationsForUser(input)
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
