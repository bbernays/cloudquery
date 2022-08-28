// Code generated by codegen; DO NOT EDIT.

package compute

import (
	"context"
	"github.com/cloudquery/plugin-sdk/schema"
	"github.com/cloudquery/plugins/source/gcp/client"
	"github.com/pkg/errors"
)

func Projects() *schema.Table {
	return &schema.Table{
		Name:      "gcp_compute_projects",
		Resolver:  fetchProjects,
		Multiplex: client.ProjectMultiplex,
		Columns: []schema.Column{
			{
				Name:     "project_id",
				Type:     schema.TypeString,
				Resolver: client.ResolveProject,
			},
			{
				Name:     "common_instance_metadata",
				Type:     schema.TypeJSON,
				Resolver: schema.PathResolver("CommonInstanceMetadata"),
			},
			{
				Name:     "creation_timestamp",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("CreationTimestamp"),
			},
			{
				Name:     "default_network_tier",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("DefaultNetworkTier"),
			},
			{
				Name:     "default_service_account",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("DefaultServiceAccount"),
			},
			{
				Name:     "description",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("Description"),
			},
			{
				Name:     "enabled_features",
				Type:     schema.TypeStringArray,
				Resolver: schema.PathResolver("EnabledFeatures"),
			},
			{
				Name:     "id",
				Type:     schema.TypeInt,
				Resolver: schema.PathResolver("Id"),
			},
			{
				Name:     "kind",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("Kind"),
			},
			{
				Name:     "name",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("Name"),
			},
			{
				Name:     "quotas",
				Type:     schema.TypeJSON,
				Resolver: schema.PathResolver("Quotas"),
			},
			{
				Name: "self_link",
				Type: schema.TypeString,
				CreationOptions: schema.ColumnCreationOptions{
					PrimaryKey: true,
				},
			},
			{
				Name:     "usage_export_location",
				Type:     schema.TypeJSON,
				Resolver: schema.PathResolver("UsageExportLocation"),
			},
			{
				Name:     "xpn_project_status",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("XpnProjectStatus"),
			},
		},
	}
}

func fetchProjects(ctx context.Context, meta schema.ClientMeta, _ *schema.Resource, res chan<- interface{}) error {
	c := meta.(*client.Client)
	output, err := c.Services.Compute.Projects.Get(c.ProjectId).Do()
	if err != nil {
		return errors.WithStack(err)
	}
	res <- output

	return nil
}
