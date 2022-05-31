// SPDX-License-Identifier: GPL-3.0
pragma solidity >=0.7.0 <0.9.0;

/**
 * @title Registry
 * @dev Registry
 */
// solc --abi registry.sol --overwrite -o registry && solc --bin registry.sol --overwrite -o registry && abigen --bin=registry/registry.bin  --abi=registry/registry.abi -pkg contracts --type=registry -out registry.go
contract Registry {

    uint256[] public supportedIndices;

    constructor(uint256[] memory _supportedIndices) {
        supportedIndices = _supportedIndices;
    }

    function getSupportedIndices() external view returns (uint256[] memory) {
        return supportedIndices;
    }
}
