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

// BaseCrowdsaleABI is the input ABI used to generate the binding from.
const BaseCrowdsaleABI = "[{\"constant\":true,\"inputs\":[],\"name\":\"maxTokensAmount\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"minBuyableAmount\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[],\"name\":\"pauseToken\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"tokenRate\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[],\"name\":\"unpause\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[],\"name\":\"unpauseToken\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"paused\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_tokenRate\",\"type\":\"uint256\"}],\"name\":\"setTokenRate\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"bankAddress\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"isFinished\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[],\"name\":\"pause\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[],\"name\":\"buyTokens\",\"outputs\":[],\"payable\":true,\"stateMutability\":\"payable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_minBuyableAmount\",\"type\":\"uint256\"}],\"name\":\"setMinBuyableAmount\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[],\"name\":\"finish\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"issuedTokensAmount\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"transferOwnership\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"token\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"name\":\"_tokenAddress\",\"type\":\"address\"},{\"name\":\"_bankAddress\",\"type\":\"address\"},{\"name\":\"_beneficiaryAddress\",\"type\":\"address\"},{\"name\":\"_tokenRate\",\"type\":\"uint256\"},{\"name\":\"_minBuyableAmount\",\"type\":\"uint256\"},{\"name\":\"_maxTokensAmount\",\"type\":\"uint256\"},{\"name\":\"_endDate\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"payable\":true,\"stateMutability\":\"payable\",\"type\":\"fallback\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"buyer\",\"type\":\"address\"},{\"indexed\":false,\"name\":\"tokens\",\"type\":\"uint256\"},{\"indexed\":false,\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"TokenBought\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[],\"name\":\"Pause\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[],\"name\":\"Unpause\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"previousOwner\",\"type\":\"address\"},{\"indexed\":true,\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"OwnershipTransferred\",\"type\":\"event\"}]"

// BaseCrowdsaleBin is the compiled bytecode used for deploying new contracts.
const BaseCrowdsaleBin = `0x60606040526000805460a060020a60ff02191681556005556009805460ff19169055341561002c57600080fd5b60405160e080610c328339810160405280805191906020018051919060200180519190602001805191906020018051919060200180519190602001805160008054600160a060020a03338116600160a060020a031992831617909255600380548c8416908316179055600280548b841690831617905560018054928a1692909116919091179055600786905591506100db905083670de0b6b3a7640000640100000000610a5a61010f82021704565b6006556100fd82670de0b6b3a7640000640100000000610a5a61010f82021704565b6004556008555061013a945050505050565b600082820283158061012b575082848281151561012857fe5b04145b151561013357fe5b9392505050565b610ae9806101496000396000f300606060405236156100e05763ffffffff60e060020a6000350416630a38466581146100ea5780632511b1821461010f5780632c3496271461012257806331711884146101495780633f4ba83a1461015c57806350669a031461016f5780635c975abb1461018257806361241c28146101955780637822ed49146101ab5780637b352962146101da5780638456cb59146101ed5780638da5cb5b14610200578063d0febe4c146100e0578063d561ef2b14610213578063d56b288914610229578063e69b414b1461023c578063f2fde38b1461024f578063fc0c546a1461026e575b6100e8610281565b005b34156100f557600080fd5b6100fd61049f565b60405190815260200160405180910390f35b341561011a57600080fd5b6100fd6104a5565b341561012d57600080fd5b6101356104ab565b604051901515815260200160405180910390f35b341561015457600080fd5b6100fd61058e565b341561016757600080fd5b6100e8610594565b341561017a57600080fd5b610135610613565b341561018d57600080fd5b6101356106dc565b34156101a057600080fd5b6100e86004356106ec565b34156101b657600080fd5b6101be610748565b604051600160a060020a03909116815260200160405180910390f35b34156101e557600080fd5b610135610757565b34156101f857600080fd5b6100e8610760565b341561020b57600080fd5b6101be6107e4565b341561021e57600080fd5b6100e86004356107f3565b341561023457600080fd5b6100e8610868565b341561024757600080fd5b6100fd6109aa565b341561025a57600080fd5b6100e8600160a060020a03600435166109b0565b341561027957600080fd5b6101be610a4b565b6009546000908190819060ff161561029857600080fd5b600454600554106102a857600080fd5b6008544211156102b757600080fd5b60005460a060020a900460ff16156102ce57600080fd5b6006543410156102dd57600080fd5b600754349350600092506102f7908463ffffffff610a5a16565b9050600454816005540111156103485760055460045461031c9163ffffffff610a8516565b905061033360075482610a9790919063ffffffff16565b9250610345348463ffffffff610a8516565b91505b60055461035b908263ffffffff610aae16565b600581905560045490111561036f57600080fd5b600354600160a060020a031663a9059cbb338360006040516020015260405160e060020a63ffffffff8516028152600160a060020a0390921660048301526024820152604401602060405180830381600087803b15156103ce57600080fd5b6102c65a03f115156103df57600080fd5b50505060405180515050600160a060020a0333167f28cab0d660ed8aedd61a8c9db00b97f6a2d67e07d87795994f440b18bc5f1aa3828560405191825260208201526040908101905180910390a2600154600160a060020a031683156108fc0284604051600060405180830381858888f19350505050151561046057600080fd5b600082111561049a57600160a060020a03331682156108fc0283604051600060405180830381858888f19350505050151561049a57600080fd5b505050565b60045481565b60065481565b6000805433600160a060020a039081169116146104c757600080fd5b600354600160a060020a0316635c975abb6000604051602001526040518163ffffffff1660e060020a028152600401602060405180830381600087803b151561050f57600080fd5b6102c65a03f1151561052057600080fd5b505050604051805115905061053457600080fd5b600354600160a060020a0316638456cb596040518163ffffffff1660e060020a028152600401600060405180830381600087803b151561057357600080fd5b6102c65a03f1151561058457600080fd5b5050506001905090565b60075481565b60005433600160a060020a039081169116146105af57600080fd5b60005460a060020a900460ff1615156105c757600080fd5b6000805474ff0000000000000000000000000000000000000000191690557f7805862f689e2f13df9f062ff482ad3ad112aca9e0847911ed832e158c525b3360405160405180910390a1565b6000805433600160a060020a0390811691161461062f57600080fd5b600354600160a060020a0316635c975abb6000604051602001526040518163ffffffff1660e060020a028152600401602060405180830381600087803b151561067757600080fd5b6102c65a03f1151561068857600080fd5b50505060405180519050151561069d57600080fd5b600354600160a060020a0316633f4ba83a6040518163ffffffff1660e060020a028152600401600060405180830381600087803b151561057357600080fd5b60005460a060020a900460ff1681565b60005433600160a060020a0390811691161461070757600080fd5b60095460ff161561071757600080fd5b6004546005541061072757600080fd5b60085442111561073657600080fd5b6000811161074357600080fd5b600755565b600254600160a060020a031681565b60095460ff1681565b60005433600160a060020a0390811691161461077b57600080fd5b60005460a060020a900460ff161561079257600080fd5b6000805474ff0000000000000000000000000000000000000000191660a060020a1790557f6985a02210a168e66602d3235cb6db0e70f92b3ba4d376a33c0f3d9434bff62560405160405180910390a1565b600054600160a060020a031681565b60005433600160a060020a0390811691161461080e57600080fd5b60095460ff161561081e57600080fd5b6004546005541061082e57600080fd5b60085442111561083d57600080fd5b6000811161084a57600080fd5b61086281670de0b6b3a764000063ffffffff610a5a16565b60065550565b60005433600160a060020a0390811691161461088357600080fd5b600454600554101580610897575060085442115b15156108a257600080fd5b60095460ff16156108b257600080fd5b6009805460ff19166001179055600354600254600160a060020a039182169163a9059cbb9116826370a082313060006040516020015260405160e060020a63ffffffff8416028152600160a060020a039091166004820152602401602060405180830381600087803b151561092657600080fd5b6102c65a03f1151561093757600080fd5b5050506040518051905060006040516020015260405160e060020a63ffffffff8516028152600160a060020a0390921660048301526024820152604401602060405180830381600087803b151561098d57600080fd5b6102c65a03f1151561099e57600080fd5b50505060405180515050565b60055481565b60005433600160a060020a039081169116146109cb57600080fd5b600160a060020a03811615156109e057600080fd5b600054600160a060020a0380831691167f8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e060405160405180910390a36000805473ffffffffffffffffffffffffffffffffffffffff1916600160a060020a0392909216919091179055565b600354600160a060020a031681565b6000828202831580610a765750828482811515610a7357fe5b04145b1515610a7e57fe5b9392505050565b600082821115610a9157fe5b50900390565b6000808284811515610aa557fe5b04949350505050565b600082820183811015610a7e57fe00a165627a7a7230582014d022b39656dd16f00b9655044cd0dd125c9986595165c4052592623577a6470029`

// DeployBaseCrowdsale deploys a new Ethereum contract, binding an instance of BaseCrowdsale to it.
func DeployBaseCrowdsale(auth *bind.TransactOpts, backend bind.ContractBackend, _tokenAddress common.Address, _bankAddress common.Address, _beneficiaryAddress common.Address, _tokenRate *big.Int, _minBuyableAmount *big.Int, _maxTokensAmount *big.Int, _endDate *big.Int) (common.Address, *types.Transaction, *BaseCrowdsale, error) {
	parsed, err := abi.JSON(strings.NewReader(BaseCrowdsaleABI))
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	address, tx, contract, err := bind.DeployContract(auth, parsed, common.FromHex(BaseCrowdsaleBin), backend, _tokenAddress, _bankAddress, _beneficiaryAddress, _tokenRate, _minBuyableAmount, _maxTokensAmount, _endDate)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &BaseCrowdsale{BaseCrowdsaleCaller: BaseCrowdsaleCaller{contract: contract}, BaseCrowdsaleTransactor: BaseCrowdsaleTransactor{contract: contract}}, nil
}

// BaseCrowdsale is an auto generated Go binding around an Ethereum contract.
type BaseCrowdsale struct {
	BaseCrowdsaleCaller     // Read-only binding to the contract
	BaseCrowdsaleTransactor // Write-only binding to the contract
}

// BaseCrowdsaleCaller is an auto generated read-only Go binding around an Ethereum contract.
type BaseCrowdsaleCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// BaseCrowdsaleTransactor is an auto generated write-only Go binding around an Ethereum contract.
type BaseCrowdsaleTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// BaseCrowdsaleSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type BaseCrowdsaleSession struct {
	Contract     *BaseCrowdsale    // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// BaseCrowdsaleCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type BaseCrowdsaleCallerSession struct {
	Contract *BaseCrowdsaleCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts        // Call options to use throughout this session
}

// BaseCrowdsaleTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type BaseCrowdsaleTransactorSession struct {
	Contract     *BaseCrowdsaleTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts        // Transaction auth options to use throughout this session
}

