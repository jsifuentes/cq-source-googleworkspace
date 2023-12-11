package users

import (
	"context"

	"github.com/cloudquery/plugin-sdk/v4/schema"
	"github.com/cloudquery/plugin-sdk/v4/transformers"
	"github.com/jsifuentes/cq-source-googleworkspace/client"
	directory "google.golang.org/api/admin/directory/v1"
)

func OrgUnitsTable() *schema.Table {
	return &schema.Table{
		Name:        "googleworkspace_org_units",
		Description: "Google Workspace Org Units",
		Transform:   transformers.TransformWithStruct(&directory.OrgUnit{}, transformers.WithPrimaryKeys("OrgUnitId")),
		Columns: []schema.Column{
			client.CustomerIDColumn,
		},
		Resolver: fetchOrgUnits,
	}
}

func fetchOrgUnits(ctx context.Context, meta schema.ClientMeta, _ *schema.Resource, res chan<- any) error {
	c := meta.(*client.Client)

	result, err := c.DirectoryService.Orgunits.List(c.Spec.CustomerID).Do()
	if err != nil {
		return err
	}

	for _, ou := range result.OrganizationUnits {
		res <- ou
	}

	return nil
}
