pragma solidity 0.8.17;

// #__DOCUMENTATION__ #__SOLIDITY__
// This code is publish into etherium blockchain
// so everybody will see it so we shouldn't put secret bussiness logic in here
//-----
// gasPrice -> amount of wei we are pay for a gas
// gasLimit -> MAX amount of gas we are willing to pay
// !!! gasLimit > Amount of gas our transaction need (ramainig gas will return back)

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
        //! Every data that modified in etherium blockchain should be in a form of TRANSACTION(
        //! this transaction should be mine and proved + cost money + take time to execute + return transaction hash)
        //
        message = newMessage;
    }

    function getMessage() public view returns (string memory){// 'view' = 'constant' means this function returns data and NOT modify contract's data
        //        'pure' -> Function will not modify or even read the contract's data
        //        'payable' -> When s1 call this function they might send ether along
        //! This function can not modify data + run instantly + free to do
        return message;
    }
}
