package linkedin

import (
	"github.com/turbot/steampipe-plugin-sdk/v5/plugin"
)

type linkedinConfig struct {
	Token *string `hcl:"token"`
}

func ConfigInstance() interface{} {
	return &linkedinConfig{}
}

// GetConfig :: retrieve and cast connection config from query data
func GetConfig(connection *plugin.Connection) linkedinConfig {
	if connection == nil || connection.Config == nil {
		return linkedinConfig{}
	}
	config, _ := connection.Config.(linkedinConfig)
	return config
}
