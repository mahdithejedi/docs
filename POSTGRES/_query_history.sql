-- Select Active Bot Orders
select order_id from orders_order inner join markets_market mm on mm.id = orders_order.market_id  where status = 'canceled' and cancel_by = 'delete_bot' and limit_dev_insert_bot_id in (
	    select id from bots_limitdevinsertbot limit_bot where limit_bot.created_at >= now() - interval '1 weeks'
	    ) and mm.name = 'FTM/USDT';


-- Select Total Market Value

WITH price_sum_ali AS (
	    select market_id, created_at, updated_at, order_id, status ,price ,sum(price) over (order by created_at, updated_at) as price_sum from orders_order where market_id = 4 and direction = 'SELL' and status = 'processing'
) select order_id from price_sum_ali where price_sum < 101.77200000000002;




