package linkedin

import (
	"context"

	"github.com/turbot/steampipe-plugin-sdk/v5/plugin"
	"github.com/turbot/steampipe-plugin-sdk/v5/plugin/transform"
)

// Plugin creates this (linkedin) plugin
func Plugin(ctx context.Context) *plugin.Plugin {
	p := &plugin.Plugin{
		Name: "steampipe-plugin-linkedin",
		ConnectionConfigSchema: &plugin.ConnectionConfigSchema{
			NewInstance: ConfigInstance,
		},
		DefaultTransform: transform.FromGo().NullIfZero(),
		TableMap: map[string]*plugin.Table{
			"linkedin_profile": tableLinkedInProfile(ctx),

			// The below tables are always returning 404 errors, so they are commented out for now.
			// "linkedin_company_employee":      tableLinkedInCompanyEmployee(ctx),
			// "linkedin_company_past_employee": tableLinkedInCompanyPastEmployee(ctx),
			// "linkedin_connection":            tableLinkedInConnection(ctx),
			// "linkedin_search_company":        tableLinkedInSearchCompany(ctx),
			// "linkedin_search_profile":        tableLinkedInSearchProfile(ctx),
		},
	}
	return p
}
