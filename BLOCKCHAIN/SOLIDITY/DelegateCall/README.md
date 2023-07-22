# What is delegate call?

The `delegatecall` function in Solidity is similar to `call`, but with one crucial difference. When you
use `delegatecall`
to execute a function from another contract, the called function runs in the context of the calling contract's storage
and state variables and caller contract will send its own storage to called contract. **So the order of variable
declaration is matter**. Because storage should be the same. we can extend memory by this way. Because it will directly
modify caller contract. Consider Called contract has 5 state variable and caller 4. When a function called Caller
contract will dedicate extra slot for that etra variable which it exists in called but not in caller.

### Sources

[medium: Solidity delegatecall usage and pitfalls](https://medium.com/@jeremythen16/solidity-delegatecall-usage-and-pitfalls-5c37eaa5bd5d)