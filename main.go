package main

import (
	"github.com/turbot/steampipe-plugin-linkedin/linkedin"
	"github.com/turbot/steampipe-plugin-sdk/v6/plugin"
)

func main() {
	plugin.Serve(&plugin.ServeOpts{PluginFunc: linkedin.Plugin})
}
