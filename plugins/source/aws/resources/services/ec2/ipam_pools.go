package ec2

import (
	"context"

	sdkTypes "github.com/cloudquery/plugin-sdk/v4/types"

	"github.com/apache/arrow/go/v13/arrow"
	"github.com/aws/aws-sdk-go-v2/service/ec2"
	"github.com/aws/aws-sdk-go-v2/service/ec2/types"
	"github.com/cloudquery/cloudquery/plugins/source/aws/client"
	"github.com/cloudquery/plugin-sdk/v4/schema"
	"github.com/cloudquery/plugin-sdk/v4/transformers"
)

func IpamPools() *schema.Table {
	tableName := "aws_ec2_ipam_pools"
	return &schema.Table{
		Name:        tableName,
		Description: `https://docs.aws.amazon.com/AWSEC2/latest/APIReference/API_IpamPool.html`,
		Resolver:    fetchIPAMPools,
		Multiplex:   client.ServiceAccountRegionMultiplexer(tableName, "ec2"),
		Transform:   transformers.TransformWithStruct(&types.IpamPool{}),
		Columns: []schema.Column{
			{
				Name:       "request_account_id",
				Type:       arrow.BinaryTypes.String,
				Resolver:   client.ResolveAWSAccount,
				PrimaryKey: true,
			},
			{
				Name:       "request_region",
				Type:       arrow.BinaryTypes.String,
				Resolver:   client.ResolveAWSRegion,
				PrimaryKey: true,
			},
			{
				Name:       "arn",
				Type:       arrow.BinaryTypes.String,
				Resolver:   schema.PathResolver("IpamPoolArn"),
				PrimaryKey: true,
			},

			{
				Name:     "tags",
				Type:     sdkTypes.ExtensionTypes.JSON,
				Resolver: client.ResolveTags,
			},
		},
		Relations: schema.Tables{
			// poolAllocations(),
			// poolCidrs(),
		},
	}
}

func fetchIPAMPools(ctx context.Context, meta schema.ClientMeta, _ *schema.Resource, res chan<- any) error {
	cl := meta.(*client.Client)
	svc := cl.Services().Ec2
	paginator := ec2.NewDescribeIpamPoolsPaginator(svc, &ec2.DescribeIpamPoolsInput{})
	for paginator.HasMorePages() {
		output, err := paginator.NextPage(ctx, func(options *ec2.Options) {
			options.Region = cl.Region
		})
		if err != nil {
			return err
		}
		res <- output.IpamPools
	}

	return nil
}
