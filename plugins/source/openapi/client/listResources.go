package client

import (
	"context"

	"github.com/cloudquery/plugin-sdk/schema"
)

func (c *Client) listTables(ctx context.Context) (schema.Tables, error) {
	var tables schema.Tables
	return tables, nil
}
