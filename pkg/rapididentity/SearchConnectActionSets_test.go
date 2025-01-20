package rapididentity

import (
	"context"
	"fmt"
	"net/http"
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
