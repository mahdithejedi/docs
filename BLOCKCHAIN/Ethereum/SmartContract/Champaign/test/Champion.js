require('@nomiclabs/hardhat-truffle5')
const assert = require('assert')

const Champion = artifacts.require('Champion')

let contract
const AGREEMENT = 60

beforeEach(async () => {
    contract = await Champion.newValueSpecifiedUnits(AGREEMENT)
})


describe("Champion tests", async () => {
    assert.strictEqual(contract.get_agreement(), AGREEMENT)
})
