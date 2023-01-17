pragma solidity 0.8.17;

contract MultiSigWallet {
    enum tnxType{Transaction, Modification}
    struct Tnx {
        address destination;
        tnxType tnx_type;
        mapping(address => bool) validated;
        uint8 validated_count;

    }

    mapping(address => bool) private Owners;
    uint8 private validator_count;
    Tnx[] private transactions;

    modifier onlyOwners(){
        require(Owners[msg.sender] == true, "You are not owner");
        _;
    }
    modifier NotOwner(){
        require(Owners[msg.sender] == false, "You are OWNER");
        _;
    }

    constructor(address[] memory owner_, uint8 counter_) {
        _add_owner_multiple(owner_);
        validator_count = counter_;
    }

    function _add_owner_multiple(address[] memory new_owners) private {
        for (uint8 i = 0; i < new_owners.length; i++) {
            Owners[new_owners[i]] = true;
        }
    }

    function add_owner(address new_owner) public onlyOwners returns (uint256){
        uint256  _tnx = transactions.length;
        transactions.push();
        newTransaction(_tnx, new_owner, tnxType.Modification);
        return _tnx;
    }
    function newTransaction(uint256 _tnx, address destination, tnxType tnx_type) private{
        Tnx storage transaction =  transactions[_tnx];
        transaction.validated[msg.sender] = true;
        transaction.tnx_type = tnx_type;
        transaction.destination = destination;
        transaction.validated_count = 1;
    }
}