SELECT amount FROM payment
WHERE EXTRACT(YEAR from payment_date) > '2006';