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
const BasicTokenBin = `0x6060604052341561000f57600080fd5b6102098061001e6000396000f3006060604052600436106100565763ffffffff7c010000000000000000000000000000000000000000000000000000000060003504166318160ddd811461005b57806370a0823114610080578063a9059cbb1461009f575b600080fd5b341561006657600080fd5b61006e6100d5565b60405190815260200160405180910390f35b341561008b57600080fd5b61006e600160a060020a03600435166100db565b34156100aa57600080fd5b6100c1600160a060020a03600435166024356100f6565b604051901515815260200160405180910390f35b60005481565b600160a060020a031660009081526001602052604090205490565b600160a060020a03331660009081526001602052604081205461011f908363ffffffff6101b516565b600160a060020a033381166000908152600160205260408082209390935590851681522054610154908363ffffffff6101c716565b600160a060020a0380851660008181526001602052604090819020939093559133909116907fddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef9085905190815260200160405180910390a350600192915050565b6000828211156101c157fe5b50900390565b6000828201838110156101d657fe5b93925050505600a165627a7a723058206d795ac68ebbd921107d6c3c127d643a60537d2876957b31e1238f6aad44387d0029`

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

// ChangeableRateCrowdsaleABI is the input ABI used to generate the binding from.
const ChangeableRateCrowdsaleABI = "[{\"constant\":false,\"inputs\":[{\"name\":\"_receiver\",\"type\":\"address\"}],\"name\":\"claimOne\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"maxTokensAmount\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"minBuyableAmount\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_receiver\",\"type\":\"address\"},{\"name\":\"_equivalentEthAmount\",\"type\":\"uint256\"}],\"name\":\"sub\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"tokenRate\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"i\",\"type\":\"uint32\"}],\"name\":\"getReceiver\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[],\"name\":\"unpause\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"rateBoundaries\",\"outputs\":[{\"name\":\"amount\",\"type\":\"uint256\"},{\"name\":\"rate\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[],\"name\":\"claim\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"paused\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_tokenRate\",\"type\":\"uint256\"}],\"name\":\"setTokenRate\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"bankAddress\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"isFinished\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[],\"name\":\"pause\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"currentBoundary\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[],\"name\":\"buy\",\"outputs\":[],\"payable\":true,\"stateMutability\":\"payable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"getReceiversCount\",\"outputs\":[{\"name\":\"\",\"type\":\"uint32\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[],\"name\":\"claimAll\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"numOfBoundaries\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[],\"name\":\"finish\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_receiver\",\"type\":\"address\"},{\"name\":\"_equivalentEthAmount\",\"type\":\"uint256\"}],\"name\":\"topUp\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"\",\"type\":\"address\"}],\"name\":\"approved\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_receiver\",\"type\":\"address\"}],\"name\":\"approve\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"\",\"type\":\"address\"}],\"name\":\"tokens\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"issuedTokensAmount\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"nextBoundaryAmount\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"transferOwnership\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_receiver\",\"type\":\"address\"},{\"name\":\"_equivalentEthAmount\",\"type\":\"uint256\"}],\"name\":\"add\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"token\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"name\":\"_tokenAddress\",\"type\":\"address\"},{\"name\":\"_bankAddress\",\"type\":\"address\"},{\"name\":\"_beneficiaryAddress\",\"type\":\"address\"},{\"name\":\"_tokenRate\",\"type\":\"uint256\"},{\"name\":\"_minBuyableAmount\",\"type\":\"uint256\"},{\"name\":\"_maxTokensAmount\",\"type\":\"uint256\"},{\"name\":\"_endDate\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"payable\":true,\"stateMutability\":\"payable\",\"type\":\"fallback\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"_buyer\",\"type\":\"address\"},{\"indexed\":false,\"name\":\"_tokens\",\"type\":\"uint256\"},{\"indexed\":false,\"name\":\"_amount\",\"type\":\"uint256\"}],\"name\":\"TokenBought\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"_receiver\",\"type\":\"address\"},{\"indexed\":false,\"name\":\"_tokens\",\"type\":\"uint256\"},{\"indexed\":false,\"name\":\"_equivalentAmount\",\"type\":\"uint256\"}],\"name\":\"TokenAdded\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"_receiver\",\"type\":\"address\"},{\"indexed\":false,\"name\":\"_tokens\",\"type\":\"uint256\"},{\"indexed\":false,\"name\":\"_equivalentAmount\",\"type\":\"uint256\"}],\"name\":\"TokenToppedUp\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"_receiver\",\"type\":\"address\"},{\"indexed\":false,\"name\":\"_tokens\",\"type\":\"uint256\"},{\"indexed\":false,\"name\":\"_equivalentAmount\",\"type\":\"uint256\"}],\"name\":\"TokenSubtracted\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"_receiver\",\"type\":\"address\"},{\"indexed\":false,\"name\":\"_tokens\",\"type\":\"uint256\"}],\"name\":\"TokenSent\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[],\"name\":\"Pause\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[],\"name\":\"Unpause\",\"type\":\"event\"}]"

// ChangeableRateCrowdsaleBin is the compiled bytecode used for deploying new contracts.
const ChangeableRateCrowdsaleBin = `0x60606040526000805460a060020a60ff021916815560058190556009805460ff19169055600f819055601055341561003657600080fd5b60405160e0806115058339810160405280805191906020018051919060200180519190602001805191906020018051919060200180519190602001805160008054600160a060020a03338116600160a060020a03199283161790925560038054838d1690831617905560028054838c1690831617905560018054928a169290911691909117905560078690556006859055600484905560088190559150604090508051908101604090815268878678326eac900000825261177060208084019190915260108054600181019091556000908152600e909152208151815560208201516001909101555060408051908101604090815269010f0cf064dd59200000825261167660208084019190915260108054600181019091556000908152600e90915220815181556020820151600190910155506040805190810160409081526901969368974c05b00000825261161260208084019190915260108054600181019091556000908152600e909152208151815560208201516001909101555060408051908101604090815269021e19e0c9bab240000082526115ae60208084019190915260108054600181019091556000908152600e90915220815181556020820151600190910155506040805190810160409081526902a5a058fc295ed00000825261154a60208084019190915260108054600181019091556000908152600e909152208151815560208201516001909101555060408051908101604090815269032d26d12e980b60000082526114e660208084019190915260108054600181019091556000908152600e90915220815181556020820151600190910155506040805190810160409081526903b4ad496106b7f00000825261148260208084019190915260108054600181019091556000908152600e909152208151815560208201516001909101555060408051908101604090815269043c33c1937564800000825261141e60208084019190915260108054600181019091556000908152600e9091522081518155602082015160019091015550604080519081016040908152690581767ba6189c400000825261138860208084019190915260108054600181019091556000908152600e909152208151815560208201516001909101555050600f546000908152600e6020526040902054601155505050505050611163806103a26000396000f3006060604052600436106101665763ffffffff60e060020a6000350416630262dc1381146101705780630a3846651461018f5780632511b182146101b457806326ffee08146101c757806331711884146101e95780633c78fe07146101fc5780633f4ba83a146102345780634a7510721461025b5780634e71d92d146102895780635c975abb1461029c57806361241c28146102af5780637822ed49146102c55780637b352962146102d85780638456cb59146102eb5780638da5cb5b146102fe578063a4bd7a2014610311578063a6f2ae3a14610166578063a72a05f714610324578063d1058e5914610350578063d4f114a614610363578063d56b288914610376578063d6f7ddf914610389578063d8b964e6146103ab578063daea85c5146103ca578063e4860339146103e9578063e69b414b14610408578063f07255e31461041b578063f2fde38b1461042e578063f5d82b6b1461044d578063fc0c546a1461046f575b61016e610482565b005b341561017b57600080fd5b61016e600160a060020a0360043516610621565b341561019a57600080fd5b6101a261065f565b60405190815260200160405180910390f35b34156101bf57600080fd5b6101a2610665565b34156101d257600080fd5b61016e600160a060020a036004351660243561066b565b34156101f457600080fd5b6101a261077b565b341561020757600080fd5b61021863ffffffff60043516610781565b604051600160a060020a03909116815260200160405180910390f35b341561023f57600080fd5b6102476107bf565b604051901515815260200160405180910390f35b341561026657600080fd5b610271600435610843565b60405191825260208201526040908101905180910390f35b341561029457600080fd5b61016e61085c565b34156102a757600080fd5b61024761087e565b34156102ba57600080fd5b61016e60043561088e565b34156102d057600080fd5b6102186108bb565b34156102e357600080fd5b6102476108ca565b34156102f657600080fd5b6102476108d3565b341561030957600080fd5b61021861095c565b341561031c57600080fd5b6101a261096b565b341561032f57600080fd5b610337610971565b60405163ffffffff909116815260200160405180910390f35b341561035b57600080fd5b61016e61099a565b341561036e57600080fd5b6101a2610a59565b341561038157600080fd5b61016e610a5f565b341561039457600080fd5b61016e600160a060020a0360043516602435610ba1565b34156103b657600080fd5b610247600160a060020a0360043516610c53565b34156103d557600080fd5b61016e600160a060020a0360043516610c68565b34156103f457600080fd5b6101a2600160a060020a0360043516610cbe565b341561041357600080fd5b6101a2610cd0565b341561042657600080fd5b6101a2610cd6565b341561043957600080fd5b61016e600160a060020a0360043516610cdc565b341561045857600080fd5b61016e600160a060020a0360043516602435610d33565b341561047a57600080fd5b610218610e14565b6009546000908190819060ff161561049957600080fd5b600454600554106104a957600080fd5b6008544211156104b857600080fd5b60005460a060020a900460ff16156104cf57600080fd5b600754349350600092506104e9908463ffffffff610e2316565b90506004548160055401111561053a5760055460045461050e9163ffffffff610e4e16565b905061052560075482610e6090919063ffffffff16565b9250610537348463ffffffff610e4e16565b91505b60055461054d908263ffffffff610e7716565b600581905560045490111561056157600080fd5b61056b3382610e86565b33600160a060020a03167f28cab0d660ed8aedd61a8c9db00b97f6a2d67e07d87795994f440b18bc5f1aa3828560405191825260208201526040908101905180910390a2600154600160a060020a031683156108fc0284604051600060405180830381858888f1935050505015156105e257600080fd5b600082111561061c57600160a060020a03331682156108fc0283604051600060405180830381858888f19350505050151561061c57600080fd5b505050565b60005433600160a060020a0390811691161461063c57600080fd5b60005460a060020a900460ff161561065357600080fd5b61065c81610eaf565b50565b60045481565b60065481565b6000805433600160a060020a0390811691161461068757600080fd5b60005460a060020a900460ff161561069e57600080fd5b6007546106b1908363ffffffff610e2316565b600160a060020a0384166000908152600a6020526040902054909150819010156106da57600080fd5b600160a060020a0383166000908152600a6020526040902054610703908263ffffffff610e4e16565b600160a060020a0384166000908152600a602052604090205560055461072f908263ffffffff610e4e16565b600555600160a060020a0383167f2da8b2becb27b189a54d21689aa110c877379a94c6f59f115cedc91d66239b33828460405191825260208201526040908101905180910390a2505050565b60075481565b6000805433600160a060020a0390811691161461079d57600080fd5b5063ffffffff166000908152600c6020526040902054600160a060020a031690565b6000805433600160a060020a039081169116146107db57600080fd5b60005460a060020a900460ff1615156107f357600080fd5b6000805474ff0000000000000000000000000000000000000000191690557f7805862f689e2f13df9f062ff482ad3ad112aca9e0847911ed832e158c525b3360405160405180910390a150600190565b600e602052600090815260409020805460019091015482565b60005460a060020a900460ff161561087357600080fd5b61087c33610eaf565b565b60005460a060020a900460ff1681565b60005433600160a060020a039081169116146108a957600080fd5b600081116108b657600080fd5b600755565b600254600160a060020a031681565b60095460ff1681565b6000805433600160a060020a039081169116146108ef57600080fd5b60005460a060020a900460ff161561090657600080fd5b6000805474ff0000000000000000000000000000000000000000191660a060020a1790557f6985a02210a168e66602d3235cb6db0e70f92b3ba4d376a33c0f3d9434bff62560405160405180910390a150600190565b600054600160a060020a031681565b600f5481565b6000805433600160a060020a0390811691161461098d57600080fd5b50600d5463ffffffff1690565b60008054819033600160a060020a039081169116146109b857600080fd5b60005460a060020a900460ff16156109cf57600080fd5b600091505b600d5463ffffffff9081169083161015610a55575063ffffffff81166000908152600c6020908152604080832054600160a060020a0316808452600b9092529091205460ff168015610a3c5750600160a060020a0381166000908152600a6020526040812054115b15610a4a57610a4a81610eaf565b6001909101906109d4565b5050565b60105481565b60005433600160a060020a03908116911614610a7a57600080fd5b600454600554101580610a8e575060085442115b1515610a9957600080fd5b60095460ff1615610aa957600080fd5b6009805460ff19166001179055600354600254600160a060020a039182169163a9059cbb9116826370a082313060006040516020015260405160e060020a63ffffffff8416028152600160a060020a039091166004820152602401602060405180830381600087803b1515610b1d57600080fd5b6102c65a03f11515610b2e57600080fd5b5050506040518051905060006040516020015260405160e060020a63ffffffff8516028152600160a060020a0390921660048301526024820152604401602060405180830381600087803b1515610b8457600080fd5b6102c65a03f11515610b9557600080fd5b50505060405180515050565b6000805433600160a060020a03908116911614610bbd57600080fd5b60005460a060020a900460ff1615610bd457600080fd5b600754610be7908363ffffffff610e2316565b600554909150610bfd908263ffffffff610e7716565b600555610c0a8382610e86565b82600160a060020a03167ff10143e1b37ebaf3698e98976d8bddabcea045adc939c03b245da18e81ea9651828460405191825260208201526040908101905180910390a2505050565b600b6020526000908152604090205460ff1681565b60005433600160a060020a03908116911614610c8357600080fd5b60005460a060020a900460ff1615610c9a57600080fd5b600160a060020a03166000908152600b60205260409020805460ff19166001179055565b600a6020526000908152604090205481565b60055481565b60115481565b60005433600160a060020a03908116911614610cf757600080fd5b600160a060020a0381161561065c5760008054600160a060020a03831673ffffffffffffffffffffffffffffffffffffffff1990911617905550565b6000805433600160a060020a03908116911614610d4f57600080fd5b60095460ff1615610d5f57600080fd5b60045460055410610d6f57600080fd5b600854421115610d7e57600080fd5b60005460a060020a900460ff1615610d9557600080fd5b600754610da8908363ffffffff610e2316565b600554909150610dbe908263ffffffff610e7716565b600555610dcb8382610e86565b82600160a060020a03167fa818a22273fc309f0a3682b642c74c5b5c25c0615ff378d07767cd231e19fffc828460405191825260208201526040908101905180910390a2505050565b600354600160a060020a031681565b6000828202831580610e3f5750828482811515610e3c57fe5b04145b1515610e4757fe5b9392505050565b600082821115610e5a57fe5b50900390565b6000808284811515610e6e57fe5b04949350505050565b600082820183811015610e4757fe5b60005460a060020a900460ff1615610e9d57600080fd5b610ea78282610ffd565b610a556110ee565b6000805460a060020a900460ff1615610ec757600080fd5b600160a060020a0382166000908152600b602052604090205460ff161515610eee57600080fd5b600160a060020a0382166000908152600a602052604081205411610f1157600080fd5b50600160a060020a038082166000908152600a60205260408082208054908390556003546002549194908116936323b872dd93929091169186918691516020015260405160e060020a63ffffffff8616028152600160a060020a0393841660048201529190921660248201526044810191909152606401602060405180830381600087803b1515610fa157600080fd5b6102c65a03f11515610fb257600080fd5b50505060405180515050600160a060020a0382167fd4ff88d569a7ad2bfd6b17da9dbe82e2ccc016fd0051e08ff1ad1e6fe8fc9b558260405190815260200160405180910390a25050565b60005460a060020a900460ff161561101457600080fd5b600160a060020a0382166000908152600a602052604090205415156110a557600d805463ffffffff9081166000908152600c60209081526040808320805473ffffffffffffffffffffffffffffffffffffffff1916600160a060020a038916908117909155855463ffffffff19811690861660010190951694909417909455918152600b90915220805460ff191690555b600160a060020a0382166000908152600a60205260409020546110ce908263ffffffff610e7716565b600160a060020a039092166000908152600a602052604090209190915550565b6011546005541061087c57600f80546001019081905560105490106111185760045460115561087c565b600f546000908152600e602052604090208054601155600101546007555600a165627a7a72305820ef59afac856d85918316c9527ea74bf6c000c249e1f96f7bbacbe46c8cee292e0029`

