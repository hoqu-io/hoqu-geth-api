// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package platform

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
const BasicTokenBin = `0x6060604052341561000f57600080fd5b6102458061001e6000396000f3006060604052600436106100565763ffffffff7c010000000000000000000000000000000000000000000000000000000060003504166318160ddd811461005b57806370a0823114610080578063a9059cbb1461009f575b600080fd5b341561006657600080fd5b61006e6100d5565b60405190815260200160405180910390f35b341561008b57600080fd5b61006e600160a060020a03600435166100db565b34156100aa57600080fd5b6100c1600160a060020a03600435166024356100f6565b604051901515815260200160405180910390f35b60005481565b600160a060020a031660009081526001602052604090205490565b6000600160a060020a038316151561010d57600080fd5b600160a060020a03331660009081526001602052604090205482111561013257600080fd5b600160a060020a03331660009081526001602052604090205461015b908363ffffffff6101f116565b600160a060020a033381166000908152600160205260408082209390935590851681522054610190908363ffffffff61020316565b600160a060020a0380851660008181526001602052604090819020939093559133909116907fddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef9085905190815260200160405180910390a350600192915050565b6000828211156101fd57fe5b50900390565b60008282018381101561021257fe5b93925050505600a165627a7a72305820c3db7338c47517d440a59a0da6844db3c54ea7763e2b7cd9377548b06904bf890029`

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

