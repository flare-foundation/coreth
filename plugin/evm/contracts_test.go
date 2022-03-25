//go:build integration
// +build integration

package evm

import (
	"math"
	"math/big"
	"strings"
	"testing"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/stretchr/testify/require"

	"github.com/flare-foundation/coreth/accounts/abi"
	"github.com/flare-foundation/coreth/accounts/abi/bind"
	"github.com/flare-foundation/coreth/accounts/abi/bind/backends"
	"github.com/flare-foundation/coreth/core"
)

const (
	// Test Store contract
	testStoreJSON = `[{"inputs":[{"internalType":"bytes32","name":"","type":"bytes32"}],"name":"items","outputs":[{"internalType":"bytes32","name":"","type":"bytes32"}],"stateMutability":"view","type":"function"},{"inputs":[{"internalType":"bytes32","name":"key","type":"bytes32"},{"internalType":"bytes32","name":"value","type":"bytes32"}],"name":"setItem","outputs":[],"stateMutability":"nonpayable","type":"function"}]`
	testStoreBin  = `0x608060405234801561001057600080fd5b50610114806100206000396000f3fe6080604052348015600f57600080fd5b506004361060325760003560e01c806348f343f3146037578063f56256c7146076575b600080fd5b606060048036036020811015604b57600080fd5b810190808035906020019092919050505060ab565b6040518082815260200191505060405180910390f35b60a960048036036040811015608a57600080fd5b81019080803590602001909291908035906020019092919050505060c3565b005b60006020528060005260406000206000915090505481565b8060008084815260200190815260200160002081905550505056fea2646970667358221220fe053d6eea044b8897233c0c52a14061e1cb5ce89c7456acb6ab673f9d039b4f64736f6c63430007060033`

	// VotePower contract
	testAbiVotePowerJSON = `[{"inputs":[{"internalType":"address[]","name":"providers","type":"address[]"},{"internalType":"uint256[]","name":"vps","type":"uint256[]"}],"stateMutability":"nonpayable","type":"constructor"},{"inputs":[{"internalType":"address","name":"provider","type":"address"}],"name":"votePowerOf","outputs":[{"internalType":"uint256","name":"","type":"uint256"}],"stateMutability":"view","type":"function"},{"inputs":[{"internalType":"address","name":"","type":"address"}],"name":"votePowers","outputs":[{"internalType":"uint256","name":"","type":"uint256"}],"stateMutability":"view","type":"function"}]`
	testAbiVotePowerBin  = `0x608060405234801561001057600080fd5b506040516103853803806103858339818101604052604081101561003357600080fd5b810190808051604051939291908464010000000082111561005357600080fd5b8382019150602082018581111561006957600080fd5b825186602082028301116401000000008211171561008657600080fd5b8083526020830192505050908051906020019060200280838360005b838110156100bd5780820151818401526020810190506100a2565b50505050905001604052602001805160405193929190846401000000008211156100e657600080fd5b838201915060208201858111156100fc57600080fd5b825186602082028301116401000000008211171561011957600080fd5b8083526020830192505050908051906020019060200280838360005b83811015610150578082015181840152602081019050610135565b505050509050016040525050506000815190508251811461017057600080fd5b60005b818110156101f15782818151811061018757fe5b602002602001015160008086848151811061019e57fe5b602002602001015173ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff168152602001908152602001600020819055508080600101915050610173565b50505050610181806102046000396000f3fe608060405234801561001057600080fd5b50600436106100365760003560e01c8063142d10181461003b5780632b2c0d0014610093575b600080fd5b61007d6004803603602081101561005157600080fd5b81019080803573ffffffffffffffffffffffffffffffffffffffff1690602001909291905050506100eb565b6040518082815260200191505060405180910390f35b6100d5600480360360208110156100a957600080fd5b81019080803573ffffffffffffffffffffffffffffffffffffffff169060200190929190505050610133565b6040518082815260200191505060405180910390f35b60008060008373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff168152602001908152602001600020549050919050565b6000602052806000526040600020600091509050548156fea264697066735822122061d621b33ebdd91d9983cb8eedc60f951f0328226d743e51982623eaa8e8cd9b64736f6c63430007060033`

	// Wnat contract
	testAbiWnatJSON = `[{"inputs":[{"internalType":"address","name":"_wNatVal","type":"address"}],"stateMutability":"nonpayable","type":"constructor"},{"inputs":[],"name":"readVotePowerContract","outputs":[{"internalType":"address","name":"","type":"address"}],"stateMutability":"view","type":"function"},{"inputs":[],"name":"wNatVal","outputs":[{"internalType":"address","name":"","type":"address"}],"stateMutability":"view","type":"function"}]`
	testAbiWnatBin  = `0x608060405234801561001057600080fd5b506040516101b23803806101b28339818101604052602081101561003357600080fd5b8101908080519060200190929190505050806000806101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff1602179055505061011e806100946000396000f3fe6080604052348015600f57600080fd5b506004361060325760003560e01c80639b3baa0e146037578063f708a3b7146069575b600080fd5b603d609b565b604051808273ffffffffffffffffffffffffffffffffffffffff16815260200191505060405180910390f35b606f60c4565b604051808273ffffffffffffffffffffffffffffffffffffffff16815260200191505060405180910390f35b60008060009054906101000a900473ffffffffffffffffffffffffffffffffffffffff16905090565b60008054906101000a900473ffffffffffffffffffffffffffffffffffffffff168156fea2646970667358221220009e7049941af6c59bd17b62a091fd720d06545233e4227f893363177f29de7364736f6c63430007060033`

	// Reward contract
	testAbiRewardJSON = `[{"inputs":[{"internalType":"address","name":"_wNatVal","type":"address"},{"internalType":"uint256[]","name":"epochs","type":"uint256[]"},{"internalType":"address[]","name":"providers","type":"address[]"},{"internalType":"uint256[]","name":"unclaimedRewards","type":"uint256[]"}],"stateMutability":"nonpayable","type":"constructor"},{"inputs":[{"internalType":"uint256","name":"epoch","type":"uint256"},{"internalType":"address","name":"provider","type":"address"}],"name":"getUnclaimedReward","outputs":[{"internalType":"uint256","name":"","type":"uint256"}],"stateMutability":"view","type":"function"},{"inputs":[{"internalType":"uint256","name":"","type":"uint256"},{"internalType":"address","name":"","type":"address"}],"name":"unclaimedReward","outputs":[{"internalType":"uint256","name":"","type":"uint256"}],"stateMutability":"view","type":"function"},{"inputs":[],"name":"wNat","outputs":[{"internalType":"address","name":"","type":"address"}],"stateMutability":"view","type":"function"},{"inputs":[],"name":"wNatVal","outputs":[{"internalType":"address","name":"","type":"address"}],"stateMutability":"view","type":"function"}]`
	testAbiRewardBin  = `0x608060405234801561001057600080fd5b506040516105953803806105958339818101604052608081101561003357600080fd5b81019080805190602001909291908051604051939291908464010000000082111561005d57600080fd5b8382019150602082018581111561007357600080fd5b825186602082028301116401000000008211171561009057600080fd5b8083526020830192505050908051906020019060200280838360005b838110156100c75780820151818401526020810190506100ac565b50505050905001604052602001805160405193929190846401000000008211156100f057600080fd5b8382019150602082018581111561010657600080fd5b825186602082028301116401000000008211171561012357600080fd5b8083526020830192505050908051906020019060200280838360005b8381101561015a57808201518184015260208101905061013f565b505050509050016040526020018051604051939291908464010000000082111561018357600080fd5b8382019150602082018581111561019957600080fd5b82518660208202830111640100000000821117156101b657600080fd5b8083526020830192505050908051906020019060200280838360005b838110156101ed5780820151818401526020810190506101d2565b50505050905001604052505050836000806101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff1602179055506000835190508251811461024d57600080fd5b8151811461025a57600080fd5b60005b818110156103005782818151811061027157fe5b60200260200101516001600087848151811061028957fe5b6020026020010151815260200190815260200160002060008684815181106102ad57fe5b602002602001015173ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002081905550808060010191505061025d565b505050505050610280806103156000396000f3fe608060405234801561001057600080fd5b506004361061004c5760003560e01c8063657d96951461005157806371510ef2146100b35780639edbf00714610115578063f708a3b714610149575b600080fd5b61009d6004803603604081101561006757600080fd5b8101908080359060200190929190803573ffffffffffffffffffffffffffffffffffffffff16906020019092919050505061017d565b6040518082815260200191505060405180910390f35b6100ff600480360360408110156100c957600080fd5b8101908080359060200190929190803573ffffffffffffffffffffffffffffffffffffffff1690602001909291905050506101d8565b6040518082815260200191505060405180910390f35b61011d6101fd565b604051808273ffffffffffffffffffffffffffffffffffffffff16815260200191505060405180910390f35b610151610226565b604051808273ffffffffffffffffffffffffffffffffffffffff16815260200191505060405180910390f35b60006001600084815260200190815260200160002060008373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002054905092915050565b6001602052816000526040600020602052806000526040600020600091509150505481565b60008060009054906101000a900473ffffffffffffffffffffffffffffffffffffffff16905090565b60008054906101000a900473ffffffffffffffffffffffffffffffffffffffff168156fea2646970667358221220a3346edd7aba614edb78b0ed19cd0b42bb6d4a4ceb683a79ca3a57c752b2c43564736f6c63430007060033`

	// Manager contract
	testAbiManagerJSON = `[{"inputs":[{"internalType":"address","name":"_rewardManager","type":"address"},{"internalType":"uint256","name":"_rewardEpochDurationSeconds","type":"uint256"},{"internalType":"uint256","name":"_rewardEpochsStartTsVal","type":"uint256"},{"internalType":"uint256","name":"_currentRewardEpoch","type":"uint256"},{"internalType":"uint256[]","name":"epochs","type":"uint256[]"},{"internalType":"uint256[]","name":"rewardEpochPowerHeight","type":"uint256[]"},{"internalType":"uint256[]","name":"rewardEpochStartHeight","type":"uint256[]"},{"internalType":"uint256[]","name":"rewardEpochStartTime","type":"uint256[]"}],"stateMutability":"nonpayable","type":"constructor"},{"inputs":[],"name":"currentRewardEpoch","outputs":[{"internalType":"uint256","name":"","type":"uint256"}],"stateMutability":"view","type":"function"},{"inputs":[],"name":"durationSeconds","outputs":[{"internalType":"uint256","name":"","type":"uint256"}],"stateMutability":"view","type":"function"},{"inputs":[],"name":"getCurrentRewardEpoch","outputs":[{"internalType":"uint256","name":"","type":"uint256"}],"stateMutability":"view","type":"function"},{"inputs":[],"name":"rewardEpochDurationSeconds","outputs":[{"internalType":"uint256","name":"","type":"uint256"}],"stateMutability":"view","type":"function"},{"inputs":[{"internalType":"uint256","name":"epoch","type":"uint256"}],"name":"rewardEpochs","outputs":[{"internalType":"uint256","name":"","type":"uint256"},{"internalType":"uint256","name":"","type":"uint256"},{"internalType":"uint256","name":"","type":"uint256"}],"stateMutability":"view","type":"function"},{"inputs":[{"internalType":"uint256","name":"","type":"uint256"},{"internalType":"uint256","name":"","type":"uint256"}],"name":"rewardEpochsMap","outputs":[{"internalType":"uint256","name":"","type":"uint256"}],"stateMutability":"view","type":"function"},{"inputs":[],"name":"rewardEpochsStartTs","outputs":[{"internalType":"uint256","name":"","type":"uint256"}],"stateMutability":"view","type":"function"},{"inputs":[],"name":"rewardEpochsStartTsVal","outputs":[{"internalType":"uint256","name":"","type":"uint256"}],"stateMutability":"view","type":"function"},{"inputs":[],"name":"rewardManager","outputs":[{"internalType":"address","name":"","type":"address"}],"stateMutability":"view","type":"function"},{"inputs":[],"name":"rewardManagerVal","outputs":[{"internalType":"address","name":"","type":"address"}],"stateMutability":"view","type":"function"}]`
	testAbiManagerBin  = `0x608060405234801561001057600080fd5b5060405161084d38038061084d833981810160405261010081101561003457600080fd5b81019080805190602001909291908051906020019092919080519060200190929190805190602001909291908051604051939291908464010000000082111561007c57600080fd5b8382019150602082018581111561009257600080fd5b82518660208202830111640100000000821117156100af57600080fd5b8083526020830192505050908051906020019060200280838360005b838110156100e65780820151818401526020810190506100cb565b505050509050016040526020018051604051939291908464010000000082111561010f57600080fd5b8382019150602082018581111561012557600080fd5b825186602082028301116401000000008211171561014257600080fd5b8083526020830192505050908051906020019060200280838360005b8381101561017957808201518184015260208101905061015e565b50505050905001604052602001805160405193929190846401000000008211156101a257600080fd5b838201915060208201858111156101b857600080fd5b82518660208202830111640100000000821117156101d557600080fd5b8083526020830192505050908051906020019060200280838360005b8381101561020c5780820151818401526020810190506101f1565b505050509050016040526020018051604051939291908464010000000082111561023557600080fd5b8382019150602082018581111561024b57600080fd5b825186602082028301116401000000008211171561026857600080fd5b8083526020830192505050908051906020019060200280838360005b8381101561029f578082015181840152602081019050610284565b50505050905001604052505050600084519050835181146102bf57600080fd5b825181146102cc57600080fd5b815181146102d957600080fd5b60005b8181101561037a5760405180606001604052808683815181106102fb57fe5b6020026020010151815260200185838151811061031457fe5b6020026020010151815260200184838151811061032d57fe5b60200260200101518152506004600088848151811061034857fe5b6020026020010151815260200190815260200160002090600361036c9291906103de565b5080806001019150506102dc565b50886000806101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff160217905550876001819055508660028190555085600381905550505050505050505050610448565b82805482825590600052602060002090810192821561041a579160200282015b828111156104195782518255916020019190600101906103fe565b5b509050610427919061042b565b5090565b5b8082111561044457600081600090555060010161042c565b5090565b6103f6806104576000396000f3fe608060405234801561001057600080fd5b506004361061009e5760003560e01c80639acba2af116100665780639acba2af14610193578063a578f55b146101b1578063a795f409146101cf578063ab3caf411461021f578063e7c830d41461023d5761009e565b80630f4ef8a6146100a35780635080247b146100d757806356ecf28b14610123578063701a48ed1461014157806385f3c9c914610175575b600080fd5b6100ab61025b565b604051808273ffffffffffffffffffffffffffffffffffffffff16815260200191505060405180910390f35b61010d600480360360408110156100ed57600080fd5b810190808035906020019092919080359060200190929190505050610284565b6040518082815260200191505060405180910390f35b61012b6102b5565b6040518082815260200191505060405180910390f35b6101496102bb565b604051808273ffffffffffffffffffffffffffffffffffffffff16815260200191505060405180910390f35b61017d6102df565b6040518082815260200191505060405180910390f35b61019b6102e9565b6040518082815260200191505060405180910390f35b6101b96102ef565b6040518082815260200191505060405180910390f35b6101fb600480360360208110156101e557600080fd5b81019080803590602001909291905050506102f9565b60405180848152602001838152602001828152602001935050505060405180910390f35b6102276103b0565b6040518082815260200191505060405180910390f35b6102456103b6565b6040518082815260200191505060405180910390f35b60008060009054906101000a900473ffffffffffffffffffffffffffffffffffffffff16905090565b600460205281600052604060002081815481106102a057600080fd5b90600052602060002001600091509150505481565b60035481565b60008054906101000a900473ffffffffffffffffffffffffffffffffffffffff1681565b6000600154905090565b60015481565b6000600254905090565b6000806000806004600086815260200190815260200160002080548060200260200160405190810160405280929190818152602001828054801561035c57602002820191906000526020600020905b815481526020019060010190808311610348575b505050505090508060008151811061037057fe5b60200260200101518160018151811061038557fe5b60200260200101518260028151811061039a57fe5b6020026020010151935093509350509193909250565b60025481565b600060035490509056fea26469706673582212204b8cfed0b8a4ad6d90d36ac9954dc09f23a91ce4b47bb8bb4aceb066812ec12964736f6c63430007060033`

	// Registry contract
	testAbiFtsoRegistryJSON = `[{"inputs":[{"internalType":"uint256[]","name":"_supportedIndices","type":"uint256[]"}],"stateMutability":"nonpayable","type":"constructor"},{"inputs":[],"name":"getSupportedIndices","outputs":[{"internalType":"uint256[]","name":"","type":"uint256[]"}],"stateMutability":"view","type":"function"},{"inputs":[{"internalType":"uint256","name":"","type":"uint256"}],"name":"supportedIndices","outputs":[{"internalType":"uint256","name":"","type":"uint256"}],"stateMutability":"view","type":"function"}]`
	testAbiFtsoRegistryBin  = `0x608060405234801561001057600080fd5b506040516102ee3803806102ee8339818101604052602081101561003357600080fd5b810190808051604051939291908464010000000082111561005357600080fd5b8382019150602082018581111561006957600080fd5b825186602082028301116401000000008211171561008657600080fd5b8083526020830192505050908051906020019060200280838360005b838110156100bd5780820151818401526020810190506100a2565b5050505090500160405250505080600090805190602001906100e09291906100e7565b5050610151565b828054828255906000526020600020908101928215610123579160200282015b82811115610122578251825591602001919060010190610107565b5b5090506101309190610134565b5090565b5b8082111561014d576000816000905550600101610135565b5090565b61018e806101606000396000f3fe608060405234801561001057600080fd5b50600436106100365760003560e01c80633b18bdcd1461003b578063798aac5b1461007d575b600080fd5b6100676004803603602081101561005157600080fd5b81019080803590602001909291905050506100dc565b6040518082815260200191505060405180910390f35b610085610100565b6040518080602001828103825283818151815260200191508051906020019060200280838360005b838110156100c85780820151818401526020810190506100ad565b505050509050019250505060405180910390f35b600081815481106100ec57600080fd5b906000526020600020016000915090505481565b6060600080548060200260200160405190810160405280929190818152602001828054801561014e57602002820191906000526020600020905b81548152602001906001019080831161013a575b505050505090509056fea2646970667358221220dfff25759b4f429af0ff119c7a5ebff831df0432cc50cf531ed71b212a30699664736f6c63430007060033`

	// Whitelist contract
	testAbiVoterWhitelisterJSON = `[{"inputs":[{"internalType":"uint256[]","name":"indices","type":"uint256[]"},{"internalType":"address[]","name":"priceProvidersAddresses","type":"address[]"}],"stateMutability":"nonpayable","type":"constructor"},{"inputs":[{"internalType":"uint256","name":"ftsoIndex","type":"uint256"}],"name":"getFtsoWhitelistedPriceProviders","outputs":[{"internalType":"address[]","name":"","type":"address[]"}],"stateMutability":"view","type":"function"},{"inputs":[{"internalType":"uint256","name":"","type":"uint256"},{"internalType":"uint256","name":"","type":"uint256"}],"name":"priceProviders","outputs":[{"internalType":"address","name":"","type":"address"}],"stateMutability":"view","type":"function"}]`
	testAbiVoterWhitelisterBin  = `0x608060405234801561001057600080fd5b506040516104b23803806104b28339818101604052604081101561003357600080fd5b810190808051604051939291908464010000000082111561005357600080fd5b8382019150602082018581111561006957600080fd5b825186602082028301116401000000008211171561008657600080fd5b8083526020830192505050908051906020019060200280838360005b838110156100bd5780820151818401526020810190506100a2565b50505050905001604052602001805160405193929190846401000000008211156100e657600080fd5b838201915060208201858111156100fc57600080fd5b825186602082028301116401000000008211171561011957600080fd5b8083526020830192505050908051906020019060200280838360005b83811015610150578082015181840152602081019050610135565b5050505090500160405250505060005b82518110156101b0578160008085848151811061017957fe5b6020026020010151815260200190815260200160002090805190602001906101a29291906101b8565b508080600101915050610160565b50505061025f565b828054828255906000526020600020908101928215610231579160200282015b828111156102305782518260006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff160217905550916020019190600101906101d8565b5b50905061023e9190610242565b5090565b5b8082111561025b576000816000905550600101610243565b5090565b6102448061026e6000396000f3fe608060405234801561001057600080fd5b50600436106100365760003560e01c806309fcb4001461003b5780635a34e814146100be575b600080fd5b6100676004803603602081101561005157600080fd5b8101908080359060200190929190505050610120565b6040518080602001828103825283818151815260200191508051906020019060200280838360005b838110156100aa57808201518184015260208101905061008f565b505050509050019250505060405180910390f35b6100f4600480360360408110156100d457600080fd5b8101908080359060200190929190803590602001909291905050506101c0565b604051808273ffffffffffffffffffffffffffffffffffffffff16815260200191505060405180910390f35b60606000808381526020019081526020016000208054806020026020016040519081016040528092919081815260200182805480156101b457602002820191906000526020600020905b8160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff168152602001906001019080831161016a575b50505050509050919050565b600060205281600052604060002081815481106101dc57600080fd5b906000526020600020016000915091509054906101000a900473ffffffffffffffffffffffffffffffffffffffff168156fea26469706673582212207a648ee953be94989df2dc1ea2bf39fa6831b7663e6a5bb37fe9267e1f8ee34a64736f6c63430007060033`

	// Validation contract
	testAbiValidationJSON = `[{"inputs":[{"internalType":"address[]","name":"dataProvidersAddresses","type":"address[]"},{"internalType":"bytes20[]","name":"_nodes","type":"bytes20[]"}],"stateMutability":"nonpayable","type":"constructor"},{"inputs":[{"internalType":"address","name":"dataProvider","type":"address"}],"name":"getNodeIdForDataProvider","outputs":[{"internalType":"bytes20","name":"","type":"bytes20"}],"stateMutability":"view","type":"function"},{"inputs":[{"internalType":"address","name":"","type":"address"}],"name":"nodes","outputs":[{"internalType":"bytes20","name":"","type":"bytes20"}],"stateMutability":"view","type":"function"}]`
	testAbiValidationBin  = `0x608060405234801561001057600080fd5b506040516103df3803806103df8339818101604052604081101561003357600080fd5b810190808051604051939291908464010000000082111561005357600080fd5b8382019150602082018581111561006957600080fd5b825186602082028301116401000000008211171561008657600080fd5b8083526020830192505050908051906020019060200280838360005b838110156100bd5780820151818401526020810190506100a2565b50505050905001604052602001805160405193929190846401000000008211156100e657600080fd5b838201915060208201858111156100fc57600080fd5b825186602082028301116401000000008211171561011957600080fd5b8083526020830192505050908051906020019060200280838360005b83811015610150578082015181840152602081019050610135565b505050509050016040525050506000815190508251811461017057600080fd5b60005b818110156102185782818151811061018757fe5b602002602001015160008086848151811061019e57fe5b602002602001015173ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908360601c02179055508080600101915050610173565b505050506101b48061022b6000396000f3fe608060405234801561001057600080fd5b50600436106100365760003560e01c8063177fbc711461003b578063189a5a17146100a2575b600080fd5b61007d6004803603602081101561005157600080fd5b81019080803573ffffffffffffffffffffffffffffffffffffffff169060200190929190505050610109565b60405180826bffffffffffffffffffffffff1916815260200191505060405180910390f35b6100e4600480360360208110156100b857600080fd5b81019080803573ffffffffffffffffffffffffffffffffffffffff16906020019092919050505061015e565b60405180826bffffffffffffffffffffffff1916815260200191505060405180910390f35b60008060008373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060009054906101000a900460601b9050919050565b60006020528060005260406000206000915054906101000a900460601b8156fea26469706673582212206fcff72ab7528e2f49e8ed7b227acd87df1719c83d37bf75dd10aaa2b8ed333f64736f6c63430007060033`

	// Submitter contract
	testAbiSubmitterJSON = `[{"inputs":[{"internalType":"address","name":"_voterWhitelister","type":"address"},{"internalType":"address","name":"_ftsoRegistry","type":"address"},{"internalType":"address","name":"_ftsoManager","type":"address"}],"stateMutability":"nonpayable","type":"constructor"},{"inputs":[],"name":"ftsoManager","outputs":[{"internalType":"address","name":"","type":"address"}],"stateMutability":"view","type":"function"},{"inputs":[],"name":"ftsoRegistry","outputs":[{"internalType":"address","name":"","type":"address"}],"stateMutability":"view","type":"function"},{"inputs":[],"name":"getFtsoManager","outputs":[{"internalType":"address","name":"","type":"address"}],"stateMutability":"view","type":"function"},{"inputs":[],"name":"getFtsoRegistry","outputs":[{"internalType":"address","name":"","type":"address"}],"stateMutability":"view","type":"function"},{"inputs":[],"name":"getVoterWhitelister","outputs":[{"internalType":"address","name":"","type":"address"}],"stateMutability":"view","type":"function"},{"inputs":[],"name":"voterWhitelister","outputs":[{"internalType":"address","name":"","type":"address"}],"stateMutability":"view","type":"function"}]`
	testAbiSubmitterBin  = `0x608060405234801561001057600080fd5b506040516103ee3803806103ee8339818101604052606081101561003357600080fd5b81019080805190602001909291908051906020019092919080519060200190929190505050826000806101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff16021790555081600160006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff16021790555080600260006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff1602179055505050506102c28061012c6000396000f3fe608060405234801561001057600080fd5b50600436106100625760003560e01c806311a7aaaa1461006757806338b5f8691461009b57806371e1fad9146100cf5780638c9d28b614610103578063b39c685814610137578063c2b0d47b1461016b575b600080fd5b61006f61019f565b604051808273ffffffffffffffffffffffffffffffffffffffff16815260200191505060405180910390f35b6100a36101c5565b604051808273ffffffffffffffffffffffffffffffffffffffff16815260200191505060405180910390f35b6100d76101eb565b604051808273ffffffffffffffffffffffffffffffffffffffff16815260200191505060405180910390f35b61010b610214565b604051808273ffffffffffffffffffffffffffffffffffffffff16815260200191505060405180910390f35b61013f61023e565b604051808273ffffffffffffffffffffffffffffffffffffffff16815260200191505060405180910390f35b610173610268565b604051808273ffffffffffffffffffffffffffffffffffffffff16815260200191505060405180910390f35b600260009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1681565b600160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1681565b60008060009054906101000a900473ffffffffffffffffffffffffffffffffffffffff16905090565b6000600160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff16905090565b6000600260009054906101000a900473ffffffffffffffffffffffffffffffffffffffff16905090565b60008054906101000a900473ffffffffffffffffffffffffffffffffffffffff168156fea264697066735822122035d0b8aa6a058839885934e180e1ddb5af32130473692bb78427ccd544f1323f64736f6c63430007060033`
)