// DeployChangeableRateCrowdsale deploys a new Ethereum contract, binding an instance of ChangeableRateCrowdsale to it.
func DeployChangeableRateCrowdsale(auth *bind.TransactOpts, backend bind.ContractBackend, _tokenAddress common.Address, _bankAddress common.Address, _beneficiaryAddress common.Address, _tokenRate *big.Int, _minBuyableAmount *big.Int, _maxTokensAmount *big.Int, _endDate *big.Int) (common.Address, *types.Transaction, *ChangeableRateCrowdsale, error) {
	parsed, err := abi.JSON(strings.NewReader(ChangeableRateCrowdsaleABI))
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	address, tx, contract, err := bind.DeployContract(auth, parsed, common.FromHex(ChangeableRateCrowdsaleBin), backend, _tokenAddress, _bankAddress, _beneficiaryAddress, _tokenRate, _minBuyableAmount, _maxTokensAmount, _endDate)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &ChangeableRateCrowdsale{ChangeableRateCrowdsaleCaller: ChangeableRateCrowdsaleCaller{contract: contract}, ChangeableRateCrowdsaleTransactor: ChangeableRateCrowdsaleTransactor{contract: contract}}, nil
}

// ChangeableRateCrowdsale is an auto generated Go binding around an Ethereum contract.
type ChangeableRateCrowdsale struct {
	ChangeableRateCrowdsaleCaller     // Read-only binding to the contract
	ChangeableRateCrowdsaleTransactor // Write-only binding to the contract
}

// ChangeableRateCrowdsaleCaller is an auto generated read-only Go binding around an Ethereum contract.
type ChangeableRateCrowdsaleCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ChangeableRateCrowdsaleTransactor is an auto generated write-only Go binding around an Ethereum contract.
type ChangeableRateCrowdsaleTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ChangeableRateCrowdsaleSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type ChangeableRateCrowdsaleSession struct {
	Contract     *ChangeableRateCrowdsale // Generic contract binding to set the session for
	CallOpts     bind.CallOpts            // Call options to use throughout this session
	TransactOpts bind.TransactOpts        // Transaction auth options to use throughout this session
}

// ChangeableRateCrowdsaleCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type ChangeableRateCrowdsaleCallerSession struct {
	Contract *ChangeableRateCrowdsaleCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts                  // Call options to use throughout this session
}

// ChangeableRateCrowdsaleTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type ChangeableRateCrowdsaleTransactorSession struct {
	Contract     *ChangeableRateCrowdsaleTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts                  // Transaction auth options to use throughout this session
}

// ChangeableRateCrowdsaleRaw is an auto generated low-level Go binding around an Ethereum contract.
type ChangeableRateCrowdsaleRaw struct {
	Contract *ChangeableRateCrowdsale // Generic contract binding to access the raw methods on
}

// ChangeableRateCrowdsaleCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type ChangeableRateCrowdsaleCallerRaw struct {
	Contract *ChangeableRateCrowdsaleCaller // Generic read-only contract binding to access the raw methods on
}

// ChangeableRateCrowdsaleTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type ChangeableRateCrowdsaleTransactorRaw struct {
	Contract *ChangeableRateCrowdsaleTransactor // Generic write-only contract binding to access the raw methods on
}

// NewChangeableRateCrowdsale creates a new instance of ChangeableRateCrowdsale, bound to a specific deployed contract.
func NewChangeableRateCrowdsale(address common.Address, backend bind.ContractBackend) (*ChangeableRateCrowdsale, error) {
	contract, err := bindChangeableRateCrowdsale(address, backend, backend)
	if err != nil {
		return nil, err
	}
	return &ChangeableRateCrowdsale{ChangeableRateCrowdsaleCaller: ChangeableRateCrowdsaleCaller{contract: contract}, ChangeableRateCrowdsaleTransactor: ChangeableRateCrowdsaleTransactor{contract: contract}}, nil
}

// NewChangeableRateCrowdsaleCaller creates a new read-only instance of ChangeableRateCrowdsale, bound to a specific deployed contract.
func NewChangeableRateCrowdsaleCaller(address common.Address, caller bind.ContractCaller) (*ChangeableRateCrowdsaleCaller, error) {
	contract, err := bindChangeableRateCrowdsale(address, caller, nil)
	if err != nil {
		return nil, err
	}
	return &ChangeableRateCrowdsaleCaller{contract: contract}, nil
}

// NewChangeableRateCrowdsaleTransactor creates a new write-only instance of ChangeableRateCrowdsale, bound to a specific deployed contract.
func NewChangeableRateCrowdsaleTransactor(address common.Address, transactor bind.ContractTransactor) (*ChangeableRateCrowdsaleTransactor, error) {
	contract, err := bindChangeableRateCrowdsale(address, nil, transactor)
	if err != nil {
		return nil, err
	}
	return &ChangeableRateCrowdsaleTransactor{contract: contract}, nil
}

// bindChangeableRateCrowdsale binds a generic wrapper to an already deployed contract.
func bindChangeableRateCrowdsale(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(ChangeableRateCrowdsaleABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_ChangeableRateCrowdsale *ChangeableRateCrowdsaleRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _ChangeableRateCrowdsale.Contract.ChangeableRateCrowdsaleCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_ChangeableRateCrowdsale *ChangeableRateCrowdsaleRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ChangeableRateCrowdsale.Contract.ChangeableRateCrowdsaleTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_ChangeableRateCrowdsale *ChangeableRateCrowdsaleRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _ChangeableRateCrowdsale.Contract.ChangeableRateCrowdsaleTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_ChangeableRateCrowdsale *ChangeableRateCrowdsaleCallerRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _ChangeableRateCrowdsale.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_ChangeableRateCrowdsale *ChangeableRateCrowdsaleTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ChangeableRateCrowdsale.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_ChangeableRateCrowdsale *ChangeableRateCrowdsaleTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _ChangeableRateCrowdsale.Contract.contract.Transact(opts, method, params...)
}

// Approved is a free data retrieval call binding the contract method 0xd8b964e6.
//
// Solidity: function approved( address) constant returns(bool)
func (_ChangeableRateCrowdsale *ChangeableRateCrowdsaleCaller) Approved(opts *bind.CallOpts, arg0 common.Address) (bool, error) {
	var (
		ret0 = new(bool)
	)
	out := ret0
	err := _ChangeableRateCrowdsale.contract.Call(opts, out, "approved", arg0)
	return *ret0, err
}

// Approved is a free data retrieval call binding the contract method 0xd8b964e6.
//
// Solidity: function approved( address) constant returns(bool)
func (_ChangeableRateCrowdsale *ChangeableRateCrowdsaleSession) Approved(arg0 common.Address) (bool, error) {
	return _ChangeableRateCrowdsale.Contract.Approved(&_ChangeableRateCrowdsale.CallOpts, arg0)
}

// Approved is a free data retrieval call binding the contract method 0xd8b964e6.
//
// Solidity: function approved( address) constant returns(bool)
func (_ChangeableRateCrowdsale *ChangeableRateCrowdsaleCallerSession) Approved(arg0 common.Address) (bool, error) {
	return _ChangeableRateCrowdsale.Contract.Approved(&_ChangeableRateCrowdsale.CallOpts, arg0)
}

// BankAddress is a free data retrieval call binding the contract method 0x7822ed49.
//
// Solidity: function bankAddress() constant returns(address)
func (_ChangeableRateCrowdsale *ChangeableRateCrowdsaleCaller) BankAddress(opts *bind.CallOpts) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _ChangeableRateCrowdsale.contract.Call(opts, out, "bankAddress")
	return *ret0, err
}

