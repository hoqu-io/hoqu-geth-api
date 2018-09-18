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
