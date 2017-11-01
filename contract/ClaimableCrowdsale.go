// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package contract

import (
	"math/big"
	"strings"

	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
)

// BasicTokenABI is the input ABI used to generate the binding from.
const BasicTokenABI = "[{\"constant\":true,\"inputs\":[],\"name\":\"totalSupply\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"_owner\",\"type\":\"address\"}],\"name\":\"balanceOf\",\"outputs\":[{\"name\":\"balance\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_to\",\"type\":\"address\"},{\"name\":\"_value\",\"type\":\"uint256\"}],\"name\":\"transfer\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"from\",\"type\":\"address\"},{\"indexed\":true,\"name\":\"to\",\"type\":\"address\"},{\"indexed\":false,\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"Transfer\",\"type\":\"event\"}]"

// BasicTokenBin is the compiled bytecode used for deploying new contracts.
const BasicTokenBin = `0x6060604052341561000f57600080fd5b61023c8061001e6000396000f300606060405263ffffffff7c010000000000000000000000000000000000000000000000000000000060003504166318160ddd811461005257806370a0823114610077578063a9059cbb1461009657600080fd5b341561005d57600080fd5b6100656100cc565b60405190815260200160405180910390f35b341561008257600080fd5b610065600160a060020a03600435166100d2565b34156100a157600080fd5b6100b8600160a060020a03600435166024356100ed565b604051901515815260200160405180910390f35b60005481565b600160a060020a031660009081526001602052604090205490565b6000600160a060020a038316151561010457600080fd5b600160a060020a03331660009081526001602052604090205482111561012957600080fd5b600160a060020a033316600090815260016020526040902054610152908363ffffffff6101e816565b600160a060020a033381166000908152600160205260408082209390935590851681522054610187908363ffffffff6101fa16565b600160a060020a0380851660008181526001602052604090819020939093559133909116907fddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef9085905190815260200160405180910390a350600192915050565b6000828211156101f457fe5b50900390565b60008282018381101561020957fe5b93925050505600a165627a7a7230582016e80afd358a5a1c5943ecf93ba1b86cd9c2f1959338d9bcc04ab31b18a508f10029`

// DeployBasicToken deploys a new Ethereum contract, binding an instance of BasicToken to it.
func DeployBasicToken(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *BasicToken, error) {
	parsed, err := abi.JSON(strings.NewReader(BasicTokenABI))
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	address, tx, contract, err := bind.DeployContract(auth, parsed, common.FromHex(BasicTokenBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &BasicToken{BasicTokenCaller: BasicTokenCaller{contract: contract}, BasicTokenTransactor: BasicTokenTransactor{contract: contract}}, nil
}

// BasicToken is an auto generated Go binding around an Ethereum contract.
type BasicToken struct {
	BasicTokenCaller     // Read-only binding to the contract
	BasicTokenTransactor // Write-only binding to the contract
}

// BasicTokenCaller is an auto generated read-only Go binding around an Ethereum contract.
type BasicTokenCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// BasicTokenTransactor is an auto generated write-only Go binding around an Ethereum contract.
type BasicTokenTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// BasicTokenSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type BasicTokenSession struct {
	Contract     *BasicToken       // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// BasicTokenCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type BasicTokenCallerSession struct {
	Contract *BasicTokenCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts     // Call options to use throughout this session
}

// BasicTokenTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type BasicTokenTransactorSession struct {
	Contract     *BasicTokenTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts     // Transaction auth options to use throughout this session
}

// BasicTokenRaw is an auto generated low-level Go binding around an Ethereum contract.
type BasicTokenRaw struct {
	Contract *BasicToken // Generic contract binding to access the raw methods on
}

// BasicTokenCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type BasicTokenCallerRaw struct {
	Contract *BasicTokenCaller // Generic read-only contract binding to access the raw methods on
}

// BasicTokenTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type BasicTokenTransactorRaw struct {
	Contract *BasicTokenTransactor // Generic write-only contract binding to access the raw methods on
}

// NewBasicToken creates a new instance of BasicToken, bound to a specific deployed contract.
func NewBasicToken(address common.Address, backend bind.ContractBackend) (*BasicToken, error) {
	contract, err := bindBasicToken(address, backend, backend)
	if err != nil {
		return nil, err
	}
	return &BasicToken{BasicTokenCaller: BasicTokenCaller{contract: contract}, BasicTokenTransactor: BasicTokenTransactor{contract: contract}}, nil
}

// NewBasicTokenCaller creates a new read-only instance of BasicToken, bound to a specific deployed contract.
func NewBasicTokenCaller(address common.Address, caller bind.ContractCaller) (*BasicTokenCaller, error) {
	contract, err := bindBasicToken(address, caller, nil)
	if err != nil {
		return nil, err
	}
	return &BasicTokenCaller{contract: contract}, nil
}

// NewBasicTokenTransactor creates a new write-only instance of BasicToken, bound to a specific deployed contract.
func NewBasicTokenTransactor(address common.Address, transactor bind.ContractTransactor) (*BasicTokenTransactor, error) {
	contract, err := bindBasicToken(address, nil, transactor)
	if err != nil {
		return nil, err
	}
	return &BasicTokenTransactor{contract: contract}, nil
}

