package rapididentity

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"strings"
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
	ctx := context.Background()
	output, err := client.GetAuthenticationPoliciesForUser(ctx, input)
	if err != nil {
		t.Errorf("got error %s, want none", err)
	}

	got := output.User.Username
	want := input.User.Username

	if got != want {
		t.Errorf("got %s. want %s", got, want)
	}
}

func TestGetBootstrapInfo(t *testing.T) {
	t.Parallel()
	client, mux := setup(t)
	mux.HandleFunc(baseUrlPath+"/bootstrapInfo", func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, "GET")
		testHeader(t, r, "Authorization", "Bearer "+mockServiceIdentity)
		testHeader(t, r, "Accept", "application/json")
		w.WriteHeader(http.StatusOK)
		fmt.Fprint(w,
			`{ 
				"tenantId": "1234"
			}`,
		)
	})
	ctx := context.Background()

	output, err := client.GetBootstrapInfo(ctx)
	if err != nil {
		t.Errorf("got error %s, want none", err)
	}

	got := output.TenantId
	want := "1234"
	if got != want {
		t.Errorf("got %s. want %s", got, want)
	}
}

func TestGetLdapAttributes(t *testing.T) {
	t.Parallel()
	client, mux := setup(t)
	mux.HandleFunc(baseUrlPath+"/admin/ldap/schema/attributes", func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, "GET")
		testHeader(t, r, "Authorization", "Bearer "+mockServiceIdentity)

		w.WriteHeader(http.StatusOK)
		fmt.Fprint(w,
			`["cn"]`,
		)
	})

	ctx := context.Background()
	output, err := client.GetRapidIdentityAttributes(ctx)
	if err != nil {
		t.Errorf("got error %s, want none", err)
	}

	got := output[0]
	want := "cn"

	if got != want {
		t.Errorf("got %s. want %s", got, want)
	}
}

func TestGetRapidIdentityAttributes_MarshalJSON_ZeroValue(t *testing.T) {
	t.Parallel()

	var list StringList = nil
	marshaledBytes, err := json.Marshal(list)
	if err != nil {
		t.Fatalf("failed to marshal StringList: %v", err)
	}

	result := string(marshaledBytes)
	want := "[]"
	if result != want {
		t.Errorf("got %s, want %s", result, want)
	}
}
func TestConfiguration_MarshalJSON_ZeroValue(t *testing.T) {
	t.Parallel()

	tests := []struct {
		name        string
		input       interface{}
		mustContain string
	}{
		{
			name: "GetAuthenticationPoliciesForUserInput with nil slices",
			input: GetAuthenticationPoliciesForUserInput{
				AuthenticationPolicyFieldsToShow: nil,
			},
			mustContain: `"authenticationPolicyFieldsToShow":[]`,
		},
		{
			name: "GetAuthenticationPoliciesForUserOutput with nil AuthenticationPolicies",
			input: GetAuthenticationPoliciesForUserOutput{
				AuthenticationPolicies: nil,
			},
			mustContain: `"authenticationPolicies":[]`,
		},
		{
			name: "AuthenticationPolicy with nil slices",
			input: AuthenticationPolicy{
				Id:       "pol-1",
				Criteria: nil,
				Methods:  nil,
			},
			mustContain: `"criteria":[]`,
		},
		{
			name: "SourceNetworkCriteria with nil Subnets",
			input: SourceNetworkCriteria{
				Subnets: nil,
			},
			mustContain: `"subnets":[]`,
		},
		{
			name: "RoleCriteria with nil Roles",
			input: RoleCriteria{
				Roles: nil,
			},
			mustContain: `"roles":[]`,
		},
		{
			name: "PictographMethod with nil ImageIds",
			input: PictographMethod{
				ImageIds: nil,
			},
			mustContain: `"imageIds":[]`,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			marshaledBytes, err := json.Marshal(tt.input)
			if err != nil {
				t.Fatalf("failed to marshal struct: %v", err)
			}

			result := string(marshaledBytes)

			if !strings.Contains(result, tt.mustContain) {
				t.Errorf("expected JSON to contain %q, but got: %s", tt.mustContain, result)
			}

			if strings.Contains(result, ":null") {
				t.Errorf("detected unexpected 'null' value in marshaled output: %s", result)
			}
		})
	}
}
