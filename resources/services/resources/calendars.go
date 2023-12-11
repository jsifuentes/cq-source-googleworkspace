package resources

import (
	"context"

	"github.com/cloudquery/plugin-sdk/v4/schema"
	"github.com/cloudquery/plugin-sdk/v4/transformers"
	"github.com/jsifuentes/cq-source-googleworkspace/client"
	directory "google.golang.org/api/admin/directory/v1"
)

func CalendarsTable() *schema.Table {
	return &schema.Table{
		Name:        "googleworkspace_resources_calendars",
		Description: "Google Workspace Resources Calendars",
		Transform:   transformers.TransformWithStruct(&directory.CalendarResource{}, transformers.WithPrimaryKeys("ResourceId")),
		Columns: []schema.Column{
			client.CustomerIDColumn,
		},
		Resolver: fetchCalendars,
	}
}

func fetchCalendars(ctx context.Context, meta schema.ClientMeta, _ *schema.Resource, res chan<- any) error {
	c := meta.(*client.Client)
	return c.DirectoryService.Resources.Calendars.List(c.Spec.CustomerID).Pages(ctx, func(calendars *directory.CalendarResources) error {
		for _, c := range calendars.Items {
			res <- c
		}
		return nil
	})
}
