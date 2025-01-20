package rapididentity

import (
	"context"
	"fmt"
	"net/http"
	"testing"
)

func TestGetConnectJobs(t *testing.T) {
	t.Parallel()
	client, mux := setup(t)
	mux.HandleFunc(baseUrlPath+"/admin/connect/jobs", func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, "GET")
		testHeader(t, r, "Authorization", "Bearer "+mockServiceIdentity)
		testQueryParam(t, r, "project", "")

		w.WriteHeader(http.StatusOK)
		fmt.Fprint(w,
			`{
				"jobs": [
					{ 
						"id": "1234"
					}
				]
			}`,
		)
	})

	input := GetConnectJobsInput{
		Project: MainProject,
	}
	ctx := context.Background()
	output, err := client.GetConnectJobs(ctx, input)
	if err != nil {
		t.Errorf("got error %s, want none", err)
	}

	got := output.Jobs[0].Id
	want := "1234"

	if got != want {
		t.Errorf("got %s. want %s", got, want)
	}
}
