package groups

import (
	"context"

	"github.com/apache/arrow/go/v14/arrow"
	"github.com/cloudquery/plugin-sdk/v4/schema"
	"github.com/cloudquery/plugin-sdk/v4/transformers"
	"github.com/jsifuentes/cq-source-googleworkspace/client"
	directory "google.golang.org/api/admin/directory/v1"
)

func GroupMembersTable() *schema.Table {
	return &schema.Table{
		Name:        "googleworkspace_group_members",
		Description: "Google Workspace Group Members",
		Transform:   transformers.TransformWithStruct(&directory.Member{}, transformers.WithPrimaryKeys("Id")),
		Resolver:    fetchGroupMembers,
		Columns: []schema.Column{
			client.CustomerIDColumn,
			{
				Name: "group_id",
				Type: arrow.BinaryTypes.String,
				Resolver: func(_ context.Context, meta schema.ClientMeta, r *schema.Resource, c schema.Column) error {
					return r.Set(c.Name, r.Parent.Item.(*directory.Group).Id)
				},
			},
			{
				Name: "group_email",
				Type: arrow.BinaryTypes.String,
				Resolver: func(_ context.Context, meta schema.ClientMeta, r *schema.Resource, c schema.Column) error {
					return r.Set(c.Name, r.Parent.Item.(*directory.Group).Email)
				},
			},
		},
	}
}

func fetchGroupMembers(ctx context.Context, meta schema.ClientMeta, resource *schema.Resource, res chan<- any) error {
	c := meta.(*client.Client)

	return c.DirectoryService.Members.List(resource.Item.(*directory.Group).Email).Pages(ctx, func(members *directory.Members) error {
		for _, m := range members.Members {
			res <- m
		}
		return nil
	})
}
