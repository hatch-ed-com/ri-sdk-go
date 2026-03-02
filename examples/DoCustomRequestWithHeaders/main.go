package main

import (
	"bytes"
	"context"
	"errors"
	"fmt"
	"io"
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
		var riError rapididentity.RapidIdentityError
		ok := errors.As(err, &riError)
		if ok {
			log.Fatalf("Request URL: %s, Status Code: %d, Message: %s", riError.ReqUrl, riError.Code, riError.Message)
		}
		log.Fatal(err)
	}
	defer client.Close()

	body := bytes.NewBufferString("plain text body")
	headers := http.Header{}
	headers.Set("Content-Type", "text/plain")

	ctx := context.Background()
	res, err := client.DoCustomRequestWithHeaders(ctx, "POST", "connect/actionSets/execute", headers, body)
	if err != nil {
		var riError rapididentity.RapidIdentityError
		ok := errors.As(err, &riError)
		if ok {
			log.Fatalf("Request URL: %s, Status Code: %d, Message: %s", riError.ReqUrl, riError.Code, riError.Message)
		}
		log.Fatal(err)
	}

	defer res.Body.Close()
	resBody, err := io.ReadAll(res.Body)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("%s\n", resBody)
}
