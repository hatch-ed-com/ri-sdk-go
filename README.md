# RapidIdentity GO SDK

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
options := rapididentity.Options{
    HTTPClient:      http.Client{},
    HostUrl:         "https://portal.us001-rapididentity.com",
    ServiceIdentity: "service_identity_key",
}

client := rapididentity.New(options)

input := rapididentity.GetConnectFilesInput{
    Path:    "/",
    Project: "sec_mgr",
}

output, err := client.GetConnectFiles(input)
if err != nil {
    fmt.Println(err)
}

fmt.Printf("%+v\n", output)
```

###### Compile and Execute

```sh
go run .
# Connect Files Output will be printed
```
