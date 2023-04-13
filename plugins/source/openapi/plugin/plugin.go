package plugin

import (
	"github.com/cloudquery/cloudquery/plugins/source/openapi/client"
	"github.com/cloudquery/cloudquery/plugins/source/openapi/resources"
	"github.com/cloudquery/plugin-sdk/plugins/source"
	"github.com/cloudquery/plugin-sdk/schema"
)

var (
	Version = "development"
)

func Plugin() *source.Plugin {
	return source.NewPlugin(
		"openapi",
		Version,
		schema.Tables{
			resources.Services(),
		},
		client.New,
	)
}
