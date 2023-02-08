# Table: linkedin_company_past_employee

List employees for a given company.

Notes:
* `company_id` must be specified in the `where` clause of queries.
* Use `query` to narrow the employee list by name, etc.
* Use `linkedin_search_company` to find a company ID.

## Examples

### List 10 past employees of a company

```sql
select
  id,
  title,
  headline
from
  linkedin_company_past_employee
where
  company_id = 7599466
limit
  10;
```

### List past employees called Dave

```sql
select
  id,
  title,
  headline
from
  linkedin_company_past_employee
where
  company_id = 7599466
  and query = 'dave';
```

### Find past employees of the company you are not directly connected to

```sql
select
  id,
  title,
  headline
from
  linkedin_company_past_employee
where
  company_id = 7599466
  and member_distance not in ('SELF', 'DISTANCE_1');
```
