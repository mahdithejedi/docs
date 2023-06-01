// Source: https://www.geeksforgeeks.org/storage-vs-memory-in-solidity/
// SPDX-License-Identifier: MIT
pragma solidity ^0.8.7;

contract MemoryVsStorage{
    int8[] public memoryArray;
    int8[] public StorageArray;
    function MemoryTest() public returns (int8[] memory){
            memoryArray.push(1);
            memoryArray.push(2);
            int8[] memory internalMemoryArray = memoryArray;
            internalMemoryArray[0] = 0;
            return memoryArray; // Result it: [1, 2]
    }
    function StorageTest() public returns (int8[] memory){
            StorageArray.push(1);
            StorageArray.push(2);
            int8[] storage internaStroageArray = StorageArray;
            internaStroageArray[0] = 0;
            return StorageArray; // Result it: [0, 2]
    }

}