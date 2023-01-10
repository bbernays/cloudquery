package bigtableadmin

import (
	pb "cloud.google.com/go/bigtable"
	"github.com/cloudquery/plugin-sdk/schema"
	"github.com/cloudquery/plugin-sdk/transformers"
	"github.com/cloudquery/plugins/source/gcp/client"
)

func Backups() *schema.Table {
	return &schema.Table{
		Name:        "gcp_bigtableadmin_backups",
		Description: `https://cloud.google.com/bigtable/docs/reference/admin/rest/v2/projects.instances.clusters.backups#Backup`,
		Resolver:    fetchBackups,
		Multiplex:   client.ProjectMultiplexEnabledServices("bigtableadmin.googleapis.com"),
		Transform:   transformers.TransformWithStruct(&pb.BackupInfo{}, client.Options()...),
		Columns: []schema.Column{
			{
				Name:     "project_id",
				Type:     schema.TypeString,
				Resolver: client.ResolveProject,
				CreationOptions: schema.ColumnCreationOptions{
					PrimaryKey: true,
				},
			},
			{
				Name:     "name",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("Name"),
				CreationOptions: schema.ColumnCreationOptions{
					PrimaryKey: true,
				},
			},
		},
	}
}
