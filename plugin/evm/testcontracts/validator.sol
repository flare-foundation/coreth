// SPDX-License-Identifier: GPL-3.0
pragma solidity >=0.7.0 <0.9.0;

/**
 * @title Validator
 * @dev Validator
 */
// solc --abi validator.sol --overwrite -o validator && solc --bin validator.sol --overwrite -o validator && abigen --bin=validator/validator.bin  --abi=validator/validator.abi -pkg contracts --type=validator -out validator.go
contract Validator {

    mapping(address => bytes20) public nodes;

    constructor(address[] memory dataProvidersAddresses, bytes20[] memory _nodes) {
        uint256 length = _nodes.length;
        require(length == dataProvidersAddresses.length);
        for (uint256 i = 0; i < length; i++)
            nodes[dataProvidersAddresses[i]] = _nodes[i];
    }

    function getNodeIdForDataProvider(address dataProvider) external view returns (bytes20) {
        return nodes[dataProvider];
    }
}
