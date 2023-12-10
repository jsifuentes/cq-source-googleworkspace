package chrome

import (
	"context"

	"github.com/cloudquery/plugin-sdk/schema"
	"github.com/cloudquery/plugin-sdk/transformers"
	"github.com/jsifuentes/cq-source-googleworkspace/client"
	directory "google.golang.org/api/admin/directory/v1"
)

func ChromeDevicesTable() *schema.Table {
	return &schema.Table{
		Name:        "googleworkspace_chrome_devices",
		Description: "Google Workspace Chrome Devices",
		Transform:   transformers.TransformWithStruct(&directory.ChromeOsDevice{}, transformers.WithPrimaryKeys("DeviceId")),
		Columns: []schema.Column{
			client.CustomerIDColumn,
		},
		Resolver: fetchChromeDevices,
	}
}

func fetchChromeDevices(ctx context.Context, meta schema.ClientMeta, _ *schema.Resource, res chan<- any) error {
	c := meta.(*client.Client)

	return c.DirectoryService.Chromeosdevices.List(c.CustomerID).Pages(ctx, func(devices *directory.ChromeOsDevices) error {
		for _, d := range devices.Chromeosdevices {
			res <- d
		}
		return nil
	})
}
