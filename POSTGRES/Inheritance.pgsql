CREATE TABLE logs (
                      inserted_date DATE NOT NULL DEFAULT now(),
                      log TEXT NOT NULL
);

CREATE TABLE connection_logs (
    user_usage INT NOT NULL
) INHERITS(logs);

-- #######
INSERT INTO logs (log)  values ('This should be LOGS');
INSERT INTO connection_logs (log, user_usage)  values ('This should be Connection Logs', 12);

-- #####
SELECT * FROM connection_logs;

-- ####
SELECT * FROM logs;
