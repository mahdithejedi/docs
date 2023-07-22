pragma solidity ^0.8.9;

contract WalletManager{
    address private creator;
    bool private _locked;
    address[] private wallets;
    event walletCreated(address new_address);
    event Received(uint256 value);

    constructor(){
        creator = msg.sender;
        _locked = false;
    }
    modifier owner(){
        require(msg.sender == creator, "Only creator can call this function");
        _;
    }

    function newWallet() public owner{
        Wallet _wallet = new Wallet(payable(address(this)));
        wallets.push(address(_wallet));
        emit walletCreated(address(_wallet));
    }

    receive() payable external{
        emit Received(msg.value);
    }
}

contract Wallet{
    address payable private owner;

    constructor(address payable _owner){
        owner = payable(_owner);
    }
    receive() external payable{
        owner.transfer(msg.value);
    }
}