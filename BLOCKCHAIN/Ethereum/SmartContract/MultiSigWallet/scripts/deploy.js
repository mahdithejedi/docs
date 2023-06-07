const hardhat = require('hardhat')

async function main(){
    const MultiSigWalletFactory = await hardhat.ethers.getContractFactory("MultiSigWallet");
    console.log("Factory", MultiSigWalletFactory);
    const MultiSigWallet = await MultiSigWalletFactory.deploy()
    console.log("Wallet", MultiSigWallet);
    await MultiSigWallet.deployed()
    return MultiSigWallet
}

export const MultiSigWallet = await main();
console.log("Deployed at", MultiSigWallet)