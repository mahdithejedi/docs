## Concurrent Index 2phase, How it works?

* Phase 1: A snapshot of the current data gets taken, and the index is built on that.
* Phase 2: Postgres then catches up with any changes (inserts, updates, or deletes) that happened during phase

## Concurrent Index can goes *WRONG*

Since concurrent Index is asynchronous, the CREATE INDEX command might fail, leaving an incomplete index behind. An
“invalid” index is ignored during querying, but this oversight can have serious consequences if not monitored.

### What can cause concurrent index to fail?

* **Deadlocks**: Index creation might conflict with ongoing transactions, leading to deadlocks.
* **Disk Space**: Large indexes may fail due to insufficient disk space.
* **Constraint Violations**: Creating unique indexes on columns with non-unique data will result in failures.

**You can find all invalid indexes by running the following:**

```postgresql
SELECT *
FROM pg_class,
     pg_index
WHERE pg_index.indisvalid = false
  AND pg_index.indexrelid = pg_class.oid;
```

### How to monitor concurrent index creation?

The system table `pg_stat_progress_create_index` can be queried for progress reporting while indexing is taking place.

```
postgres=# SELECT * FROM pg_stat_progress_create_index;
-[ RECORD 1 ]------+---------------------------------------
pid                | 896799
datid              | 16402
datname            | postgres
relid              | 17261
index_relid        | 136565
command            | CREATE INDEX CONCURRENTLY
phase              | building index: loading tuples in tree
lockers_total      | 0
lockers_done       | 0
current_locker_pid | 0
blocks_total       | 0
blocks_done        | 0
tuples_total       | 10091384
tuples_done        | 1775295
partitions_total   | 0
partitions_done    | 0
```

### How to fix invalid indexes?

Invalid indexes can be recovered using the _REINDEX_ command. It’s the same as dropping and recreating the index, except
it would also lock out reads that attempt to use that index (if not specifying CONCURRENTLY). **Note that CONCURRENTLY
reindexing isn’t supported in versions below Postgres 12.**

```postgresql
REINDEX INDEX CONCURRENTLY idx_users_email_2019;
```

# Sources

https://blog.bemi.io/indexing/
<br />
