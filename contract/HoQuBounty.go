// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package contract

import (
	"math/big"
	"strings"

	ethereum "github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/event"
)

// HoQuBountyABI is the input ABI used to generate the binding from.
const HoQuBountyABI = "[{\"constant\":false,\"inputs\":[{\"name\":\"_receiver\",\"type\":\"address\"}],\"name\":\"claimOne\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"maxTokensAmount\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"minBuyableAmount\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_receiver\",\"type\":\"address\"},{\"name\":\"_equivalentEthAmount\",\"type\":\"uint256\"}],\"name\":\"sub\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"tokenRate\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[],\"name\":\"approveAll\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"i\",\"type\":\"uint32\"}],\"name\":\"getReceiver\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[],\"name\":\"unpause\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[],\"name\":\"claim\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"paused\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_tokenRate\",\"type\":\"uint256\"}],\"name\":\"setTokenRate\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"bankAddress\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"isFinished\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[],\"name\":\"pause\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[],\"name\":\"buy\",\"outputs\":[],\"payable\":true,\"stateMutability\":\"payable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"getReceiversCount\",\"outputs\":[{\"name\":\"\",\"type\":\"uint32\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[],\"name\":\"claimAll\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[],\"name\":\"finish\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_receiver\",\"type\":\"address\"},{\"name\":\"_equivalentEthAmount\",\"type\":\"uint256\"}],\"name\":\"topUp\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"\",\"type\":\"address\"}],\"name\":\"approved\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_receiver\",\"type\":\"address\"}],\"name\":\"approve\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"\",\"type\":\"address\"}],\"name\":\"tokens\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"issuedTokensAmount\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_receiver\",\"type\":\"address\"},{\"name\":\"_tokensAmount\",\"type\":\"uint256\"}],\"name\":\"addByBounty\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"transferOwnership\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_receiver\",\"type\":\"address\"},{\"name\":\"_equivalentEthAmount\",\"type\":\"uint256\"}],\"name\":\"add\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"token\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"name\":\"_tokenAddress\",\"type\":\"address\"},{\"name\":\"_bankAddress\",\"type\":\"address\"},{\"name\":\"_beneficiaryAddress\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"payable\":true,\"stateMutability\":\"payable\",\"type\":\"fallback\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"_receiver\",\"type\":\"address\"},{\"indexed\":false,\"name\":\"_tokens\",\"type\":\"uint256\"}],\"name\":\"TokenAddedByBounty\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"_buyer\",\"type\":\"address\"},{\"indexed\":false,\"name\":\"_tokens\",\"type\":\"uint256\"},{\"indexed\":false,\"name\":\"_amount\",\"type\":\"uint256\"}],\"name\":\"TokenBought\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"_receiver\",\"type\":\"address\"},{\"indexed\":false,\"name\":\"_tokens\",\"type\":\"uint256\"},{\"indexed\":false,\"name\":\"_equivalentAmount\",\"type\":\"uint256\"}],\"name\":\"TokenAdded\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"_receiver\",\"type\":\"address\"},{\"indexed\":false,\"name\":\"_tokens\",\"type\":\"uint256\"},{\"indexed\":false,\"name\":\"_equivalentAmount\",\"type\":\"uint256\"}],\"name\":\"TokenToppedUp\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"_receiver\",\"type\":\"address\"},{\"indexed\":false,\"name\":\"_tokens\",\"type\":\"uint256\"},{\"indexed\":false,\"name\":\"_equivalentAmount\",\"type\":\"uint256\"}],\"name\":\"TokenSubtracted\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"_receiver\",\"type\":\"address\"},{\"indexed\":false,\"name\":\"_tokens\",\"type\":\"uint256\"}],\"name\":\"TokenSent\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[],\"name\":\"Pause\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[],\"name\":\"Unpause\",\"type\":\"event\"}]"

