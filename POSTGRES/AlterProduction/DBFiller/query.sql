CREATE OR REPLACE FUNCTION new_update() RETURNS TRIGGER AS $BODY$
BEGIN
        NEW.name_new = NEW.name;
        NEW.unit_new = NEW.unit;
        NEW.unit = NULL;
        NEW.name = NULL;
RETURN NEW;
end
    $BODY$ LANGUAGE plpgsql;

create TRIGGER unit_update BEFORE INSERT ON markets_coin FOR EACH ROW EXECUTE PROCEDURE new_update();
DROP TRIGGER if exists unit_update on markets_coin;
SELECT * FROM markets_coin LIMIT 10;
UPDATE markets_coin SET unit_new=unit WHERE unit_new IS NULL;
UPDATE markets_coin SET name_new=unit WHERE name_new IS NULL;


BEGIN TRANSACTION ;
LOCK TABLE markets_coin IN SHARE ROW EXCLUSIVE MODE;
ALTER TABLE markets_coin rename COLUMN name_new to name;
ALTER TABLE markets_coin rename COLUMN unit_new to unit;
DROP TRIGGER if exists unit_update on markets_coin;
END TRANSACTION ;