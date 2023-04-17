package client

import (
	"context"
	"net/url"

	"github.com/cloudquery/plugin-sdk/schema"
	"github.com/deepmap/oapi-codegen/pkg/codegen"
	"github.com/getkin/kin-openapi/openapi3"
)

func (c *Client) generateTables(ctx context.Context) (schema.Tables, error) {
	var tables schema.Tables
	return tables, nil
}

func (c *Client) generateClient(ctx context.Context) error {
	url, err := url.Parse(c.pluginSpec.OpenAPIURL)
	c.logger.Info().Msgf("err: %v", err)
	s, err := openapi3.NewLoader().LoadFromURI(url)
	c.logger.Info().Msgf("err: %v", err)
	wst, err := codegen.Generate(s, codegen.Configuration{})
	c.logger.Info().Msgf("wst: %v", wst)

	return err
}
