-- https://www.postgresqltutorial.com/postgresql-tutorial/postgresql-cte/
-- https://www.postgresql.org/docs/current/queries-with.html
WITH cte_film AS (
    SELECT
        film_id,
        title,
        (CASE
             WHEN length < 30 THEN 'Short'
             WHEN length < 90 THEN 'Medium'
             ELSE 'Long'
            END) length
    FROM
        film
)
SELECT
    film_id,
    title,
    length
FROM
    cte_film
WHERE
        length = 'Long'
ORDER BY
    title;