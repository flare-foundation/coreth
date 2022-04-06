// (c) 2019-2020, Flare Networks Limited.

package misc

import (
	"github.com/ethereum/go-ethereum/common"
	"github.com/flare-foundation/coreth/core/state"
	"github.com/flare-foundation/coreth/params"
)

var stateConnectorUpgrade = CoreContractUpdate{
	Address: params.ConnectorAddress,
	Old:     common.FromHex("0x608060405234801561001057600080fd5b50600436106101735760003560e01c80637ec93e9f116100de578063c5d64cd111610097578063dccb2d3211610071578063dccb2d32146107f2578063ee2d87371461081e578063ef4c169e1461083b578063f9c490ee1461084357610173565b8063c5d64cd11461073e578063d3b92926146107b4578063d3fb3e9f146107d157610173565b80637ec93e9f146105a85780637f582432146105d45780638abd90ae1461069c5780638b203dd4146106a45780639c531f8c1461072e578063adc890321461073657610173565b806347453f371161013057806347453f37146104355780634bdc9c8f1461045257806355d14c15146104d857806371e8d61a14610504578063741d3cb41461057457806379fd4e1a1461057c57610173565b806307003bb41461017857806322ce7387146101945780632a2434a2146101cc5780632bec6f8714610247578063388492dd1461028f5780633f57987d14610403575b600080fd5b61018061084b565b604080519115158252519081900360200190f35b6101ba600480360360208110156101aa57600080fd5b50356001600160a01b0316610854565b60408051918252519081900360200190f35b6101ef600480360360208110156101e257600080fd5b503563ffffffff16610866565b60405180876001600160401b03168152602001866001600160401b031681526020018561ffff168152602001846001600160401b03168152602001838152602001828152602001965050505050505060405180910390f35b6102736004803603604081101561025d57600080fd5b506001600160a01b038135169060200135610a98565b604080516001600160401b039092168252519081900360200190f35b610357600480360360808110156102a557600080fd5b63ffffffff823516916020810135916001600160401b0360408301351691908101906080810160608201356401000000008111156102e257600080fd5b8201836020820111156102f457600080fd5b8035906020019184600183028401116401000000008311171561031657600080fd5b91908080601f016020809104026020016040519081016040528093929190818152602001838380828437600092019190915250929550610abe945050505050565b604051808663ffffffff168152602001856001600160401b03168152602001846001600160401b0316815260200183815260200180602001828103825283818151815260200191508051906020019080838360005b838110156103c45781810151838201526020016103ac565b50505050905090810190601f1680156103f15780820380516001836020036101000a031916815260200191505b50965050505050505060405180910390f35b6101806004803603604081101561041957600080fd5b50803563ffffffff1690602001356001600160401b03166114cb565b6102736004803603602081101561044b57600080fd5b50356115c7565b61046f6004803603602081101561046857600080fd5b5035611625565b604080519a15158b5260208b0199909952898901979097526060890195909552608088019390935260a08701919091526001600160401b0390811660c08701521660e085015215156101008401526001600160a01b031661012083015251908190036101400190f35b61046f600480360360408110156104ee57600080fd5b50803590602001356001600160401b031661169a565b610548600480360360a081101561051a57600080fd5b5063ffffffff813516906020810135906040810135906001600160401b036060820135169060800135611718565b604080516001600160401b03948516815292909316602083015215158183015290519081900360600190f35b6101ba61194b565b61046f6004803603604081101561059257600080fd5b50803590602001356001600160401b0316611951565b61046f600480360360408110156105be57600080fd5b506001600160a01b0381351690602001356119cf565b610357600480360360808110156105ea57600080fd5b63ffffffff823516916020810135916001600160401b03604083013516919081019060808101606082013564010000000081111561062757600080fd5b82018360208201111561063957600080fd5b8035906020019184600183028401116401000000008311171561065b57600080fd5b91908080601f016020809104026020016040519081016040528093929190818152602001838380828437600092019190915250929550611a4c945050505050565b6101ba612460565b6106c7600480360360208110156106ba57600080fd5b503563ffffffff16612466565b604080519a15158b526001600160401b03998a1660208c01529789168a89015261ffff96871660608b015294909516608089015291861660a088015290941660c086015260e085019390935261010084015261012083019190915251908190036101400190f35b6102736124d6565b6101ba6124e5565b61077c6004803603608081101561075457600080fd5b5063ffffffff813516906001600160401b036020820135169060408101359060600135612588565b6040805163ffffffff90951685526001600160401b03909316602085015261ffff909116838301526060830152519081900360800190f35b610273600480360360208110156107ca57600080fd5b503561375d565b6107d9613778565b6040805163ffffffff9092168252519081900360200190f35b6102736004803603604081101561080857600080fd5b506001600160a01b038135169060200135613789565b61046f6004803603602081101561083457600080fd5b50356137fd565b610180613870565b6101ba614280565b60005460ff1681565b600d6020526000908152604090205481565b600080548190819081908190819060ff166108b25760405162461bcd60e51b815260040180806020018281038252603a8152602001806144ee603a913960400191505060405180910390fd5b63ffffffff8716600090815260056020526040902054879060ff1661090c576040805162461bcd60e51b815260206004820152601660248201526000805160206145b7833981519152604482015290519081900360640190fd5b600560008963ffffffff1663ffffffff168152602001908152602001600020600201549250600560008963ffffffff1663ffffffff168152602001908152602001600020600401549150600088600560008b63ffffffff1663ffffffff16815260200190815260200160002060000160159054906101000a90046001600160401b0316604051602001808363ffffffff1660e01b8152600401826001600160401b031660c01b81526008019250505060405160208183030381529060405280519060200120905060066000336001600160a01b03166001600160a01b03168152602001908152602001600020600082815260200190815260200160002060000160009054906101000a900460ff1615610a435733600090815260066020908152604080832084845290915281206004015490945092505b505063ffffffff969096166000908152600560205260409020805460019091015461010082046001600160401b0390811699600160a81b8404821699600160881b90940461ffff169850911695509193509150565b600b6020908152600092835260408084209091529082529020546001600160401b031681565b6000805481908190819060609060ff16610b095760405162461bcd60e51b815260040180806020018281038252603a8152602001806144ee603a913960400191505060405180910390fd5b63ffffffff8916600090815260056020526040902054899060ff16610b63576040805162461bcd60e51b815260206004820152601660248201526000805160206145b7833981519152604482015290519081900360640190fd5b88610baa576040805162461bcd60e51b815260206004820152601260248201527107061796d656e7448617368203d3d203078360741b604482015290519081900360640190fd5b63ffffffff8a166000908152600560205260409020600101546001600160401b0316610c075760405162461bcd60e51b81526004018080602001828103825260298152602001806144706029913960400191505060405180910390fd5b41331480610c18575041600160981b145b610c69576040805162461bcd60e51b815260206004820152601c60248201527f696e76616c696420626c6f636b2e636f696e626173652076616c756500000000604482015290519081900360640190fd5b6040805160e08c901b6001600160e01b03191660208083019190915260248083018d905283518084039091018152604490920183528151918101919091206000818152600a90925291902060060154600160801b900460ff1615610d0d576040805162461bcd60e51b81526020600482015260166024820152753830bcb6b2b73a1030b63932b0b23c90383937bb32b760511b604482015290519081900360640190fd5b63ffffffff8b166000908152600560205260409020600101546001600160401b03908116908a1610610d705760405162461bcd60e51b815260040180806020018281038252602e815260200180614442602e913960400191505060405180910390fd5b63ffffffff8b166000908152600560205260409020546001600160401b036101008204811691600160481b90041615610e515763ffffffff8c16600090815260056020526040902080546001909101546001600160401b03600160481b830481166101009093048116918116919091031611610e1d5760405162461bcd60e51b815260040180806020018281038252603981526020018061457e6039913960400191505060405180910390fd5b5063ffffffff8b16600090815260056020526040902080546001909101546001600160401b03600160481b90920482169116035b806001600160401b03168a6001600160401b03161015610eb8576040805162461bcd60e51b815260206004820152601a60248201527f6c6564676572203c20696e646578536561726368526567696f6e000000000000604482015290519081900360640190fd5b60008281526008602090815260408083206001600160401b038e16845290915281205460ff16156110415760008381526008602090815260408083206001600160401b038f168452909152902060040154421015610f475760405162461bcd60e51b81526004018080602001828103825260498152602001806143656049913960600191505060405180910390fd5b60008381526008602090815260408083206001600160401b038f1684529091529020600501548c14610fb6576040805162461bcd60e51b81526020600482015260136024820152720d2dcecc2d8d2c840e0c2f2dacadce890c2e6d606b1b604482015290519081900360640190fd5b426003546008600086815260200190815260200160002060008e6001600160401b03166001600160401b0316815260200190815260200160002060030154011161103c576040805162461bcd60e51b815260206004820152601260248201527172657665616c20697320746f6f206c61746560701b604482015290519081900360640190fd5b61107d565b413314801590611054575041600160981b145b1561107d575063ffffffff8c166000908152600560205260409020600101546001600160401b03165b4133148015611090575041600160981b14155b156114ba5760008381526008602090815260408083206001600160401b038f16845290915290205460ff16611269576040518061014001604052806001151581526020016000801b81526020016000801b8152602001428152602001600254420181526020018d81526020018c6001600160401b03168152602001836001600160401b03168152602001600015158152602001336001600160a01b03168152506008600085815260200190815260200160002060008d6001600160401b03166001600160401b0316815260200190815260200160002060008201518160000160006101000a81548160ff0219169083151502179055506020820151816001015560408201518160020155606082015181600301556080820151816004015560a0820151816005015560c08201518160060160006101000a8154816001600160401b0302191690836001600160401b0316021790555060e08201518160060160086101000a8154816001600160401b0302191690836001600160401b031602179055506101008201518160060160106101000a81548160ff0219169083151502179055506101208201518160070160006101000a8154816001600160a01b0302191690836001600160a01b031602179055509050506114ba565b6040518061014001604052806001151581526020016000801b81526020016000801b8152602001600081526020014281526020018d81526020018c6001600160401b03168152602001836001600160401b03168152602001600115158152602001336001600160a01b0316815250600a600085815260200190815260200160002060008201518160000160006101000a81548160ff0219169083151502179055506020820151816001015560408201518160020155606082015181600301556080820151816004015560a0820151816005015560c08201518160060160006101000a8154816001600160401b0302191690836001600160401b0316021790555060e08201518160060160086101000a8154816001600160401b0302191690836001600160401b031602179055506101008201518160060160106101000a81548160ff0219169083151502179055506101208201518160070160006101000a8154816001600160a01b0302191690836001600160a01b031602179055509050507ff794f5b9b6791b3ab0eb97342cf2efbb7e710a478d4512a2505ea60c93fd85d38d8c8c8f33604051808663ffffffff168152602001856001600160401b0316815260200180602001848152602001836001600160a01b03168152602001828103825285818151815260200191508051906020019080838360005b8381101561147b578181015183820152602001611463565b50505050905090810190601f1680156114a85780820380516001836020036101000a031916815260200191505b50965050505050505060405180910390a15b9b9c999b9a99975050505050505050565b6000805460ff1661150d5760405162461bcd60e51b815260040180806020018281038252603a8152602001806144ee603a913960400191505060405180910390fd5b63ffffffff8316600090815260056020526040902054839060ff16611567576040805162461bcd60e51b815260206004820152601660248201526000805160206145b7833981519152604482015290519081900360640190fd5b50506040805160e09390931b6001600160e01b03191660208085019190915260c09290921b6001600160c01b03191660248401528051808403600c018152602c909301815282519282019290922060009081526007909152205460ff1690565b6000805460ff166116095760405162461bcd60e51b815260040180806020018281038252603a8152602001806144ee603a913960400191505060405180910390fd5b506000908152600c60205260409020546001600160401b031690565b60076020819052600091825260409091208054600182015460028301546003840154600485015460058601546006870154969097015460ff95861697949693959294919391926001600160401b0380841693600160401b810490911692600160801b909104909116906001600160a01b03168a565b60086020908152600092835260408084209091529082529020805460018201546002830154600384015460048501546005860154600687015460079097015460ff96871697959694959394929391926001600160401b0380841693600160401b810490911692600160801b909104909116906001600160a01b03168a565b600080548190819060ff1661175e5760405162461bcd60e51b815260040180806020018281038252603a8152602001806144ee603a913960400191505060405180910390fd5b63ffffffff8816600090815260056020526040902054889060ff166117b8576040805162461bcd60e51b815260206004820152601660248201526000805160206145b7833981519152604482015290519081900360640190fd5b604080516001600160401b0388166020808301919091528251808303820181528284018452805190820120606083018c9052608083018b905260a083015260c08083018990528351808403909101815260e08084018552815191830191909120908d901b6001600160e01b0319166101008401526101048084018290528451808503909101815261012490930184528251928201929092206000818152600a9092529290205490919060ff166118ae576040805162461bcd60e51b81526020600482015260166024820152751c185e5b595b9d08191bd95cc81b9bdd08195e1a5cdd60521b604482015290519081900360640190fd5b6000818152600a60205260409020600501548214611909576040805162461bcd60e51b81526020600482015260136024820152720d2dcecc2d8d2c840e0c2f2dacadce890c2e6d606b1b604482015290519081900360640190fd5b6000908152600a60205260409020600601546001600160401b038082169c600160401b83049091169b50600160801b90910460ff169950975050505050505050565b60025481565b60096020908152600092835260408084209091529082529020805460018201546002830154600384015460048501546005860154600687015460079097015460ff96871697959694959394929391926001600160401b0380841693600160401b810490911692600160801b909104909116906001600160a01b03168a565b600660208181526000938452604080852090915291835291208054600182015460028301546003840154600485015460058601549686015460079096015460ff95861697949693959294919391926001600160401b0380841693600160401b810490911692600160801b909104909116906001600160a01b03168a565b6000805481908190819060609060ff16611a975760405162461bcd60e51b815260040180806020018281038252603a8152602001806144ee603a913960400191505060405180910390fd5b63ffffffff8916600090815260056020526040902054899060ff16611af1576040805162461bcd60e51b815260206004820152601660248201526000805160206145b7833981519152604482015290519081900360640190fd5b88611b38576040805162461bcd60e51b815260206004820152601260248201527107061796d656e7448617368203d3d203078360741b604482015290519081900360640190fd5b63ffffffff8a166000908152600560205260409020600101546001600160401b0316611b955760405162461bcd60e51b81526004018080602001828103825260298152602001806144706029913960400191505060405180910390fd5b41331480611ba6575041600160981b145b611bf7576040805162461bcd60e51b815260206004820152601c60248201527f696e76616c696420626c6f636b2e636f696e626173652076616c756500000000604482015290519081900360640190fd5b6040805160e08c901b6001600160e01b03191660208083019190915260248083018d905283518084039091018152604490920183528151918101919091206000818152600a90925291902060060154600160801b900460ff1615611c98576040805162461bcd60e51b81526020600482015260136024820152723a3c24b21030b63932b0b23c90383937bb32b760691b604482015290519081900360640190fd5b6000818152600a60205260409020600601546001600160401b03808b16911610611cf35760405162461bcd60e51b815260040180806020018281038252602f8152602001806143ae602f913960400191505060405180910390fd5b63ffffffff8b166000908152600560205260409020600101546001600160401b03908116908a1610611d565760405162461bcd60e51b815260040180806020018281038252602e815260200180614442602e913960400191505060405180910390fd5b63ffffffff8b166000908152600560205260409020546001600160401b036101008204811691600160481b90041615611e375763ffffffff8c16600090815260056020526040902080546001909101546001600160401b03600160481b830481166101009093048116918116919091031611611e035760405162461bcd60e51b815260040180806020018281038252603981526020018061457e6039913960400191505060405180910390fd5b5063ffffffff8b16600090815260056020526040902080546001909101546001600160401b03600160481b90920482169116035b806001600160401b03168a6001600160401b03161015611e9e576040805162461bcd60e51b815260206004820152601a60248201527f6c6564676572203c20696e646578536561726368526567696f6e000000000000604482015290519081900360640190fd5b60008281526009602090815260408083206001600160401b038e16845290915281205460ff16156120275760008381526009602090815260408083206001600160401b038f168452909152902060040154421015611f2d5760405162461bcd60e51b81526004018080602001828103825260548152602001806142876054913960600191505060405180910390fd5b60008381526009602090815260408083206001600160401b038f1684529091529020600501548c14611f9c576040805162461bcd60e51b81526020600482015260136024820152720d2dcecc2d8d2c840e0c2f2dacadce890c2e6d606b1b604482015290519081900360640190fd5b426003546009600086815260200190815260200160002060008e6001600160401b03166001600160401b03168152602001908152602001600020600301540111612022576040805162461bcd60e51b815260206004820152601260248201527172657665616c20697320746f6f206c61746560701b604482015290519081900360640190fd5b612063565b41331480159061203a575041600160981b145b15612063575063ffffffff8c166000908152600560205260409020600101546001600160401b03165b4133148015612076575041600160981b14155b156114ba5760008381526009602090815260408083206001600160401b038f16845290915290205460ff1661224f576040518061014001604052806001151581526020016000801b81526020016000801b8152602001428152602001600254420181526020018d81526020018c6001600160401b03168152602001836001600160401b03168152602001600015158152602001336001600160a01b03168152506009600085815260200190815260200160002060008d6001600160401b03166001600160401b0316815260200190815260200160002060008201518160000160006101000a81548160ff0219169083151502179055506020820151816001015560408201518160020155606082015181600301556080820151816004015560a0820151816005015560c08201518160060160006101000a8154816001600160401b0302191690836001600160401b0316021790555060e08201518160060160086101000a8154816001600160401b0302191690836001600160401b031602179055506101008201518160060160106101000a81548160ff0219169083151502179055506101208201518160070160006101000a8154816001600160a01b0302191690836001600160a01b031602179055509050506114ba565b6040518061014001604052806001151581526020016000801b81526020016000801b8152602001600081526020014281526020018d81526020018c6001600160401b03168152602001836001600160401b03168152602001600015158152602001336001600160a01b0316815250600a600085815260200190815260200160002060008201518160000160006101000a81548160ff0219169083151502179055506020820151816001015560408201518160020155606082015181600301556080820151816004015560a0820151816005015560c08201518160060160006101000a8154816001600160401b0302191690836001600160401b0316021790555060e08201518160060160086101000a8154816001600160401b0302191690836001600160401b031602179055506101008201518160060160106101000a81548160ff0219169083151502179055506101208201518160070160006101000a8154816001600160a01b0302191690836001600160a01b031602179055509050507ff794f5b9b6791b3ab0eb97342cf2efbb7e710a478d4512a2505ea60c93fd85d38d8c8c8f33604051808663ffffffff168152602001856001600160401b0316815260200180602001848152602001836001600160a01b03168152602001828103825285818151815260200191508051906020019080838360008381101561147b578181015183820152602001611463565b60035481565b6005602052600090815260409020805460018201546002830154600384015460049094015460ff8416946001600160401b036101008604811695600160481b810482169561ffff600160881b8304811696600160981b840490911695600160a81b9093048416949190931692918a565b6004546001600160401b031681565b6000805460ff166125275760405162461bcd60e51b815260040180806020018281038252603a8152602001806144ee603a913960400191505060405180910390fd5b60015442116125675760405162461bcd60e51b81526004018080602001828103825260218152602001806145286021913960400191505060405180910390fd5b6004546001546001600160401b039091169042038161258257fe5b04905090565b6000805481908190819060ff166125d05760405162461bcd60e51b815260040180806020018281038252603a8152602001806144ee603a913960400191505060405180910390fd5b63ffffffff8816600090815260056020526040902054889060ff1661262a576040805162461bcd60e51b815260206004820152601660248201526000805160206145b7833981519152604482015290519081900360640190fd5b336000908152600d602052604090205442116126775760405162461bcd60e51b81526004018080602001828103825260458152602001806143dd6045913960600191505060405180910390fd5b866126b35760405162461bcd60e51b81526004018080602001828103825260218152602001806145d76021913960400191505060405180910390fd5b856126fb576040805162461bcd60e51b81526020600482015260136024820152720636861696e54697048617368203d3d2030783606c1b604482015290519081900360640190fd5b63ffffffff8916600090815260056020526040902054600160981b900461ffff166127575760405162461bcd60e51b81526004018080602001828103825260248152602001806144ca6024913960400191505060405180910390fd5b41331480612768575041600160981b145b6127b9576040805162461bcd60e51b815260206004820152601c60248201527f696e76616c696420626c6f636b2e636f696e626173652076616c756500000000604482015290519081900360640190fd5b63ffffffff891660009081526005602052604090208054600190910154600160881b90910461ffff166001600160401b0391821601811690891614612836576040805162461bcd60e51b815260206004820152600e60248201526d34b73b30b634b2103632b233b2b960911b604482015290519081900360640190fd5b63ffffffff8916600090815260056020526040902060020154421161288c5760405162461bcd60e51b81526004018080602001828103825260358152602001806145496035913960400191505060405180910390fd5b63ffffffff89166000908152600560205260409020600381015460049091015460020210156129215763ffffffff89166000908152600560205260409020600481015460029182015491024291909103600302101561291c5760405162461bcd60e51b815260040180806020018281038252602c815260200180614339602c913960400191505060405180910390fd5b612983565b63ffffffff8916600090815260056020526040902060048101546002909101544203600f0110156129835760405162461bcd60e51b815260040180806020018281038252602c815260200180614339602c913960400191505060405180910390fd5b63ffffffff8916600090815260056020908152604080832054815160e08e901b6001600160e01b03191681850152600160a81b90910460c01b6001600160c01b03191660248201528151808203600c018152602c9091018252805190830120808452600790925290912060060154600160801b900460ff1615612a4d576040805162461bcd60e51b815260206004820152601e60248201527f6c6f636174696f6e4861736820616c72656164792066696e616c697365640000604482015290519081900360640190fd5b63ffffffff8a16600090815260056020526040902054600160a81b90046001600160401b031615612b595760008a6001600560008e63ffffffff1663ffffffff16815260200190815260200160002060000160159054906101000a90046001600160401b031603604051602001808363ffffffff1660e01b8152600401826001600160401b031660c01b8152600801925050506040516020818303038152906040528051906020012090506007600082815260200190815260200160002060060160109054906101000a900460ff16612b575760405162461bcd60e51b81526004018080602001828103825260318152602001806144996031913960400191505060405180910390fd5b505b33600090815260066020908152604080832084845290915281205460ff1615612ce157336000908152600660209081526040808320858452909152902060040154421015612bd85760405162461bcd60e51b815260040180806020018281038252605e8152602001806142db605e913960600191505060405180910390fd5b6040805133606081901b60208084019190915260348084018d90528451808503909101815260549093018452825192810192909220600091825260068352838220868352909252919091206002015414612c70576040805162461bcd60e51b81526020600482015260146024820152730d2dcecc2d8d2c840c6d0c2d2dca8d2e090c2e6d60631b604482015290519081900360640190fd5b6003805433600090815260066020908152604080832087845290915290209091015442910111612cdc576040805162461bcd60e51b815260206004820152601260248201527172657665616c20697320746f6f206c61746560701b604482015290519081900360640190fd5b612d1c565b413314801590612cf4575041600160981b145b15612d1c575063ffffffff8a16600090815260056020526040902054600160981b900461ffff165b4133148015612d2f575041600160981b14155b156137485733600090815260066020908152604080832085845290915290205460ff16612f8b5763ffffffff8b1660009081526005602052604090206004015460028054910490811015612d8257506002545b6040518061014001604052806001151581526020018b81526020018a815260200142815260200182420181526020016000801b81526020018c6001600160401b0316815260200160006001600160401b0316815260200160001515815260200160006001600160a01b031681525060066000336001600160a01b03166001600160a01b03168152602001908152602001600020600085815260200190815260200160002060008201518160000160006101000a81548160ff0219169083151502179055506020820151816001015560408201518160020155606082015181600301556080820151816004015560a0820151816005015560c08201518160060160006101000a8154816001600160401b0302191690836001600160401b0316021790555060e08201518160060160086101000a8154816001600160401b0302191690836001600160401b031602179055506101008201518160060160106101000a81548160ff0219169083151502179055506101208201518160070160006101000a8154816001600160a01b0302191690836001600160a01b031602179055509050506000805160206144228339815191528c8c600033604051808563ffffffff168152602001846001600160401b03168152602001836003811115612f6357fe5b8152602001826001600160a01b0316815260200194505050505060405180910390a150613748565b63ffffffff8b16600090815260056020526040902054600160981b810461ffff16600160a81b9091046001600160401b0316111561335c5760008b600560008e63ffffffff1663ffffffff16815260200190815260200160002060000160139054906101000a900461ffff1661ffff16600560008f63ffffffff1663ffffffff16815260200190815260200160002060000160159054906101000a90046001600160401b031603604051602001808363ffffffff1660e01b8152600401826001600160401b031660c01b815260080192505050604051602081830303815290604052805190602001209050896007600083815260200190815260200160002060050154141561321e57600061309e6124e5565b90506001600b60006007600086815260200190815260200160002060070160009054906101000a90046001600160a01b03166001600160a01b03166001600160a01b03168152602001908152602001600020600083815260200190815260200160002060008282829054906101000a90046001600160401b03160192506101000a8154816001600160401b0302191690836001600160401b031602179055506001600c600083815260200190815260200160002060008282829054906101000a90046001600160401b03160192506101000a8154816001600160401b0302191690836001600160401b031602179055506000805160206144228339815191528d8d60026007600087815260200190815260200160002060070160009054906101000a90046001600160a01b0316604051808563ffffffff168152602001846001600160401b031681526020018360038111156131f657fe5b8152602001826001600160a01b0316815260200194505050505060405180910390a150613356565b600560008d63ffffffff1663ffffffff16815260200190815260200160002060030154600560008e63ffffffff1663ffffffff16815260200190815260200160002060000160139054906101000a900461ffff1661ffff16024201600d60006007600085815260200190815260200160002060070160009054906101000a90046001600160a01b03166001600160a01b03166001600160a01b03168152602001908152602001600020819055506000805160206144228339815191528c8c60036007600086815260200190815260200160002060070160009054906101000a90046001600160a01b0316604051808563ffffffff168152602001846001600160401b0316815260200183600381111561333357fe5b8152602001826001600160a01b0316815260200194505050505060405180910390a15b506133be565b6000805160206144228339815191528b8b600133604051808563ffffffff168152602001846001600160401b0316815260200183600381111561339b57fe5b8152602001826001600160a01b0316815260200194505050505060405180910390a15b6040518061014001604052806001151581526020018a81526020016000801b815260200160066000336001600160a01b03166001600160a01b0316815260200190815260200160002060008581526020019081526020016000206003015481526020014281526020018981526020018b6001600160401b0316815260200160006001600160401b03168152602001600115158152602001336001600160a01b03168152506007600084815260200190815260200160002060008201518160000160006101000a81548160ff0219169083151502179055506020820151816001015560408201518160020155606082015181600301556080820151816004015560a0820151816005015560c08201518160060160006101000a8154816001600160401b0302191690836001600160401b0316021790555060e08201518160060160086101000a8154816001600160401b0302191690836001600160401b031602179055506101008201518160060160106101000a81548160ff0219169083151502179055506101208201518160070160006101000a8154816001600160a01b0302191690836001600160a01b031602179055509050506001600560008d63ffffffff1663ffffffff16815260200190815260200160002060000160158282829054906101000a90046001600160401b03160192506101000a8154816001600160401b0302191690836001600160401b0316021790555089600560008d63ffffffff1663ffffffff16815260200190815260200160002060010160006101000a8154816001600160401b0302191690836001600160401b0316021790555060006002600560008e63ffffffff1663ffffffff16815260200190815260200160002060040154600560008f63ffffffff1663ffffffff1681526020019081526020016000206002015460066000336001600160a01b03166001600160a01b031681526020019081526020016000206000878152602001908152602001600020600301540301816136a757fe5b63ffffffff8e1660009081526005602052604090206003015491900491506002028111156136f75763ffffffff8c1660009081526005602052604090206003810154600202600490910155613713565b63ffffffff8c1660009081526005602052604090206004018190555b5033600090815260066020908152604080832085845282528083206003015463ffffffff8f1684526005909252909120600201555b999a6000199990990199989650505050505050565b600c602052600090815260409020546001600160401b031681565b600054610100900463ffffffff1681565b6000805460ff166137cb5760405162461bcd60e51b815260040180806020018281038252603a8152602001806144ee603a913960400191505060405180910390fd5b506001600160a01b03919091166000908152600b6020908152604080832093835292905220546001600160401b031690565b600a602052600090815260409020805460018201546002830154600384015460048501546005860154600687015460079097015460ff96871697959694959394929391926001600160401b0380841693600160401b810490911692600160801b909104909116906001600160a01b03168a565b6000805460ff16156138c0576040805162461bcd60e51b8152602060048201526014602482015273696e697469616c6973656420213d2066616c736560601b604482015290519081900360640190fd5b604051806101400160405280600115158152602001620a84946001600160401b0316815260200160006001600160401b03168152602001600161ffff168152602001600461ffff16815260200160006001600160401b03168152602001620a84946001600160401b031681526020014281526020016103848152602001601e815250600560008063ffffffff16815260200190815260200160002060008201518160000160006101000a81548160ff02191690831515021790555060208201518160000160016101000a8154816001600160401b0302191690836001600160401b0316021790555060408201518160000160096101000a8154816001600160401b0302191690836001600160401b0316021790555060608201518160000160116101000a81548161ffff021916908361ffff16021790555060808201518160000160136101000a81548161ffff021916908361ffff16021790555060a08201518160000160156101000a8154816001600160401b0302191690836001600160401b0316021790555060c08201518160010160006101000a8154816001600160401b0302191690836001600160401b0316021790555060e0820151816002015561010082015181600301556101208201518160040155905050604051806101400160405280600115158152602001621fd4de6001600160401b0316815260200160006001600160401b03168152602001600161ffff168152602001600c61ffff16815260200160006001600160401b03168152602001621fd4de6001600160401b0316815260200142815260200160968152602001601e81525060056000600163ffffffff16815260200190815260200160002060008201518160000160006101000a81548160ff02191690831515021790555060208201518160000160016101000a8154816001600160401b0302191690836001600160401b0316021790555060408201518160000160096101000a8154816001600160401b0302191690836001600160401b0316021790555060608201518160000160116101000a81548161ffff021916908361ffff16021790555060808201518160000160136101000a81548161ffff021916908361ffff16021790555060a08201518160000160156101000a8154816001600160401b0302191690836001600160401b0316021790555060c08201518160010160006101000a8154816001600160401b0302191690836001600160401b0316021790555060e0820151816002015561010082015181600301556101208201518160040155905050604051806101400160405280600115158152602001623980b46001600160401b0316815260200160006001600160401b03168152602001600261ffff168152602001602861ffff16815260200160006001600160401b03168152602001623980b46001600160401b0316815260200142815260200160788152602001601e81525060056000600263ffffffff16815260200190815260200160002060008201518160000160006101000a81548160ff02191690831515021790555060208201518160000160016101000a8154816001600160401b0302191690836001600160401b0316021790555060408201518160000160096101000a8154816001600160401b0302191690836001600160401b0316021790555060608201518160000160116101000a81548161ffff021916908361ffff16021790555060808201518160000160136101000a81548161ffff021916908361ffff16021790555060a08201518160000160156101000a8154816001600160401b0302191690836001600160401b0316021790555060c08201518160010160006101000a8154816001600160401b0302191690836001600160401b0316021790555060e08201518160020155610100820151816003015561012082015181600401559050506040518061014001604052806001151581526020016303bf79006001600160401b0316815260200160006001600160401b03168152602001601e61ffff168152602001600161ffff16815260200160006001600160401b031681526020016303bf79006001600160401b0316815260200142815260200160788152602001601e81525060056000600363ffffffff16815260200190815260200160002060008201518160000160006101000a81548160ff02191690831515021790555060208201518160000160016101000a8154816001600160401b0302191690836001600160401b0316021790555060408201518160000160096101000a8154816001600160401b0302191690836001600160401b0316021790555060608201518160000160116101000a81548161ffff021916908361ffff16021790555060808201518160000160136101000a81548161ffff021916908361ffff16021790555060a08201518160000160156101000a8154816001600160401b0302191690836001600160401b0316021790555060c08201518160010160006101000a8154816001600160401b0302191690836001600160401b0316021790555060e082015181600201556101008201518160030155610120820151816004015590505060405180610140016040528060011515815260200162f6ba806001600160401b0316815260200160006001600160401b03168152602001601a61ffff168152602001600161ffff16815260200160006001600160401b0316815260200162f6ba806001600160401b0316815260200142815260200160788152602001601e81525060056000600463ffffffff16815260200190815260200160002060008201518160000160006101000a81548160ff02191690831515021790555060208201518160000160016101000a8154816001600160401b0302191690836001600160401b0316021790555060408201518160000160096101000a8154816001600160401b0302191690836001600160401b0316021790555060608201518160000160116101000a81548161ffff021916908361ffff16021790555060808201518160000160136101000a81548161ffff021916908361ffff16021790555060a08201518160000160156101000a8154816001600160401b0302191690836001600160401b0316021790555060c08201518160010160006101000a8154816001600160401b0302191690836001600160401b0316021790555060e08201518160020155610100820151816003015561012082015181600401559050506005600060016101000a81548163ffffffff021916908363ffffffff160217905550601e6002819055506201518060038190555062093a80600460006101000a8154816001600160401b0302191690836001600160401b031602179055504260018190555060016000806101000a81548160ff0219169083151502179055506001905090565b6001548156fe626c6f636b2e74696d657374616d70203c2070726f706f7365644e6f6e5061796d656e7450726f6f66735b6c6f636174696f6e486173685d5b6c65646765725d2e7065726d697474656452657665616c54696d65626c6f636b2e74696d657374616d70203c2070726f706f73656444617461417661696c6162696c69747950726f6f66735b6d73672e73656e6465725d5b6c6f636174696f6e486173685d2e7065726d697474656452657665616c54696d656e6f7420656e6f7567682074696d6520656c61707365642073696e6365207072696f722066696e616c697479626c6f636b2e74696d657374616d70203c2070726f706f7365645061796d656e7450726f6f66735b6c6f636174696f6e486173685d2e7065726d697474656452657665616c54696d6566696e616c697365645061796d656e74735b6c6f636174696f6e486173685d2e696e646578203e3d206c65646765726d73672e73656e6465722069732063757272656e746c792062616e6e656420666f722073656e64696e6720616e20756e616363657074656420636861696e546970486173685d63f7a167c8b58acdc7fe63a4a0972032b67a76b50a2d5fd8d64061f44952316c6564676572203e3d20636861696e735b636861696e49645d2e66696e616c697365644c6564676572496e646578636861696e735b636861696e49645d2e66696e616c697365644c6564676572496e646578203d3d203070726576696f75732064617461417661696c6162696c697479506572696f64206e6f74207965742066696e616c69736564636861696e735b636861696e49645d2e6e756d436f6e6669726d6174696f6e73203e2030737461746520636f6e6e6563746f72206973206e6f7420696e697469616c697365642c2072756e20696e697469616c697365436861696e732829626c6f636b2e74696d657374616d70203c3d20696e697469616c69736554696d65626c6f636b2e74696d657374616d70203c3d20636861696e735b636861696e49645d2e66696e616c6973656454696d657374616d7066696e616c697365644c6564676572496e646578202d2067656e657369734c6564676572203c3d206c6564676572486973746f727953697a65636861696e496420646f6573206e6f742065786973740000000000000000000064617461417661696c6162696c697479506572696f6448617368203d3d20307830a26469706673582212209aec6c293dc2b1ff65eed9d71736d7da7ca3d806bfbd209f092ee91d3612309464736f6c63430007060033"),
	New:     common.FromHex("0x608060405234801561001057600080fd5b50600436106100b45760003560e01c8063cfd1fdad11610071578063cfd1fdad1461015f578063eaebf6d3146101a2578063ec7424a0146101c7578063f417c9d8146101cf578063f5f59a4a146101d7578063f64b6fda146101df576100b4565b806329be4db2146100b95780634b8a125f146100e85780635f8c940d146100f057806371c5ecb1146100f857806371e24574146101155780637ff6faa61461013b575b600080fd5b6100d6600480360360208110156100cf57600080fd5b503561024f565b60408051918252519081900360200190f35b6100d661033f565b6100d6610347565b6100d66004803603602081101561010e57600080fd5b503561034c565b6100d66004803603602081101561012b57600080fd5b50356001600160a01b0316610364565b610143610379565b604080516001600160a01b039092168252519081900360200190f35b61018e6004803603608081101561017557600080fd5b508035906020810135906040810135906060013561037f565b604080519115158252519081900360200190f35b6101c5600480360360408110156101b857600080fd5b508035906020013561041e565b005b6100d66104c2565b6100d66104c8565b6100d66104ce565b6101c5600480360360208110156101f557600080fd5b81019060208101813564010000000081111561021057600080fd5b82018360208201111561022257600080fd5b8035906020019184600183028401116401000000008311171561024457600080fd5b5090925090506104d3565b60006001821161025e57600080fd5b3360009081526010602052604090206009015460001983019081111561028357600080fd5b33600090815260106020526040812060038306600381106102a057fe5b6003908102919091016002015433600090815260106020526040812091935091600019850106600381106102d057fe5b60030201600101549050816040516020018082815260200191505060405160208183030381529060405280519060200120811461030c57600080fd5b33600090815260106020526040812060036000198601066003811061032d57fe5b60030201549290921895945050505050565b636184740081565b600381565b601281611a40811061035d57600080fd5b0154905081565b60106020526000908152604090206009015481565b61dead81565b6000605a63618473ff19420104851461039757600080fd5b336000818152601060208181526040808420600981018b905581516060810183528a81528084018a905291820188905294909352529060038706600381106103db57fe5b6003020160008201518160000155602082015181600101556040820151816002015590505060115485111561041257506001610416565b5060005b949350505050565b6001821161042b57600080fd5b605a63618473ff19420104821461044157600080fd5b601154821161044f57600080fd5b334114801561045f57504161dead145b156104be576011829055806012611a40600019850106611a40811061048057fe5b0155604080518381526020810183905281517f8ffd19aa79a62d0764e560d21b1245698310783be781d7d80b38233d4d7d288c929181900390910190a15b5050565b60115481565b611a4081565b605a81565b7f5a4fad455fbfa0bb0f22d912bbfa4ef3d0887bb6933bffaf0f1f3e9fe1a12ca142838360405180848152602001806020018281038252848482818152602001925080828437600083820152604051601f909101601f1916909201829003965090945050505050a1505056fea26469706673582212207597cdc67764c23ccb2227e30ad2bf2d358d0abe1284cc2cf2534fb74bf96e1664736f6c63430007060033"),
}

func StateConnectorUpgraded(statedb *state.StateDB) bool {
	return updateApplied(statedb, stateConnectorUpgrade)
}

func UpgradeStateConnector(statedb *state.StateDB) {
	applyUpdate(statedb, stateConnectorUpgrade)
}
