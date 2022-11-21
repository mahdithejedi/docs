SELECT  * FROM markets_coin;
UPDATE markets_coin SET unit = 'newUnit' where name = 'Muxo'
SELECT * FROM markets_coin where unit = 'newUnit';
-- UPDATED BUT IN ACCESS SHARE lock you can't update
begin;
lock table markets_coin IN ACCESS SHARE MODE ;
SELECT *, pg_sleep(10) FROM markets_coin where unit = 'newUnit';
end transaction;
-- __source__ = 'https://postgreshelp.com/postgresql-locks/'
--  __source__ = 'https://hevodata.com/learn/postgresql-locks/'

-- First of all markets_coin dll
-- Column   |           Type           | Collation | Nullable |                 Default
-- ------------+--------------------------+-----------+----------+------------------------------------------
--  id         | bigint                   |           | not null | nextval('markets_coin_id_seq'::regclass)
--  updated_at | timestamp with time zone |           |          |
--  created_at | timestamp with time zone |           | not null |
--  name       | character varying(128)   |           |          |
--  unit       | character varying(128)   |           |          |
-- Indexes:
--     "markets_coin_pkey" PRIMARY KEY, btree (id)
-- Referenced by:
--     TABLE "markets_market" CONSTRAINT "markets_market_base_currency_id_0bed12db_fk_markets_coin_id" FOREIGN KEY (base_currency_id) REFERENCES markets_coin(id) DEFERRABLE INITIALLY DEFERRED
--     TABLE "markets_market" CONSTRAINT "markets_market_currency_id_6f7ca41f_fk_markets_coin_id" FOREIGN KEY (currency_id) REFERENCES markets_coin(id) DEFERRABLE INITIALLY DEFERRED
--     TABLE "users_userasset" CONSTRAINT "users_userasset_coin_id_38e54ac9_fk_markets_coin_id" FOREIGN KEY (coin_id) REFERENCES markets_coin(id) DEFERRABLE INITIALLY DEFERRED


SELECT *, pg_sleep(10) FROM markets_coin where unit = 'newUnit';

select * from pg_locks;
-- BEFORE Lock
-- +----------+--------+--------+----+-----+----------+-------------+-------+-----+--------+------------------+------+---------------+-------+--------+
-- |locktype  |database|relation|page|tuple|virtualxid|transactionid|classid|objid|objsubid|virtualtransaction|pid   |mode           |granted|fastpath|
-- +----------+--------+--------+----+-----+----------+-------------+-------+-----+--------+------------------+------+---------------+-------+--------+
-- |relation  |56759   |12143   |NULL|NULL |NULL      |NULL         |NULL   |NULL |NULL    |3/18256           |325732|AccessShareLock|true   |true    |
-- |virtualxid|NULL    |NULL    |NULL|NULL |3/18256   |NULL         |NULL   |NULL |NULL    |3/18256           |325732|ExclusiveLock  |true   |true    |
-- +----------+--------+--------+----+-----+----------+-------------+-------+-----+--------+------------------+------+---------------+-------+--------+
-- However after pg_sleep we see something interesting
+----------+--------+--------+----+-----+----------+-------------+-------+-----+--------+------------------+------+---------------+-------+--------+
-- |locktype  |database|relation|page|tuple|virtualxid|transactionid|classid|objid|objsubid|virtualtransaction|pid   |mode           |granted|fastpath|
-- +----------+--------+--------+----+-----+----------+-------------+-------+-----+--------+------------------+------+---------------+-------+--------+
-- |relation  |56759   |66590   |NULL|NULL |NULL      |NULL         |NULL   |NULL |NULL    |5/450             |333236|AccessShareLock|true   |true    |
-- |relation  |56759   |66586   |NULL|NULL |NULL      |NULL         |NULL   |NULL |NULL    |5/450             |333236|AccessShareLock|true   |true    |
-- |virtualxid|NULL    |NULL    |NULL|NULL |5/450     |NULL         |NULL   |NULL |NULL    |5/450             |333236|ExclusiveLock  |true   |true    |
-- |relation  |56759   |12143   |NULL|NULL |NULL      |NULL         |NULL   |NULL |NULL    |3/18267           |325732|AccessShareLock|true   |true    |
-- |virtualxid|NULL    |NULL    |NULL|NULL |3/18267   |NULL         |NULL   |NULL |NULL    |3/18267           |325732|ExclusiveLock  |true   |true    |
-- +----------+--------+--------+----+-----+----------+-------------+-------+-----+--------+------------------+------+---------------+-------+--------+
-- The last two are for select from pg_locks itself
-- However the three from top are for for markets_coin, one for it's table the other for it's index and the third is another story
-- Each query in postgresql > 9 are run in transaction thus it need it's lock for transaction which is virtual ( locktype -> virtualxid)
-- *****************************

-- https://www.cockroachlabs.com/blog/select-for-update/#:~:text=SELECT%20FOR%20UPDATE%20is%20a,part%20of%20has%20been%20committed.
select *, pg_sleep(10) from markets_coin for update;
-- locktype  | database | relation | page | tuple | virtualxid | transactionid | classid | objid | objsubid | virtualtransaction |  pid   |      mode       | granted | fastpath
-- ------------+----------+----------+------+-------+------------+---------------+---------+-------+----------+--------------------+--------+-----------------+---------+----------
--  relation   |    56759 |    12143 |      |       |            |               |         |       |          | 5/492              | 333236 | AccessShareLock | t       | t
--  virtualxid |          |          |      |       | 5/492      |               |         |       |          | 5/492              | 333236 | ExclusiveLock   | t       | t
--  relation   |    56759 |    66590 |      |       |            |               |         |       |          | 3/18291            | 325732 | RowShareLock    | t       | t
--  relation   |    56759 |    66586 |      |       |            |               |         |       |          | 3/18291            | 325732 | RowShareLock    | t       | t
--  virtualxid |          |          |      |       | 3/18291    |               |         |       |          | 3/18291            | 325732 | ExclusiveLock   | t       | t

-- *****************************************
-- Deadlock
-- __source__ = 'https://www.cybertec-postgresql.com/en/postgresql-understanding-deadlocks/'
-- FIRST QUERY
begin;
update markets_coin set unit = 'NewUnit12DeadLockNEW' where id = 12 returning  *, pg_sleep(15);
update markets_coin set unit = 'NewUnit13DeadLockNEW' where id = 13 returning *, pg_sleep(10);
commit;

--  SECOND QUERY
begin;
update markets_coin set unit = 'NewUnit13NotDeadLockNEW' where id = 13 returning  *, pg_sleep(10);
update markets_coin set unit = 'NewUnit12NotDeadLockNEW' where id = 12 returning *, pg_sleep(10);
commit ;

-- If you run these transactions simounteneously you will see that first transaction will get deadlock

