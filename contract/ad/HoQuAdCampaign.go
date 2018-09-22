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
const HoQuAdCampaignABI = "[{\"constant\":true,\"inputs\":[],\"name\":\"status\",\"outputs\":[{\"name\":\"\",\"type\":\"uint8\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"offerId\",\"outputs\":[{\"name\":\"\",\"type\":\"bytes16\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"payerAddress\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_status\",\"type\":\"uint8\"}],\"name\":\"setStatus\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"rater\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"id\",\"type\":\"bytes16\"},{\"name\":\"intermediaryAddress\",\"type\":\"address\"},{\"name\":\"percent\",\"type\":\"uint32\"}],\"name\":\"addLeadIntermediary\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"affiliateId\",\"outputs\":[{\"name\":\"\",\"type\":\"bytes16\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"id\",\"type\":\"bytes16\"},{\"name\":\"_dataUrl\",\"type\":\"string\"}],\"name\":\"setLeadDataUrl\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_payerAddress\",\"type\":\"address\"}],\"name\":\"setPayerAddress\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"storageAddress\",\"type\":\"address\"}],\"name\":\"setStorageAddress\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"id\",\"type\":\"bytes16\"},{\"name\":\"num\",\"type\":\"uint8\"}],\"name\":\"getLeadIntermediaryPercent\",\"outputs\":[{\"name\":\"\",\"type\":\"uint32\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"ownerAddress\",\"type\":\"address\"},{\"name\":\"id\",\"type\":\"bytes16\"}],\"name\":\"addTracker\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"\",\"type\":\"address\"}],\"name\":\"trackers\",\"outputs\":[{\"name\":\"\",\"type\":\"bytes16\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"id\",\"type\":\"bytes16\"},{\"name\":\"num\",\"type\":\"uint8\"}],\"name\":\"getOfferTariffGroup\",\"outputs\":[{\"name\":\"\",\"type\":\"bytes16\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"config\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"id\",\"type\":\"bytes16\"},{\"name\":\"trackerId\",\"type\":\"bytes16\"},{\"name\":\"meta\",\"type\":\"string\"},{\"name\":\"dataUrl\",\"type\":\"string\"},{\"name\":\"price\",\"type\":\"uint256\"}],\"name\":\"addLead\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"transactorAddress\",\"type\":\"address\"}],\"name\":\"setTransactorAddress\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"id\",\"type\":\"bytes16\"},{\"name\":\"_status\",\"type\":\"uint8\"}],\"name\":\"setLeadStatus\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"id\",\"type\":\"bytes16\"}],\"name\":\"transactLead\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"configAddress\",\"type\":\"address\"}],\"name\":\"setConfigAddress\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"transactor\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"store\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"id\",\"type\":\"bytes16\"},{\"name\":\"num\",\"type\":\"uint8\"}],\"name\":\"getUserAddress\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"id\",\"type\":\"bytes16\"},{\"name\":\"_price\",\"type\":\"uint256\"}],\"name\":\"setLeadPrice\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"adId\",\"outputs\":[{\"name\":\"\",\"type\":\"bytes16\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"raterAddress\",\"type\":\"address\"}],\"name\":\"setRaterAddress\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"beneficiaryAddress\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"id\",\"type\":\"bytes16\"},{\"name\":\"num\",\"type\":\"uint8\"}],\"name\":\"getLeadIntermediaryAddress\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"\",\"type\":\"bytes16\"}],\"name\":\"leads\",\"outputs\":[{\"name\":\"createdAt\",\"type\":\"uint256\"},{\"name\":\"trackerId\",\"type\":\"bytes16\"},{\"name\":\"dataUrl\",\"type\":\"string\"},{\"name\":\"meta\",\"type\":\"string\"},{\"name\":\"price\",\"type\":\"uint256\"},{\"name\":\"numOfIntermediaries\",\"type\":\"uint8\"},{\"name\":\"status\",\"type\":\"uint8\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_beneficiaryAddress\",\"type\":\"address\"}],\"name\":\"setBeneficiaryAddress\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"name\":\"configAddress\",\"type\":\"address\"},{\"name\":\"transactorAddress\",\"type\":\"address\"},{\"name\":\"storageAddress\",\"type\":\"address\"},{\"name\":\"raterAddress\",\"type\":\"address\"},{\"name\":\"_adId\",\"type\":\"bytes16\"},{\"name\":\"_offerId\",\"type\":\"bytes16\"},{\"name\":\"_affiliateId\",\"type\":\"bytes16\"},{\"name\":\"_beneficiaryAddress\",\"type\":\"address\"},{\"name\":\"_payerAddress\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"payerAddress\",\"type\":\"address\"},{\"indexed\":false,\"name\":\"newStatus\",\"type\":\"uint8\"}],\"name\":\"StatusChanged\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"beneficiaryAddress\",\"type\":\"address\"},{\"indexed\":true,\"name\":\"newBeneficiaryAddress\",\"type\":\"address\"}],\"name\":\"BeneficiaryAddressChanged\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"payerAddress\",\"type\":\"address\"},{\"indexed\":true,\"name\":\"newPayerAddress\",\"type\":\"address\"}],\"name\":\"PayerAddressChanged\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"beneficiaryAddress\",\"type\":\"address\"},{\"indexed\":false,\"name\":\"id\",\"type\":\"bytes16\"},{\"indexed\":false,\"name\":\"price\",\"type\":\"uint256\"},{\"indexed\":false,\"name\":\"senderAddress\",\"type\":\"address\"}],\"name\":\"LeadAdded\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"beneficiaryAddress\",\"type\":\"address\"},{\"indexed\":false,\"name\":\"id\",\"type\":\"bytes16\"},{\"indexed\":false,\"name\":\"amount\",\"type\":\"uint256\"},{\"indexed\":false,\"name\":\"senderAddress\",\"type\":\"address\"}],\"name\":\"LeadTransacted\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"beneficiaryAddress\",\"type\":\"address\"},{\"indexed\":false,\"name\":\"id\",\"type\":\"bytes16\"},{\"indexed\":false,\"name\":\"status\",\"type\":\"uint8\"}],\"name\":\"LeadStatusChanged\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"beneficiaryAddress\",\"type\":\"address\"},{\"indexed\":false,\"name\":\"id\",\"type\":\"bytes16\"},{\"indexed\":false,\"name\":\"price\",\"type\":\"uint256\"}],\"name\":\"LeadPriceChanged\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"beneficiaryAddress\",\"type\":\"address\"},{\"indexed\":false,\"name\":\"id\",\"type\":\"bytes16\"},{\"indexed\":false,\"name\":\"dataUrl\",\"type\":\"string\"}],\"name\":\"LeadDataUrlChanged\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"ownerAddress\",\"type\":\"address\"},{\"indexed\":false,\"name\":\"id\",\"type\":\"bytes16\"}],\"name\":\"TrackerAdded\",\"type\":\"event\"}]"

// HoQuAdCampaignBin is the compiled bytecode used for deploying new contracts.
const HoQuAdCampaignBin = `0x608060405234801561001057600080fd5b506040516101208061293683398101604090815281516020830151918301516060840151608085015160a086015160c087015160e08801516101009098015160008054600160a060020a03988916600160a060020a031991821617825560018054988a16988216989098179097556002805499891699881699909917909855600380549588169587169590951790945560048054700100000000000000000000000000000000938490048402948490046001608060020a0319918216176001608060020a03169490941790556005805492909104919092161790556006805495841695831695909517909455600780549290941691161760a060020a60ff02191674010000000000000000000000000000000000000000179091556127fb90819061013b90396000f30060806040526004361061017f5763ffffffff7c0100000000000000000000000000000000000000000000000000000000600035041663200d2ed2811461018457806321ce9f91146101bd5780632b2f1e14146101ef5780632e49d78b14610220578063355c9fb11461023d57806336617045146102525780633f8b35a614610289578063406865fc1461029e5780634b28676d1461030657806359b910d6146103275780635cb8f5ab146103485780636aa0082c146103895780636cb234a3146103b757806374f627c5146103d857806379502c55146104005780637a8187cf146104155780637d9da78a146104c557806381e85aa3146104e6578063835aef401461050e57806383a12de9146105305780638ecb72d114610551578063975057e714610566578063c49e2fb11461057b578063c535bd88146105a3578063cbd76c35146105c8578063ce88f04e146105dd578063d9c4870e146105fe578063da7345b714610613578063e7608fb21461063b578063ec6be06e1461077a575b600080fd5b34801561019057600080fd5b5061019961079b565b604051808260058111156101a957fe5b60ff16815260200191505060405180910390f35b3480156101c957600080fd5b506101d26107bc565b604080516001608060020a03199092168252519081900360200190f35b3480156101fb57600080fd5b506102046107cc565b60408051600160a060020a039092168252519081900360200190f35b34801561022c57600080fd5b5061023b60ff600435166107db565b005b34801561024957600080fd5b5061020461096f565b34801561025e57600080fd5b5061023b6001608060020a031960043516600160a060020a036024351663ffffffff6044351661097e565b34801561029557600080fd5b506101d2610aeb565b3480156102aa57600080fd5b5060408051602060046024803582810135601f810185900485028601850190965285855261023b9583356001608060020a031916953695604494919390910191908190840183828082843750949750610af79650505050505050565b34801561031257600080fd5b5061023b600160a060020a0360043516610cbf565b34801561033357600080fd5b5061023b600160a060020a0360043516610da3565b34801561035457600080fd5b506103706001608060020a03196004351660ff60243516610e4d565b6040805163ffffffff9092168252519081900360200190f35b34801561039557600080fd5b5061023b600160a060020a03600435166001608060020a031960243516610ec4565b3480156103c357600080fd5b506101d2600160a060020a0360043516610fb8565b3480156103e457600080fd5b506101d26001608060020a03196004351660ff60243516610fd0565b34801561040c57600080fd5b5061020461107b565b34801561042157600080fd5b50604080516020600460443581810135601f810184900484028501840190955284845261023b9482356001608060020a03199081169560248035909216953695946064949293019190819084018382808284375050604080516020601f89358b018035918201839004830284018301909452808352979a999881019791965091820194509250829150840183828082843750949750509335945061108a9350505050565b3480156104d157600080fd5b5061023b600160a060020a036004351661139a565b3480156104f257600080fd5b5061023b6001608060020a03196004351660ff60243516611444565b34801561051a57600080fd5b5061023b6001608060020a0319600435166115c2565b34801561053c57600080fd5b5061023b600160a060020a0360043516611cc5565b34801561055d57600080fd5b50610204611d6f565b34801561057257600080fd5b50610204611d7e565b34801561058757600080fd5b506102046001608060020a03196004351660ff60243516611d8d565b3480156105af57600080fd5b5061023b6001608060020a031960043516602435611e05565b3480156105d457600080fd5b506101d2611f48565b3480156105e957600080fd5b5061023b600160a060020a0360043516611f54565b34801561060a57600080fd5b50610204611ffe565b34801561061f57600080fd5b506102046001608060020a03196004351660ff6024351661200d565b34801561064757600080fd5b5061065d6001608060020a031960043516612086565b604080518881526001608060020a0319881660208201526080810185905260ff841660a08201529081016060820160c0830184600581111561069b57fe5b60ff168152602001838103835288818151815260200191508051906020019080838360005b838110156106d85781810151838201526020016106c0565b50505050905090810190601f1680156107055780820380516001836020036101000a031916815260200191505b50838103825287518152875160209182019189019080838360005b83811015610738578181015183820152602001610720565b50505050905090810190601f1680156107655780820380516001836020036101000a031916815260200191505b50995050505050505050505060405180910390f35b34801561078657600080fd5b5061023b600160a060020a03600435166121e2565b60075474010000000000000000000000000000000000000000900460ff1681565b600454608060020a908190040281565b600754600160a060020a031681565b6107e36126ce565b600080546040805160e060020a63babcc5390281523360048201529051600160a060020a039092169263babcc539926024808401936020939083900390910190829087803b15801561083457600080fd5b505af1158015610848573d6000803e3d6000fd5b505050506040513d602081101561085e57600080fd5b5051151561086b57600080fd5b6007805483919074ff00000000000000000000000000000000000000001916740100000000000000000000000000000000000000008360058111156108ac57fe5b02179055506004546108c390608060020a026122c6565b90506000816080015160058111156108d757fe5b14156108e257600080fd5b608081018260058111156108f257fe5b908160058111156108ff57fe5b90525060045461091590608060020a02826123eb565b600754604051600160a060020a03909116907f76c0bad3b7f1c74587dcb478154d4c87ae6663db97cc086b981decf4e35a797c9084908082600581111561095857fe5b60ff16815260200191505060405180910390a25050565b600354600160a060020a031681565b600080546040805160e060020a63babcc5390281523360048201529051600160a060020a039092169263babcc539926024808401936020939083900390910190829087803b1580156109cf57600080fd5b505af11580156109e3573d6000803e3d6000fd5b505050506040513d60208110156109f957600080fd5b505180610a0d575033600052600960205260015b1515610a1857600080fd5b60006001608060020a03198416600090815260086020526040902060070154610100900460ff166005811115610a4a57fe5b1415610a5557600080fd5b6001608060020a031992909216600081815260086020818152604080842060078101805460ff90811687526005830185528387208054600160a060020a031916600160a060020a039a909a169990991790985580548816865260069091018352908420805463ffffffff191663ffffffff90981697909717909655929091529052815460ff198116908216600101909116179055565b600554608060020a0281565b600080546040805160e060020a63babcc5390281523360048201529051600160a060020a039092169263babcc539926024808401936020939083900390910190829087803b158015610b4857600080fd5b505af1158015610b5c573d6000803e3d6000fd5b505050506040513d6020811015610b7257600080fd5b505180610b86575033600052600960205260015b1515610b9157600080fd5b60006001608060020a03198316600090815260086020526040902060070154610100900460ff166005811115610bc357fe5b1415610bce57600080fd5b6001608060020a0319821660009081526008602090815260409091208251610bfe926002909201918401906126fd565b50600654604080516001608060020a03198516815260208082018381528551938301939093528451600160a060020a03909416937fb7e79da6d3cc03db8c06d5600b8437e0ce2f7352f871a1a1729dd3f91ceaca6493879387939092606084019185019080838360005b83811015610c80578181015183820152602001610c68565b50505050905090810190601f168015610cad5780820380516001836020036101000a031916815260200191505b50935050505060405180910390a25050565b600080546040805160e060020a63babcc5390281523360048201529051600160a060020a039092169263babcc539926024808401936020939083900390910190829087803b158015610d1057600080fd5b505af1158015610d24573d6000803e3d6000fd5b505050506040513d6020811015610d3a57600080fd5b50511515610d4757600080fd5b600754604051600160a060020a038084169216907f650742bb22425d77490f0e51db9824c857c049c9fc9fee93b369553acb17275890600090a360078054600160a060020a031916600160a060020a0392909216919091179055565b600080546040805160e060020a63babcc5390281523360048201529051600160a060020a039092169263babcc539926024808401936020939083900390910190829087803b158015610df457600080fd5b505af1158015610e08573d6000803e3d6000fd5b505050506040513d6020811015610e1e57600080fd5b50511515610e2b57600080fd5b60018054600160a060020a031916600160a060020a0392909216919091179055565b6000806001608060020a03198416600090815260086020526040902060070154610100900460ff166005811115610e8057fe5b1415610e8b57600080fd5b506001608060020a03198216600090815260086020908152604080832060ff8516845260060190915290205463ffffffff165b92915050565b600080546040805160e060020a63babcc5390281523360048201529051600160a060020a039092169263babcc539926024808401936020939083900390910190829087803b158015610f1557600080fd5b505af1158015610f29573d6000803e3d6000fd5b505050506040513d6020811015610f3f57600080fd5b50511515610f4c57600080fd5b600160a060020a0382166000818152600960209081526040918290208054608060020a86046001608060020a0319918216179091558251908516815291517fd7793ff1544e44df2a307af241d282820620c015d2e3b8cb17b480a7e47a48be9281900390910190a25050565b600960205260009081526040902054608060020a0281565b600154604080517f74f627c50000000000000000000000000000000000000000000000000000000081526001608060020a03198516600482015260ff841660248201529051600092600160a060020a0316916374f627c591604480830192602092919082900301818787803b15801561104857600080fd5b505af115801561105c573d6000803e3d6000fd5b505050506040513d602081101561107257600080fd5b50519392505050565b600054600160a060020a031681565b61109261277b565b600080546040805160e060020a63babcc5390281523360048201529051600160a060020a039092169263babcc539926024808401936020939083900390910190829087803b1580156110e357600080fd5b505af11580156110f7573d6000803e3d6000fd5b505050506040513d602081101561110d57600080fd5b505180611121575033600052600960205260015b151561112c57600080fd5b60006001608060020a03198716600090815260086020526040902060070154610100900460ff16600581111561115e57fe5b1461116857600080fd5b611171856124c0565b905060008160a00151600581111561118557fe5b141561119057600080fd5b6040805160e0810182524281526001608060020a0319871660208201529081018490526060810185905260808101839052600060a082015260c08101600190526001608060020a03198781166000908152600860209081526040918290208451815584820151600182018054909516608060020a9091041790935590830151805161122192600285019201906126fd565b506060820151805161123d9160038401916020909101906126fd565b506080820151600482015560a082015160078201805460ff191660ff9092169190911780825560c0840151919061ff00191661010083600581111561127e57fe5b02179055505060035460048054600554604080517fc5f6dd850000000000000000000000000000000000000000000000000000000081526001608060020a0319608060020a948590048502811695820195909552848c166024820152919092029290921660448301526064820186905251600160a060020a03909216925063c5f6dd8591608480830192600092919082900301818387803b15801561132257600080fd5b505af1158015611336573d6000803e3d6000fd5b5050600654604080516001608060020a03198b1681526020810187905233818301529051600160a060020a0390921693507f1db4d44ca5cc48741d4e28698bbf31e8462db89dce45359aaa81393f41e9c4da925081900360600190a2505050505050565b600080546040805160e060020a63babcc5390281523360048201529051600160a060020a039092169263babcc539926024808401936020939083900390910190829087803b1580156113eb57600080fd5b505af11580156113ff573d6000803e3d6000fd5b505050506040513d602081101561141557600080fd5b5051151561142257600080fd5b60028054600160a060020a031916600160a060020a0392909216919091179055565b600080546040805160e060020a63babcc5390281523360048201529051600160a060020a039092169263babcc539926024808401936020939083900390910190829087803b15801561149557600080fd5b505af11580156114a9573d6000803e3d6000fd5b505050506040513d60208110156114bf57600080fd5b5051806114d3575033600052600960205260015b15156114de57600080fd5b60006001608060020a03198316600090815260086020526040902060070154610100900460ff16600581111561151057fe5b141561151b57600080fd5b6001608060020a031982166000908152600860205260409020600701805482919061ff00191661010083600581111561155057fe5b02179055506006546040516001608060020a031984168152600160a060020a03909116907ffa2a86784442baf24e238559b3ca5986aa4a588eeb7840c5ba8b27d6e663ecd99084908490602081018260058111156115aa57fe5b60ff1681526020019250505060405180910390a25050565b600080546040805160e060020a63babcc5390281523360048201529051839283928392839283928392600160a060020a039092169163babcc5399160248082019260209290919082900301818787803b15801561161e57600080fd5b505af1158015611632573d6000803e3d6000fd5b505050506040513d602081101561164857600080fd5b50518061165c575033600052600960205260015b151561166757600080fd5b60046001608060020a03198916600090815260086020526040902060070154610100900460ff16600581111561169957fe5b141580156116d7575060056001608060020a03198916600090815260086020526040902060070154610100900460ff1660058111156116d457fe5b14155b15156116e257600080fd5b6001608060020a031988166000908152600860205260408120600401541161170957600080fd5b6001608060020a03198816600090815260086020908152604080832060078101805461ff001916610400179055835482517fe14891910000000000000000000000000000000000000000000000000000000081529251919b506117ec94670de0b6b3a7640000946117e094600160a060020a039093169363e148919193600480820194918390030190829087803b1580156117a357600080fd5b505af11580156117b7573d6000803e3d6000fd5b505050506040513d60208110156117cd57600080fd5b505160048b01549063ffffffff61267e16565b9063ffffffff6126a716565b6004880154909650611804908763ffffffff6126bc16565b6002546007546004808b0154604080517ff3fef3a3000000000000000000000000000000000000000000000000000000008152600160a060020a039485169381019390935260248301919091525193985091169163f3fef3a39160448082019260009290919082900301818387803b15801561187f57600080fd5b505af1158015611893573d6000803e3d6000fd5b505060025460008054604080517f97c0262a0000000000000000000000000000000000000000000000000000000081529051600160a060020a03948516965063d0679d34955093909116926397c0262a92600480840193602093929083900390910190829087803b15801561190757600080fd5b505af115801561191b573d6000803e3d6000fd5b505050506040513d602081101561193157600080fd5b5051604080517c010000000000000000000000000000000000000000000000000000000063ffffffff8516028152600160a060020a039092166004830152602482018a905251604480830192600092919082900301818387803b15801561199757600080fd5b505af11580156119ab573d6000803e3d6000fd5b50505050600093505b600787015460ff9081169085161015611ac55760ff8416600090815260058801602090815260408083205460068b01909252909120546004890154600160a060020a03909216945063ffffffff9081169350611a1d916305f5e100916117e09190869061267e16565b600254604080517fd0679d34000000000000000000000000000000000000000000000000000000008152600160a060020a03878116600483015260248201859052915193945091169163d0679d349160448082019260009290919082900301818387803b158015611a8d57600080fd5b505af1158015611aa1573d6000803e3d6000fd5b50505050611ab881866126bc90919063ffffffff16565b94506001909301926119b4565b600254600654604080517fd0679d34000000000000000000000000000000000000000000000000000000008152600160a060020a039283166004820152602481018990529051919092169163d0679d3491604480830192600092919082900301818387803b158015611b3657600080fd5b505af1158015611b4a573d6000803e3d6000fd5b50506003546004805460018c01546005548d840154604080517f4943df380000000000000000000000000000000000000000000000000000000081526001608060020a0319608060020a96879004870281169782019790975293850286166024850152939091029093166044820152606481019290925251600160a060020a039092169350634943df38925060848082019260009290919082900301818387803b158015611bf757600080fd5b505af1158015611c0b573d6000803e3d6000fd5b5050600654604080516001608060020a03198d168152602081018a905233818301529051600160a060020a0390921693507f9ba12d15e5467b3340747345a4d3d5016d839166ebd348e1548ce30e67140fe0925081900360600190a2600654604080516001608060020a03198b168152600460208201528151600160a060020a03909316927ffa2a86784442baf24e238559b3ca5986aa4a588eeb7840c5ba8b27d6e663ecd9929181900390910190a25050505050505050565b600080546040805160e060020a63babcc5390281523360048201529051600160a060020a039092169263babcc539926024808401936020939083900390910190829087803b158015611d1657600080fd5b505af1158015611d2a573d6000803e3d6000fd5b505050506040513d6020811015611d4057600080fd5b50511515611d4d57600080fd5b60008054600160a060020a031916600160a060020a0392909216919091179055565b600254600160a060020a031681565b600154600160a060020a031681565b600154604080517fc49e2fb10000000000000000000000000000000000000000000000000000000081526001608060020a03198516600482015260ff841660248201529051600092600160a060020a03169163c49e2fb191604480830192602092919082900301818787803b15801561104857600080fd5b600080546040805160e060020a63babcc5390281523360048201529051600160a060020a039092169263babcc539926024808401936020939083900390910190829087803b158015611e5657600080fd5b505af1158015611e6a573d6000803e3d6000fd5b505050506040513d6020811015611e8057600080fd5b505180611e94575033600052600960205260015b1515611e9f57600080fd5b60006001608060020a03198316600090815260086020526040902060070154610100900460ff166005811115611ed157fe5b1415611edc57600080fd5b6001608060020a0319821660008181526008602090815260409182902060040184905560065482519384529083018490528151600160a060020a03909116927f9337eccc933bde4853c11d1f8918ffe273f426a095be22c9c2e3ba7087dd141692908290030190a25050565b600454608060020a0281565b600080546040805160e060020a63babcc5390281523360048201529051600160a060020a039092169263babcc539926024808401936020939083900390910190829087803b158015611fa557600080fd5b505af1158015611fb9573d6000803e3d6000fd5b505050506040513d6020811015611fcf57600080fd5b50511515611fdc57600080fd5b60038054600160a060020a031916600160a060020a0392909216919091179055565b600654600160a060020a031681565b6000806001608060020a03198416600090815260086020526040902060070154610100900460ff16600581111561204057fe5b141561204b57600080fd5b506001608060020a03198216600090815260086020908152604080832060ff85168452600501909152902054600160a060020a031692915050565b60086020908152600091825260409182902080546001808301546002808501805488516101009582161595909502600019011691909104601f81018790048702840187019097528683529295608060020a9091029491929183018282801561212f5780601f106121045761010080835404028352916020019161212f565b820191906000526020600020905b81548152906001019060200180831161211257829003601f168201915b5050505060038301805460408051602060026001851615610100026000190190941693909304601f81018490048402820184019092528181529495949350908301828280156121bf5780601f10612194576101008083540402835291602001916121bf565b820191906000526020600020905b8154815290600101906020018083116121a257829003601f168201915b50505050600483015460079093015491929160ff80821692506101009091041687565b600080546040805160e060020a63babcc5390281523360048201529051600160a060020a039092169263babcc539926024808401936020939083900390910190829087803b15801561223357600080fd5b505af1158015612247573d6000803e3d6000fd5b505050506040513d602081101561225d57600080fd5b5051151561226a57600080fd5b600654604051600160a060020a038084169216907fb53b7057ddbff0cd9111b0a1a501a9cc1a9c315eb3396cb0d47917908325999a90600090a360068054600160a060020a031916600160a060020a0392909216919091179055565b6122ce6126ce565b6122d66126ce565b600154604080517f890755570000000000000000000000000000000000000000000000000000000081526001608060020a0319861660048201529051600160a060020a039092169163890755579160248082019260a0929091908290030181600087803b15801561234657600080fd5b505af115801561235a573d6000803e3d6000fd5b505050506040513d60a081101561237057600080fd5b50805160208083015160408085015160608087015160809788015196979496929590949293899391840192908401919084019084018560058111156123b157fe5b60058111156123bc57fe5b905294909452600160a060020a03909416909252506001608060020a0319928316909152911690529050919050565b60015481516020830151604080850151608086015191517fa0191d200000000000000000000000000000000000000000000000000000000081526001608060020a03198089166004830190815281871660248401529085166044830152600160a060020a0383811660648401529096169563a0191d20958995909490939290919060840182600581111561247b57fe5b60ff16815260200195505050505050600060405180830381600087803b1580156124a457600080fd5b505af11580156124b8573d6000803e3d6000fd5b505050505050565b6124c861277b565b6124d061277b565b600154604080517fea2fdd8c0000000000000000000000000000000000000000000000000000000081526001608060020a0319861660048201529051600160a060020a039092169163ea2fdd8c9160248082019260009290919082900301818387803b15801561253f57600080fd5b505af1158015612553573d6000803e3d6000fd5b505050506040513d6000823e601f3d908101601f1916820160405260c081101561257c57600080fd5b81516020830151604084018051929491938201926401000000008111156125a257600080fd5b820160208101848111156125b557600080fd5b81516401000000008111828201871017156125cf57600080fd5b505092919060200180516401000000008111156125eb57600080fd5b820160208101848111156125fe57600080fd5b815164010000000081118282018710171561261857600080fd5b50506020808301516040938401519295509350909188918201908201606083016080840160a0850186600581111561264c57fe5b600581111561265757fe5b90529590955294909352939092526001608060020a03199384169052509116905292915050565b600082151561268f57506000610ebe565b5081810281838281151561269f57fe5b0414610ebe57fe5b600081838115156126b457fe5b049392505050565b6000828211156126c857fe5b50900390565b6040805160a081018252600080825260208201819052918101829052606081018290529060808201905b905290565b828054600181600116156101000203166002900490600052602060002090601f016020900481019282601f1061273e57805160ff191683800117855561276b565b8280016001018555821561276b579182015b8281111561276b578251825591602001919060010190612750565b506127779291506127b2565b5090565b6040805160c081018252600080825260208201819052606092820183905282820192909252608081018290529060a08201906126f8565b6127cc91905b8082111561277757600081556001016127b8565b905600a165627a7a723058206df8600e0945f2ddb2e6f4a39e8d21c97fa880b35b2370b0424f1b56044db1970029`

