package plugin

import (
	"strings"

	"github.com/cloudquery/plugin-sdk/caser"
	"github.com/cloudquery/plugin-sdk/plugins/source"
	"github.com/cloudquery/plugin-sdk/schema"
	"github.com/jsifuentes/cq-source-googleworkspace/client"
	"github.com/jsifuentes/cq-source-googleworkspace/resources/services/domains"
	"github.com/jsifuentes/cq-source-googleworkspace/resources/services/groups"
	"github.com/jsifuentes/cq-source-googleworkspace/resources/services/resources"
	"github.com/jsifuentes/cq-source-googleworkspace/resources/services/users"
	"golang.org/x/exp/maps"
)

var Version = "Development"

var titleTransformers = map[string]string{
	"googleworkspace": "Google Workspace",
}

func titleTransformer(table *schema.Table) string {
	if table.Title != "" {
		return table.Title
	}
	exceptions := maps.Clone(source.DefaultTitleExceptions)
	for k, v := range titleTransformers {
		exceptions[k] = v
	}
	csr := caser.New(caser.WithCustomExceptions(exceptions))
	t := csr.ToTitle(table.Name)
	return strings.Trim(strings.ReplaceAll(t, "  ", " "), " ")
}

func Plugin() *source.Plugin {
	return source.NewPlugin(
		"googleworkspace",
		Version,
		[]*schema.Table{
			users.UsersTable(),
			groups.GroupsTable(),
			domains.DomainsTable(),
			resources.FeaturesTable(),
			resources.BuildingsTable(),
			resources.CalendarsTable(),
			users.OrgUnitsTable(),
		},
		client.Configure,
		source.WithTitleTransformer(titleTransformer),
	)
}
