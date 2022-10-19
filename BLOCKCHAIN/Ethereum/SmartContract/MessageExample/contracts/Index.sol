pragma solidity 0.8.17;

contract Index {
    //  because it's public
    //it's available to all users in the world
    string public message;

    constructor(string memory initialMessage) public {
        //    Constructor which will automatically called on time
        //     when the contra\ct is created
        message = initialMessage;
    }

    function setMessage(string memory newMessage) public {// 'public' or 'private'
        message = newMessage;
    }

    function getMessage() public view returns (string memory){// 'view' = 'constant' means this function returns data and NOT modify contract's data
        //        'pure' -> Function will not modify or even read the contract's data
        //        'payable' -> When s1 call this function they might send ether along
        return message;
    }
}
