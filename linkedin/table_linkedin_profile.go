package linkedin

import (
	"context"

	"github.com/turbot/steampipe-plugin-sdk/v5/grpc/proto"
	"github.com/turbot/steampipe-plugin-sdk/v5/plugin"
	"github.com/turbot/steampipe-plugin-sdk/v5/plugin/transform"
)

func tableLinkedInProfile(ctx context.Context) *plugin.Table {
	return &plugin.Table{
		Name:        "linkedin_profile",
		Description: "Get a profile by name on LinkedIn.",
		List: &plugin.ListConfig{
			Hydrate: listProfile,
			KeyColumns: []*plugin.KeyColumn{
				{Name: "public_identifier", Require: plugin.Required},
			},
		},
		Columns: []*plugin.Column{
			// Top columns
			{Name: "id", Type: proto.ColumnType_INT, Transform: transform.FromField("ObjectUrn").Transform(extractProfileIDFromUrn), Description: "ID of the connection."},
			{Name: "first_name", Type: proto.ColumnType_STRING, Description: "First name of the profile, e.g. Jane."},
			{Name: "last_name", Type: proto.ColumnType_STRING, Description: "Last name of the profile, e.g. Bloggs."},
			{Name: "headline", Type: proto.ColumnType_STRING, Description: "Headline of the profile, e.g. CTO."},
			{Name: "industry", Type: proto.ColumnType_STRING, Transform: transform.FromField("Industry.Name"), Description: "Industry for the profile, e.g. Information Technology & Services."},
			{Name: "public_identifier", Type: proto.ColumnType_STRING, Description: "Public identifier for the profile, e.g. jbloggs."},
			{Name: "summary", Type: proto.ColumnType_STRING, Description: "Summary text about the profile."},
			// Details related to the profile
			{Name: "education", Type: proto.ColumnType_JSON, Transform: transform.FromField("ProfileEducations.Elements"), Description: "Education information for the profile."},
			{Name: "skills", Type: proto.ColumnType_JSON, Transform: transform.FromField("ProfileSkills.Elements"), Description: "Skills information for the profile."},
			{Name: "positions", Type: proto.ColumnType_JSON, Transform: transform.FromField("ProfilePositionGroups.Elements"), Description: "Position history for the profile."},
			{Name: "certifications", Type: proto.ColumnType_JSON, Transform: transform.FromField("ProfileCertifications.Elements"), Description: "Certifications for the profile."},
			// Metadata about the profile
			{Name: "location", Type: proto.ColumnType_JSON, Description: "Location (postal code, country code, etc)of the profile."},
			{Name: "entity_urn", Type: proto.ColumnType_STRING, Transform: transform.FromField("EntityUrn"), Description: "URN of the profile object, urn:li:fsd_profile:ACoABADFm8QC8-YyrOV8IGjS8vbKTPjX2vrSjPM."},
			{Name: "object_urn", Type: proto.ColumnType_STRING, Transform: transform.FromField("ObjectUrn"), Description: "URN of the profile object, urn:li:member:12341234."},
			{Name: "industry_urn", Type: proto.ColumnType_STRING, Transform: transform.FromField("Industry.EntityUrn"), Description: "Industry for the profile, e.g. urn:li:fsd_industry:96."},
			{Name: "tracking_id", Type: proto.ColumnType_STRING, Description: "Tracking ID for the profile, e.g. /OHLHgmQRXifp9//zbqrOB==."},
			{Name: "version_tag", Type: proto.ColumnType_STRING, Description: "Version tag for the profile, e.g. 775349407."},
		},
	}
}

func listProfile(ctx context.Context, d *plugin.QueryData, _ *plugin.HydrateData) (interface{}, error) {

	conn, err := connect(ctx, d)
	if err != nil {
		plugin.Logger(ctx).Error("linkedin_profile.listProfile", "connection_error", err)
		return nil, err
	}

	publicIdentifier := d.EqualsQuals["public_identifier"].GetStringValue()

	profileNode, err := conn.ProfileByUsername(publicIdentifier)
	if err != nil {
		plugin.Logger(ctx).Error("linkedin_profile.listProfile", "query_error", err, "publicIdentifier", publicIdentifier)
		return nil, err
	}

	rank := 0
	for _, profile := range profileNode.Elements {
		rank++
		d.StreamListItem(ctx, profile)
		if d.RowsRemaining(ctx) <= 0 {
			break
		}
	}

	return nil, nil
}
