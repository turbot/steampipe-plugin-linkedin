---
title: "Steampipe Table: linkedin_profile - Query LinkedIn Profiles using SQL"
description: "Allows users to query LinkedIn Profiles, specifically to retrieve detailed information about the users' professional experiences, skills, education, and more."
---

# Table: linkedin_profile - Query LinkedIn Profiles using SQL

LinkedIn is a professional networking platform that allows users to create profiles, connect with others, and share content. Users' profiles contain detailed information about their professional experiences, skills, education, and more. This information can be leveraged for various purposes such as recruitment, networking, and business development.

## Table Usage Guide

The `linkedin_profile` table provides insights into LinkedIn profiles. As a recruiter or business development professional, explore profile-specific details through this table, including professional experiences, skills, and education. Utilize it to uncover information about potential candidates, such as their qualifications, experience, and skills, or to identify potential business partners.

**Important Notes**
- You must specify the `public_identifier` in the `where` clause to query this table.

## Examples

### Get profile information
Explore personal profile details on LinkedIn by specifying a user's public identifier. This can be useful in gathering industry-specific insights or understanding professional backgrounds.

```sql+postgres
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

```sql+sqlite
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
This query is used to gain insights into the professional history of a specific LinkedIn profile. It organizes the user's past positions by company, title, and tenure, allowing for a comprehensive review of their career progression.

```sql+postgres
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

```sql+sqlite
select
  json_extract(j.value, '$.companyName') as company_name,
  json_extract(j.value, '$.dateRange.start.year') as start_year,
  json_extract(j.value, '$.dateRange.end.year') as end_year,
  json_extract(j.value, '$.title') as title,
  json_extract(j.value, '$.description') as description
from
  linkedin_profile as p,
  json_each(positions) as c,
  json_each(json_extract(c.value, '$.profilePositionInPositionGroup.elements')) as j
where
  p.public_identifier = 'nathan-wallace-86470'
order by
  start_year desc;
```

### List skills for a profile
Explore which skills are associated with a specific LinkedIn profile. This can be used to assess an individual's proficiencies and understand their professional capabilities.

```sql+postgres
select
  s ->> 'name' as skill
from
  linkedin_profile as p,
  jsonb_array_elements(skills) as s
where
  p.public_identifier = 'dboeke';
```

```sql+sqlite
select
  json_extract(s.value, '$.name') as skill
from
  linkedin_profile as p,
  json_each(skills) as s
where
  p.public_identifier = 'dboeke';
```

### List education history for a profile
Explore an individual's educational history, including the schools they attended and the degrees they obtained, in chronological order. This can be useful for background checks or understanding a person's qualifications.

```sql+postgres
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

```sql+sqlite
select
  json_extract(e.value, '$.school.name') as school_name,
  json_extract(e.value, '$.degreeName') as degree_name,
  json_extract(e.value, '$.dateRange.start.year') as start_year,
  json_extract(e.value, '$.dateRange.end.year') as end_year
from
  linkedin_profile as p,
  json_each(education) as e
where
  p.public_identifier = 'e-gineer'
order by
  start_year desc;
```

### List certifications for a profile
Discover the range of certifications associated with a specific LinkedIn profile to understand the individual's skills and qualifications. This could be useful for recruiters or hiring managers assessing a candidate's expertise in a particular field.

```sql+postgres
select
  c ->> 'name' as skill
from
  linkedin_profile as p,
  jsonb_array_elements(certifications) as c
where
  p.public_identifier = 'dglosser';
```

```sql+sqlite
select
  json_extract(c.value, '$.name') as skill
from
  linkedin_profile as p,
  json_each(certifications) as c
where
  p.public_identifier = 'dglosser';
```

### List contact details for a profile
Explore the contact information associated with a specific LinkedIn profile. This can be useful for reaching out to potential collaborators, clients, or job candidates.

```sql+postgres
select
  first_name,
  last_name,
  contact_info ->> 'emailAddress' as email,
  contact_info -> 'address' as address
from
  linkedin_profile
where
  public_identifier = 'tuhintypical';
```

```sql+sqlite
select
  first_name,
  last_name,
  json_extract(contact_info, '$.emailAddress') as email,
  json_extract(contact_info, '$.address') as address
from
  linkedin_profile
where
  public_identifier = 'tuhintypical';
```

### List additional contact details from nested arrays for a profile
Determine the additional contact information for a specific LinkedIn profile. This could be useful for expanding your network or reaching out to potential business partners.

```sql+postgres
select
  first_name,
  last_name,
  contact_info ->> 'emailAddress' as email,
  contact_info -> 'address' as address,
  twitter.value ->> 'name' as twitter_handle,
  phone.value ->> 'number' as phone_number
from
  linkedin_profile
  left join
    jsonb_array_elements(contact_info -> 'twitterHandles') as twitter
    on true
  left join
    jsonb_array_elements(contact_info -> 'phoneNumbers') as phone
    on true
where
  public_identifier = 'tuhintypical';
```

```sql+sqlite
select
  first_name,
  last_name,
  json_extract(contact_info, '$.emailAddress') as email,
  contact_info as address,
  json_extract(twitter.value, '$.name') as twitter_handle,
  json_extract(phone.value, '$.number') as phone_number
from
  linkedin_profile,
  json_each(json_extract(contact_info, '$.twitterHandles')) as twitter,
  json_each(json_extract(contact_info, '$.phoneNumbers')) as phone
where
  public_identifier = 'tuhintypical';
```