// DeployHoQuAdCampaign deploys a new Ethereum contract, binding an instance of HoQuAdCampaign to it.
func DeployHoQuAdCampaign(auth *bind.TransactOpts, backend bind.ContractBackend, configAddress common.Address, transactorAddress common.Address, storageAddress common.Address, raterAddress common.Address, _adId [16]byte, _offerId [16]byte, _affiliateId [16]byte, _beneficiaryAddress common.Address, _payerAddress common.Address) (common.Address, *types.Transaction, *HoQuAdCampaign, error) {
	parsed, err := abi.JSON(strings.NewReader(HoQuAdCampaignABI))
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	address, tx, contract, err := bind.DeployContract(auth, parsed, common.FromHex(HoQuAdCampaignBin), backend, configAddress, transactorAddress, storageAddress, raterAddress, _adId, _offerId, _affiliateId, _beneficiaryAddress, _payerAddress)
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

// GetOfferTariffGroup is a free data retrieval call binding the contract method 0x74f627c5.
//
// Solidity: function getOfferTariffGroup(id bytes16, num uint8) constant returns(bytes16)
func (_HoQuAdCampaign *HoQuAdCampaignCaller) GetOfferTariffGroup(opts *bind.CallOpts, id [16]byte, num uint8) ([16]byte, error) {
	var (
		ret0 = new([16]byte)
	)
	out := ret0
	err := _HoQuAdCampaign.contract.Call(opts, out, "getOfferTariffGroup", id, num)
	return *ret0, err
}

// GetOfferTariffGroup is a free data retrieval call binding the contract method 0x74f627c5.
//
// Solidity: function getOfferTariffGroup(id bytes16, num uint8) constant returns(bytes16)
func (_HoQuAdCampaign *HoQuAdCampaignSession) GetOfferTariffGroup(id [16]byte, num uint8) ([16]byte, error) {
	return _HoQuAdCampaign.Contract.GetOfferTariffGroup(&_HoQuAdCampaign.CallOpts, id, num)
}

// GetOfferTariffGroup is a free data retrieval call binding the contract method 0x74f627c5.
//
// Solidity: function getOfferTariffGroup(id bytes16, num uint8) constant returns(bytes16)
func (_HoQuAdCampaign *HoQuAdCampaignCallerSession) GetOfferTariffGroup(id [16]byte, num uint8) ([16]byte, error) {
	return _HoQuAdCampaign.Contract.GetOfferTariffGroup(&_HoQuAdCampaign.CallOpts, id, num)
}

// GetUserAddress is a free data retrieval call binding the contract method 0xc49e2fb1.
//
// Solidity: function getUserAddress(id bytes16, num uint8) constant returns(address)
func (_HoQuAdCampaign *HoQuAdCampaignCaller) GetUserAddress(opts *bind.CallOpts, id [16]byte, num uint8) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _HoQuAdCampaign.contract.Call(opts, out, "getUserAddress", id, num)
	return *ret0, err
}

// GetUserAddress is a free data retrieval call binding the contract method 0xc49e2fb1.
//
// Solidity: function getUserAddress(id bytes16, num uint8) constant returns(address)
func (_HoQuAdCampaign *HoQuAdCampaignSession) GetUserAddress(id [16]byte, num uint8) (common.Address, error) {
	return _HoQuAdCampaign.Contract.GetUserAddress(&_HoQuAdCampaign.CallOpts, id, num)
}

// GetUserAddress is a free data retrieval call binding the contract method 0xc49e2fb1.
//
// Solidity: function getUserAddress(id bytes16, num uint8) constant returns(address)
func (_HoQuAdCampaign *HoQuAdCampaignCallerSession) GetUserAddress(id [16]byte, num uint8) (common.Address, error) {
	return _HoQuAdCampaign.Contract.GetUserAddress(&_HoQuAdCampaign.CallOpts, id, num)
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

// Transactor is a free data retrieval call binding the contract method 0x8ecb72d1.
//
// Solidity: function transactor() constant returns(address)
func (_HoQuAdCampaign *HoQuAdCampaignCaller) Transactor(opts *bind.CallOpts) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _HoQuAdCampaign.contract.Call(opts, out, "transactor")
	return *ret0, err
}

// Transactor is a free data retrieval call binding the contract method 0x8ecb72d1.
//
// Solidity: function transactor() constant returns(address)
func (_HoQuAdCampaign *HoQuAdCampaignSession) Transactor() (common.Address, error) {
	return _HoQuAdCampaign.Contract.Transactor(&_HoQuAdCampaign.CallOpts)
}