// BaseCrowdsaleRaw is an auto generated low-level Go binding around an Ethereum contract.
type BaseCrowdsaleRaw struct {
	Contract *BaseCrowdsale // Generic contract binding to access the raw methods on
}

// BaseCrowdsaleCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type BaseCrowdsaleCallerRaw struct {
	Contract *BaseCrowdsaleCaller // Generic read-only contract binding to access the raw methods on
}

// BaseCrowdsaleTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type BaseCrowdsaleTransactorRaw struct {
	Contract *BaseCrowdsaleTransactor // Generic write-only contract binding to access the raw methods on
}

// NewBaseCrowdsale creates a new instance of BaseCrowdsale, bound to a specific deployed contract.
func NewBaseCrowdsale(address common.Address, backend bind.ContractBackend) (*BaseCrowdsale, error) {
	contract, err := bindBaseCrowdsale(address, backend, backend)
	if err != nil {
		return nil, err
	}
	return &BaseCrowdsale{BaseCrowdsaleCaller: BaseCrowdsaleCaller{contract: contract}, BaseCrowdsaleTransactor: BaseCrowdsaleTransactor{contract: contract}}, nil
}

// NewBaseCrowdsaleCaller creates a new read-only instance of BaseCrowdsale, bound to a specific deployed contract.
func NewBaseCrowdsaleCaller(address common.Address, caller bind.ContractCaller) (*BaseCrowdsaleCaller, error) {
	contract, err := bindBaseCrowdsale(address, caller, nil)
	if err != nil {
		return nil, err
	}
	return &BaseCrowdsaleCaller{contract: contract}, nil
}

// NewBaseCrowdsaleTransactor creates a new write-only instance of BaseCrowdsale, bound to a specific deployed contract.
func NewBaseCrowdsaleTransactor(address common.Address, transactor bind.ContractTransactor) (*BaseCrowdsaleTransactor, error) {
	contract, err := bindBaseCrowdsale(address, nil, transactor)
	if err != nil {
		return nil, err
	}
	return &BaseCrowdsaleTransactor{contract: contract}, nil
}

// bindBaseCrowdsale binds a generic wrapper to an already deployed contract.
func bindBaseCrowdsale(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(BaseCrowdsaleABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_BaseCrowdsale *BaseCrowdsaleRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _BaseCrowdsale.Contract.BaseCrowdsaleCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_BaseCrowdsale *BaseCrowdsaleRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _BaseCrowdsale.Contract.BaseCrowdsaleTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_BaseCrowdsale *BaseCrowdsaleRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _BaseCrowdsale.Contract.BaseCrowdsaleTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_BaseCrowdsale *BaseCrowdsaleCallerRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _BaseCrowdsale.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_BaseCrowdsale *BaseCrowdsaleTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _BaseCrowdsale.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_BaseCrowdsale *BaseCrowdsaleTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _BaseCrowdsale.Contract.contract.Transact(opts, method, params...)
}

// BankAddress is a free data retrieval call binding the contract method 0x7822ed49.
//
// Solidity: function bankAddress() constant returns(address)
func (_BaseCrowdsale *BaseCrowdsaleCaller) BankAddress(opts *bind.CallOpts) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _BaseCrowdsale.contract.Call(opts, out, "bankAddress")
	return *ret0, err
}

// BankAddress is a free data retrieval call binding the contract method 0x7822ed49.
//
// Solidity: function bankAddress() constant returns(address)
func (_BaseCrowdsale *BaseCrowdsaleSession) BankAddress() (common.Address, error) {
	return _BaseCrowdsale.Contract.BankAddress(&_BaseCrowdsale.CallOpts)
}

// BankAddress is a free data retrieval call binding the contract method 0x7822ed49.
//
// Solidity: function bankAddress() constant returns(address)
func (_BaseCrowdsale *BaseCrowdsaleCallerSession) BankAddress() (common.Address, error) {
	return _BaseCrowdsale.Contract.BankAddress(&_BaseCrowdsale.CallOpts)
}

// IsFinished is a free data retrieval call binding the contract method 0x7b352962.
//
// Solidity: function isFinished() constant returns(bool)
func (_BaseCrowdsale *BaseCrowdsaleCaller) IsFinished(opts *bind.CallOpts) (bool, error) {
	var (
		ret0 = new(bool)
	)
	out := ret0
	err := _BaseCrowdsale.contract.Call(opts, out, "isFinished")
	return *ret0, err
}

// IsFinished is a free data retrieval call binding the contract method 0x7b352962.
//
// Solidity: function isFinished() constant returns(bool)
func (_BaseCrowdsale *BaseCrowdsaleSession) IsFinished() (bool, error) {
	return _BaseCrowdsale.Contract.IsFinished(&_BaseCrowdsale.CallOpts)
}

// IsFinished is a free data retrieval call binding the contract method 0x7b352962.
//
// Solidity: function isFinished() constant returns(bool)
func (_BaseCrowdsale *BaseCrowdsaleCallerSession) IsFinished() (bool, error) {
	return _BaseCrowdsale.Contract.IsFinished(&_BaseCrowdsale.CallOpts)
}

// IssuedTokensAmount is a free data retrieval call binding the contract method 0xe69b414b.
//
// Solidity: function issuedTokensAmount() constant returns(uint256)
func (_BaseCrowdsale *BaseCrowdsaleCaller) IssuedTokensAmount(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _BaseCrowdsale.contract.Call(opts, out, "issuedTokensAmount")
	return *ret0, err
}

// IssuedTokensAmount is a free data retrieval call binding the contract method 0xe69b414b.
//
// Solidity: function issuedTokensAmount() constant returns(uint256)
func (_BaseCrowdsale *BaseCrowdsaleSession) IssuedTokensAmount() (*big.Int, error) {
	return _BaseCrowdsale.Contract.IssuedTokensAmount(&_BaseCrowdsale.CallOpts)
}

// IssuedTokensAmount is a free data retrieval call binding the contract method 0xe69b414b.
//
// Solidity: function issuedTokensAmount() constant returns(uint256)
func (_BaseCrowdsale *BaseCrowdsaleCallerSession) IssuedTokensAmount() (*big.Int, error) {
	return _BaseCrowdsale.Contract.IssuedTokensAmount(&_BaseCrowdsale.CallOpts)
}

// MaxTokensAmount is a free data retrieval call binding the contract method 0x0a384665.
//
// Solidity: function maxTokensAmount() constant returns(uint256)
func (_BaseCrowdsale *BaseCrowdsaleCaller) MaxTokensAmount(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _BaseCrowdsale.contract.Call(opts, out, "maxTokensAmount")
	return *ret0, err
}

// MaxTokensAmount is a free data retrieval call binding the contract method 0x0a384665.
//
// Solidity: function maxTokensAmount() constant returns(uint256)
func (_BaseCrowdsale *BaseCrowdsaleSession) MaxTokensAmount() (*big.Int, error) {
	return _BaseCrowdsale.Contract.MaxTokensAmount(&_BaseCrowdsale.CallOpts)
}

// MaxTokensAmount is a free data retrieval call binding the contract method 0x0a384665.
//
// Solidity: function maxTokensAmount() constant returns(uint256)
func (_BaseCrowdsale *BaseCrowdsaleCallerSession) MaxTokensAmount() (*big.Int, error) {
	return _BaseCrowdsale.Contract.MaxTokensAmount(&_BaseCrowdsale.CallOpts)
}

// MinBuyableAmount is a free data retrieval call binding the contract method 0x2511b182.
//
// Solidity: function minBuyableAmount() constant returns(uint256)
func (_BaseCrowdsale *BaseCrowdsaleCaller) MinBuyableAmount(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _BaseCrowdsale.contract.Call(opts, out, "minBuyableAmount")
	return *ret0, err
}

// MinBuyableAmount is a free data retrieval call binding the contract method 0x2511b182.
//
// Solidity: function minBuyableAmount() constant returns(uint256)
func (_BaseCrowdsale *BaseCrowdsaleSession) MinBuyableAmount() (*big.Int, error) {
	return _BaseCrowdsale.Contract.MinBuyableAmount(&_BaseCrowdsale.CallOpts)
}

