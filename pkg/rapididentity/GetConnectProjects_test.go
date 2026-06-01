package rapididentity

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"strings"
	"testing"
)

func TestGetConnectProjects(t *testing.T) {
	t.Parallel()
	client, mux := setup(t)
	mux.HandleFunc(baseUrlPath+"/admin/connect/projects", func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, "GET")
		testHeader(t, r, "Authorization", "Bearer "+mockServiceIdentity)

		w.WriteHeader(http.StatusOK)
		fmt.Fprint(w,
			`{
				"projects": [
					{ 
						"id": "1234"
					}
				]
			}`,
		)
	})

	ctx := context.Background()
	output, err := client.GetConnectProjects(ctx)
	if err != nil {
		t.Errorf("got error %s, want none", err)
	}

	got := output.Projects[0].Id
	want := "1234"

	if got != want {
		t.Errorf("got %s. want %s", got, want)
	}
}

func TestGetConnectProjectsOutput_MarshalJSON_ZeroValue(t *testing.T) {
	t.Parallel()

	tests := []struct {
		name        string
		input       interface{}
		mustContain string
	}{
		{
			name: "GetConnectProjectsOutput with nil Projects",
			input: GetConnectProjectsOutput{
				Projects: nil,
			},
			mustContain: `"projects":[]`,
		},
		{
			name: "ConnectProject with nil RestPoints",
			input: ConnectProject{
				Id: "proj-1",
				RestPoints: RestPointConfig{
					RestPoints: nil,
				},
			},
			mustContain: `"restPoints":[]`,
		},
		{
			name: "RestPoint with nil ArgMap",
			input: RestPoint{
				Id:     "rp-1",
				ArgMap: nil,
			},
			mustContain: `"argMap":[]`,
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