// HoQuBountyBin is the compiled bytecode used for deploying new contracts.
const HoQuBountyBin = `0x60606040526000805460a060020a60ff02191681556005556009805460ff19169055341561002c57600080fd5b60405160608061124283398101604052808051919060200180519190602001805160008054600160a060020a03338116600160a060020a0319928316178355600380549882169883169890981790975560028054968816968216969096179095556001805492909616919094161790935550611770600755662386f26fc100006006556a71f78cd494d9de56a00000600455635c26b90060085561116b915081906100d790396000f3006060604052600436106101505763ffffffff60e060020a6000350416630262dc13811461015a5780630a384665146101795780632511b1821461019e57806326ffee08146101b157806331711884146101d3578063380d0c08146101e65780633c78fe07146101f95780633f4ba83a146102315780634e71d92d146102585780635c975abb1461026b57806361241c281461027e5780637822ed49146102945780637b352962146102a75780638456cb59146102ba5780638da5cb5b146102cd578063a6f2ae3a14610150578063a72a05f7146102e0578063d1058e591461030c578063d56b28891461031f578063d6f7ddf914610332578063d8b964e614610354578063daea85c514610373578063e486033914610392578063e69b414b146103b1578063eee30468146103c4578063f2fde38b146103e6578063f5d82b6b14610405578063fc0c546a14610427575b61015861043a565b005b341561016557600080fd5b610158600160a060020a03600435166105d9565b341561018457600080fd5b61018c610617565b60405190815260200160405180910390f35b34156101a957600080fd5b61018c61061d565b34156101bc57600080fd5b610158600160a060020a0360043516602435610623565b34156101de57600080fd5b61018c610733565b34156101f157600080fd5b610158610739565b341561020457600080fd5b61021563ffffffff600435166107bf565b604051600160a060020a03909116815260200160405180910390f35b341561023c57600080fd5b6102446107fd565b604051901515815260200160405180910390f35b341561026357600080fd5b610158610881565b341561027657600080fd5b6102446108a3565b341561028957600080fd5b6101586004356108b3565b341561029f57600080fd5b6102156108e0565b34156102b257600080fd5b6102446108ef565b34156102c557600080fd5b6102446108f8565b34156102d857600080fd5b610215610981565b34156102eb57600080fd5b6102f3610990565b60405163ffffffff909116815260200160405180910390f35b341561031757600080fd5b6101586109b9565b341561032a57600080fd5b610158610a74565b341561033d57600080fd5b610158600160a060020a0360043516602435610b9c565b341561035f57600080fd5b610244600160a060020a0360043516610c4e565b341561037e57600080fd5b610158600160a060020a0360043516610c63565b341561039d57600080fd5b61018c600160a060020a0360043516610cb9565b34156103bc57600080fd5b61018c610ccb565b34156103cf57600080fd5b610158600160a060020a0360043516602435610cd1565b34156103f157600080fd5b610158600160a060020a0360043516610d64565b341561041057600080fd5b610158600160a060020a0360043516602435610dbb565b341561043257600080fd5b610215610e9c565b6009546000908190819060ff161561045157600080fd5b6004546005541061046157600080fd5b60085442111561047057600080fd5b60005460a060020a900460ff161561048757600080fd5b600754349350600092506104a1908463ffffffff610eab16565b9050600454816005540111156104f2576005546004546104c69163ffffffff610ed616565b90506104dd60075482610ee890919063ffffffff16565b92506104ef348463ffffffff610ed616565b91505b600554610505908263ffffffff610eff16565b600581905560045490111561051957600080fd5b6105233382610f0e565b33600160a060020a03167f28cab0d660ed8aedd61a8c9db00b97f6a2d67e07d87795994f440b18bc5f1aa3828560405191825260208201526040908101905180910390a2600154600160a060020a031683156108fc0284604051600060405180830381858888f19350505050151561059a57600080fd5b60008211156105d457600160a060020a03331682156108fc0283604051600060405180830381858888f1935050505015156105d457600080fd5b505050565b60005433600160a060020a039081169116146105f457600080fd5b60005460a060020a900460ff161561060b57600080fd5b61061481610fff565b50565b60045481565b60065481565b6000805433600160a060020a0390811691161461063f57600080fd5b60005460a060020a900460ff161561065657600080fd5b600754610669908363ffffffff610eab16565b600160a060020a0384166000908152600a60205260409020549091508190101561069257600080fd5b600160a060020a0383166000908152600a60205260409020546106bb908263ffffffff610ed616565b600160a060020a0384166000908152600a60205260409020556005546106e7908263ffffffff610ed616565b600555600160a060020a0383167f2da8b2becb27b189a54d21689aa110c877379a94c6f59f115cedc91d66239b33828460405191825260208201526040908101905180910390a2505050565b60075481565b60008054819033600160a060020a0390811691161461075757600080fd5b60005460a060020a900460ff161561076e57600080fd5b600091505b600d5463ffffffff90811690831610156107bb575063ffffffff81166000908152600c6020526040902054600160a060020a03166107b081610c63565b600190910190610773565b5050565b6000805433600160a060020a039081169116146107db57600080fd5b5063ffffffff166000908152600c6020526040902054600160a060020a031690565b6000805433600160a060020a0390811691161461081957600080fd5b60005460a060020a900460ff16151561083157600080fd5b6000805474ff0000000000000000000000000000000000000000191690557f7805862f689e2f13df9f062ff482ad3ad112aca9e0847911ed832e158c525b3360405160405180910390a150600190565b60005460a060020a900460ff161561089857600080fd5b6108a133610fff565b565b60005460a060020a900460ff1681565b60005433600160a060020a039081169116146108ce57600080fd5b600081116108db57600080fd5b600755565b600254600160a060020a031681565b60095460ff1681565b6000805433600160a060020a0390811691161461091457600080fd5b60005460a060020a900460ff161561092b57600080fd5b6000805474ff0000000000000000000000000000000000000000191660a060020a1790557f6985a02210a168e66602d3235cb6db0e70f92b3ba4d376a33c0f3d9434bff62560405160405180910390a150600190565b600054600160a060020a031681565b6000805433600160a060020a039081169116146109ac57600080fd5b50600d5463ffffffff1690565b60008054819033600160a060020a039081169116146109d757600080fd5b60005460a060020a900460ff16156109ee57600080fd5b600091505b600d5463ffffffff90811690831610156107bb575063ffffffff81166000908152600c6020908152604080832054600160a060020a0316808452600b9092529091205460ff168015610a5b5750600160a060020a0381166000908152600a6020526040812054115b15610a6957610a6981610fff565b6001909101906109f3565b60005433600160a060020a03908116911614610a8f57600080fd5b600454600554101580610aa3575060085442115b1515610aae57600080fd5b60095460ff1615610abe57600080fd5b6009805460ff19166001179055600354600254600160a060020a039182169163a9059cbb9116826370a082313060405160e060020a63ffffffff8416028152600160a060020a039091166004820152602401602060405180830381600087803b1515610b2957600080fd5b5af11515610b3657600080fd5b5050506040518051905060405160e060020a63ffffffff8516028152600160a060020a0390921660048301526024820152604401602060405180830381600087803b1515610b8357600080fd5b5af11515610b9057600080fd5b50505060405180515050565b6000805433600160a060020a03908116911614610bb857600080fd5b60005460a060020a900460ff1615610bcf57600080fd5b600754610be2908363ffffffff610eab16565b600554909150610bf8908263ffffffff610eff16565b600555610c058382610f0e565b82600160a060020a03167ff10143e1b37ebaf3698e98976d8bddabcea045adc939c03b245da18e81ea9651828460405191825260208201526040908101905180910390a2505050565b600b6020526000908152604090205460ff1681565b60005433600160a060020a03908116911614610c7e57600080fd5b60005460a060020a900460ff1615610c9557600080fd5b600160a060020a03166000908152600b60205260409020805460ff19166001179055565b600a6020526000908152604090205481565b60055481565b60005433600160a060020a03908116911614610cec57600080fd5b60005460a060020a900460ff1615610d0357600080fd5b600554610d16908263ffffffff610eff16565b600555610d238282610f0e565b81600160a060020a03167fda3b4bd5d0318cb8ec5ccd5cf9d80332d3809ba7ecac213080f0c7b22e01947d8260405190815260200160405180910390a25050565b60005433600160a060020a03908116911614610d7f57600080fd5b600160a060020a038116156106145760008054600160a060020a03831673ffffffffffffffffffffffffffffffffffffffff1990911617905550565b6000805433600160a060020a03908116911614610dd757600080fd5b60095460ff1615610de757600080fd5b60045460055410610df757600080fd5b600854421115610e0657600080fd5b60005460a060020a900460ff1615610e1d57600080fd5b600754610e30908363ffffffff610eab16565b600554909150610e46908263ffffffff610eff16565b600555610e538382610f0e565b82600160a060020a03167fa818a22273fc309f0a3682b642c74c5b5c25c0615ff378d07767cd231e19fffc828460405191825260208201526040908101905180910390a2505050565b600354600160a060020a031681565b6000828202831580610ec75750828482811515610ec457fe5b04145b1515610ecf57fe5b9392505050565b600082821115610ee257fe5b50900390565b6000808284811515610ef657fe5b04949350505050565b600082820183811015610ecf57fe5b60005460a060020a900460ff1615610f2557600080fd5b600160a060020a0382166000908152600a60205260409020541515610fb657600d805463ffffffff9081166000908152600c60209081526040808320805473ffffffffffffffffffffffffffffffffffffffff1916600160a060020a038916908117909155855463ffffffff19811690861660010190951694909417909455918152600b90915220805460ff191690555b600160a060020a0382166000908152600a6020526040902054610fdf908263ffffffff610eff16565b600160a060020a039092166000908152600a602052604090209190915550565b6000805460a060020a900460ff161561101757600080fd5b600160a060020a0382166000908152600b602052604090205460ff16151561103e57600080fd5b600160a060020a0382166000908152600a60205260408120541161106157600080fd5b50600160a060020a038082166000908152600a602052604080822080549290556003546002549293908116926323b872dd92911690859085905160e060020a63ffffffff8616028152600160a060020a0393841660048201529190921660248201526044810191909152606401602060405180830381600087803b15156110e757600080fd5b5af115156110f457600080fd5b50505060405180515050600160a060020a0382167fd4ff88d569a7ad2bfd6b17da9dbe82e2ccc016fd0051e08ff1ad1e6fe8fc9b558260405190815260200160405180910390a250505600a165627a7a72305820ccf987cdce6c55594f05070a8c73fc4b97c68f6385c501561603c8153960a2db0029`

// DeployHoQuBounty deploys a new Ethereum contract, binding an instance of HoQuBounty to it.
func DeployHoQuBounty(auth *bind.TransactOpts, backend bind.ContractBackend, _tokenAddress common.Address, _bankAddress common.Address, _beneficiaryAddress common.Address) (common.Address, *types.Transaction, *HoQuBounty, error) {
	parsed, err := abi.JSON(strings.NewReader(HoQuBountyABI))
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	address, tx, contract, err := bind.DeployContract(auth, parsed, common.FromHex(HoQuBountyBin), backend, _tokenAddress, _bankAddress, _beneficiaryAddress)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &HoQuBounty{HoQuBountyCaller: HoQuBountyCaller{contract: contract}, HoQuBountyTransactor: HoQuBountyTransactor{contract: contract}, HoQuBountyFilterer: HoQuBountyFilterer{contract: contract}}, nil
}

// HoQuBounty is an auto generated Go binding around an Ethereum contract.
type HoQuBounty struct {
	HoQuBountyCaller     // Read-only binding to the contract
	HoQuBountyTransactor // Write-only binding to the contract
	HoQuBountyFilterer   // Log filterer for contract events
}

// HoQuBountyCaller is an auto generated read-only Go binding around an Ethereum contract.
type HoQuBountyCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// HoQuBountyTransactor is an auto generated write-only Go binding around an Ethereum contract.
type HoQuBountyTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// HoQuBountyFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type HoQuBountyFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// HoQuBountySession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type HoQuBountySession struct {
	Contract     *HoQuBounty       // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// HoQuBountyCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type HoQuBountyCallerSession struct {
	Contract *HoQuBountyCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts     // Call options to use throughout this session
}

// HoQuBountyTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type HoQuBountyTransactorSession struct {
	Contract     *HoQuBountyTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts     // Transaction auth options to use throughout this session
}

// HoQuBountyRaw is an auto generated low-level Go binding around an Ethereum contract.
type HoQuBountyRaw struct {
	Contract *HoQuBounty // Generic contract binding to access the raw methods on
}

// HoQuBountyCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type HoQuBountyCallerRaw struct {
	Contract *HoQuBountyCaller // Generic read-only contract binding to access the raw methods on
}

// HoQuBountyTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type HoQuBountyTransactorRaw struct {
	Contract *HoQuBountyTransactor // Generic write-only contract binding to access the raw methods on
}

