pragma solidity 0.8.17;

contract MultiSigWallet {
    enum tnxType{Transaction, Modification}
    struct Tnx {
        address destination;
        tnxType tnx_type;
        uint amount;
        mapping(address => bool) validated;
        uint256 validated_count;

    }

    mapping(address => bool) private Owners;
    uint8 private validator_count;
    Tnx[] private transactions;

    event transferCreated(address destination, uint amount, address creator);


    modifier onlyOwners(){
        require(Owners[msg.sender] == true, "You are not owner");
        _;
    }
    modifier NotOwner(){
        require(Owners[msg.sender] == false, "You are OWNER");
        _;
    }

    constructor(address[] memory owner_, uint8 counter_) {
        Owners[msg.sender] = true;
        _add_owner_multiple(owner_);
        require(counter_ < owner_.length, "Owners are less thant validated count");
        validator_count = counter_;
    }

    function _add_owner_multiple(address[] memory new_owners) private {
        for (uint256 i = 0; i < new_owners.length; i++) {
            require(Owners[new_owners[i]] == false, "Duplicate Owner");
            Owners[new_owners[i]] = true;
        }
    }

    // function add_owner(address new_owner) public onlyOwners returns (uint256){
    //     uint256  _tnx = transactions.length;
    //     transactions.push();
    //     newTransaction(_tnx, new_owner, tnxType.Modification);
    //     return _tnx;
    // }
//    function newTransfare(uint256 _tnx, address destination, uint amount) onlyOwners private{
//        Tnx storage transaction =  transactions[_tnx];
//        transaction.validated[msg.sender] = true;
//        transaction.tnx_type = tnxType.Transaction;
//        transaction.destination = destination;
//        transaction.validated_count = 1;
//        transaction.amount = amount;
//    }
//
//    function createTransfare(address destination, uint amount) public payable onlyOwners returns(uint256){
//        uint256 _tnx = transactions.length;
//        transactions.push();
//        newTransfare(_tnx, destination, amount);
//        emit transferCreated(destination, amount, msg.sender);
//        return _tnx;
//    }
}