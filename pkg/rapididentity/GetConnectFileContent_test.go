package rapididentity

import (
	"context"
	"fmt"
	"net/http"
	"testing"
)

func TestGetConnectFileContent(t *testing.T) {
	t.Parallel()
	client, mux := setup(t)
	mux.HandleFunc(baseUrlPath+"/admin/connect/fileContent/{filePath...}", func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, "GET")
		testHeader(t, r, "Authorization", "Bearer "+mockServiceIdentity)
		testHeader(t, r, "Accept", "text/plain")
		testQueryParam(t, r, "project", "sec_mgr")
		testQueryParam(t, r, "decompress", "true")
		filePath := r.PathValue("filePath")
		if filePath == "" {
			w.WriteHeader(http.StatusBadRequest)
			fmt.Fprintln(w, "a file path is required")
			return
		}
		w.WriteHeader(http.StatusOK)
		fmt.Fprint(w, "Hello World")
	})

	input := GetConnectFileContentInput{
		Path:         "/hello/world.txt",
		Project:      "sec_mgr",
		Decompress:   true,
		ResponseType: "text/plain",
	}
	ctx := context.Background()
	output, err := client.GetConnectFileContent(ctx, input)
	if err != nil {
		t.Errorf("got error %s, want none", err)
	}

	got := output
	want := "Hello World"
	if string(got) != want {
		t.Errorf("got %s. want %s", got, want)
	}
}
