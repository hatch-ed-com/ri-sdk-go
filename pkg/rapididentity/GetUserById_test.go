package rapididentity

import (
	"context"
	"fmt"
	"net/http"
	"testing"
)

func TestGetUserById(t *testing.T) {
	t.Parallel()
	client, mux := setup(t)
	mux.HandleFunc(baseUrlPath+"/admin/ldap/users/{dnOrId}", func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, "GET")
		testHeader(t, r, "Authorization", "Bearer "+mockServiceIdentity)
		dnOrId := r.PathValue("dnOrId")
		if dnOrId == "" {
			w.WriteHeader(http.StatusBadRequest)
			fmt.Fprintln(w, "a DN or idautoID is required")
			return
		}
		w.WriteHeader(http.StatusOK)
		fmt.Fprintf(w,
			`{ 
				"id": "%s"
			}`,
			dnOrId,
		)
	})

	input := GetUserByIdInput{
		Id: "1234",
	}

	ctx := context.Background()
	output, err := client.GetUserById(ctx, input)
	if err != nil {
		t.Errorf("got error %s, want none", err)
	}

	got := output.Id
	want := input.Id

	if got != want {
		t.Errorf("got %s. want %s", got, want)
	}
}
