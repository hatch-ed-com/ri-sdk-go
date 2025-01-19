# RapidIdentity GO SDK

[![ri-sdk-go release (latest SemVer)](https://img.shields.io/github/v/release/hatch-ed-com/ri-sdk-go?sort=semver)](https://github.com/hatch-ed-com/ri-sdk-go/releases)
[![Go Reference](https://img.shields.io/static/v1?label=godoc&message=reference&color=blue)](https://pkg.go.dev/github.com/hatch-ed-com/ri-sdk-go/pkg/rapididentity)

The RapidIdentity SDK for the go programming language

## Getting Started

To get started working with the SDK setup your project for Go modules, and retrieve the SDK dependencies with go
get. Be sure to create your [Service Identity Key][1] with the appropriate permissions.
This example shows how you can use the RapidIdentity Go SDK to make an API request using the SDK's
RapidIdentithy client.

[1]: https://help.rapididentity.com/docs/service-identities-in-rapididentity

###### Initialize Project

```sh
mkdir ~/hellori
cd ~/hellori
go mod init hellori
```

###### Add Dependencies

```sh
go get github.com/hatch-ed-com/ri-sdk-go/pkg/rapididentity
```

###### Write Code

In your preferred editor add the following content to `main.go`

```go
package main

import (
	"context"
	"fmt"
	"net/http"

	"github.com/hatch-ed-com/ri-sdk-go/pkg/rapididentity"
)

func main() {
	options := rapididentity.Options{
		HTTPClient:      &http.Client{},
		BaseUrl:         "https://portal.us001-rapididentity.com",
		ServiceIdentity: "service_identity_key",
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

	input := rapididentity.GetConnectFilesInput{
		Path:    "/",
		Project: "sec_mgr",
	}

	ctx := context.Background()
	output, err := client.GetConnectFiles(ctx, input)
	if err != nil {
		riError, ok := err.(rapididentity.RapidIdentityError)
		if ok {
			log.Fatalf("Request URL: %s, Status Code: %d, Message: %s", riError.ReqUrl, riError.Code, riError.Message)
		}
		log.Fatal(err)
	}

	fmt.Printf("%+v\n", output)
}
```

###### Compile and Execute

```sh
go run .
# Connect Files Output will be printed
```
