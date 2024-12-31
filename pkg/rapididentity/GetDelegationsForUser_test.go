package rapididentity

import (
	"fmt"
	"net/http"
	"testing"
)

func TestGetDelegationsForUser(t *testing.T) {
	t.Parallel()
	client, mux := setup(t)
	mux.HandleFunc(baseUrlPath+"/profiles/aggregated/for/{userId}", func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, "GET")
		testHeader(t, r, "Authorization", "Bearer "+mockServiceIdentity)
		w.WriteHeader(http.StatusOK)
		userId := r.PathValue("userId")
		if userId == "" {
			w.WriteHeader(http.StatusBadRequest)
			fmt.Fprintln(w, "a user id is required")
			return
		}
		fmt.Fprintf(w,
			`{ 
				"aggregatedDelegation": {
					"id":"%s"
				}
			}`,
			userId,
		)
	})

	input := GetDelegationsForUserInput{
		UserId: "1234",
	}
	output, err := client.GetDelegationsForUser(input)
	if err != nil {
		t.Errorf("got error %s, want none", err)
	}

	got := output.AggregatedDelegation.Id
	want := input.UserId

	if got != want {
		t.Errorf("got %s. want %s", got, want)
	}
}
