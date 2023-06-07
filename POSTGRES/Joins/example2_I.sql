SELECT title, COUNT(title) as count FROM inventory as inv
INNER JOIN film ON inv.film_id = film.film_id
WHERE store_id = 1
GROUP BY title
ORDER BY count DESC
-- for long name we can use AS to shorten it !
