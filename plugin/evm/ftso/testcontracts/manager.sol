// SPDX-License-Identifier: GPL-3.0
pragma solidity >=0.7.0 <0.9.0;

/**
 * @title Manager
 * @dev Manager
 */

contract Manager {

    address public rewardManagerVal;
    uint256 public durationSeconds;
    uint256 public rewardEpochsStartTsVal;
    uint256 public currentRewardEpoch;
    uint256 public fraction;

    mapping(uint256 => uint256[]) public rewardEpochsMap;

    constructor(address _rewardManager, uint256 _rewardEpochDurationSeconds, uint256 _rewardEpochsStartTsVal,
        uint256 _currentRewardEpoch, uint256 _fraction, uint256[] memory epochs, uint256[] memory rewardEpochPowerHeight,
        uint256[] memory rewardEpochStartHeight, uint256[] memory rewardEpochStartTime) {

        uint256 length = epochs.length;
        require(length == rewardEpochPowerHeight.length);
        require(length == rewardEpochStartHeight.length);
        require(length == rewardEpochStartTime.length);

        for (uint256 i = 0; i < length; i++) {
            rewardEpochsMap[epochs[i]] = [rewardEpochPowerHeight[i], rewardEpochStartHeight[i], rewardEpochStartTime[i]];
        }

        rewardManagerVal = _rewardManager;
        durationSeconds = _rewardEpochDurationSeconds;
        rewardEpochsStartTsVal = _rewardEpochsStartTsVal;
        currentRewardEpoch = _currentRewardEpoch;
        fraction = _fraction;
    }

    function rewardManager() external view returns (address) {
        return rewardManagerVal;
    }

    function rewardEpochDurationSeconds() external view returns (uint256) {
        return durationSeconds;
    }

    function rewardEpochs(uint256 epoch) external view returns (uint256, uint256, uint256) {
        uint256[] memory a = rewardEpochsMap[epoch];
        return (a[0], a[1], a[2]);
    }

    function rewardEpochsStartTs() external view returns (uint256) {
        return rewardEpochsStartTsVal;
    }

    function getCurrentRewardEpoch() external view returns (uint256) {
        return currentRewardEpoch;
    }

    function settings() external view returns (uint256, uint256, uint256, uint256, uint256, uint256, uint256, bool, bool) {
        return (fraction, 0, 0, 0, 0, 0, 0, false, false);
    }
}
