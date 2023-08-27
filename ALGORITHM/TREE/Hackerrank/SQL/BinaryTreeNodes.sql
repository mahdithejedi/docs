-- __URL__ -> https://www.hackerrank.com/challenges/binary-search-tree-1/problem?isFullScreen=true
-- My answer
SELECT SubBST.NBST, (
    CASE WHEN P IS NULL THEN 'Root' ELSE (
        CASE WHEN (SELECT COUNT(*) FROM BTS WHERE BTS.P = NBST) = 0 THEN 'Leaf' ELSE 'Inner' END
        ) END
    ) FROM (
               SELECT N AS NBST, P FROM  BTS
           ) AS SubBST order by SubBST.NBST;

-- Other solutions
-- with data as(
--     SELECT N, LEVEL as lvl
--     FROM bst
--              CONNECT BY PRIOR N = P
--     Start with P is null),
--      maxL as (
--          select max(data.lvl) as m from data)
-- select dt.N, CASE WHEN dt.lvl =1 THEN 'Root'
--                   WHEN dt.lvl=maxL.m THEN 'Leaf'
--                   ELSE  'Inner' END
-- from data dt, maxL
-- order by dt.N;
-- Document `CONNECT BY` -> https://docs.oracle.com/en/database/oracle/oracle-database/19/sqlrf/Hierarchical-Queries.html#GUID-E3D35EF7-33C3-4D88-81B3-00030C47AE56