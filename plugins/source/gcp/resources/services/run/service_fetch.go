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

	locations, err := getAllRunLocations(ctx, c, runClient)
	if err != nil {
		return err
	}
	for _, l := range locations {
		call := runClient.Projects.Locations.Services.List("projects/" + c.ProjectId + "/locations/" + l.LocationId).Context(ctx)
		output, err := call.Do()
		if err != nil {
			return err
		}
		res <- output.Items
	}
	return nil
}

func getAllRunLocations(ctx context.Context, c *client.Client, runClient *run.APIService) ([]*run.Location, error) {
	var locations []*run.Location
	call := runClient.Projects.Locations.List("projects/" + c.ProjectId).Context(ctx)
	for {
		output, err := call.Do()
		if err != nil {
			return nil, err
		}
		locations = append(locations, output.Locations...)
		if output.NextPageToken == "" {
			break
		}
		call.PageToken(output.NextPageToken)
	}
	return locations, nil
}
