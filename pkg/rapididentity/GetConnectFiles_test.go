package rapididentity

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"strings"
	"testing"
)

func TestGetConnectFiles(t *testing.T) {
	t.Parallel()
	client, mux := setup(t)
	mux.HandleFunc(baseUrlPath+"/admin/connect/files/{filePath...}", func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, "GET")
		testHeader(t, r, "Authorization", "Bearer "+mockServiceIdentity)
		testQueryParam(t, r, "project", "sec_mgr")
		filePath := r.PathValue("filePath")
		if filePath == "" {
			w.WriteHeader(http.StatusBadRequest)
			fmt.Fprintln(w, "a file path is required")
			return
		}
		w.WriteHeader(http.StatusOK)
		fmt.Fprintf(w,
			`{ 
				"path": "%s"
			}`,
			filePath,
		)
	})

	input := GetConnectFilesInput{
		Path:    "log/jobs",
		Project: "sec_mgr",
	}
	ctx := context.Background()
	output, err := client.GetConnectFiles(ctx, input)
	if err != nil {
		t.Errorf("got error %s, want none", err)
	}

	got := output.Path
	want := input.Path

	if got != want {
		t.Errorf("got %s. want %s", got, want)
	}
}

func TestGetConnectFilesOutput_MarshalJSON_ZeroValue(t *testing.T) {
	t.Parallel()

	tests := []struct {
		name        string
		input       interface{}
		mustContain string
	}{
		{
			name: "GetConnectFilesOutput with nil FileEntries",
			input: GetConnectFilesOutput{
				FileEntry: FileEntry{
					Path: "test-path",
				},
				FileEntries: nil,
			},
			mustContain: `"fileEntries":[]`,
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
