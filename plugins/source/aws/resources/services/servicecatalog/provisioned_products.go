// Code generated by codegen; DO NOT EDIT.

package servicecatalog

import (
	"github.com/cloudquery/cloudquery/plugins/source/aws/client"
	"github.com/cloudquery/plugin-sdk/schema"
)

func ProvisionedProducts() *schema.Table {
	return &schema.Table{
		Name:        "aws_servicecatalog_provisioned_products",
		Description: "https://docs.aws.amazon.com/servicecatalog/latest/dg/API_ProvisionedProductAttribute.html",
		Resolver:    fetchServicecatalogProvisionedProducts,
		Multiplex:   client.ServiceAccountRegionMultiplexer("servicecatalog"),
		Columns: []schema.Column{
			{
				Name:     "account_id",
				Type:     schema.TypeString,
				Resolver: client.ResolveAWSAccount,
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
				Name:     "tags",
				Type:     schema.TypeJSON,
				Resolver: resolveProvisionedProductTags,
			},
			{
				Name:     "created_time",
				Type:     schema.TypeTimestamp,
				Resolver: schema.PathResolver("CreatedTime"),
			},
			{
				Name:     "id",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("Id"),
			},
			{
				Name:     "idempotency_token",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("IdempotencyToken"),
			},
			{
				Name:     "last_provisioning_record_id",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("LastProvisioningRecordId"),
			},
			{
				Name:     "last_record_id",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("LastRecordId"),
			},
			{
				Name:     "last_successful_provisioning_record_id",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("LastSuccessfulProvisioningRecordId"),
			},
			{
				Name:     "name",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("Name"),
			},
			{
				Name:     "physical_id",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("PhysicalId"),
			},
			{
				Name:     "product_id",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("ProductId"),
			},
			{
				Name:     "product_name",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("ProductName"),
			},
			{
				Name:     "provisioning_artifact_id",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("ProvisioningArtifactId"),
			},
			{
				Name:     "provisioning_artifact_name",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("ProvisioningArtifactName"),
			},
			{
				Name:     "status",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("Status"),
			},
			{
				Name:     "status_message",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("StatusMessage"),
			},
			{
				Name:     "type",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("Type"),
			},
			{
				Name:     "user_arn",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("UserArn"),
			},
			{
				Name:     "user_arn_session",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("UserArnSession"),
			},
		},
	}
}
