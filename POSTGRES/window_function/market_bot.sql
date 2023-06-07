-- We have order table
-- Order table has foreign key to Market
-- -------Order------    foreign key    ----- Market ----
-- |                |   ---------->    |                 |



SELECT o.*, m.name from orders_order o inner join markets_market m on o.market_id = m.id limit 10;

-- +--+---------------------------------+---------------------------------+------------------------------------+------+-----+---------+----------+-----------+---------+-------+---------+----------+-----------------------+-----------------------+------------+---------+--------+
-- |id|updated_at                       |created_at                       |order_id                            |volume|price|direction|status    |_type      |market_id|user_id|gpi_price|gpi_source|limit_dev_delete_bot_id|limit_dev_insert_bot_id|market_maker|pc_bot_id|name    |
-- +--+---------------------------------+---------------------------------+------------------------------------+------+-----+---------+----------+-----------+---------+-------+---------+----------+-----------------------+-----------------------+------------+---------+--------+
-- |27|2022-07-23 13:52:58.646502 +00:00|2022-07-23 13:52:58.600067 +00:00|0778bf10-9824-4ffa-ae1b-27041fcdd134|12    |180  |SELL     |processing|LIMIT_PRICE|4        |2      |12       |NULL      |NULL                   |NULL                   |false       |NULL     |XRP/USDT|
-- |28|2022-07-23 13:52:58.701100 +00:00|2022-07-23 13:52:58.648732 +00:00|9a08085b-355b-4660-9fe4-c0639db3b60c|13    |52   |SELL     |processing|LIMIT_PRICE|4        |2      |12       |NULL      |NULL                   |NULL                   |false       |NULL     |XRP/USDT|
-- |29|2022-07-23 13:52:59.799177 +00:00|2022-07-23 13:52:58.703349 +00:00|4288d259-7a44-4456-a810-49daea58607d|2     |74   |SELL     |processing|LIMIT_PRICE|4        |2      |12       |NULL      |NULL                   |NULL                   |false       |NULL     |XRP/USDT|
-- |30|2022-07-23 13:52:59.841328 +00:00|2022-07-23 13:52:59.802403 +00:00|83695a45-438a-49af-a9f2-b22949b5decb|10    |77   |SELL     |processing|LIMIT_PRICE|4        |2      |12       |NULL      |NULL                   |NULL                   |false       |NULL     |XRP/USDT|
-- |31|2022-07-23 13:52:59.883133 +00:00|2022-07-23 13:52:59.843069 +00:00|e80ff850-9d78-446d-9a4a-6c042a7093e0|3     |129  |SELL     |processing|LIMIT_PRICE|4        |2      |12       |NULL      |NULL                   |NULL                   |false       |NULL     |XRP/USDT|
-- |32|2022-07-23 13:52:59.924678 +00:00|2022-07-23 13:52:59.884908 +00:00|5fd532e9-040d-4a14-abef-51a8c3c6f382|10    |179  |SELL     |processing|LIMIT_PRICE|4        |2      |12       |NULL      |NULL                   |NULL                   |false       |NULL     |XRP/USDT|
-- |33|2022-07-23 13:52:59.967832 +00:00|2022-07-23 13:52:59.926218 +00:00|db76bb2c-f980-4ded-855a-2bde9b5d5623|9     |199  |SELL     |processing|LIMIT_PRICE|4        |2      |12       |NULL      |NULL                   |NULL                   |false       |NULL     |XRP/USDT|
-- |34|2022-07-23 13:53:00.014065 +00:00|2022-07-23 13:52:59.969783 +00:00|99be01e4-c6f4-4300-9b68-0c475f0cac51|7     |19   |SELL     |processing|LIMIT_PRICE|4        |2      |12       |NULL      |NULL                   |NULL                   |false       |NULL     |XRP/USDT|
-- |35|2022-07-23 13:53:00.057964 +00:00|2022-07-23 13:53:00.016012 +00:00|6984b595-2b73-47e4-bdb9-2ae113787a6f|10    |137  |SELL     |processing|LIMIT_PRICE|4        |2      |12       |NULL      |NULL                   |NULL                   |false       |NULL     |XRP/USDT|
-- |36|2022-07-23 13:53:00.104776 +00:00|2022-07-23 13:53:00.059959 +00:00|6ba6c374-854b-4080-a00c-c9040cde598f|20    |33   |SELL     |processing|LIMIT_PRICE|4        |2      |12       |NULL      |NULL                   |NULL                   |false       |NULL     |XRP/USDT|
-- +--+---------------------------------+---------------------------------+------------------------------------+------+-----+---------+----------+-----------+---------+-------+---------+----------+-----------------------+-----------------------+------------+---------+--------+
--


