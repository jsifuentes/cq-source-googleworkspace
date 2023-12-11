package domains

import (
	"context"

	"github.com/cloudquery/plugin-sdk/v4/schema"
	"github.com/cloudquery/plugin-sdk/v4/transformers"
	"github.com/jsifuentes/cq-source-googleworkspace/client"
	directory "google.golang.org/api/admin/directory/v1"
)

func DomainsTable() *schema.Table {
	return &schema.Table{
		Name:        "googleworkspace_domains",
		Description: "Google Workspace Domains",
		Transform:   transformers.TransformWithStruct(&directory.Domains{}, transformers.WithPrimaryKeys("DomainName")),
		Columns: []schema.Column{
			client.CustomerIDColumn,
		},
		Resolver: fetchDomains,
	}
}

func fetchDomains(ctx context.Context, meta schema.ClientMeta, _ *schema.Resource, res chan<- any) error {
	c := meta.(*client.Client)

	domains, err := c.DirectoryService.Domains.List(c.Spec.CustomerID).Do()
	if err != nil {
		return err
	}

	for _, d := range domains.Domains {
		res <- d
	}

	return nil
}
