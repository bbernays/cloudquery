package ec2

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/ec2"
	"github.com/aws/aws-sdk-go-v2/service/ec2/types"
	"github.com/cloudquery/cloudquery/plugins/source/aws/client"
	"github.com/cloudquery/plugin-sdk/schema"
)

func AwsRegions() *schema.Table {
	return &schema.Table{
		Name:        "aws_regions",
		Description: "Describes a Region.",
		Resolver:    fetchRegions,
		Multiplex:   client.AccountMultiplex,
		Columns: []schema.Column{
			{
				Name:        "account_id",
				Description: "The AWS Account ID of the resource.",
				Type:        schema.TypeString,
				Resolver:    client.ResolveAWSAccount,
			},
			{
				Name:        "enabled",
				Description: "Defines if region is enabled stated or not.",
				Type:        schema.TypeBool,
				Resolver:    resolveRegionEnabled,
			},
			{
				Name:        "endpoint",
				Description: "The Region service endpoint.",
				Type:        schema.TypeString,
			},
			{
				Name:        "opt_in_status",
				Description: "The Region opt-in status",
				Type:        schema.TypeString,
			},
			{
				Name:        "region",
				Description: "The name of the Region.",
				Type:        schema.TypeString,
				Resolver:    schema.PathResolver("RegionName"),
			},
			{
				Name:        "partition",
				Description: "AWS partition",
				Type:        schema.TypeString,
				Resolver:    resolveRegionPartition,
			},
		},
	}
}

// ====================================================================================================================
//                                               Table Resolver Functions
// ====================================================================================================================

func fetchRegions(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan<- interface{}) error {
	c := meta.(*client.Client)
	output, err := c.Services().EC2Instances.DescribeRegions(ctx, &ec2.DescribeRegionsInput{AllRegions: aws.Bool(true)})
	if err != nil {
		return err
	}
	res <- output.Regions
	return nil
}
func resolveRegionEnabled(ctx context.Context, meta schema.ClientMeta, resource *schema.Resource, c schema.Column) error {
	region := resource.Item.(types.Region)
	switch *region.OptInStatus {
	case "opt-in-not-required", "opted-in":
		return resource.Set(c.Name, true)
	case "not-opted-in":
		return resource.Set(c.Name, false)
	}
	return nil
}
func resolveRegionPartition(ctx context.Context, meta schema.ClientMeta, resource *schema.Resource, c schema.Column) error {
	cl := meta.(*client.Client)
	return resource.Set(c.Name, cl.Partition)
}
