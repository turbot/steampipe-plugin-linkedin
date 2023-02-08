# Table: linkedin_search_profile

Search profiles of people on LinkedIn.

Notes:
* `query` must be specified in the `where` clause of queries.
* Hard limit of 100 rows per search query.

## Examples

### Search for a profile by name

```sql
select
  id,
  title,
  headline
from
  linkedin_search_profile
where
  query = 'nathan wallace';
```

### Search for a profile by name and company

```sql
select
  id,
  title,
  headline
from
  linkedin_search_profile
where
  query = 'nathan wallace turbot';
```
