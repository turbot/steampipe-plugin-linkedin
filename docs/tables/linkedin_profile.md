# Table: linkedin_profile

Get profile information for a given user.

Notes:
* `public_identifier` must be specified in the `where` clause of queries.

## Examples

### Get profile information

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
  public_identifier = 'dboeke';
```

### List positions for a profile

```sql
select
  j ->> 'companyName' as company_name,
  (j -> 'dateRange' -> 'start' -> 'year')::int as start_year,
  (j -> 'dateRange' -> 'end' -> 'year')::int as end_year,
  j ->> 'title' as title,
  j ->> 'description' as description
from
  linkedin_profile as p,
  jsonb_array_elements(positions) as c,
  jsonb_array_elements(c -> 'profilePositionInPositionGroup' -> 'elements') as j
where
  p.public_identifier = 'nathan-wallace-86470'
order by
  start_year desc;
```

### List skills for a profile

```sql
select
  s ->> 'name' as skill
from
  linkedin_profile as p,
  jsonb_array_elements(skills) as s
where
  p.public_identifier = 'dboeke';
```

### List education history for a profile

```sql
select
  e -> 'school' ->> 'name' as school_name,
  e ->> 'degreeName' as degree_name,
  (e -> 'dateRange' -> 'start' -> 'year')::int as start_year,
  (e -> 'dateRange' -> 'end' -> 'year')::int as end_year
from
  linkedin_profile as p,
  jsonb_array_elements(education) as e
where
  p.public_identifier = 'e-gineer'
order by
  start_year desc;
```

### List certifications for a profile

```sql
select
  c ->> 'name' as skill
from
  linkedin_profile as p,
  jsonb_array_elements(certifications) as c
where
  p.public_identifier = 'dglosser';
```
