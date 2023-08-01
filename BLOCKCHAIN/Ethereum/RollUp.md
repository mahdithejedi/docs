# What is RollUp?

Rollups are one such scaling solution. They involve rolling up (hence the name) collections of transactions. The final,
rolled-up transaction is presented to the Ethereum blockchain as a single transaction.
<br />
Rollups cuts costs: the cost of an Ethereum transaction, plus the small cost of rolling up batches of transactions, is
split among users. They also speed things up: the rollup is very quick to perform and the Ethereum blockchain needs only
to process a single transaction rather than many.

## 2 Types of RollUps

* **Optimistic rollups** : make the assumption that all of this rolled-up data is valid, and that nobody is trying to fool the
  blockchain by hiding spurious transactions within rollups. The idea is that by assuming validity, things speed up. To
  protect against fraudulent transactions, optimistic rollup protocols allow people to contest bunk trades. The
  fraudulent transaction is submitted directly on the Ethereum network to check if it’s legit, and to settle the
  dispute. Both parties have ETH staked and would lose money if they are wrong or lie.
* **Zero-knowledge**: rollups (also referred to as zk-rollups) work very differently. They rely on a piece of
  cryptography called a zero-knowledge proof, which allows someone to mathematically prove that a statement is true (
  that someone is, say, Greek) without disclosing additional information about that statement (like a passport)

### Sources
[decrypt.co ETH Scaling solution](https://decrypt.co/resources/what-are-ethereum-rollups-scaling-solution-cut-transaction-costs)
<br />