--- we want to get sum price of each market

---
--   +-----+--------+----------------+
--  |price|name    |market_price_avg|
--  +-----+--------+----------------+
--  |180  |XRP/USDT|87.6984126984127|
--  |52   |XRP/USDT|87.6984126984127|
--  |74   |XRP/USDT|87.6984126984127|
--  |77   |XRP/USDT|87.6984126984127|
--  |129  |XRP/USDT|87.6984126984127|
--  |179  |XRP/USDT|87.6984126984127|
--  |199  |XRP/USDT|87.6984126984127|
--  |19   |XRP/USDT|87.6984126984127|
--  |137  |XRP/USDT|87.6984126984127|
--  |33   |XRP/USDT|87.6984126984127|
--  +-----+--------+----------------+

-- General Windows Function syntax
-- window_function(arg1, arg2,..) OVER (
--    [PARTITION BY partition_expression]
--    [ORDER BY sort_expression [ASC | DESC] [NULLS {FIRST | LAST }])


--- Consider this example
select o.price, m.name, o.created_at, o.updated_at,
-- do sum
       AVG (o.price) OVER (
           PARTITION BY o.market_id ORDER BY o.created_at, o.updated_at
           ) AS market_price_avg,
       SUM(o.price) OVER (
           PARTITION BY o.market_id ORDER BY o.created_at, o.updated_at
           ) AS market_price_sum
from orders_order o inner join markets_market m on o.market_id = m.id;

-- Instead of repeating `PARTITION BY o.market_id ORDER BY o.created_at, o.updated_at`
-- repetitively we can use
select o.price, m.name, o.created_at, o.updated_at,
       AVG (o.price) OVER market_price,
       SUM (o.price) OVER market_price
from orders_order o inner join markets_market m on o.market_id = m.id WINDOW
    market_price AS (
        PARTITION BY o.market_id ORDER BY o.created_at, o.updated_at
        );
-- +-----+--------+---------------------------------+---------------------------------+------------------+----+
-- |price|name    |created_at                       |updated_at                       |avg               |sum |
-- +-----+--------+---------------------------------+---------------------------------+------------------+----+
-- |49   |ETH/USDT|2022-07-25 07:21:33.549687 +00:00|2022-07-25 07:21:33.549916 +00:00|49                |49  |
-- |180  |ETH/USDT|2022-07-25 07:21:33.565708 +00:00|2022-07-25 07:21:33.565917 +00:00|114.5             |229 |
-- |86   |ETH/USDT|2022-07-25 07:21:33.593518 +00:00|2022-07-25 07:21:33.603761 +00:00|113.75            |455 |
-- |87   |ETH/USDT|2022-07-25 07:21:33.612408 +00:00|2022-07-25 07:21:33.612580 +00:00|108.4             |542 |
-- |91   |ETH/USDT|2022-07-25 07:21:33.621853 +00:00|2022-07-25 07:21:33.622058 +00:00|105.5             |633 |
-- |139  |BTC/ETH |2022-07-25 07:21:33.558578 +00:00|2022-07-25 07:21:33.558751 +00:00|139               |139 |
-- |27   |BTC/ETH |2022-07-25 07:21:33.614606 +00:00|2022-07-25 07:21:33.614783 +00:00|83                |166 |
-- |71   |BTC/ETH |2022-07-25 07:21:33.619310 +00:00|2022-07-25 07:21:33.619545 +00:00|79                |237 |
-- |27   |BTC/USDT|2022-07-25 07:21:33.552712 +00:00|2022-07-25 07:21:33.552887 +00:00|27                |27  |
-- |9    |BTC/USDT|2022-07-25 07:21:33.563282 +00:00|2022-07-25 07:21:33.563562 +00:00|18                |36  |
-- |115  |BTC/USDT|2022-07-25 07:21:33.577044 +00:00|2022-07-25 07:21:33.577680 +00:00|50.333333333333336|151 |
-- |89   |BTC/USDT|2022-07-25 07:21:33.585753 +00:00|2022-07-25 07:21:33.585990 +00:00|60                |240 |
-- |26   |BTC/USDT|2022-07-25 07:21:33.588357 +00:00|2022-07-25 07:21:33.588560 +00:00|53.2              |266 |
-- |158  |BTC/USDT|2022-07-25 07:21:33.595744 +00:00|2022-07-25 07:21:33.596016 +00:00|70.66666666666667 |424 |
-- |151  |BTC/USDT|2022-07-25 07:21:33.598857 +00:00|2022-07-25 07:21:33.599099 +00:00|82.14285714285714 |575 |
-- |55   |XRP/USDT|2022-07-25 07:21:33.560862 +00:00|2022-07-25 07:21:33.561037 +00:00|55                |55  |
-- |8    |XRP/USDT|2022-07-25 07:21:33.583229 +00:00|2022-07-25 07:21:33.583423 +00:00|31.5              |63  |
-- |42   |XRP/USDT|2022-07-25 07:21:33.591142 +00:00|2022-07-25 07:21:33.591347 +00:00|35                |105 |
-- |196  |XRP/USDT|2022-07-25 07:21:33.601263 +00:00|2022-07-25 07:21:33.601439 +00:00|75.25             |301 |
-- |187  |XRP/USDT|2022-07-25 07:21:33.624680 +00:00|2022-07-25 07:21:33.624882 +00:00|97.6              |488 |
-- |186  |TRX/USDT|2022-07-25 07:20:53.398062 +00:00|2022-07-25 07:20:53.398374 +00:00|186               |186 |
-- |169  |TRX/USDT|2022-07-25 07:21:33.556187 +00:00|2022-07-25 07:21:33.556361 +00:00|177.5             |355 |
-- |141  |TRX/USDT|2022-07-25 07:21:33.568024 +00:00|2022-07-25 07:21:33.568195 +00:00|165.33333333333334|496 |
-- |100  |TRX/USDT|2022-07-25 07:21:33.570251 +00:00|2022-07-25 07:21:33.570421 +00:00|149               |596 |
-- |163  |TRX/USDT|2022-07-25 07:21:33.572469 +00:00|2022-07-25 07:21:33.572639 +00:00|151.8             |759 |
-- |133  |TRX/USDT|2022-07-25 07:21:33.580841 +00:00|2022-07-25 07:21:33.581020 +00:00|148.66666666666666|892 |
-- |21   |TRX/USDT|2022-07-25 07:21:33.605744 +00:00|2022-07-25 07:21:33.605938 +00:00|130.42857142857142|913 |
-- |154  |TRX/USDT|2022-07-25 07:21:33.607961 +00:00|2022-07-25 07:21:33.608151 +00:00|133.375           |1067|
-- |34   |TRX/USDT|2022-07-25 07:21:33.610213 +00:00|2022-07-25 07:21:33.610384 +00:00|122.33333333333333|1101|
-- |3    |TRX/USDT|2022-07-25 07:21:33.616808 +00:00|2022-07-25 07:21:33.616990 +00:00|110.4             |1104|



-- ^^^^^^^^^^^^^^^^^^^^^^^^^^^
-------- EXTRA ----
-- ^^^^^^^^^^^^^^^^^^^^^^^^^^^

-- source https://www.postgresqltutorial.com/postgresql-window-function/

select market_id, created_at, updated_at, price, sum(price) over order_by_created_at as price_sum,
       -- find lowest price by market
       first_value(orders_order.price) over lowest_price_windows as lowest_price,
       -- find the highest price by market
       last_value(orders_order.price) over highest_price_windows as hiehghst_price
from orders_order where market_id = 4 window order_by_created_at as (order by created_at), lowest_price_windows as (partition by market_id order by orders_order.price), highest_price_windows as (
    partition by market_id order by orders_order.price RANGE BETWEEN UNBOUNDED PRECEDING AND UNBOUNDED FOLLOWING
    );




WITH price_sum_ali AS (
    select market_id, created_at, updated_at, price, sum(price) over (order by created_at) as price_sum from orders_order where market_id = 4
) delete from price_sum_ali where price_sum < 108;
