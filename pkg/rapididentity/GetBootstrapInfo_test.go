package rapididentity

import (
	"fmt"
	"net/http"
	"testing"
)

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

	output, err := client.GetBootstrapInfo()
	if err != nil {
		t.Errorf("got error %s, want none", err)
	}

	got := output.TenantId
	want := "1234"
	if got != want {
		t.Errorf("got %s. want %s", got, want)
	}
}
