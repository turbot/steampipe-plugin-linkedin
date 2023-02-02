package linkedin

import (
	"context"

	"github.com/tamboto2000/golinkedin"

	"github.com/turbot/steampipe-plugin-sdk/v5/grpc/proto"
	"github.com/turbot/steampipe-plugin-sdk/v5/plugin"
	"github.com/turbot/steampipe-plugin-sdk/v5/plugin/transform"
)

func tableLinkedInCompanyPastEmployee(ctx context.Context) *plugin.Table {
	return &plugin.Table{
		Name:        "linkedin_company_past_employee",
		Description: "Search former employees for a given company on LinkedIn.",
		List: &plugin.ListConfig{
			Hydrate: listCompanyPastEmployee,
			KeyColumns: []*plugin.KeyColumn{
				{Name: "company_id", Require: plugin.Required},
				{Name: "query", Require: plugin.Optional, CacheMatch: "exact"},
			},
		},
		Columns: []*plugin.Column{
			// Top columns
			{Name: "rank", Type: proto.ColumnType_INT, Description: "Rank of the employee, use for sorting."},
			{Name: "id", Type: proto.ColumnType_INT, Transform: transform.FromField("TrackingUrn").Transform(extractProfileIDFromUrn), Description: "ID of the employee."},
			{Name: "title", Type: proto.ColumnType_STRING, Transform: transform.FromField("Title.Text"), Description: "Title of the employee, e.g. Jane Bloggs."},
			{Name: "headline", Type: proto.ColumnType_STRING, Transform: transform.FromField("Headline.Text"), Description: "Headline of the employee, e.g. CTO."},
			{Name: "subline", Type: proto.ColumnType_STRING, Transform: transform.FromField("Subline.Text"), Description: "Subline of the employee, e.g. New York."},
			{Name: "social_proof_text", Type: proto.ColumnType_STRING, Description: "Text indicating social connectedness, e.g. 148 shared employees."},
			{Name: "public_identifier", Type: proto.ColumnType_STRING, Description: "Unique public identifier of the employee, e.g. jbloggs."},
			// Other columns
			{Name: "member_distance", Type: proto.ColumnType_STRING, Transform: transform.FromField("MemberDistance.Value"), Description: "Distance of this employee from the caller."},
			{Name: "navigation_url", Type: proto.ColumnType_STRING, Description: "Web URL for the employee."},
			{Name: "secondary_title", Type: proto.ColumnType_STRING, Transform: transform.FromField("SecondaryTitle.Text"), Description: "Secondary title for the employee, e.g. 1st."},
			{Name: "target_urn", Type: proto.ColumnType_STRING, Description: "Target URN of the employee."},
			{Name: "tracking_id", Type: proto.ColumnType_STRING, Description: "Tracking ID of the employee."},
			{Name: "tracking_urn", Type: proto.ColumnType_STRING, Description: "Tracking URN of the employee."},
			// Qualifiers
			{Name: "company_id", Type: proto.ColumnType_INT, Transform: transform.FromQual("company_id"), Description: "ID of the company to list employees for."},
			{Name: "query", Type: proto.ColumnType_STRING, Transform: transform.FromQual("query"), Description: "Optional query string to narrow the employees."},
		},
	}
}

func listCompanyPastEmployee(ctx context.Context, d *plugin.QueryData, _ *plugin.HydrateData) (interface{}, error) {

	conn, err := connect(ctx, d)
	if err != nil {
		plugin.Logger(ctx).Error("linkedin_company_past_employee.listCompanyPastEmployee", "connection_error", err)
		return nil, err
	}

	companyID := int(d.EqualsQuals["company_id"].GetInt64Value())
	query := d.EqualsQuals["query"].GetStringValue()

	// Search notes:
	// * Settings generally follow the defaults from https://github.com/tamboto2000/golinkedin/blob/bcc335162ca03eeba91e6581a64994c3b350bb0d/profile.go#L161-L174
	// * I can't find a way to increase the page size about 10.  The linkedin UI
	//   seems to do pages of 10 as well, so I suspect it's deliberate.

	filter := &golinkedin.PeopleSearchFilter{
		// Limit the search to people currently in the company
		PastCompany: []int{companyID},
		// Only return people
		ResultType: golinkedin.ResultPeople,
	}

	queryContext := &golinkedin.QueryContext{
		SpellCorrectionEnabled: true,
	}

	origin := golinkedin.OriginMemberProfileCannedSearch

	peopleNode, err := conn.SearchPeople(query, filter, queryContext, origin)
	if err != nil {
		plugin.Logger(ctx).Error("linkedin_company_past_employee.listCompanyPastEmployee", "query_error", err, "companyID", companyID, "query", query)
		return nil, err
	}

	rank := 0
	for {
		for _, person := range peopleNode.Elements {
			for _, i := range person.Elements {
				rank++
				row := searchProfileRow{
					i,
					rank,
				}
				d.StreamListItem(ctx, row)
			}
		}
		if !peopleNode.Next() || d.RowsRemaining(ctx) <= 0 {
			break
		}
	}

	return nil, nil
}