// HoQuPlatformABI is the input ABI used to generate the binding from.
const HoQuPlatformABI = "[{\"constant\":true,\"inputs\":[{\"name\":\"id\",\"type\":\"uint32\"},{\"name\":\"num\",\"type\":\"uint8\"}],\"name\":\"getLeadIntermediaryAddress\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"id\",\"type\":\"uint32\"},{\"name\":\"num\",\"type\":\"uint16\"}],\"name\":\"getUserKycReport\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"},{\"name\":\"\",\"type\":\"uint32\"},{\"name\":\"\",\"type\":\"uint8\"},{\"name\":\"\",\"type\":\"string\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"id\",\"type\":\"uint32\"},{\"name\":\"role\",\"type\":\"string\"},{\"name\":\"ownerAddress\",\"type\":\"address\"},{\"name\":\"pubKey\",\"type\":\"string\"}],\"name\":\"registerUser\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"id\",\"type\":\"uint32\"}],\"name\":\"sellLead\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"id\",\"type\":\"uint32\"},{\"name\":\"num\",\"type\":\"uint8\"}],\"name\":\"getUserAddress\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"id\",\"type\":\"uint32\"},{\"name\":\"pubKey\",\"type\":\"string\"}],\"name\":\"updateUserPubKey\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"\",\"type\":\"uint32\"}],\"name\":\"companies\",\"outputs\":[{\"name\":\"createdAt\",\"type\":\"uint256\"},{\"name\":\"ownerId\",\"type\":\"uint32\"},{\"name\":\"ownerAddress\",\"type\":\"address\"},{\"name\":\"name\",\"type\":\"string\"},{\"name\":\"dataUrl\",\"type\":\"string\"},{\"name\":\"status\",\"type\":\"uint8\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"\",\"type\":\"uint32\"}],\"name\":\"leads\",\"outputs\":[{\"name\":\"createdAt\",\"type\":\"uint256\"},{\"name\":\"trackerId\",\"type\":\"uint32\"},{\"name\":\"ownerId\",\"type\":\"uint32\"},{\"name\":\"beneficiaryAddress\",\"type\":\"address\"},{\"name\":\"offerId\",\"type\":\"uint32\"},{\"name\":\"dataUrl\",\"type\":\"string\"},{\"name\":\"meta\",\"type\":\"string\"},{\"name\":\"price\",\"type\":\"uint256\"},{\"name\":\"numOfIntermediaries\",\"type\":\"uint8\"},{\"name\":\"status\",\"type\":\"uint8\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"id\",\"type\":\"uint32\"},{\"name\":\"status\",\"type\":\"uint8\"}],\"name\":\"setOfferStatus\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"id\",\"type\":\"uint32\"},{\"name\":\"ownerId\",\"type\":\"uint32\"},{\"name\":\"trackerId\",\"type\":\"uint32\"},{\"name\":\"offerId\",\"type\":\"uint32\"},{\"name\":\"beneficiaryAddress\",\"type\":\"address\"},{\"name\":\"meta\",\"type\":\"string\"},{\"name\":\"dataUrl\",\"type\":\"string\"},{\"name\":\"price\",\"type\":\"uint256\"}],\"name\":\"addLead\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"id\",\"type\":\"uint32\"},{\"name\":\"status\",\"type\":\"uint8\"}],\"name\":\"setLeadStatus\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"id\",\"type\":\"uint32\"},{\"name\":\"status\",\"type\":\"uint8\"}],\"name\":\"setTrackerStatus\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"config\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"id\",\"type\":\"uint32\"},{\"name\":\"ownerAddress\",\"type\":\"address\"}],\"name\":\"addUserAddress\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"configAddress\",\"type\":\"address\"}],\"name\":\"setConfigAddress\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"id\",\"type\":\"uint32\"},{\"name\":\"companyId\",\"type\":\"uint32\"},{\"name\":\"payerAddress\",\"type\":\"address\"},{\"name\":\"name\",\"type\":\"string\"},{\"name\":\"dataUrl\",\"type\":\"string\"},{\"name\":\"cost\",\"type\":\"uint256\"}],\"name\":\"addOffer\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"\",\"type\":\"uint32\"}],\"name\":\"trackers\",\"outputs\":[{\"name\":\"createdAt\",\"type\":\"uint256\"},{\"name\":\"ownerAddress\",\"type\":\"address\"},{\"name\":\"name\",\"type\":\"string\"},{\"name\":\"dataUrl\",\"type\":\"string\"},{\"name\":\"status\",\"type\":\"uint8\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"\",\"type\":\"uint32\"}],\"name\":\"offers\",\"outputs\":[{\"name\":\"createdAt\",\"type\":\"uint256\"},{\"name\":\"companyId\",\"type\":\"uint32\"},{\"name\":\"payerAddress\",\"type\":\"address\"},{\"name\":\"name\",\"type\":\"string\"},{\"name\":\"dataUrl\",\"type\":\"string\"},{\"name\":\"cost\",\"type\":\"uint256\"},{\"name\":\"status\",\"type\":\"uint8\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"id\",\"type\":\"uint32\"},{\"name\":\"intermediaryAddress\",\"type\":\"address\"},{\"name\":\"percent\",\"type\":\"uint32\"}],\"name\":\"addLeadIntermediary\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"id\",\"type\":\"uint32\"},{\"name\":\"ownerAddress\",\"type\":\"address\"},{\"name\":\"name\",\"type\":\"string\"},{\"name\":\"dataUrl\",\"type\":\"string\"}],\"name\":\"registerTracker\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"id\",\"type\":\"uint32\"},{\"name\":\"reportId\",\"type\":\"uint32\"},{\"name\":\"kycLevel\",\"type\":\"uint8\"},{\"name\":\"dataUrl\",\"type\":\"string\"}],\"name\":\"addUserKycReport\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"id\",\"type\":\"uint32\"},{\"name\":\"status\",\"type\":\"uint8\"}],\"name\":\"setUserStatus\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"\",\"type\":\"uint32\"}],\"name\":\"users\",\"outputs\":[{\"name\":\"createdAt\",\"type\":\"uint256\"},{\"name\":\"numOfAddresses\",\"type\":\"uint8\"},{\"name\":\"role\",\"type\":\"string\"},{\"name\":\"kycLevel\",\"type\":\"uint8\"},{\"name\":\"numOfKycReports\",\"type\":\"uint16\"},{\"name\":\"pubKey\",\"type\":\"string\"},{\"name\":\"status\",\"type\":\"uint8\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"id\",\"type\":\"uint32\"},{\"name\":\"ownerId\",\"type\":\"uint32\"},{\"name\":\"ownerAddress\",\"type\":\"address\"},{\"name\":\"name\",\"type\":\"string\"},{\"name\":\"dataUrl\",\"type\":\"string\"}],\"name\":\"registerCompany\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"token\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"id\",\"type\":\"uint32\"},{\"name\":\"num\",\"type\":\"uint8\"}],\"name\":\"getLeadIntermediaryPercent\",\"outputs\":[{\"name\":\"\",\"type\":\"uint32\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"id\",\"type\":\"uint32\"},{\"name\":\"status\",\"type\":\"uint8\"}],\"name\":\"setCompanyStatus\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"name\":\"configAddress\",\"type\":\"address\"},{\"name\":\"tokenAddress\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"ownerAddress\",\"type\":\"address\"},{\"indexed\":false,\"name\":\"id\",\"type\":\"uint32\"},{\"indexed\":false,\"name\":\"role\",\"type\":\"string\"}],\"name\":\"UserRegistered\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"ownerAddress\",\"type\":\"address\"},{\"indexed\":false,\"name\":\"additionalAddress\",\"type\":\"address\"}],\"name\":\"UserAddressAdded\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"ownerAddress\",\"type\":\"address\"}],\"name\":\"UserPubKeyUpdated\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"ownerAddress\",\"type\":\"address\"},{\"indexed\":false,\"name\":\"kycLevel\",\"type\":\"uint8\"}],\"name\":\"UserKycReportAdded\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"ownerAddress\",\"type\":\"address\"},{\"indexed\":false,\"name\":\"status\",\"type\":\"uint8\"}],\"name\":\"UserStatusChanged\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"ownerAddress\",\"type\":\"address\"},{\"indexed\":false,\"name\":\"id\",\"type\":\"uint32\"},{\"indexed\":false,\"name\":\"name\",\"type\":\"string\"}],\"name\":\"CompanyRegistered\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"ownerAddress\",\"type\":\"address\"},{\"indexed\":false,\"name\":\"status\",\"type\":\"uint8\"}],\"name\":\"CompanyStatusChanged\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"ownerAddress\",\"type\":\"address\"},{\"indexed\":false,\"name\":\"id\",\"type\":\"uint32\"},{\"indexed\":false,\"name\":\"name\",\"type\":\"string\"}],\"name\":\"TrackerRegistered\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"ownerAddress\",\"type\":\"address\"},{\"indexed\":false,\"name\":\"status\",\"type\":\"uint8\"}],\"name\":\"TrackerStatusChanged\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"ownerAddress\",\"type\":\"address\"},{\"indexed\":false,\"name\":\"id\",\"type\":\"uint32\"},{\"indexed\":false,\"name\":\"name\",\"type\":\"string\"}],\"name\":\"OfferAdded\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"ownerAddress\",\"type\":\"address\"},{\"indexed\":false,\"name\":\"status\",\"type\":\"uint8\"}],\"name\":\"OfferStatusChanged\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"ownerAddress\",\"type\":\"address\"},{\"indexed\":false,\"name\":\"id\",\"type\":\"uint32\"},{\"indexed\":false,\"name\":\"price\",\"type\":\"uint256\"}],\"name\":\"LeadAdded\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"ownerAddress\",\"type\":\"address\"},{\"indexed\":false,\"name\":\"status\",\"type\":\"uint8\"}],\"name\":\"LeadStatusChanged\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"ownerAddress\",\"type\":\"address\"},{\"indexed\":false,\"name\":\"ownerAmount\",\"type\":\"uint256\"}],\"name\":\"LeadSold\",\"type\":\"event\"}]"

// HoQuPlatformBin is the compiled bytecode used for deploying new contracts.
const HoQuPlatformBin = `0x6060604052341561000f57600080fd5b604051604080613654833981016040528080519190602001805160008054600160a060020a03958616600160a060020a031991821617909155600180549590921694169390931790925550506135ea8061006a6000396000f3006060604052600436106101455763ffffffff60e060020a600035041663096e8090811461014a57806333a7140a146101885780633bc7fd2c1461024c578063416330491461031057806348412be01461032c578063533781071461034e578063548c45cd146103aa57806364c0bff8146105065780636edfb910146106885780636ef17e16146106aa57806372c9f8ac1461076b578063732808d01461078d57806379502c55146107af5780638120dcec146107c257806383a12de9146107ea5780638e9d166f1461080b578063a0801e4e146108c1578063a13bb90114610a10578063a32dd96814610b74578063aa3e2cf914610ba3578063c024b49c14610c50578063c0b5e1ae14610cbc578063e2226a5614610cde578063f4d659d714610e07578063fc0c546a14610ebb578063fc4edfab14610ece578063ffd7159914610f09575b600080fd5b341561015557600080fd5b61016c63ffffffff6004351660ff60243516610f2b565b604051600160a060020a03909116815260200160405180910390f35b341561019357600080fd5b6101ab63ffffffff6004351661ffff60243516610f9c565b60405184815263ffffffff84166020820152604081018360048111156101cd57fe5b60ff16815260200180602001828103825283818151815260200191508051906020019080838360005b8381101561020e5780820151838201526020016101f6565b50505050905090810190601f16801561023b5780820380516001836020036101000a031916815260200191505b509550505050505060405180910390f35b341561025757600080fd5b6102fc6004803563ffffffff169060446024803590810190830135806020601f82018190048102016040519081016040528181529291906020840183838082843782019150505050505091908035600160a060020a031690602001909190803590602001908201803590602001908080601f0160208091040260200160405190810160405281815292919060208401838380828437509496506110d095505050505050565b604051901515815260200160405180910390f35b341561031b57600080fd5b6102fc63ffffffff60043516611391565b341561033757600080fd5b61016c63ffffffff6004351660ff60243516611950565b341561035957600080fd5b6102fc6004803563ffffffff169060446024803590810190830135806020601f820181900481020160405190810160405281815292919060208401838380828437509496506119bc95505050505050565b34156103b557600080fd5b6103c663ffffffff60043516611afe565b60405186815263ffffffff86166020820152600160a060020a0385166040820152606081016080820160a083018460058111156103ff57fe5b60ff1681526020848203810184528754600260018216156101000260001901909116049082018190526040909101908790801561047d5780601f106104525761010080835404028352916020019161047d565b820191906000526020600020905b81548152906001019060200180831161046057829003601f168201915b50508381038252855460026000196101006001841615020190911604808252602090910190869080156104f15780601f106104c6576101008083540402835291602001916104f1565b820191906000526020600020905b8154815290600101906020018083116104d457829003601f168201915b50509850505050505050505060405180910390f35b341561051157600080fd5b61052263ffffffff60043516611b47565b6040518a815263ffffffff808b1660208301528981166040830152600160a060020a03891660608301528716608082015260e0810184905260ff831661010082015260a0810160c08201610120830184600581111561057d57fe5b60ff168152602084820381018452895460026001821615610100026000190190911604908201819052604090910190899080156105fb5780601f106105d0576101008083540402835291602001916105fb565b820191906000526020600020905b8154815290600101906020018083116105de57829003601f168201915b505083810382528754600260001961010060018416150201909116048082526020909101908890801561066f5780601f106106445761010080835404028352916020019161066f565b820191906000526020600020905b81548152906001019060200180831161065257829003601f168201915b50509c5050505050505050505050505060405180910390f35b341561069357600080fd5b6102fc63ffffffff6004351660ff60243516611bb8565b34156106b557600080fd5b6102fc63ffffffff60048035821691602480358216926044358316926064351691600160a060020a0360843516919060c49060a43590810190830135806020601f8201819004810201604051908101604052818152929190602084018383808284378201915050505050509190803590602001908201803590602001908080601f0160208091040260200160405190810160405281815292919060208401838380828437509496505093359350611d2092505050565b341561077657600080fd5b6102fc63ffffffff6004351660ff60243516612076565b341561079857600080fd5b6102fc63ffffffff6004351660ff602435166121ca565b34156107ba57600080fd5b61016c61230f565b34156107cd57600080fd5b6102fc63ffffffff60043516600160a060020a036024351661231e565b34156107f557600080fd5b610809600160a060020a0360043516612495565b005b341561081657600080fd5b6102fc63ffffffff600480358216916024803590911691600160a060020a03604435169160849060643590810190830135806020601f8201819004810201604051908101604052818152929190602084018383808284378201915050505050509190803590602001908201803590602001908080601f016020809104026020016040519081016040528181529291906020840183838082843750949650509335935061254792505050565b34156108cc57600080fd5b6108dd63ffffffff6004351661280b565b604051858152600160a060020a038516602082015260408101606082016080830184600581111561090a57fe5b60ff168152602084820381018452875460026001821615610100026000190190911604908201819052604090910190879080156109885780601f1061095d57610100808354040283529160200191610988565b820191906000526020600020905b81548152906001019060200180831161096b57829003601f168201915b50508381038252855460026000196101006001841615020190911604808252602090910190869080156109fc5780601f106109d1576101008083540402835291602001916109fc565b820191906000526020600020905b8154815290600101906020018083116109df57829003601f168201915b505097505050505050505060405180910390f35b3415610a1b57600080fd5b610a2c63ffffffff60043516612842565b60405187815263ffffffff87166020820152600160a060020a038616604082015260a08101839052606081016080820160c08301846005811115610a6c57fe5b60ff16815260208482038101845288546002600182161561010002600019019091160490820181905260409091019088908015610aea5780601f10610abf57610100808354040283529160200191610aea565b820191906000526020600020905b815481529060010190602001808311610acd57829003601f168201915b5050838103825286546002600019610100600184161502019091160480825260209091019087908015610b5e5780601f10610b3357610100808354040283529160200191610b5e565b820191906000526020600020905b815481529060010190602001808311610b4157829003601f168201915b5050995050505050505050505060405180910390f35b3415610b7f57600080fd5b6102fc63ffffffff600435811690600160a060020a03602435169060443516612892565b3415610bae57600080fd5b6102fc6004803563ffffffff169060248035600160a060020a0316919060649060443590810190830135806020601f8201819004810201604051908101604052818152929190602084018383808284378201915050505050509190803590602001908201803590602001908080601f0160208091040260200160405190810160405281815292919060208401838380828437509496506129ea95505050505050565b3415610c5b57600080fd5b6102fc63ffffffff60048035821691602480359091169160ff604435169160849060643590810190830135806020601f82018190048102016040519081016040528181529291906020840183838082843750949650612bed95505050505050565b3415610cc757600080fd5b6102fc63ffffffff6004351660ff60243516612e3c565b3415610ce957600080fd5b610cfa63ffffffff60043516612f86565b60405187815260ff871660208201526040810160608201866004811115610d1d57fe5b60ff16815261ffff861660208201526040810190606001846005811115610d4057fe5b60ff16815260208482038101845289546002600182161561010002600019019091160490820181905260409091019089908015610dbe5780601f10610d9357610100808354040283529160200191610dbe565b820191906000526020600020905b815481529060010190602001808311610da157829003601f168201915b5050838103825285546002600019610100600184161502019091160480825260209091019086908015610b5e5780601f10610b3357610100808354040283529160200191610b5e565b3415610e1257600080fd5b6102fc63ffffffff600480358216916024803590911691600160a060020a03604435169160849060643590810190830135806020601f8201819004810201604051908101604052818152929190602084018383808284378201915050505050509190803590602001908201803590602001908080601f016020809104026020016040519081016040528181529291906020840183838082843750949650612fcc95505050505050565b3415610ec657600080fd5b61016c6132ff565b3415610ed957600080fd5b610ef063ffffffff6004351660ff6024351661330e565b60405163ffffffff909116815260200160405180910390f35b3415610f1457600080fd5b6102fc63ffffffff6004351660ff60243516613376565b60008063ffffffff8416600090815260066020526040902060070154610100900460ff166005811115610f5a57fe5b1415610f6557600080fd5b5063ffffffff8216600090815260066020908152604080832060ff85168452600501909152902054600160a060020a031692915050565b6000806000610fa9613511565b63ffffffff861660009081526002602052604081206008015460ff166005811115610fd057fe5b1415610fdb57600080fd5b63ffffffff868116600090815260026020818152604080842061ffff8b16855260050182529283902080546001808301549285018054929784169664010000000090940460ff1695909485946101009385161593909302600019019093160491601f8301819004810201905190810160405280929190818152602001828054600181600116156101000203166002900480156110b85780601f1061108d576101008083540402835291602001916110b8565b820191906000526020600020905b81548152906001019060200180831161109b57829003601f168201915b50505050509050935093509350935092959194509250565b60008054600160a060020a0316633377925482604051602001526040518163ffffffff1660e060020a028152600401602060405180830381600087803b151561111857600080fd5b6102c65a03f1151561112957600080fd5b50505060405180519050600160a060020a031633600160a060020a031614151561115257600080fd5b63ffffffff851660009081526002602052604081206008015460ff16600581111561117957fe5b1461118357600080fd5b60e06040519081016040908152428252600160208301528101859052606081016000815260006020820152604081018490526060016001905263ffffffff8616600090815260026020526040902081518155602082015160028201805460ff191660ff9290921691909117905560408201518160030190805161120a929160200190613523565b5060608201518160040160006101000a81548160ff0219169083600481111561122f57fe5b0217905550608082015160068201805461ffff191661ffff9290921691909117905560a08201518160070190805161126b929160200190613523565b5060c082015160088201805460ff1916600183600581111561128957fe5b0217905550505063ffffffff8516600090815260026020908152604080832083805260010190915290819020805473ffffffffffffffffffffffffffffffffffffffff1916600160a060020a038616908117909155907fca4edd8f681565cad296907b3a63102df328c052001ddd2e937ce5d8da31b90590879087905163ffffffff8316815260406020820181815290820183818151815260200191508051906020019080838360005b8381101561134b578082015183820152602001611333565b50505050905090810190601f1680156113785780820380516001836020036101000a031916815260200191505b50935050505060405180910390a2506001949350505050565b6000805481908190819081908190819081908190600160a060020a0316633377925482604051602001526040518163ffffffff1660e060020a028152600401602060405180830381600087803b15156113e957600080fd5b6102c65a03f115156113fa57600080fd5b50505060405180519050600160a060020a031633600160a060020a031614151561142357600080fd5b600463ffffffff8b16600090815260066020526040902060070154610100900460ff16600581111561145157fe5b1415801561148b5750600563ffffffff8b16600090815260066020526040902060070154610100900460ff16600581111561148857fe5b14155b151561149657600080fd5b63ffffffff8a16600090815260066020526040812060040154116114b957600080fd5b63ffffffff8a8116600090815260066020908152604080832060078101805461ff001916610400179055600181015460e060020a9004909416835260049091528082208254939b50995061159292670de0b6b3a76400009261158692600160a060020a039092169163e1489191919051602001526040518163ffffffff1660e060020a028152600401602060405180830381600087803b151561155b57600080fd5b6102c65a03f1151561156c57600080fd5b505050604051805160048d0154915063ffffffff6134bd16565b9063ffffffff6134e816565b60048901549096506115aa908763ffffffff6134ff16565b600180549089015460048b0154929750600160a060020a03918216926323b872dd9264010000000090920490911690309060006040516020015260405160e060020a63ffffffff8616028152600160a060020a0393841660048201529190921660248201526044810191909152606401602060405180830381600087803b151561163357600080fd5b6102c65a03f1151561164457600080fd5b5050506040518051505060015460008054600160a060020a039283169263a9059cbb929116906397c0262a90604051602001526040518163ffffffff1660e060020a028152600401602060405180830381600087803b15156116a557600080fd5b6102c65a03f115156116b657600080fd5b505050604051805190508860006040516020015260405160e060020a63ffffffff8516028152600160a060020a0390921660048301526024820152604401602060405180830381600087803b151561170d57600080fd5b6102c65a03f1151561171e57600080fd5b50505060405180515060009450505b63ffffffff8a1660009081526006602052604090206007015460ff908116908516101561184d5763ffffffff8a8116600090815260066020818152604080842060ff8a1685526005810183528185205493019091529091205460048b0154600160a060020a039092169550821693506117b2916305f5e100916115869186906134bd16565b600154909150600160a060020a031663a9059cbb848360006040516020015260405160e060020a63ffffffff8516028152600160a060020a0390921660048301526024820152604401602060405180830381600087803b151561181457600080fd5b6102c65a03f1151561182557600080fd5b5050506040518051506118409050858263ffffffff6134ff16565b945060019093019261172d565b6001805490890154600160a060020a039182169163a9059cbb91680100000000000000009004168760006040516020015260405160e060020a63ffffffff8516028152600160a060020a0390921660048301526024820152604401602060405180830381600087803b15156118c157600080fd5b6102c65a03f115156118d257600080fd5b5050506040518051505063ffffffff8a166000908152600660205260409081902060010154680100000000000000009004600160a060020a0316907f80155aceba8650f79571fb220d4ebc91df4635d5b8e765fdc835a56951aa14c49087905190815260200160405180910390a25060019998505050505050505050565b60008063ffffffff841660009081526002602052604090206008015460ff16600581111561197a57fe5b141561198557600080fd5b5063ffffffff8216600090815260026020908152604080832060ff85168452600101909152902054600160a060020a031692915050565b60008054600160a060020a0316633377925482604051602001526040518163ffffffff1660e060020a028152600401602060405180830381600087803b1515611a0457600080fd5b6102c65a03f11515611a1557600080fd5b50505060405180519050600160a060020a031633600160a060020a0316141515611a3e57600080fd5b63ffffffff831660009081526002602052604081206008015460ff166005811115611a6557fe5b1415611a7057600080fd5b63ffffffff83166000908152600260205260409020600701828051611a99929160200190613523565b5063ffffffff831660009081526002602090815260408083208380526001019091529081902054600160a060020a0316907f5775531a38eee0a6720d30a54a26748b21a9053e176d628a871b33885462e15e905160405180910390a250600192915050565b6003602081905260009182526040909120805460018201546004830154919363ffffffff821693640100000000909204600160a060020a03169260028301929091019060ff1686565b60066020526000908152604090208054600182015460048301546007840154929363ffffffff808416946401000000008504821694600160a060020a03680100000000000000008204169460e060020a90910490921692600281019260039091019160ff808216916101009004168a565b60008054600160a060020a0316633377925482604051602001526040518163ffffffff1660e060020a028152600401602060405180830381600087803b1515611c0057600080fd5b6102c65a03f11515611c1157600080fd5b50505060405180519050600160a060020a031633600160a060020a0316141515611c3a57600080fd5b63ffffffff8316600090815260046020526040812060059081015460ff1690811115611c6257fe5b1415611c6d57600080fd5b63ffffffff8316600090815260046020526040902060059081018054849260ff19909116906001908490811115611ca057fe5b021790555063ffffffff831660009081526004602052604090819020600101546401000000009004600160a060020a0316907f750a2d352ffc8961cc1d60bce39c071252a74b2292ae59144609532c630cca0d9084905180826005811115611d0457fe5b60ff16815260200191505060405180910390a250600192915050565b60008054600160a060020a0316633377925482604051602001526040518163ffffffff1660e060020a028152600401602060405180830381600087803b1515611d6857600080fd5b6102c65a03f11515611d7957600080fd5b50505060405180519050600160a060020a031633600160a060020a0316141515611da257600080fd5b600063ffffffff8a16600090815260066020526040902060070154610100900460ff166005811115611dd057fe5b14611dda57600080fd5b63ffffffff881660009081526002602052604081206008015460ff166005811115611e0157fe5b1415611e0c57600080fd5b63ffffffff8716600090815260056020819052604082206004015460ff1690811115611e3457fe5b1415611e3f57600080fd5b63ffffffff8616600090815260046020526040812060059081015460ff1690811115611e6757fe5b1415611e7257600080fd5b610140604051908101604052804281526020018863ffffffff1681526020018963ffffffff16815260200186600160a060020a031681526020018763ffffffff168152602001848152602001858152602001838152602001600060ff16815260200160016005811115611ee157fe5b905263ffffffff8a16600090815260066020526040902081518155602082015160018201805463ffffffff191663ffffffff9290921691909117905560408201518160010160046101000a81548163ffffffff021916908363ffffffff16021790555060608201518160010160086101000a815481600160a060020a030219169083600160a060020a03160217905550608082015181600101601c6101000a81548163ffffffff021916908363ffffffff16021790555060a082015181600201908051611fb2929160200190613523565b5060c082015181600301908051611fcd929160200190613523565b5060e0820151816004015561010082015160078201805460ff191660ff9290921691909117905561012082015160078201805461ff00191661010083600581111561201457fe5b021790555090505084600160a060020a03167f6d07e03cac4388c92944a9c57a9da1fc2e36407d4f8e9ad1d97ea090c18a0de58a8460405163ffffffff909216825260208201526040908101905180910390a250600198975050505050505050565b60008054600160a060020a0316633377925482604051602001526040518163ffffffff1660e060020a028152600401602060405180830381600087803b15156120be57600080fd5b6102c65a03f115156120cf57600080fd5b50505060405180519050600160a060020a031633600160a060020a03161415156120f857600080fd5b600063ffffffff8416600090815260066020526040902060070154610100900460ff16600581111561212657fe5b141561213157600080fd5b63ffffffff83166000908152600660205260409020600701805483919061ff00191661010083600581111561216257fe5b021790555063ffffffff83166000908152600660205260409081902060010154680100000000000000009004600160a060020a0316907f8918986add4d2e70f3634bbbfb41e54852e84c2097407c6e77dda5b2350585f89084905180826005811115611d0457fe5b60008054600160a060020a0316633377925482604051602001526040518163ffffffff1660e060020a028152600401602060405180830381600087803b151561221257600080fd5b6102c65a03f1151561222357600080fd5b50505060405180519050600160a060020a031633600160a060020a031614151561224c57600080fd5b63ffffffff8316600090815260056020819052604082206004015460ff169081111561227457fe5b141561227f57600080fd5b63ffffffff831660009081526005602081905260409091206004018054849260ff199091169060019084908111156122b357fe5b021790555063ffffffff83166000908152600560205260409081902060010154600160a060020a0316907f8165c172797b06d1dee047ea5ca7b8b07d5d5f1571354ab1377ec50bb4ae6e9a9084905180826005811115611d0457fe5b600054600160a060020a031681565b60008054600160a060020a0316633377925482604051602001526040518163ffffffff1660e060020a028152600401602060405180830381600087803b151561236657600080fd5b6102c65a03f1151561237757600080fd5b50505060405180519050600160a060020a031633600160a060020a03161415156123a057600080fd5b63ffffffff831660009081526002602052604081206008015460ff1660058111156123c757fe5b14156123d257600080fd5b63ffffffff83166000908152600260208181526040808420928301805460ff90811686526001948501909352818520805473ffffffffffffffffffffffffffffffffffffffff1916600160a060020a0389811691909117909155815460ff198116908516909501909316939093179092558280529181902054909116907fb8a06c596f027284f088ce1ba17842629ca63619a7303c94b3645e310ae34e4490849051600160a060020a03909116815260200160405180910390a250600192915050565b60008054600160a060020a031690633377925490604051602001526040518163ffffffff1660e060020a028152600401602060405180830381600087803b15156124de57600080fd5b6102c65a03f115156124ef57600080fd5b50505060405180519050600160a060020a031633600160a060020a031614151561251857600080fd5b6000805473ffffffffffffffffffffffffffffffffffffffff1916600160a060020a0392909216919091179055565b60008054600160a060020a0316633377925482604051602001526040518163ffffffff1660e060020a028152600401602060405180830381600087803b151561258f57600080fd5b6102c65a03f115156125a057600080fd5b50505060405180519050600160a060020a031633600160a060020a03161415156125c957600080fd5b63ffffffff8716600090815260046020526040812060059081015460ff16908111156125f157fe5b146125fb57600080fd5b63ffffffff861660009081526003602052604081206004015460ff16600581111561262257fe5b141561262d57600080fd5b60e0604051908101604052804281526020018763ffffffff16815260200186600160a060020a031681526020018581526020018481526020018381526020016001600581111561267957fe5b905263ffffffff8816600090815260046020526040902081518155602082015160018201805463ffffffff191663ffffffff9290921691909117905560408201518160010160046101000a815481600160a060020a030219169083600160a060020a031602179055506060820151816002019080516126fc929160200190613523565b50608082015181600301908051612717929160200190613523565b5060a0820151816004015560c08201518160050160006101000a81548160ff0219169083600581111561274657fe5b021790555090505084600160a060020a03167f8df4f5eda431c27df93b709d2a92c920b6fd967a41434c7709e79c20fcfbe477888660405163ffffffff8316815260406020820181815290820183818151815260200191508051906020019080838360005b838110156127c35780820151838201526020016127ab565b50505050905090810190601f1680156127f05780820380516001836020036101000a031916815260200191505b50935050505060405180910390a25060019695505050505050565b60056020526000908152604090208054600182015460048301549192600160a060020a039091169160028201916003019060ff1685565b600460208190526000918252604090912080546001820154928201546005830154919363ffffffff811693640100000000909104600160a060020a03169260028201926003909201919060ff1687565b60008054600160a060020a0316633377925482604051602001526040518163ffffffff1660e060020a028152600401602060405180830381600087803b15156128da57600080fd5b6102c65a03f115156128eb57600080fd5b50505060405180519050600160a060020a031633600160a060020a031614151561291457600080fd5b600063ffffffff8516600090815260066020526040902060070154610100900460ff16600581111561294257fe5b141561294d57600080fd5b5063ffffffff928316600081815260066020818152604080842060078101805460ff9081168752600583018552838720805473ffffffffffffffffffffffffffffffffffffffff1916600160a060020a039b909b169a909a179099558054891686529084018352908420805463ffffffff19169690981695909517909655919052909252815460ff19811690821660019081019092161790915590565b60008054600160a060020a0316633377925482604051602001526040518163ffffffff1660e060020a028152600401602060405180830381600087803b1515612a3257600080fd5b6102c65a03f11515612a4357600080fd5b50505060405180519050600160a060020a031633600160a060020a0316141515612a6c57600080fd5b63ffffffff8516600090815260056020819052604082206004015460ff1690811115612a9457fe5b14612a9e57600080fd5b60a06040519081016040908152428252600160a060020a0386166020830152810184905260608101839052608081016001905263ffffffff8616600090815260056020526040902081518155602082015160018201805473ffffffffffffffffffffffffffffffffffffffff1916600160a060020a0392909216919091179055604082015181600201908051612b38929160200190613523565b50606082015181600301908051612b53929160200190613523565b50608082015160048201805460ff19166001836005811115612b7157fe5b021790555090505083600160a060020a03167fc3c3327b1be87be9f70641b266ed77609242644d48e2c95c7659798f19659c9a868560405163ffffffff8316815260406020820181815290820183818151815260200191508051906020019080838360008381101561134b578082015183820152602001611333565b60008054600160a060020a0316633377925482604051602001526040518163ffffffff1660e060020a028152600401602060405180830381600087803b1515612c3557600080fd5b6102c65a03f11515612c4657600080fd5b50505060405180519050600160a060020a031633600160a060020a0316141515612c6f57600080fd5b63ffffffff851660009081526002602052604081206008015460ff166005811115612c9657fe5b1415612ca157600080fd5b6080604051908101604052804281526020018563ffffffff168152602001846004811115612ccb57fe5b8152602090810184905263ffffffff87166000908152600282526040808220600681015461ffff1683526005019092522081518155602082015160018201805463ffffffff191663ffffffff92909216919091179055604082015160018201805464ff000000001916640100000000836004811115612d4657fe5b0217905550606082015181600201908051612d65929160200190613523565b50505063ffffffff8516600090815260026020526040902060068101805461ffff198116600161ffff928316810190921617909155600491820180548693919260ff19909116918490811115612db757fe5b021790555063ffffffff851660009081526002602090815260408083208380526001019091529081902054600160a060020a0316907f3596012bbae932a31d934ff98585be7435f26aea5acf9e50d7ef6870feb2ace69085905180826004811115612e1e57fe5b60ff16815260200191505060405180910390a2506001949350505050565b60008054600160a060020a0316633377925482604051602001526040518163ffffffff1660e060020a028152600401602060405180830381600087803b1515612e8457600080fd5b6102c65a03f11515612e9557600080fd5b50505060405180519050600160a060020a031633600160a060020a0316141515612ebe57600080fd5b63ffffffff831660009081526002602052604081206008015460ff166005811115612ee557fe5b1415612ef057600080fd5b63ffffffff83166000908152600260205260409020600801805483919060ff19166001836005811115612f1f57fe5b021790555063ffffffff831660009081526002602090815260408083208380526001019091529081902054600160a060020a0316907fb3b609c7a5e504a312603e90d906d1667fd5caff0a840ad07bf5709a3ca1a0459084905180826005811115611d0457fe5b600260208190526000918252604090912080549181015460048201546006830154600884015460ff9384169460038101949384169361ffff909316926007909101911687565b6000805481908190600160a060020a0316633377925482604051602001526040518163ffffffff1660e060020a028152600401602060405180830381600087803b151561301857600080fd5b6102c65a03f1151561302957600080fd5b50505060405180519050600160a060020a031633600160a060020a031614151561305257600080fd5b63ffffffff881660009081526003602052604081206004015460ff16600581111561307957fe5b1461308357600080fd5b63ffffffff871660009081526002602052604081206008015460ff1660058111156130aa57fe5b14156130b557600080fd5b5060009050805b63ffffffff87166000908152600260208190526040909120015460ff908116908216101561312b5763ffffffff8716600090815260026020908152604080832060ff85168452600101909152902054600160a060020a038781169116141561312357600191505b6001016130bc565b81151561313757600080fd5b60c0604051908101604052804281526020018863ffffffff16815260200187600160a060020a031681526020018681526020018581526020016001600581111561317d57fe5b905263ffffffff8916600090815260036020526040902081518155602082015160018201805463ffffffff191663ffffffff9290921691909117905560408201518160010160046101000a815481600160a060020a030219169083600160a060020a03160217905550606082015181600201908051613200929160200190613523565b5060808201518160030190805161321b929160200190613523565b5060a082015160048201805460ff1916600183600581111561323957fe5b021790555090505085600160a060020a03167f9766e50c254b00e26f81a14f5d41cf2275ab4b43e8300af9b3ccb3091479d2c4898760405163ffffffff8316815260406020820181815290820183818151815260200191508051906020019080838360005b838110156132b657808201518382015260200161329e565b50505050905090810190601f1680156132e35780820380516001836020036101000a031916815260200191505b50935050505060405180910390a2506001979650505050505050565b600154600160a060020a031681565b60008063ffffffff8416600090815260066020526040902060070154610100900460ff16600581111561333d57fe5b141561334857600080fd5b5063ffffffff918216600090815260066020818152604080842060ff95909516845293909101905220541690565b60008054600160a060020a0316633377925482604051602001526040518163ffffffff1660e060020a028152600401602060405180830381600087803b15156133be57600080fd5b6102c65a03f115156133cf57600080fd5b50505060405180519050600160a060020a031633600160a060020a03161415156133f857600080fd5b63ffffffff831660009081526003602052604081206004015460ff16600581111561341f57fe5b141561342a57600080fd5b63ffffffff83166000908152600360205260409020600401805483919060ff1916600183600581111561345957fe5b021790555063ffffffff831660009081526003602052604090819020600101546401000000009004600160a060020a0316907f22d8cc108624dd012daf3f911a3e3bc3807627cc8cfdac8459484d69623a61229084905180826005811115611d0457fe5b60008282028315806134d957508284828115156134d657fe5b04145b15156134e157fe5b9392505050565b60008082848115156134f657fe5b04949350505050565b60008282111561350b57fe5b50900390565b60206040519081016040526000815290565b828054600181600116156101000203166002900490600052602060002090601f016020900481019282601f1061356457805160ff1916838001178555613591565b82800160010185558215613591579182015b82811115613591578251825591602001919060010190613576565b5061359d9291506135a1565b5090565b6135bb91905b8082111561359d57600081556001016135a7565b905600a165627a7a72305820c0549c0c4e0346fed8262eac04eac8b8e1677202a07fadf4f7cbbf9a2cac24200029`

// DeployHoQuPlatform deploys a new Ethereum contract, binding an instance of HoQuPlatform to it.
func DeployHoQuPlatform(auth *bind.TransactOpts, backend bind.ContractBackend, configAddress common.Address, tokenAddress common.Address) (common.Address, *types.Transaction, *HoQuPlatform, error) {
	parsed, err := abi.JSON(strings.NewReader(HoQuPlatformABI))
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	address, tx, contract, err := bind.DeployContract(auth, parsed, common.FromHex(HoQuPlatformBin), backend, configAddress, tokenAddress)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &HoQuPlatform{HoQuPlatformCaller: HoQuPlatformCaller{contract: contract}, HoQuPlatformTransactor: HoQuPlatformTransactor{contract: contract}}, nil
}

// HoQuPlatform is an auto generated Go binding around an Ethereum contract.
type HoQuPlatform struct {
	HoQuPlatformCaller     // Read-only binding to the contract
	HoQuPlatformTransactor // Write-only binding to the contract
}

// HoQuPlatformCaller is an auto generated read-only Go binding around an Ethereum contract.
type HoQuPlatformCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// HoQuPlatformTransactor is an auto generated write-only Go binding around an Ethereum contract.
type HoQuPlatformTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// HoQuPlatformSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type HoQuPlatformSession struct {
	Contract     *HoQuPlatform     // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// HoQuPlatformCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type HoQuPlatformCallerSession struct {
	Contract *HoQuPlatformCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts       // Call options to use throughout this session
}

// HoQuPlatformTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type HoQuPlatformTransactorSession struct {
	Contract     *HoQuPlatformTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts       // Transaction auth options to use throughout this session
}

// HoQuPlatformRaw is an auto generated low-level Go binding around an Ethereum contract.
type HoQuPlatformRaw struct {
	Contract *HoQuPlatform // Generic contract binding to access the raw methods on
}

// HoQuPlatformCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type HoQuPlatformCallerRaw struct {
	Contract *HoQuPlatformCaller // Generic read-only contract binding to access the raw methods on
}

// HoQuPlatformTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type HoQuPlatformTransactorRaw struct {
	Contract *HoQuPlatformTransactor // Generic write-only contract binding to access the raw methods on
}

// NewHoQuPlatform creates a new instance of HoQuPlatform, bound to a specific deployed contract.
func NewHoQuPlatform(address common.Address, backend bind.ContractBackend) (*HoQuPlatform, error) {
	contract, err := bindHoQuPlatform(address, backend, backend)
	if err != nil {
		return nil, err
	}
	return &HoQuPlatform{HoQuPlatformCaller: HoQuPlatformCaller{contract: contract}, HoQuPlatformTransactor: HoQuPlatformTransactor{contract: contract}}, nil
}

// NewHoQuPlatformCaller creates a new read-only instance of HoQuPlatform, bound to a specific deployed contract.
func NewHoQuPlatformCaller(address common.Address, caller bind.ContractCaller) (*HoQuPlatformCaller, error) {
	contract, err := bindHoQuPlatform(address, caller, nil)
	if err != nil {
		return nil, err
	}
	return &HoQuPlatformCaller{contract: contract}, nil
}

// NewHoQuPlatformTransactor creates a new write-only instance of HoQuPlatform, bound to a specific deployed contract.
func NewHoQuPlatformTransactor(address common.Address, transactor bind.ContractTransactor) (*HoQuPlatformTransactor, error) {
	contract, err := bindHoQuPlatform(address, nil, transactor)
	if err != nil {
		return nil, err
	}
	return &HoQuPlatformTransactor{contract: contract}, nil
}

