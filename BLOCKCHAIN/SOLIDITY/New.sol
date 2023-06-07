// SPDX-License-Identifier: BUSL-1.1

// SPDX-License-Identifier: MIT
pragma solidity =0.7.6;
import "https://github.com/Uniswap/v3-core/blob/main/contracts/UniswapV3Pool.sol";


contract Deployer is IUniswapV3PoolDeployer{
    struct Parameters {
        address factory;
        address token0;
        address token1;
        uint24 fee;
        int24 tickSpacing;
    }
    Parameters public override parameters;
    function deploy() public  returns (address) {
        // Here in this example New UniswapV3Pool is created and deployed in a same address
        parameters = Parameters({factory: address(0x0A098Eda01Ce92ff4A4CCb7A4fFFb5A43EBC70DC), token0: address(0x006b175474e89094c44da98b954eedeac495271d0f), token1: address(0x006b175474e89094c44da98b954eedeac495271d0f), fee: 1, tickSpacing: 1});
        address _address = address(new UniswapV3Pool{salt: keccak256(abi.encode(address(0x006b175474e89094c44da98b954eedeac495271d0f), address(0x006b175474e89094c44da98b954eedeac495271d0f), uint24(1)))}());
        delete parameters;
        return _address;
    }
}