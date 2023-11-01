package linkedin

import (
	"context"
	"errors"
	"fmt"
	"os"
	"strconv"
	"strings"

	"github.com/tamboto2000/golinkedin"

	"github.com/turbot/steampipe-plugin-sdk/v5/plugin"
	"github.com/turbot/steampipe-plugin-sdk/v5/plugin/transform"
)

func connect(ctx context.Context, d *plugin.QueryData) (*golinkedin.Linkedin, error) {
	conn, err := connectCached(ctx, d, nil)
	if err != nil {
		return nil, err
	}
	return conn.(*golinkedin.Linkedin), nil
}

var connectCached = plugin.HydrateFunc(connectUncached).Memoize()

func connectUncached(ctx context.Context, d *plugin.QueryData, _ *plugin.HydrateData) (any, error) {

	conn := golinkedin.New()

	// Default to the env var settings
	token := os.Getenv("LINKEDIN_TOKEN")

	// Prefer config settings
	liConfig := GetConfig(d.Connection)
	if liConfig.Token != nil {
		token = *liConfig.Token
	}

	// Error if the minimum config is not set
	if token == "" {
		return conn, errors.New("token must be configured")
	}

	// JSESSIONID is needed to pass XSS checks, it can be any string.
	// li_at is the main authentication token.
	conn.SetCookieStr(fmt.Sprintf(`JSESSIONID="ajax:1"; li_at=%s`, token))

	return conn, nil
}

func extractProfileIDFromUrn(ctx context.Context, d *transform.TransformData) (interface{}, error) {
	urn := d.Value.(string)
	// TrackingUrn includes the profile ID
	plugin.Logger(ctx).Debug("extractProfileIDFromUrn", "urn", urn)
	parts := strings.Split(urn, ":")
	if len(parts) < 4 {
		plugin.Logger(ctx).Debug("extractProfileIDFromUrn", "urn", urn, "parts", parts)
		return nil, nil
	}
	id, err := strconv.Atoi(parts[3])
	if err != nil {
		// If the ID can't be extracted, then just return null for it and
		// log a warning.
		plugin.Logger(ctx).Warn("extractProfileIDFromUrn", "urn", urn, "id_extract_error", err)
		return nil, nil
	}
	return id, nil
}
