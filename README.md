Taking [SurrealDB][surreal] for a spin. The domain I'm interested in exploring is wide, high-branching-factor trees, like the _organizational unit_ structures you'd find in [LDAP][ldap] systems. Imagine a forest of millions of trees, where each tree can have an arbitrary height (going up to, say, 10) and potentially thousands of children (“subordinate OUs”).

[surreal]: https://surrealdb.com/
[ldap]: https://en.wikipedia.org/wiki/Lightweight_Directory_Access_Protocol

## Surreal Wishlist

- Idempotent creates (e.g., `create if not exists`)
- Scripting / Sequences `for 1..10` syntax

## Starting

Spin it up with:

```shell
docker run --rm --pull always -p 8000:8000 surrealdb/surrealdb:latest start
```

## GUI

Surrealist has the most GitHub stars — by a long shot — so visit
https://surrealist.app and start a new session.

## Query

```
-- Select ancestors and descendants of Org Unit 123
SELECT
  /* parent */
  <-belongs_to<-org_unit,

  /* grandparent */
  <-belongs_to<-org_unit<-belongs_to<-org_unit,

  /* children */
  ->belongs_to->org_unit,

  /* grandchildren */
  ->belongs_to->org_unit->belongs_to->org_unit
FROM org_unit:123;
```

## Experimenting

### Importing Dataset

Still figuring this part out... There are [`export`][export] and
[`import`][import] commands.

[export]: https://docs.surrealdb.com/docs/cli/export/
[import]: https://docs.surrealdb.com/docs/cli/import/

```
surreal import
```