// SPDX-License-Identifier: GPL-3.0
pragma solidity >=0.7.0 <0.9.0;

/**
 * @title Store
 * @dev Store
 */
// abigen --bin=store/store.bin  --abi=store/store.abi -pkg contracts --type=store -out store.go
contract Store {

    mapping(bytes32 => bytes32) public items;

    function setItem(bytes32 key, bytes32 value) external returns (bool) {
        items[key] = value;
        return true;
    }
}
