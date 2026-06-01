package rapididentity

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"strings"
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

func TestRunAuditReportPagination(t *testing.T) {
	t.Parallel()
	client, mux := setup(t)
	mux.HandleFunc(baseUrlPath+"/reporting/auditQuery", func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, "POST")
		testQueryParam(t, r, "page_size", "10")
		testQueryParam(t, r, "page_token", "abc123")
		w.WriteHeader(http.StatusOK)
		fmt.Fprint(w, `{"auditLogRecords": [], "nextPageToken": "def456"}`)
	})

	input := RunAuditReportInput{
		Query:     AuditReportQuery{OperatorType: AND},
		PageSize:  10,
		PageToken: "abc123",
	}

	ctx := context.Background()
	output, err := client.RunAuditReport(ctx, input)
	if err != nil {
		t.Errorf("got error %s, want none", err)
	}

	got := output.NextPageToken
	want := "def456"

	if got != want {
		t.Errorf("got %s, want %s", got, want)
	}
}

func TestRunAuditReportOutput_MarshalJSON_ZeroValue(t *testing.T) {
	t.Parallel()

	tests := []struct {
		name        string
		input       interface{}
		mustContain string
	}{
		{
			name: "AuditReportQuery with nil slices",
			input: AuditReportQuery{
				OperatorType: AND,
				ChildNodes:   nil,
				FieldValues:  nil,
			},
			mustContain: `"childNodes":[]`,
		},
		{
			name: "RunAuditReportOutput with nil AuditLogRecords",
			input: RunAuditReportOutput{
				AuditLogRecords: nil,
			},
			mustContain: `"auditLogRecords":[]`,
		},
		{
			name: "AuditReportResult with nil ExtendedProperties",
			input: AuditReportResult{
				Id:                 "res-1",
				ExtendedProperties: nil,
			},
			mustContain: `"extendedProperties":[]`,
		},
		{
			name: "AuditReportActionDetail with nil Categories",
			input: AuditReportActionDetail{
				AuditReportBaseDetail: AuditReportBaseDetail{Id: "act-1"},
				Categories:            nil,
			},
			mustContain: `"categories":[]`,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			marshaledBytes, err := json.Marshal(tt.input)
			if err != nil {
				t.Fatalf("failed to marshal struct: %v", err)
			}

			result := string(marshaledBytes)

			if !strings.Contains(result, tt.mustContain) {
				t.Errorf("expected JSON to contain %q, but got: %s", tt.mustContain, result)
			}

			if strings.Contains(result, ":null") {
				t.Errorf("detected unexpected 'null' value in marshaled output: %s", result)
			}
		})
	}
}
