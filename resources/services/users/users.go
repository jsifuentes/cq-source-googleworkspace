package users

import (
	"context"

	"github.com/cloudquery/plugin-sdk/schema"
	"github.com/cloudquery/plugin-sdk/transformers"
	"github.com/jsifuentes/cq-source-googleworkspace/client"
	directory "google.golang.org/api/admin/directory/v1"
)

func UsersTable() *schema.Table {
	return &schema.Table{
		Name:        "googleworkspace_users",
		Description: "Google Workspace Users",
		Transform:   transformers.TransformWithStruct(&directory.User{}, transformers.WithPrimaryKeys("Id")),
		Columns: []schema.Column{
			client.CustomerIDColumn,
			{
				Name: "first_name",
				Type: schema.TypeString,
				Resolver: func(_ context.Context, meta schema.ClientMeta, r *schema.Resource, c schema.Column) error {
					return r.Set(c.Name, r.Item.(*directory.User).Name.GivenName)
				},
			},
			{
				Name: "last_name",
				Type: schema.TypeString,
				Resolver: func(_ context.Context, meta schema.ClientMeta, r *schema.Resource, c schema.Column) error {
					return r.Set(c.Name, r.Item.(*directory.User).Name.FamilyName)
				},
			},
		},
		Relations: []*schema.Table{
			UserAliasesTable(),
		},
		Resolver: fetchUsers,
	}
}

func fetchUsers(ctx context.Context, meta schema.ClientMeta, _ *schema.Resource, res chan<- any) error {
	c := meta.(*client.Client)
	return c.DirectoryService.Users.List().Customer(c.CustomerID).Pages(ctx, func(users *directory.Users) error {
		for _, u := range users.Users {
			res <- u
		}
		return nil
	})
}