// NewHoQuBounty creates a new instance of HoQuBounty, bound to a specific deployed contract.
func NewHoQuBounty(address common.Address, backend bind.ContractBackend) (*HoQuBounty, error) {
	contract, err := bindHoQuBounty(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &HoQuBounty{HoQuBountyCaller: HoQuBountyCaller{contract: contract}, HoQuBountyTransactor: HoQuBountyTransactor{contract: contract}, HoQuBountyFilterer: HoQuBountyFilterer{contract: contract}}, nil
}

// NewHoQuBountyCaller creates a new read-only instance of HoQuBounty, bound to a specific deployed contract.
func NewHoQuBountyCaller(address common.Address, caller bind.ContractCaller) (*HoQuBountyCaller, error) {
	contract, err := bindHoQuBounty(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &HoQuBountyCaller{contract: contract}, nil
}

// NewHoQuBountyTransactor creates a new write-only instance of HoQuBounty, bound to a specific deployed contract.
func NewHoQuBountyTransactor(address common.Address, transactor bind.ContractTransactor) (*HoQuBountyTransactor, error) {
	contract, err := bindHoQuBounty(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &HoQuBountyTransactor{contract: contract}, nil
}

// NewHoQuBountyFilterer creates a new log filterer instance of HoQuBounty, bound to a specific deployed contract.
func NewHoQuBountyFilterer(address common.Address, filterer bind.ContractFilterer) (*HoQuBountyFilterer, error) {
	contract, err := bindHoQuBounty(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &HoQuBountyFilterer{contract: contract}, nil
}

// bindHoQuBounty binds a generic wrapper to an already deployed contract.
func bindHoQuBounty(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(HoQuBountyABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_HoQuBounty *HoQuBountyRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _HoQuBounty.Contract.HoQuBountyCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_HoQuBounty *HoQuBountyRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _HoQuBounty.Contract.HoQuBountyTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_HoQuBounty *HoQuBountyRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _HoQuBounty.Contract.HoQuBountyTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_HoQuBounty *HoQuBountyCallerRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _HoQuBounty.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_HoQuBounty *HoQuBountyTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _HoQuBounty.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_HoQuBounty *HoQuBountyTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _HoQuBounty.Contract.contract.Transact(opts, method, params...)
}

// Approved is a free data retrieval call binding the contract method 0xd8b964e6.
//
// Solidity: function approved( address) constant returns(bool)
func (_HoQuBounty *HoQuBountyCaller) Approved(opts *bind.CallOpts, arg0 common.Address) (bool, error) {
	var (
		ret0 = new(bool)
	)
	out := ret0
	err := _HoQuBounty.contract.Call(opts, out, "approved", arg0)
	return *ret0, err
}

// Approved is a free data retrieval call binding the contract method 0xd8b964e6.
//
// Solidity: function approved( address) constant returns(bool)
func (_HoQuBounty *HoQuBountySession) Approved(arg0 common.Address) (bool, error) {
	return _HoQuBounty.Contract.Approved(&_HoQuBounty.CallOpts, arg0)
}

// Approved is a free data retrieval call binding the contract method 0xd8b964e6.
//
// Solidity: function approved( address) constant returns(bool)
func (_HoQuBounty *HoQuBountyCallerSession) Approved(arg0 common.Address) (bool, error) {
	return _HoQuBounty.Contract.Approved(&_HoQuBounty.CallOpts, arg0)
}

// BankAddress is a free data retrieval call binding the contract method 0x7822ed49.
//
// Solidity: function bankAddress() constant returns(address)
func (_HoQuBounty *HoQuBountyCaller) BankAddress(opts *bind.CallOpts) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _HoQuBounty.contract.Call(opts, out, "bankAddress")
	return *ret0, err
}

// BankAddress is a free data retrieval call binding the contract method 0x7822ed49.
//
// Solidity: function bankAddress() constant returns(address)
func (_HoQuBounty *HoQuBountySession) BankAddress() (common.Address, error) {
	return _HoQuBounty.Contract.BankAddress(&_HoQuBounty.CallOpts)
}

// BankAddress is a free data retrieval call binding the contract method 0x7822ed49.
//
// Solidity: function bankAddress() constant returns(address)
func (_HoQuBounty *HoQuBountyCallerSession) BankAddress() (common.Address, error) {
	return _HoQuBounty.Contract.BankAddress(&_HoQuBounty.CallOpts)
}

// GetReceiver is a free data retrieval call binding the contract method 0x3c78fe07.
//
// Solidity: function getReceiver(i uint32) constant returns(address)
func (_HoQuBounty *HoQuBountyCaller) GetReceiver(opts *bind.CallOpts, i uint32) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _HoQuBounty.contract.Call(opts, out, "getReceiver", i)
	return *ret0, err
}

// GetReceiver is a free data retrieval call binding the contract method 0x3c78fe07.
//
// Solidity: function getReceiver(i uint32) constant returns(address)
func (_HoQuBounty *HoQuBountySession) GetReceiver(i uint32) (common.Address, error) {
	return _HoQuBounty.Contract.GetReceiver(&_HoQuBounty.CallOpts, i)
}

// GetReceiver is a free data retrieval call binding the contract method 0x3c78fe07.
//
// Solidity: function getReceiver(i uint32) constant returns(address)
func (_HoQuBounty *HoQuBountyCallerSession) GetReceiver(i uint32) (common.Address, error) {
	return _HoQuBounty.Contract.GetReceiver(&_HoQuBounty.CallOpts, i)
}

// GetReceiversCount is a free data retrieval call binding the contract method 0xa72a05f7.
//
// Solidity: function getReceiversCount() constant returns(uint32)
func (_HoQuBounty *HoQuBountyCaller) GetReceiversCount(opts *bind.CallOpts) (uint32, error) {
	var (
		ret0 = new(uint32)
	)
	out := ret0
	err := _HoQuBounty.contract.Call(opts, out, "getReceiversCount")
	return *ret0, err
}

// GetReceiversCount is a free data retrieval call binding the contract method 0xa72a05f7.
//
// Solidity: function getReceiversCount() constant returns(uint32)
func (_HoQuBounty *HoQuBountySession) GetReceiversCount() (uint32, error) {
	return _HoQuBounty.Contract.GetReceiversCount(&_HoQuBounty.CallOpts)
}

// GetReceiversCount is a free data retrieval call binding the contract method 0xa72a05f7.
//
// Solidity: function getReceiversCount() constant returns(uint32)
func (_HoQuBounty *HoQuBountyCallerSession) GetReceiversCount() (uint32, error) {
	return _HoQuBounty.Contract.GetReceiversCount(&_HoQuBounty.CallOpts)
}

// IsFinished is a free data retrieval call binding the contract method 0x7b352962.
//
// Solidity: function isFinished() constant returns(bool)
func (_HoQuBounty *HoQuBountyCaller) IsFinished(opts *bind.CallOpts) (bool, error) {
	var (
		ret0 = new(bool)
	)
	out := ret0
	err := _HoQuBounty.contract.Call(opts, out, "isFinished")
	return *ret0, err
}

// IsFinished is a free data retrieval call binding the contract method 0x7b352962.
//
// Solidity: function isFinished() constant returns(bool)
func (_HoQuBounty *HoQuBountySession) IsFinished() (bool, error) {
	return _HoQuBounty.Contract.IsFinished(&_HoQuBounty.CallOpts)
}

// IsFinished is a free data retrieval call binding the contract method 0x7b352962.
//
// Solidity: function isFinished() constant returns(bool)
func (_HoQuBounty *HoQuBountyCallerSession) IsFinished() (bool, error) {
	return _HoQuBounty.Contract.IsFinished(&_HoQuBounty.CallOpts)
}

// IssuedTokensAmount is a free data retrieval call binding the contract method 0xe69b414b.
//
// Solidity: function issuedTokensAmount() constant returns(uint256)
func (_HoQuBounty *HoQuBountyCaller) IssuedTokensAmount(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _HoQuBounty.contract.Call(opts, out, "issuedTokensAmount")
	return *ret0, err
}

// IssuedTokensAmount is a free data retrieval call binding the contract method 0xe69b414b.
//
// Solidity: function issuedTokensAmount() constant returns(uint256)
func (_HoQuBounty *HoQuBountySession) IssuedTokensAmount() (*big.Int, error) {
	return _HoQuBounty.Contract.IssuedTokensAmount(&_HoQuBounty.CallOpts)
}

// IssuedTokensAmount is a free data retrieval call binding the contract method 0xe69b414b.
//
// Solidity: function issuedTokensAmount() constant returns(uint256)
func (_HoQuBounty *HoQuBountyCallerSession) IssuedTokensAmount() (*big.Int, error) {
	return _HoQuBounty.Contract.IssuedTokensAmount(&_HoQuBounty.CallOpts)
}

// MaxTokensAmount is a free data retrieval call binding the contract method 0x0a384665.
//
// Solidity: function maxTokensAmount() constant returns(uint256)
func (_HoQuBounty *HoQuBountyCaller) MaxTokensAmount(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _HoQuBounty.contract.Call(opts, out, "maxTokensAmount")
	return *ret0, err
}

// MaxTokensAmount is a free data retrieval call binding the contract method 0x0a384665.
//
// Solidity: function maxTokensAmount() constant returns(uint256)
func (_HoQuBounty *HoQuBountySession) MaxTokensAmount() (*big.Int, error) {
	return _HoQuBounty.Contract.MaxTokensAmount(&_HoQuBounty.CallOpts)
}

// MaxTokensAmount is a free data retrieval call binding the contract method 0x0a384665.
//
// Solidity: function maxTokensAmount() constant returns(uint256)
func (_HoQuBounty *HoQuBountyCallerSession) MaxTokensAmount() (*big.Int, error) {
	return _HoQuBounty.Contract.MaxTokensAmount(&_HoQuBounty.CallOpts)
}

// MinBuyableAmount is a free data retrieval call binding the contract method 0x2511b182.
//
// Solidity: function minBuyableAmount() constant returns(uint256)
func (_HoQuBounty *HoQuBountyCaller) MinBuyableAmount(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _HoQuBounty.contract.Call(opts, out, "minBuyableAmount")
	return *ret0, err
}

// MinBuyableAmount is a free data retrieval call binding the contract method 0x2511b182.
//
// Solidity: function minBuyableAmount() constant returns(uint256)
func (_HoQuBounty *HoQuBountySession) MinBuyableAmount() (*big.Int, error) {
	return _HoQuBounty.Contract.MinBuyableAmount(&_HoQuBounty.CallOpts)
}

// MinBuyableAmount is a free data retrieval call binding the contract method 0x2511b182.
//
// Solidity: function minBuyableAmount() constant returns(uint256)
func (_HoQuBounty *HoQuBountyCallerSession) MinBuyableAmount() (*big.Int, error) {
	return _HoQuBounty.Contract.MinBuyableAmount(&_HoQuBounty.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() constant returns(address)
func (_HoQuBounty *HoQuBountyCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _HoQuBounty.contract.Call(opts, out, "owner")
	return *ret0, err
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() constant returns(address)
func (_HoQuBounty *HoQuBountySession) Owner() (common.Address, error) {
	return _HoQuBounty.Contract.Owner(&_HoQuBounty.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() constant returns(address)
func (_HoQuBounty *HoQuBountyCallerSession) Owner() (common.Address, error) {
	return _HoQuBounty.Contract.Owner(&_HoQuBounty.CallOpts)
}

// Paused is a free data retrieval call binding the contract method 0x5c975abb.
//
// Solidity: function paused() constant returns(bool)
func (_HoQuBounty *HoQuBountyCaller) Paused(opts *bind.CallOpts) (bool, error) {
	var (
		ret0 = new(bool)
	)
	out := ret0
	err := _HoQuBounty.contract.Call(opts, out, "paused")
	return *ret0, err
}

// Paused is a free data retrieval call binding the contract method 0x5c975abb.
//
// Solidity: function paused() constant returns(bool)
func (_HoQuBounty *HoQuBountySession) Paused() (bool, error) {
	return _HoQuBounty.Contract.Paused(&_HoQuBounty.CallOpts)
}

// Paused is a free data retrieval call binding the contract method 0x5c975abb.
//
// Solidity: function paused() constant returns(bool)
func (_HoQuBounty *HoQuBountyCallerSession) Paused() (bool, error) {
	return _HoQuBounty.Contract.Paused(&_HoQuBounty.CallOpts)
}

// Token is a free data retrieval call binding the contract method 0xfc0c546a.
//
// Solidity: function token() constant returns(address)
func (_HoQuBounty *HoQuBountyCaller) Token(opts *bind.CallOpts) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _HoQuBounty.contract.Call(opts, out, "token")
	return *ret0, err
}

// Token is a free data retrieval call binding the contract method 0xfc0c546a.
//
// Solidity: function token() constant returns(address)
func (_HoQuBounty *HoQuBountySession) Token() (common.Address, error) {
	return _HoQuBounty.Contract.Token(&_HoQuBounty.CallOpts)
}

// Token is a free data retrieval call binding the contract method 0xfc0c546a.
//
// Solidity: function token() constant returns(address)
func (_HoQuBounty *HoQuBountyCallerSession) Token() (common.Address, error) {
	return _HoQuBounty.Contract.Token(&_HoQuBounty.CallOpts)
}

// TokenRate is a free data retrieval call binding the contract method 0x31711884.
//
// Solidity: function tokenRate() constant returns(uint256)
func (_HoQuBounty *HoQuBountyCaller) TokenRate(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _HoQuBounty.contract.Call(opts, out, "tokenRate")
	return *ret0, err
}

// TokenRate is a free data retrieval call binding the contract method 0x31711884.
//
// Solidity: function tokenRate() constant returns(uint256)
func (_HoQuBounty *HoQuBountySession) TokenRate() (*big.Int, error) {
	return _HoQuBounty.Contract.TokenRate(&_HoQuBounty.CallOpts)
}

// TokenRate is a free data retrieval call binding the contract method 0x31711884.
//
// Solidity: function tokenRate() constant returns(uint256)
func (_HoQuBounty *HoQuBountyCallerSession) TokenRate() (*big.Int, error) {
	return _HoQuBounty.Contract.TokenRate(&_HoQuBounty.CallOpts)
}

// Tokens is a free data retrieval call binding the contract method 0xe4860339.
//
// Solidity: function tokens( address) constant returns(uint256)
func (_HoQuBounty *HoQuBountyCaller) Tokens(opts *bind.CallOpts, arg0 common.Address) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _HoQuBounty.contract.Call(opts, out, "tokens", arg0)
	return *ret0, err
}

// Tokens is a free data retrieval call binding the contract method 0xe4860339.
//
// Solidity: function tokens( address) constant returns(uint256)
func (_HoQuBounty *HoQuBountySession) Tokens(arg0 common.Address) (*big.Int, error) {
	return _HoQuBounty.Contract.Tokens(&_HoQuBounty.CallOpts, arg0)
}

// Tokens is a free data retrieval call binding the contract method 0xe4860339.
//
// Solidity: function tokens( address) constant returns(uint256)
func (_HoQuBounty *HoQuBountyCallerSession) Tokens(arg0 common.Address) (*big.Int, error) {
	return _HoQuBounty.Contract.Tokens(&_HoQuBounty.CallOpts, arg0)
}

// Add is a paid mutator transaction binding the contract method 0xf5d82b6b.
//
// Solidity: function add(_receiver address, _equivalentEthAmount uint256) returns()
func (_HoQuBounty *HoQuBountyTransactor) Add(opts *bind.TransactOpts, _receiver common.Address, _equivalentEthAmount *big.Int) (*types.Transaction, error) {
	return _HoQuBounty.contract.Transact(opts, "add", _receiver, _equivalentEthAmount)
}

// Add is a paid mutator transaction binding the contract method 0xf5d82b6b.
//
// Solidity: function add(_receiver address, _equivalentEthAmount uint256) returns()
func (_HoQuBounty *HoQuBountySession) Add(_receiver common.Address, _equivalentEthAmount *big.Int) (*types.Transaction, error) {
	return _HoQuBounty.Contract.Add(&_HoQuBounty.TransactOpts, _receiver, _equivalentEthAmount)
}

// Add is a paid mutator transaction binding the contract method 0xf5d82b6b.
//
// Solidity: function add(_receiver address, _equivalentEthAmount uint256) returns()
func (_HoQuBounty *HoQuBountyTransactorSession) Add(_receiver common.Address, _equivalentEthAmount *big.Int) (*types.Transaction, error) {
	return _HoQuBounty.Contract.Add(&_HoQuBounty.TransactOpts, _receiver, _equivalentEthAmount)
}

// AddByBounty is a paid mutator transaction binding the contract method 0xeee30468.
//
// Solidity: function addByBounty(_receiver address, _tokensAmount uint256) returns()
func (_HoQuBounty *HoQuBountyTransactor) AddByBounty(opts *bind.TransactOpts, _receiver common.Address, _tokensAmount *big.Int) (*types.Transaction, error) {
	return _HoQuBounty.contract.Transact(opts, "addByBounty", _receiver, _tokensAmount)
}

// AddByBounty is a paid mutator transaction binding the contract method 0xeee30468.
//
// Solidity: function addByBounty(_receiver address, _tokensAmount uint256) returns()
func (_HoQuBounty *HoQuBountySession) AddByBounty(_receiver common.Address, _tokensAmount *big.Int) (*types.Transaction, error) {
	return _HoQuBounty.Contract.AddByBounty(&_HoQuBounty.TransactOpts, _receiver, _tokensAmount)
}

// AddByBounty is a paid mutator transaction binding the contract method 0xeee30468.
//
// Solidity: function addByBounty(_receiver address, _tokensAmount uint256) returns()
func (_HoQuBounty *HoQuBountyTransactorSession) AddByBounty(_receiver common.Address, _tokensAmount *big.Int) (*types.Transaction, error) {
	return _HoQuBounty.Contract.AddByBounty(&_HoQuBounty.TransactOpts, _receiver, _tokensAmount)
}

// Approve is a paid mutator transaction binding the contract method 0xdaea85c5.
//
// Solidity: function approve(_receiver address) returns()
func (_HoQuBounty *HoQuBountyTransactor) Approve(opts *bind.TransactOpts, _receiver common.Address) (*types.Transaction, error) {
	return _HoQuBounty.contract.Transact(opts, "approve", _receiver)
}

// Approve is a paid mutator transaction binding the contract method 0xdaea85c5.
//
// Solidity: function approve(_receiver address) returns()
func (_HoQuBounty *HoQuBountySession) Approve(_receiver common.Address) (*types.Transaction, error) {
	return _HoQuBounty.Contract.Approve(&_HoQuBounty.TransactOpts, _receiver)
}

// Approve is a paid mutator transaction binding the contract method 0xdaea85c5.
//
// Solidity: function approve(_receiver address) returns()
func (_HoQuBounty *HoQuBountyTransactorSession) Approve(_receiver common.Address) (*types.Transaction, error) {
	return _HoQuBounty.Contract.Approve(&_HoQuBounty.TransactOpts, _receiver)
}

// ApproveAll is a paid mutator transaction binding the contract method 0x380d0c08.
//
// Solidity: function approveAll() returns()
func (_HoQuBounty *HoQuBountyTransactor) ApproveAll(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _HoQuBounty.contract.Transact(opts, "approveAll")
}

// ApproveAll is a paid mutator transaction binding the contract method 0x380d0c08.
//
// Solidity: function approveAll() returns()
func (_HoQuBounty *HoQuBountySession) ApproveAll() (*types.Transaction, error) {
	return _HoQuBounty.Contract.ApproveAll(&_HoQuBounty.TransactOpts)
}

// ApproveAll is a paid mutator transaction binding the contract method 0x380d0c08.
//
// Solidity: function approveAll() returns()
func (_HoQuBounty *HoQuBountyTransactorSession) ApproveAll() (*types.Transaction, error) {
	return _HoQuBounty.Contract.ApproveAll(&_HoQuBounty.TransactOpts)
}

// Buy is a paid mutator transaction binding the contract method 0xa6f2ae3a.
//
// Solidity: function buy() returns()
func (_HoQuBounty *HoQuBountyTransactor) Buy(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _HoQuBounty.contract.Transact(opts, "buy")
}

// Buy is a paid mutator transaction binding the contract method 0xa6f2ae3a.
//
// Solidity: function buy() returns()
func (_HoQuBounty *HoQuBountySession) Buy() (*types.Transaction, error) {
	return _HoQuBounty.Contract.Buy(&_HoQuBounty.TransactOpts)
}

// Buy is a paid mutator transaction binding the contract method 0xa6f2ae3a.
//
// Solidity: function buy() returns()
func (_HoQuBounty *HoQuBountyTransactorSession) Buy() (*types.Transaction, error) {
	return _HoQuBounty.Contract.Buy(&_HoQuBounty.TransactOpts)
}

// Claim is a paid mutator transaction binding the contract method 0x4e71d92d.
//
// Solidity: function claim() returns()
func (_HoQuBounty *HoQuBountyTransactor) Claim(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _HoQuBounty.contract.Transact(opts, "claim")
}

// Claim is a paid mutator transaction binding the contract method 0x4e71d92d.
//
// Solidity: function claim() returns()
func (_HoQuBounty *HoQuBountySession) Claim() (*types.Transaction, error) {
	return _HoQuBounty.Contract.Claim(&_HoQuBounty.TransactOpts)
}

// Claim is a paid mutator transaction binding the contract method 0x4e71d92d.
//
// Solidity: function claim() returns()
func (_HoQuBounty *HoQuBountyTransactorSession) Claim() (*types.Transaction, error) {
	return _HoQuBounty.Contract.Claim(&_HoQuBounty.TransactOpts)
}

// ClaimAll is a paid mutator transaction binding the contract method 0xd1058e59.
//
// Solidity: function claimAll() returns()
func (_HoQuBounty *HoQuBountyTransactor) ClaimAll(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _HoQuBounty.contract.Transact(opts, "claimAll")
}

// ClaimAll is a paid mutator transaction binding the contract method 0xd1058e59.
//
// Solidity: function claimAll() returns()
func (_HoQuBounty *HoQuBountySession) ClaimAll() (*types.Transaction, error) {
	return _HoQuBounty.Contract.ClaimAll(&_HoQuBounty.TransactOpts)
}

// ClaimAll is a paid mutator transaction binding the contract method 0xd1058e59.
//
// Solidity: function claimAll() returns()
func (_HoQuBounty *HoQuBountyTransactorSession) ClaimAll() (*types.Transaction, error) {
	return _HoQuBounty.Contract.ClaimAll(&_HoQuBounty.TransactOpts)
}

// ClaimOne is a paid mutator transaction binding the contract method 0x0262dc13.
//
// Solidity: function claimOne(_receiver address) returns()
func (_HoQuBounty *HoQuBountyTransactor) ClaimOne(opts *bind.TransactOpts, _receiver common.Address) (*types.Transaction, error) {
	return _HoQuBounty.contract.Transact(opts, "claimOne", _receiver)
}

// ClaimOne is a paid mutator transaction binding the contract method 0x0262dc13.
//
// Solidity: function claimOne(_receiver address) returns()
func (_HoQuBounty *HoQuBountySession) ClaimOne(_receiver common.Address) (*types.Transaction, error) {
	return _HoQuBounty.Contract.ClaimOne(&_HoQuBounty.TransactOpts, _receiver)
}

// ClaimOne is a paid mutator transaction binding the contract method 0x0262dc13.
//
// Solidity: function claimOne(_receiver address) returns()
func (_HoQuBounty *HoQuBountyTransactorSession) ClaimOne(_receiver common.Address) (*types.Transaction, error) {
	return _HoQuBounty.Contract.ClaimOne(&_HoQuBounty.TransactOpts, _receiver)
}

// Finish is a paid mutator transaction binding the contract method 0xd56b2889.
//
// Solidity: function finish() returns()
func (_HoQuBounty *HoQuBountyTransactor) Finish(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _HoQuBounty.contract.Transact(opts, "finish")
}

// Finish is a paid mutator transaction binding the contract method 0xd56b2889.
//
// Solidity: function finish() returns()
func (_HoQuBounty *HoQuBountySession) Finish() (*types.Transaction, error) {
	return _HoQuBounty.Contract.Finish(&_HoQuBounty.TransactOpts)
}

// Finish is a paid mutator transaction binding the contract method 0xd56b2889.
//
// Solidity: function finish() returns()
func (_HoQuBounty *HoQuBountyTransactorSession) Finish() (*types.Transaction, error) {
	return _HoQuBounty.Contract.Finish(&_HoQuBounty.TransactOpts)
}

// Pause is a paid mutator transaction binding the contract method 0x8456cb59.
//
// Solidity: function pause() returns(bool)
func (_HoQuBounty *HoQuBountyTransactor) Pause(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _HoQuBounty.contract.Transact(opts, "pause")
}

// Pause is a paid mutator transaction binding the contract method 0x8456cb59.
//
// Solidity: function pause() returns(bool)
func (_HoQuBounty *HoQuBountySession) Pause() (*types.Transaction, error) {
	return _HoQuBounty.Contract.Pause(&_HoQuBounty.TransactOpts)
}

// Pause is a paid mutator transaction binding the contract method 0x8456cb59.
//
// Solidity: function pause() returns(bool)
func (_HoQuBounty *HoQuBountyTransactorSession) Pause() (*types.Transaction, error) {
	return _HoQuBounty.Contract.Pause(&_HoQuBounty.TransactOpts)
}

// SetTokenRate is a paid mutator transaction binding the contract method 0x61241c28.
//
// Solidity: function setTokenRate(_tokenRate uint256) returns()
func (_HoQuBounty *HoQuBountyTransactor) SetTokenRate(opts *bind.TransactOpts, _tokenRate *big.Int) (*types.Transaction, error) {
	return _HoQuBounty.contract.Transact(opts, "setTokenRate", _tokenRate)
}

// SetTokenRate is a paid mutator transaction binding the contract method 0x61241c28.
//
// Solidity: function setTokenRate(_tokenRate uint256) returns()
func (_HoQuBounty *HoQuBountySession) SetTokenRate(_tokenRate *big.Int) (*types.Transaction, error) {
	return _HoQuBounty.Contract.SetTokenRate(&_HoQuBounty.TransactOpts, _tokenRate)
}

// SetTokenRate is a paid mutator transaction binding the contract method 0x61241c28.
//
// Solidity: function setTokenRate(_tokenRate uint256) returns()
func (_HoQuBounty *HoQuBountyTransactorSession) SetTokenRate(_tokenRate *big.Int) (*types.Transaction, error) {
	return _HoQuBounty.Contract.SetTokenRate(&_HoQuBounty.TransactOpts, _tokenRate)
}

// Sub is a paid mutator transaction binding the contract method 0x26ffee08.
//
// Solidity: function sub(_receiver address, _equivalentEthAmount uint256) returns()
func (_HoQuBounty *HoQuBountyTransactor) Sub(opts *bind.TransactOpts, _receiver common.Address, _equivalentEthAmount *big.Int) (*types.Transaction, error) {
	return _HoQuBounty.contract.Transact(opts, "sub", _receiver, _equivalentEthAmount)
}

// Sub is a paid mutator transaction binding the contract method 0x26ffee08.
//
// Solidity: function sub(_receiver address, _equivalentEthAmount uint256) returns()
func (_HoQuBounty *HoQuBountySession) Sub(_receiver common.Address, _equivalentEthAmount *big.Int) (*types.Transaction, error) {
	return _HoQuBounty.Contract.Sub(&_HoQuBounty.TransactOpts, _receiver, _equivalentEthAmount)
}

// Sub is a paid mutator transaction binding the contract method 0x26ffee08.
//
// Solidity: function sub(_receiver address, _equivalentEthAmount uint256) returns()
func (_HoQuBounty *HoQuBountyTransactorSession) Sub(_receiver common.Address, _equivalentEthAmount *big.Int) (*types.Transaction, error) {
	return _HoQuBounty.Contract.Sub(&_HoQuBounty.TransactOpts, _receiver, _equivalentEthAmount)
}

// TopUp is a paid mutator transaction binding the contract method 0xd6f7ddf9.
//
// Solidity: function topUp(_receiver address, _equivalentEthAmount uint256) returns()
func (_HoQuBounty *HoQuBountyTransactor) TopUp(opts *bind.TransactOpts, _receiver common.Address, _equivalentEthAmount *big.Int) (*types.Transaction, error) {
	return _HoQuBounty.contract.Transact(opts, "topUp", _receiver, _equivalentEthAmount)
}

// TopUp is a paid mutator transaction binding the contract method 0xd6f7ddf9.
//
// Solidity: function topUp(_receiver address, _equivalentEthAmount uint256) returns()
func (_HoQuBounty *HoQuBountySession) TopUp(_receiver common.Address, _equivalentEthAmount *big.Int) (*types.Transaction, error) {
	return _HoQuBounty.Contract.TopUp(&_HoQuBounty.TransactOpts, _receiver, _equivalentEthAmount)
}

// TopUp is a paid mutator transaction binding the contract method 0xd6f7ddf9.
//
// Solidity: function topUp(_receiver address, _equivalentEthAmount uint256) returns()
func (_HoQuBounty *HoQuBountyTransactorSession) TopUp(_receiver common.Address, _equivalentEthAmount *big.Int) (*types.Transaction, error) {
	return _HoQuBounty.Contract.TopUp(&_HoQuBounty.TransactOpts, _receiver, _equivalentEthAmount)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(newOwner address) returns()
func (_HoQuBounty *HoQuBountyTransactor) TransferOwnership(opts *bind.TransactOpts, newOwner common.Address) (*types.Transaction, error) {
	return _HoQuBounty.contract.Transact(opts, "transferOwnership", newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(newOwner address) returns()
func (_HoQuBounty *HoQuBountySession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _HoQuBounty.Contract.TransferOwnership(&_HoQuBounty.TransactOpts, newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(newOwner address) returns()
func (_HoQuBounty *HoQuBountyTransactorSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _HoQuBounty.Contract.TransferOwnership(&_HoQuBounty.TransactOpts, newOwner)
}

// Unpause is a paid mutator transaction binding the contract method 0x3f4ba83a.
//
// Solidity: function unpause() returns(bool)
func (_HoQuBounty *HoQuBountyTransactor) Unpause(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _HoQuBounty.contract.Transact(opts, "unpause")
}

// Unpause is a paid mutator transaction binding the contract method 0x3f4ba83a.
//
// Solidity: function unpause() returns(bool)
func (_HoQuBounty *HoQuBountySession) Unpause() (*types.Transaction, error) {
	return _HoQuBounty.Contract.Unpause(&_HoQuBounty.TransactOpts)
}

// Unpause is a paid mutator transaction binding the contract method 0x3f4ba83a.
//
// Solidity: function unpause() returns(bool)
func (_HoQuBounty *HoQuBountyTransactorSession) Unpause() (*types.Transaction, error) {
	return _HoQuBounty.Contract.Unpause(&_HoQuBounty.TransactOpts)
}

// HoQuBountyPauseIterator is returned from FilterPause and is used to iterate over the raw logs and unpacked data for Pause events raised by the HoQuBounty contract.
type HoQuBountyPauseIterator struct {
	Event *HoQuBountyPause // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *HoQuBountyPauseIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(HoQuBountyPause)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(HoQuBountyPause)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *HoQuBountyPauseIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *HoQuBountyPauseIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// HoQuBountyPause represents a Pause event raised by the HoQuBounty contract.
type HoQuBountyPause struct {
	Raw types.Log // Blockchain specific contextual infos
}

// FilterPause is a free log retrieval operation binding the contract event 0x6985a02210a168e66602d3235cb6db0e70f92b3ba4d376a33c0f3d9434bff625.
//
// Solidity: event Pause()
func (_HoQuBounty *HoQuBountyFilterer) FilterPause(opts *bind.FilterOpts) (*HoQuBountyPauseIterator, error) {

	logs, sub, err := _HoQuBounty.contract.FilterLogs(opts, "Pause")
	if err != nil {
		return nil, err
	}
	return &HoQuBountyPauseIterator{contract: _HoQuBounty.contract, event: "Pause", logs: logs, sub: sub}, nil
}

// WatchPause is a free log subscription operation binding the contract event 0x6985a02210a168e66602d3235cb6db0e70f92b3ba4d376a33c0f3d9434bff625.
//
// Solidity: event Pause()
func (_HoQuBounty *HoQuBountyFilterer) WatchPause(opts *bind.WatchOpts, sink chan<- *HoQuBountyPause) (event.Subscription, error) {

	logs, sub, err := _HoQuBounty.contract.WatchLogs(opts, "Pause")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(HoQuBountyPause)
				if err := _HoQuBounty.contract.UnpackLog(event, "Pause", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// HoQuBountyTokenAddedIterator is returned from FilterTokenAdded and is used to iterate over the raw logs and unpacked data for TokenAdded events raised by the HoQuBounty contract.
type HoQuBountyTokenAddedIterator struct {
	Event *HoQuBountyTokenAdded // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *HoQuBountyTokenAddedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(HoQuBountyTokenAdded)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(HoQuBountyTokenAdded)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *HoQuBountyTokenAddedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *HoQuBountyTokenAddedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// HoQuBountyTokenAdded represents a TokenAdded event raised by the HoQuBounty contract.
type HoQuBountyTokenAdded struct {
	Receiver         common.Address
	Tokens           *big.Int
	EquivalentAmount *big.Int
	Raw              types.Log // Blockchain specific contextual infos
}

// FilterTokenAdded is a free log retrieval operation binding the contract event 0xa818a22273fc309f0a3682b642c74c5b5c25c0615ff378d07767cd231e19fffc.
//
// Solidity: event TokenAdded(_receiver indexed address, _tokens uint256, _equivalentAmount uint256)
func (_HoQuBounty *HoQuBountyFilterer) FilterTokenAdded(opts *bind.FilterOpts, _receiver []common.Address) (*HoQuBountyTokenAddedIterator, error) {

	var _receiverRule []interface{}
	for _, _receiverItem := range _receiver {
		_receiverRule = append(_receiverRule, _receiverItem)
	}

	logs, sub, err := _HoQuBounty.contract.FilterLogs(opts, "TokenAdded", _receiverRule)
	if err != nil {
		return nil, err
	}
	return &HoQuBountyTokenAddedIterator{contract: _HoQuBounty.contract, event: "TokenAdded", logs: logs, sub: sub}, nil
}

// WatchTokenAdded is a free log subscription operation binding the contract event 0xa818a22273fc309f0a3682b642c74c5b5c25c0615ff378d07767cd231e19fffc.
//
// Solidity: event TokenAdded(_receiver indexed address, _tokens uint256, _equivalentAmount uint256)
func (_HoQuBounty *HoQuBountyFilterer) WatchTokenAdded(opts *bind.WatchOpts, sink chan<- *HoQuBountyTokenAdded, _receiver []common.Address) (event.Subscription, error) {

	var _receiverRule []interface{}
	for _, _receiverItem := range _receiver {
		_receiverRule = append(_receiverRule, _receiverItem)
	}

	logs, sub, err := _HoQuBounty.contract.WatchLogs(opts, "TokenAdded", _receiverRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(HoQuBountyTokenAdded)
				if err := _HoQuBounty.contract.UnpackLog(event, "TokenAdded", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// HoQuBountyTokenAddedByBountyIterator is returned from FilterTokenAddedByBounty and is used to iterate over the raw logs and unpacked data for TokenAddedByBounty events raised by the HoQuBounty contract.
type HoQuBountyTokenAddedByBountyIterator struct {
	Event *HoQuBountyTokenAddedByBounty // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *HoQuBountyTokenAddedByBountyIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(HoQuBountyTokenAddedByBounty)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(HoQuBountyTokenAddedByBounty)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *HoQuBountyTokenAddedByBountyIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *HoQuBountyTokenAddedByBountyIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// HoQuBountyTokenAddedByBounty represents a TokenAddedByBounty event raised by the HoQuBounty contract.
type HoQuBountyTokenAddedByBounty struct {
	Receiver common.Address
	Tokens   *big.Int
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterTokenAddedByBounty is a free log retrieval operation binding the contract event 0xda3b4bd5d0318cb8ec5ccd5cf9d80332d3809ba7ecac213080f0c7b22e01947d.
//
// Solidity: event TokenAddedByBounty(_receiver indexed address, _tokens uint256)
func (_HoQuBounty *HoQuBountyFilterer) FilterTokenAddedByBounty(opts *bind.FilterOpts, _receiver []common.Address) (*HoQuBountyTokenAddedByBountyIterator, error) {

	var _receiverRule []interface{}
	for _, _receiverItem := range _receiver {
		_receiverRule = append(_receiverRule, _receiverItem)
	}

	logs, sub, err := _HoQuBounty.contract.FilterLogs(opts, "TokenAddedByBounty", _receiverRule)
	if err != nil {
		return nil, err
	}
	return &HoQuBountyTokenAddedByBountyIterator{contract: _HoQuBounty.contract, event: "TokenAddedByBounty", logs: logs, sub: sub}, nil
}

// WatchTokenAddedByBounty is a free log subscription operation binding the contract event 0xda3b4bd5d0318cb8ec5ccd5cf9d80332d3809ba7ecac213080f0c7b22e01947d.
//
// Solidity: event TokenAddedByBounty(_receiver indexed address, _tokens uint256)
func (_HoQuBounty *HoQuBountyFilterer) WatchTokenAddedByBounty(opts *bind.WatchOpts, sink chan<- *HoQuBountyTokenAddedByBounty, _receiver []common.Address) (event.Subscription, error) {

	var _receiverRule []interface{}
	for _, _receiverItem := range _receiver {
		_receiverRule = append(_receiverRule, _receiverItem)
	}

	logs, sub, err := _HoQuBounty.contract.WatchLogs(opts, "TokenAddedByBounty", _receiverRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(HoQuBountyTokenAddedByBounty)
				if err := _HoQuBounty.contract.UnpackLog(event, "TokenAddedByBounty", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// HoQuBountyTokenBoughtIterator is returned from FilterTokenBought and is used to iterate over the raw logs and unpacked data for TokenBought events raised by the HoQuBounty contract.
type HoQuBountyTokenBoughtIterator struct {
	Event *HoQuBountyTokenBought // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *HoQuBountyTokenBoughtIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(HoQuBountyTokenBought)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(HoQuBountyTokenBought)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *HoQuBountyTokenBoughtIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *HoQuBountyTokenBoughtIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// HoQuBountyTokenBought represents a TokenBought event raised by the HoQuBounty contract.
type HoQuBountyTokenBought struct {
	Buyer  common.Address
	Tokens *big.Int
	Amount *big.Int
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterTokenBought is a free log retrieval operation binding the contract event 0x28cab0d660ed8aedd61a8c9db00b97f6a2d67e07d87795994f440b18bc5f1aa3.
//
// Solidity: event TokenBought(_buyer indexed address, _tokens uint256, _amount uint256)
func (_HoQuBounty *HoQuBountyFilterer) FilterTokenBought(opts *bind.FilterOpts, _buyer []common.Address) (*HoQuBountyTokenBoughtIterator, error) {

	var _buyerRule []interface{}
	for _, _buyerItem := range _buyer {
		_buyerRule = append(_buyerRule, _buyerItem)
	}

	logs, sub, err := _HoQuBounty.contract.FilterLogs(opts, "TokenBought", _buyerRule)
	if err != nil {
		return nil, err
	}
	return &HoQuBountyTokenBoughtIterator{contract: _HoQuBounty.contract, event: "TokenBought", logs: logs, sub: sub}, nil
}

// WatchTokenBought is a free log subscription operation binding the contract event 0x28cab0d660ed8aedd61a8c9db00b97f6a2d67e07d87795994f440b18bc5f1aa3.
//
// Solidity: event TokenBought(_buyer indexed address, _tokens uint256, _amount uint256)
func (_HoQuBounty *HoQuBountyFilterer) WatchTokenBought(opts *bind.WatchOpts, sink chan<- *HoQuBountyTokenBought, _buyer []common.Address) (event.Subscription, error) {

	var _buyerRule []interface{}
	for _, _buyerItem := range _buyer {
		_buyerRule = append(_buyerRule, _buyerItem)
	}

	logs, sub, err := _HoQuBounty.contract.WatchLogs(opts, "TokenBought", _buyerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(HoQuBountyTokenBought)
				if err := _HoQuBounty.contract.UnpackLog(event, "TokenBought", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// HoQuBountyTokenSentIterator is returned from FilterTokenSent and is used to iterate over the raw logs and unpacked data for TokenSent events raised by the HoQuBounty contract.
type HoQuBountyTokenSentIterator struct {
	Event *HoQuBountyTokenSent // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *HoQuBountyTokenSentIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(HoQuBountyTokenSent)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(HoQuBountyTokenSent)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *HoQuBountyTokenSentIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *HoQuBountyTokenSentIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// HoQuBountyTokenSent represents a TokenSent event raised by the HoQuBounty contract.
type HoQuBountyTokenSent struct {
	Receiver common.Address
	Tokens   *big.Int
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterTokenSent is a free log retrieval operation binding the contract event 0xd4ff88d569a7ad2bfd6b17da9dbe82e2ccc016fd0051e08ff1ad1e6fe8fc9b55.
//
// Solidity: event TokenSent(_receiver indexed address, _tokens uint256)
func (_HoQuBounty *HoQuBountyFilterer) FilterTokenSent(opts *bind.FilterOpts, _receiver []common.Address) (*HoQuBountyTokenSentIterator, error) {

	var _receiverRule []interface{}
	for _, _receiverItem := range _receiver {
		_receiverRule = append(_receiverRule, _receiverItem)
	}

	logs, sub, err := _HoQuBounty.contract.FilterLogs(opts, "TokenSent", _receiverRule)
	if err != nil {
		return nil, err
	}
	return &HoQuBountyTokenSentIterator{contract: _HoQuBounty.contract, event: "TokenSent", logs: logs, sub: sub}, nil
}

// WatchTokenSent is a free log subscription operation binding the contract event 0xd4ff88d569a7ad2bfd6b17da9dbe82e2ccc016fd0051e08ff1ad1e6fe8fc9b55.
//
// Solidity: event TokenSent(_receiver indexed address, _tokens uint256)
func (_HoQuBounty *HoQuBountyFilterer) WatchTokenSent(opts *bind.WatchOpts, sink chan<- *HoQuBountyTokenSent, _receiver []common.Address) (event.Subscription, error) {

	var _receiverRule []interface{}
	for _, _receiverItem := range _receiver {
		_receiverRule = append(_receiverRule, _receiverItem)
	}

	logs, sub, err := _HoQuBounty.contract.WatchLogs(opts, "TokenSent", _receiverRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(HoQuBountyTokenSent)
				if err := _HoQuBounty.contract.UnpackLog(event, "TokenSent", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// HoQuBountyTokenSubtractedIterator is returned from FilterTokenSubtracted and is used to iterate over the raw logs and unpacked data for TokenSubtracted events raised by the HoQuBounty contract.
type HoQuBountyTokenSubtractedIterator struct {
	Event *HoQuBountyTokenSubtracted // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *HoQuBountyTokenSubtractedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(HoQuBountyTokenSubtracted)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(HoQuBountyTokenSubtracted)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *HoQuBountyTokenSubtractedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *HoQuBountyTokenSubtractedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// HoQuBountyTokenSubtracted represents a TokenSubtracted event raised by the HoQuBounty contract.
type HoQuBountyTokenSubtracted struct {
	Receiver         common.Address
	Tokens           *big.Int
	EquivalentAmount *big.Int
	Raw              types.Log // Blockchain specific contextual infos
}

// FilterTokenSubtracted is a free log retrieval operation binding the contract event 0x2da8b2becb27b189a54d21689aa110c877379a94c6f59f115cedc91d66239b33.
//
// Solidity: event TokenSubtracted(_receiver indexed address, _tokens uint256, _equivalentAmount uint256)
func (_HoQuBounty *HoQuBountyFilterer) FilterTokenSubtracted(opts *bind.FilterOpts, _receiver []common.Address) (*HoQuBountyTokenSubtractedIterator, error) {

	var _receiverRule []interface{}
	for _, _receiverItem := range _receiver {
		_receiverRule = append(_receiverRule, _receiverItem)
	}

	logs, sub, err := _HoQuBounty.contract.FilterLogs(opts, "TokenSubtracted", _receiverRule)
	if err != nil {
		return nil, err
	}
	return &HoQuBountyTokenSubtractedIterator{contract: _HoQuBounty.contract, event: "TokenSubtracted", logs: logs, sub: sub}, nil
}

// WatchTokenSubtracted is a free log subscription operation binding the contract event 0x2da8b2becb27b189a54d21689aa110c877379a94c6f59f115cedc91d66239b33.
//
// Solidity: event TokenSubtracted(_receiver indexed address, _tokens uint256, _equivalentAmount uint256)
func (_HoQuBounty *HoQuBountyFilterer) WatchTokenSubtracted(opts *bind.WatchOpts, sink chan<- *HoQuBountyTokenSubtracted, _receiver []common.Address) (event.Subscription, error) {

	var _receiverRule []interface{}
	for _, _receiverItem := range _receiver {
		_receiverRule = append(_receiverRule, _receiverItem)
	}

	logs, sub, err := _HoQuBounty.contract.WatchLogs(opts, "TokenSubtracted", _receiverRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(HoQuBountyTokenSubtracted)
				if err := _HoQuBounty.contract.UnpackLog(event, "TokenSubtracted", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// HoQuBountyTokenToppedUpIterator is returned from FilterTokenToppedUp and is used to iterate over the raw logs and unpacked data for TokenToppedUp events raised by the HoQuBounty contract.
type HoQuBountyTokenToppedUpIterator struct {
	Event *HoQuBountyTokenToppedUp // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *HoQuBountyTokenToppedUpIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(HoQuBountyTokenToppedUp)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(HoQuBountyTokenToppedUp)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *HoQuBountyTokenToppedUpIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *HoQuBountyTokenToppedUpIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// HoQuBountyTokenToppedUp represents a TokenToppedUp event raised by the HoQuBounty contract.
type HoQuBountyTokenToppedUp struct {
	Receiver         common.Address
	Tokens           *big.Int
	EquivalentAmount *big.Int
	Raw              types.Log // Blockchain specific contextual infos
}

// FilterTokenToppedUp is a free log retrieval operation binding the contract event 0xf10143e1b37ebaf3698e98976d8bddabcea045adc939c03b245da18e81ea9651.
//
// Solidity: event TokenToppedUp(_receiver indexed address, _tokens uint256, _equivalentAmount uint256)
func (_HoQuBounty *HoQuBountyFilterer) FilterTokenToppedUp(opts *bind.FilterOpts, _receiver []common.Address) (*HoQuBountyTokenToppedUpIterator, error) {

	var _receiverRule []interface{}
	for _, _receiverItem := range _receiver {
		_receiverRule = append(_receiverRule, _receiverItem)
	}

	logs, sub, err := _HoQuBounty.contract.FilterLogs(opts, "TokenToppedUp", _receiverRule)
	if err != nil {
		return nil, err
	}
	return &HoQuBountyTokenToppedUpIterator{contract: _HoQuBounty.contract, event: "TokenToppedUp", logs: logs, sub: sub}, nil
}

// WatchTokenToppedUp is a free log subscription operation binding the contract event 0xf10143e1b37ebaf3698e98976d8bddabcea045adc939c03b245da18e81ea9651.
//
// Solidity: event TokenToppedUp(_receiver indexed address, _tokens uint256, _equivalentAmount uint256)
func (_HoQuBounty *HoQuBountyFilterer) WatchTokenToppedUp(opts *bind.WatchOpts, sink chan<- *HoQuBountyTokenToppedUp, _receiver []common.Address) (event.Subscription, error) {

	var _receiverRule []interface{}
	for _, _receiverItem := range _receiver {
		_receiverRule = append(_receiverRule, _receiverItem)
	}

	logs, sub, err := _HoQuBounty.contract.WatchLogs(opts, "TokenToppedUp", _receiverRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(HoQuBountyTokenToppedUp)
				if err := _HoQuBounty.contract.UnpackLog(event, "TokenToppedUp", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// HoQuBountyUnpauseIterator is returned from FilterUnpause and is used to iterate over the raw logs and unpacked data for Unpause events raised by the HoQuBounty contract.
type HoQuBountyUnpauseIterator struct {
	Event *HoQuBountyUnpause // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *HoQuBountyUnpauseIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(HoQuBountyUnpause)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(HoQuBountyUnpause)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *HoQuBountyUnpauseIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *HoQuBountyUnpauseIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// HoQuBountyUnpause represents a Unpause event raised by the HoQuBounty contract.
type HoQuBountyUnpause struct {
	Raw types.Log // Blockchain specific contextual infos
}

// FilterUnpause is a free log retrieval operation binding the contract event 0x7805862f689e2f13df9f062ff482ad3ad112aca9e0847911ed832e158c525b33.
//
// Solidity: event Unpause()
func (_HoQuBounty *HoQuBountyFilterer) FilterUnpause(opts *bind.FilterOpts) (*HoQuBountyUnpauseIterator, error) {

	logs, sub, err := _HoQuBounty.contract.FilterLogs(opts, "Unpause")
	if err != nil {
		return nil, err
	}
	return &HoQuBountyUnpauseIterator{contract: _HoQuBounty.contract, event: "Unpause", logs: logs, sub: sub}, nil
}

// WatchUnpause is a free log subscription operation binding the contract event 0x7805862f689e2f13df9f062ff482ad3ad112aca9e0847911ed832e158c525b33.
//
// Solidity: event Unpause()
func (_HoQuBounty *HoQuBountyFilterer) WatchUnpause(opts *bind.WatchOpts, sink chan<- *HoQuBountyUnpause) (event.Subscription, error) {

	logs, sub, err := _HoQuBounty.contract.WatchLogs(opts, "Unpause")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(HoQuBountyUnpause)
				if err := _HoQuBounty.contract.UnpackLog(event, "Unpause", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}
