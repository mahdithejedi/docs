-- Select Active Bot Orders
select order_id from orders_order inner join markets_market mm on mm.id = orders_order.market_id  where status = 'canceled' and cancel_by = 'delete_bot' and limit_dev_insert_bot_id in (
	    select id from bots_limitdevinsertbot limit_bot where limit_bot.created_at >= now() - interval '1 weeks'
	    ) and mm.name = 'FTM/USDT';


-- Select Total Market Value

WITH price_sum_ali AS (
	    select market_id, created_at, updated_at, order_id, status ,price ,sum(price) over (order by created_at, updated_at) as price_sum from orders_order where market_id = 4 and direction = 'SELL' and status = 'processing'
) select order_id from price_sum_ali where price_sum < 101.77200000000002;
-- Group by 1 hours interval
select count(*),
       date_trunc('hour', o.created_at)  AS ten_min_timestamp
       from orders_order o inner join markets_market mm on mm.id = o.market_id where
    cancel_by = 'limit_bot' group by ten_min_timestamp order by ten_min_timestamp desc

-- Groupy by 10 minutes interval
select count(*),
       date_trunc('hour', o.created_at) +
       (((date_part('minute', o.created_at)::integer / 10::integer) * 10::integer)
	           || ' minutes')::interval AS ten_min_timestamp
	       from orders_order o inner join markets_market mm on mm.id = o.market_id where
	    cancel_by = 'limit_bot' group by ten_min_timestamp

-- Convert TimeZone in 10 minutes interval
select count(*),
       date_trunc('hour', o.created_at) at time zone 'Iran' +
       (((date_part('minute', o.created_at)::integer / 10::integer) * 10::integer)
	           || ' minutes')::interval AS ten_min_timestamp
	       from orders_order o inner join markets_market mm on mm.id = o.market_id where
	    cancel_by = 'limit_bot' group by ten_min_timestamp order by ten_min_timestamp;
-- Last 24 hours order bot limit
select count(*) from orders_order o where o.limit_dev_insert_bot_id = 68 and o.created_at > now() - interval '24 hours';
