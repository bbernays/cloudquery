package client

import (
	"context"
	"fmt"

	"github.com/cloudquery/plugin-sdk/plugins/source"
	"github.com/cloudquery/plugin-sdk/schema"
	"github.com/cloudquery/plugin-sdk/specs"
	"github.com/rs/zerolog"
)

type Client struct {
	logger     zerolog.Logger
	spec       specs.Source
	pluginSpec Spec
	Tables     schema.Tables
}

var _ schema.ClientMeta = (*Client)(nil)

func (*Client) ID() string {
	return "source-openapi"
}

func Configure(ctx context.Context, logger zerolog.Logger, spec specs.Source, _ source.Options) (schema.ClientMeta, error) {
	c := &Client{
		logger: logger.With().Str("module", "openapi").Logger(),
	}
	var pluginSpec Spec
	c.spec = spec
	if err := spec.UnmarshalSpec(&pluginSpec); err != nil {
		return nil, fmt.Errorf("failed to unmarshal openapi spec: %w", err)
	}
	pluginSpec.SetDefaults()
	err := pluginSpec.Validate()
	if err != nil {
		return nil, err
	}
	c.pluginSpec = pluginSpec
	// get the full OpenAPI spec

	err = c.generateClient(ctx)
	if err != nil {
		return nil, fmt.Errorf("failed to get OpenAPI spec: %w", err)
	}

	c.Tables, err = c.generateTables(ctx)
	if err != nil {
		return nil, fmt.Errorf("failed to list tables: %w", err)
	}
	c.Tables, err = c.Tables.FilterDfs(spec.Tables, spec.SkipTables, spec.SkipDependentTables)
	if err != nil {
		return nil, fmt.Errorf("failed to apply config to tables: %w", err)
	}

	return c, nil
}
