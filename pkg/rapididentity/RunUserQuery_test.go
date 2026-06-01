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

func TestRunUserQuery(t *testing.T) {
	t.Parallel()
	client, mux := setup(t)
	mux.HandleFunc(baseUrlPath+"/users", func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, "POST")
		testHeader(t, r, "Content-Type", "application/json")
		testHeader(t, r, "Authorization", "Bearer "+mockServiceIdentity)
		testQueryParam(t, r, "limit", "1000")
		testQueryParam(t, r, "search", "advanced")
		testQueryParam(t, r, "did", "white_pages")
		var searchPayload AuditReportQuery
		reqBody, err := io.ReadAll(r.Body)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			fmt.Fprintln(w, err)
			return
		}
		err = json.Unmarshal(reqBody, &searchPayload)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			fmt.Fprintln(w, err)
			return
		}
		w.WriteHeader(http.StatusOK)
		fmt.Fprintf(w,
			`[
				{
					"FirstName": "%s"
				}
			]`,
			searchPayload.FieldValue)
	})

	input := RunUserQueryInput{
		DelegationIds: []string{"white_pages"},
		Query: AuditReportQuery{
			FieldName:    "givenName",
			OperatorType: EQUAL,
			FieldValue:   "Jack",
		},
	}

	ctx := context.Background()
	output, err := client.RunUserQuery(ctx, input)
	if err != nil {
		t.Errorf("got error %s, want none", err)
	}

	got := output[0].FirstName
	want := input.Query.FieldValue

	if got != want {
		t.Errorf("got %s. want %s", got, want)
	}
}

func TestRunUserQueryInput_MarshalJSON_ZeroValue(t *testing.T) {
	t.Parallel()

	tests := []struct {
		name        string
		input       RunUserQueryInput
		mustContain string
	}{
		{
			name: "RunUserQueryInput with nil DelegationIds",
			input: RunUserQueryInput{
				DelegationIds: nil,
			},
			mustContain: `"delegationIds":[]`,
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

func TestUserList_MarshalJSON_ZeroValue(t *testing.T) {
	t.Parallel()

	var list UserList = nil
	marshaledBytes, err := json.Marshal(list)
	if err != nil {
		t.Fatalf("failed to marshal UserList: %v", err)
	}

	result := string(marshaledBytes)
	want := "[]"
	if result != want {
		t.Errorf("got %s, want %s", result, want)
	}
}
