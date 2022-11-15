const assert = require('assert')
const ganache = require('ganache-cli')
const Web3 = require('web3')
const mocha = require('mocha')
const Lottery = require('../compile.js')
const web3 = new Web3(ganache.provider())


let accounts
let owner
let lottery
let OWNER_ERROR
let NO_MONEY

beforeEach(async () => {
    accounts = await web3.eth.getAccounts()
    owner = accounts[0]
    OWNER_ERROR = "Only owner can pay money"
    NO_MONEY = "You should at least have 0.001 ether"
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

transfare_money = async (lottery) => {
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
    return [total_amount, address_balance, players]

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
        let total_amount, address_balance, players
        [total_amount, address_balance , players]= await transfare_money(lottery)
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
    it('required a minimum amount of lottery', async () => {
        try {
            await lottery.methods.enter().send({
                from: accounts[0],
                value: 10
            })
            assert(false)
        } catch (err) {
            assert(err.toString().includes(NO_MONEY))
        }
    })
    it('only owner can choose winner', async () => {
        try {
            await lottery.methods.SendMoney().call({
                from: accounts[1]
            })
            assert(false)
        } catch (err) {
            assert(err.toString().includes(OWNER_ERROR))

        }
    })
    it('pick winner', async () => {
        let contract_balance = await web3.eth.getBalance(lottery.options.address)
        await transfare_money(lottery)
        let total_account_before = 0
        for (let i = 0; i < 4; i++) {
            total_account_before += await web3.eth.getBalance(accounts[i])
        }
        await lottery.methods.SendMoney().send({
            from: accounts[0]
        })
        let total_account_after = 0
        for (let i = 0; i < 4; i++) {
            total_account_after += await web3.eth.getBalance(accounts[i])
        }
        const contract_balance_after = await web3.eth.getBalance(lottery.options.address) - contract_balance
        assert.strictEqual(
            String(contract_balance_after),
            String(0)
        )
        assert.ok(
            contract_balance + total_account_before <
            total_account_after
        )
    })
})