// Transactor is a free data retrieval call binding the contract method 0x8ecb72d1.
//
// Solidity: function transactor() constant returns(address)
func (_HoQuAdCampaign *HoQuAdCampaignCallerSession) Transactor() (common.Address, error) {
	return _HoQuAdCampaign.Contract.Transactor(&_HoQuAdCampaign.CallOpts)
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

// SetLeadDataUrl is a paid mutator transaction binding the contract method 0x406865fc.
//
// Solidity: function setLeadDataUrl(id bytes16, _dataUrl string) returns()
func (_HoQuAdCampaign *HoQuAdCampaignTransactor) SetLeadDataUrl(opts *bind.TransactOpts, id [16]byte, _dataUrl string) (*types.Transaction, error) {
	return _HoQuAdCampaign.contract.Transact(opts, "setLeadDataUrl", id, _dataUrl)
}

// SetLeadDataUrl is a paid mutator transaction binding the contract method 0x406865fc.
//
// Solidity: function setLeadDataUrl(id bytes16, _dataUrl string) returns()
func (_HoQuAdCampaign *HoQuAdCampaignSession) SetLeadDataUrl(id [16]byte, _dataUrl string) (*types.Transaction, error) {
	return _HoQuAdCampaign.Contract.SetLeadDataUrl(&_HoQuAdCampaign.TransactOpts, id, _dataUrl)
}

// SetLeadDataUrl is a paid mutator transaction binding the contract method 0x406865fc.
//
// Solidity: function setLeadDataUrl(id bytes16, _dataUrl string) returns()
func (_HoQuAdCampaign *HoQuAdCampaignTransactorSession) SetLeadDataUrl(id [16]byte, _dataUrl string) (*types.Transaction, error) {
	return _HoQuAdCampaign.Contract.SetLeadDataUrl(&_HoQuAdCampaign.TransactOpts, id, _dataUrl)
}

// SetLeadPrice is a paid mutator transaction binding the contract method 0xc535bd88.
//
// Solidity: function setLeadPrice(id bytes16, _price uint256) returns()
func (_HoQuAdCampaign *HoQuAdCampaignTransactor) SetLeadPrice(opts *bind.TransactOpts, id [16]byte, _price *big.Int) (*types.Transaction, error) {
	return _HoQuAdCampaign.contract.Transact(opts, "setLeadPrice", id, _price)
}

// SetLeadPrice is a paid mutator transaction binding the contract method 0xc535bd88.
//
// Solidity: function setLeadPrice(id bytes16, _price uint256) returns()
func (_HoQuAdCampaign *HoQuAdCampaignSession) SetLeadPrice(id [16]byte, _price *big.Int) (*types.Transaction, error) {
	return _HoQuAdCampaign.Contract.SetLeadPrice(&_HoQuAdCampaign.TransactOpts, id, _price)
}

// SetLeadPrice is a paid mutator transaction binding the contract method 0xc535bd88.
//
// Solidity: function setLeadPrice(id bytes16, _price uint256) returns()
func (_HoQuAdCampaign *HoQuAdCampaignTransactorSession) SetLeadPrice(id [16]byte, _price *big.Int) (*types.Transaction, error) {
	return _HoQuAdCampaign.Contract.SetLeadPrice(&_HoQuAdCampaign.TransactOpts, id, _price)
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

// SetTransactorAddress is a paid mutator transaction binding the contract method 0x7d9da78a.
//
// Solidity: function setTransactorAddress(transactorAddress address) returns()
func (_HoQuAdCampaign *HoQuAdCampaignTransactor) SetTransactorAddress(opts *bind.TransactOpts, transactorAddress common.Address) (*types.Transaction, error) {
	return _HoQuAdCampaign.contract.Transact(opts, "setTransactorAddress", transactorAddress)
}

// SetTransactorAddress is a paid mutator transaction binding the contract method 0x7d9da78a.
//
// Solidity: function setTransactorAddress(transactorAddress address) returns()
func (_HoQuAdCampaign *HoQuAdCampaignSession) SetTransactorAddress(transactorAddress common.Address) (*types.Transaction, error) {
	return _HoQuAdCampaign.Contract.SetTransactorAddress(&_HoQuAdCampaign.TransactOpts, transactorAddress)
}

// SetTransactorAddress is a paid mutator transaction binding the contract method 0x7d9da78a.
//
// Solidity: function setTransactorAddress(transactorAddress address) returns()
func (_HoQuAdCampaign *HoQuAdCampaignTransactorSession) SetTransactorAddress(transactorAddress common.Address) (*types.Transaction, error) {
	return _HoQuAdCampaign.Contract.SetTransactorAddress(&_HoQuAdCampaign.TransactOpts, transactorAddress)
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
	SenderAddress      common.Address
	Raw                types.Log // Blockchain specific contextual infos
}

// FilterLeadAdded is a free log retrieval operation binding the contract event 0x1db4d44ca5cc48741d4e28698bbf31e8462db89dce45359aaa81393f41e9c4da.
//
// Solidity: event LeadAdded(beneficiaryAddress indexed address, id bytes16, price uint256, senderAddress address)
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

// WatchLeadAdded is a free log subscription operation binding the contract event 0x1db4d44ca5cc48741d4e28698bbf31e8462db89dce45359aaa81393f41e9c4da.
//
// Solidity: event LeadAdded(beneficiaryAddress indexed address, id bytes16, price uint256, senderAddress address)
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

// HoQuAdCampaignLeadDataUrlChangedIterator is returned from FilterLeadDataUrlChanged and is used to iterate over the raw logs and unpacked data for LeadDataUrlChanged events raised by the HoQuAdCampaign contract.
type HoQuAdCampaignLeadDataUrlChangedIterator struct {
	Event *HoQuAdCampaignLeadDataUrlChanged // Event containing the contract specifics and raw log

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
func (it *HoQuAdCampaignLeadDataUrlChangedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(HoQuAdCampaignLeadDataUrlChanged)
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
		it.Event = new(HoQuAdCampaignLeadDataUrlChanged)
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
func (it *HoQuAdCampaignLeadDataUrlChangedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *HoQuAdCampaignLeadDataUrlChangedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// HoQuAdCampaignLeadDataUrlChanged represents a LeadDataUrlChanged event raised by the HoQuAdCampaign contract.
type HoQuAdCampaignLeadDataUrlChanged struct {
	BeneficiaryAddress common.Address
	Id                 [16]byte
	DataUrl            string
	Raw                types.Log // Blockchain specific contextual infos
}

// FilterLeadDataUrlChanged is a free log retrieval operation binding the contract event 0xb7e79da6d3cc03db8c06d5600b8437e0ce2f7352f871a1a1729dd3f91ceaca64.
//
// Solidity: event LeadDataUrlChanged(beneficiaryAddress indexed address, id bytes16, dataUrl string)
func (_HoQuAdCampaign *HoQuAdCampaignFilterer) FilterLeadDataUrlChanged(opts *bind.FilterOpts, beneficiaryAddress []common.Address) (*HoQuAdCampaignLeadDataUrlChangedIterator, error) {

	var beneficiaryAddressRule []interface{}
	for _, beneficiaryAddressItem := range beneficiaryAddress {
		beneficiaryAddressRule = append(beneficiaryAddressRule, beneficiaryAddressItem)
	}

	logs, sub, err := _HoQuAdCampaign.contract.FilterLogs(opts, "LeadDataUrlChanged", beneficiaryAddressRule)
	if err != nil {
		return nil, err
	}
	return &HoQuAdCampaignLeadDataUrlChangedIterator{contract: _HoQuAdCampaign.contract, event: "LeadDataUrlChanged", logs: logs, sub: sub}, nil
}

// WatchLeadDataUrlChanged is a free log subscription operation binding the contract event 0xb7e79da6d3cc03db8c06d5600b8437e0ce2f7352f871a1a1729dd3f91ceaca64.
//
// Solidity: event LeadDataUrlChanged(beneficiaryAddress indexed address, id bytes16, dataUrl string)
func (_HoQuAdCampaign *HoQuAdCampaignFilterer) WatchLeadDataUrlChanged(opts *bind.WatchOpts, sink chan<- *HoQuAdCampaignLeadDataUrlChanged, beneficiaryAddress []common.Address) (event.Subscription, error) {

	var beneficiaryAddressRule []interface{}
	for _, beneficiaryAddressItem := range beneficiaryAddress {
		beneficiaryAddressRule = append(beneficiaryAddressRule, beneficiaryAddressItem)
	}

	logs, sub, err := _HoQuAdCampaign.contract.WatchLogs(opts, "LeadDataUrlChanged", beneficiaryAddressRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(HoQuAdCampaignLeadDataUrlChanged)
				if err := _HoQuAdCampaign.contract.UnpackLog(event, "LeadDataUrlChanged", log); err != nil {
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

// HoQuAdCampaignLeadPriceChangedIterator is returned from FilterLeadPriceChanged and is used to iterate over the raw logs and unpacked data for LeadPriceChanged events raised by the HoQuAdCampaign contract.
type HoQuAdCampaignLeadPriceChangedIterator struct {
	Event *HoQuAdCampaignLeadPriceChanged // Event containing the contract specifics and raw log

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
func (it *HoQuAdCampaignLeadPriceChangedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(HoQuAdCampaignLeadPriceChanged)
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
		it.Event = new(HoQuAdCampaignLeadPriceChanged)
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
func (it *HoQuAdCampaignLeadPriceChangedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *HoQuAdCampaignLeadPriceChangedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// HoQuAdCampaignLeadPriceChanged represents a LeadPriceChanged event raised by the HoQuAdCampaign contract.
type HoQuAdCampaignLeadPriceChanged struct {
	BeneficiaryAddress common.Address
	Id                 [16]byte
	Price              *big.Int
	Raw                types.Log // Blockchain specific contextual infos
}

// FilterLeadPriceChanged is a free log retrieval operation binding the contract event 0x9337eccc933bde4853c11d1f8918ffe273f426a095be22c9c2e3ba7087dd1416.
//
// Solidity: event LeadPriceChanged(beneficiaryAddress indexed address, id bytes16, price uint256)
func (_HoQuAdCampaign *HoQuAdCampaignFilterer) FilterLeadPriceChanged(opts *bind.FilterOpts, beneficiaryAddress []common.Address) (*HoQuAdCampaignLeadPriceChangedIterator, error) {

	var beneficiaryAddressRule []interface{}
	for _, beneficiaryAddressItem := range beneficiaryAddress {
		beneficiaryAddressRule = append(beneficiaryAddressRule, beneficiaryAddressItem)
	}

	logs, sub, err := _HoQuAdCampaign.contract.FilterLogs(opts, "LeadPriceChanged", beneficiaryAddressRule)
	if err != nil {
		return nil, err
	}
	return &HoQuAdCampaignLeadPriceChangedIterator{contract: _HoQuAdCampaign.contract, event: "LeadPriceChanged", logs: logs, sub: sub}, nil
}

// WatchLeadPriceChanged is a free log subscription operation binding the contract event 0x9337eccc933bde4853c11d1f8918ffe273f426a095be22c9c2e3ba7087dd1416.
//
// Solidity: event LeadPriceChanged(beneficiaryAddress indexed address, id bytes16, price uint256)
func (_HoQuAdCampaign *HoQuAdCampaignFilterer) WatchLeadPriceChanged(opts *bind.WatchOpts, sink chan<- *HoQuAdCampaignLeadPriceChanged, beneficiaryAddress []common.Address) (event.Subscription, error) {

	var beneficiaryAddressRule []interface{}
	for _, beneficiaryAddressItem := range beneficiaryAddress {
		beneficiaryAddressRule = append(beneficiaryAddressRule, beneficiaryAddressItem)
	}

	logs, sub, err := _HoQuAdCampaign.contract.WatchLogs(opts, "LeadPriceChanged", beneficiaryAddressRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(HoQuAdCampaignLeadPriceChanged)
				if err := _HoQuAdCampaign.contract.UnpackLog(event, "LeadPriceChanged", log); err != nil {
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
	SenderAddress      common.Address
	Raw                types.Log // Blockchain specific contextual infos
}

// FilterLeadTransacted is a free log retrieval operation binding the contract event 0x9ba12d15e5467b3340747345a4d3d5016d839166ebd348e1548ce30e67140fe0.
//
// Solidity: event LeadTransacted(beneficiaryAddress indexed address, id bytes16, amount uint256, senderAddress address)
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

// WatchLeadTransacted is a free log subscription operation binding the contract event 0x9ba12d15e5467b3340747345a4d3d5016d839166ebd348e1548ce30e67140fe0.
//
// Solidity: event LeadTransacted(beneficiaryAddress indexed address, id bytes16, amount uint256, senderAddress address)
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
const HoQuAdCampaignIABI = "[{\"constant\":false,\"inputs\":[{\"name\":\"id\",\"type\":\"bytes16\"},{\"name\":\"intermediaryAddress\",\"type\":\"address\"},{\"name\":\"percent\",\"type\":\"uint32\"}],\"name\":\"addLeadIntermediary\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"id\",\"type\":\"bytes16\"},{\"name\":\"_dataUrl\",\"type\":\"string\"}],\"name\":\"setLeadDataUrl\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"id\",\"type\":\"bytes16\"},{\"name\":\"trackerId\",\"type\":\"bytes16\"},{\"name\":\"meta\",\"type\":\"string\"},{\"name\":\"dataUrl\",\"type\":\"string\"},{\"name\":\"price\",\"type\":\"uint256\"}],\"name\":\"addLead\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"id\",\"type\":\"bytes16\"},{\"name\":\"status\",\"type\":\"uint8\"}],\"name\":\"setLeadStatus\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"id\",\"type\":\"bytes16\"}],\"name\":\"transactLead\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"id\",\"type\":\"bytes16\"},{\"name\":\"_price\",\"type\":\"uint256\"}],\"name\":\"setLeadPrice\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]"

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

// SetLeadDataUrl is a paid mutator transaction binding the contract method 0x406865fc.
//
// Solidity: function setLeadDataUrl(id bytes16, _dataUrl string) returns()
func (_HoQuAdCampaignI *HoQuAdCampaignITransactor) SetLeadDataUrl(opts *bind.TransactOpts, id [16]byte, _dataUrl string) (*types.Transaction, error) {
	return _HoQuAdCampaignI.contract.Transact(opts, "setLeadDataUrl", id, _dataUrl)
}

// SetLeadDataUrl is a paid mutator transaction binding the contract method 0x406865fc.
//
// Solidity: function setLeadDataUrl(id bytes16, _dataUrl string) returns()
func (_HoQuAdCampaignI *HoQuAdCampaignISession) SetLeadDataUrl(id [16]byte, _dataUrl string) (*types.Transaction, error) {
	return _HoQuAdCampaignI.Contract.SetLeadDataUrl(&_HoQuAdCampaignI.TransactOpts, id, _dataUrl)
}

// SetLeadDataUrl is a paid mutator transaction binding the contract method 0x406865fc.
//
// Solidity: function setLeadDataUrl(id bytes16, _dataUrl string) returns()
func (_HoQuAdCampaignI *HoQuAdCampaignITransactorSession) SetLeadDataUrl(id [16]byte, _dataUrl string) (*types.Transaction, error) {
	return _HoQuAdCampaignI.Contract.SetLeadDataUrl(&_HoQuAdCampaignI.TransactOpts, id, _dataUrl)
}

// SetLeadPrice is a paid mutator transaction binding the contract method 0xc535bd88.
//
// Solidity: function setLeadPrice(id bytes16, _price uint256) returns()
func (_HoQuAdCampaignI *HoQuAdCampaignITransactor) SetLeadPrice(opts *bind.TransactOpts, id [16]byte, _price *big.Int) (*types.Transaction, error) {
	return _HoQuAdCampaignI.contract.Transact(opts, "setLeadPrice", id, _price)
}

// SetLeadPrice is a paid mutator transaction binding the contract method 0xc535bd88.
//
// Solidity: function setLeadPrice(id bytes16, _price uint256) returns()
func (_HoQuAdCampaignI *HoQuAdCampaignISession) SetLeadPrice(id [16]byte, _price *big.Int) (*types.Transaction, error) {
	return _HoQuAdCampaignI.Contract.SetLeadPrice(&_HoQuAdCampaignI.TransactOpts, id, _price)
}

// SetLeadPrice is a paid mutator transaction binding the contract method 0xc535bd88.
//
// Solidity: function setLeadPrice(id bytes16, _price uint256) returns()
func (_HoQuAdCampaignI *HoQuAdCampaignITransactorSession) SetLeadPrice(id [16]byte, _price *big.Int) (*types.Transaction, error) {
	return _HoQuAdCampaignI.Contract.SetLeadPrice(&_HoQuAdCampaignI.TransactOpts, id, _price)
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
const HoQuConfigBin = `0x60806040526611c37937e0800060015534801561001b57600080fd5b506040516020806108cf8339810160405251600380546201000060b060020a03191633620100000217905560008054600160a060020a031916600160a060020a0390921691909117815561085a90819061007590396000f3006080604052600436106100c45763ffffffff7c0100000000000000000000000000000000000000000000000000000000600035041663355e6b4381146100c95780637065cb48146100e3578063715018a6146101045780637d60b6ce146101195780638d2573321461013a5780638da5cb5b1461015657806397c0262a14610187578063b7b603a21461019c578063b9488546146101c4578063babcc539146101f0578063c15d25b114610225578063e148919114610241578063f2fde38b14610268575b600080fd5b3480156100d557600080fd5b506100e1600435610289565b005b3480156100ef57600080fd5b506100e1600160a060020a03600435166102ff565b34801561011057600080fd5b506100e16103b1565b34801561012557600080fd5b506100e1600160a060020a036004351661042c565b34801561014657600080fd5b506100e161ffff600435166104dd565b34801561016257600080fd5b5061016b6105a8565b60408051600160a060020a039092168252519081900360200190f35b34801561019357600080fd5b5061016b6105bd565b3480156101a857600080fd5b506100e161ffff60043516600160a060020a03602435166105cc565b3480156101d057600080fd5b506101d96106bb565b6040805161ffff9092168252519081900360200190f35b3480156101fc57600080fd5b50610211600160a060020a03600435166106c5565b604080519115158252519081900360200190f35b34801561023157600080fd5b5061016b61ffff60043516610764565b34801561024d57600080fd5b5061025661077f565b60408051918252519081900360200190f35b34801561027457600080fd5b506100e1600160a060020a0360043516610785565b600354620100009004600160a060020a03163314806102ac57506102ac336106c5565b15156102b757600080fd5b600081116102c457600080fd5b60408051828152905133917ff659af9785c5d8c6128fdcbabb637f7d7ea787226f13cad89f71cbc579614fdc919081900360200190a2600155565b600354620100009004600160a060020a03163314806103225750610322336106c5565b151561032d57600080fd5b6003805461ffff908116600090815260026020526040808220805473ffffffffffffffffffffffffffffffffffffffff1916600160a060020a038716908117909155845461ffff19811690851660010190941693909317909355915190917fa91384695ab8222d2f6e2437bb5b1fefe671538a33a7e06e88fee0147fffe8a491a250565b600354620100009004600160a060020a031633146103ce57600080fd5b60035460405162010000909104600160a060020a0316907ff8df31144d9c2f0f6b59d69b8b98abd5459d07f2742c4df920b25aae33c6482090600090a26003805475ffffffffffffffffffffffffffffffffffffffff000019169055565b600354620100009004600160a060020a031633148061044f575061044f336106c5565b151561045a57600080fd5b600160a060020a038116151561046f57600080fd5b60408051600160a060020a0383168152905133917f394e7a1cfafd7b3c49218efbf47fb95da1832e5eaf0e9dc0428ea7220959d0e0919081900360200190a26000805473ffffffffffffffffffffffffffffffffffffffff1916600160a060020a0392909216919091179055565b600354620100009004600160a060020a03163314806105005750610500336106c5565b151561050b57600080fd5b61ffff8116600090815260026020526040902054600160a060020a0316151561053357600080fd5b61ffff8116600090815260026020526040808220549051600160a060020a03909116917f7ac6b089a9620017a12220abae1696205383eb917822ade81acb4eea2a424f6991a261ffff166000908152600260205260409020805473ffffffffffffffffffffffffffffffffffffffff19169055565b600354620100009004600160a060020a031681565b600054600160a060020a031681565b600354620100009004600160a060020a03163314806105ef57506105ef336106c5565b15156105fa57600080fd5b61ffff8216600090815260026020526040902054600160a060020a03161561062157600080fd5b61ffff8216600090815260026020908152604091829020548251600160a060020a038581168252935193909116927f1722097e017bb9a189312ac57bbdfff1bf251383cfd36c245f8ebabd8734c759929181900390910190a261ffff919091166000908152600260205260409020805473ffffffffffffffffffffffffffffffffffffffff1916600160a060020a03909216919091179055565b60035461ffff1681565b60008080600160a060020a03841615156106de57600080fd5b600091505b60035461ffff9081169083161015610733575061ffff8116600090815260026020526040902054600160a060020a03908116908416811415610728576001925061075d565b6001909101906106e3565b600354600160a060020a0385811662010000909204161415610758576001925061075d565b600092505b5050919050565b600260205260009081526040902054600160a060020a031681565b60015481565b600354620100009004600160a060020a031633146107a257600080fd5b600160a060020a03811615156107b757600080fd5b600354604051600160a060020a038084169262010000900416907f8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e090600090a360038054600160a060020a03909216620100000275ffffffffffffffffffffffffffffffffffffffff0000199092169190911790555600a165627a7a72305820ed1fe9b8a28cc4f754add7bee73c6992abbfa2636a99450b6e31cd7759f9a4c10029`

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

// HoQuConfigIABI is the input ABI used to generate the binding from.
const HoQuConfigIABI = "[{\"constant\":false,\"inputs\":[{\"name\":\"_commission\",\"type\":\"uint256\"}],\"name\":\"setCommission\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_owner\",\"type\":\"address\"}],\"name\":\"addOwner\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_commissionWallet\",\"type\":\"address\"}],\"name\":\"setCommissionWallet\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"i\",\"type\":\"uint16\"}],\"name\":\"deleteOwner\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"commissionWallet\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"i\",\"type\":\"uint16\"},{\"name\":\"_owner\",\"type\":\"address\"}],\"name\":\"changeOwner\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"ownersCount\",\"outputs\":[{\"name\":\"\",\"type\":\"uint16\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_owner\",\"type\":\"address\"}],\"name\":\"isAllowed\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"\",\"type\":\"uint16\"}],\"name\":\"owners\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"commission\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"}]"

// HoQuConfigIBin is the compiled bytecode used for deploying new contracts.
const HoQuConfigIBin = `0x`

// DeployHoQuConfigI deploys a new Ethereum contract, binding an instance of HoQuConfigI to it.
func DeployHoQuConfigI(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *HoQuConfigI, error) {
	parsed, err := abi.JSON(strings.NewReader(HoQuConfigIABI))
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	address, tx, contract, err := bind.DeployContract(auth, parsed, common.FromHex(HoQuConfigIBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &HoQuConfigI{HoQuConfigICaller: HoQuConfigICaller{contract: contract}, HoQuConfigITransactor: HoQuConfigITransactor{contract: contract}, HoQuConfigIFilterer: HoQuConfigIFilterer{contract: contract}}, nil
}

// HoQuConfigI is an auto generated Go binding around an Ethereum contract.
type HoQuConfigI struct {
	HoQuConfigICaller     // Read-only binding to the contract
	HoQuConfigITransactor // Write-only binding to the contract
	HoQuConfigIFilterer   // Log filterer for contract events
}

// HoQuConfigICaller is an auto generated read-only Go binding around an Ethereum contract.
type HoQuConfigICaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// HoQuConfigITransactor is an auto generated write-only Go binding around an Ethereum contract.
type HoQuConfigITransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// HoQuConfigIFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type HoQuConfigIFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// HoQuConfigISession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type HoQuConfigISession struct {
	Contract     *HoQuConfigI      // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// HoQuConfigICallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type HoQuConfigICallerSession struct {
	Contract *HoQuConfigICaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts      // Call options to use throughout this session
}

// HoQuConfigITransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type HoQuConfigITransactorSession struct {
	Contract     *HoQuConfigITransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts      // Transaction auth options to use throughout this session
}

// HoQuConfigIRaw is an auto generated low-level Go binding around an Ethereum contract.
type HoQuConfigIRaw struct {
	Contract *HoQuConfigI // Generic contract binding to access the raw methods on
}

// HoQuConfigICallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type HoQuConfigICallerRaw struct {
	Contract *HoQuConfigICaller // Generic read-only contract binding to access the raw methods on
}

// HoQuConfigITransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type HoQuConfigITransactorRaw struct {
	Contract *HoQuConfigITransactor // Generic write-only contract binding to access the raw methods on
}

// NewHoQuConfigI creates a new instance of HoQuConfigI, bound to a specific deployed contract.
func NewHoQuConfigI(address common.Address, backend bind.ContractBackend) (*HoQuConfigI, error) {
	contract, err := bindHoQuConfigI(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &HoQuConfigI{HoQuConfigICaller: HoQuConfigICaller{contract: contract}, HoQuConfigITransactor: HoQuConfigITransactor{contract: contract}, HoQuConfigIFilterer: HoQuConfigIFilterer{contract: contract}}, nil
}

// NewHoQuConfigICaller creates a new read-only instance of HoQuConfigI, bound to a specific deployed contract.
func NewHoQuConfigICaller(address common.Address, caller bind.ContractCaller) (*HoQuConfigICaller, error) {
	contract, err := bindHoQuConfigI(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &HoQuConfigICaller{contract: contract}, nil
}

// NewHoQuConfigITransactor creates a new write-only instance of HoQuConfigI, bound to a specific deployed contract.
func NewHoQuConfigITransactor(address common.Address, transactor bind.ContractTransactor) (*HoQuConfigITransactor, error) {
	contract, err := bindHoQuConfigI(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &HoQuConfigITransactor{contract: contract}, nil
}

// NewHoQuConfigIFilterer creates a new log filterer instance of HoQuConfigI, bound to a specific deployed contract.
func NewHoQuConfigIFilterer(address common.Address, filterer bind.ContractFilterer) (*HoQuConfigIFilterer, error) {
	contract, err := bindHoQuConfigI(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &HoQuConfigIFilterer{contract: contract}, nil
}

// bindHoQuConfigI binds a generic wrapper to an already deployed contract.
func bindHoQuConfigI(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(HoQuConfigIABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_HoQuConfigI *HoQuConfigIRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _HoQuConfigI.Contract.HoQuConfigICaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_HoQuConfigI *HoQuConfigIRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _HoQuConfigI.Contract.HoQuConfigITransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_HoQuConfigI *HoQuConfigIRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _HoQuConfigI.Contract.HoQuConfigITransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_HoQuConfigI *HoQuConfigICallerRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _HoQuConfigI.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_HoQuConfigI *HoQuConfigITransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _HoQuConfigI.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_HoQuConfigI *HoQuConfigITransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _HoQuConfigI.Contract.contract.Transact(opts, method, params...)
}

// Commission is a free data retrieval call binding the contract method 0xe1489191.
//
// Solidity: function commission() constant returns(uint256)
func (_HoQuConfigI *HoQuConfigICaller) Commission(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _HoQuConfigI.contract.Call(opts, out, "commission")
	return *ret0, err
}

// Commission is a free data retrieval call binding the contract method 0xe1489191.
//
// Solidity: function commission() constant returns(uint256)
func (_HoQuConfigI *HoQuConfigISession) Commission() (*big.Int, error) {
	return _HoQuConfigI.Contract.Commission(&_HoQuConfigI.CallOpts)
}

// Commission is a free data retrieval call binding the contract method 0xe1489191.
//
// Solidity: function commission() constant returns(uint256)
func (_HoQuConfigI *HoQuConfigICallerSession) Commission() (*big.Int, error) {
	return _HoQuConfigI.Contract.Commission(&_HoQuConfigI.CallOpts)
}

// CommissionWallet is a free data retrieval call binding the contract method 0x97c0262a.
//
// Solidity: function commissionWallet() constant returns(address)
func (_HoQuConfigI *HoQuConfigICaller) CommissionWallet(opts *bind.CallOpts) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _HoQuConfigI.contract.Call(opts, out, "commissionWallet")
	return *ret0, err
}

// CommissionWallet is a free data retrieval call binding the contract method 0x97c0262a.
//
// Solidity: function commissionWallet() constant returns(address)
func (_HoQuConfigI *HoQuConfigISession) CommissionWallet() (common.Address, error) {
	return _HoQuConfigI.Contract.CommissionWallet(&_HoQuConfigI.CallOpts)
}

// CommissionWallet is a free data retrieval call binding the contract method 0x97c0262a.
//
// Solidity: function commissionWallet() constant returns(address)
func (_HoQuConfigI *HoQuConfigICallerSession) CommissionWallet() (common.Address, error) {
	return _HoQuConfigI.Contract.CommissionWallet(&_HoQuConfigI.CallOpts)
}

// Owners is a free data retrieval call binding the contract method 0xc15d25b1.
//
// Solidity: function owners( uint16) constant returns(address)
func (_HoQuConfigI *HoQuConfigICaller) Owners(opts *bind.CallOpts, arg0 uint16) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _HoQuConfigI.contract.Call(opts, out, "owners", arg0)
	return *ret0, err
}

// Owners is a free data retrieval call binding the contract method 0xc15d25b1.
//
// Solidity: function owners( uint16) constant returns(address)
func (_HoQuConfigI *HoQuConfigISession) Owners(arg0 uint16) (common.Address, error) {
	return _HoQuConfigI.Contract.Owners(&_HoQuConfigI.CallOpts, arg0)
}

// Owners is a free data retrieval call binding the contract method 0xc15d25b1.
//
// Solidity: function owners( uint16) constant returns(address)
func (_HoQuConfigI *HoQuConfigICallerSession) Owners(arg0 uint16) (common.Address, error) {
	return _HoQuConfigI.Contract.Owners(&_HoQuConfigI.CallOpts, arg0)
}

// OwnersCount is a free data retrieval call binding the contract method 0xb9488546.
//
// Solidity: function ownersCount() constant returns(uint16)
func (_HoQuConfigI *HoQuConfigICaller) OwnersCount(opts *bind.CallOpts) (uint16, error) {
	var (
		ret0 = new(uint16)
	)
	out := ret0
	err := _HoQuConfigI.contract.Call(opts, out, "ownersCount")
	return *ret0, err
}

// OwnersCount is a free data retrieval call binding the contract method 0xb9488546.
//
// Solidity: function ownersCount() constant returns(uint16)
func (_HoQuConfigI *HoQuConfigISession) OwnersCount() (uint16, error) {
	return _HoQuConfigI.Contract.OwnersCount(&_HoQuConfigI.CallOpts)
}

// OwnersCount is a free data retrieval call binding the contract method 0xb9488546.
//
// Solidity: function ownersCount() constant returns(uint16)
func (_HoQuConfigI *HoQuConfigICallerSession) OwnersCount() (uint16, error) {
	return _HoQuConfigI.Contract.OwnersCount(&_HoQuConfigI.CallOpts)
}

// AddOwner is a paid mutator transaction binding the contract method 0x7065cb48.
//
// Solidity: function addOwner(_owner address) returns()
func (_HoQuConfigI *HoQuConfigITransactor) AddOwner(opts *bind.TransactOpts, _owner common.Address) (*types.Transaction, error) {
	return _HoQuConfigI.contract.Transact(opts, "addOwner", _owner)
}

// AddOwner is a paid mutator transaction binding the contract method 0x7065cb48.
//
// Solidity: function addOwner(_owner address) returns()
func (_HoQuConfigI *HoQuConfigISession) AddOwner(_owner common.Address) (*types.Transaction, error) {
	return _HoQuConfigI.Contract.AddOwner(&_HoQuConfigI.TransactOpts, _owner)
}

// AddOwner is a paid mutator transaction binding the contract method 0x7065cb48.
//
// Solidity: function addOwner(_owner address) returns()
func (_HoQuConfigI *HoQuConfigITransactorSession) AddOwner(_owner common.Address) (*types.Transaction, error) {
	return _HoQuConfigI.Contract.AddOwner(&_HoQuConfigI.TransactOpts, _owner)
}

// ChangeOwner is a paid mutator transaction binding the contract method 0xb7b603a2.
//
// Solidity: function changeOwner(i uint16, _owner address) returns()
func (_HoQuConfigI *HoQuConfigITransactor) ChangeOwner(opts *bind.TransactOpts, i uint16, _owner common.Address) (*types.Transaction, error) {
	return _HoQuConfigI.contract.Transact(opts, "changeOwner", i, _owner)
}

// ChangeOwner is a paid mutator transaction binding the contract method 0xb7b603a2.
//
// Solidity: function changeOwner(i uint16, _owner address) returns()
func (_HoQuConfigI *HoQuConfigISession) ChangeOwner(i uint16, _owner common.Address) (*types.Transaction, error) {
	return _HoQuConfigI.Contract.ChangeOwner(&_HoQuConfigI.TransactOpts, i, _owner)
}

// ChangeOwner is a paid mutator transaction binding the contract method 0xb7b603a2.
//
// Solidity: function changeOwner(i uint16, _owner address) returns()
func (_HoQuConfigI *HoQuConfigITransactorSession) ChangeOwner(i uint16, _owner common.Address) (*types.Transaction, error) {
	return _HoQuConfigI.Contract.ChangeOwner(&_HoQuConfigI.TransactOpts, i, _owner)
}

// DeleteOwner is a paid mutator transaction binding the contract method 0x8d257332.
//
// Solidity: function deleteOwner(i uint16) returns()
func (_HoQuConfigI *HoQuConfigITransactor) DeleteOwner(opts *bind.TransactOpts, i uint16) (*types.Transaction, error) {
	return _HoQuConfigI.contract.Transact(opts, "deleteOwner", i)
}

// DeleteOwner is a paid mutator transaction binding the contract method 0x8d257332.
//
// Solidity: function deleteOwner(i uint16) returns()
func (_HoQuConfigI *HoQuConfigISession) DeleteOwner(i uint16) (*types.Transaction, error) {
	return _HoQuConfigI.Contract.DeleteOwner(&_HoQuConfigI.TransactOpts, i)
}

// DeleteOwner is a paid mutator transaction binding the contract method 0x8d257332.
//
// Solidity: function deleteOwner(i uint16) returns()
func (_HoQuConfigI *HoQuConfigITransactorSession) DeleteOwner(i uint16) (*types.Transaction, error) {
	return _HoQuConfigI.Contract.DeleteOwner(&_HoQuConfigI.TransactOpts, i)
}

// IsAllowed is a paid mutator transaction binding the contract method 0xbabcc539.
//
// Solidity: function isAllowed(_owner address) returns(bool)
func (_HoQuConfigI *HoQuConfigITransactor) IsAllowed(opts *bind.TransactOpts, _owner common.Address) (*types.Transaction, error) {
	return _HoQuConfigI.contract.Transact(opts, "isAllowed", _owner)
}

// IsAllowed is a paid mutator transaction binding the contract method 0xbabcc539.
//
// Solidity: function isAllowed(_owner address) returns(bool)
func (_HoQuConfigI *HoQuConfigISession) IsAllowed(_owner common.Address) (*types.Transaction, error) {
	return _HoQuConfigI.Contract.IsAllowed(&_HoQuConfigI.TransactOpts, _owner)
}

// IsAllowed is a paid mutator transaction binding the contract method 0xbabcc539.
//
// Solidity: function isAllowed(_owner address) returns(bool)
func (_HoQuConfigI *HoQuConfigITransactorSession) IsAllowed(_owner common.Address) (*types.Transaction, error) {
	return _HoQuConfigI.Contract.IsAllowed(&_HoQuConfigI.TransactOpts, _owner)
}

// SetCommission is a paid mutator transaction binding the contract method 0x355e6b43.
//
// Solidity: function setCommission(_commission uint256) returns()
func (_HoQuConfigI *HoQuConfigITransactor) SetCommission(opts *bind.TransactOpts, _commission *big.Int) (*types.Transaction, error) {
	return _HoQuConfigI.contract.Transact(opts, "setCommission", _commission)
}

// SetCommission is a paid mutator transaction binding the contract method 0x355e6b43.
//
// Solidity: function setCommission(_commission uint256) returns()
func (_HoQuConfigI *HoQuConfigISession) SetCommission(_commission *big.Int) (*types.Transaction, error) {
	return _HoQuConfigI.Contract.SetCommission(&_HoQuConfigI.TransactOpts, _commission)
}

// SetCommission is a paid mutator transaction binding the contract method 0x355e6b43.
//
// Solidity: function setCommission(_commission uint256) returns()
func (_HoQuConfigI *HoQuConfigITransactorSession) SetCommission(_commission *big.Int) (*types.Transaction, error) {
	return _HoQuConfigI.Contract.SetCommission(&_HoQuConfigI.TransactOpts, _commission)
}

// SetCommissionWallet is a paid mutator transaction binding the contract method 0x7d60b6ce.
//
// Solidity: function setCommissionWallet(_commissionWallet address) returns()
func (_HoQuConfigI *HoQuConfigITransactor) SetCommissionWallet(opts *bind.TransactOpts, _commissionWallet common.Address) (*types.Transaction, error) {
	return _HoQuConfigI.contract.Transact(opts, "setCommissionWallet", _commissionWallet)
}

// SetCommissionWallet is a paid mutator transaction binding the contract method 0x7d60b6ce.
//
// Solidity: function setCommissionWallet(_commissionWallet address) returns()
func (_HoQuConfigI *HoQuConfigISession) SetCommissionWallet(_commissionWallet common.Address) (*types.Transaction, error) {
	return _HoQuConfigI.Contract.SetCommissionWallet(&_HoQuConfigI.TransactOpts, _commissionWallet)
}

// SetCommissionWallet is a paid mutator transaction binding the contract method 0x7d60b6ce.
//
// Solidity: function setCommissionWallet(_commissionWallet address) returns()
func (_HoQuConfigI *HoQuConfigITransactorSession) SetCommissionWallet(_commissionWallet common.Address) (*types.Transaction, error) {
	return _HoQuConfigI.Contract.SetCommissionWallet(&_HoQuConfigI.TransactOpts, _commissionWallet)
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

// HoQuStorageAccessorABI is the input ABI used to generate the binding from.
const HoQuStorageAccessorABI = "[{\"constant\":false,\"inputs\":[{\"name\":\"storageAddress\",\"type\":\"address\"}],\"name\":\"setStorageAddress\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"id\",\"type\":\"bytes16\"},{\"name\":\"num\",\"type\":\"uint8\"}],\"name\":\"getOfferTariffGroup\",\"outputs\":[{\"name\":\"\",\"type\":\"bytes16\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"config\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"configAddress\",\"type\":\"address\"}],\"name\":\"setConfigAddress\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"store\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"id\",\"type\":\"bytes16\"},{\"name\":\"num\",\"type\":\"uint8\"}],\"name\":\"getUserAddress\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"name\":\"configAddress\",\"type\":\"address\"},{\"name\":\"storageAddress\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"}]"

// HoQuStorageAccessorBin is the compiled bytecode used for deploying new contracts.
const HoQuStorageAccessorBin = `0x608060405234801561001057600080fd5b5060405160408061050e83398101604052805160209091015160008054600160a060020a03938416600160a060020a031991821617909155600180549390921692169190911790556104a7806100676000396000f3006080604052600436106100775763ffffffff7c010000000000000000000000000000000000000000000000000000000060003504166359b910d6811461007c57806374f627c51461009f57806379502c55146100f657806383a12de914610127578063975057e714610148578063c49e2fb11461015d575b600080fd5b34801561008857600080fd5b5061009d600160a060020a036004351661018e565b005b3480156100ab57600080fd5b506100d06fffffffffffffffffffffffffffffffff196004351660ff6024351661025b565b604080516fffffffffffffffffffffffffffffffff199092168252519081900360200190f35b34801561010257600080fd5b5061010b61030f565b60408051600160a060020a039092168252519081900360200190f35b34801561013357600080fd5b5061009d600160a060020a036004351661031e565b34801561015457600080fd5b5061010b6103eb565b34801561016957600080fd5b5061010b6fffffffffffffffffffffffffffffffff196004351660ff602435166103fa565b60008054604080517fbabcc5390000000000000000000000000000000000000000000000000000000081523360048201529051600160a060020a039092169263babcc539926024808401936020939083900390910190829087803b1580156101f557600080fd5b505af1158015610209573d6000803e3d6000fd5b505050506040513d602081101561021f57600080fd5b5051151561022c57600080fd5b6001805473ffffffffffffffffffffffffffffffffffffffff1916600160a060020a0392909216919091179055565b600154604080517f74f627c50000000000000000000000000000000000000000000000000000000081526fffffffffffffffffffffffffffffffff198516600482015260ff841660248201529051600092600160a060020a0316916374f627c591604480830192602092919082900301818787803b1580156102dc57600080fd5b505af11580156102f0573d6000803e3d6000fd5b505050506040513d602081101561030657600080fd5b50519392505050565b600054600160a060020a031681565b60008054604080517fbabcc5390000000000000000000000000000000000000000000000000000000081523360048201529051600160a060020a039092169263babcc539926024808401936020939083900390910190829087803b15801561038557600080fd5b505af1158015610399573d6000803e3d6000fd5b505050506040513d60208110156103af57600080fd5b505115156103bc57600080fd5b6000805473ffffffffffffffffffffffffffffffffffffffff1916600160a060020a0392909216919091179055565b600154600160a060020a031681565b600154604080517fc49e2fb10000000000000000000000000000000000000000000000000000000081526fffffffffffffffffffffffffffffffff198516600482015260ff841660248201529051600092600160a060020a03169163c49e2fb191604480830192602092919082900301818787803b1580156102dc57600080fd00a165627a7a72305820581b924c136a2819c5f2e184ba1fa248ae6eecaca99df92384d5040919b42cf70029`

// DeployHoQuStorageAccessor deploys a new Ethereum contract, binding an instance of HoQuStorageAccessor to it.
func DeployHoQuStorageAccessor(auth *bind.TransactOpts, backend bind.ContractBackend, configAddress common.Address, storageAddress common.Address) (common.Address, *types.Transaction, *HoQuStorageAccessor, error) {
	parsed, err := abi.JSON(strings.NewReader(HoQuStorageAccessorABI))
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	address, tx, contract, err := bind.DeployContract(auth, parsed, common.FromHex(HoQuStorageAccessorBin), backend, configAddress, storageAddress)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &HoQuStorageAccessor{HoQuStorageAccessorCaller: HoQuStorageAccessorCaller{contract: contract}, HoQuStorageAccessorTransactor: HoQuStorageAccessorTransactor{contract: contract}, HoQuStorageAccessorFilterer: HoQuStorageAccessorFilterer{contract: contract}}, nil
}

// HoQuStorageAccessor is an auto generated Go binding around an Ethereum contract.
type HoQuStorageAccessor struct {
	HoQuStorageAccessorCaller     // Read-only binding to the contract
	HoQuStorageAccessorTransactor // Write-only binding to the contract
	HoQuStorageAccessorFilterer   // Log filterer for contract events
}

// HoQuStorageAccessorCaller is an auto generated read-only Go binding around an Ethereum contract.
type HoQuStorageAccessorCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// HoQuStorageAccessorTransactor is an auto generated write-only Go binding around an Ethereum contract.
type HoQuStorageAccessorTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// HoQuStorageAccessorFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type HoQuStorageAccessorFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// HoQuStorageAccessorSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type HoQuStorageAccessorSession struct {
	Contract     *HoQuStorageAccessor // Generic contract binding to set the session for
	CallOpts     bind.CallOpts        // Call options to use throughout this session
	TransactOpts bind.TransactOpts    // Transaction auth options to use throughout this session
}

// HoQuStorageAccessorCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type HoQuStorageAccessorCallerSession struct {
	Contract *HoQuStorageAccessorCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts              // Call options to use throughout this session
}

// HoQuStorageAccessorTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type HoQuStorageAccessorTransactorSession struct {
	Contract     *HoQuStorageAccessorTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts              // Transaction auth options to use throughout this session
}

// HoQuStorageAccessorRaw is an auto generated low-level Go binding around an Ethereum contract.
type HoQuStorageAccessorRaw struct {
	Contract *HoQuStorageAccessor // Generic contract binding to access the raw methods on
}

// HoQuStorageAccessorCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type HoQuStorageAccessorCallerRaw struct {
	Contract *HoQuStorageAccessorCaller // Generic read-only contract binding to access the raw methods on
}

// HoQuStorageAccessorTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type HoQuStorageAccessorTransactorRaw struct {
	Contract *HoQuStorageAccessorTransactor // Generic write-only contract binding to access the raw methods on
}

// NewHoQuStorageAccessor creates a new instance of HoQuStorageAccessor, bound to a specific deployed contract.
func NewHoQuStorageAccessor(address common.Address, backend bind.ContractBackend) (*HoQuStorageAccessor, error) {
	contract, err := bindHoQuStorageAccessor(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &HoQuStorageAccessor{HoQuStorageAccessorCaller: HoQuStorageAccessorCaller{contract: contract}, HoQuStorageAccessorTransactor: HoQuStorageAccessorTransactor{contract: contract}, HoQuStorageAccessorFilterer: HoQuStorageAccessorFilterer{contract: contract}}, nil
}

// NewHoQuStorageAccessorCaller creates a new read-only instance of HoQuStorageAccessor, bound to a specific deployed contract.
func NewHoQuStorageAccessorCaller(address common.Address, caller bind.ContractCaller) (*HoQuStorageAccessorCaller, error) {
	contract, err := bindHoQuStorageAccessor(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &HoQuStorageAccessorCaller{contract: contract}, nil
}

// NewHoQuStorageAccessorTransactor creates a new write-only instance of HoQuStorageAccessor, bound to a specific deployed contract.
func NewHoQuStorageAccessorTransactor(address common.Address, transactor bind.ContractTransactor) (*HoQuStorageAccessorTransactor, error) {
	contract, err := bindHoQuStorageAccessor(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &HoQuStorageAccessorTransactor{contract: contract}, nil
}

// NewHoQuStorageAccessorFilterer creates a new log filterer instance of HoQuStorageAccessor, bound to a specific deployed contract.
func NewHoQuStorageAccessorFilterer(address common.Address, filterer bind.ContractFilterer) (*HoQuStorageAccessorFilterer, error) {
	contract, err := bindHoQuStorageAccessor(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &HoQuStorageAccessorFilterer{contract: contract}, nil
}

// bindHoQuStorageAccessor binds a generic wrapper to an already deployed contract.
func bindHoQuStorageAccessor(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(HoQuStorageAccessorABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_HoQuStorageAccessor *HoQuStorageAccessorRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _HoQuStorageAccessor.Contract.HoQuStorageAccessorCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_HoQuStorageAccessor *HoQuStorageAccessorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _HoQuStorageAccessor.Contract.HoQuStorageAccessorTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_HoQuStorageAccessor *HoQuStorageAccessorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _HoQuStorageAccessor.Contract.HoQuStorageAccessorTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_HoQuStorageAccessor *HoQuStorageAccessorCallerRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _HoQuStorageAccessor.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_HoQuStorageAccessor *HoQuStorageAccessorTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _HoQuStorageAccessor.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_HoQuStorageAccessor *HoQuStorageAccessorTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _HoQuStorageAccessor.Contract.contract.Transact(opts, method, params...)
}

// Config is a free data retrieval call binding the contract method 0x79502c55.
//
// Solidity: function config() constant returns(address)
func (_HoQuStorageAccessor *HoQuStorageAccessorCaller) Config(opts *bind.CallOpts) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _HoQuStorageAccessor.contract.Call(opts, out, "config")
	return *ret0, err
}

// Config is a free data retrieval call binding the contract method 0x79502c55.
//
// Solidity: function config() constant returns(address)
func (_HoQuStorageAccessor *HoQuStorageAccessorSession) Config() (common.Address, error) {
	return _HoQuStorageAccessor.Contract.Config(&_HoQuStorageAccessor.CallOpts)
}

// Config is a free data retrieval call binding the contract method 0x79502c55.
//
// Solidity: function config() constant returns(address)
func (_HoQuStorageAccessor *HoQuStorageAccessorCallerSession) Config() (common.Address, error) {
	return _HoQuStorageAccessor.Contract.Config(&_HoQuStorageAccessor.CallOpts)
}

// GetOfferTariffGroup is a free data retrieval call binding the contract method 0x74f627c5.
//
// Solidity: function getOfferTariffGroup(id bytes16, num uint8) constant returns(bytes16)
func (_HoQuStorageAccessor *HoQuStorageAccessorCaller) GetOfferTariffGroup(opts *bind.CallOpts, id [16]byte, num uint8) ([16]byte, error) {
	var (
		ret0 = new([16]byte)
	)
	out := ret0
	err := _HoQuStorageAccessor.contract.Call(opts, out, "getOfferTariffGroup", id, num)
	return *ret0, err
}

// GetOfferTariffGroup is a free data retrieval call binding the contract method 0x74f627c5.
//
// Solidity: function getOfferTariffGroup(id bytes16, num uint8) constant returns(bytes16)
func (_HoQuStorageAccessor *HoQuStorageAccessorSession) GetOfferTariffGroup(id [16]byte, num uint8) ([16]byte, error) {
	return _HoQuStorageAccessor.Contract.GetOfferTariffGroup(&_HoQuStorageAccessor.CallOpts, id, num)
}

// GetOfferTariffGroup is a free data retrieval call binding the contract method 0x74f627c5.
//
// Solidity: function getOfferTariffGroup(id bytes16, num uint8) constant returns(bytes16)
func (_HoQuStorageAccessor *HoQuStorageAccessorCallerSession) GetOfferTariffGroup(id [16]byte, num uint8) ([16]byte, error) {
	return _HoQuStorageAccessor.Contract.GetOfferTariffGroup(&_HoQuStorageAccessor.CallOpts, id, num)
}

// GetUserAddress is a free data retrieval call binding the contract method 0xc49e2fb1.
//
// Solidity: function getUserAddress(id bytes16, num uint8) constant returns(address)
func (_HoQuStorageAccessor *HoQuStorageAccessorCaller) GetUserAddress(opts *bind.CallOpts, id [16]byte, num uint8) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _HoQuStorageAccessor.contract.Call(opts, out, "getUserAddress", id, num)
	return *ret0, err
}

// GetUserAddress is a free data retrieval call binding the contract method 0xc49e2fb1.
//
// Solidity: function getUserAddress(id bytes16, num uint8) constant returns(address)
func (_HoQuStorageAccessor *HoQuStorageAccessorSession) GetUserAddress(id [16]byte, num uint8) (common.Address, error) {
	return _HoQuStorageAccessor.Contract.GetUserAddress(&_HoQuStorageAccessor.CallOpts, id, num)
}

// GetUserAddress is a free data retrieval call binding the contract method 0xc49e2fb1.
//
// Solidity: function getUserAddress(id bytes16, num uint8) constant returns(address)
func (_HoQuStorageAccessor *HoQuStorageAccessorCallerSession) GetUserAddress(id [16]byte, num uint8) (common.Address, error) {
	return _HoQuStorageAccessor.Contract.GetUserAddress(&_HoQuStorageAccessor.CallOpts, id, num)
}

// Store is a free data retrieval call binding the contract method 0x975057e7.
//
// Solidity: function store() constant returns(address)
func (_HoQuStorageAccessor *HoQuStorageAccessorCaller) Store(opts *bind.CallOpts) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _HoQuStorageAccessor.contract.Call(opts, out, "store")
	return *ret0, err
}

// Store is a free data retrieval call binding the contract method 0x975057e7.
//
// Solidity: function store() constant returns(address)
func (_HoQuStorageAccessor *HoQuStorageAccessorSession) Store() (common.Address, error) {
	return _HoQuStorageAccessor.Contract.Store(&_HoQuStorageAccessor.CallOpts)
}

// Store is a free data retrieval call binding the contract method 0x975057e7.
//
// Solidity: function store() constant returns(address)
func (_HoQuStorageAccessor *HoQuStorageAccessorCallerSession) Store() (common.Address, error) {
	return _HoQuStorageAccessor.Contract.Store(&_HoQuStorageAccessor.CallOpts)
}

// SetConfigAddress is a paid mutator transaction binding the contract method 0x83a12de9.
//
// Solidity: function setConfigAddress(configAddress address) returns()
func (_HoQuStorageAccessor *HoQuStorageAccessorTransactor) SetConfigAddress(opts *bind.TransactOpts, configAddress common.Address) (*types.Transaction, error) {
	return _HoQuStorageAccessor.contract.Transact(opts, "setConfigAddress", configAddress)
}

// SetConfigAddress is a paid mutator transaction binding the contract method 0x83a12de9.
//
// Solidity: function setConfigAddress(configAddress address) returns()
func (_HoQuStorageAccessor *HoQuStorageAccessorSession) SetConfigAddress(configAddress common.Address) (*types.Transaction, error) {
	return _HoQuStorageAccessor.Contract.SetConfigAddress(&_HoQuStorageAccessor.TransactOpts, configAddress)
}

// SetConfigAddress is a paid mutator transaction binding the contract method 0x83a12de9.
//
// Solidity: function setConfigAddress(configAddress address) returns()
func (_HoQuStorageAccessor *HoQuStorageAccessorTransactorSession) SetConfigAddress(configAddress common.Address) (*types.Transaction, error) {
	return _HoQuStorageAccessor.Contract.SetConfigAddress(&_HoQuStorageAccessor.TransactOpts, configAddress)
}

// SetStorageAddress is a paid mutator transaction binding the contract method 0x59b910d6.
//
// Solidity: function setStorageAddress(storageAddress address) returns()
func (_HoQuStorageAccessor *HoQuStorageAccessorTransactor) SetStorageAddress(opts *bind.TransactOpts, storageAddress common.Address) (*types.Transaction, error) {
	return _HoQuStorageAccessor.contract.Transact(opts, "setStorageAddress", storageAddress)
}

// SetStorageAddress is a paid mutator transaction binding the contract method 0x59b910d6.
//
// Solidity: function setStorageAddress(storageAddress address) returns()
func (_HoQuStorageAccessor *HoQuStorageAccessorSession) SetStorageAddress(storageAddress common.Address) (*types.Transaction, error) {
	return _HoQuStorageAccessor.Contract.SetStorageAddress(&_HoQuStorageAccessor.TransactOpts, storageAddress)
}

// SetStorageAddress is a paid mutator transaction binding the contract method 0x59b910d6.
//
// Solidity: function setStorageAddress(storageAddress address) returns()
func (_HoQuStorageAccessor *HoQuStorageAccessorTransactorSession) SetStorageAddress(storageAddress common.Address) (*types.Transaction, error) {
	return _HoQuStorageAccessor.Contract.SetStorageAddress(&_HoQuStorageAccessor.TransactOpts, storageAddress)
}

// HoQuStorageIABI is the input ABI used to generate the binding from.
const HoQuStorageIABI = "[{\"constant\":false,\"inputs\":[{\"name\":\"id\",\"type\":\"bytes16\"},{\"name\":\"ownerId\",\"type\":\"bytes16\"},{\"name\":\"name\",\"type\":\"string\"},{\"name\":\"dataUrl\",\"type\":\"string\"},{\"name\":\"status\",\"type\":\"uint8\"}],\"name\":\"setNetwork\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"id\",\"type\":\"bytes16\"},{\"name\":\"ownerAddress\",\"type\":\"address\"}],\"name\":\"addUserAddress\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"id\",\"type\":\"bytes16\"},{\"name\":\"ownerId\",\"type\":\"bytes16\"},{\"name\":\"networkId\",\"type\":\"bytes16\"},{\"name\":\"merchantId\",\"type\":\"bytes16\"},{\"name\":\"payerAddress\",\"type\":\"address\"},{\"name\":\"name\",\"type\":\"string\"},{\"name\":\"dataUrl\",\"type\":\"string\"},{\"name\":\"status\",\"type\":\"uint8\"}],\"name\":\"setOffer\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"id\",\"type\":\"bytes16\"},{\"name\":\"meta\",\"type\":\"string\"},{\"name\":\"kycLevel\",\"type\":\"uint8\"},{\"name\":\"dataUrl\",\"type\":\"string\"}],\"name\":\"addKycReport\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"id\",\"type\":\"bytes16\"},{\"name\":\"role\",\"type\":\"string\"},{\"name\":\"ownerAddress\",\"type\":\"address\"},{\"name\":\"kycLevel\",\"type\":\"uint8\"},{\"name\":\"pubKey\",\"type\":\"string\"},{\"name\":\"status\",\"type\":\"uint8\"}],\"name\":\"setUser\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"\",\"type\":\"bytes16\"}],\"name\":\"ids\",\"outputs\":[{\"name\":\"userId\",\"type\":\"bytes16\"},{\"name\":\"companyId\",\"type\":\"bytes16\"},{\"name\":\"idType\",\"type\":\"string\"},{\"name\":\"name\",\"type\":\"string\"},{\"name\":\"numOfKycReports\",\"type\":\"uint16\"},{\"name\":\"createdAt\",\"type\":\"uint256\"},{\"name\":\"status\",\"type\":\"uint8\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"id\",\"type\":\"bytes16\"},{\"name\":\"num\",\"type\":\"uint8\"},{\"name\":\"tariffGroupId\",\"type\":\"bytes16\"}],\"name\":\"setOfferTariffGroup\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"id\",\"type\":\"bytes16\"},{\"name\":\"num\",\"type\":\"uint8\"}],\"name\":\"getTariffGroupTariff\",\"outputs\":[{\"name\":\"\",\"type\":\"bytes16\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"\",\"type\":\"bytes16\"}],\"name\":\"companies\",\"outputs\":[{\"name\":\"ownerId\",\"type\":\"bytes16\"},{\"name\":\"name\",\"type\":\"string\"},{\"name\":\"dataUrl\",\"type\":\"string\"},{\"name\":\"createdAt\",\"type\":\"uint256\"},{\"name\":\"status\",\"type\":\"uint8\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"id\",\"type\":\"bytes16\"},{\"name\":\"tariffGroupId\",\"type\":\"bytes16\"},{\"name\":\"name\",\"type\":\"string\"},{\"name\":\"action\",\"type\":\"string\"},{\"name\":\"calcMethod\",\"type\":\"string\"},{\"name\":\"price\",\"type\":\"uint256\"},{\"name\":\"status\",\"type\":\"uint8\"}],\"name\":\"setTariff\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"\",\"type\":\"bytes16\"}],\"name\":\"networks\",\"outputs\":[{\"name\":\"ownerId\",\"type\":\"bytes16\"},{\"name\":\"name\",\"type\":\"string\"},{\"name\":\"dataUrl\",\"type\":\"string\"},{\"name\":\"createdAt\",\"type\":\"uint256\"},{\"name\":\"status\",\"type\":\"uint8\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"id\",\"type\":\"bytes16\"},{\"name\":\"num\",\"type\":\"uint8\"}],\"name\":\"getOfferTariffGroup\",\"outputs\":[{\"name\":\"\",\"type\":\"bytes16\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"\",\"type\":\"bytes16\"}],\"name\":\"users\",\"outputs\":[{\"name\":\"ownerAddress\",\"type\":\"address\"},{\"name\":\"numOfAddresses\",\"type\":\"uint8\"},{\"name\":\"role\",\"type\":\"string\"},{\"name\":\"kycLevel\",\"type\":\"uint8\"},{\"name\":\"pubKey\",\"type\":\"string\"},{\"name\":\"createdAt\",\"type\":\"uint256\"},{\"name\":\"status\",\"type\":\"uint8\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"id\",\"type\":\"bytes16\"},{\"name\":\"ownerId\",\"type\":\"bytes16\"},{\"name\":\"name\",\"type\":\"string\"},{\"name\":\"dataUrl\",\"type\":\"string\"},{\"name\":\"status\",\"type\":\"uint8\"}],\"name\":\"setCompany\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"\",\"type\":\"bytes16\"}],\"name\":\"adCampaigns\",\"outputs\":[{\"name\":\"ownerId\",\"type\":\"bytes16\"},{\"name\":\"offerId\",\"type\":\"bytes16\"},{\"name\":\"contractAddress\",\"type\":\"address\"},{\"name\":\"createdAt\",\"type\":\"uint256\"},{\"name\":\"status\",\"type\":\"uint8\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"\",\"type\":\"bytes16\"}],\"name\":\"tariffGroups\",\"outputs\":[{\"name\":\"ownerId\",\"type\":\"bytes16\"},{\"name\":\"name\",\"type\":\"string\"},{\"name\":\"numOfTariffs\",\"type\":\"uint8\"},{\"name\":\"createdAt\",\"type\":\"uint256\"},{\"name\":\"status\",\"type\":\"uint8\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"id\",\"type\":\"bytes16\"},{\"name\":\"ownerId\",\"type\":\"bytes16\"},{\"name\":\"offerId\",\"type\":\"bytes16\"},{\"name\":\"contractAddress\",\"type\":\"address\"},{\"name\":\"status\",\"type\":\"uint8\"}],\"name\":\"setAdCampaign\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"\",\"type\":\"bytes16\"}],\"name\":\"stats\",\"outputs\":[{\"name\":\"rating\",\"type\":\"uint256\"},{\"name\":\"volume\",\"type\":\"uint256\"},{\"name\":\"members\",\"type\":\"uint256\"},{\"name\":\"alfa\",\"type\":\"uint256\"},{\"name\":\"beta\",\"type\":\"uint256\"},{\"name\":\"status\",\"type\":\"uint8\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"\",\"type\":\"bytes16\"}],\"name\":\"tariffs\",\"outputs\":[{\"name\":\"tariffGroupId\",\"type\":\"bytes16\"},{\"name\":\"name\",\"type\":\"string\"},{\"name\":\"action\",\"type\":\"string\"},{\"name\":\"calcMethod\",\"type\":\"string\"},{\"name\":\"price\",\"type\":\"uint256\"},{\"name\":\"createdAt\",\"type\":\"uint256\"},{\"name\":\"status\",\"type\":\"uint8\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"id\",\"type\":\"bytes16\"},{\"name\":\"ownerId\",\"type\":\"bytes16\"},{\"name\":\"name\",\"type\":\"string\"},{\"name\":\"status\",\"type\":\"uint8\"}],\"name\":\"setTariffGroup\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"id\",\"type\":\"bytes16\"},{\"name\":\"userId\",\"type\":\"bytes16\"},{\"name\":\"idType\",\"type\":\"string\"},{\"name\":\"name\",\"type\":\"string\"},{\"name\":\"companyId\",\"type\":\"bytes16\"},{\"name\":\"status\",\"type\":\"uint8\"}],\"name\":\"setIdentification\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"id\",\"type\":\"bytes16\"},{\"name\":\"userId\",\"type\":\"bytes16\"},{\"name\":\"rating\",\"type\":\"uint256\"},{\"name\":\"volume\",\"type\":\"uint256\"},{\"name\":\"members\",\"type\":\"uint256\"},{\"name\":\"alfa\",\"type\":\"uint256\"},{\"name\":\"beta\",\"type\":\"uint256\"},{\"name\":\"status\",\"type\":\"uint8\"}],\"name\":\"setStats\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"\",\"type\":\"bytes16\"}],\"name\":\"offers\",\"outputs\":[{\"name\":\"ownerId\",\"type\":\"bytes16\"},{\"name\":\"networkId\",\"type\":\"bytes16\"},{\"name\":\"merchantId\",\"type\":\"bytes16\"},{\"name\":\"payerAddress\",\"type\":\"address\"},{\"name\":\"name\",\"type\":\"string\"},{\"name\":\"dataUrl\",\"type\":\"string\"},{\"name\":\"numOfTariffGroups\",\"type\":\"uint8\"},{\"name\":\"createdAt\",\"type\":\"uint256\"},{\"name\":\"status\",\"type\":\"uint8\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"id\",\"type\":\"bytes16\"},{\"name\":\"num\",\"type\":\"uint8\"}],\"name\":\"getUserAddress\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"id\",\"type\":\"bytes16\"},{\"name\":\"ownerId\",\"type\":\"bytes16\"},{\"name\":\"networkId\",\"type\":\"bytes16\"},{\"name\":\"name\",\"type\":\"string\"},{\"name\":\"dataUrl\",\"type\":\"string\"},{\"name\":\"status\",\"type\":\"uint8\"}],\"name\":\"setTracker\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"id\",\"type\":\"bytes16\"},{\"name\":\"num\",\"type\":\"uint16\"}],\"name\":\"getKycReport\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"},{\"name\":\"\",\"type\":\"string\"},{\"name\":\"\",\"type\":\"uint8\"},{\"name\":\"\",\"type\":\"string\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"\",\"type\":\"bytes16\"}],\"name\":\"trackers\",\"outputs\":[{\"name\":\"ownerId\",\"type\":\"bytes16\"},{\"name\":\"networkId\",\"type\":\"bytes16\"},{\"name\":\"name\",\"type\":\"string\"},{\"name\":\"dataUrl\",\"type\":\"string\"},{\"name\":\"createdAt\",\"type\":\"uint256\"},{\"name\":\"status\",\"type\":\"uint8\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"id\",\"type\":\"bytes16\"},{\"name\":\"tariffGroupId\",\"type\":\"bytes16\"}],\"name\":\"addOfferTariffGroup\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"id\",\"type\":\"bytes16\"},{\"name\":\"num\",\"type\":\"uint8\"},{\"name\":\"tariffId\",\"type\":\"bytes16\"}],\"name\":\"setTariffGroupTariff\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]"

// HoQuStorageIBin is the compiled bytecode used for deploying new contracts.
const HoQuStorageIBin = `0x`

// DeployHoQuStorageI deploys a new Ethereum contract, binding an instance of HoQuStorageI to it.
func DeployHoQuStorageI(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *HoQuStorageI, error) {
	parsed, err := abi.JSON(strings.NewReader(HoQuStorageIABI))
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	address, tx, contract, err := bind.DeployContract(auth, parsed, common.FromHex(HoQuStorageIBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &HoQuStorageI{HoQuStorageICaller: HoQuStorageICaller{contract: contract}, HoQuStorageITransactor: HoQuStorageITransactor{contract: contract}, HoQuStorageIFilterer: HoQuStorageIFilterer{contract: contract}}, nil
}

// HoQuStorageI is an auto generated Go binding around an Ethereum contract.
type HoQuStorageI struct {
	HoQuStorageICaller     // Read-only binding to the contract
	HoQuStorageITransactor // Write-only binding to the contract
	HoQuStorageIFilterer   // Log filterer for contract events
}

// HoQuStorageICaller is an auto generated read-only Go binding around an Ethereum contract.
type HoQuStorageICaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// HoQuStorageITransactor is an auto generated write-only Go binding around an Ethereum contract.
type HoQuStorageITransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// HoQuStorageIFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type HoQuStorageIFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// HoQuStorageISession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type HoQuStorageISession struct {
	Contract     *HoQuStorageI     // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// HoQuStorageICallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type HoQuStorageICallerSession struct {
	Contract *HoQuStorageICaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts       // Call options to use throughout this session
}

// HoQuStorageITransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type HoQuStorageITransactorSession struct {
	Contract     *HoQuStorageITransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts       // Transaction auth options to use throughout this session
}

// HoQuStorageIRaw is an auto generated low-level Go binding around an Ethereum contract.
type HoQuStorageIRaw struct {
	Contract *HoQuStorageI // Generic contract binding to access the raw methods on
}

// HoQuStorageICallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type HoQuStorageICallerRaw struct {
	Contract *HoQuStorageICaller // Generic read-only contract binding to access the raw methods on
}

// HoQuStorageITransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type HoQuStorageITransactorRaw struct {
	Contract *HoQuStorageITransactor // Generic write-only contract binding to access the raw methods on
}

// NewHoQuStorageI creates a new instance of HoQuStorageI, bound to a specific deployed contract.
func NewHoQuStorageI(address common.Address, backend bind.ContractBackend) (*HoQuStorageI, error) {
	contract, err := bindHoQuStorageI(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &HoQuStorageI{HoQuStorageICaller: HoQuStorageICaller{contract: contract}, HoQuStorageITransactor: HoQuStorageITransactor{contract: contract}, HoQuStorageIFilterer: HoQuStorageIFilterer{contract: contract}}, nil
}

// NewHoQuStorageICaller creates a new read-only instance of HoQuStorageI, bound to a specific deployed contract.
func NewHoQuStorageICaller(address common.Address, caller bind.ContractCaller) (*HoQuStorageICaller, error) {
	contract, err := bindHoQuStorageI(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &HoQuStorageICaller{contract: contract}, nil
}

// NewHoQuStorageITransactor creates a new write-only instance of HoQuStorageI, bound to a specific deployed contract.
func NewHoQuStorageITransactor(address common.Address, transactor bind.ContractTransactor) (*HoQuStorageITransactor, error) {
	contract, err := bindHoQuStorageI(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &HoQuStorageITransactor{contract: contract}, nil
}

// NewHoQuStorageIFilterer creates a new log filterer instance of HoQuStorageI, bound to a specific deployed contract.
func NewHoQuStorageIFilterer(address common.Address, filterer bind.ContractFilterer) (*HoQuStorageIFilterer, error) {
	contract, err := bindHoQuStorageI(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &HoQuStorageIFilterer{contract: contract}, nil
}

// bindHoQuStorageI binds a generic wrapper to an already deployed contract.
func bindHoQuStorageI(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(HoQuStorageIABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_HoQuStorageI *HoQuStorageIRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _HoQuStorageI.Contract.HoQuStorageICaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_HoQuStorageI *HoQuStorageIRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _HoQuStorageI.Contract.HoQuStorageITransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_HoQuStorageI *HoQuStorageIRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _HoQuStorageI.Contract.HoQuStorageITransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_HoQuStorageI *HoQuStorageICallerRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _HoQuStorageI.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_HoQuStorageI *HoQuStorageITransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _HoQuStorageI.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_HoQuStorageI *HoQuStorageITransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _HoQuStorageI.Contract.contract.Transact(opts, method, params...)
}

// AdCampaigns is a free data retrieval call binding the contract method 0x89075557.
//
// Solidity: function adCampaigns( bytes16) constant returns(ownerId bytes16, offerId bytes16, contractAddress address, createdAt uint256, status uint8)
func (_HoQuStorageI *HoQuStorageICaller) AdCampaigns(opts *bind.CallOpts, arg0 [16]byte) (struct {
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
	err := _HoQuStorageI.contract.Call(opts, out, "adCampaigns", arg0)
	return *ret, err
}

// AdCampaigns is a free data retrieval call binding the contract method 0x89075557.
//
// Solidity: function adCampaigns( bytes16) constant returns(ownerId bytes16, offerId bytes16, contractAddress address, createdAt uint256, status uint8)
func (_HoQuStorageI *HoQuStorageISession) AdCampaigns(arg0 [16]byte) (struct {
	OwnerId         [16]byte
	OfferId         [16]byte
	ContractAddress common.Address
	CreatedAt       *big.Int
	Status          uint8
}, error) {
	return _HoQuStorageI.Contract.AdCampaigns(&_HoQuStorageI.CallOpts, arg0)
}

// AdCampaigns is a free data retrieval call binding the contract method 0x89075557.
//
// Solidity: function adCampaigns( bytes16) constant returns(ownerId bytes16, offerId bytes16, contractAddress address, createdAt uint256, status uint8)
func (_HoQuStorageI *HoQuStorageICallerSession) AdCampaigns(arg0 [16]byte) (struct {
	OwnerId         [16]byte
	OfferId         [16]byte
	ContractAddress common.Address
	CreatedAt       *big.Int
	Status          uint8
}, error) {
	return _HoQuStorageI.Contract.AdCampaigns(&_HoQuStorageI.CallOpts, arg0)
}

// Companies is a free data retrieval call binding the contract method 0x4df2d847.
//
// Solidity: function companies( bytes16) constant returns(ownerId bytes16, name string, dataUrl string, createdAt uint256, status uint8)
func (_HoQuStorageI *HoQuStorageICaller) Companies(opts *bind.CallOpts, arg0 [16]byte) (struct {
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
	err := _HoQuStorageI.contract.Call(opts, out, "companies", arg0)
	return *ret, err
}

// Companies is a free data retrieval call binding the contract method 0x4df2d847.
//
// Solidity: function companies( bytes16) constant returns(ownerId bytes16, name string, dataUrl string, createdAt uint256, status uint8)
func (_HoQuStorageI *HoQuStorageISession) Companies(arg0 [16]byte) (struct {
	OwnerId   [16]byte
	Name      string
	DataUrl   string
	CreatedAt *big.Int
	Status    uint8
}, error) {
	return _HoQuStorageI.Contract.Companies(&_HoQuStorageI.CallOpts, arg0)
}

// Companies is a free data retrieval call binding the contract method 0x4df2d847.
//
// Solidity: function companies( bytes16) constant returns(ownerId bytes16, name string, dataUrl string, createdAt uint256, status uint8)
func (_HoQuStorageI *HoQuStorageICallerSession) Companies(arg0 [16]byte) (struct {
	OwnerId   [16]byte
	Name      string
	DataUrl   string
	CreatedAt *big.Int
	Status    uint8
}, error) {
	return _HoQuStorageI.Contract.Companies(&_HoQuStorageI.CallOpts, arg0)
}

// GetKycReport is a free data retrieval call binding the contract method 0xcc3436b9.
//
// Solidity: function getKycReport(id bytes16, num uint16) constant returns(uint256, string, uint8, string)
func (_HoQuStorageI *HoQuStorageICaller) GetKycReport(opts *bind.CallOpts, id [16]byte, num uint16) (*big.Int, string, uint8, string, error) {
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
	err := _HoQuStorageI.contract.Call(opts, out, "getKycReport", id, num)
	return *ret0, *ret1, *ret2, *ret3, err
}

// GetKycReport is a free data retrieval call binding the contract method 0xcc3436b9.
//
// Solidity: function getKycReport(id bytes16, num uint16) constant returns(uint256, string, uint8, string)
func (_HoQuStorageI *HoQuStorageISession) GetKycReport(id [16]byte, num uint16) (*big.Int, string, uint8, string, error) {
	return _HoQuStorageI.Contract.GetKycReport(&_HoQuStorageI.CallOpts, id, num)
}

// GetKycReport is a free data retrieval call binding the contract method 0xcc3436b9.
//
// Solidity: function getKycReport(id bytes16, num uint16) constant returns(uint256, string, uint8, string)
func (_HoQuStorageI *HoQuStorageICallerSession) GetKycReport(id [16]byte, num uint16) (*big.Int, string, uint8, string, error) {
	return _HoQuStorageI.Contract.GetKycReport(&_HoQuStorageI.CallOpts, id, num)
}

// GetOfferTariffGroup is a free data retrieval call binding the contract method 0x74f627c5.
//
// Solidity: function getOfferTariffGroup(id bytes16, num uint8) constant returns(bytes16)
func (_HoQuStorageI *HoQuStorageICaller) GetOfferTariffGroup(opts *bind.CallOpts, id [16]byte, num uint8) ([16]byte, error) {
	var (
		ret0 = new([16]byte)
	)
	out := ret0
	err := _HoQuStorageI.contract.Call(opts, out, "getOfferTariffGroup", id, num)
	return *ret0, err
}

// GetOfferTariffGroup is a free data retrieval call binding the contract method 0x74f627c5.
//
// Solidity: function getOfferTariffGroup(id bytes16, num uint8) constant returns(bytes16)
func (_HoQuStorageI *HoQuStorageISession) GetOfferTariffGroup(id [16]byte, num uint8) ([16]byte, error) {
	return _HoQuStorageI.Contract.GetOfferTariffGroup(&_HoQuStorageI.CallOpts, id, num)
}

// GetOfferTariffGroup is a free data retrieval call binding the contract method 0x74f627c5.
//
// Solidity: function getOfferTariffGroup(id bytes16, num uint8) constant returns(bytes16)
func (_HoQuStorageI *HoQuStorageICallerSession) GetOfferTariffGroup(id [16]byte, num uint8) ([16]byte, error) {
	return _HoQuStorageI.Contract.GetOfferTariffGroup(&_HoQuStorageI.CallOpts, id, num)
}

// GetTariffGroupTariff is a free data retrieval call binding the contract method 0x436e7b65.
//
// Solidity: function getTariffGroupTariff(id bytes16, num uint8) constant returns(bytes16)
func (_HoQuStorageI *HoQuStorageICaller) GetTariffGroupTariff(opts *bind.CallOpts, id [16]byte, num uint8) ([16]byte, error) {
	var (
		ret0 = new([16]byte)
	)
	out := ret0
	err := _HoQuStorageI.contract.Call(opts, out, "getTariffGroupTariff", id, num)
	return *ret0, err
}

// GetTariffGroupTariff is a free data retrieval call binding the contract method 0x436e7b65.
//
// Solidity: function getTariffGroupTariff(id bytes16, num uint8) constant returns(bytes16)
func (_HoQuStorageI *HoQuStorageISession) GetTariffGroupTariff(id [16]byte, num uint8) ([16]byte, error) {
	return _HoQuStorageI.Contract.GetTariffGroupTariff(&_HoQuStorageI.CallOpts, id, num)
}

// GetTariffGroupTariff is a free data retrieval call binding the contract method 0x436e7b65.
//
// Solidity: function getTariffGroupTariff(id bytes16, num uint8) constant returns(bytes16)
func (_HoQuStorageI *HoQuStorageICallerSession) GetTariffGroupTariff(id [16]byte, num uint8) ([16]byte, error) {
	return _HoQuStorageI.Contract.GetTariffGroupTariff(&_HoQuStorageI.CallOpts, id, num)
}

// GetUserAddress is a free data retrieval call binding the contract method 0xc49e2fb1.
//
// Solidity: function getUserAddress(id bytes16, num uint8) constant returns(address)
func (_HoQuStorageI *HoQuStorageICaller) GetUserAddress(opts *bind.CallOpts, id [16]byte, num uint8) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _HoQuStorageI.contract.Call(opts, out, "getUserAddress", id, num)
	return *ret0, err
}

// GetUserAddress is a free data retrieval call binding the contract method 0xc49e2fb1.
//
// Solidity: function getUserAddress(id bytes16, num uint8) constant returns(address)
func (_HoQuStorageI *HoQuStorageISession) GetUserAddress(id [16]byte, num uint8) (common.Address, error) {
	return _HoQuStorageI.Contract.GetUserAddress(&_HoQuStorageI.CallOpts, id, num)
}

// GetUserAddress is a free data retrieval call binding the contract method 0xc49e2fb1.
//
// Solidity: function getUserAddress(id bytes16, num uint8) constant returns(address)
func (_HoQuStorageI *HoQuStorageICallerSession) GetUserAddress(id [16]byte, num uint8) (common.Address, error) {
	return _HoQuStorageI.Contract.GetUserAddress(&_HoQuStorageI.CallOpts, id, num)
}

// Ids is a free data retrieval call binding the contract method 0x289e7470.
//
// Solidity: function ids( bytes16) constant returns(userId bytes16, companyId bytes16, idType string, name string, numOfKycReports uint16, createdAt uint256, status uint8)
func (_HoQuStorageI *HoQuStorageICaller) Ids(opts *bind.CallOpts, arg0 [16]byte) (struct {
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
	err := _HoQuStorageI.contract.Call(opts, out, "ids", arg0)
	return *ret, err
}

// Ids is a free data retrieval call binding the contract method 0x289e7470.
//
// Solidity: function ids( bytes16) constant returns(userId bytes16, companyId bytes16, idType string, name string, numOfKycReports uint16, createdAt uint256, status uint8)
func (_HoQuStorageI *HoQuStorageISession) Ids(arg0 [16]byte) (struct {
	UserId          [16]byte
	CompanyId       [16]byte
	IdType          string
	Name            string
	NumOfKycReports uint16
	CreatedAt       *big.Int
	Status          uint8
}, error) {
	return _HoQuStorageI.Contract.Ids(&_HoQuStorageI.CallOpts, arg0)
}

// Ids is a free data retrieval call binding the contract method 0x289e7470.
//
// Solidity: function ids( bytes16) constant returns(userId bytes16, companyId bytes16, idType string, name string, numOfKycReports uint16, createdAt uint256, status uint8)
func (_HoQuStorageI *HoQuStorageICallerSession) Ids(arg0 [16]byte) (struct {
	UserId          [16]byte
	CompanyId       [16]byte
	IdType          string
	Name            string
	NumOfKycReports uint16
	CreatedAt       *big.Int
	Status          uint8
}, error) {
	return _HoQuStorageI.Contract.Ids(&_HoQuStorageI.CallOpts, arg0)
}

// Networks is a free data retrieval call binding the contract method 0x72c130e4.
//
// Solidity: function networks( bytes16) constant returns(ownerId bytes16, name string, dataUrl string, createdAt uint256, status uint8)
func (_HoQuStorageI *HoQuStorageICaller) Networks(opts *bind.CallOpts, arg0 [16]byte) (struct {
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
	err := _HoQuStorageI.contract.Call(opts, out, "networks", arg0)
	return *ret, err
}

// Networks is a free data retrieval call binding the contract method 0x72c130e4.
//
// Solidity: function networks( bytes16) constant returns(ownerId bytes16, name string, dataUrl string, createdAt uint256, status uint8)
func (_HoQuStorageI *HoQuStorageISession) Networks(arg0 [16]byte) (struct {
	OwnerId   [16]byte
	Name      string
	DataUrl   string
	CreatedAt *big.Int
	Status    uint8
}, error) {
	return _HoQuStorageI.Contract.Networks(&_HoQuStorageI.CallOpts, arg0)
}

// Networks is a free data retrieval call binding the contract method 0x72c130e4.
//
// Solidity: function networks( bytes16) constant returns(ownerId bytes16, name string, dataUrl string, createdAt uint256, status uint8)
func (_HoQuStorageI *HoQuStorageICallerSession) Networks(arg0 [16]byte) (struct {
	OwnerId   [16]byte
	Name      string
	DataUrl   string
	CreatedAt *big.Int
	Status    uint8
}, error) {
	return _HoQuStorageI.Contract.Networks(&_HoQuStorageI.CallOpts, arg0)
}

// Offers is a free data retrieval call binding the contract method 0xc32d869b.
//
// Solidity: function offers( bytes16) constant returns(ownerId bytes16, networkId bytes16, merchantId bytes16, payerAddress address, name string, dataUrl string, numOfTariffGroups uint8, createdAt uint256, status uint8)
func (_HoQuStorageI *HoQuStorageICaller) Offers(opts *bind.CallOpts, arg0 [16]byte) (struct {
	OwnerId           [16]byte
	NetworkId         [16]byte
	MerchantId        [16]byte
	PayerAddress      common.Address
	Name              string
	DataUrl           string
	NumOfTariffGroups uint8
	CreatedAt         *big.Int
	Status            uint8
}, error) {
	ret := new(struct {
		OwnerId           [16]byte
		NetworkId         [16]byte
		MerchantId        [16]byte
		PayerAddress      common.Address
		Name              string
		DataUrl           string
		NumOfTariffGroups uint8
		CreatedAt         *big.Int
		Status            uint8
	})
	out := ret
	err := _HoQuStorageI.contract.Call(opts, out, "offers", arg0)
	return *ret, err
}

// Offers is a free data retrieval call binding the contract method 0xc32d869b.
//
// Solidity: function offers( bytes16) constant returns(ownerId bytes16, networkId bytes16, merchantId bytes16, payerAddress address, name string, dataUrl string, numOfTariffGroups uint8, createdAt uint256, status uint8)
func (_HoQuStorageI *HoQuStorageISession) Offers(arg0 [16]byte) (struct {
	OwnerId           [16]byte
	NetworkId         [16]byte
	MerchantId        [16]byte
	PayerAddress      common.Address
	Name              string
	DataUrl           string
	NumOfTariffGroups uint8
	CreatedAt         *big.Int
	Status            uint8
}, error) {
	return _HoQuStorageI.Contract.Offers(&_HoQuStorageI.CallOpts, arg0)
}

// Offers is a free data retrieval call binding the contract method 0xc32d869b.
//
// Solidity: function offers( bytes16) constant returns(ownerId bytes16, networkId bytes16, merchantId bytes16, payerAddress address, name string, dataUrl string, numOfTariffGroups uint8, createdAt uint256, status uint8)
func (_HoQuStorageI *HoQuStorageICallerSession) Offers(arg0 [16]byte) (struct {
	OwnerId           [16]byte
	NetworkId         [16]byte
	MerchantId        [16]byte
	PayerAddress      common.Address
	Name              string
	DataUrl           string
	NumOfTariffGroups uint8
	CreatedAt         *big.Int
	Status            uint8
}, error) {
	return _HoQuStorageI.Contract.Offers(&_HoQuStorageI.CallOpts, arg0)
}

// Stats is a free data retrieval call binding the contract method 0xa2affa4d.
//
// Solidity: function stats( bytes16) constant returns(rating uint256, volume uint256, members uint256, alfa uint256, beta uint256, status uint8)
func (_HoQuStorageI *HoQuStorageICaller) Stats(opts *bind.CallOpts, arg0 [16]byte) (struct {
	Rating  *big.Int
	Volume  *big.Int
	Members *big.Int
	Alfa    *big.Int
	Beta    *big.Int
	Status  uint8
}, error) {
	ret := new(struct {
		Rating  *big.Int
		Volume  *big.Int
		Members *big.Int
		Alfa    *big.Int
		Beta    *big.Int
		Status  uint8
	})
	out := ret
	err := _HoQuStorageI.contract.Call(opts, out, "stats", arg0)
	return *ret, err
}

// Stats is a free data retrieval call binding the contract method 0xa2affa4d.
//
// Solidity: function stats( bytes16) constant returns(rating uint256, volume uint256, members uint256, alfa uint256, beta uint256, status uint8)
func (_HoQuStorageI *HoQuStorageISession) Stats(arg0 [16]byte) (struct {
	Rating  *big.Int
	Volume  *big.Int
	Members *big.Int
	Alfa    *big.Int
	Beta    *big.Int
	Status  uint8
}, error) {
	return _HoQuStorageI.Contract.Stats(&_HoQuStorageI.CallOpts, arg0)
}

// Stats is a free data retrieval call binding the contract method 0xa2affa4d.
//
// Solidity: function stats( bytes16) constant returns(rating uint256, volume uint256, members uint256, alfa uint256, beta uint256, status uint8)
func (_HoQuStorageI *HoQuStorageICallerSession) Stats(arg0 [16]byte) (struct {
	Rating  *big.Int
	Volume  *big.Int
	Members *big.Int
	Alfa    *big.Int
	Beta    *big.Int
	Status  uint8
}, error) {
	return _HoQuStorageI.Contract.Stats(&_HoQuStorageI.CallOpts, arg0)
}

// TariffGroups is a free data retrieval call binding the contract method 0x8ade0027.
//
// Solidity: function tariffGroups( bytes16) constant returns(ownerId bytes16, name string, numOfTariffs uint8, createdAt uint256, status uint8)
func (_HoQuStorageI *HoQuStorageICaller) TariffGroups(opts *bind.CallOpts, arg0 [16]byte) (struct {
	OwnerId      [16]byte
	Name         string
	NumOfTariffs uint8
	CreatedAt    *big.Int
	Status       uint8
}, error) {
	ret := new(struct {
		OwnerId      [16]byte
		Name         string
		NumOfTariffs uint8
		CreatedAt    *big.Int
		Status       uint8
	})
	out := ret
	err := _HoQuStorageI.contract.Call(opts, out, "tariffGroups", arg0)
	return *ret, err
}

// TariffGroups is a free data retrieval call binding the contract method 0x8ade0027.
//
// Solidity: function tariffGroups( bytes16) constant returns(ownerId bytes16, name string, numOfTariffs uint8, createdAt uint256, status uint8)
func (_HoQuStorageI *HoQuStorageISession) TariffGroups(arg0 [16]byte) (struct {
	OwnerId      [16]byte
	Name         string
	NumOfTariffs uint8
	CreatedAt    *big.Int
	Status       uint8
}, error) {
	return _HoQuStorageI.Contract.TariffGroups(&_HoQuStorageI.CallOpts, arg0)
}

// TariffGroups is a free data retrieval call binding the contract method 0x8ade0027.
//
// Solidity: function tariffGroups( bytes16) constant returns(ownerId bytes16, name string, numOfTariffs uint8, createdAt uint256, status uint8)
func (_HoQuStorageI *HoQuStorageICallerSession) TariffGroups(arg0 [16]byte) (struct {
	OwnerId      [16]byte
	Name         string
	NumOfTariffs uint8
	CreatedAt    *big.Int
	Status       uint8
}, error) {
	return _HoQuStorageI.Contract.TariffGroups(&_HoQuStorageI.CallOpts, arg0)
}

// Tariffs is a free data retrieval call binding the contract method 0xa2cbda41.
//
// Solidity: function tariffs( bytes16) constant returns(tariffGroupId bytes16, name string, action string, calcMethod string, price uint256, createdAt uint256, status uint8)
func (_HoQuStorageI *HoQuStorageICaller) Tariffs(opts *bind.CallOpts, arg0 [16]byte) (struct {
	TariffGroupId [16]byte
	Name          string
	Action        string
	CalcMethod    string
	Price         *big.Int
	CreatedAt     *big.Int
	Status        uint8
}, error) {
	ret := new(struct {
		TariffGroupId [16]byte
		Name          string
		Action        string
		CalcMethod    string
		Price         *big.Int
		CreatedAt     *big.Int
		Status        uint8
	})
	out := ret
	err := _HoQuStorageI.contract.Call(opts, out, "tariffs", arg0)
	return *ret, err
}

// Tariffs is a free data retrieval call binding the contract method 0xa2cbda41.
//
// Solidity: function tariffs( bytes16) constant returns(tariffGroupId bytes16, name string, action string, calcMethod string, price uint256, createdAt uint256, status uint8)
func (_HoQuStorageI *HoQuStorageISession) Tariffs(arg0 [16]byte) (struct {
	TariffGroupId [16]byte
	Name          string
	Action        string
	CalcMethod    string
	Price         *big.Int
	CreatedAt     *big.Int
	Status        uint8
}, error) {
	return _HoQuStorageI.Contract.Tariffs(&_HoQuStorageI.CallOpts, arg0)
}

// Tariffs is a free data retrieval call binding the contract method 0xa2cbda41.
//
// Solidity: function tariffs( bytes16) constant returns(tariffGroupId bytes16, name string, action string, calcMethod string, price uint256, createdAt uint256, status uint8)
func (_HoQuStorageI *HoQuStorageICallerSession) Tariffs(arg0 [16]byte) (struct {
	TariffGroupId [16]byte
	Name          string
	Action        string
	CalcMethod    string
	Price         *big.Int
	CreatedAt     *big.Int
	Status        uint8
}, error) {
	return _HoQuStorageI.Contract.Tariffs(&_HoQuStorageI.CallOpts, arg0)
}

// Trackers is a free data retrieval call binding the contract method 0xea2fdd8c.
//
// Solidity: function trackers( bytes16) constant returns(ownerId bytes16, networkId bytes16, name string, dataUrl string, createdAt uint256, status uint8)
func (_HoQuStorageI *HoQuStorageICaller) Trackers(opts *bind.CallOpts, arg0 [16]byte) (struct {
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
	err := _HoQuStorageI.contract.Call(opts, out, "trackers", arg0)
	return *ret, err
}

// Trackers is a free data retrieval call binding the contract method 0xea2fdd8c.
//
// Solidity: function trackers( bytes16) constant returns(ownerId bytes16, networkId bytes16, name string, dataUrl string, createdAt uint256, status uint8)
func (_HoQuStorageI *HoQuStorageISession) Trackers(arg0 [16]byte) (struct {
	OwnerId   [16]byte
	NetworkId [16]byte
	Name      string
	DataUrl   string
	CreatedAt *big.Int
	Status    uint8
}, error) {
	return _HoQuStorageI.Contract.Trackers(&_HoQuStorageI.CallOpts, arg0)
}

// Trackers is a free data retrieval call binding the contract method 0xea2fdd8c.
//
// Solidity: function trackers( bytes16) constant returns(ownerId bytes16, networkId bytes16, name string, dataUrl string, createdAt uint256, status uint8)
func (_HoQuStorageI *HoQuStorageICallerSession) Trackers(arg0 [16]byte) (struct {
	OwnerId   [16]byte
	NetworkId [16]byte
	Name      string
	DataUrl   string
	CreatedAt *big.Int
	Status    uint8
}, error) {
	return _HoQuStorageI.Contract.Trackers(&_HoQuStorageI.CallOpts, arg0)
}

// Users is a free data retrieval call binding the contract method 0x7520dd14.
//
// Solidity: function users( bytes16) constant returns(ownerAddress address, numOfAddresses uint8, role string, kycLevel uint8, pubKey string, createdAt uint256, status uint8)
func (_HoQuStorageI *HoQuStorageICaller) Users(opts *bind.CallOpts, arg0 [16]byte) (struct {
	OwnerAddress   common.Address
	NumOfAddresses uint8
	Role           string
	KycLevel       uint8
	PubKey         string
	CreatedAt      *big.Int
	Status         uint8
}, error) {
	ret := new(struct {
		OwnerAddress   common.Address
		NumOfAddresses uint8
		Role           string
		KycLevel       uint8
		PubKey         string
		CreatedAt      *big.Int
		Status         uint8
	})
	out := ret
	err := _HoQuStorageI.contract.Call(opts, out, "users", arg0)
	return *ret, err
}

// Users is a free data retrieval call binding the contract method 0x7520dd14.
//
// Solidity: function users( bytes16) constant returns(ownerAddress address, numOfAddresses uint8, role string, kycLevel uint8, pubKey string, createdAt uint256, status uint8)
func (_HoQuStorageI *HoQuStorageISession) Users(arg0 [16]byte) (struct {
	OwnerAddress   common.Address
	NumOfAddresses uint8
	Role           string
	KycLevel       uint8
	PubKey         string
	CreatedAt      *big.Int
	Status         uint8
}, error) {
	return _HoQuStorageI.Contract.Users(&_HoQuStorageI.CallOpts, arg0)
}

// Users is a free data retrieval call binding the contract method 0x7520dd14.
//
// Solidity: function users( bytes16) constant returns(ownerAddress address, numOfAddresses uint8, role string, kycLevel uint8, pubKey string, createdAt uint256, status uint8)
func (_HoQuStorageI *HoQuStorageICallerSession) Users(arg0 [16]byte) (struct {
	OwnerAddress   common.Address
	NumOfAddresses uint8
	Role           string
	KycLevel       uint8
	PubKey         string
	CreatedAt      *big.Int
	Status         uint8
}, error) {
	return _HoQuStorageI.Contract.Users(&_HoQuStorageI.CallOpts, arg0)
}

// AddKycReport is a paid mutator transaction binding the contract method 0x18c8f305.
//
// Solidity: function addKycReport(id bytes16, meta string, kycLevel uint8, dataUrl string) returns()
func (_HoQuStorageI *HoQuStorageITransactor) AddKycReport(opts *bind.TransactOpts, id [16]byte, meta string, kycLevel uint8, dataUrl string) (*types.Transaction, error) {
	return _HoQuStorageI.contract.Transact(opts, "addKycReport", id, meta, kycLevel, dataUrl)
}

// AddKycReport is a paid mutator transaction binding the contract method 0x18c8f305.
//
// Solidity: function addKycReport(id bytes16, meta string, kycLevel uint8, dataUrl string) returns()
func (_HoQuStorageI *HoQuStorageISession) AddKycReport(id [16]byte, meta string, kycLevel uint8, dataUrl string) (*types.Transaction, error) {
	return _HoQuStorageI.Contract.AddKycReport(&_HoQuStorageI.TransactOpts, id, meta, kycLevel, dataUrl)
}

// AddKycReport is a paid mutator transaction binding the contract method 0x18c8f305.
//
// Solidity: function addKycReport(id bytes16, meta string, kycLevel uint8, dataUrl string) returns()
func (_HoQuStorageI *HoQuStorageITransactorSession) AddKycReport(id [16]byte, meta string, kycLevel uint8, dataUrl string) (*types.Transaction, error) {
	return _HoQuStorageI.Contract.AddKycReport(&_HoQuStorageI.TransactOpts, id, meta, kycLevel, dataUrl)
}

// AddOfferTariffGroup is a paid mutator transaction binding the contract method 0xf17eceb1.
//
// Solidity: function addOfferTariffGroup(id bytes16, tariffGroupId bytes16) returns()
func (_HoQuStorageI *HoQuStorageITransactor) AddOfferTariffGroup(opts *bind.TransactOpts, id [16]byte, tariffGroupId [16]byte) (*types.Transaction, error) {
	return _HoQuStorageI.contract.Transact(opts, "addOfferTariffGroup", id, tariffGroupId)
}

// AddOfferTariffGroup is a paid mutator transaction binding the contract method 0xf17eceb1.
//
// Solidity: function addOfferTariffGroup(id bytes16, tariffGroupId bytes16) returns()
func (_HoQuStorageI *HoQuStorageISession) AddOfferTariffGroup(id [16]byte, tariffGroupId [16]byte) (*types.Transaction, error) {
	return _HoQuStorageI.Contract.AddOfferTariffGroup(&_HoQuStorageI.TransactOpts, id, tariffGroupId)
}

// AddOfferTariffGroup is a paid mutator transaction binding the contract method 0xf17eceb1.
//
// Solidity: function addOfferTariffGroup(id bytes16, tariffGroupId bytes16) returns()
func (_HoQuStorageI *HoQuStorageITransactorSession) AddOfferTariffGroup(id [16]byte, tariffGroupId [16]byte) (*types.Transaction, error) {
	return _HoQuStorageI.Contract.AddOfferTariffGroup(&_HoQuStorageI.TransactOpts, id, tariffGroupId)
}

// AddUserAddress is a paid mutator transaction binding the contract method 0x03806028.
//
// Solidity: function addUserAddress(id bytes16, ownerAddress address) returns()
func (_HoQuStorageI *HoQuStorageITransactor) AddUserAddress(opts *bind.TransactOpts, id [16]byte, ownerAddress common.Address) (*types.Transaction, error) {
	return _HoQuStorageI.contract.Transact(opts, "addUserAddress", id, ownerAddress)
}

// AddUserAddress is a paid mutator transaction binding the contract method 0x03806028.
//
// Solidity: function addUserAddress(id bytes16, ownerAddress address) returns()
func (_HoQuStorageI *HoQuStorageISession) AddUserAddress(id [16]byte, ownerAddress common.Address) (*types.Transaction, error) {
	return _HoQuStorageI.Contract.AddUserAddress(&_HoQuStorageI.TransactOpts, id, ownerAddress)
}

// AddUserAddress is a paid mutator transaction binding the contract method 0x03806028.
//
// Solidity: function addUserAddress(id bytes16, ownerAddress address) returns()
func (_HoQuStorageI *HoQuStorageITransactorSession) AddUserAddress(id [16]byte, ownerAddress common.Address) (*types.Transaction, error) {
	return _HoQuStorageI.Contract.AddUserAddress(&_HoQuStorageI.TransactOpts, id, ownerAddress)
}

// SetAdCampaign is a paid mutator transaction binding the contract method 0xa0191d20.
//
// Solidity: function setAdCampaign(id bytes16, ownerId bytes16, offerId bytes16, contractAddress address, status uint8) returns()
func (_HoQuStorageI *HoQuStorageITransactor) SetAdCampaign(opts *bind.TransactOpts, id [16]byte, ownerId [16]byte, offerId [16]byte, contractAddress common.Address, status uint8) (*types.Transaction, error) {
	return _HoQuStorageI.contract.Transact(opts, "setAdCampaign", id, ownerId, offerId, contractAddress, status)
}

// SetAdCampaign is a paid mutator transaction binding the contract method 0xa0191d20.
//
// Solidity: function setAdCampaign(id bytes16, ownerId bytes16, offerId bytes16, contractAddress address, status uint8) returns()
func (_HoQuStorageI *HoQuStorageISession) SetAdCampaign(id [16]byte, ownerId [16]byte, offerId [16]byte, contractAddress common.Address, status uint8) (*types.Transaction, error) {
	return _HoQuStorageI.Contract.SetAdCampaign(&_HoQuStorageI.TransactOpts, id, ownerId, offerId, contractAddress, status)
}

// SetAdCampaign is a paid mutator transaction binding the contract method 0xa0191d20.
//
// Solidity: function setAdCampaign(id bytes16, ownerId bytes16, offerId bytes16, contractAddress address, status uint8) returns()
func (_HoQuStorageI *HoQuStorageITransactorSession) SetAdCampaign(id [16]byte, ownerId [16]byte, offerId [16]byte, contractAddress common.Address, status uint8) (*types.Transaction, error) {
	return _HoQuStorageI.Contract.SetAdCampaign(&_HoQuStorageI.TransactOpts, id, ownerId, offerId, contractAddress, status)
}

// SetCompany is a paid mutator transaction binding the contract method 0x768d80a8.
//
// Solidity: function setCompany(id bytes16, ownerId bytes16, name string, dataUrl string, status uint8) returns()
func (_HoQuStorageI *HoQuStorageITransactor) SetCompany(opts *bind.TransactOpts, id [16]byte, ownerId [16]byte, name string, dataUrl string, status uint8) (*types.Transaction, error) {
	return _HoQuStorageI.contract.Transact(opts, "setCompany", id, ownerId, name, dataUrl, status)
}

// SetCompany is a paid mutator transaction binding the contract method 0x768d80a8.
//
// Solidity: function setCompany(id bytes16, ownerId bytes16, name string, dataUrl string, status uint8) returns()
func (_HoQuStorageI *HoQuStorageISession) SetCompany(id [16]byte, ownerId [16]byte, name string, dataUrl string, status uint8) (*types.Transaction, error) {
	return _HoQuStorageI.Contract.SetCompany(&_HoQuStorageI.TransactOpts, id, ownerId, name, dataUrl, status)
}

// SetCompany is a paid mutator transaction binding the contract method 0x768d80a8.
//
// Solidity: function setCompany(id bytes16, ownerId bytes16, name string, dataUrl string, status uint8) returns()
func (_HoQuStorageI *HoQuStorageITransactorSession) SetCompany(id [16]byte, ownerId [16]byte, name string, dataUrl string, status uint8) (*types.Transaction, error) {
	return _HoQuStorageI.Contract.SetCompany(&_HoQuStorageI.TransactOpts, id, ownerId, name, dataUrl, status)
}

// SetIdentification is a paid mutator transaction binding the contract method 0xacc02cce.
//
// Solidity: function setIdentification(id bytes16, userId bytes16, idType string, name string, companyId bytes16, status uint8) returns()
func (_HoQuStorageI *HoQuStorageITransactor) SetIdentification(opts *bind.TransactOpts, id [16]byte, userId [16]byte, idType string, name string, companyId [16]byte, status uint8) (*types.Transaction, error) {
	return _HoQuStorageI.contract.Transact(opts, "setIdentification", id, userId, idType, name, companyId, status)
}

// SetIdentification is a paid mutator transaction binding the contract method 0xacc02cce.
//
// Solidity: function setIdentification(id bytes16, userId bytes16, idType string, name string, companyId bytes16, status uint8) returns()
func (_HoQuStorageI *HoQuStorageISession) SetIdentification(id [16]byte, userId [16]byte, idType string, name string, companyId [16]byte, status uint8) (*types.Transaction, error) {
	return _HoQuStorageI.Contract.SetIdentification(&_HoQuStorageI.TransactOpts, id, userId, idType, name, companyId, status)
}

// SetIdentification is a paid mutator transaction binding the contract method 0xacc02cce.
//
// Solidity: function setIdentification(id bytes16, userId bytes16, idType string, name string, companyId bytes16, status uint8) returns()
func (_HoQuStorageI *HoQuStorageITransactorSession) SetIdentification(id [16]byte, userId [16]byte, idType string, name string, companyId [16]byte, status uint8) (*types.Transaction, error) {
	return _HoQuStorageI.Contract.SetIdentification(&_HoQuStorageI.TransactOpts, id, userId, idType, name, companyId, status)
}

// SetNetwork is a paid mutator transaction binding the contract method 0x02884980.
//
// Solidity: function setNetwork(id bytes16, ownerId bytes16, name string, dataUrl string, status uint8) returns()
func (_HoQuStorageI *HoQuStorageITransactor) SetNetwork(opts *bind.TransactOpts, id [16]byte, ownerId [16]byte, name string, dataUrl string, status uint8) (*types.Transaction, error) {
	return _HoQuStorageI.contract.Transact(opts, "setNetwork", id, ownerId, name, dataUrl, status)
}

// SetNetwork is a paid mutator transaction binding the contract method 0x02884980.
//
// Solidity: function setNetwork(id bytes16, ownerId bytes16, name string, dataUrl string, status uint8) returns()
func (_HoQuStorageI *HoQuStorageISession) SetNetwork(id [16]byte, ownerId [16]byte, name string, dataUrl string, status uint8) (*types.Transaction, error) {
	return _HoQuStorageI.Contract.SetNetwork(&_HoQuStorageI.TransactOpts, id, ownerId, name, dataUrl, status)
}

// SetNetwork is a paid mutator transaction binding the contract method 0x02884980.
//
// Solidity: function setNetwork(id bytes16, ownerId bytes16, name string, dataUrl string, status uint8) returns()
func (_HoQuStorageI *HoQuStorageITransactorSession) SetNetwork(id [16]byte, ownerId [16]byte, name string, dataUrl string, status uint8) (*types.Transaction, error) {
	return _HoQuStorageI.Contract.SetNetwork(&_HoQuStorageI.TransactOpts, id, ownerId, name, dataUrl, status)
}

// SetOffer is a paid mutator transaction binding the contract method 0x1071d512.
//
// Solidity: function setOffer(id bytes16, ownerId bytes16, networkId bytes16, merchantId bytes16, payerAddress address, name string, dataUrl string, status uint8) returns()
func (_HoQuStorageI *HoQuStorageITransactor) SetOffer(opts *bind.TransactOpts, id [16]byte, ownerId [16]byte, networkId [16]byte, merchantId [16]byte, payerAddress common.Address, name string, dataUrl string, status uint8) (*types.Transaction, error) {
	return _HoQuStorageI.contract.Transact(opts, "setOffer", id, ownerId, networkId, merchantId, payerAddress, name, dataUrl, status)
}

// SetOffer is a paid mutator transaction binding the contract method 0x1071d512.
//
// Solidity: function setOffer(id bytes16, ownerId bytes16, networkId bytes16, merchantId bytes16, payerAddress address, name string, dataUrl string, status uint8) returns()
func (_HoQuStorageI *HoQuStorageISession) SetOffer(id [16]byte, ownerId [16]byte, networkId [16]byte, merchantId [16]byte, payerAddress common.Address, name string, dataUrl string, status uint8) (*types.Transaction, error) {
	return _HoQuStorageI.Contract.SetOffer(&_HoQuStorageI.TransactOpts, id, ownerId, networkId, merchantId, payerAddress, name, dataUrl, status)
}

// SetOffer is a paid mutator transaction binding the contract method 0x1071d512.
//
// Solidity: function setOffer(id bytes16, ownerId bytes16, networkId bytes16, merchantId bytes16, payerAddress address, name string, dataUrl string, status uint8) returns()
func (_HoQuStorageI *HoQuStorageITransactorSession) SetOffer(id [16]byte, ownerId [16]byte, networkId [16]byte, merchantId [16]byte, payerAddress common.Address, name string, dataUrl string, status uint8) (*types.Transaction, error) {
	return _HoQuStorageI.Contract.SetOffer(&_HoQuStorageI.TransactOpts, id, ownerId, networkId, merchantId, payerAddress, name, dataUrl, status)
}

// SetOfferTariffGroup is a paid mutator transaction binding the contract method 0x3e45850a.
//
// Solidity: function setOfferTariffGroup(id bytes16, num uint8, tariffGroupId bytes16) returns()
func (_HoQuStorageI *HoQuStorageITransactor) SetOfferTariffGroup(opts *bind.TransactOpts, id [16]byte, num uint8, tariffGroupId [16]byte) (*types.Transaction, error) {
	return _HoQuStorageI.contract.Transact(opts, "setOfferTariffGroup", id, num, tariffGroupId)
}

// SetOfferTariffGroup is a paid mutator transaction binding the contract method 0x3e45850a.
//
// Solidity: function setOfferTariffGroup(id bytes16, num uint8, tariffGroupId bytes16) returns()
func (_HoQuStorageI *HoQuStorageISession) SetOfferTariffGroup(id [16]byte, num uint8, tariffGroupId [16]byte) (*types.Transaction, error) {
	return _HoQuStorageI.Contract.SetOfferTariffGroup(&_HoQuStorageI.TransactOpts, id, num, tariffGroupId)
}

// SetOfferTariffGroup is a paid mutator transaction binding the contract method 0x3e45850a.
//
// Solidity: function setOfferTariffGroup(id bytes16, num uint8, tariffGroupId bytes16) returns()
func (_HoQuStorageI *HoQuStorageITransactorSession) SetOfferTariffGroup(id [16]byte, num uint8, tariffGroupId [16]byte) (*types.Transaction, error) {
	return _HoQuStorageI.Contract.SetOfferTariffGroup(&_HoQuStorageI.TransactOpts, id, num, tariffGroupId)
}

// SetStats is a paid mutator transaction binding the contract method 0xb1dbfe50.
//
// Solidity: function setStats(id bytes16, userId bytes16, rating uint256, volume uint256, members uint256, alfa uint256, beta uint256, status uint8) returns()
func (_HoQuStorageI *HoQuStorageITransactor) SetStats(opts *bind.TransactOpts, id [16]byte, userId [16]byte, rating *big.Int, volume *big.Int, members *big.Int, alfa *big.Int, beta *big.Int, status uint8) (*types.Transaction, error) {
	return _HoQuStorageI.contract.Transact(opts, "setStats", id, userId, rating, volume, members, alfa, beta, status)
}

// SetStats is a paid mutator transaction binding the contract method 0xb1dbfe50.
//
// Solidity: function setStats(id bytes16, userId bytes16, rating uint256, volume uint256, members uint256, alfa uint256, beta uint256, status uint8) returns()
func (_HoQuStorageI *HoQuStorageISession) SetStats(id [16]byte, userId [16]byte, rating *big.Int, volume *big.Int, members *big.Int, alfa *big.Int, beta *big.Int, status uint8) (*types.Transaction, error) {
	return _HoQuStorageI.Contract.SetStats(&_HoQuStorageI.TransactOpts, id, userId, rating, volume, members, alfa, beta, status)
}

// SetStats is a paid mutator transaction binding the contract method 0xb1dbfe50.
//
// Solidity: function setStats(id bytes16, userId bytes16, rating uint256, volume uint256, members uint256, alfa uint256, beta uint256, status uint8) returns()
func (_HoQuStorageI *HoQuStorageITransactorSession) SetStats(id [16]byte, userId [16]byte, rating *big.Int, volume *big.Int, members *big.Int, alfa *big.Int, beta *big.Int, status uint8) (*types.Transaction, error) {
	return _HoQuStorageI.Contract.SetStats(&_HoQuStorageI.TransactOpts, id, userId, rating, volume, members, alfa, beta, status)
}

// SetTariff is a paid mutator transaction binding the contract method 0x61fad57e.
//
// Solidity: function setTariff(id bytes16, tariffGroupId bytes16, name string, action string, calcMethod string, price uint256, status uint8) returns()
func (_HoQuStorageI *HoQuStorageITransactor) SetTariff(opts *bind.TransactOpts, id [16]byte, tariffGroupId [16]byte, name string, action string, calcMethod string, price *big.Int, status uint8) (*types.Transaction, error) {
	return _HoQuStorageI.contract.Transact(opts, "setTariff", id, tariffGroupId, name, action, calcMethod, price, status)
}

// SetTariff is a paid mutator transaction binding the contract method 0x61fad57e.
//
// Solidity: function setTariff(id bytes16, tariffGroupId bytes16, name string, action string, calcMethod string, price uint256, status uint8) returns()
func (_HoQuStorageI *HoQuStorageISession) SetTariff(id [16]byte, tariffGroupId [16]byte, name string, action string, calcMethod string, price *big.Int, status uint8) (*types.Transaction, error) {
	return _HoQuStorageI.Contract.SetTariff(&_HoQuStorageI.TransactOpts, id, tariffGroupId, name, action, calcMethod, price, status)
}

// SetTariff is a paid mutator transaction binding the contract method 0x61fad57e.
//
// Solidity: function setTariff(id bytes16, tariffGroupId bytes16, name string, action string, calcMethod string, price uint256, status uint8) returns()
func (_HoQuStorageI *HoQuStorageITransactorSession) SetTariff(id [16]byte, tariffGroupId [16]byte, name string, action string, calcMethod string, price *big.Int, status uint8) (*types.Transaction, error) {
	return _HoQuStorageI.Contract.SetTariff(&_HoQuStorageI.TransactOpts, id, tariffGroupId, name, action, calcMethod, price, status)
}

// SetTariffGroup is a paid mutator transaction binding the contract method 0xa7fbcf61.
//
// Solidity: function setTariffGroup(id bytes16, ownerId bytes16, name string, status uint8) returns()
func (_HoQuStorageI *HoQuStorageITransactor) SetTariffGroup(opts *bind.TransactOpts, id [16]byte, ownerId [16]byte, name string, status uint8) (*types.Transaction, error) {
	return _HoQuStorageI.contract.Transact(opts, "setTariffGroup", id, ownerId, name, status)
}

// SetTariffGroup is a paid mutator transaction binding the contract method 0xa7fbcf61.
//
// Solidity: function setTariffGroup(id bytes16, ownerId bytes16, name string, status uint8) returns()
func (_HoQuStorageI *HoQuStorageISession) SetTariffGroup(id [16]byte, ownerId [16]byte, name string, status uint8) (*types.Transaction, error) {
	return _HoQuStorageI.Contract.SetTariffGroup(&_HoQuStorageI.TransactOpts, id, ownerId, name, status)
}

// SetTariffGroup is a paid mutator transaction binding the contract method 0xa7fbcf61.
//
// Solidity: function setTariffGroup(id bytes16, ownerId bytes16, name string, status uint8) returns()
func (_HoQuStorageI *HoQuStorageITransactorSession) SetTariffGroup(id [16]byte, ownerId [16]byte, name string, status uint8) (*types.Transaction, error) {
	return _HoQuStorageI.Contract.SetTariffGroup(&_HoQuStorageI.TransactOpts, id, ownerId, name, status)
}

// SetTariffGroupTariff is a paid mutator transaction binding the contract method 0xfafe27d7.
//
// Solidity: function setTariffGroupTariff(id bytes16, num uint8, tariffId bytes16) returns()
func (_HoQuStorageI *HoQuStorageITransactor) SetTariffGroupTariff(opts *bind.TransactOpts, id [16]byte, num uint8, tariffId [16]byte) (*types.Transaction, error) {
	return _HoQuStorageI.contract.Transact(opts, "setTariffGroupTariff", id, num, tariffId)
}

// SetTariffGroupTariff is a paid mutator transaction binding the contract method 0xfafe27d7.
//
// Solidity: function setTariffGroupTariff(id bytes16, num uint8, tariffId bytes16) returns()
func (_HoQuStorageI *HoQuStorageISession) SetTariffGroupTariff(id [16]byte, num uint8, tariffId [16]byte) (*types.Transaction, error) {
	return _HoQuStorageI.Contract.SetTariffGroupTariff(&_HoQuStorageI.TransactOpts, id, num, tariffId)
}

// SetTariffGroupTariff is a paid mutator transaction binding the contract method 0xfafe27d7.
//
// Solidity: function setTariffGroupTariff(id bytes16, num uint8, tariffId bytes16) returns()
func (_HoQuStorageI *HoQuStorageITransactorSession) SetTariffGroupTariff(id [16]byte, num uint8, tariffId [16]byte) (*types.Transaction, error) {
	return _HoQuStorageI.Contract.SetTariffGroupTariff(&_HoQuStorageI.TransactOpts, id, num, tariffId)
}

// SetTracker is a paid mutator transaction binding the contract method 0xc4d4d887.
//
// Solidity: function setTracker(id bytes16, ownerId bytes16, networkId bytes16, name string, dataUrl string, status uint8) returns()
func (_HoQuStorageI *HoQuStorageITransactor) SetTracker(opts *bind.TransactOpts, id [16]byte, ownerId [16]byte, networkId [16]byte, name string, dataUrl string, status uint8) (*types.Transaction, error) {
	return _HoQuStorageI.contract.Transact(opts, "setTracker", id, ownerId, networkId, name, dataUrl, status)
}

// SetTracker is a paid mutator transaction binding the contract method 0xc4d4d887.
//
// Solidity: function setTracker(id bytes16, ownerId bytes16, networkId bytes16, name string, dataUrl string, status uint8) returns()
func (_HoQuStorageI *HoQuStorageISession) SetTracker(id [16]byte, ownerId [16]byte, networkId [16]byte, name string, dataUrl string, status uint8) (*types.Transaction, error) {
	return _HoQuStorageI.Contract.SetTracker(&_HoQuStorageI.TransactOpts, id, ownerId, networkId, name, dataUrl, status)
}

// SetTracker is a paid mutator transaction binding the contract method 0xc4d4d887.
//
// Solidity: function setTracker(id bytes16, ownerId bytes16, networkId bytes16, name string, dataUrl string, status uint8) returns()
func (_HoQuStorageI *HoQuStorageITransactorSession) SetTracker(id [16]byte, ownerId [16]byte, networkId [16]byte, name string, dataUrl string, status uint8) (*types.Transaction, error) {
	return _HoQuStorageI.Contract.SetTracker(&_HoQuStorageI.TransactOpts, id, ownerId, networkId, name, dataUrl, status)
}

// SetUser is a paid mutator transaction binding the contract method 0x253a9b9e.
//
// Solidity: function setUser(id bytes16, role string, ownerAddress address, kycLevel uint8, pubKey string, status uint8) returns()
func (_HoQuStorageI *HoQuStorageITransactor) SetUser(opts *bind.TransactOpts, id [16]byte, role string, ownerAddress common.Address, kycLevel uint8, pubKey string, status uint8) (*types.Transaction, error) {
	return _HoQuStorageI.contract.Transact(opts, "setUser", id, role, ownerAddress, kycLevel, pubKey, status)
}

// SetUser is a paid mutator transaction binding the contract method 0x253a9b9e.
//
// Solidity: function setUser(id bytes16, role string, ownerAddress address, kycLevel uint8, pubKey string, status uint8) returns()
func (_HoQuStorageI *HoQuStorageISession) SetUser(id [16]byte, role string, ownerAddress common.Address, kycLevel uint8, pubKey string, status uint8) (*types.Transaction, error) {
	return _HoQuStorageI.Contract.SetUser(&_HoQuStorageI.TransactOpts, id, role, ownerAddress, kycLevel, pubKey, status)
}

// SetUser is a paid mutator transaction binding the contract method 0x253a9b9e.
//
// Solidity: function setUser(id bytes16, role string, ownerAddress address, kycLevel uint8, pubKey string, status uint8) returns()
func (_HoQuStorageI *HoQuStorageITransactorSession) SetUser(id [16]byte, role string, ownerAddress common.Address, kycLevel uint8, pubKey string, status uint8) (*types.Transaction, error) {
	return _HoQuStorageI.Contract.SetUser(&_HoQuStorageI.TransactOpts, id, role, ownerAddress, kycLevel, pubKey, status)
}

// HoQuStorageSchemaABI is the input ABI used to generate the binding from.
const HoQuStorageSchemaABI = "[]"

// HoQuStorageSchemaBin is the compiled bytecode used for deploying new contracts.
const HoQuStorageSchemaBin = `0x6080604052348015600f57600080fd5b50603580601d6000396000f3006080604052600080fd00a165627a7a72305820d15b2318427fdcd8b12bc2a3fbf6043aba3ce71defebd94ad594ad4755bbc6cb0029`

// DeployHoQuStorageSchema deploys a new Ethereum contract, binding an instance of HoQuStorageSchema to it.
func DeployHoQuStorageSchema(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *HoQuStorageSchema, error) {
	parsed, err := abi.JSON(strings.NewReader(HoQuStorageSchemaABI))
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	address, tx, contract, err := bind.DeployContract(auth, parsed, common.FromHex(HoQuStorageSchemaBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &HoQuStorageSchema{HoQuStorageSchemaCaller: HoQuStorageSchemaCaller{contract: contract}, HoQuStorageSchemaTransactor: HoQuStorageSchemaTransactor{contract: contract}, HoQuStorageSchemaFilterer: HoQuStorageSchemaFilterer{contract: contract}}, nil
}

// HoQuStorageSchema is an auto generated Go binding around an Ethereum contract.
type HoQuStorageSchema struct {
	HoQuStorageSchemaCaller     // Read-only binding to the contract
	HoQuStorageSchemaTransactor // Write-only binding to the contract
	HoQuStorageSchemaFilterer   // Log filterer for contract events
}

// HoQuStorageSchemaCaller is an auto generated read-only Go binding around an Ethereum contract.
type HoQuStorageSchemaCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// HoQuStorageSchemaTransactor is an auto generated write-only Go binding around an Ethereum contract.
type HoQuStorageSchemaTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// HoQuStorageSchemaFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type HoQuStorageSchemaFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// HoQuStorageSchemaSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type HoQuStorageSchemaSession struct {
	Contract     *HoQuStorageSchema // Generic contract binding to set the session for
	CallOpts     bind.CallOpts      // Call options to use throughout this session
	TransactOpts bind.TransactOpts  // Transaction auth options to use throughout this session
}

// HoQuStorageSchemaCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type HoQuStorageSchemaCallerSession struct {
	Contract *HoQuStorageSchemaCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts            // Call options to use throughout this session
}

// HoQuStorageSchemaTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type HoQuStorageSchemaTransactorSession struct {
	Contract     *HoQuStorageSchemaTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts            // Transaction auth options to use throughout this session
}

// HoQuStorageSchemaRaw is an auto generated low-level Go binding around an Ethereum contract.
type HoQuStorageSchemaRaw struct {
	Contract *HoQuStorageSchema // Generic contract binding to access the raw methods on
}

// HoQuStorageSchemaCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type HoQuStorageSchemaCallerRaw struct {
	Contract *HoQuStorageSchemaCaller // Generic read-only contract binding to access the raw methods on
}

// HoQuStorageSchemaTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type HoQuStorageSchemaTransactorRaw struct {
	Contract *HoQuStorageSchemaTransactor // Generic write-only contract binding to access the raw methods on
}

// NewHoQuStorageSchema creates a new instance of HoQuStorageSchema, bound to a specific deployed contract.
func NewHoQuStorageSchema(address common.Address, backend bind.ContractBackend) (*HoQuStorageSchema, error) {
	contract, err := bindHoQuStorageSchema(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &HoQuStorageSchema{HoQuStorageSchemaCaller: HoQuStorageSchemaCaller{contract: contract}, HoQuStorageSchemaTransactor: HoQuStorageSchemaTransactor{contract: contract}, HoQuStorageSchemaFilterer: HoQuStorageSchemaFilterer{contract: contract}}, nil
}

// NewHoQuStorageSchemaCaller creates a new read-only instance of HoQuStorageSchema, bound to a specific deployed contract.
func NewHoQuStorageSchemaCaller(address common.Address, caller bind.ContractCaller) (*HoQuStorageSchemaCaller, error) {
	contract, err := bindHoQuStorageSchema(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &HoQuStorageSchemaCaller{contract: contract}, nil
}

// NewHoQuStorageSchemaTransactor creates a new write-only instance of HoQuStorageSchema, bound to a specific deployed contract.
func NewHoQuStorageSchemaTransactor(address common.Address, transactor bind.ContractTransactor) (*HoQuStorageSchemaTransactor, error) {
	contract, err := bindHoQuStorageSchema(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &HoQuStorageSchemaTransactor{contract: contract}, nil
}

// NewHoQuStorageSchemaFilterer creates a new log filterer instance of HoQuStorageSchema, bound to a specific deployed contract.
func NewHoQuStorageSchemaFilterer(address common.Address, filterer bind.ContractFilterer) (*HoQuStorageSchemaFilterer, error) {
	contract, err := bindHoQuStorageSchema(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &HoQuStorageSchemaFilterer{contract: contract}, nil
}

// bindHoQuStorageSchema binds a generic wrapper to an already deployed contract.
func bindHoQuStorageSchema(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(HoQuStorageSchemaABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_HoQuStorageSchema *HoQuStorageSchemaRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _HoQuStorageSchema.Contract.HoQuStorageSchemaCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_HoQuStorageSchema *HoQuStorageSchemaRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _HoQuStorageSchema.Contract.HoQuStorageSchemaTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_HoQuStorageSchema *HoQuStorageSchemaRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _HoQuStorageSchema.Contract.HoQuStorageSchemaTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_HoQuStorageSchema *HoQuStorageSchemaCallerRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _HoQuStorageSchema.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_HoQuStorageSchema *HoQuStorageSchemaTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _HoQuStorageSchema.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_HoQuStorageSchema *HoQuStorageSchemaTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _HoQuStorageSchema.Contract.contract.Transact(opts, method, params...)
}

// HoQuTransactorABI is the input ABI used to generate the binding from.
const HoQuTransactorABI = "[{\"constant\":false,\"inputs\":[{\"name\":\"tokenAddress\",\"type\":\"address\"}],\"name\":\"setTokenAddress\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"config\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"configAddress\",\"type\":\"address\"}],\"name\":\"setConfigAddress\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"approve\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"beneficiaryAddress\",\"type\":\"address\"},{\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"send\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"payerAddress\",\"type\":\"address\"},{\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"withdraw\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"token\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"name\":\"configAddress\",\"type\":\"address\"},{\"name\":\"tokenAddress\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"payerAddress\",\"type\":\"address\"},{\"indexed\":false,\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"TokenWithdrew\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"beneficiaryAddress\",\"type\":\"address\"},{\"indexed\":false,\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"TokenSent\",\"type\":\"event\"}]"

// HoQuTransactorBin is the compiled bytecode used for deploying new contracts.
const HoQuTransactorBin = `0x608060405234801561001057600080fd5b5060405160408061079a83398101604052805160209091015160008054600160a060020a03938416600160a060020a03199182161790915560018054939092169216919091179055610733806100676000396000f3006080604052600436106100825763ffffffff7c010000000000000000000000000000000000000000000000000000000060003504166326a4e8d2811461008757806379502c55146100aa57806383a12de9146100db578063b759f954146100fc578063d0679d3414610114578063f3fef3a314610138578063fc0c546a1461015c575b600080fd5b34801561009357600080fd5b506100a8600160a060020a0360043516610171565b005b3480156100b657600080fd5b506100bf610228565b60408051600160a060020a039092168252519081900360200190f35b3480156100e757600080fd5b506100a8600160a060020a0360043516610237565b34801561010857600080fd5b506100a86004356102ee565b34801561012057600080fd5b506100a8600160a060020a0360043516602435610426565b34801561014457600080fd5b506100a8600160a060020a036004351660243561058c565b34801561016857600080fd5b506100bf6106f8565b600080546040805160e060020a63babcc5390281523360048201529051600160a060020a039092169263babcc539926024808401936020939083900390910190829087803b1580156101c257600080fd5b505af11580156101d6573d6000803e3d6000fd5b505050506040513d60208110156101ec57600080fd5b505115156101f957600080fd5b6001805473ffffffffffffffffffffffffffffffffffffffff1916600160a060020a0392909216919091179055565b600054600160a060020a031681565b600080546040805160e060020a63babcc5390281523360048201529051600160a060020a039092169263babcc539926024808401936020939083900390910190829087803b15801561028857600080fd5b505af115801561029c573d6000803e3d6000fd5b505050506040513d60208110156102b257600080fd5b505115156102bf57600080fd5b6000805473ffffffffffffffffffffffffffffffffffffffff1916600160a060020a0392909216919091179055565b600154604080517f095ea7b30000000000000000000000000000000000000000000000000000000081523060048201526000602482018190529151600160a060020a039093169263095ea7b392604480840193602093929083900390910190829087803b15801561035e57600080fd5b505af1158015610372573d6000803e3d6000fd5b505050506040513d602081101561038857600080fd5b5050600154604080517f095ea7b3000000000000000000000000000000000000000000000000000000008152306004820152602481018490529051600160a060020a039092169163095ea7b3916044808201926020929091908290030181600087803b1580156103f757600080fd5b505af115801561040b573d6000803e3d6000fd5b505050506040513d602081101561042157600080fd5b505050565b600080546040805160e060020a63babcc5390281523360048201529051600160a060020a039092169263babcc539926024808401936020939083900390910190829087803b15801561047757600080fd5b505af115801561048b573d6000803e3d6000fd5b505050506040513d60208110156104a157600080fd5b505115156104ae57600080fd5b600154604080517fa9059cbb000000000000000000000000000000000000000000000000000000008152600160a060020a038581166004830152602482018590529151919092169163a9059cbb9160448083019260209291908290030181600087803b15801561051d57600080fd5b505af1158015610531573d6000803e3d6000fd5b505050506040513d602081101561054757600080fd5b5050604080518281529051600160a060020a038416917fd4ff88d569a7ad2bfd6b17da9dbe82e2ccc016fd0051e08ff1ad1e6fe8fc9b55919081900360200190a25050565b600080546040805160e060020a63babcc5390281523360048201529051600160a060020a039092169263babcc539926024808401936020939083900390910190829087803b1580156105dd57600080fd5b505af11580156105f1573d6000803e3d6000fd5b505050506040513d602081101561060757600080fd5b5051151561061457600080fd5b600154604080517f23b872dd000000000000000000000000000000000000000000000000000000008152600160a060020a03858116600483015230602483015260448201859052915191909216916323b872dd9160648083019260209291908290030181600087803b15801561068957600080fd5b505af115801561069d573d6000803e3d6000fd5b505050506040513d60208110156106b357600080fd5b5050604080518281529051600160a060020a038416917fcd6f1c60d74b50255e52371fd4b3d5678cd8cc9314a11389b82c41bf08ee28bb919081900360200190a25050565b600154600160a060020a0316815600a165627a7a723058208371be18a1549980b8678f5a5a7703d0efa528b11c8987a4832722de90be56780029`

// DeployHoQuTransactor deploys a new Ethereum contract, binding an instance of HoQuTransactor to it.
func DeployHoQuTransactor(auth *bind.TransactOpts, backend bind.ContractBackend, configAddress common.Address, tokenAddress common.Address) (common.Address, *types.Transaction, *HoQuTransactor, error) {
	parsed, err := abi.JSON(strings.NewReader(HoQuTransactorABI))
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	address, tx, contract, err := bind.DeployContract(auth, parsed, common.FromHex(HoQuTransactorBin), backend, configAddress, tokenAddress)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &HoQuTransactor{HoQuTransactorCaller: HoQuTransactorCaller{contract: contract}, HoQuTransactorTransactor: HoQuTransactorTransactor{contract: contract}, HoQuTransactorFilterer: HoQuTransactorFilterer{contract: contract}}, nil
}

// HoQuTransactor is an auto generated Go binding around an Ethereum contract.
type HoQuTransactor struct {
	HoQuTransactorCaller     // Read-only binding to the contract
	HoQuTransactorTransactor // Write-only binding to the contract
	HoQuTransactorFilterer   // Log filterer for contract events
}

// HoQuTransactorCaller is an auto generated read-only Go binding around an Ethereum contract.
type HoQuTransactorCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// HoQuTransactorTransactor is an auto generated write-only Go binding around an Ethereum contract.
type HoQuTransactorTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// HoQuTransactorFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type HoQuTransactorFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// HoQuTransactorSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type HoQuTransactorSession struct {
	Contract     *HoQuTransactor   // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// HoQuTransactorCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type HoQuTransactorCallerSession struct {
	Contract *HoQuTransactorCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts         // Call options to use throughout this session
}

// HoQuTransactorTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type HoQuTransactorTransactorSession struct {
	Contract     *HoQuTransactorTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts         // Transaction auth options to use throughout this session
}

// HoQuTransactorRaw is an auto generated low-level Go binding around an Ethereum contract.
type HoQuTransactorRaw struct {
	Contract *HoQuTransactor // Generic contract binding to access the raw methods on
}

// HoQuTransactorCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type HoQuTransactorCallerRaw struct {
	Contract *HoQuTransactorCaller // Generic read-only contract binding to access the raw methods on
}

// HoQuTransactorTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type HoQuTransactorTransactorRaw struct {
	Contract *HoQuTransactorTransactor // Generic write-only contract binding to access the raw methods on
}

// NewHoQuTransactor creates a new instance of HoQuTransactor, bound to a specific deployed contract.
func NewHoQuTransactor(address common.Address, backend bind.ContractBackend) (*HoQuTransactor, error) {
	contract, err := bindHoQuTransactor(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &HoQuTransactor{HoQuTransactorCaller: HoQuTransactorCaller{contract: contract}, HoQuTransactorTransactor: HoQuTransactorTransactor{contract: contract}, HoQuTransactorFilterer: HoQuTransactorFilterer{contract: contract}}, nil
}

// NewHoQuTransactorCaller creates a new read-only instance of HoQuTransactor, bound to a specific deployed contract.
func NewHoQuTransactorCaller(address common.Address, caller bind.ContractCaller) (*HoQuTransactorCaller, error) {
	contract, err := bindHoQuTransactor(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &HoQuTransactorCaller{contract: contract}, nil
}

// NewHoQuTransactorTransactor creates a new write-only instance of HoQuTransactor, bound to a specific deployed contract.
func NewHoQuTransactorTransactor(address common.Address, transactor bind.ContractTransactor) (*HoQuTransactorTransactor, error) {
	contract, err := bindHoQuTransactor(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &HoQuTransactorTransactor{contract: contract}, nil
}

// NewHoQuTransactorFilterer creates a new log filterer instance of HoQuTransactor, bound to a specific deployed contract.
func NewHoQuTransactorFilterer(address common.Address, filterer bind.ContractFilterer) (*HoQuTransactorFilterer, error) {
	contract, err := bindHoQuTransactor(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &HoQuTransactorFilterer{contract: contract}, nil
}

// bindHoQuTransactor binds a generic wrapper to an already deployed contract.
func bindHoQuTransactor(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(HoQuTransactorABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_HoQuTransactor *HoQuTransactorRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _HoQuTransactor.Contract.HoQuTransactorCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_HoQuTransactor *HoQuTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _HoQuTransactor.Contract.HoQuTransactorTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_HoQuTransactor *HoQuTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _HoQuTransactor.Contract.HoQuTransactorTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_HoQuTransactor *HoQuTransactorCallerRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _HoQuTransactor.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_HoQuTransactor *HoQuTransactorTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _HoQuTransactor.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_HoQuTransactor *HoQuTransactorTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _HoQuTransactor.Contract.contract.Transact(opts, method, params...)
}

// Config is a free data retrieval call binding the contract method 0x79502c55.
//
// Solidity: function config() constant returns(address)
func (_HoQuTransactor *HoQuTransactorCaller) Config(opts *bind.CallOpts) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _HoQuTransactor.contract.Call(opts, out, "config")
	return *ret0, err
}

// Config is a free data retrieval call binding the contract method 0x79502c55.
//
// Solidity: function config() constant returns(address)
func (_HoQuTransactor *HoQuTransactorSession) Config() (common.Address, error) {
	return _HoQuTransactor.Contract.Config(&_HoQuTransactor.CallOpts)
}

// Config is a free data retrieval call binding the contract method 0x79502c55.
//
// Solidity: function config() constant returns(address)
func (_HoQuTransactor *HoQuTransactorCallerSession) Config() (common.Address, error) {
	return _HoQuTransactor.Contract.Config(&_HoQuTransactor.CallOpts)
}

// Token is a free data retrieval call binding the contract method 0xfc0c546a.
//
// Solidity: function token() constant returns(address)
func (_HoQuTransactor *HoQuTransactorCaller) Token(opts *bind.CallOpts) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _HoQuTransactor.contract.Call(opts, out, "token")
	return *ret0, err
}

// Token is a free data retrieval call binding the contract method 0xfc0c546a.
//
// Solidity: function token() constant returns(address)
func (_HoQuTransactor *HoQuTransactorSession) Token() (common.Address, error) {
	return _HoQuTransactor.Contract.Token(&_HoQuTransactor.CallOpts)
}

// Token is a free data retrieval call binding the contract method 0xfc0c546a.
//
// Solidity: function token() constant returns(address)
func (_HoQuTransactor *HoQuTransactorCallerSession) Token() (common.Address, error) {
	return _HoQuTransactor.Contract.Token(&_HoQuTransactor.CallOpts)
}

// Approve is a paid mutator transaction binding the contract method 0xb759f954.
//
// Solidity: function approve(amount uint256) returns()
func (_HoQuTransactor *HoQuTransactorTransactor) Approve(opts *bind.TransactOpts, amount *big.Int) (*types.Transaction, error) {
	return _HoQuTransactor.contract.Transact(opts, "approve", amount)
}

// Approve is a paid mutator transaction binding the contract method 0xb759f954.
//
// Solidity: function approve(amount uint256) returns()
func (_HoQuTransactor *HoQuTransactorSession) Approve(amount *big.Int) (*types.Transaction, error) {
	return _HoQuTransactor.Contract.Approve(&_HoQuTransactor.TransactOpts, amount)
}

// Approve is a paid mutator transaction binding the contract method 0xb759f954.
//
// Solidity: function approve(amount uint256) returns()
func (_HoQuTransactor *HoQuTransactorTransactorSession) Approve(amount *big.Int) (*types.Transaction, error) {
	return _HoQuTransactor.Contract.Approve(&_HoQuTransactor.TransactOpts, amount)
}

// Send is a paid mutator transaction binding the contract method 0xd0679d34.
//
// Solidity: function send(beneficiaryAddress address, amount uint256) returns()
func (_HoQuTransactor *HoQuTransactorTransactor) Send(opts *bind.TransactOpts, beneficiaryAddress common.Address, amount *big.Int) (*types.Transaction, error) {
	return _HoQuTransactor.contract.Transact(opts, "send", beneficiaryAddress, amount)
}

// Send is a paid mutator transaction binding the contract method 0xd0679d34.
//
// Solidity: function send(beneficiaryAddress address, amount uint256) returns()
func (_HoQuTransactor *HoQuTransactorSession) Send(beneficiaryAddress common.Address, amount *big.Int) (*types.Transaction, error) {
	return _HoQuTransactor.Contract.Send(&_HoQuTransactor.TransactOpts, beneficiaryAddress, amount)
}

// Send is a paid mutator transaction binding the contract method 0xd0679d34.
//
// Solidity: function send(beneficiaryAddress address, amount uint256) returns()
func (_HoQuTransactor *HoQuTransactorTransactorSession) Send(beneficiaryAddress common.Address, amount *big.Int) (*types.Transaction, error) {
	return _HoQuTransactor.Contract.Send(&_HoQuTransactor.TransactOpts, beneficiaryAddress, amount)
}

// SetConfigAddress is a paid mutator transaction binding the contract method 0x83a12de9.
//
// Solidity: function setConfigAddress(configAddress address) returns()
func (_HoQuTransactor *HoQuTransactorTransactor) SetConfigAddress(opts *bind.TransactOpts, configAddress common.Address) (*types.Transaction, error) {
	return _HoQuTransactor.contract.Transact(opts, "setConfigAddress", configAddress)
}

// SetConfigAddress is a paid mutator transaction binding the contract method 0x83a12de9.
//
// Solidity: function setConfigAddress(configAddress address) returns()
func (_HoQuTransactor *HoQuTransactorSession) SetConfigAddress(configAddress common.Address) (*types.Transaction, error) {
	return _HoQuTransactor.Contract.SetConfigAddress(&_HoQuTransactor.TransactOpts, configAddress)
}

// SetConfigAddress is a paid mutator transaction binding the contract method 0x83a12de9.
//
// Solidity: function setConfigAddress(configAddress address) returns()
func (_HoQuTransactor *HoQuTransactorTransactorSession) SetConfigAddress(configAddress common.Address) (*types.Transaction, error) {
	return _HoQuTransactor.Contract.SetConfigAddress(&_HoQuTransactor.TransactOpts, configAddress)
}

// SetTokenAddress is a paid mutator transaction binding the contract method 0x26a4e8d2.
//
// Solidity: function setTokenAddress(tokenAddress address) returns()
func (_HoQuTransactor *HoQuTransactorTransactor) SetTokenAddress(opts *bind.TransactOpts, tokenAddress common.Address) (*types.Transaction, error) {
	return _HoQuTransactor.contract.Transact(opts, "setTokenAddress", tokenAddress)
}

// SetTokenAddress is a paid mutator transaction binding the contract method 0x26a4e8d2.
//
// Solidity: function setTokenAddress(tokenAddress address) returns()
func (_HoQuTransactor *HoQuTransactorSession) SetTokenAddress(tokenAddress common.Address) (*types.Transaction, error) {
	return _HoQuTransactor.Contract.SetTokenAddress(&_HoQuTransactor.TransactOpts, tokenAddress)
}

// SetTokenAddress is a paid mutator transaction binding the contract method 0x26a4e8d2.
//
// Solidity: function setTokenAddress(tokenAddress address) returns()
func (_HoQuTransactor *HoQuTransactorTransactorSession) SetTokenAddress(tokenAddress common.Address) (*types.Transaction, error) {
	return _HoQuTransactor.Contract.SetTokenAddress(&_HoQuTransactor.TransactOpts, tokenAddress)
}

// Withdraw is a paid mutator transaction binding the contract method 0xf3fef3a3.
//
// Solidity: function withdraw(payerAddress address, amount uint256) returns()
func (_HoQuTransactor *HoQuTransactorTransactor) Withdraw(opts *bind.TransactOpts, payerAddress common.Address, amount *big.Int) (*types.Transaction, error) {
	return _HoQuTransactor.contract.Transact(opts, "withdraw", payerAddress, amount)
}

// Withdraw is a paid mutator transaction binding the contract method 0xf3fef3a3.
//
// Solidity: function withdraw(payerAddress address, amount uint256) returns()
func (_HoQuTransactor *HoQuTransactorSession) Withdraw(payerAddress common.Address, amount *big.Int) (*types.Transaction, error) {
	return _HoQuTransactor.Contract.Withdraw(&_HoQuTransactor.TransactOpts, payerAddress, amount)
}

// Withdraw is a paid mutator transaction binding the contract method 0xf3fef3a3.
//
// Solidity: function withdraw(payerAddress address, amount uint256) returns()
func (_HoQuTransactor *HoQuTransactorTransactorSession) Withdraw(payerAddress common.Address, amount *big.Int) (*types.Transaction, error) {
	return _HoQuTransactor.Contract.Withdraw(&_HoQuTransactor.TransactOpts, payerAddress, amount)
}

// HoQuTransactorTokenSentIterator is returned from FilterTokenSent and is used to iterate over the raw logs and unpacked data for TokenSent events raised by the HoQuTransactor contract.
type HoQuTransactorTokenSentIterator struct {
	Event *HoQuTransactorTokenSent // Event containing the contract specifics and raw log

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
func (it *HoQuTransactorTokenSentIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(HoQuTransactorTokenSent)
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
		it.Event = new(HoQuTransactorTokenSent)
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
func (it *HoQuTransactorTokenSentIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *HoQuTransactorTokenSentIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// HoQuTransactorTokenSent represents a TokenSent event raised by the HoQuTransactor contract.
type HoQuTransactorTokenSent struct {
	BeneficiaryAddress common.Address
	Amount             *big.Int
	Raw                types.Log // Blockchain specific contextual infos
}

// FilterTokenSent is a free log retrieval operation binding the contract event 0xd4ff88d569a7ad2bfd6b17da9dbe82e2ccc016fd0051e08ff1ad1e6fe8fc9b55.
//
// Solidity: event TokenSent(beneficiaryAddress indexed address, amount uint256)
func (_HoQuTransactor *HoQuTransactorFilterer) FilterTokenSent(opts *bind.FilterOpts, beneficiaryAddress []common.Address) (*HoQuTransactorTokenSentIterator, error) {

	var beneficiaryAddressRule []interface{}
	for _, beneficiaryAddressItem := range beneficiaryAddress {
		beneficiaryAddressRule = append(beneficiaryAddressRule, beneficiaryAddressItem)
	}

	logs, sub, err := _HoQuTransactor.contract.FilterLogs(opts, "TokenSent", beneficiaryAddressRule)
	if err != nil {
		return nil, err
	}
	return &HoQuTransactorTokenSentIterator{contract: _HoQuTransactor.contract, event: "TokenSent", logs: logs, sub: sub}, nil
}

// WatchTokenSent is a free log subscription operation binding the contract event 0xd4ff88d569a7ad2bfd6b17da9dbe82e2ccc016fd0051e08ff1ad1e6fe8fc9b55.
//
// Solidity: event TokenSent(beneficiaryAddress indexed address, amount uint256)
func (_HoQuTransactor *HoQuTransactorFilterer) WatchTokenSent(opts *bind.WatchOpts, sink chan<- *HoQuTransactorTokenSent, beneficiaryAddress []common.Address) (event.Subscription, error) {

	var beneficiaryAddressRule []interface{}
	for _, beneficiaryAddressItem := range beneficiaryAddress {
		beneficiaryAddressRule = append(beneficiaryAddressRule, beneficiaryAddressItem)
	}

	logs, sub, err := _HoQuTransactor.contract.WatchLogs(opts, "TokenSent", beneficiaryAddressRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(HoQuTransactorTokenSent)
				if err := _HoQuTransactor.contract.UnpackLog(event, "TokenSent", log); err != nil {
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

// HoQuTransactorTokenWithdrewIterator is returned from FilterTokenWithdrew and is used to iterate over the raw logs and unpacked data for TokenWithdrew events raised by the HoQuTransactor contract.
type HoQuTransactorTokenWithdrewIterator struct {
	Event *HoQuTransactorTokenWithdrew // Event containing the contract specifics and raw log

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
func (it *HoQuTransactorTokenWithdrewIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(HoQuTransactorTokenWithdrew)
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
		it.Event = new(HoQuTransactorTokenWithdrew)
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
func (it *HoQuTransactorTokenWithdrewIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *HoQuTransactorTokenWithdrewIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// HoQuTransactorTokenWithdrew represents a TokenWithdrew event raised by the HoQuTransactor contract.
type HoQuTransactorTokenWithdrew struct {
	PayerAddress common.Address
	Amount       *big.Int
	Raw          types.Log // Blockchain specific contextual infos
}

// FilterTokenWithdrew is a free log retrieval operation binding the contract event 0xcd6f1c60d74b50255e52371fd4b3d5678cd8cc9314a11389b82c41bf08ee28bb.
//
// Solidity: event TokenWithdrew(payerAddress indexed address, amount uint256)
func (_HoQuTransactor *HoQuTransactorFilterer) FilterTokenWithdrew(opts *bind.FilterOpts, payerAddress []common.Address) (*HoQuTransactorTokenWithdrewIterator, error) {

	var payerAddressRule []interface{}
	for _, payerAddressItem := range payerAddress {
		payerAddressRule = append(payerAddressRule, payerAddressItem)
	}

	logs, sub, err := _HoQuTransactor.contract.FilterLogs(opts, "TokenWithdrew", payerAddressRule)
	if err != nil {
		return nil, err
	}
	return &HoQuTransactorTokenWithdrewIterator{contract: _HoQuTransactor.contract, event: "TokenWithdrew", logs: logs, sub: sub}, nil
}

// WatchTokenWithdrew is a free log subscription operation binding the contract event 0xcd6f1c60d74b50255e52371fd4b3d5678cd8cc9314a11389b82c41bf08ee28bb.
//
// Solidity: event TokenWithdrew(payerAddress indexed address, amount uint256)
func (_HoQuTransactor *HoQuTransactorFilterer) WatchTokenWithdrew(opts *bind.WatchOpts, sink chan<- *HoQuTransactorTokenWithdrew, payerAddress []common.Address) (event.Subscription, error) {

	var payerAddressRule []interface{}
	for _, payerAddressItem := range payerAddress {
		payerAddressRule = append(payerAddressRule, payerAddressItem)
	}

	logs, sub, err := _HoQuTransactor.contract.WatchLogs(opts, "TokenWithdrew", payerAddressRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(HoQuTransactorTokenWithdrew)
				if err := _HoQuTransactor.contract.UnpackLog(event, "TokenWithdrew", log); err != nil {
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

// OwnableABI is the input ABI used to generate the binding from.
const OwnableABI = "[{\"constant\":false,\"inputs\":[],\"name\":\"renounceOwnership\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"transferOwnership\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"previousOwner\",\"type\":\"address\"}],\"name\":\"OwnershipRenounced\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"previousOwner\",\"type\":\"address\"},{\"indexed\":true,\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"OwnershipTransferred\",\"type\":\"event\"}]"

// OwnableBin is the compiled bytecode used for deploying new contracts.
const OwnableBin = `0x608060405234801561001057600080fd5b5060008054600160a060020a031916331790556101ff806100326000396000f3006080604052600436106100565763ffffffff7c0100000000000000000000000000000000000000000000000000000000600035041663715018a6811461005b5780638da5cb5b14610072578063f2fde38b146100a3575b600080fd5b34801561006757600080fd5b506100706100c4565b005b34801561007e57600080fd5b50610087610130565b60408051600160a060020a039092168252519081900360200190f35b3480156100af57600080fd5b50610070600160a060020a036004351661013f565b600054600160a060020a031633146100db57600080fd5b60008054604051600160a060020a03909116917ff8df31144d9c2f0f6b59d69b8b98abd5459d07f2742c4df920b25aae33c6482091a26000805473ffffffffffffffffffffffffffffffffffffffff19169055565b600054600160a060020a031681565b600054600160a060020a0316331461015657600080fd5b600160a060020a038116151561016b57600080fd5b60008054604051600160a060020a03808516939216917f8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e091a36000805473ffffffffffffffffffffffffffffffffffffffff1916600160a060020a03929092169190911790555600a165627a7a72305820006f36e1b0f857fcaad38f1979ea697b9318680b638cabd8ba8dbba2734925910029`

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
	return address, tx, &Ownable{OwnableCaller: OwnableCaller{contract: contract}, OwnableTransactor: OwnableTransactor{contract: contract}, OwnableFilterer: OwnableFilterer{contract: contract}}, nil
}

// Ownable is an auto generated Go binding around an Ethereum contract.
type Ownable struct {
	OwnableCaller     // Read-only binding to the contract
	OwnableTransactor // Write-only binding to the contract
	OwnableFilterer   // Log filterer for contract events
}

// OwnableCaller is an auto generated read-only Go binding around an Ethereum contract.
type OwnableCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// OwnableTransactor is an auto generated write-only Go binding around an Ethereum contract.
type OwnableTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// OwnableFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type OwnableFilterer struct {
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
	contract, err := bindOwnable(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Ownable{OwnableCaller: OwnableCaller{contract: contract}, OwnableTransactor: OwnableTransactor{contract: contract}, OwnableFilterer: OwnableFilterer{contract: contract}}, nil
}

// NewOwnableCaller creates a new read-only instance of Ownable, bound to a specific deployed contract.
func NewOwnableCaller(address common.Address, caller bind.ContractCaller) (*OwnableCaller, error) {
	contract, err := bindOwnable(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &OwnableCaller{contract: contract}, nil
}

// NewOwnableTransactor creates a new write-only instance of Ownable, bound to a specific deployed contract.
func NewOwnableTransactor(address common.Address, transactor bind.ContractTransactor) (*OwnableTransactor, error) {
	contract, err := bindOwnable(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &OwnableTransactor{contract: contract}, nil
}

// NewOwnableFilterer creates a new log filterer instance of Ownable, bound to a specific deployed contract.
func NewOwnableFilterer(address common.Address, filterer bind.ContractFilterer) (*OwnableFilterer, error) {
	contract, err := bindOwnable(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &OwnableFilterer{contract: contract}, nil
}

// bindOwnable binds a generic wrapper to an already deployed contract.
func bindOwnable(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(OwnableABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
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

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_Ownable *OwnableTransactor) RenounceOwnership(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Ownable.contract.Transact(opts, "renounceOwnership")
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_Ownable *OwnableSession) RenounceOwnership() (*types.Transaction, error) {
	return _Ownable.Contract.RenounceOwnership(&_Ownable.TransactOpts)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_Ownable *OwnableTransactorSession) RenounceOwnership() (*types.Transaction, error) {
	return _Ownable.Contract.RenounceOwnership(&_Ownable.TransactOpts)
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

// OwnableOwnershipRenouncedIterator is returned from FilterOwnershipRenounced and is used to iterate over the raw logs and unpacked data for OwnershipRenounced events raised by the Ownable contract.
type OwnableOwnershipRenouncedIterator struct {
	Event *OwnableOwnershipRenounced // Event containing the contract specifics and raw log

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
func (it *OwnableOwnershipRenouncedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(OwnableOwnershipRenounced)
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
		it.Event = new(OwnableOwnershipRenounced)
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
func (it *OwnableOwnershipRenouncedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *OwnableOwnershipRenouncedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// OwnableOwnershipRenounced represents a OwnershipRenounced event raised by the Ownable contract.
type OwnableOwnershipRenounced struct {
	PreviousOwner common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterOwnershipRenounced is a free log retrieval operation binding the contract event 0xf8df31144d9c2f0f6b59d69b8b98abd5459d07f2742c4df920b25aae33c64820.
//
// Solidity: event OwnershipRenounced(previousOwner indexed address)
func (_Ownable *OwnableFilterer) FilterOwnershipRenounced(opts *bind.FilterOpts, previousOwner []common.Address) (*OwnableOwnershipRenouncedIterator, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}

	logs, sub, err := _Ownable.contract.FilterLogs(opts, "OwnershipRenounced", previousOwnerRule)
	if err != nil {
		return nil, err
	}
	return &OwnableOwnershipRenouncedIterator{contract: _Ownable.contract, event: "OwnershipRenounced", logs: logs, sub: sub}, nil
}

// WatchOwnershipRenounced is a free log subscription operation binding the contract event 0xf8df31144d9c2f0f6b59d69b8b98abd5459d07f2742c4df920b25aae33c64820.
//
// Solidity: event OwnershipRenounced(previousOwner indexed address)
func (_Ownable *OwnableFilterer) WatchOwnershipRenounced(opts *bind.WatchOpts, sink chan<- *OwnableOwnershipRenounced, previousOwner []common.Address) (event.Subscription, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}

	logs, sub, err := _Ownable.contract.WatchLogs(opts, "OwnershipRenounced", previousOwnerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(OwnableOwnershipRenounced)
				if err := _Ownable.contract.UnpackLog(event, "OwnershipRenounced", log); err != nil {
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

// OwnableOwnershipTransferredIterator is returned from FilterOwnershipTransferred and is used to iterate over the raw logs and unpacked data for OwnershipTransferred events raised by the Ownable contract.
type OwnableOwnershipTransferredIterator struct {
	Event *OwnableOwnershipTransferred // Event containing the contract specifics and raw log

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
func (it *OwnableOwnershipTransferredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(OwnableOwnershipTransferred)
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
		it.Event = new(OwnableOwnershipTransferred)
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
func (it *OwnableOwnershipTransferredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *OwnableOwnershipTransferredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// OwnableOwnershipTransferred represents a OwnershipTransferred event raised by the Ownable contract.
type OwnableOwnershipTransferred struct {
	PreviousOwner common.Address
	NewOwner      common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterOwnershipTransferred is a free log retrieval operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(previousOwner indexed address, newOwner indexed address)
func (_Ownable *OwnableFilterer) FilterOwnershipTransferred(opts *bind.FilterOpts, previousOwner []common.Address, newOwner []common.Address) (*OwnableOwnershipTransferredIterator, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _Ownable.contract.FilterLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return &OwnableOwnershipTransferredIterator{contract: _Ownable.contract, event: "OwnershipTransferred", logs: logs, sub: sub}, nil
}

// WatchOwnershipTransferred is a free log subscription operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(previousOwner indexed address, newOwner indexed address)
func (_Ownable *OwnableFilterer) WatchOwnershipTransferred(opts *bind.WatchOpts, sink chan<- *OwnableOwnershipTransferred, previousOwner []common.Address, newOwner []common.Address) (event.Subscription, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _Ownable.contract.WatchLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(OwnableOwnershipTransferred)
				if err := _Ownable.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
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

// SafeMathABI is the input ABI used to generate the binding from.
const SafeMathABI = "[]"

// SafeMathBin is the compiled bytecode used for deploying new contracts.
const SafeMathBin = `0x604c602c600b82828239805160001a60731460008114601c57601e565bfe5b5030600052607381538281f30073000000000000000000000000000000000000000030146080604052600080fd00a165627a7a7230582067a4a406ab39bef9b55ed750f9a337c5dec572a7bb91b3fdf6a96840e67956580029`

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
	return address, tx, &SafeMath{SafeMathCaller: SafeMathCaller{contract: contract}, SafeMathTransactor: SafeMathTransactor{contract: contract}, SafeMathFilterer: SafeMathFilterer{contract: contract}}, nil
}

// SafeMath is an auto generated Go binding around an Ethereum contract.
type SafeMath struct {
	SafeMathCaller     // Read-only binding to the contract
	SafeMathTransactor // Write-only binding to the contract
	SafeMathFilterer   // Log filterer for contract events
}

// SafeMathCaller is an auto generated read-only Go binding around an Ethereum contract.
type SafeMathCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// SafeMathTransactor is an auto generated write-only Go binding around an Ethereum contract.
type SafeMathTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// SafeMathFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type SafeMathFilterer struct {
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
	contract, err := bindSafeMath(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &SafeMath{SafeMathCaller: SafeMathCaller{contract: contract}, SafeMathTransactor: SafeMathTransactor{contract: contract}, SafeMathFilterer: SafeMathFilterer{contract: contract}}, nil
}

// NewSafeMathCaller creates a new read-only instance of SafeMath, bound to a specific deployed contract.
func NewSafeMathCaller(address common.Address, caller bind.ContractCaller) (*SafeMathCaller, error) {
	contract, err := bindSafeMath(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &SafeMathCaller{contract: contract}, nil
}

// NewSafeMathTransactor creates a new write-only instance of SafeMath, bound to a specific deployed contract.
func NewSafeMathTransactor(address common.Address, transactor bind.ContractTransactor) (*SafeMathTransactor, error) {
	contract, err := bindSafeMath(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &SafeMathTransactor{contract: contract}, nil
}

// NewSafeMathFilterer creates a new log filterer instance of SafeMath, bound to a specific deployed contract.
func NewSafeMathFilterer(address common.Address, filterer bind.ContractFilterer) (*SafeMathFilterer, error) {
	contract, err := bindSafeMath(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &SafeMathFilterer{contract: contract}, nil
}

// bindSafeMath binds a generic wrapper to an already deployed contract.
func bindSafeMath(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(SafeMathABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
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
