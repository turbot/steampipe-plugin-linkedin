![image](https://hub.steampipe.io/images/plugins/turbot/linkedin-social-graphic.png)

# LinkedIn Plugin for Steampipe

Use SQL to query profiles, companies, connections & more from [LinkedIn](https://linkedin.com).

* **[Get started →](https://hub.steampipe.io/plugins/turbot/linkedin)**
* Documentation: [Table definitions & examples](https://hub.steampipe.io/plugins/turbot/linkedin/tables)
* Community: [Slack Channel](https://steampipe.io/community/join)
* Get involved: [Issues](https://github.com/turbot/steampipe-plugin-linkedin/issues)

## Quick start

Install the plugin with [Steampipe](https://steampipe.io):

```shell
steampipe plugin install linkedin
```

Configure your token (from browser cookies) in `~/.steampipe/config/linkedin.spc`:

```hcl
connection "linkedin" {
  plugin  = "linkedin"
  // Set to value of the `li_at` cookie from a logged in LinkedIn browser session
  token   = "BQEDBQBCO8wBpUgWBBBBhhBtNDUBBBFFNDm4NU4BzcB32MRFFNDdx9md-Qk9_IFs6Lo4z8gTYb5HqUC5h-OjVDM22UiRUgjYX7CrYZw_IMIg6qUICOqqeCzjnb1tqIBIi7_HTZ3z2g_jq3JEqNjzqL7g"
}
```

Run steampipe:

```shell
steampipe query
```

Run a query:

```sql
select
  id,
  title,
  headline
from
  linkedin_company_employee
where
  company_id = 7599466
  and query = 'dave'
```

```
+----------+--------------+----------+
| id       | title        | headline |
+----------+--------------+----------+
| 13016000 | David Jones  | CTO      |
|  4819034 | Dave Beecham | Engineer |
+----------+--------------+----------+
```

## Developing

Prerequisites:

- [Steampipe](https://steampipe.io/downloads)
- [Golang](https://golang.org/doc/install)

Clone:

```sh
git clone https://github.com/turbot/steampipe-plugin-linkedin.git
cd steampipe-plugin-linkedin
```

Build, which automatically installs the new version to your `~/.steampipe/plugins` directory:

```
make
```

Configure the plugin:

```
cp config/* ~/.steampipe/config
vi ~/.steampipe/config/linkedin.spc
```

Try it!

```
steampipe query
> .inspect linkedin
```

Further reading:

- [Writing plugins](https://steampipe.io/docs/develop/writing-plugins)
- [Writing your first table](https://steampipe.io/docs/develop/writing-your-first-table)

## Contributing

Please see the [contribution guidelines](https://github.com/turbot/steampipe/blob/main/CONTRIBUTING.md) and our [code of conduct](https://github.com/turbot/steampipe/blob/main/CODE_OF_CONDUCT.md). All contributions are subject to the [Apache 2.0 open source license](https://github.com/turbot/steampipe-plugin-linkedin/blob/main/LICENSE).

`help wanted` issues:
- [Steampipe](https://github.com/turbot/steampipe/labels/help%20wanted)
- [LinkedIn Plugin](https://github.com/turbot/steampipe-plugin-linkedin/labels/help%20wanted)
