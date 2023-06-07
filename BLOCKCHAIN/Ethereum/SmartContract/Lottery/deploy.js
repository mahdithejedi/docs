const HDWalletProvider = require('truffle-hdwallet-provider')
const Web3 = require('web3')
const Lottery = require('./compile.js')
const assert = require("assert");

const provider = new HDWalletProvider('balance spring token ghost broom reject diagram neck banner code siege require', 'https://goerli.infura.io/v3/d206dfde28e34d009286a72673306fca')
const web3 = new Web3(provider)

const deploy = async () => {
    const accounts = await web3.eth.getAccounts()
    console.log('Attempting to deploy from account', accounts[0])

    const result = await new web3.eth.Contract(Lottery.interface).deploy({
        data: Lottery.bytecode.object,
    }).send({from: accounts[0]})

    console.log('Contract deployed to ', result.options.address)
    console.log('-----------------')
    console.log('result is', result)
}

deploy()