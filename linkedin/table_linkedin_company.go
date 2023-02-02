package linkedin

import (
	"context"

	"github.com/turbot/steampipe-plugin-sdk/v5/grpc/proto"
	"github.com/turbot/steampipe-plugin-sdk/v5/plugin"
	"github.com/turbot/steampipe-plugin-sdk/v5/plugin/transform"
)

func tableLinkedInCompany(ctx context.Context) *plugin.Table {
	return &plugin.Table{
		Name:        "linkedin_company",
		Description: "Get a company by name on LinkedIn.",
		List: &plugin.ListConfig{
			Hydrate: listCompany,
			KeyColumns: []*plugin.KeyColumn{
				{Name: "name", Require: plugin.Required},
			},
		},
		Columns: []*plugin.Column{
			// Top columns
			{Name: "name", Type: proto.ColumnType_STRING, Transform: transform.FromQual("name"), Description: "Name of the company."},
			{Name: "raw", Type: proto.ColumnType_JSON, Transform: transform.FromValue(), Description: "Company data."},
		},
	}
}

func listCompany(ctx context.Context, d *plugin.QueryData, _ *plugin.HydrateData) (interface{}, error) {

	conn, err := connect(ctx, d)
	if err != nil {
		plugin.Logger(ctx).Error("linkedin_company.listCompany", "connection_error", err)
		return nil, err
	}

	name := d.EqualsQuals["name"].GetStringValue()

	plugin.Logger(ctx).Warn("linkedin_company.listCompany", "status", "before CompanyByName")
	companyNode, err := conn.CompanyByName(name)
	plugin.Logger(ctx).Warn("linkedin_company.listCompany", "status", "after CompanyByName")
	if err != nil {
		plugin.Logger(ctx).Error("linkedin_company.listCompany", "query_error", err, "name", name)
		return nil, err
	}

	plugin.Logger(ctx).Warn("linkedin_company.listCompany", "companyNode", companyNode)
	plugin.Logger(ctx).Warn("linkedin_company.listCompany", "len(companyNode.Elements)", len(companyNode.Elements))
	//plugin.Logger(ctx).Warn("linkedin_company.listCompany", "companyNode.Elements[0]", companyNode.Elements[0])

	count := 0
	for {
		for _, company := range companyNode.Elements {
			plugin.Logger(ctx).Warn("linkedin_company.listCompany", "company loop", company)
			count++
			/*
				row := companySearchRow{
					company,
					count,
				}
			*/
			d.StreamListItem(ctx, company)
		}
		if count >= 50 || !companyNode.Next() {
			break
		}
	}

	return nil, nil
}
