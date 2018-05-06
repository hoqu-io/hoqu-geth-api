// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package rater

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

// HoQuRaterABI is the input ABI used to generate the binding from.
const HoQuRaterABI = "[{\"constant\":false,\"inputs\":[{\"name\":\"offerId\",\"type\":\"bytes16\"},{\"name\":\"trackerId\",\"type\":\"bytes16\"},{\"name\":\"affiliateId\",\"type\":\"bytes16\"},{\"name\":\"price\",\"type\":\"uint256\"}],\"name\":\"processTransactLead\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"storageAddress\",\"type\":\"address\"}],\"name\":\"setStorageAddress\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"config\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"configAddress\",\"type\":\"address\"}],\"name\":\"setConfigAddress\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"store\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"offerId\",\"type\":\"bytes16\"},{\"name\":\"trackerId\",\"type\":\"bytes16\"},{\"name\":\"affiliateId\",\"type\":\"bytes16\"},{\"name\":\"price\",\"type\":\"uint256\"}],\"name\":\"processAddLead\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"name\":\"configAddress\",\"type\":\"address\"},{\"name\":\"storageAddress\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"}]"

// HoQuRaterBin is the compiled bytecode used for deploying new contracts.
const HoQuRaterBin = `0x608060405234801561001057600080fd5b5060405160408061043b83398101604052805160209091015160008054600160a060020a03938416600160a060020a031991821617909155600180549390921692169190911790556103d4806100676000396000f3006080604052600436106100775763ffffffff7c01000000000000000000000000000000000000000000000000000000006000350416634943df38811461007c57806359b910d6146100b857806379502c55146100d957806383a12de91461010a578063975057e71461012b578063c5f6dd851461007c575b600080fd5b34801561008857600080fd5b506100b66fffffffffffffffffffffffffffffffff1960043581169060243581169060443516606435610140565b005b3480156100c457600080fd5b506100b6600160a060020a03600435166101e8565b3480156100e557600080fd5b506100ee6102b9565b60408051600160a060020a039092168252519081900360200190f35b34801561011657600080fd5b506100b6600160a060020a03600435166102c8565b34801561013757600080fd5b506100ee610399565b60008054604080517fbabcc539000000000000000000000000000000000000000000000000000000008152600160a060020a0333811660048301529151919092169263babcc53992602480820193602093909283900390910190829087803b1580156101ab57600080fd5b505af11580156101bf573d6000803e3d6000fd5b505050506040513d60208110156101d557600080fd5b505115156101e257600080fd5b50505050565b60008054604080517fbabcc539000000000000000000000000000000000000000000000000000000008152600160a060020a0333811660048301529151919092169263babcc53992602480820193602093909283900390910190829087803b15801561025357600080fd5b505af1158015610267573d6000803e3d6000fd5b505050506040513d602081101561027d57600080fd5b5051151561028a57600080fd5b6001805473ffffffffffffffffffffffffffffffffffffffff1916600160a060020a0392909216919091179055565b600054600160a060020a031681565b60008054604080517fbabcc539000000000000000000000000000000000000000000000000000000008152600160a060020a0333811660048301529151919092169263babcc53992602480820193602093909283900390910190829087803b15801561033357600080fd5b505af1158015610347573d6000803e3d6000fd5b505050506040513d602081101561035d57600080fd5b5051151561036a57600080fd5b6000805473ffffffffffffffffffffffffffffffffffffffff1916600160a060020a0392909216919091179055565b600154600160a060020a0316815600a165627a7a7230582097b4c87944770a7a281391d58b55ff75c3ef1fc5aecf162e972cfa04fbddbef20029`

// DeployHoQuRater deploys a new Ethereum contract, binding an instance of HoQuRater to it.
func DeployHoQuRater(auth *bind.TransactOpts, backend bind.ContractBackend, configAddress common.Address, storageAddress common.Address) (common.Address, *types.Transaction, *HoQuRater, error) {
	parsed, err := abi.JSON(strings.NewReader(HoQuRaterABI))
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	address, tx, contract, err := bind.DeployContract(auth, parsed, common.FromHex(HoQuRaterBin), backend, configAddress, storageAddress)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &HoQuRater{HoQuRaterCaller: HoQuRaterCaller{contract: contract}, HoQuRaterTransactor: HoQuRaterTransactor{contract: contract}, HoQuRaterFilterer: HoQuRaterFilterer{contract: contract}}, nil
}

// HoQuRater is an auto generated Go binding around an Ethereum contract.
type HoQuRater struct {
	HoQuRaterCaller     // Read-only binding to the contract
	HoQuRaterTransactor // Write-only binding to the contract
	HoQuRaterFilterer   // Log filterer for contract events
}

