const assert = require('assert')
const ganache = require('ganache-cli')
const Web3 = require('web3')
const mocha = require('mocha')
const Lottery = require('../compile.js')
const web3 = new Web3(ganache.provider())


let accounts
let owner
let lottery


beforeEach(async () => {
    accounts = await web3.eth.getAccounts()
    owner = accounts[0]
//    Deploy
    lottery = await new web3.eth.Contract(Lottery.interface).deploy({
        data: Lottery.bytecode.object,
    })
        .send({from: owner, gas: 1000000})
    console.log(lottery)

})

transfer = async (account, finney) =>{
    return web3.eth.sendTransaction({
        from: account,
        to: owner,
        value: web3.utils.toWei(web3.utils.toBN(finney), 'finney')
    });

}

describe("Lottery", () => {
    it("Deploy contract", () => {
        assert.ok(lottery.options.address)
    })
    it("check owner address", async () => {
        const contract_owner = await lottery.methods.manager().call();
        console.log(contract_owner);
        assert.strictEqual(
            owner,
            contract_owner
            ,
            'lottery owner is Wrong'
        )
    })
    it('send transaction', async() => {
        console.log(
            await transfer(
                accounts[1], 10
            )
        )
    })
})

