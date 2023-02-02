package linkedin

import (
	"context"

	"github.com/turbot/steampipe-plugin-sdk/v5/grpc/proto"
	"github.com/turbot/steampipe-plugin-sdk/v5/plugin"
	"github.com/turbot/steampipe-plugin-sdk/v5/plugin/transform"
)

func tableLinkedInSearchCompany(ctx context.Context) *plugin.Table {
	return &plugin.Table{
		Name:        "linkedin_search_company",
		Description: "Search companies on LinkedIn.",
		List: &plugin.ListConfig{
			Hydrate: listSearchCompany,
			KeyColumns: []*plugin.KeyColumn{
				{Name: "query", Require: plugin.Required, CacheMatch: "exact"},
			},
		},
		Columns: []*plugin.Column{
			// Top columns
			{Name: "rank", Type: proto.ColumnType_INT, Description: "Rank of the company among the result rows, use for sorting."},
			{Name: "id", Type: proto.ColumnType_INT, Transform: transform.FromField("TrackingUrn").Transform(extractProfileIDFromUrn), Description: "ID of the company."},
			{Name: "title", Type: proto.ColumnType_STRING, Transform: transform.FromField("Title.Text"), Description: "Title of the company."},
			{Name: "headline", Type: proto.ColumnType_STRING, Transform: transform.FromField("Headline.Text"), Description: "Headline of the company."},
			{Name: "subline", Type: proto.ColumnType_STRING, Transform: transform.FromField("Subline.Text"), Description: "Subline of the company."},
			// Company metadata
			{Name: "target_urn", Type: proto.ColumnType_STRING, Description: "URN of the target."},
			{Name: "tracking_id", Type: proto.ColumnType_STRING, Description: "Tracking ID of the company."},
			{Name: "tracking_urn", Type: proto.ColumnType_STRING, Description: "Tracking URN of the company."},
			// Qualifiers
			{Name: "query", Type: proto.ColumnType_STRING, Transform: transform.FromQual("query"), Description: "Search query string."},
		},
	}
}

func listSearchCompany(ctx context.Context, d *plugin.QueryData, _ *plugin.HydrateData) (interface{}, error) {

	conn, err := connect(ctx, d)
	if err != nil {
		plugin.Logger(ctx).Error("linkedin_my_post.listSearchCompany", "connection_error", err)
		return nil, err
	}

	query := d.EqualsQuals["query"].GetStringValue()

	// Notes:
	// * CompanyByName returns invalid data, can only use SearchCompany API. I
	//   could not work out why.
	companyNode, err := conn.SearchCompany(query)
	if err != nil {
		plugin.Logger(ctx).Error("linkedin_my_post.listSearchCompany", "query_error", err, "query", query)
		return nil, err
	}

	rank := 0
	for {
		for _, company := range companyNode.Elements {
			for _, i := range company.Elements {
				rank++
				row := searchCompanyRow{
					i,
					rank,
				}
				d.StreamListItem(ctx, row)
			}
		}
		// Stop if we've:
		// * reached the end of the list
		// * reached the limit set in the query
		// * reached 100 results (a reasonable limit for search)
		if !companyNode.Next() || d.RowsRemaining(ctx) <= 0 || rank >= 100 {
			break
		}
	}

	return nil, nil
}
