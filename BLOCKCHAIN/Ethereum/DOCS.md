# Ethereum
[Ethereum Official Documentation](https://ethereum.org/en/developers/docs/intro-to-ethereum/)
<br />
[Learn ethereum Visualize](https://eth.build/)


### Ethereum Fee
Ethereum doesn't split gas between validators instead it will first _burn_ the fee in the
given transaction and then _mint_ the gas to the validators 
<small>[source](https://ethereum.org/en/developers/docs/transactions/#on-gas) </small>

_Let's say Jordan has to pay Taylor 1 ETH. In the transaction, the gas limit is 21,000 units and the base fee is 10 gwei. Jordan includes a tip of 2 gwei.

The total fee would now be: units of gas used * (base fee + priority fee) where the base fee is a value set by the protocol and the priority fee is a value set by the user as a tip to the validator.

i.e 21,000 * (10 + 2) = 252,000 gwei or 0.000252 ETH.

When Jordan sends the money, 1.000252 ETH will be deducted from Jordan's account. Taylor will be credited 1.0000 ETH. Validator receives the tip of 0.000042 ETH. Base fee of 0.00021 ETH is burned.

Additionally, Jordan can also set a max fee (maxFeePerGas) for the transaction. The difference between the max fee and the actual fee is refunded to Jordan, i.e. refund = max fee - (base fee + priority fee). Jordan can set a maximum amount to pay for the transaction to execute and not worry about overpaying "beyond" the base fee when the transaction is executed._
<small> [source](https://ethereum.org/en/developers/docs/gas/#post-london) </small>

### Ethereum Transaction Type
    * One which result in message calls
    * One  which result in contract creation


