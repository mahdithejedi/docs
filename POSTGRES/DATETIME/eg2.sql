SELECT SUM(amount)as expenses, EXTRACT(month from payment_date) AS monthes
FROM payment
GROUP BY monthes
ORDER BY expenses DESC;