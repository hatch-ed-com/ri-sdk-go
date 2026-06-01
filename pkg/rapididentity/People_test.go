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
	ctx := context.Background()
	output, err := client.GetDelegationsForUser(ctx, input)
	if err != nil {
		t.Errorf("got error %s, want none", err)
	}

	got := output.AggregatedDelegation.Id
	want := input.UserId

	if got != want {
		t.Errorf("got %s. want %s", got, want)
	}
}

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

func TestRunUserQuery(t *testing.T) {
	t.Parallel()
	client, mux := setup(t)
	mux.HandleFunc(baseUrlPath+"/users", func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, "POST")
		testHeader(t, r, "Content-Type", "application/json")
		testHeader(t, r, "Authorization", "Bearer "+mockServiceIdentity)
		testQueryParam(t, r, "limit", "1000")
		testQueryParam(t, r, "search", "advanced")
		testQueryParam(t, r, "did", "white_pages")
		var searchPayload AuditReportQuery
		reqBody, err := io.ReadAll(r.Body)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			fmt.Fprintln(w, err)
			return
		}
		err = json.Unmarshal(reqBody, &searchPayload)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			fmt.Fprintln(w, err)
			return
		}
		w.WriteHeader(http.StatusOK)
		fmt.Fprintf(w,
			`[
				{
					"FirstName": "%s"
				}
			]`,
			searchPayload.FieldValue)
	})

	input := RunUserQueryInput{
		DelegationIds: []string{"white_pages"},
		Query: AuditReportQuery{
			FieldName:    "givenName",
			OperatorType: EQUAL,
			FieldValue:   "Jack",
		},
	}

	ctx := context.Background()
	output, err := client.RunUserQuery(ctx, input)
	if err != nil {
		t.Errorf("got error %s, want none", err)
	}

	got := output[0].FirstName
	want := input.Query.FieldValue

	if got != want {
		t.Errorf("got %s. want %s", got, want)
	}
}

func TestPeople_MarshalJSON_ZeroValue(t *testing.T) {
	t.Parallel()

	tests := []struct {
		name        string
		input       interface{}
		mustContain string
	}{
		{
			name: "AggregatedDelegation with nil slices",
			input: AggregatedDelegation{
				Id:                 "user-1",
				HelpdeskQuestions:  nil,
				DelegationProfiles: nil,
			},
			mustContain: `"helpdeskQuestions":[]`,
		},
		{
			name: "User with nil MobileNumbers",
			input: User{
				Id:            "user-1",
				MobileNumbers: nil,
			},
			mustContain: `"mobileNumbers":[]`,
		},
		{
			name: "Delegation with nil slices",
			input: Delegation{
				Id:         "del-1",
				Attributes: nil,
				Actions:    nil,
			},
			mustContain: `"attributes":[]`,
		},
		{
			name: "Profile with nil Attributes",
			input: Profile{
				Id:         "prof-1",
				Attributes: nil,
			},
			mustContain: `"attributes":[]`,
		},
		{
			name: "ProfileAttribute with nil Values",
			input: ProfileAttribute{
				Id:     "attr-1",
				Values: nil,
			},
			mustContain: `"values":[]`,
		},
		{
			name: "RunUserQueryInput with nil DelegationIds",
			input: RunUserQueryInput{
				DelegationIds: nil,
			},
			mustContain: `"delegationIds":[]`,
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
