package rapididentity

import (
	"context"
	"fmt"
	"net/http"
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
