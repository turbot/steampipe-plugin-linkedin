# Table: linkedin_search_company

Search companies on LinkedIn.

Notes:
* `query` must be specified in the where clause of queries.
* Hard limit of 100 rows per search query.

## Examples

### Search for a company by name

```sql
select
  id,
  title,
  headline,
  subline
from
  linkedin_search_company
where
  query = 'turbot'
```