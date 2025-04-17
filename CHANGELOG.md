## v1.1.0 [2025-04-17]

_Dependencies_

- Recompiled plugin with Go version `1.23.1`. ([#46](https://github.com/turbot/steampipe-plugin-linkedin/pull/46))
- Recompiled plugin with [steampipe-plugin-sdk v5.11.5](https://github.com/turbot/steampipe-plugin-sdk/blob/v5.11.5/CHANGELOG.md#v5115-2025-03-31) that addresses critical and high vulnerabilities in dependent packages. ([#46](https://github.com/turbot/steampipe-plugin-linkedin/pull/46))

## v1.0.0 [2024-10-22]

There are no significant changes in this plugin version; it has been released to align with [Steampipe's v1.0.0](https://steampipe.io/changelog/steampipe-cli-v1-0-0) release. This plugin adheres to [semantic versioning](https://semver.org/#semantic-versioning-specification-semver), ensuring backward compatibility within each major version.

_Dependencies_

- Recompiled plugin with Go version `1.22`. ([#36](https://github.com/turbot/steampipe-plugin-linkedin/pull/36))
- Recompiled plugin with [steampipe-plugin-sdk v5.10.4](https://github.com/turbot/steampipe-plugin-sdk/blob/develop/CHANGELOG.md#v5104-2024-08-29) that fixes logging in the plugin export tool. ([#36](https://github.com/turbot/steampipe-plugin-linkedin/pull/36))

## v0.5.1 [2023-12-12]

- Fixed the connection config vriable definition to only use `hcl` syntax.

## v0.5.0 [2023-12-12]

_What's new?_

- The plugin can now be downloaded and used with the [Steampipe CLI](https://steampipe.io/docs), as a [Postgres FDW](https://steampipe.io/docs/steampipe_postgres/overview), as a [SQLite extension](https://steampipe.io/docs//steampipe_sqlite/overview) and as a standalone [exporter](https://steampipe.io/docs/steampipe_export/overview). ([#23](https://github.com/turbot/steampipe-plugin-linkedin/pull/23))
- The table docs have been updated to provide corresponding example queries for Postgres FDW and SQLite extension. ([#23](https://github.com/turbot/steampipe-plugin-linkedin/pull/23))
- Docs license updated to match Steampipe [CC BY-NC-ND license](https://github.com/turbot/steampipe-plugin-linkedin/blob/main/docs/LICENSE). ([#23](https://github.com/turbot/steampipe-plugin-linkedin/pull/23))

_Dependencies_

- Recompiled plugin with [steampipe-plugin-sdk v5.8.0](https://github.com/turbot/steampipe-plugin-sdk/blob/main/CHANGELOG.md#v580-2023-12-11) that includes plugin server encapsulation for in-process and GRPC usage, adding Steampipe Plugin SDK version to `_ctx` column, and fixing connection and potential divide-by-zero bugs. ([#22](https://github.com/turbot/steampipe-plugin-linkedin/pull/22))

## v0.4.0 [2023-11-14]

_Breaking Changes_

- Removed the following tables using the search API that no longer work due to API limitations. These tables will be added back if functionality can be restored.
  - `linkedin_company_employee`
  - `linkedin_company_past_employee`
  - `linkedin_connection`
  - `linkedin_search_company`
  - `linkedin_search_profile`

## v0.3.0 [2023-11-01]

_Enhancements_

- Added the `contact_info` column to `linkedin_profile` table. ([#5](https://github.com/turbot/steampipe-plugin-linkedin/pull/5))

## v0.2.1 [2023-10-05]

_Dependencies_

- Recompiled plugin with [steampipe-plugin-sdk v5.6.2](https://github.com/turbot/steampipe-plugin-sdk/blob/main/CHANGELOG.md#v562-2023-10-03) which prevents nil pointer reference errors for implicit hydrate configs. ([#11](https://github.com/turbot/steampipe-plugin-linkedin/pull/11))

## v0.2.0 [2023-10-02]

_Dependencies_

- Upgraded to [steampipe-plugin-sdk v5.6.1](https://github.com/turbot/steampipe-plugin-sdk/blob/main/CHANGELOG.md#v561-2023-09-29) with support for rate limiters. ([#9](https://github.com/turbot/steampipe-plugin-linkedin/pull/9))
- Recompiled plugin with Go version `1.21`. ([#9](https://github.com/turbot/steampipe-plugin-linkedin/pull/9))

## v0.1.0 [2023-04-12]

_Dependencies_

- Recompiled plugin with [steampipe-plugin-sdk v5.3.0](https://github.com/turbot/steampipe-plugin-sdk/blob/main/CHANGELOG.md#v530-2023-03-16) which adds go-getter support to dynamic tables. ([#3](https://github.com/turbot/steampipe-plugin-linkedin/pull/3))

## v0.0.1 [2023-02-08]

_What's new?_

- New tables added
  - [linkedin_company_employee](https://hub.steampipe.io/plugins/turbot/linkedin/tables/linkedin_company_employee)
  - [linkedin_company_past_employee](https://hub.steampipe.io/plugins/turbot/linkedin/tables/linkedin_company_past_employee)
  - [linkedin_connection](https://hub.steampipe.io/plugins/turbot/linkedin/tables/linkedin_connection)
  - [linkedin_profile](https://hub.steampipe.io/plugins/turbot/linkedin/tables/linkedin_profile)
  - [linkedin_search_company](https://hub.steampipe.io/plugins/turbot/linkedin/tables/linkedin_search_company)
  - [linkedin_search_profile](https://hub.steampipe.io/plugins/turbot/linkedin/tables/linkedin_search_profile)
