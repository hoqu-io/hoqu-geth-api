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
const BasicTokenBin = `0x6060604052341561000f57600080fd5b6102458061001e6000396000f3006060604052600436106100565763ffffffff7c010000000000000000000000000000000000000000000000000000000060003504166318160ddd811461005b57806370a0823114610080578063a9059cbb1461009f575b600080fd5b341561006657600080fd5b61006e6100d5565b60405190815260200160405180910390f35b341561008b57600080fd5b61006e600160a060020a03600435166100db565b34156100aa57600080fd5b6100c1600160a060020a03600435166024356100f6565b604051901515815260200160405180910390f35b60005481565b600160a060020a031660009081526001602052604090205490565b6000600160a060020a038316151561010d57600080fd5b600160a060020a03331660009081526001602052604090205482111561013257600080fd5b600160a060020a03331660009081526001602052604090205461015b908363ffffffff6101f116565b600160a060020a033381166000908152600160205260408082209390935590851681522054610190908363ffffffff61020316565b600160a060020a0380851660008181526001602052604090819020939093559133909116907fddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef9085905190815260200160405180910390a350600192915050565b6000828211156101fd57fe5b50900390565b60008282018381101561021257fe5b93925050505600a165627a7a723058208cbbac6dd14122854f52638fecbb3c1aa45d6de73828ca11d0da93da7fb839660029`

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
const HoQuPlatformABI = "[{\"constant\":false,\"inputs\":[{\"name\":\"id\",\"type\":\"bytes16\"},{\"name\":\"ownerAddress\",\"type\":\"address\"}],\"name\":\"addUserAddress\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"id\",\"type\":\"bytes16\"},{\"name\":\"ownerId\",\"type\":\"bytes16\"},{\"name\":\"offerId\",\"type\":\"bytes16\"},{\"name\":\"beneficiaryAddress\",\"type\":\"address\"}],\"name\":\"addAd\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"id\",\"type\":\"bytes16\"},{\"name\":\"meta\",\"type\":\"string\"},{\"name\":\"kycLevel\",\"type\":\"uint8\"},{\"name\":\"dataUrl\",\"type\":\"string\"}],\"name\":\"addUserKycReport\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"id\",\"type\":\"bytes16\"},{\"name\":\"status\",\"type\":\"uint8\"}],\"name\":\"setAdStatus\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"id\",\"type\":\"bytes16\"},{\"name\":\"intermediaryAddress\",\"type\":\"address\"},{\"name\":\"percent\",\"type\":\"uint32\"}],\"name\":\"addLeadIntermediary\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"id\",\"type\":\"bytes16\"},{\"name\":\"status\",\"type\":\"uint8\"}],\"name\":\"setOfferStatus\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"id\",\"type\":\"bytes16\"},{\"name\":\"ownerAddress\",\"type\":\"address\"},{\"name\":\"name\",\"type\":\"string\"},{\"name\":\"dataUrl\",\"type\":\"string\"}],\"name\":\"registerTracker\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"id\",\"type\":\"bytes16\"}],\"name\":\"sellLead\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"\",\"type\":\"bytes16\"}],\"name\":\"companies\",\"outputs\":[{\"name\":\"createdAt\",\"type\":\"uint256\"},{\"name\":\"ownerId\",\"type\":\"bytes16\"},{\"name\":\"ownerAddress\",\"type\":\"address\"},{\"name\":\"name\",\"type\":\"string\"},{\"name\":\"dataUrl\",\"type\":\"string\"},{\"name\":\"status\",\"type\":\"uint8\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"id\",\"type\":\"bytes16\"},{\"name\":\"num\",\"type\":\"uint8\"}],\"name\":\"getLeadIntermediaryPercent\",\"outputs\":[{\"name\":\"\",\"type\":\"uint32\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"\",\"type\":\"bytes16\"}],\"name\":\"users\",\"outputs\":[{\"name\":\"createdAt\",\"type\":\"uint256\"},{\"name\":\"numOfAddresses\",\"type\":\"uint8\"},{\"name\":\"role\",\"type\":\"string\"},{\"name\":\"kycLevel\",\"type\":\"uint8\"},{\"name\":\"numOfKycReports\",\"type\":\"uint16\"},{\"name\":\"pubKey\",\"type\":\"string\"},{\"name\":\"status\",\"type\":\"uint8\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"config\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"id\",\"type\":\"bytes16\"},{\"name\":\"status\",\"type\":\"uint8\"}],\"name\":\"setLeadStatus\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"configAddress\",\"type\":\"address\"}],\"name\":\"setConfigAddress\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"id\",\"type\":\"bytes16\"},{\"name\":\"pubKey\",\"type\":\"string\"}],\"name\":\"updateUserPubKey\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"id\",\"type\":\"bytes16\"},{\"name\":\"adId\",\"type\":\"bytes16\"},{\"name\":\"trackerId\",\"type\":\"bytes16\"},{\"name\":\"meta\",\"type\":\"string\"},{\"name\":\"dataUrl\",\"type\":\"string\"},{\"name\":\"price\",\"type\":\"uint256\"}],\"name\":\"addLead\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"id\",\"type\":\"bytes16\"},{\"name\":\"status\",\"type\":\"uint8\"}],\"name\":\"setCompanyStatus\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"id\",\"type\":\"bytes16\"},{\"name\":\"ownerId\",\"type\":\"bytes16\"},{\"name\":\"ownerAddress\",\"type\":\"address\"},{\"name\":\"name\",\"type\":\"string\"},{\"name\":\"dataUrl\",\"type\":\"string\"}],\"name\":\"registerCompany\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"\",\"type\":\"bytes16\"}],\"name\":\"offers\",\"outputs\":[{\"name\":\"createdAt\",\"type\":\"uint256\"},{\"name\":\"companyId\",\"type\":\"bytes16\"},{\"name\":\"payerAddress\",\"type\":\"address\"},{\"name\":\"name\",\"type\":\"string\"},{\"name\":\"dataUrl\",\"type\":\"string\"},{\"name\":\"cost\",\"type\":\"uint256\"},{\"name\":\"status\",\"type\":\"uint8\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"id\",\"type\":\"bytes16\"},{\"name\":\"num\",\"type\":\"uint8\"}],\"name\":\"getUserAddress\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"\",\"type\":\"bytes16\"}],\"name\":\"ads\",\"outputs\":[{\"name\":\"createdAt\",\"type\":\"uint256\"},{\"name\":\"ownerId\",\"type\":\"bytes16\"},{\"name\":\"beneficiaryAddress\",\"type\":\"address\"},{\"name\":\"offerId\",\"type\":\"bytes16\"},{\"name\":\"status\",\"type\":\"uint8\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"id\",\"type\":\"bytes16\"},{\"name\":\"num\",\"type\":\"uint8\"}],\"name\":\"getLeadIntermediaryAddress\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"id\",\"type\":\"bytes16\"},{\"name\":\"companyId\",\"type\":\"bytes16\"},{\"name\":\"payerAddress\",\"type\":\"address\"},{\"name\":\"name\",\"type\":\"string\"},{\"name\":\"dataUrl\",\"type\":\"string\"},{\"name\":\"cost\",\"type\":\"uint256\"}],\"name\":\"addOffer\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"id\",\"type\":\"bytes16\"},{\"name\":\"status\",\"type\":\"uint8\"}],\"name\":\"setUserStatus\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"\",\"type\":\"bytes16\"}],\"name\":\"leads\",\"outputs\":[{\"name\":\"createdAt\",\"type\":\"uint256\"},{\"name\":\"trackerId\",\"type\":\"bytes16\"},{\"name\":\"adId\",\"type\":\"bytes16\"},{\"name\":\"dataUrl\",\"type\":\"string\"},{\"name\":\"meta\",\"type\":\"string\"},{\"name\":\"price\",\"type\":\"uint256\"},{\"name\":\"numOfIntermediaries\",\"type\":\"uint8\"},{\"name\":\"status\",\"type\":\"uint8\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"id\",\"type\":\"bytes16\"},{\"name\":\"status\",\"type\":\"uint8\"}],\"name\":\"setTrackerStatus\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"\",\"type\":\"bytes16\"}],\"name\":\"trackers\",\"outputs\":[{\"name\":\"createdAt\",\"type\":\"uint256\"},{\"name\":\"ownerAddress\",\"type\":\"address\"},{\"name\":\"name\",\"type\":\"string\"},{\"name\":\"dataUrl\",\"type\":\"string\"},{\"name\":\"status\",\"type\":\"uint8\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"id\",\"type\":\"bytes16\"},{\"name\":\"num\",\"type\":\"uint16\"}],\"name\":\"getUserKycReport\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"},{\"name\":\"\",\"type\":\"string\"},{\"name\":\"\",\"type\":\"uint8\"},{\"name\":\"\",\"type\":\"string\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"token\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"id\",\"type\":\"bytes16\"},{\"name\":\"role\",\"type\":\"string\"},{\"name\":\"ownerAddress\",\"type\":\"address\"},{\"name\":\"pubKey\",\"type\":\"string\"}],\"name\":\"registerUser\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"name\":\"configAddress\",\"type\":\"address\"},{\"name\":\"tokenAddress\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"ownerAddress\",\"type\":\"address\"},{\"indexed\":false,\"name\":\"id\",\"type\":\"bytes16\"},{\"indexed\":false,\"name\":\"role\",\"type\":\"string\"}],\"name\":\"UserRegistered\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"ownerAddress\",\"type\":\"address\"},{\"indexed\":false,\"name\":\"additionalAddress\",\"type\":\"address\"}],\"name\":\"UserAddressAdded\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"ownerAddress\",\"type\":\"address\"}],\"name\":\"UserPubKeyUpdated\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"ownerAddress\",\"type\":\"address\"},{\"indexed\":false,\"name\":\"kycLevel\",\"type\":\"uint8\"}],\"name\":\"UserKycReportAdded\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"ownerAddress\",\"type\":\"address\"},{\"indexed\":false,\"name\":\"id\",\"type\":\"bytes16\"},{\"indexed\":false,\"name\":\"status\",\"type\":\"uint8\"}],\"name\":\"UserStatusChanged\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"ownerAddress\",\"type\":\"address\"},{\"indexed\":false,\"name\":\"id\",\"type\":\"bytes16\"},{\"indexed\":false,\"name\":\"name\",\"type\":\"string\"}],\"name\":\"CompanyRegistered\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"ownerAddress\",\"type\":\"address\"},{\"indexed\":false,\"name\":\"id\",\"type\":\"bytes16\"},{\"indexed\":false,\"name\":\"status\",\"type\":\"uint8\"}],\"name\":\"CompanyStatusChanged\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"ownerAddress\",\"type\":\"address\"},{\"indexed\":false,\"name\":\"id\",\"type\":\"bytes16\"},{\"indexed\":false,\"name\":\"name\",\"type\":\"string\"}],\"name\":\"TrackerRegistered\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"ownerAddress\",\"type\":\"address\"},{\"indexed\":false,\"name\":\"id\",\"type\":\"bytes16\"},{\"indexed\":false,\"name\":\"status\",\"type\":\"uint8\"}],\"name\":\"TrackerStatusChanged\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"ownerAddress\",\"type\":\"address\"},{\"indexed\":false,\"name\":\"id\",\"type\":\"bytes16\"},{\"indexed\":false,\"name\":\"name\",\"type\":\"string\"}],\"name\":\"OfferAdded\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"ownerAddress\",\"type\":\"address\"},{\"indexed\":false,\"name\":\"id\",\"type\":\"bytes16\"},{\"indexed\":false,\"name\":\"status\",\"type\":\"uint8\"}],\"name\":\"OfferStatusChanged\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"ownerAddress\",\"type\":\"address\"},{\"indexed\":false,\"name\":\"id\",\"type\":\"bytes16\"}],\"name\":\"AdAdded\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"ownerAddress\",\"type\":\"address\"},{\"indexed\":false,\"name\":\"id\",\"type\":\"bytes16\"},{\"indexed\":false,\"name\":\"status\",\"type\":\"uint8\"}],\"name\":\"AdStatusChanged\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"ownerAddress\",\"type\":\"address\"},{\"indexed\":false,\"name\":\"id\",\"type\":\"bytes16\"},{\"indexed\":false,\"name\":\"price\",\"type\":\"uint256\"}],\"name\":\"LeadAdded\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"ownerAddress\",\"type\":\"address\"},{\"indexed\":false,\"name\":\"id\",\"type\":\"bytes16\"},{\"indexed\":false,\"name\":\"status\",\"type\":\"uint8\"}],\"name\":\"LeadStatusChanged\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"ownerAddress\",\"type\":\"address\"},{\"indexed\":false,\"name\":\"id\",\"type\":\"bytes16\"},{\"indexed\":false,\"name\":\"ownerAmount\",\"type\":\"uint256\"}],\"name\":\"LeadSold\",\"type\":\"event\"}]"

// HoQuPlatformBin is the compiled bytecode used for deploying new contracts.
const HoQuPlatformBin = `0x6060604052341561000f57600080fd5b604051604080613d47833981016040528080519190602001805160008054600160a060020a03958616600160a060020a03199182161790915560018054959092169416939093179092555050613cdd8061006a6000396000f3006060604052600436106101665763ffffffff60e060020a60003504166303806028811461016b57806306501a10146101ab57806311ab8673146101e35780631bf524381461029157806336617045146102b757806336d7523f146102ec57806337e830ba1461031257806347f3d978146103c35780634df2d847146103e35780635cb8f5ab146105475780637520dd141461058657806379502c55146106f457806381e85aa31461072357806383a12de9146107495780638a5463bd1461076a5780638f2e960d146107ca578063a8800f091461087d578063be95f1c2146108a3578063c32d869b1461095b578063c49e2fb114610a86578063c85a758514610aac578063da7345b714610b21578063df7e40bb14610b47578063e03d8b6414610c01578063e7608fb214610c27578063e7e5aad814610d96578063ea2fdd8c14610dbc578063ee26342214610f0f578063fc0c546a14611033578063fd3fc08f14611046575b600080fd5b341561017657600080fd5b6101976001608060020a031960043516600160a060020a03602435166110fa565b604051901515815260200160405180910390f35b34156101b657600080fd5b6101976001608060020a031960043581169060243581169060443516600160a060020a036064351661126c565b34156101ee57600080fd5b610197600480356001608060020a0319169060446024803590810190830135806020601f8201819004810201604051908101604052818152929190602084018383808284378201915050505050509190803560ff1690602001909190803590602001908201803590602001908080601f0160208091040260200160405190810160405281815292919060208401838380828437509496506114ee95505050505050565b341561029c57600080fd5b6101976001608060020a03196004351660ff60243516611739565b34156102c257600080fd5b6101976001608060020a031960043516600160a060020a036024351663ffffffff604435166118cf565b34156102f757600080fd5b6101976001608060020a03196004351660ff60243516611a20565b341561031d57600080fd5b610197600480356001608060020a0319169060248035600160a060020a0316919060649060443590810190830135806020601f8201819004810201604051908101604052818152929190602084018383808284378201915050505050509190803590602001908201803590602001908080601f016020809104026020016040519081016040528181529291906020840183838082843750949650611b7d95505050505050565b34156103ce57600080fd5b6101976001608060020a031960043516611dc6565b34156103ee57600080fd5b6104036001608060020a03196004351661237f565b6040518681526001608060020a031986166020820152600160a060020a0385166040820152606081016080820160a0830184600581111561044057fe5b60ff168152602084820381018452875460026001821615610100026000190190911604908201819052604090910190879080156104be5780601f10610493576101008083540402835291602001916104be565b820191906000526020600020905b8154815290600101906020018083116104a157829003601f168201915b50508381038252855460026000196101006001841615020190911604808252602090910190869080156105325780601f1061050757610100808354040283529160200191610532565b820191906000526020600020905b81548152906001019060200180831161051557829003601f168201915b50509850505050505050505060405180910390f35b341561055257600080fd5b61056d6001608060020a03196004351660ff602435166123c8565b60405163ffffffff909116815260200160405180910390f35b341561059157600080fd5b6105a66001608060020a03196004351661243f565b60405187815260ff8716602082015260408101606082018660048111156105c957fe5b60ff16815261ffff8616602082015260408101906060018460058111156105ec57fe5b60ff1681526020848203810184528954600260018216156101000260001901909116049082018190526040909101908990801561066a5780601f1061063f5761010080835404028352916020019161066a565b820191906000526020600020905b81548152906001019060200180831161064d57829003601f168201915b50508381038252855460026000196101006001841615020190911604808252602090910190869080156106de5780601f106106b3576101008083540402835291602001916106de565b820191906000526020600020905b8154815290600101906020018083116106c157829003601f168201915b5050995050505050505050505060405180910390f35b34156106ff57600080fd5b610707612485565b604051600160a060020a03909116815260200160405180910390f35b341561072e57600080fd5b6101976001608060020a03196004351660ff60243516612494565b341561075457600080fd5b610768600160a060020a0360043516612615565b005b341561077557600080fd5b610197600480356001608060020a0319169060446024803590810190830135806020601f820181900481020160405190810160405281815292919060208401838380828437509496506126ba95505050505050565b34156107d557600080fd5b610197600480356001608060020a0319908116916024803583169260443516919060849060643590810190830135806020601f8201819004810201604051908101604052818152929190602084018383808284378201915050505050509190803590602001908201803590602001908080601f016020809104026020016040519081016040528181529291906020840183838082843750949650509335935061280892505050565b341561088857600080fd5b6101976001608060020a03196004351660ff60243516612af7565b34156108ae57600080fd5b6101976001608060020a0319600480358216916024803590911691600160a060020a03604435169160849060643590810190830135806020601f8201819004810201604051908101604052818152929190602084018383808284378201915050505050509190803590602001908201803590602001908080601f016020809104026020016040519081016040528181529291906020840183838082843750949650612c5995505050505050565b341561096657600080fd5b61097b6001608060020a031960043516612fa4565b6040518781526001608060020a031987166020820152600160a060020a038616604082015260a08101839052606081016080820160c083018460058111156109bf57fe5b60ff16815260208482038101845288546002600182161561010002600019019091160490820181905260409091019088908015610a3d5780601f10610a1257610100808354040283529160200191610a3d565b820191906000526020600020905b815481529060010190602001808311610a2057829003601f168201915b50508381038252865460026000196101006001841615020190911604808252602090910190879080156106de5780601f106106b3576101008083540402835291602001916106de565b3415610a9157600080fd5b6107076001608060020a03196004351660ff60243516612ff2565b3415610ab757600080fd5b610acc6001608060020a031960043516613066565b6040518581526001608060020a03198086166020830152600160a060020a03851660408301528316606082015260808101826005811115610b0957fe5b60ff1681526020019550505050505060405180910390f35b3415610b2c57600080fd5b6107076001608060020a03196004351660ff602435166130a8565b3415610b5257600080fd5b6101976001608060020a0319600480358216916024803590911691600160a060020a03604435169160849060643590810190830135806020601f8201819004810201604051908101604052818152929190602084018383808284378201915050505050509190803590602001908201803590602001908080601f016020809104026020016040519081016040528181529291906020840183838082843750949650509335935061312292505050565b3415610c0c57600080fd5b6101976001608060020a03196004351660ff602435166133f8565b3415610c3257600080fd5b610c476001608060020a031960043516613560565b6040518881526001608060020a031980891660208301528716604082015260a0810184905260ff831660c0820152606081016080820160e08301846005811115610c8d57fe5b60ff16815260208482038101845289546002600182161561010002600019019091160490820181905260409091019089908015610d0b5780601f10610ce057610100808354040283529160200191610d0b565b820191906000526020600020905b815481529060010190602001808311610cee57829003601f168201915b5050838103825287546002600019610100600184161502019091160480825260209091019088908015610d7f5780601f10610d5457610100808354040283529160200191610d7f565b820191906000526020600020905b815481529060010190602001808311610d6257829003601f168201915b50509a505050505050505050505060405180910390f35b3415610da157600080fd5b6101976001608060020a03196004351660ff602435166135ab565b3415610dc757600080fd5b610ddc6001608060020a03196004351661370e565b604051858152600160a060020a0385166020820152604081016060820160808301846005811115610e0957fe5b60ff16815260208482038101845287546002600182161561010002600019019091160490820181905260409091019087908015610e875780601f10610e5c57610100808354040283529160200191610e87565b820191906000526020600020905b815481529060010190602001808311610e6a57829003601f168201915b5050838103825285546002600019610100600184161502019091160480825260209091019086908015610efb5780601f10610ed057610100808354040283529160200191610efb565b820191906000526020600020905b815481529060010190602001808311610ede57829003601f168201915b505097505050505050505060405180910390f35b3415610f1a57600080fd5b610f366001608060020a03196004351661ffff60243516613745565b6040518085815260200180602001846004811115610f5057fe5b60ff16815260200180602001838103835286818151815260200191508051906020019080838360005b83811015610f91578082015183820152602001610f79565b50505050905090810190601f168015610fbe5780820380516001836020036101000a031916815260200191505b50838103825284818151815260200191508051906020019080838360005b83811015610ff4578082015183820152602001610fdc565b50505050905090810190601f1680156110215780820380516001836020036101000a031916815260200191505b50965050505050505060405180910390f35b341561103e57600080fd5b610707613924565b341561105157600080fd5b610197600480356001608060020a0319169060446024803590810190830135806020601f82018190048102016040519081016040528181529291906020840183838082843782019150505050505091908035600160a060020a031690602001909190803590602001908201803590602001908080601f01602080910402602001604051908101604052818152929190602084018383808284375094965061393395505050505050565b60008054600160a060020a0316633377925482604051602001526040518163ffffffff1660e060020a028152600401602060405180830381600087803b151561114257600080fd5b6102c65a03f1151561115357600080fd5b50505060405180519050600160a060020a031633600160a060020a031614151561117c57600080fd5b6001608060020a0319831660009081526002602052604081206008015460ff1660058111156111a757fe5b14156111b257600080fd5b6001608060020a031983166000908152600260208181526040808420928301805460ff908116865260019485019093528185208054600160a060020a031916600160a060020a0389811691909117909155815460ff198116908516909501909316939093179092558280529181902054909116907fb8a06c596f027284f088ce1ba17842629ca63619a7303c94b3645e310ae34e4490849051600160a060020a03909116815260200160405180910390a250600192915050565b60008054600160a060020a0316633377925482604051602001526040518163ffffffff1660e060020a028152600401602060405180830381600087803b15156112b457600080fd5b6102c65a03f115156112c557600080fd5b50505060405180519050600160a060020a031633600160a060020a03161415156112ee57600080fd5b60006001608060020a03198616600090815260066020526040902060030154608060020a900460ff16600581111561132257fe5b1461132c57600080fd5b6001608060020a0319841660009081526002602052604081206008015460ff16600581111561135757fe5b141561136257600080fd5b6001608060020a0319831660009081526004602052604081206006015460ff16600581111561138d57fe5b141561139857600080fd5b60a060405190810160409081524282526001608060020a03198087166020840152600160a060020a038516918301919091528416606082015260808101600190526001608060020a0319861660009081526006602052604090208151815560208201516001820180546001608060020a031916608060020a9092049190911790556040820151600282018054600160a060020a031916600160a060020a039290921691909117905560608201516003820180546001608060020a031916608060020a909204919091179055608082015160038201805470ff000000000000000000000000000000001916608060020a83600581111561149357fe5b021790555090505081600160a060020a03167ff23d682895fab50c3556e134ca7ff98bb5c052f2e98b27872f0d698e1b44ba7f866040516001608060020a0319909116815260200160405180910390a2506001949350505050565b60008054600160a060020a0316633377925482604051602001526040518163ffffffff1660e060020a028152600401602060405180830381600087803b151561153657600080fd5b6102c65a03f1151561154757600080fd5b50505060405180519050600160a060020a031633600160a060020a031614151561157057600080fd5b6001608060020a0319851660009081526002602052604081206008015460ff16600581111561159b57fe5b14156115a657600080fd5b6080604051908101604052804281526020018581526020018460048111156115ca57fe5b815260209081018490526001608060020a031987166000908152600282526040808220600681015461ffff168352600501909252208151815560208201518160010190805161161d929160200190613c04565b50604082015160028201805460ff1916600183600481111561163b57fe5b021790555060608201518160030190805161165a929160200190613c04565b5050506001608060020a03198516600090815260026020526040902060068101805461ffff198116600161ffff928316810190921617909155600491820180548693919260ff199091169184908111156116b057fe5b02179055506001608060020a0319851660009081526002602090815260408083208380526001019091529081902054600160a060020a0316907f3596012bbae932a31d934ff98585be7435f26aea5acf9e50d7ef6870feb2ace6908590518082600481111561171b57fe5b60ff16815260200191505060405180910390a2506001949350505050565b60008054600160a060020a0316633377925482604051602001526040518163ffffffff1660e060020a028152600401602060405180830381600087803b151561178157600080fd5b6102c65a03f1151561179257600080fd5b50505060405180519050600160a060020a031633600160a060020a03161415156117bb57600080fd5b60006001608060020a03198416600090815260066020526040902060030154608060020a900460ff1660058111156117ef57fe5b14156117fa57600080fd5b6001608060020a031983166000908152600660205260409020600301805483919070ff000000000000000000000000000000001916608060020a83600581111561184057fe5b02179055506001608060020a031983166000908152600660205260409081902060020154600160a060020a0316907f0faeaed899e84c7d49a914221ec870ec4ff7a9d0f013e7cfd6360d383820e97a9085908590516001608060020a031983168152602081018260058111156118b257fe5b60ff1681526020019250505060405180910390a250600192915050565b60008054600160a060020a0316633377925482604051602001526040518163ffffffff1660e060020a028152600401602060405180830381600087803b151561191757600080fd5b6102c65a03f1151561192857600080fd5b50505060405180519050600160a060020a031633600160a060020a031614151561195157600080fd5b6001608060020a03198416600090815260076020819052604082200154610100900460ff16600581111561198157fe5b141561198c57600080fd5b506001608060020a031983166000818152600760208181526040808420808401805460ff90811687526005830185528387208054600160a060020a038c16600160a060020a031990911617905581548116875260069092018452918520805463ffffffff891663ffffffff1990911617905594909352528054808316600190810190931660ff199091161790559392505050565b60008054600160a060020a0316633377925482604051602001526040518163ffffffff1660e060020a028152600401602060405180830381600087803b1515611a6857600080fd5b6102c65a03f11515611a7957600080fd5b50505060405180519050600160a060020a031633600160a060020a0316141515611aa257600080fd5b6001608060020a0319831660009081526004602052604081206006015460ff166005811115611acd57fe5b1415611ad857600080fd5b6001608060020a031983166000908152600460205260409020600601805483919060ff19166001836005811115611b0b57fe5b02179055506001608060020a031983166000908152600460205260409081902060020154600160a060020a0316907ff9ae5f2bc88e6348608d627c5226159efd8467eb4ac3c635cdb648266e02cb129085908590516001608060020a031983168152602081018260058111156118b257fe5b60008054600160a060020a0316633377925482604051602001526040518163ffffffff1660e060020a028152600401602060405180830381600087803b1515611bc557600080fd5b6102c65a03f11515611bd657600080fd5b50505060405180519050600160a060020a031633600160a060020a0316141515611bff57600080fd5b6001608060020a03198516600090815260056020819052604082206004015460ff1690811115611c2b57fe5b14611c3557600080fd5b60a06040519081016040908152428252600160a060020a038616602083015281018490526060810183905260808101600190526001608060020a031986166000908152600560205260409020815181556020820151600182018054600160a060020a031916600160a060020a0392909216919091179055604082015181600201908051611cc6929160200190613c04565b50606082015181600301908051611ce1929160200190613c04565b50608082015160048201805460ff19166001836005811115611cff57fe5b021790555090505083600160a060020a03167fb2f67f32df451ce918abd944fd5c8a29589848693057913c5060a8e0f293b38886856040516001608060020a03198316815260406020820181815290820183818151815260200191508051906020019080838360005b83811015611d80578082015183820152602001611d68565b50505050905090810190601f168015611dad5780820380516001836020036101000a031916815260200191505b50935050505060405180910390a2506001949350505050565b6000806000806000806000806000806000809054906101000a9004600160a060020a0316600160a060020a031663337792546000604051602001526040518163ffffffff1660e060020a028152600401602060405180830381600087803b1515611e2f57600080fd5b6102c65a03f11515611e4057600080fd5b50505060405180519050600160a060020a031633600160a060020a0316141515611e6957600080fd5b60046001608060020a03198c1660009081526007602081905260409091200154610100900460ff166005811115611e9c57fe5b14158015611edb575060056001608060020a03198c1660009081526007602081905260409091200154610100900460ff166005811115611ed857fe5b14155b1515611ee657600080fd5b6001608060020a03198b1660009081526007602052604081206004015411611f0d57600080fd5b6001608060020a03198b81166000908152600760208181526040808420928301805461ff0019166104001790556001830154608060020a90819004810286168552600683528185206003810154909102909516845260049091528083208354929d50939b5092995061200292670de0b6b3a764000092611ff692600160a060020a03169163e1489191919051602001526040518163ffffffff1660e060020a028152600401602060405180830381600087803b1515611fcb57600080fd5b6102c65a03f11515611fdc57600080fd5b505050604051805160048e0154915063ffffffff613bb016565b9063ffffffff613bdb16565b60048a015490965061201a908763ffffffff613bf216565b600154600289015460048c0154929750600160a060020a03918216926323b872dd9290911690309060006040516020015260405160e060020a63ffffffff8616028152600160a060020a0393841660048201529190921660248201526044810191909152606401602060405180830381600087803b151561209a57600080fd5b6102c65a03f115156120ab57600080fd5b5050506040518051505060015460008054600160a060020a039283169263a9059cbb929116906397c0262a90604051602001526040518163ffffffff1660e060020a028152600401602060405180830381600087803b151561210c57600080fd5b6102c65a03f1151561211d57600080fd5b505050604051805190508860006040516020015260405160e060020a63ffffffff8516028152600160a060020a0390921660048301526024820152604401602060405180830381600087803b151561217457600080fd5b6102c65a03f1151561218557600080fd5b50505060405180515060009450505b600789015460ff90811690851610156122985760ff8416600090815260058a01602090815260408083205460068d019092529091205460048b0154600160a060020a03909216945063ffffffff90811693506121fd916305f5e10091611ff691908690613bb016565b600154909150600160a060020a031663a9059cbb848360006040516020015260405160e060020a63ffffffff8516028152600160a060020a0390921660048301526024820152604401602060405180830381600087803b151561225f57600080fd5b6102c65a03f1151561227057600080fd5b50505060405180515061228b9050858263ffffffff613bf216565b9450600190930192612194565b6001546002890154600160a060020a039182169163a9059cbb91168760006040516020015260405160e060020a63ffffffff8516028152600160a060020a0390921660048301526024820152604401602060405180830381600087803b151561230057600080fd5b6102c65a03f1151561231157600080fd5b505050604051805150506002880154600160a060020a03167fd2d1d52d56b0b04c337afb9725a84b2f11a9e9ecb2b71869f2ddd6f5a4e0238f8c876040516001608060020a0319909216825260208201526040908101905180910390a25060019a9950505050505050505050565b600360208190526000918252604090912080546001820154600283015460058401549294608060020a90920293600160a060020a03909116929181019160049091019060ff1686565b6000806001608060020a0319841660009081526007602081905260409091200154610100900460ff1660058111156123fc57fe5b141561240757600080fd5b506001608060020a03198216600090815260076020908152604080832060ff8516845260060190915290205463ffffffff1692915050565b600260208190526000918252604090912080549181015460048201546006830154600884015460ff9384169460038101949384169361ffff909316926007909101911687565b600054600160a060020a031681565b60008054600160a060020a0316633377925482604051602001526040518163ffffffff1660e060020a028152600401602060405180830381600087803b15156124dc57600080fd5b6102c65a03f115156124ed57600080fd5b50505060405180519050600160a060020a031633600160a060020a031614151561251657600080fd5b6001608060020a03198316600090815260076020819052604082200154610100900460ff16600581111561254657fe5b141561255157600080fd5b6001608060020a03198316600090815260076020819052604090912001805483919061ff00191661010083600581111561258757fe5b02179055506001608060020a0319808416600090815260076020908152604080832060010154608060020a908190040290931682526006905281902060020154600160a060020a0316907ffa2a86784442baf24e238559b3ca5986aa4a588eeb7840c5ba8b27d6e663ecd99085908590516001608060020a031983168152602081018260058111156118b257fe5b60008054600160a060020a031690633377925490604051602001526040518163ffffffff1660e060020a028152600401602060405180830381600087803b151561265e57600080fd5b6102c65a03f1151561266f57600080fd5b50505060405180519050600160a060020a031633600160a060020a031614151561269857600080fd5b60008054600160a060020a031916600160a060020a0392909216919091179055565b60008054600160a060020a0316633377925482604051602001526040518163ffffffff1660e060020a028152600401602060405180830381600087803b151561270257600080fd5b6102c65a03f1151561271357600080fd5b50505060405180519050600160a060020a031633600160a060020a031614151561273c57600080fd5b6001608060020a0319831660009081526002602052604081206008015460ff16600581111561276757fe5b141561277257600080fd5b6001608060020a03198316600090815260026020526040902060070182805161279f929160200190613c04565b506001608060020a0319831660009081526002602090815260408083208380526001019091529081902054600160a060020a0316907f5775531a38eee0a6720d30a54a26748b21a9053e176d628a871b33885462e15e905160405180910390a250600192915050565b60008054600160a060020a0316633377925482604051602001526040518163ffffffff1660e060020a028152600401602060405180830381600087803b151561285057600080fd5b6102c65a03f1151561286157600080fd5b50505060405180519050600160a060020a031633600160a060020a031614151561288a57600080fd5b6001608060020a03198716600090815260076020819052604082200154610100900460ff1660058111156128ba57fe5b146128c457600080fd5b60006001608060020a03198716600090815260066020526040902060030154608060020a900460ff1660058111156128f857fe5b141561290357600080fd5b6001608060020a03198516600090815260056020819052604082206004015460ff169081111561292f57fe5b141561293a57600080fd5b61010060405190810160409081524282526001608060020a03198088166020840152881690820152606081018490526080810185905260a08101839052600060c082015260e08101600190526001608060020a0319881660009081526007602052604090208151815560208201516001820180546001608060020a031916608060020a90920491909117905560408201518160010160106101000a8154816fffffffffffffffffffffffffffffffff0219169083608060020a90040217905550606082015181600201908051612a14929160200190613c04565b50608082015181600301908051612a2f929160200190613c04565b5060a0820151816004015560c082015160078201805460ff191660ff9290921691909117905560e082015160078201805461ff001916610100836005811115612a7457fe5b021790555050506001608060020a031986166000908152600660205260409081902060020154600160a060020a0316907f5a4f5a43985de2c0807c6875a069d25bfdf90255dc58958a49e07d00936537e39089908590516001608060020a0319909216825260208201526040908101905180910390a25060019695505050505050565b60008054600160a060020a0316633377925482604051602001526040518163ffffffff1660e060020a028152600401602060405180830381600087803b1515612b3f57600080fd5b6102c65a03f11515612b5057600080fd5b50505060405180519050600160a060020a031633600160a060020a0316141515612b7957600080fd5b6001608060020a03198316600090815260036020526040812060059081015460ff1690811115612ba557fe5b1415612bb057600080fd5b6001608060020a03198316600090815260036020526040902060059081018054849260ff19909116906001908490811115612be757fe5b02179055506001608060020a031983166000908152600360205260409081902060020154600160a060020a0316907f14404ec3e30e0b13c4ef42e76c8fd3afe300c9bf51761965c78d348ff45cf2f09085908590516001608060020a031983168152602081018260058111156118b257fe5b6000805481908190600160a060020a0316633377925482604051602001526040518163ffffffff1660e060020a028152600401602060405180830381600087803b1515612ca557600080fd5b6102c65a03f11515612cb657600080fd5b50505060405180519050600160a060020a031633600160a060020a0316141515612cdf57600080fd5b6001608060020a03198816600090815260036020526040812060059081015460ff1690811115612d0b57fe5b14612d1557600080fd5b6001608060020a0319871660009081526002602052604081206008015460ff166005811115612d4057fe5b1415612d4b57600080fd5b5060009050805b6001608060020a031987166000908152600260208190526040909120015460ff9081169082161015612dc9576001608060020a03198716600090815260026020908152604080832060ff85168452600101909152902054600160a060020a0387811691161415612dc157600191505b600101612d52565b811515612dd557600080fd5b60c060405190810160409081524282526001608060020a031989166020830152600160a060020a03881690820152606081018690526080810185905260a08101600190526001608060020a0319891660009081526003602052604090208151815560208201516001820180546001608060020a031916608060020a9092049190911790556040820151600282018054600160a060020a031916600160a060020a0392909216919091179055606082015181600301908051612e9a929160200190613c04565b50608082015181600401908051612eb5929160200190613c04565b5060a08201518160050160006101000a81548160ff02191690836005811115612eda57fe5b021790555090505085600160a060020a03167f0eb23d662c57754d23afdc1cffa6246fbae3bbe843eec5c42a5a5223c21414ed89876040516001608060020a03198316815260406020820181815290820183818151815260200191508051906020019080838360005b83811015612f5b578082015183820152602001612f43565b50505050905090810190601f168015612f885780820380516001836020036101000a031916815260200191505b50935050505060405180910390a2506001979650505050505050565b6004602081905260009182526040909120805460018201546002830154600584015460068501549395608060020a90930294600160a060020a03909216936003830193929092019160ff1687565b6000806001608060020a0319841660009081526002602052604090206008015460ff16600581111561302057fe5b141561302b57600080fd5b506001608060020a03198216600090815260026020908152604080832060ff85168452600101909152902054600160a060020a031692915050565b60066020526000908152604090208054600182015460028301546003909301549192608060020a91820292600160a060020a0390911691808202910460ff1685565b6000806001608060020a0319841660009081526007602081905260409091200154610100900460ff1660058111156130dc57fe5b14156130e757600080fd5b506001608060020a03198216600090815260076020908152604080832060ff85168452600501909152902054600160a060020a031692915050565b60008054600160a060020a0316633377925482604051602001526040518163ffffffff1660e060020a028152600401602060405180830381600087803b151561316a57600080fd5b6102c65a03f1151561317b57600080fd5b50505060405180519050600160a060020a031633600160a060020a03161415156131a457600080fd5b6001608060020a0319871660009081526004602052604081206006015460ff1660058111156131cf57fe5b146131d957600080fd5b6001608060020a03198616600090815260036020526040812060059081015460ff169081111561320557fe5b141561321057600080fd5b60e060405190810160405280428152602001876fffffffffffffffffffffffffffffffff1916815260200186600160a060020a031681526020018581526020018481526020018381526020016001600581111561326957fe5b90526001608060020a0319881660009081526004602052604090208151815560208201516001820180546001608060020a031916608060020a9092049190911790556040820151600282018054600160a060020a031916600160a060020a03929092169190911790556060820151816003019080516132ec929160200190613c04565b50608082015181600401908051613307929160200190613c04565b5060a0820151816005015560c082015160068201805460ff1916600183600581111561332f57fe5b021790555090505084600160a060020a03167f5c9100ad970db426d43d47b3a9639dc42d9732871d99213dcb4984f7e71509d088866040516001608060020a03198316815260406020820181815290820183818151815260200191508051906020019080838360005b838110156133b0578082015183820152602001613398565b50505050905090810190601f1680156133dd5780820380516001836020036101000a031916815260200191505b50935050505060405180910390a25060019695505050505050565b60008054600160a060020a0316633377925482604051602001526040518163ffffffff1660e060020a028152600401602060405180830381600087803b151561344057600080fd5b6102c65a03f1151561345157600080fd5b50505060405180519050600160a060020a031633600160a060020a031614151561347a57600080fd5b6001608060020a0319831660009081526002602052604081206008015460ff1660058111156134a557fe5b14156134b057600080fd5b6001608060020a031983166000908152600260205260409020600801805483919060ff191660018360058111156134e357fe5b02179055506001608060020a0319831660009081526002602090815260408083208380526001019091529081902054600160a060020a0316907f183408f9d1ce45bacc33cd1a5ef30bf3caccbae9e28fcc8916b62185fdd5ac4d9085908590516001608060020a031983168152602081018260058111156118b257fe5b6007602081905260009182526040909120805460018201546004830154938301549193608060020a808302949281900402926002830192600301919060ff8082169161010090041688565b60008054600160a060020a0316633377925482604051602001526040518163ffffffff1660e060020a028152600401602060405180830381600087803b15156135f357600080fd5b6102c65a03f1151561360457600080fd5b50505060405180519050600160a060020a031633600160a060020a031614151561362d57600080fd5b6001608060020a03198316600090815260056020819052604082206004015460ff169081111561365957fe5b141561366457600080fd5b6001608060020a0319831660009081526005602081905260409091206004018054849260ff1990911690600190849081111561369c57fe5b02179055506001608060020a031983166000908152600560205260409081902060010154600160a060020a0316907f77a747805fe662172d50df0565e43aabadea2efb1198dc3ff8e60fcd57eb74899085908590516001608060020a031983168152602081018260058111156118b257fe5b60056020526000908152604090208054600182015460048301549192600160a060020a039091169160028201916003019060ff1685565b600061374f613c82565b6000613759613c82565b6001608060020a0319861660009081526002602052604081206008015460ff16600581111561378457fe5b141561378f57600080fd5b6001608060020a03198616600090815260026020818152604080842061ffff8a168552600501825292839020805481840154600180840180549397909660ff9093169560039095019487946101009381161593909302600019019092169290920491601f83018290048202909101905190810160405280929190818152602001828054600181600116156101000203166002900480156138705780601f1061384557610100808354040283529160200191613870565b820191906000526020600020905b81548152906001019060200180831161385357829003601f168201915b50505050509250808054600181600116156101000203166002900480601f01602080910402602001604051908101604052809291908181526020018280546001816001161561010002031660029004801561390c5780601f106138e15761010080835404028352916020019161390c565b820191906000526020600020905b8154815290600101906020018083116138ef57829003601f168201915b50505050509050935093509350935092959194509250565b600154600160a060020a031681565b60008054600160a060020a0316633377925482604051602001526040518163ffffffff1660e060020a028152600401602060405180830381600087803b151561397b57600080fd5b6102c65a03f1151561398c57600080fd5b50505060405180519050600160a060020a031633600160a060020a03161415156139b557600080fd5b6001608060020a0319851660009081526002602052604081206008015460ff1660058111156139e057fe5b146139ea57600080fd5b60e0604051908101604090815242825260016020830152810185905260608101600081526000602082015260408101849052606001600190526001608060020a03198616600090815260026020526040902081518155602082015160028201805460ff191660ff92909216919091179055604082015181600301908051613a75929160200190613c04565b5060608201518160040160006101000a81548160ff02191690836004811115613a9a57fe5b0217905550608082015160068201805461ffff191661ffff9290921691909117905560a082015181600701908051613ad6929160200190613c04565b5060c082015160088201805460ff19166001836005811115613af457fe5b021790555050506001608060020a031985166000908152600260209081526040808320838052600101909152908190208054600160a060020a031916600160a060020a038616908117909155907ff661ee1472892faaf2bb68cc6874f2759c9148b4234cac6b88f6ca362652f2759087908790516001608060020a031983168152604060208201818152908201838181518152602001915080519060200190808383600083811015611d80578082015183820152602001611d68565b6000828202831580613bcc5750828482811515613bc957fe5b04145b1515613bd457fe5b9392505050565b6000808284811515613be957fe5b04949350505050565b600082821115613bfe57fe5b50900390565b828054600181600116156101000203166002900490600052602060002090601f016020900481019282601f10613c4557805160ff1916838001178555613c72565b82800160010185558215613c72579182015b82811115613c72578251825591602001919060010190613c57565b50613c7e929150613c94565b5090565b60206040519081016040526000815290565b613cae91905b80821115613c7e5760008155600101613c9a565b905600a165627a7a723058204a5559d63b13c1ccfabd02d740218c3e9dce06f84757572484f50f151e4b58ef0029`

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

// Ads is a free data retrieval call binding the contract method 0xc85a7585.
//
// Solidity: function ads( bytes16) constant returns(createdAt uint256, ownerId bytes16, beneficiaryAddress address, offerId bytes16, status uint8)
func (_HoQuPlatform *HoQuPlatformCaller) Ads(opts *bind.CallOpts, arg0 [16]byte) (struct {
	CreatedAt          *big.Int
	OwnerId            [16]byte
	BeneficiaryAddress common.Address
	OfferId            [16]byte
	Status             uint8
}, error) {
	ret := new(struct {
		CreatedAt          *big.Int
		OwnerId            [16]byte
		BeneficiaryAddress common.Address
		OfferId            [16]byte
		Status             uint8
	})
	out := ret
	err := _HoQuPlatform.contract.Call(opts, out, "ads", arg0)
	return *ret, err
}

// Ads is a free data retrieval call binding the contract method 0xc85a7585.
//
// Solidity: function ads( bytes16) constant returns(createdAt uint256, ownerId bytes16, beneficiaryAddress address, offerId bytes16, status uint8)
func (_HoQuPlatform *HoQuPlatformSession) Ads(arg0 [16]byte) (struct {
	CreatedAt          *big.Int
	OwnerId            [16]byte
	BeneficiaryAddress common.Address
	OfferId            [16]byte
	Status             uint8
}, error) {
	return _HoQuPlatform.Contract.Ads(&_HoQuPlatform.CallOpts, arg0)
}

// Ads is a free data retrieval call binding the contract method 0xc85a7585.
//
// Solidity: function ads( bytes16) constant returns(createdAt uint256, ownerId bytes16, beneficiaryAddress address, offerId bytes16, status uint8)
func (_HoQuPlatform *HoQuPlatformCallerSession) Ads(arg0 [16]byte) (struct {
	CreatedAt          *big.Int
	OwnerId            [16]byte
	BeneficiaryAddress common.Address
	OfferId            [16]byte
	Status             uint8
}, error) {
	return _HoQuPlatform.Contract.Ads(&_HoQuPlatform.CallOpts, arg0)
}

// Companies is a free data retrieval call binding the contract method 0x4df2d847.
//
// Solidity: function companies( bytes16) constant returns(createdAt uint256, ownerId bytes16, ownerAddress address, name string, dataUrl string, status uint8)
func (_HoQuPlatform *HoQuPlatformCaller) Companies(opts *bind.CallOpts, arg0 [16]byte) (struct {
	CreatedAt    *big.Int
	OwnerId      [16]byte
	OwnerAddress common.Address
	Name         string
	DataUrl      string
	Status       uint8
}, error) {
	ret := new(struct {
		CreatedAt    *big.Int
		OwnerId      [16]byte
		OwnerAddress common.Address
		Name         string
		DataUrl      string
		Status       uint8
	})
	out := ret
	err := _HoQuPlatform.contract.Call(opts, out, "companies", arg0)
	return *ret, err
}

// Companies is a free data retrieval call binding the contract method 0x4df2d847.
//
// Solidity: function companies( bytes16) constant returns(createdAt uint256, ownerId bytes16, ownerAddress address, name string, dataUrl string, status uint8)
func (_HoQuPlatform *HoQuPlatformSession) Companies(arg0 [16]byte) (struct {
	CreatedAt    *big.Int
	OwnerId      [16]byte
	OwnerAddress common.Address
	Name         string
	DataUrl      string
	Status       uint8
}, error) {
	return _HoQuPlatform.Contract.Companies(&_HoQuPlatform.CallOpts, arg0)
}

// Companies is a free data retrieval call binding the contract method 0x4df2d847.
//
// Solidity: function companies( bytes16) constant returns(createdAt uint256, ownerId bytes16, ownerAddress address, name string, dataUrl string, status uint8)
func (_HoQuPlatform *HoQuPlatformCallerSession) Companies(arg0 [16]byte) (struct {
	CreatedAt    *big.Int
	OwnerId      [16]byte
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

// GetLeadIntermediaryAddress is a free data retrieval call binding the contract method 0xda7345b7.
//
// Solidity: function getLeadIntermediaryAddress(id bytes16, num uint8) constant returns(address)
func (_HoQuPlatform *HoQuPlatformCaller) GetLeadIntermediaryAddress(opts *bind.CallOpts, id [16]byte, num uint8) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _HoQuPlatform.contract.Call(opts, out, "getLeadIntermediaryAddress", id, num)
	return *ret0, err
}

// GetLeadIntermediaryAddress is a free data retrieval call binding the contract method 0xda7345b7.
//
// Solidity: function getLeadIntermediaryAddress(id bytes16, num uint8) constant returns(address)
func (_HoQuPlatform *HoQuPlatformSession) GetLeadIntermediaryAddress(id [16]byte, num uint8) (common.Address, error) {
	return _HoQuPlatform.Contract.GetLeadIntermediaryAddress(&_HoQuPlatform.CallOpts, id, num)
}

// GetLeadIntermediaryAddress is a free data retrieval call binding the contract method 0xda7345b7.
//
// Solidity: function getLeadIntermediaryAddress(id bytes16, num uint8) constant returns(address)
func (_HoQuPlatform *HoQuPlatformCallerSession) GetLeadIntermediaryAddress(id [16]byte, num uint8) (common.Address, error) {
	return _HoQuPlatform.Contract.GetLeadIntermediaryAddress(&_HoQuPlatform.CallOpts, id, num)
}

// GetLeadIntermediaryPercent is a free data retrieval call binding the contract method 0x5cb8f5ab.
//
// Solidity: function getLeadIntermediaryPercent(id bytes16, num uint8) constant returns(uint32)
func (_HoQuPlatform *HoQuPlatformCaller) GetLeadIntermediaryPercent(opts *bind.CallOpts, id [16]byte, num uint8) (uint32, error) {
	var (
		ret0 = new(uint32)
	)
	out := ret0
	err := _HoQuPlatform.contract.Call(opts, out, "getLeadIntermediaryPercent", id, num)
	return *ret0, err
}

// GetLeadIntermediaryPercent is a free data retrieval call binding the contract method 0x5cb8f5ab.
//
// Solidity: function getLeadIntermediaryPercent(id bytes16, num uint8) constant returns(uint32)
func (_HoQuPlatform *HoQuPlatformSession) GetLeadIntermediaryPercent(id [16]byte, num uint8) (uint32, error) {
	return _HoQuPlatform.Contract.GetLeadIntermediaryPercent(&_HoQuPlatform.CallOpts, id, num)
}

// GetLeadIntermediaryPercent is a free data retrieval call binding the contract method 0x5cb8f5ab.
//
// Solidity: function getLeadIntermediaryPercent(id bytes16, num uint8) constant returns(uint32)
func (_HoQuPlatform *HoQuPlatformCallerSession) GetLeadIntermediaryPercent(id [16]byte, num uint8) (uint32, error) {
	return _HoQuPlatform.Contract.GetLeadIntermediaryPercent(&_HoQuPlatform.CallOpts, id, num)
}

// GetUserAddress is a free data retrieval call binding the contract method 0xc49e2fb1.
//
// Solidity: function getUserAddress(id bytes16, num uint8) constant returns(address)
func (_HoQuPlatform *HoQuPlatformCaller) GetUserAddress(opts *bind.CallOpts, id [16]byte, num uint8) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _HoQuPlatform.contract.Call(opts, out, "getUserAddress", id, num)
	return *ret0, err
}

// GetUserAddress is a free data retrieval call binding the contract method 0xc49e2fb1.
//
// Solidity: function getUserAddress(id bytes16, num uint8) constant returns(address)
func (_HoQuPlatform *HoQuPlatformSession) GetUserAddress(id [16]byte, num uint8) (common.Address, error) {
	return _HoQuPlatform.Contract.GetUserAddress(&_HoQuPlatform.CallOpts, id, num)
}

// GetUserAddress is a free data retrieval call binding the contract method 0xc49e2fb1.
//
// Solidity: function getUserAddress(id bytes16, num uint8) constant returns(address)
func (_HoQuPlatform *HoQuPlatformCallerSession) GetUserAddress(id [16]byte, num uint8) (common.Address, error) {
	return _HoQuPlatform.Contract.GetUserAddress(&_HoQuPlatform.CallOpts, id, num)
}

// GetUserKycReport is a free data retrieval call binding the contract method 0xee263422.
//
// Solidity: function getUserKycReport(id bytes16, num uint16) constant returns(uint256, string, uint8, string)
func (_HoQuPlatform *HoQuPlatformCaller) GetUserKycReport(opts *bind.CallOpts, id [16]byte, num uint16) (*big.Int, string, uint8, string, error) {
	var (
		ret0 = new(*big.Int)
		ret1 = new(string)
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

// GetUserKycReport is a free data retrieval call binding the contract method 0xee263422.
//
// Solidity: function getUserKycReport(id bytes16, num uint16) constant returns(uint256, string, uint8, string)
func (_HoQuPlatform *HoQuPlatformSession) GetUserKycReport(id [16]byte, num uint16) (*big.Int, string, uint8, string, error) {
	return _HoQuPlatform.Contract.GetUserKycReport(&_HoQuPlatform.CallOpts, id, num)
}

// GetUserKycReport is a free data retrieval call binding the contract method 0xee263422.
//
// Solidity: function getUserKycReport(id bytes16, num uint16) constant returns(uint256, string, uint8, string)
func (_HoQuPlatform *HoQuPlatformCallerSession) GetUserKycReport(id [16]byte, num uint16) (*big.Int, string, uint8, string, error) {
	return _HoQuPlatform.Contract.GetUserKycReport(&_HoQuPlatform.CallOpts, id, num)
}

// Leads is a free data retrieval call binding the contract method 0xe7608fb2.
//
// Solidity: function leads( bytes16) constant returns(createdAt uint256, trackerId bytes16, adId bytes16, dataUrl string, meta string, price uint256, numOfIntermediaries uint8, status uint8)
func (_HoQuPlatform *HoQuPlatformCaller) Leads(opts *bind.CallOpts, arg0 [16]byte) (struct {
	CreatedAt           *big.Int
	TrackerId           [16]byte
	AdId                [16]byte
	DataUrl             string
	Meta                string
	Price               *big.Int
	NumOfIntermediaries uint8
	Status              uint8
}, error) {
	ret := new(struct {
		CreatedAt           *big.Int
		TrackerId           [16]byte
		AdId                [16]byte
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

// Leads is a free data retrieval call binding the contract method 0xe7608fb2.
//
// Solidity: function leads( bytes16) constant returns(createdAt uint256, trackerId bytes16, adId bytes16, dataUrl string, meta string, price uint256, numOfIntermediaries uint8, status uint8)
func (_HoQuPlatform *HoQuPlatformSession) Leads(arg0 [16]byte) (struct {
	CreatedAt           *big.Int
	TrackerId           [16]byte
	AdId                [16]byte
	DataUrl             string
	Meta                string
	Price               *big.Int
	NumOfIntermediaries uint8
	Status              uint8
}, error) {
	return _HoQuPlatform.Contract.Leads(&_HoQuPlatform.CallOpts, arg0)
}

// Leads is a free data retrieval call binding the contract method 0xe7608fb2.
//
// Solidity: function leads( bytes16) constant returns(createdAt uint256, trackerId bytes16, adId bytes16, dataUrl string, meta string, price uint256, numOfIntermediaries uint8, status uint8)
func (_HoQuPlatform *HoQuPlatformCallerSession) Leads(arg0 [16]byte) (struct {
	CreatedAt           *big.Int
	TrackerId           [16]byte
	AdId                [16]byte
	DataUrl             string
	Meta                string
	Price               *big.Int
	NumOfIntermediaries uint8
	Status              uint8
}, error) {
	return _HoQuPlatform.Contract.Leads(&_HoQuPlatform.CallOpts, arg0)
}

// Offers is a free data retrieval call binding the contract method 0xc32d869b.
//
// Solidity: function offers( bytes16) constant returns(createdAt uint256, companyId bytes16, payerAddress address, name string, dataUrl string, cost uint256, status uint8)
func (_HoQuPlatform *HoQuPlatformCaller) Offers(opts *bind.CallOpts, arg0 [16]byte) (struct {
	CreatedAt    *big.Int
	CompanyId    [16]byte
	PayerAddress common.Address
	Name         string
	DataUrl      string
	Cost         *big.Int
	Status       uint8
}, error) {
	ret := new(struct {
		CreatedAt    *big.Int
		CompanyId    [16]byte
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

// Offers is a free data retrieval call binding the contract method 0xc32d869b.
//
// Solidity: function offers( bytes16) constant returns(createdAt uint256, companyId bytes16, payerAddress address, name string, dataUrl string, cost uint256, status uint8)
func (_HoQuPlatform *HoQuPlatformSession) Offers(arg0 [16]byte) (struct {
	CreatedAt    *big.Int
	CompanyId    [16]byte
	PayerAddress common.Address
	Name         string
	DataUrl      string
	Cost         *big.Int
	Status       uint8
}, error) {
	return _HoQuPlatform.Contract.Offers(&_HoQuPlatform.CallOpts, arg0)
}

// Offers is a free data retrieval call binding the contract method 0xc32d869b.
//
// Solidity: function offers( bytes16) constant returns(createdAt uint256, companyId bytes16, payerAddress address, name string, dataUrl string, cost uint256, status uint8)
func (_HoQuPlatform *HoQuPlatformCallerSession) Offers(arg0 [16]byte) (struct {
	CreatedAt    *big.Int
	CompanyId    [16]byte
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

// Trackers is a free data retrieval call binding the contract method 0xea2fdd8c.
//
// Solidity: function trackers( bytes16) constant returns(createdAt uint256, ownerAddress address, name string, dataUrl string, status uint8)
func (_HoQuPlatform *HoQuPlatformCaller) Trackers(opts *bind.CallOpts, arg0 [16]byte) (struct {
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

// Trackers is a free data retrieval call binding the contract method 0xea2fdd8c.
//
// Solidity: function trackers( bytes16) constant returns(createdAt uint256, ownerAddress address, name string, dataUrl string, status uint8)
func (_HoQuPlatform *HoQuPlatformSession) Trackers(arg0 [16]byte) (struct {
	CreatedAt    *big.Int
	OwnerAddress common.Address
	Name         string
	DataUrl      string
	Status       uint8
}, error) {
	return _HoQuPlatform.Contract.Trackers(&_HoQuPlatform.CallOpts, arg0)
}

// Trackers is a free data retrieval call binding the contract method 0xea2fdd8c.
//
// Solidity: function trackers( bytes16) constant returns(createdAt uint256, ownerAddress address, name string, dataUrl string, status uint8)
func (_HoQuPlatform *HoQuPlatformCallerSession) Trackers(arg0 [16]byte) (struct {
	CreatedAt    *big.Int
	OwnerAddress common.Address
	Name         string
	DataUrl      string
	Status       uint8
}, error) {
	return _HoQuPlatform.Contract.Trackers(&_HoQuPlatform.CallOpts, arg0)
}

// Users is a free data retrieval call binding the contract method 0x7520dd14.
//
// Solidity: function users( bytes16) constant returns(createdAt uint256, numOfAddresses uint8, role string, kycLevel uint8, numOfKycReports uint16, pubKey string, status uint8)
func (_HoQuPlatform *HoQuPlatformCaller) Users(opts *bind.CallOpts, arg0 [16]byte) (struct {
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

// Users is a free data retrieval call binding the contract method 0x7520dd14.
//
// Solidity: function users( bytes16) constant returns(createdAt uint256, numOfAddresses uint8, role string, kycLevel uint8, numOfKycReports uint16, pubKey string, status uint8)
func (_HoQuPlatform *HoQuPlatformSession) Users(arg0 [16]byte) (struct {
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

// Users is a free data retrieval call binding the contract method 0x7520dd14.
//
// Solidity: function users( bytes16) constant returns(createdAt uint256, numOfAddresses uint8, role string, kycLevel uint8, numOfKycReports uint16, pubKey string, status uint8)
func (_HoQuPlatform *HoQuPlatformCallerSession) Users(arg0 [16]byte) (struct {
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

// AddAd is a paid mutator transaction binding the contract method 0x06501a10.
//
// Solidity: function addAd(id bytes16, ownerId bytes16, offerId bytes16, beneficiaryAddress address) returns(bool)
func (_HoQuPlatform *HoQuPlatformTransactor) AddAd(opts *bind.TransactOpts, id [16]byte, ownerId [16]byte, offerId [16]byte, beneficiaryAddress common.Address) (*types.Transaction, error) {
	return _HoQuPlatform.contract.Transact(opts, "addAd", id, ownerId, offerId, beneficiaryAddress)
}

// AddAd is a paid mutator transaction binding the contract method 0x06501a10.
//
// Solidity: function addAd(id bytes16, ownerId bytes16, offerId bytes16, beneficiaryAddress address) returns(bool)
func (_HoQuPlatform *HoQuPlatformSession) AddAd(id [16]byte, ownerId [16]byte, offerId [16]byte, beneficiaryAddress common.Address) (*types.Transaction, error) {
	return _HoQuPlatform.Contract.AddAd(&_HoQuPlatform.TransactOpts, id, ownerId, offerId, beneficiaryAddress)
}

// AddAd is a paid mutator transaction binding the contract method 0x06501a10.
//
// Solidity: function addAd(id bytes16, ownerId bytes16, offerId bytes16, beneficiaryAddress address) returns(bool)
func (_HoQuPlatform *HoQuPlatformTransactorSession) AddAd(id [16]byte, ownerId [16]byte, offerId [16]byte, beneficiaryAddress common.Address) (*types.Transaction, error) {
	return _HoQuPlatform.Contract.AddAd(&_HoQuPlatform.TransactOpts, id, ownerId, offerId, beneficiaryAddress)
}

// AddLead is a paid mutator transaction binding the contract method 0x8f2e960d.
//
// Solidity: function addLead(id bytes16, adId bytes16, trackerId bytes16, meta string, dataUrl string, price uint256) returns(bool)
func (_HoQuPlatform *HoQuPlatformTransactor) AddLead(opts *bind.TransactOpts, id [16]byte, adId [16]byte, trackerId [16]byte, meta string, dataUrl string, price *big.Int) (*types.Transaction, error) {
	return _HoQuPlatform.contract.Transact(opts, "addLead", id, adId, trackerId, meta, dataUrl, price)
}

// AddLead is a paid mutator transaction binding the contract method 0x8f2e960d.
//
// Solidity: function addLead(id bytes16, adId bytes16, trackerId bytes16, meta string, dataUrl string, price uint256) returns(bool)
func (_HoQuPlatform *HoQuPlatformSession) AddLead(id [16]byte, adId [16]byte, trackerId [16]byte, meta string, dataUrl string, price *big.Int) (*types.Transaction, error) {
	return _HoQuPlatform.Contract.AddLead(&_HoQuPlatform.TransactOpts, id, adId, trackerId, meta, dataUrl, price)
}

// AddLead is a paid mutator transaction binding the contract method 0x8f2e960d.
//
// Solidity: function addLead(id bytes16, adId bytes16, trackerId bytes16, meta string, dataUrl string, price uint256) returns(bool)
func (_HoQuPlatform *HoQuPlatformTransactorSession) AddLead(id [16]byte, adId [16]byte, trackerId [16]byte, meta string, dataUrl string, price *big.Int) (*types.Transaction, error) {
	return _HoQuPlatform.Contract.AddLead(&_HoQuPlatform.TransactOpts, id, adId, trackerId, meta, dataUrl, price)
}

// AddLeadIntermediary is a paid mutator transaction binding the contract method 0x36617045.
//
// Solidity: function addLeadIntermediary(id bytes16, intermediaryAddress address, percent uint32) returns(bool)
func (_HoQuPlatform *HoQuPlatformTransactor) AddLeadIntermediary(opts *bind.TransactOpts, id [16]byte, intermediaryAddress common.Address, percent uint32) (*types.Transaction, error) {
	return _HoQuPlatform.contract.Transact(opts, "addLeadIntermediary", id, intermediaryAddress, percent)
}

// AddLeadIntermediary is a paid mutator transaction binding the contract method 0x36617045.
//
// Solidity: function addLeadIntermediary(id bytes16, intermediaryAddress address, percent uint32) returns(bool)
func (_HoQuPlatform *HoQuPlatformSession) AddLeadIntermediary(id [16]byte, intermediaryAddress common.Address, percent uint32) (*types.Transaction, error) {
	return _HoQuPlatform.Contract.AddLeadIntermediary(&_HoQuPlatform.TransactOpts, id, intermediaryAddress, percent)
}

// AddLeadIntermediary is a paid mutator transaction binding the contract method 0x36617045.
//
// Solidity: function addLeadIntermediary(id bytes16, intermediaryAddress address, percent uint32) returns(bool)
func (_HoQuPlatform *HoQuPlatformTransactorSession) AddLeadIntermediary(id [16]byte, intermediaryAddress common.Address, percent uint32) (*types.Transaction, error) {
	return _HoQuPlatform.Contract.AddLeadIntermediary(&_HoQuPlatform.TransactOpts, id, intermediaryAddress, percent)
}

// AddOffer is a paid mutator transaction binding the contract method 0xdf7e40bb.
//
// Solidity: function addOffer(id bytes16, companyId bytes16, payerAddress address, name string, dataUrl string, cost uint256) returns(bool)
func (_HoQuPlatform *HoQuPlatformTransactor) AddOffer(opts *bind.TransactOpts, id [16]byte, companyId [16]byte, payerAddress common.Address, name string, dataUrl string, cost *big.Int) (*types.Transaction, error) {
	return _HoQuPlatform.contract.Transact(opts, "addOffer", id, companyId, payerAddress, name, dataUrl, cost)
}

// AddOffer is a paid mutator transaction binding the contract method 0xdf7e40bb.
//
// Solidity: function addOffer(id bytes16, companyId bytes16, payerAddress address, name string, dataUrl string, cost uint256) returns(bool)
func (_HoQuPlatform *HoQuPlatformSession) AddOffer(id [16]byte, companyId [16]byte, payerAddress common.Address, name string, dataUrl string, cost *big.Int) (*types.Transaction, error) {
	return _HoQuPlatform.Contract.AddOffer(&_HoQuPlatform.TransactOpts, id, companyId, payerAddress, name, dataUrl, cost)
}

// AddOffer is a paid mutator transaction binding the contract method 0xdf7e40bb.
//
// Solidity: function addOffer(id bytes16, companyId bytes16, payerAddress address, name string, dataUrl string, cost uint256) returns(bool)
func (_HoQuPlatform *HoQuPlatformTransactorSession) AddOffer(id [16]byte, companyId [16]byte, payerAddress common.Address, name string, dataUrl string, cost *big.Int) (*types.Transaction, error) {
	return _HoQuPlatform.Contract.AddOffer(&_HoQuPlatform.TransactOpts, id, companyId, payerAddress, name, dataUrl, cost)
}

// AddUserAddress is a paid mutator transaction binding the contract method 0x03806028.
//
// Solidity: function addUserAddress(id bytes16, ownerAddress address) returns(bool)
func (_HoQuPlatform *HoQuPlatformTransactor) AddUserAddress(opts *bind.TransactOpts, id [16]byte, ownerAddress common.Address) (*types.Transaction, error) {
	return _HoQuPlatform.contract.Transact(opts, "addUserAddress", id, ownerAddress)
}

// AddUserAddress is a paid mutator transaction binding the contract method 0x03806028.
//
// Solidity: function addUserAddress(id bytes16, ownerAddress address) returns(bool)
func (_HoQuPlatform *HoQuPlatformSession) AddUserAddress(id [16]byte, ownerAddress common.Address) (*types.Transaction, error) {
	return _HoQuPlatform.Contract.AddUserAddress(&_HoQuPlatform.TransactOpts, id, ownerAddress)
}

// AddUserAddress is a paid mutator transaction binding the contract method 0x03806028.
//
// Solidity: function addUserAddress(id bytes16, ownerAddress address) returns(bool)
func (_HoQuPlatform *HoQuPlatformTransactorSession) AddUserAddress(id [16]byte, ownerAddress common.Address) (*types.Transaction, error) {
	return _HoQuPlatform.Contract.AddUserAddress(&_HoQuPlatform.TransactOpts, id, ownerAddress)
}

// AddUserKycReport is a paid mutator transaction binding the contract method 0x11ab8673.
//
// Solidity: function addUserKycReport(id bytes16, meta string, kycLevel uint8, dataUrl string) returns(bool)
func (_HoQuPlatform *HoQuPlatformTransactor) AddUserKycReport(opts *bind.TransactOpts, id [16]byte, meta string, kycLevel uint8, dataUrl string) (*types.Transaction, error) {
	return _HoQuPlatform.contract.Transact(opts, "addUserKycReport", id, meta, kycLevel, dataUrl)
}

// AddUserKycReport is a paid mutator transaction binding the contract method 0x11ab8673.
//
// Solidity: function addUserKycReport(id bytes16, meta string, kycLevel uint8, dataUrl string) returns(bool)
func (_HoQuPlatform *HoQuPlatformSession) AddUserKycReport(id [16]byte, meta string, kycLevel uint8, dataUrl string) (*types.Transaction, error) {
	return _HoQuPlatform.Contract.AddUserKycReport(&_HoQuPlatform.TransactOpts, id, meta, kycLevel, dataUrl)
}

// AddUserKycReport is a paid mutator transaction binding the contract method 0x11ab8673.
//
// Solidity: function addUserKycReport(id bytes16, meta string, kycLevel uint8, dataUrl string) returns(bool)
func (_HoQuPlatform *HoQuPlatformTransactorSession) AddUserKycReport(id [16]byte, meta string, kycLevel uint8, dataUrl string) (*types.Transaction, error) {
	return _HoQuPlatform.Contract.AddUserKycReport(&_HoQuPlatform.TransactOpts, id, meta, kycLevel, dataUrl)
}

// RegisterCompany is a paid mutator transaction binding the contract method 0xbe95f1c2.
//
// Solidity: function registerCompany(id bytes16, ownerId bytes16, ownerAddress address, name string, dataUrl string) returns(bool)
func (_HoQuPlatform *HoQuPlatformTransactor) RegisterCompany(opts *bind.TransactOpts, id [16]byte, ownerId [16]byte, ownerAddress common.Address, name string, dataUrl string) (*types.Transaction, error) {
	return _HoQuPlatform.contract.Transact(opts, "registerCompany", id, ownerId, ownerAddress, name, dataUrl)
}

// RegisterCompany is a paid mutator transaction binding the contract method 0xbe95f1c2.
//
// Solidity: function registerCompany(id bytes16, ownerId bytes16, ownerAddress address, name string, dataUrl string) returns(bool)
func (_HoQuPlatform *HoQuPlatformSession) RegisterCompany(id [16]byte, ownerId [16]byte, ownerAddress common.Address, name string, dataUrl string) (*types.Transaction, error) {
	return _HoQuPlatform.Contract.RegisterCompany(&_HoQuPlatform.TransactOpts, id, ownerId, ownerAddress, name, dataUrl)
}

// RegisterCompany is a paid mutator transaction binding the contract method 0xbe95f1c2.
//
// Solidity: function registerCompany(id bytes16, ownerId bytes16, ownerAddress address, name string, dataUrl string) returns(bool)
func (_HoQuPlatform *HoQuPlatformTransactorSession) RegisterCompany(id [16]byte, ownerId [16]byte, ownerAddress common.Address, name string, dataUrl string) (*types.Transaction, error) {
	return _HoQuPlatform.Contract.RegisterCompany(&_HoQuPlatform.TransactOpts, id, ownerId, ownerAddress, name, dataUrl)
}

// RegisterTracker is a paid mutator transaction binding the contract method 0x37e830ba.
//
// Solidity: function registerTracker(id bytes16, ownerAddress address, name string, dataUrl string) returns(bool)
func (_HoQuPlatform *HoQuPlatformTransactor) RegisterTracker(opts *bind.TransactOpts, id [16]byte, ownerAddress common.Address, name string, dataUrl string) (*types.Transaction, error) {
	return _HoQuPlatform.contract.Transact(opts, "registerTracker", id, ownerAddress, name, dataUrl)
}

// RegisterTracker is a paid mutator transaction binding the contract method 0x37e830ba.
//
// Solidity: function registerTracker(id bytes16, ownerAddress address, name string, dataUrl string) returns(bool)
func (_HoQuPlatform *HoQuPlatformSession) RegisterTracker(id [16]byte, ownerAddress common.Address, name string, dataUrl string) (*types.Transaction, error) {
	return _HoQuPlatform.Contract.RegisterTracker(&_HoQuPlatform.TransactOpts, id, ownerAddress, name, dataUrl)
}

// RegisterTracker is a paid mutator transaction binding the contract method 0x37e830ba.
//
// Solidity: function registerTracker(id bytes16, ownerAddress address, name string, dataUrl string) returns(bool)
func (_HoQuPlatform *HoQuPlatformTransactorSession) RegisterTracker(id [16]byte, ownerAddress common.Address, name string, dataUrl string) (*types.Transaction, error) {
	return _HoQuPlatform.Contract.RegisterTracker(&_HoQuPlatform.TransactOpts, id, ownerAddress, name, dataUrl)
}

// RegisterUser is a paid mutator transaction binding the contract method 0xfd3fc08f.
//
// Solidity: function registerUser(id bytes16, role string, ownerAddress address, pubKey string) returns(bool)
func (_HoQuPlatform *HoQuPlatformTransactor) RegisterUser(opts *bind.TransactOpts, id [16]byte, role string, ownerAddress common.Address, pubKey string) (*types.Transaction, error) {
	return _HoQuPlatform.contract.Transact(opts, "registerUser", id, role, ownerAddress, pubKey)
}

// RegisterUser is a paid mutator transaction binding the contract method 0xfd3fc08f.
//
// Solidity: function registerUser(id bytes16, role string, ownerAddress address, pubKey string) returns(bool)
func (_HoQuPlatform *HoQuPlatformSession) RegisterUser(id [16]byte, role string, ownerAddress common.Address, pubKey string) (*types.Transaction, error) {
	return _HoQuPlatform.Contract.RegisterUser(&_HoQuPlatform.TransactOpts, id, role, ownerAddress, pubKey)
}

// RegisterUser is a paid mutator transaction binding the contract method 0xfd3fc08f.
//
// Solidity: function registerUser(id bytes16, role string, ownerAddress address, pubKey string) returns(bool)
func (_HoQuPlatform *HoQuPlatformTransactorSession) RegisterUser(id [16]byte, role string, ownerAddress common.Address, pubKey string) (*types.Transaction, error) {
	return _HoQuPlatform.Contract.RegisterUser(&_HoQuPlatform.TransactOpts, id, role, ownerAddress, pubKey)
}

// SellLead is a paid mutator transaction binding the contract method 0x47f3d978.
//
// Solidity: function sellLead(id bytes16) returns(bool)
func (_HoQuPlatform *HoQuPlatformTransactor) SellLead(opts *bind.TransactOpts, id [16]byte) (*types.Transaction, error) {
	return _HoQuPlatform.contract.Transact(opts, "sellLead", id)
}

// SellLead is a paid mutator transaction binding the contract method 0x47f3d978.
//
// Solidity: function sellLead(id bytes16) returns(bool)
func (_HoQuPlatform *HoQuPlatformSession) SellLead(id [16]byte) (*types.Transaction, error) {
	return _HoQuPlatform.Contract.SellLead(&_HoQuPlatform.TransactOpts, id)
}

// SellLead is a paid mutator transaction binding the contract method 0x47f3d978.
//
// Solidity: function sellLead(id bytes16) returns(bool)
func (_HoQuPlatform *HoQuPlatformTransactorSession) SellLead(id [16]byte) (*types.Transaction, error) {
	return _HoQuPlatform.Contract.SellLead(&_HoQuPlatform.TransactOpts, id)
}

// SetAdStatus is a paid mutator transaction binding the contract method 0x1bf52438.
//
// Solidity: function setAdStatus(id bytes16, status uint8) returns(bool)
func (_HoQuPlatform *HoQuPlatformTransactor) SetAdStatus(opts *bind.TransactOpts, id [16]byte, status uint8) (*types.Transaction, error) {
	return _HoQuPlatform.contract.Transact(opts, "setAdStatus", id, status)
}

// SetAdStatus is a paid mutator transaction binding the contract method 0x1bf52438.
//
// Solidity: function setAdStatus(id bytes16, status uint8) returns(bool)
func (_HoQuPlatform *HoQuPlatformSession) SetAdStatus(id [16]byte, status uint8) (*types.Transaction, error) {
	return _HoQuPlatform.Contract.SetAdStatus(&_HoQuPlatform.TransactOpts, id, status)
}

// SetAdStatus is a paid mutator transaction binding the contract method 0x1bf52438.
//
// Solidity: function setAdStatus(id bytes16, status uint8) returns(bool)
func (_HoQuPlatform *HoQuPlatformTransactorSession) SetAdStatus(id [16]byte, status uint8) (*types.Transaction, error) {
	return _HoQuPlatform.Contract.SetAdStatus(&_HoQuPlatform.TransactOpts, id, status)
}

// SetCompanyStatus is a paid mutator transaction binding the contract method 0xa8800f09.
//
// Solidity: function setCompanyStatus(id bytes16, status uint8) returns(bool)
func (_HoQuPlatform *HoQuPlatformTransactor) SetCompanyStatus(opts *bind.TransactOpts, id [16]byte, status uint8) (*types.Transaction, error) {
	return _HoQuPlatform.contract.Transact(opts, "setCompanyStatus", id, status)
}

// SetCompanyStatus is a paid mutator transaction binding the contract method 0xa8800f09.
//
// Solidity: function setCompanyStatus(id bytes16, status uint8) returns(bool)
func (_HoQuPlatform *HoQuPlatformSession) SetCompanyStatus(id [16]byte, status uint8) (*types.Transaction, error) {
	return _HoQuPlatform.Contract.SetCompanyStatus(&_HoQuPlatform.TransactOpts, id, status)
}

// SetCompanyStatus is a paid mutator transaction binding the contract method 0xa8800f09.
//
// Solidity: function setCompanyStatus(id bytes16, status uint8) returns(bool)
func (_HoQuPlatform *HoQuPlatformTransactorSession) SetCompanyStatus(id [16]byte, status uint8) (*types.Transaction, error) {
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

// SetLeadStatus is a paid mutator transaction binding the contract method 0x81e85aa3.
//
// Solidity: function setLeadStatus(id bytes16, status uint8) returns(bool)
func (_HoQuPlatform *HoQuPlatformTransactor) SetLeadStatus(opts *bind.TransactOpts, id [16]byte, status uint8) (*types.Transaction, error) {
	return _HoQuPlatform.contract.Transact(opts, "setLeadStatus", id, status)
}

// SetLeadStatus is a paid mutator transaction binding the contract method 0x81e85aa3.
//
// Solidity: function setLeadStatus(id bytes16, status uint8) returns(bool)
func (_HoQuPlatform *HoQuPlatformSession) SetLeadStatus(id [16]byte, status uint8) (*types.Transaction, error) {
	return _HoQuPlatform.Contract.SetLeadStatus(&_HoQuPlatform.TransactOpts, id, status)
}

// SetLeadStatus is a paid mutator transaction binding the contract method 0x81e85aa3.
//
// Solidity: function setLeadStatus(id bytes16, status uint8) returns(bool)
func (_HoQuPlatform *HoQuPlatformTransactorSession) SetLeadStatus(id [16]byte, status uint8) (*types.Transaction, error) {
	return _HoQuPlatform.Contract.SetLeadStatus(&_HoQuPlatform.TransactOpts, id, status)
}

// SetOfferStatus is a paid mutator transaction binding the contract method 0x36d7523f.
//
// Solidity: function setOfferStatus(id bytes16, status uint8) returns(bool)
func (_HoQuPlatform *HoQuPlatformTransactor) SetOfferStatus(opts *bind.TransactOpts, id [16]byte, status uint8) (*types.Transaction, error) {
	return _HoQuPlatform.contract.Transact(opts, "setOfferStatus", id, status)
}

// SetOfferStatus is a paid mutator transaction binding the contract method 0x36d7523f.
//
// Solidity: function setOfferStatus(id bytes16, status uint8) returns(bool)
func (_HoQuPlatform *HoQuPlatformSession) SetOfferStatus(id [16]byte, status uint8) (*types.Transaction, error) {
	return _HoQuPlatform.Contract.SetOfferStatus(&_HoQuPlatform.TransactOpts, id, status)
}

// SetOfferStatus is a paid mutator transaction binding the contract method 0x36d7523f.
//
// Solidity: function setOfferStatus(id bytes16, status uint8) returns(bool)
func (_HoQuPlatform *HoQuPlatformTransactorSession) SetOfferStatus(id [16]byte, status uint8) (*types.Transaction, error) {
	return _HoQuPlatform.Contract.SetOfferStatus(&_HoQuPlatform.TransactOpts, id, status)
}

// SetTrackerStatus is a paid mutator transaction binding the contract method 0xe7e5aad8.
//
// Solidity: function setTrackerStatus(id bytes16, status uint8) returns(bool)
func (_HoQuPlatform *HoQuPlatformTransactor) SetTrackerStatus(opts *bind.TransactOpts, id [16]byte, status uint8) (*types.Transaction, error) {
	return _HoQuPlatform.contract.Transact(opts, "setTrackerStatus", id, status)
}

// SetTrackerStatus is a paid mutator transaction binding the contract method 0xe7e5aad8.
//
// Solidity: function setTrackerStatus(id bytes16, status uint8) returns(bool)
func (_HoQuPlatform *HoQuPlatformSession) SetTrackerStatus(id [16]byte, status uint8) (*types.Transaction, error) {
	return _HoQuPlatform.Contract.SetTrackerStatus(&_HoQuPlatform.TransactOpts, id, status)
}

// SetTrackerStatus is a paid mutator transaction binding the contract method 0xe7e5aad8.
//
// Solidity: function setTrackerStatus(id bytes16, status uint8) returns(bool)
func (_HoQuPlatform *HoQuPlatformTransactorSession) SetTrackerStatus(id [16]byte, status uint8) (*types.Transaction, error) {
	return _HoQuPlatform.Contract.SetTrackerStatus(&_HoQuPlatform.TransactOpts, id, status)
}

// SetUserStatus is a paid mutator transaction binding the contract method 0xe03d8b64.
//
// Solidity: function setUserStatus(id bytes16, status uint8) returns(bool)
func (_HoQuPlatform *HoQuPlatformTransactor) SetUserStatus(opts *bind.TransactOpts, id [16]byte, status uint8) (*types.Transaction, error) {
	return _HoQuPlatform.contract.Transact(opts, "setUserStatus", id, status)
}

// SetUserStatus is a paid mutator transaction binding the contract method 0xe03d8b64.
//
// Solidity: function setUserStatus(id bytes16, status uint8) returns(bool)
func (_HoQuPlatform *HoQuPlatformSession) SetUserStatus(id [16]byte, status uint8) (*types.Transaction, error) {
	return _HoQuPlatform.Contract.SetUserStatus(&_HoQuPlatform.TransactOpts, id, status)
}

// SetUserStatus is a paid mutator transaction binding the contract method 0xe03d8b64.
//
// Solidity: function setUserStatus(id bytes16, status uint8) returns(bool)
func (_HoQuPlatform *HoQuPlatformTransactorSession) SetUserStatus(id [16]byte, status uint8) (*types.Transaction, error) {
	return _HoQuPlatform.Contract.SetUserStatus(&_HoQuPlatform.TransactOpts, id, status)
}

// UpdateUserPubKey is a paid mutator transaction binding the contract method 0x8a5463bd.
//
// Solidity: function updateUserPubKey(id bytes16, pubKey string) returns(bool)
func (_HoQuPlatform *HoQuPlatformTransactor) UpdateUserPubKey(opts *bind.TransactOpts, id [16]byte, pubKey string) (*types.Transaction, error) {
	return _HoQuPlatform.contract.Transact(opts, "updateUserPubKey", id, pubKey)
}

// UpdateUserPubKey is a paid mutator transaction binding the contract method 0x8a5463bd.
//
// Solidity: function updateUserPubKey(id bytes16, pubKey string) returns(bool)
func (_HoQuPlatform *HoQuPlatformSession) UpdateUserPubKey(id [16]byte, pubKey string) (*types.Transaction, error) {
	return _HoQuPlatform.Contract.UpdateUserPubKey(&_HoQuPlatform.TransactOpts, id, pubKey)
}

// UpdateUserPubKey is a paid mutator transaction binding the contract method 0x8a5463bd.
//
// Solidity: function updateUserPubKey(id bytes16, pubKey string) returns(bool)
func (_HoQuPlatform *HoQuPlatformTransactorSession) UpdateUserPubKey(id [16]byte, pubKey string) (*types.Transaction, error) {
	return _HoQuPlatform.Contract.UpdateUserPubKey(&_HoQuPlatform.TransactOpts, id, pubKey)
}

// HoQuPlatformConfigABI is the input ABI used to generate the binding from.
const HoQuPlatformConfigABI = "[{\"constant\":true,\"inputs\":[],\"name\":\"systemOwner\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_commission\",\"type\":\"uint256\"}],\"name\":\"setCommission\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_systemOwner\",\"type\":\"address\"}],\"name\":\"setSystemOwner\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_commissionWallet\",\"type\":\"address\"}],\"name\":\"setCommissionWallet\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"commissionWallet\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"commission\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"transferOwnership\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"name\":\"_systemOwner\",\"type\":\"address\"},{\"name\":\"_commissionWallet\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"previousOwner\",\"type\":\"address\"},{\"indexed\":false,\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"SystemOwnerChanged\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"changedBy\",\"type\":\"address\"},{\"indexed\":false,\"name\":\"commissionWallet\",\"type\":\"address\"}],\"name\":\"CommissionWalletChanged\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"changedBy\",\"type\":\"address\"},{\"indexed\":false,\"name\":\"commission\",\"type\":\"uint256\"}],\"name\":\"CommissionChanged\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"previousOwner\",\"type\":\"address\"},{\"indexed\":true,\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"OwnershipTransferred\",\"type\":\"event\"}]"

// HoQuPlatformConfigBin is the compiled bytecode used for deploying new contracts.
const HoQuPlatformConfigBin = `0x60606040526611c37937e08000600355341561001a57600080fd5b60405160408061050a833981016040528080519190602001805160008054600160a060020a03338116600160a060020a031992831617909255600180549683169682169690961790955560028054919092169416939093179092555050610484806100866000396000f30060606040526004361061008d5763ffffffff7c0100000000000000000000000000000000000000000000000000000000600035041663337792548114610092578063355e6b43146100c1578063621eb9a2146100d95780637d60b6ce146100f85780638da5cb5b1461011757806397c0262a1461012a578063e14891911461013d578063f2fde38b14610162575b600080fd5b341561009d57600080fd5b6100a5610181565b604051600160a060020a03909116815260200160405180910390f35b34156100cc57600080fd5b6100d7600435610190565b005b34156100e457600080fd5b6100d7600160a060020a0360043516610215565b341561010357600080fd5b6100d7600160a060020a03600435166102d8565b341561012257600080fd5b6100a5610399565b341561013557600080fd5b6100a56103a8565b341561014857600080fd5b6101506103b7565b60405190815260200160405180910390f35b341561016d57600080fd5b6100d7600160a060020a03600435166103bd565b600154600160a060020a031681565b60005433600160a060020a03908116911614806101bb575060015433600160a060020a039081169116145b15156101c657600080fd5b600081116101d357600080fd5b33600160a060020a03167ff659af9785c5d8c6128fdcbabb637f7d7ea787226f13cad89f71cbc579614fdc8260405190815260200160405180910390a2600355565b60005433600160a060020a0390811691161480610240575060015433600160a060020a039081169116145b151561024b57600080fd5b600160a060020a038116151561026057600080fd5b600154600160a060020a03167f1722097e017bb9a189312ac57bbdfff1bf251383cfd36c245f8ebabd8734c75982604051600160a060020a03909116815260200160405180910390a26001805473ffffffffffffffffffffffffffffffffffffffff1916600160a060020a0392909216919091179055565b60005433600160a060020a0390811691161480610303575060015433600160a060020a039081169116145b151561030e57600080fd5b600160a060020a038116151561032357600080fd5b33600160a060020a03167f394e7a1cfafd7b3c49218efbf47fb95da1832e5eaf0e9dc0428ea7220959d0e082604051600160a060020a03909116815260200160405180910390a26002805473ffffffffffffffffffffffffffffffffffffffff1916600160a060020a0392909216919091179055565b600054600160a060020a031681565b600254600160a060020a031681565b60035481565b60005433600160a060020a039081169116146103d857600080fd5b600160a060020a03811615156103ed57600080fd5b600054600160a060020a0380831691167f8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e060405160405180910390a36000805473ffffffffffffffffffffffffffffffffffffffff1916600160a060020a03929092169190911790555600a165627a7a72305820e8da2a87c337b6f3e5c07475dbad0a7ea4039c2896fd0f435f3c372edef1e2bb0029`

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
const HoQuTokenBin = `0x60606040526003805460a060020a60ff0219169055341561001f57600080fd5b604051602080610b998339810160405280805160038054600160a060020a03191633600160a060020a03161790559150506000811161005d57600080fd5b600160a060020a033316600090815260016020526040812082905555610b11806100886000396000f3006060604052600436106100e55763ffffffff7c010000000000000000000000000000000000000000000000000000000060003504166306fdde0381146100ea578063095ea7b31461017457806318160ddd146101aa57806323b872dd146101cf578063313ce567146101f75780633f4ba83a146102235780635c975abb14610238578063661884631461024b57806370a082311461026d5780638456cb591461028c5780638da5cb5b1461029f57806395d89b41146102ce578063a9059cbb146102e1578063d73dd62314610303578063dd62ed3e14610325578063f2fde38b1461034a575b600080fd5b34156100f557600080fd5b6100fd610369565b60405160208082528190810183818151815260200191508051906020019080838360005b83811015610139578082015183820152602001610121565b50505050905090810190601f1680156101665780820380516001836020036101000a031916815260200191505b509250505060405180910390f35b341561017f57600080fd5b610196600160a060020a03600435166024356103a0565b604051901515815260200160405180910390f35b34156101b557600080fd5b6101bd61040c565b60405190815260200160405180910390f35b34156101da57600080fd5b610196600160a060020a0360043581169060243516604435610412565b341561020257600080fd5b61020a61043f565b60405163ffffffff909116815260200160405180910390f35b341561022e57600080fd5b610236610444565b005b341561024357600080fd5b6101966104c3565b341561025657600080fd5b610196600160a060020a03600435166024356104d3565b341561027857600080fd5b6101bd600160a060020a03600435166105cd565b341561029757600080fd5b6102366105e8565b34156102aa57600080fd5b6102b261066c565b604051600160a060020a03909116815260200160405180910390f35b34156102d957600080fd5b6100fd61067b565b34156102ec57600080fd5b610196600160a060020a03600435166024356106b2565b341561030e57600080fd5b610196600160a060020a03600435166024356106dd565b341561033057600080fd5b6101bd600160a060020a0360043581169060243516610781565b341561035557600080fd5b610236600160a060020a03600435166107ac565b60408051908101604052600a81527f484f515520546f6b656e00000000000000000000000000000000000000000000602082015281565b600160a060020a03338116600081815260026020908152604080832094871680845294909152808220859055909291907f8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b9259085905190815260200160405180910390a350600192915050565b60005481565b60035460009060a060020a900460ff161561042c57600080fd5b610437848484610847565b949350505050565b601281565b60035433600160a060020a0390811691161461045f57600080fd5b60035460a060020a900460ff16151561047757600080fd5b6003805474ff0000000000000000000000000000000000000000191690557f7805862f689e2f13df9f062ff482ad3ad112aca9e0847911ed832e158c525b3360405160405180910390a1565b60035460a060020a900460ff1681565b600160a060020a0333811660009081526002602090815260408083209386168352929052908120548083111561053057600160a060020a033381166000908152600260209081526040808320938816835292905290812055610567565b610540818463ffffffff6109c916565b600160a060020a033381166000908152600260209081526040808320938916835292905220555b600160a060020a0333811660008181526002602090815260408083209489168084529490915290819020547f8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925915190815260200160405180910390a35060019392505050565b600160a060020a031660009081526001602052604090205490565b60035433600160a060020a0390811691161461060357600080fd5b60035460a060020a900460ff161561061a57600080fd5b6003805474ff0000000000000000000000000000000000000000191660a060020a1790557f6985a02210a168e66602d3235cb6db0e70f92b3ba4d376a33c0f3d9434bff62560405160405180910390a1565b600354600160a060020a031681565b60408051908101604052600381527f4851580000000000000000000000000000000000000000000000000000000000602082015281565b60035460009060a060020a900460ff16156106cc57600080fd5b6106d683836109db565b9392505050565b600160a060020a033381166000908152600260209081526040808320938616835292905290812054610715908363ffffffff610ad616565b600160a060020a0333811660008181526002602090815260408083209489168084529490915290819020849055919290917f8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b92591905190815260200160405180910390a350600192915050565b600160a060020a03918216600090815260026020908152604080832093909416825291909152205490565b60035433600160a060020a039081169116146107c757600080fd5b600160a060020a03811615156107dc57600080fd5b600354600160a060020a0380831691167f8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e060405160405180910390a36003805473ffffffffffffffffffffffffffffffffffffffff1916600160a060020a0392909216919091179055565b6000600160a060020a038316151561085e57600080fd5b600160a060020a03841660009081526001602052604090205482111561088357600080fd5b600160a060020a03808516600090815260026020908152604080832033909416835292905220548211156108b657600080fd5b600160a060020a0384166000908152600160205260409020546108df908363ffffffff6109c916565b600160a060020a038086166000908152600160205260408082209390935590851681522054610914908363ffffffff610ad616565b600160a060020a0380851660009081526001602090815260408083209490945587831682526002815283822033909316825291909152205461095c908363ffffffff6109c916565b600160a060020a03808616600081815260026020908152604080832033861684529091529081902093909355908516917fddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef9085905190815260200160405180910390a35060019392505050565b6000828211156109d557fe5b50900390565b6000600160a060020a03831615156109f257600080fd5b600160a060020a033316600090815260016020526040902054821115610a1757600080fd5b600160a060020a033316600090815260016020526040902054610a40908363ffffffff6109c916565b600160a060020a033381166000908152600160205260408082209390935590851681522054610a75908363ffffffff610ad616565b600160a060020a0380851660008181526001602052604090819020939093559133909116907fddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef9085905190815260200160405180910390a350600192915050565b6000828201838110156106d657fe00a165627a7a723058209dc51ff288a171e16fbdc5727f743bb83ec2c1228d298186e4c2a30926a8a2b70029`

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
const OwnableBin = `0x6060604052341561000f57600080fd5b60008054600160a060020a033316600160a060020a03199091161790556101768061003b6000396000f30060606040526004361061004b5763ffffffff7c01000000000000000000000000000000000000000000000000000000006000350416638da5cb5b8114610050578063f2fde38b1461007f575b600080fd5b341561005b57600080fd5b6100636100a0565b604051600160a060020a03909116815260200160405180910390f35b341561008a57600080fd5b61009e600160a060020a03600435166100af565b005b600054600160a060020a031681565b60005433600160a060020a039081169116146100ca57600080fd5b600160a060020a03811615156100df57600080fd5b600054600160a060020a0380831691167f8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e060405160405180910390a36000805473ffffffffffffffffffffffffffffffffffffffff1916600160a060020a03929092169190911790555600a165627a7a72305820e6aa5a01d9e4049290779debc9d05a8ce023346b4e48d4920d9f0d7bf7e7af680029`

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
const PausableBin = `0x606060405260008054600160a060020a033316600160a860020a03199091161790556102f7806100306000396000f30060606040526004361061006c5763ffffffff7c01000000000000000000000000000000000000000000000000000000006000350416633f4ba83a81146100715780635c975abb146100865780638456cb59146100ad5780638da5cb5b146100c0578063f2fde38b146100ef575b600080fd5b341561007c57600080fd5b61008461010e565b005b341561009157600080fd5b61009961018d565b604051901515815260200160405180910390f35b34156100b857600080fd5b61008461019d565b34156100cb57600080fd5b6100d3610221565b604051600160a060020a03909116815260200160405180910390f35b34156100fa57600080fd5b610084600160a060020a0360043516610230565b60005433600160a060020a0390811691161461012957600080fd5b60005460a060020a900460ff16151561014157600080fd5b6000805474ff0000000000000000000000000000000000000000191690557f7805862f689e2f13df9f062ff482ad3ad112aca9e0847911ed832e158c525b3360405160405180910390a1565b60005460a060020a900460ff1681565b60005433600160a060020a039081169116146101b857600080fd5b60005460a060020a900460ff16156101cf57600080fd5b6000805474ff0000000000000000000000000000000000000000191660a060020a1790557f6985a02210a168e66602d3235cb6db0e70f92b3ba4d376a33c0f3d9434bff62560405160405180910390a1565b600054600160a060020a031681565b60005433600160a060020a0390811691161461024b57600080fd5b600160a060020a038116151561026057600080fd5b600054600160a060020a0380831691167f8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e060405160405180910390a36000805473ffffffffffffffffffffffffffffffffffffffff1916600160a060020a03929092169190911790555600a165627a7a7230582027f76a2878fd3af7f1b05b341fc39d2977895a957c8e22f194270efc92ad93360029`

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
const SafeMathBin = `0x60606040523415600e57600080fd5b603580601b6000396000f3006060604052600080fd00a165627a7a72305820ca74ed078225c9747b390a8202dbaf2f038ea1c9552fa4e62f4f84391effcebf0029`

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
const StandardTokenBin = `0x6060604052341561000f57600080fd5b6106e68061001e6000396000f30060606040526004361061008d5763ffffffff7c0100000000000000000000000000000000000000000000000000000000600035041663095ea7b3811461009257806318160ddd146100c857806323b872dd146100ed578063661884631461011557806370a0823114610137578063a9059cbb14610156578063d73dd62314610178578063dd62ed3e1461019a575b600080fd5b341561009d57600080fd5b6100b4600160a060020a03600435166024356101bf565b604051901515815260200160405180910390f35b34156100d357600080fd5b6100db61022b565b60405190815260200160405180910390f35b34156100f857600080fd5b6100b4600160a060020a0360043581169060243516604435610231565b341561012057600080fd5b6100b4600160a060020a03600435166024356103b3565b341561014257600080fd5b6100db600160a060020a03600435166104ad565b341561016157600080fd5b6100b4600160a060020a03600435166024356104c8565b341561018357600080fd5b6100b4600160a060020a03600435166024356105c3565b34156101a557600080fd5b6100db600160a060020a0360043581169060243516610667565b600160a060020a03338116600081815260026020908152604080832094871680845294909152808220859055909291907f8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b9259085905190815260200160405180910390a350600192915050565b60005481565b6000600160a060020a038316151561024857600080fd5b600160a060020a03841660009081526001602052604090205482111561026d57600080fd5b600160a060020a03808516600090815260026020908152604080832033909416835292905220548211156102a057600080fd5b600160a060020a0384166000908152600160205260409020546102c9908363ffffffff61069216565b600160a060020a0380861660009081526001602052604080822093909355908516815220546102fe908363ffffffff6106a416565b600160a060020a03808516600090815260016020908152604080832094909455878316825260028152838220339093168252919091522054610346908363ffffffff61069216565b600160a060020a03808616600081815260026020908152604080832033861684529091529081902093909355908516917fddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef9085905190815260200160405180910390a35060019392505050565b600160a060020a0333811660009081526002602090815260408083209386168352929052908120548083111561041057600160a060020a033381166000908152600260209081526040808320938816835292905290812055610447565b610420818463ffffffff61069216565b600160a060020a033381166000908152600260209081526040808320938916835292905220555b600160a060020a0333811660008181526002602090815260408083209489168084529490915290819020547f8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925915190815260200160405180910390a35060019392505050565b600160a060020a031660009081526001602052604090205490565b6000600160a060020a03831615156104df57600080fd5b600160a060020a03331660009081526001602052604090205482111561050457600080fd5b600160a060020a03331660009081526001602052604090205461052d908363ffffffff61069216565b600160a060020a033381166000908152600160205260408082209390935590851681522054610562908363ffffffff6106a416565b600160a060020a0380851660008181526001602052604090819020939093559133909116907fddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef9085905190815260200160405180910390a350600192915050565b600160a060020a0333811660009081526002602090815260408083209386168352929052908120546105fb908363ffffffff6106a416565b600160a060020a0333811660008181526002602090815260408083209489168084529490915290819020849055919290917f8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b92591905190815260200160405180910390a350600192915050565b600160a060020a03918216600090815260026020908152604080832093909416825291909152205490565b60008282111561069e57fe5b50900390565b6000828201838110156106b357fe5b93925050505600a165627a7a723058208116fdca829779590809dc5e40c6eb2d5cbb4df9fc0bc5bc82ba3ed24c3043a90029`

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
