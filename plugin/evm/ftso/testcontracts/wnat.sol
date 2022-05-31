// SPDX-License-Identifier: GPL-3.0
pragma solidity >=0.7.0 <0.9.0;

/**
 * @title Wnat
 * @dev Wnat
 */

// solc --abi wnat.sol --overwrite -o wnat && solc --bin wnat.sol --overwrite -o wnat && abigen --bin=wnat/wnat.bin  --abi=wnat/wnat.abi -pkg contracts --type=wnat -out wnat.go
contract Wnat {

    address public wNatVal;
    uint256 public totalSupplyVal;

    constructor(address _wNatVal, uint256 _totalSupplyVal) {
        wNatVal = _wNatVal;
        totalSupplyVal = _totalSupplyVal;
    }

    function readVotePowerContract() external view returns (address) {
        return wNatVal;
    }

    function totalSupply() external view returns (uint256) {
        return totalSupplyVal;
    }
}
