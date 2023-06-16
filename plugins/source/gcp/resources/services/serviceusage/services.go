package serviceusage

import (
	"context"

	serviceusage "cloud.google.com/go/serviceusage/apiv1"
	pb "cloud.google.com/go/serviceusage/apiv1/serviceusagepb"
	"github.com/apache/arrow/go/v13/arrow"
	"github.com/cloudquery/plugin-sdk/v3/schema"
	"github.com/cloudquery/plugin-sdk/v3/transformers"
	"github.com/cloudquery/plugins/source/gcp/client"
	"google.golang.org/api/iterator"
)

func Services() *schema.Table {
	return &schema.Table{
		Name:                "gcp_serviceusage_services",
		Description:         `https://cloud.google.com/service-usage/docs/reference/rest/v1/services#Service`,
		Resolver:            listServices,
		PreResourceResolver: getService,
		Multiplex:           client.ProjectMultiplexEnabledServices("serviceusage.googleapis.com"),
		Transform:           client.TransformWithStruct(&pb.Service{}, transformers.WithPrimaryKeys("Name")),
		Columns: []schema.Column{
			{
				Name:     "project_id",
				Type:     arrow.BinaryTypes.String,
				Resolver: client.ResolveProject,
			},
		},
	}
}

func listServices(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan<- any) error {
	c := meta.(*client.Client)
	if len(c.EnabledServices) > 0 {
		for _, svc := range c.EnabledServices[c.ProjectId] {
			res <- svc.(*pb.Service)
		}
		return nil
	}
	req := &pb.ListServicesRequest{
		Parent:   "projects/" + c.ProjectId,
		PageSize: 200,
		Filter:   "state:ENABLED",
	}
	gcpClient, err := serviceusage.NewClient(ctx, c.ClientOptions...)
	if err != nil {
		return err
	}
	it := gcpClient.ListServices(ctx, req, c.CallOptions...)
	for {
		resp, err := it.Next()
		if err == iterator.Done {
			break
		}
		if err != nil {
			return err
		}

		res <- resp
	}
	return nil
}

func getService(ctx context.Context, meta schema.ClientMeta, r *schema.Resource) error {
	c := meta.(*client.Client)
	gcpClient, err := serviceusage.NewClient(ctx, c.ClientOptions...)
	if err != nil {
		return err
	}
	item, err := gcpClient.GetService(ctx, &pb.GetServiceRequest{Name: r.Item.(*pb.Service).Name}, c.CallOptions...)
	if err != nil {
		return err
	}
	r.SetItem(item)
	return nil
}
