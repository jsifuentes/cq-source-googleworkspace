package main

import (
	"github.com/cloudquery/plugin-sdk/serve"
	"github.com/jsifuentes/cq-source-googleworkspace/resources/plugin"
)

func main() {
	serve.Source(plugin.Plugin()) // serve.WithSourceSentryDSN(sentryDSN),
}
