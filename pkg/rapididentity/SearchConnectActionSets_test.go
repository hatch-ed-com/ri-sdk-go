package rapididentity

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"strings"
	"testing"
)

func TestSearchConnectActionSets(t *testing.T) {
	t.Parallel()
	client, mux := setup(t)
	mux.HandleFunc(baseUrlPath+"/admin/connect/search/actions", func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, "GET")
		testHeader(t, r, "Authorization", "Bearer "+mockServiceIdentity)
		testQueryParam(t, r, "project", "sec_mgr")
		testQueryParam(t, r, "searchString", "FnAuthenticate")
		testQueryParam(t, r, "matchAction", "false")
		testQueryParam(t, r, "matchCase", "false")
		testQueryParam(t, r, "regex", "false")
		w.WriteHeader(http.StatusOK)
		fmt.Fprint(w,
			`{ 
				"name": "result"
			}`,
		)
	})

	input := SearchConnectActionSetsInput{
		SearchString: "FnAuthenticate",
		Project:      "sec_mgr",
	}
	ctx := context.Background()
	output, err := client.SearchConnectActionSets(ctx, input)
	if err != nil {
		t.Errorf("got error %s, want none", err)
	}

	got := output.Name
	want := "result"

	if got != want {
		t.Errorf("got %s. want %s", got, want)
	}
}

func TestActionDefSlice_MarshalJSON_ZeroValue(t *testing.T) {
	t.Parallel()

	// Define a test case matrix for all types that should force `[]`
	tests := []struct {
		name        string
		input       interface{}
		mustContain string
	}{
		{
			name: "SearchConnectActionSetsOutput with nil ActionDefs",
			input: SearchConnectActionSetsOutput{
				Name:       "test-query",
				ActionDefs: nil,
			},
			mustContain: `"actionDefs":[]`,
		},
		{
			name: "ActionDef with nil ArgDefs and Actions",
			input: ActionDef{
				Id:      "action-1",
				ArgDefs: nil,
				Actions: nil,
			},
			mustContain: `"argDefs":[]`,
		},
		{
			name: "ConnectAction with nil Args",
			input: ConnectAction{
				Id:   "connect-1",
				Args: nil,
			},
			mustContain: `"args":[]`,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			marshaledBytes, err := json.Marshal(tt.input)
			if err != nil {
				t.Fatalf("failed to marshal struct: %v", err)
			}

			result := string(marshaledBytes)

			// Assert that it overrode null with an empty array literal
			if !strings.Contains(result, tt.mustContain) {
				t.Errorf("expected JSON to contain %q, but got: %s", tt.mustContain, result)
			}

			// Defensive check: ensure "null" didn't sneak in anywhere
			if strings.Contains(result, ":null") {
				t.Errorf("detected unexpected 'null' value in marshaled output: %s", result)
			}
		})
	}
}