var (
	testAbiFtsoRegistry     abi.ABI
	testAbiManager          abi.ABI
	testAbiReward           abi.ABI
	testAbiStore            abi.ABI
	testAbiSubmitter        abi.ABI
	testAbiValidation       abi.ABI
	testAbiVotePower        abi.ABI
	testAbiVoterWhitelister abi.ABI
	testAbiWnat             abi.ABI
)

func initTestContracts(t *testing.T) {
	t.Helper()

	var err error

	testAbiStore, err = abi.JSON(strings.NewReader(testStoreJSON))
	require.NoError(t, err)

	testAbiVotePower, err = abi.JSON(strings.NewReader(testAbiVotePowerJSON))
	require.NoError(t, err)

	testAbiWnat, err = abi.JSON(strings.NewReader(testAbiWnatJSON))
	require.NoError(t, err)

	testAbiReward, err = abi.JSON(strings.NewReader(testAbiRewardJSON))
	require.NoError(t, err)

	testAbiManager, err = abi.JSON(strings.NewReader(testAbiManagerJSON))
	require.NoError(t, err)

	testAbiFtsoRegistry, err = abi.JSON(strings.NewReader(testAbiFtsoRegistryJSON))
	require.NoError(t, err)

	testAbiVoterWhitelister, err = abi.JSON(strings.NewReader(testAbiVoterWhitelisterJSON))
	require.NoError(t, err)

	testAbiValidation, err = abi.JSON(strings.NewReader(testAbiValidationJSON))
	require.NoError(t, err)

	testAbiSubmitter, err = abi.JSON(strings.NewReader(testAbiSubmitterJSON))
	require.NoError(t, err)
}

