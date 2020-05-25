SELECT SUM(amount) as t_exp, EXTRACT(MONTH from payment_date) AS pay_months
FROM payment
GROUP BY pay_months -- Here we group the all the rows in payment by month and the calculate the sum
ORDER BY t_exp DESC

/*
POINT: you think that the first think is GROUP BY
first the pyment_date will group by monthes then after that 
the SUM of pay_month or the amount of all the same monthes will be execute
*/