// BankAddress is a free data retrieval call binding the contract method 0x7822ed49.
//
// Solidity: function bankAddress() constant returns(address)
func (_ChangeableRateCrowdsale *ChangeableRateCrowdsaleSession) BankAddress() (common.Address, error) {
	return _ChangeableRateCrowdsale.Contract.BankAddress(&_ChangeableRateCrowdsale.CallOpts)
}

// BankAddress is a free data retrieval call binding the contract method 0x7822ed49.
//
// Solidity: function bankAddress() constant returns(address)
func (_ChangeableRateCrowdsale *ChangeableRateCrowdsaleCallerSession) BankAddress() (common.Address, error) {
	return _ChangeableRateCrowdsale.Contract.BankAddress(&_ChangeableRateCrowdsale.CallOpts)
}

// CurrentBoundary is a free data retrieval call binding the contract method 0xa4bd7a20.
//
// Solidity: function currentBoundary() constant returns(uint256)
func (_ChangeableRateCrowdsale *ChangeableRateCrowdsaleCaller) CurrentBoundary(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _ChangeableRateCrowdsale.contract.Call(opts, out, "currentBoundary")
	return *ret0, err
}

// CurrentBoundary is a free data retrieval call binding the contract method 0xa4bd7a20.
//
// Solidity: function currentBoundary() constant returns(uint256)
func (_ChangeableRateCrowdsale *ChangeableRateCrowdsaleSession) CurrentBoundary() (*big.Int, error) {
	return _ChangeableRateCrowdsale.Contract.CurrentBoundary(&_ChangeableRateCrowdsale.CallOpts)
}

// CurrentBoundary is a free data retrieval call binding the contract method 0xa4bd7a20.
//
// Solidity: function currentBoundary() constant returns(uint256)
func (_ChangeableRateCrowdsale *ChangeableRateCrowdsaleCallerSession) CurrentBoundary() (*big.Int, error) {
	return _ChangeableRateCrowdsale.Contract.CurrentBoundary(&_ChangeableRateCrowdsale.CallOpts)
}

// GetReceiver is a free data retrieval call binding the contract method 0x3c78fe07.
//
// Solidity: function getReceiver(i uint32) constant returns(address)
func (_ChangeableRateCrowdsale *ChangeableRateCrowdsaleCaller) GetReceiver(opts *bind.CallOpts, i uint32) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _ChangeableRateCrowdsale.contract.Call(opts, out, "getReceiver", i)
	return *ret0, err
}

// GetReceiver is a free data retrieval call binding the contract method 0x3c78fe07.
//
// Solidity: function getReceiver(i uint32) constant returns(address)
func (_ChangeableRateCrowdsale *ChangeableRateCrowdsaleSession) GetReceiver(i uint32) (common.Address, error) {
	return _ChangeableRateCrowdsale.Contract.GetReceiver(&_ChangeableRateCrowdsale.CallOpts, i)
}

// GetReceiver is a free data retrieval call binding the contract method 0x3c78fe07.
//
// Solidity: function getReceiver(i uint32) constant returns(address)
func (_ChangeableRateCrowdsale *ChangeableRateCrowdsaleCallerSession) GetReceiver(i uint32) (common.Address, error) {
	return _ChangeableRateCrowdsale.Contract.GetReceiver(&_ChangeableRateCrowdsale.CallOpts, i)
}

// GetReceiversCount is a free data retrieval call binding the contract method 0xa72a05f7.
//
// Solidity: function getReceiversCount() constant returns(uint32)
func (_ChangeableRateCrowdsale *ChangeableRateCrowdsaleCaller) GetReceiversCount(opts *bind.CallOpts) (uint32, error) {
	var (
		ret0 = new(uint32)
	)
	out := ret0
	err := _ChangeableRateCrowdsale.contract.Call(opts, out, "getReceiversCount")
	return *ret0, err
}

// GetReceiversCount is a free data retrieval call binding the contract method 0xa72a05f7.
//
// Solidity: function getReceiversCount() constant returns(uint32)
func (_ChangeableRateCrowdsale *ChangeableRateCrowdsaleSession) GetReceiversCount() (uint32, error) {
	return _ChangeableRateCrowdsale.Contract.GetReceiversCount(&_ChangeableRateCrowdsale.CallOpts)
}

// GetReceiversCount is a free data retrieval call binding the contract method 0xa72a05f7.
//
// Solidity: function getReceiversCount() constant returns(uint32)
func (_ChangeableRateCrowdsale *ChangeableRateCrowdsaleCallerSession) GetReceiversCount() (uint32, error) {
	return _ChangeableRateCrowdsale.Contract.GetReceiversCount(&_ChangeableRateCrowdsale.CallOpts)
}

// IsFinished is a free data retrieval call binding the contract method 0x7b352962.
//
// Solidity: function isFinished() constant returns(bool)
func (_ChangeableRateCrowdsale *ChangeableRateCrowdsaleCaller) IsFinished(opts *bind.CallOpts) (bool, error) {
	var (
		ret0 = new(bool)
	)
	out := ret0
	err := _ChangeableRateCrowdsale.contract.Call(opts, out, "isFinished")
	return *ret0, err
}

// IsFinished is a free data retrieval call binding the contract method 0x7b352962.
//
// Solidity: function isFinished() constant returns(bool)
func (_ChangeableRateCrowdsale *ChangeableRateCrowdsaleSession) IsFinished() (bool, error) {
	return _ChangeableRateCrowdsale.Contract.IsFinished(&_ChangeableRateCrowdsale.CallOpts)
}

// IsFinished is a free data retrieval call binding the contract method 0x7b352962.
//
// Solidity: function isFinished() constant returns(bool)
func (_ChangeableRateCrowdsale *ChangeableRateCrowdsaleCallerSession) IsFinished() (bool, error) {
	return _ChangeableRateCrowdsale.Contract.IsFinished(&_ChangeableRateCrowdsale.CallOpts)
}

// IssuedTokensAmount is a free data retrieval call binding the contract method 0xe69b414b.
//
// Solidity: function issuedTokensAmount() constant returns(uint256)
func (_ChangeableRateCrowdsale *ChangeableRateCrowdsaleCaller) IssuedTokensAmount(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _ChangeableRateCrowdsale.contract.Call(opts, out, "issuedTokensAmount")
	return *ret0, err
}

// IssuedTokensAmount is a free data retrieval call binding the contract method 0xe69b414b.
//
// Solidity: function issuedTokensAmount() constant returns(uint256)
func (_ChangeableRateCrowdsale *ChangeableRateCrowdsaleSession) IssuedTokensAmount() (*big.Int, error) {
	return _ChangeableRateCrowdsale.Contract.IssuedTokensAmount(&_ChangeableRateCrowdsale.CallOpts)
}

// IssuedTokensAmount is a free data retrieval call binding the contract method 0xe69b414b.
//
// Solidity: function issuedTokensAmount() constant returns(uint256)
func (_ChangeableRateCrowdsale *ChangeableRateCrowdsaleCallerSession) IssuedTokensAmount() (*big.Int, error) {
	return _ChangeableRateCrowdsale.Contract.IssuedTokensAmount(&_ChangeableRateCrowdsale.CallOpts)
}

// MaxTokensAmount is a free data retrieval call binding the contract method 0x0a384665.
//
// Solidity: function maxTokensAmount() constant returns(uint256)
func (_ChangeableRateCrowdsale *ChangeableRateCrowdsaleCaller) MaxTokensAmount(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _ChangeableRateCrowdsale.contract.Call(opts, out, "maxTokensAmount")
	return *ret0, err
}

// MaxTokensAmount is a free data retrieval call binding the contract method 0x0a384665.
//
// Solidity: function maxTokensAmount() constant returns(uint256)
func (_ChangeableRateCrowdsale *ChangeableRateCrowdsaleSession) MaxTokensAmount() (*big.Int, error) {
	return _ChangeableRateCrowdsale.Contract.MaxTokensAmount(&_ChangeableRateCrowdsale.CallOpts)
}

// MaxTokensAmount is a free data retrieval call binding the contract method 0x0a384665.
//
// Solidity: function maxTokensAmount() constant returns(uint256)
func (_ChangeableRateCrowdsale *ChangeableRateCrowdsaleCallerSession) MaxTokensAmount() (*big.Int, error) {
	return _ChangeableRateCrowdsale.Contract.MaxTokensAmount(&_ChangeableRateCrowdsale.CallOpts)
}

// MinBuyableAmount is a free data retrieval call binding the contract method 0x2511b182.
//
// Solidity: function minBuyableAmount() constant returns(uint256)
func (_ChangeableRateCrowdsale *ChangeableRateCrowdsaleCaller) MinBuyableAmount(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _ChangeableRateCrowdsale.contract.Call(opts, out, "minBuyableAmount")
	return *ret0, err
}

// MinBuyableAmount is a free data retrieval call binding the contract method 0x2511b182.
//
// Solidity: function minBuyableAmount() constant returns(uint256)
func (_ChangeableRateCrowdsale *ChangeableRateCrowdsaleSession) MinBuyableAmount() (*big.Int, error) {
	return _ChangeableRateCrowdsale.Contract.MinBuyableAmount(&_ChangeableRateCrowdsale.CallOpts)
}

// MinBuyableAmount is a free data retrieval call binding the contract method 0x2511b182.
//
// Solidity: function minBuyableAmount() constant returns(uint256)
func (_ChangeableRateCrowdsale *ChangeableRateCrowdsaleCallerSession) MinBuyableAmount() (*big.Int, error) {
	return _ChangeableRateCrowdsale.Contract.MinBuyableAmount(&_ChangeableRateCrowdsale.CallOpts)
}

// NextBoundaryAmount is a free data retrieval call binding the contract method 0xf07255e3.
//
// Solidity: function nextBoundaryAmount() constant returns(uint256)
func (_ChangeableRateCrowdsale *ChangeableRateCrowdsaleCaller) NextBoundaryAmount(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _ChangeableRateCrowdsale.contract.Call(opts, out, "nextBoundaryAmount")
	return *ret0, err
}

// NextBoundaryAmount is a free data retrieval call binding the contract method 0xf07255e3.
//
// Solidity: function nextBoundaryAmount() constant returns(uint256)
func (_ChangeableRateCrowdsale *ChangeableRateCrowdsaleSession) NextBoundaryAmount() (*big.Int, error) {
	return _ChangeableRateCrowdsale.Contract.NextBoundaryAmount(&_ChangeableRateCrowdsale.CallOpts)
}

// NextBoundaryAmount is a free data retrieval call binding the contract method 0xf07255e3.
//
// Solidity: function nextBoundaryAmount() constant returns(uint256)
func (_ChangeableRateCrowdsale *ChangeableRateCrowdsaleCallerSession) NextBoundaryAmount() (*big.Int, error) {
	return _ChangeableRateCrowdsale.Contract.NextBoundaryAmount(&_ChangeableRateCrowdsale.CallOpts)
}