// Slices epochs, rewardEpochPowerHeight, rewardEpochStartHeight and rewardEpochStartTime
// should all have the same length.
//
// When calling RewardEpoch method using 'epochs[i]' values 'rewardEpochPowerHeight[i]',
// 'rewardEpochStartHeight[i]', 'rewardEpochStartTime[i]' will be returned
//
func deployManagerContract(t *testing.T, auth *bind.TransactOpts, be *backends.SimulatedBackend,
	rewardManager common.Address, rewardEpochDurationSeconds *big.Int, rewardEpochsStartTs *big.Int, currentRewardEpoch *big.Int,
	epochs []*big.Int, rewardEpochPowerHeight []*big.Int, rewardEpochStartHeight []*big.Int, rewardEpochStartTime []*big.Int,
) common.Address {

	t.Helper()

	return deployTestContract(t, auth, be, testAbiManager, testAbiManagerBin, rewardManager, rewardEpochDurationSeconds,
		rewardEpochsStartTs, currentRewardEpoch, epochs, rewardEpochPowerHeight, rewardEpochStartHeight, rewardEpochStartTime)

}

func deployFTSORegistryContract(t *testing.T, auth *bind.TransactOpts, be *backends.SimulatedBackend, supportedIndices []*big.Int) common.Address {
	t.Helper()

	return deployTestContract(t, auth, be, testAbiFtsoRegistry, testAbiFtsoRegistryBin, supportedIndices)

}

