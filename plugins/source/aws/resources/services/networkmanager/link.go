package networkmanager

import (
	"context"

	"github.com/apache/arrow/go/v13/arrow"
	"github.com/aws/aws-sdk-go-v2/service/networkmanager"
	"github.com/aws/aws-sdk-go-v2/service/networkmanager/types"
	"github.com/cloudquery/cloudquery/plugins/source/aws/client"
	"github.com/cloudquery/plugin-sdk/v3/schema"
	"github.com/cloudquery/plugin-sdk/v3/transformers"
)

func links() *schema.Table {
	tableName := "aws_networkmanager_links"
	return &schema.Table{
		Name:        tableName,
		Description: `https://docs.aws.amazon.com/networkmanager/latest/APIReference/API_Link.html`,
		Resolver:    fetchLinks,
		Transform:   transformers.TransformWithStruct(&types.Link{}, transformers.WithPrimaryKeys("GlobalNetworkId")),
		Columns: []schema.Column{
			client.DefaultAccountIDColumn(false),
			client.DefaultRegionColumn(true),
			{
				Name:     "arn",
				Type:     arrow.BinaryTypes.String,
				Resolver: schema.PathResolver("LinkArn"),
			},
		},
		Relations: schema.Tables{},
	}
}

func fetchLinks(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan<- any) error {
	cl := meta.(*client.Client)
	svc := cl.Services().Networkmanager
	globalNetwork := parent.Item.(types.GlobalNetwork)
	input := &networkmanager.GetLinksInput{
		GlobalNetworkId: globalNetwork.GlobalNetworkId,
	}
	paginator := networkmanager.NewGetLinksPaginator(svc, input)
	for paginator.HasMorePages() {
		page, err := paginator.NextPage(ctx, func(options *networkmanager.Options) {
			options.Region = cl.Region
		})
		if err != nil {
			return err
		}
		res <- page.Links
	}
	return nil
}
