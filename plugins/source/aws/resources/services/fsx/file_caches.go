package fsx

import (
	"github.com/aws/aws-sdk-go-v2/service/fsx/types"
	"github.com/cloudquery/cloudquery/plugins/source/aws/client"
	"github.com/cloudquery/plugin-sdk/schema"
	"github.com/cloudquery/plugin-sdk/transformers"
)

func FileCaches() *schema.Table {
	return &schema.Table{
		Name:        "aws_fsx_file_caches",
		Description: `https://docs.aws.amazon.com/fsx/latest/APIReference/API_FileCache.html`,
		Resolver:    fetchFsxFileCaches,
		Transform:   transformers.TransformWithStruct(&types.FileCache{}),
		Multiplex:   client.ServiceAccountRegionMultiplexer("fsx"),
		Columns: []schema.Column{
			client.AccountPKColumn(false),
			client.RegionPKColumn(false),
			{
				Name:     "arn",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("ResourceARN"),
				CreationOptions: schema.ColumnCreationOptions{
					PrimaryKey: true,
				},
			},
			{
				Name:     "tags",
				Type:     schema.TypeJSON,
				Resolver: resolveFileCacheTags,
			},
		},
	}
}
