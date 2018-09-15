// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package platform

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

// HoQuStorageABI is the input ABI used to generate the binding from.
const HoQuStorageABI = "[{\"constant\":false,\"inputs\":[{\"name\":\"id\",\"type\":\"bytes16\"},{\"name\":\"ownerId\",\"type\":\"bytes16\"},{\"name\":\"name\",\"type\":\"string\"},{\"name\":\"dataUrl\",\"type\":\"string\"},{\"name\":\"status\",\"type\":\"uint8\"}],\"name\":\"setNetwork\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"id\",\"type\":\"bytes16\"},{\"name\":\"ownerAddress\",\"type\":\"address\"}],\"name\":\"addUserAddress\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"id\",\"type\":\"bytes16\"},{\"name\":\"ownerId\",\"type\":\"bytes16\"},{\"name\":\"networkId\",\"type\":\"bytes16\"},{\"name\":\"merchantId\",\"type\":\"bytes16\"},{\"name\":\"payerAddress\",\"type\":\"address\"},{\"name\":\"name\",\"type\":\"string\"},{\"name\":\"dataUrl\",\"type\":\"string\"},{\"name\":\"status\",\"type\":\"uint8\"}],\"name\":\"setOffer\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"id\",\"type\":\"bytes16\"},{\"name\":\"meta\",\"type\":\"string\"},{\"name\":\"kycLevel\",\"type\":\"uint8\"},{\"name\":\"dataUrl\",\"type\":\"string\"}],\"name\":\"addKycReport\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"id\",\"type\":\"bytes16\"},{\"name\":\"role\",\"type\":\"string\"},{\"name\":\"ownerAddress\",\"type\":\"address\"},{\"name\":\"kycLevel\",\"type\":\"uint8\"},{\"name\":\"pubKey\",\"type\":\"string\"},{\"name\":\"status\",\"type\":\"uint8\"}],\"name\":\"setUser\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"\",\"type\":\"bytes16\"}],\"name\":\"ids\",\"outputs\":[{\"name\":\"userId\",\"type\":\"bytes16\"},{\"name\":\"companyId\",\"type\":\"bytes16\"},{\"name\":\"idType\",\"type\":\"string\"},{\"name\":\"name\",\"type\":\"string\"},{\"name\":\"numOfKycReports\",\"type\":\"uint16\"},{\"name\":\"createdAt\",\"type\":\"uint256\"},{\"name\":\"status\",\"type\":\"uint8\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"id\",\"type\":\"bytes16\"},{\"name\":\"num\",\"type\":\"uint8\"},{\"name\":\"tariffGroupId\",\"type\":\"bytes16\"}],\"name\":\"setOfferTariffGroup\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"id\",\"type\":\"bytes16\"},{\"name\":\"num\",\"type\":\"uint8\"}],\"name\":\"getTariffGroupTariff\",\"outputs\":[{\"name\":\"\",\"type\":\"bytes16\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"\",\"type\":\"bytes16\"}],\"name\":\"companies\",\"outputs\":[{\"name\":\"ownerId\",\"type\":\"bytes16\"},{\"name\":\"name\",\"type\":\"string\"},{\"name\":\"dataUrl\",\"type\":\"string\"},{\"name\":\"createdAt\",\"type\":\"uint256\"},{\"name\":\"status\",\"type\":\"uint8\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"id\",\"type\":\"bytes16\"},{\"name\":\"tariffGroupId\",\"type\":\"bytes16\"},{\"name\":\"name\",\"type\":\"string\"},{\"name\":\"action\",\"type\":\"string\"},{\"name\":\"calcMethod\",\"type\":\"string\"},{\"name\":\"price\",\"type\":\"uint256\"},{\"name\":\"status\",\"type\":\"uint8\"}],\"name\":\"setTariff\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"\",\"type\":\"bytes16\"}],\"name\":\"networks\",\"outputs\":[{\"name\":\"ownerId\",\"type\":\"bytes16\"},{\"name\":\"name\",\"type\":\"string\"},{\"name\":\"dataUrl\",\"type\":\"string\"},{\"name\":\"createdAt\",\"type\":\"uint256\"},{\"name\":\"status\",\"type\":\"uint8\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"id\",\"type\":\"bytes16\"},{\"name\":\"num\",\"type\":\"uint8\"}],\"name\":\"getOfferTariffGroup\",\"outputs\":[{\"name\":\"\",\"type\":\"bytes16\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"\",\"type\":\"bytes16\"}],\"name\":\"users\",\"outputs\":[{\"name\":\"ownerAddress\",\"type\":\"address\"},{\"name\":\"numOfAddresses\",\"type\":\"uint8\"},{\"name\":\"role\",\"type\":\"string\"},{\"name\":\"kycLevel\",\"type\":\"uint8\"},{\"name\":\"pubKey\",\"type\":\"string\"},{\"name\":\"createdAt\",\"type\":\"uint256\"},{\"name\":\"status\",\"type\":\"uint8\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"id\",\"type\":\"bytes16\"},{\"name\":\"ownerId\",\"type\":\"bytes16\"},{\"name\":\"name\",\"type\":\"string\"},{\"name\":\"dataUrl\",\"type\":\"string\"},{\"name\":\"status\",\"type\":\"uint8\"}],\"name\":\"setCompany\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"config\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"configAddress\",\"type\":\"address\"}],\"name\":\"setConfigAddress\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"\",\"type\":\"bytes16\"}],\"name\":\"adCampaigns\",\"outputs\":[{\"name\":\"ownerId\",\"type\":\"bytes16\"},{\"name\":\"offerId\",\"type\":\"bytes16\"},{\"name\":\"contractAddress\",\"type\":\"address\"},{\"name\":\"createdAt\",\"type\":\"uint256\"},{\"name\":\"status\",\"type\":\"uint8\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"\",\"type\":\"bytes16\"}],\"name\":\"tariffGroups\",\"outputs\":[{\"name\":\"ownerId\",\"type\":\"bytes16\"},{\"name\":\"name\",\"type\":\"string\"},{\"name\":\"numOfTariffs\",\"type\":\"uint8\"},{\"name\":\"createdAt\",\"type\":\"uint256\"},{\"name\":\"status\",\"type\":\"uint8\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"id\",\"type\":\"bytes16\"},{\"name\":\"ownerId\",\"type\":\"bytes16\"},{\"name\":\"offerId\",\"type\":\"bytes16\"},{\"name\":\"contractAddress\",\"type\":\"address\"},{\"name\":\"status\",\"type\":\"uint8\"}],\"name\":\"setAdCampaign\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"\",\"type\":\"bytes16\"}],\"name\":\"stats\",\"outputs\":[{\"name\":\"rating\",\"type\":\"uint256\"},{\"name\":\"volume\",\"type\":\"uint256\"},{\"name\":\"members\",\"type\":\"uint256\"},{\"name\":\"alfa\",\"type\":\"uint256\"},{\"name\":\"beta\",\"type\":\"uint256\"},{\"name\":\"status\",\"type\":\"uint8\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"\",\"type\":\"bytes16\"}],\"name\":\"tariffs\",\"outputs\":[{\"name\":\"tariffGroupId\",\"type\":\"bytes16\"},{\"name\":\"name\",\"type\":\"string\"},{\"name\":\"action\",\"type\":\"string\"},{\"name\":\"calcMethod\",\"type\":\"string\"},{\"name\":\"price\",\"type\":\"uint256\"},{\"name\":\"createdAt\",\"type\":\"uint256\"},{\"name\":\"status\",\"type\":\"uint8\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"id\",\"type\":\"bytes16\"},{\"name\":\"ownerId\",\"type\":\"bytes16\"},{\"name\":\"name\",\"type\":\"string\"},{\"name\":\"status\",\"type\":\"uint8\"}],\"name\":\"setTariffGroup\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"id\",\"type\":\"bytes16\"},{\"name\":\"userId\",\"type\":\"bytes16\"},{\"name\":\"idType\",\"type\":\"string\"},{\"name\":\"name\",\"type\":\"string\"},{\"name\":\"companyId\",\"type\":\"bytes16\"},{\"name\":\"status\",\"type\":\"uint8\"}],\"name\":\"setIdentification\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"id\",\"type\":\"bytes16\"},{\"name\":\"userId\",\"type\":\"bytes16\"},{\"name\":\"rating\",\"type\":\"uint256\"},{\"name\":\"volume\",\"type\":\"uint256\"},{\"name\":\"members\",\"type\":\"uint256\"},{\"name\":\"alfa\",\"type\":\"uint256\"},{\"name\":\"beta\",\"type\":\"uint256\"},{\"name\":\"status\",\"type\":\"uint8\"}],\"name\":\"setStats\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"\",\"type\":\"bytes16\"}],\"name\":\"offers\",\"outputs\":[{\"name\":\"ownerId\",\"type\":\"bytes16\"},{\"name\":\"networkId\",\"type\":\"bytes16\"},{\"name\":\"merchantId\",\"type\":\"bytes16\"},{\"name\":\"payerAddress\",\"type\":\"address\"},{\"name\":\"name\",\"type\":\"string\"},{\"name\":\"dataUrl\",\"type\":\"string\"},{\"name\":\"numOfTariffGroups\",\"type\":\"uint8\"},{\"name\":\"createdAt\",\"type\":\"uint256\"},{\"name\":\"status\",\"type\":\"uint8\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"id\",\"type\":\"bytes16\"},{\"name\":\"num\",\"type\":\"uint8\"}],\"name\":\"getUserAddress\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"id\",\"type\":\"bytes16\"},{\"name\":\"ownerId\",\"type\":\"bytes16\"},{\"name\":\"networkId\",\"type\":\"bytes16\"},{\"name\":\"name\",\"type\":\"string\"},{\"name\":\"dataUrl\",\"type\":\"string\"},{\"name\":\"status\",\"type\":\"uint8\"}],\"name\":\"setTracker\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"id\",\"type\":\"bytes16\"},{\"name\":\"num\",\"type\":\"uint16\"}],\"name\":\"getKycReport\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"},{\"name\":\"\",\"type\":\"string\"},{\"name\":\"\",\"type\":\"uint8\"},{\"name\":\"\",\"type\":\"string\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"\",\"type\":\"bytes16\"}],\"name\":\"trackers\",\"outputs\":[{\"name\":\"ownerId\",\"type\":\"bytes16\"},{\"name\":\"networkId\",\"type\":\"bytes16\"},{\"name\":\"name\",\"type\":\"string\"},{\"name\":\"dataUrl\",\"type\":\"string\"},{\"name\":\"createdAt\",\"type\":\"uint256\"},{\"name\":\"status\",\"type\":\"uint8\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"id\",\"type\":\"bytes16\"},{\"name\":\"tariffGroupId\",\"type\":\"bytes16\"}],\"name\":\"addOfferTariffGroup\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"id\",\"type\":\"bytes16\"},{\"name\":\"num\",\"type\":\"uint8\"},{\"name\":\"tariffId\",\"type\":\"bytes16\"}],\"name\":\"setTariffGroupTariff\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"name\":\"configAddress\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"ownerAddress\",\"type\":\"address\"},{\"indexed\":false,\"name\":\"id\",\"type\":\"bytes16\"},{\"indexed\":false,\"name\":\"role\",\"type\":\"string\"}],\"name\":\"UserRegistered\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"ownerAddress\",\"type\":\"address\"},{\"indexed\":false,\"name\":\"additionalAddress\",\"type\":\"address\"},{\"indexed\":false,\"name\":\"id\",\"type\":\"bytes16\"}],\"name\":\"UserAddressAdded\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"ownerAddress\",\"type\":\"address\"},{\"indexed\":false,\"name\":\"id\",\"type\":\"bytes16\"},{\"indexed\":false,\"name\":\"rating\",\"type\":\"uint256\"}],\"name\":\"StatsChanged\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"ownerAddress\",\"type\":\"address\"},{\"indexed\":false,\"name\":\"id\",\"type\":\"bytes16\"},{\"indexed\":false,\"name\":\"userId\",\"type\":\"bytes16\"},{\"indexed\":false,\"name\":\"name\",\"type\":\"string\"}],\"name\":\"IdentificationAdded\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"ownerAddress\",\"type\":\"address\"},{\"indexed\":false,\"name\":\"kycLevel\",\"type\":\"uint8\"}],\"name\":\"KycReportAdded\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"ownerAddress\",\"type\":\"address\"},{\"indexed\":false,\"name\":\"id\",\"type\":\"bytes16\"},{\"indexed\":false,\"name\":\"name\",\"type\":\"string\"}],\"name\":\"CompanyRegistered\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"ownerAddress\",\"type\":\"address\"},{\"indexed\":false,\"name\":\"id\",\"type\":\"bytes16\"},{\"indexed\":false,\"name\":\"name\",\"type\":\"string\"}],\"name\":\"NetworkRegistered\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"ownerAddress\",\"type\":\"address\"},{\"indexed\":false,\"name\":\"id\",\"type\":\"bytes16\"},{\"indexed\":false,\"name\":\"name\",\"type\":\"string\"}],\"name\":\"TrackerRegistered\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"ownerAddress\",\"type\":\"address\"},{\"indexed\":false,\"name\":\"id\",\"type\":\"bytes16\"},{\"indexed\":false,\"name\":\"name\",\"type\":\"string\"}],\"name\":\"OfferAdded\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"ownerAddress\",\"type\":\"address\"},{\"indexed\":false,\"name\":\"id\",\"type\":\"bytes16\"},{\"indexed\":false,\"name\":\"tariff_id\",\"type\":\"bytes16\"}],\"name\":\"OfferTariffGroupAdded\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"ownerAddress\",\"type\":\"address\"},{\"indexed\":false,\"name\":\"id\",\"type\":\"bytes16\"},{\"indexed\":false,\"name\":\"contractAddress\",\"type\":\"address\"}],\"name\":\"AdCampaignAdded\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"ownerAddress\",\"type\":\"address\"},{\"indexed\":false,\"name\":\"id\",\"type\":\"bytes16\"},{\"indexed\":false,\"name\":\"name\",\"type\":\"string\"}],\"name\":\"TariffGroupAdded\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"ownerAddress\",\"type\":\"address\"},{\"indexed\":false,\"name\":\"id\",\"type\":\"bytes16\"},{\"indexed\":false,\"name\":\"name\",\"type\":\"string\"}],\"name\":\"TariffAdded\",\"type\":\"event\"}]"

// HoQuStorageBin is the compiled bytecode used for deploying new contracts.
const HoQuStorageBin = `0x608060405234801561001057600080fd5b5060405160208061516d8339810160405251600a8054600160a060020a031916600160a060020a0390921691909117905561511d806100506000396000f30060806040526004361061018a5763ffffffff7c010000000000000000000000000000000000000000000000000000000060003504166302884980811461018f57806303806028146102445780631071d5121461027257806318c8f3051461033e578063253a9b9e146103ed578063289e7470146104b15780633e45850a146105f3578063436e7b65146106225780634df2d8471461066757806361fad57e1461079657806372c130e41461088d57806374f627c5146108af5780637520dd14146108d7578063768d80a8146109df57806379502c5514610a9257806383a12de914610ac35780638907555714610ae45780638ade002714610b5f578063a0191d2014610c34578063a2affa4d14610c74578063a2cbda4114610cdd578063a7fbcf6114610e81578063acc02cce14610ef6578063b1dbfe5014610fb6578063c32d869b14610ff3578063c49e2fb114611151578063c4d4d88714611179578063cc3436b91461122f578063ea2fdd8c14611312578063f17eceb114611449578063fafe27d714611471575b600080fd5b34801561019b57600080fd5b50604080516020600460443581810135601f81018490048402850184019095528484526102429482356001608060020a03199081169560248035909216953695946064949293019190819084018382808284375050604080516020601f89358b018035918201839004830284018301909452808352979a9998810197919650918201945092508291508401838280828437509497505050923560ff1693506114a092505050565b005b34801561025057600080fd5b506102426001608060020a031960043516600160a060020a036024351661182d565b34801561027e57600080fd5b50604080516020600460a43581810135601f81018490048402850184019095528484526102429482356001608060020a0319908116956024803583169660443584169660643590941695608435600160a060020a03169536959460c4949093920191819084018382808284375050604080516020601f89358b018035918201839004830284018301909452808352979a9998810197919650918201945092508291508401838280828437509497505050923560ff16935061199692505050565b34801561034a57600080fd5b5060408051602060046024803582810135601f81018590048502860185019096528585526102429583356001608060020a03191695369560449491939091019190819084018382808284375050604080516020601f818a01358b0180359182018390048302840183018552818452989b60ff8b35169b909a909994019750919550918201935091508190840183828082843750949750611ea99650505050505050565b3480156103f957600080fd5b5060408051602060046024803582810135601f81018590048502860185019096528585526102429583356001608060020a03191695369560449491939091019190819084018382808284375050604080516020888301358a018035601f8101839004830284018301909452838352979a8935600160a060020a03169a8a83013560ff169a919990985060609091019650919450908101925081908401838280828437509497505050923560ff1693506120d192505050565b3480156104bd57600080fd5b506104d36001608060020a031960043516612494565b604080516001608060020a0319808a1682528816602082015261ffff8516608082015260a081018490529081016060820160c0830184600581111561051457fe5b60ff168152602001838103835288818151815260200191508051906020019080838360005b83811015610551578181015183820152602001610539565b50505050905090810190601f16801561057e5780820380516001836020036101000a031916815260200191505b50838103825287518152875160209182019189019080838360005b838110156105b1578181015183820152602001610599565b50505050905090810190601f1680156105de5780820380516001836020036101000a031916815260200191505b50995050505050505050505060405180910390f35b3480156105ff57600080fd5b506102426001608060020a031960043581169060ff6024351690604435166125f7565b34801561062e57600080fd5b5061064a6001608060020a03196004351660ff60243516612798565b604080516001608060020a03199092168252519081900360200190f35b34801561067357600080fd5b506106896001608060020a03196004351661280a565b604080516001608060020a031987168152606081018490529060208201908201608083018460058111156106b957fe5b60ff168152602001838103835287818151815260200191508051906020019080838360005b838110156106f65781810151838201526020016106de565b50505050905090810190601f1680156107235780820380516001836020036101000a031916815260200191505b50838103825286518152865160209182019188019080838360005b8381101561075657818101518382015260200161073e565b50505050905090810190601f1680156107835780820380516001836020036101000a031916815260200191505b5097505050505050505060405180910390f35b3480156107a257600080fd5b50604080516020600460443581810135601f81018490048402850184019095528484526102429482356001608060020a03199081169560248035909216953695946064949293019190819084018382808284375050604080516020601f89358b018035918201839004830284018301909452808352979a99988101979196509182019450925082915084018382808284375050604080516020601f89358b018035918201839004830284018301909452808352979a999881019791965091820194509250829150840183828082843750949750508435955050505060209091013560ff16905061295d565b34801561089957600080fd5b506106896001608060020a031960043516612e5d565b3480156108bb57600080fd5b5061064a6001608060020a03196004351660ff60243516612ed9565b3480156108e357600080fd5b506108f96001608060020a031960043516612f4b565b60408051600160a060020a038916815260ff881660208201529081016060820186600581111561092557fe5b60ff1681526020018060200185815260200184600581111561094357fe5b60ff168152602001838103835288818151815260200191508051906020019080838360005b83811015610980578181015183820152602001610968565b50505050905090810190601f1680156109ad5780820380516001836020036101000a031916815260200191505b5083810382528651815286516020918201918801908083836000838110156105b1578181015183820152602001610599565b3480156109eb57600080fd5b50604080516020600460443581810135601f81018490048402850184019095528484526102429482356001608060020a03199081169560248035909216953695946064949293019190819084018382808284375050604080516020601f89358b018035918201839004830284018301909452808352979a9998810197919650918201945092508291508401838280828437509497505050923560ff1693506130ae92505050565b348015610a9e57600080fd5b50610aa76133ec565b60408051600160a060020a039092168252519081900360200190f35b348015610acf57600080fd5b50610242600160a060020a03600435166133fb565b348015610af057600080fd5b50610b066001608060020a0319600435166134a4565b604080516001608060020a0319808816825286166020820152600160a060020a038516918101919091526060810183905260808101826005811115610b4757fe5b60ff1681526020019550505050505060405180910390f35b348015610b6b57600080fd5b50610b816001608060020a0319600435166134e6565b604080516001608060020a03198716815260ff851691810191909152606081018390526020810160808201836005811115610bb857fe5b60ff168152602001828103825286818151815260200191508051906020019080838360005b83811015610bf5578181015183820152602001610bdd565b50505050905090810190601f168015610c225780820380516001836020036101000a031916815260200191505b50965050505050505060405180910390f35b348015610c4057600080fd5b506102426001608060020a031960043581169060243581169060443516600160a060020a036064351660ff608435166135b0565b348015610c8057600080fd5b50610c966001608060020a0319600435166138df565b60405180878152602001868152602001858152602001848152602001838152602001826005811115610cc457fe5b60ff168152602001965050505050505060405180910390f35b348015610ce957600080fd5b50610cff6001608060020a031960043516613919565b60405180886001608060020a0319166001608060020a0319168152602001806020018060200180602001878152602001868152602001856005811115610d4157fe5b60ff16815260200184810384528a818151815260200191508051906020019080838360005b83811015610d7e578181015183820152602001610d66565b50505050905090810190601f168015610dab5780820380516001836020036101000a031916815260200191505b5084810383528951815289516020918201918b019080838360005b83811015610dde578181015183820152602001610dc6565b50505050905090810190601f168015610e0b5780820380516001836020036101000a031916815260200191505b5084810382528851815288516020918201918a019080838360005b83811015610e3e578181015183820152602001610e26565b50505050905090810190601f168015610e6b5780820380516001836020036101000a031916815260200191505b509a505050505050505050505060405180910390f35b348015610e8d57600080fd5b50604080516020600460443581810135601f81018490048402850184019095528484526102429482356001608060020a0319908116956024803590921695369594606494929301919081908401838280828437509497505050923560ff169350613b0392505050565b348015610f0257600080fd5b50604080516020600460443581810135601f81018490048402850184019095528484526102429482356001608060020a03199081169560248035909216953695946064949293019190819084018382808284375050604080516020601f89358b018035918201839004830284018301909452808352979a999881019791965091820194509250829150840183828082843750949750505083356001608060020a03191694505050506020013560ff16613dbf565b348015610fc257600080fd5b506102426001608060020a03196004358116906024351660443560643560843560a43560c43560ff60e4351661413a565b348015610fff57600080fd5b506110156001608060020a0319600435166144ed565b604080516001608060020a0319808c1682528a81166020830152891691810191909152600160a060020a038716606082015260ff841660c082015260e081018390526080810160a08201610100830184600581111561107057fe5b60ff168152602001838103835288818151815260200191508051906020019080838360005b838110156110ad578181015183820152602001611095565b50505050905090810190601f1680156110da5780820380516001836020036101000a031916815260200191505b50838103825287518152875160209182019189019080838360005b8381101561110d5781810151838201526020016110f5565b50505050905090810190601f16801561113a5780820380516001836020036101000a031916815260200191505b509b50505050505050505050505060405180910390f35b34801561115d57600080fd5b50610aa76001608060020a03196004351660ff60243516614662565b34801561118557600080fd5b50604080516020601f606435600481810135928301849004840285018401909552818452610242946001608060020a0319813581169560248035831696604435909316953695608494920191819084018382808284375050604080516020601f89358b018035918201839004830284018301909452808352979a9998810197919650918201945092508291508401838280828437509497505050923560ff1693506146d492505050565b34801561123b57600080fd5b506112586001608060020a03196004351661ffff60243516614a97565b604051808581526020018060200184600581111561127257fe5b60ff16815260200180602001838103835286818151815260200191508051906020019080838360005b838110156112b357818101518382015260200161129b565b50505050905090810190601f1680156112e05780820380516001836020036101000a031916815260200191505b508381038252845181528451602091820191860190808383600083811015610bf5578181015183820152602001610bdd565b34801561131e57600080fd5b506113346001608060020a031960043516614c3e565b604080516001608060020a0319808916825287166020820152608081018490529081016060820160a0830184600581111561136b57fe5b60ff168152602001838103835287818151815260200191508051906020019080838360005b838110156113a8578181015183820152602001611390565b50505050905090810190601f1680156113d55780820380516001836020036101000a031916815260200191505b50838103825286518152865160209182019188019080838360005b838110156114085781810151838201526020016113f0565b50505050905090810190601f1680156114355780820380516001836020036101000a031916815260200191505b509850505050505050505060405180910390f35b34801561145557600080fd5b506102426001608060020a031960043581169060243516614d96565b34801561147d57600080fd5b506102426001608060020a031960043581169060ff602435169060443516614f18565b600a546040805160e060020a63babcc5390281523360048201529051600160a060020a039092169163babcc539916024808201926020929091908290030181600087803b1580156114f057600080fd5b505af1158015611504573d6000803e3d6000fd5b505050506040513d602081101561151a57600080fd5b5051151561152757600080fd5b6001608060020a0319851660009081526004602081905260408220015460ff16600581111561155257fe5b1415611767576001608060020a0319841660009081526020819052604081206007015460ff16600581111561158357fe5b141561158e57600080fd5b6001608060020a03198416600090815260208181526040808320838052600101909152902054600160a060020a031615156115c857600080fd5b6040805160a0810182526001608060020a0319861681526020810185905290810183905242606082015260808101600190526001608060020a0319808716600090815260046020908152604090912083518154608060020a90910493169290921782558281015180516116419260018501920190615056565b506040820151805161165d916002840191602090910190615056565b5060608201516003820155608082015160048201805460ff1916600183600581111561168557fe5b021790555050506001608060020a03198085166000908152602081815260408083208380526001018252808320548151948a1685528483018281528851928601929092528751600160a060020a03909116947f592e83ede8c846dffd6e4edcbe77983b153bec820ac697aa34a236147b1341e0948b948a9492936060850192918601918190849084905b8381101561172757818101518382015260200161170f565b50505050905090810190601f1680156117545780820380516001836020036101000a031916815260200191505b50935050505060405180910390a2611826565b8251156117a0576001608060020a031985166000908152600460209081526040909120845161179e92600190920191860190615056565b505b8151156117d9576001608060020a03198516600090815260046020908152604090912083516117d792600290920191850190615056565b505b60008160058111156117e757fe5b14611826576001608060020a03198516600090815260046020819052604090912001805482919060ff1916600183600581111561182057fe5b02179055505b5050505050565b600a546040805160e060020a63babcc5390281523360048201529051600160a060020a039092169163babcc539916024808201926020929091908290030181600087803b15801561187d57600080fd5b505af1158015611891573d6000803e3d6000fd5b505050506040513d60208110156118a757600080fd5b505115156118b457600080fd5b6001608060020a0319821660009081526020819052604081206007015460ff1660058111156118df57fe5b14156118ea57600080fd5b6001608060020a0319821660008181526020818152604080832060028101805460ff9081168652600192830185528386208054600160a060020a031916600160a060020a038a8116918217909255835460ff198116908416909501909216939093179091558480529382902054825194855292840194909452805191909316927fb38b406c043d4a3feaa0192d2e1b02f317fa73c284eb3e70b9a1fde8612af1b3928290030190a25050565b600a546040805160e060020a63babcc5390281523360048201529051600160a060020a039092169163babcc539916024808201926020929091908290030181600087803b1580156119e657600080fd5b505af11580156119fa573d6000803e3d6000fd5b505050506040513d6020811015611a1057600080fd5b50511515611a1d57600080fd5b6001608060020a0319861660009081526004602081905260408220015460ff166005811115611a4857fe5b1415611a5357600080fd5b6001608060020a0319851660009081526020819052604081206007015460ff166005811115611a7e57fe5b1415611a8957600080fd5b6001608060020a0319881660009081526006602052604081206004015460ff166005811115611ab457fe5b1415611d52576001608060020a0319871660009081526020819052604081206007015460ff166005811115611ae557fe5b1415611af057600080fd5b6001608060020a03198716600090815260208181526040808320838052600101909152902054600160a060020a03161515611b2a57600080fd5b60408051610120810182526001608060020a0319808a1682528881166020830152871691810191909152600160a060020a03851660608201526080810184905260a08101839052600060c08201524260e08201526101008101600190526001608060020a031989811660009081526005602090815260409182902084518154868401516001608060020a03918716608060020a93849004179190911690829004820217825592850151600182018054909516939004929092179092556060830151600282018054600160a060020a031916600160a060020a03909216919091179055608083015180519192611c2792600385019290910190615056565b5060a08201518051611c43916004840191602090910190615056565b5060c082015160068201805460ff90921660ff1992831617905560e0830151600783015561010083015160088301805491929091166001836005811115611c8657fe5b021790555090505083600160a060020a03167f5c9100ad970db426d43d47b3a9639dc42d9732871d99213dcb4984f7e71509d0898560405180836001608060020a0319166001608060020a031916815260200180602001828103825283818151815260200191508051906020019080838360005b83811015611d12578181015183820152602001611cfa565b50505050905090810190601f168015611d3f5780820380516001836020036101000a031916815260200191505b50935050505060405180910390a2611e9f565b6001608060020a0319888116600090815260056020526040902080546001608060020a0316608060020a808a04810291909117825560019091018054909216908704179055600160a060020a03841615611ddc576001608060020a0319881660009081526005602052604090206002018054600160a060020a031916600160a060020a0386161790555b825115611e15576001608060020a0319881660009081526005602090815260409091208451611e1392600390920191860190615056565b505b815115611e4e576001608060020a0319881660009081526005602090815260409091208351611e4c92600490920191850190615056565b505b6000816005811115611e5c57fe5b14611e9f576001608060020a0319881660009081526005602081905260409091206008018054839260ff19909116906001908490811115611e9957fe5b02179055505b5050505050505050565b600a546040805160e060020a63babcc5390281523360048201529051600160a060020a039092169163babcc539916024808201926020929091908290030181600087803b158015611ef957600080fd5b505af1158015611f0d573d6000803e3d6000fd5b505050506040513d6020811015611f2357600080fd5b50511515611f3057600080fd5b6001608060020a0319841660009081526001602052604081206006015460ff166005811115611f5b57fe5b1415611f6657600080fd5b608060405190810160405280848152602001836005811115611f8457fe5b81526020808201849052426040928301526001608060020a03198716600090815260018252828120600481015461ffff168252600301825291909120825180519192611fd592849290910190615056565b50602082015160018083018054909160ff1990911690836005811115611ff757fe5b021790555060408201518051612017916002840191602090910190615056565b50606091909101516003909101556001608060020a0319808516600090815260016020818152604080842060048101805461ffff80821687011661ffff1990911617905554608060020a0290941683528281528383208380529091019052819020549051600160a060020a03909116907f92262b4b81e23d81eaec5e9e8c9439f0a59929da8c5d49b32bca92c112f4172a908490808260058111156120b857fe5b60ff16815260200191505060405180910390a250505050565b600a546040805160e060020a63babcc5390281523360048201529051600160a060020a039092169163babcc539916024808201926020929091908290030181600087803b15801561212157600080fd5b505af1158015612135573d6000803e3d6000fd5b505050506040513d602081101561214b57600080fd5b5051151561215857600080fd5b6001608060020a0319861660009081526020819052604081206007015460ff16600581111561218357fe5b1415612386576040805160e081018252600160a060020a0386168152600160208201819052918101879052906060820190815260208101849052426040820152606001600190526001608060020a031987166000908152602081815260409182902083518154600160a060020a031916600160a060020a039091161781558382015160028201805460ff191660ff9092169190911790559183015180516122309260038501920190615056565b50606082015160048201805460ff1916600183600581111561224e57fe5b02179055506080820151805161226e916005840191602090910190615056565b5060a0820151600682015560c082015160078201805460ff1916600183600581111561229657fe5b021790555050506001608060020a0319861660008181526020818152604080832083805260010182528083208054600160a060020a031916600160a060020a038a1690811790915581519485528483018281528a5192860192909252895190947ff661ee1472892faaf2bb68cc6874f2759c9148b4234cac6b88f6ca362652f275948c948c9492936060850192918601918190849084905b8381101561234657818101518382015260200161232e565b50505050905090810190601f1680156123735780820380516001836020036101000a031916815260200191505b50935050505060405180910390a261248c565b8451156123bd576001608060020a0319861660009081526020818152604090912086516123bb92600390920191880190615056565b505b60008360058111156123cb57fe5b14612409576001608060020a031986166000908152602081905260409020600401805484919060ff1916600183600581111561240357fe5b02179055505b815115612440576001608060020a03198616600090815260208181526040909120835161243e92600590920191850190615056565b505b600081600581111561244e57fe5b1461248c576001608060020a031986166000908152602081905260409020600701805482919060ff1916600183600581111561248657fe5b02179055505b505050505050565b600160208181526000928352604092839020805481840180548651600261010097831615979097026000190190911695909504601f8101859004850286018501909652858552608060020a80830296928190040294929390919083018282801561253f5780601f106125145761010080835404028352916020019161253f565b820191906000526020600020905b81548152906001019060200180831161252257829003601f168201915b50505060028085018054604080516020601f60001961010060018716150201909416959095049283018590048502810185019091528181529596959450909250908301828280156125d15780601f106125a6576101008083540402835291602001916125d1565b820191906000526020600020905b8154815290600101906020018083116125b457829003601f168201915b5050505060048301546005840154600690940154929361ffff9091169290915060ff1687565b600a546040805160e060020a63babcc5390281523360048201529051600092600160a060020a03169163babcc53991602480830192602092919082900301818787803b15801561264657600080fd5b505af115801561265a573d6000803e3d6000fd5b505050506040513d602081101561267057600080fd5b5051151561267d57600080fd5b6001608060020a03198416600090815260056020819052604082206008015460ff16908111156126a957fe5b14156126b457600080fd5b6001608060020a0319808516600090815260056020818152604080842060ff891685529092019052902054608060020a021615156126f157600080fd5b6001608060020a0319808516600081815260056020818152604080842060ff8a16855280840183529084208054608060020a808b049190981617905593835252905461273e920290614662565b604080516001608060020a03198088168252851660208201528151929350600160a060020a038416927f695164079ea85cb7933398c55f763b3dd70fe1cfe293a3b90aa22f9cf6270dab929181900390910190a250505050565b6000806001608060020a03198416600090815260096020526040902060059081015460ff16908111156127c757fe5b14156127d257600080fd5b506001608060020a03198216600090815260096020908152604080832060ff85168452600201909152902054608060020a0292915050565b6003602090815260009182526040918290208054600180830180548651600261010094831615949094026000190190911692909204601f8101869004860283018601909652858252608060020a9092029492939092908301828280156128b15780601f10612886576101008083540402835291602001916128b1565b820191906000526020600020905b81548152906001019060200180831161289457829003601f168201915b50505060028085018054604080516020601f60001961010060018716150201909416959095049283018590048502810185019091528181529596959450909250908301828280156129435780601f1061291857610100808354040283529160200191612943565b820191906000526020600020905b81548152906001019060200180831161292657829003601f168201915b50505050600383015460049093015491929160ff16905085565b600a546040805160e060020a63babcc5390281523360048201529051600092600160a060020a03169163babcc53991602480830192602092919082900301818787803b1580156129ac57600080fd5b505af11580156129c0573d6000803e3d6000fd5b505050506040513d60208110156129d657600080fd5b505115156129e357600080fd5b6001608060020a03198716600090815260096020526040812060059081015460ff1690811115612a0f57fe5b1415612a1a57600080fd5b6001608060020a0319881660009081526008602052604081206006015460ff166005811115612a4557fe5b1415612d465760006001608060020a0319808916600090815260096020908152604080832054608060020a029093168252819052206007015460ff166005811115612a8c57fe5b1415612a9757600080fd5b6001608060020a03198716600090815260096020526040812054612ac391608060020a90910290614662565b905060e060405190810160405280886001608060020a031916815260200187815260200186815260200185815260200184815260200142815260200160016005811115612b0c57fe5b90526001608060020a0319808a16600090815260086020908152604090912083518154608060020a9091049316929092178255828101518051612b559260018501920190615056565b5060408201518051612b71916002840191602090910190615056565b5060608201518051612b8d916003840191602090910190615056565b506080820151816004015560a0820151816005015560c08201518160060160006101000a81548160ff02191690836005811115612bc657fe5b02179055509050508760096000896001608060020a0319166001608060020a03191681526020019081526020016000206002016000600960008b6001608060020a0319166001608060020a031916815260200190815260200160002060030160009054906101000a900460ff1660ff1660ff16815260200190815260200160002060006101000a8154816001608060020a030219169083608060020a9004021790555060096000886001608060020a0319166001608060020a0319168152602001908152602001600020600301600081819054906101000a900460ff168092919060010191906101000a81548160ff021916908360ff1602179055505080600160a060020a03167f4ad20f56c3c96698f2e84d660e8c5e39341fd79929743804d261beb65ffc5b47898860405180836001608060020a0319166001608060020a0319168152602001806020018281038252838181518152602001915080519060200190808383600083811015611d12578181015183820152602001611cfa565b855115612d7f576001608060020a0319881660009081526008602090815260409091208751612d7d92600190920191890190615056565b505b845115612db8576001608060020a0319881660009081526008602090815260409091208651612db692600290920191880190615056565b505b835115612df1576001608060020a0319881660009081526008602090815260409091208551612def92600390920191870190615056565b505b8215612e17576001608060020a0319881660009081526008602052604090206004018390555b6000826005811115612e2557fe5b14611e9f576001608060020a031988166000908152600860205260409020600601805483919060ff19166001836005811115611e9957fe5b6004602090815260009182526040918290208054600180830180548651600261010094831615949094026000190190911692909204601f8101869004860283018601909652858252608060020a9092029492939092908301828280156128b15780601f10612886576101008083540402835291602001916128b1565b6000806001608060020a0319841660009081526005602081905260409091206008015460ff1690811115612f0957fe5b1415612f1457600080fd5b506001608060020a03198216600090815260056020818152604080842060ff861685529092019052902054608060020a0292915050565b60006020818152918152604090819020805460028083015460038401805486516101006001831615026000190190911693909304601f8101889004880284018801909652858352600160a060020a039093169560ff9091169491929190830182828015612ff95780601f10612fce57610100808354040283529160200191612ff9565b820191906000526020600020905b815481529060010190602001808311612fdc57829003601f168201915b50505050600483015460058401805460408051602060026101006001861615026000190190941693909304601f8101849004840282018401909252818152959660ff90941695939450908301828280156130945780601f1061306957610100808354040283529160200191613094565b820191906000526020600020905b81548152906001019060200180831161307757829003601f168201915b50505050600683015460079093015491929160ff16905087565b600a546040805160e060020a63babcc5390281523360048201529051600160a060020a039092169163babcc539916024808201926020929091908290030181600087803b1580156130fe57600080fd5b505af1158015613112573d6000803e3d6000fd5b505050506040513d602081101561312857600080fd5b5051151561313557600080fd5b6001608060020a0319851660009081526003602052604081206004015460ff16600581111561316057fe5b1415613334576001608060020a0319841660009081526020819052604081206007015460ff16600581111561319157fe5b141561319c57600080fd5b6001608060020a03198416600090815260208181526040808320838052600101909152902054600160a060020a031615156131d657600080fd5b6040805160a0810182526001608060020a0319861681526020810185905290810183905242606082015260808101600190526001608060020a0319808716600090815260036020908152604090912083518154608060020a909104931692909217825582810151805161324f9260018501920190615056565b506040820151805161326b916002840191602090910190615056565b5060608201516003820155608082015160048201805460ff1916600183600581111561329357fe5b021790555050506001608060020a03198085166000908152602081815260408083208380526001018252808320548151948a1685528483018281528851928601929092528751600160a060020a03909116947f0eb23d662c57754d23afdc1cffa6246fbae3bbe843eec5c42a5a5223c21414ed948b948a9492936060850192918601918190849084908381101561172757818101518382015260200161170f565b82511561336d576001608060020a031985166000908152600360209081526040909120845161336b92600190920191860190615056565b505b8151156133a6576001608060020a03198516600090815260036020908152604090912083516133a492600290920191850190615056565b505b60008160058111156133b457fe5b14611826576001608060020a031985166000908152600360205260409020600401805482919060ff1916600183600581111561182057fe5b600a54600160a060020a031681565b600a546040805160e060020a63babcc5390281523360048201529051600160a060020a039092169163babcc539916024808201926020929091908290030181600087803b15801561344b57600080fd5b505af115801561345f573d6000803e3d6000fd5b505050506040513d602081101561347557600080fd5b5051151561348257600080fd5b600a8054600160a060020a031916600160a060020a0392909216919091179055565b6007602052600090815260409020805460018201546002830154600390930154608060020a80840294938190040292600160a060020a03909216919060ff1685565b6009602090815260009182526040918290208054600180830180548651600261010094831615949094026000190190911692909204601f8101869004860283018601909652858252608060020a90920294929390929083018282801561358d5780601f106135625761010080835404028352916020019161358d565b820191906000526020600020905b81548152906001019060200180831161357057829003601f168201915b5050505060038301546004840154600590940154929360ff918216939092501685565b600a546040805160e060020a63babcc5390281523360048201529051600092600160a060020a03169163babcc53991602480830192602092919082900301818787803b1580156135ff57600080fd5b505af1158015613613573d6000803e3d6000fd5b505050506040513d602081101561362957600080fd5b5051151561363657600080fd5b6001608060020a0319861660009081526007602052604081206003015460ff16600581111561366157fe5b1415613854576001608060020a0319851660009081526020819052604081206007015460ff16600581111561369257fe5b141561369d57600080fd5b6001608060020a03198516600090815260208181526040808320838052600101909152902054600160a060020a031615156136d757600080fd5b6001608060020a03198416600090815260056020819052604082206008015460ff169081111561370357fe5b141561370e57600080fd5b613719856000614662565b6040805160a0810182526001608060020a0319808916825287166020820152600160a060020a0386169181019190915242606082015290915060808101600190526001608060020a03198088166000908152600760209081526040918290208451815492860151608060020a908190048102910492909416919091176001608060020a031692909217825582015160018083018054600160a060020a03909316600160a060020a0319909316929092179091556060830151600283015560808301516003830180549192909160ff1916908360058111156137f657fe5b021790555050604080516001608060020a031989168152600160a060020a038681166020830152825190851693507f62a609593cded926aabb6aef1e5b884804ad29d1bd1116ac4ce32085690d5ab2929181900390910190a261248c565b600160a060020a03831615613899576001608060020a0319861660009081526007602052604090206001018054600160a060020a031916600160a060020a0385161790555b60008260058111156138a757fe5b1461248c576001608060020a031986166000908152600760205260409020600301805483919060ff1916600183600581111561248657fe5b6002602081905260009182526040909120805460018201549282015460038301546004840154600590940154929493919290919060ff1686565b6008602090815260009182526040918290208054600180830180548651600261010094831615949094026000190190911692909204601f8101869004860283018601909652858252608060020a9092029492939092908301828280156139c05780601f10613995576101008083540402835291602001916139c0565b820191906000526020600020905b8154815290600101906020018083116139a357829003601f168201915b50505060028085018054604080516020601f6000196101006001871615020190941695909504928301859004850281018501909152818152959695945090925090830182828015613a525780601f10613a2757610100808354040283529160200191613a52565b820191906000526020600020905b815481529060010190602001808311613a3557829003601f168201915b5050505060038301805460408051602060026001851615610100026000190190941693909304601f8101849004840282018401909252818152949594935090830182828015613ae25780601f10613ab757610100808354040283529160200191613ae2565b820191906000526020600020905b815481529060010190602001808311613ac557829003601f168201915b50505050600483015460058401546006909401549293909290915060ff1687565b600a546040805160e060020a63babcc5390281523360048201529051600092600160a060020a03169163babcc53991602480830192602092919082900301818787803b158015613b5257600080fd5b505af1158015613b66573d6000803e3d6000fd5b505050506040513d6020811015613b7c57600080fd5b50511515613b8957600080fd5b6001608060020a03198516600090815260096020526040812060059081015460ff1690811115613bb557fe5b1415613d3c576001608060020a0319841660009081526020819052604081206007015460ff166005811115613be657fe5b1415613bf157600080fd5b613bfc846000614662565b6040805160a0810182526001608060020a03198781168252602080830188815260008486018190524260608601526001608086018190528c851682526009845295902084518154608060020a909104941693909317835551805195965092949193613c6d9390850192910190615056565b50604082015160038201805460ff90921660ff19928316179055606083015160048301556080830151600580840180549293909216906001908490811115613cb157fe5b021790555090505080600160a060020a03167f0d945b4e88c22df3f6d39d8866fac2dd4964090d9d9c04fc74f33c5695e561df868560405180836001608060020a0319166001608060020a031916815260200180602001828103825283818151815260200191508051906020019080838360008381101561172757818101518382015260200161170f565b825115613d75576001608060020a0319851660009081526009602090815260409091208451613d7392600190920191860190615056565b505b6000826005811115613d8357fe5b14611826576001608060020a03198516600090815260096020526040902060059081018054849260ff1990911690600190849081111561182057fe5b600a546040805160e060020a63babcc5390281523360048201529051600092600160a060020a03169163babcc53991602480830192602092919082900301818787803b158015613e0e57600080fd5b505af1158015613e22573d6000803e3d6000fd5b505050506040513d6020811015613e3857600080fd5b50511515613e4557600080fd5b6001608060020a0319871660009081526001602052604081206006015460ff166005811115613e7057fe5b141561407b57613e81866000614662565b6040805160e0810182526001608060020a0319808a1682528616602082015290810187905260608101869052600060808201524260a082015290915060c08101600190526001608060020a03198089166000908152600160208181526040928390208551815487840151608060020a90819004810292049616959095176001608060020a0316949094178455918401518051613f2593928501929190910190615056565b5060608201518051613f41916002840191602090910190615056565b5060808201518160040160006101000a81548161ffff021916908361ffff16021790555060a0820151816005015560c08201518160060160006101000a81548160ff02191690836005811115613f9357fe5b021790555090505080600160a060020a03167f8c03773782e9ed7fe9111421165a6cf12ce5855a4181cd1a8166d47d719ac07888888760405180846001608060020a0319166001608060020a0319168152602001836001608060020a0319166001608060020a031916815260200180602001828103825283818151815260200191508051906020019080838360005b8381101561403a578181015183820152602001614022565b50505050905090810190601f1680156140675780820380516001836020036101000a031916815260200191505b5094505050505060405180910390a2614131565b8451156140b4576001608060020a03198716600090815260016020818152604090922087516140b293919092019190880190615056565b505b6001608060020a03198716600090815260016020526040812080546001608060020a0316608060020a808704021790558260058111156140f057fe5b14614131576001608060020a0319871660009081526001602081905260409091206006018054849260ff199091169083600581111561412b57fe5b02179055505b50505050505050565b600a546040805160e060020a63babcc5390281523360048201529051600092600160a060020a03169163babcc53991602480830192602092919082900301818787803b15801561418957600080fd5b505af115801561419d573d6000803e3d6000fd5b505050506040513d60208110156141b357600080fd5b505115156141c057600080fd5b6001608060020a03198916600090815260026020526040812060059081015460ff16908111156141ec57fe5b1415614390576141fd886000614662565b905060c0604051908101604052808881526020018781526020018681526020018581526020018481526020016001600581111561423657fe5b815250600260008b6001608060020a0319166001608060020a0319168152602001908152602001600020600082015181600001556020820151816001015560408201518160020155606082015181600301556080820151816004015560a08201518160050160006101000a81548160ff021916908360058111156142b657fe5b021790555050506001608060020a0319808a166000908152600260208190526040808320938c1683529091208254815560018084015481830155838301549282019290925560038084015490820155600480840154908201556005808401548183018054939460ff90921693909260ff199091169190849081111561433757fe5b02179055505050604080516001608060020a03198b168152602081018990528151600160a060020a038416927fd2e9a0e1a68c0320c3b58608fcefa40787236bd0d2153ebbd82690b3fa064ad4928290030190a26144e2565b86156143be576001608060020a03198981166000908152600260205260408082208a9055918a168152208790555b85156143f3576001608060020a031989811660009081526002602052604080822060019081018a9055928b1682529020018690555b8415614428576001608060020a031989811660009081526002602081905260408083208201899055928b168252919020018590555b831561445d576001608060020a03198981166000908152600260205260408082206003908101889055928b1682529020018490555b8215614492576001608060020a03198981166000908152600260205260408082206004908101879055928b1682529020018390555b60008260058111156144a057fe5b146144e2576001608060020a03198916600090815260026020526040902060059081018054849260ff199091169060019084908111156144dc57fe5b02179055505b505050505050505050565b600560209081526000918252604091829020805460018083015460028085015460038601805489516101009682161596909602600019011692909204601f8101889004880285018801909852878452608060020a8086029895819004810297930295600160a060020a03909116949093928301828280156145af5780601f10614584576101008083540402835291602001916145af565b820191906000526020600020905b81548152906001019060200180831161459257829003601f168201915b5050505060048301805460408051602060026001851615610100026000190190941693909304601f810184900484028201840190925281815294959493509083018282801561463f5780601f106146145761010080835404028352916020019161463f565b820191906000526020600020905b81548152906001019060200180831161462257829003601f168201915b5050505060068301546007840154600890940154929360ff918216939092501689565b6000806001608060020a0319841660009081526020819052604090206007015460ff16600581111561469057fe5b141561469b57600080fd5b506001608060020a0319821660009081526020818152604080832060ff85168452600101909152902054600160a060020a031692915050565b600a546040805160e060020a63babcc5390281523360048201529051600160a060020a039092169163babcc539916024808201926020929091908290030181600087803b15801561472457600080fd5b505af1158015614738573d6000803e3d6000fd5b505050506040513d602081101561474e57600080fd5b5051151561475b57600080fd5b6001608060020a0319841660009081526004602081905260408220015460ff16600581111561478657fe5b141561479157600080fd5b6001608060020a0319861660009081526006602052604081206004015460ff1660058111156147bc57fe5b14156149af576001608060020a0319851660009081526020819052604081206007015460ff1660058111156147ed57fe5b14156147f857600080fd5b6001608060020a03198516600090815260208181526040808320838052600101909152902054600160a060020a0316151561483257600080fd5b6040805160c0810182526001608060020a03198088168252861660208201529081018490526060810183905242608082015260a08101600190526001608060020a03198088166000908152600660209081526040918290208451815486840151608060020a90819004810292049516949094176001608060020a03169390931783559083015180516148ca9260018501920190615056565b50606082015180516148e6916002840191602090910190615056565b506080820151600382015560a082015160048201805460ff1916600183600581111561490e57fe5b021790555050506001608060020a03198086166000908152602081815260408083208380526001018252808320548151948b1685528483018281528851928601929092528751600160a060020a03909116947fb2f67f32df451ce918abd944fd5c8a29589848693057913c5060a8e0f293b388948c948a9492936060850192918601918190849084908381101561234657818101518382015260200161232e565b6001608060020a03198616600090815260066020526040902080546001608060020a0316608060020a80870402179055825115614a18576001608060020a0319861660009081526006602090815260409091208451614a1692600190920191860190615056565b505b815115614a51576001608060020a0319861660009081526006602090815260409091208351614a4f92600290920191850190615056565b505b6000816005811115614a5f57fe5b1461248c576001608060020a031986166000908152600660205260409020600401805482919060ff1916600183600581111561248657fe5b600060608181816001608060020a0319871660009081526001602052604090206006015460ff166005811115614ac957fe5b1415614ad457600080fd5b6001608060020a03198616600090815260016020818152604080842061ffff8a16855260039081018352938190209384015484840154855483516002610100978316159790970260001901909116869004601f810186900486028201860190945283815291959460ff90911693908501928591830182828015614b985780601f10614b6d57610100808354040283529160200191614b98565b820191906000526020600020905b815481529060010190602001808311614b7b57829003601f168201915b5050845460408051602060026001851615610100026000190190941693909304601f810184900484028201840190925281815295985086945092508401905082828015614c265780601f10614bfb57610100808354040283529160200191614c26565b820191906000526020600020905b815481529060010190602001808311614c0957829003601f168201915b50505050509050935093509350935092959194509250565b6006602090815260009182526040918290208054600180830180548651600261010094831615949094026000190190911692909204601f8101869004860283018601909652858252608060020a8084029693819004029491929091830182828015614cea5780601f10614cbf57610100808354040283529160200191614cea565b820191906000526020600020905b815481529060010190602001808311614ccd57829003601f168201915b50505060028085018054604080516020601f6000196101006001871615020190941695909504928301859004850281018501909152818152959695945090925090830182828015614d7c5780601f10614d5157610100808354040283529160200191614d7c565b820191906000526020600020905b815481529060010190602001808311614d5f57829003601f168201915b50505050600383015460049093015491929160ff16905086565b600a546040805160e060020a63babcc5390281523360048201529051600092600160a060020a03169163babcc53991602480830192602092919082900301818787803b158015614de557600080fd5b505af1158015614df9573d6000803e3d6000fd5b505050506040513d6020811015614e0f57600080fd5b50511515614e1c57600080fd5b6001608060020a03198316600090815260056020819052604082206008015460ff1690811115614e4857fe5b1415614e5357600080fd5b6001608060020a0319808416600081815260056020818152604080842060068101805460ff908116875282860185529286208054608060020a808d0491909a1617905595855292909152835480821660010190911660ff19909116179092559054614ebf920290614662565b604080516001608060020a03198087168252851660208201528151929350600160a060020a038416927f695164079ea85cb7933398c55f763b3dd70fe1cfe293a3b90aa22f9cf6270dab929181900390910190a2505050565b600a546040805160e060020a63babcc5390281523360048201529051600160a060020a039092169163babcc539916024808201926020929091908290030181600087803b158015614f6857600080fd5b505af1158015614f7c573d6000803e3d6000fd5b505050506040513d6020811015614f9257600080fd5b50511515614f9f57600080fd5b6001608060020a03198316600090815260096020526040812060059081015460ff1690811115614fcb57fe5b1415614fd657600080fd5b6001608060020a0319808416600090815260096020908152604080832060ff87168452600201909152902054608060020a0216151561501457600080fd5b6001608060020a0319928316600090815260096020908152604080832060ff909516835260029094019052919091208054608060020a90920491909216179055565b828054600181600116156101000203166002900490600052602060002090601f016020900481019282601f1061509757805160ff19168380011785556150c4565b828001600101855582156150c4579182015b828111156150c45782518255916020019190600101906150a9565b506150d09291506150d4565b5090565b6150ee91905b808211156150d057600081556001016150da565b905600a165627a7a72305820f3cd9a9f2beae61a26183f8f19442ccc6162f54c22143b98fa6300a7acc6d8830029`

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

// GetOfferTariffGroup is a free data retrieval call binding the contract method 0x74f627c5.
//
// Solidity: function getOfferTariffGroup(id bytes16, num uint8) constant returns(bytes16)
func (_HoQuStorage *HoQuStorageCaller) GetOfferTariffGroup(opts *bind.CallOpts, id [16]byte, num uint8) ([16]byte, error) {
	var (
		ret0 = new([16]byte)
	)
	out := ret0
	err := _HoQuStorage.contract.Call(opts, out, "getOfferTariffGroup", id, num)
	return *ret0, err
}

// GetOfferTariffGroup is a free data retrieval call binding the contract method 0x74f627c5.
//
// Solidity: function getOfferTariffGroup(id bytes16, num uint8) constant returns(bytes16)
func (_HoQuStorage *HoQuStorageSession) GetOfferTariffGroup(id [16]byte, num uint8) ([16]byte, error) {
	return _HoQuStorage.Contract.GetOfferTariffGroup(&_HoQuStorage.CallOpts, id, num)
}

// GetOfferTariffGroup is a free data retrieval call binding the contract method 0x74f627c5.
//
// Solidity: function getOfferTariffGroup(id bytes16, num uint8) constant returns(bytes16)
func (_HoQuStorage *HoQuStorageCallerSession) GetOfferTariffGroup(id [16]byte, num uint8) ([16]byte, error) {
	return _HoQuStorage.Contract.GetOfferTariffGroup(&_HoQuStorage.CallOpts, id, num)
}

// GetTariffGroupTariff is a free data retrieval call binding the contract method 0x436e7b65.
//
// Solidity: function getTariffGroupTariff(id bytes16, num uint8) constant returns(bytes16)
func (_HoQuStorage *HoQuStorageCaller) GetTariffGroupTariff(opts *bind.CallOpts, id [16]byte, num uint8) ([16]byte, error) {
	var (
		ret0 = new([16]byte)
	)
	out := ret0
	err := _HoQuStorage.contract.Call(opts, out, "getTariffGroupTariff", id, num)
	return *ret0, err
}

// GetTariffGroupTariff is a free data retrieval call binding the contract method 0x436e7b65.
//
// Solidity: function getTariffGroupTariff(id bytes16, num uint8) constant returns(bytes16)
func (_HoQuStorage *HoQuStorageSession) GetTariffGroupTariff(id [16]byte, num uint8) ([16]byte, error) {
	return _HoQuStorage.Contract.GetTariffGroupTariff(&_HoQuStorage.CallOpts, id, num)
}

// GetTariffGroupTariff is a free data retrieval call binding the contract method 0x436e7b65.
//
// Solidity: function getTariffGroupTariff(id bytes16, num uint8) constant returns(bytes16)
func (_HoQuStorage *HoQuStorageCallerSession) GetTariffGroupTariff(id [16]byte, num uint8) ([16]byte, error) {
	return _HoQuStorage.Contract.GetTariffGroupTariff(&_HoQuStorage.CallOpts, id, num)
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
// Solidity: function offers( bytes16) constant returns(ownerId bytes16, networkId bytes16, merchantId bytes16, payerAddress address, name string, dataUrl string, numOfTariffGroups uint8, createdAt uint256, status uint8)
func (_HoQuStorage *HoQuStorageCaller) Offers(opts *bind.CallOpts, arg0 [16]byte) (struct {
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
	err := _HoQuStorage.contract.Call(opts, out, "offers", arg0)
	return *ret, err
}

// Offers is a free data retrieval call binding the contract method 0xc32d869b.
//
// Solidity: function offers( bytes16) constant returns(ownerId bytes16, networkId bytes16, merchantId bytes16, payerAddress address, name string, dataUrl string, numOfTariffGroups uint8, createdAt uint256, status uint8)
func (_HoQuStorage *HoQuStorageSession) Offers(arg0 [16]byte) (struct {
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
	return _HoQuStorage.Contract.Offers(&_HoQuStorage.CallOpts, arg0)
}

// Offers is a free data retrieval call binding the contract method 0xc32d869b.
//
// Solidity: function offers( bytes16) constant returns(ownerId bytes16, networkId bytes16, merchantId bytes16, payerAddress address, name string, dataUrl string, numOfTariffGroups uint8, createdAt uint256, status uint8)
func (_HoQuStorage *HoQuStorageCallerSession) Offers(arg0 [16]byte) (struct {
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
	return _HoQuStorage.Contract.Offers(&_HoQuStorage.CallOpts, arg0)
}

// Stats is a free data retrieval call binding the contract method 0xa2affa4d.
//
// Solidity: function stats( bytes16) constant returns(rating uint256, volume uint256, members uint256, alfa uint256, beta uint256, status uint8)
func (_HoQuStorage *HoQuStorageCaller) Stats(opts *bind.CallOpts, arg0 [16]byte) (struct {
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
	err := _HoQuStorage.contract.Call(opts, out, "stats", arg0)
	return *ret, err
}

// Stats is a free data retrieval call binding the contract method 0xa2affa4d.
//
// Solidity: function stats( bytes16) constant returns(rating uint256, volume uint256, members uint256, alfa uint256, beta uint256, status uint8)
func (_HoQuStorage *HoQuStorageSession) Stats(arg0 [16]byte) (struct {
	Rating  *big.Int
	Volume  *big.Int
	Members *big.Int
	Alfa    *big.Int
	Beta    *big.Int
	Status  uint8
}, error) {
	return _HoQuStorage.Contract.Stats(&_HoQuStorage.CallOpts, arg0)
}

// Stats is a free data retrieval call binding the contract method 0xa2affa4d.
//
// Solidity: function stats( bytes16) constant returns(rating uint256, volume uint256, members uint256, alfa uint256, beta uint256, status uint8)
func (_HoQuStorage *HoQuStorageCallerSession) Stats(arg0 [16]byte) (struct {
	Rating  *big.Int
	Volume  *big.Int
	Members *big.Int
	Alfa    *big.Int
	Beta    *big.Int
	Status  uint8
}, error) {
	return _HoQuStorage.Contract.Stats(&_HoQuStorage.CallOpts, arg0)
}

// TariffGroups is a free data retrieval call binding the contract method 0x8ade0027.
//
// Solidity: function tariffGroups( bytes16) constant returns(ownerId bytes16, name string, numOfTariffs uint8, createdAt uint256, status uint8)
func (_HoQuStorage *HoQuStorageCaller) TariffGroups(opts *bind.CallOpts, arg0 [16]byte) (struct {
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
	err := _HoQuStorage.contract.Call(opts, out, "tariffGroups", arg0)
	return *ret, err
}

// TariffGroups is a free data retrieval call binding the contract method 0x8ade0027.
//
// Solidity: function tariffGroups( bytes16) constant returns(ownerId bytes16, name string, numOfTariffs uint8, createdAt uint256, status uint8)
func (_HoQuStorage *HoQuStorageSession) TariffGroups(arg0 [16]byte) (struct {
	OwnerId      [16]byte
	Name         string
	NumOfTariffs uint8
	CreatedAt    *big.Int
	Status       uint8
}, error) {
	return _HoQuStorage.Contract.TariffGroups(&_HoQuStorage.CallOpts, arg0)
}

// TariffGroups is a free data retrieval call binding the contract method 0x8ade0027.
//
// Solidity: function tariffGroups( bytes16) constant returns(ownerId bytes16, name string, numOfTariffs uint8, createdAt uint256, status uint8)
func (_HoQuStorage *HoQuStorageCallerSession) TariffGroups(arg0 [16]byte) (struct {
	OwnerId      [16]byte
	Name         string
	NumOfTariffs uint8
	CreatedAt    *big.Int
	Status       uint8
}, error) {
	return _HoQuStorage.Contract.TariffGroups(&_HoQuStorage.CallOpts, arg0)
}

// Tariffs is a free data retrieval call binding the contract method 0xa2cbda41.
//
// Solidity: function tariffs( bytes16) constant returns(tariffGroupId bytes16, name string, action string, calcMethod string, price uint256, createdAt uint256, status uint8)
func (_HoQuStorage *HoQuStorageCaller) Tariffs(opts *bind.CallOpts, arg0 [16]byte) (struct {
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
	err := _HoQuStorage.contract.Call(opts, out, "tariffs", arg0)
	return *ret, err
}

// Tariffs is a free data retrieval call binding the contract method 0xa2cbda41.
//
// Solidity: function tariffs( bytes16) constant returns(tariffGroupId bytes16, name string, action string, calcMethod string, price uint256, createdAt uint256, status uint8)
func (_HoQuStorage *HoQuStorageSession) Tariffs(arg0 [16]byte) (struct {
	TariffGroupId [16]byte
	Name          string
	Action        string
	CalcMethod    string
	Price         *big.Int
	CreatedAt     *big.Int
	Status        uint8
}, error) {
	return _HoQuStorage.Contract.Tariffs(&_HoQuStorage.CallOpts, arg0)
}

// Tariffs is a free data retrieval call binding the contract method 0xa2cbda41.
//
// Solidity: function tariffs( bytes16) constant returns(tariffGroupId bytes16, name string, action string, calcMethod string, price uint256, createdAt uint256, status uint8)
func (_HoQuStorage *HoQuStorageCallerSession) Tariffs(arg0 [16]byte) (struct {
	TariffGroupId [16]byte
	Name          string
	Action        string
	CalcMethod    string
	Price         *big.Int
	CreatedAt     *big.Int
	Status        uint8
}, error) {
	return _HoQuStorage.Contract.Tariffs(&_HoQuStorage.CallOpts, arg0)
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
// Solidity: function users( bytes16) constant returns(ownerAddress address, numOfAddresses uint8, role string, kycLevel uint8, pubKey string, createdAt uint256, status uint8)
func (_HoQuStorage *HoQuStorageCaller) Users(opts *bind.CallOpts, arg0 [16]byte) (struct {
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
	err := _HoQuStorage.contract.Call(opts, out, "users", arg0)
	return *ret, err
}

// Users is a free data retrieval call binding the contract method 0x7520dd14.
//
// Solidity: function users( bytes16) constant returns(ownerAddress address, numOfAddresses uint8, role string, kycLevel uint8, pubKey string, createdAt uint256, status uint8)
func (_HoQuStorage *HoQuStorageSession) Users(arg0 [16]byte) (struct {
	OwnerAddress   common.Address
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
// Solidity: function users( bytes16) constant returns(ownerAddress address, numOfAddresses uint8, role string, kycLevel uint8, pubKey string, createdAt uint256, status uint8)
func (_HoQuStorage *HoQuStorageCallerSession) Users(arg0 [16]byte) (struct {
	OwnerAddress   common.Address
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

// AddOfferTariffGroup is a paid mutator transaction binding the contract method 0xf17eceb1.
//
// Solidity: function addOfferTariffGroup(id bytes16, tariffGroupId bytes16) returns()
func (_HoQuStorage *HoQuStorageTransactor) AddOfferTariffGroup(opts *bind.TransactOpts, id [16]byte, tariffGroupId [16]byte) (*types.Transaction, error) {
	return _HoQuStorage.contract.Transact(opts, "addOfferTariffGroup", id, tariffGroupId)
}

// AddOfferTariffGroup is a paid mutator transaction binding the contract method 0xf17eceb1.
//
// Solidity: function addOfferTariffGroup(id bytes16, tariffGroupId bytes16) returns()
func (_HoQuStorage *HoQuStorageSession) AddOfferTariffGroup(id [16]byte, tariffGroupId [16]byte) (*types.Transaction, error) {
	return _HoQuStorage.Contract.AddOfferTariffGroup(&_HoQuStorage.TransactOpts, id, tariffGroupId)
}

// AddOfferTariffGroup is a paid mutator transaction binding the contract method 0xf17eceb1.
//
// Solidity: function addOfferTariffGroup(id bytes16, tariffGroupId bytes16) returns()
func (_HoQuStorage *HoQuStorageTransactorSession) AddOfferTariffGroup(id [16]byte, tariffGroupId [16]byte) (*types.Transaction, error) {
	return _HoQuStorage.Contract.AddOfferTariffGroup(&_HoQuStorage.TransactOpts, id, tariffGroupId)
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

// SetOffer is a paid mutator transaction binding the contract method 0x1071d512.
//
// Solidity: function setOffer(id bytes16, ownerId bytes16, networkId bytes16, merchantId bytes16, payerAddress address, name string, dataUrl string, status uint8) returns()
func (_HoQuStorage *HoQuStorageTransactor) SetOffer(opts *bind.TransactOpts, id [16]byte, ownerId [16]byte, networkId [16]byte, merchantId [16]byte, payerAddress common.Address, name string, dataUrl string, status uint8) (*types.Transaction, error) {
	return _HoQuStorage.contract.Transact(opts, "setOffer", id, ownerId, networkId, merchantId, payerAddress, name, dataUrl, status)
}

// SetOffer is a paid mutator transaction binding the contract method 0x1071d512.
//
// Solidity: function setOffer(id bytes16, ownerId bytes16, networkId bytes16, merchantId bytes16, payerAddress address, name string, dataUrl string, status uint8) returns()
func (_HoQuStorage *HoQuStorageSession) SetOffer(id [16]byte, ownerId [16]byte, networkId [16]byte, merchantId [16]byte, payerAddress common.Address, name string, dataUrl string, status uint8) (*types.Transaction, error) {
	return _HoQuStorage.Contract.SetOffer(&_HoQuStorage.TransactOpts, id, ownerId, networkId, merchantId, payerAddress, name, dataUrl, status)
}

// SetOffer is a paid mutator transaction binding the contract method 0x1071d512.
//
// Solidity: function setOffer(id bytes16, ownerId bytes16, networkId bytes16, merchantId bytes16, payerAddress address, name string, dataUrl string, status uint8) returns()
func (_HoQuStorage *HoQuStorageTransactorSession) SetOffer(id [16]byte, ownerId [16]byte, networkId [16]byte, merchantId [16]byte, payerAddress common.Address, name string, dataUrl string, status uint8) (*types.Transaction, error) {
	return _HoQuStorage.Contract.SetOffer(&_HoQuStorage.TransactOpts, id, ownerId, networkId, merchantId, payerAddress, name, dataUrl, status)
}

// SetOfferTariffGroup is a paid mutator transaction binding the contract method 0x3e45850a.
//
// Solidity: function setOfferTariffGroup(id bytes16, num uint8, tariffGroupId bytes16) returns()
func (_HoQuStorage *HoQuStorageTransactor) SetOfferTariffGroup(opts *bind.TransactOpts, id [16]byte, num uint8, tariffGroupId [16]byte) (*types.Transaction, error) {
	return _HoQuStorage.contract.Transact(opts, "setOfferTariffGroup", id, num, tariffGroupId)
}

// SetOfferTariffGroup is a paid mutator transaction binding the contract method 0x3e45850a.
//
// Solidity: function setOfferTariffGroup(id bytes16, num uint8, tariffGroupId bytes16) returns()
func (_HoQuStorage *HoQuStorageSession) SetOfferTariffGroup(id [16]byte, num uint8, tariffGroupId [16]byte) (*types.Transaction, error) {
	return _HoQuStorage.Contract.SetOfferTariffGroup(&_HoQuStorage.TransactOpts, id, num, tariffGroupId)
}

// SetOfferTariffGroup is a paid mutator transaction binding the contract method 0x3e45850a.
//
// Solidity: function setOfferTariffGroup(id bytes16, num uint8, tariffGroupId bytes16) returns()
func (_HoQuStorage *HoQuStorageTransactorSession) SetOfferTariffGroup(id [16]byte, num uint8, tariffGroupId [16]byte) (*types.Transaction, error) {
	return _HoQuStorage.Contract.SetOfferTariffGroup(&_HoQuStorage.TransactOpts, id, num, tariffGroupId)
}

// SetStats is a paid mutator transaction binding the contract method 0xb1dbfe50.
//
// Solidity: function setStats(id bytes16, userId bytes16, rating uint256, volume uint256, members uint256, alfa uint256, beta uint256, status uint8) returns()
func (_HoQuStorage *HoQuStorageTransactor) SetStats(opts *bind.TransactOpts, id [16]byte, userId [16]byte, rating *big.Int, volume *big.Int, members *big.Int, alfa *big.Int, beta *big.Int, status uint8) (*types.Transaction, error) {
	return _HoQuStorage.contract.Transact(opts, "setStats", id, userId, rating, volume, members, alfa, beta, status)
}

// SetStats is a paid mutator transaction binding the contract method 0xb1dbfe50.
//
// Solidity: function setStats(id bytes16, userId bytes16, rating uint256, volume uint256, members uint256, alfa uint256, beta uint256, status uint8) returns()
func (_HoQuStorage *HoQuStorageSession) SetStats(id [16]byte, userId [16]byte, rating *big.Int, volume *big.Int, members *big.Int, alfa *big.Int, beta *big.Int, status uint8) (*types.Transaction, error) {
	return _HoQuStorage.Contract.SetStats(&_HoQuStorage.TransactOpts, id, userId, rating, volume, members, alfa, beta, status)
}

// SetStats is a paid mutator transaction binding the contract method 0xb1dbfe50.
//
// Solidity: function setStats(id bytes16, userId bytes16, rating uint256, volume uint256, members uint256, alfa uint256, beta uint256, status uint8) returns()
func (_HoQuStorage *HoQuStorageTransactorSession) SetStats(id [16]byte, userId [16]byte, rating *big.Int, volume *big.Int, members *big.Int, alfa *big.Int, beta *big.Int, status uint8) (*types.Transaction, error) {
	return _HoQuStorage.Contract.SetStats(&_HoQuStorage.TransactOpts, id, userId, rating, volume, members, alfa, beta, status)
}

// SetTariff is a paid mutator transaction binding the contract method 0x61fad57e.
//
// Solidity: function setTariff(id bytes16, tariffGroupId bytes16, name string, action string, calcMethod string, price uint256, status uint8) returns()
func (_HoQuStorage *HoQuStorageTransactor) SetTariff(opts *bind.TransactOpts, id [16]byte, tariffGroupId [16]byte, name string, action string, calcMethod string, price *big.Int, status uint8) (*types.Transaction, error) {
	return _HoQuStorage.contract.Transact(opts, "setTariff", id, tariffGroupId, name, action, calcMethod, price, status)
}

// SetTariff is a paid mutator transaction binding the contract method 0x61fad57e.
//
// Solidity: function setTariff(id bytes16, tariffGroupId bytes16, name string, action string, calcMethod string, price uint256, status uint8) returns()
func (_HoQuStorage *HoQuStorageSession) SetTariff(id [16]byte, tariffGroupId [16]byte, name string, action string, calcMethod string, price *big.Int, status uint8) (*types.Transaction, error) {
	return _HoQuStorage.Contract.SetTariff(&_HoQuStorage.TransactOpts, id, tariffGroupId, name, action, calcMethod, price, status)
}

// SetTariff is a paid mutator transaction binding the contract method 0x61fad57e.
//
// Solidity: function setTariff(id bytes16, tariffGroupId bytes16, name string, action string, calcMethod string, price uint256, status uint8) returns()
func (_HoQuStorage *HoQuStorageTransactorSession) SetTariff(id [16]byte, tariffGroupId [16]byte, name string, action string, calcMethod string, price *big.Int, status uint8) (*types.Transaction, error) {
	return _HoQuStorage.Contract.SetTariff(&_HoQuStorage.TransactOpts, id, tariffGroupId, name, action, calcMethod, price, status)
}

// SetTariffGroup is a paid mutator transaction binding the contract method 0xa7fbcf61.
//
// Solidity: function setTariffGroup(id bytes16, ownerId bytes16, name string, status uint8) returns()
func (_HoQuStorage *HoQuStorageTransactor) SetTariffGroup(opts *bind.TransactOpts, id [16]byte, ownerId [16]byte, name string, status uint8) (*types.Transaction, error) {
	return _HoQuStorage.contract.Transact(opts, "setTariffGroup", id, ownerId, name, status)
}

// SetTariffGroup is a paid mutator transaction binding the contract method 0xa7fbcf61.
//
// Solidity: function setTariffGroup(id bytes16, ownerId bytes16, name string, status uint8) returns()
func (_HoQuStorage *HoQuStorageSession) SetTariffGroup(id [16]byte, ownerId [16]byte, name string, status uint8) (*types.Transaction, error) {
	return _HoQuStorage.Contract.SetTariffGroup(&_HoQuStorage.TransactOpts, id, ownerId, name, status)
}

// SetTariffGroup is a paid mutator transaction binding the contract method 0xa7fbcf61.
//
// Solidity: function setTariffGroup(id bytes16, ownerId bytes16, name string, status uint8) returns()
func (_HoQuStorage *HoQuStorageTransactorSession) SetTariffGroup(id [16]byte, ownerId [16]byte, name string, status uint8) (*types.Transaction, error) {
	return _HoQuStorage.Contract.SetTariffGroup(&_HoQuStorage.TransactOpts, id, ownerId, name, status)
}

// SetTariffGroupTariff is a paid mutator transaction binding the contract method 0xfafe27d7.
//
// Solidity: function setTariffGroupTariff(id bytes16, num uint8, tariffId bytes16) returns()
func (_HoQuStorage *HoQuStorageTransactor) SetTariffGroupTariff(opts *bind.TransactOpts, id [16]byte, num uint8, tariffId [16]byte) (*types.Transaction, error) {
	return _HoQuStorage.contract.Transact(opts, "setTariffGroupTariff", id, num, tariffId)
}

// SetTariffGroupTariff is a paid mutator transaction binding the contract method 0xfafe27d7.
//
// Solidity: function setTariffGroupTariff(id bytes16, num uint8, tariffId bytes16) returns()
func (_HoQuStorage *HoQuStorageSession) SetTariffGroupTariff(id [16]byte, num uint8, tariffId [16]byte) (*types.Transaction, error) {
	return _HoQuStorage.Contract.SetTariffGroupTariff(&_HoQuStorage.TransactOpts, id, num, tariffId)
}

// SetTariffGroupTariff is a paid mutator transaction binding the contract method 0xfafe27d7.
//
// Solidity: function setTariffGroupTariff(id bytes16, num uint8, tariffId bytes16) returns()
func (_HoQuStorage *HoQuStorageTransactorSession) SetTariffGroupTariff(id [16]byte, num uint8, tariffId [16]byte) (*types.Transaction, error) {
	return _HoQuStorage.Contract.SetTariffGroupTariff(&_HoQuStorage.TransactOpts, id, num, tariffId)
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
	UserId       [16]byte
	Name         string
	Raw          types.Log // Blockchain specific contextual infos
}

// FilterIdentificationAdded is a free log retrieval operation binding the contract event 0x8c03773782e9ed7fe9111421165a6cf12ce5855a4181cd1a8166d47d719ac078.
//
// Solidity: event IdentificationAdded(ownerAddress indexed address, id bytes16, userId bytes16, name string)
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

// WatchIdentificationAdded is a free log subscription operation binding the contract event 0x8c03773782e9ed7fe9111421165a6cf12ce5855a4181cd1a8166d47d719ac078.
//
// Solidity: event IdentificationAdded(ownerAddress indexed address, id bytes16, userId bytes16, name string)
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

// HoQuStorageOfferTariffGroupAddedIterator is returned from FilterOfferTariffGroupAdded and is used to iterate over the raw logs and unpacked data for OfferTariffGroupAdded events raised by the HoQuStorage contract.
type HoQuStorageOfferTariffGroupAddedIterator struct {
	Event *HoQuStorageOfferTariffGroupAdded // Event containing the contract specifics and raw log

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
func (it *HoQuStorageOfferTariffGroupAddedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(HoQuStorageOfferTariffGroupAdded)
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
		it.Event = new(HoQuStorageOfferTariffGroupAdded)
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
func (it *HoQuStorageOfferTariffGroupAddedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *HoQuStorageOfferTariffGroupAddedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// HoQuStorageOfferTariffGroupAdded represents a OfferTariffGroupAdded event raised by the HoQuStorage contract.
type HoQuStorageOfferTariffGroupAdded struct {
	OwnerAddress common.Address
	Id           [16]byte
	TariffId     [16]byte
	Raw          types.Log // Blockchain specific contextual infos
}

// FilterOfferTariffGroupAdded is a free log retrieval operation binding the contract event 0x695164079ea85cb7933398c55f763b3dd70fe1cfe293a3b90aa22f9cf6270dab.
//
// Solidity: event OfferTariffGroupAdded(ownerAddress indexed address, id bytes16, tariff_id bytes16)
func (_HoQuStorage *HoQuStorageFilterer) FilterOfferTariffGroupAdded(opts *bind.FilterOpts, ownerAddress []common.Address) (*HoQuStorageOfferTariffGroupAddedIterator, error) {

	var ownerAddressRule []interface{}
	for _, ownerAddressItem := range ownerAddress {
		ownerAddressRule = append(ownerAddressRule, ownerAddressItem)
	}

	logs, sub, err := _HoQuStorage.contract.FilterLogs(opts, "OfferTariffGroupAdded", ownerAddressRule)
	if err != nil {
		return nil, err
	}
	return &HoQuStorageOfferTariffGroupAddedIterator{contract: _HoQuStorage.contract, event: "OfferTariffGroupAdded", logs: logs, sub: sub}, nil
}

// WatchOfferTariffGroupAdded is a free log subscription operation binding the contract event 0x695164079ea85cb7933398c55f763b3dd70fe1cfe293a3b90aa22f9cf6270dab.
//
// Solidity: event OfferTariffGroupAdded(ownerAddress indexed address, id bytes16, tariff_id bytes16)
func (_HoQuStorage *HoQuStorageFilterer) WatchOfferTariffGroupAdded(opts *bind.WatchOpts, sink chan<- *HoQuStorageOfferTariffGroupAdded, ownerAddress []common.Address) (event.Subscription, error) {

	var ownerAddressRule []interface{}
	for _, ownerAddressItem := range ownerAddress {
		ownerAddressRule = append(ownerAddressRule, ownerAddressItem)
	}

	logs, sub, err := _HoQuStorage.contract.WatchLogs(opts, "OfferTariffGroupAdded", ownerAddressRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(HoQuStorageOfferTariffGroupAdded)
				if err := _HoQuStorage.contract.UnpackLog(event, "OfferTariffGroupAdded", log); err != nil {
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

// HoQuStorageTariffAddedIterator is returned from FilterTariffAdded and is used to iterate over the raw logs and unpacked data for TariffAdded events raised by the HoQuStorage contract.
type HoQuStorageTariffAddedIterator struct {
	Event *HoQuStorageTariffAdded // Event containing the contract specifics and raw log

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
func (it *HoQuStorageTariffAddedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(HoQuStorageTariffAdded)
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
		it.Event = new(HoQuStorageTariffAdded)
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
func (it *HoQuStorageTariffAddedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *HoQuStorageTariffAddedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// HoQuStorageTariffAdded represents a TariffAdded event raised by the HoQuStorage contract.
type HoQuStorageTariffAdded struct {
	OwnerAddress common.Address
	Id           [16]byte
	Name         string
	Raw          types.Log // Blockchain specific contextual infos
}

// FilterTariffAdded is a free log retrieval operation binding the contract event 0x4ad20f56c3c96698f2e84d660e8c5e39341fd79929743804d261beb65ffc5b47.
//
// Solidity: event TariffAdded(ownerAddress indexed address, id bytes16, name string)
func (_HoQuStorage *HoQuStorageFilterer) FilterTariffAdded(opts *bind.FilterOpts, ownerAddress []common.Address) (*HoQuStorageTariffAddedIterator, error) {

	var ownerAddressRule []interface{}
	for _, ownerAddressItem := range ownerAddress {
		ownerAddressRule = append(ownerAddressRule, ownerAddressItem)
	}

	logs, sub, err := _HoQuStorage.contract.FilterLogs(opts, "TariffAdded", ownerAddressRule)
	if err != nil {
		return nil, err
	}
	return &HoQuStorageTariffAddedIterator{contract: _HoQuStorage.contract, event: "TariffAdded", logs: logs, sub: sub}, nil
}

// WatchTariffAdded is a free log subscription operation binding the contract event 0x4ad20f56c3c96698f2e84d660e8c5e39341fd79929743804d261beb65ffc5b47.
//
// Solidity: event TariffAdded(ownerAddress indexed address, id bytes16, name string)
func (_HoQuStorage *HoQuStorageFilterer) WatchTariffAdded(opts *bind.WatchOpts, sink chan<- *HoQuStorageTariffAdded, ownerAddress []common.Address) (event.Subscription, error) {

	var ownerAddressRule []interface{}
	for _, ownerAddressItem := range ownerAddress {
		ownerAddressRule = append(ownerAddressRule, ownerAddressItem)
	}

	logs, sub, err := _HoQuStorage.contract.WatchLogs(opts, "TariffAdded", ownerAddressRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(HoQuStorageTariffAdded)
				if err := _HoQuStorage.contract.UnpackLog(event, "TariffAdded", log); err != nil {
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

// HoQuStorageTariffGroupAddedIterator is returned from FilterTariffGroupAdded and is used to iterate over the raw logs and unpacked data for TariffGroupAdded events raised by the HoQuStorage contract.
type HoQuStorageTariffGroupAddedIterator struct {
	Event *HoQuStorageTariffGroupAdded // Event containing the contract specifics and raw log

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
func (it *HoQuStorageTariffGroupAddedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(HoQuStorageTariffGroupAdded)
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
		it.Event = new(HoQuStorageTariffGroupAdded)
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
func (it *HoQuStorageTariffGroupAddedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *HoQuStorageTariffGroupAddedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// HoQuStorageTariffGroupAdded represents a TariffGroupAdded event raised by the HoQuStorage contract.
type HoQuStorageTariffGroupAdded struct {
	OwnerAddress common.Address
	Id           [16]byte
	Name         string
	Raw          types.Log // Blockchain specific contextual infos
}

// FilterTariffGroupAdded is a free log retrieval operation binding the contract event 0x0d945b4e88c22df3f6d39d8866fac2dd4964090d9d9c04fc74f33c5695e561df.
//
// Solidity: event TariffGroupAdded(ownerAddress indexed address, id bytes16, name string)
func (_HoQuStorage *HoQuStorageFilterer) FilterTariffGroupAdded(opts *bind.FilterOpts, ownerAddress []common.Address) (*HoQuStorageTariffGroupAddedIterator, error) {

	var ownerAddressRule []interface{}
	for _, ownerAddressItem := range ownerAddress {
		ownerAddressRule = append(ownerAddressRule, ownerAddressItem)
	}

	logs, sub, err := _HoQuStorage.contract.FilterLogs(opts, "TariffGroupAdded", ownerAddressRule)
	if err != nil {
		return nil, err
	}
	return &HoQuStorageTariffGroupAddedIterator{contract: _HoQuStorage.contract, event: "TariffGroupAdded", logs: logs, sub: sub}, nil
}

// WatchTariffGroupAdded is a free log subscription operation binding the contract event 0x0d945b4e88c22df3f6d39d8866fac2dd4964090d9d9c04fc74f33c5695e561df.
//
// Solidity: event TariffGroupAdded(ownerAddress indexed address, id bytes16, name string)
func (_HoQuStorage *HoQuStorageFilterer) WatchTariffGroupAdded(opts *bind.WatchOpts, sink chan<- *HoQuStorageTariffGroupAdded, ownerAddress []common.Address) (event.Subscription, error) {

	var ownerAddressRule []interface{}
	for _, ownerAddressItem := range ownerAddress {
		ownerAddressRule = append(ownerAddressRule, ownerAddressItem)
	}

	logs, sub, err := _HoQuStorage.contract.WatchLogs(opts, "TariffGroupAdded", ownerAddressRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(HoQuStorageTariffGroupAdded)
				if err := _HoQuStorage.contract.UnpackLog(event, "TariffGroupAdded", log); err != nil {
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

// HoQuStorageTrackerRegisteredIterator is returned from FilterTrackerRegistered and is used to iterate over the raw logs and unpacked data for TrackerRegistered events raised by the HoQuStorage contract.
type HoQuStorageTrackerRegisteredIterator struct {
	Event *HoQuStorageTrackerRegistered // Event containing the contract specifics and raw log

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
func (it *HoQuStorageTrackerRegisteredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(HoQuStorageTrackerRegistered)
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
		it.Event = new(HoQuStorageTrackerRegistered)
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
func (it *HoQuStorageTrackerRegisteredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *HoQuStorageTrackerRegisteredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// HoQuStorageTrackerRegistered represents a TrackerRegistered event raised by the HoQuStorage contract.
type HoQuStorageTrackerRegistered struct {
	OwnerAddress common.Address
	Id           [16]byte
	Name         string
	Raw          types.Log // Blockchain specific contextual infos
}

// FilterTrackerRegistered is a free log retrieval operation binding the contract event 0xb2f67f32df451ce918abd944fd5c8a29589848693057913c5060a8e0f293b388.
//
// Solidity: event TrackerRegistered(ownerAddress indexed address, id bytes16, name string)
func (_HoQuStorage *HoQuStorageFilterer) FilterTrackerRegistered(opts *bind.FilterOpts, ownerAddress []common.Address) (*HoQuStorageTrackerRegisteredIterator, error) {

	var ownerAddressRule []interface{}
	for _, ownerAddressItem := range ownerAddress {
		ownerAddressRule = append(ownerAddressRule, ownerAddressItem)
	}

	logs, sub, err := _HoQuStorage.contract.FilterLogs(opts, "TrackerRegistered", ownerAddressRule)
	if err != nil {
		return nil, err
	}
	return &HoQuStorageTrackerRegisteredIterator{contract: _HoQuStorage.contract, event: "TrackerRegistered", logs: logs, sub: sub}, nil
}

// WatchTrackerRegistered is a free log subscription operation binding the contract event 0xb2f67f32df451ce918abd944fd5c8a29589848693057913c5060a8e0f293b388.
//
// Solidity: event TrackerRegistered(ownerAddress indexed address, id bytes16, name string)
func (_HoQuStorage *HoQuStorageFilterer) WatchTrackerRegistered(opts *bind.WatchOpts, sink chan<- *HoQuStorageTrackerRegistered, ownerAddress []common.Address) (event.Subscription, error) {

	var ownerAddressRule []interface{}
	for _, ownerAddressItem := range ownerAddress {
		ownerAddressRule = append(ownerAddressRule, ownerAddressItem)
	}

	logs, sub, err := _HoQuStorage.contract.WatchLogs(opts, "TrackerRegistered", ownerAddressRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(HoQuStorageTrackerRegistered)
				if err := _HoQuStorage.contract.UnpackLog(event, "TrackerRegistered", log); err != nil {
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

// HoQuStorageUserAddressAddedIterator is returned from FilterUserAddressAdded and is used to iterate over the raw logs and unpacked data for UserAddressAdded events raised by the HoQuStorage contract.
type HoQuStorageUserAddressAddedIterator struct {
	Event *HoQuStorageUserAddressAdded // Event containing the contract specifics and raw log

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
func (it *HoQuStorageUserAddressAddedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(HoQuStorageUserAddressAdded)
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
		it.Event = new(HoQuStorageUserAddressAdded)
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
func (it *HoQuStorageUserAddressAddedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *HoQuStorageUserAddressAddedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// HoQuStorageUserAddressAdded represents a UserAddressAdded event raised by the HoQuStorage contract.
type HoQuStorageUserAddressAdded struct {
	OwnerAddress      common.Address
	AdditionalAddress common.Address
	Id                [16]byte
	Raw               types.Log // Blockchain specific contextual infos
}

// FilterUserAddressAdded is a free log retrieval operation binding the contract event 0xb38b406c043d4a3feaa0192d2e1b02f317fa73c284eb3e70b9a1fde8612af1b3.
//
// Solidity: event UserAddressAdded(ownerAddress indexed address, additionalAddress address, id bytes16)
func (_HoQuStorage *HoQuStorageFilterer) FilterUserAddressAdded(opts *bind.FilterOpts, ownerAddress []common.Address) (*HoQuStorageUserAddressAddedIterator, error) {

	var ownerAddressRule []interface{}
	for _, ownerAddressItem := range ownerAddress {
		ownerAddressRule = append(ownerAddressRule, ownerAddressItem)
	}

	logs, sub, err := _HoQuStorage.contract.FilterLogs(opts, "UserAddressAdded", ownerAddressRule)
	if err != nil {
		return nil, err
	}
	return &HoQuStorageUserAddressAddedIterator{contract: _HoQuStorage.contract, event: "UserAddressAdded", logs: logs, sub: sub}, nil
}

// WatchUserAddressAdded is a free log subscription operation binding the contract event 0xb38b406c043d4a3feaa0192d2e1b02f317fa73c284eb3e70b9a1fde8612af1b3.
//
// Solidity: event UserAddressAdded(ownerAddress indexed address, additionalAddress address, id bytes16)
func (_HoQuStorage *HoQuStorageFilterer) WatchUserAddressAdded(opts *bind.WatchOpts, sink chan<- *HoQuStorageUserAddressAdded, ownerAddress []common.Address) (event.Subscription, error) {

	var ownerAddressRule []interface{}
	for _, ownerAddressItem := range ownerAddress {
		ownerAddressRule = append(ownerAddressRule, ownerAddressItem)
	}

	logs, sub, err := _HoQuStorage.contract.WatchLogs(opts, "UserAddressAdded", ownerAddressRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(HoQuStorageUserAddressAdded)
				if err := _HoQuStorage.contract.UnpackLog(event, "UserAddressAdded", log); err != nil {
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

// HoQuStorageUserRegisteredIterator is returned from FilterUserRegistered and is used to iterate over the raw logs and unpacked data for UserRegistered events raised by the HoQuStorage contract.
type HoQuStorageUserRegisteredIterator struct {
	Event *HoQuStorageUserRegistered // Event containing the contract specifics and raw log

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
func (it *HoQuStorageUserRegisteredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(HoQuStorageUserRegistered)
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
		it.Event = new(HoQuStorageUserRegistered)
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
func (it *HoQuStorageUserRegisteredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *HoQuStorageUserRegisteredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// HoQuStorageUserRegistered represents a UserRegistered event raised by the HoQuStorage contract.
type HoQuStorageUserRegistered struct {
	OwnerAddress common.Address
	Id           [16]byte
	Role         string
	Raw          types.Log // Blockchain specific contextual infos
}

// FilterUserRegistered is a free log retrieval operation binding the contract event 0xf661ee1472892faaf2bb68cc6874f2759c9148b4234cac6b88f6ca362652f275.
//
// Solidity: event UserRegistered(ownerAddress indexed address, id bytes16, role string)
func (_HoQuStorage *HoQuStorageFilterer) FilterUserRegistered(opts *bind.FilterOpts, ownerAddress []common.Address) (*HoQuStorageUserRegisteredIterator, error) {

	var ownerAddressRule []interface{}
	for _, ownerAddressItem := range ownerAddress {
		ownerAddressRule = append(ownerAddressRule, ownerAddressItem)
	}

	logs, sub, err := _HoQuStorage.contract.FilterLogs(opts, "UserRegistered", ownerAddressRule)
	if err != nil {
		return nil, err
	}
	return &HoQuStorageUserRegisteredIterator{contract: _HoQuStorage.contract, event: "UserRegistered", logs: logs, sub: sub}, nil
}

// WatchUserRegistered is a free log subscription operation binding the contract event 0xf661ee1472892faaf2bb68cc6874f2759c9148b4234cac6b88f6ca362652f275.
//
// Solidity: event UserRegistered(ownerAddress indexed address, id bytes16, role string)
func (_HoQuStorage *HoQuStorageFilterer) WatchUserRegistered(opts *bind.WatchOpts, sink chan<- *HoQuStorageUserRegistered, ownerAddress []common.Address) (event.Subscription, error) {

	var ownerAddressRule []interface{}
	for _, ownerAddressItem := range ownerAddress {
		ownerAddressRule = append(ownerAddressRule, ownerAddressItem)
	}

	logs, sub, err := _HoQuStorage.contract.WatchLogs(opts, "UserRegistered", ownerAddressRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(HoQuStorageUserRegistered)
				if err := _HoQuStorage.contract.UnpackLog(event, "UserRegistered", log); err != nil {
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
