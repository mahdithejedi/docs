-- __source__ = 'https://shiroyasha.io/selecting-for-share-and-update-in-postgresql.html'

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

CREATE OR REPLACE  FUNCTION update_coin_with_return(newName text, beginWith text) RETURNS INT as
    $BODY$
    DECLARE
        nameRows text[];
    BEGIN
        nameRows := array (SELECT name::TEXT FROM markets_coin WHERE name ILIKE beginWith || '%s' for update )::text[];
        IF array_length(nameRows, 0) THEN return 0; END IF;
--         SELECT pg_sleep(10);
        UPDATE markets_coin SET name=newName where name in (select unnest(nameRows));
        return 1;
    END;
    $BODY$ LANGUAGE plpgsql;

select update_coin_with_return('12MahdiNew12', 'j');

UPDATE markets_coin SET name='Mahdi' where name in (select unnest(array (SELECT name::TEXT FROM markets_coin WHERE name ILIKE 'M' || '%s')::text[])) returning *;