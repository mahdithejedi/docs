# Include in Postgresql Index
## description
In PostgreSQL, a B-Tree index creates a multi-level tree structure where each level can be used as a doubly-linked list of pages. Leaf pages are those at the lowest level of a tree, that point to rows of tables.

With covering indexes, records of the columns mentioned in the INCLUDE clause are included in the leaf pages of the B-Tree as "payload" and are not part of the search key.

### in detail

Consider this example:
```postgresql
EXPLAIN ANALYZE
SELECT
    first_name,
    last_name,
    email
FROM
    "bankdb"
WHERE
    email = 'd-abbott3425@google.edu';
```
for this index
```postgresql
CREATE UNIQUE INDEX emails_idx
ON bankdb(email);
```
postgresql first will go to search in index and find related email,
then retrive data from table
however in this case:
```postgresql
CREATE UNIQUE INDEX emails_idx
ON bankdb(email)
INCLUDE(first_name,last_name);
```
postgresql will keep _first\_name_ and _last\_name_ in index itself now if you see
query plan:
```postgresql
QUERY PLAN
---------------------------------------
 Index Only Scan using emails_idx on bankdb  (cost=0.28..4.29 rows=1 width=37) (ac
tual time=0.228..0.231 rows=1 loops=1)
   Index Cond: (email = 'd-abbott3425@google.edu'::text)
   Heap Fetches: 0
 Planning Time: 0.233 ms
 Execution Time: 0.283 ms
(5 rows)
```
the _Heap Fetches_ is **0** because postgresql keep first_name and last_name inside of index itself.

# Points
1. **A non-key column (INCLUDE) cannot be used in an index scan search qualification, and it is disregarded for purposes of any uniqueness or exclusion constraint enforced by the index.**
# Sources
[atlasgo](https://atlasgo.io/guides/postgres/included-columns)
[postgresql documentation]()