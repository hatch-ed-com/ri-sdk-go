package rapididentity

import (
	"context"
	"fmt"
	"net/http"
	"testing"
)

func TestGetLdapAttributes(t *testing.T) {
	t.Parallel()
	client, mux := setup(t)
	mux.HandleFunc(baseUrlPath+"/admin/ldap/schema/attribute", func(w http.ResponseWriter, r *http.Request) {
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
