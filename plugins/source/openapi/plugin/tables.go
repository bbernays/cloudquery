package plugin

import (
	"context"

	"github.com/cloudquery/cloudquery/plugins/source/openapi/client"
	"github.com/cloudquery/plugin-sdk/schema"
)

func getDynamicTables(ctx context.Context, c schema.ClientMeta) (schema.Tables, error) {
	cl := c.(*client.Client)
	return cl.Tables, nil
}
