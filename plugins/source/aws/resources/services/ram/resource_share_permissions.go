// Code generated by codegen; DO NOT EDIT.

package ram

import (
	"github.com/cloudquery/cloudquery/plugins/source/aws/client"
	"github.com/cloudquery/plugin-sdk/schema"
)

func ResourceSharePermissions() *schema.Table {
	return &schema.Table{
		Name:        "aws_ram_resource_share_permissions",
		Description: `https://docs.aws.amazon.com/ram/latest/APIReference/API_ResourceSharePermissionSummary.html`,
		Resolver:    fetchRamResourceSharePermissions,
		Multiplex:   client.ServiceAccountRegionMultiplexer("ram"),
		Columns: []schema.Column{
			{
				Name:     "account_id",
				Type:     schema.TypeString,
				Resolver: client.ResolveAWSAccount,
			},
			{
				Name:     "region",
				Type:     schema.TypeString,
				Resolver: client.ResolveAWSRegion,
			},
			{
				Name:     "permission",
				Type:     schema.TypeJSON,
				Resolver: resolveResourceSharePermissionDetailPermission,
			},
			{
				Name:     "arn",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("Arn"),
				CreationOptions: schema.ColumnCreationOptions{
					PrimaryKey: true,
				},
			},
			{
				Name:     "creation_time",
				Type:     schema.TypeTimestamp,
				Resolver: schema.PathResolver("CreationTime"),
			},
			{
				Name:     "default_version",
				Type:     schema.TypeBool,
				Resolver: schema.PathResolver("DefaultVersion"),
			},
			{
				Name:     "is_resource_type_default",
				Type:     schema.TypeBool,
				Resolver: schema.PathResolver("IsResourceTypeDefault"),
			},
			{
				Name:     "last_updated_time",
				Type:     schema.TypeTimestamp,
				Resolver: schema.PathResolver("LastUpdatedTime"),
			},
			{
				Name:     "name",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("Name"),
			},
			{
				Name:     "resource_type",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("ResourceType"),
			},
			{
				Name:     "status",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("Status"),
			},
			{
				Name:     "version",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("Version"),
			},
		},
	}
}
