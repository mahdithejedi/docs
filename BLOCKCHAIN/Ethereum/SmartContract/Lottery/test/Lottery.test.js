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

})

ether_to_wie = (ether) => {
    return web3.utils.toWei(
        ether, 'ether'
    )
}


describe("Lottery", () => {
    it("Deploy contract", () => {
        assert.ok(lottery.options.address)
    })
    it("check owner address", async () => {
        const contract_owner = await lottery.methods.manager().call();
        assert.strictEqual(
            owner,
            contract_owner
            ,
            'lottery owner is Wrong'
        )
    })
    it('check transactions count', async () => {
        let total_amount = 0
        let players = []


        await lottery.methods.enter().send({
            from: accounts[1],
            value: ether_to_wie('0.2')
        })
        total_amount += 0.2
        players.push(accounts[1])

        await lottery.methods.enter().send({
            from: accounts[2],
            value: ether_to_wie('0.3')
        })
        total_amount += 0.3
        players.push(accounts[2])


        await lottery.methods.enter().send({
            from: accounts[3],
            value: ether_to_wie('0.5')
        })
        total_amount += 0.5
        players.push(accounts[3])
        let address_balance = await web3.eth.getBalance(
            lottery.options.address
        )
        //    check total amount
        assert.strictEqual(
            web3.utils.toWei(String(total_amount), 'ether'),
            address_balance
        )

        // check players array
        let contract_players = await lottery.methods.getPlayers().call();
        assert.deepStrictEqual(
            players,
            contract_players
        )


    })

    it('Enter lottery', async () => {

    })
})

