package evm

import (
	"fmt"
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

var (
	// VotePower contract
	testAbiVotePowerJSON = `[{"inputs":[{"internalType":"uint256","name":"_vpInt","type":"uint256"}],"stateMutability":"nonpayable","type":"constructor"},{"inputs":[{"internalType":"address","name":"provider","type":"address"}],"name":"votePowerOf","outputs":[{"internalType":"uint256","name":"","type":"uint256"}],"stateMutability":"view","type":"function"},{"inputs":[],"name":"vpInt","outputs":[{"internalType":"uint256","name":"","type":"uint256"}],"stateMutability":"view","type":"function"}]`
	testAbiVotePowerBin  = `0x608060405234801561001057600080fd5b506040516101493803806101498339818101604052602081101561003357600080fd5b8101908080519060200190929190505050806000819055505060ef8061005a6000396000f3fe6080604052348015600f57600080fd5b506004361060325760003560e01c8063142d10181460375780638966b43214608c575b600080fd5b607660048036036020811015604b57600080fd5b81019080803573ffffffffffffffffffffffffffffffffffffffff16906020019092919050505060a8565b6040518082815260200191505060405180910390f35b609260b3565b6040518082815260200191505060405180910390f35b600080549050919050565b6000548156fea264697066735822122043d0eedff37b596f65fc5d3f9b98f8444bac96fcf123c979b81ece56c2991d2664736f6c63430007060033`
	testAbiVotePower, _  = abi.JSON(strings.NewReader(testAbiVotePowerJSON))

	// Wnat contract
	testAbiWnatJSON = `[{"inputs":[{"internalType":"address","name":"_wNatVal","type":"address"}],"stateMutability":"nonpayable","type":"constructor"},{"inputs":[],"name":"readVotePowerContract","outputs":[{"internalType":"address","name":"","type":"address"}],"stateMutability":"view","type":"function"},{"inputs":[],"name":"wNatVal","outputs":[{"internalType":"address","name":"","type":"address"}],"stateMutability":"view","type":"function"}]`
	testAbiWnatBin  = `0x608060405234801561001057600080fd5b506040516101b23803806101b28339818101604052602081101561003357600080fd5b8101908080519060200190929190505050806000806101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff1602179055505061011e806100946000396000f3fe6080604052348015600f57600080fd5b506004361060325760003560e01c80639b3baa0e146037578063f708a3b7146069575b600080fd5b603d609b565b604051808273ffffffffffffffffffffffffffffffffffffffff16815260200191505060405180910390f35b606f60c4565b604051808273ffffffffffffffffffffffffffffffffffffffff16815260200191505060405180910390f35b60008060009054906101000a900473ffffffffffffffffffffffffffffffffffffffff16905090565b60008054906101000a900473ffffffffffffffffffffffffffffffffffffffff168156fea2646970667358221220009e7049941af6c59bd17b62a091fd720d06545233e4227f893363177f29de7364736f6c63430007060033`
	testAbiWnat, _  = abi.JSON(strings.NewReader(testAbiWnatJSON))

	// Reward contract
	testAbiRewardJSON = `[{"inputs":[{"internalType":"address","name":"_wNatVal","type":"address"},{"internalType":"uint256","name":"_unclaimedReward","type":"uint256"}],"stateMutability":"nonpayable","type":"constructor"},{"inputs":[{"internalType":"uint256","name":"epoch","type":"uint256"},{"internalType":"address","name":"provider","type":"address"}],"name":"getUnclaimedReward","outputs":[{"internalType":"uint256","name":"","type":"uint256"}],"stateMutability":"view","type":"function"},{"inputs":[],"name":"unclaimedReward","outputs":[{"internalType":"uint256","name":"","type":"uint256"}],"stateMutability":"view","type":"function"},{"inputs":[],"name":"wNat","outputs":[{"internalType":"address","name":"","type":"address"}],"stateMutability":"view","type":"function"},{"inputs":[],"name":"wNatVal","outputs":[{"internalType":"address","name":"","type":"address"}],"stateMutability":"view","type":"function"}]`
	testAbiRewardBin  = `0x608060405234801561001057600080fd5b506040516102753803806102758339818101604052604081101561003357600080fd5b810190808051906020019092919080519060200190929190505050816000806101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff1602179055508060018190555050506101cf806100a66000396000f3fe608060405234801561001057600080fd5b506004361061004c5760003560e01c8063657d9695146100515780639edbf007146100b3578063cee83cee146100e7578063f708a3b714610105575b600080fd5b61009d6004803603604081101561006757600080fd5b8101908080359060200190929190803573ffffffffffffffffffffffffffffffffffffffff169060200190929190505050610139565b6040518082815260200191505060405180910390f35b6100bb610146565b604051808273ffffffffffffffffffffffffffffffffffffffff16815260200191505060405180910390f35b6100ef61016f565b6040518082815260200191505060405180910390f35b61010d610175565b604051808273ffffffffffffffffffffffffffffffffffffffff16815260200191505060405180910390f35b6000600154905092915050565b60008060009054906101000a900473ffffffffffffffffffffffffffffffffffffffff16905090565b60015481565b60008054906101000a900473ffffffffffffffffffffffffffffffffffffffff168156fea26469706673582212202de85ff60e4f8dc712b18a2ac0bc0b78455bfc02a03d036526eefa766354001664736f6c63430007060033`
	testAbiReward, _  = abi.JSON(strings.NewReader(testAbiRewardJSON))

	// Manager contract
	testAbiManagerJSON = `[{"inputs":[{"internalType":"address","name":"_rewardManager","type":"address"},{"internalType":"uint256","name":"_rewardEpochDurationSeconds","type":"uint256"},{"internalType":"uint256","name":"_rewardEpochPowerHeight","type":"uint256"},{"internalType":"uint256","name":"_rewardEpochStartHeight","type":"uint256"},{"internalType":"uint256","name":"_rewardEpochStartTime","type":"uint256"},{"internalType":"uint256","name":"_rewardEpochsStartTsVal","type":"uint256"}],"stateMutability":"nonpayable","type":"constructor"},{"inputs":[],"name":"durationSeconds","outputs":[{"internalType":"uint256","name":"","type":"uint256"}],"stateMutability":"view","type":"function"},{"inputs":[],"name":"rewardEpochDurationSeconds","outputs":[{"internalType":"uint256","name":"","type":"uint256"}],"stateMutability":"view","type":"function"},{"inputs":[],"name":"rewardEpochPowerHeight","outputs":[{"internalType":"uint256","name":"","type":"uint256"}],"stateMutability":"view","type":"function"},{"inputs":[],"name":"rewardEpochStartHeight","outputs":[{"internalType":"uint256","name":"","type":"uint256"}],"stateMutability":"view","type":"function"},{"inputs":[],"name":"rewardEpochStartTime","outputs":[{"internalType":"uint256","name":"","type":"uint256"}],"stateMutability":"view","type":"function"},{"inputs":[{"internalType":"uint256","name":"epoch","type":"uint256"}],"name":"rewardEpochs","outputs":[{"internalType":"uint256","name":"","type":"uint256"},{"internalType":"uint256","name":"","type":"uint256"},{"internalType":"uint256","name":"","type":"uint256"}],"stateMutability":"view","type":"function"},{"inputs":[],"name":"rewardEpochsStartTs","outputs":[{"internalType":"uint256","name":"","type":"uint256"}],"stateMutability":"view","type":"function"},{"inputs":[],"name":"rewardEpochsStartTsVal","outputs":[{"internalType":"uint256","name":"","type":"uint256"}],"stateMutability":"view","type":"function"},{"inputs":[],"name":"rewardManager","outputs":[{"internalType":"address","name":"","type":"address"}],"stateMutability":"view","type":"function"},{"inputs":[],"name":"rewardManagerVal","outputs":[{"internalType":"address","name":"","type":"address"}],"stateMutability":"view","type":"function"}]`
	testAbiManagerBin  = `0x608060405234801561001057600080fd5b506040516103eb3803806103eb833981810160405260c081101561003357600080fd5b81019080805190602001909291908051906020019092919080519060200190929190805190602001909291908051906020019092919080519060200190929190505050856000806101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff16021790555084600181905550836002819055508260038190555081600481905550806005819055505050505050506102fd806100ee6000396000f3fe608060405234801561001057600080fd5b506004361061009e5760003560e01c80639acba2af116100665780639acba2af14610165578063a578f55b14610183578063a795f409146101a1578063ab3caf41146101f1578063d5ea4a5c1461020f5761009e565b80630f4ef8a6146100a35780632f1b6b97146100d757806340e43150146100f5578063701a48ed1461011357806385f3c9c914610147575b600080fd5b6100ab61022d565b604051808273ffffffffffffffffffffffffffffffffffffffff16815260200191505060405180910390f35b6100df610256565b6040518082815260200191505060405180910390f35b6100fd61025c565b6040518082815260200191505060405180910390f35b61011b610262565b604051808273ffffffffffffffffffffffffffffffffffffffff16815260200191505060405180910390f35b61014f610286565b6040518082815260200191505060405180910390f35b61016d610290565b6040518082815260200191505060405180910390f35b61018b610296565b6040518082815260200191505060405180910390f35b6101cd600480360360208110156101b757600080fd5b81019080803590602001909291905050506102a0565b60405180848152602001838152602001828152602001935050505060405180910390f35b6101f96102bb565b6040518082815260200191505060405180910390f35b6102176102c1565b6040518082815260200191505060405180910390f35b60008060009054906101000a900473ffffffffffffffffffffffffffffffffffffffff16905090565b60025481565b60045481565b60008054906101000a900473ffffffffffffffffffffffffffffffffffffffff1681565b6000600154905090565b60015481565b6000600554905090565b60008060006002546003546004549250925092509193909250565b60055481565b6003548156fea2646970667358221220866af91a16ff6f65927d80f94fcb8949718a433aa06b7bba18a468e30981f92d64736f6c63430007060033`
	testAbiManager, _  = abi.JSON(strings.NewReader(testAbiManagerJSON))

	// Registry contract
	testAbiFtsoRegistryJSON = `[{"inputs":[{"internalType":"uint256[]","name":"_supportedIndices","type":"uint256[]"}],"stateMutability":"nonpayable","type":"constructor"},{"inputs":[],"name":"getSupportedIndices","outputs":[{"internalType":"uint256[]","name":"","type":"uint256[]"}],"stateMutability":"view","type":"function"},{"inputs":[{"internalType":"uint256","name":"","type":"uint256"}],"name":"supportedIndices","outputs":[{"internalType":"uint256","name":"","type":"uint256"}],"stateMutability":"view","type":"function"}]`
	testAbiFtsoRegistryBin  = `0x608060405234801561001057600080fd5b506040516102ee3803806102ee8339818101604052602081101561003357600080fd5b810190808051604051939291908464010000000082111561005357600080fd5b8382019150602082018581111561006957600080fd5b825186602082028301116401000000008211171561008657600080fd5b8083526020830192505050908051906020019060200280838360005b838110156100bd5780820151818401526020810190506100a2565b5050505090500160405250505080600090805190602001906100e09291906100e7565b5050610151565b828054828255906000526020600020908101928215610123579160200282015b82811115610122578251825591602001919060010190610107565b5b5090506101309190610134565b5090565b5b8082111561014d576000816000905550600101610135565b5090565b61018e806101606000396000f3fe608060405234801561001057600080fd5b50600436106100365760003560e01c80633b18bdcd1461003b578063798aac5b1461007d575b600080fd5b6100676004803603602081101561005157600080fd5b81019080803590602001909291905050506100dc565b6040518082815260200191505060405180910390f35b610085610100565b6040518080602001828103825283818151815260200191508051906020019060200280838360005b838110156100c85780820151818401526020810190506100ad565b505050509050019250505060405180910390f35b600081815481106100ec57600080fd5b906000526020600020016000915090505481565b6060600080548060200260200160405190810160405280929190818152602001828054801561014e57602002820191906000526020600020905b81548152602001906001019080831161013a575b505050505090509056fea2646970667358221220a76e31fdca4b4e549b9284ab4ea797a49c13a7ecab7f4100ea5d046ae6c3ef2464736f6c63430007060033`
	testAbiFtsoRegistry, _  = abi.JSON(strings.NewReader(testAbiFtsoRegistryJSON))

	// Whitelist contract
	testAbiVoterWhitelisterJSON = `[{"inputs":[{"internalType":"uint256[]","name":"indices","type":"uint256[]"},{"internalType":"address[]","name":"priceProvidersAddresses","type":"address[]"}],"stateMutability":"nonpayable","type":"constructor"},{"inputs":[{"internalType":"uint256","name":"ftsoIndex","type":"uint256"}],"name":"getFtsoWhitelistedPriceProviders","outputs":[{"internalType":"address[]","name":"","type":"address[]"}],"stateMutability":"view","type":"function"},{"inputs":[{"internalType":"uint256","name":"","type":"uint256"},{"internalType":"uint256","name":"","type":"uint256"}],"name":"priceProviders","outputs":[{"internalType":"address","name":"","type":"address"}],"stateMutability":"view","type":"function"}]`
	testAbiVoterWhitelisterBin  = `608060405234801561001057600080fd5b506040516104c53803806104c58339818101604052604081101561003357600080fd5b810190808051604051939291908464010000000082111561005357600080fd5b8382019150602082018581111561006957600080fd5b825186602082028301116401000000008211171561008657600080fd5b8083526020830192505050908051906020019060200280838360005b838110156100bd5780820151818401526020810190506100a2565b50505050905001604052602001805160405193929190846401000000008211156100e657600080fd5b838201915060208201858111156100fc57600080fd5b825186602082028301116401000000008211171561011957600080fd5b8083526020830192505050908051906020019060200280838360005b83811015610150578082015181840152602081019050610135565b505050509050016040525050506000825190508151811461017057600080fd5b60005b818110156101c2578260008086848151811061018b57fe5b6020026020010151815260200190815260200160002090805190602001906101b49291906101cb565b508080600101915050610173565b50505050610272565b828054828255906000526020600020908101928215610244579160200282015b828111156102435782518260006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff160217905550916020019190600101906101eb565b5b5090506102519190610255565b5090565b5b8082111561026e576000816000905550600101610256565b5090565b610244806102816000396000f3fe608060405234801561001057600080fd5b50600436106100365760003560e01c806309fcb4001461003b5780635a34e814146100be575b600080fd5b6100676004803603602081101561005157600080fd5b8101908080359060200190929190505050610120565b6040518080602001828103825283818151815260200191508051906020019060200280838360005b838110156100aa57808201518184015260208101905061008f565b505050509050019250505060405180910390f35b6100f4600480360360408110156100d457600080fd5b8101908080359060200190929190803590602001909291905050506101c0565b604051808273ffffffffffffffffffffffffffffffffffffffff16815260200191505060405180910390f35b60606000808381526020019081526020016000208054806020026020016040519081016040528092919081815260200182805480156101b457602002820191906000526020600020905b8160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff168152602001906001019080831161016a575b50505050509050919050565b600060205281600052604060002081815481106101dc57600080fd5b906000526020600020016000915091509054906101000a900473ffffffffffffffffffffffffffffffffffffffff168156fea264697066735822122096b13fe73494f9b3c29d2b4ee5689fc9b738c8a2b7b741123af130cffb34a9ea64736f6c63430007060033`
	testAbiVoterWhitelister, _  = abi.JSON(strings.NewReader(testAbiVoterWhitelisterJSON))

	// Validation contract
	testAbiValidationJSON = `[{"inputs":[{"internalType":"address[]","name":"dataProvidersAddresses","type":"address[]"},{"internalType":"bytes20[]","name":"_nodes","type":"bytes20[]"}],"stateMutability":"nonpayable","type":"constructor"},{"inputs":[{"internalType":"address","name":"dataProvider","type":"address"}],"name":"getNodeIdForDataProvider","outputs":[{"internalType":"bytes20","name":"","type":"bytes20"}],"stateMutability":"view","type":"function"},{"inputs":[{"internalType":"address","name":"","type":"address"}],"name":"nodes","outputs":[{"internalType":"bytes20","name":"","type":"bytes20"}],"stateMutability":"view","type":"function"}]`
	testAbiValidationBin  = `0x608060405234801561001057600080fd5b506040516103df3803806103df8339818101604052604081101561003357600080fd5b810190808051604051939291908464010000000082111561005357600080fd5b8382019150602082018581111561006957600080fd5b825186602082028301116401000000008211171561008657600080fd5b8083526020830192505050908051906020019060200280838360005b838110156100bd5780820151818401526020810190506100a2565b50505050905001604052602001805160405193929190846401000000008211156100e657600080fd5b838201915060208201858111156100fc57600080fd5b825186602082028301116401000000008211171561011957600080fd5b8083526020830192505050908051906020019060200280838360005b83811015610150578082015181840152602081019050610135565b505050509050016040525050506000815190508251811461017057600080fd5b60005b818110156102185782818151811061018757fe5b602002602001015160008086848151811061019e57fe5b602002602001015173ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908360601c02179055508080600101915050610173565b505050506101b48061022b6000396000f3fe608060405234801561001057600080fd5b50600436106100365760003560e01c8063177fbc711461003b578063189a5a17146100a2575b600080fd5b61007d6004803603602081101561005157600080fd5b81019080803573ffffffffffffffffffffffffffffffffffffffff169060200190929190505050610109565b60405180826bffffffffffffffffffffffff1916815260200191505060405180910390f35b6100e4600480360360208110156100b857600080fd5b81019080803573ffffffffffffffffffffffffffffffffffffffff16906020019092919050505061015e565b60405180826bffffffffffffffffffffffff1916815260200191505060405180910390f35b60008060008373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060009054906101000a900460601b9050919050565b60006020528060005260406000206000915054906101000a900460601b8156fea26469706673582212206fcff72ab7528e2f49e8ed7b227acd87df1719c83d37bf75dd10aaa2b8ed333f64736f6c63430007060033`
	testAbiValidation, _  = abi.JSON(strings.NewReader(testAbiValidationJSON))

	// Submitter contract
	testAbiSubmitterJSON = `[{"inputs":[{"internalType":"address","name":"_voterWhitelister","type":"address"},{"internalType":"address","name":"_ftsoRegistry","type":"address"},{"internalType":"address","name":"_ftsoManager","type":"address"}],"stateMutability":"nonpayable","type":"constructor"},{"inputs":[],"name":"ftsoManager","outputs":[{"internalType":"address","name":"","type":"address"}],"stateMutability":"view","type":"function"},{"inputs":[],"name":"ftsoRegistry","outputs":[{"internalType":"address","name":"","type":"address"}],"stateMutability":"view","type":"function"},{"inputs":[],"name":"getFtsoManager","outputs":[{"internalType":"address","name":"","type":"address"}],"stateMutability":"view","type":"function"},{"inputs":[],"name":"getFtsoRegistry","outputs":[{"internalType":"address","name":"","type":"address"}],"stateMutability":"view","type":"function"},{"inputs":[],"name":"getVoterWhitelister","outputs":[{"internalType":"address","name":"","type":"address"}],"stateMutability":"view","type":"function"},{"inputs":[],"name":"voterWhitelister","outputs":[{"internalType":"address","name":"","type":"address"}],"stateMutability":"view","type":"function"}]`
	testAbiSubmitterBin  = `0x608060405234801561001057600080fd5b506040516103ee3803806103ee8339818101604052606081101561003357600080fd5b81019080805190602001909291908051906020019092919080519060200190929190505050826000806101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff16021790555081600160006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff16021790555080600260006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff1602179055505050506102c28061012c6000396000f3fe608060405234801561001057600080fd5b50600436106100625760003560e01c806311a7aaaa1461006757806338b5f8691461009b57806371e1fad9146100cf5780638c9d28b614610103578063b39c685814610137578063c2b0d47b1461016b575b600080fd5b61006f61019f565b604051808273ffffffffffffffffffffffffffffffffffffffff16815260200191505060405180910390f35b6100a36101c5565b604051808273ffffffffffffffffffffffffffffffffffffffff16815260200191505060405180910390f35b6100d76101eb565b604051808273ffffffffffffffffffffffffffffffffffffffff16815260200191505060405180910390f35b61010b610214565b604051808273ffffffffffffffffffffffffffffffffffffffff16815260200191505060405180910390f35b61013f61023e565b604051808273ffffffffffffffffffffffffffffffffffffffff16815260200191505060405180910390f35b610173610268565b604051808273ffffffffffffffffffffffffffffffffffffffff16815260200191505060405180910390f35b600260009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1681565b600160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1681565b60008060009054906101000a900473ffffffffffffffffffffffffffffffffffffffff16905090565b6000600160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff16905090565b6000600260009054906101000a900473ffffffffffffffffffffffffffffffffffffffff16905090565b60008054906101000a900473ffffffffffffffffffffffffffffffffffffffff168156fea264697066735822122035d0b8aa6a058839885934e180e1ddb5af32130473692bb78427ccd544f1323f64736f6c63430007060033`
	testAbiSubmitter, _  = abi.JSON(strings.NewReader(testAbiSubmitterJSON))
)

func TestFTSOSystem_Contracts(t *testing.T) {
	t.Run("common case", func(t *testing.T) {
		auth, be := simulatedBlockchain(t)

		// VotePower contract
		vpInt := big.NewInt(100)
		votePowerAddr := deployTestContract(t, auth, be, testAbiVotePower, testAbiVotePowerBin, vpInt)

		// Wnat contract
		wnatAddr := deployTestContract(t, auth, be, testAbiWnat, testAbiWnatBin, votePowerAddr)

		// Reward contract
		unclaimedReward := big.NewInt(100)
		rewardAddr := deployTestContract(t, auth, be, testAbiReward, testAbiRewardBin, wnatAddr, unclaimedReward)

		// Manager contract
		rewardEpochDurationSeconds := big.NewInt(100)
		rewardEpochPowerHeight := big.NewInt(100)
		rewardEpochStartHeight := big.NewInt(100)
		rewardEpochStartTime := big.NewInt(100)
		rewardEpochsStartTs := big.NewInt(100)
		ftsoManagerAddr := deployTestContract(t, auth, be, testAbiManager, testAbiManagerBin, rewardAddr, rewardEpochDurationSeconds, rewardEpochPowerHeight, rewardEpochStartHeight, rewardEpochStartTime, rewardEpochsStartTs)

		// Registry contract
		index1 := big.NewInt(123)
		index2 := big.NewInt(321)
		supportedIndices := []*big.Int{index1, index2}
		ftsoRegistryAddr := deployTestContract(t, auth, be, testAbiFtsoRegistry, testAbiFtsoRegistryBin, supportedIndices)

		// Whitelist contract
		indices := []*big.Int{big.NewInt(123)}
		priceProvidersAddresses := []common.Address{common.BigToAddress(big.NewInt(222))}
		voterWhitelisterAddr := deployTestContract(t, auth, be, testAbiVoterWhitelister, testAbiVoterWhitelisterBin, indices, priceProvidersAddresses)

		// Validation contract
		providerAddresses := []common.Address{common.BigToAddress(big.NewInt(222))}
		nodes := [][20]byte{{1, 2}}
		validationAddr := deployTestContract(t, auth, be, testAbiValidation, testAbiValidationBin, providerAddresses, nodes)

		// Submitter contract
		submitterAddr := deployTestContract(t, auth, be, testAbiSubmitter, testAbiSubmitterBin, voterWhitelisterAddr, ftsoRegistryAddr, ftsoManagerAddr)

		be.Commit(true)

		ftsoSystem := testFTSOSystem(t, be, submitterAddr, validationAddr)

		latestBlock := be.Blockchain().LastAcceptedBlock()

		contracts, err := ftsoSystem.Contracts(latestBlock.Hash())

		fmt.Println(contracts)
		fmt.Println(err)
	})

}

func testFTSOSystem(t *testing.T, be *backends.SimulatedBackend, submitterAddr, validationAddr common.Address) *FTSOSystem {
	t.Helper()

	submitter := EVMContract{
		address: submitterAddr,
		abi:     testAbiSubmitter,
	}

	validation := EVMContract{
		address: validationAddr,
		abi:     testAbiValidation,
	}

	abis := FTSOABIs{
		Registry:  testAbiFtsoRegistry,
		Manager:   testAbiManager,
		Rewards:   testAbiReward,
		WNAT:      testAbiWnat,
		Whitelist: testAbiVoterWhitelister,
		Votepower: testAbiVotePower,
	}

	f := FTSOSystem{
		blockchain: be.Blockchain(),
		submitter:  submitter,
		validation: validation,
		abis:       abis,
	}

	return &f
}

func deployTestContract(t *testing.T, auth *bind.TransactOpts, be *backends.SimulatedBackend, testAbi abi.ABI, testAbiBin string, params ...interface{}) common.Address {
	t.Helper()

	addr, _, _, err := bind.DeployContract(auth, testAbi, common.FromHex(testAbiBin), be, params...)
	require.NoError(t, err)

	return addr
}

func simulatedBlockchain(t *testing.T) (*bind.TransactOpts, *backends.SimulatedBackend) {
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

	return auth, be
}
