const assert = require('assert')
const ganache = require('ganache-cli')
const Web3 = require('web3')
const mocha = require('mocha')
const Inbox = require('../compile.js')
const web3 = new Web3(ganache.provider())


let accounts
let inbox
const INITIAL_MESSAGE = 'Hi there!'

beforeEach(async () => {
//    Get list of accounts
//     web3.eth.getAccounts().then(account => {
//         console.log(account)
//     })
    accounts = await web3.eth.getAccounts()
//    Deploy
    inbox = await new web3.eth.Contract(Inbox.interface).deploy({
        data: Inbox.bytecode.object,
        arguments: [INITIAL_MESSAGE]
    })
        .send({from: accounts[0], gas: 1000000})

})

describe("Inbox", () => {
    it("Deploy contract", () => {
        assert.ok(inbox.options.address)
    })
    it("had a default message", async () => {
        //inbox.methods => public methods
        // message() => set arguments of function
        // call () => customize transaction(for example who is going to pay transaction or gas price or ..)
        const message = await inbox.methods.message().call();
        assert.strictEqual(message, INITIAL_MESSAGE)
    })
    it("Test change message", async () => {
        const TEST_MSG = 'Test'
        const transaction_data = await inbox.methods.setMessage('Test').send({
            from: accounts[0]
        })
        console.log('---------------------transaction data ------------------')
        console.log(transaction_data)
        const message = await inbox.methods.message().call();
        assert.strictEqual(message, TEST_MSG)

    })
    // it("Check inbox", () => {
    //     console.log(inbox)
    // })
})