// MinBuyableAmount is a free data retrieval call binding the contract method 0x2511b182.
//
// Solidity: function minBuyableAmount() constant returns(uint256)
func (_BaseCrowdsale *BaseCrowdsaleCallerSession) MinBuyableAmount() (*big.Int, error) {
	return _BaseCrowdsale.Contract.MinBuyableAmount(&_BaseCrowdsale.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() constant returns(address)
func (_BaseCrowdsale *BaseCrowdsaleCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _BaseCrowdsale.contract.Call(opts, out, "owner")
	return *ret0, err
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() constant returns(address)
func (_BaseCrowdsale *BaseCrowdsaleSession) Owner() (common.Address, error) {
	return _BaseCrowdsale.Contract.Owner(&_BaseCrowdsale.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() constant returns(address)
func (_BaseCrowdsale *BaseCrowdsaleCallerSession) Owner() (common.Address, error) {
	return _BaseCrowdsale.Contract.Owner(&_BaseCrowdsale.CallOpts)
}

// Paused is a free data retrieval call binding the contract method 0x5c975abb.
//
// Solidity: function paused() constant returns(bool)
func (_BaseCrowdsale *BaseCrowdsaleCaller) Paused(opts *bind.CallOpts) (bool, error) {
	var (
		ret0 = new(bool)
	)
	out := ret0
	err := _BaseCrowdsale.contract.Call(opts, out, "paused")
	return *ret0, err
}

// Paused is a free data retrieval call binding the contract method 0x5c975abb.
//
// Solidity: function paused() constant returns(bool)
func (_BaseCrowdsale *BaseCrowdsaleSession) Paused() (bool, error) {
	return _BaseCrowdsale.Contract.Paused(&_BaseCrowdsale.CallOpts)
}

// Paused is a free data retrieval call binding the contract method 0x5c975abb.
//
// Solidity: function paused() constant returns(bool)
func (_BaseCrowdsale *BaseCrowdsaleCallerSession) Paused() (bool, error) {
	return _BaseCrowdsale.Contract.Paused(&_BaseCrowdsale.CallOpts)
}

// Token is a free data retrieval call binding the contract method 0xfc0c546a.
//
// Solidity: function token() constant returns(address)
func (_BaseCrowdsale *BaseCrowdsaleCaller) Token(opts *bind.CallOpts) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _BaseCrowdsale.contract.Call(opts, out, "token")
	return *ret0, err
}

// Token is a free data retrieval call binding the contract method 0xfc0c546a.
//
// Solidity: function token() constant returns(address)
func (_BaseCrowdsale *BaseCrowdsaleSession) Token() (common.Address, error) {
	return _BaseCrowdsale.Contract.Token(&_BaseCrowdsale.CallOpts)
}

// Token is a free data retrieval call binding the contract method 0xfc0c546a.
//
// Solidity: function token() constant returns(address)
func (_BaseCrowdsale *BaseCrowdsaleCallerSession) Token() (common.Address, error) {
	return _BaseCrowdsale.Contract.Token(&_BaseCrowdsale.CallOpts)
}

// TokenRate is a free data retrieval call binding the contract method 0x31711884.
//
// Solidity: function tokenRate() constant returns(uint256)
func (_BaseCrowdsale *BaseCrowdsaleCaller) TokenRate(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _BaseCrowdsale.contract.Call(opts, out, "tokenRate")
	return *ret0, err
}

// TokenRate is a free data retrieval call binding the contract method 0x31711884.
//
// Solidity: function tokenRate() constant returns(uint256)
func (_BaseCrowdsale *BaseCrowdsaleSession) TokenRate() (*big.Int, error) {
	return _BaseCrowdsale.Contract.TokenRate(&_BaseCrowdsale.CallOpts)
}

// TokenRate is a free data retrieval call binding the contract method 0x31711884.
//
// Solidity: function tokenRate() constant returns(uint256)
func (_BaseCrowdsale *BaseCrowdsaleCallerSession) TokenRate() (*big.Int, error) {
	return _BaseCrowdsale.Contract.TokenRate(&_BaseCrowdsale.CallOpts)
}

// BuyTokens is a paid mutator transaction binding the contract method 0xd0febe4c.
//
// Solidity: function buyTokens() returns()
func (_BaseCrowdsale *BaseCrowdsaleTransactor) BuyTokens(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _BaseCrowdsale.contract.Transact(opts, "buyTokens")
}

// BuyTokens is a paid mutator transaction binding the contract method 0xd0febe4c.
//
// Solidity: function buyTokens() returns()
func (_BaseCrowdsale *BaseCrowdsaleSession) BuyTokens() (*types.Transaction, error) {
	return _BaseCrowdsale.Contract.BuyTokens(&_BaseCrowdsale.TransactOpts)
}

// BuyTokens is a paid mutator transaction binding the contract method 0xd0febe4c.
//
// Solidity: function buyTokens() returns()
func (_BaseCrowdsale *BaseCrowdsaleTransactorSession) BuyTokens() (*types.Transaction, error) {
	return _BaseCrowdsale.Contract.BuyTokens(&_BaseCrowdsale.TransactOpts)
}

// Finish is a paid mutator transaction binding the contract method 0xd56b2889.
//
// Solidity: function finish() returns()
func (_BaseCrowdsale *BaseCrowdsaleTransactor) Finish(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _BaseCrowdsale.contract.Transact(opts, "finish")
}

// Finish is a paid mutator transaction binding the contract method 0xd56b2889.
//
// Solidity: function finish() returns()
func (_BaseCrowdsale *BaseCrowdsaleSession) Finish() (*types.Transaction, error) {
	return _BaseCrowdsale.Contract.Finish(&_BaseCrowdsale.TransactOpts)
}

// Finish is a paid mutator transaction binding the contract method 0xd56b2889.
//
// Solidity: function finish() returns()
func (_BaseCrowdsale *BaseCrowdsaleTransactorSession) Finish() (*types.Transaction, error) {
	return _BaseCrowdsale.Contract.Finish(&_BaseCrowdsale.TransactOpts)
}

// Pause is a paid mutator transaction binding the contract method 0x8456cb59.
//
// Solidity: function pause() returns()
func (_BaseCrowdsale *BaseCrowdsaleTransactor) Pause(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _BaseCrowdsale.contract.Transact(opts, "pause")
}

// Pause is a paid mutator transaction binding the contract method 0x8456cb59.
//
// Solidity: function pause() returns()
func (_BaseCrowdsale *BaseCrowdsaleSession) Pause() (*types.Transaction, error) {
	return _BaseCrowdsale.Contract.Pause(&_BaseCrowdsale.TransactOpts)
}

// Pause is a paid mutator transaction binding the contract method 0x8456cb59.
//
// Solidity: function pause() returns()
func (_BaseCrowdsale *BaseCrowdsaleTransactorSession) Pause() (*types.Transaction, error) {
	return _BaseCrowdsale.Contract.Pause(&_BaseCrowdsale.TransactOpts)
}

// PauseToken is a paid mutator transaction binding the contract method 0x2c349627.
//
// Solidity: function pauseToken() returns(bool)
func (_BaseCrowdsale *BaseCrowdsaleTransactor) PauseToken(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _BaseCrowdsale.contract.Transact(opts, "pauseToken")
}

// PauseToken is a paid mutator transaction binding the contract method 0x2c349627.
//
// Solidity: function pauseToken() returns(bool)
func (_BaseCrowdsale *BaseCrowdsaleSession) PauseToken() (*types.Transaction, error) {
	return _BaseCrowdsale.Contract.PauseToken(&_BaseCrowdsale.TransactOpts)
}

// PauseToken is a paid mutator transaction binding the contract method 0x2c349627.
//
// Solidity: function pauseToken() returns(bool)
func (_BaseCrowdsale *BaseCrowdsaleTransactorSession) PauseToken() (*types.Transaction, error) {
	return _BaseCrowdsale.Contract.PauseToken(&_BaseCrowdsale.TransactOpts)
}

// SetMinBuyableAmount is a paid mutator transaction binding the contract method 0xd561ef2b.
//
// Solidity: function setMinBuyableAmount(_minBuyableAmount uint256) returns()
func (_BaseCrowdsale *BaseCrowdsaleTransactor) SetMinBuyableAmount(opts *bind.TransactOpts, _minBuyableAmount *big.Int) (*types.Transaction, error) {
	return _BaseCrowdsale.contract.Transact(opts, "setMinBuyableAmount", _minBuyableAmount)
}

// SetMinBuyableAmount is a paid mutator transaction binding the contract method 0xd561ef2b.
//
// Solidity: function setMinBuyableAmount(_minBuyableAmount uint256) returns()
func (_BaseCrowdsale *BaseCrowdsaleSession) SetMinBuyableAmount(_minBuyableAmount *big.Int) (*types.Transaction, error) {
	return _BaseCrowdsale.Contract.SetMinBuyableAmount(&_BaseCrowdsale.TransactOpts, _minBuyableAmount)
}

// SetMinBuyableAmount is a paid mutator transaction binding the contract method 0xd561ef2b.
//
// Solidity: function setMinBuyableAmount(_minBuyableAmount uint256) returns()
func (_BaseCrowdsale *BaseCrowdsaleTransactorSession) SetMinBuyableAmount(_minBuyableAmount *big.Int) (*types.Transaction, error) {
	return _BaseCrowdsale.Contract.SetMinBuyableAmount(&_BaseCrowdsale.TransactOpts, _minBuyableAmount)
}

// SetTokenRate is a paid mutator transaction binding the contract method 0x61241c28.
//
// Solidity: function setTokenRate(_tokenRate uint256) returns()
func (_BaseCrowdsale *BaseCrowdsaleTransactor) SetTokenRate(opts *bind.TransactOpts, _tokenRate *big.Int) (*types.Transaction, error) {
	return _BaseCrowdsale.contract.Transact(opts, "setTokenRate", _tokenRate)
}

// SetTokenRate is a paid mutator transaction binding the contract method 0x61241c28.
//
// Solidity: function setTokenRate(_tokenRate uint256) returns()
func (_BaseCrowdsale *BaseCrowdsaleSession) SetTokenRate(_tokenRate *big.Int) (*types.Transaction, error) {
	return _BaseCrowdsale.Contract.SetTokenRate(&_BaseCrowdsale.TransactOpts, _tokenRate)
}

// SetTokenRate is a paid mutator transaction binding the contract method 0x61241c28.
//
// Solidity: function setTokenRate(_tokenRate uint256) returns()
func (_BaseCrowdsale *BaseCrowdsaleTransactorSession) SetTokenRate(_tokenRate *big.Int) (*types.Transaction, error) {
	return _BaseCrowdsale.Contract.SetTokenRate(&_BaseCrowdsale.TransactOpts, _tokenRate)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(newOwner address) returns()
func (_BaseCrowdsale *BaseCrowdsaleTransactor) TransferOwnership(opts *bind.TransactOpts, newOwner common.Address) (*types.Transaction, error) {
	return _BaseCrowdsale.contract.Transact(opts, "transferOwnership", newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(newOwner address) returns()
func (_BaseCrowdsale *BaseCrowdsaleSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _BaseCrowdsale.Contract.TransferOwnership(&_BaseCrowdsale.TransactOpts, newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(newOwner address) returns()
func (_BaseCrowdsale *BaseCrowdsaleTransactorSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _BaseCrowdsale.Contract.TransferOwnership(&_BaseCrowdsale.TransactOpts, newOwner)
}

// Unpause is a paid mutator transaction binding the contract method 0x3f4ba83a.
//
// Solidity: function unpause() returns()
func (_BaseCrowdsale *BaseCrowdsaleTransactor) Unpause(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _BaseCrowdsale.contract.Transact(opts, "unpause")
}

// Unpause is a paid mutator transaction binding the contract method 0x3f4ba83a.
//
// Solidity: function unpause() returns()
func (_BaseCrowdsale *BaseCrowdsaleSession) Unpause() (*types.Transaction, error) {
	return _BaseCrowdsale.Contract.Unpause(&_BaseCrowdsale.TransactOpts)
}

// Unpause is a paid mutator transaction binding the contract method 0x3f4ba83a.
//
// Solidity: function unpause() returns()
func (_BaseCrowdsale *BaseCrowdsaleTransactorSession) Unpause() (*types.Transaction, error) {
	return _BaseCrowdsale.Contract.Unpause(&_BaseCrowdsale.TransactOpts)
}

// UnpauseToken is a paid mutator transaction binding the contract method 0x50669a03.
//
// Solidity: function unpauseToken() returns(bool)
func (_BaseCrowdsale *BaseCrowdsaleTransactor) UnpauseToken(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _BaseCrowdsale.contract.Transact(opts, "unpauseToken")
}

// UnpauseToken is a paid mutator transaction binding the contract method 0x50669a03.
//
// Solidity: function unpauseToken() returns(bool)
func (_BaseCrowdsale *BaseCrowdsaleSession) UnpauseToken() (*types.Transaction, error) {
	return _BaseCrowdsale.Contract.UnpauseToken(&_BaseCrowdsale.TransactOpts)
}

// UnpauseToken is a paid mutator transaction binding the contract method 0x50669a03.
//
// Solidity: function unpauseToken() returns(bool)
func (_BaseCrowdsale *BaseCrowdsaleTransactorSession) UnpauseToken() (*types.Transaction, error) {
	return _BaseCrowdsale.Contract.UnpauseToken(&_BaseCrowdsale.TransactOpts)
}

// PrivatePlacementABI is the input ABI used to generate the binding from.
const PrivatePlacementABI = "[{\"constant\":true,\"inputs\":[],\"name\":\"maxTokensAmount\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"supportAddress\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[],\"name\":\"allocateInternalWallets\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"totalSupply\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"minBuyableAmount\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[],\"name\":\"pauseToken\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"tokenRate\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[],\"name\":\"unpause\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"initialFoundersAmount\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[],\"name\":\"unpauseToken\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"paused\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_tokenRate\",\"type\":\"uint256\"}],\"name\":\"setTokenRate\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"bankAddress\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"isFinished\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[],\"name\":\"pause\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"initialSupportAmount\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"bountyAddress\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[],\"name\":\"buyTokens\",\"outputs\":[],\"payable\":true,\"stateMutability\":\"payable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_minBuyableAmount\",\"type\":\"uint256\"}],\"name\":\"setMinBuyableAmount\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[],\"name\":\"finish\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"issuedTokensAmount\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"foundersAddress\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"transferOwnership\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"initialBountyAmount\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"token\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"name\":\"_bankAddress\",\"type\":\"address\"},{\"name\":\"_foundersAddress\",\"type\":\"address\"},{\"name\":\"_supportAddress\",\"type\":\"address\"},{\"name\":\"_bountyAddress\",\"type\":\"address\"},{\"name\":\"_beneficiaryAddress\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"payable\":true,\"stateMutability\":\"payable\",\"type\":\"fallback\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"buyer\",\"type\":\"address\"},{\"indexed\":false,\"name\":\"tokens\",\"type\":\"uint256\"},{\"indexed\":false,\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"TokenBought\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[],\"name\":\"Pause\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[],\"name\":\"Unpause\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"previousOwner\",\"type\":\"address\"},{\"indexed\":true,\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"OwnershipTransferred\",\"type\":\"event\"}]"

// PrivatePlacementBin is the compiled bytecode used for deploying new contracts.
const PrivatePlacementBin = `0x60606040526000805460a060020a60ff021990811682556005919091556009805460ff19169055600b8054909116905534156200003b57600080fd5b60405160a080620027b0833981016040528080519190602001805191906020018051919060200180519190602001805191506200009490506b02df455b08e99ea2b300000064010000000062000e006200019a82021704565b60008054600160a060020a03338116600160a060020a0319928316179092556003805484841690831617905560028054838a1690831617905560018054928516929091169190911790556127106007819055869083906064630160a5b06359d030006200011883670de0b6b3a764000064010000000062000d9d620001ca82021704565b6006556200013d82670de0b6b3a764000064010000000062000d9d620001ca82021704565b60045560085550506009805461010060a860020a031916610100600160a060020a039a8b16021790555050600a8054600160a060020a03199081169688169690961790555050600b80549093169190931617905550620002099050565b600081620001a7620001f8565b908152602001604051809103906000f0801515620001c457600080fd5b92915050565b6000828202831580620001e85750828482811515620001e557fe5b04145b1515620001f157fe5b9392505050565b604051610b978062001c1983390190565b611a0080620002196000396000f300606060405236156101385763ffffffff60e060020a6000350416630a38466581146101425780630f6d37d6146101675780631135b3ac1461019657806318160ddd146101a95780632511b182146101bc5780632c349627146101cf57806331711884146101f65780633f4ba83a146102095780634d39ed061461021c57806350669a031461022f5780635c975abb1461024257806361241c28146102555780637822ed491461026b5780637b3529621461027e5780638456cb59146102915780638da5cb5b146102a4578063ab28c704146102b7578063c516358f146102ca578063d0febe4c14610138578063d561ef2b146102dd578063d56b2889146102f3578063e69b414b14610306578063ebfdc65714610319578063f2fde38b1461032c578063f338c9841461034b578063fc0c546a1461035e575b610140610371565b005b341561014d57600080fd5b61015561058f565b60405190815260200160405180910390f35b341561017257600080fd5b61017a610595565b604051600160a060020a03909116815260200160405180910390f35b34156101a157600080fd5b6101406105a4565b34156101b457600080fd5b6101556107a5565b34156101c757600080fd5b6101556107b5565b34156101da57600080fd5b6101e26107bb565b604051901515815260200160405180910390f35b341561020157600080fd5b61015561089e565b341561021457600080fd5b6101406108a4565b341561022757600080fd5b610155610923565b341561023a57600080fd5b6101e2610932565b341561024d57600080fd5b6101e26109fb565b341561026057600080fd5b610140600435610a0b565b341561027657600080fd5b61017a610a67565b341561028957600080fd5b6101e2610a76565b341561029c57600080fd5b610140610a7f565b34156102af57600080fd5b61017a610b03565b34156102c257600080fd5b610155610b12565b34156102d557600080fd5b61017a610b21565b34156102e857600080fd5b610140600435610b30565b34156102fe57600080fd5b610140610ba5565b341561031157600080fd5b610155610cca565b341561032457600080fd5b61017a610cd0565b341561033757600080fd5b610140600160a060020a0360043516610ce4565b341561035657600080fd5b610155610d7f565b341561036957600080fd5b61017a610d8e565b6009546000908190819060ff161561038857600080fd5b6004546005541061039857600080fd5b6008544211156103a757600080fd5b60005460a060020a900460ff16156103be57600080fd5b6006543410156103cd57600080fd5b600754349350600092506103e7908463ffffffff610d9d16565b9050600454816005540111156104385760055460045461040c9163ffffffff610dc816565b905061042360075482610dda90919063ffffffff16565b9250610435348463ffffffff610dc816565b91505b60055461044b908263ffffffff610df116565b600581905560045490111561045f57600080fd5b600354600160a060020a031663a9059cbb338360006040516020015260405160e060020a63ffffffff8516028152600160a060020a0390921660048301526024820152604401602060405180830381600087803b15156104be57600080fd5b6102c65a03f115156104cf57600080fd5b50505060405180515050600160a060020a0333167f28cab0d660ed8aedd61a8c9db00b97f6a2d67e07d87795994f440b18bc5f1aa3828560405191825260208201526040908101905180910390a2600154600160a060020a031683156108fc0284604051600060405180830381858888f19350505050151561055057600080fd5b600082111561058a57600160a060020a03331682156108fc0283604051600060405180830381858888f19350505050151561058a57600080fd5b505050565b60045481565b600a54600160a060020a031681565b60005433600160a060020a039081169116146105bf57600080fd5b600b5460a060020a900460ff16156105d657600080fd5b600b805474ff0000000000000000000000000000000000000000191660a060020a179055600354600954600160a060020a039182169163a9059cbb916101009004166adc94ce82ac7c640280000060006040516020015260405160e060020a63ffffffff8516028152600160a060020a0390921660048301526024820152604401602060405180830381600087803b151561067057600080fd5b6102c65a03f1151561068157600080fd5b50505060405180515050600354600a54600160a060020a039182169163a9059cbb91166a075a4b267d3758aac0000060006040516020015260405160e060020a63ffffffff8516028152600160a060020a0390921660048301526024820152604401602060405180830381600087803b15156106fc57600080fd5b6102c65a03f1151561070d57600080fd5b50505060405180515050600354600b54600160a060020a039182169163a9059cbb91166a1d692c99f4dd62ab00000060006040516020015260405160e060020a63ffffffff8516028152600160a060020a0390921660048301526024820152604401602060405180830381600087803b151561078857600080fd5b6102c65a03f1151561079957600080fd5b50505060405180515050565b6b02df455b08e99ea2b300000081565b60065481565b6000805433600160a060020a039081169116146107d757600080fd5b600354600160a060020a0316635c975abb6000604051602001526040518163ffffffff1660e060020a028152600401602060405180830381600087803b151561081f57600080fd5b6102c65a03f1151561083057600080fd5b505050604051805115905061084457600080fd5b600354600160a060020a0316638456cb596040518163ffffffff1660e060020a028152600401600060405180830381600087803b151561088357600080fd5b6102c65a03f1151561089457600080fd5b5050506001905090565b60075481565b60005433600160a060020a039081169116146108bf57600080fd5b60005460a060020a900460ff1615156108d757600080fd5b6000805474ff0000000000000000000000000000000000000000191690557f7805862f689e2f13df9f062ff482ad3ad112aca9e0847911ed832e158c525b3360405160405180910390a1565b6adc94ce82ac7c640280000081565b6000805433600160a060020a0390811691161461094e57600080fd5b600354600160a060020a0316635c975abb6000604051602001526040518163ffffffff1660e060020a028152600401602060405180830381600087803b151561099657600080fd5b6102c65a03f115156109a757600080fd5b5050506040518051905015156109bc57600080fd5b600354600160a060020a0316633f4ba83a6040518163ffffffff1660e060020a028152600401600060405180830381600087803b151561088357600080fd5b60005460a060020a900460ff1681565b60005433600160a060020a03908116911614610a2657600080fd5b60095460ff1615610a3657600080fd5b60045460055410610a4657600080fd5b600854421115610a5557600080fd5b60008111610a6257600080fd5b600755565b600254600160a060020a031681565b60095460ff1681565b60005433600160a060020a03908116911614610a9a57600080fd5b60005460a060020a900460ff1615610ab157600080fd5b6000805474ff0000000000000000000000000000000000000000191660a060020a1790557f6985a02210a168e66602d3235cb6db0e70f92b3ba4d376a33c0f3d9434bff62560405160405180910390a1565b600054600160a060020a031681565b6a075a4b267d3758aac0000081565b600b54600160a060020a031681565b60005433600160a060020a03908116911614610b4b57600080fd5b60095460ff1615610b5b57600080fd5b60045460055410610b6b57600080fd5b600854421115610b7a57600080fd5b60008111610b8757600080fd5b610b9f81670de0b6b3a764000063ffffffff610d9d16565b60065550565b60005433600160a060020a03908116911614610bc057600080fd5b600454600554101580610bd4575060085442115b1515610bdf57600080fd5b60095460ff1615610bef57600080fd5b6009805460ff19166001179055600354600254600160a060020a039182169163a9059cbb9116826370a082313060006040516020015260405160e060020a63ffffffff8416028152600160a060020a039091166004820152602401602060405180830381600087803b1515610c6357600080fd5b6102c65a03f11515610c7457600080fd5b5050506040518051905060006040516020015260405160e060020a63ffffffff8516028152600160a060020a0390921660048301526024820152604401602060405180830381600087803b151561078857600080fd5b60055481565b6009546101009004600160a060020a031681565b60005433600160a060020a03908116911614610cff57600080fd5b600160a060020a0381161515610d1457600080fd5b600054600160a060020a0380831691167f8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e060405160405180910390a36000805473ffffffffffffffffffffffffffffffffffffffff1916600160a060020a0392909216919091179055565b6a1d692c99f4dd62ab00000081565b600354600160a060020a031681565b6000828202831580610db95750828482811515610db657fe5b04145b1515610dc157fe5b9392505050565b600082821115610dd457fe5b50900390565b6000808284811515610de857fe5b04949350505050565b600082820183811015610dc157fe5b600081610e0b610e2d565b908152602001604051809103906000f0801515610e2757600080fd5b92915050565b604051610b9780610e3e83390190560060606040526003805460a060020a60ff0219169055341561001f57600080fd5b604051602080610b978339810160405280805160038054600160a060020a03191633600160a060020a03161790559150506000811161005d57600080fd5b600160a060020a033316600090815260016020526040812082905555610b0f806100886000396000f300606060405236156100e35763ffffffff7c010000000000000000000000000000000000000000000000000000000060003504166306fdde0381146100e8578063095ea7b31461017257806318160ddd146101a857806323b872dd146101cd578063313ce567146101f55780633f4ba83a146102215780635c975abb14610236578063661884631461024957806370a082311461026b5780638456cb591461028a5780638da5cb5b1461029d57806395d89b41146102cc578063a9059cbb146102df578063d73dd62314610301578063dd62ed3e14610323578063f2fde38b14610348575b600080fd5b34156100f357600080fd5b6100fb610367565b60405160208082528190810183818151815260200191508051906020019080838360005b8381101561013757808201518382015260200161011f565b50505050905090810190601f1680156101645780820380516001836020036101000a031916815260200191505b509250505060405180910390f35b341561017d57600080fd5b610194600160a060020a036004351660243561039e565b604051901515815260200160405180910390f35b34156101b357600080fd5b6101bb61040a565b60405190815260200160405180910390f35b34156101d857600080fd5b610194600160a060020a0360043581169060243516604435610410565b341561020057600080fd5b61020861043d565b60405163ffffffff909116815260200160405180910390f35b341561022c57600080fd5b610234610442565b005b341561024157600080fd5b6101946104c1565b341561025457600080fd5b610194600160a060020a03600435166024356104d1565b341561027657600080fd5b6101bb600160a060020a03600435166105cb565b341561029557600080fd5b6102346105e6565b34156102a857600080fd5b6102b061066a565b604051600160a060020a03909116815260200160405180910390f35b34156102d757600080fd5b6100fb610679565b34156102ea57600080fd5b610194600160a060020a03600435166024356106b0565b341561030c57600080fd5b610194600160a060020a03600435166024356106db565b341561032e57600080fd5b6101bb600160a060020a036004358116906024351661077f565b341561035357600080fd5b610234600160a060020a03600435166107aa565b60408051908101604052600a81527f484f515520546f6b656e00000000000000000000000000000000000000000000602082015281565b600160a060020a03338116600081815260026020908152604080832094871680845294909152808220859055909291907f8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b9259085905190815260200160405180910390a350600192915050565b60005481565b60035460009060a060020a900460ff161561042a57600080fd5b610435848484610845565b949350505050565b601281565b60035433600160a060020a0390811691161461045d57600080fd5b60035460a060020a900460ff16151561047557600080fd5b6003805474ff0000000000000000000000000000000000000000191690557f7805862f689e2f13df9f062ff482ad3ad112aca9e0847911ed832e158c525b3360405160405180910390a1565b60035460a060020a900460ff1681565b600160a060020a0333811660009081526002602090815260408083209386168352929052908120548083111561052e57600160a060020a033381166000908152600260209081526040808320938816835292905290812055610565565b61053e818463ffffffff6109c716565b600160a060020a033381166000908152600260209081526040808320938916835292905220555b600160a060020a0333811660008181526002602090815260408083209489168084529490915290819020547f8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925915190815260200160405180910390a35060019392505050565b600160a060020a031660009081526001602052604090205490565b60035433600160a060020a0390811691161461060157600080fd5b60035460a060020a900460ff161561061857600080fd5b6003805474ff0000000000000000000000000000000000000000191660a060020a1790557f6985a02210a168e66602d3235cb6db0e70f92b3ba4d376a33c0f3d9434bff62560405160405180910390a1565b600354600160a060020a031681565b60408051908101604052600381527f4851580000000000000000000000000000000000000000000000000000000000602082015281565b60035460009060a060020a900460ff16156106ca57600080fd5b6106d483836109d9565b9392505050565b600160a060020a033381166000908152600260209081526040808320938616835292905290812054610713908363ffffffff610ad416565b600160a060020a0333811660008181526002602090815260408083209489168084529490915290819020849055919290917f8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b92591905190815260200160405180910390a350600192915050565b600160a060020a03918216600090815260026020908152604080832093909416825291909152205490565b60035433600160a060020a039081169116146107c557600080fd5b600160a060020a03811615156107da57600080fd5b600354600160a060020a0380831691167f8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e060405160405180910390a36003805473ffffffffffffffffffffffffffffffffffffffff1916600160a060020a0392909216919091179055565b6000600160a060020a038316151561085c57600080fd5b600160a060020a03841660009081526001602052604090205482111561088157600080fd5b600160a060020a03808516600090815260026020908152604080832033909416835292905220548211156108b457600080fd5b600160a060020a0384166000908152600160205260409020546108dd908363ffffffff6109c716565b600160a060020a038086166000908152600160205260408082209390935590851681522054610912908363ffffffff610ad416565b600160a060020a0380851660009081526001602090815260408083209490945587831682526002815283822033909316825291909152205461095a908363ffffffff6109c716565b600160a060020a03808616600081815260026020908152604080832033861684529091529081902093909355908516917fddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef9085905190815260200160405180910390a35060019392505050565b6000828211156109d357fe5b50900390565b6000600160a060020a03831615156109f057600080fd5b600160a060020a033316600090815260016020526040902054821115610a1557600080fd5b600160a060020a033316600090815260016020526040902054610a3e908363ffffffff6109c716565b600160a060020a033381166000908152600160205260408082209390935590851681522054610a73908363ffffffff610ad416565b600160a060020a0380851660008181526001602052604090819020939093559133909116907fddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef9085905190815260200160405180910390a350600192915050565b6000828201838110156106d457fe00a165627a7a723058209b6a03e89247e92eae0caf598688102b6b8174c57df8a77361c359233f8e072f0029a165627a7a723058207d4b5bc3af1d3de3ded3b1875220cb21650ccf29124e29b152a823140181dd02002960606040526003805460a060020a60ff0219169055341561001f57600080fd5b604051602080610b978339810160405280805160038054600160a060020a03191633600160a060020a03161790559150506000811161005d57600080fd5b600160a060020a033316600090815260016020526040812082905555610b0f806100886000396000f300606060405236156100e35763ffffffff7c010000000000000000000000000000000000000000000000000000000060003504166306fdde0381146100e8578063095ea7b31461017257806318160ddd146101a857806323b872dd146101cd578063313ce567146101f55780633f4ba83a146102215780635c975abb14610236578063661884631461024957806370a082311461026b5780638456cb591461028a5780638da5cb5b1461029d57806395d89b41146102cc578063a9059cbb146102df578063d73dd62314610301578063dd62ed3e14610323578063f2fde38b14610348575b600080fd5b34156100f357600080fd5b6100fb610367565b60405160208082528190810183818151815260200191508051906020019080838360005b8381101561013757808201518382015260200161011f565b50505050905090810190601f1680156101645780820380516001836020036101000a031916815260200191505b509250505060405180910390f35b341561017d57600080fd5b610194600160a060020a036004351660243561039e565b604051901515815260200160405180910390f35b34156101b357600080fd5b6101bb61040a565b60405190815260200160405180910390f35b34156101d857600080fd5b610194600160a060020a0360043581169060243516604435610410565b341561020057600080fd5b61020861043d565b60405163ffffffff909116815260200160405180910390f35b341561022c57600080fd5b610234610442565b005b341561024157600080fd5b6101946104c1565b341561025457600080fd5b610194600160a060020a03600435166024356104d1565b341561027657600080fd5b6101bb600160a060020a03600435166105cb565b341561029557600080fd5b6102346105e6565b34156102a857600080fd5b6102b061066a565b604051600160a060020a03909116815260200160405180910390f35b34156102d757600080fd5b6100fb610679565b34156102ea57600080fd5b610194600160a060020a03600435166024356106b0565b341561030c57600080fd5b610194600160a060020a03600435166024356106db565b341561032e57600080fd5b6101bb600160a060020a036004358116906024351661077f565b341561035357600080fd5b610234600160a060020a03600435166107aa565b60408051908101604052600a81527f484f515520546f6b656e00000000000000000000000000000000000000000000602082015281565b600160a060020a03338116600081815260026020908152604080832094871680845294909152808220859055909291907f8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b9259085905190815260200160405180910390a350600192915050565b60005481565b60035460009060a060020a900460ff161561042a57600080fd5b610435848484610845565b949350505050565b601281565b60035433600160a060020a0390811691161461045d57600080fd5b60035460a060020a900460ff16151561047557600080fd5b6003805474ff0000000000000000000000000000000000000000191690557f7805862f689e2f13df9f062ff482ad3ad112aca9e0847911ed832e158c525b3360405160405180910390a1565b60035460a060020a900460ff1681565b600160a060020a0333811660009081526002602090815260408083209386168352929052908120548083111561052e57600160a060020a033381166000908152600260209081526040808320938816835292905290812055610565565b61053e818463ffffffff6109c716565b600160a060020a033381166000908152600260209081526040808320938916835292905220555b600160a060020a0333811660008181526002602090815260408083209489168084529490915290819020547f8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925915190815260200160405180910390a35060019392505050565b600160a060020a031660009081526001602052604090205490565b60035433600160a060020a0390811691161461060157600080fd5b60035460a060020a900460ff161561061857600080fd5b6003805474ff0000000000000000000000000000000000000000191660a060020a1790557f6985a02210a168e66602d3235cb6db0e70f92b3ba4d376a33c0f3d9434bff62560405160405180910390a1565b600354600160a060020a031681565b60408051908101604052600381527f4851580000000000000000000000000000000000000000000000000000000000602082015281565b60035460009060a060020a900460ff16156106ca57600080fd5b6106d483836109d9565b9392505050565b600160a060020a033381166000908152600260209081526040808320938616835292905290812054610713908363ffffffff610ad416565b600160a060020a0333811660008181526002602090815260408083209489168084529490915290819020849055919290917f8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b92591905190815260200160405180910390a350600192915050565b600160a060020a03918216600090815260026020908152604080832093909416825291909152205490565b60035433600160a060020a039081169116146107c557600080fd5b600160a060020a03811615156107da57600080fd5b600354600160a060020a0380831691167f8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e060405160405180910390a36003805473ffffffffffffffffffffffffffffffffffffffff1916600160a060020a0392909216919091179055565b6000600160a060020a038316151561085c57600080fd5b600160a060020a03841660009081526001602052604090205482111561088157600080fd5b600160a060020a03808516600090815260026020908152604080832033909416835292905220548211156108b457600080fd5b600160a060020a0384166000908152600160205260409020546108dd908363ffffffff6109c716565b600160a060020a038086166000908152600160205260408082209390935590851681522054610912908363ffffffff610ad416565b600160a060020a0380851660009081526001602090815260408083209490945587831682526002815283822033909316825291909152205461095a908363ffffffff6109c716565b600160a060020a03808616600081815260026020908152604080832033861684529091529081902093909355908516917fddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef9085905190815260200160405180910390a35060019392505050565b6000828211156109d357fe5b50900390565b6000600160a060020a03831615156109f057600080fd5b600160a060020a033316600090815260016020526040902054821115610a1557600080fd5b600160a060020a033316600090815260016020526040902054610a3e908363ffffffff6109c716565b600160a060020a033381166000908152600160205260408082209390935590851681522054610a73908363ffffffff610ad416565b600160a060020a0380851660008181526001602052604090819020939093559133909116907fddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef9085905190815260200160405180910390a350600192915050565b6000828201838110156106d457fe00a165627a7a723058209b6a03e89247e92eae0caf598688102b6b8174c57df8a77361c359233f8e072f0029`

// DeployPrivatePlacement deploys a new Ethereum contract, binding an instance of PrivatePlacement to it.
func DeployPrivatePlacement(auth *bind.TransactOpts, backend bind.ContractBackend, _bankAddress common.Address, _foundersAddress common.Address, _supportAddress common.Address, _bountyAddress common.Address, _beneficiaryAddress common.Address) (common.Address, *types.Transaction, *PrivatePlacement, error) {
	parsed, err := abi.JSON(strings.NewReader(PrivatePlacementABI))
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	address, tx, contract, err := bind.DeployContract(auth, parsed, common.FromHex(PrivatePlacementBin), backend, _bankAddress, _foundersAddress, _supportAddress, _bountyAddress, _beneficiaryAddress)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &PrivatePlacement{PrivatePlacementCaller: PrivatePlacementCaller{contract: contract}, PrivatePlacementTransactor: PrivatePlacementTransactor{contract: contract}}, nil
}

// PrivatePlacement is an auto generated Go binding around an Ethereum contract.
type PrivatePlacement struct {
	PrivatePlacementCaller     // Read-only binding to the contract
	PrivatePlacementTransactor // Write-only binding to the contract
}

// PrivatePlacementCaller is an auto generated read-only Go binding around an Ethereum contract.
type PrivatePlacementCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// PrivatePlacementTransactor is an auto generated write-only Go binding around an Ethereum contract.
type PrivatePlacementTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// PrivatePlacementSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type PrivatePlacementSession struct {
	Contract     *PrivatePlacement // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// PrivatePlacementCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type PrivatePlacementCallerSession struct {
	Contract *PrivatePlacementCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts           // Call options to use throughout this session
}

// PrivatePlacementTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type PrivatePlacementTransactorSession struct {
	Contract     *PrivatePlacementTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts           // Transaction auth options to use throughout this session
}

// PrivatePlacementRaw is an auto generated low-level Go binding around an Ethereum contract.
type PrivatePlacementRaw struct {
	Contract *PrivatePlacement // Generic contract binding to access the raw methods on
}

// PrivatePlacementCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type PrivatePlacementCallerRaw struct {
	Contract *PrivatePlacementCaller // Generic read-only contract binding to access the raw methods on
}

// PrivatePlacementTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type PrivatePlacementTransactorRaw struct {
	Contract *PrivatePlacementTransactor // Generic write-only contract binding to access the raw methods on
}

// NewPrivatePlacement creates a new instance of PrivatePlacement, bound to a specific deployed contract.
func NewPrivatePlacement(address common.Address, backend bind.ContractBackend) (*PrivatePlacement, error) {
	contract, err := bindPrivatePlacement(address, backend, backend)
	if err != nil {
		return nil, err
	}
	return &PrivatePlacement{PrivatePlacementCaller: PrivatePlacementCaller{contract: contract}, PrivatePlacementTransactor: PrivatePlacementTransactor{contract: contract}}, nil
}

// NewPrivatePlacementCaller creates a new read-only instance of PrivatePlacement, bound to a specific deployed contract.
func NewPrivatePlacementCaller(address common.Address, caller bind.ContractCaller) (*PrivatePlacementCaller, error) {
	contract, err := bindPrivatePlacement(address, caller, nil)
	if err != nil {
		return nil, err
	}
	return &PrivatePlacementCaller{contract: contract}, nil
}

// NewPrivatePlacementTransactor creates a new write-only instance of PrivatePlacement, bound to a specific deployed contract.
func NewPrivatePlacementTransactor(address common.Address, transactor bind.ContractTransactor) (*PrivatePlacementTransactor, error) {
	contract, err := bindPrivatePlacement(address, nil, transactor)
	if err != nil {
		return nil, err
	}
	return &PrivatePlacementTransactor{contract: contract}, nil
}

// bindPrivatePlacement binds a generic wrapper to an already deployed contract.
func bindPrivatePlacement(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(PrivatePlacementABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_PrivatePlacement *PrivatePlacementRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _PrivatePlacement.Contract.PrivatePlacementCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_PrivatePlacement *PrivatePlacementRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _PrivatePlacement.Contract.PrivatePlacementTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_PrivatePlacement *PrivatePlacementRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _PrivatePlacement.Contract.PrivatePlacementTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_PrivatePlacement *PrivatePlacementCallerRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _PrivatePlacement.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_PrivatePlacement *PrivatePlacementTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _PrivatePlacement.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_PrivatePlacement *PrivatePlacementTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _PrivatePlacement.Contract.contract.Transact(opts, method, params...)
}

// BankAddress is a free data retrieval call binding the contract method 0x7822ed49.
//
// Solidity: function bankAddress() constant returns(address)
func (_PrivatePlacement *PrivatePlacementCaller) BankAddress(opts *bind.CallOpts) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _PrivatePlacement.contract.Call(opts, out, "bankAddress")
	return *ret0, err
}

// BankAddress is a free data retrieval call binding the contract method 0x7822ed49.
//
// Solidity: function bankAddress() constant returns(address)
func (_PrivatePlacement *PrivatePlacementSession) BankAddress() (common.Address, error) {
	return _PrivatePlacement.Contract.BankAddress(&_PrivatePlacement.CallOpts)
}

// BankAddress is a free data retrieval call binding the contract method 0x7822ed49.
//
// Solidity: function bankAddress() constant returns(address)
func (_PrivatePlacement *PrivatePlacementCallerSession) BankAddress() (common.Address, error) {
	return _PrivatePlacement.Contract.BankAddress(&_PrivatePlacement.CallOpts)
}

// BountyAddress is a free data retrieval call binding the contract method 0xc516358f.
//
// Solidity: function bountyAddress() constant returns(address)
func (_PrivatePlacement *PrivatePlacementCaller) BountyAddress(opts *bind.CallOpts) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _PrivatePlacement.contract.Call(opts, out, "bountyAddress")
	return *ret0, err
}

// BountyAddress is a free data retrieval call binding the contract method 0xc516358f.
//
// Solidity: function bountyAddress() constant returns(address)
func (_PrivatePlacement *PrivatePlacementSession) BountyAddress() (common.Address, error) {
	return _PrivatePlacement.Contract.BountyAddress(&_PrivatePlacement.CallOpts)
}

// BountyAddress is a free data retrieval call binding the contract method 0xc516358f.
//
// Solidity: function bountyAddress() constant returns(address)
func (_PrivatePlacement *PrivatePlacementCallerSession) BountyAddress() (common.Address, error) {
	return _PrivatePlacement.Contract.BountyAddress(&_PrivatePlacement.CallOpts)
}

// FoundersAddress is a free data retrieval call binding the contract method 0xebfdc657.
//
// Solidity: function foundersAddress() constant returns(address)
func (_PrivatePlacement *PrivatePlacementCaller) FoundersAddress(opts *bind.CallOpts) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _PrivatePlacement.contract.Call(opts, out, "foundersAddress")
	return *ret0, err
}

// FoundersAddress is a free data retrieval call binding the contract method 0xebfdc657.
//
// Solidity: function foundersAddress() constant returns(address)
func (_PrivatePlacement *PrivatePlacementSession) FoundersAddress() (common.Address, error) {
	return _PrivatePlacement.Contract.FoundersAddress(&_PrivatePlacement.CallOpts)
}

// FoundersAddress is a free data retrieval call binding the contract method 0xebfdc657.
//
// Solidity: function foundersAddress() constant returns(address)
func (_PrivatePlacement *PrivatePlacementCallerSession) FoundersAddress() (common.Address, error) {
	return _PrivatePlacement.Contract.FoundersAddress(&_PrivatePlacement.CallOpts)
}

// InitialBountyAmount is a free data retrieval call binding the contract method 0xf338c984.
//
// Solidity: function initialBountyAmount() constant returns(uint256)
func (_PrivatePlacement *PrivatePlacementCaller) InitialBountyAmount(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _PrivatePlacement.contract.Call(opts, out, "initialBountyAmount")
	return *ret0, err
}

// InitialBountyAmount is a free data retrieval call binding the contract method 0xf338c984.
//
// Solidity: function initialBountyAmount() constant returns(uint256)
func (_PrivatePlacement *PrivatePlacementSession) InitialBountyAmount() (*big.Int, error) {
	return _PrivatePlacement.Contract.InitialBountyAmount(&_PrivatePlacement.CallOpts)
}

// InitialBountyAmount is a free data retrieval call binding the contract method 0xf338c984.
//
// Solidity: function initialBountyAmount() constant returns(uint256)
func (_PrivatePlacement *PrivatePlacementCallerSession) InitialBountyAmount() (*big.Int, error) {
	return _PrivatePlacement.Contract.InitialBountyAmount(&_PrivatePlacement.CallOpts)
}

// InitialFoundersAmount is a free data retrieval call binding the contract method 0x4d39ed06.
//
// Solidity: function initialFoundersAmount() constant returns(uint256)
func (_PrivatePlacement *PrivatePlacementCaller) InitialFoundersAmount(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _PrivatePlacement.contract.Call(opts, out, "initialFoundersAmount")
	return *ret0, err
}

// InitialFoundersAmount is a free data retrieval call binding the contract method 0x4d39ed06.
//
// Solidity: function initialFoundersAmount() constant returns(uint256)
func (_PrivatePlacement *PrivatePlacementSession) InitialFoundersAmount() (*big.Int, error) {
	return _PrivatePlacement.Contract.InitialFoundersAmount(&_PrivatePlacement.CallOpts)
}

// InitialFoundersAmount is a free data retrieval call binding the contract method 0x4d39ed06.
//
// Solidity: function initialFoundersAmount() constant returns(uint256)
func (_PrivatePlacement *PrivatePlacementCallerSession) InitialFoundersAmount() (*big.Int, error) {
	return _PrivatePlacement.Contract.InitialFoundersAmount(&_PrivatePlacement.CallOpts)
}

// InitialSupportAmount is a free data retrieval call binding the contract method 0xab28c704.
//
// Solidity: function initialSupportAmount() constant returns(uint256)
func (_PrivatePlacement *PrivatePlacementCaller) InitialSupportAmount(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _PrivatePlacement.contract.Call(opts, out, "initialSupportAmount")
	return *ret0, err
}

// InitialSupportAmount is a free data retrieval call binding the contract method 0xab28c704.
//
// Solidity: function initialSupportAmount() constant returns(uint256)
func (_PrivatePlacement *PrivatePlacementSession) InitialSupportAmount() (*big.Int, error) {
	return _PrivatePlacement.Contract.InitialSupportAmount(&_PrivatePlacement.CallOpts)
}

// InitialSupportAmount is a free data retrieval call binding the contract method 0xab28c704.
//
// Solidity: function initialSupportAmount() constant returns(uint256)
func (_PrivatePlacement *PrivatePlacementCallerSession) InitialSupportAmount() (*big.Int, error) {
	return _PrivatePlacement.Contract.InitialSupportAmount(&_PrivatePlacement.CallOpts)
}

// IsFinished is a free data retrieval call binding the contract method 0x7b352962.
//
// Solidity: function isFinished() constant returns(bool)
func (_PrivatePlacement *PrivatePlacementCaller) IsFinished(opts *bind.CallOpts) (bool, error) {
	var (
		ret0 = new(bool)
	)
	out := ret0
	err := _PrivatePlacement.contract.Call(opts, out, "isFinished")
	return *ret0, err
}

// IsFinished is a free data retrieval call binding the contract method 0x7b352962.
//
// Solidity: function isFinished() constant returns(bool)
func (_PrivatePlacement *PrivatePlacementSession) IsFinished() (bool, error) {
	return _PrivatePlacement.Contract.IsFinished(&_PrivatePlacement.CallOpts)
}

// IsFinished is a free data retrieval call binding the contract method 0x7b352962.
//
// Solidity: function isFinished() constant returns(bool)
func (_PrivatePlacement *PrivatePlacementCallerSession) IsFinished() (bool, error) {
	return _PrivatePlacement.Contract.IsFinished(&_PrivatePlacement.CallOpts)
}

// IssuedTokensAmount is a free data retrieval call binding the contract method 0xe69b414b.
//
// Solidity: function issuedTokensAmount() constant returns(uint256)
func (_PrivatePlacement *PrivatePlacementCaller) IssuedTokensAmount(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _PrivatePlacement.contract.Call(opts, out, "issuedTokensAmount")
	return *ret0, err
}

// IssuedTokensAmount is a free data retrieval call binding the contract method 0xe69b414b.
//
// Solidity: function issuedTokensAmount() constant returns(uint256)
func (_PrivatePlacement *PrivatePlacementSession) IssuedTokensAmount() (*big.Int, error) {
	return _PrivatePlacement.Contract.IssuedTokensAmount(&_PrivatePlacement.CallOpts)
}

// IssuedTokensAmount is a free data retrieval call binding the contract method 0xe69b414b.
//
// Solidity: function issuedTokensAmount() constant returns(uint256)
func (_PrivatePlacement *PrivatePlacementCallerSession) IssuedTokensAmount() (*big.Int, error) {
	return _PrivatePlacement.Contract.IssuedTokensAmount(&_PrivatePlacement.CallOpts)
}

// MaxTokensAmount is a free data retrieval call binding the contract method 0x0a384665.
//
// Solidity: function maxTokensAmount() constant returns(uint256)
func (_PrivatePlacement *PrivatePlacementCaller) MaxTokensAmount(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _PrivatePlacement.contract.Call(opts, out, "maxTokensAmount")
	return *ret0, err
}

// MaxTokensAmount is a free data retrieval call binding the contract method 0x0a384665.
//
// Solidity: function maxTokensAmount() constant returns(uint256)
func (_PrivatePlacement *PrivatePlacementSession) MaxTokensAmount() (*big.Int, error) {
	return _PrivatePlacement.Contract.MaxTokensAmount(&_PrivatePlacement.CallOpts)
}

// MaxTokensAmount is a free data retrieval call binding the contract method 0x0a384665.
//
// Solidity: function maxTokensAmount() constant returns(uint256)
func (_PrivatePlacement *PrivatePlacementCallerSession) MaxTokensAmount() (*big.Int, error) {
	return _PrivatePlacement.Contract.MaxTokensAmount(&_PrivatePlacement.CallOpts)
}

// MinBuyableAmount is a free data retrieval call binding the contract method 0x2511b182.
//
// Solidity: function minBuyableAmount() constant returns(uint256)
func (_PrivatePlacement *PrivatePlacementCaller) MinBuyableAmount(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _PrivatePlacement.contract.Call(opts, out, "minBuyableAmount")
	return *ret0, err
}

// MinBuyableAmount is a free data retrieval call binding the contract method 0x2511b182.
//
// Solidity: function minBuyableAmount() constant returns(uint256)
func (_PrivatePlacement *PrivatePlacementSession) MinBuyableAmount() (*big.Int, error) {
	return _PrivatePlacement.Contract.MinBuyableAmount(&_PrivatePlacement.CallOpts)
}

// MinBuyableAmount is a free data retrieval call binding the contract method 0x2511b182.
//
// Solidity: function minBuyableAmount() constant returns(uint256)
func (_PrivatePlacement *PrivatePlacementCallerSession) MinBuyableAmount() (*big.Int, error) {
	return _PrivatePlacement.Contract.MinBuyableAmount(&_PrivatePlacement.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() constant returns(address)
func (_PrivatePlacement *PrivatePlacementCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _PrivatePlacement.contract.Call(opts, out, "owner")
	return *ret0, err
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() constant returns(address)
func (_PrivatePlacement *PrivatePlacementSession) Owner() (common.Address, error) {
	return _PrivatePlacement.Contract.Owner(&_PrivatePlacement.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() constant returns(address)
func (_PrivatePlacement *PrivatePlacementCallerSession) Owner() (common.Address, error) {
	return _PrivatePlacement.Contract.Owner(&_PrivatePlacement.CallOpts)
}

// Paused is a free data retrieval call binding the contract method 0x5c975abb.
//
// Solidity: function paused() constant returns(bool)
func (_PrivatePlacement *PrivatePlacementCaller) Paused(opts *bind.CallOpts) (bool, error) {
	var (
		ret0 = new(bool)
	)
	out := ret0
	err := _PrivatePlacement.contract.Call(opts, out, "paused")
	return *ret0, err
}

// Paused is a free data retrieval call binding the contract method 0x5c975abb.
//
// Solidity: function paused() constant returns(bool)
func (_PrivatePlacement *PrivatePlacementSession) Paused() (bool, error) {
	return _PrivatePlacement.Contract.Paused(&_PrivatePlacement.CallOpts)
}

// Paused is a free data retrieval call binding the contract method 0x5c975abb.
//
// Solidity: function paused() constant returns(bool)
func (_PrivatePlacement *PrivatePlacementCallerSession) Paused() (bool, error) {
	return _PrivatePlacement.Contract.Paused(&_PrivatePlacement.CallOpts)
}

// SupportAddress is a free data retrieval call binding the contract method 0x0f6d37d6.
//
// Solidity: function supportAddress() constant returns(address)
func (_PrivatePlacement *PrivatePlacementCaller) SupportAddress(opts *bind.CallOpts) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _PrivatePlacement.contract.Call(opts, out, "supportAddress")
	return *ret0, err
}

// SupportAddress is a free data retrieval call binding the contract method 0x0f6d37d6.
//
// Solidity: function supportAddress() constant returns(address)
func (_PrivatePlacement *PrivatePlacementSession) SupportAddress() (common.Address, error) {
	return _PrivatePlacement.Contract.SupportAddress(&_PrivatePlacement.CallOpts)
}

// SupportAddress is a free data retrieval call binding the contract method 0x0f6d37d6.
//
// Solidity: function supportAddress() constant returns(address)
func (_PrivatePlacement *PrivatePlacementCallerSession) SupportAddress() (common.Address, error) {
	return _PrivatePlacement.Contract.SupportAddress(&_PrivatePlacement.CallOpts)
}

// Token is a free data retrieval call binding the contract method 0xfc0c546a.
//
// Solidity: function token() constant returns(address)
func (_PrivatePlacement *PrivatePlacementCaller) Token(opts *bind.CallOpts) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _PrivatePlacement.contract.Call(opts, out, "token")
	return *ret0, err
}

// Token is a free data retrieval call binding the contract method 0xfc0c546a.
//
// Solidity: function token() constant returns(address)
func (_PrivatePlacement *PrivatePlacementSession) Token() (common.Address, error) {
	return _PrivatePlacement.Contract.Token(&_PrivatePlacement.CallOpts)
}

// Token is a free data retrieval call binding the contract method 0xfc0c546a.
//
// Solidity: function token() constant returns(address)
func (_PrivatePlacement *PrivatePlacementCallerSession) Token() (common.Address, error) {
	return _PrivatePlacement.Contract.Token(&_PrivatePlacement.CallOpts)
}

// TokenRate is a free data retrieval call binding the contract method 0x31711884.
//
// Solidity: function tokenRate() constant returns(uint256)
func (_PrivatePlacement *PrivatePlacementCaller) TokenRate(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _PrivatePlacement.contract.Call(opts, out, "tokenRate")
	return *ret0, err
}

// TokenRate is a free data retrieval call binding the contract method 0x31711884.
//
// Solidity: function tokenRate() constant returns(uint256)
func (_PrivatePlacement *PrivatePlacementSession) TokenRate() (*big.Int, error) {
	return _PrivatePlacement.Contract.TokenRate(&_PrivatePlacement.CallOpts)
}

// TokenRate is a free data retrieval call binding the contract method 0x31711884.
//
// Solidity: function tokenRate() constant returns(uint256)
func (_PrivatePlacement *PrivatePlacementCallerSession) TokenRate() (*big.Int, error) {
	return _PrivatePlacement.Contract.TokenRate(&_PrivatePlacement.CallOpts)
}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() constant returns(uint256)
func (_PrivatePlacement *PrivatePlacementCaller) TotalSupply(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _PrivatePlacement.contract.Call(opts, out, "totalSupply")
	return *ret0, err
}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() constant returns(uint256)
func (_PrivatePlacement *PrivatePlacementSession) TotalSupply() (*big.Int, error) {
	return _PrivatePlacement.Contract.TotalSupply(&_PrivatePlacement.CallOpts)
}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() constant returns(uint256)
func (_PrivatePlacement *PrivatePlacementCallerSession) TotalSupply() (*big.Int, error) {
	return _PrivatePlacement.Contract.TotalSupply(&_PrivatePlacement.CallOpts)
}

// AllocateInternalWallets is a paid mutator transaction binding the contract method 0x1135b3ac.
//
// Solidity: function allocateInternalWallets() returns()
func (_PrivatePlacement *PrivatePlacementTransactor) AllocateInternalWallets(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _PrivatePlacement.contract.Transact(opts, "allocateInternalWallets")
}

// AllocateInternalWallets is a paid mutator transaction binding the contract method 0x1135b3ac.
//
// Solidity: function allocateInternalWallets() returns()
func (_PrivatePlacement *PrivatePlacementSession) AllocateInternalWallets() (*types.Transaction, error) {
	return _PrivatePlacement.Contract.AllocateInternalWallets(&_PrivatePlacement.TransactOpts)
}

// AllocateInternalWallets is a paid mutator transaction binding the contract method 0x1135b3ac.
//
// Solidity: function allocateInternalWallets() returns()
func (_PrivatePlacement *PrivatePlacementTransactorSession) AllocateInternalWallets() (*types.Transaction, error) {
	return _PrivatePlacement.Contract.AllocateInternalWallets(&_PrivatePlacement.TransactOpts)
}

// BuyTokens is a paid mutator transaction binding the contract method 0xd0febe4c.
//
// Solidity: function buyTokens() returns()
func (_PrivatePlacement *PrivatePlacementTransactor) BuyTokens(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _PrivatePlacement.contract.Transact(opts, "buyTokens")
}

// BuyTokens is a paid mutator transaction binding the contract method 0xd0febe4c.
//
// Solidity: function buyTokens() returns()
func (_PrivatePlacement *PrivatePlacementSession) BuyTokens() (*types.Transaction, error) {
	return _PrivatePlacement.Contract.BuyTokens(&_PrivatePlacement.TransactOpts)
}

// BuyTokens is a paid mutator transaction binding the contract method 0xd0febe4c.
//
// Solidity: function buyTokens() returns()
func (_PrivatePlacement *PrivatePlacementTransactorSession) BuyTokens() (*types.Transaction, error) {
	return _PrivatePlacement.Contract.BuyTokens(&_PrivatePlacement.TransactOpts)
}

// Finish is a paid mutator transaction binding the contract method 0xd56b2889.
//
// Solidity: function finish() returns()
func (_PrivatePlacement *PrivatePlacementTransactor) Finish(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _PrivatePlacement.contract.Transact(opts, "finish")
}

// Finish is a paid mutator transaction binding the contract method 0xd56b2889.
//
// Solidity: function finish() returns()
func (_PrivatePlacement *PrivatePlacementSession) Finish() (*types.Transaction, error) {
	return _PrivatePlacement.Contract.Finish(&_PrivatePlacement.TransactOpts)
}

// Finish is a paid mutator transaction binding the contract method 0xd56b2889.
//
// Solidity: function finish() returns()
func (_PrivatePlacement *PrivatePlacementTransactorSession) Finish() (*types.Transaction, error) {
	return _PrivatePlacement.Contract.Finish(&_PrivatePlacement.TransactOpts)
}

// Pause is a paid mutator transaction binding the contract method 0x8456cb59.
//
// Solidity: function pause() returns()
func (_PrivatePlacement *PrivatePlacementTransactor) Pause(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _PrivatePlacement.contract.Transact(opts, "pause")
}

// Pause is a paid mutator transaction binding the contract method 0x8456cb59.
//
// Solidity: function pause() returns()
func (_PrivatePlacement *PrivatePlacementSession) Pause() (*types.Transaction, error) {
	return _PrivatePlacement.Contract.Pause(&_PrivatePlacement.TransactOpts)
}

// Pause is a paid mutator transaction binding the contract method 0x8456cb59.
//
// Solidity: function pause() returns()
func (_PrivatePlacement *PrivatePlacementTransactorSession) Pause() (*types.Transaction, error) {
	return _PrivatePlacement.Contract.Pause(&_PrivatePlacement.TransactOpts)
}

// PauseToken is a paid mutator transaction binding the contract method 0x2c349627.
//
// Solidity: function pauseToken() returns(bool)
func (_PrivatePlacement *PrivatePlacementTransactor) PauseToken(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _PrivatePlacement.contract.Transact(opts, "pauseToken")
}

// PauseToken is a paid mutator transaction binding the contract method 0x2c349627.
//
// Solidity: function pauseToken() returns(bool)
func (_PrivatePlacement *PrivatePlacementSession) PauseToken() (*types.Transaction, error) {
	return _PrivatePlacement.Contract.PauseToken(&_PrivatePlacement.TransactOpts)
}

// PauseToken is a paid mutator transaction binding the contract method 0x2c349627.
//
// Solidity: function pauseToken() returns(bool)
func (_PrivatePlacement *PrivatePlacementTransactorSession) PauseToken() (*types.Transaction, error) {
	return _PrivatePlacement.Contract.PauseToken(&_PrivatePlacement.TransactOpts)
}

// SetMinBuyableAmount is a paid mutator transaction binding the contract method 0xd561ef2b.
//
// Solidity: function setMinBuyableAmount(_minBuyableAmount uint256) returns()
func (_PrivatePlacement *PrivatePlacementTransactor) SetMinBuyableAmount(opts *bind.TransactOpts, _minBuyableAmount *big.Int) (*types.Transaction, error) {
	return _PrivatePlacement.contract.Transact(opts, "setMinBuyableAmount", _minBuyableAmount)
}

// SetMinBuyableAmount is a paid mutator transaction binding the contract method 0xd561ef2b.
//
// Solidity: function setMinBuyableAmount(_minBuyableAmount uint256) returns()
func (_PrivatePlacement *PrivatePlacementSession) SetMinBuyableAmount(_minBuyableAmount *big.Int) (*types.Transaction, error) {
	return _PrivatePlacement.Contract.SetMinBuyableAmount(&_PrivatePlacement.TransactOpts, _minBuyableAmount)
}

// SetMinBuyableAmount is a paid mutator transaction binding the contract method 0xd561ef2b.
//
// Solidity: function setMinBuyableAmount(_minBuyableAmount uint256) returns()
func (_PrivatePlacement *PrivatePlacementTransactorSession) SetMinBuyableAmount(_minBuyableAmount *big.Int) (*types.Transaction, error) {
	return _PrivatePlacement.Contract.SetMinBuyableAmount(&_PrivatePlacement.TransactOpts, _minBuyableAmount)
}

// SetTokenRate is a paid mutator transaction binding the contract method 0x61241c28.
//
// Solidity: function setTokenRate(_tokenRate uint256) returns()
func (_PrivatePlacement *PrivatePlacementTransactor) SetTokenRate(opts *bind.TransactOpts, _tokenRate *big.Int) (*types.Transaction, error) {
	return _PrivatePlacement.contract.Transact(opts, "setTokenRate", _tokenRate)
}

// SetTokenRate is a paid mutator transaction binding the contract method 0x61241c28.
//
// Solidity: function setTokenRate(_tokenRate uint256) returns()
func (_PrivatePlacement *PrivatePlacementSession) SetTokenRate(_tokenRate *big.Int) (*types.Transaction, error) {
	return _PrivatePlacement.Contract.SetTokenRate(&_PrivatePlacement.TransactOpts, _tokenRate)
}

// SetTokenRate is a paid mutator transaction binding the contract method 0x61241c28.
//
// Solidity: function setTokenRate(_tokenRate uint256) returns()
func (_PrivatePlacement *PrivatePlacementTransactorSession) SetTokenRate(_tokenRate *big.Int) (*types.Transaction, error) {
	return _PrivatePlacement.Contract.SetTokenRate(&_PrivatePlacement.TransactOpts, _tokenRate)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(newOwner address) returns()
func (_PrivatePlacement *PrivatePlacementTransactor) TransferOwnership(opts *bind.TransactOpts, newOwner common.Address) (*types.Transaction, error) {
	return _PrivatePlacement.contract.Transact(opts, "transferOwnership", newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(newOwner address) returns()
func (_PrivatePlacement *PrivatePlacementSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _PrivatePlacement.Contract.TransferOwnership(&_PrivatePlacement.TransactOpts, newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(newOwner address) returns()
func (_PrivatePlacement *PrivatePlacementTransactorSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _PrivatePlacement.Contract.TransferOwnership(&_PrivatePlacement.TransactOpts, newOwner)
}

// Unpause is a paid mutator transaction binding the contract method 0x3f4ba83a.
//
// Solidity: function unpause() returns()
func (_PrivatePlacement *PrivatePlacementTransactor) Unpause(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _PrivatePlacement.contract.Transact(opts, "unpause")
}

// Unpause is a paid mutator transaction binding the contract method 0x3f4ba83a.
//
// Solidity: function unpause() returns()
func (_PrivatePlacement *PrivatePlacementSession) Unpause() (*types.Transaction, error) {
	return _PrivatePlacement.Contract.Unpause(&_PrivatePlacement.TransactOpts)
}

// Unpause is a paid mutator transaction binding the contract method 0x3f4ba83a.
//
// Solidity: function unpause() returns()
func (_PrivatePlacement *PrivatePlacementTransactorSession) Unpause() (*types.Transaction, error) {
	return _PrivatePlacement.Contract.Unpause(&_PrivatePlacement.TransactOpts)
}

// UnpauseToken is a paid mutator transaction binding the contract method 0x50669a03.
//
// Solidity: function unpauseToken() returns(bool)
func (_PrivatePlacement *PrivatePlacementTransactor) UnpauseToken(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _PrivatePlacement.contract.Transact(opts, "unpauseToken")
}

// UnpauseToken is a paid mutator transaction binding the contract method 0x50669a03.
//
// Solidity: function unpauseToken() returns(bool)
func (_PrivatePlacement *PrivatePlacementSession) UnpauseToken() (*types.Transaction, error) {
	return _PrivatePlacement.Contract.UnpauseToken(&_PrivatePlacement.TransactOpts)
}

// UnpauseToken is a paid mutator transaction binding the contract method 0x50669a03.
//
// Solidity: function unpauseToken() returns(bool)
func (_PrivatePlacement *PrivatePlacementTransactorSession) UnpauseToken() (*types.Transaction, error) {
	return _PrivatePlacement.Contract.UnpauseToken(&_PrivatePlacement.TransactOpts)
}
