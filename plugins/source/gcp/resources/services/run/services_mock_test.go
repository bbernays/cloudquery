package run

import (
	"encoding/json"
	"net/http"
	"testing"

	"github.com/cloudquery/plugin-sdk/faker"
	"github.com/cloudquery/plugins/source/gcp/client"
	"github.com/julienschmidt/httprouter"
	"google.golang.org/api/run/v1"
)

func createRunServices(mux *httprouter.Router) error {
	var spec *run.ServiceSpec

	if err := faker.FakeObject(&spec); err != nil {
		return err
	}
	var status *run.ServiceStatus

	if err := faker.FakeObject(&status); err != nil {
		return err
	}

	var meta *run.ObjectMeta

	if err := faker.FakeObject(&meta); err != nil {
		return err
	}
	mux.GET("/v1/projects/testProject/locations/-/services", func(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
		resp := run.ListServicesResponse{
			Items: []*run.Service{{
				Spec:     spec,
				Status:   status,
				Metadata: meta,
			},
			},
		}
		b, err := json.Marshal(resp)
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

func TestServices(t *testing.T) {
	client.MockTestRestHelper(t, Services(), createRunServices, client.TestOptions{})
}
