[//]: # (__POSTGRESQL__)
[//]: # (__POSTGRES__)
[//]: # (__POSTGRESQLPERFORMANCE__)
[//]: # (__POSTGRESQL_PERFORMANCE__)
[//]: # (__POSTGRESQL_WORK_MEM__)
[//]: # (__POSTGRESQLWORK_MEM__)
[//]: # (__POSTGRESPERFORMANCE__)
[//]: # (__POSTGRESPERFORMANCE__)
[//]: # (__PERFORMANCE__)
# Performance TIP

1- **update the default value for [`work_mem`](https://www.postgresql.org/docs/current/runtime-config-resource.html#GUC-WORK-MEM)** . This setting governs how much memory is available to each query operation
before it must start writing data to temporary files on disk, and can have a huge impact on performance. with the formula <small> [source](https://philbooth.me/blog/nine-ways-to-shoot-yourself-in-the-foot-with-postgresql) </small>
<br /> 
_work_mem = ($YOUR_INSTANCE_MEMORY * 0.8 - shared_buffers) / $YOUR_ACTIVE_CONNECTION_COUNT_
___
2- **Use recursive CTEs for time-critical queries** <small> [source](https://philbooth.me/blog/nine-ways-to-shoot-yourself-in-the-foot-with-postgresql) </small>
___
3- **_h1fre_ comment in Hackernews** <small> [source](https://news.ycombinator.com/item?id=35684220#35698301) </small>
- Configure Vacuum and maintenance_work_mem regularly if your DB size increases, if you allocate too much or too often it can clog up your memory.

- If you plan on deleting more than a 10000 rows regularly, maybe you should look at partition, it's surprisingly very slow to delete that "much" data. And even more with foreign key.

- Index on Boolean is useless, it's an easy mistake that will take memory and space disk for nothing.

- Broad indices are easier to maintain but if you can have multiple smaller indices with WHERE condition it will be much faster

- You can speed up, by a huge margin, big string indices with md5/hash index (only relevant for exact match)

- Postgres as a queue is definitely working and scales pretty far

- Related: be sure to understand the difference between transaction vs explicit locking, a lot of people assume too much from transaction and it will eventually breaks in prod.
___