// bindHoQuPlatform binds a generic wrapper to an already deployed contract.
func bindHoQuPlatform(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(HoQuPlatformABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_HoQuPlatform *HoQuPlatformRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _HoQuPlatform.Contract.HoQuPlatformCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_HoQuPlatform *HoQuPlatformRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _HoQuPlatform.Contract.HoQuPlatformTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_HoQuPlatform *HoQuPlatformRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _HoQuPlatform.Contract.HoQuPlatformTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_HoQuPlatform *HoQuPlatformCallerRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _HoQuPlatform.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_HoQuPlatform *HoQuPlatformTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _HoQuPlatform.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_HoQuPlatform *HoQuPlatformTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _HoQuPlatform.Contract.contract.Transact(opts, method, params...)
}

// Companies is a free data retrieval call binding the contract method 0x548c45cd.
//
// Solidity: function companies( uint32) constant returns(createdAt uint256, ownerId uint32, ownerAddress address, name string, dataUrl string, status uint8)
func (_HoQuPlatform *HoQuPlatformCaller) Companies(opts *bind.CallOpts, arg0 uint32) (struct {
	CreatedAt    *big.Int
	OwnerId      uint32
	OwnerAddress common.Address
	Name         string
	DataUrl      string
	Status       uint8
}, error) {
	ret := new(struct {
		CreatedAt    *big.Int
		OwnerId      uint32
		OwnerAddress common.Address
		Name         string
		DataUrl      string
		Status       uint8
	})
	out := ret
	err := _HoQuPlatform.contract.Call(opts, out, "companies", arg0)
	return *ret, err
}

// Companies is a free data retrieval call binding the contract method 0x548c45cd.
//
// Solidity: function companies( uint32) constant returns(createdAt uint256, ownerId uint32, ownerAddress address, name string, dataUrl string, status uint8)
func (_HoQuPlatform *HoQuPlatformSession) Companies(arg0 uint32) (struct {
	CreatedAt    *big.Int
	OwnerId      uint32
	OwnerAddress common.Address
	Name         string
	DataUrl      string
	Status       uint8
}, error) {
	return _HoQuPlatform.Contract.Companies(&_HoQuPlatform.CallOpts, arg0)
}

// Companies is a free data retrieval call binding the contract method 0x548c45cd.
//
// Solidity: function companies( uint32) constant returns(createdAt uint256, ownerId uint32, ownerAddress address, name string, dataUrl string, status uint8)
func (_HoQuPlatform *HoQuPlatformCallerSession) Companies(arg0 uint32) (struct {
	CreatedAt    *big.Int
	OwnerId      uint32
	OwnerAddress common.Address
	Name         string
	DataUrl      string
	Status       uint8
}, error) {
	return _HoQuPlatform.Contract.Companies(&_HoQuPlatform.CallOpts, arg0)
}

// Config is a free data retrieval call binding the contract method 0x79502c55.
//
// Solidity: function config() constant returns(address)
func (_HoQuPlatform *HoQuPlatformCaller) Config(opts *bind.CallOpts) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _HoQuPlatform.contract.Call(opts, out, "config")
	return *ret0, err
}

// Config is a free data retrieval call binding the contract method 0x79502c55.
//
// Solidity: function config() constant returns(address)
func (_HoQuPlatform *HoQuPlatformSession) Config() (common.Address, error) {
	return _HoQuPlatform.Contract.Config(&_HoQuPlatform.CallOpts)
}

// Config is a free data retrieval call binding the contract method 0x79502c55.
//
// Solidity: function config() constant returns(address)
func (_HoQuPlatform *HoQuPlatformCallerSession) Config() (common.Address, error) {
	return _HoQuPlatform.Contract.Config(&_HoQuPlatform.CallOpts)
}

// GetLeadIntermediaryAddress is a free data retrieval call binding the contract method 0x096e8090.
//
// Solidity: function getLeadIntermediaryAddress(id uint32, num uint8) constant returns(address)
func (_HoQuPlatform *HoQuPlatformCaller) GetLeadIntermediaryAddress(opts *bind.CallOpts, id uint32, num uint8) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _HoQuPlatform.contract.Call(opts, out, "getLeadIntermediaryAddress", id, num)
	return *ret0, err
}

// GetLeadIntermediaryAddress is a free data retrieval call binding the contract method 0x096e8090.
//
// Solidity: function getLeadIntermediaryAddress(id uint32, num uint8) constant returns(address)
func (_HoQuPlatform *HoQuPlatformSession) GetLeadIntermediaryAddress(id uint32, num uint8) (common.Address, error) {
	return _HoQuPlatform.Contract.GetLeadIntermediaryAddress(&_HoQuPlatform.CallOpts, id, num)
}

// GetLeadIntermediaryAddress is a free data retrieval call binding the contract method 0x096e8090.
//
// Solidity: function getLeadIntermediaryAddress(id uint32, num uint8) constant returns(address)
func (_HoQuPlatform *HoQuPlatformCallerSession) GetLeadIntermediaryAddress(id uint32, num uint8) (common.Address, error) {
	return _HoQuPlatform.Contract.GetLeadIntermediaryAddress(&_HoQuPlatform.CallOpts, id, num)
}

// GetLeadIntermediaryPercent is a free data retrieval call binding the contract method 0xfc4edfab.
//
// Solidity: function getLeadIntermediaryPercent(id uint32, num uint8) constant returns(uint32)
func (_HoQuPlatform *HoQuPlatformCaller) GetLeadIntermediaryPercent(opts *bind.CallOpts, id uint32, num uint8) (uint32, error) {
	var (
		ret0 = new(uint32)
	)
	out := ret0
	err := _HoQuPlatform.contract.Call(opts, out, "getLeadIntermediaryPercent", id, num)
	return *ret0, err
}

// GetLeadIntermediaryPercent is a free data retrieval call binding the contract method 0xfc4edfab.
//
// Solidity: function getLeadIntermediaryPercent(id uint32, num uint8) constant returns(uint32)
func (_HoQuPlatform *HoQuPlatformSession) GetLeadIntermediaryPercent(id uint32, num uint8) (uint32, error) {
	return _HoQuPlatform.Contract.GetLeadIntermediaryPercent(&_HoQuPlatform.CallOpts, id, num)
}

// GetLeadIntermediaryPercent is a free data retrieval call binding the contract method 0xfc4edfab.
//
// Solidity: function getLeadIntermediaryPercent(id uint32, num uint8) constant returns(uint32)
func (_HoQuPlatform *HoQuPlatformCallerSession) GetLeadIntermediaryPercent(id uint32, num uint8) (uint32, error) {
	return _HoQuPlatform.Contract.GetLeadIntermediaryPercent(&_HoQuPlatform.CallOpts, id, num)
}

// GetUserAddress is a free data retrieval call binding the contract method 0x48412be0.
//
// Solidity: function getUserAddress(id uint32, num uint8) constant returns(address)
func (_HoQuPlatform *HoQuPlatformCaller) GetUserAddress(opts *bind.CallOpts, id uint32, num uint8) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _HoQuPlatform.contract.Call(opts, out, "getUserAddress", id, num)
	return *ret0, err
}

// GetUserAddress is a free data retrieval call binding the contract method 0x48412be0.
//
// Solidity: function getUserAddress(id uint32, num uint8) constant returns(address)
func (_HoQuPlatform *HoQuPlatformSession) GetUserAddress(id uint32, num uint8) (common.Address, error) {
	return _HoQuPlatform.Contract.GetUserAddress(&_HoQuPlatform.CallOpts, id, num)
}

// GetUserAddress is a free data retrieval call binding the contract method 0x48412be0.
//
// Solidity: function getUserAddress(id uint32, num uint8) constant returns(address)
func (_HoQuPlatform *HoQuPlatformCallerSession) GetUserAddress(id uint32, num uint8) (common.Address, error) {
	return _HoQuPlatform.Contract.GetUserAddress(&_HoQuPlatform.CallOpts, id, num)
}

// GetUserKycReport is a free data retrieval call binding the contract method 0x33a7140a.
//
// Solidity: function getUserKycReport(id uint32, num uint16) constant returns(uint256, uint32, uint8, string)
func (_HoQuPlatform *HoQuPlatformCaller) GetUserKycReport(opts *bind.CallOpts, id uint32, num uint16) (*big.Int, uint32, uint8, string, error) {
	var (
		ret0 = new(*big.Int)
		ret1 = new(uint32)
		ret2 = new(uint8)
		ret3 = new(string)
	)
	out := &[]interface{}{
		ret0,
		ret1,
		ret2,
		ret3,
	}
	err := _HoQuPlatform.contract.Call(opts, out, "getUserKycReport", id, num)
	return *ret0, *ret1, *ret2, *ret3, err
}

// GetUserKycReport is a free data retrieval call binding the contract method 0x33a7140a.
//
// Solidity: function getUserKycReport(id uint32, num uint16) constant returns(uint256, uint32, uint8, string)
func (_HoQuPlatform *HoQuPlatformSession) GetUserKycReport(id uint32, num uint16) (*big.Int, uint32, uint8, string, error) {
	return _HoQuPlatform.Contract.GetUserKycReport(&_HoQuPlatform.CallOpts, id, num)
}

// GetUserKycReport is a free data retrieval call binding the contract method 0x33a7140a.
//
// Solidity: function getUserKycReport(id uint32, num uint16) constant returns(uint256, uint32, uint8, string)
func (_HoQuPlatform *HoQuPlatformCallerSession) GetUserKycReport(id uint32, num uint16) (*big.Int, uint32, uint8, string, error) {
	return _HoQuPlatform.Contract.GetUserKycReport(&_HoQuPlatform.CallOpts, id, num)
}

// Leads is a free data retrieval call binding the contract method 0x64c0bff8.
//
// Solidity: function leads( uint32) constant returns(createdAt uint256, trackerId uint32, ownerId uint32, beneficiaryAddress address, offerId uint32, dataUrl string, meta string, price uint256, numOfIntermediaries uint8, status uint8)
func (_HoQuPlatform *HoQuPlatformCaller) Leads(opts *bind.CallOpts, arg0 uint32) (struct {
	CreatedAt           *big.Int
	TrackerId           uint32
	OwnerId             uint32
	BeneficiaryAddress  common.Address
	OfferId             uint32
	DataUrl             string
	Meta                string
	Price               *big.Int
	NumOfIntermediaries uint8
	Status              uint8
}, error) {
	ret := new(struct {
		CreatedAt           *big.Int
		TrackerId           uint32
		OwnerId             uint32
		BeneficiaryAddress  common.Address
		OfferId             uint32
		DataUrl             string
		Meta                string
		Price               *big.Int
		NumOfIntermediaries uint8
		Status              uint8
	})
	out := ret
	err := _HoQuPlatform.contract.Call(opts, out, "leads", arg0)
	return *ret, err
}

// Leads is a free data retrieval call binding the contract method 0x64c0bff8.
//
// Solidity: function leads( uint32) constant returns(createdAt uint256, trackerId uint32, ownerId uint32, beneficiaryAddress address, offerId uint32, dataUrl string, meta string, price uint256, numOfIntermediaries uint8, status uint8)
func (_HoQuPlatform *HoQuPlatformSession) Leads(arg0 uint32) (struct {
	CreatedAt           *big.Int
	TrackerId           uint32
	OwnerId             uint32
	BeneficiaryAddress  common.Address
	OfferId             uint32
	DataUrl             string
	Meta                string
	Price               *big.Int
	NumOfIntermediaries uint8
	Status              uint8
}, error) {
	return _HoQuPlatform.Contract.Leads(&_HoQuPlatform.CallOpts, arg0)
}

// Leads is a free data retrieval call binding the contract method 0x64c0bff8.
//
// Solidity: function leads( uint32) constant returns(createdAt uint256, trackerId uint32, ownerId uint32, beneficiaryAddress address, offerId uint32, dataUrl string, meta string, price uint256, numOfIntermediaries uint8, status uint8)
func (_HoQuPlatform *HoQuPlatformCallerSession) Leads(arg0 uint32) (struct {
	CreatedAt           *big.Int
	TrackerId           uint32
	OwnerId             uint32
	BeneficiaryAddress  common.Address
	OfferId             uint32
	DataUrl             string
	Meta                string
	Price               *big.Int
	NumOfIntermediaries uint8
	Status              uint8
}, error) {
	return _HoQuPlatform.Contract.Leads(&_HoQuPlatform.CallOpts, arg0)
}

// Offers is a free data retrieval call binding the contract method 0xa13bb901.
//
// Solidity: function offers( uint32) constant returns(createdAt uint256, companyId uint32, payerAddress address, name string, dataUrl string, cost uint256, status uint8)
func (_HoQuPlatform *HoQuPlatformCaller) Offers(opts *bind.CallOpts, arg0 uint32) (struct {
	CreatedAt    *big.Int
	CompanyId    uint32
	PayerAddress common.Address
	Name         string
	DataUrl      string
	Cost         *big.Int
	Status       uint8
}, error) {
	ret := new(struct {
		CreatedAt    *big.Int
		CompanyId    uint32
		PayerAddress common.Address
		Name         string
		DataUrl      string
		Cost         *big.Int
		Status       uint8
	})
	out := ret
	err := _HoQuPlatform.contract.Call(opts, out, "offers", arg0)
	return *ret, err
}

// Offers is a free data retrieval call binding the contract method 0xa13bb901.
//
// Solidity: function offers( uint32) constant returns(createdAt uint256, companyId uint32, payerAddress address, name string, dataUrl string, cost uint256, status uint8)
func (_HoQuPlatform *HoQuPlatformSession) Offers(arg0 uint32) (struct {
	CreatedAt    *big.Int
	CompanyId    uint32
	PayerAddress common.Address
	Name         string
	DataUrl      string
	Cost         *big.Int
	Status       uint8
}, error) {
	return _HoQuPlatform.Contract.Offers(&_HoQuPlatform.CallOpts, arg0)
}

// Offers is a free data retrieval call binding the contract method 0xa13bb901.
//
// Solidity: function offers( uint32) constant returns(createdAt uint256, companyId uint32, payerAddress address, name string, dataUrl string, cost uint256, status uint8)
func (_HoQuPlatform *HoQuPlatformCallerSession) Offers(arg0 uint32) (struct {
	CreatedAt    *big.Int
	CompanyId    uint32
	PayerAddress common.Address
	Name         string
	DataUrl      string
	Cost         *big.Int
	Status       uint8
}, error) {
	return _HoQuPlatform.Contract.Offers(&_HoQuPlatform.CallOpts, arg0)
}

// Token is a free data retrieval call binding the contract method 0xfc0c546a.
//
// Solidity: function token() constant returns(address)
func (_HoQuPlatform *HoQuPlatformCaller) Token(opts *bind.CallOpts) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _HoQuPlatform.contract.Call(opts, out, "token")
	return *ret0, err
}

// Token is a free data retrieval call binding the contract method 0xfc0c546a.
//
// Solidity: function token() constant returns(address)
func (_HoQuPlatform *HoQuPlatformSession) Token() (common.Address, error) {
	return _HoQuPlatform.Contract.Token(&_HoQuPlatform.CallOpts)
}

// Token is a free data retrieval call binding the contract method 0xfc0c546a.
//
// Solidity: function token() constant returns(address)
func (_HoQuPlatform *HoQuPlatformCallerSession) Token() (common.Address, error) {
	return _HoQuPlatform.Contract.Token(&_HoQuPlatform.CallOpts)
}

// Trackers is a free data retrieval call binding the contract method 0xa0801e4e.
//
// Solidity: function trackers( uint32) constant returns(createdAt uint256, ownerAddress address, name string, dataUrl string, status uint8)
func (_HoQuPlatform *HoQuPlatformCaller) Trackers(opts *bind.CallOpts, arg0 uint32) (struct {
	CreatedAt    *big.Int
	OwnerAddress common.Address
	Name         string
	DataUrl      string
	Status       uint8
}, error) {
	ret := new(struct {
		CreatedAt    *big.Int
		OwnerAddress common.Address
		Name         string
		DataUrl      string
		Status       uint8
	})
	out := ret
	err := _HoQuPlatform.contract.Call(opts, out, "trackers", arg0)
	return *ret, err
}

