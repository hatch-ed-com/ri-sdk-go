package rapididentity

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"testing"
)

func TestRunAuditReport(t *testing.T) {
	t.Parallel()
	client, mux := setup(t)
	mux.HandleFunc(baseUrlPath+"/reporting/auditQuery", func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, "POST")
		testHeader(t, r, "Content-Type", "application/json")
		testHeader(t, r, "Authorization", "Bearer "+mockServiceIdentity)
		var auditReportQuery AuditReportQuery
		reqBody, err := io.ReadAll(r.Body)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			fmt.Fprintln(w, err)
			return
		}
		err = json.Unmarshal(reqBody, &auditReportQuery)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			fmt.Fprintln(w, err)
			return
		}
		w.WriteHeader(http.StatusOK)
		fmt.Fprintf(w,
			`{ 
				"auditLogRecords": [
					{
						"id": "%s"
					}
				]
			}`,
			auditReportQuery.ChildNodes[0].FieldValues[0].Id)
	})

	input := RunAuditReportInput{
		Query: AuditReportQuery{
			ChildNodes: []AuditReportQuery{
				{
					FieldName: "target",
					FieldValues: []AuditReportFieldValue{
						{
							Dn:                   "idautoID=1234,dc=accounts,dc=meta",
							FieldNameAndServerId: "Person",
							Id:                   "1234",
							Name:                 "John Doe",
						},
					},
				},
				{
					FieldName: "timestamp",
					FieldValues: []AuditReportFieldValue{
						{
							Dn:                   "LAST_7_DAYS",
							FieldNameAndServerId: "LAST_7_DAYS",
							Id:                   "LAST_7_DAYS",
							Name:                 "LAST_7_DAYS",
						},
					},
				},
			},
			OperatorType: AND,
		},
	}

	ctx := context.Background()
	output, err := client.RunAuditReport(ctx, input)
	if err != nil {
		t.Errorf("got error %s, want none", err)
	}

	got := output.AuditLogRecords[0].Id
	want := input.Query.ChildNodes[0].FieldValues[0].Id

	if got != want {
		t.Errorf("got %s. want %s", got, want)
	}
}
