package ssm

import (
	"github.com/aws/aws-sdk-go-v2/service/ssm/types"
	"github.com/cloudquery/cloudquery/plugins/source/aws/client"
	"github.com/cloudquery/plugin-sdk/schema"
	"github.com/cloudquery/plugin-sdk/transformers"
)

func InstanceComplianceItems() *schema.Table {
	return &schema.Table{
		Name:        "aws_ssm_instance_compliance_items",
		Description: `https://docs.aws.amazon.com/systems-manager/latest/APIReference/API_ComplianceItem.html`,
		Resolver:    fetchSsmInstanceComplianceItems,
		Transform:   transformers.TransformWithStruct(&types.ComplianceItem{}),
		Multiplex:   client.ServiceAccountRegionMultiplexer("ssm"),
		Columns: []schema.Column{
			client.AccountPKColumn(false),
			client.RegionPKColumn(false),
			{
				Name:     "id",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("Id"),
				CreationOptions: schema.ColumnCreationOptions{
					PrimaryKey: true,
				},
			},
			{
				Name:     "instance_arn",
				Type:     schema.TypeString,
				Resolver: resolveInstanceComplianceItemInstanceARN,
				CreationOptions: schema.ColumnCreationOptions{
					PrimaryKey: true,
				},
			},
		},
	}
}
