package client

import (
	"context"

	"github.com/apache/arrow/go/v14/arrow"
	"github.com/cloudquery/plugin-sdk/v4/schema"
)

var (
	CustomerIDColumn = schema.Column{
		Name: "customer_id",
		Type: arrow.BinaryTypes.String,
		Resolver: func(ctx context.Context, meta schema.ClientMeta, resource *schema.Resource, c schema.Column) error {
			return resource.Set(c.Name, meta.(*Client).Spec.CustomerID)
		},
	}
)
