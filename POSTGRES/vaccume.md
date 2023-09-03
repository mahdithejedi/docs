# What is _vacuum_?

A vacuum is used for recovering space occupied by “dead tuples” in a table. A dead tuple is created when a record is
either deleted or updated (a delete followed by an insert). PostgreSQL doesn’t physically remove the old row from the
table but puts a “marker” on it so that queries don’t return that row. When a vacuum process runs, the space occupied by
these dead tuples is marked reusable by other tuples.

## usage
```postgresql
VACUUM [FULL] [FREEZE] [VERBOSE] [table_name ];
```
* **FULL**: _Optional_. If specified, the database writes the full contents of the table into a new file. This reclaims all unused space and requires an exclusive lock on each table that is vacuumed.
* **FREEZE**: _Optional_. If specified, the tuples are aggressively frozen when the table is vacuumed. This is the default behavior when FULL is specified, so it is redundant to specify both FULL and FREEZE.
* **VERBOSE**: _Optional_. If specified, an activity report will be printed detailing the vacuum activity for each table.
* **ANALYZE**: _Optional_. If specified, the statistics used by the planner will be updated. These statistics are used to determine the most efficient plan for executing a particular query.
* **table_name**: _Optional_. If specified, only the table listed will be vacuumed. If not specified, all tables in the database will be vacuumed.
* **col1, col2, ... col_n**: _Optional_. If specified, these are the columns that will be analyzed

<br />

## Points
### Don’t Run Manual VACUUM or ANALYZE Without Reason

PostgreSQL vacuuming (autovacuum or manual vacuum) minimizes table bloats and prevents transaction ID wraparound.
Autovacuum does not recover the disk space taken up by dead tuples. However, running a VACUUM FULL command will do so.
VACUUM FULL has its performance implication, though. The target table is exclusively locked during the operation,
preventing even reads on the table.

It’s also a best practice to not run manual vacuums too often on the entire database; the target database could be
already optimally vacuumed by the autovacuum process. As a result, a manual vacuum may not remove any dead tuples but
cause unnecessary I/O loads or CPU spikes. If necessary, manual vacuums should be only run on a table-by-table basis
when there’s a need for it, like low ratios of live rows to dead rows, or large gaps between autovacuums. Also, manual
vacuums should be run when user activity is minimum.

Autovacuum also keeps a table’s data distribution statistics up-to-date (it doesn’t rebuild them). When manually run,
the ANALYZE command actually rebuilds these statistics instead of updating them. Again, rebuilding statistics when
they’re already optimally updated by a regular autovacuum might cause unnecessary pressure on system resources. 

**The time when you must run ANALYZE manually is immediately after bulk loading data into the target table.**


#### Sources
[Enterprisedb vacuum](https://www.enterprisedb.com/blog/postgresql-vacuum-and-analyze-best-practice-tips)
< br/>