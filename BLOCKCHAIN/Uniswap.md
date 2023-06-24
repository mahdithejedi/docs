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


### Uniswap V3 Oracle
#### SqrtRation
* IN uniswap V2 we calculate _TWAP_, in V3 the return of TWAP is _tick_. what it means is that the answer of the
<br />
**`(TickCommulative2 - TickCommulative1) / (Timestamp2 - Timestamp1) = tick-TWAP`**

* On the other-hand we have this formula in which it calculates tick:
**`(1.0001)^tick=price`**
* This is the reason behind the function of [`getSqrtRatioAtTick`](https://github.com/Uniswap/v3-core/blob/main/contracts/libraries/TickMath.sol#L23C14-L23C32)


## More Sources and References
[Medium-A Guide on Uniswap v3 TWAP Oracle](https://tienshaoku.medium.com/a-guide-on-uniswap-v3-twap-oracle-2aa74a4a97c5)
<br />










