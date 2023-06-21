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