// Trackers is a free data retrieval call binding the contract method 0xa0801e4e.
//
// Solidity: function trackers( uint32) constant returns(createdAt uint256, ownerAddress address, name string, dataUrl string, status uint8)
func (_HoQuPlatform *HoQuPlatformSession) Trackers(arg0 uint32) (struct {
	CreatedAt    *big.Int
	OwnerAddress common.Address
	Name         string
	DataUrl      string
	Status       uint8
}, error) {
	return _HoQuPlatform.Contract.Trackers(&_HoQuPlatform.CallOpts, arg0)
}

// Trackers is a free data retrieval call binding the contract method 0xa0801e4e.
//
// Solidity: function trackers( uint32) constant returns(createdAt uint256, ownerAddress address, name string, dataUrl string, status uint8)
func (_HoQuPlatform *HoQuPlatformCallerSession) Trackers(arg0 uint32) (struct {
	CreatedAt    *big.Int
	OwnerAddress common.Address
	Name         string
	DataUrl      string
	Status       uint8
}, error) {
	return _HoQuPlatform.Contract.Trackers(&_HoQuPlatform.CallOpts, arg0)
}

// Users is a free data retrieval call binding the contract method 0xe2226a56.
//
// Solidity: function users( uint32) constant returns(createdAt uint256, numOfAddresses uint8, role string, kycLevel uint8, numOfKycReports uint16, pubKey string, status uint8)
func (_HoQuPlatform *HoQuPlatformCaller) Users(opts *bind.CallOpts, arg0 uint32) (struct {
	CreatedAt       *big.Int
	NumOfAddresses  uint8
	Role            string
	KycLevel        uint8
	NumOfKycReports uint16
	PubKey          string
	Status          uint8
}, error) {
	ret := new(struct {
		CreatedAt       *big.Int
		NumOfAddresses  uint8
		Role            string
		KycLevel        uint8
		NumOfKycReports uint16
		PubKey          string
		Status          uint8
	})
	out := ret
	err := _HoQuPlatform.contract.Call(opts, out, "users", arg0)
	return *ret, err
}

// Users is a free data retrieval call binding the contract method 0xe2226a56.
//
// Solidity: function users( uint32) constant returns(createdAt uint256, numOfAddresses uint8, role string, kycLevel uint8, numOfKycReports uint16, pubKey string, status uint8)
func (_HoQuPlatform *HoQuPlatformSession) Users(arg0 uint32) (struct {
	CreatedAt       *big.Int
	NumOfAddresses  uint8
	Role            string
	KycLevel        uint8
	NumOfKycReports uint16
	PubKey          string
	Status          uint8
}, error) {
	return _HoQuPlatform.Contract.Users(&_HoQuPlatform.CallOpts, arg0)
}

// Users is a free data retrieval call binding the contract method 0xe2226a56.
//
// Solidity: function users( uint32) constant returns(createdAt uint256, numOfAddresses uint8, role string, kycLevel uint8, numOfKycReports uint16, pubKey string, status uint8)
func (_HoQuPlatform *HoQuPlatformCallerSession) Users(arg0 uint32) (struct {
	CreatedAt       *big.Int
	NumOfAddresses  uint8
	Role            string
	KycLevel        uint8
	NumOfKycReports uint16
	PubKey          string
	Status          uint8
}, error) {
	return _HoQuPlatform.Contract.Users(&_HoQuPlatform.CallOpts, arg0)
}

// AddLead is a paid mutator transaction binding the contract method 0x6ef17e16.
//
// Solidity: function addLead(id uint32, ownerId uint32, trackerId uint32, offerId uint32, beneficiaryAddress address, meta string, dataUrl string, price uint256) returns(bool)
func (_HoQuPlatform *HoQuPlatformTransactor) AddLead(opts *bind.TransactOpts, id uint32, ownerId uint32, trackerId uint32, offerId uint32, beneficiaryAddress common.Address, meta string, dataUrl string, price *big.Int) (*types.Transaction, error) {
	return _HoQuPlatform.contract.Transact(opts, "addLead", id, ownerId, trackerId, offerId, beneficiaryAddress, meta, dataUrl, price)
}

// AddLead is a paid mutator transaction binding the contract method 0x6ef17e16.
//
// Solidity: function addLead(id uint32, ownerId uint32, trackerId uint32, offerId uint32, beneficiaryAddress address, meta string, dataUrl string, price uint256) returns(bool)
func (_HoQuPlatform *HoQuPlatformSession) AddLead(id uint32, ownerId uint32, trackerId uint32, offerId uint32, beneficiaryAddress common.Address, meta string, dataUrl string, price *big.Int) (*types.Transaction, error) {
	return _HoQuPlatform.Contract.AddLead(&_HoQuPlatform.TransactOpts, id, ownerId, trackerId, offerId, beneficiaryAddress, meta, dataUrl, price)
}

// AddLead is a paid mutator transaction binding the contract method 0x6ef17e16.
//
// Solidity: function addLead(id uint32, ownerId uint32, trackerId uint32, offerId uint32, beneficiaryAddress address, meta string, dataUrl string, price uint256) returns(bool)
func (_HoQuPlatform *HoQuPlatformTransactorSession) AddLead(id uint32, ownerId uint32, trackerId uint32, offerId uint32, beneficiaryAddress common.Address, meta string, dataUrl string, price *big.Int) (*types.Transaction, error) {
	return _HoQuPlatform.Contract.AddLead(&_HoQuPlatform.TransactOpts, id, ownerId, trackerId, offerId, beneficiaryAddress, meta, dataUrl, price)
}

// AddLeadIntermediary is a paid mutator transaction binding the contract method 0xa32dd968.
//
// Solidity: function addLeadIntermediary(id uint32, intermediaryAddress address, percent uint32) returns(bool)
func (_HoQuPlatform *HoQuPlatformTransactor) AddLeadIntermediary(opts *bind.TransactOpts, id uint32, intermediaryAddress common.Address, percent uint32) (*types.Transaction, error) {
	return _HoQuPlatform.contract.Transact(opts, "addLeadIntermediary", id, intermediaryAddress, percent)
}

// AddLeadIntermediary is a paid mutator transaction binding the contract method 0xa32dd968.
//
// Solidity: function addLeadIntermediary(id uint32, intermediaryAddress address, percent uint32) returns(bool)
func (_HoQuPlatform *HoQuPlatformSession) AddLeadIntermediary(id uint32, intermediaryAddress common.Address, percent uint32) (*types.Transaction, error) {
	return _HoQuPlatform.Contract.AddLeadIntermediary(&_HoQuPlatform.TransactOpts, id, intermediaryAddress, percent)
}

// AddLeadIntermediary is a paid mutator transaction binding the contract method 0xa32dd968.
//
// Solidity: function addLeadIntermediary(id uint32, intermediaryAddress address, percent uint32) returns(bool)
func (_HoQuPlatform *HoQuPlatformTransactorSession) AddLeadIntermediary(id uint32, intermediaryAddress common.Address, percent uint32) (*types.Transaction, error) {
	return _HoQuPlatform.Contract.AddLeadIntermediary(&_HoQuPlatform.TransactOpts, id, intermediaryAddress, percent)
}

// AddOffer is a paid mutator transaction binding the contract method 0x8e9d166f.
//
// Solidity: function addOffer(id uint32, companyId uint32, payerAddress address, name string, dataUrl string, cost uint256) returns(bool)
func (_HoQuPlatform *HoQuPlatformTransactor) AddOffer(opts *bind.TransactOpts, id uint32, companyId uint32, payerAddress common.Address, name string, dataUrl string, cost *big.Int) (*types.Transaction, error) {
	return _HoQuPlatform.contract.Transact(opts, "addOffer", id, companyId, payerAddress, name, dataUrl, cost)
}

// AddOffer is a paid mutator transaction binding the contract method 0x8e9d166f.
//
// Solidity: function addOffer(id uint32, companyId uint32, payerAddress address, name string, dataUrl string, cost uint256) returns(bool)
func (_HoQuPlatform *HoQuPlatformSession) AddOffer(id uint32, companyId uint32, payerAddress common.Address, name string, dataUrl string, cost *big.Int) (*types.Transaction, error) {
	return _HoQuPlatform.Contract.AddOffer(&_HoQuPlatform.TransactOpts, id, companyId, payerAddress, name, dataUrl, cost)
}

// AddOffer is a paid mutator transaction binding the contract method 0x8e9d166f.
//
// Solidity: function addOffer(id uint32, companyId uint32, payerAddress address, name string, dataUrl string, cost uint256) returns(bool)
func (_HoQuPlatform *HoQuPlatformTransactorSession) AddOffer(id uint32, companyId uint32, payerAddress common.Address, name string, dataUrl string, cost *big.Int) (*types.Transaction, error) {
	return _HoQuPlatform.Contract.AddOffer(&_HoQuPlatform.TransactOpts, id, companyId, payerAddress, name, dataUrl, cost)
}

// AddUserAddress is a paid mutator transaction binding the contract method 0x8120dcec.
//
// Solidity: function addUserAddress(id uint32, ownerAddress address) returns(bool)
func (_HoQuPlatform *HoQuPlatformTransactor) AddUserAddress(opts *bind.TransactOpts, id uint32, ownerAddress common.Address) (*types.Transaction, error) {
	return _HoQuPlatform.contract.Transact(opts, "addUserAddress", id, ownerAddress)
}

// AddUserAddress is a paid mutator transaction binding the contract method 0x8120dcec.
//
// Solidity: function addUserAddress(id uint32, ownerAddress address) returns(bool)
func (_HoQuPlatform *HoQuPlatformSession) AddUserAddress(id uint32, ownerAddress common.Address) (*types.Transaction, error) {
	return _HoQuPlatform.Contract.AddUserAddress(&_HoQuPlatform.TransactOpts, id, ownerAddress)
}

// AddUserAddress is a paid mutator transaction binding the contract method 0x8120dcec.
//
// Solidity: function addUserAddress(id uint32, ownerAddress address) returns(bool)
func (_HoQuPlatform *HoQuPlatformTransactorSession) AddUserAddress(id uint32, ownerAddress common.Address) (*types.Transaction, error) {
	return _HoQuPlatform.Contract.AddUserAddress(&_HoQuPlatform.TransactOpts, id, ownerAddress)
}

// AddUserKycReport is a paid mutator transaction binding the contract method 0xc024b49c.
//
// Solidity: function addUserKycReport(id uint32, reportId uint32, kycLevel uint8, dataUrl string) returns(bool)
func (_HoQuPlatform *HoQuPlatformTransactor) AddUserKycReport(opts *bind.TransactOpts, id uint32, reportId uint32, kycLevel uint8, dataUrl string) (*types.Transaction, error) {
	return _HoQuPlatform.contract.Transact(opts, "addUserKycReport", id, reportId, kycLevel, dataUrl)
}

// AddUserKycReport is a paid mutator transaction binding the contract method 0xc024b49c.
//
// Solidity: function addUserKycReport(id uint32, reportId uint32, kycLevel uint8, dataUrl string) returns(bool)
func (_HoQuPlatform *HoQuPlatformSession) AddUserKycReport(id uint32, reportId uint32, kycLevel uint8, dataUrl string) (*types.Transaction, error) {
	return _HoQuPlatform.Contract.AddUserKycReport(&_HoQuPlatform.TransactOpts, id, reportId, kycLevel, dataUrl)
}

// AddUserKycReport is a paid mutator transaction binding the contract method 0xc024b49c.
//
// Solidity: function addUserKycReport(id uint32, reportId uint32, kycLevel uint8, dataUrl string) returns(bool)
func (_HoQuPlatform *HoQuPlatformTransactorSession) AddUserKycReport(id uint32, reportId uint32, kycLevel uint8, dataUrl string) (*types.Transaction, error) {
	return _HoQuPlatform.Contract.AddUserKycReport(&_HoQuPlatform.TransactOpts, id, reportId, kycLevel, dataUrl)
}

// RegisterCompany is a paid mutator transaction binding the contract method 0xf4d659d7.
//
// Solidity: function registerCompany(id uint32, ownerId uint32, ownerAddress address, name string, dataUrl string) returns(bool)
func (_HoQuPlatform *HoQuPlatformTransactor) RegisterCompany(opts *bind.TransactOpts, id uint32, ownerId uint32, ownerAddress common.Address, name string, dataUrl string) (*types.Transaction, error) {
	return _HoQuPlatform.contract.Transact(opts, "registerCompany", id, ownerId, ownerAddress, name, dataUrl)
}

// RegisterCompany is a paid mutator transaction binding the contract method 0xf4d659d7.
//
// Solidity: function registerCompany(id uint32, ownerId uint32, ownerAddress address, name string, dataUrl string) returns(bool)
func (_HoQuPlatform *HoQuPlatformSession) RegisterCompany(id uint32, ownerId uint32, ownerAddress common.Address, name string, dataUrl string) (*types.Transaction, error) {
	return _HoQuPlatform.Contract.RegisterCompany(&_HoQuPlatform.TransactOpts, id, ownerId, ownerAddress, name, dataUrl)
}

// RegisterCompany is a paid mutator transaction binding the contract method 0xf4d659d7.
//
// Solidity: function registerCompany(id uint32, ownerId uint32, ownerAddress address, name string, dataUrl string) returns(bool)
func (_HoQuPlatform *HoQuPlatformTransactorSession) RegisterCompany(id uint32, ownerId uint32, ownerAddress common.Address, name string, dataUrl string) (*types.Transaction, error) {
	return _HoQuPlatform.Contract.RegisterCompany(&_HoQuPlatform.TransactOpts, id, ownerId, ownerAddress, name, dataUrl)
}

// RegisterTracker is a paid mutator transaction binding the contract method 0xaa3e2cf9.
//
// Solidity: function registerTracker(id uint32, ownerAddress address, name string, dataUrl string) returns(bool)
func (_HoQuPlatform *HoQuPlatformTransactor) RegisterTracker(opts *bind.TransactOpts, id uint32, ownerAddress common.Address, name string, dataUrl string) (*types.Transaction, error) {
	return _HoQuPlatform.contract.Transact(opts, "registerTracker", id, ownerAddress, name, dataUrl)
}

// RegisterTracker is a paid mutator transaction binding the contract method 0xaa3e2cf9.
//
// Solidity: function registerTracker(id uint32, ownerAddress address, name string, dataUrl string) returns(bool)
func (_HoQuPlatform *HoQuPlatformSession) RegisterTracker(id uint32, ownerAddress common.Address, name string, dataUrl string) (*types.Transaction, error) {
	return _HoQuPlatform.Contract.RegisterTracker(&_HoQuPlatform.TransactOpts, id, ownerAddress, name, dataUrl)
}

