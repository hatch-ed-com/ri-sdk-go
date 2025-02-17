package rapididentity

import (
	"context"
	"fmt"
	"net/http"
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
