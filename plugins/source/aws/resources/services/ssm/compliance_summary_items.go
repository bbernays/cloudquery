package ssm

import (
	"github.com/aws/aws-sdk-go-v2/service/ssm/types"
	"github.com/cloudquery/cloudquery/plugins/source/aws/client"
	"github.com/cloudquery/plugin-sdk/schema"
	"github.com/cloudquery/plugin-sdk/transformers"
)

func ComplianceSummaryItems() *schema.Table {
	return &schema.Table{
		Name:        "aws_ssm_compliance_summary_items",
		Description: `https://docs.aws.amazon.com/systems-manager/latest/APIReference/API_ComplianceSummaryItem.html`,
		Resolver:    fetchSsmComplianceSummaryItems,
		Transform:   transformers.TransformWithStruct(&types.ComplianceSummaryItem{}),
		Multiplex:   client.ServiceAccountRegionMultiplexer("ssm"),
		Columns: []schema.Column{
			client.AccountPKColumn(true),
			client.RegionPKColumn(true),
			{
				Name: "compliance_type",
				Type: schema.TypeString,
				CreationOptions: schema.ColumnCreationOptions{
					PrimaryKey: true,
				},
			},
		},
	}
}