// NumOfBoundaries is a free data retrieval call binding the contract method 0xd4f114a6.
//
// Solidity: function numOfBoundaries() constant returns(uint256)
func (_ChangeableRateCrowdsale *ChangeableRateCrowdsaleCaller) NumOfBoundaries(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _ChangeableRateCrowdsale.contract.Call(opts, out, "numOfBoundaries")
	return *ret0, err
}

// NumOfBoundaries is a free data retrieval call binding the contract method 0xd4f114a6.
//
// Solidity: function numOfBoundaries() constant returns(uint256)
func (_ChangeableRateCrowdsale *ChangeableRateCrowdsaleSession) NumOfBoundaries() (*big.Int, error) {
	return _ChangeableRateCrowdsale.Contract.NumOfBoundaries(&_ChangeableRateCrowdsale.CallOpts)
}

// NumOfBoundaries is a free data retrieval call binding the contract method 0xd4f114a6.
//
// Solidity: function numOfBoundaries() constant returns(uint256)
func (_ChangeableRateCrowdsale *ChangeableRateCrowdsaleCallerSession) NumOfBoundaries() (*big.Int, error) {
	return _ChangeableRateCrowdsale.Contract.NumOfBoundaries(&_ChangeableRateCrowdsale.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() constant returns(address)
func (_ChangeableRateCrowdsale *ChangeableRateCrowdsaleCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _ChangeableRateCrowdsale.contract.Call(opts, out, "owner")
	return *ret0, err
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() constant returns(address)
func (_ChangeableRateCrowdsale *ChangeableRateCrowdsaleSession) Owner() (common.Address, error) {
	return _ChangeableRateCrowdsale.Contract.Owner(&_ChangeableRateCrowdsale.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() constant returns(address)
func (_ChangeableRateCrowdsale *ChangeableRateCrowdsaleCallerSession) Owner() (common.Address, error) {
	return _ChangeableRateCrowdsale.Contract.Owner(&_ChangeableRateCrowdsale.CallOpts)
}

// Paused is a free data retrieval call binding the contract method 0x5c975abb.
//
// Solidity: function paused() constant returns(bool)
func (_ChangeableRateCrowdsale *ChangeableRateCrowdsaleCaller) Paused(opts *bind.CallOpts) (bool, error) {
	var (
		ret0 = new(bool)
	)
	out := ret0
	err := _ChangeableRateCrowdsale.contract.Call(opts, out, "paused")
	return *ret0, err
}

// Paused is a free data retrieval call binding the contract method 0x5c975abb.
//
// Solidity: function paused() constant returns(bool)
func (_ChangeableRateCrowdsale *ChangeableRateCrowdsaleSession) Paused() (bool, error) {
	return _ChangeableRateCrowdsale.Contract.Paused(&_ChangeableRateCrowdsale.CallOpts)
}

// Paused is a free data retrieval call binding the contract method 0x5c975abb.
//
// Solidity: function paused() constant returns(bool)
func (_ChangeableRateCrowdsale *ChangeableRateCrowdsaleCallerSession) Paused() (bool, error) {
	return _ChangeableRateCrowdsale.Contract.Paused(&_ChangeableRateCrowdsale.CallOpts)
}

// RateBoundaries is a free data retrieval call binding the contract method 0x4a751072.
//
// Solidity: function rateBoundaries( uint256) constant returns(amount uint256, rate uint256)
func (_ChangeableRateCrowdsale *ChangeableRateCrowdsaleCaller) RateBoundaries(opts *bind.CallOpts, arg0 *big.Int) (struct {
	Amount *big.Int
	Rate   *big.Int
}, error) {
	ret := new(struct {
		Amount *big.Int
		Rate   *big.Int
	})
	out := ret
	err := _ChangeableRateCrowdsale.contract.Call(opts, out, "rateBoundaries", arg0)
	return *ret, err
}

// RateBoundaries is a free data retrieval call binding the contract method 0x4a751072.
//
// Solidity: function rateBoundaries( uint256) constant returns(amount uint256, rate uint256)
func (_ChangeableRateCrowdsale *ChangeableRateCrowdsaleSession) RateBoundaries(arg0 *big.Int) (struct {
	Amount *big.Int
	Rate   *big.Int
}, error) {
	return _ChangeableRateCrowdsale.Contract.RateBoundaries(&_ChangeableRateCrowdsale.CallOpts, arg0)
}

// RateBoundaries is a free data retrieval call binding the contract method 0x4a751072.
//
// Solidity: function rateBoundaries( uint256) constant returns(amount uint256, rate uint256)
func (_ChangeableRateCrowdsale *ChangeableRateCrowdsaleCallerSession) RateBoundaries(arg0 *big.Int) (struct {
	Amount *big.Int
	Rate   *big.Int
}, error) {
	return _ChangeableRateCrowdsale.Contract.RateBoundaries(&_ChangeableRateCrowdsale.CallOpts, arg0)
}

// Token is a free data retrieval call binding the contract method 0xfc0c546a.
//
// Solidity: function token() constant returns(address)
func (_ChangeableRateCrowdsale *ChangeableRateCrowdsaleCaller) Token(opts *bind.CallOpts) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _ChangeableRateCrowdsale.contract.Call(opts, out, "token")
	return *ret0, err
}

// Token is a free data retrieval call binding the contract method 0xfc0c546a.
//
// Solidity: function token() constant returns(address)
func (_ChangeableRateCrowdsale *ChangeableRateCrowdsaleSession) Token() (common.Address, error) {
	return _ChangeableRateCrowdsale.Contract.Token(&_ChangeableRateCrowdsale.CallOpts)
}

// Token is a free data retrieval call binding the contract method 0xfc0c546a.
//
// Solidity: function token() constant returns(address)
func (_ChangeableRateCrowdsale *ChangeableRateCrowdsaleCallerSession) Token() (common.Address, error) {
	return _ChangeableRateCrowdsale.Contract.Token(&_ChangeableRateCrowdsale.CallOpts)
}

// TokenRate is a free data retrieval call binding the contract method 0x31711884.
//
// Solidity: function tokenRate() constant returns(uint256)
func (_ChangeableRateCrowdsale *ChangeableRateCrowdsaleCaller) TokenRate(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _ChangeableRateCrowdsale.contract.Call(opts, out, "tokenRate")
	return *ret0, err
}

// TokenRate is a free data retrieval call binding the contract method 0x31711884.
//
// Solidity: function tokenRate() constant returns(uint256)
func (_ChangeableRateCrowdsale *ChangeableRateCrowdsaleSession) TokenRate() (*big.Int, error) {
	return _ChangeableRateCrowdsale.Contract.TokenRate(&_ChangeableRateCrowdsale.CallOpts)
}

// TokenRate is a free data retrieval call binding the contract method 0x31711884.
//
// Solidity: function tokenRate() constant returns(uint256)
func (_ChangeableRateCrowdsale *ChangeableRateCrowdsaleCallerSession) TokenRate() (*big.Int, error) {
	return _ChangeableRateCrowdsale.Contract.TokenRate(&_ChangeableRateCrowdsale.CallOpts)
}

// Tokens is a free data retrieval call binding the contract method 0xe4860339.
//
// Solidity: function tokens( address) constant returns(uint256)
func (_ChangeableRateCrowdsale *ChangeableRateCrowdsaleCaller) Tokens(opts *bind.CallOpts, arg0 common.Address) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _ChangeableRateCrowdsale.contract.Call(opts, out, "tokens", arg0)
	return *ret0, err
}

// Tokens is a free data retrieval call binding the contract method 0xe4860339.
//
// Solidity: function tokens( address) constant returns(uint256)
func (_ChangeableRateCrowdsale *ChangeableRateCrowdsaleSession) Tokens(arg0 common.Address) (*big.Int, error) {
	return _ChangeableRateCrowdsale.Contract.Tokens(&_ChangeableRateCrowdsale.CallOpts, arg0)
}

// Tokens is a free data retrieval call binding the contract method 0xe4860339.
//
// Solidity: function tokens( address) constant returns(uint256)
func (_ChangeableRateCrowdsale *ChangeableRateCrowdsaleCallerSession) Tokens(arg0 common.Address) (*big.Int, error) {
	return _ChangeableRateCrowdsale.Contract.Tokens(&_ChangeableRateCrowdsale.CallOpts, arg0)
}

// Add is a paid mutator transaction binding the contract method 0xf5d82b6b.
//
// Solidity: function add(_receiver address, _equivalentEthAmount uint256) returns()
func (_ChangeableRateCrowdsale *ChangeableRateCrowdsaleTransactor) Add(opts *bind.TransactOpts, _receiver common.Address, _equivalentEthAmount *big.Int) (*types.Transaction, error) {
	return _ChangeableRateCrowdsale.contract.Transact(opts, "add", _receiver, _equivalentEthAmount)
}

// Add is a paid mutator transaction binding the contract method 0xf5d82b6b.
//
// Solidity: function add(_receiver address, _equivalentEthAmount uint256) returns()
func (_ChangeableRateCrowdsale *ChangeableRateCrowdsaleSession) Add(_receiver common.Address, _equivalentEthAmount *big.Int) (*types.Transaction, error) {
	return _ChangeableRateCrowdsale.Contract.Add(&_ChangeableRateCrowdsale.TransactOpts, _receiver, _equivalentEthAmount)
}

// Add is a paid mutator transaction binding the contract method 0xf5d82b6b.
//
// Solidity: function add(_receiver address, _equivalentEthAmount uint256) returns()
func (_ChangeableRateCrowdsale *ChangeableRateCrowdsaleTransactorSession) Add(_receiver common.Address, _equivalentEthAmount *big.Int) (*types.Transaction, error) {
	return _ChangeableRateCrowdsale.Contract.Add(&_ChangeableRateCrowdsale.TransactOpts, _receiver, _equivalentEthAmount)
}

// Approve is a paid mutator transaction binding the contract method 0xdaea85c5.
//
// Solidity: function approve(_receiver address) returns()
func (_ChangeableRateCrowdsale *ChangeableRateCrowdsaleTransactor) Approve(opts *bind.TransactOpts, _receiver common.Address) (*types.Transaction, error) {
	return _ChangeableRateCrowdsale.contract.Transact(opts, "approve", _receiver)
}

// Approve is a paid mutator transaction binding the contract method 0xdaea85c5.
//
// Solidity: function approve(_receiver address) returns()
func (_ChangeableRateCrowdsale *ChangeableRateCrowdsaleSession) Approve(_receiver common.Address) (*types.Transaction, error) {
	return _ChangeableRateCrowdsale.Contract.Approve(&_ChangeableRateCrowdsale.TransactOpts, _receiver)
}

// Approve is a paid mutator transaction binding the contract method 0xdaea85c5.
//
// Solidity: function approve(_receiver address) returns()
func (_ChangeableRateCrowdsale *ChangeableRateCrowdsaleTransactorSession) Approve(_receiver common.Address) (*types.Transaction, error) {
	return _ChangeableRateCrowdsale.Contract.Approve(&_ChangeableRateCrowdsale.TransactOpts, _receiver)
}

// Buy is a paid mutator transaction binding the contract method 0xa6f2ae3a.
//
// Solidity: function buy() returns()
func (_ChangeableRateCrowdsale *ChangeableRateCrowdsaleTransactor) Buy(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ChangeableRateCrowdsale.contract.Transact(opts, "buy")
}

// Buy is a paid mutator transaction binding the contract method 0xa6f2ae3a.
//
// Solidity: function buy() returns()
func (_ChangeableRateCrowdsale *ChangeableRateCrowdsaleSession) Buy() (*types.Transaction, error) {
	return _ChangeableRateCrowdsale.Contract.Buy(&_ChangeableRateCrowdsale.TransactOpts)
}

// Buy is a paid mutator transaction binding the contract method 0xa6f2ae3a.
//
// Solidity: function buy() returns()
func (_ChangeableRateCrowdsale *ChangeableRateCrowdsaleTransactorSession) Buy() (*types.Transaction, error) {
	return _ChangeableRateCrowdsale.Contract.Buy(&_ChangeableRateCrowdsale.TransactOpts)
}

// Claim is a paid mutator transaction binding the contract method 0x4e71d92d.
//
// Solidity: function claim() returns()
func (_ChangeableRateCrowdsale *ChangeableRateCrowdsaleTransactor) Claim(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ChangeableRateCrowdsale.contract.Transact(opts, "claim")
}

// Claim is a paid mutator transaction binding the contract method 0x4e71d92d.
//
// Solidity: function claim() returns()
func (_ChangeableRateCrowdsale *ChangeableRateCrowdsaleSession) Claim() (*types.Transaction, error) {
	return _ChangeableRateCrowdsale.Contract.Claim(&_ChangeableRateCrowdsale.TransactOpts)
}

// Claim is a paid mutator transaction binding the contract method 0x4e71d92d.
//
// Solidity: function claim() returns()
func (_ChangeableRateCrowdsale *ChangeableRateCrowdsaleTransactorSession) Claim() (*types.Transaction, error) {
	return _ChangeableRateCrowdsale.Contract.Claim(&_ChangeableRateCrowdsale.TransactOpts)
}

// ClaimAll is a paid mutator transaction binding the contract method 0xd1058e59.
//
// Solidity: function claimAll() returns()
func (_ChangeableRateCrowdsale *ChangeableRateCrowdsaleTransactor) ClaimAll(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ChangeableRateCrowdsale.contract.Transact(opts, "claimAll")
}

// ClaimAll is a paid mutator transaction binding the contract method 0xd1058e59.
//
// Solidity: function claimAll() returns()
func (_ChangeableRateCrowdsale *ChangeableRateCrowdsaleSession) ClaimAll() (*types.Transaction, error) {
	return _ChangeableRateCrowdsale.Contract.ClaimAll(&_ChangeableRateCrowdsale.TransactOpts)
}

// ClaimAll is a paid mutator transaction binding the contract method 0xd1058e59.
//
// Solidity: function claimAll() returns()
func (_ChangeableRateCrowdsale *ChangeableRateCrowdsaleTransactorSession) ClaimAll() (*types.Transaction, error) {
	return _ChangeableRateCrowdsale.Contract.ClaimAll(&_ChangeableRateCrowdsale.TransactOpts)
}

// ClaimOne is a paid mutator transaction binding the contract method 0x0262dc13.
//
// Solidity: function claimOne(_receiver address) returns()
func (_ChangeableRateCrowdsale *ChangeableRateCrowdsaleTransactor) ClaimOne(opts *bind.TransactOpts, _receiver common.Address) (*types.Transaction, error) {
	return _ChangeableRateCrowdsale.contract.Transact(opts, "claimOne", _receiver)
}

// ClaimOne is a paid mutator transaction binding the contract method 0x0262dc13.
//
// Solidity: function claimOne(_receiver address) returns()
func (_ChangeableRateCrowdsale *ChangeableRateCrowdsaleSession) ClaimOne(_receiver common.Address) (*types.Transaction, error) {
	return _ChangeableRateCrowdsale.Contract.ClaimOne(&_ChangeableRateCrowdsale.TransactOpts, _receiver)
}

// ClaimOne is a paid mutator transaction binding the contract method 0x0262dc13.
//
// Solidity: function claimOne(_receiver address) returns()
func (_ChangeableRateCrowdsale *ChangeableRateCrowdsaleTransactorSession) ClaimOne(_receiver common.Address) (*types.Transaction, error) {
	return _ChangeableRateCrowdsale.Contract.ClaimOne(&_ChangeableRateCrowdsale.TransactOpts, _receiver)
}

// Finish is a paid mutator transaction binding the contract method 0xd56b2889.
//
// Solidity: function finish() returns()
func (_ChangeableRateCrowdsale *ChangeableRateCrowdsaleTransactor) Finish(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ChangeableRateCrowdsale.contract.Transact(opts, "finish")
}

// Finish is a paid mutator transaction binding the contract method 0xd56b2889.
//
// Solidity: function finish() returns()
func (_ChangeableRateCrowdsale *ChangeableRateCrowdsaleSession) Finish() (*types.Transaction, error) {
	return _ChangeableRateCrowdsale.Contract.Finish(&_ChangeableRateCrowdsale.TransactOpts)
}

// Finish is a paid mutator transaction binding the contract method 0xd56b2889.
//
// Solidity: function finish() returns()
func (_ChangeableRateCrowdsale *ChangeableRateCrowdsaleTransactorSession) Finish() (*types.Transaction, error) {
	return _ChangeableRateCrowdsale.Contract.Finish(&_ChangeableRateCrowdsale.TransactOpts)
}

// Pause is a paid mutator transaction binding the contract method 0x8456cb59.
//
// Solidity: function pause() returns(bool)
func (_ChangeableRateCrowdsale *ChangeableRateCrowdsaleTransactor) Pause(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ChangeableRateCrowdsale.contract.Transact(opts, "pause")
}

// Pause is a paid mutator transaction binding the contract method 0x8456cb59.
//
// Solidity: function pause() returns(bool)
func (_ChangeableRateCrowdsale *ChangeableRateCrowdsaleSession) Pause() (*types.Transaction, error) {
	return _ChangeableRateCrowdsale.Contract.Pause(&_ChangeableRateCrowdsale.TransactOpts)
}

// Pause is a paid mutator transaction binding the contract method 0x8456cb59.
//
// Solidity: function pause() returns(bool)
func (_ChangeableRateCrowdsale *ChangeableRateCrowdsaleTransactorSession) Pause() (*types.Transaction, error) {
	return _ChangeableRateCrowdsale.Contract.Pause(&_ChangeableRateCrowdsale.TransactOpts)
}

// SetTokenRate is a paid mutator transaction binding the contract method 0x61241c28.
//
// Solidity: function setTokenRate(_tokenRate uint256) returns()
func (_ChangeableRateCrowdsale *ChangeableRateCrowdsaleTransactor) SetTokenRate(opts *bind.TransactOpts, _tokenRate *big.Int) (*types.Transaction, error) {
	return _ChangeableRateCrowdsale.contract.Transact(opts, "setTokenRate", _tokenRate)
}

// SetTokenRate is a paid mutator transaction binding the contract method 0x61241c28.
//
// Solidity: function setTokenRate(_tokenRate uint256) returns()
func (_ChangeableRateCrowdsale *ChangeableRateCrowdsaleSession) SetTokenRate(_tokenRate *big.Int) (*types.Transaction, error) {
	return _ChangeableRateCrowdsale.Contract.SetTokenRate(&_ChangeableRateCrowdsale.TransactOpts, _tokenRate)
}

// SetTokenRate is a paid mutator transaction binding the contract method 0x61241c28.
//
// Solidity: function setTokenRate(_tokenRate uint256) returns()
func (_ChangeableRateCrowdsale *ChangeableRateCrowdsaleTransactorSession) SetTokenRate(_tokenRate *big.Int) (*types.Transaction, error) {
	return _ChangeableRateCrowdsale.Contract.SetTokenRate(&_ChangeableRateCrowdsale.TransactOpts, _tokenRate)
}

// Sub is a paid mutator transaction binding the contract method 0x26ffee08.
//
// Solidity: function sub(_receiver address, _equivalentEthAmount uint256) returns()
func (_ChangeableRateCrowdsale *ChangeableRateCrowdsaleTransactor) Sub(opts *bind.TransactOpts, _receiver common.Address, _equivalentEthAmount *big.Int) (*types.Transaction, error) {
	return _ChangeableRateCrowdsale.contract.Transact(opts, "sub", _receiver, _equivalentEthAmount)
}

// Sub is a paid mutator transaction binding the contract method 0x26ffee08.
//
// Solidity: function sub(_receiver address, _equivalentEthAmount uint256) returns()
func (_ChangeableRateCrowdsale *ChangeableRateCrowdsaleSession) Sub(_receiver common.Address, _equivalentEthAmount *big.Int) (*types.Transaction, error) {
	return _ChangeableRateCrowdsale.Contract.Sub(&_ChangeableRateCrowdsale.TransactOpts, _receiver, _equivalentEthAmount)
}

// Sub is a paid mutator transaction binding the contract method 0x26ffee08.
//
// Solidity: function sub(_receiver address, _equivalentEthAmount uint256) returns()
func (_ChangeableRateCrowdsale *ChangeableRateCrowdsaleTransactorSession) Sub(_receiver common.Address, _equivalentEthAmount *big.Int) (*types.Transaction, error) {
	return _ChangeableRateCrowdsale.Contract.Sub(&_ChangeableRateCrowdsale.TransactOpts, _receiver, _equivalentEthAmount)
}

// TopUp is a paid mutator transaction binding the contract method 0xd6f7ddf9.
//
// Solidity: function topUp(_receiver address, _equivalentEthAmount uint256) returns()
func (_ChangeableRateCrowdsale *ChangeableRateCrowdsaleTransactor) TopUp(opts *bind.TransactOpts, _receiver common.Address, _equivalentEthAmount *big.Int) (*types.Transaction, error) {
	return _ChangeableRateCrowdsale.contract.Transact(opts, "topUp", _receiver, _equivalentEthAmount)
}

// TopUp is a paid mutator transaction binding the contract method 0xd6f7ddf9.
//
// Solidity: function topUp(_receiver address, _equivalentEthAmount uint256) returns()
func (_ChangeableRateCrowdsale *ChangeableRateCrowdsaleSession) TopUp(_receiver common.Address, _equivalentEthAmount *big.Int) (*types.Transaction, error) {
	return _ChangeableRateCrowdsale.Contract.TopUp(&_ChangeableRateCrowdsale.TransactOpts, _receiver, _equivalentEthAmount)
}

// TopUp is a paid mutator transaction binding the contract method 0xd6f7ddf9.
//
// Solidity: function topUp(_receiver address, _equivalentEthAmount uint256) returns()
func (_ChangeableRateCrowdsale *ChangeableRateCrowdsaleTransactorSession) TopUp(_receiver common.Address, _equivalentEthAmount *big.Int) (*types.Transaction, error) {
	return _ChangeableRateCrowdsale.Contract.TopUp(&_ChangeableRateCrowdsale.TransactOpts, _receiver, _equivalentEthAmount)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(newOwner address) returns()
func (_ChangeableRateCrowdsale *ChangeableRateCrowdsaleTransactor) TransferOwnership(opts *bind.TransactOpts, newOwner common.Address) (*types.Transaction, error) {
	return _ChangeableRateCrowdsale.contract.Transact(opts, "transferOwnership", newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(newOwner address) returns()
func (_ChangeableRateCrowdsale *ChangeableRateCrowdsaleSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _ChangeableRateCrowdsale.Contract.TransferOwnership(&_ChangeableRateCrowdsale.TransactOpts, newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(newOwner address) returns()
func (_ChangeableRateCrowdsale *ChangeableRateCrowdsaleTransactorSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _ChangeableRateCrowdsale.Contract.TransferOwnership(&_ChangeableRateCrowdsale.TransactOpts, newOwner)
}

// Unpause is a paid mutator transaction binding the contract method 0x3f4ba83a.
//
// Solidity: function unpause() returns(bool)
func (_ChangeableRateCrowdsale *ChangeableRateCrowdsaleTransactor) Unpause(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ChangeableRateCrowdsale.contract.Transact(opts, "unpause")
}

// Unpause is a paid mutator transaction binding the contract method 0x3f4ba83a.
//
// Solidity: function unpause() returns(bool)
func (_ChangeableRateCrowdsale *ChangeableRateCrowdsaleSession) Unpause() (*types.Transaction, error) {
	return _ChangeableRateCrowdsale.Contract.Unpause(&_ChangeableRateCrowdsale.TransactOpts)
}

// Unpause is a paid mutator transaction binding the contract method 0x3f4ba83a.
//
// Solidity: function unpause() returns(bool)
func (_ChangeableRateCrowdsale *ChangeableRateCrowdsaleTransactorSession) Unpause() (*types.Transaction, error) {
	return _ChangeableRateCrowdsale.Contract.Unpause(&_ChangeableRateCrowdsale.TransactOpts)
}

// ClaimableCrowdsaleABI is the input ABI used to generate the binding from.
const ClaimableCrowdsaleABI = "[{\"constant\":false,\"inputs\":[{\"name\":\"_receiver\",\"type\":\"address\"}],\"name\":\"claimOne\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"maxTokensAmount\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"minBuyableAmount\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_receiver\",\"type\":\"address\"},{\"name\":\"_equivalentEthAmount\",\"type\":\"uint256\"}],\"name\":\"sub\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"tokenRate\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"i\",\"type\":\"uint32\"}],\"name\":\"getReceiver\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[],\"name\":\"unpause\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[],\"name\":\"claim\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"paused\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_tokenRate\",\"type\":\"uint256\"}],\"name\":\"setTokenRate\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"bankAddress\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"isFinished\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[],\"name\":\"pause\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[],\"name\":\"buy\",\"outputs\":[],\"payable\":true,\"stateMutability\":\"payable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"getReceiversCount\",\"outputs\":[{\"name\":\"\",\"type\":\"uint32\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[],\"name\":\"claimAll\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[],\"name\":\"finish\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_receiver\",\"type\":\"address\"},{\"name\":\"_equivalentEthAmount\",\"type\":\"uint256\"}],\"name\":\"topUp\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"\",\"type\":\"address\"}],\"name\":\"approved\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_receiver\",\"type\":\"address\"}],\"name\":\"approve\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"\",\"type\":\"address\"}],\"name\":\"tokens\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"issuedTokensAmount\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"transferOwnership\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_receiver\",\"type\":\"address\"},{\"name\":\"_equivalentEthAmount\",\"type\":\"uint256\"}],\"name\":\"add\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"token\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"name\":\"_tokenAddress\",\"type\":\"address\"},{\"name\":\"_bankAddress\",\"type\":\"address\"},{\"name\":\"_beneficiaryAddress\",\"type\":\"address\"},{\"name\":\"_tokenRate\",\"type\":\"uint256\"},{\"name\":\"_minBuyableAmount\",\"type\":\"uint256\"},{\"name\":\"_maxTokensAmount\",\"type\":\"uint256\"},{\"name\":\"_endDate\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"payable\":true,\"stateMutability\":\"payable\",\"type\":\"fallback\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"_buyer\",\"type\":\"address\"},{\"indexed\":false,\"name\":\"_tokens\",\"type\":\"uint256\"},{\"indexed\":false,\"name\":\"_amount\",\"type\":\"uint256\"}],\"name\":\"TokenBought\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"_receiver\",\"type\":\"address\"},{\"indexed\":false,\"name\":\"_tokens\",\"type\":\"uint256\"},{\"indexed\":false,\"name\":\"_equivalentAmount\",\"type\":\"uint256\"}],\"name\":\"TokenAdded\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"_receiver\",\"type\":\"address\"},{\"indexed\":false,\"name\":\"_tokens\",\"type\":\"uint256\"},{\"indexed\":false,\"name\":\"_equivalentAmount\",\"type\":\"uint256\"}],\"name\":\"TokenToppedUp\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"_receiver\",\"type\":\"address\"},{\"indexed\":false,\"name\":\"_tokens\",\"type\":\"uint256\"},{\"indexed\":false,\"name\":\"_equivalentAmount\",\"type\":\"uint256\"}],\"name\":\"TokenSubtracted\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"_receiver\",\"type\":\"address\"},{\"indexed\":false,\"name\":\"_tokens\",\"type\":\"uint256\"}],\"name\":\"TokenSent\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[],\"name\":\"Pause\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[],\"name\":\"Unpause\",\"type\":\"event\"}]"

// ClaimableCrowdsaleBin is the compiled bytecode used for deploying new contracts.
const ClaimableCrowdsaleBin = `0x60606040526000805460a060020a60ff02191681556005556009805460ff19169055341561002c57600080fd5b60405160e0806111108339810160405280805191906020018051919060200180519190602001805191906020018051919060200180519190602001805160008054600160a060020a03338116600160a060020a031992831617909255600380549b83169b82169b909b17909a5560028054998216998b1699909917909855600180549790981696909816959095179095556007929092556006556004555050600855611033806100dd6000396000f30060606040526004361061013a5763ffffffff60e060020a6000350416630262dc1381146101445780630a384665146101635780632511b1821461018857806326ffee081461019b57806331711884146101bd5780633c78fe07146101d05780633f4ba83a146102085780634e71d92d1461022f5780635c975abb1461024257806361241c28146102555780637822ed491461026b5780637b3529621461027e5780638456cb59146102915780638da5cb5b146102a4578063a6f2ae3a1461013a578063a72a05f7146102b7578063d1058e59146102e3578063d56b2889146102f6578063d6f7ddf914610309578063d8b964e61461032b578063daea85c51461034a578063e486033914610369578063e69b414b14610388578063f2fde38b1461039b578063f5d82b6b146103ba578063fc0c546a146103dc575b6101426103ef565b005b341561014f57600080fd5b610142600160a060020a036004351661058e565b341561016e57600080fd5b6101766105cc565b60405190815260200160405180910390f35b341561019357600080fd5b6101766105d2565b34156101a657600080fd5b610142600160a060020a03600435166024356105d8565b34156101c857600080fd5b6101766106e8565b34156101db57600080fd5b6101ec63ffffffff600435166106ee565b604051600160a060020a03909116815260200160405180910390f35b341561021357600080fd5b61021b61072c565b604051901515815260200160405180910390f35b341561023a57600080fd5b6101426107b0565b341561024d57600080fd5b61021b6107d2565b341561026057600080fd5b6101426004356107e2565b341561027657600080fd5b6101ec61080f565b341561028957600080fd5b61021b61081e565b341561029c57600080fd5b61021b610827565b34156102af57600080fd5b6101ec6108b0565b34156102c257600080fd5b6102ca6108bf565b60405163ffffffff909116815260200160405180910390f35b34156102ee57600080fd5b6101426108e8565b341561030157600080fd5b6101426109a7565b341561031457600080fd5b610142600160a060020a0360043516602435610ae9565b341561033657600080fd5b61021b600160a060020a0360043516610b9b565b341561035557600080fd5b610142600160a060020a0360043516610bb0565b341561037457600080fd5b610176600160a060020a0360043516610c06565b341561039357600080fd5b610176610c18565b34156103a657600080fd5b610142600160a060020a0360043516610c1e565b34156103c557600080fd5b610142600160a060020a0360043516602435610c75565b34156103e757600080fd5b6101ec610d56565b6009546000908190819060ff161561040657600080fd5b6004546005541061041657600080fd5b60085442111561042557600080fd5b60005460a060020a900460ff161561043c57600080fd5b60075434935060009250610456908463ffffffff610d6516565b9050600454816005540111156104a75760055460045461047b9163ffffffff610d9016565b905061049260075482610da290919063ffffffff16565b92506104a4348463ffffffff610d9016565b91505b6005546104ba908263ffffffff610db916565b60058190556004549011156104ce57600080fd5b6104d83382610dc8565b33600160a060020a03167f28cab0d660ed8aedd61a8c9db00b97f6a2d67e07d87795994f440b18bc5f1aa3828560405191825260208201526040908101905180910390a2600154600160a060020a031683156108fc0284604051600060405180830381858888f19350505050151561054f57600080fd5b600082111561058957600160a060020a03331682156108fc0283604051600060405180830381858888f19350505050151561058957600080fd5b505050565b60005433600160a060020a039081169116146105a957600080fd5b60005460a060020a900460ff16156105c057600080fd5b6105c981610eb9565b50565b60045481565b60065481565b6000805433600160a060020a039081169116146105f457600080fd5b60005460a060020a900460ff161561060b57600080fd5b60075461061e908363ffffffff610d6516565b600160a060020a0384166000908152600a60205260409020549091508190101561064757600080fd5b600160a060020a0383166000908152600a6020526040902054610670908263ffffffff610d9016565b600160a060020a0384166000908152600a602052604090205560055461069c908263ffffffff610d9016565b600555600160a060020a0383167f2da8b2becb27b189a54d21689aa110c877379a94c6f59f115cedc91d66239b33828460405191825260208201526040908101905180910390a2505050565b60075481565b6000805433600160a060020a0390811691161461070a57600080fd5b5063ffffffff166000908152600c6020526040902054600160a060020a031690565b6000805433600160a060020a0390811691161461074857600080fd5b60005460a060020a900460ff16151561076057600080fd5b6000805474ff0000000000000000000000000000000000000000191690557f7805862f689e2f13df9f062ff482ad3ad112aca9e0847911ed832e158c525b3360405160405180910390a150600190565b60005460a060020a900460ff16156107c757600080fd5b6107d033610eb9565b565b60005460a060020a900460ff1681565b60005433600160a060020a039081169116146107fd57600080fd5b6000811161080a57600080fd5b600755565b600254600160a060020a031681565b60095460ff1681565b6000805433600160a060020a0390811691161461084357600080fd5b60005460a060020a900460ff161561085a57600080fd5b6000805474ff0000000000000000000000000000000000000000191660a060020a1790557f6985a02210a168e66602d3235cb6db0e70f92b3ba4d376a33c0f3d9434bff62560405160405180910390a150600190565b600054600160a060020a031681565b6000805433600160a060020a039081169116146108db57600080fd5b50600d5463ffffffff1690565b60008054819033600160a060020a0390811691161461090657600080fd5b60005460a060020a900460ff161561091d57600080fd5b600091505b600d5463ffffffff90811690831610156109a3575063ffffffff81166000908152600c6020908152604080832054600160a060020a0316808452600b9092529091205460ff16801561098a5750600160a060020a0381166000908152600a6020526040812054115b156109985761099881610eb9565b600190910190610922565b5050565b60005433600160a060020a039081169116146109c257600080fd5b6004546005541015806109d6575060085442115b15156109e157600080fd5b60095460ff16156109f157600080fd5b6009805460ff19166001179055600354600254600160a060020a039182169163a9059cbb9116826370a082313060006040516020015260405160e060020a63ffffffff8416028152600160a060020a039091166004820152602401602060405180830381600087803b1515610a6557600080fd5b6102c65a03f11515610a7657600080fd5b5050506040518051905060006040516020015260405160e060020a63ffffffff8516028152600160a060020a0390921660048301526024820152604401602060405180830381600087803b1515610acc57600080fd5b6102c65a03f11515610add57600080fd5b50505060405180515050565b6000805433600160a060020a03908116911614610b0557600080fd5b60005460a060020a900460ff1615610b1c57600080fd5b600754610b2f908363ffffffff610d6516565b600554909150610b45908263ffffffff610db916565b600555610b528382610dc8565b82600160a060020a03167ff10143e1b37ebaf3698e98976d8bddabcea045adc939c03b245da18e81ea9651828460405191825260208201526040908101905180910390a2505050565b600b6020526000908152604090205460ff1681565b60005433600160a060020a03908116911614610bcb57600080fd5b60005460a060020a900460ff1615610be257600080fd5b600160a060020a03166000908152600b60205260409020805460ff19166001179055565b600a6020526000908152604090205481565b60055481565b60005433600160a060020a03908116911614610c3957600080fd5b600160a060020a038116156105c95760008054600160a060020a03831673ffffffffffffffffffffffffffffffffffffffff1990911617905550565b6000805433600160a060020a03908116911614610c9157600080fd5b60095460ff1615610ca157600080fd5b60045460055410610cb157600080fd5b600854421115610cc057600080fd5b60005460a060020a900460ff1615610cd757600080fd5b600754610cea908363ffffffff610d6516565b600554909150610d00908263ffffffff610db916565b600555610d0d8382610dc8565b82600160a060020a03167fa818a22273fc309f0a3682b642c74c5b5c25c0615ff378d07767cd231e19fffc828460405191825260208201526040908101905180910390a2505050565b600354600160a060020a031681565b6000828202831580610d815750828482811515610d7e57fe5b04145b1515610d8957fe5b9392505050565b600082821115610d9c57fe5b50900390565b6000808284811515610db057fe5b04949350505050565b600082820183811015610d8957fe5b60005460a060020a900460ff1615610ddf57600080fd5b600160a060020a0382166000908152600a60205260409020541515610e7057600d805463ffffffff9081166000908152600c60209081526040808320805473ffffffffffffffffffffffffffffffffffffffff1916600160a060020a038916908117909155855463ffffffff19811690861660010190951694909417909455918152600b90915220805460ff191690555b600160a060020a0382166000908152600a6020526040902054610e99908263ffffffff610db916565b600160a060020a039092166000908152600a602052604090209190915550565b6000805460a060020a900460ff1615610ed157600080fd5b600160a060020a0382166000908152600b602052604090205460ff161515610ef857600080fd5b600160a060020a0382166000908152600a602052604081205411610f1b57600080fd5b50600160a060020a038082166000908152600a60205260408082208054908390556003546002549194908116936323b872dd93929091169186918691516020015260405160e060020a63ffffffff8616028152600160a060020a0393841660048201529190921660248201526044810191909152606401602060405180830381600087803b1515610fab57600080fd5b6102c65a03f11515610fbc57600080fd5b50505060405180515050600160a060020a0382167fd4ff88d569a7ad2bfd6b17da9dbe82e2ccc016fd0051e08ff1ad1e6fe8fc9b558260405190815260200160405180910390a250505600a165627a7a723058206d3fa00f8df6a7cb9d3ba65b2815f2523e915659e71815dcee254637fbcf8ba80029`

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
// Solidity: function pause() returns(bool)
func (_ClaimableCrowdsale *ClaimableCrowdsaleTransactor) Pause(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ClaimableCrowdsale.contract.Transact(opts, "pause")
}

// Pause is a paid mutator transaction binding the contract method 0x8456cb59.
//
// Solidity: function pause() returns(bool)
func (_ClaimableCrowdsale *ClaimableCrowdsaleSession) Pause() (*types.Transaction, error) {
	return _ClaimableCrowdsale.Contract.Pause(&_ClaimableCrowdsale.TransactOpts)
}

// Pause is a paid mutator transaction binding the contract method 0x8456cb59.
//
// Solidity: function pause() returns(bool)
func (_ClaimableCrowdsale *ClaimableCrowdsaleTransactorSession) Pause() (*types.Transaction, error) {
	return _ClaimableCrowdsale.Contract.Pause(&_ClaimableCrowdsale.TransactOpts)
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

// Sub is a paid mutator transaction binding the contract method 0x26ffee08.
//
// Solidity: function sub(_receiver address, _equivalentEthAmount uint256) returns()
func (_ClaimableCrowdsale *ClaimableCrowdsaleTransactor) Sub(opts *bind.TransactOpts, _receiver common.Address, _equivalentEthAmount *big.Int) (*types.Transaction, error) {
	return _ClaimableCrowdsale.contract.Transact(opts, "sub", _receiver, _equivalentEthAmount)
}

// Sub is a paid mutator transaction binding the contract method 0x26ffee08.
//
// Solidity: function sub(_receiver address, _equivalentEthAmount uint256) returns()
func (_ClaimableCrowdsale *ClaimableCrowdsaleSession) Sub(_receiver common.Address, _equivalentEthAmount *big.Int) (*types.Transaction, error) {
	return _ClaimableCrowdsale.Contract.Sub(&_ClaimableCrowdsale.TransactOpts, _receiver, _equivalentEthAmount)
}

// Sub is a paid mutator transaction binding the contract method 0x26ffee08.
//
// Solidity: function sub(_receiver address, _equivalentEthAmount uint256) returns()
func (_ClaimableCrowdsale *ClaimableCrowdsaleTransactorSession) Sub(_receiver common.Address, _equivalentEthAmount *big.Int) (*types.Transaction, error) {
	return _ClaimableCrowdsale.Contract.Sub(&_ClaimableCrowdsale.TransactOpts, _receiver, _equivalentEthAmount)
}

// TopUp is a paid mutator transaction binding the contract method 0xd6f7ddf9.
//
// Solidity: function topUp(_receiver address, _equivalentEthAmount uint256) returns()
func (_ClaimableCrowdsale *ClaimableCrowdsaleTransactor) TopUp(opts *bind.TransactOpts, _receiver common.Address, _equivalentEthAmount *big.Int) (*types.Transaction, error) {
	return _ClaimableCrowdsale.contract.Transact(opts, "topUp", _receiver, _equivalentEthAmount)
}

// TopUp is a paid mutator transaction binding the contract method 0xd6f7ddf9.
//
// Solidity: function topUp(_receiver address, _equivalentEthAmount uint256) returns()
func (_ClaimableCrowdsale *ClaimableCrowdsaleSession) TopUp(_receiver common.Address, _equivalentEthAmount *big.Int) (*types.Transaction, error) {
	return _ClaimableCrowdsale.Contract.TopUp(&_ClaimableCrowdsale.TransactOpts, _receiver, _equivalentEthAmount)
}

// TopUp is a paid mutator transaction binding the contract method 0xd6f7ddf9.
//
// Solidity: function topUp(_receiver address, _equivalentEthAmount uint256) returns()
func (_ClaimableCrowdsale *ClaimableCrowdsaleTransactorSession) TopUp(_receiver common.Address, _equivalentEthAmount *big.Int) (*types.Transaction, error) {
	return _ClaimableCrowdsale.Contract.TopUp(&_ClaimableCrowdsale.TransactOpts, _receiver, _equivalentEthAmount)
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
// Solidity: function unpause() returns(bool)
func (_ClaimableCrowdsale *ClaimableCrowdsaleTransactor) Unpause(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ClaimableCrowdsale.contract.Transact(opts, "unpause")
}

// Unpause is a paid mutator transaction binding the contract method 0x3f4ba83a.
//
// Solidity: function unpause() returns(bool)
func (_ClaimableCrowdsale *ClaimableCrowdsaleSession) Unpause() (*types.Transaction, error) {
	return _ClaimableCrowdsale.Contract.Unpause(&_ClaimableCrowdsale.TransactOpts)
}

// Unpause is a paid mutator transaction binding the contract method 0x3f4ba83a.
//
// Solidity: function unpause() returns(bool)
func (_ClaimableCrowdsale *ClaimableCrowdsaleTransactorSession) Unpause() (*types.Transaction, error) {
	return _ClaimableCrowdsale.Contract.Unpause(&_ClaimableCrowdsale.TransactOpts)
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
const HoQuTokenABI = "[{\"constant\":true,\"inputs\":[],\"name\":\"name\",\"outputs\":[{\"name\":\"\",\"type\":\"string\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_spender\",\"type\":\"address\"},{\"name\":\"_value\",\"type\":\"uint256\"}],\"name\":\"approve\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"totalSupply\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_from\",\"type\":\"address\"},{\"name\":\"_to\",\"type\":\"address\"},{\"name\":\"_value\",\"type\":\"uint256\"}],\"name\":\"transferFrom\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"decimals\",\"outputs\":[{\"name\":\"\",\"type\":\"uint32\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[],\"name\":\"unpause\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"paused\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"_owner\",\"type\":\"address\"}],\"name\":\"balanceOf\",\"outputs\":[{\"name\":\"balance\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[],\"name\":\"pause\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"symbol\",\"outputs\":[{\"name\":\"\",\"type\":\"string\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_to\",\"type\":\"address\"},{\"name\":\"_value\",\"type\":\"uint256\"}],\"name\":\"transfer\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"_owner\",\"type\":\"address\"},{\"name\":\"_spender\",\"type\":\"address\"}],\"name\":\"allowance\",\"outputs\":[{\"name\":\"remaining\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"transferOwnership\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"name\":\"_totalSupply\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[],\"name\":\"Pause\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[],\"name\":\"Unpause\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"owner\",\"type\":\"address\"},{\"indexed\":true,\"name\":\"spender\",\"type\":\"address\"},{\"indexed\":false,\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"Approval\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"from\",\"type\":\"address\"},{\"indexed\":true,\"name\":\"to\",\"type\":\"address\"},{\"indexed\":false,\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"Transfer\",\"type\":\"event\"}]"

// HoQuTokenBin is the compiled bytecode used for deploying new contracts.
const HoQuTokenBin = `0x60606040526003805460a060020a60ff0219169055341561001f57600080fd5b6040516020806108f98339810160405280805160038054600160a060020a03191633600160a060020a03161790559150506000811161005d57600080fd5b600160a060020a033316600090815260016020526040812082905555610871806100886000396000f3006060604052600436106100cf5763ffffffff7c010000000000000000000000000000000000000000000000000000000060003504166306fdde0381146100d4578063095ea7b31461015e57806318160ddd1461019457806323b872dd146101b9578063313ce567146101e15780633f4ba83a1461020d5780635c975abb1461022057806370a08231146102335780638456cb59146102525780638da5cb5b1461026557806395d89b4114610294578063a9059cbb146102a7578063dd62ed3e146102c9578063f2fde38b146102ee575b600080fd5b34156100df57600080fd5b6100e761030f565b60405160208082528190810183818151815260200191508051906020019080838360005b8381101561012357808201518382015260200161010b565b50505050905090810190601f1680156101505780820380516001836020036101000a031916815260200191505b509250505060405180910390f35b341561016957600080fd5b610180600160a060020a0360043516602435610346565b604051901515815260200160405180910390f35b341561019f57600080fd5b6101a76103ec565b60405190815260200160405180910390f35b34156101c457600080fd5b610180600160a060020a03600435811690602435166044356103f2565b34156101ec57600080fd5b6101f461041f565b60405163ffffffff909116815260200160405180910390f35b341561021857600080fd5b610180610424565b341561022b57600080fd5b6101806104aa565b341561023e57600080fd5b6101a7600160a060020a03600435166104ba565b341561025d57600080fd5b6101806104d5565b341561027057600080fd5b610278610560565b604051600160a060020a03909116815260200160405180910390f35b341561029f57600080fd5b6100e761056f565b34156102b257600080fd5b610180600160a060020a03600435166024356105a6565b34156102d457600080fd5b6101a7600160a060020a03600435811690602435166105d1565b34156102f957600080fd5b61030d600160a060020a03600435166105fc565b005b60408051908101604052600a81527f484f515520546f6b656e00000000000000000000000000000000000000000000602082015281565b60008115806103785750600160a060020a03338116600090815260026020908152604080832093871683529290522054155b151561038357600080fd5b600160a060020a03338116600081815260026020908152604080832094881680845294909152908190208590557f8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b9259085905190815260200160405180910390a350600192915050565b60005481565b60035460009060a060020a900460ff161561040c57600080fd5b610417848484610652565b949350505050565b601281565b60035460009033600160a060020a0390811691161461044257600080fd5b60035460a060020a900460ff16151561045a57600080fd5b6003805474ff0000000000000000000000000000000000000000191690557f7805862f689e2f13df9f062ff482ad3ad112aca9e0847911ed832e158c525b3360405160405180910390a150600190565b60035460a060020a900460ff1681565b600160a060020a031660009081526001602052604090205490565b60035460009033600160a060020a039081169116146104f357600080fd5b60035460a060020a900460ff161561050a57600080fd5b6003805474ff0000000000000000000000000000000000000000191660a060020a1790557f6985a02210a168e66602d3235cb6db0e70f92b3ba4d376a33c0f3d9434bff62560405160405180910390a150600190565b600354600160a060020a031681565b60408051908101604052600381527f4851580000000000000000000000000000000000000000000000000000000000602082015281565b60035460009060a060020a900460ff16156105c057600080fd5b6105ca8383610765565b9392505050565b600160a060020a03918216600090815260026020908152604080832093909416825291909152205490565b60035433600160a060020a0390811691161461061757600080fd5b600160a060020a0381161561064f576003805473ffffffffffffffffffffffffffffffffffffffff1916600160a060020a0383161790555b50565b600160a060020a038084166000908152600260209081526040808320338516845282528083205493861683526001909152812054909190610699908463ffffffff61082416565b600160a060020a0380861660009081526001602052604080822093909355908716815220546106ce908463ffffffff61083316565b600160a060020a0386166000908152600160205260409020556106f7818463ffffffff61083316565b600160a060020a03808716600081815260026020908152604080832033861684529091529081902093909355908616917fddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef9086905190815260200160405180910390a3506001949350505050565b600160a060020a03331660009081526001602052604081205461078e908363ffffffff61083316565b600160a060020a0333811660009081526001602052604080822093909355908516815220546107c3908363ffffffff61082416565b600160a060020a0380851660008181526001602052604090819020939093559133909116907fddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef9085905190815260200160405180910390a350600192915050565b6000828201838110156105ca57fe5b60008282111561083f57fe5b509003905600a165627a7a72305820b0ae4b3f09a7123418f449dfe02d2cadeba8d150ff67bb4f860d3e96d3ab34010029`

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

// Pause is a paid mutator transaction binding the contract method 0x8456cb59.
//
// Solidity: function pause() returns(bool)
func (_HoQuToken *HoQuTokenTransactor) Pause(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _HoQuToken.contract.Transact(opts, "pause")
}

// Pause is a paid mutator transaction binding the contract method 0x8456cb59.
//
// Solidity: function pause() returns(bool)
func (_HoQuToken *HoQuTokenSession) Pause() (*types.Transaction, error) {
	return _HoQuToken.Contract.Pause(&_HoQuToken.TransactOpts)
}

// Pause is a paid mutator transaction binding the contract method 0x8456cb59.
//
// Solidity: function pause() returns(bool)
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
// Solidity: function unpause() returns(bool)
func (_HoQuToken *HoQuTokenTransactor) Unpause(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _HoQuToken.contract.Transact(opts, "unpause")
}

// Unpause is a paid mutator transaction binding the contract method 0x3f4ba83a.
//
// Solidity: function unpause() returns(bool)
func (_HoQuToken *HoQuTokenSession) Unpause() (*types.Transaction, error) {
	return _HoQuToken.Contract.Unpause(&_HoQuToken.TransactOpts)
}

// Unpause is a paid mutator transaction binding the contract method 0x3f4ba83a.
//
// Solidity: function unpause() returns(bool)
func (_HoQuToken *HoQuTokenTransactorSession) Unpause() (*types.Transaction, error) {
	return _HoQuToken.Contract.Unpause(&_HoQuToken.TransactOpts)
}

// OwnableABI is the input ABI used to generate the binding from.
const OwnableABI = "[{\"constant\":true,\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"transferOwnership\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"}]"

// OwnableBin is the compiled bytecode used for deploying new contracts.
const OwnableBin = `0x6060604052341561000f57600080fd5b60008054600160a060020a033316600160a060020a03199091161790556101268061003b6000396000f30060606040526004361060485763ffffffff7c01000000000000000000000000000000000000000000000000000000006000350416638da5cb5b8114604d578063f2fde38b146079575b600080fd5b3415605757600080fd5b605d6097565b604051600160a060020a03909116815260200160405180910390f35b3415608357600080fd5b6095600160a060020a036004351660a6565b005b600054600160a060020a031681565b60005433600160a060020a0390811691161460c057600080fd5b600160a060020a0381161560f7576000805473ffffffffffffffffffffffffffffffffffffffff1916600160a060020a0383161790555b505600a165627a7a7230582080718c8e0cb92310c1ff594709942218339b9e7aaa6f9eb3b3f48037ca9afc440029`

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
const PausableABI = "[{\"constant\":false,\"inputs\":[],\"name\":\"unpause\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"paused\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[],\"name\":\"pause\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"transferOwnership\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"anonymous\":false,\"inputs\":[],\"name\":\"Pause\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[],\"name\":\"Unpause\",\"type\":\"event\"}]"

// PausableBin is the compiled bytecode used for deploying new contracts.
const PausableBin = `0x606060405260008054600160a060020a033316600160a860020a03199091161790556102bc806100306000396000f30060606040526004361061006c5763ffffffff7c01000000000000000000000000000000000000000000000000000000006000350416633f4ba83a81146100715780635c975abb146100985780638456cb59146100ab5780638da5cb5b146100be578063f2fde38b146100ed575b600080fd5b341561007c57600080fd5b61008461010e565b604051901515815260200160405180910390f35b34156100a357600080fd5b610084610192565b34156100b657600080fd5b6100846101a2565b34156100c957600080fd5b6100d161022b565b604051600160a060020a03909116815260200160405180910390f35b34156100f857600080fd5b61010c600160a060020a036004351661023a565b005b6000805433600160a060020a0390811691161461012a57600080fd5b60005460a060020a900460ff16151561014257600080fd5b6000805474ff0000000000000000000000000000000000000000191690557f7805862f689e2f13df9f062ff482ad3ad112aca9e0847911ed832e158c525b3360405160405180910390a150600190565b60005460a060020a900460ff1681565b6000805433600160a060020a039081169116146101be57600080fd5b60005460a060020a900460ff16156101d557600080fd5b6000805474ff0000000000000000000000000000000000000000191660a060020a1790557f6985a02210a168e66602d3235cb6db0e70f92b3ba4d376a33c0f3d9434bff62560405160405180910390a150600190565b600054600160a060020a031681565b60005433600160a060020a0390811691161461025557600080fd5b600160a060020a0381161561028d576000805473ffffffffffffffffffffffffffffffffffffffff1916600160a060020a0383161790555b505600a165627a7a723058202b8ff84eed3e7e036ca8883bcd8902eb421dc2d9add268a8075494fbb9f531200029`

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
// Solidity: function pause() returns(bool)
func (_Pausable *PausableTransactor) Pause(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Pausable.contract.Transact(opts, "pause")
}

// Pause is a paid mutator transaction binding the contract method 0x8456cb59.
//
// Solidity: function pause() returns(bool)
func (_Pausable *PausableSession) Pause() (*types.Transaction, error) {
	return _Pausable.Contract.Pause(&_Pausable.TransactOpts)
}

// Pause is a paid mutator transaction binding the contract method 0x8456cb59.
//
// Solidity: function pause() returns(bool)
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
// Solidity: function unpause() returns(bool)
func (_Pausable *PausableTransactor) Unpause(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Pausable.contract.Transact(opts, "unpause")
}

// Unpause is a paid mutator transaction binding the contract method 0x3f4ba83a.
//
// Solidity: function unpause() returns(bool)
func (_Pausable *PausableSession) Unpause() (*types.Transaction, error) {
	return _Pausable.Contract.Unpause(&_Pausable.TransactOpts)
}

// Unpause is a paid mutator transaction binding the contract method 0x3f4ba83a.
//
// Solidity: function unpause() returns(bool)
func (_Pausable *PausableTransactorSession) Unpause() (*types.Transaction, error) {
	return _Pausable.Contract.Unpause(&_Pausable.TransactOpts)
}

// SafeMathABI is the input ABI used to generate the binding from.
const SafeMathABI = "[]"

// SafeMathBin is the compiled bytecode used for deploying new contracts.
const SafeMathBin = `0x60606040523415600e57600080fd5b603580601b6000396000f3006060604052600080fd00a165627a7a72305820b35624bf17cdffc8d9945ec6b1d40f683b77e1e9b5f205195eba716569991fb20029`

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
const StandardTokenABI = "[{\"constant\":false,\"inputs\":[{\"name\":\"_spender\",\"type\":\"address\"},{\"name\":\"_value\",\"type\":\"uint256\"}],\"name\":\"approve\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"totalSupply\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_from\",\"type\":\"address\"},{\"name\":\"_to\",\"type\":\"address\"},{\"name\":\"_value\",\"type\":\"uint256\"}],\"name\":\"transferFrom\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"_owner\",\"type\":\"address\"}],\"name\":\"balanceOf\",\"outputs\":[{\"name\":\"balance\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_to\",\"type\":\"address\"},{\"name\":\"_value\",\"type\":\"uint256\"}],\"name\":\"transfer\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"_owner\",\"type\":\"address\"},{\"name\":\"_spender\",\"type\":\"address\"}],\"name\":\"allowance\",\"outputs\":[{\"name\":\"remaining\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"owner\",\"type\":\"address\"},{\"indexed\":true,\"name\":\"spender\",\"type\":\"address\"},{\"indexed\":false,\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"Approval\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"from\",\"type\":\"address\"},{\"indexed\":true,\"name\":\"to\",\"type\":\"address\"},{\"indexed\":false,\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"Transfer\",\"type\":\"event\"}]"

// StandardTokenBin is the compiled bytecode used for deploying new contracts.
const StandardTokenBin = `0x6060604052341561000f57600080fd5b61047d8061001e6000396000f3006060604052600436106100775763ffffffff7c0100000000000000000000000000000000000000000000000000000000600035041663095ea7b3811461007c57806318160ddd146100b257806323b872dd146100d757806370a08231146100ff578063a9059cbb1461011e578063dd62ed3e14610140575b600080fd5b341561008757600080fd5b61009e600160a060020a0360043516602435610165565b604051901515815260200160405180910390f35b34156100bd57600080fd5b6100c561020b565b60405190815260200160405180910390f35b34156100e257600080fd5b61009e600160a060020a0360043581169060243516604435610211565b341561010a57600080fd5b6100c5600160a060020a0360043516610324565b341561012957600080fd5b61009e600160a060020a036004351660243561033f565b341561014b57600080fd5b6100c5600160a060020a03600435811690602435166103fe565b60008115806101975750600160a060020a03338116600090815260026020908152604080832093871683529290522054155b15156101a257600080fd5b600160a060020a03338116600081815260026020908152604080832094881680845294909152908190208590557f8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b9259085905190815260200160405180910390a350600192915050565b60005481565b600160a060020a038084166000908152600260209081526040808320338516845282528083205493861683526001909152812054909190610258908463ffffffff61042916565b600160a060020a03808616600090815260016020526040808220939093559087168152205461028d908463ffffffff61043f16565b600160a060020a0386166000908152600160205260409020556102b6818463ffffffff61043f16565b600160a060020a03808716600081815260026020908152604080832033861684529091529081902093909355908616917fddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef9086905190815260200160405180910390a3506001949350505050565b600160a060020a031660009081526001602052604090205490565b600160a060020a033316600090815260016020526040812054610368908363ffffffff61043f16565b600160a060020a03338116600090815260016020526040808220939093559085168152205461039d908363ffffffff61042916565b600160a060020a0380851660008181526001602052604090819020939093559133909116907fddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef9085905190815260200160405180910390a350600192915050565b600160a060020a03918216600090815260026020908152604080832093909416825291909152205490565b60008282018381101561043857fe5b9392505050565b60008282111561044b57fe5b509003905600a165627a7a72305820f5160f19df5db525591709eca55d43847b73d0963f10f70fedc25efa10c0066d0029`

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
