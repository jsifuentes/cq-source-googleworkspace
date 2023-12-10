package resources

import (
	"context"

	"github.com/cloudquery/plugin-sdk/schema"
	"github.com/cloudquery/plugin-sdk/transformers"
	"github.com/jsifuentes/cq-source-googleworkspace/client"
	directory "google.golang.org/api/admin/directory/v1"
)

func FeaturesTable() *schema.Table {
	return &schema.Table{
		Name:        "googleworkspace_resources_features",
		Description: "Google Workspace Resources Features",
		Transform:   transformers.TransformWithStruct(&directory.Feature{}, transformers.WithPrimaryKeys("Name")),
		Columns: []schema.Column{
			client.CustomerIDColumn,
		},
		Resolver: fetchFeatures,
	}
}

func fetchFeatures(ctx context.Context, meta schema.ClientMeta, _ *schema.Resource, res chan<- any) error {
	c := meta.(*client.Client)

	return c.DirectoryService.Resources.Features.List(c.CustomerID).Pages(ctx, func(features *directory.Features) error {
		for _, f := range features.Features {
			res <- f
		}
		return nil
	})
}
