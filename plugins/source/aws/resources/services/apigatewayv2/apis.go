// Code generated by codegen; DO NOT EDIT.

package apigatewayv2

import (
	"github.com/cloudquery/cloudquery/plugins/source/aws/client"
	"github.com/cloudquery/plugin-sdk/schema"
)

func Apis() *schema.Table {
	return &schema.Table{
		Name:        "aws_apigatewayv2_apis",
		Description: "https://docs.aws.amazon.com/apigateway/latest/api/API_Api.html",
		Resolver:    fetchApigatewayv2Apis,
		Multiplex:   client.ServiceAccountRegionMultiplexer("apigateway"),
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
				Name:     "arn",
				Type:     schema.TypeString,
				Resolver: resolveApiArn(),
			},
			{
				Name:     "id",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("ApiId"),
			},
			{
				Name:     "name",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("Name"),
			},
			{
				Name:     "protocol_type",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("ProtocolType"),
			},
			{
				Name:     "route_selection_expression",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("RouteSelectionExpression"),
			},
			{
				Name:     "api_endpoint",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("ApiEndpoint"),
			},
			{
				Name:     "api_gateway_managed",
				Type:     schema.TypeBool,
				Resolver: schema.PathResolver("ApiGatewayManaged"),
			},
			{
				Name:     "api_key_selection_expression",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("ApiKeySelectionExpression"),
			},
			{
				Name:     "cors_configuration",
				Type:     schema.TypeJSON,
				Resolver: schema.PathResolver("CorsConfiguration"),
			},
			{
				Name:     "created_date",
				Type:     schema.TypeTimestamp,
				Resolver: schema.PathResolver("CreatedDate"),
			},
			{
				Name:     "description",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("Description"),
			},
			{
				Name:     "disable_execute_api_endpoint",
				Type:     schema.TypeBool,
				Resolver: schema.PathResolver("DisableExecuteApiEndpoint"),
			},
			{
				Name:     "disable_schema_validation",
				Type:     schema.TypeBool,
				Resolver: schema.PathResolver("DisableSchemaValidation"),
			},
			{
				Name:     "import_info",
				Type:     schema.TypeStringArray,
				Resolver: schema.PathResolver("ImportInfo"),
			},
			{
				Name:     "tags",
				Type:     schema.TypeJSON,
				Resolver: schema.PathResolver("Tags"),
			},
			{
				Name:     "version",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("Version"),
			},
			{
				Name:     "warnings",
				Type:     schema.TypeStringArray,
				Resolver: schema.PathResolver("Warnings"),
			},
		},

		Relations: []*schema.Table{
			ApiAuthorizers(),
			ApiDeployments(),
			ApiIntegrations(),
			ApiModels(),
			ApiRoutes(),
			ApiStages(),
		},
	}
}
