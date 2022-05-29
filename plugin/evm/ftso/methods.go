package ftso

// The constants below define the various function names for calls against the
// FTSO smart contracts, based on the respective ABIs.
const (
	getAddressManager   = "getFtsoManager"
	getAddressRegistry  = "getFtsoRegistry"
	getAddressRewards   = "rewardManager"
	getAddressWhitelist = "getVoterWhitelister"
	getAddressWNAT      = "wNat"
	getAddressVotepower = "readVotePowerContract"

	getFTSOSupply    = "totalSupply"
	getFTSOIndices   = "getSupportedIndices"
	getFTSOSettings  = "settings"
	getFTSOProviders = "getFtsoWhitelistedPriceProviders"

	getEpochCurrent = "getCurrentRewardEpoch"
	getEpochInfo    = "rewardEpochs"

	getProviderVotepower = "votePowerOfAt"
	getProviderRewards   = "getUnclaimedReward"
)
