// SPDX-License-Identifier: GPL-3.0
pragma solidity >=0.7.0 <0.9.0;

/**
 * @title Reward
 * @dev Reward
 */
// solc --abi reward.sol --overwrite -o reward && solc --bin reward.sol --overwrite -o reward && abigen --bin=reward/reward.bin  --abi=reward/reward.abi -pkg contracts --type=reward -out reward.go

contract Reward {

    address public wNatVal;

    mapping(uint256 => mapping(address => uint256)) public unclaimedReward;

    constructor(address _wNatVal, uint256[] memory epochs, address[] memory providers, uint256[] memory unclaimedRewards) {
        wNatVal = _wNatVal;

        uint256 length = epochs.length;
        require(length == providers.length);
        require(length == unclaimedRewards.length);

        for (uint256 i = 0; i < length; i++) {
            unclaimedReward[epochs[i]][providers[i]] = unclaimedRewards[i];
        }
    }

    function wNat() external view returns (address) {
        return wNatVal;
    }

    function getUnclaimedReward(uint256 epoch, address provider) external view returns (uint256, uint256) {
        return (unclaimedReward[epoch][provider], 0);
    }
}
