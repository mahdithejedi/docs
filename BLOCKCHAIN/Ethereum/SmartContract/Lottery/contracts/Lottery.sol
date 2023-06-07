pragma solidity 0.8.17;

// #__DOCUMENTATION__ #__SOLIDITY__

// #__POINT__ => We CAN have nested arrays in solidity but not in web3 SO
//WE CAN NOT HAVE AN ARRAY OF STRINGs (string itself if an array of string)

//        __MSG__ => is a global variable which is available in whole project

pragma solidity 0.8.17;

contract Lottery{
    address payable public manager;
    address payable[] public members;


    constructor(){
        manager = payable(msg.sender);
    }

    modifier onyOwner(){
        require(manager == msg.sender, 'Only owner can pay money');
        // It's like we put all the function in '_'
        _;
    }

    function enter() public payable {
        require(msg.value > .001 ether, 'You should at least have 0.001 ether');
        members.push(payable(msg.sender));
    }

    function SendMoney() public  onyOwner{
        pickWinner().transfer(address(this).balance);
        // Initate new address with zero default element
        members = new address payable[](0);

    }


    function pickWinner() private view onyOwner returns (address payable){
        uint ran = random() % members.length;
        return members[ran];
    }

    function random() private view returns (uint){
        return uint(
            keccak256(
                abi.encodePacked(
                    block.difficulty, block.timestamp, members
                )
            )
        );
    }
    function getPlayers() public view returns(address payable[] memory){
        return members;
    }

}