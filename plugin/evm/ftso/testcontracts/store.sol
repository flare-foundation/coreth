// SPDX-License-Identifier: GPL-3.0
pragma solidity >=0.7.0 <0.9.0;

/**
 * @title Store
 * @dev Store
 */
// solc --abi store.sol --overwrite -o store && solc --bin store.sol --overwrite -o store
contract Store {

    mapping(bytes32 => bytes32) public items;
    uint public totalEntries = 0;

    function setItem(bytes32 key, bytes32 value) external returns (uint) {
        items[key] = value;
        return ++totalEntries;
    }
}
