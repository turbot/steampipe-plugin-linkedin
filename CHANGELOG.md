## v0.4.0 [2023-11-14]

_Breaking Changes_

- Removed the following tables using the search API that no longer work due to API limitations. These tables will be added back if functionality can be restoerd.
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
