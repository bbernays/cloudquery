package run

import (
	"context"

	"github.com/cloudquery/plugin-sdk/schema"
	"github.com/cloudquery/plugins/source/gcp/client"
	"google.golang.org/api/run/v1"
)

func fetchServices(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan<- any) error {
	c := meta.(*client.Client)
	runClient, err := run.NewService(ctx, c.ClientOptions...)
	if err != nil {
		return err
	}

	call := runClient.Projects.Locations.Services.List("projects/" + c.ProjectId + "/locations/-").Context(ctx)
	output, err := call.Do()
	if err != nil {
		return err
	}
	res <- output.Items

	return nil
}
