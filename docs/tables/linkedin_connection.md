# Table: linkedin_connection

List connections (1st level only) for a given profile.

Notes:
* `profile_id` must be specified in the `where` clause of queries.
* Use `query` to narrow the connection list by name, etc.

## Examples

### List 10 connections for a profile

```sql
select
  c.id,
  c.title,
  c.headline
from
  linkedin_connection as c,
  linkedin_profile as p
where
  p.username = 'e-gineer'
  and c.profile_id = p.id
limit
  10;
```

### Find connections called Dave for a profile (by ID)

```sql
select
  id,
  title,
  headline
from
  linkedin_connection
where
  profile_id = 146380
  and query = 'dave';
```
