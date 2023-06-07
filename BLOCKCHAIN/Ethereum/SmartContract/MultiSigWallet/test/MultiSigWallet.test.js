const mocha = require('mocha')
const assert = require('assert')
const Web3 = require('web3')
// const ganache = require('ganache-cli')
const hardhat = require("hardhat");
const ethers = hardhat.ethers
const chai = require('chai')
const expect = chai.expect
// const web3 = new Web3(ganache.provider())

// let MultiSigWallet
let creator
let signers
let contract


const DUPLICATE_ERROR = "Duplicate Owner"

getRandom = (max) => {
    return Math.floor(Math.random() * max)
}
getNewWallet = () => {
    return ethers.Wallet.createRandom(getRandom(12000))
}

async function deploy(owners, counter) {
    const MultiSigWalletFactory = await hardhat.ethers.getContractFactory("MultiSigWallet");
    const MultiSigWallet = await MultiSigWalletFactory.deploy(owners, counter)
    await MultiSigWallet.deployed()
    return MultiSigWallet
}

beforeEach(async () => {
    creator = getNewWallet();
    signers = {
        'first': getNewWallet(),
        'second': getNewWallet(),
        'third': getNewWallet()
    }
    contract = await deploy(
        [signers.first.address, signers.second.address, signers.third.address],
        3
    )
    console.log("Contract deployed -> Address is", contract.address)

})

describe("Constructor", async () => {
    it("Duplicate error", async () => {
        await expect(await deploy(
            [signers.first.address, signers.second.address, signers.second.address],
            3
        )).to.revertedWith(DUPLICATE_ERROR)
        // try {
        //     await deploy(
        //         [signers.first.address, signers.second.address, signers.second.address],
        //         3
        //     )
        // } catch (error) {
        //     assert.strictEqual(error, DUPLICATE_ERROR)
        // }
    })
})