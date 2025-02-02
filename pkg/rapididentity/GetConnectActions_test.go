package rapididentity

import (
	"context"
	"fmt"
	"net/http"
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
