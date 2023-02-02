package linkedin

import (
	"context"

	"github.com/tamboto2000/golinkedin"

	"github.com/turbot/steampipe-plugin-sdk/v5/grpc/proto"
	"github.com/turbot/steampipe-plugin-sdk/v5/plugin"
	"github.com/turbot/steampipe-plugin-sdk/v5/plugin/transform"
)

func tableLinkedInSearchProfile(ctx context.Context) *plugin.Table {
	return &plugin.Table{
		Name:        "linkedin_search_profile",
		Description: "Search people profiles on LinkedIn.",
		List: &plugin.ListConfig{
			Hydrate: listSearchProfile,
			KeyColumns: []*plugin.KeyColumn{
				{Name: "query", Require: plugin.Required, CacheMatch: "exact"},
			},
		},
		Columns: []*plugin.Column{
			// Top columns
			{Name: "rank", Type: proto.ColumnType_INT, Description: "Rank of the connection, use for sorting."},
			{Name: "id", Type: proto.ColumnType_INT, Transform: transform.FromField("TrackingUrn").Transform(extractProfileIDFromUrn), Description: "ID of the connection."},
			{Name: "title", Type: proto.ColumnType_STRING, Transform: transform.FromField("Title.Text"), Description: "Title of the connection, e.g. Jane Bloggs."},
			{Name: "headline", Type: proto.ColumnType_STRING, Transform: transform.FromField("Headline.Text"), Description: "Headline of the connection, e.g. CTO."},
			{Name: "subline", Type: proto.ColumnType_STRING, Transform: transform.FromField("Subline.Text"), Description: "Subline of the connection, e.g. New York."},
			{Name: "social_proof_text", Type: proto.ColumnType_STRING, Description: "Text indicating social connectedness, e.g. 148 shared connections."},
			{Name: "public_identifier", Type: proto.ColumnType_STRING, Description: "Unique public identifier of the profile, e.g. jbloggs."},
			// Other columns
			{Name: "member_distance", Type: proto.ColumnType_STRING, Transform: transform.FromField("MemberDistance.Value"), Description: "Distance of this connection from the profile_id provided."},
			{Name: "navigation_url", Type: proto.ColumnType_STRING, Description: "Web URL for the connection."},
			{Name: "secondary_title", Type: proto.ColumnType_STRING, Transform: transform.FromField("SecondaryTitle.Text"), Description: "Secondary title for the connection, e.g. 1st."},
			{Name: "target_urn", Type: proto.ColumnType_STRING, Description: "Target URN of the connection."},
			{Name: "tracking_id", Type: proto.ColumnType_STRING, Description: "Tracking ID of the connection."},
			{Name: "tracking_urn", Type: proto.ColumnType_STRING, Description: "Tracking URN of the connection."},
			// Qualifiers
			{Name: "query", Type: proto.ColumnType_STRING, Transform: transform.FromQual("query"), Description: "Optional query string to narrow the connections."},
		},
	}
}

func listSearchProfile(ctx context.Context, d *plugin.QueryData, _ *plugin.HydrateData) (interface{}, error) {

	conn, err := connect(ctx, d)
	if err != nil {
		plugin.Logger(ctx).Error("linkedin_search_profile.listSearchProfile", "connection_error", err)
		return nil, err
	}

	query := d.EqualsQuals["query"].GetStringValue()

	// Search notes:
	// * Settings generally follow the defaults from https://github.com/tamboto2000/golinkedin/blob/bcc335162ca03eeba91e6581a64994c3b350bb0d/profile.go#L161-L174
	// * I can't find a way to increase the page size about 10.  The linkedin UI
	//   seems to do pages of 10 as well, so I suspect it's deliberate.

	filter := &golinkedin.PeopleSearchFilter{
		ResultType: golinkedin.ResultPeople,
	}

	queryContext := &golinkedin.QueryContext{
		SpellCorrectionEnabled: true,
	}

	origin := golinkedin.OriginMemberProfileCannedSearch

	peopleNode, err := conn.SearchPeople(query, filter, queryContext, origin)
	if err != nil {
		plugin.Logger(ctx).Error("linkedin_search_profile.listSearchProfile", "query_error", err, "query", query)
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
		// Stop if we've:
		// * reached the end of the list
		// * reached the limit set in the query
		// * reached 100 results (a reasonable limit for search)
		if !peopleNode.Next() || d.RowsRemaining(ctx) <= 0 || rank >= 100 {
			break
		}
	}

	return nil, nil
}
