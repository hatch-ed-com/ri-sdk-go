package main

import (
	"fmt"
	"net/http"

	"github.com/hatch-ed-com/ri-sdk-go/pkg/rapididentity"
)

func main() {
	options := rapididentity.Options{
		HTTPClient:      http.Client{},
		HostUrl:         "https://portal.us001-rapididentity.com",
		ServiceIdentity: "service_identity_key",
	}

	client := rapididentity.New(options)

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
		fmt.Println(err)
	}

	fmt.Printf("%+v\n", output)

}
