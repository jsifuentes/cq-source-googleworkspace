package users

import (
	"context"

	"github.com/cloudquery/plugin-sdk/v4/schema"
	"github.com/cloudquery/plugin-sdk/v4/transformers"
	directory "google.golang.org/api/admin/directory/v1"
)

type UserAliasRow struct {
	UserId           string `json:"user_id"`
	UserPrimaryEmail string `json:"user_primary_email"`
	Alias            string `json:"alias"`
}

func UserAliasesTable() *schema.Table {
	return &schema.Table{
		Name:        "googleworkspace_user_aliases",
		Description: "Google Workspace User Aliases",
		Transform:   transformers.TransformWithStruct(&UserAliasRow{}, transformers.WithPrimaryKeys("UserId", "Alias")),
		Resolver:    fetchUserAliases,
	}
}

func fetchUserAliases(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan<- any) error {
	user := parent.Item.(*directory.User)

	for _, alias := range user.Aliases {
		aliasRow := UserAliasRow{
			UserId:           user.Id,
			UserPrimaryEmail: user.PrimaryEmail,
			Alias:            alias,
		}
		res <- aliasRow
	}

	return nil
}
