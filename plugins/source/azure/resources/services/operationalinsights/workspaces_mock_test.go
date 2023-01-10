// Code generated by codegen; DO NOT EDIT.
package operationalinsights

import (
	"encoding/json"
	"net/http"

	"github.com/cloudquery/cloudquery/plugins/source/azure/client"
	"testing"

	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/operationalinsights/armoperationalinsights"
	"github.com/cloudquery/plugin-sdk/faker"
	"github.com/gorilla/mux"
)

func createWorkspaces(router *mux.Router) error {
	var item armoperationalinsights.WorkspacesClientListResponse
	if err := faker.FakeObject(&item); err != nil {
		return err
	}

	router.HandleFunc("/subscriptions/{subscriptionId}/providers/Microsoft.OperationalInsights/workspaces", func(w http.ResponseWriter, r *http.Request) {
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

func TestWorkspaces(t *testing.T) {
	client.MockTestHelper(t, Workspaces(), createWorkspaces)
}