// Slices epochs, providers and unclaimedRewards should all have the same length.
//
// When calling ProviderRewards method using 'epochs[i]' and 'providers[i]'
// value 'unclaimedRewards[i]' will be returned.
//
func deployRewardContract(t *testing.T, auth *bind.TransactOpts, be *backends.SimulatedBackend,
	wNat common.Address, epochs []*big.Int, providers []common.Address, unclaimedRewards []*big.Int,
) common.Address {

	t.Helper()

	return deployTestContract(t, auth, be, testAbiReward, testAbiRewardBin, wNat, epochs, providers, unclaimedRewards)

}

func deploySubmitterContract(t *testing.T, auth *bind.TransactOpts, be *backends.SimulatedBackend,
	voterWhitelister common.Address, ftsoRegistry common.Address, ftsoManager common.Address,
) common.Address {

	t.Helper()

	return deployTestContract(t, auth, be, testAbiSubmitter, testAbiSubmitterBin, voterWhitelister, ftsoRegistry, ftsoManager)
}

// Slice dataProvidersAddresses should have the same length as nodes outer slice
//
// When calling ProviderNode method using 'dataProvidersAddresses[i]'
// value 'nodes[i]' will be returned.
//
func deployValidatorContract(t *testing.T, auth *bind.TransactOpts, be *backends.SimulatedBackend,
	dataProvidersAddresses []common.Address, nodes [][20]byte,
) common.Address {

	t.Helper()

	return deployTestContract(t, auth, be, testAbiValidation, testAbiValidationBin, dataProvidersAddresses, nodes)
}

