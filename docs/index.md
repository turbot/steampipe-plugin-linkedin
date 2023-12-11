---
organization: Turbot
category: ["media"]
icon_url: "/images/plugins/turbot/linkedin.svg"
brand_color: "#0077B5"
display_name: LinkedIn
name: linkedin
description: Steampipe plugin to query LinkedIn profiles.
og_description: Query LinkedIn with SQL! Open source CLI. No DB required.
og_image: "/images/plugins/turbot/linkedin-social-graphic.png"
engines: ["steampipe", "sqlite", "postgres", "export"]
---

# LinkedIn + Steampipe

[Steampipe](https://steampipe.io) is an open source CLI to instantly query cloud APIs using SQL.

[LinkedIn](https://linkedin.com) is a business and employment-focused social media platform.

Example query:

```sql
select
  first_name,
  last_name,
  headline,
  public_identifier,
  industry
from
  linkedin_profile
where
  public_identifier = 'dave';
```

```
+------------+-----------+------------------------------------+-------------------+-------------------+
| first_name | last_name | headline                           | public_identifier | industry          |
+------------+-----------+------------------------------------+-------------------+-------------------+
| Dave       | Jones     | Engineering Manager                | dave              | Computer Software |
+------------+-----------+------------------------------------+-------------------+-------------------+
```

## Documentation

- **[Table definitions & examples →](/plugins/turbot/linkedin/tables)**

## Get started

### Install

Download and install the latest LinkedIn plugin:

```bash
steampipe plugin install linkedin
```

### Configuration

Installing the latest linkedin plugin will create a config file (`~/.steampipe/config/linkedin.spc`) with a single connection named `linkedin`:

```hcl
connection "linkedin" {
  plugin  = "linkedin"
  // Set to value of the `li_at` cookie from a logged in LinkedIn browser session
  token   = "BQEDBQBCO8wBpUgWBBBBhhBtNDUBBBFFNDm4NU4BzcB32MRFFNDdx9md-Qk9_IFs6Lo4z8gTYb5HqUC5h-OjVDM22UiRUgjYX7CrYZw_IMIg6qUICOqqeCzjnb1tqIBIi7_HTZ3z2g_jq3JEqNjzqL7g"
}
```

LinkedIn does not offer an official API. This plugin uses the APIs normally
called by your browser when viewing the website. So, credentials are setup
using a browser based session cookie. To get your authentication token for
the plugin:
1. Browse to [LinkedIn](https://linkedin.com) and login.
2. Open your browser developer tools.
3. Open the Network view to see and analyze the network requests that make up each individual page load within a single user's session.
4. Open any request to a www.linked.com URL from the list and check the Cookies tab to get the list of request cookies.
5. The token is the value of the cookie named `li_at`.

Environment variables are also available as an alternate configuration method:
* `LINKEDIN_TOKEN`

## Get involved

* Open source: https://github.com/turbot/steampipe-plugin-linkedin
* Community: [Join #steampipe on Slack →](https://turbot.com/community/join)
