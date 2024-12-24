// # The RapidIdentity SDK for the go programming language
//
// ## Getting Started
//
// To get started working with the SDK setup your project for Go modules, and retrieve the SDK dependencies with go
// get. Be sure to create your Service Identity Key with the appropriate permissions.
// This example shows how you can use the RapidIdentity Go SDK to make an API request using the SDK's
// RapidIdentithy client.
//
//	go get github.com/hatch-ed-com/ri-sdk-go/pkg/rapididentity
//
// ## Hello RapidIdentity
//
//	package main
//
//	import (
//		"fmt"
//		"net/http"
//
//		"github.com/hatch-ed-com/ri-sdk-go/pkg/rapididentity"
//	)
//
//	func main() {
//		options := rapididentity.Options{
//			HTTPClient:      http.Client{},
//			HostUrl:         "https://portal.us001-rapididentity.com",
//			ServiceIdentity: "service_identity_key",
//		}
//
//		client := rapididentity.New(options)
//
//		input := rapididentity.GetConnectFilesInput{
//			Path:    "/",
//			Project: "sec_mgr",
//		}
//
//		output, err := client.GetConnectFiles(input)
//		if err != nil {
//			fmt.Println(err)
//		}
//
//		fmt.Printf("%+v\n", output)
//	}
package sdk
