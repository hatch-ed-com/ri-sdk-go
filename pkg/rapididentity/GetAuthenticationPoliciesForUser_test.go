package rapididentity

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"testing"
)

func TestGetAuthenticationPoliciesForUser(t *testing.T) {
	t.Parallel()
	client, mux := setup(t)
	mux.HandleFunc(baseUrlPath+"/authn/v1/username", func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, "POST")
		testHeader(t, r, "Content-Type", "application/json")
		testHeader(t, r, "Authorization", "Bearer "+mockServiceIdentity)
		testQueryParam(t, r, "claim", "true")
		testQueryParam(t, r, "authenticationPolicies", "true")
		testQueryParam(t, r, "authenticationPolicyField", "methods")
		var userPayload GetAuthenticationPoliciesForUserPayload
		reqBody, err := io.ReadAll(r.Body)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			fmt.Fprintln(w, err)
			return
		}
		err = json.Unmarshal(reqBody, &userPayload)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			fmt.Fprintln(w, err)
			return
		}
		w.WriteHeader(http.StatusOK)
		fmt.Fprintf(w,
			`{ 
				"user": {
						"username": "%s"
					}
			}`,
			userPayload.Username)
	})

	input := GetAuthenticationPoliciesForUserInput{
		ShowAuthenticationPolicies:       true,
		ShowClaims:                       true,
		AuthenticationPolicyFieldsToShow: []string{"methods"},
		User: GetAuthenticationPoliciesForUserPayload{
			Username: "user@example.com",
		},
	}
	output, err := client.GetAuthenticationPoliciesForUser(input)
	if err != nil {
		t.Errorf("got error %s, want none", err)
	}

	got := output.User.Username
	want := input.User.Username

	if got != want {
		t.Errorf("got %s. want %s", got, want)
	}
}
