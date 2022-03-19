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

func initDummyContracts(t *testing.T, auth *bind.TransactOpts, be *backends.SimulatedBackend) *FTSOSystem {
	t.Helper()

	abiVotePowerJSON := `[{"inputs":[{"internalType":"uint256","name":"_vpInt","type":"uint256"}],"stateMutability":"nonpayable","type":"constructor"},{"inputs":[{"internalType":"address","name":"provider","type":"address"}],"name":"votePowerOf","outputs":[{"internalType":"uint256","name":"","type":"uint256"}],"stateMutability":"view","type":"function"},{"inputs":[],"name":"vpInt","outputs":[{"internalType":"uint256","name":"","type":"uint256"}],"stateMutability":"view","type":"function"}]`
	abiVotePowerBin := `0x608060405234801561001057600080fd5b506040516101493803806101498339818101604052602081101561003357600080fd5b8101908080519060200190929190505050806000819055505060ef8061005a6000396000f3fe6080604052348015600f57600080fd5b506004361060325760003560e01c8063142d10181460375780638966b43214608c575b600080fd5b607660048036036020811015604b57600080fd5b81019080803573ffffffffffffffffffffffffffffffffffffffff16906020019092919050505060a8565b6040518082815260200191505060405180910390f35b609260b3565b6040518082815260200191505060405180910390f35b600080549050919050565b6000548156fea264697066735822122043d0eedff37b596f65fc5d3f9b98f8444bac96fcf123c979b81ece56c2991d2664736f6c63430007060033`

	abiVotePower, err := abi.JSON(strings.NewReader(abiVotePowerJSON))
	require.NoError(t, err)

	vpInt := big.NewInt(100)
	votePowerAddr, _, _, err := bind.DeployContract(auth, abiVotePower, common.FromHex(abiVotePowerBin), be, vpInt)
	require.NoError(t, err)

	abiWnatJSON := `[{"inputs":[{"internalType":"address","name":"_wNatVal","type":"address"}],"stateMutability":"nonpayable","type":"constructor"},{"inputs":[],"name":"readVotePowerContract","outputs":[{"internalType":"address","name":"","type":"address"}],"stateMutability":"view","type":"function"},{"inputs":[],"name":"wNatVal","outputs":[{"internalType":"address","name":"","type":"address"}],"stateMutability":"view","type":"function"}]`
	abiWnatBin := `0x608060405234801561001057600080fd5b506040516101b23803806101b28339818101604052602081101561003357600080fd5b8101908080519060200190929190505050806000806101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff1602179055505061011e806100946000396000f3fe6080604052348015600f57600080fd5b506004361060325760003560e01c80639b3baa0e146037578063f708a3b7146069575b600080fd5b603d609b565b604051808273ffffffffffffffffffffffffffffffffffffffff16815260200191505060405180910390f35b606f60c4565b604051808273ffffffffffffffffffffffffffffffffffffffff16815260200191505060405180910390f35b60008060009054906101000a900473ffffffffffffffffffffffffffffffffffffffff16905090565b60008054906101000a900473ffffffffffffffffffffffffffffffffffffffff168156fea2646970667358221220009e7049941af6c59bd17b62a091fd720d06545233e4227f893363177f29de7364736f6c63430007060033`

	abiWnat, err := abi.JSON(strings.NewReader(abiWnatJSON))
	require.NoError(t, err)

	wnatAddr, _, _, err := bind.DeployContract(auth, abiWnat, common.FromHex(abiWnatBin), be, votePowerAddr)

	// Reward contract
	abiRewardJSON := `[{"inputs":[{"internalType":"address","name":"_wNatVal","type":"address"},{"internalType":"uint256","name":"_unclaimedReward","type":"uint256"}],"stateMutability":"nonpayable","type":"constructor"},{"inputs":[{"internalType":"uint256","name":"epoch","type":"uint256"},{"internalType":"address","name":"provider","type":"address"}],"name":"getUnclaimedReward","outputs":[{"internalType":"uint256","name":"","type":"uint256"}],"stateMutability":"view","type":"function"},{"inputs":[],"name":"unclaimedReward","outputs":[{"internalType":"uint256","name":"","type":"uint256"}],"stateMutability":"view","type":"function"},{"inputs":[],"name":"wNat","outputs":[{"internalType":"address","name":"","type":"address"}],"stateMutability":"view","type":"function"},{"inputs":[],"name":"wNatVal","outputs":[{"internalType":"address","name":"","type":"address"}],"stateMutability":"view","type":"function"}]`
	abiRewardBin := `0x608060405234801561001057600080fd5b506040516102753803806102758339818101604052604081101561003357600080fd5b810190808051906020019092919080519060200190929190505050816000806101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff1602179055508060018190555050506101cf806100a66000396000f3fe608060405234801561001057600080fd5b506004361061004c5760003560e01c8063657d9695146100515780639edbf007146100b3578063cee83cee146100e7578063f708a3b714610105575b600080fd5b61009d6004803603604081101561006757600080fd5b8101908080359060200190929190803573ffffffffffffffffffffffffffffffffffffffff169060200190929190505050610139565b6040518082815260200191505060405180910390f35b6100bb610146565b604051808273ffffffffffffffffffffffffffffffffffffffff16815260200191505060405180910390f35b6100ef61016f565b6040518082815260200191505060405180910390f35b61010d610175565b604051808273ffffffffffffffffffffffffffffffffffffffff16815260200191505060405180910390f35b6000600154905092915050565b60008060009054906101000a900473ffffffffffffffffffffffffffffffffffffffff16905090565b60015481565b60008054906101000a900473ffffffffffffffffffffffffffffffffffffffff168156fea26469706673582212202de85ff60e4f8dc712b18a2ac0bc0b78455bfc02a03d036526eefa766354001664736f6c63430007060033`

	abiReward, err := abi.JSON(strings.NewReader(abiRewardJSON))
	require.NoError(t, err)

	unclaimedReward := big.NewInt(100)
	rewardAddr, _, _, err := bind.DeployContract(auth, abiReward, common.FromHex(abiRewardBin), be, wnatAddr, unclaimedReward)
	require.NoError(t, err)

	// Manager contract
	abiManagerJSON := `[{"inputs":[{"internalType":"address","name":"_rewardManager","type":"address"},{"internalType":"uint256","name":"_rewardEpochDurationSeconds","type":"uint256"},{"internalType":"uint256","name":"_rewardEpochPowerHeight","type":"uint256"},{"internalType":"uint256","name":"_rewardEpochStartHeight","type":"uint256"},{"internalType":"uint256","name":"_rewardEpochStartTime","type":"uint256"}],"stateMutability":"nonpayable","type":"constructor"},{"inputs":[],"name":"durationSeconds","outputs":[{"internalType":"uint256","name":"","type":"uint256"}],"stateMutability":"view","type":"function"},{"inputs":[],"name":"rewardEpochDurationSeconds","outputs":[{"internalType":"uint256","name":"","type":"uint256"}],"stateMutability":"view","type":"function"},{"inputs":[],"name":"rewardEpochPowerHeight","outputs":[{"internalType":"uint256","name":"","type":"uint256"}],"stateMutability":"view","type":"function"},{"inputs":[],"name":"rewardEpochStartHeight","outputs":[{"internalType":"uint256","name":"","type":"uint256"}],"stateMutability":"view","type":"function"},{"inputs":[],"name":"rewardEpochStartTime","outputs":[{"internalType":"uint256","name":"","type":"uint256"}],"stateMutability":"view","type":"function"},{"inputs":[{"internalType":"uint256","name":"epoch","type":"uint256"}],"name":"rewardEpochs","outputs":[{"internalType":"uint256","name":"","type":"uint256"},{"internalType":"uint256","name":"","type":"uint256"},{"internalType":"uint256","name":"","type":"uint256"}],"stateMutability":"view","type":"function"},{"inputs":[],"name":"rewardManager","outputs":[{"internalType":"address","name":"","type":"address"}],"stateMutability":"view","type":"function"},{"inputs":[],"name":"rewardManagerVal","outputs":[{"internalType":"address","name":"","type":"address"}],"stateMutability":"view","type":"function"}]`
	abiManagerBin := `0x608060405234801561001057600080fd5b50604051610377380380610377833981810160405260a081101561003357600080fd5b810190808051906020019092919080519060200190929190805190602001909291908051906020019092919080519060200190929190505050846000806101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff16021790555083600181905550826002819055508160038190555080600481905550505050505061029b806100dc6000396000f3fe608060405234801561001057600080fd5b50600436106100885760003560e01c806385f3c9c91161005b57806385f3c9c9146101315780639acba2af1461014f578063a795f4091461016d578063d5ea4a5c146101bd57610088565b80630f4ef8a61461008d5780632f1b6b97146100c157806340e43150146100df578063701a48ed146100fd575b600080fd5b6100956101db565b604051808273ffffffffffffffffffffffffffffffffffffffff16815260200191505060405180910390f35b6100c9610204565b6040518082815260200191505060405180910390f35b6100e761020a565b6040518082815260200191505060405180910390f35b610105610210565b604051808273ffffffffffffffffffffffffffffffffffffffff16815260200191505060405180910390f35b610139610234565b6040518082815260200191505060405180910390f35b61015761023e565b6040518082815260200191505060405180910390f35b6101996004803603602081101561018357600080fd5b8101908080359060200190929190505050610244565b60405180848152602001838152602001828152602001935050505060405180910390f35b6101c561025f565b6040518082815260200191505060405180910390f35b60008060009054906101000a900473ffffffffffffffffffffffffffffffffffffffff16905090565b60025481565b60045481565b60008054906101000a900473ffffffffffffffffffffffffffffffffffffffff1681565b6000600154905090565b60015481565b60008060006002546003546004549250925092509193909250565b6003548156fea26469706673582212207fb59f53ffcdad2c687354af1a47bc3edacd1830068127683a254006c33fc76764736f6c63430007060033`

	abiManager, err := abi.JSON(strings.NewReader(abiManagerJSON))
	require.NoError(t, err)

	rewardEpochDurationSeconds := big.NewInt(100)
	rewardEpochPowerHeight := big.NewInt(100)
	rewardEpochStartHeight := big.NewInt(100)
	rewardEpochStartTime := big.NewInt(100)
	ftsoManagerAddr, _, _, err := bind.DeployContract(auth, abiManager, common.FromHex(abiManagerBin), be, rewardAddr, rewardEpochDurationSeconds, rewardEpochPowerHeight, rewardEpochStartHeight, rewardEpochStartTime)
	require.NoError(t, err)

	// Registry
	abiFtsoRegistryJSON := `[{"inputs":[{"internalType":"uint256[]","name":"_supportedIndices","type":"uint256[]"}],"stateMutability":"nonpayable","type":"constructor"},{"inputs":[],"name":"getSupportedIndices","outputs":[{"internalType":"uint256[]","name":"","type":"uint256[]"}],"stateMutability":"view","type":"function"},{"inputs":[{"internalType":"uint256","name":"","type":"uint256"}],"name":"supportedIndices","outputs":[{"internalType":"uint256","name":"","type":"uint256"}],"stateMutability":"view","type":"function"}]`
	abiFtsoRegistryBin := `0x608060405234801561001057600080fd5b506040516102ee3803806102ee8339818101604052602081101561003357600080fd5b810190808051604051939291908464010000000082111561005357600080fd5b8382019150602082018581111561006957600080fd5b825186602082028301116401000000008211171561008657600080fd5b8083526020830192505050908051906020019060200280838360005b838110156100bd5780820151818401526020810190506100a2565b5050505090500160405250505080600090805190602001906100e09291906100e7565b5050610151565b828054828255906000526020600020908101928215610123579160200282015b82811115610122578251825591602001919060010190610107565b5b5090506101309190610134565b5090565b5b8082111561014d576000816000905550600101610135565b5090565b61018e806101606000396000f3fe608060405234801561001057600080fd5b50600436106100365760003560e01c80633b18bdcd1461003b578063798aac5b1461007d575b600080fd5b6100676004803603602081101561005157600080fd5b81019080803590602001909291905050506100dc565b6040518082815260200191505060405180910390f35b610085610100565b6040518080602001828103825283818151815260200191508051906020019060200280838360005b838110156100c85780820151818401526020810190506100ad565b505050509050019250505060405180910390f35b600081815481106100ec57600080fd5b906000526020600020016000915090505481565b6060600080548060200260200160405190810160405280929190818152602001828054801561014e57602002820191906000526020600020905b81548152602001906001019080831161013a575b505050505090509056fea2646970667358221220a76e31fdca4b4e549b9284ab4ea797a49c13a7ecab7f4100ea5d046ae6c3ef2464736f6c63430007060033`

	abiFtsoRegistry, err := abi.JSON(strings.NewReader(abiFtsoRegistryJSON))
	require.NoError(t, err)

	index1 := big.NewInt(123)
	index2 := big.NewInt(321)
	supportedIndices := []*big.Int{index1, index2}
	ftsoRegistryAddr, _, _, err := bind.DeployContract(auth, abiFtsoRegistry, common.FromHex(abiFtsoRegistryBin), be, supportedIndices)
	require.NoError(t, err)

	// Whitelist
	abiVoterWhitelisterJSON := `[{"inputs":[{"internalType":"uint256[]","name":"indices","type":"uint256[]"},{"internalType":"address[]","name":"priceProvidersAddresses","type":"address[]"}],"stateMutability":"nonpayable","type":"constructor"},{"inputs":[{"internalType":"uint256","name":"ftsoIndex","type":"uint256"}],"name":"getFtsoWhitelistedPriceProviders","outputs":[{"internalType":"address[]","name":"","type":"address[]"}],"stateMutability":"view","type":"function"},{"inputs":[{"internalType":"uint256","name":"","type":"uint256"},{"internalType":"uint256","name":"","type":"uint256"}],"name":"priceProviders","outputs":[{"internalType":"address","name":"","type":"address"}],"stateMutability":"view","type":"function"}]`
	abiVoterWhitelisterBin := `608060405234801561001057600080fd5b506040516104c53803806104c58339818101604052604081101561003357600080fd5b810190808051604051939291908464010000000082111561005357600080fd5b8382019150602082018581111561006957600080fd5b825186602082028301116401000000008211171561008657600080fd5b8083526020830192505050908051906020019060200280838360005b838110156100bd5780820151818401526020810190506100a2565b50505050905001604052602001805160405193929190846401000000008211156100e657600080fd5b838201915060208201858111156100fc57600080fd5b825186602082028301116401000000008211171561011957600080fd5b8083526020830192505050908051906020019060200280838360005b83811015610150578082015181840152602081019050610135565b505050509050016040525050506000825190508151811461017057600080fd5b60005b818110156101c2578260008086848151811061018b57fe5b6020026020010151815260200190815260200160002090805190602001906101b49291906101cb565b508080600101915050610173565b50505050610272565b828054828255906000526020600020908101928215610244579160200282015b828111156102435782518260006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff160217905550916020019190600101906101eb565b5b5090506102519190610255565b5090565b5b8082111561026e576000816000905550600101610256565b5090565b610244806102816000396000f3fe608060405234801561001057600080fd5b50600436106100365760003560e01c806309fcb4001461003b5780635a34e814146100be575b600080fd5b6100676004803603602081101561005157600080fd5b8101908080359060200190929190505050610120565b6040518080602001828103825283818151815260200191508051906020019060200280838360005b838110156100aa57808201518184015260208101905061008f565b505050509050019250505060405180910390f35b6100f4600480360360408110156100d457600080fd5b8101908080359060200190929190803590602001909291905050506101c0565b604051808273ffffffffffffffffffffffffffffffffffffffff16815260200191505060405180910390f35b60606000808381526020019081526020016000208054806020026020016040519081016040528092919081815260200182805480156101b457602002820191906000526020600020905b8160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff168152602001906001019080831161016a575b50505050509050919050565b600060205281600052604060002081815481106101dc57600080fd5b906000526020600020016000915091509054906101000a900473ffffffffffffffffffffffffffffffffffffffff168156fea264697066735822122096b13fe73494f9b3c29d2b4ee5689fc9b738c8a2b7b741123af130cffb34a9ea64736f6c63430007060033`

	abiVoterWhitelister, err := abi.JSON(strings.NewReader(abiVoterWhitelisterJSON))
	require.NoError(t, err)

	indices := []*big.Int{big.NewInt(123)}
	priceProvidersAddresses := []common.Address{common.BigToAddress(big.NewInt(222))}
	voterWhitelisterAddr, _, _, err := bind.DeployContract(auth, abiVoterWhitelister, common.FromHex(abiVoterWhitelisterBin), be, indices, priceProvidersAddresses)
	require.NoError(t, err)

	// Validation
	abiValidationJSON := `[{"inputs":[{"internalType":"address[]","name":"dataProvidersAddresses","type":"address[]"},{"internalType":"bytes20[]","name":"_nodes","type":"bytes20[]"}],"stateMutability":"nonpayable","type":"constructor"},{"inputs":[{"internalType":"address","name":"dataProvider","type":"address"}],"name":"getNodeIdForDataProvider","outputs":[{"internalType":"bytes20","name":"","type":"bytes20"}],"stateMutability":"view","type":"function"},{"inputs":[{"internalType":"address","name":"","type":"address"}],"name":"nodes","outputs":[{"internalType":"bytes20","name":"","type":"bytes20"}],"stateMutability":"view","type":"function"}]`
	abiValidationBin := `0x608060405234801561001057600080fd5b506040516103df3803806103df8339818101604052604081101561003357600080fd5b810190808051604051939291908464010000000082111561005357600080fd5b8382019150602082018581111561006957600080fd5b825186602082028301116401000000008211171561008657600080fd5b8083526020830192505050908051906020019060200280838360005b838110156100bd5780820151818401526020810190506100a2565b50505050905001604052602001805160405193929190846401000000008211156100e657600080fd5b838201915060208201858111156100fc57600080fd5b825186602082028301116401000000008211171561011957600080fd5b8083526020830192505050908051906020019060200280838360005b83811015610150578082015181840152602081019050610135565b505050509050016040525050506000815190508251811461017057600080fd5b60005b818110156102185782818151811061018757fe5b602002602001015160008086848151811061019e57fe5b602002602001015173ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908360601c02179055508080600101915050610173565b505050506101b48061022b6000396000f3fe608060405234801561001057600080fd5b50600436106100365760003560e01c8063177fbc711461003b578063189a5a17146100a2575b600080fd5b61007d6004803603602081101561005157600080fd5b81019080803573ffffffffffffffffffffffffffffffffffffffff169060200190929190505050610109565b60405180826bffffffffffffffffffffffff1916815260200191505060405180910390f35b6100e4600480360360208110156100b857600080fd5b81019080803573ffffffffffffffffffffffffffffffffffffffff16906020019092919050505061015e565b60405180826bffffffffffffffffffffffff1916815260200191505060405180910390f35b60008060008373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060009054906101000a900460601b9050919050565b60006020528060005260406000206000915054906101000a900460601b8156fea26469706673582212206fcff72ab7528e2f49e8ed7b227acd87df1719c83d37bf75dd10aaa2b8ed333f64736f6c63430007060033`

	abiValidation, err := abi.JSON(strings.NewReader(abiValidationJSON))
	require.NoError(t, err)

	providerAddresses := []common.Address{common.BigToAddress(big.NewInt(222))}
	nodes := [][20]byte{{1, 2}}
	addressValidation, _, _, err := bind.DeployContract(auth, abiValidation, common.FromHex(abiValidationBin), be, providerAddresses, nodes)
	require.NoError(t, err)

	// Submitter contract
	abiSubmitterJSON := `[{"inputs":[{"internalType":"address","name":"_voterWhitelister","type":"address"},{"internalType":"address","name":"_ftsoRegistry","type":"address"},{"internalType":"address","name":"_ftsoManager","type":"address"}],"stateMutability":"nonpayable","type":"constructor"},{"inputs":[],"name":"ftsoManager","outputs":[{"internalType":"address","name":"","type":"address"}],"stateMutability":"view","type":"function"},{"inputs":[],"name":"ftsoRegistry","outputs":[{"internalType":"address","name":"","type":"address"}],"stateMutability":"view","type":"function"},{"inputs":[],"name":"getFtsoManager","outputs":[{"internalType":"address","name":"","type":"address"}],"stateMutability":"view","type":"function"},{"inputs":[],"name":"getFtsoRegistry","outputs":[{"internalType":"address","name":"","type":"address"}],"stateMutability":"view","type":"function"},{"inputs":[],"name":"getVoterWhitelister","outputs":[{"internalType":"address","name":"","type":"address"}],"stateMutability":"view","type":"function"},{"inputs":[],"name":"voterWhitelister","outputs":[{"internalType":"address","name":"","type":"address"}],"stateMutability":"view","type":"function"}]`
	abiSubmitterBin := `0x608060405234801561001057600080fd5b506040516103ee3803806103ee8339818101604052606081101561003357600080fd5b81019080805190602001909291908051906020019092919080519060200190929190505050826000806101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff16021790555081600160006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff16021790555080600260006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff1602179055505050506102c28061012c6000396000f3fe608060405234801561001057600080fd5b50600436106100625760003560e01c806311a7aaaa1461006757806338b5f8691461009b57806371e1fad9146100cf5780638c9d28b614610103578063b39c685814610137578063c2b0d47b1461016b575b600080fd5b61006f61019f565b604051808273ffffffffffffffffffffffffffffffffffffffff16815260200191505060405180910390f35b6100a36101c5565b604051808273ffffffffffffffffffffffffffffffffffffffff16815260200191505060405180910390f35b6100d76101eb565b604051808273ffffffffffffffffffffffffffffffffffffffff16815260200191505060405180910390f35b61010b610214565b604051808273ffffffffffffffffffffffffffffffffffffffff16815260200191505060405180910390f35b61013f61023e565b604051808273ffffffffffffffffffffffffffffffffffffffff16815260200191505060405180910390f35b610173610268565b604051808273ffffffffffffffffffffffffffffffffffffffff16815260200191505060405180910390f35b600260009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1681565b600160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1681565b60008060009054906101000a900473ffffffffffffffffffffffffffffffffffffffff16905090565b6000600160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff16905090565b6000600260009054906101000a900473ffffffffffffffffffffffffffffffffffffffff16905090565b60008054906101000a900473ffffffffffffffffffffffffffffffffffffffff168156fea264697066735822122035d0b8aa6a058839885934e180e1ddb5af32130473692bb78427ccd544f1323f64736f6c63430007060033`

	abiSubmitter, err := abi.JSON(strings.NewReader(abiSubmitterJSON))
	require.NoError(t, err)

	submitterAddr, _, _, err := bind.DeployContract(auth, abiSubmitter, common.FromHex(abiSubmitterBin), be, voterWhitelisterAddr, ftsoRegistryAddr, ftsoManagerAddr)
	require.NoError(t, err)

	submitter := EVMContract{
		address: submitterAddr,
		abi:     abiSubmitter,
	}

	validation := EVMContract{
		address: addressValidation,
		abi:     abiValidation,
	}

	abis := FTSOABIs{
		Registry:  abiFtsoRegistry,
		Manager:   abiManager,
		Rewards:   abiReward,
		WNAT:      abiWnat,
		Whitelist: abiVoterWhitelister,
		Votepower: abiVotePower,
	}

	f := FTSOSystem{
		blockchain: be.Blockchain(),
		submitter:  submitter,
		validation: validation,
		abis:       abis,
	}

	return &f
}

func TestFTSOSystem_Contracts(t *testing.T) {
	key, err := crypto.GenerateKey()
	require.NoError(t, err)

	auth, err := bind.NewKeyedTransactorWithChainID(key, big.NewInt(1337))
	require.NoError(t, err)

	balance := new(big.Int)
	balance.SetUint64(math.MaxUint64)
	alloc := make(core.GenesisAlloc)
	alloc[auth.From] = core.GenesisAccount{
		Balance: balance,
	}
	var gasLimit uint64 = math.MaxUint64

	be := backends.NewSimulatedBackend(alloc, gasLimit)

	ftsoSystem := initDummyContracts(t, auth, be)
	_ = ftsoSystem

	be.Commit(true)

}