// RegisterTracker is a paid mutator transaction binding the contract method 0xaa3e2cf9.
//
// Solidity: function registerTracker(id uint32, ownerAddress address, name string, dataUrl string) returns(bool)
func (_HoQuPlatform *HoQuPlatformTransactorSession) RegisterTracker(id uint32, ownerAddress common.Address, name string, dataUrl string) (*types.Transaction, error) {
	return _HoQuPlatform.Contract.RegisterTracker(&_HoQuPlatform.TransactOpts, id, ownerAddress, name, dataUrl)
}

// RegisterUser is a paid mutator transaction binding the contract method 0x3bc7fd2c.
//
// Solidity: function registerUser(id uint32, role string, ownerAddress address, pubKey string) returns(bool)
func (_HoQuPlatform *HoQuPlatformTransactor) RegisterUser(opts *bind.TransactOpts, id uint32, role string, ownerAddress common.Address, pubKey string) (*types.Transaction, error) {
	return _HoQuPlatform.contract.Transact(opts, "registerUser", id, role, ownerAddress, pubKey)
}

// RegisterUser is a paid mutator transaction binding the contract method 0x3bc7fd2c.
//
// Solidity: function registerUser(id uint32, role string, ownerAddress address, pubKey string) returns(bool)
func (_HoQuPlatform *HoQuPlatformSession) RegisterUser(id uint32, role string, ownerAddress common.Address, pubKey string) (*types.Transaction, error) {
	return _HoQuPlatform.Contract.RegisterUser(&_HoQuPlatform.TransactOpts, id, role, ownerAddress, pubKey)
}

// RegisterUser is a paid mutator transaction binding the contract method 0x3bc7fd2c.
//
// Solidity: function registerUser(id uint32, role string, ownerAddress address, pubKey string) returns(bool)
func (_HoQuPlatform *HoQuPlatformTransactorSession) RegisterUser(id uint32, role string, ownerAddress common.Address, pubKey string) (*types.Transaction, error) {
	return _HoQuPlatform.Contract.RegisterUser(&_HoQuPlatform.TransactOpts, id, role, ownerAddress, pubKey)
}

// SellLead is a paid mutator transaction binding the contract method 0x41633049.
//
// Solidity: function sellLead(id uint32) returns(bool)
func (_HoQuPlatform *HoQuPlatformTransactor) SellLead(opts *bind.TransactOpts, id uint32) (*types.Transaction, error) {
	return _HoQuPlatform.contract.Transact(opts, "sellLead", id)
}

// SellLead is a paid mutator transaction binding the contract method 0x41633049.
//
// Solidity: function sellLead(id uint32) returns(bool)
func (_HoQuPlatform *HoQuPlatformSession) SellLead(id uint32) (*types.Transaction, error) {
	return _HoQuPlatform.Contract.SellLead(&_HoQuPlatform.TransactOpts, id)
}

// SellLead is a paid mutator transaction binding the contract method 0x41633049.
//
// Solidity: function sellLead(id uint32) returns(bool)
func (_HoQuPlatform *HoQuPlatformTransactorSession) SellLead(id uint32) (*types.Transaction, error) {
	return _HoQuPlatform.Contract.SellLead(&_HoQuPlatform.TransactOpts, id)
}

// SetCompanyStatus is a paid mutator transaction binding the contract method 0xffd71599.
//
// Solidity: function setCompanyStatus(id uint32, status uint8) returns(bool)
func (_HoQuPlatform *HoQuPlatformTransactor) SetCompanyStatus(opts *bind.TransactOpts, id uint32, status uint8) (*types.Transaction, error) {
	return _HoQuPlatform.contract.Transact(opts, "setCompanyStatus", id, status)
}

// SetCompanyStatus is a paid mutator transaction binding the contract method 0xffd71599.
//
// Solidity: function setCompanyStatus(id uint32, status uint8) returns(bool)
func (_HoQuPlatform *HoQuPlatformSession) SetCompanyStatus(id uint32, status uint8) (*types.Transaction, error) {
	return _HoQuPlatform.Contract.SetCompanyStatus(&_HoQuPlatform.TransactOpts, id, status)
}

// SetCompanyStatus is a paid mutator transaction binding the contract method 0xffd71599.
//
// Solidity: function setCompanyStatus(id uint32, status uint8) returns(bool)
func (_HoQuPlatform *HoQuPlatformTransactorSession) SetCompanyStatus(id uint32, status uint8) (*types.Transaction, error) {
	return _HoQuPlatform.Contract.SetCompanyStatus(&_HoQuPlatform.TransactOpts, id, status)
}

// SetConfigAddress is a paid mutator transaction binding the contract method 0x83a12de9.
//
// Solidity: function setConfigAddress(configAddress address) returns()
func (_HoQuPlatform *HoQuPlatformTransactor) SetConfigAddress(opts *bind.TransactOpts, configAddress common.Address) (*types.Transaction, error) {
	return _HoQuPlatform.contract.Transact(opts, "setConfigAddress", configAddress)
}

// SetConfigAddress is a paid mutator transaction binding the contract method 0x83a12de9.
//
// Solidity: function setConfigAddress(configAddress address) returns()
func (_HoQuPlatform *HoQuPlatformSession) SetConfigAddress(configAddress common.Address) (*types.Transaction, error) {
	return _HoQuPlatform.Contract.SetConfigAddress(&_HoQuPlatform.TransactOpts, configAddress)
}

// SetConfigAddress is a paid mutator transaction binding the contract method 0x83a12de9.
//
// Solidity: function setConfigAddress(configAddress address) returns()
func (_HoQuPlatform *HoQuPlatformTransactorSession) SetConfigAddress(configAddress common.Address) (*types.Transaction, error) {
	return _HoQuPlatform.Contract.SetConfigAddress(&_HoQuPlatform.TransactOpts, configAddress)
}

// SetLeadStatus is a paid mutator transaction binding the contract method 0x72c9f8ac.
//
// Solidity: function setLeadStatus(id uint32, status uint8) returns(bool)
func (_HoQuPlatform *HoQuPlatformTransactor) SetLeadStatus(opts *bind.TransactOpts, id uint32, status uint8) (*types.Transaction, error) {
	return _HoQuPlatform.contract.Transact(opts, "setLeadStatus", id, status)
}

// SetLeadStatus is a paid mutator transaction binding the contract method 0x72c9f8ac.
//
// Solidity: function setLeadStatus(id uint32, status uint8) returns(bool)
func (_HoQuPlatform *HoQuPlatformSession) SetLeadStatus(id uint32, status uint8) (*types.Transaction, error) {
	return _HoQuPlatform.Contract.SetLeadStatus(&_HoQuPlatform.TransactOpts, id, status)
}

// SetLeadStatus is a paid mutator transaction binding the contract method 0x72c9f8ac.
//
// Solidity: function setLeadStatus(id uint32, status uint8) returns(bool)
func (_HoQuPlatform *HoQuPlatformTransactorSession) SetLeadStatus(id uint32, status uint8) (*types.Transaction, error) {
	return _HoQuPlatform.Contract.SetLeadStatus(&_HoQuPlatform.TransactOpts, id, status)
}

// SetOfferStatus is a paid mutator transaction binding the contract method 0x6edfb910.
//
// Solidity: function setOfferStatus(id uint32, status uint8) returns(bool)
func (_HoQuPlatform *HoQuPlatformTransactor) SetOfferStatus(opts *bind.TransactOpts, id uint32, status uint8) (*types.Transaction, error) {
	return _HoQuPlatform.contract.Transact(opts, "setOfferStatus", id, status)
}

// SetOfferStatus is a paid mutator transaction binding the contract method 0x6edfb910.
//
// Solidity: function setOfferStatus(id uint32, status uint8) returns(bool)
func (_HoQuPlatform *HoQuPlatformSession) SetOfferStatus(id uint32, status uint8) (*types.Transaction, error) {
	return _HoQuPlatform.Contract.SetOfferStatus(&_HoQuPlatform.TransactOpts, id, status)
}

// SetOfferStatus is a paid mutator transaction binding the contract method 0x6edfb910.
//
// Solidity: function setOfferStatus(id uint32, status uint8) returns(bool)
func (_HoQuPlatform *HoQuPlatformTransactorSession) SetOfferStatus(id uint32, status uint8) (*types.Transaction, error) {
	return _HoQuPlatform.Contract.SetOfferStatus(&_HoQuPlatform.TransactOpts, id, status)
}

// SetTrackerStatus is a paid mutator transaction binding the contract method 0x732808d0.
//
// Solidity: function setTrackerStatus(id uint32, status uint8) returns(bool)
func (_HoQuPlatform *HoQuPlatformTransactor) SetTrackerStatus(opts *bind.TransactOpts, id uint32, status uint8) (*types.Transaction, error) {
	return _HoQuPlatform.contract.Transact(opts, "setTrackerStatus", id, status)
}

// SetTrackerStatus is a paid mutator transaction binding the contract method 0x732808d0.
//
// Solidity: function setTrackerStatus(id uint32, status uint8) returns(bool)
func (_HoQuPlatform *HoQuPlatformSession) SetTrackerStatus(id uint32, status uint8) (*types.Transaction, error) {
	return _HoQuPlatform.Contract.SetTrackerStatus(&_HoQuPlatform.TransactOpts, id, status)
}

// SetTrackerStatus is a paid mutator transaction binding the contract method 0x732808d0.
//
// Solidity: function setTrackerStatus(id uint32, status uint8) returns(bool)
func (_HoQuPlatform *HoQuPlatformTransactorSession) SetTrackerStatus(id uint32, status uint8) (*types.Transaction, error) {
	return _HoQuPlatform.Contract.SetTrackerStatus(&_HoQuPlatform.TransactOpts, id, status)
}

// SetUserStatus is a paid mutator transaction binding the contract method 0xc0b5e1ae.
//
// Solidity: function setUserStatus(id uint32, status uint8) returns(bool)
func (_HoQuPlatform *HoQuPlatformTransactor) SetUserStatus(opts *bind.TransactOpts, id uint32, status uint8) (*types.Transaction, error) {
	return _HoQuPlatform.contract.Transact(opts, "setUserStatus", id, status)
}

// SetUserStatus is a paid mutator transaction binding the contract method 0xc0b5e1ae.
//
// Solidity: function setUserStatus(id uint32, status uint8) returns(bool)
func (_HoQuPlatform *HoQuPlatformSession) SetUserStatus(id uint32, status uint8) (*types.Transaction, error) {
	return _HoQuPlatform.Contract.SetUserStatus(&_HoQuPlatform.TransactOpts, id, status)
}

// SetUserStatus is a paid mutator transaction binding the contract method 0xc0b5e1ae.
//
// Solidity: function setUserStatus(id uint32, status uint8) returns(bool)
func (_HoQuPlatform *HoQuPlatformTransactorSession) SetUserStatus(id uint32, status uint8) (*types.Transaction, error) {
	return _HoQuPlatform.Contract.SetUserStatus(&_HoQuPlatform.TransactOpts, id, status)
}

// UpdateUserPubKey is a paid mutator transaction binding the contract method 0x53378107.
//
// Solidity: function updateUserPubKey(id uint32, pubKey string) returns(bool)
func (_HoQuPlatform *HoQuPlatformTransactor) UpdateUserPubKey(opts *bind.TransactOpts, id uint32, pubKey string) (*types.Transaction, error) {
	return _HoQuPlatform.contract.Transact(opts, "updateUserPubKey", id, pubKey)
}

// UpdateUserPubKey is a paid mutator transaction binding the contract method 0x53378107.
//
// Solidity: function updateUserPubKey(id uint32, pubKey string) returns(bool)
func (_HoQuPlatform *HoQuPlatformSession) UpdateUserPubKey(id uint32, pubKey string) (*types.Transaction, error) {
	return _HoQuPlatform.Contract.UpdateUserPubKey(&_HoQuPlatform.TransactOpts, id, pubKey)
}

// UpdateUserPubKey is a paid mutator transaction binding the contract method 0x53378107.
//
// Solidity: function updateUserPubKey(id uint32, pubKey string) returns(bool)
func (_HoQuPlatform *HoQuPlatformTransactorSession) UpdateUserPubKey(id uint32, pubKey string) (*types.Transaction, error) {
	return _HoQuPlatform.Contract.UpdateUserPubKey(&_HoQuPlatform.TransactOpts, id, pubKey)
}

// HoQuPlatformConfigABI is the input ABI used to generate the binding from.
const HoQuPlatformConfigABI = "[{\"constant\":true,\"inputs\":[],\"name\":\"systemOwner\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_commission\",\"type\":\"uint256\"}],\"name\":\"setCommission\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_systemOwner\",\"type\":\"address\"}],\"name\":\"setSystemOwner\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_commissionWallet\",\"type\":\"address\"}],\"name\":\"setCommissionWallet\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"commissionWallet\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"commission\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"transferOwnership\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"name\":\"_systemOwner\",\"type\":\"address\"},{\"name\":\"_commissionWallet\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"previousOwner\",\"type\":\"address\"},{\"indexed\":false,\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"SystemOwnerChanged\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"changedBy\",\"type\":\"address\"},{\"indexed\":false,\"name\":\"commissionWallet\",\"type\":\"address\"}],\"name\":\"CommissionWalletChanged\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"changedBy\",\"type\":\"address\"},{\"indexed\":false,\"name\":\"commission\",\"type\":\"uint256\"}],\"name\":\"CommissionChanged\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"previousOwner\",\"type\":\"address\"},{\"indexed\":true,\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"OwnershipTransferred\",\"type\":\"event\"}]"

