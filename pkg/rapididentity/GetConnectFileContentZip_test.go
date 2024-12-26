package rapididentity

import (
	"fmt"
	"net/http"
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
	output, err := client.GetConnectFileContentZip(input)
	if err != nil {
		t.Errorf("got error %s, want none", err)
	}

	got := output
	want := "Hello World"
	if string(got) != want {
		t.Errorf("got %s. want %s", got, want)
	}
}
