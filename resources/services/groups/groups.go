package groups

import (
	"context"

	"github.com/cloudquery/plugin-sdk/v4/schema"
	"github.com/cloudquery/plugin-sdk/v4/transformers"
	"github.com/jsifuentes/cq-source-googleworkspace/client"
	directory "google.golang.org/api/admin/directory/v1"
)

func GroupsTable() *schema.Table {
	return &schema.Table{
		Name:        "googleworkspace_groups",
		Description: "Google Workspace Groups",
		Transform:   transformers.TransformWithStruct(&directory.Group{}, transformers.WithPrimaryKeys("Id")),
		Columns: []schema.Column{
			client.CustomerIDColumn,
		},
		Relations: []*schema.Table{
			GroupAliasesTable(),
			GroupMembersTable(),
		},
		Resolver: fetchGroups,
	}
}

func fetchGroups(ctx context.Context, meta schema.ClientMeta, _ *schema.Resource, res chan<- any) error {
	c := meta.(*client.Client)

	return c.DirectoryService.Groups.List().Customer(c.Spec.CustomerID).Pages(ctx, func(groups *directory.Groups) error {
		if groups == nil {
			return nil
		}
		for _, g := range groups.Groups {
			res <- g
		}
		return nil
	})
}