// HoQuPlatformConfigBin is the compiled bytecode used for deploying new contracts.
const HoQuPlatformConfigBin = `0x60606040526611c37937e08000600355341561001a57600080fd5b60405160408061050a833981016040528080519190602001805160008054600160a060020a03338116600160a060020a031992831617909255600180549683169682169690961790955560028054919092169416939093179092555050610484806100866000396000f30060606040526004361061008d5763ffffffff7c0100000000000000000000000000000000000000000000000000000000600035041663337792548114610092578063355e6b43146100c1578063621eb9a2146100d95780637d60b6ce146100f85780638da5cb5b1461011757806397c0262a1461012a578063e14891911461013d578063f2fde38b14610162575b600080fd5b341561009d57600080fd5b6100a5610181565b604051600160a060020a03909116815260200160405180910390f35b34156100cc57600080fd5b6100d7600435610190565b005b34156100e457600080fd5b6100d7600160a060020a0360043516610215565b341561010357600080fd5b6100d7600160a060020a03600435166102d8565b341561012257600080fd5b6100a5610399565b341561013557600080fd5b6100a56103a8565b341561014857600080fd5b6101506103b7565b60405190815260200160405180910390f35b341561016d57600080fd5b6100d7600160a060020a03600435166103bd565b600154600160a060020a031681565b60005433600160a060020a03908116911614806101bb575060015433600160a060020a039081169116145b15156101c657600080fd5b600081116101d357600080fd5b33600160a060020a03167ff659af9785c5d8c6128fdcbabb637f7d7ea787226f13cad89f71cbc579614fdc8260405190815260200160405180910390a2600355565b60005433600160a060020a0390811691161480610240575060015433600160a060020a039081169116145b151561024b57600080fd5b600160a060020a038116151561026057600080fd5b600154600160a060020a03167f1722097e017bb9a189312ac57bbdfff1bf251383cfd36c245f8ebabd8734c75982604051600160a060020a03909116815260200160405180910390a26001805473ffffffffffffffffffffffffffffffffffffffff1916600160a060020a0392909216919091179055565b60005433600160a060020a0390811691161480610303575060015433600160a060020a039081169116145b151561030e57600080fd5b600160a060020a038116151561032357600080fd5b33600160a060020a03167f394e7a1cfafd7b3c49218efbf47fb95da1832e5eaf0e9dc0428ea7220959d0e082604051600160a060020a03909116815260200160405180910390a26002805473ffffffffffffffffffffffffffffffffffffffff1916600160a060020a0392909216919091179055565b600054600160a060020a031681565b600254600160a060020a031681565b60035481565b60005433600160a060020a039081169116146103d857600080fd5b600160a060020a03811615156103ed57600080fd5b600054600160a060020a0380831691167f8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e060405160405180910390a36000805473ffffffffffffffffffffffffffffffffffffffff1916600160a060020a03929092169190911790555600a165627a7a723058203b87527eb39e7df3de05ee7616badc791f9f3d910191089454a06a2270e17c7a0029`

// DeployHoQuPlatformConfig deploys a new Ethereum contract, binding an instance of HoQuPlatformConfig to it.
func DeployHoQuPlatformConfig(auth *bind.TransactOpts, backend bind.ContractBackend, _systemOwner common.Address, _commissionWallet common.Address) (common.Address, *types.Transaction, *HoQuPlatformConfig, error) {
	parsed, err := abi.JSON(strings.NewReader(HoQuPlatformConfigABI))
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	address, tx, contract, err := bind.DeployContract(auth, parsed, common.FromHex(HoQuPlatformConfigBin), backend, _systemOwner, _commissionWallet)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &HoQuPlatformConfig{HoQuPlatformConfigCaller: HoQuPlatformConfigCaller{contract: contract}, HoQuPlatformConfigTransactor: HoQuPlatformConfigTransactor{contract: contract}}, nil
}

// HoQuPlatformConfig is an auto generated Go binding around an Ethereum contract.
type HoQuPlatformConfig struct {
	HoQuPlatformConfigCaller     // Read-only binding to the contract
	HoQuPlatformConfigTransactor // Write-only binding to the contract
}

// HoQuPlatformConfigCaller is an auto generated read-only Go binding around an Ethereum contract.
type HoQuPlatformConfigCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// HoQuPlatformConfigTransactor is an auto generated write-only Go binding around an Ethereum contract.
type HoQuPlatformConfigTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// HoQuPlatformConfigSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type HoQuPlatformConfigSession struct {
	Contract     *HoQuPlatformConfig // Generic contract binding to set the session for
	CallOpts     bind.CallOpts       // Call options to use throughout this session
	TransactOpts bind.TransactOpts   // Transaction auth options to use throughout this session
}

// HoQuPlatformConfigCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type HoQuPlatformConfigCallerSession struct {
	Contract *HoQuPlatformConfigCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts             // Call options to use throughout this session
}

// HoQuPlatformConfigTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type HoQuPlatformConfigTransactorSession struct {
	Contract     *HoQuPlatformConfigTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts             // Transaction auth options to use throughout this session
}

// HoQuPlatformConfigRaw is an auto generated low-level Go binding around an Ethereum contract.
type HoQuPlatformConfigRaw struct {
	Contract *HoQuPlatformConfig // Generic contract binding to access the raw methods on
}

// HoQuPlatformConfigCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type HoQuPlatformConfigCallerRaw struct {
	Contract *HoQuPlatformConfigCaller // Generic read-only contract binding to access the raw methods on
}

// HoQuPlatformConfigTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type HoQuPlatformConfigTransactorRaw struct {
	Contract *HoQuPlatformConfigTransactor // Generic write-only contract binding to access the raw methods on
}

// NewHoQuPlatformConfig creates a new instance of HoQuPlatformConfig, bound to a specific deployed contract.
func NewHoQuPlatformConfig(address common.Address, backend bind.ContractBackend) (*HoQuPlatformConfig, error) {
	contract, err := bindHoQuPlatformConfig(address, backend, backend)
	if err != nil {
		return nil, err
	}
	return &HoQuPlatformConfig{HoQuPlatformConfigCaller: HoQuPlatformConfigCaller{contract: contract}, HoQuPlatformConfigTransactor: HoQuPlatformConfigTransactor{contract: contract}}, nil
}

// NewHoQuPlatformConfigCaller creates a new read-only instance of HoQuPlatformConfig, bound to a specific deployed contract.
func NewHoQuPlatformConfigCaller(address common.Address, caller bind.ContractCaller) (*HoQuPlatformConfigCaller, error) {
	contract, err := bindHoQuPlatformConfig(address, caller, nil)
	if err != nil {
		return nil, err
	}
	return &HoQuPlatformConfigCaller{contract: contract}, nil
}

// NewHoQuPlatformConfigTransactor creates a new write-only instance of HoQuPlatformConfig, bound to a specific deployed contract.
func NewHoQuPlatformConfigTransactor(address common.Address, transactor bind.ContractTransactor) (*HoQuPlatformConfigTransactor, error) {
	contract, err := bindHoQuPlatformConfig(address, nil, transactor)
	if err != nil {
		return nil, err
	}
	return &HoQuPlatformConfigTransactor{contract: contract}, nil
}

// bindHoQuPlatformConfig binds a generic wrapper to an already deployed contract.
func bindHoQuPlatformConfig(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(HoQuPlatformConfigABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_HoQuPlatformConfig *HoQuPlatformConfigRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _HoQuPlatformConfig.Contract.HoQuPlatformConfigCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_HoQuPlatformConfig *HoQuPlatformConfigRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _HoQuPlatformConfig.Contract.HoQuPlatformConfigTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_HoQuPlatformConfig *HoQuPlatformConfigRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _HoQuPlatformConfig.Contract.HoQuPlatformConfigTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_HoQuPlatformConfig *HoQuPlatformConfigCallerRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _HoQuPlatformConfig.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_HoQuPlatformConfig *HoQuPlatformConfigTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _HoQuPlatformConfig.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_HoQuPlatformConfig *HoQuPlatformConfigTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _HoQuPlatformConfig.Contract.contract.Transact(opts, method, params...)
}

// Commission is a free data retrieval call binding the contract method 0xe1489191.
//
// Solidity: function commission() constant returns(uint256)
func (_HoQuPlatformConfig *HoQuPlatformConfigCaller) Commission(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _HoQuPlatformConfig.contract.Call(opts, out, "commission")
	return *ret0, err
}

// Commission is a free data retrieval call binding the contract method 0xe1489191.
//
// Solidity: function commission() constant returns(uint256)
func (_HoQuPlatformConfig *HoQuPlatformConfigSession) Commission() (*big.Int, error) {
	return _HoQuPlatformConfig.Contract.Commission(&_HoQuPlatformConfig.CallOpts)
}

// Commission is a free data retrieval call binding the contract method 0xe1489191.
//
// Solidity: function commission() constant returns(uint256)
func (_HoQuPlatformConfig *HoQuPlatformConfigCallerSession) Commission() (*big.Int, error) {
	return _HoQuPlatformConfig.Contract.Commission(&_HoQuPlatformConfig.CallOpts)
}

// CommissionWallet is a free data retrieval call binding the contract method 0x97c0262a.
//
// Solidity: function commissionWallet() constant returns(address)
func (_HoQuPlatformConfig *HoQuPlatformConfigCaller) CommissionWallet(opts *bind.CallOpts) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _HoQuPlatformConfig.contract.Call(opts, out, "commissionWallet")
	return *ret0, err
}

// CommissionWallet is a free data retrieval call binding the contract method 0x97c0262a.
//
// Solidity: function commissionWallet() constant returns(address)
func (_HoQuPlatformConfig *HoQuPlatformConfigSession) CommissionWallet() (common.Address, error) {
	return _HoQuPlatformConfig.Contract.CommissionWallet(&_HoQuPlatformConfig.CallOpts)
}

