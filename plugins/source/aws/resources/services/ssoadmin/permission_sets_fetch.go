package ssoadmin

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/service/ssoadmin"
	"github.com/aws/aws-sdk-go-v2/service/ssoadmin/types"
	"github.com/cloudquery/cloudquery/plugins/source/aws/client"
	"github.com/cloudquery/plugin-sdk/schema"
)

func fetchSsoadminPermissionSets(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan<- interface{}) error {
	svc := meta.(*client.Client).Services().Ssoadmin
	obj := parent.Item.(types.InstanceMetadata)
	config := ssoadmin.ListPermissionSetsInput{
		InstanceArn: obj.InstanceArn,
	}
	response, err := svc.ListPermissionSets(ctx, &config)
	if err != nil {
		return err
	}
	for _, i := range response.PermissionSets {
		res <- i
	}
	return nil
}

func getSSOAdminPermissionSet(ctx context.Context, meta schema.ClientMeta, resource *schema.Resource) error {
	svc := meta.(*client.Client).Services().Ssoadmin
	permission_set_arn := resource.Item.(string)
	instance_arn := resource.Parent.Item.(types.InstanceMetadata).InstanceArn
	config := ssoadmin.DescribePermissionSetInput{
		InstanceArn:      instance_arn,
		PermissionSetArn: &permission_set_arn,
	}
	resp, err := svc.DescribePermissionSet(ctx, &config)
	if err != nil {
		return err
	}
	resource.Item = resp.PermissionSet
	return nil

}
