# Uniswap

[Uniswap Version 3 youtube Architecture Explained](https://youtu.be/Ehm-OYBmlPM)
<br />
[Uniswap v3 Features Explained in Depth](https://medium.com/taipei-ethereum-meetup/uniswap-v3-features-explained-in-depth-178cfe45f223)

## What is slippage?

[Medium Article](https://dexenetwork.medium.com/what-is-slippage-and-why-does-it-matter-uniswap-example-43e32d712651)
<br />

[//]: # (__UNISWAP__)

[//]: # (__TWAP_ORACLE__)

[//]: # (__TWAP_ORACLE_UNISWAP__)


[//]: # (__UNISWAP__)

[//]: # (__UNISWAP_PRICING_FUNCTION__)

### Uniswap V2 Oracle

Uniswap V2 includes several improvements for supporting manipulation-resistant public price
feeds. <small> [source](https://docs.uniswap.org/contracts/v2/concepts/core-concepts/oracles) </small>

* First, every pair measures (but does not store) the market price at the beginning of each block, before any trades
  take place
* Second, Uniswap V2 adds this end-of-block price to a single cumulative-price variable in the core contract weighted by
  the amount of time this price existed. This variable represents a sum of the Uniswap price for every second in the
  entire history of the contract.

![TWAP formula ](https://docs.uniswap.org/assets/images/v2_twap-fdc82ab82856196510db6b421cce9204.png)

[TWAP Contract Exmaple](https://github.com/Uniswap/v2-periphery/blob/master/contracts/examples/ExampleOracleSimple.sol)

# Uniswap V3

## What is NEW in V3?

* **Concentrated Liquidity**: Liquidity providers (LPs) are given the ability to concentrate their liquidity by
  “bounding" it within an arbitrary price range. This improves the pool’s capital efficiency and allows LPs to
  approximate their preferred reserves curve, while still being efficiently aggregated with the rest of the pool
* **Flexible Fees**: The swap fee is no longer locked at 0.30%. Rather, the fee tier for each pool (of which there can
  be multiple per asset pair) is set on initialization (Section 3.1). The initially supported fee tiers are 0.05%,
  0.30%, and 1%.
* **Protocol Fee Governance**: UNI governance has more flexibility in setting the fraction of swap fees collected by the
  protocol
* **Improved Price Oracle**: Uniswap v3 provides a way for users to query recent price accumulator values, thus avoiding
  the need to checkpoint the accumulator value at the exact beginning and end of the period for which a TWAP is being
  measured.
* **Liquidity Oracle**: The contracts expose a time-weighted average liquidity oracle
  <br />
  <small> [From Uniswap V3 Whitepaper](https://uniswap.org/whitepaper-v3.pdf) </small>

### Uniswap V3  ARCHITECTURAL CHANGES

* *Non-Fungible Liquidity*:
    * *Non-Compounding Fees*: due to the non-fungible nature of positions, this is no longer possible. Instead, fee
      earnings are stored separately and held as the tokens in which the fees are paid

## Uniswap V3 Fee

Fees collected separately from the pool and must be manually redeemed when the owner wishes to collect their
fees. <small> [source](https://docs.uniswap.org/concepts/protocol/fees) </small>
<br/>
As you read
the [code](https://github.com/Uniswap/v3-core/blob/d8b1c635c275d2a9450bd6a78f3fa2484fef73eb/contracts/UniswapV3Pool.sol#L635)
you will find out that fees are calculated separately in each tick(or swap state). First Fee is _0_ and then fee added
with a correlation with
`feeGrowthGlobal0X128` and
`feeGrowthGlobal1X128` in each tick (swap state).

### Uniswap V3 Oracle

#### SqrtRation

* IN uniswap V2 we calculate _TWAP_, in V3 the return of TWAP is _tick_. what it means is that the answer of the
  <br />
  **`(TickCommulative2 - TickCommulative1) / (Timestamp2 - Timestamp1) = tick-TWAP`**

* On the other-hand we have this formula in which it calculates tick:
  **`(1.0001)^tick=price`**
* This is the reason behind the function
  of [`getSqrtRatioAtTick`](https://github.com/Uniswap/v3-core/blob/main/contracts/libraries/TickMath.sol#L23C14-L23C32)

# [Uniswap V3 Book](https://uniswapv3book.com/)

The Uniswap V3 books is followed in the [`Uniswap V3 Book`](./Uniswap%20V3%20Book) Directory.

<iframe src="https://www.desmos.com/calculator/m1dlhizkr5?embed" width="500" height="500" style="border: 1px solid #ccc" frameborder=0></iframe>

1. The purple line is the curve, the axes are the reserves of a pool (notice that they’re equal at the start price).
2. tart price is 1.
3. We’re selling 200 of token 0. If we use only the start price, we expect to get 200 of token 1. **However, the
  execution price is 0.666, so we get only 133.333 of token 1!**

The law of supply and demand tells us that when demand is high (and supply is constant) the price is also high. And when
demand is low, the price is also lower. This is how markets work. And, magically, the constant product function
implements this mechanism! Demand is defined by the amount you want to buy, and supply is the pool reserves. When you
want to buy a big amount relative to pool reserves the price is higher than when you want to buy a smaller amount. Such
a simple formula guarantees such a powerful mechanism!

### Slope of actual price

1. Before a trade, there’s a spot price. It’s equal to the relation of reserves, y/x or x/y depending on the
  direction of the trade. This price is also the slope of the tangent line at the starting point.
2. After a trade, there’s a new spot price, at a different point on the curve. And it’s the slope of the tangent line at this new point.
3. The actual price of the trade is the slope of the line connecting the two points!

### Position and Tick
Position: Positions represent an owner address' liquidity between a lower and upper tick boundary 


Liquidity providers can provide liquidity in a range between any two ticks (which need not be adjacent)

## More Sources and References

[Medium-A Guide on Uniswap v3 TWAP Oracle](https://tienshaoku.medium.com/a-guide-on-uniswap-v3-twap-oracle-2aa74a4a97c5)
<br />