// bindBasicToken binds a generic wrapper to an already deployed contract.
func bindBasicToken(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(BasicTokenABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_BasicToken *BasicTokenRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _BasicToken.Contract.BasicTokenCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_BasicToken *BasicTokenRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _BasicToken.Contract.BasicTokenTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_BasicToken *BasicTokenRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _BasicToken.Contract.BasicTokenTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_BasicToken *BasicTokenCallerRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _BasicToken.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_BasicToken *BasicTokenTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _BasicToken.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_BasicToken *BasicTokenTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _BasicToken.Contract.contract.Transact(opts, method, params...)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(_owner address) constant returns(balance uint256)
func (_BasicToken *BasicTokenCaller) BalanceOf(opts *bind.CallOpts, _owner common.Address) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _BasicToken.contract.Call(opts, out, "balanceOf", _owner)
	return *ret0, err
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(_owner address) constant returns(balance uint256)
func (_BasicToken *BasicTokenSession) BalanceOf(_owner common.Address) (*big.Int, error) {
	return _BasicToken.Contract.BalanceOf(&_BasicToken.CallOpts, _owner)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(_owner address) constant returns(balance uint256)
func (_BasicToken *BasicTokenCallerSession) BalanceOf(_owner common.Address) (*big.Int, error) {
	return _BasicToken.Contract.BalanceOf(&_BasicToken.CallOpts, _owner)
}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() constant returns(uint256)
func (_BasicToken *BasicTokenCaller) TotalSupply(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _BasicToken.contract.Call(opts, out, "totalSupply")
	return *ret0, err
}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() constant returns(uint256)
func (_BasicToken *BasicTokenSession) TotalSupply() (*big.Int, error) {
	return _BasicToken.Contract.TotalSupply(&_BasicToken.CallOpts)
}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() constant returns(uint256)
func (_BasicToken *BasicTokenCallerSession) TotalSupply() (*big.Int, error) {
	return _BasicToken.Contract.TotalSupply(&_BasicToken.CallOpts)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(_to address, _value uint256) returns(bool)
func (_BasicToken *BasicTokenTransactor) Transfer(opts *bind.TransactOpts, _to common.Address, _value *big.Int) (*types.Transaction, error) {
	return _BasicToken.contract.Transact(opts, "transfer", _to, _value)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(_to address, _value uint256) returns(bool)
func (_BasicToken *BasicTokenSession) Transfer(_to common.Address, _value *big.Int) (*types.Transaction, error) {
	return _BasicToken.Contract.Transfer(&_BasicToken.TransactOpts, _to, _value)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(_to address, _value uint256) returns(bool)
func (_BasicToken *BasicTokenTransactorSession) Transfer(_to common.Address, _value *big.Int) (*types.Transaction, error) {
	return _BasicToken.Contract.Transfer(&_BasicToken.TransactOpts, _to, _value)
}

// ClaimableCrowdsaleABI is the input ABI used to generate the binding from.
const ClaimableCrowdsaleABI = "[{\"constant\":false,\"inputs\":[{\"name\":\"_receiver\",\"type\":\"address\"}],\"name\":\"claimOne\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"maxTokensAmount\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"minBuyableAmount\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[],\"name\":\"pauseToken\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"tokenRate\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"i\",\"type\":\"uint32\"}],\"name\":\"getReceiver\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[],\"name\":\"unpause\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[],\"name\":\"claim\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[],\"name\":\"unpauseToken\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"paused\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_tokenRate\",\"type\":\"uint256\"}],\"name\":\"setTokenRate\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"bankAddress\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"isFinished\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[],\"name\":\"pause\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[],\"name\":\"buy\",\"outputs\":[],\"payable\":true,\"stateMutability\":\"payable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"getReceiversCount\",\"outputs\":[{\"name\":\"\",\"type\":\"uint32\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[],\"name\":\"claimAll\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_minBuyableAmount\",\"type\":\"uint256\"}],\"name\":\"setMinBuyableAmount\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[],\"name\":\"finish\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"\",\"type\":\"address\"}],\"name\":\"approved\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_receiver\",\"type\":\"address\"}],\"name\":\"approve\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"\",\"type\":\"address\"}],\"name\":\"tokens\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"issuedTokensAmount\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"transferOwnership\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_receiver\",\"type\":\"address\"},{\"name\":\"_equivalentEthAmount\",\"type\":\"uint256\"}],\"name\":\"add\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"token\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"name\":\"_tokenAddress\",\"type\":\"address\"},{\"name\":\"_bankAddress\",\"type\":\"address\"},{\"name\":\"_beneficiaryAddress\",\"type\":\"address\"},{\"name\":\"_tokenRate\",\"type\":\"uint256\"},{\"name\":\"_minBuyableAmount\",\"type\":\"uint256\"},{\"name\":\"_maxTokensAmount\",\"type\":\"uint256\"},{\"name\":\"_endDate\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"payable\":true,\"stateMutability\":\"payable\",\"type\":\"fallback\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"_buyer\",\"type\":\"address\"},{\"indexed\":false,\"name\":\"_tokens\",\"type\":\"uint256\"},{\"indexed\":false,\"name\":\"_amount\",\"type\":\"uint256\"}],\"name\":\"TokenBought\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"_receiver\",\"type\":\"address\"},{\"indexed\":false,\"name\":\"_tokens\",\"type\":\"uint256\"},{\"indexed\":false,\"name\":\"_equivalentAmount\",\"type\":\"uint256\"}],\"name\":\"TokenAdded\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"_receiver\",\"type\":\"address\"},{\"indexed\":false,\"name\":\"_tokens\",\"type\":\"uint256\"}],\"name\":\"TokenSent\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[],\"name\":\"Pause\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[],\"name\":\"Unpause\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"previousOwner\",\"type\":\"address\"},{\"indexed\":true,\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"OwnershipTransferred\",\"type\":\"event\"}]"

// ClaimableCrowdsaleBin is the compiled bytecode used for deploying new contracts.
const ClaimableCrowdsaleBin = `0x60606040526000805460a060020a60ff02191681556005556009805460ff19169055341561002c57600080fd5b60405160e0806111458339810160405280805191906020018051919060200180519190602001805191906020018051919060200180519190602001805160008054600160a060020a03338116600160a060020a031992831617909255600380549b83169b82169b909b17909a5560028054998216998b1699909917909855600180549790981696909816959095179095556007929092556006556004555050600855611068806100dd6000396000f300606060405236156101435763ffffffff60e060020a6000350416630262dc13811461014d5780630a3846651461016c5780632511b182146101915780632c349627146101a457806331711884146101cb5780633c78fe07146101de5780633f4ba83a146102165780634e71d92d1461022957806350669a031461023c5780635c975abb1461024f57806361241c28146102625780637822ed49146102785780637b3529621461028b5780638456cb591461029e5780638da5cb5b146102b1578063a6f2ae3a14610143578063a72a05f7146102c4578063d1058e59146102f0578063d561ef2b14610303578063d56b288914610319578063d8b964e61461032c578063daea85c51461034b578063e48603391461036a578063e69b414b14610389578063f2fde38b1461039c578063f5d82b6b146103bb578063fc0c546a146103dd575b61014b6103f0565b005b341561015857600080fd5b61014b600160a060020a036004351661059e565b341561017757600080fd5b61017f6105dc565b60405190815260200160405180910390f35b341561019c57600080fd5b61017f6105e2565b34156101af57600080fd5b6101b76105e8565b604051901515815260200160405180910390f35b34156101d657600080fd5b61017f6106cb565b34156101e957600080fd5b6101fa63ffffffff600435166106d1565b604051600160a060020a03909116815260200160405180910390f35b341561022157600080fd5b61014b61070f565b341561023457600080fd5b61014b61078e565b341561024757600080fd5b6101b76107b0565b341561025a57600080fd5b6101b7610879565b341561026d57600080fd5b61014b600435610889565b341561028357600080fd5b6101fa6108b6565b341561029657600080fd5b6101b76108c5565b34156102a957600080fd5b61014b6108ce565b34156102bc57600080fd5b6101fa610952565b34156102cf57600080fd5b6102d7610961565b60405163ffffffff909116815260200160405180910390f35b34156102fb57600080fd5b61014b61098a565b341561030e57600080fd5b61014b600435610a49565b341561032457600080fd5b61014b610a8f565b341561033757600080fd5b6101b7600160a060020a0360043516610bd1565b341561035657600080fd5b61014b600160a060020a0360043516610be6565b341561037557600080fd5b61017f600160a060020a0360043516610c3c565b341561039457600080fd5b61017f610c4e565b34156103a757600080fd5b61014b600160a060020a0360043516610c54565b34156103c657600080fd5b61014b600160a060020a0360043516602435610cef565b34156103e857600080fd5b6101fa610da1565b6009546000908190819060ff161561040757600080fd5b6004546005541061041757600080fd5b60085442111561042657600080fd5b60005460a060020a900460ff161561043d57600080fd5b60065434101561044c57600080fd5b60075434935060009250610466908463ffffffff610db016565b9050600454816005540111156104b75760055460045461048b9163ffffffff610ddb16565b90506104a260075482610ded90919063ffffffff16565b92506104b4348463ffffffff610ddb16565b91505b6005546104ca908263ffffffff610e0416565b60058190556004549011156104de57600080fd5b6104e83382610e13565b33600160a060020a03167f28cab0d660ed8aedd61a8c9db00b97f6a2d67e07d87795994f440b18bc5f1aa3828560405191825260208201526040908101905180910390a2600154600160a060020a031683156108fc0284604051600060405180830381858888f19350505050151561055f57600080fd5b600082111561059957600160a060020a03331682156108fc0283604051600060405180830381858888f19350505050151561059957600080fd5b505050565b60005433600160a060020a039081169116146105b957600080fd5b60005460a060020a900460ff16156105d057600080fd5b6105d981610f04565b50565b60045481565b60065481565b6000805433600160a060020a0390811691161461060457600080fd5b600354600160a060020a0316635c975abb6000604051602001526040518163ffffffff1660e060020a028152600401602060405180830381600087803b151561064c57600080fd5b6102c65a03f1151561065d57600080fd5b505050604051805115905061067157600080fd5b600354600160a060020a0316638456cb596040518163ffffffff1660e060020a028152600401600060405180830381600087803b15156106b057600080fd5b6102c65a03f115156106c157600080fd5b5050506001905090565b60075481565b6000805433600160a060020a039081169116146106ed57600080fd5b5063ffffffff166000908152600c6020526040902054600160a060020a031690565b60005433600160a060020a0390811691161461072a57600080fd5b60005460a060020a900460ff16151561074257600080fd5b6000805474ff0000000000000000000000000000000000000000191690557f7805862f689e2f13df9f062ff482ad3ad112aca9e0847911ed832e158c525b3360405160405180910390a1565b60005460a060020a900460ff16156107a557600080fd5b6107ae33610f04565b565b6000805433600160a060020a039081169116146107cc57600080fd5b600354600160a060020a0316635c975abb6000604051602001526040518163ffffffff1660e060020a028152600401602060405180830381600087803b151561081457600080fd5b6102c65a03f1151561082557600080fd5b50505060405180519050151561083a57600080fd5b600354600160a060020a0316633f4ba83a6040518163ffffffff1660e060020a028152600401600060405180830381600087803b15156106b057600080fd5b60005460a060020a900460ff1681565b60005433600160a060020a039081169116146108a457600080fd5b600081116108b157600080fd5b600755565b600254600160a060020a031681565b60095460ff1681565b60005433600160a060020a039081169116146108e957600080fd5b60005460a060020a900460ff161561090057600080fd5b6000805474ff0000000000000000000000000000000000000000191660a060020a1790557f6985a02210a168e66602d3235cb6db0e70f92b3ba4d376a33c0f3d9434bff62560405160405180910390a1565b600054600160a060020a031681565b6000805433600160a060020a0390811691161461097d57600080fd5b50600d5463ffffffff1690565b60008054819033600160a060020a039081169116146109a857600080fd5b60005460a060020a900460ff16156109bf57600080fd5b600091505b600d5463ffffffff9081169083161015610a45575063ffffffff81166000908152600c6020908152604080832054600160a060020a0316808452600b9092529091205460ff168015610a2c5750600160a060020a0381166000908152600a6020526040812054115b15610a3a57610a3a81610f04565b6001909101906109c4565b5050565b60005433600160a060020a03908116911614610a6457600080fd5b60008111610a7157600080fd5b610a8981670de0b6b3a764000063ffffffff610db016565b60065550565b60005433600160a060020a03908116911614610aaa57600080fd5b600454600554101580610abe575060085442115b1515610ac957600080fd5b60095460ff1615610ad957600080fd5b6009805460ff19166001179055600354600254600160a060020a039182169163a9059cbb9116826370a082313060006040516020015260405160e060020a63ffffffff8416028152600160a060020a039091166004820152602401602060405180830381600087803b1515610b4d57600080fd5b6102c65a03f11515610b5e57600080fd5b5050506040518051905060006040516020015260405160e060020a63ffffffff8516028152600160a060020a0390921660048301526024820152604401602060405180830381600087803b1515610bb457600080fd5b6102c65a03f11515610bc557600080fd5b50505060405180515050565b600b6020526000908152604090205460ff1681565b60005433600160a060020a03908116911614610c0157600080fd5b60005460a060020a900460ff1615610c1857600080fd5b600160a060020a03166000908152600b60205260409020805460ff19166001179055565b600a6020526000908152604090205481565b60055481565b60005433600160a060020a03908116911614610c6f57600080fd5b600160a060020a0381161515610c8457600080fd5b600054600160a060020a0380831691167f8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e060405160405180910390a36000805473ffffffffffffffffffffffffffffffffffffffff1916600160a060020a0392909216919091179055565b6000805433600160a060020a03908116911614610d0b57600080fd5b60005460a060020a900460ff1615610d2257600080fd5b600754610d35908363ffffffff610db016565b600554909150610d4b908263ffffffff610e0416565b600555610d588382610e13565b82600160a060020a03167fa818a22273fc309f0a3682b642c74c5b5c25c0615ff378d07767cd231e19fffc828460405191825260208201526040908101905180910390a2505050565b600354600160a060020a031681565b6000828202831580610dcc5750828482811515610dc957fe5b04145b1515610dd457fe5b9392505050565b600082821115610de757fe5b50900390565b6000808284811515610dfb57fe5b04949350505050565b600082820183811015610dd457fe5b60005460a060020a900460ff1615610e2a57600080fd5b600160a060020a0382166000908152600a60205260409020541515610ebb57600d805463ffffffff9081166000908152600c60209081526040808320805473ffffffffffffffffffffffffffffffffffffffff1916600160a060020a038916908117909155855463ffffffff19811690861660010190951694909417909455918152600b90915220805460ff191690555b600160a060020a0382166000908152600a6020526040902054610ee4908263ffffffff610e0416565b600160a060020a039092166000908152600a602052604090209190915550565b6000805460a060020a900460ff1615610f1c57600080fd5b600160a060020a0382166000908152600b602052604090205460ff161515610f4357600080fd5b600160a060020a0382166000908152600a602052604081205411610f6657600080fd5b50600160a060020a038082166000908152600a60205260408082208054908390556003549093169163a9059cbb9185918591516020015260405160e060020a63ffffffff8516028152600160a060020a0390921660048301526024820152604401602060405180830381600087803b1515610fe057600080fd5b6102c65a03f11515610ff157600080fd5b50505060405180515050600160a060020a0382167fd4ff88d569a7ad2bfd6b17da9dbe82e2ccc016fd0051e08ff1ad1e6fe8fc9b558260405190815260200160405180910390a250505600a165627a7a72305820da695b8ba79e7ceaf1977044704dbeb5e6fceb6322292acefd3378bd438c5fa20029`

// DeployClaimableCrowdsale deploys a new Ethereum contract, binding an instance of ClaimableCrowdsale to it.
func DeployClaimableCrowdsale(auth *bind.TransactOpts, backend bind.ContractBackend, _tokenAddress common.Address, _bankAddress common.Address, _beneficiaryAddress common.Address, _tokenRate *big.Int, _minBuyableAmount *big.Int, _maxTokensAmount *big.Int, _endDate *big.Int) (common.Address, *types.Transaction, *ClaimableCrowdsale, error) {
	parsed, err := abi.JSON(strings.NewReader(ClaimableCrowdsaleABI))
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	address, tx, contract, err := bind.DeployContract(auth, parsed, common.FromHex(ClaimableCrowdsaleBin), backend, _tokenAddress, _bankAddress, _beneficiaryAddress, _tokenRate, _minBuyableAmount, _maxTokensAmount, _endDate)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &ClaimableCrowdsale{ClaimableCrowdsaleCaller: ClaimableCrowdsaleCaller{contract: contract}, ClaimableCrowdsaleTransactor: ClaimableCrowdsaleTransactor{contract: contract}}, nil
}

// ClaimableCrowdsale is an auto generated Go binding around an Ethereum contract.
type ClaimableCrowdsale struct {
	ClaimableCrowdsaleCaller     // Read-only binding to the contract
	ClaimableCrowdsaleTransactor // Write-only binding to the contract
}

// ClaimableCrowdsaleCaller is an auto generated read-only Go binding around an Ethereum contract.
type ClaimableCrowdsaleCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ClaimableCrowdsaleTransactor is an auto generated write-only Go binding around an Ethereum contract.
type ClaimableCrowdsaleTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ClaimableCrowdsaleSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type ClaimableCrowdsaleSession struct {
	Contract     *ClaimableCrowdsale // Generic contract binding to set the session for
	CallOpts     bind.CallOpts       // Call options to use throughout this session
	TransactOpts bind.TransactOpts   // Transaction auth options to use throughout this session
}

// ClaimableCrowdsaleCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type ClaimableCrowdsaleCallerSession struct {
	Contract *ClaimableCrowdsaleCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts             // Call options to use throughout this session
}

// ClaimableCrowdsaleTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type ClaimableCrowdsaleTransactorSession struct {
	Contract     *ClaimableCrowdsaleTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts             // Transaction auth options to use throughout this session
}

// ClaimableCrowdsaleRaw is an auto generated low-level Go binding around an Ethereum contract.
type ClaimableCrowdsaleRaw struct {
	Contract *ClaimableCrowdsale // Generic contract binding to access the raw methods on
}

// ClaimableCrowdsaleCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type ClaimableCrowdsaleCallerRaw struct {
	Contract *ClaimableCrowdsaleCaller // Generic read-only contract binding to access the raw methods on
}

// ClaimableCrowdsaleTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type ClaimableCrowdsaleTransactorRaw struct {
	Contract *ClaimableCrowdsaleTransactor // Generic write-only contract binding to access the raw methods on
}

// NewClaimableCrowdsale creates a new instance of ClaimableCrowdsale, bound to a specific deployed contract.
func NewClaimableCrowdsale(address common.Address, backend bind.ContractBackend) (*ClaimableCrowdsale, error) {
	contract, err := bindClaimableCrowdsale(address, backend, backend)
	if err != nil {
		return nil, err
	}
	return &ClaimableCrowdsale{ClaimableCrowdsaleCaller: ClaimableCrowdsaleCaller{contract: contract}, ClaimableCrowdsaleTransactor: ClaimableCrowdsaleTransactor{contract: contract}}, nil
}

// NewClaimableCrowdsaleCaller creates a new read-only instance of ClaimableCrowdsale, bound to a specific deployed contract.
func NewClaimableCrowdsaleCaller(address common.Address, caller bind.ContractCaller) (*ClaimableCrowdsaleCaller, error) {
	contract, err := bindClaimableCrowdsale(address, caller, nil)
	if err != nil {
		return nil, err
	}
	return &ClaimableCrowdsaleCaller{contract: contract}, nil
}

// NewClaimableCrowdsaleTransactor creates a new write-only instance of ClaimableCrowdsale, bound to a specific deployed contract.
func NewClaimableCrowdsaleTransactor(address common.Address, transactor bind.ContractTransactor) (*ClaimableCrowdsaleTransactor, error) {
	contract, err := bindClaimableCrowdsale(address, nil, transactor)
	if err != nil {
		return nil, err
	}
	return &ClaimableCrowdsaleTransactor{contract: contract}, nil
}

// bindClaimableCrowdsale binds a generic wrapper to an already deployed contract.
func bindClaimableCrowdsale(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(ClaimableCrowdsaleABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_ClaimableCrowdsale *ClaimableCrowdsaleRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _ClaimableCrowdsale.Contract.ClaimableCrowdsaleCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_ClaimableCrowdsale *ClaimableCrowdsaleRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ClaimableCrowdsale.Contract.ClaimableCrowdsaleTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_ClaimableCrowdsale *ClaimableCrowdsaleRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _ClaimableCrowdsale.Contract.ClaimableCrowdsaleTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_ClaimableCrowdsale *ClaimableCrowdsaleCallerRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _ClaimableCrowdsale.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_ClaimableCrowdsale *ClaimableCrowdsaleTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ClaimableCrowdsale.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_ClaimableCrowdsale *ClaimableCrowdsaleTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _ClaimableCrowdsale.Contract.contract.Transact(opts, method, params...)
}

// Approved is a free data retrieval call binding the contract method 0xd8b964e6.
//
// Solidity: function approved( address) constant returns(bool)
func (_ClaimableCrowdsale *ClaimableCrowdsaleCaller) Approved(opts *bind.CallOpts, arg0 common.Address) (bool, error) {
	var (
		ret0 = new(bool)
	)
	out := ret0
	err := _ClaimableCrowdsale.contract.Call(opts, out, "approved", arg0)
	return *ret0, err
}

// Approved is a free data retrieval call binding the contract method 0xd8b964e6.
//
// Solidity: function approved( address) constant returns(bool)
func (_ClaimableCrowdsale *ClaimableCrowdsaleSession) Approved(arg0 common.Address) (bool, error) {
	return _ClaimableCrowdsale.Contract.Approved(&_ClaimableCrowdsale.CallOpts, arg0)
}

// Approved is a free data retrieval call binding the contract method 0xd8b964e6.
//
// Solidity: function approved( address) constant returns(bool)
func (_ClaimableCrowdsale *ClaimableCrowdsaleCallerSession) Approved(arg0 common.Address) (bool, error) {
	return _ClaimableCrowdsale.Contract.Approved(&_ClaimableCrowdsale.CallOpts, arg0)
}

// BankAddress is a free data retrieval call binding the contract method 0x7822ed49.
//
// Solidity: function bankAddress() constant returns(address)
func (_ClaimableCrowdsale *ClaimableCrowdsaleCaller) BankAddress(opts *bind.CallOpts) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _ClaimableCrowdsale.contract.Call(opts, out, "bankAddress")
	return *ret0, err
}

// BankAddress is a free data retrieval call binding the contract method 0x7822ed49.
//
// Solidity: function bankAddress() constant returns(address)
func (_ClaimableCrowdsale *ClaimableCrowdsaleSession) BankAddress() (common.Address, error) {
	return _ClaimableCrowdsale.Contract.BankAddress(&_ClaimableCrowdsale.CallOpts)
}

// BankAddress is a free data retrieval call binding the contract method 0x7822ed49.
//
// Solidity: function bankAddress() constant returns(address)
func (_ClaimableCrowdsale *ClaimableCrowdsaleCallerSession) BankAddress() (common.Address, error) {
	return _ClaimableCrowdsale.Contract.BankAddress(&_ClaimableCrowdsale.CallOpts)
}

// GetReceiver is a free data retrieval call binding the contract method 0x3c78fe07.
//
// Solidity: function getReceiver(i uint32) constant returns(address)
func (_ClaimableCrowdsale *ClaimableCrowdsaleCaller) GetReceiver(opts *bind.CallOpts, i uint32) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _ClaimableCrowdsale.contract.Call(opts, out, "getReceiver", i)
	return *ret0, err
}

// GetReceiver is a free data retrieval call binding the contract method 0x3c78fe07.
//
// Solidity: function getReceiver(i uint32) constant returns(address)
func (_ClaimableCrowdsale *ClaimableCrowdsaleSession) GetReceiver(i uint32) (common.Address, error) {
	return _ClaimableCrowdsale.Contract.GetReceiver(&_ClaimableCrowdsale.CallOpts, i)
}

// GetReceiver is a free data retrieval call binding the contract method 0x3c78fe07.
//
// Solidity: function getReceiver(i uint32) constant returns(address)
func (_ClaimableCrowdsale *ClaimableCrowdsaleCallerSession) GetReceiver(i uint32) (common.Address, error) {
	return _ClaimableCrowdsale.Contract.GetReceiver(&_ClaimableCrowdsale.CallOpts, i)
}

// GetReceiversCount is a free data retrieval call binding the contract method 0xa72a05f7.
//
// Solidity: function getReceiversCount() constant returns(uint32)
func (_ClaimableCrowdsale *ClaimableCrowdsaleCaller) GetReceiversCount(opts *bind.CallOpts) (uint32, error) {
	var (
		ret0 = new(uint32)
	)
	out := ret0
	err := _ClaimableCrowdsale.contract.Call(opts, out, "getReceiversCount")
	return *ret0, err
}

// GetReceiversCount is a free data retrieval call binding the contract method 0xa72a05f7.
//
// Solidity: function getReceiversCount() constant returns(uint32)
func (_ClaimableCrowdsale *ClaimableCrowdsaleSession) GetReceiversCount() (uint32, error) {
	return _ClaimableCrowdsale.Contract.GetReceiversCount(&_ClaimableCrowdsale.CallOpts)
}

// GetReceiversCount is a free data retrieval call binding the contract method 0xa72a05f7.
//
// Solidity: function getReceiversCount() constant returns(uint32)
func (_ClaimableCrowdsale *ClaimableCrowdsaleCallerSession) GetReceiversCount() (uint32, error) {
	return _ClaimableCrowdsale.Contract.GetReceiversCount(&_ClaimableCrowdsale.CallOpts)
}

// IsFinished is a free data retrieval call binding the contract method 0x7b352962.
//
// Solidity: function isFinished() constant returns(bool)
func (_ClaimableCrowdsale *ClaimableCrowdsaleCaller) IsFinished(opts *bind.CallOpts) (bool, error) {
	var (
		ret0 = new(bool)
	)
	out := ret0
	err := _ClaimableCrowdsale.contract.Call(opts, out, "isFinished")
	return *ret0, err
}

// IsFinished is a free data retrieval call binding the contract method 0x7b352962.
//
// Solidity: function isFinished() constant returns(bool)
func (_ClaimableCrowdsale *ClaimableCrowdsaleSession) IsFinished() (bool, error) {
	return _ClaimableCrowdsale.Contract.IsFinished(&_ClaimableCrowdsale.CallOpts)
}

// IsFinished is a free data retrieval call binding the contract method 0x7b352962.
//
// Solidity: function isFinished() constant returns(bool)
func (_ClaimableCrowdsale *ClaimableCrowdsaleCallerSession) IsFinished() (bool, error) {
	return _ClaimableCrowdsale.Contract.IsFinished(&_ClaimableCrowdsale.CallOpts)
}

// IssuedTokensAmount is a free data retrieval call binding the contract method 0xe69b414b.
//
// Solidity: function issuedTokensAmount() constant returns(uint256)
func (_ClaimableCrowdsale *ClaimableCrowdsaleCaller) IssuedTokensAmount(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _ClaimableCrowdsale.contract.Call(opts, out, "issuedTokensAmount")
	return *ret0, err
}

// IssuedTokensAmount is a free data retrieval call binding the contract method 0xe69b414b.
//
// Solidity: function issuedTokensAmount() constant returns(uint256)
func (_ClaimableCrowdsale *ClaimableCrowdsaleSession) IssuedTokensAmount() (*big.Int, error) {
	return _ClaimableCrowdsale.Contract.IssuedTokensAmount(&_ClaimableCrowdsale.CallOpts)
}

// IssuedTokensAmount is a free data retrieval call binding the contract method 0xe69b414b.
//
// Solidity: function issuedTokensAmount() constant returns(uint256)
func (_ClaimableCrowdsale *ClaimableCrowdsaleCallerSession) IssuedTokensAmount() (*big.Int, error) {
	return _ClaimableCrowdsale.Contract.IssuedTokensAmount(&_ClaimableCrowdsale.CallOpts)
}

// MaxTokensAmount is a free data retrieval call binding the contract method 0x0a384665.
//
// Solidity: function maxTokensAmount() constant returns(uint256)
func (_ClaimableCrowdsale *ClaimableCrowdsaleCaller) MaxTokensAmount(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _ClaimableCrowdsale.contract.Call(opts, out, "maxTokensAmount")
	return *ret0, err
}

// MaxTokensAmount is a free data retrieval call binding the contract method 0x0a384665.
//
// Solidity: function maxTokensAmount() constant returns(uint256)
func (_ClaimableCrowdsale *ClaimableCrowdsaleSession) MaxTokensAmount() (*big.Int, error) {
	return _ClaimableCrowdsale.Contract.MaxTokensAmount(&_ClaimableCrowdsale.CallOpts)
}

// MaxTokensAmount is a free data retrieval call binding the contract method 0x0a384665.
//
// Solidity: function maxTokensAmount() constant returns(uint256)
func (_ClaimableCrowdsale *ClaimableCrowdsaleCallerSession) MaxTokensAmount() (*big.Int, error) {
	return _ClaimableCrowdsale.Contract.MaxTokensAmount(&_ClaimableCrowdsale.CallOpts)
}

// MinBuyableAmount is a free data retrieval call binding the contract method 0x2511b182.
//
// Solidity: function minBuyableAmount() constant returns(uint256)
func (_ClaimableCrowdsale *ClaimableCrowdsaleCaller) MinBuyableAmount(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _ClaimableCrowdsale.contract.Call(opts, out, "minBuyableAmount")
	return *ret0, err
}

// MinBuyableAmount is a free data retrieval call binding the contract method 0x2511b182.
//
// Solidity: function minBuyableAmount() constant returns(uint256)
func (_ClaimableCrowdsale *ClaimableCrowdsaleSession) MinBuyableAmount() (*big.Int, error) {
	return _ClaimableCrowdsale.Contract.MinBuyableAmount(&_ClaimableCrowdsale.CallOpts)
}

// MinBuyableAmount is a free data retrieval call binding the contract method 0x2511b182.
//
// Solidity: function minBuyableAmount() constant returns(uint256)
func (_ClaimableCrowdsale *ClaimableCrowdsaleCallerSession) MinBuyableAmount() (*big.Int, error) {
	return _ClaimableCrowdsale.Contract.MinBuyableAmount(&_ClaimableCrowdsale.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() constant returns(address)
func (_ClaimableCrowdsale *ClaimableCrowdsaleCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _ClaimableCrowdsale.contract.Call(opts, out, "owner")
	return *ret0, err
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() constant returns(address)
func (_ClaimableCrowdsale *ClaimableCrowdsaleSession) Owner() (common.Address, error) {
	return _ClaimableCrowdsale.Contract.Owner(&_ClaimableCrowdsale.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() constant returns(address)
func (_ClaimableCrowdsale *ClaimableCrowdsaleCallerSession) Owner() (common.Address, error) {
	return _ClaimableCrowdsale.Contract.Owner(&_ClaimableCrowdsale.CallOpts)
}

// Paused is a free data retrieval call binding the contract method 0x5c975abb.
//
// Solidity: function paused() constant returns(bool)
func (_ClaimableCrowdsale *ClaimableCrowdsaleCaller) Paused(opts *bind.CallOpts) (bool, error) {
	var (
		ret0 = new(bool)
	)
	out := ret0
	err := _ClaimableCrowdsale.contract.Call(opts, out, "paused")
	return *ret0, err
}

// Paused is a free data retrieval call binding the contract method 0x5c975abb.
//
// Solidity: function paused() constant returns(bool)
func (_ClaimableCrowdsale *ClaimableCrowdsaleSession) Paused() (bool, error) {
	return _ClaimableCrowdsale.Contract.Paused(&_ClaimableCrowdsale.CallOpts)
}

// Paused is a free data retrieval call binding the contract method 0x5c975abb.
//
// Solidity: function paused() constant returns(bool)
func (_ClaimableCrowdsale *ClaimableCrowdsaleCallerSession) Paused() (bool, error) {
	return _ClaimableCrowdsale.Contract.Paused(&_ClaimableCrowdsale.CallOpts)
}

// Token is a free data retrieval call binding the contract method 0xfc0c546a.
//
// Solidity: function token() constant returns(address)
func (_ClaimableCrowdsale *ClaimableCrowdsaleCaller) Token(opts *bind.CallOpts) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _ClaimableCrowdsale.contract.Call(opts, out, "token")
	return *ret0, err
}

// Token is a free data retrieval call binding the contract method 0xfc0c546a.
//
// Solidity: function token() constant returns(address)
func (_ClaimableCrowdsale *ClaimableCrowdsaleSession) Token() (common.Address, error) {
	return _ClaimableCrowdsale.Contract.Token(&_ClaimableCrowdsale.CallOpts)
}

// Token is a free data retrieval call binding the contract method 0xfc0c546a.
//
// Solidity: function token() constant returns(address)
func (_ClaimableCrowdsale *ClaimableCrowdsaleCallerSession) Token() (common.Address, error) {
	return _ClaimableCrowdsale.Contract.Token(&_ClaimableCrowdsale.CallOpts)
}

// TokenRate is a free data retrieval call binding the contract method 0x31711884.
//
// Solidity: function tokenRate() constant returns(uint256)
func (_ClaimableCrowdsale *ClaimableCrowdsaleCaller) TokenRate(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _ClaimableCrowdsale.contract.Call(opts, out, "tokenRate")
	return *ret0, err
}

// TokenRate is a free data retrieval call binding the contract method 0x31711884.
//
// Solidity: function tokenRate() constant returns(uint256)
func (_ClaimableCrowdsale *ClaimableCrowdsaleSession) TokenRate() (*big.Int, error) {
	return _ClaimableCrowdsale.Contract.TokenRate(&_ClaimableCrowdsale.CallOpts)
}

// TokenRate is a free data retrieval call binding the contract method 0x31711884.
//
// Solidity: function tokenRate() constant returns(uint256)
func (_ClaimableCrowdsale *ClaimableCrowdsaleCallerSession) TokenRate() (*big.Int, error) {
	return _ClaimableCrowdsale.Contract.TokenRate(&_ClaimableCrowdsale.CallOpts)
}

// Tokens is a free data retrieval call binding the contract method 0xe4860339.
//
// Solidity: function tokens( address) constant returns(uint256)
func (_ClaimableCrowdsale *ClaimableCrowdsaleCaller) Tokens(opts *bind.CallOpts, arg0 common.Address) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _ClaimableCrowdsale.contract.Call(opts, out, "tokens", arg0)
	return *ret0, err
}

// Tokens is a free data retrieval call binding the contract method 0xe4860339.
//
// Solidity: function tokens( address) constant returns(uint256)
func (_ClaimableCrowdsale *ClaimableCrowdsaleSession) Tokens(arg0 common.Address) (*big.Int, error) {
	return _ClaimableCrowdsale.Contract.Tokens(&_ClaimableCrowdsale.CallOpts, arg0)
}

// Tokens is a free data retrieval call binding the contract method 0xe4860339.
//
// Solidity: function tokens( address) constant returns(uint256)
func (_ClaimableCrowdsale *ClaimableCrowdsaleCallerSession) Tokens(arg0 common.Address) (*big.Int, error) {
	return _ClaimableCrowdsale.Contract.Tokens(&_ClaimableCrowdsale.CallOpts, arg0)
}

// Add is a paid mutator transaction binding the contract method 0xf5d82b6b.
//
// Solidity: function add(_receiver address, _equivalentEthAmount uint256) returns()
func (_ClaimableCrowdsale *ClaimableCrowdsaleTransactor) Add(opts *bind.TransactOpts, _receiver common.Address, _equivalentEthAmount *big.Int) (*types.Transaction, error) {
	return _ClaimableCrowdsale.contract.Transact(opts, "add", _receiver, _equivalentEthAmount)
}

// Add is a paid mutator transaction binding the contract method 0xf5d82b6b.
//
// Solidity: function add(_receiver address, _equivalentEthAmount uint256) returns()
func (_ClaimableCrowdsale *ClaimableCrowdsaleSession) Add(_receiver common.Address, _equivalentEthAmount *big.Int) (*types.Transaction, error) {
	return _ClaimableCrowdsale.Contract.Add(&_ClaimableCrowdsale.TransactOpts, _receiver, _equivalentEthAmount)
}

// Add is a paid mutator transaction binding the contract method 0xf5d82b6b.
//
// Solidity: function add(_receiver address, _equivalentEthAmount uint256) returns()
func (_ClaimableCrowdsale *ClaimableCrowdsaleTransactorSession) Add(_receiver common.Address, _equivalentEthAmount *big.Int) (*types.Transaction, error) {
	return _ClaimableCrowdsale.Contract.Add(&_ClaimableCrowdsale.TransactOpts, _receiver, _equivalentEthAmount)
}

// Approve is a paid mutator transaction binding the contract method 0xdaea85c5.
//
// Solidity: function approve(_receiver address) returns()
func (_ClaimableCrowdsale *ClaimableCrowdsaleTransactor) Approve(opts *bind.TransactOpts, _receiver common.Address) (*types.Transaction, error) {
	return _ClaimableCrowdsale.contract.Transact(opts, "approve", _receiver)
}

// Approve is a paid mutator transaction binding the contract method 0xdaea85c5.
//
// Solidity: function approve(_receiver address) returns()
func (_ClaimableCrowdsale *ClaimableCrowdsaleSession) Approve(_receiver common.Address) (*types.Transaction, error) {
	return _ClaimableCrowdsale.Contract.Approve(&_ClaimableCrowdsale.TransactOpts, _receiver)
}

// Approve is a paid mutator transaction binding the contract method 0xdaea85c5.
//
// Solidity: function approve(_receiver address) returns()
func (_ClaimableCrowdsale *ClaimableCrowdsaleTransactorSession) Approve(_receiver common.Address) (*types.Transaction, error) {
	return _ClaimableCrowdsale.Contract.Approve(&_ClaimableCrowdsale.TransactOpts, _receiver)
}

// Buy is a paid mutator transaction binding the contract method 0xa6f2ae3a.
//
// Solidity: function buy() returns()
func (_ClaimableCrowdsale *ClaimableCrowdsaleTransactor) Buy(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ClaimableCrowdsale.contract.Transact(opts, "buy")
}

// Buy is a paid mutator transaction binding the contract method 0xa6f2ae3a.
//
// Solidity: function buy() returns()
func (_ClaimableCrowdsale *ClaimableCrowdsaleSession) Buy() (*types.Transaction, error) {
	return _ClaimableCrowdsale.Contract.Buy(&_ClaimableCrowdsale.TransactOpts)
}

// Buy is a paid mutator transaction binding the contract method 0xa6f2ae3a.
//
// Solidity: function buy() returns()
func (_ClaimableCrowdsale *ClaimableCrowdsaleTransactorSession) Buy() (*types.Transaction, error) {
	return _ClaimableCrowdsale.Contract.Buy(&_ClaimableCrowdsale.TransactOpts)
}

// Claim is a paid mutator transaction binding the contract method 0x4e71d92d.
//
// Solidity: function claim() returns()
func (_ClaimableCrowdsale *ClaimableCrowdsaleTransactor) Claim(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ClaimableCrowdsale.contract.Transact(opts, "claim")
}

// Claim is a paid mutator transaction binding the contract method 0x4e71d92d.
//
// Solidity: function claim() returns()
func (_ClaimableCrowdsale *ClaimableCrowdsaleSession) Claim() (*types.Transaction, error) {
	return _ClaimableCrowdsale.Contract.Claim(&_ClaimableCrowdsale.TransactOpts)
}

// Claim is a paid mutator transaction binding the contract method 0x4e71d92d.
//
// Solidity: function claim() returns()
func (_ClaimableCrowdsale *ClaimableCrowdsaleTransactorSession) Claim() (*types.Transaction, error) {
	return _ClaimableCrowdsale.Contract.Claim(&_ClaimableCrowdsale.TransactOpts)
}

// ClaimAll is a paid mutator transaction binding the contract method 0xd1058e59.
//
// Solidity: function claimAll() returns()
func (_ClaimableCrowdsale *ClaimableCrowdsaleTransactor) ClaimAll(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ClaimableCrowdsale.contract.Transact(opts, "claimAll")
}

// ClaimAll is a paid mutator transaction binding the contract method 0xd1058e59.
//
// Solidity: function claimAll() returns()
func (_ClaimableCrowdsale *ClaimableCrowdsaleSession) ClaimAll() (*types.Transaction, error) {
	return _ClaimableCrowdsale.Contract.ClaimAll(&_ClaimableCrowdsale.TransactOpts)
}

// ClaimAll is a paid mutator transaction binding the contract method 0xd1058e59.
//
// Solidity: function claimAll() returns()
func (_ClaimableCrowdsale *ClaimableCrowdsaleTransactorSession) ClaimAll() (*types.Transaction, error) {
	return _ClaimableCrowdsale.Contract.ClaimAll(&_ClaimableCrowdsale.TransactOpts)
}

// ClaimOne is a paid mutator transaction binding the contract method 0x0262dc13.
//
// Solidity: function claimOne(_receiver address) returns()
func (_ClaimableCrowdsale *ClaimableCrowdsaleTransactor) ClaimOne(opts *bind.TransactOpts, _receiver common.Address) (*types.Transaction, error) {
	return _ClaimableCrowdsale.contract.Transact(opts, "claimOne", _receiver)
}

// ClaimOne is a paid mutator transaction binding the contract method 0x0262dc13.
//
// Solidity: function claimOne(_receiver address) returns()
func (_ClaimableCrowdsale *ClaimableCrowdsaleSession) ClaimOne(_receiver common.Address) (*types.Transaction, error) {
	return _ClaimableCrowdsale.Contract.ClaimOne(&_ClaimableCrowdsale.TransactOpts, _receiver)
}

// ClaimOne is a paid mutator transaction binding the contract method 0x0262dc13.
//
// Solidity: function claimOne(_receiver address) returns()
func (_ClaimableCrowdsale *ClaimableCrowdsaleTransactorSession) ClaimOne(_receiver common.Address) (*types.Transaction, error) {
	return _ClaimableCrowdsale.Contract.ClaimOne(&_ClaimableCrowdsale.TransactOpts, _receiver)
}

// Finish is a paid mutator transaction binding the contract method 0xd56b2889.
//
// Solidity: function finish() returns()
func (_ClaimableCrowdsale *ClaimableCrowdsaleTransactor) Finish(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ClaimableCrowdsale.contract.Transact(opts, "finish")
}

// Finish is a paid mutator transaction binding the contract method 0xd56b2889.
//
// Solidity: function finish() returns()
func (_ClaimableCrowdsale *ClaimableCrowdsaleSession) Finish() (*types.Transaction, error) {
	return _ClaimableCrowdsale.Contract.Finish(&_ClaimableCrowdsale.TransactOpts)
}

// Finish is a paid mutator transaction binding the contract method 0xd56b2889.
//
// Solidity: function finish() returns()
func (_ClaimableCrowdsale *ClaimableCrowdsaleTransactorSession) Finish() (*types.Transaction, error) {
	return _ClaimableCrowdsale.Contract.Finish(&_ClaimableCrowdsale.TransactOpts)
}

// Pause is a paid mutator transaction binding the contract method 0x8456cb59.
//
// Solidity: function pause() returns()
func (_ClaimableCrowdsale *ClaimableCrowdsaleTransactor) Pause(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ClaimableCrowdsale.contract.Transact(opts, "pause")
}

// Pause is a paid mutator transaction binding the contract method 0x8456cb59.
//
// Solidity: function pause() returns()
func (_ClaimableCrowdsale *ClaimableCrowdsaleSession) Pause() (*types.Transaction, error) {
	return _ClaimableCrowdsale.Contract.Pause(&_ClaimableCrowdsale.TransactOpts)
}

// Pause is a paid mutator transaction binding the contract method 0x8456cb59.
//
// Solidity: function pause() returns()
func (_ClaimableCrowdsale *ClaimableCrowdsaleTransactorSession) Pause() (*types.Transaction, error) {
	return _ClaimableCrowdsale.Contract.Pause(&_ClaimableCrowdsale.TransactOpts)
}

// PauseToken is a paid mutator transaction binding the contract method 0x2c349627.
//
// Solidity: function pauseToken() returns(bool)
func (_ClaimableCrowdsale *ClaimableCrowdsaleTransactor) PauseToken(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ClaimableCrowdsale.contract.Transact(opts, "pauseToken")
}

// PauseToken is a paid mutator transaction binding the contract method 0x2c349627.
//
// Solidity: function pauseToken() returns(bool)
func (_ClaimableCrowdsale *ClaimableCrowdsaleSession) PauseToken() (*types.Transaction, error) {
	return _ClaimableCrowdsale.Contract.PauseToken(&_ClaimableCrowdsale.TransactOpts)
}

// PauseToken is a paid mutator transaction binding the contract method 0x2c349627.
//
// Solidity: function pauseToken() returns(bool)
func (_ClaimableCrowdsale *ClaimableCrowdsaleTransactorSession) PauseToken() (*types.Transaction, error) {
	return _ClaimableCrowdsale.Contract.PauseToken(&_ClaimableCrowdsale.TransactOpts)
}

// SetMinBuyableAmount is a paid mutator transaction binding the contract method 0xd561ef2b.
//
// Solidity: function setMinBuyableAmount(_minBuyableAmount uint256) returns()
func (_ClaimableCrowdsale *ClaimableCrowdsaleTransactor) SetMinBuyableAmount(opts *bind.TransactOpts, _minBuyableAmount *big.Int) (*types.Transaction, error) {
	return _ClaimableCrowdsale.contract.Transact(opts, "setMinBuyableAmount", _minBuyableAmount)
}

// SetMinBuyableAmount is a paid mutator transaction binding the contract method 0xd561ef2b.
//
// Solidity: function setMinBuyableAmount(_minBuyableAmount uint256) returns()
func (_ClaimableCrowdsale *ClaimableCrowdsaleSession) SetMinBuyableAmount(_minBuyableAmount *big.Int) (*types.Transaction, error) {
	return _ClaimableCrowdsale.Contract.SetMinBuyableAmount(&_ClaimableCrowdsale.TransactOpts, _minBuyableAmount)
}

// SetMinBuyableAmount is a paid mutator transaction binding the contract method 0xd561ef2b.
//
// Solidity: function setMinBuyableAmount(_minBuyableAmount uint256) returns()
func (_ClaimableCrowdsale *ClaimableCrowdsaleTransactorSession) SetMinBuyableAmount(_minBuyableAmount *big.Int) (*types.Transaction, error) {
	return _ClaimableCrowdsale.Contract.SetMinBuyableAmount(&_ClaimableCrowdsale.TransactOpts, _minBuyableAmount)
}

// SetTokenRate is a paid mutator transaction binding the contract method 0x61241c28.
//
// Solidity: function setTokenRate(_tokenRate uint256) returns()
func (_ClaimableCrowdsale *ClaimableCrowdsaleTransactor) SetTokenRate(opts *bind.TransactOpts, _tokenRate *big.Int) (*types.Transaction, error) {
	return _ClaimableCrowdsale.contract.Transact(opts, "setTokenRate", _tokenRate)
}

// SetTokenRate is a paid mutator transaction binding the contract method 0x61241c28.
//
// Solidity: function setTokenRate(_tokenRate uint256) returns()
func (_ClaimableCrowdsale *ClaimableCrowdsaleSession) SetTokenRate(_tokenRate *big.Int) (*types.Transaction, error) {
	return _ClaimableCrowdsale.Contract.SetTokenRate(&_ClaimableCrowdsale.TransactOpts, _tokenRate)
}

// SetTokenRate is a paid mutator transaction binding the contract method 0x61241c28.
//
// Solidity: function setTokenRate(_tokenRate uint256) returns()
func (_ClaimableCrowdsale *ClaimableCrowdsaleTransactorSession) SetTokenRate(_tokenRate *big.Int) (*types.Transaction, error) {
	return _ClaimableCrowdsale.Contract.SetTokenRate(&_ClaimableCrowdsale.TransactOpts, _tokenRate)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(newOwner address) returns()
func (_ClaimableCrowdsale *ClaimableCrowdsaleTransactor) TransferOwnership(opts *bind.TransactOpts, newOwner common.Address) (*types.Transaction, error) {
	return _ClaimableCrowdsale.contract.Transact(opts, "transferOwnership", newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(newOwner address) returns()
func (_ClaimableCrowdsale *ClaimableCrowdsaleSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _ClaimableCrowdsale.Contract.TransferOwnership(&_ClaimableCrowdsale.TransactOpts, newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(newOwner address) returns()
func (_ClaimableCrowdsale *ClaimableCrowdsaleTransactorSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _ClaimableCrowdsale.Contract.TransferOwnership(&_ClaimableCrowdsale.TransactOpts, newOwner)
}

// Unpause is a paid mutator transaction binding the contract method 0x3f4ba83a.
//
// Solidity: function unpause() returns()
func (_ClaimableCrowdsale *ClaimableCrowdsaleTransactor) Unpause(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ClaimableCrowdsale.contract.Transact(opts, "unpause")
}

// Unpause is a paid mutator transaction binding the contract method 0x3f4ba83a.
//
// Solidity: function unpause() returns()
func (_ClaimableCrowdsale *ClaimableCrowdsaleSession) Unpause() (*types.Transaction, error) {
	return _ClaimableCrowdsale.Contract.Unpause(&_ClaimableCrowdsale.TransactOpts)
}

// Unpause is a paid mutator transaction binding the contract method 0x3f4ba83a.
//
// Solidity: function unpause() returns()
func (_ClaimableCrowdsale *ClaimableCrowdsaleTransactorSession) Unpause() (*types.Transaction, error) {
	return _ClaimableCrowdsale.Contract.Unpause(&_ClaimableCrowdsale.TransactOpts)
}

// UnpauseToken is a paid mutator transaction binding the contract method 0x50669a03.
//
// Solidity: function unpauseToken() returns(bool)
func (_ClaimableCrowdsale *ClaimableCrowdsaleTransactor) UnpauseToken(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ClaimableCrowdsale.contract.Transact(opts, "unpauseToken")
}

// UnpauseToken is a paid mutator transaction binding the contract method 0x50669a03.
//
// Solidity: function unpauseToken() returns(bool)
func (_ClaimableCrowdsale *ClaimableCrowdsaleSession) UnpauseToken() (*types.Transaction, error) {
	return _ClaimableCrowdsale.Contract.UnpauseToken(&_ClaimableCrowdsale.TransactOpts)
}

// UnpauseToken is a paid mutator transaction binding the contract method 0x50669a03.
//
// Solidity: function unpauseToken() returns(bool)
func (_ClaimableCrowdsale *ClaimableCrowdsaleTransactorSession) UnpauseToken() (*types.Transaction, error) {
	return _ClaimableCrowdsale.Contract.UnpauseToken(&_ClaimableCrowdsale.TransactOpts)
}

// ERC20ABI is the input ABI used to generate the binding from.
const ERC20ABI = "[{\"constant\":false,\"inputs\":[{\"name\":\"spender\",\"type\":\"address\"},{\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"approve\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"totalSupply\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"from\",\"type\":\"address\"},{\"name\":\"to\",\"type\":\"address\"},{\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"transferFrom\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"who\",\"type\":\"address\"}],\"name\":\"balanceOf\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"to\",\"type\":\"address\"},{\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"transfer\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"owner\",\"type\":\"address\"},{\"name\":\"spender\",\"type\":\"address\"}],\"name\":\"allowance\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"owner\",\"type\":\"address\"},{\"indexed\":true,\"name\":\"spender\",\"type\":\"address\"},{\"indexed\":false,\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"Approval\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"from\",\"type\":\"address\"},{\"indexed\":true,\"name\":\"to\",\"type\":\"address\"},{\"indexed\":false,\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"Transfer\",\"type\":\"event\"}]"

// ERC20Bin is the compiled bytecode used for deploying new contracts.
const ERC20Bin = `0x`

// DeployERC20 deploys a new Ethereum contract, binding an instance of ERC20 to it.
func DeployERC20(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *ERC20, error) {
	parsed, err := abi.JSON(strings.NewReader(ERC20ABI))
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	address, tx, contract, err := bind.DeployContract(auth, parsed, common.FromHex(ERC20Bin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &ERC20{ERC20Caller: ERC20Caller{contract: contract}, ERC20Transactor: ERC20Transactor{contract: contract}}, nil
}

// ERC20 is an auto generated Go binding around an Ethereum contract.
type ERC20 struct {
	ERC20Caller     // Read-only binding to the contract
	ERC20Transactor // Write-only binding to the contract
}

// ERC20Caller is an auto generated read-only Go binding around an Ethereum contract.
type ERC20Caller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ERC20Transactor is an auto generated write-only Go binding around an Ethereum contract.
type ERC20Transactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ERC20Session is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type ERC20Session struct {
	Contract     *ERC20            // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// ERC20CallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type ERC20CallerSession struct {
	Contract *ERC20Caller  // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts // Call options to use throughout this session
}

// ERC20TransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type ERC20TransactorSession struct {
	Contract     *ERC20Transactor  // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// ERC20Raw is an auto generated low-level Go binding around an Ethereum contract.
type ERC20Raw struct {
	Contract *ERC20 // Generic contract binding to access the raw methods on
}

// ERC20CallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type ERC20CallerRaw struct {
	Contract *ERC20Caller // Generic read-only contract binding to access the raw methods on
}

// ERC20TransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type ERC20TransactorRaw struct {
	Contract *ERC20Transactor // Generic write-only contract binding to access the raw methods on
}

// NewERC20 creates a new instance of ERC20, bound to a specific deployed contract.
func NewERC20(address common.Address, backend bind.ContractBackend) (*ERC20, error) {
	contract, err := bindERC20(address, backend, backend)
	if err != nil {
		return nil, err
	}
	return &ERC20{ERC20Caller: ERC20Caller{contract: contract}, ERC20Transactor: ERC20Transactor{contract: contract}}, nil
}

// NewERC20Caller creates a new read-only instance of ERC20, bound to a specific deployed contract.
func NewERC20Caller(address common.Address, caller bind.ContractCaller) (*ERC20Caller, error) {
	contract, err := bindERC20(address, caller, nil)
	if err != nil {
		return nil, err
	}
	return &ERC20Caller{contract: contract}, nil
}

// NewERC20Transactor creates a new write-only instance of ERC20, bound to a specific deployed contract.
func NewERC20Transactor(address common.Address, transactor bind.ContractTransactor) (*ERC20Transactor, error) {
	contract, err := bindERC20(address, nil, transactor)
	if err != nil {
		return nil, err
	}
	return &ERC20Transactor{contract: contract}, nil
}

// bindERC20 binds a generic wrapper to an already deployed contract.
func bindERC20(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(ERC20ABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_ERC20 *ERC20Raw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _ERC20.Contract.ERC20Caller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_ERC20 *ERC20Raw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ERC20.Contract.ERC20Transactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_ERC20 *ERC20Raw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _ERC20.Contract.ERC20Transactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_ERC20 *ERC20CallerRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _ERC20.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_ERC20 *ERC20TransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ERC20.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_ERC20 *ERC20TransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _ERC20.Contract.contract.Transact(opts, method, params...)
}

// Allowance is a free data retrieval call binding the contract method 0xdd62ed3e.
//
// Solidity: function allowance(owner address, spender address) constant returns(uint256)
func (_ERC20 *ERC20Caller) Allowance(opts *bind.CallOpts, owner common.Address, spender common.Address) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _ERC20.contract.Call(opts, out, "allowance", owner, spender)
	return *ret0, err
}

// Allowance is a free data retrieval call binding the contract method 0xdd62ed3e.
//
// Solidity: function allowance(owner address, spender address) constant returns(uint256)
func (_ERC20 *ERC20Session) Allowance(owner common.Address, spender common.Address) (*big.Int, error) {
	return _ERC20.Contract.Allowance(&_ERC20.CallOpts, owner, spender)
}

// Allowance is a free data retrieval call binding the contract method 0xdd62ed3e.
//
// Solidity: function allowance(owner address, spender address) constant returns(uint256)
func (_ERC20 *ERC20CallerSession) Allowance(owner common.Address, spender common.Address) (*big.Int, error) {
	return _ERC20.Contract.Allowance(&_ERC20.CallOpts, owner, spender)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(who address) constant returns(uint256)
func (_ERC20 *ERC20Caller) BalanceOf(opts *bind.CallOpts, who common.Address) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _ERC20.contract.Call(opts, out, "balanceOf", who)
	return *ret0, err
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(who address) constant returns(uint256)
func (_ERC20 *ERC20Session) BalanceOf(who common.Address) (*big.Int, error) {
	return _ERC20.Contract.BalanceOf(&_ERC20.CallOpts, who)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(who address) constant returns(uint256)
func (_ERC20 *ERC20CallerSession) BalanceOf(who common.Address) (*big.Int, error) {
	return _ERC20.Contract.BalanceOf(&_ERC20.CallOpts, who)
}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() constant returns(uint256)
func (_ERC20 *ERC20Caller) TotalSupply(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _ERC20.contract.Call(opts, out, "totalSupply")
	return *ret0, err
}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() constant returns(uint256)
func (_ERC20 *ERC20Session) TotalSupply() (*big.Int, error) {
	return _ERC20.Contract.TotalSupply(&_ERC20.CallOpts)
}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() constant returns(uint256)
func (_ERC20 *ERC20CallerSession) TotalSupply() (*big.Int, error) {
	return _ERC20.Contract.TotalSupply(&_ERC20.CallOpts)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(spender address, value uint256) returns(bool)
func (_ERC20 *ERC20Transactor) Approve(opts *bind.TransactOpts, spender common.Address, value *big.Int) (*types.Transaction, error) {
	return _ERC20.contract.Transact(opts, "approve", spender, value)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(spender address, value uint256) returns(bool)
func (_ERC20 *ERC20Session) Approve(spender common.Address, value *big.Int) (*types.Transaction, error) {
	return _ERC20.Contract.Approve(&_ERC20.TransactOpts, spender, value)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(spender address, value uint256) returns(bool)
func (_ERC20 *ERC20TransactorSession) Approve(spender common.Address, value *big.Int) (*types.Transaction, error) {
	return _ERC20.Contract.Approve(&_ERC20.TransactOpts, spender, value)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(to address, value uint256) returns(bool)
func (_ERC20 *ERC20Transactor) Transfer(opts *bind.TransactOpts, to common.Address, value *big.Int) (*types.Transaction, error) {
	return _ERC20.contract.Transact(opts, "transfer", to, value)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(to address, value uint256) returns(bool)
func (_ERC20 *ERC20Session) Transfer(to common.Address, value *big.Int) (*types.Transaction, error) {
	return _ERC20.Contract.Transfer(&_ERC20.TransactOpts, to, value)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(to address, value uint256) returns(bool)
func (_ERC20 *ERC20TransactorSession) Transfer(to common.Address, value *big.Int) (*types.Transaction, error) {
	return _ERC20.Contract.Transfer(&_ERC20.TransactOpts, to, value)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(from address, to address, value uint256) returns(bool)
func (_ERC20 *ERC20Transactor) TransferFrom(opts *bind.TransactOpts, from common.Address, to common.Address, value *big.Int) (*types.Transaction, error) {
	return _ERC20.contract.Transact(opts, "transferFrom", from, to, value)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(from address, to address, value uint256) returns(bool)
func (_ERC20 *ERC20Session) TransferFrom(from common.Address, to common.Address, value *big.Int) (*types.Transaction, error) {
	return _ERC20.Contract.TransferFrom(&_ERC20.TransactOpts, from, to, value)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(from address, to address, value uint256) returns(bool)
func (_ERC20 *ERC20TransactorSession) TransferFrom(from common.Address, to common.Address, value *big.Int) (*types.Transaction, error) {
	return _ERC20.Contract.TransferFrom(&_ERC20.TransactOpts, from, to, value)
}

// ERC20BasicABI is the input ABI used to generate the binding from.
const ERC20BasicABI = "[{\"constant\":true,\"inputs\":[],\"name\":\"totalSupply\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"who\",\"type\":\"address\"}],\"name\":\"balanceOf\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"to\",\"type\":\"address\"},{\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"transfer\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"from\",\"type\":\"address\"},{\"indexed\":true,\"name\":\"to\",\"type\":\"address\"},{\"indexed\":false,\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"Transfer\",\"type\":\"event\"}]"

// ERC20BasicBin is the compiled bytecode used for deploying new contracts.
const ERC20BasicBin = `0x`

// DeployERC20Basic deploys a new Ethereum contract, binding an instance of ERC20Basic to it.
func DeployERC20Basic(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *ERC20Basic, error) {
	parsed, err := abi.JSON(strings.NewReader(ERC20BasicABI))
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	address, tx, contract, err := bind.DeployContract(auth, parsed, common.FromHex(ERC20BasicBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &ERC20Basic{ERC20BasicCaller: ERC20BasicCaller{contract: contract}, ERC20BasicTransactor: ERC20BasicTransactor{contract: contract}}, nil
}

// ERC20Basic is an auto generated Go binding around an Ethereum contract.
type ERC20Basic struct {
	ERC20BasicCaller     // Read-only binding to the contract
	ERC20BasicTransactor // Write-only binding to the contract
}

// ERC20BasicCaller is an auto generated read-only Go binding around an Ethereum contract.
type ERC20BasicCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ERC20BasicTransactor is an auto generated write-only Go binding around an Ethereum contract.
type ERC20BasicTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ERC20BasicSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type ERC20BasicSession struct {
	Contract     *ERC20Basic       // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// ERC20BasicCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type ERC20BasicCallerSession struct {
	Contract *ERC20BasicCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts     // Call options to use throughout this session
}

// ERC20BasicTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type ERC20BasicTransactorSession struct {
	Contract     *ERC20BasicTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts     // Transaction auth options to use throughout this session
}

// ERC20BasicRaw is an auto generated low-level Go binding around an Ethereum contract.
type ERC20BasicRaw struct {
	Contract *ERC20Basic // Generic contract binding to access the raw methods on
}

// ERC20BasicCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type ERC20BasicCallerRaw struct {
	Contract *ERC20BasicCaller // Generic read-only contract binding to access the raw methods on
}

// ERC20BasicTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type ERC20BasicTransactorRaw struct {
	Contract *ERC20BasicTransactor // Generic write-only contract binding to access the raw methods on
}

// NewERC20Basic creates a new instance of ERC20Basic, bound to a specific deployed contract.
func NewERC20Basic(address common.Address, backend bind.ContractBackend) (*ERC20Basic, error) {
	contract, err := bindERC20Basic(address, backend, backend)
	if err != nil {
		return nil, err
	}
	return &ERC20Basic{ERC20BasicCaller: ERC20BasicCaller{contract: contract}, ERC20BasicTransactor: ERC20BasicTransactor{contract: contract}}, nil
}

// NewERC20BasicCaller creates a new read-only instance of ERC20Basic, bound to a specific deployed contract.
func NewERC20BasicCaller(address common.Address, caller bind.ContractCaller) (*ERC20BasicCaller, error) {
	contract, err := bindERC20Basic(address, caller, nil)
	if err != nil {
		return nil, err
	}
	return &ERC20BasicCaller{contract: contract}, nil
}

// NewERC20BasicTransactor creates a new write-only instance of ERC20Basic, bound to a specific deployed contract.
func NewERC20BasicTransactor(address common.Address, transactor bind.ContractTransactor) (*ERC20BasicTransactor, error) {
	contract, err := bindERC20Basic(address, nil, transactor)
	if err != nil {
		return nil, err
	}
	return &ERC20BasicTransactor{contract: contract}, nil
}

// bindERC20Basic binds a generic wrapper to an already deployed contract.
func bindERC20Basic(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(ERC20BasicABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_ERC20Basic *ERC20BasicRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _ERC20Basic.Contract.ERC20BasicCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_ERC20Basic *ERC20BasicRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ERC20Basic.Contract.ERC20BasicTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_ERC20Basic *ERC20BasicRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _ERC20Basic.Contract.ERC20BasicTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_ERC20Basic *ERC20BasicCallerRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _ERC20Basic.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_ERC20Basic *ERC20BasicTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ERC20Basic.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_ERC20Basic *ERC20BasicTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _ERC20Basic.Contract.contract.Transact(opts, method, params...)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(who address) constant returns(uint256)
func (_ERC20Basic *ERC20BasicCaller) BalanceOf(opts *bind.CallOpts, who common.Address) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _ERC20Basic.contract.Call(opts, out, "balanceOf", who)
	return *ret0, err
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(who address) constant returns(uint256)
func (_ERC20Basic *ERC20BasicSession) BalanceOf(who common.Address) (*big.Int, error) {
	return _ERC20Basic.Contract.BalanceOf(&_ERC20Basic.CallOpts, who)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(who address) constant returns(uint256)
func (_ERC20Basic *ERC20BasicCallerSession) BalanceOf(who common.Address) (*big.Int, error) {
	return _ERC20Basic.Contract.BalanceOf(&_ERC20Basic.CallOpts, who)
}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() constant returns(uint256)
func (_ERC20Basic *ERC20BasicCaller) TotalSupply(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _ERC20Basic.contract.Call(opts, out, "totalSupply")
	return *ret0, err
}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() constant returns(uint256)
func (_ERC20Basic *ERC20BasicSession) TotalSupply() (*big.Int, error) {
	return _ERC20Basic.Contract.TotalSupply(&_ERC20Basic.CallOpts)
}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() constant returns(uint256)
func (_ERC20Basic *ERC20BasicCallerSession) TotalSupply() (*big.Int, error) {
	return _ERC20Basic.Contract.TotalSupply(&_ERC20Basic.CallOpts)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(to address, value uint256) returns(bool)
func (_ERC20Basic *ERC20BasicTransactor) Transfer(opts *bind.TransactOpts, to common.Address, value *big.Int) (*types.Transaction, error) {
	return _ERC20Basic.contract.Transact(opts, "transfer", to, value)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(to address, value uint256) returns(bool)
func (_ERC20Basic *ERC20BasicSession) Transfer(to common.Address, value *big.Int) (*types.Transaction, error) {
	return _ERC20Basic.Contract.Transfer(&_ERC20Basic.TransactOpts, to, value)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(to address, value uint256) returns(bool)
func (_ERC20Basic *ERC20BasicTransactorSession) Transfer(to common.Address, value *big.Int) (*types.Transaction, error) {
	return _ERC20Basic.Contract.Transfer(&_ERC20Basic.TransactOpts, to, value)
}

// HoQuTokenABI is the input ABI used to generate the binding from.
const HoQuTokenABI = "[{\"constant\":true,\"inputs\":[],\"name\":\"name\",\"outputs\":[{\"name\":\"\",\"type\":\"string\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_spender\",\"type\":\"address\"},{\"name\":\"_value\",\"type\":\"uint256\"}],\"name\":\"approve\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"totalSupply\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_from\",\"type\":\"address\"},{\"name\":\"_to\",\"type\":\"address\"},{\"name\":\"_value\",\"type\":\"uint256\"}],\"name\":\"transferFrom\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"decimals\",\"outputs\":[{\"name\":\"\",\"type\":\"uint32\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[],\"name\":\"unpause\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"paused\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_spender\",\"type\":\"address\"},{\"name\":\"_subtractedValue\",\"type\":\"uint256\"}],\"name\":\"decreaseApproval\",\"outputs\":[{\"name\":\"success\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"_owner\",\"type\":\"address\"}],\"name\":\"balanceOf\",\"outputs\":[{\"name\":\"balance\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[],\"name\":\"pause\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"symbol\",\"outputs\":[{\"name\":\"\",\"type\":\"string\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_to\",\"type\":\"address\"},{\"name\":\"_value\",\"type\":\"uint256\"}],\"name\":\"transfer\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_spender\",\"type\":\"address\"},{\"name\":\"_addedValue\",\"type\":\"uint256\"}],\"name\":\"increaseApproval\",\"outputs\":[{\"name\":\"success\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"_owner\",\"type\":\"address\"},{\"name\":\"_spender\",\"type\":\"address\"}],\"name\":\"allowance\",\"outputs\":[{\"name\":\"remaining\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"transferOwnership\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"name\":\"_totalSupply\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[],\"name\":\"Pause\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[],\"name\":\"Unpause\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"previousOwner\",\"type\":\"address\"},{\"indexed\":true,\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"OwnershipTransferred\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"owner\",\"type\":\"address\"},{\"indexed\":true,\"name\":\"spender\",\"type\":\"address\"},{\"indexed\":false,\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"Approval\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"from\",\"type\":\"address\"},{\"indexed\":true,\"name\":\"to\",\"type\":\"address\"},{\"indexed\":false,\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"Transfer\",\"type\":\"event\"}]"

// HoQuTokenBin is the compiled bytecode used for deploying new contracts.
const HoQuTokenBin = `0x60606040526003805460a060020a60ff0219169055341561001f57600080fd5b604051602080610b978339810160405280805160038054600160a060020a03191633600160a060020a03161790559150506000811161005d57600080fd5b600160a060020a033316600090815260016020526040812082905555610b0f806100886000396000f300606060405236156100e35763ffffffff7c010000000000000000000000000000000000000000000000000000000060003504166306fdde0381146100e8578063095ea7b31461017257806318160ddd146101a857806323b872dd146101cd578063313ce567146101f55780633f4ba83a146102215780635c975abb14610236578063661884631461024957806370a082311461026b5780638456cb591461028a5780638da5cb5b1461029d57806395d89b41146102cc578063a9059cbb146102df578063d73dd62314610301578063dd62ed3e14610323578063f2fde38b14610348575b600080fd5b34156100f357600080fd5b6100fb610367565b60405160208082528190810183818151815260200191508051906020019080838360005b8381101561013757808201518382015260200161011f565b50505050905090810190601f1680156101645780820380516001836020036101000a031916815260200191505b509250505060405180910390f35b341561017d57600080fd5b610194600160a060020a036004351660243561039e565b604051901515815260200160405180910390f35b34156101b357600080fd5b6101bb61040a565b60405190815260200160405180910390f35b34156101d857600080fd5b610194600160a060020a0360043581169060243516604435610410565b341561020057600080fd5b61020861043d565b60405163ffffffff909116815260200160405180910390f35b341561022c57600080fd5b610234610442565b005b341561024157600080fd5b6101946104c1565b341561025457600080fd5b610194600160a060020a03600435166024356104d1565b341561027657600080fd5b6101bb600160a060020a03600435166105cb565b341561029557600080fd5b6102346105e6565b34156102a857600080fd5b6102b061066a565b604051600160a060020a03909116815260200160405180910390f35b34156102d757600080fd5b6100fb610679565b34156102ea57600080fd5b610194600160a060020a03600435166024356106b0565b341561030c57600080fd5b610194600160a060020a03600435166024356106db565b341561032e57600080fd5b6101bb600160a060020a036004358116906024351661077f565b341561035357600080fd5b610234600160a060020a03600435166107aa565b60408051908101604052600a81527f484f515520546f6b656e00000000000000000000000000000000000000000000602082015281565b600160a060020a03338116600081815260026020908152604080832094871680845294909152808220859055909291907f8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b9259085905190815260200160405180910390a350600192915050565b60005481565b60035460009060a060020a900460ff161561042a57600080fd5b610435848484610845565b949350505050565b601281565b60035433600160a060020a0390811691161461045d57600080fd5b60035460a060020a900460ff16151561047557600080fd5b6003805474ff0000000000000000000000000000000000000000191690557f7805862f689e2f13df9f062ff482ad3ad112aca9e0847911ed832e158c525b3360405160405180910390a1565b60035460a060020a900460ff1681565b600160a060020a0333811660009081526002602090815260408083209386168352929052908120548083111561052e57600160a060020a033381166000908152600260209081526040808320938816835292905290812055610565565b61053e818463ffffffff6109c716565b600160a060020a033381166000908152600260209081526040808320938916835292905220555b600160a060020a0333811660008181526002602090815260408083209489168084529490915290819020547f8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925915190815260200160405180910390a35060019392505050565b600160a060020a031660009081526001602052604090205490565b60035433600160a060020a0390811691161461060157600080fd5b60035460a060020a900460ff161561061857600080fd5b6003805474ff0000000000000000000000000000000000000000191660a060020a1790557f6985a02210a168e66602d3235cb6db0e70f92b3ba4d376a33c0f3d9434bff62560405160405180910390a1565b600354600160a060020a031681565b60408051908101604052600381527f4851580000000000000000000000000000000000000000000000000000000000602082015281565b60035460009060a060020a900460ff16156106ca57600080fd5b6106d483836109d9565b9392505050565b600160a060020a033381166000908152600260209081526040808320938616835292905290812054610713908363ffffffff610ad416565b600160a060020a0333811660008181526002602090815260408083209489168084529490915290819020849055919290917f8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b92591905190815260200160405180910390a350600192915050565b600160a060020a03918216600090815260026020908152604080832093909416825291909152205490565b60035433600160a060020a039081169116146107c557600080fd5b600160a060020a03811615156107da57600080fd5b600354600160a060020a0380831691167f8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e060405160405180910390a36003805473ffffffffffffffffffffffffffffffffffffffff1916600160a060020a0392909216919091179055565b6000600160a060020a038316151561085c57600080fd5b600160a060020a03841660009081526001602052604090205482111561088157600080fd5b600160a060020a03808516600090815260026020908152604080832033909416835292905220548211156108b457600080fd5b600160a060020a0384166000908152600160205260409020546108dd908363ffffffff6109c716565b600160a060020a038086166000908152600160205260408082209390935590851681522054610912908363ffffffff610ad416565b600160a060020a0380851660009081526001602090815260408083209490945587831682526002815283822033909316825291909152205461095a908363ffffffff6109c716565b600160a060020a03808616600081815260026020908152604080832033861684529091529081902093909355908516917fddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef9085905190815260200160405180910390a35060019392505050565b6000828211156109d357fe5b50900390565b6000600160a060020a03831615156109f057600080fd5b600160a060020a033316600090815260016020526040902054821115610a1557600080fd5b600160a060020a033316600090815260016020526040902054610a3e908363ffffffff6109c716565b600160a060020a033381166000908152600160205260408082209390935590851681522054610a73908363ffffffff610ad416565b600160a060020a0380851660008181526001602052604090819020939093559133909116907fddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef9085905190815260200160405180910390a350600192915050565b6000828201838110156106d457fe00a165627a7a723058209b6a03e89247e92eae0caf598688102b6b8174c57df8a77361c359233f8e072f0029`

// DeployHoQuToken deploys a new Ethereum contract, binding an instance of HoQuToken to it.
func DeployHoQuToken(auth *bind.TransactOpts, backend bind.ContractBackend, _totalSupply *big.Int) (common.Address, *types.Transaction, *HoQuToken, error) {
	parsed, err := abi.JSON(strings.NewReader(HoQuTokenABI))
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	address, tx, contract, err := bind.DeployContract(auth, parsed, common.FromHex(HoQuTokenBin), backend, _totalSupply)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &HoQuToken{HoQuTokenCaller: HoQuTokenCaller{contract: contract}, HoQuTokenTransactor: HoQuTokenTransactor{contract: contract}}, nil
}

// HoQuToken is an auto generated Go binding around an Ethereum contract.
type HoQuToken struct {
	HoQuTokenCaller     // Read-only binding to the contract
	HoQuTokenTransactor // Write-only binding to the contract
}

// HoQuTokenCaller is an auto generated read-only Go binding around an Ethereum contract.
type HoQuTokenCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// HoQuTokenTransactor is an auto generated write-only Go binding around an Ethereum contract.
type HoQuTokenTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// HoQuTokenSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type HoQuTokenSession struct {
	Contract     *HoQuToken        // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// HoQuTokenCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type HoQuTokenCallerSession struct {
	Contract *HoQuTokenCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts    // Call options to use throughout this session
}

// HoQuTokenTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type HoQuTokenTransactorSession struct {
	Contract     *HoQuTokenTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts    // Transaction auth options to use throughout this session
}

// HoQuTokenRaw is an auto generated low-level Go binding around an Ethereum contract.
type HoQuTokenRaw struct {
	Contract *HoQuToken // Generic contract binding to access the raw methods on
}

// HoQuTokenCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type HoQuTokenCallerRaw struct {
	Contract *HoQuTokenCaller // Generic read-only contract binding to access the raw methods on
}

// HoQuTokenTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type HoQuTokenTransactorRaw struct {
	Contract *HoQuTokenTransactor // Generic write-only contract binding to access the raw methods on
}

// NewHoQuToken creates a new instance of HoQuToken, bound to a specific deployed contract.
func NewHoQuToken(address common.Address, backend bind.ContractBackend) (*HoQuToken, error) {
	contract, err := bindHoQuToken(address, backend, backend)
	if err != nil {
		return nil, err
	}
	return &HoQuToken{HoQuTokenCaller: HoQuTokenCaller{contract: contract}, HoQuTokenTransactor: HoQuTokenTransactor{contract: contract}}, nil
}

// NewHoQuTokenCaller creates a new read-only instance of HoQuToken, bound to a specific deployed contract.
func NewHoQuTokenCaller(address common.Address, caller bind.ContractCaller) (*HoQuTokenCaller, error) {
	contract, err := bindHoQuToken(address, caller, nil)
	if err != nil {
		return nil, err
	}
	return &HoQuTokenCaller{contract: contract}, nil
}

// NewHoQuTokenTransactor creates a new write-only instance of HoQuToken, bound to a specific deployed contract.
func NewHoQuTokenTransactor(address common.Address, transactor bind.ContractTransactor) (*HoQuTokenTransactor, error) {
	contract, err := bindHoQuToken(address, nil, transactor)
	if err != nil {
		return nil, err
	}
	return &HoQuTokenTransactor{contract: contract}, nil
}

// bindHoQuToken binds a generic wrapper to an already deployed contract.
func bindHoQuToken(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(HoQuTokenABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_HoQuToken *HoQuTokenRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _HoQuToken.Contract.HoQuTokenCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_HoQuToken *HoQuTokenRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _HoQuToken.Contract.HoQuTokenTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_HoQuToken *HoQuTokenRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _HoQuToken.Contract.HoQuTokenTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_HoQuToken *HoQuTokenCallerRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _HoQuToken.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_HoQuToken *HoQuTokenTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _HoQuToken.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_HoQuToken *HoQuTokenTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _HoQuToken.Contract.contract.Transact(opts, method, params...)
}

// Allowance is a free data retrieval call binding the contract method 0xdd62ed3e.
//
// Solidity: function allowance(_owner address, _spender address) constant returns(remaining uint256)
func (_HoQuToken *HoQuTokenCaller) Allowance(opts *bind.CallOpts, _owner common.Address, _spender common.Address) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _HoQuToken.contract.Call(opts, out, "allowance", _owner, _spender)
	return *ret0, err
}

// Allowance is a free data retrieval call binding the contract method 0xdd62ed3e.
//
// Solidity: function allowance(_owner address, _spender address) constant returns(remaining uint256)
func (_HoQuToken *HoQuTokenSession) Allowance(_owner common.Address, _spender common.Address) (*big.Int, error) {
	return _HoQuToken.Contract.Allowance(&_HoQuToken.CallOpts, _owner, _spender)
}

// Allowance is a free data retrieval call binding the contract method 0xdd62ed3e.
//
// Solidity: function allowance(_owner address, _spender address) constant returns(remaining uint256)
func (_HoQuToken *HoQuTokenCallerSession) Allowance(_owner common.Address, _spender common.Address) (*big.Int, error) {
	return _HoQuToken.Contract.Allowance(&_HoQuToken.CallOpts, _owner, _spender)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(_owner address) constant returns(balance uint256)
func (_HoQuToken *HoQuTokenCaller) BalanceOf(opts *bind.CallOpts, _owner common.Address) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _HoQuToken.contract.Call(opts, out, "balanceOf", _owner)
	return *ret0, err
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(_owner address) constant returns(balance uint256)
func (_HoQuToken *HoQuTokenSession) BalanceOf(_owner common.Address) (*big.Int, error) {
	return _HoQuToken.Contract.BalanceOf(&_HoQuToken.CallOpts, _owner)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(_owner address) constant returns(balance uint256)
func (_HoQuToken *HoQuTokenCallerSession) BalanceOf(_owner common.Address) (*big.Int, error) {
	return _HoQuToken.Contract.BalanceOf(&_HoQuToken.CallOpts, _owner)
}

// Decimals is a free data retrieval call binding the contract method 0x313ce567.
//
// Solidity: function decimals() constant returns(uint32)
func (_HoQuToken *HoQuTokenCaller) Decimals(opts *bind.CallOpts) (uint32, error) {
	var (
		ret0 = new(uint32)
	)
	out := ret0
	err := _HoQuToken.contract.Call(opts, out, "decimals")
	return *ret0, err
}

// Decimals is a free data retrieval call binding the contract method 0x313ce567.
//
// Solidity: function decimals() constant returns(uint32)
func (_HoQuToken *HoQuTokenSession) Decimals() (uint32, error) {
	return _HoQuToken.Contract.Decimals(&_HoQuToken.CallOpts)
}

// Decimals is a free data retrieval call binding the contract method 0x313ce567.
//
// Solidity: function decimals() constant returns(uint32)
func (_HoQuToken *HoQuTokenCallerSession) Decimals() (uint32, error) {
	return _HoQuToken.Contract.Decimals(&_HoQuToken.CallOpts)
}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() constant returns(string)
func (_HoQuToken *HoQuTokenCaller) Name(opts *bind.CallOpts) (string, error) {
	var (
		ret0 = new(string)
	)
	out := ret0
	err := _HoQuToken.contract.Call(opts, out, "name")
	return *ret0, err
}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() constant returns(string)
func (_HoQuToken *HoQuTokenSession) Name() (string, error) {
	return _HoQuToken.Contract.Name(&_HoQuToken.CallOpts)
}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() constant returns(string)
func (_HoQuToken *HoQuTokenCallerSession) Name() (string, error) {
	return _HoQuToken.Contract.Name(&_HoQuToken.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() constant returns(address)
func (_HoQuToken *HoQuTokenCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _HoQuToken.contract.Call(opts, out, "owner")
	return *ret0, err
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() constant returns(address)
func (_HoQuToken *HoQuTokenSession) Owner() (common.Address, error) {
	return _HoQuToken.Contract.Owner(&_HoQuToken.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() constant returns(address)
func (_HoQuToken *HoQuTokenCallerSession) Owner() (common.Address, error) {
	return _HoQuToken.Contract.Owner(&_HoQuToken.CallOpts)
}

// Paused is a free data retrieval call binding the contract method 0x5c975abb.
//
// Solidity: function paused() constant returns(bool)
func (_HoQuToken *HoQuTokenCaller) Paused(opts *bind.CallOpts) (bool, error) {
	var (
		ret0 = new(bool)
	)
	out := ret0
	err := _HoQuToken.contract.Call(opts, out, "paused")
	return *ret0, err
}

// Paused is a free data retrieval call binding the contract method 0x5c975abb.
//
// Solidity: function paused() constant returns(bool)
func (_HoQuToken *HoQuTokenSession) Paused() (bool, error) {
	return _HoQuToken.Contract.Paused(&_HoQuToken.CallOpts)
}

// Paused is a free data retrieval call binding the contract method 0x5c975abb.
//
// Solidity: function paused() constant returns(bool)
func (_HoQuToken *HoQuTokenCallerSession) Paused() (bool, error) {
	return _HoQuToken.Contract.Paused(&_HoQuToken.CallOpts)
}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() constant returns(string)
func (_HoQuToken *HoQuTokenCaller) Symbol(opts *bind.CallOpts) (string, error) {
	var (
		ret0 = new(string)
	)
	out := ret0
	err := _HoQuToken.contract.Call(opts, out, "symbol")
	return *ret0, err
}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() constant returns(string)
func (_HoQuToken *HoQuTokenSession) Symbol() (string, error) {
	return _HoQuToken.Contract.Symbol(&_HoQuToken.CallOpts)
}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() constant returns(string)
func (_HoQuToken *HoQuTokenCallerSession) Symbol() (string, error) {
	return _HoQuToken.Contract.Symbol(&_HoQuToken.CallOpts)
}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() constant returns(uint256)
func (_HoQuToken *HoQuTokenCaller) TotalSupply(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _HoQuToken.contract.Call(opts, out, "totalSupply")
	return *ret0, err
}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() constant returns(uint256)
func (_HoQuToken *HoQuTokenSession) TotalSupply() (*big.Int, error) {
	return _HoQuToken.Contract.TotalSupply(&_HoQuToken.CallOpts)
}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() constant returns(uint256)
func (_HoQuToken *HoQuTokenCallerSession) TotalSupply() (*big.Int, error) {
	return _HoQuToken.Contract.TotalSupply(&_HoQuToken.CallOpts)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(_spender address, _value uint256) returns(bool)
func (_HoQuToken *HoQuTokenTransactor) Approve(opts *bind.TransactOpts, _spender common.Address, _value *big.Int) (*types.Transaction, error) {
	return _HoQuToken.contract.Transact(opts, "approve", _spender, _value)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(_spender address, _value uint256) returns(bool)
func (_HoQuToken *HoQuTokenSession) Approve(_spender common.Address, _value *big.Int) (*types.Transaction, error) {
	return _HoQuToken.Contract.Approve(&_HoQuToken.TransactOpts, _spender, _value)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(_spender address, _value uint256) returns(bool)
func (_HoQuToken *HoQuTokenTransactorSession) Approve(_spender common.Address, _value *big.Int) (*types.Transaction, error) {
	return _HoQuToken.Contract.Approve(&_HoQuToken.TransactOpts, _spender, _value)
}

// DecreaseApproval is a paid mutator transaction binding the contract method 0x66188463.
//
// Solidity: function decreaseApproval(_spender address, _subtractedValue uint256) returns(success bool)
func (_HoQuToken *HoQuTokenTransactor) DecreaseApproval(opts *bind.TransactOpts, _spender common.Address, _subtractedValue *big.Int) (*types.Transaction, error) {
	return _HoQuToken.contract.Transact(opts, "decreaseApproval", _spender, _subtractedValue)
}

// DecreaseApproval is a paid mutator transaction binding the contract method 0x66188463.
//
// Solidity: function decreaseApproval(_spender address, _subtractedValue uint256) returns(success bool)
func (_HoQuToken *HoQuTokenSession) DecreaseApproval(_spender common.Address, _subtractedValue *big.Int) (*types.Transaction, error) {
	return _HoQuToken.Contract.DecreaseApproval(&_HoQuToken.TransactOpts, _spender, _subtractedValue)
}

// DecreaseApproval is a paid mutator transaction binding the contract method 0x66188463.
//
// Solidity: function decreaseApproval(_spender address, _subtractedValue uint256) returns(success bool)
func (_HoQuToken *HoQuTokenTransactorSession) DecreaseApproval(_spender common.Address, _subtractedValue *big.Int) (*types.Transaction, error) {
	return _HoQuToken.Contract.DecreaseApproval(&_HoQuToken.TransactOpts, _spender, _subtractedValue)
}

// IncreaseApproval is a paid mutator transaction binding the contract method 0xd73dd623.
//
// Solidity: function increaseApproval(_spender address, _addedValue uint256) returns(success bool)
func (_HoQuToken *HoQuTokenTransactor) IncreaseApproval(opts *bind.TransactOpts, _spender common.Address, _addedValue *big.Int) (*types.Transaction, error) {
	return _HoQuToken.contract.Transact(opts, "increaseApproval", _spender, _addedValue)
}

// IncreaseApproval is a paid mutator transaction binding the contract method 0xd73dd623.
//
// Solidity: function increaseApproval(_spender address, _addedValue uint256) returns(success bool)
func (_HoQuToken *HoQuTokenSession) IncreaseApproval(_spender common.Address, _addedValue *big.Int) (*types.Transaction, error) {
	return _HoQuToken.Contract.IncreaseApproval(&_HoQuToken.TransactOpts, _spender, _addedValue)
}

// IncreaseApproval is a paid mutator transaction binding the contract method 0xd73dd623.
//
// Solidity: function increaseApproval(_spender address, _addedValue uint256) returns(success bool)
func (_HoQuToken *HoQuTokenTransactorSession) IncreaseApproval(_spender common.Address, _addedValue *big.Int) (*types.Transaction, error) {
	return _HoQuToken.Contract.IncreaseApproval(&_HoQuToken.TransactOpts, _spender, _addedValue)
}

// Pause is a paid mutator transaction binding the contract method 0x8456cb59.
//
// Solidity: function pause() returns()
func (_HoQuToken *HoQuTokenTransactor) Pause(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _HoQuToken.contract.Transact(opts, "pause")
}

// Pause is a paid mutator transaction binding the contract method 0x8456cb59.
//
// Solidity: function pause() returns()
func (_HoQuToken *HoQuTokenSession) Pause() (*types.Transaction, error) {
	return _HoQuToken.Contract.Pause(&_HoQuToken.TransactOpts)
}

// Pause is a paid mutator transaction binding the contract method 0x8456cb59.
//
// Solidity: function pause() returns()
func (_HoQuToken *HoQuTokenTransactorSession) Pause() (*types.Transaction, error) {
	return _HoQuToken.Contract.Pause(&_HoQuToken.TransactOpts)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(_to address, _value uint256) returns(bool)
func (_HoQuToken *HoQuTokenTransactor) Transfer(opts *bind.TransactOpts, _to common.Address, _value *big.Int) (*types.Transaction, error) {
	return _HoQuToken.contract.Transact(opts, "transfer", _to, _value)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(_to address, _value uint256) returns(bool)
func (_HoQuToken *HoQuTokenSession) Transfer(_to common.Address, _value *big.Int) (*types.Transaction, error) {
	return _HoQuToken.Contract.Transfer(&_HoQuToken.TransactOpts, _to, _value)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(_to address, _value uint256) returns(bool)
func (_HoQuToken *HoQuTokenTransactorSession) Transfer(_to common.Address, _value *big.Int) (*types.Transaction, error) {
	return _HoQuToken.Contract.Transfer(&_HoQuToken.TransactOpts, _to, _value)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(_from address, _to address, _value uint256) returns(bool)
func (_HoQuToken *HoQuTokenTransactor) TransferFrom(opts *bind.TransactOpts, _from common.Address, _to common.Address, _value *big.Int) (*types.Transaction, error) {
	return _HoQuToken.contract.Transact(opts, "transferFrom", _from, _to, _value)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(_from address, _to address, _value uint256) returns(bool)
func (_HoQuToken *HoQuTokenSession) TransferFrom(_from common.Address, _to common.Address, _value *big.Int) (*types.Transaction, error) {
	return _HoQuToken.Contract.TransferFrom(&_HoQuToken.TransactOpts, _from, _to, _value)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(_from address, _to address, _value uint256) returns(bool)
func (_HoQuToken *HoQuTokenTransactorSession) TransferFrom(_from common.Address, _to common.Address, _value *big.Int) (*types.Transaction, error) {
	return _HoQuToken.Contract.TransferFrom(&_HoQuToken.TransactOpts, _from, _to, _value)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(newOwner address) returns()
func (_HoQuToken *HoQuTokenTransactor) TransferOwnership(opts *bind.TransactOpts, newOwner common.Address) (*types.Transaction, error) {
	return _HoQuToken.contract.Transact(opts, "transferOwnership", newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(newOwner address) returns()
func (_HoQuToken *HoQuTokenSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _HoQuToken.Contract.TransferOwnership(&_HoQuToken.TransactOpts, newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(newOwner address) returns()
func (_HoQuToken *HoQuTokenTransactorSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _HoQuToken.Contract.TransferOwnership(&_HoQuToken.TransactOpts, newOwner)
}

// Unpause is a paid mutator transaction binding the contract method 0x3f4ba83a.
//
// Solidity: function unpause() returns()
func (_HoQuToken *HoQuTokenTransactor) Unpause(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _HoQuToken.contract.Transact(opts, "unpause")
}

// Unpause is a paid mutator transaction binding the contract method 0x3f4ba83a.
//
// Solidity: function unpause() returns()
func (_HoQuToken *HoQuTokenSession) Unpause() (*types.Transaction, error) {
	return _HoQuToken.Contract.Unpause(&_HoQuToken.TransactOpts)
}

// Unpause is a paid mutator transaction binding the contract method 0x3f4ba83a.
//
// Solidity: function unpause() returns()
func (_HoQuToken *HoQuTokenTransactorSession) Unpause() (*types.Transaction, error) {
	return _HoQuToken.Contract.Unpause(&_HoQuToken.TransactOpts)
}

// OwnableABI is the input ABI used to generate the binding from.
const OwnableABI = "[{\"constant\":true,\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"transferOwnership\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"previousOwner\",\"type\":\"address\"},{\"indexed\":true,\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"OwnershipTransferred\",\"type\":\"event\"}]"

// OwnableBin is the compiled bytecode used for deploying new contracts.
const OwnableBin = `0x6060604052341561000f57600080fd5b60008054600160a060020a033316600160a060020a031990911617905561016d8061003b6000396000f300606060405263ffffffff7c01000000000000000000000000000000000000000000000000000000006000350416638da5cb5b8114610047578063f2fde38b1461007657600080fd5b341561005257600080fd5b61005a610097565b604051600160a060020a03909116815260200160405180910390f35b341561008157600080fd5b610095600160a060020a03600435166100a6565b005b600054600160a060020a031681565b60005433600160a060020a039081169116146100c157600080fd5b600160a060020a03811615156100d657600080fd5b600054600160a060020a0380831691167f8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e060405160405180910390a36000805473ffffffffffffffffffffffffffffffffffffffff1916600160a060020a03929092169190911790555600a165627a7a72305820327236c255905ffe9ca63652d3435145c0a86627e614727e6368572ac6e1dd7d0029`

// DeployOwnable deploys a new Ethereum contract, binding an instance of Ownable to it.
func DeployOwnable(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *Ownable, error) {
	parsed, err := abi.JSON(strings.NewReader(OwnableABI))
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	address, tx, contract, err := bind.DeployContract(auth, parsed, common.FromHex(OwnableBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &Ownable{OwnableCaller: OwnableCaller{contract: contract}, OwnableTransactor: OwnableTransactor{contract: contract}}, nil
}

// Ownable is an auto generated Go binding around an Ethereum contract.
type Ownable struct {
	OwnableCaller     // Read-only binding to the contract
	OwnableTransactor // Write-only binding to the contract
}

// OwnableCaller is an auto generated read-only Go binding around an Ethereum contract.
type OwnableCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// OwnableTransactor is an auto generated write-only Go binding around an Ethereum contract.
type OwnableTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// OwnableSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type OwnableSession struct {
	Contract     *Ownable          // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// OwnableCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type OwnableCallerSession struct {
	Contract *OwnableCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts  // Call options to use throughout this session
}

// OwnableTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type OwnableTransactorSession struct {
	Contract     *OwnableTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts  // Transaction auth options to use throughout this session
}

// OwnableRaw is an auto generated low-level Go binding around an Ethereum contract.
type OwnableRaw struct {
	Contract *Ownable // Generic contract binding to access the raw methods on
}

// OwnableCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type OwnableCallerRaw struct {
	Contract *OwnableCaller // Generic read-only contract binding to access the raw methods on
}

// OwnableTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type OwnableTransactorRaw struct {
	Contract *OwnableTransactor // Generic write-only contract binding to access the raw methods on
}

// NewOwnable creates a new instance of Ownable, bound to a specific deployed contract.
func NewOwnable(address common.Address, backend bind.ContractBackend) (*Ownable, error) {
	contract, err := bindOwnable(address, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Ownable{OwnableCaller: OwnableCaller{contract: contract}, OwnableTransactor: OwnableTransactor{contract: contract}}, nil
}

// NewOwnableCaller creates a new read-only instance of Ownable, bound to a specific deployed contract.
func NewOwnableCaller(address common.Address, caller bind.ContractCaller) (*OwnableCaller, error) {
	contract, err := bindOwnable(address, caller, nil)
	if err != nil {
		return nil, err
	}
	return &OwnableCaller{contract: contract}, nil
}

// NewOwnableTransactor creates a new write-only instance of Ownable, bound to a specific deployed contract.
func NewOwnableTransactor(address common.Address, transactor bind.ContractTransactor) (*OwnableTransactor, error) {
	contract, err := bindOwnable(address, nil, transactor)
	if err != nil {
		return nil, err
	}
	return &OwnableTransactor{contract: contract}, nil
}

// bindOwnable binds a generic wrapper to an already deployed contract.
func bindOwnable(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(OwnableABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Ownable *OwnableRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _Ownable.Contract.OwnableCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Ownable *OwnableRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Ownable.Contract.OwnableTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Ownable *OwnableRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Ownable.Contract.OwnableTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Ownable *OwnableCallerRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _Ownable.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Ownable *OwnableTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Ownable.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Ownable *OwnableTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Ownable.Contract.contract.Transact(opts, method, params...)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() constant returns(address)
func (_Ownable *OwnableCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _Ownable.contract.Call(opts, out, "owner")
	return *ret0, err
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() constant returns(address)
func (_Ownable *OwnableSession) Owner() (common.Address, error) {
	return _Ownable.Contract.Owner(&_Ownable.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() constant returns(address)
func (_Ownable *OwnableCallerSession) Owner() (common.Address, error) {
	return _Ownable.Contract.Owner(&_Ownable.CallOpts)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(newOwner address) returns()
func (_Ownable *OwnableTransactor) TransferOwnership(opts *bind.TransactOpts, newOwner common.Address) (*types.Transaction, error) {
	return _Ownable.contract.Transact(opts, "transferOwnership", newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(newOwner address) returns()
func (_Ownable *OwnableSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _Ownable.Contract.TransferOwnership(&_Ownable.TransactOpts, newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(newOwner address) returns()
func (_Ownable *OwnableTransactorSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _Ownable.Contract.TransferOwnership(&_Ownable.TransactOpts, newOwner)
}

// PausableABI is the input ABI used to generate the binding from.
const PausableABI = "[{\"constant\":false,\"inputs\":[],\"name\":\"unpause\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"paused\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[],\"name\":\"pause\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"transferOwnership\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"anonymous\":false,\"inputs\":[],\"name\":\"Pause\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[],\"name\":\"Unpause\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"previousOwner\",\"type\":\"address\"},{\"indexed\":true,\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"OwnershipTransferred\",\"type\":\"event\"}]"

// PausableBin is the compiled bytecode used for deploying new contracts.
const PausableBin = `0x606060405260008054600160a060020a033316600160a860020a03199091161790556102ee806100306000396000f300606060405263ffffffff7c01000000000000000000000000000000000000000000000000000000006000350416633f4ba83a81146100685780635c975abb1461007d5780638456cb59146100a45780638da5cb5b146100b7578063f2fde38b146100e657600080fd5b341561007357600080fd5b61007b610105565b005b341561008857600080fd5b610090610184565b604051901515815260200160405180910390f35b34156100af57600080fd5b61007b610194565b34156100c257600080fd5b6100ca610218565b604051600160a060020a03909116815260200160405180910390f35b34156100f157600080fd5b61007b600160a060020a0360043516610227565b60005433600160a060020a0390811691161461012057600080fd5b60005460a060020a900460ff16151561013857600080fd5b6000805474ff0000000000000000000000000000000000000000191690557f7805862f689e2f13df9f062ff482ad3ad112aca9e0847911ed832e158c525b3360405160405180910390a1565b60005460a060020a900460ff1681565b60005433600160a060020a039081169116146101af57600080fd5b60005460a060020a900460ff16156101c657600080fd5b6000805474ff0000000000000000000000000000000000000000191660a060020a1790557f6985a02210a168e66602d3235cb6db0e70f92b3ba4d376a33c0f3d9434bff62560405160405180910390a1565b600054600160a060020a031681565b60005433600160a060020a0390811691161461024257600080fd5b600160a060020a038116151561025757600080fd5b600054600160a060020a0380831691167f8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e060405160405180910390a36000805473ffffffffffffffffffffffffffffffffffffffff1916600160a060020a03929092169190911790555600a165627a7a7230582041f71d89f9ba4a1736516c0bb17e9394b7de355486611bce649d37446c4ceb280029`

// DeployPausable deploys a new Ethereum contract, binding an instance of Pausable to it.
func DeployPausable(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *Pausable, error) {
	parsed, err := abi.JSON(strings.NewReader(PausableABI))
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	address, tx, contract, err := bind.DeployContract(auth, parsed, common.FromHex(PausableBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &Pausable{PausableCaller: PausableCaller{contract: contract}, PausableTransactor: PausableTransactor{contract: contract}}, nil
}

// Pausable is an auto generated Go binding around an Ethereum contract.
type Pausable struct {
	PausableCaller     // Read-only binding to the contract
	PausableTransactor // Write-only binding to the contract
}

// PausableCaller is an auto generated read-only Go binding around an Ethereum contract.
type PausableCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// PausableTransactor is an auto generated write-only Go binding around an Ethereum contract.
type PausableTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// PausableSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type PausableSession struct {
	Contract     *Pausable         // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// PausableCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type PausableCallerSession struct {
	Contract *PausableCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts   // Call options to use throughout this session
}

// PausableTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type PausableTransactorSession struct {
	Contract     *PausableTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts   // Transaction auth options to use throughout this session
}

// PausableRaw is an auto generated low-level Go binding around an Ethereum contract.
type PausableRaw struct {
	Contract *Pausable // Generic contract binding to access the raw methods on
}

// PausableCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type PausableCallerRaw struct {
	Contract *PausableCaller // Generic read-only contract binding to access the raw methods on
}

// PausableTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type PausableTransactorRaw struct {
	Contract *PausableTransactor // Generic write-only contract binding to access the raw methods on
}

// NewPausable creates a new instance of Pausable, bound to a specific deployed contract.
func NewPausable(address common.Address, backend bind.ContractBackend) (*Pausable, error) {
	contract, err := bindPausable(address, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Pausable{PausableCaller: PausableCaller{contract: contract}, PausableTransactor: PausableTransactor{contract: contract}}, nil
}

// NewPausableCaller creates a new read-only instance of Pausable, bound to a specific deployed contract.
func NewPausableCaller(address common.Address, caller bind.ContractCaller) (*PausableCaller, error) {
	contract, err := bindPausable(address, caller, nil)
	if err != nil {
		return nil, err
	}
	return &PausableCaller{contract: contract}, nil
}

// NewPausableTransactor creates a new write-only instance of Pausable, bound to a specific deployed contract.
func NewPausableTransactor(address common.Address, transactor bind.ContractTransactor) (*PausableTransactor, error) {
	contract, err := bindPausable(address, nil, transactor)
	if err != nil {
		return nil, err
	}
	return &PausableTransactor{contract: contract}, nil
}

// bindPausable binds a generic wrapper to an already deployed contract.
func bindPausable(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(PausableABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Pausable *PausableRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _Pausable.Contract.PausableCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Pausable *PausableRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Pausable.Contract.PausableTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Pausable *PausableRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Pausable.Contract.PausableTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Pausable *PausableCallerRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _Pausable.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Pausable *PausableTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Pausable.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Pausable *PausableTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Pausable.Contract.contract.Transact(opts, method, params...)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() constant returns(address)
func (_Pausable *PausableCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _Pausable.contract.Call(opts, out, "owner")
	return *ret0, err
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() constant returns(address)
func (_Pausable *PausableSession) Owner() (common.Address, error) {
	return _Pausable.Contract.Owner(&_Pausable.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() constant returns(address)
func (_Pausable *PausableCallerSession) Owner() (common.Address, error) {
	return _Pausable.Contract.Owner(&_Pausable.CallOpts)
}

// Paused is a free data retrieval call binding the contract method 0x5c975abb.
//
// Solidity: function paused() constant returns(bool)
func (_Pausable *PausableCaller) Paused(opts *bind.CallOpts) (bool, error) {
	var (
		ret0 = new(bool)
	)
	out := ret0
	err := _Pausable.contract.Call(opts, out, "paused")
	return *ret0, err
}

// Paused is a free data retrieval call binding the contract method 0x5c975abb.
//
// Solidity: function paused() constant returns(bool)
func (_Pausable *PausableSession) Paused() (bool, error) {
	return _Pausable.Contract.Paused(&_Pausable.CallOpts)
}

// Paused is a free data retrieval call binding the contract method 0x5c975abb.
//
// Solidity: function paused() constant returns(bool)
func (_Pausable *PausableCallerSession) Paused() (bool, error) {
	return _Pausable.Contract.Paused(&_Pausable.CallOpts)
}

// Pause is a paid mutator transaction binding the contract method 0x8456cb59.
//
// Solidity: function pause() returns()
func (_Pausable *PausableTransactor) Pause(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Pausable.contract.Transact(opts, "pause")
}

// Pause is a paid mutator transaction binding the contract method 0x8456cb59.
//
// Solidity: function pause() returns()
func (_Pausable *PausableSession) Pause() (*types.Transaction, error) {
	return _Pausable.Contract.Pause(&_Pausable.TransactOpts)
}

// Pause is a paid mutator transaction binding the contract method 0x8456cb59.
//
// Solidity: function pause() returns()
func (_Pausable *PausableTransactorSession) Pause() (*types.Transaction, error) {
	return _Pausable.Contract.Pause(&_Pausable.TransactOpts)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(newOwner address) returns()
func (_Pausable *PausableTransactor) TransferOwnership(opts *bind.TransactOpts, newOwner common.Address) (*types.Transaction, error) {
	return _Pausable.contract.Transact(opts, "transferOwnership", newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(newOwner address) returns()
func (_Pausable *PausableSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _Pausable.Contract.TransferOwnership(&_Pausable.TransactOpts, newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(newOwner address) returns()
func (_Pausable *PausableTransactorSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _Pausable.Contract.TransferOwnership(&_Pausable.TransactOpts, newOwner)
}

// Unpause is a paid mutator transaction binding the contract method 0x3f4ba83a.
//
// Solidity: function unpause() returns()
func (_Pausable *PausableTransactor) Unpause(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Pausable.contract.Transact(opts, "unpause")
}

// Unpause is a paid mutator transaction binding the contract method 0x3f4ba83a.
//
// Solidity: function unpause() returns()
func (_Pausable *PausableSession) Unpause() (*types.Transaction, error) {
	return _Pausable.Contract.Unpause(&_Pausable.TransactOpts)
}

// Unpause is a paid mutator transaction binding the contract method 0x3f4ba83a.
//
// Solidity: function unpause() returns()
func (_Pausable *PausableTransactorSession) Unpause() (*types.Transaction, error) {
	return _Pausable.Contract.Unpause(&_Pausable.TransactOpts)
}

// SafeMathABI is the input ABI used to generate the binding from.
const SafeMathABI = "[]"

// SafeMathBin is the compiled bytecode used for deploying new contracts.
const SafeMathBin = `0x60606040523415600e57600080fd5b603580601b6000396000f3006060604052600080fd00a165627a7a7230582022bb8387f0043980ba9fa96c3c9ae61ec2dd881153a38510d3f4666d08e078520029`

// DeploySafeMath deploys a new Ethereum contract, binding an instance of SafeMath to it.
func DeploySafeMath(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *SafeMath, error) {
	parsed, err := abi.JSON(strings.NewReader(SafeMathABI))
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	address, tx, contract, err := bind.DeployContract(auth, parsed, common.FromHex(SafeMathBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &SafeMath{SafeMathCaller: SafeMathCaller{contract: contract}, SafeMathTransactor: SafeMathTransactor{contract: contract}}, nil
}

// SafeMath is an auto generated Go binding around an Ethereum contract.
type SafeMath struct {
	SafeMathCaller     // Read-only binding to the contract
	SafeMathTransactor // Write-only binding to the contract
}

// SafeMathCaller is an auto generated read-only Go binding around an Ethereum contract.
type SafeMathCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// SafeMathTransactor is an auto generated write-only Go binding around an Ethereum contract.
type SafeMathTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// SafeMathSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type SafeMathSession struct {
	Contract     *SafeMath         // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// SafeMathCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type SafeMathCallerSession struct {
	Contract *SafeMathCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts   // Call options to use throughout this session
}

// SafeMathTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type SafeMathTransactorSession struct {
	Contract     *SafeMathTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts   // Transaction auth options to use throughout this session
}

// SafeMathRaw is an auto generated low-level Go binding around an Ethereum contract.
type SafeMathRaw struct {
	Contract *SafeMath // Generic contract binding to access the raw methods on
}

// SafeMathCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type SafeMathCallerRaw struct {
	Contract *SafeMathCaller // Generic read-only contract binding to access the raw methods on
}

// SafeMathTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type SafeMathTransactorRaw struct {
	Contract *SafeMathTransactor // Generic write-only contract binding to access the raw methods on
}

// NewSafeMath creates a new instance of SafeMath, bound to a specific deployed contract.
func NewSafeMath(address common.Address, backend bind.ContractBackend) (*SafeMath, error) {
	contract, err := bindSafeMath(address, backend, backend)
	if err != nil {
		return nil, err
	}
	return &SafeMath{SafeMathCaller: SafeMathCaller{contract: contract}, SafeMathTransactor: SafeMathTransactor{contract: contract}}, nil
}

// NewSafeMathCaller creates a new read-only instance of SafeMath, bound to a specific deployed contract.
func NewSafeMathCaller(address common.Address, caller bind.ContractCaller) (*SafeMathCaller, error) {
	contract, err := bindSafeMath(address, caller, nil)
	if err != nil {
		return nil, err
	}
	return &SafeMathCaller{contract: contract}, nil
}

// NewSafeMathTransactor creates a new write-only instance of SafeMath, bound to a specific deployed contract.
func NewSafeMathTransactor(address common.Address, transactor bind.ContractTransactor) (*SafeMathTransactor, error) {
	contract, err := bindSafeMath(address, nil, transactor)
	if err != nil {
		return nil, err
	}
	return &SafeMathTransactor{contract: contract}, nil
}

// bindSafeMath binds a generic wrapper to an already deployed contract.
func bindSafeMath(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(SafeMathABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_SafeMath *SafeMathRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _SafeMath.Contract.SafeMathCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_SafeMath *SafeMathRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _SafeMath.Contract.SafeMathTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_SafeMath *SafeMathRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _SafeMath.Contract.SafeMathTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_SafeMath *SafeMathCallerRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _SafeMath.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_SafeMath *SafeMathTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _SafeMath.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_SafeMath *SafeMathTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _SafeMath.Contract.contract.Transact(opts, method, params...)
}

// StandardTokenABI is the input ABI used to generate the binding from.
const StandardTokenABI = "[{\"constant\":false,\"inputs\":[{\"name\":\"_spender\",\"type\":\"address\"},{\"name\":\"_value\",\"type\":\"uint256\"}],\"name\":\"approve\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"totalSupply\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_from\",\"type\":\"address\"},{\"name\":\"_to\",\"type\":\"address\"},{\"name\":\"_value\",\"type\":\"uint256\"}],\"name\":\"transferFrom\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_spender\",\"type\":\"address\"},{\"name\":\"_subtractedValue\",\"type\":\"uint256\"}],\"name\":\"decreaseApproval\",\"outputs\":[{\"name\":\"success\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"_owner\",\"type\":\"address\"}],\"name\":\"balanceOf\",\"outputs\":[{\"name\":\"balance\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_to\",\"type\":\"address\"},{\"name\":\"_value\",\"type\":\"uint256\"}],\"name\":\"transfer\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_spender\",\"type\":\"address\"},{\"name\":\"_addedValue\",\"type\":\"uint256\"}],\"name\":\"increaseApproval\",\"outputs\":[{\"name\":\"success\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"_owner\",\"type\":\"address\"},{\"name\":\"_spender\",\"type\":\"address\"}],\"name\":\"allowance\",\"outputs\":[{\"name\":\"remaining\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"owner\",\"type\":\"address\"},{\"indexed\":true,\"name\":\"spender\",\"type\":\"address\"},{\"indexed\":false,\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"Approval\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"from\",\"type\":\"address\"},{\"indexed\":true,\"name\":\"to\",\"type\":\"address\"},{\"indexed\":false,\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"Transfer\",\"type\":\"event\"}]"

// StandardTokenBin is the compiled bytecode used for deploying new contracts.
const StandardTokenBin = `0x6060604052341561000f57600080fd5b6106e48061001e6000396000f3006060604052361561008b5763ffffffff7c0100000000000000000000000000000000000000000000000000000000600035041663095ea7b3811461009057806318160ddd146100c657806323b872dd146100eb578063661884631461011357806370a0823114610135578063a9059cbb14610154578063d73dd62314610176578063dd62ed3e14610198575b600080fd5b341561009b57600080fd5b6100b2600160a060020a03600435166024356101bd565b604051901515815260200160405180910390f35b34156100d157600080fd5b6100d9610229565b60405190815260200160405180910390f35b34156100f657600080fd5b6100b2600160a060020a036004358116906024351660443561022f565b341561011e57600080fd5b6100b2600160a060020a03600435166024356103b1565b341561014057600080fd5b6100d9600160a060020a03600435166104ab565b341561015f57600080fd5b6100b2600160a060020a03600435166024356104c6565b341561018157600080fd5b6100b2600160a060020a03600435166024356105c1565b34156101a357600080fd5b6100d9600160a060020a0360043581169060243516610665565b600160a060020a03338116600081815260026020908152604080832094871680845294909152808220859055909291907f8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b9259085905190815260200160405180910390a350600192915050565b60005481565b6000600160a060020a038316151561024657600080fd5b600160a060020a03841660009081526001602052604090205482111561026b57600080fd5b600160a060020a038085166000908152600260209081526040808320339094168352929052205482111561029e57600080fd5b600160a060020a0384166000908152600160205260409020546102c7908363ffffffff61069016565b600160a060020a0380861660009081526001602052604080822093909355908516815220546102fc908363ffffffff6106a216565b600160a060020a03808516600090815260016020908152604080832094909455878316825260028152838220339093168252919091522054610344908363ffffffff61069016565b600160a060020a03808616600081815260026020908152604080832033861684529091529081902093909355908516917fddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef9085905190815260200160405180910390a35060019392505050565b600160a060020a0333811660009081526002602090815260408083209386168352929052908120548083111561040e57600160a060020a033381166000908152600260209081526040808320938816835292905290812055610445565b61041e818463ffffffff61069016565b600160a060020a033381166000908152600260209081526040808320938916835292905220555b600160a060020a0333811660008181526002602090815260408083209489168084529490915290819020547f8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925915190815260200160405180910390a35060019392505050565b600160a060020a031660009081526001602052604090205490565b6000600160a060020a03831615156104dd57600080fd5b600160a060020a03331660009081526001602052604090205482111561050257600080fd5b600160a060020a03331660009081526001602052604090205461052b908363ffffffff61069016565b600160a060020a033381166000908152600160205260408082209390935590851681522054610560908363ffffffff6106a216565b600160a060020a0380851660008181526001602052604090819020939093559133909116907fddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef9085905190815260200160405180910390a350600192915050565b600160a060020a0333811660009081526002602090815260408083209386168352929052908120546105f9908363ffffffff6106a216565b600160a060020a0333811660008181526002602090815260408083209489168084529490915290819020849055919290917f8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b92591905190815260200160405180910390a350600192915050565b600160a060020a03918216600090815260026020908152604080832093909416825291909152205490565b60008282111561069c57fe5b50900390565b6000828201838110156106b157fe5b93925050505600a165627a7a72305820e16dce13437f8a1170cc892471c6777be0b39136204307cf691642bdc7558bb50029`

// DeployStandardToken deploys a new Ethereum contract, binding an instance of StandardToken to it.
func DeployStandardToken(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *StandardToken, error) {
	parsed, err := abi.JSON(strings.NewReader(StandardTokenABI))
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	address, tx, contract, err := bind.DeployContract(auth, parsed, common.FromHex(StandardTokenBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &StandardToken{StandardTokenCaller: StandardTokenCaller{contract: contract}, StandardTokenTransactor: StandardTokenTransactor{contract: contract}}, nil
}

// StandardToken is an auto generated Go binding around an Ethereum contract.
type StandardToken struct {
	StandardTokenCaller     // Read-only binding to the contract
	StandardTokenTransactor // Write-only binding to the contract
}

// StandardTokenCaller is an auto generated read-only Go binding around an Ethereum contract.
type StandardTokenCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// StandardTokenTransactor is an auto generated write-only Go binding around an Ethereum contract.
type StandardTokenTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// StandardTokenSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type StandardTokenSession struct {
	Contract     *StandardToken    // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// StandardTokenCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type StandardTokenCallerSession struct {
	Contract *StandardTokenCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts        // Call options to use throughout this session
}

// StandardTokenTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type StandardTokenTransactorSession struct {
	Contract     *StandardTokenTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts        // Transaction auth options to use throughout this session
}

// StandardTokenRaw is an auto generated low-level Go binding around an Ethereum contract.
type StandardTokenRaw struct {
	Contract *StandardToken // Generic contract binding to access the raw methods on
}

// StandardTokenCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type StandardTokenCallerRaw struct {
	Contract *StandardTokenCaller // Generic read-only contract binding to access the raw methods on
}

// StandardTokenTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type StandardTokenTransactorRaw struct {
	Contract *StandardTokenTransactor // Generic write-only contract binding to access the raw methods on
}

// NewStandardToken creates a new instance of StandardToken, bound to a specific deployed contract.
func NewStandardToken(address common.Address, backend bind.ContractBackend) (*StandardToken, error) {
	contract, err := bindStandardToken(address, backend, backend)
	if err != nil {
		return nil, err
	}
	return &StandardToken{StandardTokenCaller: StandardTokenCaller{contract: contract}, StandardTokenTransactor: StandardTokenTransactor{contract: contract}}, nil
}

// NewStandardTokenCaller creates a new read-only instance of StandardToken, bound to a specific deployed contract.
func NewStandardTokenCaller(address common.Address, caller bind.ContractCaller) (*StandardTokenCaller, error) {
	contract, err := bindStandardToken(address, caller, nil)
	if err != nil {
		return nil, err
	}
	return &StandardTokenCaller{contract: contract}, nil
}

// NewStandardTokenTransactor creates a new write-only instance of StandardToken, bound to a specific deployed contract.
func NewStandardTokenTransactor(address common.Address, transactor bind.ContractTransactor) (*StandardTokenTransactor, error) {
	contract, err := bindStandardToken(address, nil, transactor)
	if err != nil {
		return nil, err
	}
	return &StandardTokenTransactor{contract: contract}, nil
}

// bindStandardToken binds a generic wrapper to an already deployed contract.
func bindStandardToken(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(StandardTokenABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_StandardToken *StandardTokenRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _StandardToken.Contract.StandardTokenCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_StandardToken *StandardTokenRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _StandardToken.Contract.StandardTokenTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_StandardToken *StandardTokenRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _StandardToken.Contract.StandardTokenTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_StandardToken *StandardTokenCallerRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _StandardToken.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_StandardToken *StandardTokenTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _StandardToken.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_StandardToken *StandardTokenTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _StandardToken.Contract.contract.Transact(opts, method, params...)
}

// Allowance is a free data retrieval call binding the contract method 0xdd62ed3e.
//
// Solidity: function allowance(_owner address, _spender address) constant returns(remaining uint256)
func (_StandardToken *StandardTokenCaller) Allowance(opts *bind.CallOpts, _owner common.Address, _spender common.Address) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _StandardToken.contract.Call(opts, out, "allowance", _owner, _spender)
	return *ret0, err
}

// Allowance is a free data retrieval call binding the contract method 0xdd62ed3e.
//
// Solidity: function allowance(_owner address, _spender address) constant returns(remaining uint256)
func (_StandardToken *StandardTokenSession) Allowance(_owner common.Address, _spender common.Address) (*big.Int, error) {
	return _StandardToken.Contract.Allowance(&_StandardToken.CallOpts, _owner, _spender)
}

// Allowance is a free data retrieval call binding the contract method 0xdd62ed3e.
//
// Solidity: function allowance(_owner address, _spender address) constant returns(remaining uint256)
func (_StandardToken *StandardTokenCallerSession) Allowance(_owner common.Address, _spender common.Address) (*big.Int, error) {
	return _StandardToken.Contract.Allowance(&_StandardToken.CallOpts, _owner, _spender)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(_owner address) constant returns(balance uint256)
func (_StandardToken *StandardTokenCaller) BalanceOf(opts *bind.CallOpts, _owner common.Address) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _StandardToken.contract.Call(opts, out, "balanceOf", _owner)
	return *ret0, err
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(_owner address) constant returns(balance uint256)
func (_StandardToken *StandardTokenSession) BalanceOf(_owner common.Address) (*big.Int, error) {
	return _StandardToken.Contract.BalanceOf(&_StandardToken.CallOpts, _owner)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(_owner address) constant returns(balance uint256)
func (_StandardToken *StandardTokenCallerSession) BalanceOf(_owner common.Address) (*big.Int, error) {
	return _StandardToken.Contract.BalanceOf(&_StandardToken.CallOpts, _owner)
}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() constant returns(uint256)
func (_StandardToken *StandardTokenCaller) TotalSupply(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _StandardToken.contract.Call(opts, out, "totalSupply")
	return *ret0, err
}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() constant returns(uint256)
func (_StandardToken *StandardTokenSession) TotalSupply() (*big.Int, error) {
	return _StandardToken.Contract.TotalSupply(&_StandardToken.CallOpts)
}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() constant returns(uint256)
func (_StandardToken *StandardTokenCallerSession) TotalSupply() (*big.Int, error) {
	return _StandardToken.Contract.TotalSupply(&_StandardToken.CallOpts)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(_spender address, _value uint256) returns(bool)
func (_StandardToken *StandardTokenTransactor) Approve(opts *bind.TransactOpts, _spender common.Address, _value *big.Int) (*types.Transaction, error) {
	return _StandardToken.contract.Transact(opts, "approve", _spender, _value)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(_spender address, _value uint256) returns(bool)
func (_StandardToken *StandardTokenSession) Approve(_spender common.Address, _value *big.Int) (*types.Transaction, error) {
	return _StandardToken.Contract.Approve(&_StandardToken.TransactOpts, _spender, _value)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(_spender address, _value uint256) returns(bool)
func (_StandardToken *StandardTokenTransactorSession) Approve(_spender common.Address, _value *big.Int) (*types.Transaction, error) {
	return _StandardToken.Contract.Approve(&_StandardToken.TransactOpts, _spender, _value)
}

// DecreaseApproval is a paid mutator transaction binding the contract method 0x66188463.
//
// Solidity: function decreaseApproval(_spender address, _subtractedValue uint256) returns(success bool)
func (_StandardToken *StandardTokenTransactor) DecreaseApproval(opts *bind.TransactOpts, _spender common.Address, _subtractedValue *big.Int) (*types.Transaction, error) {
	return _StandardToken.contract.Transact(opts, "decreaseApproval", _spender, _subtractedValue)
}

// DecreaseApproval is a paid mutator transaction binding the contract method 0x66188463.
//
// Solidity: function decreaseApproval(_spender address, _subtractedValue uint256) returns(success bool)
func (_StandardToken *StandardTokenSession) DecreaseApproval(_spender common.Address, _subtractedValue *big.Int) (*types.Transaction, error) {
	return _StandardToken.Contract.DecreaseApproval(&_StandardToken.TransactOpts, _spender, _subtractedValue)
}

// DecreaseApproval is a paid mutator transaction binding the contract method 0x66188463.
//
// Solidity: function decreaseApproval(_spender address, _subtractedValue uint256) returns(success bool)
func (_StandardToken *StandardTokenTransactorSession) DecreaseApproval(_spender common.Address, _subtractedValue *big.Int) (*types.Transaction, error) {
	return _StandardToken.Contract.DecreaseApproval(&_StandardToken.TransactOpts, _spender, _subtractedValue)
}

// IncreaseApproval is a paid mutator transaction binding the contract method 0xd73dd623.
//
// Solidity: function increaseApproval(_spender address, _addedValue uint256) returns(success bool)
func (_StandardToken *StandardTokenTransactor) IncreaseApproval(opts *bind.TransactOpts, _spender common.Address, _addedValue *big.Int) (*types.Transaction, error) {
	return _StandardToken.contract.Transact(opts, "increaseApproval", _spender, _addedValue)
}

// IncreaseApproval is a paid mutator transaction binding the contract method 0xd73dd623.
//
// Solidity: function increaseApproval(_spender address, _addedValue uint256) returns(success bool)
func (_StandardToken *StandardTokenSession) IncreaseApproval(_spender common.Address, _addedValue *big.Int) (*types.Transaction, error) {
	return _StandardToken.Contract.IncreaseApproval(&_StandardToken.TransactOpts, _spender, _addedValue)
}

// IncreaseApproval is a paid mutator transaction binding the contract method 0xd73dd623.
//
// Solidity: function increaseApproval(_spender address, _addedValue uint256) returns(success bool)
func (_StandardToken *StandardTokenTransactorSession) IncreaseApproval(_spender common.Address, _addedValue *big.Int) (*types.Transaction, error) {
	return _StandardToken.Contract.IncreaseApproval(&_StandardToken.TransactOpts, _spender, _addedValue)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(_to address, _value uint256) returns(bool)
func (_StandardToken *StandardTokenTransactor) Transfer(opts *bind.TransactOpts, _to common.Address, _value *big.Int) (*types.Transaction, error) {
	return _StandardToken.contract.Transact(opts, "transfer", _to, _value)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(_to address, _value uint256) returns(bool)
func (_StandardToken *StandardTokenSession) Transfer(_to common.Address, _value *big.Int) (*types.Transaction, error) {
	return _StandardToken.Contract.Transfer(&_StandardToken.TransactOpts, _to, _value)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(_to address, _value uint256) returns(bool)
func (_StandardToken *StandardTokenTransactorSession) Transfer(_to common.Address, _value *big.Int) (*types.Transaction, error) {
	return _StandardToken.Contract.Transfer(&_StandardToken.TransactOpts, _to, _value)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(_from address, _to address, _value uint256) returns(bool)
func (_StandardToken *StandardTokenTransactor) TransferFrom(opts *bind.TransactOpts, _from common.Address, _to common.Address, _value *big.Int) (*types.Transaction, error) {
	return _StandardToken.contract.Transact(opts, "transferFrom", _from, _to, _value)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(_from address, _to address, _value uint256) returns(bool)
func (_StandardToken *StandardTokenSession) TransferFrom(_from common.Address, _to common.Address, _value *big.Int) (*types.Transaction, error) {
	return _StandardToken.Contract.TransferFrom(&_StandardToken.TransactOpts, _from, _to, _value)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(_from address, _to address, _value uint256) returns(bool)
func (_StandardToken *StandardTokenTransactorSession) TransferFrom(_from common.Address, _to common.Address, _value *big.Int) (*types.Transaction, error) {
	return _StandardToken.Contract.TransferFrom(&_StandardToken.TransactOpts, _from, _to, _value)
}
