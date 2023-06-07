-- **** BETWEEN STATEMENT ***
-- value BETWEEN low AND high == value >=low AND value <= high
-- value NOT BETWEEN low AND hight  -> is value out of range..?
-- Example ->

SELECT * FROM payment 
WHERE amount BETWEEN 8 AND 12; -- or WHERE amount NOT BETWEEN 10 AND 12;


-- **** IN STATEMENT ****
-- value <opt:NOT:> IN (value1,value,....)
-- Exmaple ->

SELECT payment_id, customer_id FROM payment
WHERE amount IN (8.99,9.99);


-- ***** <opt:NOT> LIKE STATEMENT *****
-- check if just a part of is match !! What I've been said 
-- SmallExmaple you think you want a customer and it's name start with joe
-- You can say WHERE first_name LIKE 'Joe%' /*** This technick called Paterned pactching ***\
-- Note : '%' match any seuqnce of characters -- '_' match any single character
-- Exmaple ->
SELECT first_name ,last_name 
FROM customer
WHERE first_name LIKE '_a%'
ORDER BY first_name;

-- POINT: LIKE is CASEsensitive so for 
-- for checking nonsensitivity we can use ILIKE