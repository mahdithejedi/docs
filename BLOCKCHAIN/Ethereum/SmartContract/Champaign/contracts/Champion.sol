// SPDX-License-Identifier: UNLICENSED
pragma solidity ^0.8.9;

// Uncomment this line to use console.log
// import "hardhat/console.sol";

contract Champion{

    address owner;

    uint8 agreement;


    modifier ownerOnly(){
        require(msg.sender == owner, 'Only owner can change');
        _;
    }
    modifier agreementCheck(uint8 _agreement){
        require(_agreement < 100, 'agreement can be between 0 upto 100');
        _;
    }

    constructor(uint8 _agreement) agreementCheck(_agreement){
        agreement = _agreement;
        owner = msg.sender;
    }


    function set_agreement(uint8 new_agreement) public ownerOnly agreementCheck(new_agreement){
        require(new_agreement < 100, 'new_agreement can be between 0 upto 100');
        agreement = new_agreement;
    }
    function get_agreement() public view returns (uint8){
        return agreement;
    }

}
