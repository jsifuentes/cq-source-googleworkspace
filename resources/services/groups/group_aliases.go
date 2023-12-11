package groups

import (
	"context"

	"github.com/cloudquery/plugin-sdk/v4/schema"
	"github.com/cloudquery/plugin-sdk/v4/transformers"
	directory "google.golang.org/api/admin/directory/v1"
)

type GroupAliasRow struct {
	GroupId           string `json:"group_id"`
	GroupPrimaryEmail string `json:"group_primary_email"`
	Alias             string `json:"alias"`
}

func GroupAliasesTable() *schema.Table {
	return &schema.Table{
		Name:        "googleworkspace_group_aliases",
		Description: "Google Workspace Group Aliases",
		Transform:   transformers.TransformWithStruct(&GroupAliasRow{}, transformers.WithPrimaryKeys("GroupId", "Alias")),
		Resolver:    fetchGroupAliases,
	}
}

func fetchGroupAliases(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan<- any) error {
	group := parent.Item.(*directory.Group)

	for _, alias := range group.Aliases {
		aliasRow := GroupAliasRow{
			GroupId:           group.Id,
			GroupPrimaryEmail: group.Email,
			Alias:             alias,
		}
		res <- aliasRow
	}

	return nil
}
