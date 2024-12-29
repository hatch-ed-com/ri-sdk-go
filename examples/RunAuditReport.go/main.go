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

	client, err := rapididentity.New(options)
	if err != nil {
		riError, ok := err.(rapididentity.RapidIdentityError)
		if ok {
			log.Fatalf("Request URL: %s, Status Code: %d, Message: %s", riError.ReqUrl, riError.Code, riError.Message)
		}
		log.Fatal(err)
	}

	input := rapididentity.RunAuditReportInput{
		Query: rapididentity.AuditReportQuery{
			ChildNodes: []rapididentity.AuditReportQuery{
				{
					ChildNodes: []rapididentity.AuditReportQuery{
						{
							FieldName:          "action.displayName",
							FieldSecondaryName: "action.displayName",
							FieldValue:         "IdP Authentication",
							OperatorType:       rapididentity.EQUAL,
						},
						{
							FieldName:          "action.displayName",
							FieldSecondaryName: "action.displayName",
							FieldValue:         "Change Password",
							OperatorType:       rapididentity.EQUAL,
						},
					},
					OperatorType: rapididentity.OR,
				},
				{
					FieldName:          "timestamp",
					FieldSecondaryName: "timestamp",
					FieldValues: []rapididentity.AuditReportFieldValue{
						{
							Dn:                   "LAST_7_DAYS",
							FieldNameAndServerId: "LAST_7_DAYS",
							Id:                   "LAST_7_DAYS",
							Name:                 "LAST_7_DAYS",
						},
					},
					OperatorType: rapididentity.EQUAL,
				},
			},
			OperatorType: rapididentity.AND,
		},
	}

	output, err := client.RunAuditReport(input)
	if err != nil {
		riError, ok := err.(rapididentity.RapidIdentityError)
		if ok {
			log.Fatalf("Request URL: %s, Status Code: %d, Message: %s", riError.ReqUrl, riError.Code, riError.Message)
		}
		log.Fatal(err)
	}

	fmt.Printf("%+v\n", output)

}
