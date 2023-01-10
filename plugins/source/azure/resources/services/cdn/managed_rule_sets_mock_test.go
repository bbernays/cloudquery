// Code generated by codegen; DO NOT EDIT.
package cdn

import (
	"encoding/json"
	"net/http"

	"github.com/cloudquery/cloudquery/plugins/source/azure/client"
	"testing"

	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/cdn/armcdn"
	"github.com/cloudquery/plugin-sdk/faker"
	"github.com/gorilla/mux"
)

func createManagedRuleSets(router *mux.Router) error {
	var item armcdn.ManagedRuleSetsClientListResponse
	if err := faker.FakeObject(&item); err != nil {
		return err
	}

	emptyStr := ""
	item.NextLink = &emptyStr

	router.HandleFunc("/subscriptions/{subscriptionId}/providers/Microsoft.Cdn/cdnWebApplicationFirewallManagedRuleSets", func(w http.ResponseWriter, r *http.Request) {
		b, err := json.Marshal(&item)
		if err != nil {
			http.Error(w, "unable to marshal request: "+err.Error(), http.StatusBadRequest)
			return
		}
		if _, err := w.Write(b); err != nil {
			http.Error(w, "failed to write", http.StatusBadRequest)
			return
		}
	})

	return nil
}

func TestManagedRuleSets(t *testing.T) {
	client.MockTestHelper(t, ManagedRuleSets(), createManagedRuleSets)
}
