package rapididentity

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"strings"
	"testing"
)

func TestGetConnectJobs(t *testing.T) {
	t.Parallel()
	client, mux := setup(t)
	mux.HandleFunc(baseUrlPath+"/admin/connect/jobs", func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, "GET")
		testHeader(t, r, "Authorization", "Bearer "+mockServiceIdentity)
		testQueryParam(t, r, "project", "")

		w.WriteHeader(http.StatusOK)
		fmt.Fprint(w,
			`{
				"jobs": [
					{ 
						"id": "1234"
					}
				]
			}`,
		)
	})

	input := GetConnectJobsInput{
		Project: MainProject,
	}
	ctx := context.Background()
	output, err := client.GetConnectJobs(ctx, input)
	if err != nil {
		t.Errorf("got error %s, want none", err)
	}

	got := output.Jobs[0].Id
	want := "1234"

	if got != want {
		t.Errorf("got %s. want %s", got, want)
	}
}

func TestGetConnectJobsOutput_MarshalJSON_ZeroValue(t *testing.T) {
	t.Parallel()

	tests := []struct {
		name        string
		input       interface{}
		mustContain string
	}{
		{
			name: "GetConnectJobsOutput with nil Jobs",
			input: GetConnectJobsOutput{
				Jobs: nil,
			},
			mustContain: `"jobs":[]`,
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
