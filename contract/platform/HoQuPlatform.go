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

// HoQuPlatformABI is the input ABI used to generate the binding from.
const HoQuPlatformABI = "[{\"constant\":false,\"inputs\":[{\"name\":\"id\",\"type\":\"bytes16\"},{\"name\":\"status\",\"type\":\"uint8\"}],\"name\":\"setNetworkStatus\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"id\",\"type\":\"bytes16\"},{\"name\":\"ownerAddress\",\"type\":\"address\"}],\"name\":\"addUserAddress\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"id\",\"type\":\"bytes16\"},{\"name\":\"adCampaignId\",\"type\":\"bytes16\"}],\"name\":\"transactLead\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"id\",\"type\":\"bytes16\"},{\"name\":\"ownerId\",\"type\":\"bytes16\"},{\"name\":\"networkId\",\"type\":\"bytes16\"},{\"name\":\"merchantId\",\"type\":\"bytes16\"},{\"name\":\"payerAddress\",\"type\":\"address\"},{\"name\":\"name\",\"type\":\"string\"},{\"name\":\"dataUrl\",\"type\":\"string\"}],\"name\":\"addOffer\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"id\",\"type\":\"bytes16\"},{\"name\":\"userId\",\"type\":\"bytes16\"},{\"name\":\"idType\",\"type\":\"string\"},{\"name\":\"name\",\"type\":\"string\"},{\"name\":\"companyId\",\"type\":\"bytes16\"}],\"name\":\"addIdentification\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"id\",\"type\":\"bytes16\"},{\"name\":\"meta\",\"type\":\"string\"},{\"name\":\"kycLevel\",\"type\":\"uint8\"},{\"name\":\"dataUrl\",\"type\":\"string\"}],\"name\":\"addKycReport\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"id\",\"type\":\"bytes16\"},{\"name\":\"tariffGroupId\",\"type\":\"bytes16\"},{\"name\":\"name\",\"type\":\"string\"},{\"name\":\"action\",\"type\":\"string\"},{\"name\":\"calcMethod\",\"type\":\"string\"},{\"name\":\"price\",\"type\":\"uint256\"}],\"name\":\"addTariff\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"id\",\"type\":\"bytes16\"},{\"name\":\"payerAddress\",\"type\":\"address\"}],\"name\":\"setOfferPayerAddress\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"id\",\"type\":\"bytes16\"},{\"name\":\"status\",\"type\":\"uint8\"}],\"name\":\"setTariffGroupStatus\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"id\",\"type\":\"bytes16\"},{\"name\":\"status\",\"type\":\"uint8\"}],\"name\":\"setOfferStatus\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"id\",\"type\":\"bytes16\"},{\"name\":\"num\",\"type\":\"uint8\"},{\"name\":\"tariffGroupId\",\"type\":\"bytes16\"}],\"name\":\"setOfferTariffGroup\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"id\",\"type\":\"bytes16\"},{\"name\":\"name\",\"type\":\"string\"}],\"name\":\"setOfferName\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"id\",\"type\":\"bytes16\"},{\"name\":\"ownerId\",\"type\":\"bytes16\"},{\"name\":\"offerId\",\"type\":\"bytes16\"},{\"name\":\"contractAddress\",\"type\":\"address\"}],\"name\":\"addAdCampaign\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"id\",\"type\":\"bytes16\"},{\"name\":\"ownerId\",\"type\":\"bytes16\"},{\"name\":\"networkId\",\"type\":\"bytes16\"},{\"name\":\"name\",\"type\":\"string\"},{\"name\":\"dataUrl\",\"type\":\"string\"}],\"name\":\"registerTracker\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"storageAddress\",\"type\":\"address\"}],\"name\":\"setStorageAddress\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"id\",\"type\":\"bytes16\"},{\"name\":\"adCampaignId\",\"type\":\"bytes16\"},{\"name\":\"_status\",\"type\":\"uint8\"}],\"name\":\"setLeadStatus\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"id\",\"type\":\"bytes16\"},{\"name\":\"num\",\"type\":\"uint8\"}],\"name\":\"getOfferTariffGroup\",\"outputs\":[{\"name\":\"\",\"type\":\"bytes16\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"id\",\"type\":\"bytes16\"},{\"name\":\"ownerId\",\"type\":\"bytes16\"},{\"name\":\"name\",\"type\":\"string\"},{\"name\":\"dataUrl\",\"type\":\"string\"}],\"name\":\"registerNetwork\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"config\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"id\",\"type\":\"bytes16\"},{\"name\":\"adCampaignId\",\"type\":\"bytes16\"},{\"name\":\"_price\",\"type\":\"uint256\"}],\"name\":\"setLeadPrice\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"configAddress\",\"type\":\"address\"}],\"name\":\"setConfigAddress\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"id\",\"type\":\"bytes16\"},{\"name\":\"pubKey\",\"type\":\"string\"}],\"name\":\"updateUserPubKey\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"id\",\"type\":\"bytes16\"},{\"name\":\"adCampaignId\",\"type\":\"bytes16\"},{\"name\":\"trackerId\",\"type\":\"bytes16\"},{\"name\":\"meta\",\"type\":\"string\"},{\"name\":\"dataUrl\",\"type\":\"string\"},{\"name\":\"price\",\"type\":\"uint256\"}],\"name\":\"addLead\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"id\",\"type\":\"bytes16\"},{\"name\":\"ownerId\",\"type\":\"bytes16\"},{\"name\":\"name\",\"type\":\"string\"}],\"name\":\"addTariffGroup\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"store\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"id\",\"type\":\"bytes16\"},{\"name\":\"ownerId\",\"type\":\"bytes16\"},{\"name\":\"name\",\"type\":\"string\"},{\"name\":\"dataUrl\",\"type\":\"string\"}],\"name\":\"registerCompany\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"id\",\"type\":\"bytes16\"},{\"name\":\"adCampaignId\",\"type\":\"bytes16\"},{\"name\":\"_dataUrl\",\"type\":\"string\"}],\"name\":\"setLeadDataUrl\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"id\",\"type\":\"bytes16\"},{\"name\":\"adCampaignId\",\"type\":\"bytes16\"},{\"name\":\"intermediaryAddress\",\"type\":\"address\"},{\"name\":\"percent\",\"type\":\"uint32\"}],\"name\":\"addLeadIntermediary\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"id\",\"type\":\"bytes16\"},{\"name\":\"status\",\"type\":\"uint8\"}],\"name\":\"setCompanyStatus\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"id\",\"type\":\"bytes16\"},{\"name\":\"num\",\"type\":\"uint8\"}],\"name\":\"getUserAddress\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"id\",\"type\":\"bytes16\"},{\"name\":\"num\",\"type\":\"uint16\"}],\"name\":\"getKycReport\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"},{\"name\":\"\",\"type\":\"string\"},{\"name\":\"\",\"type\":\"uint8\"},{\"name\":\"\",\"type\":\"string\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"id\",\"type\":\"bytes16\"},{\"name\":\"status\",\"type\":\"uint8\"}],\"name\":\"setAdCampaignStatus\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"id\",\"type\":\"bytes16\"},{\"name\":\"status\",\"type\":\"uint8\"}],\"name\":\"setUserStatus\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"id\",\"type\":\"bytes16\"},{\"name\":\"status\",\"type\":\"uint8\"}],\"name\":\"setTariffStatus\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"id\",\"type\":\"bytes16\"},{\"name\":\"status\",\"type\":\"uint8\"}],\"name\":\"setTrackerStatus\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"id\",\"type\":\"bytes16\"},{\"name\":\"tariffGroupId\",\"type\":\"bytes16\"}],\"name\":\"addOfferTariffGroup\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"id\",\"type\":\"bytes16\"},{\"name\":\"role\",\"type\":\"string\"},{\"name\":\"ownerAddress\",\"type\":\"address\"},{\"name\":\"pubKey\",\"type\":\"string\"}],\"name\":\"registerUser\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"name\":\"configAddress\",\"type\":\"address\"},{\"name\":\"storageAddress\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"ownerAddress\",\"type\":\"address\"},{\"indexed\":false,\"name\":\"id\",\"type\":\"bytes16\"},{\"indexed\":false,\"name\":\"role\",\"type\":\"string\"}],\"name\":\"UserRegistered\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"ownerAddress\",\"type\":\"address\"},{\"indexed\":false,\"name\":\"additionalAddress\",\"type\":\"address\"},{\"indexed\":false,\"name\":\"id\",\"type\":\"bytes16\"}],\"name\":\"UserAddressAdded\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"ownerAddress\",\"type\":\"address\"},{\"indexed\":false,\"name\":\"id\",\"type\":\"bytes16\"}],\"name\":\"UserPubKeyUpdated\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"ownerAddress\",\"type\":\"address\"},{\"indexed\":false,\"name\":\"id\",\"type\":\"bytes16\"},{\"indexed\":false,\"name\":\"status\",\"type\":\"uint8\"}],\"name\":\"UserStatusChanged\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"ownerAddress\",\"type\":\"address\"},{\"indexed\":false,\"name\":\"id\",\"type\":\"bytes16\"},{\"indexed\":false,\"name\":\"userId\",\"type\":\"bytes16\"},{\"indexed\":false,\"name\":\"name\",\"type\":\"string\"}],\"name\":\"IdentificationAdded\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"ownerAddress\",\"type\":\"address\"},{\"indexed\":false,\"name\":\"kycLevel\",\"type\":\"uint8\"},{\"indexed\":false,\"name\":\"id\",\"type\":\"bytes16\"},{\"indexed\":false,\"name\":\"userId\",\"type\":\"bytes16\"}],\"name\":\"KycReportAdded\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"ownerAddress\",\"type\":\"address\"},{\"indexed\":false,\"name\":\"id\",\"type\":\"bytes16\"},{\"indexed\":false,\"name\":\"name\",\"type\":\"string\"}],\"name\":\"CompanyRegistered\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"ownerAddress\",\"type\":\"address\"},{\"indexed\":false,\"name\":\"id\",\"type\":\"bytes16\"},{\"indexed\":false,\"name\":\"status\",\"type\":\"uint8\"}],\"name\":\"CompanyStatusChanged\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"ownerAddress\",\"type\":\"address\"},{\"indexed\":false,\"name\":\"id\",\"type\":\"bytes16\"},{\"indexed\":false,\"name\":\"name\",\"type\":\"string\"}],\"name\":\"NetworkRegistered\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"ownerAddress\",\"type\":\"address\"},{\"indexed\":false,\"name\":\"id\",\"type\":\"bytes16\"},{\"indexed\":false,\"name\":\"status\",\"type\":\"uint8\"}],\"name\":\"NetworkStatusChanged\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"ownerAddress\",\"type\":\"address\"},{\"indexed\":false,\"name\":\"id\",\"type\":\"bytes16\"},{\"indexed\":false,\"name\":\"name\",\"type\":\"string\"}],\"name\":\"TrackerRegistered\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"ownerAddress\",\"type\":\"address\"},{\"indexed\":false,\"name\":\"id\",\"type\":\"bytes16\"},{\"indexed\":false,\"name\":\"status\",\"type\":\"uint8\"}],\"name\":\"TrackerStatusChanged\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"ownerAddress\",\"type\":\"address\"},{\"indexed\":false,\"name\":\"id\",\"type\":\"bytes16\"},{\"indexed\":false,\"name\":\"name\",\"type\":\"string\"}],\"name\":\"OfferAdded\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"ownerAddress\",\"type\":\"address\"},{\"indexed\":false,\"name\":\"id\",\"type\":\"bytes16\"},{\"indexed\":false,\"name\":\"tariff_id\",\"type\":\"bytes16\"}],\"name\":\"OfferTariffGroupAdded\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"ownerAddress\",\"type\":\"address\"},{\"indexed\":false,\"name\":\"id\",\"type\":\"bytes16\"},{\"indexed\":false,\"name\":\"status\",\"type\":\"uint8\"}],\"name\":\"OfferStatusChanged\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"ownerAddress\",\"type\":\"address\"},{\"indexed\":false,\"name\":\"id\",\"type\":\"bytes16\"},{\"indexed\":false,\"name\":\"contractAddress\",\"type\":\"address\"}],\"name\":\"AdCampaignAdded\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"ownerAddress\",\"type\":\"address\"},{\"indexed\":false,\"name\":\"id\",\"type\":\"bytes16\"},{\"indexed\":false,\"name\":\"status\",\"type\":\"uint8\"}],\"name\":\"AdCampaignStatusChanged\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"contractAddress\",\"type\":\"address\"},{\"indexed\":false,\"name\":\"adCampaignId\",\"type\":\"bytes16\"},{\"indexed\":false,\"name\":\"id\",\"type\":\"bytes16\"}],\"name\":\"LeadAdded\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"contractAddress\",\"type\":\"address\"},{\"indexed\":false,\"name\":\"adCampaignId\",\"type\":\"bytes16\"},{\"indexed\":false,\"name\":\"id\",\"type\":\"bytes16\"}],\"name\":\"LeadTransacted\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"contractAddress\",\"type\":\"address\"},{\"indexed\":false,\"name\":\"adCampaignId\",\"type\":\"bytes16\"},{\"indexed\":false,\"name\":\"id\",\"type\":\"bytes16\"},{\"indexed\":false,\"name\":\"status\",\"type\":\"uint8\"}],\"name\":\"LeadStatusChanged\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"contractAddress\",\"type\":\"address\"},{\"indexed\":false,\"name\":\"adCampaignId\",\"type\":\"bytes16\"},{\"indexed\":false,\"name\":\"id\",\"type\":\"bytes16\"},{\"indexed\":false,\"name\":\"price\",\"type\":\"uint256\"}],\"name\":\"LeadPriceChanged\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"contractAddress\",\"type\":\"address\"},{\"indexed\":false,\"name\":\"adCampaignId\",\"type\":\"bytes16\"},{\"indexed\":false,\"name\":\"id\",\"type\":\"bytes16\"},{\"indexed\":false,\"name\":\"dataUrl\",\"type\":\"string\"}],\"name\":\"LeadDataUrlChanged\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"ownerAddress\",\"type\":\"address\"},{\"indexed\":false,\"name\":\"id\",\"type\":\"bytes16\"},{\"indexed\":false,\"name\":\"name\",\"type\":\"string\"}],\"name\":\"TariffGroupAdded\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"ownerAddress\",\"type\":\"address\"},{\"indexed\":false,\"name\":\"id\",\"type\":\"bytes16\"},{\"indexed\":false,\"name\":\"status\",\"type\":\"uint8\"}],\"name\":\"TariffGroupStatusChanged\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"ownerAddress\",\"type\":\"address\"},{\"indexed\":false,\"name\":\"id\",\"type\":\"bytes16\"},{\"indexed\":false,\"name\":\"name\",\"type\":\"string\"}],\"name\":\"TariffAdded\",\"type\":\"event\"}]"

// HoQuPlatformBin is the compiled bytecode used for deploying new contracts.
const HoQuPlatformBin = `0x608060405234801561001057600080fd5b5060405160408062005fd383398101604052805160209091015160008054600160a060020a03938416600160a060020a03199182161790915560018054939092169216919091179055615f6a80620000696000396000f3006080604052600436106101cb5763ffffffff7c01000000000000000000000000000000000000000000000000000000006000350416627c9b5181146101d057806303806028146101fa57806305b8c83e146102285780630c189e10146102505780630c2ec3851461031757806318c8f305146103d15780631a16f4471461048057806322e1d7631461056e578063310eea121461059c57806336d7523f146105c45780633e45850a146105ec5780634af4c3fd1461061b57806351a4d8a6146106835780635392a3f2146106bd57806359b910d61461076e578063689708541461078f57806374f627c5146107bd5780637618c8101461080257806379502c55146108b057806379b4330b146108e157806383a12de91461090c5780638a5463bd1461092d5780638f2e960d146109955780639228d91714610a48578063975057e714610ab85780639f5dec0614610acd578063a6ccf89d14610b7b578063a779dcac14610beb578063a8800f0914610c28578063c49e2fb114610c50578063cc3436b914610c78578063d7915dfb14610d9b578063e03d8b6414610dc3578063e1ef893514610deb578063e7e5aad814610e13578063f17eceb114610e3b578063fd3fc08f14610e63575b600080fd5b3480156101dc57600080fd5b506101f86001608060020a03196004351660ff60243516610f18565b005b34801561020657600080fd5b506101f86001608060020a031960043516600160a060020a0360243516611079565b34801561023457600080fd5b506101f86001608060020a03196004358116906024351661116e565b34801561025c57600080fd5b50604080516020600460a43581810135601f81018490048402850184019095528484526101f89482356001608060020a0319908116956024803583169660443584169660643590941695608435600160a060020a03169536959460c4949093920191819084018382808284375050604080516020601f89358b018035918201839004830284018301909452808352979a9998810197919650918201945092508291508401838280828437509497506112da9650505050505050565b34801561032357600080fd5b50604080516020600460443581810135601f81018490048402850184019095528484526101f89482356001608060020a03199081169560248035909216953695946064949293019190819084018382808284375050604080516020601f89358b018035918201839004830284018301909452808352979a999881019791965091820194509250829150840183828082843750949750505092356001608060020a03191693506114ac92505050565b3480156103dd57600080fd5b5060408051602060046024803582810135601f81018590048502860185019096528585526101f89583356001608060020a03191695369560449491939091019190819084018382808284375050604080516020601f818a01358b0180359182018390048302840183018552818452989b60ff8b35169b909a9099940197509195509182019350915081908401838280828437509497506116809650505050505050565b34801561048c57600080fd5b50604080516020600460443581810135601f81018490048402850184019095528484526101f89482356001608060020a03199081169560248035909216953695946064949293019190819084018382808284375050604080516020601f89358b018035918201839004830284018301909452808352979a99988101979196509182019450925082915084018382808284375050604080516020601f89358b018035918201839004830284018301909452808352979a99988101979196509182019450925082915084018382808284375094975050933594506118259350505050565b34801561057a57600080fd5b506101f86001608060020a031960043516600160a060020a03602435166119d0565b3480156105a857600080fd5b506101f86001608060020a03196004351660ff60243516611ab6565b3480156105d057600080fd5b506101f86001608060020a03196004351660ff60243516611bfc565b3480156105f857600080fd5b506101f86001608060020a031960043581169060ff602435169060443516611d45565b34801561062757600080fd5b5060408051602060046024803582810135601f81018590048502860185019096528585526101f89583356001608060020a031916953695604494919390910191908190840183828082843750949750611e509650505050505050565b34801561068f57600080fd5b506101f86001608060020a031960043581169060243581169060443516600160a060020a0360643516611f28565b3480156106c957600080fd5b50604080516020601f6064356004818101359283018490048402850184019095528184526101f8946001608060020a0319813581169560248035831696604435909316953695608494920191819084018382808284375050604080516020601f89358b018035918201839004830284018301909452808352979a9998810197919650918201945092508291508401838280828437509497506120709650505050505050565b34801561077a57600080fd5b506101f8600160a060020a0360043516612228565b34801561079b57600080fd5b506101f86001608060020a03196004358116906024351660ff604435166122df565b3480156107c957600080fd5b506107e56001608060020a03196004351660ff60243516612481565b604080516001608060020a03199092168252519081900360200190f35b34801561080e57600080fd5b50604080516020600460443581810135601f81018490048402850184019095528484526101f89482356001608060020a03199081169560248035909216953695946064949293019190819084018382808284375050604080516020601f89358b018035918201839004830284018301909452808352979a99988101979196509182019450925082915084018382808284375094975061252c9650505050505050565b3480156108bc57600080fd5b506108c56126db565b60408051600160a060020a039092168252519081900360200190f35b3480156108ed57600080fd5b506101f86001608060020a0319600435811690602435166044356126ea565b34801561091857600080fd5b506101f8600160a060020a0360043516612864565b34801561093957600080fd5b5060408051602060046024803582810135601f81018590048502860185019096528585526101f89583356001608060020a03191695369560449491939091019190819084018382808284375094975061291b9650505050505050565b3480156109a157600080fd5b50604080516020601f6064356004818101359283018490048402850184019095528184526101f8946001608060020a0319813581169560248035831696604435909316953695608494920191819084018382808284375050604080516020601f89358b018035918201839004830284018301909452808352979a9998810197919650918201945092508291508401838280828437509497505093359450612a259350505050565b348015610a5457600080fd5b50604080516020600460443581810135601f81018490048402850184019095528484526101f89482356001608060020a031990811695602480359092169536959460649492930191908190840183828082843750949750612c7e9650505050505050565b348015610ac457600080fd5b506108c5612e25565b348015610ad957600080fd5b50604080516020600460443581810135601f81018490048402850184019095528484526101f89482356001608060020a03199081169560248035909216953695946064949293019190819084018382808284375050604080516020601f89358b018035918201839004830284018301909452808352979a999881019791965091820194509250829150840183828082843750949750612e349650505050505050565b348015610b8757600080fd5b50604080516020600460443581810135601f81018490048402850184019095528484526101f89482356001608060020a031990811695602480359092169536959460649492930191908190840183828082843750949750612f9f9650505050505050565b348015610bf757600080fd5b506101f86001608060020a031960043581169060243516600160a060020a036044351663ffffffff60643516613200565b348015610c3457600080fd5b506101f86001608060020a03196004351660ff60243516613332565b348015610c5c57600080fd5b506108c56001608060020a03196004351660ff60243516613479565b348015610c8457600080fd5b50610ca16001608060020a03196004351661ffff602435166134f1565b6040518085815260200180602001846005811115610cbb57fe5b60ff16815260200180602001838103835286818151815260200191508051906020019080838360005b83811015610cfc578181015183820152602001610ce4565b50505050905090810190601f168015610d295780820380516001836020036101000a031916815260200191505b50838103825284518152845160209182019186019080838360005b83811015610d5c578181015183820152602001610d44565b50505050905090810190601f168015610d895780820380516001836020036101000a031916815260200191505b50965050505050505060405180910390f35b348015610da757600080fd5b506101f86001608060020a03196004351660ff60243516613664565b348015610dcf57600080fd5b506101f86001608060020a03196004351660ff602435166137ab565b348015610df757600080fd5b506101f86001608060020a03196004351660ff602435166138d6565b348015610e1f57600080fd5b506101f86001608060020a03196004351660ff602435166139bd565b348015610e4757600080fd5b506101f86001608060020a031960043581169060243516613b04565b348015610e6f57600080fd5b5060408051602060046024803582810135601f81018590048502860185019096528585526101f89583356001608060020a03191695369560449491939091019190819084018382808284375050604080516020601f818a01358b0180359182018390048302840183018552818452989b600160a060020a038b35169b909a909994019750919550918201935091508190840183828082843750949750613c0d9650505050505050565b610f20615d38565b600080546040805160e060020a63babcc5390281523360048201529051600160a060020a039092169163babcc5399160248082019260209290919082900301818787803b158015610f7057600080fd5b505af1158015610f84573d6000803e3d6000fd5b505050506040513d6020811015610f9a57600080fd5b50511515610fa757600080fd5b610fb084613d6b565b9150600082608001516005811115610fc457fe5b1415610fcf57600080fd5b8151610fdc906000613479565b905060808201836005811115610fee57fe5b90816005811115610ffb57fe5b9052506110088483613f1c565b80600160a060020a03167fce8cdf87291e737b80351c93570f8555ca6a0cc77c3fb29176e56af69dccdd84858560405180836001608060020a0319166001608060020a031916815260200182600581111561105f57fe5b60ff1681526020019250505060405180910390a250505050565b600080546040805160e060020a63babcc5390281523360048201529051600160a060020a039092169163babcc5399160248082019260209290919082900301818787803b1580156110c957600080fd5b505af11580156110dd573d6000803e3d6000fd5b505050506040513d60208110156110f357600080fd5b5051151561110057600080fd5b61110a83836140b0565b611115836000613479565b60408051600160a060020a0385811682526001608060020a03198716602083015282519394508416927fb38b406c043d4a3feaa0192d2e1b02f317fa73c284eb3e70b9a1fde8612af1b3929181900390910190a2505050565b600080546040805160e060020a63babcc5390281523360048201529051600160a060020a039092169163babcc5399160248082019260209290919082900301818787803b1580156111be57600080fd5b505af11580156111d2573d6000803e3d6000fd5b505050506040513d60208110156111e857600080fd5b505115156111f557600080fd5b6111fe82614127565b604080517f835aef400000000000000000000000000000000000000000000000000000000081526001608060020a0319861660048201529051919250600160a060020a0383169163835aef409160248082019260009290919082900301818387803b15801561126c57600080fd5b505af1158015611280573d6000803e3d6000fd5b5050604080516001608060020a03198087168252871660208201528151600160a060020a03861694507f3d86ba643ff86fc12dc0ce1d161b9407a79cb7131f80001f21c054658de5e1e893509081900390910190a2505050565b6112e2615d68565b600080546040805160e060020a63babcc5390281523360048201529051600160a060020a039092169163babcc5399160248082019260209290919082900301818787803b15801561133257600080fd5b505af1158015611346573d6000803e3d6000fd5b505050506040513d602081101561135c57600080fd5b5051151561136957600080fd5b61137289614163565b91506000826101000151600581111561138757fe5b1461139157600080fd5b61139c886000613479565b6001608060020a0319808a168452888116602085015287166040840152600160a060020a03861660608401526080830185905260a0830184905290506113e2898361449b565b80600160a060020a03167f5c9100ad970db426d43d47b3a9639dc42d9732871d99213dcb4984f7e71509d08a8660405180836001608060020a0319166001608060020a031916815260200180602001828103825283818151815260200191508051906020019080838360005b8381101561146657818101518382015260200161144e565b50505050905090810190601f1680156114935780820380516001836020036101000a031916815260200191505b50935050505060405180910390a2505050505050505050565b6114b4615db3565b600080546040805160e060020a63babcc5390281523360048201529051600160a060020a039092169163babcc5399160248082019260209290919082900301818787803b15801561150457600080fd5b505af1158015611518573d6000803e3d6000fd5b505050506040513d602081101561152e57600080fd5b5051151561153b57600080fd5b61154487614641565b915060008260c00151600581111561155857fe5b1461156257600080fd5b61156d866000613479565b6001608060020a03198088168452604084018790526060840186905284166020840152905061159c8783614813565b80600160a060020a03167f8c03773782e9ed7fe9111421165a6cf12ce5855a4181cd1a8166d47d719ac07888888760405180846001608060020a0319166001608060020a0319168152602001836001608060020a0319166001608060020a031916815260200180602001828103825283818151815260200191508051906020019080838360005b8381101561163b578181015183820152602001611623565b50505050905090810190601f1680156116685780820380516001836020036101000a031916815260200191505b5094505050505060405180910390a250505050505050565b611688615db3565b6000611692615df1565b61169a615e1a565b600080546040805160e060020a63babcc5390281523360048201529051600160a060020a039092169263babcc539926024808401936020939083900390910190829087803b1580156116eb57600080fd5b505af11580156116ff573d6000803e3d6000fd5b505050506040513d602081101561171557600080fd5b5051151561172257600080fd5b61172b88614641565b935061173c84600001516000613479565b87835292506020820186600581111561175157fe5b9081600581111561175e57fe5b905250604082018590526117728883614999565b61177b88614b00565b90506060810186600581111561178d57fe5b9081600581111561179a57fe5b90525083516117a99082614ce9565b82600160a060020a03167f09d7204a29a3e39bb12fb2c813212dd90eb69edaf26f867689c3e3a2eb6eb7b9878a8760000151604051808460058111156117eb57fe5b60ff1681526001608060020a0319938416602082015291909216604080830191909152519081900360600192509050a25050505050505050565b61182d615e5d565b611835615e8c565b600080546040805160e060020a63babcc5390281523360048201529051600160a060020a039092169163babcc5399160248082019260209290919082900301818787803b15801561188557600080fd5b505af1158015611899573d6000803e3d6000fd5b505050506040513d60208110156118af57600080fd5b505115156118bc57600080fd5b6118c588614e25565b92506000836080015160058111156118d957fe5b14156118e457600080fd5b6118ed89615088565b915060008260c00151600581111561190157fe5b1461190b57600080fd5b8251611918906000613479565b6001608060020a03198916835260208301889052604083018790526060830186905260808301859052905061194d8983615402565b80600160a060020a03167f4ad20f56c3c96698f2e84d660e8c5e39341fd79929743804d261beb65ffc5b478a8960405180836001608060020a0319166001608060020a031916815260200180602001828103825283818151815260200191508051906020019080838360008381101561146657818101518382015260200161144e565b6119d8615d68565b600080546040805160e060020a63babcc5390281523360048201529051600160a060020a039092169163babcc5399160248082019260209290919082900301818787803b158015611a2857600080fd5b505af1158015611a3c573d6000803e3d6000fd5b505050506040513d6020811015611a5257600080fd5b50511515611a5f57600080fd5b611a6884614163565b915060008261010001516005811115611a7d57fe5b1415611a8857600080fd5b8151611a95906000613479565b600160a060020a03841660608401529050611ab0848361449b565b50505050565b611abe615e5d565b600080546040805160e060020a63babcc5390281523360048201529051600160a060020a039092169163babcc5399160248082019260209290919082900301818787803b158015611b0e57600080fd5b505af1158015611b22573d6000803e3d6000fd5b505050506040513d6020811015611b3857600080fd5b50511515611b4557600080fd5b611b4e84614e25565b9150600082608001516005811115611b6257fe5b1415611b6d57600080fd5b8151611b7a906000613479565b905060808201836005811115611b8c57fe5b90816005811115611b9957fe5b905250611ba684836155a0565b80600160a060020a03167eb686e81dae287e0f688bd472beb9aa4c8b029ef31da062379d0d46713dfde6858560405180836001608060020a0319166001608060020a031916815260200182600581111561105f57fe5b611c04615d68565b600080546040805160e060020a63babcc5390281523360048201529051600160a060020a039092169163babcc5399160248082019260209290919082900301818787803b158015611c5457600080fd5b505af1158015611c68573d6000803e3d6000fd5b505050506040513d6020811015611c7e57600080fd5b50511515611c8b57600080fd5b611c9484614163565b915060008261010001516005811115611ca957fe5b1415611cb457600080fd5b8151611cc1906000613479565b90506101008201836005811115611cd457fe5b90816005811115611ce157fe5b905250611cee848361449b565b80600160a060020a03167ff9ae5f2bc88e6348608d627c5226159efd8467eb4ac3c635cdb648266e02cb12858560405180836001608060020a0319166001608060020a031916815260200182600581111561105f57fe5b611d4d615d68565b600080546040805160e060020a63babcc5390281523360048201529051600160a060020a039092169163babcc5399160248082019260209290919082900301818787803b158015611d9d57600080fd5b505af1158015611db1573d6000803e3d6000fd5b505050506040513d6020811015611dc757600080fd5b50511515611dd457600080fd5b611ddd85614163565b9150611dee82600001516000613479565b9050611dfb8585856156ab565b604080516001608060020a03198088168252851660208201528151600160a060020a038416927f695164079ea85cb7933398c55f763b3dd70fe1cfe293a3b90aa22f9cf6270dab928290030190a25050505050565b611e58615d68565b600080546040805160e060020a63babcc5390281523360048201529051600160a060020a039092169163babcc5399160248082019260209290919082900301818787803b158015611ea857600080fd5b505af1158015611ebc573d6000803e3d6000fd5b505050506040513d6020811015611ed257600080fd5b50511515611edf57600080fd5b611ee884614163565b915060008261010001516005811115611efd57fe5b1415611f0857600080fd5b8151611f15906000613479565b608083018490529050611ab0848361449b565b611f30615ed9565b600080546040805160e060020a63babcc5390281523360048201529051600160a060020a039092169163babcc5399160248082019260209290919082900301818787803b158015611f8057600080fd5b505af1158015611f94573d6000803e3d6000fd5b505050506040513d6020811015611faa57600080fd5b50511515611fb757600080fd5b611fc086615748565b9150600082608001516005811115611fd457fe5b14611fde57600080fd5b611fe9856000613479565b6001608060020a0319808716845285166020840152600160a060020a03841660408401529050612019868361586d565b604080516001608060020a031988168152600160a060020a0385811660208301528251908416927f62a609593cded926aabb6aef1e5b884804ad29d1bd1116ac4ce32085690d5ab2928290030190a2505050505050565b612078615f07565b600080546040805160e060020a63babcc5390281523360048201529051600160a060020a039092169163babcc5399160248082019260209290919082900301818787803b1580156120c857600080fd5b505af11580156120dc573d6000803e3d6000fd5b505050506040513d60208110156120f257600080fd5b505115156120ff57600080fd5b61210887615926565b915060008260a00151600581111561211c57fe5b1461212657600080fd5b612131866000613479565b6001608060020a0319808816845286166020840152604083018590526060830184905290506121608783615ae4565b80600160a060020a03167fb2f67f32df451ce918abd944fd5c8a29589848693057913c5060a8e0f293b388888660405180836001608060020a0319166001608060020a031916815260200180602001828103825283818151815260200191508051906020019080838360005b838110156121e45781810151838201526020016121cc565b50505050905090810190601f1680156122115780820380516001836020036101000a031916815260200191505b50935050505060405180910390a250505050505050565b600080546040805160e060020a63babcc5390281523360048201529051600160a060020a039092169263babcc539926024808401936020939083900390910190829087803b15801561227957600080fd5b505af115801561228d573d6000803e3d6000fd5b505050506040513d60208110156122a357600080fd5b505115156122b057600080fd5b6001805473ffffffffffffffffffffffffffffffffffffffff1916600160a060020a0392909216919091179055565b600080546040805160e060020a63babcc5390281523360048201529051600160a060020a039092169163babcc5399160248082019260209290919082900301818787803b15801561232f57600080fd5b505af1158015612343573d6000803e3d6000fd5b505050506040513d602081101561235957600080fd5b5051151561236657600080fd5b61236f83614127565b6040517f81e85aa30000000000000000000000000000000000000000000000000000000081526001608060020a0319861660048201908152919250600160a060020a038316916381e85aa391879186916024018260058111156123ce57fe5b60ff16815260200192505050600060405180830381600087803b1580156123f457600080fd5b505af1158015612408573d6000803e3d6000fd5b5050604080516001608060020a0319808816825288166020820152600160a060020a03851693507fcde6c64834d04acd2555be68ee4f9a2b92fa8c30fc842f42f5ed581214c8dfff9250869188918791810182600581111561246657fe5b60ff168152602001935050505060405180910390a250505050565b600154604080517f74f627c50000000000000000000000000000000000000000000000000000000081526001608060020a03198516600482015260ff841660248201529051600092600160a060020a0316916374f627c591604480830192602092919082900301818787803b1580156124f957600080fd5b505af115801561250d573d6000803e3d6000fd5b505050506040513d602081101561252357600080fd5b50519392505050565b612534615d38565b600080546040805160e060020a63babcc5390281523360048201529051600160a060020a039092169163babcc5399160248082019260209290919082900301818787803b15801561258457600080fd5b505af1158015612598573d6000803e3d6000fd5b505050506040513d60208110156125ae57600080fd5b505115156125bb57600080fd5b6125c486613d6b565b91506000826080015160058111156125d857fe5b146125e257600080fd5b6125ed856000613479565b6001608060020a031986168352602083018590526040830184905290506126148683613f1c565b80600160a060020a03167f592e83ede8c846dffd6e4edcbe77983b153bec820ac697aa34a236147b1341e0878660405180836001608060020a0319166001608060020a031916815260200180602001828103825283818151815260200191508051906020019080838360005b83811015612698578181015183820152602001612680565b50505050905090810190601f1680156126c55780820380516001836020036101000a031916815260200191505b50935050505060405180910390a2505050505050565b600054600160a060020a031681565b600080546040805160e060020a63babcc5390281523360048201529051600160a060020a039092169163babcc5399160248082019260209290919082900301818787803b15801561273a57600080fd5b505af115801561274e573d6000803e3d6000fd5b505050506040513d602081101561276457600080fd5b5051151561277157600080fd5b61277a83614127565b604080517fc535bd880000000000000000000000000000000000000000000000000000000081526001608060020a031987166004820152602481018590529051919250600160a060020a0383169163c535bd889160448082019260009290919082900301818387803b1580156127ef57600080fd5b505af1158015612803573d6000803e3d6000fd5b5050604080516001608060020a03198088168252881660208201528082018690529051600160a060020a03851693507f10d48255e0ae3847d38a5e163a9d9afb2964ac9c892d76cea6572eca867daa3992509081900360600190a250505050565b600080546040805160e060020a63babcc5390281523360048201529051600160a060020a039092169263babcc539926024808401936020939083900390910190829087803b1580156128b557600080fd5b505af11580156128c9573d6000803e3d6000fd5b505050506040513d60208110156128df57600080fd5b505115156128ec57600080fd5b6000805473ffffffffffffffffffffffffffffffffffffffff1916600160a060020a0392909216919091179055565b6000612925615e1a565b600080546040805160e060020a63babcc5390281523360048201529051600160a060020a039092169263babcc539926024808401936020939083900390910190829087803b15801561297657600080fd5b505af115801561298a573d6000803e3d6000fd5b505050506040513d60208110156129a057600080fd5b505115156129ad57600080fd5b6129b8846000613479565b91506129c384614b00565b6080810184905290506129d68482614ce9565b604080516001608060020a0319861681529051600160a060020a038416917f85e1546b8c89c1a56c4237d2608010366813f87b475163aadf4074e5bc29339b919081900360200190a250505050565b600080546040805160e060020a63babcc5390281523360048201529051600160a060020a039092169163babcc5399160248082019260209290919082900301818787803b158015612a7557600080fd5b505af1158015612a89573d6000803e3d6000fd5b505050506040513d6020811015612a9f57600080fd5b50511515612aac57600080fd5b612ab586614127565b6040517f7a8187cf0000000000000000000000000000000000000000000000000000000081526001608060020a0319808a166004830190815290881660248301526084820185905260a060448301908152875160a48401528751939450600160a060020a03851693637a8187cf938c938b938b938b938b939291606482019160c40190602088019080838360005b83811015612b5b578181015183820152602001612b43565b50505050905090810190601f168015612b885780820380516001836020036101000a031916815260200191505b50838103825285518152855160209182019187019080838360005b83811015612bbb578181015183820152602001612ba3565b50505050905090810190601f168015612be85780820380516001836020036101000a031916815260200191505b50975050505050505050600060405180830381600087803b158015612c0c57600080fd5b505af1158015612c20573d6000803e3d6000fd5b5050604080516001608060020a0319808b1682528b1660208201528151600160a060020a03861694507f4d5402ce2635084c7def60402936ba18370a5a5d055ecdddc10721afa736ce4293509081900390910190a250505050505050565b612c86615e5d565b600080546040805160e060020a63babcc5390281523360048201529051600160a060020a039092169163babcc5399160248082019260209290919082900301818787803b158015612cd657600080fd5b505af1158015612cea573d6000803e3d6000fd5b505050506040513d6020811015612d0057600080fd5b50511515612d0d57600080fd5b612d1685614e25565b9150600082608001516005811115612d2a57fe5b14612d3457600080fd5b612d3f846000613479565b6001608060020a031985168352602083018490529050612d5f85836155a0565b80600160a060020a03167f0d945b4e88c22df3f6d39d8866fac2dd4964090d9d9c04fc74f33c5695e561df868560405180836001608060020a0319166001608060020a031916815260200180602001828103825283818151815260200191508051906020019080838360005b83811015612de3578181015183820152602001612dcb565b50505050905090810190601f168015612e105780820380516001836020036101000a031916815260200191505b50935050505060405180910390a25050505050565b600154600160a060020a031681565b612e3c615d38565b600080546040805160e060020a63babcc5390281523360048201529051600160a060020a039092169163babcc5399160248082019260209290919082900301818787803b158015612e8c57600080fd5b505af1158015612ea0573d6000803e3d6000fd5b505050506040513d6020811015612eb657600080fd5b50511515612ec357600080fd5b612ecc86615bb8565b9150600082608001516005811115612ee057fe5b14612eea57600080fd5b612ef5856000613479565b6001608060020a03198616835260208301859052604083018490529050612f1c8683615c37565b80600160a060020a03167f0eb23d662c57754d23afdc1cffa6246fbae3bbe843eec5c42a5a5223c21414ed878660405180836001608060020a0319166001608060020a0319168152602001806020018281038252838181518152602001915080519060200190808383600083811015612698578181015183820152602001612680565b600080546040805160e060020a63babcc5390281523360048201529051600160a060020a039092169163babcc5399160248082019260209290919082900301818787803b158015612fef57600080fd5b505af1158015613003573d6000803e3d6000fd5b505050506040513d602081101561301957600080fd5b5051151561302657600080fd5b61302f83614127565b604080517f406865fc0000000000000000000000000000000000000000000000000000000081526001608060020a031987166004820190815260248201928352855160448301528551939450600160a060020a0385169363406865fc938993889392606490910190602085019080838360005b838110156130ba5781810151838201526020016130a2565b50505050905090810190601f1680156130e75780820380516001836020036101000a031916815260200191505b509350505050600060405180830381600087803b15801561310757600080fd5b505af115801561311b573d6000803e3d6000fd5b5050505080600160a060020a03167f030cd1b359a7520d9c7419a712a01660bf8f8a953e84d3c917f8f87a68214a4184868560405180846001608060020a0319166001608060020a0319168152602001836001608060020a0319166001608060020a031916815260200180602001828103825283818151815260200191508051906020019080838360005b838110156131be5781810151838201526020016131a6565b50505050905090810190601f1680156131eb5780820380516001836020036101000a031916815260200191505b5094505050505060405180910390a250505050565b600080546040805160e060020a63babcc5390281523360048201529051600160a060020a039092169163babcc5399160248082019260209290919082900301818787803b15801561325057600080fd5b505af1158015613264573d6000803e3d6000fd5b505050506040513d602081101561327a57600080fd5b5051151561328757600080fd5b61329084614127565b604080517f366170450000000000000000000000000000000000000000000000000000000081526001608060020a031988166004820152600160a060020a03868116602483015263ffffffff8616604483015291519293509083169163366170459160648082019260009290919082900301818387803b15801561331357600080fd5b505af1158015613327573d6000803e3d6000fd5b505050505050505050565b61333a615d38565b600080546040805160e060020a63babcc5390281523360048201529051600160a060020a039092169163babcc5399160248082019260209290919082900301818787803b15801561338a57600080fd5b505af115801561339e573d6000803e3d6000fd5b505050506040513d60208110156133b457600080fd5b505115156133c157600080fd5b6133ca84615bb8565b91506000826080015160058111156133de57fe5b14156133e957600080fd5b81516133f6906000613479565b90506080820183600581111561340857fe5b9081600581111561341557fe5b9052506134228483615c37565b80600160a060020a03167f14404ec3e30e0b13c4ef42e76c8fd3afe300c9bf51761965c78d348ff45cf2f0858560405180836001608060020a0319166001608060020a031916815260200182600581111561105f57fe5b600154604080517fc49e2fb10000000000000000000000000000000000000000000000000000000081526001608060020a03198516600482015260ff841660248201529051600092600160a060020a03169163c49e2fb191604480830192602092919082900301818787803b1580156124f957600080fd5b600154604080517fcc3436b90000000000000000000000000000000000000000000000000000000081526001608060020a03198516600482015261ffff84166024820152905160009260609284928492600160a060020a03169163cc3436b9916044808301928792919082900301818387803b15801561357057600080fd5b505af1158015613584573d6000803e3d6000fd5b505050506040513d6000823e601f3d908101601f1916820160405260808110156135ad57600080fd5b8151602083018051919392830192916401000000008111156135ce57600080fd5b820160208101848111156135e157600080fd5b81516401000000008111828201871017156135fb57600080fd5b5050602082015160409092018051919492939164010000000081111561362057600080fd5b8201602081018481111561363357600080fd5b815164010000000081118282018710171561364d57600080fd5b50969d959c50939a50929850929650505050505050565b61366c615ed9565b600080546040805160e060020a63babcc5390281523360048201529051600160a060020a039092169163babcc5399160248082019260209290919082900301818787803b1580156136bc57600080fd5b505af11580156136d0573d6000803e3d6000fd5b505050506040513d60208110156136e657600080fd5b505115156136f357600080fd5b6136fc84615748565b915060008260800151600581111561371057fe5b141561371b57600080fd5b8151613728906000613479565b90506080820183600581111561373a57fe5b9081600581111561374757fe5b905250613754848361586d565b80600160a060020a03167fa25922380013ceba1e6f7c02184e30818a465ac4d7107ac4357b5883887b5af4858560405180836001608060020a0319166001608060020a031916815260200182600581111561105f57fe5b60006137b5615e1a565b600080546040805160e060020a63babcc5390281523360048201529051600160a060020a039092169263babcc539926024808401936020939083900390910190829087803b15801561380657600080fd5b505af115801561381a573d6000803e3d6000fd5b505050506040513d602081101561383057600080fd5b5051151561383d57600080fd5b613848846000613479565b915061385384614b00565b905060c0810183600581111561386557fe5b9081600581111561387257fe5b90525061387f8482614ce9565b81600160a060020a03167f183408f9d1ce45bacc33cd1a5ef30bf3caccbae9e28fcc8916b62185fdd5ac4d858560405180836001608060020a0319166001608060020a031916815260200182600581111561105f57fe5b6138de615e8c565b600080546040805160e060020a63babcc5390281523360048201529051600160a060020a039092169263babcc539926024808401936020939083900390910190829087803b15801561392f57600080fd5b505af1158015613943573d6000803e3d6000fd5b505050506040513d602081101561395957600080fd5b5051151561396657600080fd5b61396f83615088565b905060008160c00151600581111561398357fe5b141561398e57600080fd5b60c0810182600581111561399e57fe5b908160058111156139ab57fe5b9052506139b88382615402565b505050565b6139c5615f07565b600080546040805160e060020a63babcc5390281523360048201529051600160a060020a039092169163babcc5399160248082019260209290919082900301818787803b158015613a1557600080fd5b505af1158015613a29573d6000803e3d6000fd5b505050506040513d6020811015613a3f57600080fd5b50511515613a4c57600080fd5b613a5584615926565b915060008260a001516005811115613a6957fe5b1415613a7457600080fd5b8151613a81906000613479565b905060a08201836005811115613a9357fe5b90816005811115613aa057fe5b905250613aad8483615ae4565b80600160a060020a03167f77a747805fe662172d50df0565e43aabadea2efb1198dc3ff8e60fcd57eb7489858560405180836001608060020a0319166001608060020a031916815260200182600581111561105f57fe5b613b0c615d68565b600080546040805160e060020a63babcc5390281523360048201529051600160a060020a039092169163babcc5399160248082019260209290919082900301818787803b158015613b5c57600080fd5b505af1158015613b70573d6000803e3d6000fd5b505050506040513d6020811015613b8657600080fd5b50511515613b9357600080fd5b613b9c84614163565b9150613bad82600001516000613479565b9050613bb98484615cc1565b604080516001608060020a03198087168252851660208201528151600160a060020a038416927f695164079ea85cb7933398c55f763b3dd70fe1cfe293a3b90aa22f9cf6270dab928290030190a250505050565b613c15615e1a565b600080546040805160e060020a63babcc5390281523360048201529051600160a060020a039092169263babcc539926024808401936020939083900390910190829087803b158015613c6657600080fd5b505af1158015613c7a573d6000803e3d6000fd5b505050506040513d6020811015613c9057600080fd5b50511515613c9d57600080fd5b613ca685614b00565b905060008160c001516005811115613cba57fe5b14613cc457600080fd5b600160a060020a03831681526040810184905260808101829052613ce88582614ce9565b82600160a060020a03167ff661ee1472892faaf2bb68cc6874f2759c9148b4234cac6b88f6ca362652f275868660405180836001608060020a0319166001608060020a0319168152602001806020018281038252838181518152602001915080519060200190808383600083811015612de3578181015183820152602001612dcb565b613d73615d38565b613d7b615d38565b600154604080517f72c130e40000000000000000000000000000000000000000000000000000000081526001608060020a0319861660048201529051600160a060020a03909216916372c130e49160248082019260009290919082900301818387803b158015613dea57600080fd5b505af1158015613dfe573d6000803e3d6000fd5b505050506040513d6000823e601f3d908101601f1916820160405260a0811015613e2757600080fd5b815160208301805191939283019291640100000000811115613e4857600080fd5b82016020810184811115613e5b57600080fd5b8151640100000000811182820187101715613e7557600080fd5b50509291906020018051640100000000811115613e9157600080fd5b82016020810184811115613ea457600080fd5b8151640100000000811182820187101715613ebe57600080fd5b505060208083015160409384015192955093509091879182019082016060830160808401856005811115613eee57fe5b6005811115613ef957fe5b90529490945293909252509190526001608060020a031990911690529050919050565b60015481516020830151604080850151608086015191517f028849800000000000000000000000000000000000000000000000000000000081526001608060020a0319808916600483019081529086166024830152600160a060020a03909616956302884980958995909490939290919060448101906064810190608401846005811115613fa657fe5b60ff168152602001838103835286818151815260200191508051906020019080838360005b83811015613fe3578181015183820152602001613fcb565b50505050905090810190601f1680156140105780820380516001836020036101000a031916815260200191505b50838103825285518152855160209182019187019080838360005b8381101561404357818101518382015260200161402b565b50505050905090810190601f1680156140705780820380516001836020036101000a031916815260200191505b50975050505050505050600060405180830381600087803b15801561409457600080fd5b505af11580156140a8573d6000803e3d6000fd5b505050505050565b600154604080517f038060280000000000000000000000000000000000000000000000000000000081526001608060020a031985166004820152600160a060020a03848116602483015291519190921691630380602891604480830192600092919082900301818387803b15801561409457600080fd5b6000614131615ed9565b61413a83615748565b905060008160800151600581111561414e57fe5b141561415957600080fd5b6040015192915050565b61416b615d68565b614173615d68565b600154604080517fc32d869b0000000000000000000000000000000000000000000000000000000081526001608060020a0319861660048201529051600160a060020a039092169163c32d869b9160248082019260009290919082900301818387803b1580156141e257600080fd5b505af11580156141f6573d6000803e3d6000fd5b505050506040513d6000823e601f3d908101601f1916820160405261012081101561422057600080fd5b81516020830151604084015160608501516080860180519496939592949193928301929164010000000081111561425657600080fd5b8201602081018481111561426957600080fd5b815164010000000081118282018710171561428357600080fd5b5050929190602001805164010000000081111561429f57600080fd5b820160208101848111156142b257600080fd5b81516401000000008111828201871017156142cc57600080fd5b505060a08a01525050506080860152600160a060020a0390811660608601526001608060020a03199182166040868101919091529282166020860152928116845260015482517fc32d869b00000000000000000000000000000000000000000000000000000000815291871660048301529151919092169163c32d869b91602480830192600092919082900301818387803b15801561436a57600080fd5b505af115801561437e573d6000803e3d6000fd5b505050506040513d6000823e601f3d908101601f191682016040526101208110156143a857600080fd5b8151602083015160408401516060850151608086018051949693959294919392830192916401000000008111156143de57600080fd5b820160208101848111156143f157600080fd5b815164010000000081118282018710171561440b57600080fd5b5050929190602001805164010000000081111561442757600080fd5b8201602081018481111561443a57600080fd5b815164010000000081118282018710171561445457600080fd5b50505060408101516060909101519097509550505060e086019250506101008501905082600581111561448357fe5b600581111561448e57fe5b9052919091525092915050565b600154815160208301516040808501516060860151608087015160a088015161010089015194517f1071d5120000000000000000000000000000000000000000000000000000000081526001608060020a0319808c1660048301908152818a16602484015281891660448401529086166064830152600160a060020a03858116608484015290991698631071d512988c98909790969594939290919060a481019060c481019060e40184600581111561455057fe5b60ff168152602001838103835286818151815260200191508051906020019080838360005b8381101561458d578181015183820152602001614575565b50505050905090810190601f1680156145ba5780820380516001836020036101000a031916815260200191505b50838103825285518152855160209182019187019080838360005b838110156145ed5781810151838201526020016145d5565b50505050905090810190601f16801561461a5780820380516001836020036101000a031916815260200191505b509a5050505050505050505050600060405180830381600087803b15801561409457600080fd5b614649615db3565b614651615db3565b600154604080517f289e74700000000000000000000000000000000000000000000000000000000081526001608060020a0319861660048201529051600160a060020a039092169163289e74709160248082019260009290919082900301818387803b1580156146c057600080fd5b505af11580156146d4573d6000803e3d6000fd5b505050506040513d6000823e601f3d908101601f1916820160405260e08110156146fd57600080fd5b815160208301516040840180519294919382019264010000000081111561472357600080fd5b8201602081018481111561473657600080fd5b815164010000000081118282018710171561475057600080fd5b5050929190602001805164010000000081111561476c57600080fd5b8201602081018481111561477f57600080fd5b815164010000000081118282018710171561479957600080fd5b505060208083015160408085015160609586015194975091955090938a9283019183019083016080840160a0850160c086018760058111156147d757fe5b60058111156147e257fe5b90529690965261ffff9096169094529490935250929092526001608060020a03199283169052911690529050919050565b60015481516040808401516060850151602086015160c087015193517facc02cce0000000000000000000000000000000000000000000000000000000081526001608060020a0319808a166004830190815281881660248401529083166084830152600160a060020a039097169663acc02cce968a9690959493929091906044810190606481019060a4018460058111156148aa57fe5b60ff168152602001838103835287818151815260200191508051906020019080838360005b838110156148e75781810151838201526020016148cf565b50505050905090810190601f1680156149145780820380516001836020036101000a031916815260200191505b50838103825286518152865160209182019188019080838360005b8381101561494757818101518382015260200161492f565b50505050905090810190601f1680156149745780820380516001836020036101000a031916815260200191505b5098505050505050505050600060405180830381600087803b15801561409457600080fd5b6001548151602083015160408085015190517f18c8f3050000000000000000000000000000000000000000000000000000000081526001608060020a0319871660048201908152600160a060020a03909516946318c8f305948894909390929091906024810190604401846005811115614a0f57fe5b60ff16815260200180602001838103835286818151815260200191508051906020019080838360005b83811015614a50578181015183820152602001614a38565b50505050905090810190601f168015614a7d5780820380516001836020036101000a031916815260200191505b50838103825284518152845160209182019186019080838360005b83811015614ab0578181015183820152602001614a98565b50505050905090810190601f168015614add5780820380516001836020036101000a031916815260200191505b509650505050505050600060405180830381600087803b15801561409457600080fd5b614b08615e1a565b614b10615e1a565b600154604080517f7520dd140000000000000000000000000000000000000000000000000000000081526001608060020a0319861660048201529051600160a060020a0390921691637520dd149160248082019260009290919082900301818387803b158015614b7f57600080fd5b505af1158015614b93573d6000803e3d6000fd5b505050506040513d6000823e601f3d908101601f1916820160405260e0811015614bbc57600080fd5b8151602083015160408401805192949193820192640100000000811115614be257600080fd5b82016020810184811115614bf557600080fd5b8151640100000000811182820187101715614c0f57600080fd5b50506020820151604090920180519194929391640100000000811115614c3457600080fd5b82016020810184811115614c4757600080fd5b8151640100000000811182820187101715614c6157600080fd5b50506020808301516040938401519295509350909189918201908201606083016080840160a0850160c08601876005811115614c9957fe5b6005811115614ca457fe5b9052879052879052876005811115614cb857fe5b6005811115614cc357fe5b90529690965260ff909616909452505050600160a060020a039092169052509050919050565b60015460408083015183516060850151608086015160c087015194517f253a9b9e0000000000000000000000000000000000000000000000000000000081526001608060020a0319891660048201908152600160a060020a0385811660448401529097169663253a9b9e968a96959493929091906024810190606401856005811115614d7157fe5b60ff16815260200180602001846005811115614d8957fe5b60ff168152602001838103835288818151815260200191508051906020019080838360005b83811015614dc6578181015183820152602001614dae565b50505050905090810190601f168015614df35780820380516001836020036101000a031916815260200191505b50838103825285518152855160209182019187019080838360008381101561494757818101518382015260200161492f565b614e2d615e5d565b614e35615e5d565b600154604080517f8ade00270000000000000000000000000000000000000000000000000000000081526001608060020a0319861660048201529051600160a060020a0390921691638ade00279160248082019260009290919082900301818387803b158015614ea457600080fd5b505af1158015614eb8573d6000803e3d6000fd5b505050506040513d6000823e601f3d908101601f1916820160405260a0811015614ee157600080fd5b815160208301805191939283019291640100000000811115614f0257600080fd5b82016020810184811115614f1557600080fd5b8151640100000000811182820187101715614f2f57600080fd5b505060208601525050506001608060020a03199081168252600154604080517f8ade0027000000000000000000000000000000000000000000000000000000008152928616600484015251600160a060020a0390911691638ade002791602480830192600092919082900301818387803b158015614fac57600080fd5b505af1158015614fc0573d6000803e3d6000fd5b505050506040513d6000823e601f3d908101601f1916820160405260a0811015614fe957600080fd5b81516020830180519193928301929164010000000081111561500a57600080fd5b8201602081018481111561501d57600080fd5b815164010000000081118282018710171561503757600080fd5b5050506020810151604080830151606093840151929650945090925085019085016080860183600581111561506857fe5b600581111561507357fe5b9052929092525060ff90911690529050919050565b615090615e8c565b615098615e8c565b600154604080517fa2cbda410000000000000000000000000000000000000000000000000000000081526001608060020a0319861660048201529051600160a060020a039092169163a2cbda419160248082019260009290919082900301818387803b15801561510757600080fd5b505af115801561511b573d6000803e3d6000fd5b505050506040513d6000823e601f3d908101601f1916820160405260e081101561514457600080fd5b81516020830180519193928301929164010000000081111561516557600080fd5b8201602081018481111561517857600080fd5b815164010000000081118282018710171561519257600080fd5b505092919060200180516401000000008111156151ae57600080fd5b820160208101848111156151c157600080fd5b81516401000000008111828201871017156151db57600080fd5b505092919060200180516401000000008111156151f757600080fd5b8201602081018481111561520a57600080fd5b815164010000000081118282018710171561522457600080fd5b505060209182015160808901526060880152604087810194909452860193909352506001608060020a0319928316845260015481517fa2cbda4100000000000000000000000000000000000000000000000000000000815293871660048501529051600160a060020a039091169263a2cbda41925060248082019260009290919082900301818387803b1580156152ba57600080fd5b505af11580156152ce573d6000803e3d6000fd5b505050506040513d6000823e601f3d908101601f1916820160405260e08110156152f757600080fd5b81516020830180519193928301929164010000000081111561531857600080fd5b8201602081018481111561532b57600080fd5b815164010000000081118282018710171561534557600080fd5b5050929190602001805164010000000081111561536157600080fd5b8201602081018481111561537457600080fd5b815164010000000081118282018710171561538e57600080fd5b505092919060200180516401000000008111156153aa57600080fd5b820160208101848111156153bd57600080fd5b81516401000000008111828201871017156153d757600080fd5b505050604081015160609091015190955093505060a0850191505060c0840182600581111561448357fe5b600154815160208301516040808501516060860151608087015160c088015193517f61fad57e0000000000000000000000000000000000000000000000000000000081526001608060020a0319808b1660048301908152908816602483015260a48201839052600160a060020a03909816976361fad57e978b979096909594939290919060448101906064810190608481019060c4018560058111156154a457fe5b60ff168152602001848103845289818151815260200191508051906020019080838360005b838110156154e15781810151838201526020016154c9565b50505050905090810190601f16801561550e5780820380516001836020036101000a031916815260200191505b5084810383528851815288516020918201918a019080838360005b83811015615541578181015183820152602001615529565b50505050905090810190601f16801561556e5780820380516001836020036101000a031916815260200191505b5084810382528751815287516020918201918901908083836000838110156145ed5781810151838201526020016145d5565b6001548151602083015160808401516040517fa7fbcf610000000000000000000000000000000000000000000000000000000081526001608060020a0319808816600483019081529085166024830152600160a060020a039095169463a7fbcf6194889490939092909190604481019060640183600581111561561f57fe5b60ff168152602001828103825284818151815260200191508051906020019080838360005b8381101561565c578181015183820152602001615644565b50505050905090810190601f1680156156895780820380516001836020036101000a031916815260200191505b5095505050505050600060405180830381600087803b15801561409457600080fd5b600154604080517f3e45850a0000000000000000000000000000000000000000000000000000000081526001608060020a0319808716600483015260ff86166024830152841660448201529051600160a060020a0390921691633e45850a9160648082019260009290919082900301818387803b15801561572b57600080fd5b505af115801561573f573d6000803e3d6000fd5b50505050505050565b615750615ed9565b615758615ed9565b600154604080517f890755570000000000000000000000000000000000000000000000000000000081526001608060020a0319861660048201529051600160a060020a039092169163890755579160248082019260a0929091908290030181600087803b1580156157c857600080fd5b505af11580156157dc573d6000803e3d6000fd5b505050506040513d60a08110156157f257600080fd5b508051602080830151604080850151606080870151608097880151969794969295909492938993918401929084019190840190840185600581111561583357fe5b600581111561583e57fe5b905294909452600160a060020a03909416909252506001608060020a0319928316909152911690529050919050565b60015481516020830151604080850151608086015191517fa0191d200000000000000000000000000000000000000000000000000000000081526001608060020a03198089166004830190815281871660248401529085166044830152600160a060020a0383811660648401529096169563a0191d2095899590949093929091906084018260058111156158fd57fe5b60ff16815260200195505050505050600060405180830381600087803b15801561409457600080fd5b61592e615f07565b615936615f07565b600154604080517fea2fdd8c0000000000000000000000000000000000000000000000000000000081526001608060020a0319861660048201529051600160a060020a039092169163ea2fdd8c9160248082019260009290919082900301818387803b1580156159a557600080fd5b505af11580156159b9573d6000803e3d6000fd5b505050506040513d6000823e601f3d908101601f1916820160405260c08110156159e257600080fd5b8151602083015160408401805192949193820192640100000000811115615a0857600080fd5b82016020810184811115615a1b57600080fd5b8151640100000000811182820187101715615a3557600080fd5b50509291906020018051640100000000811115615a5157600080fd5b82016020810184811115615a6457600080fd5b8151640100000000811182820187101715615a7e57600080fd5b50506020808301516040938401519295509350909188918201908201606083016080840160a08501866005811115615ab257fe5b6005811115615abd57fe5b90529590955294909352939092526001608060020a03199384169052509116905292915050565b60015481516020830151604080850151606086015160a087015192517fc4d4d8870000000000000000000000000000000000000000000000000000000081526001608060020a0319808a166004830190815281881660248401529086166044830152600160a060020a039097169663c4d4d887968a969095909493929091906064810190608481019060a401846005811115615b7c57fe5b60ff1681526020018381038352868181518152602001915080519060200190808383600083811015614dc6578181015183820152602001614dae565b615bc0615d38565b615bc8615d38565b600154604080517f4df2d8470000000000000000000000000000000000000000000000000000000081526001608060020a0319861660048201529051600160a060020a0390921691634df2d8479160248082019260009290919082900301818387803b158015613dea57600080fd5b60015481516020830151604080850151608086015191517f768d80a80000000000000000000000000000000000000000000000000000000081526001608060020a0319808916600483019081529086166024830152600160a060020a039096169563768d80a8958995909490939290919060448101906064810190608401846005811115613fa657fe5b600154604080517ff17eceb10000000000000000000000000000000000000000000000000000000081526001608060020a03198086166004830152841660248201529051600160a060020a039092169163f17eceb19160448082019260009290919082900301818387803b15801561409457600080fd5b6040805160a08101825260008082526060602083018190529282018390529181018290529060808201905b905290565b604080516101208101825260008082526020820181905291810182905260608082018390526080820181905260a082015260c0810182905260e0810182905290610100820190615d63565b6040805160e0810182526000808252602082018190526060928201839052828201929092526080810182905260a081018290529060c0820190615d63565b604080516080810190915260608152602081016000815260200160608152602001600081525090565b6040805160e081018252600080825260208201819052606092820183905290918201908152602001606081526020016000815260200160006005811115615d6357fe5b6040805160a0810182526000808252606060208301819052928201819052918101829052906080820190615d63565b60e06040519081016040528060006001608060020a0319168152602001606081526020016060815260200160608152602001600081526020016000815260200160006005811115615d6357fe5b6040805160a08101825260008082526020820181905291810182905260608101829052906080820190615d63565b6040805160c081018252600080825260208201819052606092820183905282820192909252608081018290529060a0820190615d635600a165627a7a723058205261672ef6f11de32057d0cb5d888bc02db4c10963239ba473c258500d1866fe0029`

// DeployHoQuPlatform deploys a new Ethereum contract, binding an instance of HoQuPlatform to it.
func DeployHoQuPlatform(auth *bind.TransactOpts, backend bind.ContractBackend, configAddress common.Address, storageAddress common.Address) (common.Address, *types.Transaction, *HoQuPlatform, error) {
	parsed, err := abi.JSON(strings.NewReader(HoQuPlatformABI))
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	address, tx, contract, err := bind.DeployContract(auth, parsed, common.FromHex(HoQuPlatformBin), backend, configAddress, storageAddress)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &HoQuPlatform{HoQuPlatformCaller: HoQuPlatformCaller{contract: contract}, HoQuPlatformTransactor: HoQuPlatformTransactor{contract: contract}, HoQuPlatformFilterer: HoQuPlatformFilterer{contract: contract}}, nil
}

// HoQuPlatform is an auto generated Go binding around an Ethereum contract.
type HoQuPlatform struct {
	HoQuPlatformCaller     // Read-only binding to the contract
	HoQuPlatformTransactor // Write-only binding to the contract
	HoQuPlatformFilterer   // Log filterer for contract events
}

// HoQuPlatformCaller is an auto generated read-only Go binding around an Ethereum contract.
type HoQuPlatformCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// HoQuPlatformTransactor is an auto generated write-only Go binding around an Ethereum contract.
type HoQuPlatformTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// HoQuPlatformFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type HoQuPlatformFilterer struct {
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
	contract, err := bindHoQuPlatform(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &HoQuPlatform{HoQuPlatformCaller: HoQuPlatformCaller{contract: contract}, HoQuPlatformTransactor: HoQuPlatformTransactor{contract: contract}, HoQuPlatformFilterer: HoQuPlatformFilterer{contract: contract}}, nil
}

// NewHoQuPlatformCaller creates a new read-only instance of HoQuPlatform, bound to a specific deployed contract.
func NewHoQuPlatformCaller(address common.Address, caller bind.ContractCaller) (*HoQuPlatformCaller, error) {
	contract, err := bindHoQuPlatform(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &HoQuPlatformCaller{contract: contract}, nil
}

// NewHoQuPlatformTransactor creates a new write-only instance of HoQuPlatform, bound to a specific deployed contract.
func NewHoQuPlatformTransactor(address common.Address, transactor bind.ContractTransactor) (*HoQuPlatformTransactor, error) {
	contract, err := bindHoQuPlatform(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &HoQuPlatformTransactor{contract: contract}, nil
}

// NewHoQuPlatformFilterer creates a new log filterer instance of HoQuPlatform, bound to a specific deployed contract.
func NewHoQuPlatformFilterer(address common.Address, filterer bind.ContractFilterer) (*HoQuPlatformFilterer, error) {
	contract, err := bindHoQuPlatform(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &HoQuPlatformFilterer{contract: contract}, nil
}

// bindHoQuPlatform binds a generic wrapper to an already deployed contract.
func bindHoQuPlatform(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(HoQuPlatformABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
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

// GetKycReport is a free data retrieval call binding the contract method 0xcc3436b9.
//
// Solidity: function getKycReport(id bytes16, num uint16) constant returns(uint256, string, uint8, string)
func (_HoQuPlatform *HoQuPlatformCaller) GetKycReport(opts *bind.CallOpts, id [16]byte, num uint16) (*big.Int, string, uint8, string, error) {
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
	err := _HoQuPlatform.contract.Call(opts, out, "getKycReport", id, num)
	return *ret0, *ret1, *ret2, *ret3, err
}

// GetKycReport is a free data retrieval call binding the contract method 0xcc3436b9.
//
// Solidity: function getKycReport(id bytes16, num uint16) constant returns(uint256, string, uint8, string)
func (_HoQuPlatform *HoQuPlatformSession) GetKycReport(id [16]byte, num uint16) (*big.Int, string, uint8, string, error) {
	return _HoQuPlatform.Contract.GetKycReport(&_HoQuPlatform.CallOpts, id, num)
}

// GetKycReport is a free data retrieval call binding the contract method 0xcc3436b9.
//
// Solidity: function getKycReport(id bytes16, num uint16) constant returns(uint256, string, uint8, string)
func (_HoQuPlatform *HoQuPlatformCallerSession) GetKycReport(id [16]byte, num uint16) (*big.Int, string, uint8, string, error) {
	return _HoQuPlatform.Contract.GetKycReport(&_HoQuPlatform.CallOpts, id, num)
}

// GetOfferTariffGroup is a free data retrieval call binding the contract method 0x74f627c5.
//
// Solidity: function getOfferTariffGroup(id bytes16, num uint8) constant returns(bytes16)
func (_HoQuPlatform *HoQuPlatformCaller) GetOfferTariffGroup(opts *bind.CallOpts, id [16]byte, num uint8) ([16]byte, error) {
	var (
		ret0 = new([16]byte)
	)
	out := ret0
	err := _HoQuPlatform.contract.Call(opts, out, "getOfferTariffGroup", id, num)
	return *ret0, err
}

// GetOfferTariffGroup is a free data retrieval call binding the contract method 0x74f627c5.
//
// Solidity: function getOfferTariffGroup(id bytes16, num uint8) constant returns(bytes16)
func (_HoQuPlatform *HoQuPlatformSession) GetOfferTariffGroup(id [16]byte, num uint8) ([16]byte, error) {
	return _HoQuPlatform.Contract.GetOfferTariffGroup(&_HoQuPlatform.CallOpts, id, num)
}

// GetOfferTariffGroup is a free data retrieval call binding the contract method 0x74f627c5.
//
// Solidity: function getOfferTariffGroup(id bytes16, num uint8) constant returns(bytes16)
func (_HoQuPlatform *HoQuPlatformCallerSession) GetOfferTariffGroup(id [16]byte, num uint8) ([16]byte, error) {
	return _HoQuPlatform.Contract.GetOfferTariffGroup(&_HoQuPlatform.CallOpts, id, num)
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

// Store is a free data retrieval call binding the contract method 0x975057e7.
//
// Solidity: function store() constant returns(address)
func (_HoQuPlatform *HoQuPlatformCaller) Store(opts *bind.CallOpts) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _HoQuPlatform.contract.Call(opts, out, "store")
	return *ret0, err
}

// Store is a free data retrieval call binding the contract method 0x975057e7.
//
// Solidity: function store() constant returns(address)
func (_HoQuPlatform *HoQuPlatformSession) Store() (common.Address, error) {
	return _HoQuPlatform.Contract.Store(&_HoQuPlatform.CallOpts)
}

// Store is a free data retrieval call binding the contract method 0x975057e7.
//
// Solidity: function store() constant returns(address)
func (_HoQuPlatform *HoQuPlatformCallerSession) Store() (common.Address, error) {
	return _HoQuPlatform.Contract.Store(&_HoQuPlatform.CallOpts)
}

// AddAdCampaign is a paid mutator transaction binding the contract method 0x51a4d8a6.
//
// Solidity: function addAdCampaign(id bytes16, ownerId bytes16, offerId bytes16, contractAddress address) returns()
func (_HoQuPlatform *HoQuPlatformTransactor) AddAdCampaign(opts *bind.TransactOpts, id [16]byte, ownerId [16]byte, offerId [16]byte, contractAddress common.Address) (*types.Transaction, error) {
	return _HoQuPlatform.contract.Transact(opts, "addAdCampaign", id, ownerId, offerId, contractAddress)
}

// AddAdCampaign is a paid mutator transaction binding the contract method 0x51a4d8a6.
//
// Solidity: function addAdCampaign(id bytes16, ownerId bytes16, offerId bytes16, contractAddress address) returns()
func (_HoQuPlatform *HoQuPlatformSession) AddAdCampaign(id [16]byte, ownerId [16]byte, offerId [16]byte, contractAddress common.Address) (*types.Transaction, error) {
	return _HoQuPlatform.Contract.AddAdCampaign(&_HoQuPlatform.TransactOpts, id, ownerId, offerId, contractAddress)
}

// AddAdCampaign is a paid mutator transaction binding the contract method 0x51a4d8a6.
//
// Solidity: function addAdCampaign(id bytes16, ownerId bytes16, offerId bytes16, contractAddress address) returns()
func (_HoQuPlatform *HoQuPlatformTransactorSession) AddAdCampaign(id [16]byte, ownerId [16]byte, offerId [16]byte, contractAddress common.Address) (*types.Transaction, error) {
	return _HoQuPlatform.Contract.AddAdCampaign(&_HoQuPlatform.TransactOpts, id, ownerId, offerId, contractAddress)
}

// AddIdentification is a paid mutator transaction binding the contract method 0x0c2ec385.
//
// Solidity: function addIdentification(id bytes16, userId bytes16, idType string, name string, companyId bytes16) returns()
func (_HoQuPlatform *HoQuPlatformTransactor) AddIdentification(opts *bind.TransactOpts, id [16]byte, userId [16]byte, idType string, name string, companyId [16]byte) (*types.Transaction, error) {
	return _HoQuPlatform.contract.Transact(opts, "addIdentification", id, userId, idType, name, companyId)
}

// AddIdentification is a paid mutator transaction binding the contract method 0x0c2ec385.
//
// Solidity: function addIdentification(id bytes16, userId bytes16, idType string, name string, companyId bytes16) returns()
func (_HoQuPlatform *HoQuPlatformSession) AddIdentification(id [16]byte, userId [16]byte, idType string, name string, companyId [16]byte) (*types.Transaction, error) {
	return _HoQuPlatform.Contract.AddIdentification(&_HoQuPlatform.TransactOpts, id, userId, idType, name, companyId)
}

// AddIdentification is a paid mutator transaction binding the contract method 0x0c2ec385.
//
// Solidity: function addIdentification(id bytes16, userId bytes16, idType string, name string, companyId bytes16) returns()
func (_HoQuPlatform *HoQuPlatformTransactorSession) AddIdentification(id [16]byte, userId [16]byte, idType string, name string, companyId [16]byte) (*types.Transaction, error) {
	return _HoQuPlatform.Contract.AddIdentification(&_HoQuPlatform.TransactOpts, id, userId, idType, name, companyId)
}

// AddKycReport is a paid mutator transaction binding the contract method 0x18c8f305.
//
// Solidity: function addKycReport(id bytes16, meta string, kycLevel uint8, dataUrl string) returns()
func (_HoQuPlatform *HoQuPlatformTransactor) AddKycReport(opts *bind.TransactOpts, id [16]byte, meta string, kycLevel uint8, dataUrl string) (*types.Transaction, error) {
	return _HoQuPlatform.contract.Transact(opts, "addKycReport", id, meta, kycLevel, dataUrl)
}

// AddKycReport is a paid mutator transaction binding the contract method 0x18c8f305.
//
// Solidity: function addKycReport(id bytes16, meta string, kycLevel uint8, dataUrl string) returns()
func (_HoQuPlatform *HoQuPlatformSession) AddKycReport(id [16]byte, meta string, kycLevel uint8, dataUrl string) (*types.Transaction, error) {
	return _HoQuPlatform.Contract.AddKycReport(&_HoQuPlatform.TransactOpts, id, meta, kycLevel, dataUrl)
}

// AddKycReport is a paid mutator transaction binding the contract method 0x18c8f305.
//
// Solidity: function addKycReport(id bytes16, meta string, kycLevel uint8, dataUrl string) returns()
func (_HoQuPlatform *HoQuPlatformTransactorSession) AddKycReport(id [16]byte, meta string, kycLevel uint8, dataUrl string) (*types.Transaction, error) {
	return _HoQuPlatform.Contract.AddKycReport(&_HoQuPlatform.TransactOpts, id, meta, kycLevel, dataUrl)
}

// AddLead is a paid mutator transaction binding the contract method 0x8f2e960d.
//
// Solidity: function addLead(id bytes16, adCampaignId bytes16, trackerId bytes16, meta string, dataUrl string, price uint256) returns()
func (_HoQuPlatform *HoQuPlatformTransactor) AddLead(opts *bind.TransactOpts, id [16]byte, adCampaignId [16]byte, trackerId [16]byte, meta string, dataUrl string, price *big.Int) (*types.Transaction, error) {
	return _HoQuPlatform.contract.Transact(opts, "addLead", id, adCampaignId, trackerId, meta, dataUrl, price)
}

// AddLead is a paid mutator transaction binding the contract method 0x8f2e960d.
//
// Solidity: function addLead(id bytes16, adCampaignId bytes16, trackerId bytes16, meta string, dataUrl string, price uint256) returns()
func (_HoQuPlatform *HoQuPlatformSession) AddLead(id [16]byte, adCampaignId [16]byte, trackerId [16]byte, meta string, dataUrl string, price *big.Int) (*types.Transaction, error) {
	return _HoQuPlatform.Contract.AddLead(&_HoQuPlatform.TransactOpts, id, adCampaignId, trackerId, meta, dataUrl, price)
}

// AddLead is a paid mutator transaction binding the contract method 0x8f2e960d.
//
// Solidity: function addLead(id bytes16, adCampaignId bytes16, trackerId bytes16, meta string, dataUrl string, price uint256) returns()
func (_HoQuPlatform *HoQuPlatformTransactorSession) AddLead(id [16]byte, adCampaignId [16]byte, trackerId [16]byte, meta string, dataUrl string, price *big.Int) (*types.Transaction, error) {
	return _HoQuPlatform.Contract.AddLead(&_HoQuPlatform.TransactOpts, id, adCampaignId, trackerId, meta, dataUrl, price)
}

// AddLeadIntermediary is a paid mutator transaction binding the contract method 0xa779dcac.
//
// Solidity: function addLeadIntermediary(id bytes16, adCampaignId bytes16, intermediaryAddress address, percent uint32) returns()
func (_HoQuPlatform *HoQuPlatformTransactor) AddLeadIntermediary(opts *bind.TransactOpts, id [16]byte, adCampaignId [16]byte, intermediaryAddress common.Address, percent uint32) (*types.Transaction, error) {
	return _HoQuPlatform.contract.Transact(opts, "addLeadIntermediary", id, adCampaignId, intermediaryAddress, percent)
}

// AddLeadIntermediary is a paid mutator transaction binding the contract method 0xa779dcac.
//
// Solidity: function addLeadIntermediary(id bytes16, adCampaignId bytes16, intermediaryAddress address, percent uint32) returns()
func (_HoQuPlatform *HoQuPlatformSession) AddLeadIntermediary(id [16]byte, adCampaignId [16]byte, intermediaryAddress common.Address, percent uint32) (*types.Transaction, error) {
	return _HoQuPlatform.Contract.AddLeadIntermediary(&_HoQuPlatform.TransactOpts, id, adCampaignId, intermediaryAddress, percent)
}

// AddLeadIntermediary is a paid mutator transaction binding the contract method 0xa779dcac.
//
// Solidity: function addLeadIntermediary(id bytes16, adCampaignId bytes16, intermediaryAddress address, percent uint32) returns()
func (_HoQuPlatform *HoQuPlatformTransactorSession) AddLeadIntermediary(id [16]byte, adCampaignId [16]byte, intermediaryAddress common.Address, percent uint32) (*types.Transaction, error) {
	return _HoQuPlatform.Contract.AddLeadIntermediary(&_HoQuPlatform.TransactOpts, id, adCampaignId, intermediaryAddress, percent)
}

// AddOffer is a paid mutator transaction binding the contract method 0x0c189e10.
//
// Solidity: function addOffer(id bytes16, ownerId bytes16, networkId bytes16, merchantId bytes16, payerAddress address, name string, dataUrl string) returns()
func (_HoQuPlatform *HoQuPlatformTransactor) AddOffer(opts *bind.TransactOpts, id [16]byte, ownerId [16]byte, networkId [16]byte, merchantId [16]byte, payerAddress common.Address, name string, dataUrl string) (*types.Transaction, error) {
	return _HoQuPlatform.contract.Transact(opts, "addOffer", id, ownerId, networkId, merchantId, payerAddress, name, dataUrl)
}

// AddOffer is a paid mutator transaction binding the contract method 0x0c189e10.
//
// Solidity: function addOffer(id bytes16, ownerId bytes16, networkId bytes16, merchantId bytes16, payerAddress address, name string, dataUrl string) returns()
func (_HoQuPlatform *HoQuPlatformSession) AddOffer(id [16]byte, ownerId [16]byte, networkId [16]byte, merchantId [16]byte, payerAddress common.Address, name string, dataUrl string) (*types.Transaction, error) {
	return _HoQuPlatform.Contract.AddOffer(&_HoQuPlatform.TransactOpts, id, ownerId, networkId, merchantId, payerAddress, name, dataUrl)
}

// AddOffer is a paid mutator transaction binding the contract method 0x0c189e10.
//
// Solidity: function addOffer(id bytes16, ownerId bytes16, networkId bytes16, merchantId bytes16, payerAddress address, name string, dataUrl string) returns()
func (_HoQuPlatform *HoQuPlatformTransactorSession) AddOffer(id [16]byte, ownerId [16]byte, networkId [16]byte, merchantId [16]byte, payerAddress common.Address, name string, dataUrl string) (*types.Transaction, error) {
	return _HoQuPlatform.Contract.AddOffer(&_HoQuPlatform.TransactOpts, id, ownerId, networkId, merchantId, payerAddress, name, dataUrl)
}

// AddOfferTariffGroup is a paid mutator transaction binding the contract method 0xf17eceb1.
//
// Solidity: function addOfferTariffGroup(id bytes16, tariffGroupId bytes16) returns()
func (_HoQuPlatform *HoQuPlatformTransactor) AddOfferTariffGroup(opts *bind.TransactOpts, id [16]byte, tariffGroupId [16]byte) (*types.Transaction, error) {
	return _HoQuPlatform.contract.Transact(opts, "addOfferTariffGroup", id, tariffGroupId)
}

// AddOfferTariffGroup is a paid mutator transaction binding the contract method 0xf17eceb1.
//
// Solidity: function addOfferTariffGroup(id bytes16, tariffGroupId bytes16) returns()
func (_HoQuPlatform *HoQuPlatformSession) AddOfferTariffGroup(id [16]byte, tariffGroupId [16]byte) (*types.Transaction, error) {
	return _HoQuPlatform.Contract.AddOfferTariffGroup(&_HoQuPlatform.TransactOpts, id, tariffGroupId)
}

// AddOfferTariffGroup is a paid mutator transaction binding the contract method 0xf17eceb1.
//
// Solidity: function addOfferTariffGroup(id bytes16, tariffGroupId bytes16) returns()
func (_HoQuPlatform *HoQuPlatformTransactorSession) AddOfferTariffGroup(id [16]byte, tariffGroupId [16]byte) (*types.Transaction, error) {
	return _HoQuPlatform.Contract.AddOfferTariffGroup(&_HoQuPlatform.TransactOpts, id, tariffGroupId)
}

// AddTariff is a paid mutator transaction binding the contract method 0x1a16f447.
//
// Solidity: function addTariff(id bytes16, tariffGroupId bytes16, name string, action string, calcMethod string, price uint256) returns()
func (_HoQuPlatform *HoQuPlatformTransactor) AddTariff(opts *bind.TransactOpts, id [16]byte, tariffGroupId [16]byte, name string, action string, calcMethod string, price *big.Int) (*types.Transaction, error) {
	return _HoQuPlatform.contract.Transact(opts, "addTariff", id, tariffGroupId, name, action, calcMethod, price)
}

// AddTariff is a paid mutator transaction binding the contract method 0x1a16f447.
//
// Solidity: function addTariff(id bytes16, tariffGroupId bytes16, name string, action string, calcMethod string, price uint256) returns()
func (_HoQuPlatform *HoQuPlatformSession) AddTariff(id [16]byte, tariffGroupId [16]byte, name string, action string, calcMethod string, price *big.Int) (*types.Transaction, error) {
	return _HoQuPlatform.Contract.AddTariff(&_HoQuPlatform.TransactOpts, id, tariffGroupId, name, action, calcMethod, price)
}

// AddTariff is a paid mutator transaction binding the contract method 0x1a16f447.
//
// Solidity: function addTariff(id bytes16, tariffGroupId bytes16, name string, action string, calcMethod string, price uint256) returns()
func (_HoQuPlatform *HoQuPlatformTransactorSession) AddTariff(id [16]byte, tariffGroupId [16]byte, name string, action string, calcMethod string, price *big.Int) (*types.Transaction, error) {
	return _HoQuPlatform.Contract.AddTariff(&_HoQuPlatform.TransactOpts, id, tariffGroupId, name, action, calcMethod, price)
}

// AddTariffGroup is a paid mutator transaction binding the contract method 0x9228d917.
//
// Solidity: function addTariffGroup(id bytes16, ownerId bytes16, name string) returns()
func (_HoQuPlatform *HoQuPlatformTransactor) AddTariffGroup(opts *bind.TransactOpts, id [16]byte, ownerId [16]byte, name string) (*types.Transaction, error) {
	return _HoQuPlatform.contract.Transact(opts, "addTariffGroup", id, ownerId, name)
}

// AddTariffGroup is a paid mutator transaction binding the contract method 0x9228d917.
//
// Solidity: function addTariffGroup(id bytes16, ownerId bytes16, name string) returns()
func (_HoQuPlatform *HoQuPlatformSession) AddTariffGroup(id [16]byte, ownerId [16]byte, name string) (*types.Transaction, error) {
	return _HoQuPlatform.Contract.AddTariffGroup(&_HoQuPlatform.TransactOpts, id, ownerId, name)
}

// AddTariffGroup is a paid mutator transaction binding the contract method 0x9228d917.
//
// Solidity: function addTariffGroup(id bytes16, ownerId bytes16, name string) returns()
func (_HoQuPlatform *HoQuPlatformTransactorSession) AddTariffGroup(id [16]byte, ownerId [16]byte, name string) (*types.Transaction, error) {
	return _HoQuPlatform.Contract.AddTariffGroup(&_HoQuPlatform.TransactOpts, id, ownerId, name)
}

// AddUserAddress is a paid mutator transaction binding the contract method 0x03806028.
//
// Solidity: function addUserAddress(id bytes16, ownerAddress address) returns()
func (_HoQuPlatform *HoQuPlatformTransactor) AddUserAddress(opts *bind.TransactOpts, id [16]byte, ownerAddress common.Address) (*types.Transaction, error) {
	return _HoQuPlatform.contract.Transact(opts, "addUserAddress", id, ownerAddress)
}

// AddUserAddress is a paid mutator transaction binding the contract method 0x03806028.
//
// Solidity: function addUserAddress(id bytes16, ownerAddress address) returns()
func (_HoQuPlatform *HoQuPlatformSession) AddUserAddress(id [16]byte, ownerAddress common.Address) (*types.Transaction, error) {
	return _HoQuPlatform.Contract.AddUserAddress(&_HoQuPlatform.TransactOpts, id, ownerAddress)
}

// AddUserAddress is a paid mutator transaction binding the contract method 0x03806028.
//
// Solidity: function addUserAddress(id bytes16, ownerAddress address) returns()
func (_HoQuPlatform *HoQuPlatformTransactorSession) AddUserAddress(id [16]byte, ownerAddress common.Address) (*types.Transaction, error) {
	return _HoQuPlatform.Contract.AddUserAddress(&_HoQuPlatform.TransactOpts, id, ownerAddress)
}

// RegisterCompany is a paid mutator transaction binding the contract method 0x9f5dec06.
//
// Solidity: function registerCompany(id bytes16, ownerId bytes16, name string, dataUrl string) returns()
func (_HoQuPlatform *HoQuPlatformTransactor) RegisterCompany(opts *bind.TransactOpts, id [16]byte, ownerId [16]byte, name string, dataUrl string) (*types.Transaction, error) {
	return _HoQuPlatform.contract.Transact(opts, "registerCompany", id, ownerId, name, dataUrl)
}

// RegisterCompany is a paid mutator transaction binding the contract method 0x9f5dec06.
//
// Solidity: function registerCompany(id bytes16, ownerId bytes16, name string, dataUrl string) returns()
func (_HoQuPlatform *HoQuPlatformSession) RegisterCompany(id [16]byte, ownerId [16]byte, name string, dataUrl string) (*types.Transaction, error) {
	return _HoQuPlatform.Contract.RegisterCompany(&_HoQuPlatform.TransactOpts, id, ownerId, name, dataUrl)
}

// RegisterCompany is a paid mutator transaction binding the contract method 0x9f5dec06.
//
// Solidity: function registerCompany(id bytes16, ownerId bytes16, name string, dataUrl string) returns()
func (_HoQuPlatform *HoQuPlatformTransactorSession) RegisterCompany(id [16]byte, ownerId [16]byte, name string, dataUrl string) (*types.Transaction, error) {
	return _HoQuPlatform.Contract.RegisterCompany(&_HoQuPlatform.TransactOpts, id, ownerId, name, dataUrl)
}

// RegisterNetwork is a paid mutator transaction binding the contract method 0x7618c810.
//
// Solidity: function registerNetwork(id bytes16, ownerId bytes16, name string, dataUrl string) returns()
func (_HoQuPlatform *HoQuPlatformTransactor) RegisterNetwork(opts *bind.TransactOpts, id [16]byte, ownerId [16]byte, name string, dataUrl string) (*types.Transaction, error) {
	return _HoQuPlatform.contract.Transact(opts, "registerNetwork", id, ownerId, name, dataUrl)
}

// RegisterNetwork is a paid mutator transaction binding the contract method 0x7618c810.
//
// Solidity: function registerNetwork(id bytes16, ownerId bytes16, name string, dataUrl string) returns()
func (_HoQuPlatform *HoQuPlatformSession) RegisterNetwork(id [16]byte, ownerId [16]byte, name string, dataUrl string) (*types.Transaction, error) {
	return _HoQuPlatform.Contract.RegisterNetwork(&_HoQuPlatform.TransactOpts, id, ownerId, name, dataUrl)
}

// RegisterNetwork is a paid mutator transaction binding the contract method 0x7618c810.
//
// Solidity: function registerNetwork(id bytes16, ownerId bytes16, name string, dataUrl string) returns()
func (_HoQuPlatform *HoQuPlatformTransactorSession) RegisterNetwork(id [16]byte, ownerId [16]byte, name string, dataUrl string) (*types.Transaction, error) {
	return _HoQuPlatform.Contract.RegisterNetwork(&_HoQuPlatform.TransactOpts, id, ownerId, name, dataUrl)
}

// RegisterTracker is a paid mutator transaction binding the contract method 0x5392a3f2.
//
// Solidity: function registerTracker(id bytes16, ownerId bytes16, networkId bytes16, name string, dataUrl string) returns()
func (_HoQuPlatform *HoQuPlatformTransactor) RegisterTracker(opts *bind.TransactOpts, id [16]byte, ownerId [16]byte, networkId [16]byte, name string, dataUrl string) (*types.Transaction, error) {
	return _HoQuPlatform.contract.Transact(opts, "registerTracker", id, ownerId, networkId, name, dataUrl)
}

// RegisterTracker is a paid mutator transaction binding the contract method 0x5392a3f2.
//
// Solidity: function registerTracker(id bytes16, ownerId bytes16, networkId bytes16, name string, dataUrl string) returns()
func (_HoQuPlatform *HoQuPlatformSession) RegisterTracker(id [16]byte, ownerId [16]byte, networkId [16]byte, name string, dataUrl string) (*types.Transaction, error) {
	return _HoQuPlatform.Contract.RegisterTracker(&_HoQuPlatform.TransactOpts, id, ownerId, networkId, name, dataUrl)
}

// RegisterTracker is a paid mutator transaction binding the contract method 0x5392a3f2.
//
// Solidity: function registerTracker(id bytes16, ownerId bytes16, networkId bytes16, name string, dataUrl string) returns()
func (_HoQuPlatform *HoQuPlatformTransactorSession) RegisterTracker(id [16]byte, ownerId [16]byte, networkId [16]byte, name string, dataUrl string) (*types.Transaction, error) {
	return _HoQuPlatform.Contract.RegisterTracker(&_HoQuPlatform.TransactOpts, id, ownerId, networkId, name, dataUrl)
}

// RegisterUser is a paid mutator transaction binding the contract method 0xfd3fc08f.
//
// Solidity: function registerUser(id bytes16, role string, ownerAddress address, pubKey string) returns()
func (_HoQuPlatform *HoQuPlatformTransactor) RegisterUser(opts *bind.TransactOpts, id [16]byte, role string, ownerAddress common.Address, pubKey string) (*types.Transaction, error) {
	return _HoQuPlatform.contract.Transact(opts, "registerUser", id, role, ownerAddress, pubKey)
}

// RegisterUser is a paid mutator transaction binding the contract method 0xfd3fc08f.
//
// Solidity: function registerUser(id bytes16, role string, ownerAddress address, pubKey string) returns()
func (_HoQuPlatform *HoQuPlatformSession) RegisterUser(id [16]byte, role string, ownerAddress common.Address, pubKey string) (*types.Transaction, error) {
	return _HoQuPlatform.Contract.RegisterUser(&_HoQuPlatform.TransactOpts, id, role, ownerAddress, pubKey)
}

// RegisterUser is a paid mutator transaction binding the contract method 0xfd3fc08f.
//
// Solidity: function registerUser(id bytes16, role string, ownerAddress address, pubKey string) returns()
func (_HoQuPlatform *HoQuPlatformTransactorSession) RegisterUser(id [16]byte, role string, ownerAddress common.Address, pubKey string) (*types.Transaction, error) {
	return _HoQuPlatform.Contract.RegisterUser(&_HoQuPlatform.TransactOpts, id, role, ownerAddress, pubKey)
}

// SetAdCampaignStatus is a paid mutator transaction binding the contract method 0xd7915dfb.
//
// Solidity: function setAdCampaignStatus(id bytes16, status uint8) returns()
func (_HoQuPlatform *HoQuPlatformTransactor) SetAdCampaignStatus(opts *bind.TransactOpts, id [16]byte, status uint8) (*types.Transaction, error) {
	return _HoQuPlatform.contract.Transact(opts, "setAdCampaignStatus", id, status)
}

// SetAdCampaignStatus is a paid mutator transaction binding the contract method 0xd7915dfb.
//
// Solidity: function setAdCampaignStatus(id bytes16, status uint8) returns()
func (_HoQuPlatform *HoQuPlatformSession) SetAdCampaignStatus(id [16]byte, status uint8) (*types.Transaction, error) {
	return _HoQuPlatform.Contract.SetAdCampaignStatus(&_HoQuPlatform.TransactOpts, id, status)
}

// SetAdCampaignStatus is a paid mutator transaction binding the contract method 0xd7915dfb.
//
// Solidity: function setAdCampaignStatus(id bytes16, status uint8) returns()
func (_HoQuPlatform *HoQuPlatformTransactorSession) SetAdCampaignStatus(id [16]byte, status uint8) (*types.Transaction, error) {
	return _HoQuPlatform.Contract.SetAdCampaignStatus(&_HoQuPlatform.TransactOpts, id, status)
}

// SetCompanyStatus is a paid mutator transaction binding the contract method 0xa8800f09.
//
// Solidity: function setCompanyStatus(id bytes16, status uint8) returns()
func (_HoQuPlatform *HoQuPlatformTransactor) SetCompanyStatus(opts *bind.TransactOpts, id [16]byte, status uint8) (*types.Transaction, error) {
	return _HoQuPlatform.contract.Transact(opts, "setCompanyStatus", id, status)
}

// SetCompanyStatus is a paid mutator transaction binding the contract method 0xa8800f09.
//
// Solidity: function setCompanyStatus(id bytes16, status uint8) returns()
func (_HoQuPlatform *HoQuPlatformSession) SetCompanyStatus(id [16]byte, status uint8) (*types.Transaction, error) {
	return _HoQuPlatform.Contract.SetCompanyStatus(&_HoQuPlatform.TransactOpts, id, status)
}

// SetCompanyStatus is a paid mutator transaction binding the contract method 0xa8800f09.
//
// Solidity: function setCompanyStatus(id bytes16, status uint8) returns()
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

// SetLeadDataUrl is a paid mutator transaction binding the contract method 0xa6ccf89d.
//
// Solidity: function setLeadDataUrl(id bytes16, adCampaignId bytes16, _dataUrl string) returns()
func (_HoQuPlatform *HoQuPlatformTransactor) SetLeadDataUrl(opts *bind.TransactOpts, id [16]byte, adCampaignId [16]byte, _dataUrl string) (*types.Transaction, error) {
	return _HoQuPlatform.contract.Transact(opts, "setLeadDataUrl", id, adCampaignId, _dataUrl)
}

// SetLeadDataUrl is a paid mutator transaction binding the contract method 0xa6ccf89d.
//
// Solidity: function setLeadDataUrl(id bytes16, adCampaignId bytes16, _dataUrl string) returns()
func (_HoQuPlatform *HoQuPlatformSession) SetLeadDataUrl(id [16]byte, adCampaignId [16]byte, _dataUrl string) (*types.Transaction, error) {
	return _HoQuPlatform.Contract.SetLeadDataUrl(&_HoQuPlatform.TransactOpts, id, adCampaignId, _dataUrl)
}

// SetLeadDataUrl is a paid mutator transaction binding the contract method 0xa6ccf89d.
//
// Solidity: function setLeadDataUrl(id bytes16, adCampaignId bytes16, _dataUrl string) returns()
func (_HoQuPlatform *HoQuPlatformTransactorSession) SetLeadDataUrl(id [16]byte, adCampaignId [16]byte, _dataUrl string) (*types.Transaction, error) {
	return _HoQuPlatform.Contract.SetLeadDataUrl(&_HoQuPlatform.TransactOpts, id, adCampaignId, _dataUrl)
}

// SetLeadPrice is a paid mutator transaction binding the contract method 0x79b4330b.
//
// Solidity: function setLeadPrice(id bytes16, adCampaignId bytes16, _price uint256) returns()
func (_HoQuPlatform *HoQuPlatformTransactor) SetLeadPrice(opts *bind.TransactOpts, id [16]byte, adCampaignId [16]byte, _price *big.Int) (*types.Transaction, error) {
	return _HoQuPlatform.contract.Transact(opts, "setLeadPrice", id, adCampaignId, _price)
}

// SetLeadPrice is a paid mutator transaction binding the contract method 0x79b4330b.
//
// Solidity: function setLeadPrice(id bytes16, adCampaignId bytes16, _price uint256) returns()
func (_HoQuPlatform *HoQuPlatformSession) SetLeadPrice(id [16]byte, adCampaignId [16]byte, _price *big.Int) (*types.Transaction, error) {
	return _HoQuPlatform.Contract.SetLeadPrice(&_HoQuPlatform.TransactOpts, id, adCampaignId, _price)
}

// SetLeadPrice is a paid mutator transaction binding the contract method 0x79b4330b.
//
// Solidity: function setLeadPrice(id bytes16, adCampaignId bytes16, _price uint256) returns()
func (_HoQuPlatform *HoQuPlatformTransactorSession) SetLeadPrice(id [16]byte, adCampaignId [16]byte, _price *big.Int) (*types.Transaction, error) {
	return _HoQuPlatform.Contract.SetLeadPrice(&_HoQuPlatform.TransactOpts, id, adCampaignId, _price)
}

// SetLeadStatus is a paid mutator transaction binding the contract method 0x68970854.
//
// Solidity: function setLeadStatus(id bytes16, adCampaignId bytes16, _status uint8) returns()
func (_HoQuPlatform *HoQuPlatformTransactor) SetLeadStatus(opts *bind.TransactOpts, id [16]byte, adCampaignId [16]byte, _status uint8) (*types.Transaction, error) {
	return _HoQuPlatform.contract.Transact(opts, "setLeadStatus", id, adCampaignId, _status)
}

// SetLeadStatus is a paid mutator transaction binding the contract method 0x68970854.
//
// Solidity: function setLeadStatus(id bytes16, adCampaignId bytes16, _status uint8) returns()
func (_HoQuPlatform *HoQuPlatformSession) SetLeadStatus(id [16]byte, adCampaignId [16]byte, _status uint8) (*types.Transaction, error) {
	return _HoQuPlatform.Contract.SetLeadStatus(&_HoQuPlatform.TransactOpts, id, adCampaignId, _status)
}

// SetLeadStatus is a paid mutator transaction binding the contract method 0x68970854.
//
// Solidity: function setLeadStatus(id bytes16, adCampaignId bytes16, _status uint8) returns()
func (_HoQuPlatform *HoQuPlatformTransactorSession) SetLeadStatus(id [16]byte, adCampaignId [16]byte, _status uint8) (*types.Transaction, error) {
	return _HoQuPlatform.Contract.SetLeadStatus(&_HoQuPlatform.TransactOpts, id, adCampaignId, _status)
}

// SetNetworkStatus is a paid mutator transaction binding the contract method 0x007c9b51.
//
// Solidity: function setNetworkStatus(id bytes16, status uint8) returns()
func (_HoQuPlatform *HoQuPlatformTransactor) SetNetworkStatus(opts *bind.TransactOpts, id [16]byte, status uint8) (*types.Transaction, error) {
	return _HoQuPlatform.contract.Transact(opts, "setNetworkStatus", id, status)
}

// SetNetworkStatus is a paid mutator transaction binding the contract method 0x007c9b51.
//
// Solidity: function setNetworkStatus(id bytes16, status uint8) returns()
func (_HoQuPlatform *HoQuPlatformSession) SetNetworkStatus(id [16]byte, status uint8) (*types.Transaction, error) {
	return _HoQuPlatform.Contract.SetNetworkStatus(&_HoQuPlatform.TransactOpts, id, status)
}

// SetNetworkStatus is a paid mutator transaction binding the contract method 0x007c9b51.
//
// Solidity: function setNetworkStatus(id bytes16, status uint8) returns()
func (_HoQuPlatform *HoQuPlatformTransactorSession) SetNetworkStatus(id [16]byte, status uint8) (*types.Transaction, error) {
	return _HoQuPlatform.Contract.SetNetworkStatus(&_HoQuPlatform.TransactOpts, id, status)
}

// SetOfferName is a paid mutator transaction binding the contract method 0x4af4c3fd.
//
// Solidity: function setOfferName(id bytes16, name string) returns()
func (_HoQuPlatform *HoQuPlatformTransactor) SetOfferName(opts *bind.TransactOpts, id [16]byte, name string) (*types.Transaction, error) {
	return _HoQuPlatform.contract.Transact(opts, "setOfferName", id, name)
}

// SetOfferName is a paid mutator transaction binding the contract method 0x4af4c3fd.
//
// Solidity: function setOfferName(id bytes16, name string) returns()
func (_HoQuPlatform *HoQuPlatformSession) SetOfferName(id [16]byte, name string) (*types.Transaction, error) {
	return _HoQuPlatform.Contract.SetOfferName(&_HoQuPlatform.TransactOpts, id, name)
}

// SetOfferName is a paid mutator transaction binding the contract method 0x4af4c3fd.
//
// Solidity: function setOfferName(id bytes16, name string) returns()
func (_HoQuPlatform *HoQuPlatformTransactorSession) SetOfferName(id [16]byte, name string) (*types.Transaction, error) {
	return _HoQuPlatform.Contract.SetOfferName(&_HoQuPlatform.TransactOpts, id, name)
}

// SetOfferPayerAddress is a paid mutator transaction binding the contract method 0x22e1d763.
//
// Solidity: function setOfferPayerAddress(id bytes16, payerAddress address) returns()
func (_HoQuPlatform *HoQuPlatformTransactor) SetOfferPayerAddress(opts *bind.TransactOpts, id [16]byte, payerAddress common.Address) (*types.Transaction, error) {
	return _HoQuPlatform.contract.Transact(opts, "setOfferPayerAddress", id, payerAddress)
}

// SetOfferPayerAddress is a paid mutator transaction binding the contract method 0x22e1d763.
//
// Solidity: function setOfferPayerAddress(id bytes16, payerAddress address) returns()
func (_HoQuPlatform *HoQuPlatformSession) SetOfferPayerAddress(id [16]byte, payerAddress common.Address) (*types.Transaction, error) {
	return _HoQuPlatform.Contract.SetOfferPayerAddress(&_HoQuPlatform.TransactOpts, id, payerAddress)
}

// SetOfferPayerAddress is a paid mutator transaction binding the contract method 0x22e1d763.
//
// Solidity: function setOfferPayerAddress(id bytes16, payerAddress address) returns()
func (_HoQuPlatform *HoQuPlatformTransactorSession) SetOfferPayerAddress(id [16]byte, payerAddress common.Address) (*types.Transaction, error) {
	return _HoQuPlatform.Contract.SetOfferPayerAddress(&_HoQuPlatform.TransactOpts, id, payerAddress)
}

// SetOfferStatus is a paid mutator transaction binding the contract method 0x36d7523f.
//
// Solidity: function setOfferStatus(id bytes16, status uint8) returns()
func (_HoQuPlatform *HoQuPlatformTransactor) SetOfferStatus(opts *bind.TransactOpts, id [16]byte, status uint8) (*types.Transaction, error) {
	return _HoQuPlatform.contract.Transact(opts, "setOfferStatus", id, status)
}

// SetOfferStatus is a paid mutator transaction binding the contract method 0x36d7523f.
//
// Solidity: function setOfferStatus(id bytes16, status uint8) returns()
func (_HoQuPlatform *HoQuPlatformSession) SetOfferStatus(id [16]byte, status uint8) (*types.Transaction, error) {
	return _HoQuPlatform.Contract.SetOfferStatus(&_HoQuPlatform.TransactOpts, id, status)
}

// SetOfferStatus is a paid mutator transaction binding the contract method 0x36d7523f.
//
// Solidity: function setOfferStatus(id bytes16, status uint8) returns()
func (_HoQuPlatform *HoQuPlatformTransactorSession) SetOfferStatus(id [16]byte, status uint8) (*types.Transaction, error) {
	return _HoQuPlatform.Contract.SetOfferStatus(&_HoQuPlatform.TransactOpts, id, status)
}

// SetOfferTariffGroup is a paid mutator transaction binding the contract method 0x3e45850a.
//
// Solidity: function setOfferTariffGroup(id bytes16, num uint8, tariffGroupId bytes16) returns()
func (_HoQuPlatform *HoQuPlatformTransactor) SetOfferTariffGroup(opts *bind.TransactOpts, id [16]byte, num uint8, tariffGroupId [16]byte) (*types.Transaction, error) {
	return _HoQuPlatform.contract.Transact(opts, "setOfferTariffGroup", id, num, tariffGroupId)
}

// SetOfferTariffGroup is a paid mutator transaction binding the contract method 0x3e45850a.
//
// Solidity: function setOfferTariffGroup(id bytes16, num uint8, tariffGroupId bytes16) returns()
func (_HoQuPlatform *HoQuPlatformSession) SetOfferTariffGroup(id [16]byte, num uint8, tariffGroupId [16]byte) (*types.Transaction, error) {
	return _HoQuPlatform.Contract.SetOfferTariffGroup(&_HoQuPlatform.TransactOpts, id, num, tariffGroupId)
}

// SetOfferTariffGroup is a paid mutator transaction binding the contract method 0x3e45850a.
//
// Solidity: function setOfferTariffGroup(id bytes16, num uint8, tariffGroupId bytes16) returns()
func (_HoQuPlatform *HoQuPlatformTransactorSession) SetOfferTariffGroup(id [16]byte, num uint8, tariffGroupId [16]byte) (*types.Transaction, error) {
	return _HoQuPlatform.Contract.SetOfferTariffGroup(&_HoQuPlatform.TransactOpts, id, num, tariffGroupId)
}

// SetStorageAddress is a paid mutator transaction binding the contract method 0x59b910d6.
//
// Solidity: function setStorageAddress(storageAddress address) returns()
func (_HoQuPlatform *HoQuPlatformTransactor) SetStorageAddress(opts *bind.TransactOpts, storageAddress common.Address) (*types.Transaction, error) {
	return _HoQuPlatform.contract.Transact(opts, "setStorageAddress", storageAddress)
}

// SetStorageAddress is a paid mutator transaction binding the contract method 0x59b910d6.
//
// Solidity: function setStorageAddress(storageAddress address) returns()
func (_HoQuPlatform *HoQuPlatformSession) SetStorageAddress(storageAddress common.Address) (*types.Transaction, error) {
	return _HoQuPlatform.Contract.SetStorageAddress(&_HoQuPlatform.TransactOpts, storageAddress)
}

// SetStorageAddress is a paid mutator transaction binding the contract method 0x59b910d6.
//
// Solidity: function setStorageAddress(storageAddress address) returns()
func (_HoQuPlatform *HoQuPlatformTransactorSession) SetStorageAddress(storageAddress common.Address) (*types.Transaction, error) {
	return _HoQuPlatform.Contract.SetStorageAddress(&_HoQuPlatform.TransactOpts, storageAddress)
}

// SetTariffGroupStatus is a paid mutator transaction binding the contract method 0x310eea12.
//
// Solidity: function setTariffGroupStatus(id bytes16, status uint8) returns()
func (_HoQuPlatform *HoQuPlatformTransactor) SetTariffGroupStatus(opts *bind.TransactOpts, id [16]byte, status uint8) (*types.Transaction, error) {
	return _HoQuPlatform.contract.Transact(opts, "setTariffGroupStatus", id, status)
}

// SetTariffGroupStatus is a paid mutator transaction binding the contract method 0x310eea12.
//
// Solidity: function setTariffGroupStatus(id bytes16, status uint8) returns()
func (_HoQuPlatform *HoQuPlatformSession) SetTariffGroupStatus(id [16]byte, status uint8) (*types.Transaction, error) {
	return _HoQuPlatform.Contract.SetTariffGroupStatus(&_HoQuPlatform.TransactOpts, id, status)
}

// SetTariffGroupStatus is a paid mutator transaction binding the contract method 0x310eea12.
//
// Solidity: function setTariffGroupStatus(id bytes16, status uint8) returns()
func (_HoQuPlatform *HoQuPlatformTransactorSession) SetTariffGroupStatus(id [16]byte, status uint8) (*types.Transaction, error) {
	return _HoQuPlatform.Contract.SetTariffGroupStatus(&_HoQuPlatform.TransactOpts, id, status)
}

// SetTariffStatus is a paid mutator transaction binding the contract method 0xe1ef8935.
//
// Solidity: function setTariffStatus(id bytes16, status uint8) returns()
func (_HoQuPlatform *HoQuPlatformTransactor) SetTariffStatus(opts *bind.TransactOpts, id [16]byte, status uint8) (*types.Transaction, error) {
	return _HoQuPlatform.contract.Transact(opts, "setTariffStatus", id, status)
}

// SetTariffStatus is a paid mutator transaction binding the contract method 0xe1ef8935.
//
// Solidity: function setTariffStatus(id bytes16, status uint8) returns()
func (_HoQuPlatform *HoQuPlatformSession) SetTariffStatus(id [16]byte, status uint8) (*types.Transaction, error) {
	return _HoQuPlatform.Contract.SetTariffStatus(&_HoQuPlatform.TransactOpts, id, status)
}

// SetTariffStatus is a paid mutator transaction binding the contract method 0xe1ef8935.
//
// Solidity: function setTariffStatus(id bytes16, status uint8) returns()
func (_HoQuPlatform *HoQuPlatformTransactorSession) SetTariffStatus(id [16]byte, status uint8) (*types.Transaction, error) {
	return _HoQuPlatform.Contract.SetTariffStatus(&_HoQuPlatform.TransactOpts, id, status)
}

// SetTrackerStatus is a paid mutator transaction binding the contract method 0xe7e5aad8.
//
// Solidity: function setTrackerStatus(id bytes16, status uint8) returns()
func (_HoQuPlatform *HoQuPlatformTransactor) SetTrackerStatus(opts *bind.TransactOpts, id [16]byte, status uint8) (*types.Transaction, error) {
	return _HoQuPlatform.contract.Transact(opts, "setTrackerStatus", id, status)
}

// SetTrackerStatus is a paid mutator transaction binding the contract method 0xe7e5aad8.
//
// Solidity: function setTrackerStatus(id bytes16, status uint8) returns()
func (_HoQuPlatform *HoQuPlatformSession) SetTrackerStatus(id [16]byte, status uint8) (*types.Transaction, error) {
	return _HoQuPlatform.Contract.SetTrackerStatus(&_HoQuPlatform.TransactOpts, id, status)
}

// SetTrackerStatus is a paid mutator transaction binding the contract method 0xe7e5aad8.
//
// Solidity: function setTrackerStatus(id bytes16, status uint8) returns()
func (_HoQuPlatform *HoQuPlatformTransactorSession) SetTrackerStatus(id [16]byte, status uint8) (*types.Transaction, error) {
	return _HoQuPlatform.Contract.SetTrackerStatus(&_HoQuPlatform.TransactOpts, id, status)
}

// SetUserStatus is a paid mutator transaction binding the contract method 0xe03d8b64.
//
// Solidity: function setUserStatus(id bytes16, status uint8) returns()
func (_HoQuPlatform *HoQuPlatformTransactor) SetUserStatus(opts *bind.TransactOpts, id [16]byte, status uint8) (*types.Transaction, error) {
	return _HoQuPlatform.contract.Transact(opts, "setUserStatus", id, status)
}

// SetUserStatus is a paid mutator transaction binding the contract method 0xe03d8b64.
//
// Solidity: function setUserStatus(id bytes16, status uint8) returns()
func (_HoQuPlatform *HoQuPlatformSession) SetUserStatus(id [16]byte, status uint8) (*types.Transaction, error) {
	return _HoQuPlatform.Contract.SetUserStatus(&_HoQuPlatform.TransactOpts, id, status)
}

// SetUserStatus is a paid mutator transaction binding the contract method 0xe03d8b64.
//
// Solidity: function setUserStatus(id bytes16, status uint8) returns()
func (_HoQuPlatform *HoQuPlatformTransactorSession) SetUserStatus(id [16]byte, status uint8) (*types.Transaction, error) {
	return _HoQuPlatform.Contract.SetUserStatus(&_HoQuPlatform.TransactOpts, id, status)
}

// TransactLead is a paid mutator transaction binding the contract method 0x05b8c83e.
//
// Solidity: function transactLead(id bytes16, adCampaignId bytes16) returns()
func (_HoQuPlatform *HoQuPlatformTransactor) TransactLead(opts *bind.TransactOpts, id [16]byte, adCampaignId [16]byte) (*types.Transaction, error) {
	return _HoQuPlatform.contract.Transact(opts, "transactLead", id, adCampaignId)
}

// TransactLead is a paid mutator transaction binding the contract method 0x05b8c83e.
//
// Solidity: function transactLead(id bytes16, adCampaignId bytes16) returns()
func (_HoQuPlatform *HoQuPlatformSession) TransactLead(id [16]byte, adCampaignId [16]byte) (*types.Transaction, error) {
	return _HoQuPlatform.Contract.TransactLead(&_HoQuPlatform.TransactOpts, id, adCampaignId)
}

// TransactLead is a paid mutator transaction binding the contract method 0x05b8c83e.
//
// Solidity: function transactLead(id bytes16, adCampaignId bytes16) returns()
func (_HoQuPlatform *HoQuPlatformTransactorSession) TransactLead(id [16]byte, adCampaignId [16]byte) (*types.Transaction, error) {
	return _HoQuPlatform.Contract.TransactLead(&_HoQuPlatform.TransactOpts, id, adCampaignId)
}

// UpdateUserPubKey is a paid mutator transaction binding the contract method 0x8a5463bd.
//
// Solidity: function updateUserPubKey(id bytes16, pubKey string) returns()
func (_HoQuPlatform *HoQuPlatformTransactor) UpdateUserPubKey(opts *bind.TransactOpts, id [16]byte, pubKey string) (*types.Transaction, error) {
	return _HoQuPlatform.contract.Transact(opts, "updateUserPubKey", id, pubKey)
}

// UpdateUserPubKey is a paid mutator transaction binding the contract method 0x8a5463bd.
//
// Solidity: function updateUserPubKey(id bytes16, pubKey string) returns()
func (_HoQuPlatform *HoQuPlatformSession) UpdateUserPubKey(id [16]byte, pubKey string) (*types.Transaction, error) {
	return _HoQuPlatform.Contract.UpdateUserPubKey(&_HoQuPlatform.TransactOpts, id, pubKey)
}

// UpdateUserPubKey is a paid mutator transaction binding the contract method 0x8a5463bd.
//
// Solidity: function updateUserPubKey(id bytes16, pubKey string) returns()
func (_HoQuPlatform *HoQuPlatformTransactorSession) UpdateUserPubKey(id [16]byte, pubKey string) (*types.Transaction, error) {
	return _HoQuPlatform.Contract.UpdateUserPubKey(&_HoQuPlatform.TransactOpts, id, pubKey)
}

// HoQuPlatformAdCampaignAddedIterator is returned from FilterAdCampaignAdded and is used to iterate over the raw logs and unpacked data for AdCampaignAdded events raised by the HoQuPlatform contract.
type HoQuPlatformAdCampaignAddedIterator struct {
	Event *HoQuPlatformAdCampaignAdded // Event containing the contract specifics and raw log

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
func (it *HoQuPlatformAdCampaignAddedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(HoQuPlatformAdCampaignAdded)
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
		it.Event = new(HoQuPlatformAdCampaignAdded)
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
func (it *HoQuPlatformAdCampaignAddedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *HoQuPlatformAdCampaignAddedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// HoQuPlatformAdCampaignAdded represents a AdCampaignAdded event raised by the HoQuPlatform contract.
type HoQuPlatformAdCampaignAdded struct {
	OwnerAddress    common.Address
	Id              [16]byte
	ContractAddress common.Address
	Raw             types.Log // Blockchain specific contextual infos
}

// FilterAdCampaignAdded is a free log retrieval operation binding the contract event 0x62a609593cded926aabb6aef1e5b884804ad29d1bd1116ac4ce32085690d5ab2.
//
// Solidity: event AdCampaignAdded(ownerAddress indexed address, id bytes16, contractAddress address)
func (_HoQuPlatform *HoQuPlatformFilterer) FilterAdCampaignAdded(opts *bind.FilterOpts, ownerAddress []common.Address) (*HoQuPlatformAdCampaignAddedIterator, error) {

	var ownerAddressRule []interface{}
	for _, ownerAddressItem := range ownerAddress {
		ownerAddressRule = append(ownerAddressRule, ownerAddressItem)
	}

	logs, sub, err := _HoQuPlatform.contract.FilterLogs(opts, "AdCampaignAdded", ownerAddressRule)
	if err != nil {
		return nil, err
	}
	return &HoQuPlatformAdCampaignAddedIterator{contract: _HoQuPlatform.contract, event: "AdCampaignAdded", logs: logs, sub: sub}, nil
}

// WatchAdCampaignAdded is a free log subscription operation binding the contract event 0x62a609593cded926aabb6aef1e5b884804ad29d1bd1116ac4ce32085690d5ab2.
//
// Solidity: event AdCampaignAdded(ownerAddress indexed address, id bytes16, contractAddress address)
func (_HoQuPlatform *HoQuPlatformFilterer) WatchAdCampaignAdded(opts *bind.WatchOpts, sink chan<- *HoQuPlatformAdCampaignAdded, ownerAddress []common.Address) (event.Subscription, error) {

	var ownerAddressRule []interface{}
	for _, ownerAddressItem := range ownerAddress {
		ownerAddressRule = append(ownerAddressRule, ownerAddressItem)
	}

	logs, sub, err := _HoQuPlatform.contract.WatchLogs(opts, "AdCampaignAdded", ownerAddressRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(HoQuPlatformAdCampaignAdded)
				if err := _HoQuPlatform.contract.UnpackLog(event, "AdCampaignAdded", log); err != nil {
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

// HoQuPlatformAdCampaignStatusChangedIterator is returned from FilterAdCampaignStatusChanged and is used to iterate over the raw logs and unpacked data for AdCampaignStatusChanged events raised by the HoQuPlatform contract.
type HoQuPlatformAdCampaignStatusChangedIterator struct {
	Event *HoQuPlatformAdCampaignStatusChanged // Event containing the contract specifics and raw log

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
func (it *HoQuPlatformAdCampaignStatusChangedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(HoQuPlatformAdCampaignStatusChanged)
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
		it.Event = new(HoQuPlatformAdCampaignStatusChanged)
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
func (it *HoQuPlatformAdCampaignStatusChangedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *HoQuPlatformAdCampaignStatusChangedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// HoQuPlatformAdCampaignStatusChanged represents a AdCampaignStatusChanged event raised by the HoQuPlatform contract.
type HoQuPlatformAdCampaignStatusChanged struct {
	OwnerAddress common.Address
	Id           [16]byte
	Status       uint8
	Raw          types.Log // Blockchain specific contextual infos
}

// FilterAdCampaignStatusChanged is a free log retrieval operation binding the contract event 0xa25922380013ceba1e6f7c02184e30818a465ac4d7107ac4357b5883887b5af4.
//
// Solidity: event AdCampaignStatusChanged(ownerAddress indexed address, id bytes16, status uint8)
func (_HoQuPlatform *HoQuPlatformFilterer) FilterAdCampaignStatusChanged(opts *bind.FilterOpts, ownerAddress []common.Address) (*HoQuPlatformAdCampaignStatusChangedIterator, error) {

	var ownerAddressRule []interface{}
	for _, ownerAddressItem := range ownerAddress {
		ownerAddressRule = append(ownerAddressRule, ownerAddressItem)
	}

	logs, sub, err := _HoQuPlatform.contract.FilterLogs(opts, "AdCampaignStatusChanged", ownerAddressRule)
	if err != nil {
		return nil, err
	}
	return &HoQuPlatformAdCampaignStatusChangedIterator{contract: _HoQuPlatform.contract, event: "AdCampaignStatusChanged", logs: logs, sub: sub}, nil
}

// WatchAdCampaignStatusChanged is a free log subscription operation binding the contract event 0xa25922380013ceba1e6f7c02184e30818a465ac4d7107ac4357b5883887b5af4.
//
// Solidity: event AdCampaignStatusChanged(ownerAddress indexed address, id bytes16, status uint8)
func (_HoQuPlatform *HoQuPlatformFilterer) WatchAdCampaignStatusChanged(opts *bind.WatchOpts, sink chan<- *HoQuPlatformAdCampaignStatusChanged, ownerAddress []common.Address) (event.Subscription, error) {

	var ownerAddressRule []interface{}
	for _, ownerAddressItem := range ownerAddress {
		ownerAddressRule = append(ownerAddressRule, ownerAddressItem)
	}

	logs, sub, err := _HoQuPlatform.contract.WatchLogs(opts, "AdCampaignStatusChanged", ownerAddressRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(HoQuPlatformAdCampaignStatusChanged)
				if err := _HoQuPlatform.contract.UnpackLog(event, "AdCampaignStatusChanged", log); err != nil {
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

// HoQuPlatformCompanyRegisteredIterator is returned from FilterCompanyRegistered and is used to iterate over the raw logs and unpacked data for CompanyRegistered events raised by the HoQuPlatform contract.
type HoQuPlatformCompanyRegisteredIterator struct {
	Event *HoQuPlatformCompanyRegistered // Event containing the contract specifics and raw log

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
func (it *HoQuPlatformCompanyRegisteredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(HoQuPlatformCompanyRegistered)
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
		it.Event = new(HoQuPlatformCompanyRegistered)
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
func (it *HoQuPlatformCompanyRegisteredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *HoQuPlatformCompanyRegisteredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// HoQuPlatformCompanyRegistered represents a CompanyRegistered event raised by the HoQuPlatform contract.
type HoQuPlatformCompanyRegistered struct {
	OwnerAddress common.Address
	Id           [16]byte
	Name         string
	Raw          types.Log // Blockchain specific contextual infos
}

// FilterCompanyRegistered is a free log retrieval operation binding the contract event 0x0eb23d662c57754d23afdc1cffa6246fbae3bbe843eec5c42a5a5223c21414ed.
//
// Solidity: event CompanyRegistered(ownerAddress indexed address, id bytes16, name string)
func (_HoQuPlatform *HoQuPlatformFilterer) FilterCompanyRegistered(opts *bind.FilterOpts, ownerAddress []common.Address) (*HoQuPlatformCompanyRegisteredIterator, error) {

	var ownerAddressRule []interface{}
	for _, ownerAddressItem := range ownerAddress {
		ownerAddressRule = append(ownerAddressRule, ownerAddressItem)
	}

	logs, sub, err := _HoQuPlatform.contract.FilterLogs(opts, "CompanyRegistered", ownerAddressRule)
	if err != nil {
		return nil, err
	}
	return &HoQuPlatformCompanyRegisteredIterator{contract: _HoQuPlatform.contract, event: "CompanyRegistered", logs: logs, sub: sub}, nil
}

// WatchCompanyRegistered is a free log subscription operation binding the contract event 0x0eb23d662c57754d23afdc1cffa6246fbae3bbe843eec5c42a5a5223c21414ed.
//
// Solidity: event CompanyRegistered(ownerAddress indexed address, id bytes16, name string)
func (_HoQuPlatform *HoQuPlatformFilterer) WatchCompanyRegistered(opts *bind.WatchOpts, sink chan<- *HoQuPlatformCompanyRegistered, ownerAddress []common.Address) (event.Subscription, error) {

	var ownerAddressRule []interface{}
	for _, ownerAddressItem := range ownerAddress {
		ownerAddressRule = append(ownerAddressRule, ownerAddressItem)
	}

	logs, sub, err := _HoQuPlatform.contract.WatchLogs(opts, "CompanyRegistered", ownerAddressRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(HoQuPlatformCompanyRegistered)
				if err := _HoQuPlatform.contract.UnpackLog(event, "CompanyRegistered", log); err != nil {
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

// HoQuPlatformCompanyStatusChangedIterator is returned from FilterCompanyStatusChanged and is used to iterate over the raw logs and unpacked data for CompanyStatusChanged events raised by the HoQuPlatform contract.
type HoQuPlatformCompanyStatusChangedIterator struct {
	Event *HoQuPlatformCompanyStatusChanged // Event containing the contract specifics and raw log

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
func (it *HoQuPlatformCompanyStatusChangedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(HoQuPlatformCompanyStatusChanged)
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
		it.Event = new(HoQuPlatformCompanyStatusChanged)
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
func (it *HoQuPlatformCompanyStatusChangedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *HoQuPlatformCompanyStatusChangedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// HoQuPlatformCompanyStatusChanged represents a CompanyStatusChanged event raised by the HoQuPlatform contract.
type HoQuPlatformCompanyStatusChanged struct {
	OwnerAddress common.Address
	Id           [16]byte
	Status       uint8
	Raw          types.Log // Blockchain specific contextual infos
}

// FilterCompanyStatusChanged is a free log retrieval operation binding the contract event 0x14404ec3e30e0b13c4ef42e76c8fd3afe300c9bf51761965c78d348ff45cf2f0.
//
// Solidity: event CompanyStatusChanged(ownerAddress indexed address, id bytes16, status uint8)
func (_HoQuPlatform *HoQuPlatformFilterer) FilterCompanyStatusChanged(opts *bind.FilterOpts, ownerAddress []common.Address) (*HoQuPlatformCompanyStatusChangedIterator, error) {

	var ownerAddressRule []interface{}
	for _, ownerAddressItem := range ownerAddress {
		ownerAddressRule = append(ownerAddressRule, ownerAddressItem)
	}

	logs, sub, err := _HoQuPlatform.contract.FilterLogs(opts, "CompanyStatusChanged", ownerAddressRule)
	if err != nil {
		return nil, err
	}
	return &HoQuPlatformCompanyStatusChangedIterator{contract: _HoQuPlatform.contract, event: "CompanyStatusChanged", logs: logs, sub: sub}, nil
}

// WatchCompanyStatusChanged is a free log subscription operation binding the contract event 0x14404ec3e30e0b13c4ef42e76c8fd3afe300c9bf51761965c78d348ff45cf2f0.
//
// Solidity: event CompanyStatusChanged(ownerAddress indexed address, id bytes16, status uint8)
func (_HoQuPlatform *HoQuPlatformFilterer) WatchCompanyStatusChanged(opts *bind.WatchOpts, sink chan<- *HoQuPlatformCompanyStatusChanged, ownerAddress []common.Address) (event.Subscription, error) {

	var ownerAddressRule []interface{}
	for _, ownerAddressItem := range ownerAddress {
		ownerAddressRule = append(ownerAddressRule, ownerAddressItem)
	}

	logs, sub, err := _HoQuPlatform.contract.WatchLogs(opts, "CompanyStatusChanged", ownerAddressRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(HoQuPlatformCompanyStatusChanged)
				if err := _HoQuPlatform.contract.UnpackLog(event, "CompanyStatusChanged", log); err != nil {
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

// HoQuPlatformIdentificationAddedIterator is returned from FilterIdentificationAdded and is used to iterate over the raw logs and unpacked data for IdentificationAdded events raised by the HoQuPlatform contract.
type HoQuPlatformIdentificationAddedIterator struct {
	Event *HoQuPlatformIdentificationAdded // Event containing the contract specifics and raw log

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
func (it *HoQuPlatformIdentificationAddedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(HoQuPlatformIdentificationAdded)
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
		it.Event = new(HoQuPlatformIdentificationAdded)
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
func (it *HoQuPlatformIdentificationAddedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *HoQuPlatformIdentificationAddedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// HoQuPlatformIdentificationAdded represents a IdentificationAdded event raised by the HoQuPlatform contract.
type HoQuPlatformIdentificationAdded struct {
	OwnerAddress common.Address
	Id           [16]byte
	UserId       [16]byte
	Name         string
	Raw          types.Log // Blockchain specific contextual infos
}

// FilterIdentificationAdded is a free log retrieval operation binding the contract event 0x8c03773782e9ed7fe9111421165a6cf12ce5855a4181cd1a8166d47d719ac078.
//
// Solidity: event IdentificationAdded(ownerAddress indexed address, id bytes16, userId bytes16, name string)
func (_HoQuPlatform *HoQuPlatformFilterer) FilterIdentificationAdded(opts *bind.FilterOpts, ownerAddress []common.Address) (*HoQuPlatformIdentificationAddedIterator, error) {

	var ownerAddressRule []interface{}
	for _, ownerAddressItem := range ownerAddress {
		ownerAddressRule = append(ownerAddressRule, ownerAddressItem)
	}

	logs, sub, err := _HoQuPlatform.contract.FilterLogs(opts, "IdentificationAdded", ownerAddressRule)
	if err != nil {
		return nil, err
	}
	return &HoQuPlatformIdentificationAddedIterator{contract: _HoQuPlatform.contract, event: "IdentificationAdded", logs: logs, sub: sub}, nil
}

// WatchIdentificationAdded is a free log subscription operation binding the contract event 0x8c03773782e9ed7fe9111421165a6cf12ce5855a4181cd1a8166d47d719ac078.
//
// Solidity: event IdentificationAdded(ownerAddress indexed address, id bytes16, userId bytes16, name string)
func (_HoQuPlatform *HoQuPlatformFilterer) WatchIdentificationAdded(opts *bind.WatchOpts, sink chan<- *HoQuPlatformIdentificationAdded, ownerAddress []common.Address) (event.Subscription, error) {

	var ownerAddressRule []interface{}
	for _, ownerAddressItem := range ownerAddress {
		ownerAddressRule = append(ownerAddressRule, ownerAddressItem)
	}

	logs, sub, err := _HoQuPlatform.contract.WatchLogs(opts, "IdentificationAdded", ownerAddressRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(HoQuPlatformIdentificationAdded)
				if err := _HoQuPlatform.contract.UnpackLog(event, "IdentificationAdded", log); err != nil {
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

// HoQuPlatformKycReportAddedIterator is returned from FilterKycReportAdded and is used to iterate over the raw logs and unpacked data for KycReportAdded events raised by the HoQuPlatform contract.
type HoQuPlatformKycReportAddedIterator struct {
	Event *HoQuPlatformKycReportAdded // Event containing the contract specifics and raw log

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
func (it *HoQuPlatformKycReportAddedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(HoQuPlatformKycReportAdded)
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
		it.Event = new(HoQuPlatformKycReportAdded)
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
func (it *HoQuPlatformKycReportAddedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *HoQuPlatformKycReportAddedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// HoQuPlatformKycReportAdded represents a KycReportAdded event raised by the HoQuPlatform contract.
type HoQuPlatformKycReportAdded struct {
	OwnerAddress common.Address
	KycLevel     uint8
	Id           [16]byte
	UserId       [16]byte
	Raw          types.Log // Blockchain specific contextual infos
}

// FilterKycReportAdded is a free log retrieval operation binding the contract event 0x09d7204a29a3e39bb12fb2c813212dd90eb69edaf26f867689c3e3a2eb6eb7b9.
//
// Solidity: event KycReportAdded(ownerAddress indexed address, kycLevel uint8, id bytes16, userId bytes16)
func (_HoQuPlatform *HoQuPlatformFilterer) FilterKycReportAdded(opts *bind.FilterOpts, ownerAddress []common.Address) (*HoQuPlatformKycReportAddedIterator, error) {

	var ownerAddressRule []interface{}
	for _, ownerAddressItem := range ownerAddress {
		ownerAddressRule = append(ownerAddressRule, ownerAddressItem)
	}

	logs, sub, err := _HoQuPlatform.contract.FilterLogs(opts, "KycReportAdded", ownerAddressRule)
	if err != nil {
		return nil, err
	}
	return &HoQuPlatformKycReportAddedIterator{contract: _HoQuPlatform.contract, event: "KycReportAdded", logs: logs, sub: sub}, nil
}

// WatchKycReportAdded is a free log subscription operation binding the contract event 0x09d7204a29a3e39bb12fb2c813212dd90eb69edaf26f867689c3e3a2eb6eb7b9.
//
// Solidity: event KycReportAdded(ownerAddress indexed address, kycLevel uint8, id bytes16, userId bytes16)
func (_HoQuPlatform *HoQuPlatformFilterer) WatchKycReportAdded(opts *bind.WatchOpts, sink chan<- *HoQuPlatformKycReportAdded, ownerAddress []common.Address) (event.Subscription, error) {

	var ownerAddressRule []interface{}
	for _, ownerAddressItem := range ownerAddress {
		ownerAddressRule = append(ownerAddressRule, ownerAddressItem)
	}

	logs, sub, err := _HoQuPlatform.contract.WatchLogs(opts, "KycReportAdded", ownerAddressRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(HoQuPlatformKycReportAdded)
				if err := _HoQuPlatform.contract.UnpackLog(event, "KycReportAdded", log); err != nil {
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

// HoQuPlatformLeadAddedIterator is returned from FilterLeadAdded and is used to iterate over the raw logs and unpacked data for LeadAdded events raised by the HoQuPlatform contract.
type HoQuPlatformLeadAddedIterator struct {
	Event *HoQuPlatformLeadAdded // Event containing the contract specifics and raw log

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
func (it *HoQuPlatformLeadAddedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(HoQuPlatformLeadAdded)
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
		it.Event = new(HoQuPlatformLeadAdded)
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
func (it *HoQuPlatformLeadAddedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *HoQuPlatformLeadAddedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// HoQuPlatformLeadAdded represents a LeadAdded event raised by the HoQuPlatform contract.
type HoQuPlatformLeadAdded struct {
	ContractAddress common.Address
	AdCampaignId    [16]byte
	Id              [16]byte
	Raw             types.Log // Blockchain specific contextual infos
}

// FilterLeadAdded is a free log retrieval operation binding the contract event 0x4d5402ce2635084c7def60402936ba18370a5a5d055ecdddc10721afa736ce42.
//
// Solidity: event LeadAdded(contractAddress indexed address, adCampaignId bytes16, id bytes16)
func (_HoQuPlatform *HoQuPlatformFilterer) FilterLeadAdded(opts *bind.FilterOpts, contractAddress []common.Address) (*HoQuPlatformLeadAddedIterator, error) {

	var contractAddressRule []interface{}
	for _, contractAddressItem := range contractAddress {
		contractAddressRule = append(contractAddressRule, contractAddressItem)
	}

	logs, sub, err := _HoQuPlatform.contract.FilterLogs(opts, "LeadAdded", contractAddressRule)
	if err != nil {
		return nil, err
	}
	return &HoQuPlatformLeadAddedIterator{contract: _HoQuPlatform.contract, event: "LeadAdded", logs: logs, sub: sub}, nil
}

// WatchLeadAdded is a free log subscription operation binding the contract event 0x4d5402ce2635084c7def60402936ba18370a5a5d055ecdddc10721afa736ce42.
//
// Solidity: event LeadAdded(contractAddress indexed address, adCampaignId bytes16, id bytes16)
func (_HoQuPlatform *HoQuPlatformFilterer) WatchLeadAdded(opts *bind.WatchOpts, sink chan<- *HoQuPlatformLeadAdded, contractAddress []common.Address) (event.Subscription, error) {

	var contractAddressRule []interface{}
	for _, contractAddressItem := range contractAddress {
		contractAddressRule = append(contractAddressRule, contractAddressItem)
	}

	logs, sub, err := _HoQuPlatform.contract.WatchLogs(opts, "LeadAdded", contractAddressRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(HoQuPlatformLeadAdded)
				if err := _HoQuPlatform.contract.UnpackLog(event, "LeadAdded", log); err != nil {
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

// HoQuPlatformLeadDataUrlChangedIterator is returned from FilterLeadDataUrlChanged and is used to iterate over the raw logs and unpacked data for LeadDataUrlChanged events raised by the HoQuPlatform contract.
type HoQuPlatformLeadDataUrlChangedIterator struct {
	Event *HoQuPlatformLeadDataUrlChanged // Event containing the contract specifics and raw log

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
func (it *HoQuPlatformLeadDataUrlChangedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(HoQuPlatformLeadDataUrlChanged)
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
		it.Event = new(HoQuPlatformLeadDataUrlChanged)
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
func (it *HoQuPlatformLeadDataUrlChangedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *HoQuPlatformLeadDataUrlChangedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// HoQuPlatformLeadDataUrlChanged represents a LeadDataUrlChanged event raised by the HoQuPlatform contract.
type HoQuPlatformLeadDataUrlChanged struct {
	ContractAddress common.Address
	AdCampaignId    [16]byte
	Id              [16]byte
	DataUrl         string
	Raw             types.Log // Blockchain specific contextual infos
}

// FilterLeadDataUrlChanged is a free log retrieval operation binding the contract event 0x030cd1b359a7520d9c7419a712a01660bf8f8a953e84d3c917f8f87a68214a41.
//
// Solidity: event LeadDataUrlChanged(contractAddress indexed address, adCampaignId bytes16, id bytes16, dataUrl string)
func (_HoQuPlatform *HoQuPlatformFilterer) FilterLeadDataUrlChanged(opts *bind.FilterOpts, contractAddress []common.Address) (*HoQuPlatformLeadDataUrlChangedIterator, error) {

	var contractAddressRule []interface{}
	for _, contractAddressItem := range contractAddress {
		contractAddressRule = append(contractAddressRule, contractAddressItem)
	}

	logs, sub, err := _HoQuPlatform.contract.FilterLogs(opts, "LeadDataUrlChanged", contractAddressRule)
	if err != nil {
		return nil, err
	}
	return &HoQuPlatformLeadDataUrlChangedIterator{contract: _HoQuPlatform.contract, event: "LeadDataUrlChanged", logs: logs, sub: sub}, nil
}

// WatchLeadDataUrlChanged is a free log subscription operation binding the contract event 0x030cd1b359a7520d9c7419a712a01660bf8f8a953e84d3c917f8f87a68214a41.
//
// Solidity: event LeadDataUrlChanged(contractAddress indexed address, adCampaignId bytes16, id bytes16, dataUrl string)
func (_HoQuPlatform *HoQuPlatformFilterer) WatchLeadDataUrlChanged(opts *bind.WatchOpts, sink chan<- *HoQuPlatformLeadDataUrlChanged, contractAddress []common.Address) (event.Subscription, error) {

	var contractAddressRule []interface{}
	for _, contractAddressItem := range contractAddress {
		contractAddressRule = append(contractAddressRule, contractAddressItem)
	}

	logs, sub, err := _HoQuPlatform.contract.WatchLogs(opts, "LeadDataUrlChanged", contractAddressRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(HoQuPlatformLeadDataUrlChanged)
				if err := _HoQuPlatform.contract.UnpackLog(event, "LeadDataUrlChanged", log); err != nil {
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

// HoQuPlatformLeadPriceChangedIterator is returned from FilterLeadPriceChanged and is used to iterate over the raw logs and unpacked data for LeadPriceChanged events raised by the HoQuPlatform contract.
type HoQuPlatformLeadPriceChangedIterator struct {
	Event *HoQuPlatformLeadPriceChanged // Event containing the contract specifics and raw log

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
func (it *HoQuPlatformLeadPriceChangedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(HoQuPlatformLeadPriceChanged)
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
		it.Event = new(HoQuPlatformLeadPriceChanged)
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
func (it *HoQuPlatformLeadPriceChangedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *HoQuPlatformLeadPriceChangedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// HoQuPlatformLeadPriceChanged represents a LeadPriceChanged event raised by the HoQuPlatform contract.
type HoQuPlatformLeadPriceChanged struct {
	ContractAddress common.Address
	AdCampaignId    [16]byte
	Id              [16]byte
	Price           *big.Int
	Raw             types.Log // Blockchain specific contextual infos
}

// FilterLeadPriceChanged is a free log retrieval operation binding the contract event 0x10d48255e0ae3847d38a5e163a9d9afb2964ac9c892d76cea6572eca867daa39.
//
// Solidity: event LeadPriceChanged(contractAddress indexed address, adCampaignId bytes16, id bytes16, price uint256)
func (_HoQuPlatform *HoQuPlatformFilterer) FilterLeadPriceChanged(opts *bind.FilterOpts, contractAddress []common.Address) (*HoQuPlatformLeadPriceChangedIterator, error) {

	var contractAddressRule []interface{}
	for _, contractAddressItem := range contractAddress {
		contractAddressRule = append(contractAddressRule, contractAddressItem)
	}

	logs, sub, err := _HoQuPlatform.contract.FilterLogs(opts, "LeadPriceChanged", contractAddressRule)
	if err != nil {
		return nil, err
	}
	return &HoQuPlatformLeadPriceChangedIterator{contract: _HoQuPlatform.contract, event: "LeadPriceChanged", logs: logs, sub: sub}, nil
}

// WatchLeadPriceChanged is a free log subscription operation binding the contract event 0x10d48255e0ae3847d38a5e163a9d9afb2964ac9c892d76cea6572eca867daa39.
//
// Solidity: event LeadPriceChanged(contractAddress indexed address, adCampaignId bytes16, id bytes16, price uint256)
func (_HoQuPlatform *HoQuPlatformFilterer) WatchLeadPriceChanged(opts *bind.WatchOpts, sink chan<- *HoQuPlatformLeadPriceChanged, contractAddress []common.Address) (event.Subscription, error) {

	var contractAddressRule []interface{}
	for _, contractAddressItem := range contractAddress {
		contractAddressRule = append(contractAddressRule, contractAddressItem)
	}

	logs, sub, err := _HoQuPlatform.contract.WatchLogs(opts, "LeadPriceChanged", contractAddressRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(HoQuPlatformLeadPriceChanged)
				if err := _HoQuPlatform.contract.UnpackLog(event, "LeadPriceChanged", log); err != nil {
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

// HoQuPlatformLeadStatusChangedIterator is returned from FilterLeadStatusChanged and is used to iterate over the raw logs and unpacked data for LeadStatusChanged events raised by the HoQuPlatform contract.
type HoQuPlatformLeadStatusChangedIterator struct {
	Event *HoQuPlatformLeadStatusChanged // Event containing the contract specifics and raw log

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
func (it *HoQuPlatformLeadStatusChangedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(HoQuPlatformLeadStatusChanged)
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
		it.Event = new(HoQuPlatformLeadStatusChanged)
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
func (it *HoQuPlatformLeadStatusChangedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *HoQuPlatformLeadStatusChangedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// HoQuPlatformLeadStatusChanged represents a LeadStatusChanged event raised by the HoQuPlatform contract.
type HoQuPlatformLeadStatusChanged struct {
	ContractAddress common.Address
	AdCampaignId    [16]byte
	Id              [16]byte
	Status          uint8
	Raw             types.Log // Blockchain specific contextual infos
}

// FilterLeadStatusChanged is a free log retrieval operation binding the contract event 0xcde6c64834d04acd2555be68ee4f9a2b92fa8c30fc842f42f5ed581214c8dfff.
//
// Solidity: event LeadStatusChanged(contractAddress indexed address, adCampaignId bytes16, id bytes16, status uint8)
func (_HoQuPlatform *HoQuPlatformFilterer) FilterLeadStatusChanged(opts *bind.FilterOpts, contractAddress []common.Address) (*HoQuPlatformLeadStatusChangedIterator, error) {

	var contractAddressRule []interface{}
	for _, contractAddressItem := range contractAddress {
		contractAddressRule = append(contractAddressRule, contractAddressItem)
	}

	logs, sub, err := _HoQuPlatform.contract.FilterLogs(opts, "LeadStatusChanged", contractAddressRule)
	if err != nil {
		return nil, err
	}
	return &HoQuPlatformLeadStatusChangedIterator{contract: _HoQuPlatform.contract, event: "LeadStatusChanged", logs: logs, sub: sub}, nil
}

// WatchLeadStatusChanged is a free log subscription operation binding the contract event 0xcde6c64834d04acd2555be68ee4f9a2b92fa8c30fc842f42f5ed581214c8dfff.
//
// Solidity: event LeadStatusChanged(contractAddress indexed address, adCampaignId bytes16, id bytes16, status uint8)
func (_HoQuPlatform *HoQuPlatformFilterer) WatchLeadStatusChanged(opts *bind.WatchOpts, sink chan<- *HoQuPlatformLeadStatusChanged, contractAddress []common.Address) (event.Subscription, error) {

	var contractAddressRule []interface{}
	for _, contractAddressItem := range contractAddress {
		contractAddressRule = append(contractAddressRule, contractAddressItem)
	}

	logs, sub, err := _HoQuPlatform.contract.WatchLogs(opts, "LeadStatusChanged", contractAddressRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(HoQuPlatformLeadStatusChanged)
				if err := _HoQuPlatform.contract.UnpackLog(event, "LeadStatusChanged", log); err != nil {
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

// HoQuPlatformLeadTransactedIterator is returned from FilterLeadTransacted and is used to iterate over the raw logs and unpacked data for LeadTransacted events raised by the HoQuPlatform contract.
type HoQuPlatformLeadTransactedIterator struct {
	Event *HoQuPlatformLeadTransacted // Event containing the contract specifics and raw log

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
func (it *HoQuPlatformLeadTransactedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(HoQuPlatformLeadTransacted)
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
		it.Event = new(HoQuPlatformLeadTransacted)
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
func (it *HoQuPlatformLeadTransactedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *HoQuPlatformLeadTransactedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// HoQuPlatformLeadTransacted represents a LeadTransacted event raised by the HoQuPlatform contract.
type HoQuPlatformLeadTransacted struct {
	ContractAddress common.Address
	AdCampaignId    [16]byte
	Id              [16]byte
	Raw             types.Log // Blockchain specific contextual infos
}

// FilterLeadTransacted is a free log retrieval operation binding the contract event 0x3d86ba643ff86fc12dc0ce1d161b9407a79cb7131f80001f21c054658de5e1e8.
//
// Solidity: event LeadTransacted(contractAddress indexed address, adCampaignId bytes16, id bytes16)
func (_HoQuPlatform *HoQuPlatformFilterer) FilterLeadTransacted(opts *bind.FilterOpts, contractAddress []common.Address) (*HoQuPlatformLeadTransactedIterator, error) {

	var contractAddressRule []interface{}
	for _, contractAddressItem := range contractAddress {
		contractAddressRule = append(contractAddressRule, contractAddressItem)
	}

	logs, sub, err := _HoQuPlatform.contract.FilterLogs(opts, "LeadTransacted", contractAddressRule)
	if err != nil {
		return nil, err
	}
	return &HoQuPlatformLeadTransactedIterator{contract: _HoQuPlatform.contract, event: "LeadTransacted", logs: logs, sub: sub}, nil
}

// WatchLeadTransacted is a free log subscription operation binding the contract event 0x3d86ba643ff86fc12dc0ce1d161b9407a79cb7131f80001f21c054658de5e1e8.
//
// Solidity: event LeadTransacted(contractAddress indexed address, adCampaignId bytes16, id bytes16)
func (_HoQuPlatform *HoQuPlatformFilterer) WatchLeadTransacted(opts *bind.WatchOpts, sink chan<- *HoQuPlatformLeadTransacted, contractAddress []common.Address) (event.Subscription, error) {

	var contractAddressRule []interface{}
	for _, contractAddressItem := range contractAddress {
		contractAddressRule = append(contractAddressRule, contractAddressItem)
	}

	logs, sub, err := _HoQuPlatform.contract.WatchLogs(opts, "LeadTransacted", contractAddressRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(HoQuPlatformLeadTransacted)
				if err := _HoQuPlatform.contract.UnpackLog(event, "LeadTransacted", log); err != nil {
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

// HoQuPlatformNetworkRegisteredIterator is returned from FilterNetworkRegistered and is used to iterate over the raw logs and unpacked data for NetworkRegistered events raised by the HoQuPlatform contract.
type HoQuPlatformNetworkRegisteredIterator struct {
	Event *HoQuPlatformNetworkRegistered // Event containing the contract specifics and raw log

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
func (it *HoQuPlatformNetworkRegisteredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(HoQuPlatformNetworkRegistered)
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
		it.Event = new(HoQuPlatformNetworkRegistered)
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
func (it *HoQuPlatformNetworkRegisteredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *HoQuPlatformNetworkRegisteredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// HoQuPlatformNetworkRegistered represents a NetworkRegistered event raised by the HoQuPlatform contract.
type HoQuPlatformNetworkRegistered struct {
	OwnerAddress common.Address
	Id           [16]byte
	Name         string
	Raw          types.Log // Blockchain specific contextual infos
}

// FilterNetworkRegistered is a free log retrieval operation binding the contract event 0x592e83ede8c846dffd6e4edcbe77983b153bec820ac697aa34a236147b1341e0.
//
// Solidity: event NetworkRegistered(ownerAddress indexed address, id bytes16, name string)
func (_HoQuPlatform *HoQuPlatformFilterer) FilterNetworkRegistered(opts *bind.FilterOpts, ownerAddress []common.Address) (*HoQuPlatformNetworkRegisteredIterator, error) {

	var ownerAddressRule []interface{}
	for _, ownerAddressItem := range ownerAddress {
		ownerAddressRule = append(ownerAddressRule, ownerAddressItem)
	}

	logs, sub, err := _HoQuPlatform.contract.FilterLogs(opts, "NetworkRegistered", ownerAddressRule)
	if err != nil {
		return nil, err
	}
	return &HoQuPlatformNetworkRegisteredIterator{contract: _HoQuPlatform.contract, event: "NetworkRegistered", logs: logs, sub: sub}, nil
}

// WatchNetworkRegistered is a free log subscription operation binding the contract event 0x592e83ede8c846dffd6e4edcbe77983b153bec820ac697aa34a236147b1341e0.
//
// Solidity: event NetworkRegistered(ownerAddress indexed address, id bytes16, name string)
func (_HoQuPlatform *HoQuPlatformFilterer) WatchNetworkRegistered(opts *bind.WatchOpts, sink chan<- *HoQuPlatformNetworkRegistered, ownerAddress []common.Address) (event.Subscription, error) {

	var ownerAddressRule []interface{}
	for _, ownerAddressItem := range ownerAddress {
		ownerAddressRule = append(ownerAddressRule, ownerAddressItem)
	}

	logs, sub, err := _HoQuPlatform.contract.WatchLogs(opts, "NetworkRegistered", ownerAddressRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(HoQuPlatformNetworkRegistered)
				if err := _HoQuPlatform.contract.UnpackLog(event, "NetworkRegistered", log); err != nil {
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

// HoQuPlatformNetworkStatusChangedIterator is returned from FilterNetworkStatusChanged and is used to iterate over the raw logs and unpacked data for NetworkStatusChanged events raised by the HoQuPlatform contract.
type HoQuPlatformNetworkStatusChangedIterator struct {
	Event *HoQuPlatformNetworkStatusChanged // Event containing the contract specifics and raw log

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
func (it *HoQuPlatformNetworkStatusChangedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(HoQuPlatformNetworkStatusChanged)
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
		it.Event = new(HoQuPlatformNetworkStatusChanged)
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
func (it *HoQuPlatformNetworkStatusChangedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *HoQuPlatformNetworkStatusChangedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// HoQuPlatformNetworkStatusChanged represents a NetworkStatusChanged event raised by the HoQuPlatform contract.
type HoQuPlatformNetworkStatusChanged struct {
	OwnerAddress common.Address
	Id           [16]byte
	Status       uint8
	Raw          types.Log // Blockchain specific contextual infos
}

// FilterNetworkStatusChanged is a free log retrieval operation binding the contract event 0xce8cdf87291e737b80351c93570f8555ca6a0cc77c3fb29176e56af69dccdd84.
//
// Solidity: event NetworkStatusChanged(ownerAddress indexed address, id bytes16, status uint8)
func (_HoQuPlatform *HoQuPlatformFilterer) FilterNetworkStatusChanged(opts *bind.FilterOpts, ownerAddress []common.Address) (*HoQuPlatformNetworkStatusChangedIterator, error) {

	var ownerAddressRule []interface{}
	for _, ownerAddressItem := range ownerAddress {
		ownerAddressRule = append(ownerAddressRule, ownerAddressItem)
	}

	logs, sub, err := _HoQuPlatform.contract.FilterLogs(opts, "NetworkStatusChanged", ownerAddressRule)
	if err != nil {
		return nil, err
	}
	return &HoQuPlatformNetworkStatusChangedIterator{contract: _HoQuPlatform.contract, event: "NetworkStatusChanged", logs: logs, sub: sub}, nil
}

// WatchNetworkStatusChanged is a free log subscription operation binding the contract event 0xce8cdf87291e737b80351c93570f8555ca6a0cc77c3fb29176e56af69dccdd84.
//
// Solidity: event NetworkStatusChanged(ownerAddress indexed address, id bytes16, status uint8)
func (_HoQuPlatform *HoQuPlatformFilterer) WatchNetworkStatusChanged(opts *bind.WatchOpts, sink chan<- *HoQuPlatformNetworkStatusChanged, ownerAddress []common.Address) (event.Subscription, error) {

	var ownerAddressRule []interface{}
	for _, ownerAddressItem := range ownerAddress {
		ownerAddressRule = append(ownerAddressRule, ownerAddressItem)
	}

	logs, sub, err := _HoQuPlatform.contract.WatchLogs(opts, "NetworkStatusChanged", ownerAddressRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(HoQuPlatformNetworkStatusChanged)
				if err := _HoQuPlatform.contract.UnpackLog(event, "NetworkStatusChanged", log); err != nil {
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

// HoQuPlatformOfferAddedIterator is returned from FilterOfferAdded and is used to iterate over the raw logs and unpacked data for OfferAdded events raised by the HoQuPlatform contract.
type HoQuPlatformOfferAddedIterator struct {
	Event *HoQuPlatformOfferAdded // Event containing the contract specifics and raw log

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
func (it *HoQuPlatformOfferAddedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(HoQuPlatformOfferAdded)
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
		it.Event = new(HoQuPlatformOfferAdded)
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
func (it *HoQuPlatformOfferAddedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *HoQuPlatformOfferAddedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// HoQuPlatformOfferAdded represents a OfferAdded event raised by the HoQuPlatform contract.
type HoQuPlatformOfferAdded struct {
	OwnerAddress common.Address
	Id           [16]byte
	Name         string
	Raw          types.Log // Blockchain specific contextual infos
}

// FilterOfferAdded is a free log retrieval operation binding the contract event 0x5c9100ad970db426d43d47b3a9639dc42d9732871d99213dcb4984f7e71509d0.
//
// Solidity: event OfferAdded(ownerAddress indexed address, id bytes16, name string)
func (_HoQuPlatform *HoQuPlatformFilterer) FilterOfferAdded(opts *bind.FilterOpts, ownerAddress []common.Address) (*HoQuPlatformOfferAddedIterator, error) {

	var ownerAddressRule []interface{}
	for _, ownerAddressItem := range ownerAddress {
		ownerAddressRule = append(ownerAddressRule, ownerAddressItem)
	}

	logs, sub, err := _HoQuPlatform.contract.FilterLogs(opts, "OfferAdded", ownerAddressRule)
	if err != nil {
		return nil, err
	}
	return &HoQuPlatformOfferAddedIterator{contract: _HoQuPlatform.contract, event: "OfferAdded", logs: logs, sub: sub}, nil
}

// WatchOfferAdded is a free log subscription operation binding the contract event 0x5c9100ad970db426d43d47b3a9639dc42d9732871d99213dcb4984f7e71509d0.
//
// Solidity: event OfferAdded(ownerAddress indexed address, id bytes16, name string)
func (_HoQuPlatform *HoQuPlatformFilterer) WatchOfferAdded(opts *bind.WatchOpts, sink chan<- *HoQuPlatformOfferAdded, ownerAddress []common.Address) (event.Subscription, error) {

	var ownerAddressRule []interface{}
	for _, ownerAddressItem := range ownerAddress {
		ownerAddressRule = append(ownerAddressRule, ownerAddressItem)
	}

	logs, sub, err := _HoQuPlatform.contract.WatchLogs(opts, "OfferAdded", ownerAddressRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(HoQuPlatformOfferAdded)
				if err := _HoQuPlatform.contract.UnpackLog(event, "OfferAdded", log); err != nil {
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

// HoQuPlatformOfferStatusChangedIterator is returned from FilterOfferStatusChanged and is used to iterate over the raw logs and unpacked data for OfferStatusChanged events raised by the HoQuPlatform contract.
type HoQuPlatformOfferStatusChangedIterator struct {
	Event *HoQuPlatformOfferStatusChanged // Event containing the contract specifics and raw log

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
func (it *HoQuPlatformOfferStatusChangedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(HoQuPlatformOfferStatusChanged)
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
		it.Event = new(HoQuPlatformOfferStatusChanged)
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
func (it *HoQuPlatformOfferStatusChangedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *HoQuPlatformOfferStatusChangedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// HoQuPlatformOfferStatusChanged represents a OfferStatusChanged event raised by the HoQuPlatform contract.
type HoQuPlatformOfferStatusChanged struct {
	OwnerAddress common.Address
	Id           [16]byte
	Status       uint8
	Raw          types.Log // Blockchain specific contextual infos
}

// FilterOfferStatusChanged is a free log retrieval operation binding the contract event 0xf9ae5f2bc88e6348608d627c5226159efd8467eb4ac3c635cdb648266e02cb12.
//
// Solidity: event OfferStatusChanged(ownerAddress indexed address, id bytes16, status uint8)
func (_HoQuPlatform *HoQuPlatformFilterer) FilterOfferStatusChanged(opts *bind.FilterOpts, ownerAddress []common.Address) (*HoQuPlatformOfferStatusChangedIterator, error) {

	var ownerAddressRule []interface{}
	for _, ownerAddressItem := range ownerAddress {
		ownerAddressRule = append(ownerAddressRule, ownerAddressItem)
	}

	logs, sub, err := _HoQuPlatform.contract.FilterLogs(opts, "OfferStatusChanged", ownerAddressRule)
	if err != nil {
		return nil, err
	}
	return &HoQuPlatformOfferStatusChangedIterator{contract: _HoQuPlatform.contract, event: "OfferStatusChanged", logs: logs, sub: sub}, nil
}

// WatchOfferStatusChanged is a free log subscription operation binding the contract event 0xf9ae5f2bc88e6348608d627c5226159efd8467eb4ac3c635cdb648266e02cb12.
//
// Solidity: event OfferStatusChanged(ownerAddress indexed address, id bytes16, status uint8)
func (_HoQuPlatform *HoQuPlatformFilterer) WatchOfferStatusChanged(opts *bind.WatchOpts, sink chan<- *HoQuPlatformOfferStatusChanged, ownerAddress []common.Address) (event.Subscription, error) {

	var ownerAddressRule []interface{}
	for _, ownerAddressItem := range ownerAddress {
		ownerAddressRule = append(ownerAddressRule, ownerAddressItem)
	}

	logs, sub, err := _HoQuPlatform.contract.WatchLogs(opts, "OfferStatusChanged", ownerAddressRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(HoQuPlatformOfferStatusChanged)
				if err := _HoQuPlatform.contract.UnpackLog(event, "OfferStatusChanged", log); err != nil {
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

// HoQuPlatformOfferTariffGroupAddedIterator is returned from FilterOfferTariffGroupAdded and is used to iterate over the raw logs and unpacked data for OfferTariffGroupAdded events raised by the HoQuPlatform contract.
type HoQuPlatformOfferTariffGroupAddedIterator struct {
	Event *HoQuPlatformOfferTariffGroupAdded // Event containing the contract specifics and raw log

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
func (it *HoQuPlatformOfferTariffGroupAddedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(HoQuPlatformOfferTariffGroupAdded)
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
		it.Event = new(HoQuPlatformOfferTariffGroupAdded)
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
func (it *HoQuPlatformOfferTariffGroupAddedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *HoQuPlatformOfferTariffGroupAddedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// HoQuPlatformOfferTariffGroupAdded represents a OfferTariffGroupAdded event raised by the HoQuPlatform contract.
type HoQuPlatformOfferTariffGroupAdded struct {
	OwnerAddress common.Address
	Id           [16]byte
	TariffId     [16]byte
	Raw          types.Log // Blockchain specific contextual infos
}

// FilterOfferTariffGroupAdded is a free log retrieval operation binding the contract event 0x695164079ea85cb7933398c55f763b3dd70fe1cfe293a3b90aa22f9cf6270dab.
//
// Solidity: event OfferTariffGroupAdded(ownerAddress indexed address, id bytes16, tariff_id bytes16)
func (_HoQuPlatform *HoQuPlatformFilterer) FilterOfferTariffGroupAdded(opts *bind.FilterOpts, ownerAddress []common.Address) (*HoQuPlatformOfferTariffGroupAddedIterator, error) {

	var ownerAddressRule []interface{}
	for _, ownerAddressItem := range ownerAddress {
		ownerAddressRule = append(ownerAddressRule, ownerAddressItem)
	}

	logs, sub, err := _HoQuPlatform.contract.FilterLogs(opts, "OfferTariffGroupAdded", ownerAddressRule)
	if err != nil {
		return nil, err
	}
	return &HoQuPlatformOfferTariffGroupAddedIterator{contract: _HoQuPlatform.contract, event: "OfferTariffGroupAdded", logs: logs, sub: sub}, nil
}

// WatchOfferTariffGroupAdded is a free log subscription operation binding the contract event 0x695164079ea85cb7933398c55f763b3dd70fe1cfe293a3b90aa22f9cf6270dab.
//
// Solidity: event OfferTariffGroupAdded(ownerAddress indexed address, id bytes16, tariff_id bytes16)
func (_HoQuPlatform *HoQuPlatformFilterer) WatchOfferTariffGroupAdded(opts *bind.WatchOpts, sink chan<- *HoQuPlatformOfferTariffGroupAdded, ownerAddress []common.Address) (event.Subscription, error) {

	var ownerAddressRule []interface{}
	for _, ownerAddressItem := range ownerAddress {
		ownerAddressRule = append(ownerAddressRule, ownerAddressItem)
	}

	logs, sub, err := _HoQuPlatform.contract.WatchLogs(opts, "OfferTariffGroupAdded", ownerAddressRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(HoQuPlatformOfferTariffGroupAdded)
				if err := _HoQuPlatform.contract.UnpackLog(event, "OfferTariffGroupAdded", log); err != nil {
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

// HoQuPlatformTariffAddedIterator is returned from FilterTariffAdded and is used to iterate over the raw logs and unpacked data for TariffAdded events raised by the HoQuPlatform contract.
type HoQuPlatformTariffAddedIterator struct {
	Event *HoQuPlatformTariffAdded // Event containing the contract specifics and raw log

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
func (it *HoQuPlatformTariffAddedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(HoQuPlatformTariffAdded)
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
		it.Event = new(HoQuPlatformTariffAdded)
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
func (it *HoQuPlatformTariffAddedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *HoQuPlatformTariffAddedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// HoQuPlatformTariffAdded represents a TariffAdded event raised by the HoQuPlatform contract.
type HoQuPlatformTariffAdded struct {
	OwnerAddress common.Address
	Id           [16]byte
	Name         string
	Raw          types.Log // Blockchain specific contextual infos
}

// FilterTariffAdded is a free log retrieval operation binding the contract event 0x4ad20f56c3c96698f2e84d660e8c5e39341fd79929743804d261beb65ffc5b47.
//
// Solidity: event TariffAdded(ownerAddress indexed address, id bytes16, name string)
func (_HoQuPlatform *HoQuPlatformFilterer) FilterTariffAdded(opts *bind.FilterOpts, ownerAddress []common.Address) (*HoQuPlatformTariffAddedIterator, error) {

	var ownerAddressRule []interface{}
	for _, ownerAddressItem := range ownerAddress {
		ownerAddressRule = append(ownerAddressRule, ownerAddressItem)
	}

	logs, sub, err := _HoQuPlatform.contract.FilterLogs(opts, "TariffAdded", ownerAddressRule)
	if err != nil {
		return nil, err
	}
	return &HoQuPlatformTariffAddedIterator{contract: _HoQuPlatform.contract, event: "TariffAdded", logs: logs, sub: sub}, nil
}

// WatchTariffAdded is a free log subscription operation binding the contract event 0x4ad20f56c3c96698f2e84d660e8c5e39341fd79929743804d261beb65ffc5b47.
//
// Solidity: event TariffAdded(ownerAddress indexed address, id bytes16, name string)
func (_HoQuPlatform *HoQuPlatformFilterer) WatchTariffAdded(opts *bind.WatchOpts, sink chan<- *HoQuPlatformTariffAdded, ownerAddress []common.Address) (event.Subscription, error) {

	var ownerAddressRule []interface{}
	for _, ownerAddressItem := range ownerAddress {
		ownerAddressRule = append(ownerAddressRule, ownerAddressItem)
	}

	logs, sub, err := _HoQuPlatform.contract.WatchLogs(opts, "TariffAdded", ownerAddressRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(HoQuPlatformTariffAdded)
				if err := _HoQuPlatform.contract.UnpackLog(event, "TariffAdded", log); err != nil {
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

// HoQuPlatformTariffGroupAddedIterator is returned from FilterTariffGroupAdded and is used to iterate over the raw logs and unpacked data for TariffGroupAdded events raised by the HoQuPlatform contract.
type HoQuPlatformTariffGroupAddedIterator struct {
	Event *HoQuPlatformTariffGroupAdded // Event containing the contract specifics and raw log

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
func (it *HoQuPlatformTariffGroupAddedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(HoQuPlatformTariffGroupAdded)
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
		it.Event = new(HoQuPlatformTariffGroupAdded)
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
func (it *HoQuPlatformTariffGroupAddedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *HoQuPlatformTariffGroupAddedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// HoQuPlatformTariffGroupAdded represents a TariffGroupAdded event raised by the HoQuPlatform contract.
type HoQuPlatformTariffGroupAdded struct {
	OwnerAddress common.Address
	Id           [16]byte
	Name         string
	Raw          types.Log // Blockchain specific contextual infos
}

// FilterTariffGroupAdded is a free log retrieval operation binding the contract event 0x0d945b4e88c22df3f6d39d8866fac2dd4964090d9d9c04fc74f33c5695e561df.
//
// Solidity: event TariffGroupAdded(ownerAddress indexed address, id bytes16, name string)
func (_HoQuPlatform *HoQuPlatformFilterer) FilterTariffGroupAdded(opts *bind.FilterOpts, ownerAddress []common.Address) (*HoQuPlatformTariffGroupAddedIterator, error) {

	var ownerAddressRule []interface{}
	for _, ownerAddressItem := range ownerAddress {
		ownerAddressRule = append(ownerAddressRule, ownerAddressItem)
	}

	logs, sub, err := _HoQuPlatform.contract.FilterLogs(opts, "TariffGroupAdded", ownerAddressRule)
	if err != nil {
		return nil, err
	}
	return &HoQuPlatformTariffGroupAddedIterator{contract: _HoQuPlatform.contract, event: "TariffGroupAdded", logs: logs, sub: sub}, nil
}

// WatchTariffGroupAdded is a free log subscription operation binding the contract event 0x0d945b4e88c22df3f6d39d8866fac2dd4964090d9d9c04fc74f33c5695e561df.
//
// Solidity: event TariffGroupAdded(ownerAddress indexed address, id bytes16, name string)
func (_HoQuPlatform *HoQuPlatformFilterer) WatchTariffGroupAdded(opts *bind.WatchOpts, sink chan<- *HoQuPlatformTariffGroupAdded, ownerAddress []common.Address) (event.Subscription, error) {

	var ownerAddressRule []interface{}
	for _, ownerAddressItem := range ownerAddress {
		ownerAddressRule = append(ownerAddressRule, ownerAddressItem)
	}

	logs, sub, err := _HoQuPlatform.contract.WatchLogs(opts, "TariffGroupAdded", ownerAddressRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(HoQuPlatformTariffGroupAdded)
				if err := _HoQuPlatform.contract.UnpackLog(event, "TariffGroupAdded", log); err != nil {
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

// HoQuPlatformTariffGroupStatusChangedIterator is returned from FilterTariffGroupStatusChanged and is used to iterate over the raw logs and unpacked data for TariffGroupStatusChanged events raised by the HoQuPlatform contract.
type HoQuPlatformTariffGroupStatusChangedIterator struct {
	Event *HoQuPlatformTariffGroupStatusChanged // Event containing the contract specifics and raw log

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
func (it *HoQuPlatformTariffGroupStatusChangedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(HoQuPlatformTariffGroupStatusChanged)
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
		it.Event = new(HoQuPlatformTariffGroupStatusChanged)
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
func (it *HoQuPlatformTariffGroupStatusChangedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *HoQuPlatformTariffGroupStatusChangedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// HoQuPlatformTariffGroupStatusChanged represents a TariffGroupStatusChanged event raised by the HoQuPlatform contract.
type HoQuPlatformTariffGroupStatusChanged struct {
	OwnerAddress common.Address
	Id           [16]byte
	Status       uint8
	Raw          types.Log // Blockchain specific contextual infos
}

// FilterTariffGroupStatusChanged is a free log retrieval operation binding the contract event 0x00b686e81dae287e0f688bd472beb9aa4c8b029ef31da062379d0d46713dfde6.
//
// Solidity: event TariffGroupStatusChanged(ownerAddress indexed address, id bytes16, status uint8)
func (_HoQuPlatform *HoQuPlatformFilterer) FilterTariffGroupStatusChanged(opts *bind.FilterOpts, ownerAddress []common.Address) (*HoQuPlatformTariffGroupStatusChangedIterator, error) {

	var ownerAddressRule []interface{}
	for _, ownerAddressItem := range ownerAddress {
		ownerAddressRule = append(ownerAddressRule, ownerAddressItem)
	}

	logs, sub, err := _HoQuPlatform.contract.FilterLogs(opts, "TariffGroupStatusChanged", ownerAddressRule)
	if err != nil {
		return nil, err
	}
	return &HoQuPlatformTariffGroupStatusChangedIterator{contract: _HoQuPlatform.contract, event: "TariffGroupStatusChanged", logs: logs, sub: sub}, nil
}

// WatchTariffGroupStatusChanged is a free log subscription operation binding the contract event 0x00b686e81dae287e0f688bd472beb9aa4c8b029ef31da062379d0d46713dfde6.
//
// Solidity: event TariffGroupStatusChanged(ownerAddress indexed address, id bytes16, status uint8)
func (_HoQuPlatform *HoQuPlatformFilterer) WatchTariffGroupStatusChanged(opts *bind.WatchOpts, sink chan<- *HoQuPlatformTariffGroupStatusChanged, ownerAddress []common.Address) (event.Subscription, error) {

	var ownerAddressRule []interface{}
	for _, ownerAddressItem := range ownerAddress {
		ownerAddressRule = append(ownerAddressRule, ownerAddressItem)
	}

	logs, sub, err := _HoQuPlatform.contract.WatchLogs(opts, "TariffGroupStatusChanged", ownerAddressRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(HoQuPlatformTariffGroupStatusChanged)
				if err := _HoQuPlatform.contract.UnpackLog(event, "TariffGroupStatusChanged", log); err != nil {
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

// HoQuPlatformTrackerRegisteredIterator is returned from FilterTrackerRegistered and is used to iterate over the raw logs and unpacked data for TrackerRegistered events raised by the HoQuPlatform contract.
type HoQuPlatformTrackerRegisteredIterator struct {
	Event *HoQuPlatformTrackerRegistered // Event containing the contract specifics and raw log

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
func (it *HoQuPlatformTrackerRegisteredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(HoQuPlatformTrackerRegistered)
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
		it.Event = new(HoQuPlatformTrackerRegistered)
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
func (it *HoQuPlatformTrackerRegisteredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *HoQuPlatformTrackerRegisteredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// HoQuPlatformTrackerRegistered represents a TrackerRegistered event raised by the HoQuPlatform contract.
type HoQuPlatformTrackerRegistered struct {
	OwnerAddress common.Address
	Id           [16]byte
	Name         string
	Raw          types.Log // Blockchain specific contextual infos
}

// FilterTrackerRegistered is a free log retrieval operation binding the contract event 0xb2f67f32df451ce918abd944fd5c8a29589848693057913c5060a8e0f293b388.
//
// Solidity: event TrackerRegistered(ownerAddress indexed address, id bytes16, name string)
func (_HoQuPlatform *HoQuPlatformFilterer) FilterTrackerRegistered(opts *bind.FilterOpts, ownerAddress []common.Address) (*HoQuPlatformTrackerRegisteredIterator, error) {

	var ownerAddressRule []interface{}
	for _, ownerAddressItem := range ownerAddress {
		ownerAddressRule = append(ownerAddressRule, ownerAddressItem)
	}

	logs, sub, err := _HoQuPlatform.contract.FilterLogs(opts, "TrackerRegistered", ownerAddressRule)
	if err != nil {
		return nil, err
	}
	return &HoQuPlatformTrackerRegisteredIterator{contract: _HoQuPlatform.contract, event: "TrackerRegistered", logs: logs, sub: sub}, nil
}

// WatchTrackerRegistered is a free log subscription operation binding the contract event 0xb2f67f32df451ce918abd944fd5c8a29589848693057913c5060a8e0f293b388.
//
// Solidity: event TrackerRegistered(ownerAddress indexed address, id bytes16, name string)
func (_HoQuPlatform *HoQuPlatformFilterer) WatchTrackerRegistered(opts *bind.WatchOpts, sink chan<- *HoQuPlatformTrackerRegistered, ownerAddress []common.Address) (event.Subscription, error) {

	var ownerAddressRule []interface{}
	for _, ownerAddressItem := range ownerAddress {
		ownerAddressRule = append(ownerAddressRule, ownerAddressItem)
	}

	logs, sub, err := _HoQuPlatform.contract.WatchLogs(opts, "TrackerRegistered", ownerAddressRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(HoQuPlatformTrackerRegistered)
				if err := _HoQuPlatform.contract.UnpackLog(event, "TrackerRegistered", log); err != nil {
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

// HoQuPlatformTrackerStatusChangedIterator is returned from FilterTrackerStatusChanged and is used to iterate over the raw logs and unpacked data for TrackerStatusChanged events raised by the HoQuPlatform contract.
type HoQuPlatformTrackerStatusChangedIterator struct {
	Event *HoQuPlatformTrackerStatusChanged // Event containing the contract specifics and raw log

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
func (it *HoQuPlatformTrackerStatusChangedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(HoQuPlatformTrackerStatusChanged)
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
		it.Event = new(HoQuPlatformTrackerStatusChanged)
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
func (it *HoQuPlatformTrackerStatusChangedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *HoQuPlatformTrackerStatusChangedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// HoQuPlatformTrackerStatusChanged represents a TrackerStatusChanged event raised by the HoQuPlatform contract.
type HoQuPlatformTrackerStatusChanged struct {
	OwnerAddress common.Address
	Id           [16]byte
	Status       uint8
	Raw          types.Log // Blockchain specific contextual infos
}

// FilterTrackerStatusChanged is a free log retrieval operation binding the contract event 0x77a747805fe662172d50df0565e43aabadea2efb1198dc3ff8e60fcd57eb7489.
//
// Solidity: event TrackerStatusChanged(ownerAddress indexed address, id bytes16, status uint8)
func (_HoQuPlatform *HoQuPlatformFilterer) FilterTrackerStatusChanged(opts *bind.FilterOpts, ownerAddress []common.Address) (*HoQuPlatformTrackerStatusChangedIterator, error) {

	var ownerAddressRule []interface{}
	for _, ownerAddressItem := range ownerAddress {
		ownerAddressRule = append(ownerAddressRule, ownerAddressItem)
	}

	logs, sub, err := _HoQuPlatform.contract.FilterLogs(opts, "TrackerStatusChanged", ownerAddressRule)
	if err != nil {
		return nil, err
	}
	return &HoQuPlatformTrackerStatusChangedIterator{contract: _HoQuPlatform.contract, event: "TrackerStatusChanged", logs: logs, sub: sub}, nil
}

// WatchTrackerStatusChanged is a free log subscription operation binding the contract event 0x77a747805fe662172d50df0565e43aabadea2efb1198dc3ff8e60fcd57eb7489.
//
// Solidity: event TrackerStatusChanged(ownerAddress indexed address, id bytes16, status uint8)
func (_HoQuPlatform *HoQuPlatformFilterer) WatchTrackerStatusChanged(opts *bind.WatchOpts, sink chan<- *HoQuPlatformTrackerStatusChanged, ownerAddress []common.Address) (event.Subscription, error) {

	var ownerAddressRule []interface{}
	for _, ownerAddressItem := range ownerAddress {
		ownerAddressRule = append(ownerAddressRule, ownerAddressItem)
	}

	logs, sub, err := _HoQuPlatform.contract.WatchLogs(opts, "TrackerStatusChanged", ownerAddressRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(HoQuPlatformTrackerStatusChanged)
				if err := _HoQuPlatform.contract.UnpackLog(event, "TrackerStatusChanged", log); err != nil {
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

// HoQuPlatformUserAddressAddedIterator is returned from FilterUserAddressAdded and is used to iterate over the raw logs and unpacked data for UserAddressAdded events raised by the HoQuPlatform contract.
type HoQuPlatformUserAddressAddedIterator struct {
	Event *HoQuPlatformUserAddressAdded // Event containing the contract specifics and raw log

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
func (it *HoQuPlatformUserAddressAddedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(HoQuPlatformUserAddressAdded)
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
		it.Event = new(HoQuPlatformUserAddressAdded)
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
func (it *HoQuPlatformUserAddressAddedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *HoQuPlatformUserAddressAddedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// HoQuPlatformUserAddressAdded represents a UserAddressAdded event raised by the HoQuPlatform contract.
type HoQuPlatformUserAddressAdded struct {
	OwnerAddress      common.Address
	AdditionalAddress common.Address
	Id                [16]byte
	Raw               types.Log // Blockchain specific contextual infos
}

// FilterUserAddressAdded is a free log retrieval operation binding the contract event 0xb38b406c043d4a3feaa0192d2e1b02f317fa73c284eb3e70b9a1fde8612af1b3.
//
// Solidity: event UserAddressAdded(ownerAddress indexed address, additionalAddress address, id bytes16)
func (_HoQuPlatform *HoQuPlatformFilterer) FilterUserAddressAdded(opts *bind.FilterOpts, ownerAddress []common.Address) (*HoQuPlatformUserAddressAddedIterator, error) {

	var ownerAddressRule []interface{}
	for _, ownerAddressItem := range ownerAddress {
		ownerAddressRule = append(ownerAddressRule, ownerAddressItem)
	}

	logs, sub, err := _HoQuPlatform.contract.FilterLogs(opts, "UserAddressAdded", ownerAddressRule)
	if err != nil {
		return nil, err
	}
	return &HoQuPlatformUserAddressAddedIterator{contract: _HoQuPlatform.contract, event: "UserAddressAdded", logs: logs, sub: sub}, nil
}

// WatchUserAddressAdded is a free log subscription operation binding the contract event 0xb38b406c043d4a3feaa0192d2e1b02f317fa73c284eb3e70b9a1fde8612af1b3.
//
// Solidity: event UserAddressAdded(ownerAddress indexed address, additionalAddress address, id bytes16)
func (_HoQuPlatform *HoQuPlatformFilterer) WatchUserAddressAdded(opts *bind.WatchOpts, sink chan<- *HoQuPlatformUserAddressAdded, ownerAddress []common.Address) (event.Subscription, error) {

	var ownerAddressRule []interface{}
	for _, ownerAddressItem := range ownerAddress {
		ownerAddressRule = append(ownerAddressRule, ownerAddressItem)
	}

	logs, sub, err := _HoQuPlatform.contract.WatchLogs(opts, "UserAddressAdded", ownerAddressRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(HoQuPlatformUserAddressAdded)
				if err := _HoQuPlatform.contract.UnpackLog(event, "UserAddressAdded", log); err != nil {
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

// HoQuPlatformUserPubKeyUpdatedIterator is returned from FilterUserPubKeyUpdated and is used to iterate over the raw logs and unpacked data for UserPubKeyUpdated events raised by the HoQuPlatform contract.
type HoQuPlatformUserPubKeyUpdatedIterator struct {
	Event *HoQuPlatformUserPubKeyUpdated // Event containing the contract specifics and raw log

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
func (it *HoQuPlatformUserPubKeyUpdatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(HoQuPlatformUserPubKeyUpdated)
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
		it.Event = new(HoQuPlatformUserPubKeyUpdated)
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
func (it *HoQuPlatformUserPubKeyUpdatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *HoQuPlatformUserPubKeyUpdatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// HoQuPlatformUserPubKeyUpdated represents a UserPubKeyUpdated event raised by the HoQuPlatform contract.
type HoQuPlatformUserPubKeyUpdated struct {
	OwnerAddress common.Address
	Id           [16]byte
	Raw          types.Log // Blockchain specific contextual infos
}

// FilterUserPubKeyUpdated is a free log retrieval operation binding the contract event 0x85e1546b8c89c1a56c4237d2608010366813f87b475163aadf4074e5bc29339b.
//
// Solidity: event UserPubKeyUpdated(ownerAddress indexed address, id bytes16)
func (_HoQuPlatform *HoQuPlatformFilterer) FilterUserPubKeyUpdated(opts *bind.FilterOpts, ownerAddress []common.Address) (*HoQuPlatformUserPubKeyUpdatedIterator, error) {

	var ownerAddressRule []interface{}
	for _, ownerAddressItem := range ownerAddress {
		ownerAddressRule = append(ownerAddressRule, ownerAddressItem)
	}

	logs, sub, err := _HoQuPlatform.contract.FilterLogs(opts, "UserPubKeyUpdated", ownerAddressRule)
	if err != nil {
		return nil, err
	}
	return &HoQuPlatformUserPubKeyUpdatedIterator{contract: _HoQuPlatform.contract, event: "UserPubKeyUpdated", logs: logs, sub: sub}, nil
}

// WatchUserPubKeyUpdated is a free log subscription operation binding the contract event 0x85e1546b8c89c1a56c4237d2608010366813f87b475163aadf4074e5bc29339b.
//
// Solidity: event UserPubKeyUpdated(ownerAddress indexed address, id bytes16)
func (_HoQuPlatform *HoQuPlatformFilterer) WatchUserPubKeyUpdated(opts *bind.WatchOpts, sink chan<- *HoQuPlatformUserPubKeyUpdated, ownerAddress []common.Address) (event.Subscription, error) {

	var ownerAddressRule []interface{}
	for _, ownerAddressItem := range ownerAddress {
		ownerAddressRule = append(ownerAddressRule, ownerAddressItem)
	}

	logs, sub, err := _HoQuPlatform.contract.WatchLogs(opts, "UserPubKeyUpdated", ownerAddressRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(HoQuPlatformUserPubKeyUpdated)
				if err := _HoQuPlatform.contract.UnpackLog(event, "UserPubKeyUpdated", log); err != nil {
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

// HoQuPlatformUserRegisteredIterator is returned from FilterUserRegistered and is used to iterate over the raw logs and unpacked data for UserRegistered events raised by the HoQuPlatform contract.
type HoQuPlatformUserRegisteredIterator struct {
	Event *HoQuPlatformUserRegistered // Event containing the contract specifics and raw log

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
func (it *HoQuPlatformUserRegisteredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(HoQuPlatformUserRegistered)
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
		it.Event = new(HoQuPlatformUserRegistered)
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
func (it *HoQuPlatformUserRegisteredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *HoQuPlatformUserRegisteredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// HoQuPlatformUserRegistered represents a UserRegistered event raised by the HoQuPlatform contract.
type HoQuPlatformUserRegistered struct {
	OwnerAddress common.Address
	Id           [16]byte
	Role         string
	Raw          types.Log // Blockchain specific contextual infos
}

// FilterUserRegistered is a free log retrieval operation binding the contract event 0xf661ee1472892faaf2bb68cc6874f2759c9148b4234cac6b88f6ca362652f275.
//
// Solidity: event UserRegistered(ownerAddress indexed address, id bytes16, role string)
func (_HoQuPlatform *HoQuPlatformFilterer) FilterUserRegistered(opts *bind.FilterOpts, ownerAddress []common.Address) (*HoQuPlatformUserRegisteredIterator, error) {

	var ownerAddressRule []interface{}
	for _, ownerAddressItem := range ownerAddress {
		ownerAddressRule = append(ownerAddressRule, ownerAddressItem)
	}

	logs, sub, err := _HoQuPlatform.contract.FilterLogs(opts, "UserRegistered", ownerAddressRule)
	if err != nil {
		return nil, err
	}
	return &HoQuPlatformUserRegisteredIterator{contract: _HoQuPlatform.contract, event: "UserRegistered", logs: logs, sub: sub}, nil
}

// WatchUserRegistered is a free log subscription operation binding the contract event 0xf661ee1472892faaf2bb68cc6874f2759c9148b4234cac6b88f6ca362652f275.
//
// Solidity: event UserRegistered(ownerAddress indexed address, id bytes16, role string)
func (_HoQuPlatform *HoQuPlatformFilterer) WatchUserRegistered(opts *bind.WatchOpts, sink chan<- *HoQuPlatformUserRegistered, ownerAddress []common.Address) (event.Subscription, error) {

	var ownerAddressRule []interface{}
	for _, ownerAddressItem := range ownerAddress {
		ownerAddressRule = append(ownerAddressRule, ownerAddressItem)
	}

	logs, sub, err := _HoQuPlatform.contract.WatchLogs(opts, "UserRegistered", ownerAddressRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(HoQuPlatformUserRegistered)
				if err := _HoQuPlatform.contract.UnpackLog(event, "UserRegistered", log); err != nil {
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

// HoQuPlatformUserStatusChangedIterator is returned from FilterUserStatusChanged and is used to iterate over the raw logs and unpacked data for UserStatusChanged events raised by the HoQuPlatform contract.
type HoQuPlatformUserStatusChangedIterator struct {
	Event *HoQuPlatformUserStatusChanged // Event containing the contract specifics and raw log

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
func (it *HoQuPlatformUserStatusChangedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(HoQuPlatformUserStatusChanged)
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
		it.Event = new(HoQuPlatformUserStatusChanged)
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
func (it *HoQuPlatformUserStatusChangedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *HoQuPlatformUserStatusChangedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// HoQuPlatformUserStatusChanged represents a UserStatusChanged event raised by the HoQuPlatform contract.
type HoQuPlatformUserStatusChanged struct {
	OwnerAddress common.Address
	Id           [16]byte
	Status       uint8
	Raw          types.Log // Blockchain specific contextual infos
}

// FilterUserStatusChanged is a free log retrieval operation binding the contract event 0x183408f9d1ce45bacc33cd1a5ef30bf3caccbae9e28fcc8916b62185fdd5ac4d.
//
// Solidity: event UserStatusChanged(ownerAddress indexed address, id bytes16, status uint8)
func (_HoQuPlatform *HoQuPlatformFilterer) FilterUserStatusChanged(opts *bind.FilterOpts, ownerAddress []common.Address) (*HoQuPlatformUserStatusChangedIterator, error) {

	var ownerAddressRule []interface{}
	for _, ownerAddressItem := range ownerAddress {
		ownerAddressRule = append(ownerAddressRule, ownerAddressItem)
	}

	logs, sub, err := _HoQuPlatform.contract.FilterLogs(opts, "UserStatusChanged", ownerAddressRule)
	if err != nil {
		return nil, err
	}
	return &HoQuPlatformUserStatusChangedIterator{contract: _HoQuPlatform.contract, event: "UserStatusChanged", logs: logs, sub: sub}, nil
}

// WatchUserStatusChanged is a free log subscription operation binding the contract event 0x183408f9d1ce45bacc33cd1a5ef30bf3caccbae9e28fcc8916b62185fdd5ac4d.
//
// Solidity: event UserStatusChanged(ownerAddress indexed address, id bytes16, status uint8)
func (_HoQuPlatform *HoQuPlatformFilterer) WatchUserStatusChanged(opts *bind.WatchOpts, sink chan<- *HoQuPlatformUserStatusChanged, ownerAddress []common.Address) (event.Subscription, error) {

	var ownerAddressRule []interface{}
	for _, ownerAddressItem := range ownerAddress {
		ownerAddressRule = append(ownerAddressRule, ownerAddressItem)
	}

	logs, sub, err := _HoQuPlatform.contract.WatchLogs(opts, "UserStatusChanged", ownerAddressRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(HoQuPlatformUserStatusChanged)
				if err := _HoQuPlatform.contract.UnpackLog(event, "UserStatusChanged", log); err != nil {
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
