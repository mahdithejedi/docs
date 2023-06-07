--   /*****GROUP BY*****\
--is used for grouping the rows with the same amount
SELECT rating,round(AVG(rental_rate),2)
FROM film
GROUP BY rating
ORDER BY rating DESC;
-- In this example we say to SQL to get the avarage of replacement_cost from the rating coloume 
-- WHERE the amount of rating in the rows are same
-- that means at first find the rows which thier rating amount are same and do AVG on them
-- SmallPoint -> as we've said in the last episot we use ROUND for round the number of ** :)
SELECT customer_id, round(SUM(amount),2) FROM payment
GROUP BY customer_id
ORDER BY SUM(amount) DESC
LIMIT 5;




-- /***** HAVING ****\
-- if  I want to explain it in a frankly way 
-- Havinh is a condition for GROUP BY something like WHERE
-- but the difrent beween them is WHERE aplied **BEFORE** GROUP BY applies while
-- WHERE is applies after GROUP BY applied
SELECT rating,AVG(rental_rate)
FROM film
GROUP BY rating
HAVING rating IN ('R', 'G', 'PG');