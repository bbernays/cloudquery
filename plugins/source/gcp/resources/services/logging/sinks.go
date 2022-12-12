// Code generated by codegen; DO NOT EDIT.

package logging

import (
	"context"
	"google.golang.org/api/iterator"

	pb "google.golang.org/genproto/googleapis/logging/v2"

	"github.com/cloudquery/plugin-sdk/schema"
	"github.com/cloudquery/plugins/source/gcp/client"

	"cloud.google.com/go/logging/apiv2"
)

func Sinks() *schema.Table {
	return &schema.Table{
		Name:      "gcp_logging_sinks",
		Resolver:  fetchSinks,
		Multiplex: client.ProjectMultiplex,
		Columns: []schema.Column{
			{
				Name:     "project_id",
				Type:     schema.TypeString,
				Resolver: client.ResolveProject,
			},
			{
				Name:     "name",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("Name"),
				CreationOptions: schema.ColumnCreationOptions{
					PrimaryKey: true,
				},
			},
			{
				Name:     "destination",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("Destination"),
			},
			{
				Name:     "filter",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("Filter"),
			},
			{
				Name:     "description",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("Description"),
			},
			{
				Name:     "disabled",
				Type:     schema.TypeBool,
				Resolver: schema.PathResolver("Disabled"),
			},
			{
				Name:     "exclusions",
				Type:     schema.TypeJSON,
				Resolver: schema.PathResolver("Exclusions"),
			},
			{
				Name:     "output_version_format",
				Type:     schema.TypeString,
				Resolver: client.ResolveProtoEnum("OutputVersionFormat"),
			},
			{
				Name:     "writer_identity",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("WriterIdentity"),
			},
			{
				Name:     "include_children",
				Type:     schema.TypeBool,
				Resolver: schema.PathResolver("IncludeChildren"),
			},
			{
				Name:     "create_time",
				Type:     schema.TypeTimestamp,
				Resolver: client.ResolveProtoTimestamp("CreateTime"),
			},
			{
				Name:     "update_time",
				Type:     schema.TypeTimestamp,
				Resolver: client.ResolveProtoTimestamp("UpdateTime"),
			},
		},
	}
}

func fetchSinks(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan<- interface{}) error {
	c := meta.(*client.Client)
	req := &pb.ListSinksRequest{
		Parent: "projects/" + c.ProjectId,
	}
	gcpClient, err := logging.NewConfigClient(ctx, c.ClientOptions...)
	if err != nil {
		return err
	}
	it := gcpClient.ListSinks(ctx, req, c.CallOptions...)
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
