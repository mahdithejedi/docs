SELECT 
language.language_id, language.last_update, -- from language table
title, description,rental_duration, release_year, rental_duration, rental_rate,rating,
film_id
FROM film
INNER JOIN language ON language.language_id = film.language_id
ORDER BY rental_duration 