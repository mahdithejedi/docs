const path = require('path');
const fs = require('fs');
const solc = require('solc');

const inboxPath = path.resolve(__dirname, 'contracts', 'Inbox.sol');
const source = fs.readFileSync(inboxPath, 'utf8');

var input = {
    language: 'Solidity',
    sources: {
        'Inbox.sol': {
            content: source
        }
    },
    settings: {
        outputSelection: {
            '*': {
                '*': ['*']
            }
        }
    }
};

_compiled = solc.compile(JSON.stringify(input))
output = JSON.parse(_compiled)
const bytecode = output.contracts['Inbox.sol'].Index.evm['bytecode']
exports.bytecode = bytecode
const _interface = output.contracts['Inbox.sol'].Index.abi
exports.interface = _interface