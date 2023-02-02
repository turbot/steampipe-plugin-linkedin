package linkedin

import (
	"context"
	"fmt"

	"github.com/tamboto2000/golinkedin"

	"github.com/turbot/steampipe-plugin-sdk/v5/grpc/proto"
	"github.com/turbot/steampipe-plugin-sdk/v5/plugin"
	"github.com/turbot/steampipe-plugin-sdk/v5/plugin/transform"
)

func tableLinkedInConnection(ctx context.Context) *plugin.Table {
	return &plugin.Table{
		Name:        "linkedin_connection",
		Description: "Connections (1st level) for a given profile.",
		List: &plugin.ListConfig{
			Hydrate: listConnection,
			KeyColumns: []*plugin.KeyColumn{
				{Name: "profile_id", Require: plugin.Required},
				{Name: "query", Require: plugin.Optional, CacheMatch: "exact"},
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
			{Name: "public_identifier", Type: proto.ColumnType_STRING, Description: "Unique public identifier of the connection, e.g. jbloggs."},
			// Other columns
			{Name: "member_distance", Type: proto.ColumnType_STRING, Transform: transform.FromField("MemberDistance.Value"), Description: "Distance of this connection from the profile_id provided."},
			{Name: "navigation_url", Type: proto.ColumnType_STRING, Description: "Web URL for the connection."},
			{Name: "secondary_title", Type: proto.ColumnType_STRING, Transform: transform.FromField("SecondaryTitle.Text"), Description: "Secondary title for the connection, e.g. 1st."},
			{Name: "target_urn", Type: proto.ColumnType_STRING, Description: "Target URN of the connection."},
			{Name: "tracking_id", Type: proto.ColumnType_STRING, Description: "Tracking ID of the connection."},
			{Name: "tracking_urn", Type: proto.ColumnType_STRING, Description: "Tracking URN of the connection."},
			// Qualifiers
			{Name: "profile_id", Type: proto.ColumnType_INT, Transform: transform.FromQual("profile_id"), Description: "ID of the profile for whom the list of connections is retrived. Required in the where clause."},
			{Name: "query", Type: proto.ColumnType_STRING, Transform: transform.FromQual("query"), Description: "Optional query string to narrow the connections."},
		},
	}
}

func listConnection(ctx context.Context, d *plugin.QueryData, _ *plugin.HydrateData) (interface{}, error) {

	conn, err := connect(ctx, d)
	if err != nil {
		plugin.Logger(ctx).Error("linkedin_connection.listConnection", "connection_error", err)
		return nil, err
	}

	profileID := d.EqualsQuals["profile_id"].GetInt64Value()
	query := d.EqualsQuals["query"].GetStringValue()

	// Search notes:
	// * Settings generally follow the defaults from https://github.com/tamboto2000/golinkedin/blob/bcc335162ca03eeba91e6581a64994c3b350bb0d/profile.go#L161-L174
	// * I can't find a way to increase the page size about 10.  The linkedin UI
	//   seems to do pages of 10 as well, so I suspect it's deliberate.

	filter := &golinkedin.PeopleSearchFilter{
		// Limit the search to connections for the given profileID
		ConnectionOf: fmt.Sprintf("%d", profileID),
		// Target F(irst) level connections only.
		// First and Second = []string{"F", "S"}
		// First, Second & Third+ = []string{"F", "S", "O"}
		Network: []string{golinkedin.Rank1},
		// Only return people
		ResultType: golinkedin.ResultPeople,
	}

	queryContext := &golinkedin.QueryContext{
		SpellCorrectionEnabled: true,
	}

	origin := golinkedin.OriginMemberProfileCannedSearch

	peopleNode, err := conn.SearchPeople(query, filter, queryContext, origin)
	if err != nil {
		plugin.Logger(ctx).Error("linkedin_connection.listConnection", "query_error", err, "profileID", profileID, "query", query)
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
