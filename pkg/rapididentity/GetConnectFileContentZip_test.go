package rapididentity

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"strings"
	"testing"
)

func TestGetConnectFileContentZip(t *testing.T) {
	t.Parallel()
	client, mux := setup(t)
	mux.HandleFunc(baseUrlPath+"/admin/connect/fileContentZip", func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, "GET")
		testHeader(t, r, "Authorization", "Bearer "+mockServiceIdentity)
		testHeader(t, r, "Accept", "application/zip")
		testQueryParam(t, r, "project", "sec_mgr")
		testQueryParam(t, r, "path", "/hello/world.txt")
		w.WriteHeader(http.StatusOK)
		fmt.Fprint(w, "Hello World")
	})

	input := GetConnectFileContentZipInput{
		PathList: []string{"/hello/world.txt", "/foo/bar.txt"},
		Project:  "sec_mgr",
	}
	ctx := context.Background()
	output, err := client.GetConnectFileContentZip(ctx, input)
	if err != nil {
		t.Errorf("got error %s, want none", err)
	}

	got := output
	want := "Hello World"
	if string(got) != want {
		t.Errorf("got %s. want %s", got, want)
	}
}

func TestGetConnectFileContentZipInput_MarshalJSON_ZeroValue(t *testing.T) {
	t.Parallel()

	tests := []struct {
		name        string
		input       GetConnectFileContentZipInput
		mustContain string
	}{
		{
			name: "GetConnectFileContentZipInput with nil PathList",
			input: GetConnectFileContentZipInput{
				PathList: nil,
			},
			mustContain: `"pathList":[]`,
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
