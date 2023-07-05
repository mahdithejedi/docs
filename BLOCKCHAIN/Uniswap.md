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

### Ticks and Ranges

Liquidity providers can provide liquidity in a range between any two ticks (which need not be adjacent)

## More Sources and References

[Medium-A Guide on Uniswap v3 TWAP Oracle](https://tienshaoku.medium.com/a-guide-on-uniswap-v3-twap-oracle-2aa74a4a97c5)
<br />