// CommissionWallet is a free data retrieval call binding the contract method 0x97c0262a.
//
// Solidity: function commissionWallet() constant returns(address)
func (_HoQuPlatformConfig *HoQuPlatformConfigCallerSession) CommissionWallet() (common.Address, error) {
	return _HoQuPlatformConfig.Contract.CommissionWallet(&_HoQuPlatformConfig.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() constant returns(address)
func (_HoQuPlatformConfig *HoQuPlatformConfigCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _HoQuPlatformConfig.contract.Call(opts, out, "owner")
	return *ret0, err
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() constant returns(address)
func (_HoQuPlatformConfig *HoQuPlatformConfigSession) Owner() (common.Address, error) {
	return _HoQuPlatformConfig.Contract.Owner(&_HoQuPlatformConfig.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() constant returns(address)
func (_HoQuPlatformConfig *HoQuPlatformConfigCallerSession) Owner() (common.Address, error) {
	return _HoQuPlatformConfig.Contract.Owner(&_HoQuPlatformConfig.CallOpts)
}

// SystemOwner is a free data retrieval call binding the contract method 0x33779254.
//
// Solidity: function systemOwner() constant returns(address)
func (_HoQuPlatformConfig *HoQuPlatformConfigCaller) SystemOwner(opts *bind.CallOpts) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _HoQuPlatformConfig.contract.Call(opts, out, "systemOwner")
	return *ret0, err
}

// SystemOwner is a free data retrieval call binding the contract method 0x33779254.
//
// Solidity: function systemOwner() constant returns(address)
func (_HoQuPlatformConfig *HoQuPlatformConfigSession) SystemOwner() (common.Address, error) {
	return _HoQuPlatformConfig.Contract.SystemOwner(&_HoQuPlatformConfig.CallOpts)
}

// SystemOwner is a free data retrieval call binding the contract method 0x33779254.
//
// Solidity: function systemOwner() constant returns(address)
func (_HoQuPlatformConfig *HoQuPlatformConfigCallerSession) SystemOwner() (common.Address, error) {
	return _HoQuPlatformConfig.Contract.SystemOwner(&_HoQuPlatformConfig.CallOpts)
}

// SetCommission is a paid mutator transaction binding the contract method 0x355e6b43.
//
// Solidity: function setCommission(_commission uint256) returns()
func (_HoQuPlatformConfig *HoQuPlatformConfigTransactor) SetCommission(opts *bind.TransactOpts, _commission *big.Int) (*types.Transaction, error) {
	return _HoQuPlatformConfig.contract.Transact(opts, "setCommission", _commission)
}

// SetCommission is a paid mutator transaction binding the contract method 0x355e6b43.
//
// Solidity: function setCommission(_commission uint256) returns()
func (_HoQuPlatformConfig *HoQuPlatformConfigSession) SetCommission(_commission *big.Int) (*types.Transaction, error) {
	return _HoQuPlatformConfig.Contract.SetCommission(&_HoQuPlatformConfig.TransactOpts, _commission)
}

// SetCommission is a paid mutator transaction binding the contract method 0x355e6b43.
//
// Solidity: function setCommission(_commission uint256) returns()
func (_HoQuPlatformConfig *HoQuPlatformConfigTransactorSession) SetCommission(_commission *big.Int) (*types.Transaction, error) {
	return _HoQuPlatformConfig.Contract.SetCommission(&_HoQuPlatformConfig.TransactOpts, _commission)
}

// SetCommissionWallet is a paid mutator transaction binding the contract method 0x7d60b6ce.
//
// Solidity: function setCommissionWallet(_commissionWallet address) returns()
func (_HoQuPlatformConfig *HoQuPlatformConfigTransactor) SetCommissionWallet(opts *bind.TransactOpts, _commissionWallet common.Address) (*types.Transaction, error) {
	return _HoQuPlatformConfig.contract.Transact(opts, "setCommissionWallet", _commissionWallet)
}

// SetCommissionWallet is a paid mutator transaction binding the contract method 0x7d60b6ce.
//
// Solidity: function setCommissionWallet(_commissionWallet address) returns()
func (_HoQuPlatformConfig *HoQuPlatformConfigSession) SetCommissionWallet(_commissionWallet common.Address) (*types.Transaction, error) {
	return _HoQuPlatformConfig.Contract.SetCommissionWallet(&_HoQuPlatformConfig.TransactOpts, _commissionWallet)
}

// SetCommissionWallet is a paid mutator transaction binding the contract method 0x7d60b6ce.
//
// Solidity: function setCommissionWallet(_commissionWallet address) returns()
func (_HoQuPlatformConfig *HoQuPlatformConfigTransactorSession) SetCommissionWallet(_commissionWallet common.Address) (*types.Transaction, error) {
	return _HoQuPlatformConfig.Contract.SetCommissionWallet(&_HoQuPlatformConfig.TransactOpts, _commissionWallet)
}

// SetSystemOwner is a paid mutator transaction binding the contract method 0x621eb9a2.
//
// Solidity: function setSystemOwner(_systemOwner address) returns()
func (_HoQuPlatformConfig *HoQuPlatformConfigTransactor) SetSystemOwner(opts *bind.TransactOpts, _systemOwner common.Address) (*types.Transaction, error) {
	return _HoQuPlatformConfig.contract.Transact(opts, "setSystemOwner", _systemOwner)
}

// SetSystemOwner is a paid mutator transaction binding the contract method 0x621eb9a2.
//
// Solidity: function setSystemOwner(_systemOwner address) returns()
func (_HoQuPlatformConfig *HoQuPlatformConfigSession) SetSystemOwner(_systemOwner common.Address) (*types.Transaction, error) {
	return _HoQuPlatformConfig.Contract.SetSystemOwner(&_HoQuPlatformConfig.TransactOpts, _systemOwner)
}

// SetSystemOwner is a paid mutator transaction binding the contract method 0x621eb9a2.
//
// Solidity: function setSystemOwner(_systemOwner address) returns()
func (_HoQuPlatformConfig *HoQuPlatformConfigTransactorSession) SetSystemOwner(_systemOwner common.Address) (*types.Transaction, error) {
	return _HoQuPlatformConfig.Contract.SetSystemOwner(&_HoQuPlatformConfig.TransactOpts, _systemOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(newOwner address) returns()
func (_HoQuPlatformConfig *HoQuPlatformConfigTransactor) TransferOwnership(opts *bind.TransactOpts, newOwner common.Address) (*types.Transaction, error) {
	return _HoQuPlatformConfig.contract.Transact(opts, "transferOwnership", newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(newOwner address) returns()
func (_HoQuPlatformConfig *HoQuPlatformConfigSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _HoQuPlatformConfig.Contract.TransferOwnership(&_HoQuPlatformConfig.TransactOpts, newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(newOwner address) returns()
func (_HoQuPlatformConfig *HoQuPlatformConfigTransactorSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _HoQuPlatformConfig.Contract.TransferOwnership(&_HoQuPlatformConfig.TransactOpts, newOwner)
}

// HoQuTokenABI is the input ABI used to generate the binding from.
const HoQuTokenABI = "[{\"constant\":true,\"inputs\":[],\"name\":\"name\",\"outputs\":[{\"name\":\"\",\"type\":\"string\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_spender\",\"type\":\"address\"},{\"name\":\"_value\",\"type\":\"uint256\"}],\"name\":\"approve\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"totalSupply\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_from\",\"type\":\"address\"},{\"name\":\"_to\",\"type\":\"address\"},{\"name\":\"_value\",\"type\":\"uint256\"}],\"name\":\"transferFrom\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"decimals\",\"outputs\":[{\"name\":\"\",\"type\":\"uint32\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[],\"name\":\"unpause\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"paused\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_spender\",\"type\":\"address\"},{\"name\":\"_subtractedValue\",\"type\":\"uint256\"}],\"name\":\"decreaseApproval\",\"outputs\":[{\"name\":\"success\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"_owner\",\"type\":\"address\"}],\"name\":\"balanceOf\",\"outputs\":[{\"name\":\"balance\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[],\"name\":\"pause\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"symbol\",\"outputs\":[{\"name\":\"\",\"type\":\"string\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_to\",\"type\":\"address\"},{\"name\":\"_value\",\"type\":\"uint256\"}],\"name\":\"transfer\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_spender\",\"type\":\"address\"},{\"name\":\"_addedValue\",\"type\":\"uint256\"}],\"name\":\"increaseApproval\",\"outputs\":[{\"name\":\"success\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"_owner\",\"type\":\"address\"},{\"name\":\"_spender\",\"type\":\"address\"}],\"name\":\"allowance\",\"outputs\":[{\"name\":\"remaining\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"transferOwnership\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"name\":\"_totalSupply\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[],\"name\":\"Pause\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[],\"name\":\"Unpause\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"previousOwner\",\"type\":\"address\"},{\"indexed\":true,\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"OwnershipTransferred\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"owner\",\"type\":\"address\"},{\"indexed\":true,\"name\":\"spender\",\"type\":\"address\"},{\"indexed\":false,\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"Approval\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"from\",\"type\":\"address\"},{\"indexed\":true,\"name\":\"to\",\"type\":\"address\"},{\"indexed\":false,\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"Transfer\",\"type\":\"event\"}]"

// HoQuTokenBin is the compiled bytecode used for deploying new contracts.
const HoQuTokenBin = `0x60606040526003805460a060020a60ff0219169055341561001f57600080fd5b604051602080610b998339810160405280805160038054600160a060020a03191633600160a060020a03161790559150506000811161005d57600080fd5b600160a060020a033316600090815260016020526040812082905555610b11806100886000396000f3006060604052600436106100e55763ffffffff7c010000000000000000000000000000000000000000000000000000000060003504166306fdde0381146100ea578063095ea7b31461017457806318160ddd146101aa57806323b872dd146101cf578063313ce567146101f75780633f4ba83a146102235780635c975abb14610238578063661884631461024b57806370a082311461026d5780638456cb591461028c5780638da5cb5b1461029f57806395d89b41146102ce578063a9059cbb146102e1578063d73dd62314610303578063dd62ed3e14610325578063f2fde38b1461034a575b600080fd5b34156100f557600080fd5b6100fd610369565b60405160208082528190810183818151815260200191508051906020019080838360005b83811015610139578082015183820152602001610121565b50505050905090810190601f1680156101665780820380516001836020036101000a031916815260200191505b509250505060405180910390f35b341561017f57600080fd5b610196600160a060020a03600435166024356103a0565b604051901515815260200160405180910390f35b34156101b557600080fd5b6101bd61040c565b60405190815260200160405180910390f35b34156101da57600080fd5b610196600160a060020a0360043581169060243516604435610412565b341561020257600080fd5b61020a61043f565b60405163ffffffff909116815260200160405180910390f35b341561022e57600080fd5b610236610444565b005b341561024357600080fd5b6101966104c3565b341561025657600080fd5b610196600160a060020a03600435166024356104d3565b341561027857600080fd5b6101bd600160a060020a03600435166105cd565b341561029757600080fd5b6102366105e8565b34156102aa57600080fd5b6102b261066c565b604051600160a060020a03909116815260200160405180910390f35b34156102d957600080fd5b6100fd61067b565b34156102ec57600080fd5b610196600160a060020a03600435166024356106b2565b341561030e57600080fd5b610196600160a060020a03600435166024356106dd565b341561033057600080fd5b6101bd600160a060020a0360043581169060243516610781565b341561035557600080fd5b610236600160a060020a03600435166107ac565b60408051908101604052600a81527f484f515520546f6b656e00000000000000000000000000000000000000000000602082015281565b600160a060020a03338116600081815260026020908152604080832094871680845294909152808220859055909291907f8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b9259085905190815260200160405180910390a350600192915050565b60005481565b60035460009060a060020a900460ff161561042c57600080fd5b610437848484610847565b949350505050565b601281565b60035433600160a060020a0390811691161461045f57600080fd5b60035460a060020a900460ff16151561047757600080fd5b6003805474ff0000000000000000000000000000000000000000191690557f7805862f689e2f13df9f062ff482ad3ad112aca9e0847911ed832e158c525b3360405160405180910390a1565b60035460a060020a900460ff1681565b600160a060020a0333811660009081526002602090815260408083209386168352929052908120548083111561053057600160a060020a033381166000908152600260209081526040808320938816835292905290812055610567565b610540818463ffffffff6109c916565b600160a060020a033381166000908152600260209081526040808320938916835292905220555b600160a060020a0333811660008181526002602090815260408083209489168084529490915290819020547f8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925915190815260200160405180910390a35060019392505050565b600160a060020a031660009081526001602052604090205490565b60035433600160a060020a0390811691161461060357600080fd5b60035460a060020a900460ff161561061a57600080fd5b6003805474ff0000000000000000000000000000000000000000191660a060020a1790557f6985a02210a168e66602d3235cb6db0e70f92b3ba4d376a33c0f3d9434bff62560405160405180910390a1565b600354600160a060020a031681565b60408051908101604052600381527f4851580000000000000000000000000000000000000000000000000000000000602082015281565b60035460009060a060020a900460ff16156106cc57600080fd5b6106d683836109db565b9392505050565b600160a060020a033381166000908152600260209081526040808320938616835292905290812054610715908363ffffffff610ad616565b600160a060020a0333811660008181526002602090815260408083209489168084529490915290819020849055919290917f8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b92591905190815260200160405180910390a350600192915050565b600160a060020a03918216600090815260026020908152604080832093909416825291909152205490565b60035433600160a060020a039081169116146107c757600080fd5b600160a060020a03811615156107dc57600080fd5b600354600160a060020a0380831691167f8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e060405160405180910390a36003805473ffffffffffffffffffffffffffffffffffffffff1916600160a060020a0392909216919091179055565b6000600160a060020a038316151561085e57600080fd5b600160a060020a03841660009081526001602052604090205482111561088357600080fd5b600160a060020a03808516600090815260026020908152604080832033909416835292905220548211156108b657600080fd5b600160a060020a0384166000908152600160205260409020546108df908363ffffffff6109c916565b600160a060020a038086166000908152600160205260408082209390935590851681522054610914908363ffffffff610ad616565b600160a060020a0380851660009081526001602090815260408083209490945587831682526002815283822033909316825291909152205461095c908363ffffffff6109c916565b600160a060020a03808616600081815260026020908152604080832033861684529091529081902093909355908516917fddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef9085905190815260200160405180910390a35060019392505050565b6000828211156109d557fe5b50900390565b6000600160a060020a03831615156109f257600080fd5b600160a060020a033316600090815260016020526040902054821115610a1757600080fd5b600160a060020a033316600090815260016020526040902054610a40908363ffffffff6109c916565b600160a060020a033381166000908152600160205260408082209390935590851681522054610a75908363ffffffff610ad616565b600160a060020a0380851660008181526001602052604090819020939093559133909116907fddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef9085905190815260200160405180910390a350600192915050565b6000828201838110156106d657fe00a165627a7a72305820490b105f18036764a3f4dd5962fa2231f0bf8b00e36eb7ae5222ef96fc8aa4160029`

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
const OwnableBin = `0x6060604052341561000f57600080fd5b60008054600160a060020a033316600160a060020a03199091161790556101768061003b6000396000f30060606040526004361061004b5763ffffffff7c01000000000000000000000000000000000000000000000000000000006000350416638da5cb5b8114610050578063f2fde38b1461007f575b600080fd5b341561005b57600080fd5b6100636100a0565b604051600160a060020a03909116815260200160405180910390f35b341561008a57600080fd5b61009e600160a060020a03600435166100af565b005b600054600160a060020a031681565b60005433600160a060020a039081169116146100ca57600080fd5b600160a060020a03811615156100df57600080fd5b600054600160a060020a0380831691167f8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e060405160405180910390a36000805473ffffffffffffffffffffffffffffffffffffffff1916600160a060020a03929092169190911790555600a165627a7a72305820657c8479493b7ea3d0f380c21cc411d601b750afe9d568a3fb4e27c7b2f709540029`

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
const PausableBin = `0x606060405260008054600160a060020a033316600160a860020a03199091161790556102f7806100306000396000f30060606040526004361061006c5763ffffffff7c01000000000000000000000000000000000000000000000000000000006000350416633f4ba83a81146100715780635c975abb146100865780638456cb59146100ad5780638da5cb5b146100c0578063f2fde38b146100ef575b600080fd5b341561007c57600080fd5b61008461010e565b005b341561009157600080fd5b61009961018d565b604051901515815260200160405180910390f35b34156100b857600080fd5b61008461019d565b34156100cb57600080fd5b6100d3610221565b604051600160a060020a03909116815260200160405180910390f35b34156100fa57600080fd5b610084600160a060020a0360043516610230565b60005433600160a060020a0390811691161461012957600080fd5b60005460a060020a900460ff16151561014157600080fd5b6000805474ff0000000000000000000000000000000000000000191690557f7805862f689e2f13df9f062ff482ad3ad112aca9e0847911ed832e158c525b3360405160405180910390a1565b60005460a060020a900460ff1681565b60005433600160a060020a039081169116146101b857600080fd5b60005460a060020a900460ff16156101cf57600080fd5b6000805474ff0000000000000000000000000000000000000000191660a060020a1790557f6985a02210a168e66602d3235cb6db0e70f92b3ba4d376a33c0f3d9434bff62560405160405180910390a1565b600054600160a060020a031681565b60005433600160a060020a0390811691161461024b57600080fd5b600160a060020a038116151561026057600080fd5b600054600160a060020a0380831691167f8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e060405160405180910390a36000805473ffffffffffffffffffffffffffffffffffffffff1916600160a060020a03929092169190911790555600a165627a7a723058202d512f32af36adf94792fe507441839ffe94795ca1b124f68f207b609361fb9e0029`

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
const SafeMathBin = `0x60606040523415600e57600080fd5b603580601b6000396000f3006060604052600080fd00a165627a7a7230582031dcc88147137cdb69a2ccf52939c5b94fe671c6c6024fe05aff2c021658bb140029`

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
const StandardTokenBin = `0x6060604052341561000f57600080fd5b6106e68061001e6000396000f30060606040526004361061008d5763ffffffff7c0100000000000000000000000000000000000000000000000000000000600035041663095ea7b3811461009257806318160ddd146100c857806323b872dd146100ed578063661884631461011557806370a0823114610137578063a9059cbb14610156578063d73dd62314610178578063dd62ed3e1461019a575b600080fd5b341561009d57600080fd5b6100b4600160a060020a03600435166024356101bf565b604051901515815260200160405180910390f35b34156100d357600080fd5b6100db61022b565b60405190815260200160405180910390f35b34156100f857600080fd5b6100b4600160a060020a0360043581169060243516604435610231565b341561012057600080fd5b6100b4600160a060020a03600435166024356103b3565b341561014257600080fd5b6100db600160a060020a03600435166104ad565b341561016157600080fd5b6100b4600160a060020a03600435166024356104c8565b341561018357600080fd5b6100b4600160a060020a03600435166024356105c3565b34156101a557600080fd5b6100db600160a060020a0360043581169060243516610667565b600160a060020a03338116600081815260026020908152604080832094871680845294909152808220859055909291907f8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b9259085905190815260200160405180910390a350600192915050565b60005481565b6000600160a060020a038316151561024857600080fd5b600160a060020a03841660009081526001602052604090205482111561026d57600080fd5b600160a060020a03808516600090815260026020908152604080832033909416835292905220548211156102a057600080fd5b600160a060020a0384166000908152600160205260409020546102c9908363ffffffff61069216565b600160a060020a0380861660009081526001602052604080822093909355908516815220546102fe908363ffffffff6106a416565b600160a060020a03808516600090815260016020908152604080832094909455878316825260028152838220339093168252919091522054610346908363ffffffff61069216565b600160a060020a03808616600081815260026020908152604080832033861684529091529081902093909355908516917fddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef9085905190815260200160405180910390a35060019392505050565b600160a060020a0333811660009081526002602090815260408083209386168352929052908120548083111561041057600160a060020a033381166000908152600260209081526040808320938816835292905290812055610447565b610420818463ffffffff61069216565b600160a060020a033381166000908152600260209081526040808320938916835292905220555b600160a060020a0333811660008181526002602090815260408083209489168084529490915290819020547f8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925915190815260200160405180910390a35060019392505050565b600160a060020a031660009081526001602052604090205490565b6000600160a060020a03831615156104df57600080fd5b600160a060020a03331660009081526001602052604090205482111561050457600080fd5b600160a060020a03331660009081526001602052604090205461052d908363ffffffff61069216565b600160a060020a033381166000908152600160205260408082209390935590851681522054610562908363ffffffff6106a416565b600160a060020a0380851660008181526001602052604090819020939093559133909116907fddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef9085905190815260200160405180910390a350600192915050565b600160a060020a0333811660009081526002602090815260408083209386168352929052908120546105fb908363ffffffff6106a416565b600160a060020a0333811660008181526002602090815260408083209489168084529490915290819020849055919290917f8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b92591905190815260200160405180910390a350600192915050565b600160a060020a03918216600090815260026020908152604080832093909416825291909152205490565b60008282111561069e57fe5b50900390565b6000828201838110156106b357fe5b93925050505600a165627a7a7230582043933a435a4f3343077f322830ade151cde8716e0d1cf59946a06d10325f612e0029`

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
