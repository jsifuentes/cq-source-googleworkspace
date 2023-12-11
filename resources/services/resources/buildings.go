package resources

import (
	"context"

	"github.com/cloudquery/plugin-sdk/v4/schema"
	"github.com/cloudquery/plugin-sdk/v4/transformers"
	"github.com/jsifuentes/cq-source-googleworkspace/client"
	directory "google.golang.org/api/admin/directory/v1"
)

func BuildingsTable() *schema.Table {
	return &schema.Table{
		Name:        "googleworkspace_resources_buildings",
		Description: "Google Workspace Resources Buildings",
		Transform:   transformers.TransformWithStruct(&directory.Building{}, transformers.WithPrimaryKeys("BuildingId")),
		Columns: []schema.Column{
			client.CustomerIDColumn,
		},
		Resolver: fetchBuildings,
	}
}

func fetchBuildings(ctx context.Context, meta schema.ClientMeta, _ *schema.Resource, res chan<- any) error {
	c := meta.(*client.Client)
	return c.DirectoryService.Resources.Buildings.List(c.Spec.CustomerID).Pages(ctx, func(buildings *directory.Buildings) error {
		for _, b := range buildings.Buildings {
			res <- b
		}
		return nil
	})
}
