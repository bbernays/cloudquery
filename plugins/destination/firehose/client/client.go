package client

import (
	"context"
	"fmt"

	"github.com/aws/aws-sdk-go-v2/aws/arn"
	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/service/firehose"

	"github.com/cloudquery/plugin-pb-go/specs"
	"github.com/cloudquery/plugin-sdk/v3/plugins/destination"
	"github.com/rs/zerolog"
)

type Client struct {
	destination.UnimplementedManagedWriter
	logger         zerolog.Logger
	spec           specs.Destination
	pluginSpec     Spec
	metrics        destination.Metrics
	firehoseClient *firehose.Client
}

var _ destination.Client = (*Client)(nil)

func New(ctx context.Context, logger zerolog.Logger, spec specs.Destination) (destination.Client, error) {
	if spec.WriteMode != specs.WriteModeAppend {
		return nil, fmt.Errorf("destination only supports append mode")
	}
	c := &Client{
		logger: logger.With().Str("module", "firehose").Logger(),
		spec:   spec,
	}

	if err := spec.UnmarshalSpec(&c.pluginSpec); err != nil {
		return nil, fmt.Errorf("failed to unmarshal firehose spec: %w", err)
	}
	if err := c.pluginSpec.Validate(); err != nil {
		return nil, err
	}
	c.pluginSpec.SetDefaults()
	parsedARN, err := arn.Parse(c.pluginSpec.StreamARN)
	if err != nil {
		return nil, fmt.Errorf("failed to parse firehose stream ARN: %w", err)
	}
	cfg, err := config.LoadDefaultConfig(ctx, config.WithRegion(parsedARN.Region))
	if err != nil {
		return nil, fmt.Errorf("unable to load AWS SDK config: %w", err)
	}

	c.firehoseClient = firehose.NewFromConfig(cfg)

	return c, nil
}

func (*Client) Close(ctx context.Context) error {
	return nil
}
