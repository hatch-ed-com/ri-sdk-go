package rapididentity

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"strings"
	"testing"
)

func TestGetConnectActions(t *testing.T) {
	t.Parallel()
	client, mux := setup(t)
	mux.HandleFunc(baseUrlPath+"/admin/connect/actions", func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, "GET")
		testHeader(t, r, "Authorization", "Bearer "+mockServiceIdentity)
		testQueryParam(t, r, "project", "sec_mgr")
		testQueryParam(t, r, "metaDataOnly", "true")
		w.WriteHeader(http.StatusOK)
		fmt.Fprint(w,
			`{ 
				"name": "all"
			}`,
		)
	})

	input := GetConnectActionsInput{
		Project:      "sec_mgr",
		MetaDataOnly: true,
	}
	ctx := context.Background()
	output, err := client.GetConnectActions(ctx, input)
	if err != nil {
		t.Errorf("got error %s, want none", err)
	}

	got := output.Name
	want := "all"

	if got != want {
		t.Errorf("got %s. want %s", got, want)
	}
}

func TestGetConnectActionsOutput_MarshalJSON_ZeroValue(t *testing.T) {
	t.Parallel()

	tests := []struct {
		name        string
		input       interface{}
		mustContain string
	}{
		{
			name: "GetConnectActionsOutput with nil ActionDefs",
			input: GetConnectActionsOutput{
				Name:       "test-actions",
				ActionDefs: nil,
			},
			mustContain: `"actionDefs":[]`,
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