// HoQuRaterCaller is an auto generated read-only Go binding around an Ethereum contract.
type HoQuRaterCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// HoQuRaterTransactor is an auto generated write-only Go binding around an Ethereum contract.
type HoQuRaterTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// HoQuRaterFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type HoQuRaterFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// HoQuRaterSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type HoQuRaterSession struct {
	Contract     *HoQuRater        // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// HoQuRaterCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type HoQuRaterCallerSession struct {
	Contract *HoQuRaterCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts    // Call options to use throughout this session
}

// HoQuRaterTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type HoQuRaterTransactorSession struct {
	Contract     *HoQuRaterTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts    // Transaction auth options to use throughout this session
}

// HoQuRaterRaw is an auto generated low-level Go binding around an Ethereum contract.
type HoQuRaterRaw struct {
	Contract *HoQuRater // Generic contract binding to access the raw methods on
}

// HoQuRaterCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type HoQuRaterCallerRaw struct {
	Contract *HoQuRaterCaller // Generic read-only contract binding to access the raw methods on
}

// HoQuRaterTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type HoQuRaterTransactorRaw struct {
	Contract *HoQuRaterTransactor // Generic write-only contract binding to access the raw methods on
}

// NewHoQuRater creates a new instance of HoQuRater, bound to a specific deployed contract.
func NewHoQuRater(address common.Address, backend bind.ContractBackend) (*HoQuRater, error) {
	contract, err := bindHoQuRater(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &HoQuRater{HoQuRaterCaller: HoQuRaterCaller{contract: contract}, HoQuRaterTransactor: HoQuRaterTransactor{contract: contract}, HoQuRaterFilterer: HoQuRaterFilterer{contract: contract}}, nil
}

// NewHoQuRaterCaller creates a new read-only instance of HoQuRater, bound to a specific deployed contract.
func NewHoQuRaterCaller(address common.Address, caller bind.ContractCaller) (*HoQuRaterCaller, error) {
	contract, err := bindHoQuRater(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &HoQuRaterCaller{contract: contract}, nil
}

// NewHoQuRaterTransactor creates a new write-only instance of HoQuRater, bound to a specific deployed contract.
func NewHoQuRaterTransactor(address common.Address, transactor bind.ContractTransactor) (*HoQuRaterTransactor, error) {
	contract, err := bindHoQuRater(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &HoQuRaterTransactor{contract: contract}, nil
}

// NewHoQuRaterFilterer creates a new log filterer instance of HoQuRater, bound to a specific deployed contract.
func NewHoQuRaterFilterer(address common.Address, filterer bind.ContractFilterer) (*HoQuRaterFilterer, error) {
	contract, err := bindHoQuRater(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &HoQuRaterFilterer{contract: contract}, nil
}

// bindHoQuRater binds a generic wrapper to an already deployed contract.
func bindHoQuRater(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(HoQuRaterABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_HoQuRater *HoQuRaterRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _HoQuRater.Contract.HoQuRaterCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_HoQuRater *HoQuRaterRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _HoQuRater.Contract.HoQuRaterTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_HoQuRater *HoQuRaterRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _HoQuRater.Contract.HoQuRaterTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_HoQuRater *HoQuRaterCallerRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _HoQuRater.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_HoQuRater *HoQuRaterTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _HoQuRater.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_HoQuRater *HoQuRaterTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _HoQuRater.Contract.contract.Transact(opts, method, params...)
}

// Config is a free data retrieval call binding the contract method 0x79502c55.
//
// Solidity: function config() constant returns(address)
func (_HoQuRater *HoQuRaterCaller) Config(opts *bind.CallOpts) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _HoQuRater.contract.Call(opts, out, "config")
	return *ret0, err
}

// Config is a free data retrieval call binding the contract method 0x79502c55.
//
// Solidity: function config() constant returns(address)
func (_HoQuRater *HoQuRaterSession) Config() (common.Address, error) {
	return _HoQuRater.Contract.Config(&_HoQuRater.CallOpts)
}

// Config is a free data retrieval call binding the contract method 0x79502c55.
//
// Solidity: function config() constant returns(address)
func (_HoQuRater *HoQuRaterCallerSession) Config() (common.Address, error) {
	return _HoQuRater.Contract.Config(&_HoQuRater.CallOpts)
}

// Store is a free data retrieval call binding the contract method 0x975057e7.
//
// Solidity: function store() constant returns(address)
func (_HoQuRater *HoQuRaterCaller) Store(opts *bind.CallOpts) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _HoQuRater.contract.Call(opts, out, "store")
	return *ret0, err
}

// Store is a free data retrieval call binding the contract method 0x975057e7.
//
// Solidity: function store() constant returns(address)
func (_HoQuRater *HoQuRaterSession) Store() (common.Address, error) {
	return _HoQuRater.Contract.Store(&_HoQuRater.CallOpts)
}

// Store is a free data retrieval call binding the contract method 0x975057e7.
//
// Solidity: function store() constant returns(address)
func (_HoQuRater *HoQuRaterCallerSession) Store() (common.Address, error) {
	return _HoQuRater.Contract.Store(&_HoQuRater.CallOpts)
}

// ProcessAddLead is a paid mutator transaction binding the contract method 0xc5f6dd85.
//
// Solidity: function processAddLead(offerId bytes16, trackerId bytes16, affiliateId bytes16, price uint256) returns()
func (_HoQuRater *HoQuRaterTransactor) ProcessAddLead(opts *bind.TransactOpts, offerId [16]byte, trackerId [16]byte, affiliateId [16]byte, price *big.Int) (*types.Transaction, error) {
	return _HoQuRater.contract.Transact(opts, "processAddLead", offerId, trackerId, affiliateId, price)
}

// ProcessAddLead is a paid mutator transaction binding the contract method 0xc5f6dd85.
//
// Solidity: function processAddLead(offerId bytes16, trackerId bytes16, affiliateId bytes16, price uint256) returns()
func (_HoQuRater *HoQuRaterSession) ProcessAddLead(offerId [16]byte, trackerId [16]byte, affiliateId [16]byte, price *big.Int) (*types.Transaction, error) {
	return _HoQuRater.Contract.ProcessAddLead(&_HoQuRater.TransactOpts, offerId, trackerId, affiliateId, price)
}

// ProcessAddLead is a paid mutator transaction binding the contract method 0xc5f6dd85.
//
// Solidity: function processAddLead(offerId bytes16, trackerId bytes16, affiliateId bytes16, price uint256) returns()
func (_HoQuRater *HoQuRaterTransactorSession) ProcessAddLead(offerId [16]byte, trackerId [16]byte, affiliateId [16]byte, price *big.Int) (*types.Transaction, error) {
	return _HoQuRater.Contract.ProcessAddLead(&_HoQuRater.TransactOpts, offerId, trackerId, affiliateId, price)
}

// ProcessTransactLead is a paid mutator transaction binding the contract method 0x4943df38.
//
// Solidity: function processTransactLead(offerId bytes16, trackerId bytes16, affiliateId bytes16, price uint256) returns()
func (_HoQuRater *HoQuRaterTransactor) ProcessTransactLead(opts *bind.TransactOpts, offerId [16]byte, trackerId [16]byte, affiliateId [16]byte, price *big.Int) (*types.Transaction, error) {
	return _HoQuRater.contract.Transact(opts, "processTransactLead", offerId, trackerId, affiliateId, price)
}

// ProcessTransactLead is a paid mutator transaction binding the contract method 0x4943df38.
//
// Solidity: function processTransactLead(offerId bytes16, trackerId bytes16, affiliateId bytes16, price uint256) returns()
func (_HoQuRater *HoQuRaterSession) ProcessTransactLead(offerId [16]byte, trackerId [16]byte, affiliateId [16]byte, price *big.Int) (*types.Transaction, error) {
	return _HoQuRater.Contract.ProcessTransactLead(&_HoQuRater.TransactOpts, offerId, trackerId, affiliateId, price)
}

// ProcessTransactLead is a paid mutator transaction binding the contract method 0x4943df38.
//
// Solidity: function processTransactLead(offerId bytes16, trackerId bytes16, affiliateId bytes16, price uint256) returns()
func (_HoQuRater *HoQuRaterTransactorSession) ProcessTransactLead(offerId [16]byte, trackerId [16]byte, affiliateId [16]byte, price *big.Int) (*types.Transaction, error) {
	return _HoQuRater.Contract.ProcessTransactLead(&_HoQuRater.TransactOpts, offerId, trackerId, affiliateId, price)
}

// SetConfigAddress is a paid mutator transaction binding the contract method 0x83a12de9.
//
// Solidity: function setConfigAddress(configAddress address) returns()
func (_HoQuRater *HoQuRaterTransactor) SetConfigAddress(opts *bind.TransactOpts, configAddress common.Address) (*types.Transaction, error) {
	return _HoQuRater.contract.Transact(opts, "setConfigAddress", configAddress)
}

// SetConfigAddress is a paid mutator transaction binding the contract method 0x83a12de9.
//
// Solidity: function setConfigAddress(configAddress address) returns()
func (_HoQuRater *HoQuRaterSession) SetConfigAddress(configAddress common.Address) (*types.Transaction, error) {
	return _HoQuRater.Contract.SetConfigAddress(&_HoQuRater.TransactOpts, configAddress)
}

// SetConfigAddress is a paid mutator transaction binding the contract method 0x83a12de9.
//
// Solidity: function setConfigAddress(configAddress address) returns()
func (_HoQuRater *HoQuRaterTransactorSession) SetConfigAddress(configAddress common.Address) (*types.Transaction, error) {
	return _HoQuRater.Contract.SetConfigAddress(&_HoQuRater.TransactOpts, configAddress)
}

// SetStorageAddress is a paid mutator transaction binding the contract method 0x59b910d6.
//
// Solidity: function setStorageAddress(storageAddress address) returns()
func (_HoQuRater *HoQuRaterTransactor) SetStorageAddress(opts *bind.TransactOpts, storageAddress common.Address) (*types.Transaction, error) {
	return _HoQuRater.contract.Transact(opts, "setStorageAddress", storageAddress)
}

// SetStorageAddress is a paid mutator transaction binding the contract method 0x59b910d6.
//
// Solidity: function setStorageAddress(storageAddress address) returns()
func (_HoQuRater *HoQuRaterSession) SetStorageAddress(storageAddress common.Address) (*types.Transaction, error) {
	return _HoQuRater.Contract.SetStorageAddress(&_HoQuRater.TransactOpts, storageAddress)
}

// SetStorageAddress is a paid mutator transaction binding the contract method 0x59b910d6.
//
// Solidity: function setStorageAddress(storageAddress address) returns()
func (_HoQuRater *HoQuRaterTransactorSession) SetStorageAddress(storageAddress common.Address) (*types.Transaction, error) {
	return _HoQuRater.Contract.SetStorageAddress(&_HoQuRater.TransactOpts, storageAddress)
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
	Raw               types.Log // Blockchain specific contextual infos
}

// FilterUserAddressAdded is a free log retrieval operation binding the contract event 0xb8a06c596f027284f088ce1ba17842629ca63619a7303c94b3645e310ae34e44.
//
// Solidity: event UserAddressAdded(ownerAddress indexed address, additionalAddress address)
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

// WatchUserAddressAdded is a free log subscription operation binding the contract event 0xb8a06c596f027284f088ce1ba17842629ca63619a7303c94b3645e310ae34e44.
//
// Solidity: event UserAddressAdded(ownerAddress indexed address, additionalAddress address)
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

// HoQuStorageSchemaABI is the input ABI used to generate the binding from.
const HoQuStorageSchemaABI = "[]"

// HoQuStorageSchemaBin is the compiled bytecode used for deploying new contracts.
const HoQuStorageSchemaBin = `0x6080604052348015600f57600080fd5b50603580601d6000396000f3006080604052600080fd00a165627a7a72305820c8282cfff2cab821c02062bf28c10a97d6a88745315e2b1c90a6cb168fcc18610029`

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

// OwnableABI is the input ABI used to generate the binding from.
const OwnableABI = "[{\"constant\":false,\"inputs\":[],\"name\":\"renounceOwnership\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"transferOwnership\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"previousOwner\",\"type\":\"address\"}],\"name\":\"OwnershipRenounced\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"previousOwner\",\"type\":\"address\"},{\"indexed\":true,\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"OwnershipTransferred\",\"type\":\"event\"}]"

// OwnableBin is the compiled bytecode used for deploying new contracts.
const OwnableBin = `0x608060405234801561001057600080fd5b5060008054600160a060020a033316600160a060020a03199091161790556102078061003d6000396000f3006080604052600436106100565763ffffffff7c0100000000000000000000000000000000000000000000000000000000600035041663715018a6811461005b5780638da5cb5b14610072578063f2fde38b146100a3575b600080fd5b34801561006757600080fd5b506100706100c4565b005b34801561007e57600080fd5b50610087610134565b60408051600160a060020a039092168252519081900360200190f35b3480156100af57600080fd5b50610070600160a060020a0360043516610143565b60005433600160a060020a039081169116146100df57600080fd5b60008054604051600160a060020a03909116917ff8df31144d9c2f0f6b59d69b8b98abd5459d07f2742c4df920b25aae33c6482091a26000805473ffffffffffffffffffffffffffffffffffffffff19169055565b600054600160a060020a031681565b60005433600160a060020a0390811691161461015e57600080fd5b600160a060020a038116151561017357600080fd5b60008054604051600160a060020a03808516939216917f8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e091a36000805473ffffffffffffffffffffffffffffffffffffffff1916600160a060020a03929092169190911790555600a165627a7a723058201b7d2787c2c8497b6e40a1fccc152f9690c33ed956149f32f23f8e6cffe7b30b0029`

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
const SafeMathBin = `0x604c602c600b82828239805160001a60731460008114601c57601e565bfe5b5030600052607381538281f30073000000000000000000000000000000000000000030146080604052600080fd00a165627a7a723058204b0833ca02fe1aa398f2bd537624b0972c3af7a3a4e3fcd6ba952f7fa3886c310029`

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
