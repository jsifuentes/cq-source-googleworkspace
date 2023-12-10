package client

import (
	"context"

	"github.com/cloudquery/plugin-sdk/schema"
)

var (
	CustomerIDColumn = schema.Column{
		Name: "customer_id",
		Type: schema.TypeString,
		Resolver: func(_ context.Context, meta schema.ClientMeta, r *schema.Resource, c schema.Column) error {
			return r.Set(c.Name, meta.(*Client).CustomerID)
		},
	}
)
