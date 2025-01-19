package rapididentity

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
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

	output, err := client.RunUserQuery(input)
	if err != nil {
		t.Errorf("got error %s, want none", err)
	}

	got := output[0].FirstName
	want := input.Query.FieldValue

	if got != want {
		t.Errorf("got %s. want %s", got, want)
	}
}
