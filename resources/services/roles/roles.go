package resources

import (
	"context"

	"github.com/cloudquery/plugin-sdk/v4/schema"
	"github.com/cloudquery/plugin-sdk/v4/transformers"
	"github.com/jsifuentes/cq-source-googleworkspace/client"
	directory "google.golang.org/api/admin/directory/v1"
)

func RolesTable() *schema.Table {
	return &schema.Table{
		Name:        "googleworkspace_roles",
		Description: "Google Workspace Roles",
		Transform:   transformers.TransformWithStruct(&directory.Role{}, transformers.WithPrimaryKeys("RoleId")),
		Columns: []schema.Column{
			client.CustomerIDColumn,
		},
		Resolver: fetchRoles,
	}
}

func fetchRoles(ctx context.Context, meta schema.ClientMeta, _ *schema.Resource, res chan<- any) error {
	c := meta.(*client.Client)
	return c.DirectoryService.Roles.List(c.Spec.CustomerID).Pages(ctx, func(roles *directory.Roles) error {
		for _, r := range roles.Items {
			res <- r
		}
		return nil
	})
}
