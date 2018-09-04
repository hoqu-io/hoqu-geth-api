// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package ad

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
	return address, tx, &ERC20{ERC20Caller: ERC20Caller{contract: contract}, ERC20Transactor: ERC20Transactor{contract: contract}, ERC20Filterer: ERC20Filterer{contract: contract}}, nil
}

// ERC20 is an auto generated Go binding around an Ethereum contract.
type ERC20 struct {
	ERC20Caller     // Read-only binding to the contract
	ERC20Transactor // Write-only binding to the contract
	ERC20Filterer   // Log filterer for contract events
}

// ERC20Caller is an auto generated read-only Go binding around an Ethereum contract.
type ERC20Caller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ERC20Transactor is an auto generated write-only Go binding around an Ethereum contract.
type ERC20Transactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ERC20Filterer is an auto generated log filtering Go binding around an Ethereum contract events.
type ERC20Filterer struct {
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
	contract, err := bindERC20(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &ERC20{ERC20Caller: ERC20Caller{contract: contract}, ERC20Transactor: ERC20Transactor{contract: contract}, ERC20Filterer: ERC20Filterer{contract: contract}}, nil
}

// NewERC20Caller creates a new read-only instance of ERC20, bound to a specific deployed contract.
func NewERC20Caller(address common.Address, caller bind.ContractCaller) (*ERC20Caller, error) {
	contract, err := bindERC20(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &ERC20Caller{contract: contract}, nil
}

// NewERC20Transactor creates a new write-only instance of ERC20, bound to a specific deployed contract.
func NewERC20Transactor(address common.Address, transactor bind.ContractTransactor) (*ERC20Transactor, error) {
	contract, err := bindERC20(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &ERC20Transactor{contract: contract}, nil
}

// NewERC20Filterer creates a new log filterer instance of ERC20, bound to a specific deployed contract.
func NewERC20Filterer(address common.Address, filterer bind.ContractFilterer) (*ERC20Filterer, error) {
	contract, err := bindERC20(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &ERC20Filterer{contract: contract}, nil
}

// bindERC20 binds a generic wrapper to an already deployed contract.
func bindERC20(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(ERC20ABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
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

// ERC20ApprovalIterator is returned from FilterApproval and is used to iterate over the raw logs and unpacked data for Approval events raised by the ERC20 contract.
type ERC20ApprovalIterator struct {
	Event *ERC20Approval // Event containing the contract specifics and raw log

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
func (it *ERC20ApprovalIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ERC20Approval)
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
		it.Event = new(ERC20Approval)
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
func (it *ERC20ApprovalIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ERC20ApprovalIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ERC20Approval represents a Approval event raised by the ERC20 contract.
type ERC20Approval struct {
	Owner   common.Address
	Spender common.Address
	Value   *big.Int
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterApproval is a free log retrieval operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: event Approval(owner indexed address, spender indexed address, value uint256)
func (_ERC20 *ERC20Filterer) FilterApproval(opts *bind.FilterOpts, owner []common.Address, spender []common.Address) (*ERC20ApprovalIterator, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var spenderRule []interface{}
	for _, spenderItem := range spender {
		spenderRule = append(spenderRule, spenderItem)
	}

	logs, sub, err := _ERC20.contract.FilterLogs(opts, "Approval", ownerRule, spenderRule)
	if err != nil {
		return nil, err
	}
	return &ERC20ApprovalIterator{contract: _ERC20.contract, event: "Approval", logs: logs, sub: sub}, nil
}

// WatchApproval is a free log subscription operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: event Approval(owner indexed address, spender indexed address, value uint256)
func (_ERC20 *ERC20Filterer) WatchApproval(opts *bind.WatchOpts, sink chan<- *ERC20Approval, owner []common.Address, spender []common.Address) (event.Subscription, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var spenderRule []interface{}
	for _, spenderItem := range spender {
		spenderRule = append(spenderRule, spenderItem)
	}

	logs, sub, err := _ERC20.contract.WatchLogs(opts, "Approval", ownerRule, spenderRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ERC20Approval)
				if err := _ERC20.contract.UnpackLog(event, "Approval", log); err != nil {
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

// ERC20TransferIterator is returned from FilterTransfer and is used to iterate over the raw logs and unpacked data for Transfer events raised by the ERC20 contract.
type ERC20TransferIterator struct {
	Event *ERC20Transfer // Event containing the contract specifics and raw log

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
func (it *ERC20TransferIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ERC20Transfer)
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
		it.Event = new(ERC20Transfer)
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
func (it *ERC20TransferIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ERC20TransferIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ERC20Transfer represents a Transfer event raised by the ERC20 contract.
type ERC20Transfer struct {
	From  common.Address
	To    common.Address
	Value *big.Int
	Raw   types.Log // Blockchain specific contextual infos
}

// FilterTransfer is a free log retrieval operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: event Transfer(from indexed address, to indexed address, value uint256)
func (_ERC20 *ERC20Filterer) FilterTransfer(opts *bind.FilterOpts, from []common.Address, to []common.Address) (*ERC20TransferIterator, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _ERC20.contract.FilterLogs(opts, "Transfer", fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return &ERC20TransferIterator{contract: _ERC20.contract, event: "Transfer", logs: logs, sub: sub}, nil
}

// WatchTransfer is a free log subscription operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: event Transfer(from indexed address, to indexed address, value uint256)
func (_ERC20 *ERC20Filterer) WatchTransfer(opts *bind.WatchOpts, sink chan<- *ERC20Transfer, from []common.Address, to []common.Address) (event.Subscription, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _ERC20.contract.WatchLogs(opts, "Transfer", fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ERC20Transfer)
				if err := _ERC20.contract.UnpackLog(event, "Transfer", log); err != nil {
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
	return address, tx, &ERC20Basic{ERC20BasicCaller: ERC20BasicCaller{contract: contract}, ERC20BasicTransactor: ERC20BasicTransactor{contract: contract}, ERC20BasicFilterer: ERC20BasicFilterer{contract: contract}}, nil
}

// ERC20Basic is an auto generated Go binding around an Ethereum contract.
type ERC20Basic struct {
	ERC20BasicCaller     // Read-only binding to the contract
	ERC20BasicTransactor // Write-only binding to the contract
	ERC20BasicFilterer   // Log filterer for contract events
}

// ERC20BasicCaller is an auto generated read-only Go binding around an Ethereum contract.
type ERC20BasicCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ERC20BasicTransactor is an auto generated write-only Go binding around an Ethereum contract.
type ERC20BasicTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ERC20BasicFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type ERC20BasicFilterer struct {
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
	contract, err := bindERC20Basic(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &ERC20Basic{ERC20BasicCaller: ERC20BasicCaller{contract: contract}, ERC20BasicTransactor: ERC20BasicTransactor{contract: contract}, ERC20BasicFilterer: ERC20BasicFilterer{contract: contract}}, nil
}

// NewERC20BasicCaller creates a new read-only instance of ERC20Basic, bound to a specific deployed contract.
func NewERC20BasicCaller(address common.Address, caller bind.ContractCaller) (*ERC20BasicCaller, error) {
	contract, err := bindERC20Basic(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &ERC20BasicCaller{contract: contract}, nil
}

// NewERC20BasicTransactor creates a new write-only instance of ERC20Basic, bound to a specific deployed contract.
func NewERC20BasicTransactor(address common.Address, transactor bind.ContractTransactor) (*ERC20BasicTransactor, error) {
	contract, err := bindERC20Basic(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &ERC20BasicTransactor{contract: contract}, nil
}

// NewERC20BasicFilterer creates a new log filterer instance of ERC20Basic, bound to a specific deployed contract.
func NewERC20BasicFilterer(address common.Address, filterer bind.ContractFilterer) (*ERC20BasicFilterer, error) {
	contract, err := bindERC20Basic(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &ERC20BasicFilterer{contract: contract}, nil
}

// bindERC20Basic binds a generic wrapper to an already deployed contract.
func bindERC20Basic(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(ERC20BasicABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
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

// ERC20BasicTransferIterator is returned from FilterTransfer and is used to iterate over the raw logs and unpacked data for Transfer events raised by the ERC20Basic contract.
type ERC20BasicTransferIterator struct {
	Event *ERC20BasicTransfer // Event containing the contract specifics and raw log

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
func (it *ERC20BasicTransferIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ERC20BasicTransfer)
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
		it.Event = new(ERC20BasicTransfer)
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
func (it *ERC20BasicTransferIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ERC20BasicTransferIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ERC20BasicTransfer represents a Transfer event raised by the ERC20Basic contract.
type ERC20BasicTransfer struct {
	From  common.Address
	To    common.Address
	Value *big.Int
	Raw   types.Log // Blockchain specific contextual infos
}

// FilterTransfer is a free log retrieval operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: event Transfer(from indexed address, to indexed address, value uint256)
func (_ERC20Basic *ERC20BasicFilterer) FilterTransfer(opts *bind.FilterOpts, from []common.Address, to []common.Address) (*ERC20BasicTransferIterator, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _ERC20Basic.contract.FilterLogs(opts, "Transfer", fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return &ERC20BasicTransferIterator{contract: _ERC20Basic.contract, event: "Transfer", logs: logs, sub: sub}, nil
}

// WatchTransfer is a free log subscription operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: event Transfer(from indexed address, to indexed address, value uint256)
func (_ERC20Basic *ERC20BasicFilterer) WatchTransfer(opts *bind.WatchOpts, sink chan<- *ERC20BasicTransfer, from []common.Address, to []common.Address) (event.Subscription, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _ERC20Basic.contract.WatchLogs(opts, "Transfer", fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ERC20BasicTransfer)
				if err := _ERC20Basic.contract.UnpackLog(event, "Transfer", log); err != nil {
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

// HoQuAdCampaignABI is the input ABI used to generate the binding from.
const HoQuAdCampaignABI = "[{\"constant\":true,\"inputs\":[],\"name\":\"status\",\"outputs\":[{\"name\":\"\",\"type\":\"uint8\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"offerId\",\"outputs\":[{\"name\":\"\",\"type\":\"bytes16\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"tokenAddress\",\"type\":\"address\"}],\"name\":\"setTokenAddress\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"payerAddress\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_status\",\"type\":\"uint8\"}],\"name\":\"setStatus\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"rater\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"id\",\"type\":\"bytes16\"},{\"name\":\"intermediaryAddress\",\"type\":\"address\"},{\"name\":\"percent\",\"type\":\"uint32\"}],\"name\":\"addLeadIntermediary\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"affiliateId\",\"outputs\":[{\"name\":\"\",\"type\":\"bytes16\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_payerAddress\",\"type\":\"address\"}],\"name\":\"setPayerAddress\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"storageAddress\",\"type\":\"address\"}],\"name\":\"setStorageAddress\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"id\",\"type\":\"bytes16\"},{\"name\":\"num\",\"type\":\"uint8\"}],\"name\":\"getLeadIntermediaryPercent\",\"outputs\":[{\"name\":\"\",\"type\":\"uint32\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"ownerAddress\",\"type\":\"address\"},{\"name\":\"id\",\"type\":\"bytes16\"}],\"name\":\"addTracker\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"\",\"type\":\"address\"}],\"name\":\"trackers\",\"outputs\":[{\"name\":\"\",\"type\":\"bytes16\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"config\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"id\",\"type\":\"bytes16\"},{\"name\":\"trackerId\",\"type\":\"bytes16\"},{\"name\":\"meta\",\"type\":\"string\"},{\"name\":\"dataUrl\",\"type\":\"string\"},{\"name\":\"price\",\"type\":\"uint256\"}],\"name\":\"addLead\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"id\",\"type\":\"bytes16\"},{\"name\":\"_status\",\"type\":\"uint8\"}],\"name\":\"setLeadStatus\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"id\",\"type\":\"bytes16\"}],\"name\":\"transactLead\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"configAddress\",\"type\":\"address\"}],\"name\":\"setConfigAddress\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"store\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"adId\",\"outputs\":[{\"name\":\"\",\"type\":\"bytes16\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"raterAddress\",\"type\":\"address\"}],\"name\":\"setRaterAddress\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"beneficiaryAddress\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"id\",\"type\":\"bytes16\"},{\"name\":\"num\",\"type\":\"uint8\"}],\"name\":\"getLeadIntermediaryAddress\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"\",\"type\":\"bytes16\"}],\"name\":\"leads\",\"outputs\":[{\"name\":\"createdAt\",\"type\":\"uint256\"},{\"name\":\"trackerId\",\"type\":\"bytes16\"},{\"name\":\"dataUrl\",\"type\":\"string\"},{\"name\":\"meta\",\"type\":\"string\"},{\"name\":\"price\",\"type\":\"uint256\"},{\"name\":\"numOfIntermediaries\",\"type\":\"uint8\"},{\"name\":\"status\",\"type\":\"uint8\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_beneficiaryAddress\",\"type\":\"address\"}],\"name\":\"setBeneficiaryAddress\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"token\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"name\":\"configAddress\",\"type\":\"address\"},{\"name\":\"tokenAddress\",\"type\":\"address\"},{\"name\":\"storageAddress\",\"type\":\"address\"},{\"name\":\"raterAddress\",\"type\":\"address\"},{\"name\":\"_adId\",\"type\":\"bytes16\"},{\"name\":\"_offerId\",\"type\":\"bytes16\"},{\"name\":\"_affiliateId\",\"type\":\"bytes16\"},{\"name\":\"_beneficiaryAddress\",\"type\":\"address\"},{\"name\":\"_payerAddress\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"payerAddress\",\"type\":\"address\"},{\"indexed\":false,\"name\":\"newStatus\",\"type\":\"uint8\"}],\"name\":\"StatusChanged\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"beneficiaryAddress\",\"type\":\"address\"},{\"indexed\":true,\"name\":\"newBeneficiaryAddress\",\"type\":\"address\"}],\"name\":\"BeneficiaryAddressChanged\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"payerAddress\",\"type\":\"address\"},{\"indexed\":true,\"name\":\"newPayerAddress\",\"type\":\"address\"}],\"name\":\"PayerAddressChanged\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"beneficiaryAddress\",\"type\":\"address\"},{\"indexed\":false,\"name\":\"id\",\"type\":\"bytes16\"},{\"indexed\":false,\"name\":\"price\",\"type\":\"uint256\"}],\"name\":\"LeadAdded\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"beneficiaryAddress\",\"type\":\"address\"},{\"indexed\":false,\"name\":\"id\",\"type\":\"bytes16\"},{\"indexed\":false,\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"LeadTransacted\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"beneficiaryAddress\",\"type\":\"address\"},{\"indexed\":false,\"name\":\"id\",\"type\":\"bytes16\"},{\"indexed\":false,\"name\":\"status\",\"type\":\"uint8\"}],\"name\":\"LeadStatusChanged\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"ownerAddress\",\"type\":\"address\"},{\"indexed\":false,\"name\":\"id\",\"type\":\"bytes16\"}],\"name\":\"TrackerAdded\",\"type\":\"event\"}]"

// HoQuAdCampaignBin is the compiled bytecode used for deploying new contracts.
const HoQuAdCampaignBin = `0x608060405234801561001057600080fd5b50604051610120806122bb83398101604090815281516020830151918301516060840151608085015160a086015160c087015160e08801516101009098015160008054600160a060020a03988916600160a060020a03199182161790915560018054998916998216999099179098556002805496881696891696909617909555600380549487169488169490941790935560048054700100000000000000000000000000000000928390048302938390046001608060020a0319918216176001608060020a0316939093179055600580549190930491161790556006805494831694841694909417909355600780549190931691161790556121a4806101176000396000f3006080604052600436106101535763ffffffff7c0100000000000000000000000000000000000000000000000000000000600035041663200d2ed2811461015857806321ce9f911461019157806326a4e8d2146101c35780632b2f1e14146101e65780632e49d78b14610217578063355c9fb11461023257806336617045146102475780633f8b35a61461027e5780634b28676d1461029357806359b910d6146102b45780635cb8f5ab146102d55780636aa0082c146103165780636cb234a31461034457806379502c55146103655780637a8187cf1461037a57806381e85aa31461042a578063835aef401461045257806383a12de914610474578063975057e714610495578063cbd76c35146104aa578063ce88f04e146104bf578063d9c4870e146104e0578063da7345b7146104f5578063e7608fb21461051d578063ec6be06e1461065c578063fc0c546a1461067d575b600080fd5b34801561016457600080fd5b5061016d610692565b6040518082600581111561017d57fe5b60ff16815260200191505060405180910390f35b34801561019d57600080fd5b506101a66106b3565b604080516001608060020a03199092168252519081900360200190f35b3480156101cf57600080fd5b506101e4600160a060020a03600435166106c3565b005b3480156101f257600080fd5b506101fb610771565b60408051600160a060020a039092168252519081900360200190f35b34801561022357600080fd5b506101e460ff60043516610780565b34801561023e57600080fd5b506101fb610a3a565b34801561025357600080fd5b506101e46001608060020a031960043516600160a060020a036024351663ffffffff60443516610a49565b34801561028a57600080fd5b506101a6610bc3565b34801561029f57600080fd5b506101e4600160a060020a0360043516610bcf565b3480156102c057600080fd5b506101e4600160a060020a0360043516610cb7565b3480156102e157600080fd5b506102fd6001608060020a03196004351660ff60243516610d65565b6040805163ffffffff9092168252519081900360200190f35b34801561032257600080fd5b506101e4600160a060020a03600435166001608060020a031960243516610ddc565b34801561035057600080fd5b506101a6600160a060020a0360043516610ed4565b34801561037157600080fd5b506101fb610eec565b34801561038657600080fd5b50604080516020600460443581810135601f81018490048402850184019095528484526101e49482356001608060020a03199081169560248035909216953695946064949293019190819084018382808284375050604080516020601f89358b018035918201839004830284018301909452808352979a9998810197919650918201945092508291508401838280828437509497505093359450610efb9350505050565b34801561043657600080fd5b506101e46001608060020a03196004351660ff60243516611354565b34801561045e57600080fd5b506101e46001608060020a0319600435166114df565b34801561048057600080fd5b506101e4600160a060020a0360043516611c3b565b3480156104a157600080fd5b506101fb611ce9565b3480156104b657600080fd5b506101a6611cf8565b3480156104cb57600080fd5b506101e4600160a060020a0360043516611d04565b3480156104ec57600080fd5b506101fb611db2565b34801561050157600080fd5b506101fb6001608060020a03196004351660ff60243516611dc1565b34801561052957600080fd5b5061053f6001608060020a031960043516611e3a565b604080518881526001608060020a0319881660208201526080810185905260ff841660a08201529081016060820160c0830184600581111561057d57fe5b60ff168152602001838103835288818151815260200191508051906020019080838360005b838110156105ba5781810151838201526020016105a2565b50505050905090810190601f1680156105e75780820380516001836020036101000a031916815260200191505b50838103825287518152875160209182019189019080838360005b8381101561061a578181015183820152602001610602565b50505050905090810190601f1680156106475780820380516001836020036101000a031916815260200191505b50995050505050505050505060405180910390f35b34801561066857600080fd5b506101e4600160a060020a0360043516611f96565b34801561068957600080fd5b506101fb61207e565b60075474010000000000000000000000000000000000000000900460ff1681565b600454608060020a908190040281565b600080546040805160e060020a63babcc539028152600160a060020a0333811660048301529151919092169263babcc53992602480820193602093909283900390910190829087803b15801561071857600080fd5b505af115801561072c573d6000803e3d6000fd5b505050506040513d602081101561074257600080fd5b5051151561074f57600080fd5b60018054600160a060020a031916600160a060020a0392909216919091179055565b600754600160a060020a031681565b600080546040805160e060020a63babcc539028152600160a060020a0333811660048301529151919092169163babcc53991602480830192602092919082900301818787803b1580156107d257600080fd5b505af11580156107e6573d6000803e3d6000fd5b505050506040513d60208110156107fc57600080fd5b5051151561080957600080fd5b60025460048054604080517f89075557000000000000000000000000000000000000000000000000000000008152608060020a9092026001608060020a031916928201929092529051600160a060020a039092169163890755579160248082019260a0929091908290030181600087803b15801561088657600080fd5b505af115801561089a573d6000803e3d6000fd5b505050506040513d60a08110156108b057600080fd5b5060800151905060008160058111156108c557fe5b14156108d057600080fd5b600254600480546040517fa0191d200000000000000000000000000000000000000000000000000000000081526001608060020a0319608060020a9092029182169281019283526000602482018190526044820181905260648201819052600160a060020a039094169363a0191d209390918291829189919060840182600581111561095857fe5b60ff16815260200195505050505050600060405180830381600087803b15801561098157600080fd5b505af1158015610995573d6000803e3d6000fd5b50506007805485935090915074ff00000000000000000000000000000000000000001916740100000000000000000000000000000000000000008360058111156109db57fe5b0217905550600754604051600160a060020a03909116907f76c0bad3b7f1c74587dcb478154d4c87ae6663db97cc086b981decf4e35a797c90849080826005811115610a2357fe5b60ff16815260200191505060405180910390a25050565b600354600160a060020a031681565b600080546040805160e060020a63babcc539028152600160a060020a0333811660048301529151919092169263babcc53992602480820193602093909283900390910190829087803b158015610a9e57600080fd5b505af1158015610ab2573d6000803e3d6000fd5b505050506040513d6020811015610ac857600080fd5b505180610ae55750600160a060020a033316600052600960205260015b1515610af057600080fd5b60006001608060020a03198416600090815260086020526040902060070154610100900460ff166005811115610b2257fe5b1415610b2d57600080fd5b6001608060020a031992909216600081815260086020818152604080842060078101805460ff90811687526005830185528387208054600160a060020a031916600160a060020a039a909a169990991790985580548816865260069091018352908420805463ffffffff191663ffffffff90981697909717909655929091529052815460ff198116908216600101909116179055565b600554608060020a0281565b600080546040805160e060020a63babcc539028152600160a060020a0333811660048301529151919092169263babcc53992602480820193602093909283900390910190829087803b158015610c2457600080fd5b505af1158015610c38573d6000803e3d6000fd5b505050506040513d6020811015610c4e57600080fd5b50511515610c5b57600080fd5b600754604051600160a060020a038084169216907f650742bb22425d77490f0e51db9824c857c049c9fc9fee93b369553acb17275890600090a360078054600160a060020a031916600160a060020a0392909216919091179055565b600080546040805160e060020a63babcc539028152600160a060020a0333811660048301529151919092169263babcc53992602480820193602093909283900390910190829087803b158015610d0c57600080fd5b505af1158015610d20573d6000803e3d6000fd5b505050506040513d6020811015610d3657600080fd5b50511515610d4357600080fd5b60028054600160a060020a031916600160a060020a0392909216919091179055565b6000806001608060020a03198416600090815260086020526040902060070154610100900460ff166005811115610d9857fe5b1415610da357600080fd5b506001608060020a03198216600090815260086020908152604080832060ff8516845260060190915290205463ffffffff165b92915050565b600080546040805160e060020a63babcc539028152600160a060020a0333811660048301529151919092169263babcc53992602480820193602093909283900390910190829087803b158015610e3157600080fd5b505af1158015610e45573d6000803e3d6000fd5b505050506040513d6020811015610e5b57600080fd5b50511515610e6857600080fd5b600160a060020a0382166000818152600960209081526040918290208054608060020a86046001608060020a0319918216179091558251908516815291517fd7793ff1544e44df2a307af241d282820620c015d2e3b8cb17b480a7e47a48be9281900390910190a25050565b600960205260009081526040902054608060020a0281565b600054600160a060020a031681565b600080546040805160e060020a63babcc539028152600160a060020a0333811660048301529151919092169163babcc53991602480830192602092919082900301818787803b158015610f4d57600080fd5b505af1158015610f61573d6000803e3d6000fd5b505050506040513d6020811015610f7757600080fd5b505180610f945750600160a060020a033316600052600960205260015b1515610f9f57600080fd5b60006001608060020a03198716600090815260086020526040902060070154610100900460ff166005811115610fd157fe5b14610fdb57600080fd5b600254604080517fea2fdd8c0000000000000000000000000000000000000000000000000000000081526001608060020a0319881660048201529051600160a060020a039092169163ea2fdd8c9160248082019260009290919082900301818387803b15801561104a57600080fd5b505af115801561105e573d6000803e3d6000fd5b505050506040513d6000823e601f3d908101601f1916820160405260c081101561108757600080fd5b81516020830151604084018051929491938201926401000000008111156110ad57600080fd5b820160208101848111156110c057600080fd5b81516401000000008111828201871017156110da57600080fd5b505092919060200180516401000000008111156110f657600080fd5b8201602081018481111561110957600080fd5b815164010000000081118282018710171561112357600080fd5b5050506040015195506000945061113a9350505050565b81600581111561114657fe5b141561115157600080fd5b6040805160e0810182524281526001608060020a0319871660208201529081018490526060810185905260808101839052600060a082015260c08101600190526001608060020a03198781166000908152600860209081526040918290208451815584820151600182018054909516608060020a909104179093559083015180516111e292600285019201906120dd565b50606082015180516111fe9160038401916020909101906120dd565b506080820151600482015560a082015160078201805460ff191660ff9092169190911780825560c0840151919061ff00191661010083600581111561123f57fe5b02179055505060035460048054600554604080517fc5f6dd850000000000000000000000000000000000000000000000000000000081526001608060020a0319608060020a948590048502811695820195909552848c166024820152919092029290921660448301526064820186905251600160a060020a03909216925063c5f6dd8591608480830192600092919082900301818387803b1580156112e357600080fd5b505af11580156112f7573d6000803e3d6000fd5b5050600654604080516001608060020a03198b168152602081018790528151600160a060020a0390931694507f5a4f5a43985de2c0807c6875a069d25bfdf90255dc58958a49e07d00936537e393508290030190a2505050505050565b600080546040805160e060020a63babcc539028152600160a060020a0333811660048301529151919092169263babcc53992602480820193602093909283900390910190829087803b1580156113a957600080fd5b505af11580156113bd573d6000803e3d6000fd5b505050506040513d60208110156113d357600080fd5b5051806113f05750600160a060020a033316600052600960205260015b15156113fb57600080fd5b60006001608060020a03198316600090815260086020526040902060070154610100900460ff16600581111561142d57fe5b141561143857600080fd5b6001608060020a031982166000908152600860205260409020600701805482919061ff00191661010083600581111561146d57fe5b02179055506006546040516001608060020a031984168152600160a060020a03909116907ffa2a86784442baf24e238559b3ca5986aa4a588eeb7840c5ba8b27d6e663ecd99084908490602081018260058111156114c757fe5b60ff1681526020019250505060405180910390a25050565b600080546040805160e060020a63babcc539028152600160a060020a0333811660048301529151849384938493849384938493169163babcc53991602480830192602092919082900301818787803b15801561153a57600080fd5b505af115801561154e573d6000803e3d6000fd5b505050506040513d602081101561156457600080fd5b5051806115815750600160a060020a033316600052600960205260015b151561158c57600080fd5b60046001608060020a03198916600090815260086020526040902060070154610100900460ff1660058111156115be57fe5b141580156115fc575060056001608060020a03198916600090815260086020526040902060070154610100900460ff1660058111156115f957fe5b14155b151561160757600080fd5b6001608060020a031988166000908152600860205260408120600401541161162e57600080fd5b6001608060020a03198816600090815260086020908152604080832060078101805461ff001916610400179055835482517fe14891910000000000000000000000000000000000000000000000000000000081529251919b5061171194670de0b6b3a76400009461170594600160a060020a039093169363e148919193600480820194918390030190829087803b1580156116c857600080fd5b505af11580156116dc573d6000803e3d6000fd5b505050506040513d60208110156116f257600080fd5b505160048b01549063ffffffff61208d16565b9063ffffffff6120b616565b6004880154909650611729908763ffffffff6120cb16565b6001546007546004808b0154604080517f23b872dd000000000000000000000000000000000000000000000000000000008152600160a060020a039485169381019390935230841660248401526044830191909152519398509116916323b872dd916064808201926020929091908290030181600087803b1580156117ad57600080fd5b505af11580156117c1573d6000803e3d6000fd5b505050506040513d60208110156117d757600080fd5b505060015460008054604080517f97c0262a0000000000000000000000000000000000000000000000000000000081529051600160a060020a039485169463a9059cbb949316926397c0262a92600480820193602093909283900390910190829087803b15801561184757600080fd5b505af115801561185b573d6000803e3d6000fd5b505050506040513d602081101561187157600080fd5b5051604080517c010000000000000000000000000000000000000000000000000000000063ffffffff8516028152600160a060020a039092166004830152602482018a90525160448083019260209291908290030181600087803b1580156118d857600080fd5b505af11580156118ec573d6000803e3d6000fd5b505050506040513d602081101561190257600080fd5b5060009450505b600787015460ff9081169085161015611a2d5760ff8416600090815260058801602090815260408083205460068b01909252909120546004890154600160a060020a03909216945063ffffffff9081169350611972916305f5e100916117059190869061208d16565b600154604080517fa9059cbb000000000000000000000000000000000000000000000000000000008152600160a060020a03878116600483015260248201859052915193945091169163a9059cbb916044808201926020929091908290030181600087803b1580156119e357600080fd5b505af11580156119f7573d6000803e3d6000fd5b505050506040513d6020811015611a0d57600080fd5b50611a209050858263ffffffff6120cb16565b9450600190930192611909565b600154600654604080517fa9059cbb000000000000000000000000000000000000000000000000000000008152600160a060020a039283166004820152602481018990529051919092169163a9059cbb9160448083019260209291908290030181600087803b158015611a9f57600080fd5b505af1158015611ab3573d6000803e3d6000fd5b505050506040513d6020811015611ac957600080fd5b50506003546004805460018a01546005548b840154604080517f4943df380000000000000000000000000000000000000000000000000000000081526001608060020a0319608060020a96879004870281169782019790975293850286166024850152939091029093166044820152606481019290925251600160a060020a0390921691634943df389160848082019260009290919082900301818387803b158015611b7457600080fd5b505af1158015611b88573d6000803e3d6000fd5b5050600654604080516001608060020a03198d168152602081018a90528151600160a060020a0390931694507f8e1b43b752aaaf23ee6b52413a1dd24c21f4571e7ff32d88456e52befde6dd0493508290030190a2600654604080516001608060020a03198b168152600460208201528151600160a060020a03909316927ffa2a86784442baf24e238559b3ca5986aa4a588eeb7840c5ba8b27d6e663ecd9929181900390910190a25050505050505050565b600080546040805160e060020a63babcc539028152600160a060020a0333811660048301529151919092169263babcc53992602480820193602093909283900390910190829087803b158015611c9057600080fd5b505af1158015611ca4573d6000803e3d6000fd5b505050506040513d6020811015611cba57600080fd5b50511515611cc757600080fd5b60008054600160a060020a031916600160a060020a0392909216919091179055565b600254600160a060020a031681565b600454608060020a0281565b600080546040805160e060020a63babcc539028152600160a060020a0333811660048301529151919092169263babcc53992602480820193602093909283900390910190829087803b158015611d5957600080fd5b505af1158015611d6d573d6000803e3d6000fd5b505050506040513d6020811015611d8357600080fd5b50511515611d9057600080fd5b60038054600160a060020a031916600160a060020a0392909216919091179055565b600654600160a060020a031681565b6000806001608060020a03198416600090815260086020526040902060070154610100900460ff166005811115611df457fe5b1415611dff57600080fd5b506001608060020a03198216600090815260086020908152604080832060ff85168452600501909152902054600160a060020a031692915050565b60086020908152600091825260409182902080546001808301546002808501805488516101009582161595909502600019011691909104601f81018790048702840187019097528683529295608060020a90910294919291830182828015611ee35780601f10611eb857610100808354040283529160200191611ee3565b820191906000526020600020905b815481529060010190602001808311611ec657829003601f168201915b5050505060038301805460408051602060026001851615610100026000190190941693909304601f8101849004840282018401909252818152949594935090830182828015611f735780601f10611f4857610100808354040283529160200191611f73565b820191906000526020600020905b815481529060010190602001808311611f5657829003601f168201915b50505050600483015460079093015491929160ff80821692506101009091041687565b600080546040805160e060020a63babcc539028152600160a060020a0333811660048301529151919092169263babcc53992602480820193602093909283900390910190829087803b158015611feb57600080fd5b505af1158015611fff573d6000803e3d6000fd5b505050506040513d602081101561201557600080fd5b5051151561202257600080fd5b600654604051600160a060020a038084169216907fb53b7057ddbff0cd9111b0a1a501a9cc1a9c315eb3396cb0d47917908325999a90600090a360068054600160a060020a031916600160a060020a0392909216919091179055565b600154600160a060020a031681565b600082151561209e57506000610dd6565b508181028183828115156120ae57fe5b0414610dd657fe5b600081838115156120c357fe5b049392505050565b6000828211156120d757fe5b50900390565b828054600181600116156101000203166002900490600052602060002090601f016020900481019282601f1061211e57805160ff191683800117855561214b565b8280016001018555821561214b579182015b8281111561214b578251825591602001919060010190612130565b5061215792915061215b565b5090565b61217591905b808211156121575760008155600101612161565b905600a165627a7a72305820e2491bdb9ba1babe301415d84ea6de7d01a2b839b91462b23d21f5154b79577a0029`

// DeployHoQuAdCampaign deploys a new Ethereum contract, binding an instance of HoQuAdCampaign to it.
func DeployHoQuAdCampaign(auth *bind.TransactOpts, backend bind.ContractBackend, configAddress common.Address, tokenAddress common.Address, storageAddress common.Address, raterAddress common.Address, _adId [16]byte, _offerId [16]byte, _affiliateId [16]byte, _beneficiaryAddress common.Address, _payerAddress common.Address) (common.Address, *types.Transaction, *HoQuAdCampaign, error) {
	parsed, err := abi.JSON(strings.NewReader(HoQuAdCampaignABI))
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	address, tx, contract, err := bind.DeployContract(auth, parsed, common.FromHex(HoQuAdCampaignBin), backend, configAddress, tokenAddress, storageAddress, raterAddress, _adId, _offerId, _affiliateId, _beneficiaryAddress, _payerAddress)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &HoQuAdCampaign{HoQuAdCampaignCaller: HoQuAdCampaignCaller{contract: contract}, HoQuAdCampaignTransactor: HoQuAdCampaignTransactor{contract: contract}, HoQuAdCampaignFilterer: HoQuAdCampaignFilterer{contract: contract}}, nil
}

// HoQuAdCampaign is an auto generated Go binding around an Ethereum contract.
type HoQuAdCampaign struct {
	HoQuAdCampaignCaller     // Read-only binding to the contract
	HoQuAdCampaignTransactor // Write-only binding to the contract
	HoQuAdCampaignFilterer   // Log filterer for contract events
}

// HoQuAdCampaignCaller is an auto generated read-only Go binding around an Ethereum contract.
type HoQuAdCampaignCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// HoQuAdCampaignTransactor is an auto generated write-only Go binding around an Ethereum contract.
type HoQuAdCampaignTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// HoQuAdCampaignFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type HoQuAdCampaignFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// HoQuAdCampaignSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type HoQuAdCampaignSession struct {
	Contract     *HoQuAdCampaign   // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// HoQuAdCampaignCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type HoQuAdCampaignCallerSession struct {
	Contract *HoQuAdCampaignCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts         // Call options to use throughout this session
}

// HoQuAdCampaignTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type HoQuAdCampaignTransactorSession struct {
	Contract     *HoQuAdCampaignTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts         // Transaction auth options to use throughout this session
}

// HoQuAdCampaignRaw is an auto generated low-level Go binding around an Ethereum contract.
type HoQuAdCampaignRaw struct {
	Contract *HoQuAdCampaign // Generic contract binding to access the raw methods on
}

// HoQuAdCampaignCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type HoQuAdCampaignCallerRaw struct {
	Contract *HoQuAdCampaignCaller // Generic read-only contract binding to access the raw methods on
}

// HoQuAdCampaignTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type HoQuAdCampaignTransactorRaw struct {
	Contract *HoQuAdCampaignTransactor // Generic write-only contract binding to access the raw methods on
}

// NewHoQuAdCampaign creates a new instance of HoQuAdCampaign, bound to a specific deployed contract.
func NewHoQuAdCampaign(address common.Address, backend bind.ContractBackend) (*HoQuAdCampaign, error) {
	contract, err := bindHoQuAdCampaign(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &HoQuAdCampaign{HoQuAdCampaignCaller: HoQuAdCampaignCaller{contract: contract}, HoQuAdCampaignTransactor: HoQuAdCampaignTransactor{contract: contract}, HoQuAdCampaignFilterer: HoQuAdCampaignFilterer{contract: contract}}, nil
}

// NewHoQuAdCampaignCaller creates a new read-only instance of HoQuAdCampaign, bound to a specific deployed contract.
func NewHoQuAdCampaignCaller(address common.Address, caller bind.ContractCaller) (*HoQuAdCampaignCaller, error) {
	contract, err := bindHoQuAdCampaign(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &HoQuAdCampaignCaller{contract: contract}, nil
}

// NewHoQuAdCampaignTransactor creates a new write-only instance of HoQuAdCampaign, bound to a specific deployed contract.
func NewHoQuAdCampaignTransactor(address common.Address, transactor bind.ContractTransactor) (*HoQuAdCampaignTransactor, error) {
	contract, err := bindHoQuAdCampaign(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &HoQuAdCampaignTransactor{contract: contract}, nil
}

// NewHoQuAdCampaignFilterer creates a new log filterer instance of HoQuAdCampaign, bound to a specific deployed contract.
func NewHoQuAdCampaignFilterer(address common.Address, filterer bind.ContractFilterer) (*HoQuAdCampaignFilterer, error) {
	contract, err := bindHoQuAdCampaign(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &HoQuAdCampaignFilterer{contract: contract}, nil
}

// bindHoQuAdCampaign binds a generic wrapper to an already deployed contract.
func bindHoQuAdCampaign(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(HoQuAdCampaignABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_HoQuAdCampaign *HoQuAdCampaignRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _HoQuAdCampaign.Contract.HoQuAdCampaignCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_HoQuAdCampaign *HoQuAdCampaignRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _HoQuAdCampaign.Contract.HoQuAdCampaignTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_HoQuAdCampaign *HoQuAdCampaignRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _HoQuAdCampaign.Contract.HoQuAdCampaignTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_HoQuAdCampaign *HoQuAdCampaignCallerRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _HoQuAdCampaign.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_HoQuAdCampaign *HoQuAdCampaignTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _HoQuAdCampaign.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_HoQuAdCampaign *HoQuAdCampaignTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _HoQuAdCampaign.Contract.contract.Transact(opts, method, params...)
}

// AdId is a free data retrieval call binding the contract method 0xcbd76c35.
//
// Solidity: function adId() constant returns(bytes16)
func (_HoQuAdCampaign *HoQuAdCampaignCaller) AdId(opts *bind.CallOpts) ([16]byte, error) {
	var (
		ret0 = new([16]byte)
	)
	out := ret0
	err := _HoQuAdCampaign.contract.Call(opts, out, "adId")
	return *ret0, err
}

// AdId is a free data retrieval call binding the contract method 0xcbd76c35.
//
// Solidity: function adId() constant returns(bytes16)
func (_HoQuAdCampaign *HoQuAdCampaignSession) AdId() ([16]byte, error) {
	return _HoQuAdCampaign.Contract.AdId(&_HoQuAdCampaign.CallOpts)
}

// AdId is a free data retrieval call binding the contract method 0xcbd76c35.
//
// Solidity: function adId() constant returns(bytes16)
func (_HoQuAdCampaign *HoQuAdCampaignCallerSession) AdId() ([16]byte, error) {
	return _HoQuAdCampaign.Contract.AdId(&_HoQuAdCampaign.CallOpts)
}

// AffiliateId is a free data retrieval call binding the contract method 0x3f8b35a6.
//
// Solidity: function affiliateId() constant returns(bytes16)
func (_HoQuAdCampaign *HoQuAdCampaignCaller) AffiliateId(opts *bind.CallOpts) ([16]byte, error) {
	var (
		ret0 = new([16]byte)
	)
	out := ret0
	err := _HoQuAdCampaign.contract.Call(opts, out, "affiliateId")
	return *ret0, err
}

// AffiliateId is a free data retrieval call binding the contract method 0x3f8b35a6.
//
// Solidity: function affiliateId() constant returns(bytes16)
func (_HoQuAdCampaign *HoQuAdCampaignSession) AffiliateId() ([16]byte, error) {
	return _HoQuAdCampaign.Contract.AffiliateId(&_HoQuAdCampaign.CallOpts)
}

// AffiliateId is a free data retrieval call binding the contract method 0x3f8b35a6.
//
// Solidity: function affiliateId() constant returns(bytes16)
func (_HoQuAdCampaign *HoQuAdCampaignCallerSession) AffiliateId() ([16]byte, error) {
	return _HoQuAdCampaign.Contract.AffiliateId(&_HoQuAdCampaign.CallOpts)
}

// BeneficiaryAddress is a free data retrieval call binding the contract method 0xd9c4870e.
//
// Solidity: function beneficiaryAddress() constant returns(address)
func (_HoQuAdCampaign *HoQuAdCampaignCaller) BeneficiaryAddress(opts *bind.CallOpts) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _HoQuAdCampaign.contract.Call(opts, out, "beneficiaryAddress")
	return *ret0, err
}

// BeneficiaryAddress is a free data retrieval call binding the contract method 0xd9c4870e.
//
// Solidity: function beneficiaryAddress() constant returns(address)
func (_HoQuAdCampaign *HoQuAdCampaignSession) BeneficiaryAddress() (common.Address, error) {
	return _HoQuAdCampaign.Contract.BeneficiaryAddress(&_HoQuAdCampaign.CallOpts)
}

// BeneficiaryAddress is a free data retrieval call binding the contract method 0xd9c4870e.
//
// Solidity: function beneficiaryAddress() constant returns(address)
func (_HoQuAdCampaign *HoQuAdCampaignCallerSession) BeneficiaryAddress() (common.Address, error) {
	return _HoQuAdCampaign.Contract.BeneficiaryAddress(&_HoQuAdCampaign.CallOpts)
}

// Config is a free data retrieval call binding the contract method 0x79502c55.
//
// Solidity: function config() constant returns(address)
func (_HoQuAdCampaign *HoQuAdCampaignCaller) Config(opts *bind.CallOpts) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _HoQuAdCampaign.contract.Call(opts, out, "config")
	return *ret0, err
}

// Config is a free data retrieval call binding the contract method 0x79502c55.
//
// Solidity: function config() constant returns(address)
func (_HoQuAdCampaign *HoQuAdCampaignSession) Config() (common.Address, error) {
	return _HoQuAdCampaign.Contract.Config(&_HoQuAdCampaign.CallOpts)
}

// Config is a free data retrieval call binding the contract method 0x79502c55.
//
// Solidity: function config() constant returns(address)
func (_HoQuAdCampaign *HoQuAdCampaignCallerSession) Config() (common.Address, error) {
	return _HoQuAdCampaign.Contract.Config(&_HoQuAdCampaign.CallOpts)
}

// GetLeadIntermediaryAddress is a free data retrieval call binding the contract method 0xda7345b7.
//
// Solidity: function getLeadIntermediaryAddress(id bytes16, num uint8) constant returns(address)
func (_HoQuAdCampaign *HoQuAdCampaignCaller) GetLeadIntermediaryAddress(opts *bind.CallOpts, id [16]byte, num uint8) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _HoQuAdCampaign.contract.Call(opts, out, "getLeadIntermediaryAddress", id, num)
	return *ret0, err
}

// GetLeadIntermediaryAddress is a free data retrieval call binding the contract method 0xda7345b7.
//
// Solidity: function getLeadIntermediaryAddress(id bytes16, num uint8) constant returns(address)
func (_HoQuAdCampaign *HoQuAdCampaignSession) GetLeadIntermediaryAddress(id [16]byte, num uint8) (common.Address, error) {
	return _HoQuAdCampaign.Contract.GetLeadIntermediaryAddress(&_HoQuAdCampaign.CallOpts, id, num)
}

// GetLeadIntermediaryAddress is a free data retrieval call binding the contract method 0xda7345b7.
//
// Solidity: function getLeadIntermediaryAddress(id bytes16, num uint8) constant returns(address)
func (_HoQuAdCampaign *HoQuAdCampaignCallerSession) GetLeadIntermediaryAddress(id [16]byte, num uint8) (common.Address, error) {
	return _HoQuAdCampaign.Contract.GetLeadIntermediaryAddress(&_HoQuAdCampaign.CallOpts, id, num)
}

// GetLeadIntermediaryPercent is a free data retrieval call binding the contract method 0x5cb8f5ab.
//
// Solidity: function getLeadIntermediaryPercent(id bytes16, num uint8) constant returns(uint32)
func (_HoQuAdCampaign *HoQuAdCampaignCaller) GetLeadIntermediaryPercent(opts *bind.CallOpts, id [16]byte, num uint8) (uint32, error) {
	var (
		ret0 = new(uint32)
	)
	out := ret0
	err := _HoQuAdCampaign.contract.Call(opts, out, "getLeadIntermediaryPercent", id, num)
	return *ret0, err
}

// GetLeadIntermediaryPercent is a free data retrieval call binding the contract method 0x5cb8f5ab.
//
// Solidity: function getLeadIntermediaryPercent(id bytes16, num uint8) constant returns(uint32)
func (_HoQuAdCampaign *HoQuAdCampaignSession) GetLeadIntermediaryPercent(id [16]byte, num uint8) (uint32, error) {
	return _HoQuAdCampaign.Contract.GetLeadIntermediaryPercent(&_HoQuAdCampaign.CallOpts, id, num)
}

// GetLeadIntermediaryPercent is a free data retrieval call binding the contract method 0x5cb8f5ab.
//
// Solidity: function getLeadIntermediaryPercent(id bytes16, num uint8) constant returns(uint32)
func (_HoQuAdCampaign *HoQuAdCampaignCallerSession) GetLeadIntermediaryPercent(id [16]byte, num uint8) (uint32, error) {
	return _HoQuAdCampaign.Contract.GetLeadIntermediaryPercent(&_HoQuAdCampaign.CallOpts, id, num)
}

// Leads is a free data retrieval call binding the contract method 0xe7608fb2.
//
// Solidity: function leads( bytes16) constant returns(createdAt uint256, trackerId bytes16, dataUrl string, meta string, price uint256, numOfIntermediaries uint8, status uint8)
func (_HoQuAdCampaign *HoQuAdCampaignCaller) Leads(opts *bind.CallOpts, arg0 [16]byte) (struct {
	CreatedAt           *big.Int
	TrackerId           [16]byte
	DataUrl             string
	Meta                string
	Price               *big.Int
	NumOfIntermediaries uint8
	Status              uint8
}, error) {
	ret := new(struct {
		CreatedAt           *big.Int
		TrackerId           [16]byte
		DataUrl             string
		Meta                string
		Price               *big.Int
		NumOfIntermediaries uint8
		Status              uint8
	})
	out := ret
	err := _HoQuAdCampaign.contract.Call(opts, out, "leads", arg0)
	return *ret, err
}

// Leads is a free data retrieval call binding the contract method 0xe7608fb2.
//
// Solidity: function leads( bytes16) constant returns(createdAt uint256, trackerId bytes16, dataUrl string, meta string, price uint256, numOfIntermediaries uint8, status uint8)
func (_HoQuAdCampaign *HoQuAdCampaignSession) Leads(arg0 [16]byte) (struct {
	CreatedAt           *big.Int
	TrackerId           [16]byte
	DataUrl             string
	Meta                string
	Price               *big.Int
	NumOfIntermediaries uint8
	Status              uint8
}, error) {
	return _HoQuAdCampaign.Contract.Leads(&_HoQuAdCampaign.CallOpts, arg0)
}

// Leads is a free data retrieval call binding the contract method 0xe7608fb2.
//
// Solidity: function leads( bytes16) constant returns(createdAt uint256, trackerId bytes16, dataUrl string, meta string, price uint256, numOfIntermediaries uint8, status uint8)
func (_HoQuAdCampaign *HoQuAdCampaignCallerSession) Leads(arg0 [16]byte) (struct {
	CreatedAt           *big.Int
	TrackerId           [16]byte
	DataUrl             string
	Meta                string
	Price               *big.Int
	NumOfIntermediaries uint8
	Status              uint8
}, error) {
	return _HoQuAdCampaign.Contract.Leads(&_HoQuAdCampaign.CallOpts, arg0)
}

// OfferId is a free data retrieval call binding the contract method 0x21ce9f91.
//
// Solidity: function offerId() constant returns(bytes16)
func (_HoQuAdCampaign *HoQuAdCampaignCaller) OfferId(opts *bind.CallOpts) ([16]byte, error) {
	var (
		ret0 = new([16]byte)
	)
	out := ret0
	err := _HoQuAdCampaign.contract.Call(opts, out, "offerId")
	return *ret0, err
}

// OfferId is a free data retrieval call binding the contract method 0x21ce9f91.
//
// Solidity: function offerId() constant returns(bytes16)
func (_HoQuAdCampaign *HoQuAdCampaignSession) OfferId() ([16]byte, error) {
	return _HoQuAdCampaign.Contract.OfferId(&_HoQuAdCampaign.CallOpts)
}

// OfferId is a free data retrieval call binding the contract method 0x21ce9f91.
//
// Solidity: function offerId() constant returns(bytes16)
func (_HoQuAdCampaign *HoQuAdCampaignCallerSession) OfferId() ([16]byte, error) {
	return _HoQuAdCampaign.Contract.OfferId(&_HoQuAdCampaign.CallOpts)
}

// PayerAddress is a free data retrieval call binding the contract method 0x2b2f1e14.
//
// Solidity: function payerAddress() constant returns(address)
func (_HoQuAdCampaign *HoQuAdCampaignCaller) PayerAddress(opts *bind.CallOpts) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _HoQuAdCampaign.contract.Call(opts, out, "payerAddress")
	return *ret0, err
}

// PayerAddress is a free data retrieval call binding the contract method 0x2b2f1e14.
//
// Solidity: function payerAddress() constant returns(address)
func (_HoQuAdCampaign *HoQuAdCampaignSession) PayerAddress() (common.Address, error) {
	return _HoQuAdCampaign.Contract.PayerAddress(&_HoQuAdCampaign.CallOpts)
}

// PayerAddress is a free data retrieval call binding the contract method 0x2b2f1e14.
//
// Solidity: function payerAddress() constant returns(address)
func (_HoQuAdCampaign *HoQuAdCampaignCallerSession) PayerAddress() (common.Address, error) {
	return _HoQuAdCampaign.Contract.PayerAddress(&_HoQuAdCampaign.CallOpts)
}

// Rater is a free data retrieval call binding the contract method 0x355c9fb1.
//
// Solidity: function rater() constant returns(address)
func (_HoQuAdCampaign *HoQuAdCampaignCaller) Rater(opts *bind.CallOpts) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _HoQuAdCampaign.contract.Call(opts, out, "rater")
	return *ret0, err
}

// Rater is a free data retrieval call binding the contract method 0x355c9fb1.
//
// Solidity: function rater() constant returns(address)
func (_HoQuAdCampaign *HoQuAdCampaignSession) Rater() (common.Address, error) {
	return _HoQuAdCampaign.Contract.Rater(&_HoQuAdCampaign.CallOpts)
}

// Rater is a free data retrieval call binding the contract method 0x355c9fb1.
//
// Solidity: function rater() constant returns(address)
func (_HoQuAdCampaign *HoQuAdCampaignCallerSession) Rater() (common.Address, error) {
	return _HoQuAdCampaign.Contract.Rater(&_HoQuAdCampaign.CallOpts)
}

// Status is a free data retrieval call binding the contract method 0x200d2ed2.
//
// Solidity: function status() constant returns(uint8)
func (_HoQuAdCampaign *HoQuAdCampaignCaller) Status(opts *bind.CallOpts) (uint8, error) {
	var (
		ret0 = new(uint8)
	)
	out := ret0
	err := _HoQuAdCampaign.contract.Call(opts, out, "status")
	return *ret0, err
}

// Status is a free data retrieval call binding the contract method 0x200d2ed2.
//
// Solidity: function status() constant returns(uint8)
func (_HoQuAdCampaign *HoQuAdCampaignSession) Status() (uint8, error) {
	return _HoQuAdCampaign.Contract.Status(&_HoQuAdCampaign.CallOpts)
}

// Status is a free data retrieval call binding the contract method 0x200d2ed2.
//
// Solidity: function status() constant returns(uint8)
func (_HoQuAdCampaign *HoQuAdCampaignCallerSession) Status() (uint8, error) {
	return _HoQuAdCampaign.Contract.Status(&_HoQuAdCampaign.CallOpts)
}

// Store is a free data retrieval call binding the contract method 0x975057e7.
//
// Solidity: function store() constant returns(address)
func (_HoQuAdCampaign *HoQuAdCampaignCaller) Store(opts *bind.CallOpts) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _HoQuAdCampaign.contract.Call(opts, out, "store")
	return *ret0, err
}

// Store is a free data retrieval call binding the contract method 0x975057e7.
//
// Solidity: function store() constant returns(address)
func (_HoQuAdCampaign *HoQuAdCampaignSession) Store() (common.Address, error) {
	return _HoQuAdCampaign.Contract.Store(&_HoQuAdCampaign.CallOpts)
}

// Store is a free data retrieval call binding the contract method 0x975057e7.
//
// Solidity: function store() constant returns(address)
func (_HoQuAdCampaign *HoQuAdCampaignCallerSession) Store() (common.Address, error) {
	return _HoQuAdCampaign.Contract.Store(&_HoQuAdCampaign.CallOpts)
}

// Token is a free data retrieval call binding the contract method 0xfc0c546a.
//
// Solidity: function token() constant returns(address)
func (_HoQuAdCampaign *HoQuAdCampaignCaller) Token(opts *bind.CallOpts) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _HoQuAdCampaign.contract.Call(opts, out, "token")
	return *ret0, err
}

// Token is a free data retrieval call binding the contract method 0xfc0c546a.
//
// Solidity: function token() constant returns(address)
func (_HoQuAdCampaign *HoQuAdCampaignSession) Token() (common.Address, error) {
	return _HoQuAdCampaign.Contract.Token(&_HoQuAdCampaign.CallOpts)
}

// Token is a free data retrieval call binding the contract method 0xfc0c546a.
//
// Solidity: function token() constant returns(address)
func (_HoQuAdCampaign *HoQuAdCampaignCallerSession) Token() (common.Address, error) {
	return _HoQuAdCampaign.Contract.Token(&_HoQuAdCampaign.CallOpts)
}

// Trackers is a free data retrieval call binding the contract method 0x6cb234a3.
//
// Solidity: function trackers( address) constant returns(bytes16)
func (_HoQuAdCampaign *HoQuAdCampaignCaller) Trackers(opts *bind.CallOpts, arg0 common.Address) ([16]byte, error) {
	var (
		ret0 = new([16]byte)
	)
	out := ret0
	err := _HoQuAdCampaign.contract.Call(opts, out, "trackers", arg0)
	return *ret0, err
}

// Trackers is a free data retrieval call binding the contract method 0x6cb234a3.
//
// Solidity: function trackers( address) constant returns(bytes16)
func (_HoQuAdCampaign *HoQuAdCampaignSession) Trackers(arg0 common.Address) ([16]byte, error) {
	return _HoQuAdCampaign.Contract.Trackers(&_HoQuAdCampaign.CallOpts, arg0)
}

// Trackers is a free data retrieval call binding the contract method 0x6cb234a3.
//
// Solidity: function trackers( address) constant returns(bytes16)
func (_HoQuAdCampaign *HoQuAdCampaignCallerSession) Trackers(arg0 common.Address) ([16]byte, error) {
	return _HoQuAdCampaign.Contract.Trackers(&_HoQuAdCampaign.CallOpts, arg0)
}

// AddLead is a paid mutator transaction binding the contract method 0x7a8187cf.
//
// Solidity: function addLead(id bytes16, trackerId bytes16, meta string, dataUrl string, price uint256) returns()
func (_HoQuAdCampaign *HoQuAdCampaignTransactor) AddLead(opts *bind.TransactOpts, id [16]byte, trackerId [16]byte, meta string, dataUrl string, price *big.Int) (*types.Transaction, error) {
	return _HoQuAdCampaign.contract.Transact(opts, "addLead", id, trackerId, meta, dataUrl, price)
}

// AddLead is a paid mutator transaction binding the contract method 0x7a8187cf.
//
// Solidity: function addLead(id bytes16, trackerId bytes16, meta string, dataUrl string, price uint256) returns()
func (_HoQuAdCampaign *HoQuAdCampaignSession) AddLead(id [16]byte, trackerId [16]byte, meta string, dataUrl string, price *big.Int) (*types.Transaction, error) {
	return _HoQuAdCampaign.Contract.AddLead(&_HoQuAdCampaign.TransactOpts, id, trackerId, meta, dataUrl, price)
}

// AddLead is a paid mutator transaction binding the contract method 0x7a8187cf.
//
// Solidity: function addLead(id bytes16, trackerId bytes16, meta string, dataUrl string, price uint256) returns()
func (_HoQuAdCampaign *HoQuAdCampaignTransactorSession) AddLead(id [16]byte, trackerId [16]byte, meta string, dataUrl string, price *big.Int) (*types.Transaction, error) {
	return _HoQuAdCampaign.Contract.AddLead(&_HoQuAdCampaign.TransactOpts, id, trackerId, meta, dataUrl, price)
}

// AddLeadIntermediary is a paid mutator transaction binding the contract method 0x36617045.
//
// Solidity: function addLeadIntermediary(id bytes16, intermediaryAddress address, percent uint32) returns()
func (_HoQuAdCampaign *HoQuAdCampaignTransactor) AddLeadIntermediary(opts *bind.TransactOpts, id [16]byte, intermediaryAddress common.Address, percent uint32) (*types.Transaction, error) {
	return _HoQuAdCampaign.contract.Transact(opts, "addLeadIntermediary", id, intermediaryAddress, percent)
}

// AddLeadIntermediary is a paid mutator transaction binding the contract method 0x36617045.
//
// Solidity: function addLeadIntermediary(id bytes16, intermediaryAddress address, percent uint32) returns()
func (_HoQuAdCampaign *HoQuAdCampaignSession) AddLeadIntermediary(id [16]byte, intermediaryAddress common.Address, percent uint32) (*types.Transaction, error) {
	return _HoQuAdCampaign.Contract.AddLeadIntermediary(&_HoQuAdCampaign.TransactOpts, id, intermediaryAddress, percent)
}

// AddLeadIntermediary is a paid mutator transaction binding the contract method 0x36617045.
//
// Solidity: function addLeadIntermediary(id bytes16, intermediaryAddress address, percent uint32) returns()
func (_HoQuAdCampaign *HoQuAdCampaignTransactorSession) AddLeadIntermediary(id [16]byte, intermediaryAddress common.Address, percent uint32) (*types.Transaction, error) {
	return _HoQuAdCampaign.Contract.AddLeadIntermediary(&_HoQuAdCampaign.TransactOpts, id, intermediaryAddress, percent)
}

// AddTracker is a paid mutator transaction binding the contract method 0x6aa0082c.
//
// Solidity: function addTracker(ownerAddress address, id bytes16) returns()
func (_HoQuAdCampaign *HoQuAdCampaignTransactor) AddTracker(opts *bind.TransactOpts, ownerAddress common.Address, id [16]byte) (*types.Transaction, error) {
	return _HoQuAdCampaign.contract.Transact(opts, "addTracker", ownerAddress, id)
}

// AddTracker is a paid mutator transaction binding the contract method 0x6aa0082c.
//
// Solidity: function addTracker(ownerAddress address, id bytes16) returns()
func (_HoQuAdCampaign *HoQuAdCampaignSession) AddTracker(ownerAddress common.Address, id [16]byte) (*types.Transaction, error) {
	return _HoQuAdCampaign.Contract.AddTracker(&_HoQuAdCampaign.TransactOpts, ownerAddress, id)
}

// AddTracker is a paid mutator transaction binding the contract method 0x6aa0082c.
//
// Solidity: function addTracker(ownerAddress address, id bytes16) returns()
func (_HoQuAdCampaign *HoQuAdCampaignTransactorSession) AddTracker(ownerAddress common.Address, id [16]byte) (*types.Transaction, error) {
	return _HoQuAdCampaign.Contract.AddTracker(&_HoQuAdCampaign.TransactOpts, ownerAddress, id)
}

// SetBeneficiaryAddress is a paid mutator transaction binding the contract method 0xec6be06e.
//
// Solidity: function setBeneficiaryAddress(_beneficiaryAddress address) returns()
func (_HoQuAdCampaign *HoQuAdCampaignTransactor) SetBeneficiaryAddress(opts *bind.TransactOpts, _beneficiaryAddress common.Address) (*types.Transaction, error) {
	return _HoQuAdCampaign.contract.Transact(opts, "setBeneficiaryAddress", _beneficiaryAddress)
}

// SetBeneficiaryAddress is a paid mutator transaction binding the contract method 0xec6be06e.
//
// Solidity: function setBeneficiaryAddress(_beneficiaryAddress address) returns()
func (_HoQuAdCampaign *HoQuAdCampaignSession) SetBeneficiaryAddress(_beneficiaryAddress common.Address) (*types.Transaction, error) {
	return _HoQuAdCampaign.Contract.SetBeneficiaryAddress(&_HoQuAdCampaign.TransactOpts, _beneficiaryAddress)
}

// SetBeneficiaryAddress is a paid mutator transaction binding the contract method 0xec6be06e.
//
// Solidity: function setBeneficiaryAddress(_beneficiaryAddress address) returns()
func (_HoQuAdCampaign *HoQuAdCampaignTransactorSession) SetBeneficiaryAddress(_beneficiaryAddress common.Address) (*types.Transaction, error) {
	return _HoQuAdCampaign.Contract.SetBeneficiaryAddress(&_HoQuAdCampaign.TransactOpts, _beneficiaryAddress)
}

// SetConfigAddress is a paid mutator transaction binding the contract method 0x83a12de9.
//
// Solidity: function setConfigAddress(configAddress address) returns()
func (_HoQuAdCampaign *HoQuAdCampaignTransactor) SetConfigAddress(opts *bind.TransactOpts, configAddress common.Address) (*types.Transaction, error) {
	return _HoQuAdCampaign.contract.Transact(opts, "setConfigAddress", configAddress)
}

// SetConfigAddress is a paid mutator transaction binding the contract method 0x83a12de9.
//
// Solidity: function setConfigAddress(configAddress address) returns()
func (_HoQuAdCampaign *HoQuAdCampaignSession) SetConfigAddress(configAddress common.Address) (*types.Transaction, error) {
	return _HoQuAdCampaign.Contract.SetConfigAddress(&_HoQuAdCampaign.TransactOpts, configAddress)
}

// SetConfigAddress is a paid mutator transaction binding the contract method 0x83a12de9.
//
// Solidity: function setConfigAddress(configAddress address) returns()
func (_HoQuAdCampaign *HoQuAdCampaignTransactorSession) SetConfigAddress(configAddress common.Address) (*types.Transaction, error) {
	return _HoQuAdCampaign.Contract.SetConfigAddress(&_HoQuAdCampaign.TransactOpts, configAddress)
}

// SetLeadStatus is a paid mutator transaction binding the contract method 0x81e85aa3.
//
// Solidity: function setLeadStatus(id bytes16, _status uint8) returns()
func (_HoQuAdCampaign *HoQuAdCampaignTransactor) SetLeadStatus(opts *bind.TransactOpts, id [16]byte, _status uint8) (*types.Transaction, error) {
	return _HoQuAdCampaign.contract.Transact(opts, "setLeadStatus", id, _status)
}

// SetLeadStatus is a paid mutator transaction binding the contract method 0x81e85aa3.
//
// Solidity: function setLeadStatus(id bytes16, _status uint8) returns()
func (_HoQuAdCampaign *HoQuAdCampaignSession) SetLeadStatus(id [16]byte, _status uint8) (*types.Transaction, error) {
	return _HoQuAdCampaign.Contract.SetLeadStatus(&_HoQuAdCampaign.TransactOpts, id, _status)
}

// SetLeadStatus is a paid mutator transaction binding the contract method 0x81e85aa3.
//
// Solidity: function setLeadStatus(id bytes16, _status uint8) returns()
func (_HoQuAdCampaign *HoQuAdCampaignTransactorSession) SetLeadStatus(id [16]byte, _status uint8) (*types.Transaction, error) {
	return _HoQuAdCampaign.Contract.SetLeadStatus(&_HoQuAdCampaign.TransactOpts, id, _status)
}

// SetPayerAddress is a paid mutator transaction binding the contract method 0x4b28676d.
//
// Solidity: function setPayerAddress(_payerAddress address) returns()
func (_HoQuAdCampaign *HoQuAdCampaignTransactor) SetPayerAddress(opts *bind.TransactOpts, _payerAddress common.Address) (*types.Transaction, error) {
	return _HoQuAdCampaign.contract.Transact(opts, "setPayerAddress", _payerAddress)
}

// SetPayerAddress is a paid mutator transaction binding the contract method 0x4b28676d.
//
// Solidity: function setPayerAddress(_payerAddress address) returns()
func (_HoQuAdCampaign *HoQuAdCampaignSession) SetPayerAddress(_payerAddress common.Address) (*types.Transaction, error) {
	return _HoQuAdCampaign.Contract.SetPayerAddress(&_HoQuAdCampaign.TransactOpts, _payerAddress)
}

// SetPayerAddress is a paid mutator transaction binding the contract method 0x4b28676d.
//
// Solidity: function setPayerAddress(_payerAddress address) returns()
func (_HoQuAdCampaign *HoQuAdCampaignTransactorSession) SetPayerAddress(_payerAddress common.Address) (*types.Transaction, error) {
	return _HoQuAdCampaign.Contract.SetPayerAddress(&_HoQuAdCampaign.TransactOpts, _payerAddress)
}

// SetRaterAddress is a paid mutator transaction binding the contract method 0xce88f04e.
//
// Solidity: function setRaterAddress(raterAddress address) returns()
func (_HoQuAdCampaign *HoQuAdCampaignTransactor) SetRaterAddress(opts *bind.TransactOpts, raterAddress common.Address) (*types.Transaction, error) {
	return _HoQuAdCampaign.contract.Transact(opts, "setRaterAddress", raterAddress)
}

// SetRaterAddress is a paid mutator transaction binding the contract method 0xce88f04e.
//
// Solidity: function setRaterAddress(raterAddress address) returns()
func (_HoQuAdCampaign *HoQuAdCampaignSession) SetRaterAddress(raterAddress common.Address) (*types.Transaction, error) {
	return _HoQuAdCampaign.Contract.SetRaterAddress(&_HoQuAdCampaign.TransactOpts, raterAddress)
}

// SetRaterAddress is a paid mutator transaction binding the contract method 0xce88f04e.
//
// Solidity: function setRaterAddress(raterAddress address) returns()
func (_HoQuAdCampaign *HoQuAdCampaignTransactorSession) SetRaterAddress(raterAddress common.Address) (*types.Transaction, error) {
	return _HoQuAdCampaign.Contract.SetRaterAddress(&_HoQuAdCampaign.TransactOpts, raterAddress)
}

// SetStatus is a paid mutator transaction binding the contract method 0x2e49d78b.
//
// Solidity: function setStatus(_status uint8) returns()
func (_HoQuAdCampaign *HoQuAdCampaignTransactor) SetStatus(opts *bind.TransactOpts, _status uint8) (*types.Transaction, error) {
	return _HoQuAdCampaign.contract.Transact(opts, "setStatus", _status)
}

// SetStatus is a paid mutator transaction binding the contract method 0x2e49d78b.
//
// Solidity: function setStatus(_status uint8) returns()
func (_HoQuAdCampaign *HoQuAdCampaignSession) SetStatus(_status uint8) (*types.Transaction, error) {
	return _HoQuAdCampaign.Contract.SetStatus(&_HoQuAdCampaign.TransactOpts, _status)
}

// SetStatus is a paid mutator transaction binding the contract method 0x2e49d78b.
//
// Solidity: function setStatus(_status uint8) returns()
func (_HoQuAdCampaign *HoQuAdCampaignTransactorSession) SetStatus(_status uint8) (*types.Transaction, error) {
	return _HoQuAdCampaign.Contract.SetStatus(&_HoQuAdCampaign.TransactOpts, _status)
}

// SetStorageAddress is a paid mutator transaction binding the contract method 0x59b910d6.
//
// Solidity: function setStorageAddress(storageAddress address) returns()
func (_HoQuAdCampaign *HoQuAdCampaignTransactor) SetStorageAddress(opts *bind.TransactOpts, storageAddress common.Address) (*types.Transaction, error) {
	return _HoQuAdCampaign.contract.Transact(opts, "setStorageAddress", storageAddress)
}

// SetStorageAddress is a paid mutator transaction binding the contract method 0x59b910d6.
//
// Solidity: function setStorageAddress(storageAddress address) returns()
func (_HoQuAdCampaign *HoQuAdCampaignSession) SetStorageAddress(storageAddress common.Address) (*types.Transaction, error) {
	return _HoQuAdCampaign.Contract.SetStorageAddress(&_HoQuAdCampaign.TransactOpts, storageAddress)
}

// SetStorageAddress is a paid mutator transaction binding the contract method 0x59b910d6.
//
// Solidity: function setStorageAddress(storageAddress address) returns()
func (_HoQuAdCampaign *HoQuAdCampaignTransactorSession) SetStorageAddress(storageAddress common.Address) (*types.Transaction, error) {
	return _HoQuAdCampaign.Contract.SetStorageAddress(&_HoQuAdCampaign.TransactOpts, storageAddress)
}

// SetTokenAddress is a paid mutator transaction binding the contract method 0x26a4e8d2.
//
// Solidity: function setTokenAddress(tokenAddress address) returns()
func (_HoQuAdCampaign *HoQuAdCampaignTransactor) SetTokenAddress(opts *bind.TransactOpts, tokenAddress common.Address) (*types.Transaction, error) {
	return _HoQuAdCampaign.contract.Transact(opts, "setTokenAddress", tokenAddress)
}

// SetTokenAddress is a paid mutator transaction binding the contract method 0x26a4e8d2.
//
// Solidity: function setTokenAddress(tokenAddress address) returns()
func (_HoQuAdCampaign *HoQuAdCampaignSession) SetTokenAddress(tokenAddress common.Address) (*types.Transaction, error) {
	return _HoQuAdCampaign.Contract.SetTokenAddress(&_HoQuAdCampaign.TransactOpts, tokenAddress)
}

// SetTokenAddress is a paid mutator transaction binding the contract method 0x26a4e8d2.
//
// Solidity: function setTokenAddress(tokenAddress address) returns()
func (_HoQuAdCampaign *HoQuAdCampaignTransactorSession) SetTokenAddress(tokenAddress common.Address) (*types.Transaction, error) {
	return _HoQuAdCampaign.Contract.SetTokenAddress(&_HoQuAdCampaign.TransactOpts, tokenAddress)
}

// TransactLead is a paid mutator transaction binding the contract method 0x835aef40.
//
// Solidity: function transactLead(id bytes16) returns()
func (_HoQuAdCampaign *HoQuAdCampaignTransactor) TransactLead(opts *bind.TransactOpts, id [16]byte) (*types.Transaction, error) {
	return _HoQuAdCampaign.contract.Transact(opts, "transactLead", id)
}

// TransactLead is a paid mutator transaction binding the contract method 0x835aef40.
//
// Solidity: function transactLead(id bytes16) returns()
func (_HoQuAdCampaign *HoQuAdCampaignSession) TransactLead(id [16]byte) (*types.Transaction, error) {
	return _HoQuAdCampaign.Contract.TransactLead(&_HoQuAdCampaign.TransactOpts, id)
}

// TransactLead is a paid mutator transaction binding the contract method 0x835aef40.
//
// Solidity: function transactLead(id bytes16) returns()
func (_HoQuAdCampaign *HoQuAdCampaignTransactorSession) TransactLead(id [16]byte) (*types.Transaction, error) {
	return _HoQuAdCampaign.Contract.TransactLead(&_HoQuAdCampaign.TransactOpts, id)
}

// HoQuAdCampaignBeneficiaryAddressChangedIterator is returned from FilterBeneficiaryAddressChanged and is used to iterate over the raw logs and unpacked data for BeneficiaryAddressChanged events raised by the HoQuAdCampaign contract.
type HoQuAdCampaignBeneficiaryAddressChangedIterator struct {
	Event *HoQuAdCampaignBeneficiaryAddressChanged // Event containing the contract specifics and raw log

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
func (it *HoQuAdCampaignBeneficiaryAddressChangedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(HoQuAdCampaignBeneficiaryAddressChanged)
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
		it.Event = new(HoQuAdCampaignBeneficiaryAddressChanged)
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
func (it *HoQuAdCampaignBeneficiaryAddressChangedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *HoQuAdCampaignBeneficiaryAddressChangedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// HoQuAdCampaignBeneficiaryAddressChanged represents a BeneficiaryAddressChanged event raised by the HoQuAdCampaign contract.
type HoQuAdCampaignBeneficiaryAddressChanged struct {
	BeneficiaryAddress    common.Address
	NewBeneficiaryAddress common.Address
	Raw                   types.Log // Blockchain specific contextual infos
}

// FilterBeneficiaryAddressChanged is a free log retrieval operation binding the contract event 0xb53b7057ddbff0cd9111b0a1a501a9cc1a9c315eb3396cb0d47917908325999a.
//
// Solidity: event BeneficiaryAddressChanged(beneficiaryAddress indexed address, newBeneficiaryAddress indexed address)
func (_HoQuAdCampaign *HoQuAdCampaignFilterer) FilterBeneficiaryAddressChanged(opts *bind.FilterOpts, beneficiaryAddress []common.Address, newBeneficiaryAddress []common.Address) (*HoQuAdCampaignBeneficiaryAddressChangedIterator, error) {

	var beneficiaryAddressRule []interface{}
	for _, beneficiaryAddressItem := range beneficiaryAddress {
		beneficiaryAddressRule = append(beneficiaryAddressRule, beneficiaryAddressItem)
	}
	var newBeneficiaryAddressRule []interface{}
	for _, newBeneficiaryAddressItem := range newBeneficiaryAddress {
		newBeneficiaryAddressRule = append(newBeneficiaryAddressRule, newBeneficiaryAddressItem)
	}

	logs, sub, err := _HoQuAdCampaign.contract.FilterLogs(opts, "BeneficiaryAddressChanged", beneficiaryAddressRule, newBeneficiaryAddressRule)
	if err != nil {
		return nil, err
	}
	return &HoQuAdCampaignBeneficiaryAddressChangedIterator{contract: _HoQuAdCampaign.contract, event: "BeneficiaryAddressChanged", logs: logs, sub: sub}, nil
}

// WatchBeneficiaryAddressChanged is a free log subscription operation binding the contract event 0xb53b7057ddbff0cd9111b0a1a501a9cc1a9c315eb3396cb0d47917908325999a.
//
// Solidity: event BeneficiaryAddressChanged(beneficiaryAddress indexed address, newBeneficiaryAddress indexed address)
func (_HoQuAdCampaign *HoQuAdCampaignFilterer) WatchBeneficiaryAddressChanged(opts *bind.WatchOpts, sink chan<- *HoQuAdCampaignBeneficiaryAddressChanged, beneficiaryAddress []common.Address, newBeneficiaryAddress []common.Address) (event.Subscription, error) {

	var beneficiaryAddressRule []interface{}
	for _, beneficiaryAddressItem := range beneficiaryAddress {
		beneficiaryAddressRule = append(beneficiaryAddressRule, beneficiaryAddressItem)
	}
	var newBeneficiaryAddressRule []interface{}
	for _, newBeneficiaryAddressItem := range newBeneficiaryAddress {
		newBeneficiaryAddressRule = append(newBeneficiaryAddressRule, newBeneficiaryAddressItem)
	}

	logs, sub, err := _HoQuAdCampaign.contract.WatchLogs(opts, "BeneficiaryAddressChanged", beneficiaryAddressRule, newBeneficiaryAddressRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(HoQuAdCampaignBeneficiaryAddressChanged)
				if err := _HoQuAdCampaign.contract.UnpackLog(event, "BeneficiaryAddressChanged", log); err != nil {
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

// HoQuAdCampaignLeadAddedIterator is returned from FilterLeadAdded and is used to iterate over the raw logs and unpacked data for LeadAdded events raised by the HoQuAdCampaign contract.
type HoQuAdCampaignLeadAddedIterator struct {
	Event *HoQuAdCampaignLeadAdded // Event containing the contract specifics and raw log

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
func (it *HoQuAdCampaignLeadAddedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(HoQuAdCampaignLeadAdded)
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
		it.Event = new(HoQuAdCampaignLeadAdded)
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
func (it *HoQuAdCampaignLeadAddedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *HoQuAdCampaignLeadAddedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// HoQuAdCampaignLeadAdded represents a LeadAdded event raised by the HoQuAdCampaign contract.
type HoQuAdCampaignLeadAdded struct {
	BeneficiaryAddress common.Address
	Id                 [16]byte
	Price              *big.Int
	Raw                types.Log // Blockchain specific contextual infos
}

// FilterLeadAdded is a free log retrieval operation binding the contract event 0x5a4f5a43985de2c0807c6875a069d25bfdf90255dc58958a49e07d00936537e3.
//
// Solidity: event LeadAdded(beneficiaryAddress indexed address, id bytes16, price uint256)
func (_HoQuAdCampaign *HoQuAdCampaignFilterer) FilterLeadAdded(opts *bind.FilterOpts, beneficiaryAddress []common.Address) (*HoQuAdCampaignLeadAddedIterator, error) {

	var beneficiaryAddressRule []interface{}
	for _, beneficiaryAddressItem := range beneficiaryAddress {
		beneficiaryAddressRule = append(beneficiaryAddressRule, beneficiaryAddressItem)
	}

	logs, sub, err := _HoQuAdCampaign.contract.FilterLogs(opts, "LeadAdded", beneficiaryAddressRule)
	if err != nil {
		return nil, err
	}
	return &HoQuAdCampaignLeadAddedIterator{contract: _HoQuAdCampaign.contract, event: "LeadAdded", logs: logs, sub: sub}, nil
}

// WatchLeadAdded is a free log subscription operation binding the contract event 0x5a4f5a43985de2c0807c6875a069d25bfdf90255dc58958a49e07d00936537e3.
//
// Solidity: event LeadAdded(beneficiaryAddress indexed address, id bytes16, price uint256)
func (_HoQuAdCampaign *HoQuAdCampaignFilterer) WatchLeadAdded(opts *bind.WatchOpts, sink chan<- *HoQuAdCampaignLeadAdded, beneficiaryAddress []common.Address) (event.Subscription, error) {

	var beneficiaryAddressRule []interface{}
	for _, beneficiaryAddressItem := range beneficiaryAddress {
		beneficiaryAddressRule = append(beneficiaryAddressRule, beneficiaryAddressItem)
	}

	logs, sub, err := _HoQuAdCampaign.contract.WatchLogs(opts, "LeadAdded", beneficiaryAddressRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(HoQuAdCampaignLeadAdded)
				if err := _HoQuAdCampaign.contract.UnpackLog(event, "LeadAdded", log); err != nil {
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

// HoQuAdCampaignLeadStatusChangedIterator is returned from FilterLeadStatusChanged and is used to iterate over the raw logs and unpacked data for LeadStatusChanged events raised by the HoQuAdCampaign contract.
type HoQuAdCampaignLeadStatusChangedIterator struct {
	Event *HoQuAdCampaignLeadStatusChanged // Event containing the contract specifics and raw log

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
func (it *HoQuAdCampaignLeadStatusChangedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(HoQuAdCampaignLeadStatusChanged)
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
		it.Event = new(HoQuAdCampaignLeadStatusChanged)
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
func (it *HoQuAdCampaignLeadStatusChangedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *HoQuAdCampaignLeadStatusChangedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// HoQuAdCampaignLeadStatusChanged represents a LeadStatusChanged event raised by the HoQuAdCampaign contract.
type HoQuAdCampaignLeadStatusChanged struct {
	BeneficiaryAddress common.Address
	Id                 [16]byte
	Status             uint8
	Raw                types.Log // Blockchain specific contextual infos
}

// FilterLeadStatusChanged is a free log retrieval operation binding the contract event 0xfa2a86784442baf24e238559b3ca5986aa4a588eeb7840c5ba8b27d6e663ecd9.
//
// Solidity: event LeadStatusChanged(beneficiaryAddress indexed address, id bytes16, status uint8)
func (_HoQuAdCampaign *HoQuAdCampaignFilterer) FilterLeadStatusChanged(opts *bind.FilterOpts, beneficiaryAddress []common.Address) (*HoQuAdCampaignLeadStatusChangedIterator, error) {

	var beneficiaryAddressRule []interface{}
	for _, beneficiaryAddressItem := range beneficiaryAddress {
		beneficiaryAddressRule = append(beneficiaryAddressRule, beneficiaryAddressItem)
	}

	logs, sub, err := _HoQuAdCampaign.contract.FilterLogs(opts, "LeadStatusChanged", beneficiaryAddressRule)
	if err != nil {
		return nil, err
	}
	return &HoQuAdCampaignLeadStatusChangedIterator{contract: _HoQuAdCampaign.contract, event: "LeadStatusChanged", logs: logs, sub: sub}, nil
}

// WatchLeadStatusChanged is a free log subscription operation binding the contract event 0xfa2a86784442baf24e238559b3ca5986aa4a588eeb7840c5ba8b27d6e663ecd9.
//
// Solidity: event LeadStatusChanged(beneficiaryAddress indexed address, id bytes16, status uint8)
func (_HoQuAdCampaign *HoQuAdCampaignFilterer) WatchLeadStatusChanged(opts *bind.WatchOpts, sink chan<- *HoQuAdCampaignLeadStatusChanged, beneficiaryAddress []common.Address) (event.Subscription, error) {

	var beneficiaryAddressRule []interface{}
	for _, beneficiaryAddressItem := range beneficiaryAddress {
		beneficiaryAddressRule = append(beneficiaryAddressRule, beneficiaryAddressItem)
	}

	logs, sub, err := _HoQuAdCampaign.contract.WatchLogs(opts, "LeadStatusChanged", beneficiaryAddressRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(HoQuAdCampaignLeadStatusChanged)
				if err := _HoQuAdCampaign.contract.UnpackLog(event, "LeadStatusChanged", log); err != nil {
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

// HoQuAdCampaignLeadTransactedIterator is returned from FilterLeadTransacted and is used to iterate over the raw logs and unpacked data for LeadTransacted events raised by the HoQuAdCampaign contract.
type HoQuAdCampaignLeadTransactedIterator struct {
	Event *HoQuAdCampaignLeadTransacted // Event containing the contract specifics and raw log

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
func (it *HoQuAdCampaignLeadTransactedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(HoQuAdCampaignLeadTransacted)
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
		it.Event = new(HoQuAdCampaignLeadTransacted)
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
func (it *HoQuAdCampaignLeadTransactedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *HoQuAdCampaignLeadTransactedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// HoQuAdCampaignLeadTransacted represents a LeadTransacted event raised by the HoQuAdCampaign contract.
type HoQuAdCampaignLeadTransacted struct {
	BeneficiaryAddress common.Address
	Id                 [16]byte
	Amount             *big.Int
	Raw                types.Log // Blockchain specific contextual infos
}

// FilterLeadTransacted is a free log retrieval operation binding the contract event 0x8e1b43b752aaaf23ee6b52413a1dd24c21f4571e7ff32d88456e52befde6dd04.
//
// Solidity: event LeadTransacted(beneficiaryAddress indexed address, id bytes16, amount uint256)
func (_HoQuAdCampaign *HoQuAdCampaignFilterer) FilterLeadTransacted(opts *bind.FilterOpts, beneficiaryAddress []common.Address) (*HoQuAdCampaignLeadTransactedIterator, error) {

	var beneficiaryAddressRule []interface{}
	for _, beneficiaryAddressItem := range beneficiaryAddress {
		beneficiaryAddressRule = append(beneficiaryAddressRule, beneficiaryAddressItem)
	}

	logs, sub, err := _HoQuAdCampaign.contract.FilterLogs(opts, "LeadTransacted", beneficiaryAddressRule)
	if err != nil {
		return nil, err
	}
	return &HoQuAdCampaignLeadTransactedIterator{contract: _HoQuAdCampaign.contract, event: "LeadTransacted", logs: logs, sub: sub}, nil
}

// WatchLeadTransacted is a free log subscription operation binding the contract event 0x8e1b43b752aaaf23ee6b52413a1dd24c21f4571e7ff32d88456e52befde6dd04.
//
// Solidity: event LeadTransacted(beneficiaryAddress indexed address, id bytes16, amount uint256)
func (_HoQuAdCampaign *HoQuAdCampaignFilterer) WatchLeadTransacted(opts *bind.WatchOpts, sink chan<- *HoQuAdCampaignLeadTransacted, beneficiaryAddress []common.Address) (event.Subscription, error) {

	var beneficiaryAddressRule []interface{}
	for _, beneficiaryAddressItem := range beneficiaryAddress {
		beneficiaryAddressRule = append(beneficiaryAddressRule, beneficiaryAddressItem)
	}

	logs, sub, err := _HoQuAdCampaign.contract.WatchLogs(opts, "LeadTransacted", beneficiaryAddressRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(HoQuAdCampaignLeadTransacted)
				if err := _HoQuAdCampaign.contract.UnpackLog(event, "LeadTransacted", log); err != nil {
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

// HoQuAdCampaignPayerAddressChangedIterator is returned from FilterPayerAddressChanged and is used to iterate over the raw logs and unpacked data for PayerAddressChanged events raised by the HoQuAdCampaign contract.
type HoQuAdCampaignPayerAddressChangedIterator struct {
	Event *HoQuAdCampaignPayerAddressChanged // Event containing the contract specifics and raw log

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
func (it *HoQuAdCampaignPayerAddressChangedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(HoQuAdCampaignPayerAddressChanged)
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
		it.Event = new(HoQuAdCampaignPayerAddressChanged)
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
func (it *HoQuAdCampaignPayerAddressChangedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *HoQuAdCampaignPayerAddressChangedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// HoQuAdCampaignPayerAddressChanged represents a PayerAddressChanged event raised by the HoQuAdCampaign contract.
type HoQuAdCampaignPayerAddressChanged struct {
	PayerAddress    common.Address
	NewPayerAddress common.Address
	Raw             types.Log // Blockchain specific contextual infos
}

// FilterPayerAddressChanged is a free log retrieval operation binding the contract event 0x650742bb22425d77490f0e51db9824c857c049c9fc9fee93b369553acb172758.
//
// Solidity: event PayerAddressChanged(payerAddress indexed address, newPayerAddress indexed address)
func (_HoQuAdCampaign *HoQuAdCampaignFilterer) FilterPayerAddressChanged(opts *bind.FilterOpts, payerAddress []common.Address, newPayerAddress []common.Address) (*HoQuAdCampaignPayerAddressChangedIterator, error) {

	var payerAddressRule []interface{}
	for _, payerAddressItem := range payerAddress {
		payerAddressRule = append(payerAddressRule, payerAddressItem)
	}
	var newPayerAddressRule []interface{}
	for _, newPayerAddressItem := range newPayerAddress {
		newPayerAddressRule = append(newPayerAddressRule, newPayerAddressItem)
	}

	logs, sub, err := _HoQuAdCampaign.contract.FilterLogs(opts, "PayerAddressChanged", payerAddressRule, newPayerAddressRule)
	if err != nil {
		return nil, err
	}
	return &HoQuAdCampaignPayerAddressChangedIterator{contract: _HoQuAdCampaign.contract, event: "PayerAddressChanged", logs: logs, sub: sub}, nil
}

// WatchPayerAddressChanged is a free log subscription operation binding the contract event 0x650742bb22425d77490f0e51db9824c857c049c9fc9fee93b369553acb172758.
//
// Solidity: event PayerAddressChanged(payerAddress indexed address, newPayerAddress indexed address)
func (_HoQuAdCampaign *HoQuAdCampaignFilterer) WatchPayerAddressChanged(opts *bind.WatchOpts, sink chan<- *HoQuAdCampaignPayerAddressChanged, payerAddress []common.Address, newPayerAddress []common.Address) (event.Subscription, error) {

	var payerAddressRule []interface{}
	for _, payerAddressItem := range payerAddress {
		payerAddressRule = append(payerAddressRule, payerAddressItem)
	}
	var newPayerAddressRule []interface{}
	for _, newPayerAddressItem := range newPayerAddress {
		newPayerAddressRule = append(newPayerAddressRule, newPayerAddressItem)
	}

	logs, sub, err := _HoQuAdCampaign.contract.WatchLogs(opts, "PayerAddressChanged", payerAddressRule, newPayerAddressRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(HoQuAdCampaignPayerAddressChanged)
				if err := _HoQuAdCampaign.contract.UnpackLog(event, "PayerAddressChanged", log); err != nil {
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

// HoQuAdCampaignStatusChangedIterator is returned from FilterStatusChanged and is used to iterate over the raw logs and unpacked data for StatusChanged events raised by the HoQuAdCampaign contract.
type HoQuAdCampaignStatusChangedIterator struct {
	Event *HoQuAdCampaignStatusChanged // Event containing the contract specifics and raw log

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
func (it *HoQuAdCampaignStatusChangedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(HoQuAdCampaignStatusChanged)
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
		it.Event = new(HoQuAdCampaignStatusChanged)
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
func (it *HoQuAdCampaignStatusChangedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *HoQuAdCampaignStatusChangedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// HoQuAdCampaignStatusChanged represents a StatusChanged event raised by the HoQuAdCampaign contract.
type HoQuAdCampaignStatusChanged struct {
	PayerAddress common.Address
	NewStatus    uint8
	Raw          types.Log // Blockchain specific contextual infos
}

// FilterStatusChanged is a free log retrieval operation binding the contract event 0x76c0bad3b7f1c74587dcb478154d4c87ae6663db97cc086b981decf4e35a797c.
//
// Solidity: event StatusChanged(payerAddress indexed address, newStatus uint8)
func (_HoQuAdCampaign *HoQuAdCampaignFilterer) FilterStatusChanged(opts *bind.FilterOpts, payerAddress []common.Address) (*HoQuAdCampaignStatusChangedIterator, error) {

	var payerAddressRule []interface{}
	for _, payerAddressItem := range payerAddress {
		payerAddressRule = append(payerAddressRule, payerAddressItem)
	}

	logs, sub, err := _HoQuAdCampaign.contract.FilterLogs(opts, "StatusChanged", payerAddressRule)
	if err != nil {
		return nil, err
	}
	return &HoQuAdCampaignStatusChangedIterator{contract: _HoQuAdCampaign.contract, event: "StatusChanged", logs: logs, sub: sub}, nil
}

// WatchStatusChanged is a free log subscription operation binding the contract event 0x76c0bad3b7f1c74587dcb478154d4c87ae6663db97cc086b981decf4e35a797c.
//
// Solidity: event StatusChanged(payerAddress indexed address, newStatus uint8)
func (_HoQuAdCampaign *HoQuAdCampaignFilterer) WatchStatusChanged(opts *bind.WatchOpts, sink chan<- *HoQuAdCampaignStatusChanged, payerAddress []common.Address) (event.Subscription, error) {

	var payerAddressRule []interface{}
	for _, payerAddressItem := range payerAddress {
		payerAddressRule = append(payerAddressRule, payerAddressItem)
	}

	logs, sub, err := _HoQuAdCampaign.contract.WatchLogs(opts, "StatusChanged", payerAddressRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(HoQuAdCampaignStatusChanged)
				if err := _HoQuAdCampaign.contract.UnpackLog(event, "StatusChanged", log); err != nil {
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

// HoQuAdCampaignTrackerAddedIterator is returned from FilterTrackerAdded and is used to iterate over the raw logs and unpacked data for TrackerAdded events raised by the HoQuAdCampaign contract.
type HoQuAdCampaignTrackerAddedIterator struct {
	Event *HoQuAdCampaignTrackerAdded // Event containing the contract specifics and raw log

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
func (it *HoQuAdCampaignTrackerAddedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(HoQuAdCampaignTrackerAdded)
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
		it.Event = new(HoQuAdCampaignTrackerAdded)
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
func (it *HoQuAdCampaignTrackerAddedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *HoQuAdCampaignTrackerAddedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// HoQuAdCampaignTrackerAdded represents a TrackerAdded event raised by the HoQuAdCampaign contract.
type HoQuAdCampaignTrackerAdded struct {
	OwnerAddress common.Address
	Id           [16]byte
	Raw          types.Log // Blockchain specific contextual infos
}

// FilterTrackerAdded is a free log retrieval operation binding the contract event 0xd7793ff1544e44df2a307af241d282820620c015d2e3b8cb17b480a7e47a48be.
//
// Solidity: event TrackerAdded(ownerAddress indexed address, id bytes16)
func (_HoQuAdCampaign *HoQuAdCampaignFilterer) FilterTrackerAdded(opts *bind.FilterOpts, ownerAddress []common.Address) (*HoQuAdCampaignTrackerAddedIterator, error) {

	var ownerAddressRule []interface{}
	for _, ownerAddressItem := range ownerAddress {
		ownerAddressRule = append(ownerAddressRule, ownerAddressItem)
	}

	logs, sub, err := _HoQuAdCampaign.contract.FilterLogs(opts, "TrackerAdded", ownerAddressRule)
	if err != nil {
		return nil, err
	}
	return &HoQuAdCampaignTrackerAddedIterator{contract: _HoQuAdCampaign.contract, event: "TrackerAdded", logs: logs, sub: sub}, nil
}

// WatchTrackerAdded is a free log subscription operation binding the contract event 0xd7793ff1544e44df2a307af241d282820620c015d2e3b8cb17b480a7e47a48be.
//
// Solidity: event TrackerAdded(ownerAddress indexed address, id bytes16)
func (_HoQuAdCampaign *HoQuAdCampaignFilterer) WatchTrackerAdded(opts *bind.WatchOpts, sink chan<- *HoQuAdCampaignTrackerAdded, ownerAddress []common.Address) (event.Subscription, error) {

	var ownerAddressRule []interface{}
	for _, ownerAddressItem := range ownerAddress {
		ownerAddressRule = append(ownerAddressRule, ownerAddressItem)
	}

	logs, sub, err := _HoQuAdCampaign.contract.WatchLogs(opts, "TrackerAdded", ownerAddressRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(HoQuAdCampaignTrackerAdded)
				if err := _HoQuAdCampaign.contract.UnpackLog(event, "TrackerAdded", log); err != nil {
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

// HoQuAdCampaignIABI is the input ABI used to generate the binding from.
const HoQuAdCampaignIABI = "[{\"constant\":false,\"inputs\":[{\"name\":\"id\",\"type\":\"bytes16\"},{\"name\":\"intermediaryAddress\",\"type\":\"address\"},{\"name\":\"percent\",\"type\":\"uint32\"}],\"name\":\"addLeadIntermediary\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"id\",\"type\":\"bytes16\"},{\"name\":\"trackerId\",\"type\":\"bytes16\"},{\"name\":\"meta\",\"type\":\"string\"},{\"name\":\"dataUrl\",\"type\":\"string\"},{\"name\":\"price\",\"type\":\"uint256\"}],\"name\":\"addLead\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"id\",\"type\":\"bytes16\"},{\"name\":\"status\",\"type\":\"uint8\"}],\"name\":\"setLeadStatus\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"id\",\"type\":\"bytes16\"}],\"name\":\"transactLead\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]"

// HoQuAdCampaignIBin is the compiled bytecode used for deploying new contracts.
const HoQuAdCampaignIBin = `0x`

// DeployHoQuAdCampaignI deploys a new Ethereum contract, binding an instance of HoQuAdCampaignI to it.
func DeployHoQuAdCampaignI(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *HoQuAdCampaignI, error) {
	parsed, err := abi.JSON(strings.NewReader(HoQuAdCampaignIABI))
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	address, tx, contract, err := bind.DeployContract(auth, parsed, common.FromHex(HoQuAdCampaignIBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &HoQuAdCampaignI{HoQuAdCampaignICaller: HoQuAdCampaignICaller{contract: contract}, HoQuAdCampaignITransactor: HoQuAdCampaignITransactor{contract: contract}, HoQuAdCampaignIFilterer: HoQuAdCampaignIFilterer{contract: contract}}, nil
}

// HoQuAdCampaignI is an auto generated Go binding around an Ethereum contract.
type HoQuAdCampaignI struct {
	HoQuAdCampaignICaller     // Read-only binding to the contract
	HoQuAdCampaignITransactor // Write-only binding to the contract
	HoQuAdCampaignIFilterer   // Log filterer for contract events
}

// HoQuAdCampaignICaller is an auto generated read-only Go binding around an Ethereum contract.
type HoQuAdCampaignICaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// HoQuAdCampaignITransactor is an auto generated write-only Go binding around an Ethereum contract.
type HoQuAdCampaignITransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// HoQuAdCampaignIFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type HoQuAdCampaignIFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// HoQuAdCampaignISession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type HoQuAdCampaignISession struct {
	Contract     *HoQuAdCampaignI  // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// HoQuAdCampaignICallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type HoQuAdCampaignICallerSession struct {
	Contract *HoQuAdCampaignICaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts          // Call options to use throughout this session
}

// HoQuAdCampaignITransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type HoQuAdCampaignITransactorSession struct {
	Contract     *HoQuAdCampaignITransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts          // Transaction auth options to use throughout this session
}

// HoQuAdCampaignIRaw is an auto generated low-level Go binding around an Ethereum contract.
type HoQuAdCampaignIRaw struct {
	Contract *HoQuAdCampaignI // Generic contract binding to access the raw methods on
}

// HoQuAdCampaignICallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type HoQuAdCampaignICallerRaw struct {
	Contract *HoQuAdCampaignICaller // Generic read-only contract binding to access the raw methods on
}

// HoQuAdCampaignITransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type HoQuAdCampaignITransactorRaw struct {
	Contract *HoQuAdCampaignITransactor // Generic write-only contract binding to access the raw methods on
}

// NewHoQuAdCampaignI creates a new instance of HoQuAdCampaignI, bound to a specific deployed contract.
func NewHoQuAdCampaignI(address common.Address, backend bind.ContractBackend) (*HoQuAdCampaignI, error) {
	contract, err := bindHoQuAdCampaignI(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &HoQuAdCampaignI{HoQuAdCampaignICaller: HoQuAdCampaignICaller{contract: contract}, HoQuAdCampaignITransactor: HoQuAdCampaignITransactor{contract: contract}, HoQuAdCampaignIFilterer: HoQuAdCampaignIFilterer{contract: contract}}, nil
}

// NewHoQuAdCampaignICaller creates a new read-only instance of HoQuAdCampaignI, bound to a specific deployed contract.
func NewHoQuAdCampaignICaller(address common.Address, caller bind.ContractCaller) (*HoQuAdCampaignICaller, error) {
	contract, err := bindHoQuAdCampaignI(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &HoQuAdCampaignICaller{contract: contract}, nil
}

// NewHoQuAdCampaignITransactor creates a new write-only instance of HoQuAdCampaignI, bound to a specific deployed contract.
func NewHoQuAdCampaignITransactor(address common.Address, transactor bind.ContractTransactor) (*HoQuAdCampaignITransactor, error) {
	contract, err := bindHoQuAdCampaignI(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &HoQuAdCampaignITransactor{contract: contract}, nil
}

// NewHoQuAdCampaignIFilterer creates a new log filterer instance of HoQuAdCampaignI, bound to a specific deployed contract.
func NewHoQuAdCampaignIFilterer(address common.Address, filterer bind.ContractFilterer) (*HoQuAdCampaignIFilterer, error) {
	contract, err := bindHoQuAdCampaignI(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &HoQuAdCampaignIFilterer{contract: contract}, nil
}

// bindHoQuAdCampaignI binds a generic wrapper to an already deployed contract.
func bindHoQuAdCampaignI(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(HoQuAdCampaignIABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_HoQuAdCampaignI *HoQuAdCampaignIRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _HoQuAdCampaignI.Contract.HoQuAdCampaignICaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_HoQuAdCampaignI *HoQuAdCampaignIRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _HoQuAdCampaignI.Contract.HoQuAdCampaignITransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_HoQuAdCampaignI *HoQuAdCampaignIRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _HoQuAdCampaignI.Contract.HoQuAdCampaignITransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_HoQuAdCampaignI *HoQuAdCampaignICallerRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _HoQuAdCampaignI.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_HoQuAdCampaignI *HoQuAdCampaignITransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _HoQuAdCampaignI.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_HoQuAdCampaignI *HoQuAdCampaignITransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _HoQuAdCampaignI.Contract.contract.Transact(opts, method, params...)
}

// AddLead is a paid mutator transaction binding the contract method 0x7a8187cf.
//
// Solidity: function addLead(id bytes16, trackerId bytes16, meta string, dataUrl string, price uint256) returns()
func (_HoQuAdCampaignI *HoQuAdCampaignITransactor) AddLead(opts *bind.TransactOpts, id [16]byte, trackerId [16]byte, meta string, dataUrl string, price *big.Int) (*types.Transaction, error) {
	return _HoQuAdCampaignI.contract.Transact(opts, "addLead", id, trackerId, meta, dataUrl, price)
}

// AddLead is a paid mutator transaction binding the contract method 0x7a8187cf.
//
// Solidity: function addLead(id bytes16, trackerId bytes16, meta string, dataUrl string, price uint256) returns()
func (_HoQuAdCampaignI *HoQuAdCampaignISession) AddLead(id [16]byte, trackerId [16]byte, meta string, dataUrl string, price *big.Int) (*types.Transaction, error) {
	return _HoQuAdCampaignI.Contract.AddLead(&_HoQuAdCampaignI.TransactOpts, id, trackerId, meta, dataUrl, price)
}

// AddLead is a paid mutator transaction binding the contract method 0x7a8187cf.
//
// Solidity: function addLead(id bytes16, trackerId bytes16, meta string, dataUrl string, price uint256) returns()
func (_HoQuAdCampaignI *HoQuAdCampaignITransactorSession) AddLead(id [16]byte, trackerId [16]byte, meta string, dataUrl string, price *big.Int) (*types.Transaction, error) {
	return _HoQuAdCampaignI.Contract.AddLead(&_HoQuAdCampaignI.TransactOpts, id, trackerId, meta, dataUrl, price)
}

// AddLeadIntermediary is a paid mutator transaction binding the contract method 0x36617045.
//
// Solidity: function addLeadIntermediary(id bytes16, intermediaryAddress address, percent uint32) returns()
func (_HoQuAdCampaignI *HoQuAdCampaignITransactor) AddLeadIntermediary(opts *bind.TransactOpts, id [16]byte, intermediaryAddress common.Address, percent uint32) (*types.Transaction, error) {
	return _HoQuAdCampaignI.contract.Transact(opts, "addLeadIntermediary", id, intermediaryAddress, percent)
}

// AddLeadIntermediary is a paid mutator transaction binding the contract method 0x36617045.
//
// Solidity: function addLeadIntermediary(id bytes16, intermediaryAddress address, percent uint32) returns()
func (_HoQuAdCampaignI *HoQuAdCampaignISession) AddLeadIntermediary(id [16]byte, intermediaryAddress common.Address, percent uint32) (*types.Transaction, error) {
	return _HoQuAdCampaignI.Contract.AddLeadIntermediary(&_HoQuAdCampaignI.TransactOpts, id, intermediaryAddress, percent)
}

// AddLeadIntermediary is a paid mutator transaction binding the contract method 0x36617045.
//
// Solidity: function addLeadIntermediary(id bytes16, intermediaryAddress address, percent uint32) returns()
func (_HoQuAdCampaignI *HoQuAdCampaignITransactorSession) AddLeadIntermediary(id [16]byte, intermediaryAddress common.Address, percent uint32) (*types.Transaction, error) {
	return _HoQuAdCampaignI.Contract.AddLeadIntermediary(&_HoQuAdCampaignI.TransactOpts, id, intermediaryAddress, percent)
}

// SetLeadStatus is a paid mutator transaction binding the contract method 0x81e85aa3.
//
// Solidity: function setLeadStatus(id bytes16, status uint8) returns()
func (_HoQuAdCampaignI *HoQuAdCampaignITransactor) SetLeadStatus(opts *bind.TransactOpts, id [16]byte, status uint8) (*types.Transaction, error) {
	return _HoQuAdCampaignI.contract.Transact(opts, "setLeadStatus", id, status)
}

// SetLeadStatus is a paid mutator transaction binding the contract method 0x81e85aa3.
//
// Solidity: function setLeadStatus(id bytes16, status uint8) returns()
func (_HoQuAdCampaignI *HoQuAdCampaignISession) SetLeadStatus(id [16]byte, status uint8) (*types.Transaction, error) {
	return _HoQuAdCampaignI.Contract.SetLeadStatus(&_HoQuAdCampaignI.TransactOpts, id, status)
}

// SetLeadStatus is a paid mutator transaction binding the contract method 0x81e85aa3.
//
// Solidity: function setLeadStatus(id bytes16, status uint8) returns()
func (_HoQuAdCampaignI *HoQuAdCampaignITransactorSession) SetLeadStatus(id [16]byte, status uint8) (*types.Transaction, error) {
	return _HoQuAdCampaignI.Contract.SetLeadStatus(&_HoQuAdCampaignI.TransactOpts, id, status)
}

// TransactLead is a paid mutator transaction binding the contract method 0x835aef40.
//
// Solidity: function transactLead(id bytes16) returns()
func (_HoQuAdCampaignI *HoQuAdCampaignITransactor) TransactLead(opts *bind.TransactOpts, id [16]byte) (*types.Transaction, error) {
	return _HoQuAdCampaignI.contract.Transact(opts, "transactLead", id)
}

// TransactLead is a paid mutator transaction binding the contract method 0x835aef40.
//
// Solidity: function transactLead(id bytes16) returns()
func (_HoQuAdCampaignI *HoQuAdCampaignISession) TransactLead(id [16]byte) (*types.Transaction, error) {
	return _HoQuAdCampaignI.Contract.TransactLead(&_HoQuAdCampaignI.TransactOpts, id)
}

// TransactLead is a paid mutator transaction binding the contract method 0x835aef40.
//
// Solidity: function transactLead(id bytes16) returns()
func (_HoQuAdCampaignI *HoQuAdCampaignITransactorSession) TransactLead(id [16]byte) (*types.Transaction, error) {
	return _HoQuAdCampaignI.Contract.TransactLead(&_HoQuAdCampaignI.TransactOpts, id)
}

// HoQuConfigABI is the input ABI used to generate the binding from.
const HoQuConfigABI = "[{\"constant\":false,\"inputs\":[{\"name\":\"_commission\",\"type\":\"uint256\"}],\"name\":\"setCommission\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_owner\",\"type\":\"address\"}],\"name\":\"addOwner\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[],\"name\":\"renounceOwnership\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_commissionWallet\",\"type\":\"address\"}],\"name\":\"setCommissionWallet\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"i\",\"type\":\"uint16\"}],\"name\":\"deleteOwner\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"commissionWallet\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"i\",\"type\":\"uint16\"},{\"name\":\"_owner\",\"type\":\"address\"}],\"name\":\"changeOwner\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"ownersCount\",\"outputs\":[{\"name\":\"\",\"type\":\"uint16\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_owner\",\"type\":\"address\"}],\"name\":\"isAllowed\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"\",\"type\":\"uint16\"}],\"name\":\"owners\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"commission\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"transferOwnership\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"name\":\"_commissionWallet\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"changedBy\",\"type\":\"address\"},{\"indexed\":false,\"name\":\"commissionWallet\",\"type\":\"address\"}],\"name\":\"CommissionWalletChanged\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"changedBy\",\"type\":\"address\"},{\"indexed\":false,\"name\":\"commission\",\"type\":\"uint256\"}],\"name\":\"CommissionChanged\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"SystemOwnerAdded\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"previousOwner\",\"type\":\"address\"},{\"indexed\":false,\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"SystemOwnerChanged\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"deletedOwner\",\"type\":\"address\"}],\"name\":\"SystemOwnerDeleted\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"previousOwner\",\"type\":\"address\"}],\"name\":\"OwnershipRenounced\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"previousOwner\",\"type\":\"address\"},{\"indexed\":true,\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"OwnershipTransferred\",\"type\":\"event\"}]"

// HoQuConfigBin is the compiled bytecode used for deploying new contracts.
const HoQuConfigBin = `0x60806040526611c37937e0800060025534801561001b57600080fd5b5060405160208061089d833981016040525160008054600160a060020a03338116600160a060020a03199283161790925560018054929093169116179055610835806100686000396000f3006080604052600436106100c45763ffffffff7c0100000000000000000000000000000000000000000000000000000000600035041663355e6b4381146100c95780637065cb48146100e3578063715018a6146101045780637d60b6ce146101195780638d2573321461013a5780638da5cb5b1461015657806397c0262a14610187578063b7b603a21461019c578063b9488546146101c4578063babcc539146101f0578063c15d25b114610225578063e148919114610241578063f2fde38b14610268575b600080fd5b3480156100d557600080fd5b506100e1600435610289565b005b3480156100ef57600080fd5b506100e1600160a060020a0360043516610306565b34801561011057600080fd5b506100e16103b6565b34801561012557600080fd5b506100e1600160a060020a0360043516610426565b34801561014657600080fd5b506100e161ffff600435166104d9565b34801561016257600080fd5b5061016b6105a2565b60408051600160a060020a039092168252519081900360200190f35b34801561019357600080fd5b5061016b6105b1565b3480156101a857600080fd5b506100e161ffff60043516600160a060020a03602435166105c0565b3480156101d057600080fd5b506101d96106ad565b6040805161ffff9092168252519081900360200190f35b3480156101fc57600080fd5b50610211600160a060020a03600435166106b7565b604080519115158252519081900360200190f35b34801561023157600080fd5b5061016b61ffff60043516610750565b34801561024d57600080fd5b5061025661076b565b60408051918252519081900360200190f35b34801561027457600080fd5b506100e1600160a060020a0360043516610771565b60005433600160a060020a03908116911614806102aa57506102aa336106b7565b15156102b557600080fd5b600081116102c257600080fd5b604080518281529051600160a060020a033316917ff659af9785c5d8c6128fdcbabb637f7d7ea787226f13cad89f71cbc579614fdc919081900360200190a2600255565b60005433600160a060020a03908116911614806103275750610327336106b7565b151561033257600080fd5b6004805461ffff908116600090815260036020526040808220805473ffffffffffffffffffffffffffffffffffffffff1916600160a060020a038716908117909155845461ffff19811690851660010190941693909317909355915190917fa91384695ab8222d2f6e2437bb5b1fefe671538a33a7e06e88fee0147fffe8a491a250565b60005433600160a060020a039081169116146103d157600080fd5b60008054604051600160a060020a03909116917ff8df31144d9c2f0f6b59d69b8b98abd5459d07f2742c4df920b25aae33c6482091a26000805473ffffffffffffffffffffffffffffffffffffffff19169055565b60005433600160a060020a03908116911614806104475750610447336106b7565b151561045257600080fd5b600160a060020a038116151561046757600080fd5b60408051600160a060020a03838116825291513392909216917f394e7a1cfafd7b3c49218efbf47fb95da1832e5eaf0e9dc0428ea7220959d0e09181900360200190a26001805473ffffffffffffffffffffffffffffffffffffffff1916600160a060020a0392909216919091179055565b60005433600160a060020a03908116911614806104fa57506104fa336106b7565b151561050557600080fd5b61ffff8116600090815260036020526040902054600160a060020a0316151561052d57600080fd5b61ffff8116600090815260036020526040808220549051600160a060020a03909116917f7ac6b089a9620017a12220abae1696205383eb917822ade81acb4eea2a424f6991a261ffff166000908152600360205260409020805473ffffffffffffffffffffffffffffffffffffffff19169055565b600054600160a060020a031681565b600154600160a060020a031681565b60005433600160a060020a03908116911614806105e157506105e1336106b7565b15156105ec57600080fd5b61ffff8216600090815260036020526040902054600160a060020a03161561061357600080fd5b61ffff8216600090815260036020908152604091829020548251600160a060020a038581168252935193909116927f1722097e017bb9a189312ac57bbdfff1bf251383cfd36c245f8ebabd8734c759929181900390910190a261ffff919091166000908152600360205260409020805473ffffffffffffffffffffffffffffffffffffffff1916600160a060020a03909216919091179055565b60045461ffff1681565b60008080600160a060020a03841615156106d057600080fd5b600091505b60045461ffff9081169083161015610725575061ffff8116600090815260036020526040902054600160a060020a0390811690841681141561071a5760019250610749565b6001909101906106d5565b600054600160a060020a03858116911614156107445760019250610749565b600092505b5050919050565b600360205260009081526040902054600160a060020a031681565b60025481565b60005433600160a060020a0390811691161461078c57600080fd5b600160a060020a03811615156107a157600080fd5b60008054604051600160a060020a03808516939216917f8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e091a36000805473ffffffffffffffffffffffffffffffffffffffff1916600160a060020a03929092169190911790555600a165627a7a72305820cea82f55d49da6825e1ae2fe9b5ae267907413da5c080c1df68782659f27294f0029`

// DeployHoQuConfig deploys a new Ethereum contract, binding an instance of HoQuConfig to it.
func DeployHoQuConfig(auth *bind.TransactOpts, backend bind.ContractBackend, _commissionWallet common.Address) (common.Address, *types.Transaction, *HoQuConfig, error) {
	parsed, err := abi.JSON(strings.NewReader(HoQuConfigABI))
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	address, tx, contract, err := bind.DeployContract(auth, parsed, common.FromHex(HoQuConfigBin), backend, _commissionWallet)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &HoQuConfig{HoQuConfigCaller: HoQuConfigCaller{contract: contract}, HoQuConfigTransactor: HoQuConfigTransactor{contract: contract}, HoQuConfigFilterer: HoQuConfigFilterer{contract: contract}}, nil
}

// HoQuConfig is an auto generated Go binding around an Ethereum contract.
type HoQuConfig struct {
	HoQuConfigCaller     // Read-only binding to the contract
	HoQuConfigTransactor // Write-only binding to the contract
	HoQuConfigFilterer   // Log filterer for contract events
}

// HoQuConfigCaller is an auto generated read-only Go binding around an Ethereum contract.
type HoQuConfigCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// HoQuConfigTransactor is an auto generated write-only Go binding around an Ethereum contract.
type HoQuConfigTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// HoQuConfigFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type HoQuConfigFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// HoQuConfigSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type HoQuConfigSession struct {
	Contract     *HoQuConfig       // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// HoQuConfigCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type HoQuConfigCallerSession struct {
	Contract *HoQuConfigCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts     // Call options to use throughout this session
}

// HoQuConfigTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type HoQuConfigTransactorSession struct {
	Contract     *HoQuConfigTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts     // Transaction auth options to use throughout this session
}

// HoQuConfigRaw is an auto generated low-level Go binding around an Ethereum contract.
type HoQuConfigRaw struct {
	Contract *HoQuConfig // Generic contract binding to access the raw methods on
}

// HoQuConfigCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type HoQuConfigCallerRaw struct {
	Contract *HoQuConfigCaller // Generic read-only contract binding to access the raw methods on
}

// HoQuConfigTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type HoQuConfigTransactorRaw struct {
	Contract *HoQuConfigTransactor // Generic write-only contract binding to access the raw methods on
}

// NewHoQuConfig creates a new instance of HoQuConfig, bound to a specific deployed contract.
func NewHoQuConfig(address common.Address, backend bind.ContractBackend) (*HoQuConfig, error) {
	contract, err := bindHoQuConfig(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &HoQuConfig{HoQuConfigCaller: HoQuConfigCaller{contract: contract}, HoQuConfigTransactor: HoQuConfigTransactor{contract: contract}, HoQuConfigFilterer: HoQuConfigFilterer{contract: contract}}, nil
}

// NewHoQuConfigCaller creates a new read-only instance of HoQuConfig, bound to a specific deployed contract.
func NewHoQuConfigCaller(address common.Address, caller bind.ContractCaller) (*HoQuConfigCaller, error) {
	contract, err := bindHoQuConfig(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &HoQuConfigCaller{contract: contract}, nil
}

// NewHoQuConfigTransactor creates a new write-only instance of HoQuConfig, bound to a specific deployed contract.
func NewHoQuConfigTransactor(address common.Address, transactor bind.ContractTransactor) (*HoQuConfigTransactor, error) {
	contract, err := bindHoQuConfig(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &HoQuConfigTransactor{contract: contract}, nil
}

// NewHoQuConfigFilterer creates a new log filterer instance of HoQuConfig, bound to a specific deployed contract.
func NewHoQuConfigFilterer(address common.Address, filterer bind.ContractFilterer) (*HoQuConfigFilterer, error) {
	contract, err := bindHoQuConfig(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &HoQuConfigFilterer{contract: contract}, nil
}

// bindHoQuConfig binds a generic wrapper to an already deployed contract.
func bindHoQuConfig(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(HoQuConfigABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_HoQuConfig *HoQuConfigRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _HoQuConfig.Contract.HoQuConfigCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_HoQuConfig *HoQuConfigRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _HoQuConfig.Contract.HoQuConfigTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_HoQuConfig *HoQuConfigRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _HoQuConfig.Contract.HoQuConfigTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_HoQuConfig *HoQuConfigCallerRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _HoQuConfig.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_HoQuConfig *HoQuConfigTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _HoQuConfig.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_HoQuConfig *HoQuConfigTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _HoQuConfig.Contract.contract.Transact(opts, method, params...)
}

// Commission is a free data retrieval call binding the contract method 0xe1489191.
//
// Solidity: function commission() constant returns(uint256)
func (_HoQuConfig *HoQuConfigCaller) Commission(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _HoQuConfig.contract.Call(opts, out, "commission")
	return *ret0, err
}

// Commission is a free data retrieval call binding the contract method 0xe1489191.
//
// Solidity: function commission() constant returns(uint256)
func (_HoQuConfig *HoQuConfigSession) Commission() (*big.Int, error) {
	return _HoQuConfig.Contract.Commission(&_HoQuConfig.CallOpts)
}

// Commission is a free data retrieval call binding the contract method 0xe1489191.
//
// Solidity: function commission() constant returns(uint256)
func (_HoQuConfig *HoQuConfigCallerSession) Commission() (*big.Int, error) {
	return _HoQuConfig.Contract.Commission(&_HoQuConfig.CallOpts)
}

// CommissionWallet is a free data retrieval call binding the contract method 0x97c0262a.
//
// Solidity: function commissionWallet() constant returns(address)
func (_HoQuConfig *HoQuConfigCaller) CommissionWallet(opts *bind.CallOpts) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _HoQuConfig.contract.Call(opts, out, "commissionWallet")
	return *ret0, err
}

// CommissionWallet is a free data retrieval call binding the contract method 0x97c0262a.
//
// Solidity: function commissionWallet() constant returns(address)
func (_HoQuConfig *HoQuConfigSession) CommissionWallet() (common.Address, error) {
	return _HoQuConfig.Contract.CommissionWallet(&_HoQuConfig.CallOpts)
}

// CommissionWallet is a free data retrieval call binding the contract method 0x97c0262a.
//
// Solidity: function commissionWallet() constant returns(address)
func (_HoQuConfig *HoQuConfigCallerSession) CommissionWallet() (common.Address, error) {
	return _HoQuConfig.Contract.CommissionWallet(&_HoQuConfig.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() constant returns(address)
func (_HoQuConfig *HoQuConfigCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _HoQuConfig.contract.Call(opts, out, "owner")
	return *ret0, err
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() constant returns(address)
func (_HoQuConfig *HoQuConfigSession) Owner() (common.Address, error) {
	return _HoQuConfig.Contract.Owner(&_HoQuConfig.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() constant returns(address)
func (_HoQuConfig *HoQuConfigCallerSession) Owner() (common.Address, error) {
	return _HoQuConfig.Contract.Owner(&_HoQuConfig.CallOpts)
}

// Owners is a free data retrieval call binding the contract method 0xc15d25b1.
//
// Solidity: function owners( uint16) constant returns(address)
func (_HoQuConfig *HoQuConfigCaller) Owners(opts *bind.CallOpts, arg0 uint16) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _HoQuConfig.contract.Call(opts, out, "owners", arg0)
	return *ret0, err
}

// Owners is a free data retrieval call binding the contract method 0xc15d25b1.
//
// Solidity: function owners( uint16) constant returns(address)
func (_HoQuConfig *HoQuConfigSession) Owners(arg0 uint16) (common.Address, error) {
	return _HoQuConfig.Contract.Owners(&_HoQuConfig.CallOpts, arg0)
}

// Owners is a free data retrieval call binding the contract method 0xc15d25b1.
//
// Solidity: function owners( uint16) constant returns(address)
func (_HoQuConfig *HoQuConfigCallerSession) Owners(arg0 uint16) (common.Address, error) {
	return _HoQuConfig.Contract.Owners(&_HoQuConfig.CallOpts, arg0)
}

// OwnersCount is a free data retrieval call binding the contract method 0xb9488546.
//
// Solidity: function ownersCount() constant returns(uint16)
func (_HoQuConfig *HoQuConfigCaller) OwnersCount(opts *bind.CallOpts) (uint16, error) {
	var (
		ret0 = new(uint16)
	)
	out := ret0
	err := _HoQuConfig.contract.Call(opts, out, "ownersCount")
	return *ret0, err
}

// OwnersCount is a free data retrieval call binding the contract method 0xb9488546.
//
// Solidity: function ownersCount() constant returns(uint16)
func (_HoQuConfig *HoQuConfigSession) OwnersCount() (uint16, error) {
	return _HoQuConfig.Contract.OwnersCount(&_HoQuConfig.CallOpts)
}

// OwnersCount is a free data retrieval call binding the contract method 0xb9488546.
//
// Solidity: function ownersCount() constant returns(uint16)
func (_HoQuConfig *HoQuConfigCallerSession) OwnersCount() (uint16, error) {
	return _HoQuConfig.Contract.OwnersCount(&_HoQuConfig.CallOpts)
}

// AddOwner is a paid mutator transaction binding the contract method 0x7065cb48.
//
// Solidity: function addOwner(_owner address) returns()
func (_HoQuConfig *HoQuConfigTransactor) AddOwner(opts *bind.TransactOpts, _owner common.Address) (*types.Transaction, error) {
	return _HoQuConfig.contract.Transact(opts, "addOwner", _owner)
}

// AddOwner is a paid mutator transaction binding the contract method 0x7065cb48.
//
// Solidity: function addOwner(_owner address) returns()
func (_HoQuConfig *HoQuConfigSession) AddOwner(_owner common.Address) (*types.Transaction, error) {
	return _HoQuConfig.Contract.AddOwner(&_HoQuConfig.TransactOpts, _owner)
}

// AddOwner is a paid mutator transaction binding the contract method 0x7065cb48.
//
// Solidity: function addOwner(_owner address) returns()
func (_HoQuConfig *HoQuConfigTransactorSession) AddOwner(_owner common.Address) (*types.Transaction, error) {
	return _HoQuConfig.Contract.AddOwner(&_HoQuConfig.TransactOpts, _owner)
}

// ChangeOwner is a paid mutator transaction binding the contract method 0xb7b603a2.
//
// Solidity: function changeOwner(i uint16, _owner address) returns()
func (_HoQuConfig *HoQuConfigTransactor) ChangeOwner(opts *bind.TransactOpts, i uint16, _owner common.Address) (*types.Transaction, error) {
	return _HoQuConfig.contract.Transact(opts, "changeOwner", i, _owner)
}

// ChangeOwner is a paid mutator transaction binding the contract method 0xb7b603a2.
//
// Solidity: function changeOwner(i uint16, _owner address) returns()
func (_HoQuConfig *HoQuConfigSession) ChangeOwner(i uint16, _owner common.Address) (*types.Transaction, error) {
	return _HoQuConfig.Contract.ChangeOwner(&_HoQuConfig.TransactOpts, i, _owner)
}

// ChangeOwner is a paid mutator transaction binding the contract method 0xb7b603a2.
//
// Solidity: function changeOwner(i uint16, _owner address) returns()
func (_HoQuConfig *HoQuConfigTransactorSession) ChangeOwner(i uint16, _owner common.Address) (*types.Transaction, error) {
	return _HoQuConfig.Contract.ChangeOwner(&_HoQuConfig.TransactOpts, i, _owner)
}

// DeleteOwner is a paid mutator transaction binding the contract method 0x8d257332.
//
// Solidity: function deleteOwner(i uint16) returns()
func (_HoQuConfig *HoQuConfigTransactor) DeleteOwner(opts *bind.TransactOpts, i uint16) (*types.Transaction, error) {
	return _HoQuConfig.contract.Transact(opts, "deleteOwner", i)
}

// DeleteOwner is a paid mutator transaction binding the contract method 0x8d257332.
//
// Solidity: function deleteOwner(i uint16) returns()
func (_HoQuConfig *HoQuConfigSession) DeleteOwner(i uint16) (*types.Transaction, error) {
	return _HoQuConfig.Contract.DeleteOwner(&_HoQuConfig.TransactOpts, i)
}

// DeleteOwner is a paid mutator transaction binding the contract method 0x8d257332.
//
// Solidity: function deleteOwner(i uint16) returns()
func (_HoQuConfig *HoQuConfigTransactorSession) DeleteOwner(i uint16) (*types.Transaction, error) {
	return _HoQuConfig.Contract.DeleteOwner(&_HoQuConfig.TransactOpts, i)
}

// IsAllowed is a paid mutator transaction binding the contract method 0xbabcc539.
//
// Solidity: function isAllowed(_owner address) returns(bool)
func (_HoQuConfig *HoQuConfigTransactor) IsAllowed(opts *bind.TransactOpts, _owner common.Address) (*types.Transaction, error) {
	return _HoQuConfig.contract.Transact(opts, "isAllowed", _owner)
}

// IsAllowed is a paid mutator transaction binding the contract method 0xbabcc539.
//
// Solidity: function isAllowed(_owner address) returns(bool)
func (_HoQuConfig *HoQuConfigSession) IsAllowed(_owner common.Address) (*types.Transaction, error) {
	return _HoQuConfig.Contract.IsAllowed(&_HoQuConfig.TransactOpts, _owner)
}

// IsAllowed is a paid mutator transaction binding the contract method 0xbabcc539.
//
// Solidity: function isAllowed(_owner address) returns(bool)
func (_HoQuConfig *HoQuConfigTransactorSession) IsAllowed(_owner common.Address) (*types.Transaction, error) {
	return _HoQuConfig.Contract.IsAllowed(&_HoQuConfig.TransactOpts, _owner)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_HoQuConfig *HoQuConfigTransactor) RenounceOwnership(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _HoQuConfig.contract.Transact(opts, "renounceOwnership")
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_HoQuConfig *HoQuConfigSession) RenounceOwnership() (*types.Transaction, error) {
	return _HoQuConfig.Contract.RenounceOwnership(&_HoQuConfig.TransactOpts)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_HoQuConfig *HoQuConfigTransactorSession) RenounceOwnership() (*types.Transaction, error) {
	return _HoQuConfig.Contract.RenounceOwnership(&_HoQuConfig.TransactOpts)
}

// SetCommission is a paid mutator transaction binding the contract method 0x355e6b43.
//
// Solidity: function setCommission(_commission uint256) returns()
func (_HoQuConfig *HoQuConfigTransactor) SetCommission(opts *bind.TransactOpts, _commission *big.Int) (*types.Transaction, error) {
	return _HoQuConfig.contract.Transact(opts, "setCommission", _commission)
}

// SetCommission is a paid mutator transaction binding the contract method 0x355e6b43.
//
// Solidity: function setCommission(_commission uint256) returns()
func (_HoQuConfig *HoQuConfigSession) SetCommission(_commission *big.Int) (*types.Transaction, error) {
	return _HoQuConfig.Contract.SetCommission(&_HoQuConfig.TransactOpts, _commission)
}

// SetCommission is a paid mutator transaction binding the contract method 0x355e6b43.
//
// Solidity: function setCommission(_commission uint256) returns()
func (_HoQuConfig *HoQuConfigTransactorSession) SetCommission(_commission *big.Int) (*types.Transaction, error) {
	return _HoQuConfig.Contract.SetCommission(&_HoQuConfig.TransactOpts, _commission)
}

// SetCommissionWallet is a paid mutator transaction binding the contract method 0x7d60b6ce.
//
// Solidity: function setCommissionWallet(_commissionWallet address) returns()
func (_HoQuConfig *HoQuConfigTransactor) SetCommissionWallet(opts *bind.TransactOpts, _commissionWallet common.Address) (*types.Transaction, error) {
	return _HoQuConfig.contract.Transact(opts, "setCommissionWallet", _commissionWallet)
}

// SetCommissionWallet is a paid mutator transaction binding the contract method 0x7d60b6ce.
//
// Solidity: function setCommissionWallet(_commissionWallet address) returns()
func (_HoQuConfig *HoQuConfigSession) SetCommissionWallet(_commissionWallet common.Address) (*types.Transaction, error) {
	return _HoQuConfig.Contract.SetCommissionWallet(&_HoQuConfig.TransactOpts, _commissionWallet)
}

// SetCommissionWallet is a paid mutator transaction binding the contract method 0x7d60b6ce.
//
// Solidity: function setCommissionWallet(_commissionWallet address) returns()
func (_HoQuConfig *HoQuConfigTransactorSession) SetCommissionWallet(_commissionWallet common.Address) (*types.Transaction, error) {
	return _HoQuConfig.Contract.SetCommissionWallet(&_HoQuConfig.TransactOpts, _commissionWallet)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(newOwner address) returns()
func (_HoQuConfig *HoQuConfigTransactor) TransferOwnership(opts *bind.TransactOpts, newOwner common.Address) (*types.Transaction, error) {
	return _HoQuConfig.contract.Transact(opts, "transferOwnership", newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(newOwner address) returns()
func (_HoQuConfig *HoQuConfigSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _HoQuConfig.Contract.TransferOwnership(&_HoQuConfig.TransactOpts, newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(newOwner address) returns()
func (_HoQuConfig *HoQuConfigTransactorSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _HoQuConfig.Contract.TransferOwnership(&_HoQuConfig.TransactOpts, newOwner)
}

// HoQuConfigCommissionChangedIterator is returned from FilterCommissionChanged and is used to iterate over the raw logs and unpacked data for CommissionChanged events raised by the HoQuConfig contract.
type HoQuConfigCommissionChangedIterator struct {
	Event *HoQuConfigCommissionChanged // Event containing the contract specifics and raw log

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
func (it *HoQuConfigCommissionChangedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(HoQuConfigCommissionChanged)
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
		it.Event = new(HoQuConfigCommissionChanged)
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
func (it *HoQuConfigCommissionChangedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *HoQuConfigCommissionChangedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// HoQuConfigCommissionChanged represents a CommissionChanged event raised by the HoQuConfig contract.
type HoQuConfigCommissionChanged struct {
	ChangedBy  common.Address
	Commission *big.Int
	Raw        types.Log // Blockchain specific contextual infos
}

// FilterCommissionChanged is a free log retrieval operation binding the contract event 0xf659af9785c5d8c6128fdcbabb637f7d7ea787226f13cad89f71cbc579614fdc.
//
// Solidity: event CommissionChanged(changedBy indexed address, commission uint256)
func (_HoQuConfig *HoQuConfigFilterer) FilterCommissionChanged(opts *bind.FilterOpts, changedBy []common.Address) (*HoQuConfigCommissionChangedIterator, error) {

	var changedByRule []interface{}
	for _, changedByItem := range changedBy {
		changedByRule = append(changedByRule, changedByItem)
	}

	logs, sub, err := _HoQuConfig.contract.FilterLogs(opts, "CommissionChanged", changedByRule)
	if err != nil {
		return nil, err
	}
	return &HoQuConfigCommissionChangedIterator{contract: _HoQuConfig.contract, event: "CommissionChanged", logs: logs, sub: sub}, nil
}

// WatchCommissionChanged is a free log subscription operation binding the contract event 0xf659af9785c5d8c6128fdcbabb637f7d7ea787226f13cad89f71cbc579614fdc.
//
// Solidity: event CommissionChanged(changedBy indexed address, commission uint256)
func (_HoQuConfig *HoQuConfigFilterer) WatchCommissionChanged(opts *bind.WatchOpts, sink chan<- *HoQuConfigCommissionChanged, changedBy []common.Address) (event.Subscription, error) {

	var changedByRule []interface{}
	for _, changedByItem := range changedBy {
		changedByRule = append(changedByRule, changedByItem)
	}

	logs, sub, err := _HoQuConfig.contract.WatchLogs(opts, "CommissionChanged", changedByRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(HoQuConfigCommissionChanged)
				if err := _HoQuConfig.contract.UnpackLog(event, "CommissionChanged", log); err != nil {
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

// HoQuConfigCommissionWalletChangedIterator is returned from FilterCommissionWalletChanged and is used to iterate over the raw logs and unpacked data for CommissionWalletChanged events raised by the HoQuConfig contract.
type HoQuConfigCommissionWalletChangedIterator struct {
	Event *HoQuConfigCommissionWalletChanged // Event containing the contract specifics and raw log

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
func (it *HoQuConfigCommissionWalletChangedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(HoQuConfigCommissionWalletChanged)
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
		it.Event = new(HoQuConfigCommissionWalletChanged)
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
func (it *HoQuConfigCommissionWalletChangedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *HoQuConfigCommissionWalletChangedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// HoQuConfigCommissionWalletChanged represents a CommissionWalletChanged event raised by the HoQuConfig contract.
type HoQuConfigCommissionWalletChanged struct {
	ChangedBy        common.Address
	CommissionWallet common.Address
	Raw              types.Log // Blockchain specific contextual infos
}

// FilterCommissionWalletChanged is a free log retrieval operation binding the contract event 0x394e7a1cfafd7b3c49218efbf47fb95da1832e5eaf0e9dc0428ea7220959d0e0.
//
// Solidity: event CommissionWalletChanged(changedBy indexed address, commissionWallet address)
func (_HoQuConfig *HoQuConfigFilterer) FilterCommissionWalletChanged(opts *bind.FilterOpts, changedBy []common.Address) (*HoQuConfigCommissionWalletChangedIterator, error) {

	var changedByRule []interface{}
	for _, changedByItem := range changedBy {
		changedByRule = append(changedByRule, changedByItem)
	}

	logs, sub, err := _HoQuConfig.contract.FilterLogs(opts, "CommissionWalletChanged", changedByRule)
	if err != nil {
		return nil, err
	}
	return &HoQuConfigCommissionWalletChangedIterator{contract: _HoQuConfig.contract, event: "CommissionWalletChanged", logs: logs, sub: sub}, nil
}

// WatchCommissionWalletChanged is a free log subscription operation binding the contract event 0x394e7a1cfafd7b3c49218efbf47fb95da1832e5eaf0e9dc0428ea7220959d0e0.
//
// Solidity: event CommissionWalletChanged(changedBy indexed address, commissionWallet address)
func (_HoQuConfig *HoQuConfigFilterer) WatchCommissionWalletChanged(opts *bind.WatchOpts, sink chan<- *HoQuConfigCommissionWalletChanged, changedBy []common.Address) (event.Subscription, error) {

	var changedByRule []interface{}
	for _, changedByItem := range changedBy {
		changedByRule = append(changedByRule, changedByItem)
	}

	logs, sub, err := _HoQuConfig.contract.WatchLogs(opts, "CommissionWalletChanged", changedByRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(HoQuConfigCommissionWalletChanged)
				if err := _HoQuConfig.contract.UnpackLog(event, "CommissionWalletChanged", log); err != nil {
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

// HoQuConfigOwnershipRenouncedIterator is returned from FilterOwnershipRenounced and is used to iterate over the raw logs and unpacked data for OwnershipRenounced events raised by the HoQuConfig contract.
type HoQuConfigOwnershipRenouncedIterator struct {
	Event *HoQuConfigOwnershipRenounced // Event containing the contract specifics and raw log

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
func (it *HoQuConfigOwnershipRenouncedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(HoQuConfigOwnershipRenounced)
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
		it.Event = new(HoQuConfigOwnershipRenounced)
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
func (it *HoQuConfigOwnershipRenouncedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *HoQuConfigOwnershipRenouncedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// HoQuConfigOwnershipRenounced represents a OwnershipRenounced event raised by the HoQuConfig contract.
type HoQuConfigOwnershipRenounced struct {
	PreviousOwner common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterOwnershipRenounced is a free log retrieval operation binding the contract event 0xf8df31144d9c2f0f6b59d69b8b98abd5459d07f2742c4df920b25aae33c64820.
//
// Solidity: event OwnershipRenounced(previousOwner indexed address)
func (_HoQuConfig *HoQuConfigFilterer) FilterOwnershipRenounced(opts *bind.FilterOpts, previousOwner []common.Address) (*HoQuConfigOwnershipRenouncedIterator, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}

	logs, sub, err := _HoQuConfig.contract.FilterLogs(opts, "OwnershipRenounced", previousOwnerRule)
	if err != nil {
		return nil, err
	}
	return &HoQuConfigOwnershipRenouncedIterator{contract: _HoQuConfig.contract, event: "OwnershipRenounced", logs: logs, sub: sub}, nil
}

// WatchOwnershipRenounced is a free log subscription operation binding the contract event 0xf8df31144d9c2f0f6b59d69b8b98abd5459d07f2742c4df920b25aae33c64820.
//
// Solidity: event OwnershipRenounced(previousOwner indexed address)
func (_HoQuConfig *HoQuConfigFilterer) WatchOwnershipRenounced(opts *bind.WatchOpts, sink chan<- *HoQuConfigOwnershipRenounced, previousOwner []common.Address) (event.Subscription, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}

	logs, sub, err := _HoQuConfig.contract.WatchLogs(opts, "OwnershipRenounced", previousOwnerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(HoQuConfigOwnershipRenounced)
				if err := _HoQuConfig.contract.UnpackLog(event, "OwnershipRenounced", log); err != nil {
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

// HoQuConfigOwnershipTransferredIterator is returned from FilterOwnershipTransferred and is used to iterate over the raw logs and unpacked data for OwnershipTransferred events raised by the HoQuConfig contract.
type HoQuConfigOwnershipTransferredIterator struct {
	Event *HoQuConfigOwnershipTransferred // Event containing the contract specifics and raw log

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
func (it *HoQuConfigOwnershipTransferredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(HoQuConfigOwnershipTransferred)
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
		it.Event = new(HoQuConfigOwnershipTransferred)
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
func (it *HoQuConfigOwnershipTransferredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *HoQuConfigOwnershipTransferredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// HoQuConfigOwnershipTransferred represents a OwnershipTransferred event raised by the HoQuConfig contract.
type HoQuConfigOwnershipTransferred struct {
	PreviousOwner common.Address
	NewOwner      common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterOwnershipTransferred is a free log retrieval operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(previousOwner indexed address, newOwner indexed address)
func (_HoQuConfig *HoQuConfigFilterer) FilterOwnershipTransferred(opts *bind.FilterOpts, previousOwner []common.Address, newOwner []common.Address) (*HoQuConfigOwnershipTransferredIterator, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _HoQuConfig.contract.FilterLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return &HoQuConfigOwnershipTransferredIterator{contract: _HoQuConfig.contract, event: "OwnershipTransferred", logs: logs, sub: sub}, nil
}

// WatchOwnershipTransferred is a free log subscription operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(previousOwner indexed address, newOwner indexed address)
func (_HoQuConfig *HoQuConfigFilterer) WatchOwnershipTransferred(opts *bind.WatchOpts, sink chan<- *HoQuConfigOwnershipTransferred, previousOwner []common.Address, newOwner []common.Address) (event.Subscription, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _HoQuConfig.contract.WatchLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(HoQuConfigOwnershipTransferred)
				if err := _HoQuConfig.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
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

// HoQuConfigSystemOwnerAddedIterator is returned from FilterSystemOwnerAdded and is used to iterate over the raw logs and unpacked data for SystemOwnerAdded events raised by the HoQuConfig contract.
type HoQuConfigSystemOwnerAddedIterator struct {
	Event *HoQuConfigSystemOwnerAdded // Event containing the contract specifics and raw log

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
func (it *HoQuConfigSystemOwnerAddedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(HoQuConfigSystemOwnerAdded)
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
		it.Event = new(HoQuConfigSystemOwnerAdded)
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
func (it *HoQuConfigSystemOwnerAddedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *HoQuConfigSystemOwnerAddedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// HoQuConfigSystemOwnerAdded represents a SystemOwnerAdded event raised by the HoQuConfig contract.
type HoQuConfigSystemOwnerAdded struct {
	NewOwner common.Address
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterSystemOwnerAdded is a free log retrieval operation binding the contract event 0xa91384695ab8222d2f6e2437bb5b1fefe671538a33a7e06e88fee0147fffe8a4.
//
// Solidity: event SystemOwnerAdded(newOwner indexed address)
func (_HoQuConfig *HoQuConfigFilterer) FilterSystemOwnerAdded(opts *bind.FilterOpts, newOwner []common.Address) (*HoQuConfigSystemOwnerAddedIterator, error) {

	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _HoQuConfig.contract.FilterLogs(opts, "SystemOwnerAdded", newOwnerRule)
	if err != nil {
		return nil, err
	}
	return &HoQuConfigSystemOwnerAddedIterator{contract: _HoQuConfig.contract, event: "SystemOwnerAdded", logs: logs, sub: sub}, nil
}

// WatchSystemOwnerAdded is a free log subscription operation binding the contract event 0xa91384695ab8222d2f6e2437bb5b1fefe671538a33a7e06e88fee0147fffe8a4.
//
// Solidity: event SystemOwnerAdded(newOwner indexed address)
func (_HoQuConfig *HoQuConfigFilterer) WatchSystemOwnerAdded(opts *bind.WatchOpts, sink chan<- *HoQuConfigSystemOwnerAdded, newOwner []common.Address) (event.Subscription, error) {

	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _HoQuConfig.contract.WatchLogs(opts, "SystemOwnerAdded", newOwnerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(HoQuConfigSystemOwnerAdded)
				if err := _HoQuConfig.contract.UnpackLog(event, "SystemOwnerAdded", log); err != nil {
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

// HoQuConfigSystemOwnerChangedIterator is returned from FilterSystemOwnerChanged and is used to iterate over the raw logs and unpacked data for SystemOwnerChanged events raised by the HoQuConfig contract.
type HoQuConfigSystemOwnerChangedIterator struct {
	Event *HoQuConfigSystemOwnerChanged // Event containing the contract specifics and raw log

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
func (it *HoQuConfigSystemOwnerChangedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(HoQuConfigSystemOwnerChanged)
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
		it.Event = new(HoQuConfigSystemOwnerChanged)
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
func (it *HoQuConfigSystemOwnerChangedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *HoQuConfigSystemOwnerChangedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// HoQuConfigSystemOwnerChanged represents a SystemOwnerChanged event raised by the HoQuConfig contract.
type HoQuConfigSystemOwnerChanged struct {
	PreviousOwner common.Address
	NewOwner      common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterSystemOwnerChanged is a free log retrieval operation binding the contract event 0x1722097e017bb9a189312ac57bbdfff1bf251383cfd36c245f8ebabd8734c759.
//
// Solidity: event SystemOwnerChanged(previousOwner indexed address, newOwner address)
func (_HoQuConfig *HoQuConfigFilterer) FilterSystemOwnerChanged(opts *bind.FilterOpts, previousOwner []common.Address) (*HoQuConfigSystemOwnerChangedIterator, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}

	logs, sub, err := _HoQuConfig.contract.FilterLogs(opts, "SystemOwnerChanged", previousOwnerRule)
	if err != nil {
		return nil, err
	}
	return &HoQuConfigSystemOwnerChangedIterator{contract: _HoQuConfig.contract, event: "SystemOwnerChanged", logs: logs, sub: sub}, nil
}

// WatchSystemOwnerChanged is a free log subscription operation binding the contract event 0x1722097e017bb9a189312ac57bbdfff1bf251383cfd36c245f8ebabd8734c759.
//
// Solidity: event SystemOwnerChanged(previousOwner indexed address, newOwner address)
func (_HoQuConfig *HoQuConfigFilterer) WatchSystemOwnerChanged(opts *bind.WatchOpts, sink chan<- *HoQuConfigSystemOwnerChanged, previousOwner []common.Address) (event.Subscription, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}

	logs, sub, err := _HoQuConfig.contract.WatchLogs(opts, "SystemOwnerChanged", previousOwnerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(HoQuConfigSystemOwnerChanged)
				if err := _HoQuConfig.contract.UnpackLog(event, "SystemOwnerChanged", log); err != nil {
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

// HoQuConfigSystemOwnerDeletedIterator is returned from FilterSystemOwnerDeleted and is used to iterate over the raw logs and unpacked data for SystemOwnerDeleted events raised by the HoQuConfig contract.
type HoQuConfigSystemOwnerDeletedIterator struct {
	Event *HoQuConfigSystemOwnerDeleted // Event containing the contract specifics and raw log

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
func (it *HoQuConfigSystemOwnerDeletedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(HoQuConfigSystemOwnerDeleted)
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
		it.Event = new(HoQuConfigSystemOwnerDeleted)
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
func (it *HoQuConfigSystemOwnerDeletedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *HoQuConfigSystemOwnerDeletedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// HoQuConfigSystemOwnerDeleted represents a SystemOwnerDeleted event raised by the HoQuConfig contract.
type HoQuConfigSystemOwnerDeleted struct {
	DeletedOwner common.Address
	Raw          types.Log // Blockchain specific contextual infos
}

// FilterSystemOwnerDeleted is a free log retrieval operation binding the contract event 0x7ac6b089a9620017a12220abae1696205383eb917822ade81acb4eea2a424f69.
//
// Solidity: event SystemOwnerDeleted(deletedOwner indexed address)
func (_HoQuConfig *HoQuConfigFilterer) FilterSystemOwnerDeleted(opts *bind.FilterOpts, deletedOwner []common.Address) (*HoQuConfigSystemOwnerDeletedIterator, error) {

	var deletedOwnerRule []interface{}
	for _, deletedOwnerItem := range deletedOwner {
		deletedOwnerRule = append(deletedOwnerRule, deletedOwnerItem)
	}

	logs, sub, err := _HoQuConfig.contract.FilterLogs(opts, "SystemOwnerDeleted", deletedOwnerRule)
	if err != nil {
		return nil, err
	}
	return &HoQuConfigSystemOwnerDeletedIterator{contract: _HoQuConfig.contract, event: "SystemOwnerDeleted", logs: logs, sub: sub}, nil
}

// WatchSystemOwnerDeleted is a free log subscription operation binding the contract event 0x7ac6b089a9620017a12220abae1696205383eb917822ade81acb4eea2a424f69.
//
// Solidity: event SystemOwnerDeleted(deletedOwner indexed address)
func (_HoQuConfig *HoQuConfigFilterer) WatchSystemOwnerDeleted(opts *bind.WatchOpts, sink chan<- *HoQuConfigSystemOwnerDeleted, deletedOwner []common.Address) (event.Subscription, error) {

	var deletedOwnerRule []interface{}
	for _, deletedOwnerItem := range deletedOwner {
		deletedOwnerRule = append(deletedOwnerRule, deletedOwnerItem)
	}

	logs, sub, err := _HoQuConfig.contract.WatchLogs(opts, "SystemOwnerDeleted", deletedOwnerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(HoQuConfigSystemOwnerDeleted)
				if err := _HoQuConfig.contract.UnpackLog(event, "SystemOwnerDeleted", log); err != nil {
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

// HoQuRaterIABI is the input ABI used to generate the binding from.
const HoQuRaterIABI = "[{\"constant\":false,\"inputs\":[{\"name\":\"offerId\",\"type\":\"bytes16\"},{\"name\":\"trackerId\",\"type\":\"bytes16\"},{\"name\":\"affiliateId\",\"type\":\"bytes16\"},{\"name\":\"price\",\"type\":\"uint256\"}],\"name\":\"processTransactLead\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"offerId\",\"type\":\"bytes16\"},{\"name\":\"trackerId\",\"type\":\"bytes16\"},{\"name\":\"affiliateId\",\"type\":\"bytes16\"},{\"name\":\"price\",\"type\":\"uint256\"}],\"name\":\"processAddLead\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]"

// HoQuRaterIBin is the compiled bytecode used for deploying new contracts.
const HoQuRaterIBin = `0x`

// DeployHoQuRaterI deploys a new Ethereum contract, binding an instance of HoQuRaterI to it.
func DeployHoQuRaterI(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *HoQuRaterI, error) {
	parsed, err := abi.JSON(strings.NewReader(HoQuRaterIABI))
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	address, tx, contract, err := bind.DeployContract(auth, parsed, common.FromHex(HoQuRaterIBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &HoQuRaterI{HoQuRaterICaller: HoQuRaterICaller{contract: contract}, HoQuRaterITransactor: HoQuRaterITransactor{contract: contract}, HoQuRaterIFilterer: HoQuRaterIFilterer{contract: contract}}, nil
}

// HoQuRaterI is an auto generated Go binding around an Ethereum contract.
type HoQuRaterI struct {
	HoQuRaterICaller     // Read-only binding to the contract
	HoQuRaterITransactor // Write-only binding to the contract
	HoQuRaterIFilterer   // Log filterer for contract events
}

// HoQuRaterICaller is an auto generated read-only Go binding around an Ethereum contract.
type HoQuRaterICaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// HoQuRaterITransactor is an auto generated write-only Go binding around an Ethereum contract.
type HoQuRaterITransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// HoQuRaterIFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type HoQuRaterIFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// HoQuRaterISession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type HoQuRaterISession struct {
	Contract     *HoQuRaterI       // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// HoQuRaterICallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type HoQuRaterICallerSession struct {
	Contract *HoQuRaterICaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts     // Call options to use throughout this session
}

// HoQuRaterITransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type HoQuRaterITransactorSession struct {
	Contract     *HoQuRaterITransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts     // Transaction auth options to use throughout this session
}

// HoQuRaterIRaw is an auto generated low-level Go binding around an Ethereum contract.
type HoQuRaterIRaw struct {
	Contract *HoQuRaterI // Generic contract binding to access the raw methods on
}

// HoQuRaterICallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type HoQuRaterICallerRaw struct {
	Contract *HoQuRaterICaller // Generic read-only contract binding to access the raw methods on
}

// HoQuRaterITransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type HoQuRaterITransactorRaw struct {
	Contract *HoQuRaterITransactor // Generic write-only contract binding to access the raw methods on
}

// NewHoQuRaterI creates a new instance of HoQuRaterI, bound to a specific deployed contract.
func NewHoQuRaterI(address common.Address, backend bind.ContractBackend) (*HoQuRaterI, error) {
	contract, err := bindHoQuRaterI(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &HoQuRaterI{HoQuRaterICaller: HoQuRaterICaller{contract: contract}, HoQuRaterITransactor: HoQuRaterITransactor{contract: contract}, HoQuRaterIFilterer: HoQuRaterIFilterer{contract: contract}}, nil
}

// NewHoQuRaterICaller creates a new read-only instance of HoQuRaterI, bound to a specific deployed contract.
func NewHoQuRaterICaller(address common.Address, caller bind.ContractCaller) (*HoQuRaterICaller, error) {
	contract, err := bindHoQuRaterI(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &HoQuRaterICaller{contract: contract}, nil
}

// NewHoQuRaterITransactor creates a new write-only instance of HoQuRaterI, bound to a specific deployed contract.
func NewHoQuRaterITransactor(address common.Address, transactor bind.ContractTransactor) (*HoQuRaterITransactor, error) {
	contract, err := bindHoQuRaterI(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &HoQuRaterITransactor{contract: contract}, nil
}

// NewHoQuRaterIFilterer creates a new log filterer instance of HoQuRaterI, bound to a specific deployed contract.
func NewHoQuRaterIFilterer(address common.Address, filterer bind.ContractFilterer) (*HoQuRaterIFilterer, error) {
	contract, err := bindHoQuRaterI(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &HoQuRaterIFilterer{contract: contract}, nil
}

// bindHoQuRaterI binds a generic wrapper to an already deployed contract.
func bindHoQuRaterI(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(HoQuRaterIABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_HoQuRaterI *HoQuRaterIRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _HoQuRaterI.Contract.HoQuRaterICaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_HoQuRaterI *HoQuRaterIRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _HoQuRaterI.Contract.HoQuRaterITransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_HoQuRaterI *HoQuRaterIRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _HoQuRaterI.Contract.HoQuRaterITransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_HoQuRaterI *HoQuRaterICallerRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _HoQuRaterI.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_HoQuRaterI *HoQuRaterITransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _HoQuRaterI.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_HoQuRaterI *HoQuRaterITransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _HoQuRaterI.Contract.contract.Transact(opts, method, params...)
}

// ProcessAddLead is a paid mutator transaction binding the contract method 0xc5f6dd85.
//
// Solidity: function processAddLead(offerId bytes16, trackerId bytes16, affiliateId bytes16, price uint256) returns()
func (_HoQuRaterI *HoQuRaterITransactor) ProcessAddLead(opts *bind.TransactOpts, offerId [16]byte, trackerId [16]byte, affiliateId [16]byte, price *big.Int) (*types.Transaction, error) {
	return _HoQuRaterI.contract.Transact(opts, "processAddLead", offerId, trackerId, affiliateId, price)
}

// ProcessAddLead is a paid mutator transaction binding the contract method 0xc5f6dd85.
//
// Solidity: function processAddLead(offerId bytes16, trackerId bytes16, affiliateId bytes16, price uint256) returns()
func (_HoQuRaterI *HoQuRaterISession) ProcessAddLead(offerId [16]byte, trackerId [16]byte, affiliateId [16]byte, price *big.Int) (*types.Transaction, error) {
	return _HoQuRaterI.Contract.ProcessAddLead(&_HoQuRaterI.TransactOpts, offerId, trackerId, affiliateId, price)
}

// ProcessAddLead is a paid mutator transaction binding the contract method 0xc5f6dd85.
//
// Solidity: function processAddLead(offerId bytes16, trackerId bytes16, affiliateId bytes16, price uint256) returns()
func (_HoQuRaterI *HoQuRaterITransactorSession) ProcessAddLead(offerId [16]byte, trackerId [16]byte, affiliateId [16]byte, price *big.Int) (*types.Transaction, error) {
	return _HoQuRaterI.Contract.ProcessAddLead(&_HoQuRaterI.TransactOpts, offerId, trackerId, affiliateId, price)
}

// ProcessTransactLead is a paid mutator transaction binding the contract method 0x4943df38.
//
// Solidity: function processTransactLead(offerId bytes16, trackerId bytes16, affiliateId bytes16, price uint256) returns()
func (_HoQuRaterI *HoQuRaterITransactor) ProcessTransactLead(opts *bind.TransactOpts, offerId [16]byte, trackerId [16]byte, affiliateId [16]byte, price *big.Int) (*types.Transaction, error) {
	return _HoQuRaterI.contract.Transact(opts, "processTransactLead", offerId, trackerId, affiliateId, price)
}

// ProcessTransactLead is a paid mutator transaction binding the contract method 0x4943df38.
//
// Solidity: function processTransactLead(offerId bytes16, trackerId bytes16, affiliateId bytes16, price uint256) returns()
func (_HoQuRaterI *HoQuRaterISession) ProcessTransactLead(offerId [16]byte, trackerId [16]byte, affiliateId [16]byte, price *big.Int) (*types.Transaction, error) {
	return _HoQuRaterI.Contract.ProcessTransactLead(&_HoQuRaterI.TransactOpts, offerId, trackerId, affiliateId, price)
}

// ProcessTransactLead is a paid mutator transaction binding the contract method 0x4943df38.
//
// Solidity: function processTransactLead(offerId bytes16, trackerId bytes16, affiliateId bytes16, price uint256) returns()
func (_HoQuRaterI *HoQuRaterITransactorSession) ProcessTransactLead(offerId [16]byte, trackerId [16]byte, affiliateId [16]byte, price *big.Int) (*types.Transaction, error) {
	return _HoQuRaterI.Contract.ProcessTransactLead(&_HoQuRaterI.TransactOpts, offerId, trackerId, affiliateId, price)
}

// HoQuStorageABI is the input ABI used to generate the binding from.
const HoQuStorageABI = "[{\"constant\":false,\"inputs\":[{\"name\":\"id\",\"type\":\"bytes16\"},{\"name\":\"ownerId\",\"type\":\"bytes16\"},{\"name\":\"name\",\"type\":\"string\"},{\"name\":\"dataUrl\",\"type\":\"string\"},{\"name\":\"status\",\"type\":\"uint8\"}],\"name\":\"setNetwork\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"id\",\"type\":\"bytes16\"},{\"name\":\"ownerAddress\",\"type\":\"address\"}],\"name\":\"addUserAddress\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"id\",\"type\":\"bytes16\"},{\"name\":\"meta\",\"type\":\"string\"},{\"name\":\"kycLevel\",\"type\":\"uint8\"},{\"name\":\"dataUrl\",\"type\":\"string\"}],\"name\":\"addKycReport\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"id\",\"type\":\"bytes16\"},{\"name\":\"role\",\"type\":\"string\"},{\"name\":\"ownerAddress\",\"type\":\"address\"},{\"name\":\"kycLevel\",\"type\":\"uint8\"},{\"name\":\"pubKey\",\"type\":\"string\"},{\"name\":\"status\",\"type\":\"uint8\"}],\"name\":\"setUser\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"\",\"type\":\"bytes16\"}],\"name\":\"ids\",\"outputs\":[{\"name\":\"userId\",\"type\":\"bytes16\"},{\"name\":\"companyId\",\"type\":\"bytes16\"},{\"name\":\"idType\",\"type\":\"string\"},{\"name\":\"name\",\"type\":\"string\"},{\"name\":\"numOfKycReports\",\"type\":\"uint16\"},{\"name\":\"createdAt\",\"type\":\"uint256\"},{\"name\":\"status\",\"type\":\"uint8\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"\",\"type\":\"bytes16\"}],\"name\":\"companies\",\"outputs\":[{\"name\":\"ownerId\",\"type\":\"bytes16\"},{\"name\":\"name\",\"type\":\"string\"},{\"name\":\"dataUrl\",\"type\":\"string\"},{\"name\":\"createdAt\",\"type\":\"uint256\"},{\"name\":\"status\",\"type\":\"uint8\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"\",\"type\":\"bytes16\"}],\"name\":\"networks\",\"outputs\":[{\"name\":\"ownerId\",\"type\":\"bytes16\"},{\"name\":\"name\",\"type\":\"string\"},{\"name\":\"dataUrl\",\"type\":\"string\"},{\"name\":\"createdAt\",\"type\":\"uint256\"},{\"name\":\"status\",\"type\":\"uint8\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"\",\"type\":\"bytes16\"}],\"name\":\"users\",\"outputs\":[{\"name\":\"numOfAddresses\",\"type\":\"uint8\"},{\"name\":\"role\",\"type\":\"string\"},{\"name\":\"kycLevel\",\"type\":\"uint8\"},{\"name\":\"pubKey\",\"type\":\"string\"},{\"name\":\"createdAt\",\"type\":\"uint256\"},{\"name\":\"status\",\"type\":\"uint8\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"id\",\"type\":\"bytes16\"},{\"name\":\"ownerId\",\"type\":\"bytes16\"},{\"name\":\"name\",\"type\":\"string\"},{\"name\":\"dataUrl\",\"type\":\"string\"},{\"name\":\"status\",\"type\":\"uint8\"}],\"name\":\"setCompany\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"config\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"configAddress\",\"type\":\"address\"}],\"name\":\"setConfigAddress\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"\",\"type\":\"bytes16\"}],\"name\":\"adCampaigns\",\"outputs\":[{\"name\":\"ownerId\",\"type\":\"bytes16\"},{\"name\":\"offerId\",\"type\":\"bytes16\"},{\"name\":\"contractAddress\",\"type\":\"address\"},{\"name\":\"createdAt\",\"type\":\"uint256\"},{\"name\":\"status\",\"type\":\"uint8\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"id\",\"type\":\"bytes16\"},{\"name\":\"ownerId\",\"type\":\"bytes16\"},{\"name\":\"offerId\",\"type\":\"bytes16\"},{\"name\":\"contractAddress\",\"type\":\"address\"},{\"name\":\"status\",\"type\":\"uint8\"}],\"name\":\"setAdCampaign\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"\",\"type\":\"bytes16\"}],\"name\":\"stats\",\"outputs\":[{\"name\":\"rating\",\"type\":\"uint256\"},{\"name\":\"volume\",\"type\":\"uint256\"},{\"name\":\"contragents\",\"type\":\"uint256\"},{\"name\":\"stat1\",\"type\":\"uint256\"},{\"name\":\"stat2\",\"type\":\"uint256\"},{\"name\":\"status\",\"type\":\"uint8\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"id\",\"type\":\"bytes16\"},{\"name\":\"userId\",\"type\":\"bytes16\"},{\"name\":\"idType\",\"type\":\"string\"},{\"name\":\"name\",\"type\":\"string\"},{\"name\":\"companyId\",\"type\":\"bytes16\"},{\"name\":\"status\",\"type\":\"uint8\"}],\"name\":\"setIdentification\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"id\",\"type\":\"bytes16\"},{\"name\":\"userId\",\"type\":\"bytes16\"},{\"name\":\"rating\",\"type\":\"uint256\"},{\"name\":\"volume\",\"type\":\"uint256\"},{\"name\":\"contragents\",\"type\":\"uint256\"},{\"name\":\"stat1\",\"type\":\"uint256\"},{\"name\":\"stat2\",\"type\":\"uint256\"},{\"name\":\"status\",\"type\":\"uint8\"}],\"name\":\"setStats\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"\",\"type\":\"bytes16\"}],\"name\":\"offers\",\"outputs\":[{\"name\":\"ownerId\",\"type\":\"bytes16\"},{\"name\":\"networkId\",\"type\":\"bytes16\"},{\"name\":\"merchantId\",\"type\":\"bytes16\"},{\"name\":\"payerAddress\",\"type\":\"address\"},{\"name\":\"name\",\"type\":\"string\"},{\"name\":\"dataUrl\",\"type\":\"string\"},{\"name\":\"cost\",\"type\":\"uint256\"},{\"name\":\"createdAt\",\"type\":\"uint256\"},{\"name\":\"status\",\"type\":\"uint8\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"id\",\"type\":\"bytes16\"},{\"name\":\"num\",\"type\":\"uint8\"}],\"name\":\"getUserAddress\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"id\",\"type\":\"bytes16\"},{\"name\":\"ownerId\",\"type\":\"bytes16\"},{\"name\":\"networkId\",\"type\":\"bytes16\"},{\"name\":\"name\",\"type\":\"string\"},{\"name\":\"dataUrl\",\"type\":\"string\"},{\"name\":\"status\",\"type\":\"uint8\"}],\"name\":\"setTracker\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"id\",\"type\":\"bytes16\"},{\"name\":\"num\",\"type\":\"uint16\"}],\"name\":\"getKycReport\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"},{\"name\":\"\",\"type\":\"string\"},{\"name\":\"\",\"type\":\"uint8\"},{\"name\":\"\",\"type\":\"string\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"\",\"type\":\"bytes16\"}],\"name\":\"trackers\",\"outputs\":[{\"name\":\"ownerId\",\"type\":\"bytes16\"},{\"name\":\"networkId\",\"type\":\"bytes16\"},{\"name\":\"name\",\"type\":\"string\"},{\"name\":\"dataUrl\",\"type\":\"string\"},{\"name\":\"createdAt\",\"type\":\"uint256\"},{\"name\":\"status\",\"type\":\"uint8\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"id\",\"type\":\"bytes16\"},{\"name\":\"ownerId\",\"type\":\"bytes16\"},{\"name\":\"networkId\",\"type\":\"bytes16\"},{\"name\":\"merchantId\",\"type\":\"bytes16\"},{\"name\":\"payerAddress\",\"type\":\"address\"},{\"name\":\"name\",\"type\":\"string\"},{\"name\":\"dataUrl\",\"type\":\"string\"},{\"name\":\"cost\",\"type\":\"uint256\"},{\"name\":\"status\",\"type\":\"uint8\"}],\"name\":\"setOffer\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"name\":\"configAddress\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"ownerAddress\",\"type\":\"address\"},{\"indexed\":false,\"name\":\"id\",\"type\":\"bytes16\"},{\"indexed\":false,\"name\":\"role\",\"type\":\"string\"}],\"name\":\"UserRegistered\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"ownerAddress\",\"type\":\"address\"},{\"indexed\":false,\"name\":\"additionalAddress\",\"type\":\"address\"}],\"name\":\"UserAddressAdded\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"ownerAddress\",\"type\":\"address\"},{\"indexed\":false,\"name\":\"id\",\"type\":\"bytes16\"},{\"indexed\":false,\"name\":\"rating\",\"type\":\"uint256\"}],\"name\":\"StatsChanged\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"ownerAddress\",\"type\":\"address\"},{\"indexed\":false,\"name\":\"id\",\"type\":\"bytes16\"},{\"indexed\":false,\"name\":\"name\",\"type\":\"string\"}],\"name\":\"IdentificationAdded\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"ownerAddress\",\"type\":\"address\"},{\"indexed\":false,\"name\":\"kycLevel\",\"type\":\"uint8\"}],\"name\":\"KycReportAdded\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"ownerAddress\",\"type\":\"address\"},{\"indexed\":false,\"name\":\"id\",\"type\":\"bytes16\"},{\"indexed\":false,\"name\":\"name\",\"type\":\"string\"}],\"name\":\"CompanyRegistered\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"ownerAddress\",\"type\":\"address\"},{\"indexed\":false,\"name\":\"id\",\"type\":\"bytes16\"},{\"indexed\":false,\"name\":\"name\",\"type\":\"string\"}],\"name\":\"NetworkRegistered\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"ownerAddress\",\"type\":\"address\"},{\"indexed\":false,\"name\":\"id\",\"type\":\"bytes16\"},{\"indexed\":false,\"name\":\"name\",\"type\":\"string\"}],\"name\":\"TrackerRegistered\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"ownerAddress\",\"type\":\"address\"},{\"indexed\":false,\"name\":\"id\",\"type\":\"bytes16\"},{\"indexed\":false,\"name\":\"name\",\"type\":\"string\"}],\"name\":\"OfferAdded\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"ownerAddress\",\"type\":\"address\"},{\"indexed\":false,\"name\":\"id\",\"type\":\"bytes16\"},{\"indexed\":false,\"name\":\"contractAddress\",\"type\":\"address\"}],\"name\":\"AdCampaignAdded\",\"type\":\"event\"}]"

// HoQuStorageBin is the compiled bytecode used for deploying new contracts.
const HoQuStorageBin = `0x608060405234801561001057600080fd5b50604051602080613c38833981016040525160008054600160a060020a03909216600160a060020a0319909216919091179055613be6806100526000396000f3006080604052600436106101275763ffffffff7c010000000000000000000000000000000000000000000000000000000060003504166302884980811461012c57806303806028146101e157806318c8f3051461020f578063253a9b9e146102be578063289e7470146103825780634df2d847146104c457806372c130e4146105f35780637520dd1414610615578063768d80a81461075357806379502c551461080657806383a12de9146108375780638907555714610858578063a0191d20146108d3578063a2affa4d14610913578063acc02cce1461097c578063b1dbfe5014610a3c578063c32d869b14610a79578063c49e2fb114610bd5578063c4d4d88714610bfd578063cc3436b914610cb3578063ea2fdd8c14610dd6578063f3b412f114610e6b575b600080fd5b34801561013857600080fd5b50604080516020600460443581810135601f81018490048402850184019095528484526101df9482356001608060020a03199081169560248035909216953695946064949293019190819084018382808284375050604080516020601f89358b018035918201839004830284018301909452808352979a9998810197919650918201945092508291508401838280828437509497505050923560ff169350610f3d92505050565b005b3480156101ed57600080fd5b506101df6001608060020a031960043516600160a060020a03602435166112d2565b34801561021b57600080fd5b5060408051602060046024803582810135601f81018590048502860185019096528585526101df9583356001608060020a03191695369560449491939091019190819084018382808284375050604080516020601f818a01358b0180359182018390048302840183018552818452989b60ff8b35169b909a9099940197509195509182019350915081908401838280828437509497506114369650505050505050565b3480156102ca57600080fd5b5060408051602060046024803582810135601f81018590048502860185019096528585526101df9583356001608060020a03191695369560449491939091019190819084018382808284375050604080516020888301358a018035601f8101839004830284018301909452838352979a8935600160a060020a03169a8a83013560ff169a919990985060609091019650919450908101925081908401838280828437509497505050923560ff16935061166592505050565b34801561038e57600080fd5b506103a46001608060020a031960043516611a0d565b604080516001608060020a0319808a1682528816602082015261ffff8516608082015260a081018490529081016060820160c083018460058111156103e557fe5b60ff168152602001838103835288818151815260200191508051906020019080838360005b8381101561042257818101518382015260200161040a565b50505050905090810190601f16801561044f5780820380516001836020036101000a031916815260200191505b50838103825287518152875160209182019189019080838360005b8381101561048257818101518382015260200161046a565b50505050905090810190601f1680156104af5780820380516001836020036101000a031916815260200191505b50995050505050505050505060405180910390f35b3480156104d057600080fd5b506104e66001608060020a031960043516611b6e565b604080516001608060020a0319871681526060810184905290602082019082016080830184600581111561051657fe5b60ff168152602001838103835287818151815260200191508051906020019080838360005b8381101561055357818101518382015260200161053b565b50505050905090810190601f1680156105805780820380516001836020036101000a031916815260200191505b50838103825286518152865160209182019188019080838360005b838110156105b357818101518382015260200161059b565b50505050905090810190601f1680156105e05780820380516001836020036101000a031916815260200191505b5097505050505050505060405180910390f35b3480156105ff57600080fd5b506104e66001608060020a031960043516611cc1565b34801561062157600080fd5b506106376001608060020a031960043516611d3d565b604051808760ff1660ff1681526020018060200186600581111561065757fe5b60ff1681526020018060200185815260200184600581111561067557fe5b60ff168152602001838103835288818151815260200191508051906020019080838360005b838110156106b257818101518382015260200161069a565b50505050905090810190601f1680156106df5780820380516001836020036101000a031916815260200191505b50838103825286518152865160209182019188019080838360005b838110156107125781810151838201526020016106fa565b50505050905090810190601f16801561073f5780820380516001836020036101000a031916815260200191505b509850505050505050505060405180910390f35b34801561075f57600080fd5b50604080516020600460443581810135601f81018490048402850184019095528484526101df9482356001608060020a03199081169560248035909216953695946064949293019190819084018382808284375050604080516020601f89358b018035918201839004830284018301909452808352979a9998810197919650918201945092508291508401838280828437509497505050923560ff169350611e9592505050565b34801561081257600080fd5b5061081b6121d7565b60408051600160a060020a039092168252519081900360200190f35b34801561084357600080fd5b506101df600160a060020a03600435166121e6565b34801561086457600080fd5b5061087a6001608060020a031960043516612294565b604080516001608060020a0319808816825286166020820152600160a060020a0385169181019190915260608101839052608081018260058111156108bb57fe5b60ff1681526020019550505050505060405180910390f35b3480156108df57600080fd5b506101df6001608060020a031960043581169060243581169060443516600160a060020a036064351660ff608435166122d6565b34801561091f57600080fd5b506109356001608060020a031960043516612606565b6040518087815260200186815260200185815260200184815260200183815260200182600581111561096357fe5b60ff168152602001965050505050505060405180910390f35b34801561098857600080fd5b50604080516020600460443581810135601f81018490048402850184019095528484526101df9482356001608060020a03199081169560248035909216953695946064949293019190819084018382808284375050604080516020601f89358b018035918201839004830284018301909452808352979a999881019791965091820194509250829150840183828082843750949750505083356001608060020a03191694505050506020013560ff16612641565b348015610a4857600080fd5b506101df6001608060020a03196004358116906024351660443560643560843560a43560c43560ff60e4351661299e565b348015610a8557600080fd5b50610a9b6001608060020a031960043516612d54565b604080516001608060020a0319808c1682528a81166020830152891691810191909152600160a060020a038716606082015260c0810184905260e081018390526080810160a082016101008301846005811115610af457fe5b60ff168152602001838103835288818151815260200191508051906020019080838360005b83811015610b31578181015183820152602001610b19565b50505050905090810190601f168015610b5e5780820380516001836020036101000a031916815260200191505b50838103825287518152875160209182019189019080838360005b83811015610b91578181015183820152602001610b79565b50505050905090810190601f168015610bbe5780820380516001836020036101000a031916815260200191505b509b50505050505050505050505060405180910390f35b348015610be157600080fd5b5061081b6001608060020a03196004351660ff60243516612ec7565b348015610c0957600080fd5b50604080516020601f6064356004818101359283018490048402850184019095528184526101df946001608060020a0319813581169560248035831696604435909316953695608494920191819084018382808284375050604080516020601f89358b018035918201839004830284018301909452808352979a9998810197919650918201945092508291508401838280828437509497505050923560ff169350612f3892505050565b348015610cbf57600080fd5b50610cdc6001608060020a03196004351661ffff602435166132ff565b6040518085815260200180602001846005811115610cf657fe5b60ff16815260200180602001838103835286818151815260200191508051906020019080838360005b83811015610d37578181015183820152602001610d1f565b50505050905090810190601f168015610d645780820380516001836020036101000a031916815260200191505b50838103825284518152845160209182019186019080838360005b83811015610d97578181015183820152602001610d7f565b50505050905090810190601f168015610dc45780820380516001836020036101000a031916815260200191505b50965050505050505060405180910390f35b348015610de257600080fd5b50610df86001608060020a0319600435166134a7565b604080516001608060020a0319808916825287166020820152608081018490529081016060820160a08301846005811115610e2f57fe5b60ff16815260200183810383528781815181526020019150805190602001908083836000838110156106b257818101518382015260200161069a565b348015610e7757600080fd5b50604080516020600460a43581810135601f81018490048402850184019095528484526101df9482356001608060020a0319908116956024803583169660443584169660643590941695608435600160a060020a03169536959460c4949093920191819084018382808284375050604080516020601f89358b018035918201839004830284018301909452808352979a999881019791965091820194509250829150840183828082843750949750508435955050505060209091013560ff1690506135ff565b600080546040805160e060020a63babcc539028152600160a060020a0333811660048301529151919092169263babcc53992602480820193602093909283900390910190829087803b158015610f9257600080fd5b505af1158015610fa6573d6000803e3d6000fd5b505050506040513d6020811015610fbc57600080fd5b50511515610fc957600080fd5b6001608060020a0319841660009081526001602052604081206006015460ff166005811115610ff457fe5b1415610fff57600080fd5b6001608060020a031984166000908152600160209081526040808320838052909152902054600160a060020a0316151561103857600080fd5b6001608060020a03198516600090815260056020819052604082206004015460ff169081111561106457fe5b1415611208576040805160a0810182526001608060020a0319861681526020810185905290810183905242606082015260808101600190526001608060020a0319808716600090815260056020908152604090912083518154608060020a90910493169290921782558281015180516110e39260018501920190613b1f565b50604082015180516110ff916002840191602090910190613b1f565b5060608201516003820155608082015160048201805460ff1916600183600581111561112757fe5b021790555050506001608060020a031980851660009081526001602090815260408083208380528252808320548151948a1685528483018281528851928601929092528751600160a060020a03909116947f592e83ede8c846dffd6e4edcbe77983b153bec820ac697aa34a236147b1341e0948b948a9492936060850192918601918190849084905b838110156111c85781810151838201526020016111b0565b50505050905090810190601f1680156111f55780820380516001836020036101000a031916815260200191505b50935050505060405180910390a26112cb565b825115611241576001608060020a031985166000908152600560209081526040909120845161123f92600190920191860190613b1f565b505b81511561127a576001608060020a031985166000908152600560209081526040909120835161127892600290920191850190613b1f565b505b600081600581111561128857fe5b146112cb576001608060020a0319851660009081526005602081905260409091206004018054839260ff199091169060019084908111156112c557fe5b02179055505b5050505050565b600080546040805160e060020a63babcc539028152600160a060020a0333811660048301529151919092169263babcc53992602480820193602093909283900390910190829087803b15801561132757600080fd5b505af115801561133b573d6000803e3d6000fd5b505050506040513d602081101561135157600080fd5b5051151561135e57600080fd5b6001608060020a0319821660009081526001602052604081206006015460ff16600581111561138957fe5b141561139457600080fd5b6001608060020a031982166000908152600160208181526040808420808401805460ff90811687529184528286208054600160a060020a031916600160a060020a03898116918217909255825460ff1981169085169097019093169590951790558480529381902054815194855290519216927fb8a06c596f027284f088ce1ba17842629ca63619a7303c94b3645e310ae34e44929081900390910190a25050565b600080546040805160e060020a63babcc539028152600160a060020a0333811660048301529151919092169263babcc53992602480820193602093909283900390910190829087803b15801561148b57600080fd5b505af115801561149f573d6000803e3d6000fd5b505050506040513d60208110156114b557600080fd5b505115156114c257600080fd5b6001608060020a0319841660009081526002602052604081206006015460ff1660058111156114ed57fe5b14156114f857600080fd5b60806040519081016040528084815260200183600581111561151657fe5b81526020808201849052426040928301526001608060020a03198716600090815260028252828120600481015461ffff16825260030182529190912082518051919261156792849290910190613b1f565b50602082015160018083018054909160ff199091169083600581111561158957fe5b0217905550604082015180516115a9916002840191602090910190613b1f565b50606091909101516003909101556001608060020a0319848116600090815260026020908152604080832060048101805461ffff198116600161ffff9283168101909216179091559054608060020a0290941683529281528282208280529052819020549051600160a060020a03909116907f92262b4b81e23d81eaec5e9e8c9439f0a59929da8c5d49b32bca92c112f4172a9084908082600581111561164c57fe5b60ff16815260200191505060405180910390a250505050565b600080546040805160e060020a63babcc539028152600160a060020a0333811660048301529151919092169263babcc53992602480820193602093909283900390910190829087803b1580156116ba57600080fd5b505af11580156116ce573d6000803e3d6000fd5b505050506040513d60208110156116e457600080fd5b505115156116f157600080fd5b6001608060020a0319861660009081526001602052604081206006015460ff16600581111561171c57fe5b14156118f5576040805160c081018252600180825260208083018981528385018390526060840187905242608085015260a084018390526001608060020a03198b166000908152838352949094208351928101805460ff191660ff9094169390931790925592518051929391926117999260028501920190613b1f565b50604082015160038201805460ff191660018360058111156117b757fe5b0217905550606082015180516117d7916004840191602090910190613b1f565b506080820151816005015560a08201518160060160006101000a81548160ff0219169083600581111561180657fe5b021790555050506001608060020a03198616600081815260016020908152604080832083805282528083208054600160a060020a031916600160a060020a038a1690811790915581519485528483018281528a5192860192909252895190947ff661ee1472892faaf2bb68cc6874f2759c9148b4234cac6b88f6ca362652f275948c948c9492936060850192918601918190849084905b838110156118b557818101518382015260200161189d565b50505050905090810190601f1680156118e25780820380516001836020036101000a031916815260200191505b50935050505060405180910390a2611a05565b84511561192e576001608060020a031986166000908152600160209081526040909120865161192c92600290920191880190613b1f565b505b600083600581111561193c57fe5b1461197d576001608060020a0319861660009081526001602081905260409091206003018054859260ff199091169083600581111561197757fe5b02179055505b8151156119b6576001608060020a03198616600090815260016020908152604090912083516119b492600490920191850190613b1f565b505b60008160058111156119c457fe5b14611a05576001608060020a0319861660009081526001602081905260409091206006018054839260ff19909116908360058111156119ff57fe5b02179055505b505050505050565b6002602081815260009283526040928390208054600180830180548751601f60001994831615610100029490940190911696909604918201859004850286018501909652808552608060020a808302969281900402949293830182828015611ab65780601f10611a8b57610100808354040283529160200191611ab6565b820191906000526020600020905b815481529060010190602001808311611a9957829003601f168201915b50505060028085018054604080516020601f6000196101006001871615020190941695909504928301859004850281018501909152818152959695945090925090830182828015611b485780601f10611b1d57610100808354040283529160200191611b48565b820191906000526020600020905b815481529060010190602001808311611b2b57829003601f168201915b5050505060048301546005840154600690940154929361ffff9091169290915060ff1687565b6004602090815260009182526040918290208054600180830180548651600261010094831615949094026000190190911692909204601f8101869004860283018601909652858252608060020a909202949293909290830182828015611c155780601f10611bea57610100808354040283529160200191611c15565b820191906000526020600020905b815481529060010190602001808311611bf857829003601f168201915b50505060028085018054604080516020601f6000196101006001871615020190941695909504928301859004850281018501909152818152959695945090925090830182828015611ca75780601f10611c7c57610100808354040283529160200191611ca7565b820191906000526020600020905b815481529060010190602001808311611c8a57829003601f168201915b50505050600383015460049093015491929160ff16905085565b6005602090815260009182526040918290208054600180830180548651600261010094831615949094026000190190911692909204601f8101869004860283018601909652858252608060020a909202949293909290830182828015611c155780601f10611bea57610100808354040283529160200191611c15565b600160208181526000928352604092839020808301546002808301805487516101009782161597909702600019011691909104601f810185900485028601850190965285855260ff90911694919392909190830182828015611de05780601f10611db557610100808354040283529160200191611de0565b820191906000526020600020905b815481529060010190602001808311611dc357829003601f168201915b50505050600383015460048401805460408051602060026101006001861615026000190190941693909304601f8101849004840282018401909252818152959660ff9094169593945090830182828015611e7b5780601f10611e5057610100808354040283529160200191611e7b565b820191906000526020600020905b815481529060010190602001808311611e5e57829003601f168201915b50505050600583015460069093015491929160ff16905086565b600080546040805160e060020a63babcc539028152600160a060020a0333811660048301529151919092169263babcc53992602480820193602093909283900390910190829087803b158015611eea57600080fd5b505af1158015611efe573d6000803e3d6000fd5b505050506040513d6020811015611f1457600080fd5b50511515611f2157600080fd5b6001608060020a0319841660009081526001602052604081206006015460ff166005811115611f4c57fe5b1415611f5757600080fd5b6001608060020a031984166000908152600160209081526040808320838052909152902054600160a060020a03161515611f9057600080fd5b6001608060020a0319851660009081526004602081905260408220015460ff166005811115611fbb57fe5b141561211e576040805160a0810182526001608060020a0319861681526020810185905290810183905242606082015260808101600190526001608060020a0319808716600090815260046020908152604090912083518154608060020a909104931692909217825582810151805161203a9260018501920190613b1f565b5060408201518051612056916002840191602090910190613b1f565b5060608201516003820155608082015160048201805460ff1916600183600581111561207e57fe5b021790555050506001608060020a031980851660009081526001602090815260408083208380528252808320548151948a1685528483018281528851928601929092528751600160a060020a03909116947f0eb23d662c57754d23afdc1cffa6246fbae3bbe843eec5c42a5a5223c21414ed948b948a949293606085019291860191819084908490838110156111c85781810151838201526020016111b0565b825115612157576001608060020a031985166000908152600460209081526040909120845161215592600190920191860190613b1f565b505b815115612190576001608060020a031985166000908152600460209081526040909120835161218e92600290920191850190613b1f565b505b600081600581111561219e57fe5b146112cb576001608060020a03198516600090815260046020819052604090912001805482919060ff191660018360058111156112c557fe5b600054600160a060020a031681565b600080546040805160e060020a63babcc539028152600160a060020a0333811660048301529151919092169263babcc53992602480820193602093909283900390910190829087803b15801561223b57600080fd5b505af115801561224f573d6000803e3d6000fd5b505050506040513d602081101561226557600080fd5b5051151561227257600080fd5b60008054600160a060020a031916600160a060020a0392909216919091179055565b6008602052600090815260409020805460018201546002830154600390930154608060020a80840294938190040292600160a060020a03909216919060ff1685565b600080546040805160e060020a63babcc539028152600160a060020a0333811660048301529151919092169163babcc53991602480830192602092919082900301818787803b15801561232857600080fd5b505af115801561233c573d6000803e3d6000fd5b505050506040513d602081101561235257600080fd5b5051151561235f57600080fd5b6001608060020a0319851660009081526001602052604081206006015460ff16600581111561238a57fe5b141561239557600080fd5b6001608060020a031985166000908152600160209081526040808320838052909152902054600160a060020a031615156123ce57600080fd5b6001608060020a0319841660009081526006602052604081206007015460ff1660058111156123f957fe5b141561240457600080fd5b6001608060020a0319861660009081526007602052604081206004015460ff16600581111561242f57fe5b141561257b57612440856000612ec7565b6040805160a0810182526001608060020a0319808916825287166020820152600160a060020a0386169181019190915242606082015290915060808101600190526001608060020a03198088166000908152600860209081526040918290208451815492860151608060020a908190048102910492909416919091176001608060020a031692909217825582015160018083018054600160a060020a03909316600160a060020a0319909316929092179091556060830151600283015560808301516003830180549192909160ff19169083600581111561251d57fe5b021790555050604080516001608060020a031989168152600160a060020a038681166020830152825190851693507f62a609593cded926aabb6aef1e5b884804ad29d1bd1116ac4ce32085690d5ab2929181900390910190a2611a05565b600160a060020a038316156125c0576001608060020a0319861660009081526008602052604090206001018054600160a060020a031916600160a060020a0385161790555b60008260058111156125ce57fe5b14611a05576001608060020a031986166000908152600860205260409020600301805483919060ff191660018360058111156119ff57fe5b600360208190526000918252604090912080546001820154600283015493830154600484015460059094015492949193919290919060ff1686565b600080546040805160e060020a63babcc539028152600160a060020a0333811660048301529151919092169163babcc53991602480830192602092919082900301818787803b15801561269357600080fd5b505af11580156126a7573d6000803e3d6000fd5b505050506040513d60208110156126bd57600080fd5b505115156126ca57600080fd5b6001608060020a0319871660009081526002602052604081206006015460ff1660058111156126f557fe5b14156128e257612706866000612ec7565b6040805160e0810182526001608060020a0319808a1682528616602082015290810187905260608101869052600060808201524260a082015290915060c08101600190526001608060020a03198089166000908152600260209081526040918290208451815486840151608060020a90819004810292049516949094176001608060020a03169390931783559083015180516127a89260018501920190613b1f565b50606082015180516127c4916002840191602090910190613b1f565b5060808201518160040160006101000a81548161ffff021916908361ffff16021790555060a0820151816005015560c08201518160060160006101000a81548160ff0219169083600581111561281657fe5b021790555090505080600160a060020a03167fa5c73c5ece99b6456d1b224c43e2a1fc268c1e3a14156f135f3e8bdec2ff9a5a888660405180836001608060020a0319166001608060020a031916815260200180602001828103825283818151815260200191508051906020019080838360005b838110156128a257818101518382015260200161288a565b50505050905090810190601f1680156128cf5780820380516001836020036101000a031916815260200191505b50935050505060405180910390a2612995565b84511561291b576001608060020a031987166000908152600260209081526040909120865161291992600190920191880190613b1f565b505b6001608060020a03198716600090815260026020526040812080546001608060020a0316608060020a8087040217905582600581111561295757fe5b14612995576001608060020a031987166000908152600260205260409020600601805483919060ff1916600183600581111561298f57fe5b02179055505b50505050505050565b600080546040805160e060020a63babcc539028152600160a060020a0333811660048301529151919092169163babcc53991602480830192602092919082900301818787803b1580156129f057600080fd5b505af1158015612a04573d6000803e3d6000fd5b505050506040513d6020811015612a1a57600080fd5b50511515612a2757600080fd5b6001608060020a03198916600090815260036020526040812060059081015460ff1690811115612a5357fe5b1415612bf757612a64886000612ec7565b905060c06040519081016040528088815260200187815260200186815260200185815260200184815260200160016005811115612a9d57fe5b815250600360008b6001608060020a0319166001608060020a0319168152602001908152602001600020600082015181600001556020820151816001015560408201518160020155606082015181600301556080820151816004015560a08201518160050160006101000a81548160ff02191690836005811115612b1d57fe5b021790555050506001608060020a0319808a166000908152600360208190526040808320938c1683529091208254815560018084015481830155600280850154908301558383015492820192909255600480840154908201556005808401548183018054939460ff90921693909260ff1990911691908490811115612b9e57fe5b02179055505050604080516001608060020a03198b168152602081018990528151600160a060020a038416927fd2e9a0e1a68c0320c3b58608fcefa40787236bd0d2153ebbd82690b3fa064ad4928290030190a2612d49565b8615612c25576001608060020a03198981166000908152600360205260408082208a9055918a168152208790555b8515612c5a576001608060020a031989811660009081526003602052604080822060019081018a9055928b1682529020018690555b8415612c8f576001608060020a03198981166000908152600360205260408082206002908101899055928b1682529020018590555b8315612cc4576001608060020a031989811660009081526003602081905260408083208201889055928b168252919020018490555b8215612cf9576001608060020a03198981166000908152600360205260408082206004908101879055928b1682529020018390555b6000826005811115612d0757fe5b14612d49576001608060020a03198916600090815260036020526040902060059081018054849260ff19909116906001908490811115612d4357fe5b02179055505b505050505050505050565b600660209081526000918252604091829020805460018083015460028085015460038601805489516101009682161596909602600019011692909204601f8101889004880285018801909852878452608060020a8086029895819004810297930295600160a060020a0390911694909392830182828015612e165780601f10612deb57610100808354040283529160200191612e16565b820191906000526020600020905b815481529060010190602001808311612df957829003601f168201915b5050505060048301805460408051602060026001851615610100026000190190941693909304601f8101849004840282018401909252818152949594935090830182828015612ea65780601f10612e7b57610100808354040283529160200191612ea6565b820191906000526020600020905b815481529060010190602001808311612e8957829003601f168201915b50505050600583015460068401546007909401549293909290915060ff1689565b6000806001608060020a0319841660009081526001602052604090206006015460ff166005811115612ef557fe5b1415612f0057600080fd5b506001608060020a03198216600090815260016020908152604080832060ff85168452909152902054600160a060020a031692915050565b600080546040805160e060020a63babcc539028152600160a060020a0333811660048301529151919092169263babcc53992602480820193602093909283900390910190829087803b158015612f8d57600080fd5b505af1158015612fa1573d6000803e3d6000fd5b505050506040513d6020811015612fb757600080fd5b50511515612fc457600080fd5b6001608060020a0319851660009081526001602052604081206006015460ff166005811115612fef57fe5b1415612ffa57600080fd5b6001608060020a031985166000908152600160209081526040808320838052909152902054600160a060020a0316151561303357600080fd5b6001608060020a03198416600090815260056020819052604082206004015460ff169081111561305f57fe5b141561306a57600080fd5b6001608060020a0319861660009081526007602052604081206004015460ff16600581111561309557fe5b1415613217576040805160c0810182526001608060020a03198088168252861660208201529081018490526060810183905242608082015260a08101600190526001608060020a03198088166000908152600760209081526040918290208451815486840151608060020a90819004810292049516949094176001608060020a03169390931783559083015180516131339260018501920190613b1f565b506060820151805161314f916002840191602090910190613b1f565b506080820151600382015560a082015160048201805460ff1916600183600581111561317757fe5b021790555050506001608060020a031980861660009081526001602090815260408083208380528252808320548151948b1685528483018281528851928601929092528751600160a060020a03909116947fb2f67f32df451ce918abd944fd5c8a29589848693057913c5060a8e0f293b388948c948a949293606085019291860191819084908490838110156118b557818101518382015260200161189d565b6001608060020a03198616600090815260076020526040902080546001608060020a0316608060020a80870402179055825115613280576001608060020a031986166000908152600760209081526040909120845161327e92600190920191860190613b1f565b505b8151156132b9576001608060020a03198616600090815260076020908152604090912083516132b792600290920191850190613b1f565b505b60008160058111156132c757fe5b14611a05576001608060020a031986166000908152600760205260409020600401805482919060ff191660018360058111156119ff57fe5b600060608181816001608060020a0319871660009081526002602052604090206006015460ff16600581111561333157fe5b141561333c57600080fd5b6001608060020a03198616600090815260026020818152604080842061ffff8a1685526003908101835293819020938401546001808601548654845161010093821615939093026000190116869004601f810186900486028301860190945283825291959460ff909216939185019290918591908301828280156134015780601f106133d657610100808354040283529160200191613401565b820191906000526020600020905b8154815290600101906020018083116133e457829003601f168201915b5050845460408051602060026001851615610100026000190190941693909304601f81018490048402820184019092528181529598508694509250840190508282801561348f5780601f106134645761010080835404028352916020019161348f565b820191906000526020600020905b81548152906001019060200180831161347257829003601f168201915b50505050509050935093509350935092959194509250565b6007602090815260009182526040918290208054600180830180548651600261010094831615949094026000190190911692909204601f8101869004860283018601909652858252608060020a80840296938190040294919290918301828280156135535780601f1061352857610100808354040283529160200191613553565b820191906000526020600020905b81548152906001019060200180831161353657829003601f168201915b50505060028085018054604080516020601f60001961010060018716150201909416959095049283018590048502810185019091528181529596959450909250908301828280156135e55780601f106135ba576101008083540402835291602001916135e5565b820191906000526020600020905b8154815290600101906020018083116135c857829003601f168201915b50505050600383015460049093015491929160ff16905086565b600080546040805160e060020a63babcc539028152600160a060020a0333811660048301529151919092169263babcc53992602480820193602093909283900390910190829087803b15801561365457600080fd5b505af1158015613668573d6000803e3d6000fd5b505050506040513d602081101561367e57600080fd5b5051151561368b57600080fd5b6001608060020a0319881660009081526001602052604081206006015460ff1660058111156136b657fe5b14156136c157600080fd5b6001608060020a031988166000908152600160209081526040808320838052909152902054600160a060020a031615156136fa57600080fd5b6001608060020a03198716600090815260056020819052604082206004015460ff169081111561372657fe5b141561373157600080fd5b6001608060020a0319861660009081526001602052604081206006015460ff16600581111561375c57fe5b141561376757600080fd5b6001608060020a0319891660009081526007602052604081206004015460ff16600581111561379257fe5b14156139b75760408051610120810182526001608060020a0319808b1682528981166020830152881691810191909152600160a060020a03861660608201526080810185905260a0810184905260c081018390524260e08201526101008101600190526001608060020a03198a811660009081526006602090815260409182902084518154868401516001608060020a03918716608060020a93849004179190911690829004820217825592850151600182018054909516939004929092179092556060830151600282018054600160a060020a031916600160a060020a0390921691909117905560808301518051919261389592600385019290910190613b1f565b5060a082015180516138b1916004840191602090910190613b1f565b5060c0820151816005015560e082015181600601556101008201518160070160006101000a81548160ff021916908360058111156138eb57fe5b021790555090505084600160a060020a03167f5c9100ad970db426d43d47b3a9639dc42d9732871d99213dcb4984f7e71509d08a8660405180836001608060020a0319166001608060020a031916815260200180602001828103825283818151815260200191508051906020019080838360005b8381101561397757818101518382015260200161395f565b50505050905090810190601f1680156139a45780820380516001836020036101000a031916815260200191505b50935050505060405180910390a2612d49565b6001608060020a0319898116600090815260066020526040902080546001608060020a0316608060020a808b04810291909117825560019091018054909216908804179055600160a060020a03851615613a41576001608060020a0319891660009081526006602052604090206002018054600160a060020a031916600160a060020a0387161790555b835115613a7a576001608060020a0319891660009081526006602090815260409091208551613a7892600390920191870190613b1f565b505b825115613ab3576001608060020a0319891660009081526006602090815260409091208451613ab192600490920191860190613b1f565b505b8115613ad9576001608060020a0319891660009081526006602052604090206005018290555b6000816005811115613ae757fe5b14612d49576001608060020a031989166000908152600660205260409020600701805482919060ff19166001836005811115612d4357fe5b828054600181600116156101000203166002900490600052602060002090601f016020900481019282601f10613b6057805160ff1916838001178555613b8d565b82800160010185558215613b8d579182015b82811115613b8d578251825591602001919060010190613b72565b50613b99929150613b9d565b5090565b613bb791905b80821115613b995760008155600101613ba3565b905600a165627a7a72305820e606b389ca60f3215d8bcc4ca08c5ef3a5667acacde73855dc024ec06c41a0360029`

// DeployHoQuStorage deploys a new Ethereum contract, binding an instance of HoQuStorage to it.
func DeployHoQuStorage(auth *bind.TransactOpts, backend bind.ContractBackend, configAddress common.Address) (common.Address, *types.Transaction, *HoQuStorage, error) {
	parsed, err := abi.JSON(strings.NewReader(HoQuStorageABI))
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	address, tx, contract, err := bind.DeployContract(auth, parsed, common.FromHex(HoQuStorageBin), backend, configAddress)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &HoQuStorage{HoQuStorageCaller: HoQuStorageCaller{contract: contract}, HoQuStorageTransactor: HoQuStorageTransactor{contract: contract}, HoQuStorageFilterer: HoQuStorageFilterer{contract: contract}}, nil
}

// HoQuStorage is an auto generated Go binding around an Ethereum contract.
type HoQuStorage struct {
	HoQuStorageCaller     // Read-only binding to the contract
	HoQuStorageTransactor // Write-only binding to the contract
	HoQuStorageFilterer   // Log filterer for contract events
}

// HoQuStorageCaller is an auto generated read-only Go binding around an Ethereum contract.
type HoQuStorageCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// HoQuStorageTransactor is an auto generated write-only Go binding around an Ethereum contract.
type HoQuStorageTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// HoQuStorageFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type HoQuStorageFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// HoQuStorageSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type HoQuStorageSession struct {
	Contract     *HoQuStorage      // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// HoQuStorageCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type HoQuStorageCallerSession struct {
	Contract *HoQuStorageCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts      // Call options to use throughout this session
}

// HoQuStorageTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type HoQuStorageTransactorSession struct {
	Contract     *HoQuStorageTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts      // Transaction auth options to use throughout this session
}

// HoQuStorageRaw is an auto generated low-level Go binding around an Ethereum contract.
type HoQuStorageRaw struct {
	Contract *HoQuStorage // Generic contract binding to access the raw methods on
}

// HoQuStorageCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type HoQuStorageCallerRaw struct {
	Contract *HoQuStorageCaller // Generic read-only contract binding to access the raw methods on
}

// HoQuStorageTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type HoQuStorageTransactorRaw struct {
	Contract *HoQuStorageTransactor // Generic write-only contract binding to access the raw methods on
}

// NewHoQuStorage creates a new instance of HoQuStorage, bound to a specific deployed contract.
func NewHoQuStorage(address common.Address, backend bind.ContractBackend) (*HoQuStorage, error) {
	contract, err := bindHoQuStorage(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &HoQuStorage{HoQuStorageCaller: HoQuStorageCaller{contract: contract}, HoQuStorageTransactor: HoQuStorageTransactor{contract: contract}, HoQuStorageFilterer: HoQuStorageFilterer{contract: contract}}, nil
}

// NewHoQuStorageCaller creates a new read-only instance of HoQuStorage, bound to a specific deployed contract.
func NewHoQuStorageCaller(address common.Address, caller bind.ContractCaller) (*HoQuStorageCaller, error) {
	contract, err := bindHoQuStorage(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &HoQuStorageCaller{contract: contract}, nil
}

// NewHoQuStorageTransactor creates a new write-only instance of HoQuStorage, bound to a specific deployed contract.
func NewHoQuStorageTransactor(address common.Address, transactor bind.ContractTransactor) (*HoQuStorageTransactor, error) {
	contract, err := bindHoQuStorage(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &HoQuStorageTransactor{contract: contract}, nil
}

// NewHoQuStorageFilterer creates a new log filterer instance of HoQuStorage, bound to a specific deployed contract.
func NewHoQuStorageFilterer(address common.Address, filterer bind.ContractFilterer) (*HoQuStorageFilterer, error) {
	contract, err := bindHoQuStorage(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &HoQuStorageFilterer{contract: contract}, nil
}

// bindHoQuStorage binds a generic wrapper to an already deployed contract.
func bindHoQuStorage(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(HoQuStorageABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_HoQuStorage *HoQuStorageRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _HoQuStorage.Contract.HoQuStorageCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_HoQuStorage *HoQuStorageRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _HoQuStorage.Contract.HoQuStorageTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_HoQuStorage *HoQuStorageRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _HoQuStorage.Contract.HoQuStorageTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_HoQuStorage *HoQuStorageCallerRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _HoQuStorage.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_HoQuStorage *HoQuStorageTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _HoQuStorage.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_HoQuStorage *HoQuStorageTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _HoQuStorage.Contract.contract.Transact(opts, method, params...)
}

// AdCampaigns is a free data retrieval call binding the contract method 0x89075557.
//
// Solidity: function adCampaigns( bytes16) constant returns(ownerId bytes16, offerId bytes16, contractAddress address, createdAt uint256, status uint8)
func (_HoQuStorage *HoQuStorageCaller) AdCampaigns(opts *bind.CallOpts, arg0 [16]byte) (struct {
	OwnerId         [16]byte
	OfferId         [16]byte
	ContractAddress common.Address
	CreatedAt       *big.Int
	Status          uint8
}, error) {
	ret := new(struct {
		OwnerId         [16]byte
		OfferId         [16]byte
		ContractAddress common.Address
		CreatedAt       *big.Int
		Status          uint8
	})
	out := ret
	err := _HoQuStorage.contract.Call(opts, out, "adCampaigns", arg0)
	return *ret, err
}

// AdCampaigns is a free data retrieval call binding the contract method 0x89075557.
//
// Solidity: function adCampaigns( bytes16) constant returns(ownerId bytes16, offerId bytes16, contractAddress address, createdAt uint256, status uint8)
func (_HoQuStorage *HoQuStorageSession) AdCampaigns(arg0 [16]byte) (struct {
	OwnerId         [16]byte
	OfferId         [16]byte
	ContractAddress common.Address
	CreatedAt       *big.Int
	Status          uint8
}, error) {
	return _HoQuStorage.Contract.AdCampaigns(&_HoQuStorage.CallOpts, arg0)
}

// AdCampaigns is a free data retrieval call binding the contract method 0x89075557.
//
// Solidity: function adCampaigns( bytes16) constant returns(ownerId bytes16, offerId bytes16, contractAddress address, createdAt uint256, status uint8)
func (_HoQuStorage *HoQuStorageCallerSession) AdCampaigns(arg0 [16]byte) (struct {
	OwnerId         [16]byte
	OfferId         [16]byte
	ContractAddress common.Address
	CreatedAt       *big.Int
	Status          uint8
}, error) {
	return _HoQuStorage.Contract.AdCampaigns(&_HoQuStorage.CallOpts, arg0)
}

// Companies is a free data retrieval call binding the contract method 0x4df2d847.
//
// Solidity: function companies( bytes16) constant returns(ownerId bytes16, name string, dataUrl string, createdAt uint256, status uint8)
func (_HoQuStorage *HoQuStorageCaller) Companies(opts *bind.CallOpts, arg0 [16]byte) (struct {
	OwnerId   [16]byte
	Name      string
	DataUrl   string
	CreatedAt *big.Int
	Status    uint8
}, error) {
	ret := new(struct {
		OwnerId   [16]byte
		Name      string
		DataUrl   string
		CreatedAt *big.Int
		Status    uint8
	})
	out := ret
	err := _HoQuStorage.contract.Call(opts, out, "companies", arg0)
	return *ret, err
}

// Companies is a free data retrieval call binding the contract method 0x4df2d847.
//
// Solidity: function companies( bytes16) constant returns(ownerId bytes16, name string, dataUrl string, createdAt uint256, status uint8)
func (_HoQuStorage *HoQuStorageSession) Companies(arg0 [16]byte) (struct {
	OwnerId   [16]byte
	Name      string
	DataUrl   string
	CreatedAt *big.Int
	Status    uint8
}, error) {
	return _HoQuStorage.Contract.Companies(&_HoQuStorage.CallOpts, arg0)
}

// Companies is a free data retrieval call binding the contract method 0x4df2d847.
//
// Solidity: function companies( bytes16) constant returns(ownerId bytes16, name string, dataUrl string, createdAt uint256, status uint8)
func (_HoQuStorage *HoQuStorageCallerSession) Companies(arg0 [16]byte) (struct {
	OwnerId   [16]byte
	Name      string
	DataUrl   string
	CreatedAt *big.Int
	Status    uint8
}, error) {
	return _HoQuStorage.Contract.Companies(&_HoQuStorage.CallOpts, arg0)
}

// Config is a free data retrieval call binding the contract method 0x79502c55.
//
// Solidity: function config() constant returns(address)
func (_HoQuStorage *HoQuStorageCaller) Config(opts *bind.CallOpts) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _HoQuStorage.contract.Call(opts, out, "config")
	return *ret0, err
}

// Config is a free data retrieval call binding the contract method 0x79502c55.
//
// Solidity: function config() constant returns(address)
func (_HoQuStorage *HoQuStorageSession) Config() (common.Address, error) {
	return _HoQuStorage.Contract.Config(&_HoQuStorage.CallOpts)
}

// Config is a free data retrieval call binding the contract method 0x79502c55.
//
// Solidity: function config() constant returns(address)
func (_HoQuStorage *HoQuStorageCallerSession) Config() (common.Address, error) {
	return _HoQuStorage.Contract.Config(&_HoQuStorage.CallOpts)
}

// GetKycReport is a free data retrieval call binding the contract method 0xcc3436b9.
//
// Solidity: function getKycReport(id bytes16, num uint16) constant returns(uint256, string, uint8, string)
func (_HoQuStorage *HoQuStorageCaller) GetKycReport(opts *bind.CallOpts, id [16]byte, num uint16) (*big.Int, string, uint8, string, error) {
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
	err := _HoQuStorage.contract.Call(opts, out, "getKycReport", id, num)
	return *ret0, *ret1, *ret2, *ret3, err
}

// GetKycReport is a free data retrieval call binding the contract method 0xcc3436b9.
//
// Solidity: function getKycReport(id bytes16, num uint16) constant returns(uint256, string, uint8, string)
func (_HoQuStorage *HoQuStorageSession) GetKycReport(id [16]byte, num uint16) (*big.Int, string, uint8, string, error) {
	return _HoQuStorage.Contract.GetKycReport(&_HoQuStorage.CallOpts, id, num)
}

// GetKycReport is a free data retrieval call binding the contract method 0xcc3436b9.
//
// Solidity: function getKycReport(id bytes16, num uint16) constant returns(uint256, string, uint8, string)
func (_HoQuStorage *HoQuStorageCallerSession) GetKycReport(id [16]byte, num uint16) (*big.Int, string, uint8, string, error) {
	return _HoQuStorage.Contract.GetKycReport(&_HoQuStorage.CallOpts, id, num)
}

// GetUserAddress is a free data retrieval call binding the contract method 0xc49e2fb1.
//
// Solidity: function getUserAddress(id bytes16, num uint8) constant returns(address)
func (_HoQuStorage *HoQuStorageCaller) GetUserAddress(opts *bind.CallOpts, id [16]byte, num uint8) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _HoQuStorage.contract.Call(opts, out, "getUserAddress", id, num)
	return *ret0, err
}

// GetUserAddress is a free data retrieval call binding the contract method 0xc49e2fb1.
//
// Solidity: function getUserAddress(id bytes16, num uint8) constant returns(address)
func (_HoQuStorage *HoQuStorageSession) GetUserAddress(id [16]byte, num uint8) (common.Address, error) {
	return _HoQuStorage.Contract.GetUserAddress(&_HoQuStorage.CallOpts, id, num)
}

// GetUserAddress is a free data retrieval call binding the contract method 0xc49e2fb1.
//
// Solidity: function getUserAddress(id bytes16, num uint8) constant returns(address)
func (_HoQuStorage *HoQuStorageCallerSession) GetUserAddress(id [16]byte, num uint8) (common.Address, error) {
	return _HoQuStorage.Contract.GetUserAddress(&_HoQuStorage.CallOpts, id, num)
}

// Ids is a free data retrieval call binding the contract method 0x289e7470.
//
// Solidity: function ids( bytes16) constant returns(userId bytes16, companyId bytes16, idType string, name string, numOfKycReports uint16, createdAt uint256, status uint8)
func (_HoQuStorage *HoQuStorageCaller) Ids(opts *bind.CallOpts, arg0 [16]byte) (struct {
	UserId          [16]byte
	CompanyId       [16]byte
	IdType          string
	Name            string
	NumOfKycReports uint16
	CreatedAt       *big.Int
	Status          uint8
}, error) {
	ret := new(struct {
		UserId          [16]byte
		CompanyId       [16]byte
		IdType          string
		Name            string
		NumOfKycReports uint16
		CreatedAt       *big.Int
		Status          uint8
	})
	out := ret
	err := _HoQuStorage.contract.Call(opts, out, "ids", arg0)
	return *ret, err
}

// Ids is a free data retrieval call binding the contract method 0x289e7470.
//
// Solidity: function ids( bytes16) constant returns(userId bytes16, companyId bytes16, idType string, name string, numOfKycReports uint16, createdAt uint256, status uint8)
func (_HoQuStorage *HoQuStorageSession) Ids(arg0 [16]byte) (struct {
	UserId          [16]byte
	CompanyId       [16]byte
	IdType          string
	Name            string
	NumOfKycReports uint16
	CreatedAt       *big.Int
	Status          uint8
}, error) {
	return _HoQuStorage.Contract.Ids(&_HoQuStorage.CallOpts, arg0)
}

// Ids is a free data retrieval call binding the contract method 0x289e7470.
//
// Solidity: function ids( bytes16) constant returns(userId bytes16, companyId bytes16, idType string, name string, numOfKycReports uint16, createdAt uint256, status uint8)
func (_HoQuStorage *HoQuStorageCallerSession) Ids(arg0 [16]byte) (struct {
	UserId          [16]byte
	CompanyId       [16]byte
	IdType          string
	Name            string
	NumOfKycReports uint16
	CreatedAt       *big.Int
	Status          uint8
}, error) {
	return _HoQuStorage.Contract.Ids(&_HoQuStorage.CallOpts, arg0)
}

// Networks is a free data retrieval call binding the contract method 0x72c130e4.
//
// Solidity: function networks( bytes16) constant returns(ownerId bytes16, name string, dataUrl string, createdAt uint256, status uint8)
func (_HoQuStorage *HoQuStorageCaller) Networks(opts *bind.CallOpts, arg0 [16]byte) (struct {
	OwnerId   [16]byte
	Name      string
	DataUrl   string
	CreatedAt *big.Int
	Status    uint8
}, error) {
	ret := new(struct {
		OwnerId   [16]byte
		Name      string
		DataUrl   string
		CreatedAt *big.Int
		Status    uint8
	})
	out := ret
	err := _HoQuStorage.contract.Call(opts, out, "networks", arg0)
	return *ret, err
}

// Networks is a free data retrieval call binding the contract method 0x72c130e4.
//
// Solidity: function networks( bytes16) constant returns(ownerId bytes16, name string, dataUrl string, createdAt uint256, status uint8)
func (_HoQuStorage *HoQuStorageSession) Networks(arg0 [16]byte) (struct {
	OwnerId   [16]byte
	Name      string
	DataUrl   string
	CreatedAt *big.Int
	Status    uint8
}, error) {
	return _HoQuStorage.Contract.Networks(&_HoQuStorage.CallOpts, arg0)
}

// Networks is a free data retrieval call binding the contract method 0x72c130e4.
//
// Solidity: function networks( bytes16) constant returns(ownerId bytes16, name string, dataUrl string, createdAt uint256, status uint8)
func (_HoQuStorage *HoQuStorageCallerSession) Networks(arg0 [16]byte) (struct {
	OwnerId   [16]byte
	Name      string
	DataUrl   string
	CreatedAt *big.Int
	Status    uint8
}, error) {
	return _HoQuStorage.Contract.Networks(&_HoQuStorage.CallOpts, arg0)
}

// Offers is a free data retrieval call binding the contract method 0xc32d869b.
//
// Solidity: function offers( bytes16) constant returns(ownerId bytes16, networkId bytes16, merchantId bytes16, payerAddress address, name string, dataUrl string, cost uint256, createdAt uint256, status uint8)
func (_HoQuStorage *HoQuStorageCaller) Offers(opts *bind.CallOpts, arg0 [16]byte) (struct {
	OwnerId      [16]byte
	NetworkId    [16]byte
	MerchantId   [16]byte
	PayerAddress common.Address
	Name         string
	DataUrl      string
	Cost         *big.Int
	CreatedAt    *big.Int
	Status       uint8
}, error) {
	ret := new(struct {
		OwnerId      [16]byte
		NetworkId    [16]byte
		MerchantId   [16]byte
		PayerAddress common.Address
		Name         string
		DataUrl      string
		Cost         *big.Int
		CreatedAt    *big.Int
		Status       uint8
	})
	out := ret
	err := _HoQuStorage.contract.Call(opts, out, "offers", arg0)
	return *ret, err
}

// Offers is a free data retrieval call binding the contract method 0xc32d869b.
//
// Solidity: function offers( bytes16) constant returns(ownerId bytes16, networkId bytes16, merchantId bytes16, payerAddress address, name string, dataUrl string, cost uint256, createdAt uint256, status uint8)
func (_HoQuStorage *HoQuStorageSession) Offers(arg0 [16]byte) (struct {
	OwnerId      [16]byte
	NetworkId    [16]byte
	MerchantId   [16]byte
	PayerAddress common.Address
	Name         string
	DataUrl      string
	Cost         *big.Int
	CreatedAt    *big.Int
	Status       uint8
}, error) {
	return _HoQuStorage.Contract.Offers(&_HoQuStorage.CallOpts, arg0)
}

// Offers is a free data retrieval call binding the contract method 0xc32d869b.
//
// Solidity: function offers( bytes16) constant returns(ownerId bytes16, networkId bytes16, merchantId bytes16, payerAddress address, name string, dataUrl string, cost uint256, createdAt uint256, status uint8)
func (_HoQuStorage *HoQuStorageCallerSession) Offers(arg0 [16]byte) (struct {
	OwnerId      [16]byte
	NetworkId    [16]byte
	MerchantId   [16]byte
	PayerAddress common.Address
	Name         string
	DataUrl      string
	Cost         *big.Int
	CreatedAt    *big.Int
	Status       uint8
}, error) {
	return _HoQuStorage.Contract.Offers(&_HoQuStorage.CallOpts, arg0)
}

// Stats is a free data retrieval call binding the contract method 0xa2affa4d.
//
// Solidity: function stats( bytes16) constant returns(rating uint256, volume uint256, contragents uint256, stat1 uint256, stat2 uint256, status uint8)
func (_HoQuStorage *HoQuStorageCaller) Stats(opts *bind.CallOpts, arg0 [16]byte) (struct {
	Rating      *big.Int
	Volume      *big.Int
	Contragents *big.Int
	Stat1       *big.Int
	Stat2       *big.Int
	Status      uint8
}, error) {
	ret := new(struct {
		Rating      *big.Int
		Volume      *big.Int
		Contragents *big.Int
		Stat1       *big.Int
		Stat2       *big.Int
		Status      uint8
	})
	out := ret
	err := _HoQuStorage.contract.Call(opts, out, "stats", arg0)
	return *ret, err
}

// Stats is a free data retrieval call binding the contract method 0xa2affa4d.
//
// Solidity: function stats( bytes16) constant returns(rating uint256, volume uint256, contragents uint256, stat1 uint256, stat2 uint256, status uint8)
func (_HoQuStorage *HoQuStorageSession) Stats(arg0 [16]byte) (struct {
	Rating      *big.Int
	Volume      *big.Int
	Contragents *big.Int
	Stat1       *big.Int
	Stat2       *big.Int
	Status      uint8
}, error) {
	return _HoQuStorage.Contract.Stats(&_HoQuStorage.CallOpts, arg0)
}

// Stats is a free data retrieval call binding the contract method 0xa2affa4d.
//
// Solidity: function stats( bytes16) constant returns(rating uint256, volume uint256, contragents uint256, stat1 uint256, stat2 uint256, status uint8)
func (_HoQuStorage *HoQuStorageCallerSession) Stats(arg0 [16]byte) (struct {
	Rating      *big.Int
	Volume      *big.Int
	Contragents *big.Int
	Stat1       *big.Int
	Stat2       *big.Int
	Status      uint8
}, error) {
	return _HoQuStorage.Contract.Stats(&_HoQuStorage.CallOpts, arg0)
}

// Trackers is a free data retrieval call binding the contract method 0xea2fdd8c.
//
// Solidity: function trackers( bytes16) constant returns(ownerId bytes16, networkId bytes16, name string, dataUrl string, createdAt uint256, status uint8)
func (_HoQuStorage *HoQuStorageCaller) Trackers(opts *bind.CallOpts, arg0 [16]byte) (struct {
	OwnerId   [16]byte
	NetworkId [16]byte
	Name      string
	DataUrl   string
	CreatedAt *big.Int
	Status    uint8
}, error) {
	ret := new(struct {
		OwnerId   [16]byte
		NetworkId [16]byte
		Name      string
		DataUrl   string
		CreatedAt *big.Int
		Status    uint8
	})
	out := ret
	err := _HoQuStorage.contract.Call(opts, out, "trackers", arg0)
	return *ret, err
}

// Trackers is a free data retrieval call binding the contract method 0xea2fdd8c.
//
// Solidity: function trackers( bytes16) constant returns(ownerId bytes16, networkId bytes16, name string, dataUrl string, createdAt uint256, status uint8)
func (_HoQuStorage *HoQuStorageSession) Trackers(arg0 [16]byte) (struct {
	OwnerId   [16]byte
	NetworkId [16]byte
	Name      string
	DataUrl   string
	CreatedAt *big.Int
	Status    uint8
}, error) {
	return _HoQuStorage.Contract.Trackers(&_HoQuStorage.CallOpts, arg0)
}

// Trackers is a free data retrieval call binding the contract method 0xea2fdd8c.
//
// Solidity: function trackers( bytes16) constant returns(ownerId bytes16, networkId bytes16, name string, dataUrl string, createdAt uint256, status uint8)
func (_HoQuStorage *HoQuStorageCallerSession) Trackers(arg0 [16]byte) (struct {
	OwnerId   [16]byte
	NetworkId [16]byte
	Name      string
	DataUrl   string
	CreatedAt *big.Int
	Status    uint8
}, error) {
	return _HoQuStorage.Contract.Trackers(&_HoQuStorage.CallOpts, arg0)
}

// Users is a free data retrieval call binding the contract method 0x7520dd14.
//
// Solidity: function users( bytes16) constant returns(numOfAddresses uint8, role string, kycLevel uint8, pubKey string, createdAt uint256, status uint8)
func (_HoQuStorage *HoQuStorageCaller) Users(opts *bind.CallOpts, arg0 [16]byte) (struct {
	NumOfAddresses uint8
	Role           string
	KycLevel       uint8
	PubKey         string
	CreatedAt      *big.Int
	Status         uint8
}, error) {
	ret := new(struct {
		NumOfAddresses uint8
		Role           string
		KycLevel       uint8
		PubKey         string
		CreatedAt      *big.Int
		Status         uint8
	})
	out := ret
	err := _HoQuStorage.contract.Call(opts, out, "users", arg0)
	return *ret, err
}

// Users is a free data retrieval call binding the contract method 0x7520dd14.
//
// Solidity: function users( bytes16) constant returns(numOfAddresses uint8, role string, kycLevel uint8, pubKey string, createdAt uint256, status uint8)
func (_HoQuStorage *HoQuStorageSession) Users(arg0 [16]byte) (struct {
	NumOfAddresses uint8
	Role           string
	KycLevel       uint8
	PubKey         string
	CreatedAt      *big.Int
	Status         uint8
}, error) {
	return _HoQuStorage.Contract.Users(&_HoQuStorage.CallOpts, arg0)
}

// Users is a free data retrieval call binding the contract method 0x7520dd14.
//
// Solidity: function users( bytes16) constant returns(numOfAddresses uint8, role string, kycLevel uint8, pubKey string, createdAt uint256, status uint8)
func (_HoQuStorage *HoQuStorageCallerSession) Users(arg0 [16]byte) (struct {
	NumOfAddresses uint8
	Role           string
	KycLevel       uint8
	PubKey         string
	CreatedAt      *big.Int
	Status         uint8
}, error) {
	return _HoQuStorage.Contract.Users(&_HoQuStorage.CallOpts, arg0)
}

// AddKycReport is a paid mutator transaction binding the contract method 0x18c8f305.
//
// Solidity: function addKycReport(id bytes16, meta string, kycLevel uint8, dataUrl string) returns()
func (_HoQuStorage *HoQuStorageTransactor) AddKycReport(opts *bind.TransactOpts, id [16]byte, meta string, kycLevel uint8, dataUrl string) (*types.Transaction, error) {
	return _HoQuStorage.contract.Transact(opts, "addKycReport", id, meta, kycLevel, dataUrl)
}

// AddKycReport is a paid mutator transaction binding the contract method 0x18c8f305.
//
// Solidity: function addKycReport(id bytes16, meta string, kycLevel uint8, dataUrl string) returns()
func (_HoQuStorage *HoQuStorageSession) AddKycReport(id [16]byte, meta string, kycLevel uint8, dataUrl string) (*types.Transaction, error) {
	return _HoQuStorage.Contract.AddKycReport(&_HoQuStorage.TransactOpts, id, meta, kycLevel, dataUrl)
}

// AddKycReport is a paid mutator transaction binding the contract method 0x18c8f305.
//
// Solidity: function addKycReport(id bytes16, meta string, kycLevel uint8, dataUrl string) returns()
func (_HoQuStorage *HoQuStorageTransactorSession) AddKycReport(id [16]byte, meta string, kycLevel uint8, dataUrl string) (*types.Transaction, error) {
	return _HoQuStorage.Contract.AddKycReport(&_HoQuStorage.TransactOpts, id, meta, kycLevel, dataUrl)
}

// AddUserAddress is a paid mutator transaction binding the contract method 0x03806028.
//
// Solidity: function addUserAddress(id bytes16, ownerAddress address) returns()
func (_HoQuStorage *HoQuStorageTransactor) AddUserAddress(opts *bind.TransactOpts, id [16]byte, ownerAddress common.Address) (*types.Transaction, error) {
	return _HoQuStorage.contract.Transact(opts, "addUserAddress", id, ownerAddress)
}

// AddUserAddress is a paid mutator transaction binding the contract method 0x03806028.
//
// Solidity: function addUserAddress(id bytes16, ownerAddress address) returns()
func (_HoQuStorage *HoQuStorageSession) AddUserAddress(id [16]byte, ownerAddress common.Address) (*types.Transaction, error) {
	return _HoQuStorage.Contract.AddUserAddress(&_HoQuStorage.TransactOpts, id, ownerAddress)
}

// AddUserAddress is a paid mutator transaction binding the contract method 0x03806028.
//
// Solidity: function addUserAddress(id bytes16, ownerAddress address) returns()
func (_HoQuStorage *HoQuStorageTransactorSession) AddUserAddress(id [16]byte, ownerAddress common.Address) (*types.Transaction, error) {
	return _HoQuStorage.Contract.AddUserAddress(&_HoQuStorage.TransactOpts, id, ownerAddress)
}

// SetAdCampaign is a paid mutator transaction binding the contract method 0xa0191d20.
//
// Solidity: function setAdCampaign(id bytes16, ownerId bytes16, offerId bytes16, contractAddress address, status uint8) returns()
func (_HoQuStorage *HoQuStorageTransactor) SetAdCampaign(opts *bind.TransactOpts, id [16]byte, ownerId [16]byte, offerId [16]byte, contractAddress common.Address, status uint8) (*types.Transaction, error) {
	return _HoQuStorage.contract.Transact(opts, "setAdCampaign", id, ownerId, offerId, contractAddress, status)
}

// SetAdCampaign is a paid mutator transaction binding the contract method 0xa0191d20.
//
// Solidity: function setAdCampaign(id bytes16, ownerId bytes16, offerId bytes16, contractAddress address, status uint8) returns()
func (_HoQuStorage *HoQuStorageSession) SetAdCampaign(id [16]byte, ownerId [16]byte, offerId [16]byte, contractAddress common.Address, status uint8) (*types.Transaction, error) {
	return _HoQuStorage.Contract.SetAdCampaign(&_HoQuStorage.TransactOpts, id, ownerId, offerId, contractAddress, status)
}

// SetAdCampaign is a paid mutator transaction binding the contract method 0xa0191d20.
//
// Solidity: function setAdCampaign(id bytes16, ownerId bytes16, offerId bytes16, contractAddress address, status uint8) returns()
func (_HoQuStorage *HoQuStorageTransactorSession) SetAdCampaign(id [16]byte, ownerId [16]byte, offerId [16]byte, contractAddress common.Address, status uint8) (*types.Transaction, error) {
	return _HoQuStorage.Contract.SetAdCampaign(&_HoQuStorage.TransactOpts, id, ownerId, offerId, contractAddress, status)
}

// SetCompany is a paid mutator transaction binding the contract method 0x768d80a8.
//
// Solidity: function setCompany(id bytes16, ownerId bytes16, name string, dataUrl string, status uint8) returns()
func (_HoQuStorage *HoQuStorageTransactor) SetCompany(opts *bind.TransactOpts, id [16]byte, ownerId [16]byte, name string, dataUrl string, status uint8) (*types.Transaction, error) {
	return _HoQuStorage.contract.Transact(opts, "setCompany", id, ownerId, name, dataUrl, status)
}

// SetCompany is a paid mutator transaction binding the contract method 0x768d80a8.
//
// Solidity: function setCompany(id bytes16, ownerId bytes16, name string, dataUrl string, status uint8) returns()
func (_HoQuStorage *HoQuStorageSession) SetCompany(id [16]byte, ownerId [16]byte, name string, dataUrl string, status uint8) (*types.Transaction, error) {
	return _HoQuStorage.Contract.SetCompany(&_HoQuStorage.TransactOpts, id, ownerId, name, dataUrl, status)
}

// SetCompany is a paid mutator transaction binding the contract method 0x768d80a8.
//
// Solidity: function setCompany(id bytes16, ownerId bytes16, name string, dataUrl string, status uint8) returns()
func (_HoQuStorage *HoQuStorageTransactorSession) SetCompany(id [16]byte, ownerId [16]byte, name string, dataUrl string, status uint8) (*types.Transaction, error) {
	return _HoQuStorage.Contract.SetCompany(&_HoQuStorage.TransactOpts, id, ownerId, name, dataUrl, status)
}

// SetConfigAddress is a paid mutator transaction binding the contract method 0x83a12de9.
//
// Solidity: function setConfigAddress(configAddress address) returns()
func (_HoQuStorage *HoQuStorageTransactor) SetConfigAddress(opts *bind.TransactOpts, configAddress common.Address) (*types.Transaction, error) {
	return _HoQuStorage.contract.Transact(opts, "setConfigAddress", configAddress)
}

// SetConfigAddress is a paid mutator transaction binding the contract method 0x83a12de9.
//
// Solidity: function setConfigAddress(configAddress address) returns()
func (_HoQuStorage *HoQuStorageSession) SetConfigAddress(configAddress common.Address) (*types.Transaction, error) {
	return _HoQuStorage.Contract.SetConfigAddress(&_HoQuStorage.TransactOpts, configAddress)
}

// SetConfigAddress is a paid mutator transaction binding the contract method 0x83a12de9.
//
// Solidity: function setConfigAddress(configAddress address) returns()
func (_HoQuStorage *HoQuStorageTransactorSession) SetConfigAddress(configAddress common.Address) (*types.Transaction, error) {
	return _HoQuStorage.Contract.SetConfigAddress(&_HoQuStorage.TransactOpts, configAddress)
}

// SetIdentification is a paid mutator transaction binding the contract method 0xacc02cce.
//
// Solidity: function setIdentification(id bytes16, userId bytes16, idType string, name string, companyId bytes16, status uint8) returns()
func (_HoQuStorage *HoQuStorageTransactor) SetIdentification(opts *bind.TransactOpts, id [16]byte, userId [16]byte, idType string, name string, companyId [16]byte, status uint8) (*types.Transaction, error) {
	return _HoQuStorage.contract.Transact(opts, "setIdentification", id, userId, idType, name, companyId, status)
}

// SetIdentification is a paid mutator transaction binding the contract method 0xacc02cce.
//
// Solidity: function setIdentification(id bytes16, userId bytes16, idType string, name string, companyId bytes16, status uint8) returns()
func (_HoQuStorage *HoQuStorageSession) SetIdentification(id [16]byte, userId [16]byte, idType string, name string, companyId [16]byte, status uint8) (*types.Transaction, error) {
	return _HoQuStorage.Contract.SetIdentification(&_HoQuStorage.TransactOpts, id, userId, idType, name, companyId, status)
}

// SetIdentification is a paid mutator transaction binding the contract method 0xacc02cce.
//
// Solidity: function setIdentification(id bytes16, userId bytes16, idType string, name string, companyId bytes16, status uint8) returns()
func (_HoQuStorage *HoQuStorageTransactorSession) SetIdentification(id [16]byte, userId [16]byte, idType string, name string, companyId [16]byte, status uint8) (*types.Transaction, error) {
	return _HoQuStorage.Contract.SetIdentification(&_HoQuStorage.TransactOpts, id, userId, idType, name, companyId, status)
}

// SetNetwork is a paid mutator transaction binding the contract method 0x02884980.
//
// Solidity: function setNetwork(id bytes16, ownerId bytes16, name string, dataUrl string, status uint8) returns()
func (_HoQuStorage *HoQuStorageTransactor) SetNetwork(opts *bind.TransactOpts, id [16]byte, ownerId [16]byte, name string, dataUrl string, status uint8) (*types.Transaction, error) {
	return _HoQuStorage.contract.Transact(opts, "setNetwork", id, ownerId, name, dataUrl, status)
}

// SetNetwork is a paid mutator transaction binding the contract method 0x02884980.
//
// Solidity: function setNetwork(id bytes16, ownerId bytes16, name string, dataUrl string, status uint8) returns()
func (_HoQuStorage *HoQuStorageSession) SetNetwork(id [16]byte, ownerId [16]byte, name string, dataUrl string, status uint8) (*types.Transaction, error) {
	return _HoQuStorage.Contract.SetNetwork(&_HoQuStorage.TransactOpts, id, ownerId, name, dataUrl, status)
}

// SetNetwork is a paid mutator transaction binding the contract method 0x02884980.
//
// Solidity: function setNetwork(id bytes16, ownerId bytes16, name string, dataUrl string, status uint8) returns()
func (_HoQuStorage *HoQuStorageTransactorSession) SetNetwork(id [16]byte, ownerId [16]byte, name string, dataUrl string, status uint8) (*types.Transaction, error) {
	return _HoQuStorage.Contract.SetNetwork(&_HoQuStorage.TransactOpts, id, ownerId, name, dataUrl, status)
}

// SetOffer is a paid mutator transaction binding the contract method 0xf3b412f1.
//
// Solidity: function setOffer(id bytes16, ownerId bytes16, networkId bytes16, merchantId bytes16, payerAddress address, name string, dataUrl string, cost uint256, status uint8) returns()
func (_HoQuStorage *HoQuStorageTransactor) SetOffer(opts *bind.TransactOpts, id [16]byte, ownerId [16]byte, networkId [16]byte, merchantId [16]byte, payerAddress common.Address, name string, dataUrl string, cost *big.Int, status uint8) (*types.Transaction, error) {
	return _HoQuStorage.contract.Transact(opts, "setOffer", id, ownerId, networkId, merchantId, payerAddress, name, dataUrl, cost, status)
}

// SetOffer is a paid mutator transaction binding the contract method 0xf3b412f1.
//
// Solidity: function setOffer(id bytes16, ownerId bytes16, networkId bytes16, merchantId bytes16, payerAddress address, name string, dataUrl string, cost uint256, status uint8) returns()
func (_HoQuStorage *HoQuStorageSession) SetOffer(id [16]byte, ownerId [16]byte, networkId [16]byte, merchantId [16]byte, payerAddress common.Address, name string, dataUrl string, cost *big.Int, status uint8) (*types.Transaction, error) {
	return _HoQuStorage.Contract.SetOffer(&_HoQuStorage.TransactOpts, id, ownerId, networkId, merchantId, payerAddress, name, dataUrl, cost, status)
}

// SetOffer is a paid mutator transaction binding the contract method 0xf3b412f1.
//
// Solidity: function setOffer(id bytes16, ownerId bytes16, networkId bytes16, merchantId bytes16, payerAddress address, name string, dataUrl string, cost uint256, status uint8) returns()
func (_HoQuStorage *HoQuStorageTransactorSession) SetOffer(id [16]byte, ownerId [16]byte, networkId [16]byte, merchantId [16]byte, payerAddress common.Address, name string, dataUrl string, cost *big.Int, status uint8) (*types.Transaction, error) {
	return _HoQuStorage.Contract.SetOffer(&_HoQuStorage.TransactOpts, id, ownerId, networkId, merchantId, payerAddress, name, dataUrl, cost, status)
}

// SetStats is a paid mutator transaction binding the contract method 0xb1dbfe50.
//
// Solidity: function setStats(id bytes16, userId bytes16, rating uint256, volume uint256, contragents uint256, stat1 uint256, stat2 uint256, status uint8) returns()
func (_HoQuStorage *HoQuStorageTransactor) SetStats(opts *bind.TransactOpts, id [16]byte, userId [16]byte, rating *big.Int, volume *big.Int, contragents *big.Int, stat1 *big.Int, stat2 *big.Int, status uint8) (*types.Transaction, error) {
	return _HoQuStorage.contract.Transact(opts, "setStats", id, userId, rating, volume, contragents, stat1, stat2, status)
}

// SetStats is a paid mutator transaction binding the contract method 0xb1dbfe50.
//
// Solidity: function setStats(id bytes16, userId bytes16, rating uint256, volume uint256, contragents uint256, stat1 uint256, stat2 uint256, status uint8) returns()
func (_HoQuStorage *HoQuStorageSession) SetStats(id [16]byte, userId [16]byte, rating *big.Int, volume *big.Int, contragents *big.Int, stat1 *big.Int, stat2 *big.Int, status uint8) (*types.Transaction, error) {
	return _HoQuStorage.Contract.SetStats(&_HoQuStorage.TransactOpts, id, userId, rating, volume, contragents, stat1, stat2, status)
}

// SetStats is a paid mutator transaction binding the contract method 0xb1dbfe50.
//
// Solidity: function setStats(id bytes16, userId bytes16, rating uint256, volume uint256, contragents uint256, stat1 uint256, stat2 uint256, status uint8) returns()
func (_HoQuStorage *HoQuStorageTransactorSession) SetStats(id [16]byte, userId [16]byte, rating *big.Int, volume *big.Int, contragents *big.Int, stat1 *big.Int, stat2 *big.Int, status uint8) (*types.Transaction, error) {
	return _HoQuStorage.Contract.SetStats(&_HoQuStorage.TransactOpts, id, userId, rating, volume, contragents, stat1, stat2, status)
}

// SetTracker is a paid mutator transaction binding the contract method 0xc4d4d887.
//
// Solidity: function setTracker(id bytes16, ownerId bytes16, networkId bytes16, name string, dataUrl string, status uint8) returns()
func (_HoQuStorage *HoQuStorageTransactor) SetTracker(opts *bind.TransactOpts, id [16]byte, ownerId [16]byte, networkId [16]byte, name string, dataUrl string, status uint8) (*types.Transaction, error) {
	return _HoQuStorage.contract.Transact(opts, "setTracker", id, ownerId, networkId, name, dataUrl, status)
}

// SetTracker is a paid mutator transaction binding the contract method 0xc4d4d887.
//
// Solidity: function setTracker(id bytes16, ownerId bytes16, networkId bytes16, name string, dataUrl string, status uint8) returns()
func (_HoQuStorage *HoQuStorageSession) SetTracker(id [16]byte, ownerId [16]byte, networkId [16]byte, name string, dataUrl string, status uint8) (*types.Transaction, error) {
	return _HoQuStorage.Contract.SetTracker(&_HoQuStorage.TransactOpts, id, ownerId, networkId, name, dataUrl, status)
}

// SetTracker is a paid mutator transaction binding the contract method 0xc4d4d887.
//
// Solidity: function setTracker(id bytes16, ownerId bytes16, networkId bytes16, name string, dataUrl string, status uint8) returns()
func (_HoQuStorage *HoQuStorageTransactorSession) SetTracker(id [16]byte, ownerId [16]byte, networkId [16]byte, name string, dataUrl string, status uint8) (*types.Transaction, error) {
	return _HoQuStorage.Contract.SetTracker(&_HoQuStorage.TransactOpts, id, ownerId, networkId, name, dataUrl, status)
}

// SetUser is a paid mutator transaction binding the contract method 0x253a9b9e.
//
// Solidity: function setUser(id bytes16, role string, ownerAddress address, kycLevel uint8, pubKey string, status uint8) returns()
func (_HoQuStorage *HoQuStorageTransactor) SetUser(opts *bind.TransactOpts, id [16]byte, role string, ownerAddress common.Address, kycLevel uint8, pubKey string, status uint8) (*types.Transaction, error) {
	return _HoQuStorage.contract.Transact(opts, "setUser", id, role, ownerAddress, kycLevel, pubKey, status)
}

// SetUser is a paid mutator transaction binding the contract method 0x253a9b9e.
//
// Solidity: function setUser(id bytes16, role string, ownerAddress address, kycLevel uint8, pubKey string, status uint8) returns()
func (_HoQuStorage *HoQuStorageSession) SetUser(id [16]byte, role string, ownerAddress common.Address, kycLevel uint8, pubKey string, status uint8) (*types.Transaction, error) {
	return _HoQuStorage.Contract.SetUser(&_HoQuStorage.TransactOpts, id, role, ownerAddress, kycLevel, pubKey, status)
}

// SetUser is a paid mutator transaction binding the contract method 0x253a9b9e.
//
// Solidity: function setUser(id bytes16, role string, ownerAddress address, kycLevel uint8, pubKey string, status uint8) returns()
func (_HoQuStorage *HoQuStorageTransactorSession) SetUser(id [16]byte, role string, ownerAddress common.Address, kycLevel uint8, pubKey string, status uint8) (*types.Transaction, error) {
	return _HoQuStorage.Contract.SetUser(&_HoQuStorage.TransactOpts, id, role, ownerAddress, kycLevel, pubKey, status)
}

// HoQuStorageAdCampaignAddedIterator is returned from FilterAdCampaignAdded and is used to iterate over the raw logs and unpacked data for AdCampaignAdded events raised by the HoQuStorage contract.
type HoQuStorageAdCampaignAddedIterator struct {
	Event *HoQuStorageAdCampaignAdded // Event containing the contract specifics and raw log

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
func (it *HoQuStorageAdCampaignAddedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(HoQuStorageAdCampaignAdded)
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
		it.Event = new(HoQuStorageAdCampaignAdded)
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
func (it *HoQuStorageAdCampaignAddedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *HoQuStorageAdCampaignAddedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// HoQuStorageAdCampaignAdded represents a AdCampaignAdded event raised by the HoQuStorage contract.
type HoQuStorageAdCampaignAdded struct {
	OwnerAddress    common.Address
	Id              [16]byte
	ContractAddress common.Address
	Raw             types.Log // Blockchain specific contextual infos
}

// FilterAdCampaignAdded is a free log retrieval operation binding the contract event 0x62a609593cded926aabb6aef1e5b884804ad29d1bd1116ac4ce32085690d5ab2.
//
// Solidity: event AdCampaignAdded(ownerAddress indexed address, id bytes16, contractAddress address)
func (_HoQuStorage *HoQuStorageFilterer) FilterAdCampaignAdded(opts *bind.FilterOpts, ownerAddress []common.Address) (*HoQuStorageAdCampaignAddedIterator, error) {

	var ownerAddressRule []interface{}
	for _, ownerAddressItem := range ownerAddress {
		ownerAddressRule = append(ownerAddressRule, ownerAddressItem)
	}

	logs, sub, err := _HoQuStorage.contract.FilterLogs(opts, "AdCampaignAdded", ownerAddressRule)
	if err != nil {
		return nil, err
	}
	return &HoQuStorageAdCampaignAddedIterator{contract: _HoQuStorage.contract, event: "AdCampaignAdded", logs: logs, sub: sub}, nil
}

// WatchAdCampaignAdded is a free log subscription operation binding the contract event 0x62a609593cded926aabb6aef1e5b884804ad29d1bd1116ac4ce32085690d5ab2.
//
// Solidity: event AdCampaignAdded(ownerAddress indexed address, id bytes16, contractAddress address)
func (_HoQuStorage *HoQuStorageFilterer) WatchAdCampaignAdded(opts *bind.WatchOpts, sink chan<- *HoQuStorageAdCampaignAdded, ownerAddress []common.Address) (event.Subscription, error) {

	var ownerAddressRule []interface{}
	for _, ownerAddressItem := range ownerAddress {
		ownerAddressRule = append(ownerAddressRule, ownerAddressItem)
	}

	logs, sub, err := _HoQuStorage.contract.WatchLogs(opts, "AdCampaignAdded", ownerAddressRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(HoQuStorageAdCampaignAdded)
				if err := _HoQuStorage.contract.UnpackLog(event, "AdCampaignAdded", log); err != nil {
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

// HoQuStorageCompanyRegisteredIterator is returned from FilterCompanyRegistered and is used to iterate over the raw logs and unpacked data for CompanyRegistered events raised by the HoQuStorage contract.
type HoQuStorageCompanyRegisteredIterator struct {
	Event *HoQuStorageCompanyRegistered // Event containing the contract specifics and raw log

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
func (it *HoQuStorageCompanyRegisteredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(HoQuStorageCompanyRegistered)
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
		it.Event = new(HoQuStorageCompanyRegistered)
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
func (it *HoQuStorageCompanyRegisteredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *HoQuStorageCompanyRegisteredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// HoQuStorageCompanyRegistered represents a CompanyRegistered event raised by the HoQuStorage contract.
type HoQuStorageCompanyRegistered struct {
	OwnerAddress common.Address
	Id           [16]byte
	Name         string
	Raw          types.Log // Blockchain specific contextual infos
}

// FilterCompanyRegistered is a free log retrieval operation binding the contract event 0x0eb23d662c57754d23afdc1cffa6246fbae3bbe843eec5c42a5a5223c21414ed.
//
// Solidity: event CompanyRegistered(ownerAddress indexed address, id bytes16, name string)
func (_HoQuStorage *HoQuStorageFilterer) FilterCompanyRegistered(opts *bind.FilterOpts, ownerAddress []common.Address) (*HoQuStorageCompanyRegisteredIterator, error) {

	var ownerAddressRule []interface{}
	for _, ownerAddressItem := range ownerAddress {
		ownerAddressRule = append(ownerAddressRule, ownerAddressItem)
	}

	logs, sub, err := _HoQuStorage.contract.FilterLogs(opts, "CompanyRegistered", ownerAddressRule)
	if err != nil {
		return nil, err
	}
	return &HoQuStorageCompanyRegisteredIterator{contract: _HoQuStorage.contract, event: "CompanyRegistered", logs: logs, sub: sub}, nil
}

// WatchCompanyRegistered is a free log subscription operation binding the contract event 0x0eb23d662c57754d23afdc1cffa6246fbae3bbe843eec5c42a5a5223c21414ed.
//
// Solidity: event CompanyRegistered(ownerAddress indexed address, id bytes16, name string)
func (_HoQuStorage *HoQuStorageFilterer) WatchCompanyRegistered(opts *bind.WatchOpts, sink chan<- *HoQuStorageCompanyRegistered, ownerAddress []common.Address) (event.Subscription, error) {

	var ownerAddressRule []interface{}
	for _, ownerAddressItem := range ownerAddress {
		ownerAddressRule = append(ownerAddressRule, ownerAddressItem)
	}

	logs, sub, err := _HoQuStorage.contract.WatchLogs(opts, "CompanyRegistered", ownerAddressRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(HoQuStorageCompanyRegistered)
				if err := _HoQuStorage.contract.UnpackLog(event, "CompanyRegistered", log); err != nil {
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

// HoQuStorageIdentificationAddedIterator is returned from FilterIdentificationAdded and is used to iterate over the raw logs and unpacked data for IdentificationAdded events raised by the HoQuStorage contract.
type HoQuStorageIdentificationAddedIterator struct {
	Event *HoQuStorageIdentificationAdded // Event containing the contract specifics and raw log

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
func (it *HoQuStorageIdentificationAddedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(HoQuStorageIdentificationAdded)
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
		it.Event = new(HoQuStorageIdentificationAdded)
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
func (it *HoQuStorageIdentificationAddedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *HoQuStorageIdentificationAddedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// HoQuStorageIdentificationAdded represents a IdentificationAdded event raised by the HoQuStorage contract.
type HoQuStorageIdentificationAdded struct {
	OwnerAddress common.Address
	Id           [16]byte
	Name         string
	Raw          types.Log // Blockchain specific contextual infos
}

// FilterIdentificationAdded is a free log retrieval operation binding the contract event 0xa5c73c5ece99b6456d1b224c43e2a1fc268c1e3a14156f135f3e8bdec2ff9a5a.
//
// Solidity: event IdentificationAdded(ownerAddress indexed address, id bytes16, name string)
func (_HoQuStorage *HoQuStorageFilterer) FilterIdentificationAdded(opts *bind.FilterOpts, ownerAddress []common.Address) (*HoQuStorageIdentificationAddedIterator, error) {

	var ownerAddressRule []interface{}
	for _, ownerAddressItem := range ownerAddress {
		ownerAddressRule = append(ownerAddressRule, ownerAddressItem)
	}

	logs, sub, err := _HoQuStorage.contract.FilterLogs(opts, "IdentificationAdded", ownerAddressRule)
	if err != nil {
		return nil, err
	}
	return &HoQuStorageIdentificationAddedIterator{contract: _HoQuStorage.contract, event: "IdentificationAdded", logs: logs, sub: sub}, nil
}

// WatchIdentificationAdded is a free log subscription operation binding the contract event 0xa5c73c5ece99b6456d1b224c43e2a1fc268c1e3a14156f135f3e8bdec2ff9a5a.
//
// Solidity: event IdentificationAdded(ownerAddress indexed address, id bytes16, name string)
func (_HoQuStorage *HoQuStorageFilterer) WatchIdentificationAdded(opts *bind.WatchOpts, sink chan<- *HoQuStorageIdentificationAdded, ownerAddress []common.Address) (event.Subscription, error) {

	var ownerAddressRule []interface{}
	for _, ownerAddressItem := range ownerAddress {
		ownerAddressRule = append(ownerAddressRule, ownerAddressItem)
	}

	logs, sub, err := _HoQuStorage.contract.WatchLogs(opts, "IdentificationAdded", ownerAddressRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(HoQuStorageIdentificationAdded)
				if err := _HoQuStorage.contract.UnpackLog(event, "IdentificationAdded", log); err != nil {
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

// HoQuStorageKycReportAddedIterator is returned from FilterKycReportAdded and is used to iterate over the raw logs and unpacked data for KycReportAdded events raised by the HoQuStorage contract.
type HoQuStorageKycReportAddedIterator struct {
	Event *HoQuStorageKycReportAdded // Event containing the contract specifics and raw log

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
func (it *HoQuStorageKycReportAddedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(HoQuStorageKycReportAdded)
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
		it.Event = new(HoQuStorageKycReportAdded)
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
func (it *HoQuStorageKycReportAddedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *HoQuStorageKycReportAddedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// HoQuStorageKycReportAdded represents a KycReportAdded event raised by the HoQuStorage contract.
type HoQuStorageKycReportAdded struct {
	OwnerAddress common.Address
	KycLevel     uint8
	Raw          types.Log // Blockchain specific contextual infos
}

// FilterKycReportAdded is a free log retrieval operation binding the contract event 0x92262b4b81e23d81eaec5e9e8c9439f0a59929da8c5d49b32bca92c112f4172a.
//
// Solidity: event KycReportAdded(ownerAddress indexed address, kycLevel uint8)
func (_HoQuStorage *HoQuStorageFilterer) FilterKycReportAdded(opts *bind.FilterOpts, ownerAddress []common.Address) (*HoQuStorageKycReportAddedIterator, error) {

	var ownerAddressRule []interface{}
	for _, ownerAddressItem := range ownerAddress {
		ownerAddressRule = append(ownerAddressRule, ownerAddressItem)
	}

	logs, sub, err := _HoQuStorage.contract.FilterLogs(opts, "KycReportAdded", ownerAddressRule)
	if err != nil {
		return nil, err
	}
	return &HoQuStorageKycReportAddedIterator{contract: _HoQuStorage.contract, event: "KycReportAdded", logs: logs, sub: sub}, nil
}

// WatchKycReportAdded is a free log subscription operation binding the contract event 0x92262b4b81e23d81eaec5e9e8c9439f0a59929da8c5d49b32bca92c112f4172a.
//
// Solidity: event KycReportAdded(ownerAddress indexed address, kycLevel uint8)
func (_HoQuStorage *HoQuStorageFilterer) WatchKycReportAdded(opts *bind.WatchOpts, sink chan<- *HoQuStorageKycReportAdded, ownerAddress []common.Address) (event.Subscription, error) {

	var ownerAddressRule []interface{}
	for _, ownerAddressItem := range ownerAddress {
		ownerAddressRule = append(ownerAddressRule, ownerAddressItem)
	}

	logs, sub, err := _HoQuStorage.contract.WatchLogs(opts, "KycReportAdded", ownerAddressRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(HoQuStorageKycReportAdded)
				if err := _HoQuStorage.contract.UnpackLog(event, "KycReportAdded", log); err != nil {
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

// HoQuStorageNetworkRegisteredIterator is returned from FilterNetworkRegistered and is used to iterate over the raw logs and unpacked data for NetworkRegistered events raised by the HoQuStorage contract.
type HoQuStorageNetworkRegisteredIterator struct {
	Event *HoQuStorageNetworkRegistered // Event containing the contract specifics and raw log

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
func (it *HoQuStorageNetworkRegisteredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(HoQuStorageNetworkRegistered)
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
		it.Event = new(HoQuStorageNetworkRegistered)
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
func (it *HoQuStorageNetworkRegisteredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *HoQuStorageNetworkRegisteredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// HoQuStorageNetworkRegistered represents a NetworkRegistered event raised by the HoQuStorage contract.
type HoQuStorageNetworkRegistered struct {
	OwnerAddress common.Address
	Id           [16]byte
	Name         string
	Raw          types.Log // Blockchain specific contextual infos
}

// FilterNetworkRegistered is a free log retrieval operation binding the contract event 0x592e83ede8c846dffd6e4edcbe77983b153bec820ac697aa34a236147b1341e0.
//
// Solidity: event NetworkRegistered(ownerAddress indexed address, id bytes16, name string)
func (_HoQuStorage *HoQuStorageFilterer) FilterNetworkRegistered(opts *bind.FilterOpts, ownerAddress []common.Address) (*HoQuStorageNetworkRegisteredIterator, error) {

	var ownerAddressRule []interface{}
	for _, ownerAddressItem := range ownerAddress {
		ownerAddressRule = append(ownerAddressRule, ownerAddressItem)
	}

	logs, sub, err := _HoQuStorage.contract.FilterLogs(opts, "NetworkRegistered", ownerAddressRule)
	if err != nil {
		return nil, err
	}
	return &HoQuStorageNetworkRegisteredIterator{contract: _HoQuStorage.contract, event: "NetworkRegistered", logs: logs, sub: sub}, nil
}

// WatchNetworkRegistered is a free log subscription operation binding the contract event 0x592e83ede8c846dffd6e4edcbe77983b153bec820ac697aa34a236147b1341e0.
//
// Solidity: event NetworkRegistered(ownerAddress indexed address, id bytes16, name string)
func (_HoQuStorage *HoQuStorageFilterer) WatchNetworkRegistered(opts *bind.WatchOpts, sink chan<- *HoQuStorageNetworkRegistered, ownerAddress []common.Address) (event.Subscription, error) {

	var ownerAddressRule []interface{}
	for _, ownerAddressItem := range ownerAddress {
		ownerAddressRule = append(ownerAddressRule, ownerAddressItem)
	}

	logs, sub, err := _HoQuStorage.contract.WatchLogs(opts, "NetworkRegistered", ownerAddressRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(HoQuStorageNetworkRegistered)
				if err := _HoQuStorage.contract.UnpackLog(event, "NetworkRegistered", log); err != nil {
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

// HoQuStorageOfferAddedIterator is returned from FilterOfferAdded and is used to iterate over the raw logs and unpacked data for OfferAdded events raised by the HoQuStorage contract.
type HoQuStorageOfferAddedIterator struct {
	Event *HoQuStorageOfferAdded // Event containing the contract specifics and raw log

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
func (it *HoQuStorageOfferAddedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(HoQuStorageOfferAdded)
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
		it.Event = new(HoQuStorageOfferAdded)
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
func (it *HoQuStorageOfferAddedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *HoQuStorageOfferAddedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// HoQuStorageOfferAdded represents a OfferAdded event raised by the HoQuStorage contract.
type HoQuStorageOfferAdded struct {
	OwnerAddress common.Address
	Id           [16]byte
	Name         string
	Raw          types.Log // Blockchain specific contextual infos
}

// FilterOfferAdded is a free log retrieval operation binding the contract event 0x5c9100ad970db426d43d47b3a9639dc42d9732871d99213dcb4984f7e71509d0.
//
// Solidity: event OfferAdded(ownerAddress indexed address, id bytes16, name string)
func (_HoQuStorage *HoQuStorageFilterer) FilterOfferAdded(opts *bind.FilterOpts, ownerAddress []common.Address) (*HoQuStorageOfferAddedIterator, error) {

	var ownerAddressRule []interface{}
	for _, ownerAddressItem := range ownerAddress {
		ownerAddressRule = append(ownerAddressRule, ownerAddressItem)
	}

	logs, sub, err := _HoQuStorage.contract.FilterLogs(opts, "OfferAdded", ownerAddressRule)
	if err != nil {
		return nil, err
	}
	return &HoQuStorageOfferAddedIterator{contract: _HoQuStorage.contract, event: "OfferAdded", logs: logs, sub: sub}, nil
}

// WatchOfferAdded is a free log subscription operation binding the contract event 0x5c9100ad970db426d43d47b3a9639dc42d9732871d99213dcb4984f7e71509d0.
//
// Solidity: event OfferAdded(ownerAddress indexed address, id bytes16, name string)
func (_HoQuStorage *HoQuStorageFilterer) WatchOfferAdded(opts *bind.WatchOpts, sink chan<- *HoQuStorageOfferAdded, ownerAddress []common.Address) (event.Subscription, error) {

	var ownerAddressRule []interface{}
	for _, ownerAddressItem := range ownerAddress {
		ownerAddressRule = append(ownerAddressRule, ownerAddressItem)
	}

	logs, sub, err := _HoQuStorage.contract.WatchLogs(opts, "OfferAdded", ownerAddressRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(HoQuStorageOfferAdded)
				if err := _HoQuStorage.contract.UnpackLog(event, "OfferAdded", log); err != nil {
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

// HoQuStorageStatsChangedIterator is returned from FilterStatsChanged and is used to iterate over the raw logs and unpacked data for StatsChanged events raised by the HoQuStorage contract.
type HoQuStorageStatsChangedIterator struct {
	Event *HoQuStorageStatsChanged // Event containing the contract specifics and raw log

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
func (it *HoQuStorageStatsChangedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(HoQuStorageStatsChanged)
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
		it.Event = new(HoQuStorageStatsChanged)
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
func (it *HoQuStorageStatsChangedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *HoQuStorageStatsChangedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// HoQuStorageStatsChanged represents a StatsChanged event raised by the HoQuStorage contract.
type HoQuStorageStatsChanged struct {
	OwnerAddress common.Address
	Id           [16]byte
	Rating       *big.Int
	Raw          types.Log // Blockchain specific contextual infos
}

// FilterStatsChanged is a free log retrieval operation binding the contract event 0xd2e9a0e1a68c0320c3b58608fcefa40787236bd0d2153ebbd82690b3fa064ad4.
//
// Solidity: event StatsChanged(ownerAddress indexed address, id bytes16, rating uint256)
func (_HoQuStorage *HoQuStorageFilterer) FilterStatsChanged(opts *bind.FilterOpts, ownerAddress []common.Address) (*HoQuStorageStatsChangedIterator, error) {

	var ownerAddressRule []interface{}
	for _, ownerAddressItem := range ownerAddress {
		ownerAddressRule = append(ownerAddressRule, ownerAddressItem)
	}

	logs, sub, err := _HoQuStorage.contract.FilterLogs(opts, "StatsChanged", ownerAddressRule)
	if err != nil {
		return nil, err
	}
	return &HoQuStorageStatsChangedIterator{contract: _HoQuStorage.contract, event: "StatsChanged", logs: logs, sub: sub}, nil
}

// WatchStatsChanged is a free log subscription operation binding the contract event 0xd2e9a0e1a68c0320c3b58608fcefa40787236bd0d2153ebbd82690b3fa064ad4.
//
// Solidity: event StatsChanged(ownerAddress indexed address, id bytes16, rating uint256)
func (_HoQuStorage *HoQuStorageFilterer) WatchStatsChanged(opts *bind.WatchOpts, sink chan<- *HoQuStorageStatsChanged, ownerAddress []common.Address) (event.Subscription, error) {

	var ownerAddressRule []interface{}
	for _, ownerAddressItem := range ownerAddress {
		ownerAddressRule = append(ownerAddressRule, ownerAddressItem)
	}

	logs, sub, err := _HoQuStorage.contract.WatchLogs(opts, "StatsChanged", ownerAddressRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(HoQuStorageStatsChanged)
				if err := _HoQuStorage.contract.UnpackLog(event, "StatsChanged", log); err != nil {
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
