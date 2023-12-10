package client

import (
	"context"

	"github.com/cloudquery/plugin-sdk/plugins/source"
	"github.com/cloudquery/plugin-sdk/schema"
	"github.com/cloudquery/plugin-sdk/specs"
	"github.com/rs/zerolog"
	"golang.org/x/oauth2/google"
	directory "google.golang.org/api/admin/directory/v1"
	"google.golang.org/api/option"
)

type Client struct {
	DirectoryService *directory.Service
	CustomerID       string
	logger           zerolog.Logger
}

var _ schema.ClientMeta = (*Client)(nil)

func (c *Client) Logger() *zerolog.Logger {
	return &c.logger
}

func (c *Client) ID() string {
	return "googleworkspace:customer:{" + c.CustomerID + "}"
}

func Configure(ctx context.Context, logger zerolog.Logger, srcSpec specs.Source, options source.Options) (schema.ClientMeta, error) {
	spec := new(Spec)
	if err := srcSpec.UnmarshalSpec(&spec); err != nil {
		return nil, err
	}

	if err := spec.validate(); err != nil {
		return nil, err
	}

	opts := []option.ClientOption{
		// option.WithRequestReason("cloudquery resource fetch"),
		// we disable telemetry to boost performance and be on the safe side with telemetry
		// option.WithTelemetryDisabled(),
	}

	if spec.OAuth != nil {
		tokenSource, err := spec.OAuth.getTokenSource(ctx, google.Endpoint)
		if err != nil {
			return nil, err
		}
		opts = append(opts, option.WithTokenSource(tokenSource))
	}

	svc, err := directory.NewService(ctx, opts...)
	if err != nil {
		return nil, err
	}

	svc.UserAgent = "cloudquery:source-googleworkspace/" + srcSpec.Version

	c := &Client{
		DirectoryService: svc,
		CustomerID:       spec.CustomerID,
		logger: logger.With().
			Str("plugin", "googleworkspace").
			Str("customer_id", spec.CustomerID).
			Logger(),
	}

	return c, nil
}
