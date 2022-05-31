// SPDX-License-Identifier: GPL-3.0
pragma solidity >=0.7.0 <0.9.0;

/**
 * @title Votepower
 * @dev Votepower
 */
// solc --abi votepower.sol --overwrite -o votepower && solc --bin votepower.sol --overwrite -o votepower && abigen --bin=votepower/votepower.bin  --abi=votepower/votepower.abi -pkg contracts --type=votepower -out votepower.go

contract Votepower {

    mapping(address => mapping(uint256 => uint256)) public votePowers;

    constructor(address[] memory providers, uint256[] memory heights, uint256[] memory vps) {
        uint256 length = vps.length;
        require(length == providers.length);
        require(length == heights.length);
        for (uint256 i = 0; i < length; i++) {
            votePowers[providers[i]][heights[i]] = vps[i];
        }
    }

    function votePowerOfAt(address provider, uint256 height) external view returns (uint256) {
        return votePowers[provider][height];
    }
}
