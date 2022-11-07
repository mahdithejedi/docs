pragma solidity 0.8.17;

// #__DOCUMENTATION__ #__SOLIDITY__

// #__POINT__ => We CAN have nested arrays in solidity but not in web3 SO
//WE CAN NOT HAVE AN ARRAY OF STRINGs (string itself if an array of string)

//        __MSG__ => is a global variable which is available in whole project

contract Lottery{
    address public manager;
    address[] public members;


    constructor() public{
        manager = msg.sender;
    }
    function enter() public payable {
        require(msg.value > .001 ether, 'You should at least have 0.001 ether');
        members.push(msg.sender);
    }

}