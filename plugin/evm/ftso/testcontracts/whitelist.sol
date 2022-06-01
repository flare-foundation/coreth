// SPDX-License-Identifier: GPL-3.0
pragma solidity >=0.7.0 <0.9.0;

/**
 * @title Whitelist
 * @dev Whitelist
 */

// solc --abi whitelist.sol --overwrite -o whitelist && solc --bin whitelist.sol --overwrite -o whitelist && abigen --bin=whitelist/whitelist.bin  --abi=whitelist/whitelist.abi -pkg contracts --type=whitelist -out whitelist.go
contract Whitelist {

    mapping(uint256 => address[]) public priceProviders;

    constructor(uint256[] memory indices, address[] memory priceProvidersAddresses) {
        for (uint256 i = 0; i < indices.length; i++)
            priceProviders[indices[i]] = priceProvidersAddresses;
    }

    function getFtsoWhitelistedPriceProviders(uint256 ftsoIndex) external view returns (address[] memory) {
        return priceProviders[ftsoIndex];
    }
}