// Slices providers and vps should all have the same length.
//
// When calling ProviderVotepower method using 'providers[i]'
// value 'vps[i]' will be returned.
//
func deployVotepowerContract(t *testing.T, auth *bind.TransactOpts, be *backends.SimulatedBackend,
	providers []common.Address, vps []*big.Int,
) common.Address {

	t.Helper()

	return deployTestContract(t, auth, be, testAbiVotePower, testAbiVotePowerBin, providers, vps)
}

func deployWhitelistContract(t *testing.T, auth *bind.TransactOpts, be *backends.SimulatedBackend,
	ftsoIndices []*big.Int, priceProvidersAddresses []common.Address,
) common.Address {

	t.Helper()

	return deployTestContract(t, auth, be, testAbiVoterWhitelister, testAbiVoterWhitelisterBin, ftsoIndices, priceProvidersAddresses)
}

func deployWnatContract(t *testing.T, auth *bind.TransactOpts, be *backends.SimulatedBackend, wNat common.Address) common.Address {
	t.Helper()

	return deployTestContract(t, auth, be, testAbiWnat, testAbiWnatBin, wNat)
}

func deployTestContract(t *testing.T, auth *bind.TransactOpts, be *backends.SimulatedBackend, testAbi abi.ABI, testAbiBin string, params ...interface{}) common.Address {
	t.Helper()

	addr, _, _, err := bind.DeployContract(auth, testAbi, common.FromHex(testAbiBin), be, params...)
	require.NoError(t, err)

	return addr
}

func simulatedBlockchain(t *testing.T) (*bind.TransactOpts, *backends.SimulatedBackend) {
	t.Helper()

	key, err := crypto.GenerateKey()
	require.NoError(t, err)

	auth, err := bind.NewKeyedTransactorWithChainID(key, big.NewInt(1337))
	require.NoError(t, err)

	balance := big.NewInt(0).SetUint64(math.MaxUint64)
	alloc := make(core.GenesisAlloc)
	alloc[auth.From] = core.GenesisAccount{
		Balance: balance,
	}
	gasLimit := uint64(math.MaxUint64)

	be := backends.NewSimulatedBackend(alloc, gasLimit)

	return auth, be
}
