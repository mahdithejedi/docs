const mocha = require('mocha')
const assert = require('assert')
const Web3 = require('web3')
const ganache = require('ganache-cli')
const hardhat = require("hardhat");
const web3 = new Web3(ganache.provider())

let MultiSigWallet
let creator
let signers
let contract




let DUPLICATE_ERROR = "Duplicate Owner"


async function deploy(){
    const MultiSigWalletFactory = await hardhat.ethers.getContractFactory("MultiSigWallet");
    console.log("Factory", MultiSigWalletFactory);
    const MultiSigWallet = await MultiSigWalletFactory.deploy()
    console.log("Wallet", MultiSigWallet);
    await MultiSigWallet.deployed()
    return MultiSigWallet
}

beforeEach(async () => {
    MultiSigWallet = await deploy();
    creator = web3.eth.accounts[0];
    signers = {
        'first': web3.eth.accounts[1],
        'second': web3.eth.accounts[2],
        'third': web3.eth.accounts
    }
    contract = await new web3.eth.Contract(MultiSigWallet.abi).deploy({
        'owner_': [signers.first, signers.second, signers.third],
        'counter_': 3
    })
    console.log("Address is", MultiSigWallet)

})

describe("Constructor", async () => {
    it("Duplicate error", async () => {
        try {
            await new web3.eth.Contract(abi).deploy({
                'owner_': [signers.first, signers.second, signers.second],
                'counter_': 3
            })
        } catch (error) {
            assert.strictEqual(error, DUPLICATE_ERROR)
        }
    })
})