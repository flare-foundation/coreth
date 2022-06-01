// SPDX-License-Identifier: GPL-3.0
pragma solidity >=0.7.0 <0.9.0;

/**
 * @title Submitter
 * @dev Submitter
 */

// solc --abi submitter.sol --overwrite -o submitter && solc --bin submitter.sol --overwrite -o submitter && abigen --bin=submitter/submitter.bin  --abi=submitter/submitter.abi -pkg contracts --type=submitter -out submitter.go
contract Submitter {

    address public voterWhitelister;
    address public ftsoRegistry;
    address public ftsoManager;

    constructor(address _voterWhitelister, address _ftsoRegistry, address _ftsoManager) {
        voterWhitelister = _voterWhitelister;
        ftsoRegistry = _ftsoRegistry;
        ftsoManager = _ftsoManager;
    }

    function getVoterWhitelister() external view returns (address) {
        return voterWhitelister;
    }

    function getFtsoRegistry() external view returns (address) {
        return ftsoRegistry;
    }

    function getFtsoManager() external view returns (address) {
        return ftsoManager;
    }
}
