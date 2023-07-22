pragma solidity ^0.8.9;

contract A {
    event Called(address thisVar, address MsgSender);

    function checkCaller() public {
        emit Called(address(this), address(msg.sender));
    }
}


contract B {
    event CalledB(address ThisBVar, address MsgBSender);

    function Check() public {
        A a;
        a = new A();
        a.checkCaller();
        emit CalledB(address(this), msg.sender);
    }
}


//contract NoDelegateCall{
//    address private original;
//
//    constructor(){
//        original = address(this);
//    }
//    function NoDelegateCallFunction() internal {
//        require(original == address(this), "delegate call is not allowed!");
//    }
//    modifier NoDelegateAllowed(){
//        NoDelegateCallFunction();
//        _;
//    }
//}

//contract MainDelegate is NoDelegateCall{
contract MainDelegate {
    address public sender;
    uint8 public id;
    address public selfThis;

    event CheckAssignMain(
        address DelSender,
        uint8 id,
        address DelselfThis
    );

    function Assign(uint8 _id) public returns (uint8){
        sender = msg.sender;
        id = _id;
        selfThis = address(this);
        emit CheckAssignMain(sender, id, selfThis);
        return id;
    }
}

contract DelegateMainCall {
    address public sender;
    uint8 public id;
    address public selfThis;

    event CheckAssign(
        address DelSender,
        uint8 id,
        address DelselfThis
    );

    address public main;

    constructor(){
        MainDelegate con = new MainDelegate();
        main = address(con);
    }

    function DelegatceCall() public {
        uint8 _newId = 15;
        (bool success,) = main.delegatecall(abi.encodeWithSignature("Assign(uint8)", _newId));
        emit CheckAssign(msg.sender, id, address(this));
        // what we will see is `CheckAssignMain` this & sender address is exactly the same as CheckAssign,
        // it means like Assign is a DelegateMainCall function rather than DelegateMainCall
        //        [
        //        {
        //        "logIndex": 1,
        //        "blockNumber": 25,
        //        "blockHash": "0xa33694ccb42d0850b1b718e9b1691f0072430d2062093851b350046267c87685",
        //    "transactionHash": "0x2823d967078eea9bcf0ed419b82a66f5b31736e567d210012a202d44aaf10e9c",
        //    "transactionIndex": 0,
        //        "address": "0xD4Fc541236927E2EAf8F27606bD7309C1Fc2cbee",
        //        "data": "0x0000000000000000000000005b38da6a701c568545dcfcb03fcb875f56beddc4000000000000000000000000000000000000000000000000000000000000000f000000000000000000000000d4fc541236927e2eaf8f27606bd7309c1fc2cbee",
        //        "topics": [
        //    "0x90e395f2a158073afaa940071a399f96946a4cbfeb5a6a264261577974b0688c"
        //    ],
        //    "id": "log_5c2b717b"
        //        },
        //        {
        //        "logIndex": 1,
        //        "blockNumber": 25,
        //        "blockHash": "0xa33694ccb42d0850b1b718e9b1691f0072430d2062093851b350046267c87685",
        //    "transactionHash": "0x2823d967078eea9bcf0ed419b82a66f5b31736e567d210012a202d44aaf10e9c",
        //    "transactionIndex": 0,
        //        "address": "0xD4Fc541236927E2EAf8F27606bD7309C1Fc2cbee",
        //        "data": "0x0000000000000000000000005b38da6a701c568545dcfcb03fcb875f56beddc40000000000000000000000000000000000000000000000000000000000000000000000000000000000000000d4fc541236927e2eaf8f27606bd7309c1fc2cbee",
        //        "topics": [
        //    "0x4de9627489cd3996ff101b1ce0fcce6bcd07c3b5646f012a192cb230ee4099c5"
        //    ],
        //    "id": "log_5c2b717b"
        //        }
        //    ]
    require(success, "not successfull");

    }

}