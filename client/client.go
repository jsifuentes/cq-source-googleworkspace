package client

import (
	"context"
	"fmt"

	"github.com/rs/zerolog"
	"golang.org/x/oauth2/google"
	directory "google.golang.org/api/admin/directory/v1"
	"google.golang.org/api/option"
)

type Client struct {
	logger           zerolog.Logger
	Spec             Spec
	DirectoryService *directory.Service
}

func (c *Client) ID() string {
	// TODO: Change to either your plugin name or a unique dynamic identifier
	return "googleworkspace:customer:{" + c.Spec.CustomerID + "}"
}

func (c *Client) Logger() *zerolog.Logger {
	return &c.logger
}

func New(ctx context.Context, logger zerolog.Logger, s *Spec) (Client, error) {
	c := Client{
		logger: logger.With().
			Str("plugin", "googleworkspace").
			Str("customer_id", s.CustomerID).
			Logger(),
		Spec: *s,
	}

	if err := s.validate(); err != nil {
		return c, fmt.Errorf("invalid spec: %w", err)
	}

	opts := []option.ClientOption{}

	if c.Spec.OAuth != nil {
		tokenSource, err := c.Spec.OAuth.getTokenSource(ctx, logger, google.Endpoint)
		if err != nil {
			return c, err
		}
		opts = append(opts, option.WithTokenSource(tokenSource))
	} else if c.Spec.ServiceAccount != nil {
		tokenSource, err := c.Spec.ServiceAccount.getTokenSource(ctx)
		if err != nil {
			return c, err
		}
		opts = append(opts, option.WithTokenSource(tokenSource))
	}

	service, err := directory.NewService(ctx, opts...)
	if err != nil {
		return c, err
	}

	service.UserAgent = "cloudquery:source-googleworkspace"
	c.DirectoryService = service

	return c, nil
}
