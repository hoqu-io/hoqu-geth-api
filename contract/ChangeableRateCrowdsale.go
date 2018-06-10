// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package contract

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

// ChangeableRateCrowdsaleABI is the input ABI used to generate the binding from.
const ChangeableRateCrowdsaleABI = "[{\"constant\":false,\"inputs\":[{\"name\":\"_receiver\",\"type\":\"address\"}],\"name\":\"claimOne\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"maxTokensAmount\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"minBuyableAmount\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_receiver\",\"type\":\"address\"},{\"name\":\"_equivalentEthAmount\",\"type\":\"uint256\"}],\"name\":\"sub\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"tokenRate\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"i\",\"type\":\"uint32\"}],\"name\":\"getReceiver\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[],\"name\":\"unpause\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"rateBoundaries\",\"outputs\":[{\"name\":\"amount\",\"type\":\"uint256\"},{\"name\":\"rate\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[],\"name\":\"claim\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"paused\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_tokenRate\",\"type\":\"uint256\"}],\"name\":\"setTokenRate\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"bankAddress\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"isFinished\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[],\"name\":\"pause\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"currentBoundary\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[],\"name\":\"buy\",\"outputs\":[],\"payable\":true,\"stateMutability\":\"payable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"getReceiversCount\",\"outputs\":[{\"name\":\"\",\"type\":\"uint32\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[],\"name\":\"claimAll\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"numOfBoundaries\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[],\"name\":\"finish\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_receiver\",\"type\":\"address\"},{\"name\":\"_equivalentEthAmount\",\"type\":\"uint256\"}],\"name\":\"topUp\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"\",\"type\":\"address\"}],\"name\":\"approved\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_receiver\",\"type\":\"address\"}],\"name\":\"approve\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"\",\"type\":\"address\"}],\"name\":\"tokens\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"issuedTokensAmount\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"nextBoundaryAmount\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"transferOwnership\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_receiver\",\"type\":\"address\"},{\"name\":\"_equivalentEthAmount\",\"type\":\"uint256\"}],\"name\":\"add\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"token\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"name\":\"_tokenAddress\",\"type\":\"address\"},{\"name\":\"_bankAddress\",\"type\":\"address\"},{\"name\":\"_beneficiaryAddress\",\"type\":\"address\"},{\"name\":\"_tokenRate\",\"type\":\"uint256\"},{\"name\":\"_minBuyableAmount\",\"type\":\"uint256\"},{\"name\":\"_maxTokensAmount\",\"type\":\"uint256\"},{\"name\":\"_endDate\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"payable\":true,\"stateMutability\":\"payable\",\"type\":\"fallback\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"_buyer\",\"type\":\"address\"},{\"indexed\":false,\"name\":\"_tokens\",\"type\":\"uint256\"},{\"indexed\":false,\"name\":\"_amount\",\"type\":\"uint256\"}],\"name\":\"TokenBought\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"_receiver\",\"type\":\"address\"},{\"indexed\":false,\"name\":\"_tokens\",\"type\":\"uint256\"},{\"indexed\":false,\"name\":\"_equivalentAmount\",\"type\":\"uint256\"}],\"name\":\"TokenAdded\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"_receiver\",\"type\":\"address\"},{\"indexed\":false,\"name\":\"_tokens\",\"type\":\"uint256\"},{\"indexed\":false,\"name\":\"_equivalentAmount\",\"type\":\"uint256\"}],\"name\":\"TokenToppedUp\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"_receiver\",\"type\":\"address\"},{\"indexed\":false,\"name\":\"_tokens\",\"type\":\"uint256\"},{\"indexed\":false,\"name\":\"_equivalentAmount\",\"type\":\"uint256\"}],\"name\":\"TokenSubtracted\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"_receiver\",\"type\":\"address\"},{\"indexed\":false,\"name\":\"_tokens\",\"type\":\"uint256\"}],\"name\":\"TokenSent\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[],\"name\":\"Pause\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[],\"name\":\"Unpause\",\"type\":\"event\"}]"

// ChangeableRateCrowdsaleBin is the compiled bytecode used for deploying new contracts.
const ChangeableRateCrowdsaleBin = `0x60606040526000805460a060020a60ff021916815560058190556009805460ff19169055600f819055601055341561003657600080fd5b60405160e0806114e78339810160405280805191906020018051919060200180519190602001805191906020018051919060200180519190602001805160008054600160a060020a03338116600160a060020a03199283161790925560038054838d1690831617905560028054838c1690831617905560018054928a16929091169190911790556007869055600685905560048490556008819055915060409050805190810160409081526a0b658e154215c96f100000825261177060208084019190915260108054600181019091556000908152600e90915220815181556020820151600190910155506040805190810160409081526a16cb1c2a842b92de200000825261167660208084019190915260108054600181019091556000908152600e90915220815181556020820151600190910155506040805190810160409081526a2230aa3fc6415c4d300000825261161260208084019190915260108054600181019091556000908152600e90915220815181556020820151600190910155506040805190810160409081526a2d963855085725bc40000082526115ae60208084019190915260108054600181019091556000908152600e90915220815181556020820151600190910155506040805190810160409081526a38fbc66a4a6cef2b500000825261154a60208084019190915260108054600181019091556000908152600e90915220815181556020820151600190910155506040805190810160409081526a4461547f8c82b89a60000082526114e660208084019190915260108054600181019091556000908152600e90915220815181556020820151600190910155506040805190810160409081526a4fc6e294ce988209700000825261148260208084019190915260108054600181019091556000908152600e90915220815181556020820151600190910155506040805190810160409081526a5b2c70aa10ae4b78800000825261141e60208084019190915260108054600181019091556000908152600e90915220815181556020820151600190910155506040805190810160409081526a71f78cd494d9de56a00000825261138860208084019190915260108054600181019091556000908152600e909152208151815560208201516001909101555050600f546000908152600e602052604090205460115550505050505061113b806103ac6000396000f3006060604052600436106101665763ffffffff60e060020a6000350416630262dc1381146101705780630a3846651461018f5780632511b182146101b457806326ffee08146101c757806331711884146101e95780633c78fe07146101fc5780633f4ba83a146102345780634a7510721461025b5780634e71d92d146102895780635c975abb1461029c57806361241c28146102af5780637822ed49146102c55780637b352962146102d85780638456cb59146102eb5780638da5cb5b146102fe578063a4bd7a2014610311578063a6f2ae3a14610166578063a72a05f714610324578063d1058e5914610350578063d4f114a614610363578063d56b288914610376578063d6f7ddf914610389578063d8b964e6146103ab578063daea85c5146103ca578063e4860339146103e9578063e69b414b14610408578063f07255e31461041b578063f2fde38b1461042e578063f5d82b6b1461044d578063fc0c546a1461046f575b61016e610482565b005b341561017b57600080fd5b61016e600160a060020a0360043516610621565b341561019a57600080fd5b6101a261065f565b60405190815260200160405180910390f35b34156101bf57600080fd5b6101a2610665565b34156101d257600080fd5b61016e600160a060020a036004351660243561066b565b34156101f457600080fd5b6101a261077b565b341561020757600080fd5b61021863ffffffff60043516610781565b604051600160a060020a03909116815260200160405180910390f35b341561023f57600080fd5b6102476107bf565b604051901515815260200160405180910390f35b341561026657600080fd5b610271600435610843565b60405191825260208201526040908101905180910390f35b341561029457600080fd5b61016e61085c565b34156102a757600080fd5b61024761087e565b34156102ba57600080fd5b61016e60043561088e565b34156102d057600080fd5b6102186108bb565b34156102e357600080fd5b6102476108ca565b34156102f657600080fd5b6102476108d3565b341561030957600080fd5b61021861095c565b341561031c57600080fd5b6101a261096b565b341561032f57600080fd5b610337610971565b60405163ffffffff909116815260200160405180910390f35b341561035b57600080fd5b61016e61099a565b341561036e57600080fd5b6101a2610a59565b341561038157600080fd5b61016e610a5f565b341561039457600080fd5b61016e600160a060020a0360043516602435610b87565b34156103b657600080fd5b610247600160a060020a0360043516610c39565b34156103d557600080fd5b61016e600160a060020a0360043516610c4e565b34156103f457600080fd5b6101a2600160a060020a0360043516610ca4565b341561041357600080fd5b6101a2610cb6565b341561042657600080fd5b6101a2610cbc565b341561043957600080fd5b61016e600160a060020a0360043516610cc2565b341561045857600080fd5b61016e600160a060020a0360043516602435610d19565b341561047a57600080fd5b610218610dfa565b6009546000908190819060ff161561049957600080fd5b600454600554106104a957600080fd5b6008544211156104b857600080fd5b60005460a060020a900460ff16156104cf57600080fd5b600754349350600092506104e9908463ffffffff610e0916565b90506004548160055401111561053a5760055460045461050e9163ffffffff610e3416565b905061052560075482610e4690919063ffffffff16565b9250610537348463ffffffff610e3416565b91505b60055461054d908263ffffffff610e5d16565b600581905560045490111561056157600080fd5b61056b3382610e6c565b33600160a060020a03167f28cab0d660ed8aedd61a8c9db00b97f6a2d67e07d87795994f440b18bc5f1aa3828560405191825260208201526040908101905180910390a2600154600160a060020a031683156108fc0284604051600060405180830381858888f1935050505015156105e257600080fd5b600082111561061c57600160a060020a03331682156108fc0283604051600060405180830381858888f19350505050151561061c57600080fd5b505050565b60005433600160a060020a0390811691161461063c57600080fd5b60005460a060020a900460ff161561065357600080fd5b61065c81610e95565b50565b60045481565b60065481565b6000805433600160a060020a0390811691161461068757600080fd5b60005460a060020a900460ff161561069e57600080fd5b6007546106b1908363ffffffff610e0916565b600160a060020a0384166000908152600a6020526040902054909150819010156106da57600080fd5b600160a060020a0383166000908152600a6020526040902054610703908263ffffffff610e3416565b600160a060020a0384166000908152600a602052604090205560055461072f908263ffffffff610e3416565b600555600160a060020a0383167f2da8b2becb27b189a54d21689aa110c877379a94c6f59f115cedc91d66239b33828460405191825260208201526040908101905180910390a2505050565b60075481565b6000805433600160a060020a0390811691161461079d57600080fd5b5063ffffffff166000908152600c6020526040902054600160a060020a031690565b6000805433600160a060020a039081169116146107db57600080fd5b60005460a060020a900460ff1615156107f357600080fd5b6000805474ff0000000000000000000000000000000000000000191690557f7805862f689e2f13df9f062ff482ad3ad112aca9e0847911ed832e158c525b3360405160405180910390a150600190565b600e602052600090815260409020805460019091015482565b60005460a060020a900460ff161561087357600080fd5b61087c33610e95565b565b60005460a060020a900460ff1681565b60005433600160a060020a039081169116146108a957600080fd5b600081116108b657600080fd5b600755565b600254600160a060020a031681565b60095460ff1681565b6000805433600160a060020a039081169116146108ef57600080fd5b60005460a060020a900460ff161561090657600080fd5b6000805474ff0000000000000000000000000000000000000000191660a060020a1790557f6985a02210a168e66602d3235cb6db0e70f92b3ba4d376a33c0f3d9434bff62560405160405180910390a150600190565b600054600160a060020a031681565b600f5481565b6000805433600160a060020a0390811691161461098d57600080fd5b50600d5463ffffffff1690565b60008054819033600160a060020a039081169116146109b857600080fd5b60005460a060020a900460ff16156109cf57600080fd5b600091505b600d5463ffffffff9081169083161015610a55575063ffffffff81166000908152600c6020908152604080832054600160a060020a0316808452600b9092529091205460ff168015610a3c5750600160a060020a0381166000908152600a6020526040812054115b15610a4a57610a4a81610e95565b6001909101906109d4565b5050565b60105481565b60005433600160a060020a03908116911614610a7a57600080fd5b600454600554101580610a8e575060085442115b1515610a9957600080fd5b60095460ff1615610aa957600080fd5b6009805460ff19166001179055600354600254600160a060020a039182169163a9059cbb9116826370a082313060405160e060020a63ffffffff8416028152600160a060020a039091166004820152602401602060405180830381600087803b1515610b1457600080fd5b5af11515610b2157600080fd5b5050506040518051905060405160e060020a63ffffffff8516028152600160a060020a0390921660048301526024820152604401602060405180830381600087803b1515610b6e57600080fd5b5af11515610b7b57600080fd5b50505060405180515050565b6000805433600160a060020a03908116911614610ba357600080fd5b60005460a060020a900460ff1615610bba57600080fd5b600754610bcd908363ffffffff610e0916565b600554909150610be3908263ffffffff610e5d16565b600555610bf08382610e6c565b82600160a060020a03167ff10143e1b37ebaf3698e98976d8bddabcea045adc939c03b245da18e81ea9651828460405191825260208201526040908101905180910390a2505050565b600b6020526000908152604090205460ff1681565b60005433600160a060020a03908116911614610c6957600080fd5b60005460a060020a900460ff1615610c8057600080fd5b600160a060020a03166000908152600b60205260409020805460ff19166001179055565b600a6020526000908152604090205481565b60055481565b60115481565b60005433600160a060020a03908116911614610cdd57600080fd5b600160a060020a0381161561065c5760008054600160a060020a03831673ffffffffffffffffffffffffffffffffffffffff1990911617905550565b6000805433600160a060020a03908116911614610d3557600080fd5b60095460ff1615610d4557600080fd5b60045460055410610d5557600080fd5b600854421115610d6457600080fd5b60005460a060020a900460ff1615610d7b57600080fd5b600754610d8e908363ffffffff610e0916565b600554909150610da4908263ffffffff610e5d16565b600555610db18382610e6c565b82600160a060020a03167fa818a22273fc309f0a3682b642c74c5b5c25c0615ff378d07767cd231e19fffc828460405191825260208201526040908101905180910390a2505050565b600354600160a060020a031681565b6000828202831580610e255750828482811515610e2257fe5b04145b1515610e2d57fe5b9392505050565b600082821115610e4057fe5b50900390565b6000808284811515610e5457fe5b04949350505050565b600082820183811015610e2d57fe5b60005460a060020a900460ff1615610e8357600080fd5b610e8d8282610fd5565b610a556110c6565b6000805460a060020a900460ff1615610ead57600080fd5b600160a060020a0382166000908152600b602052604090205460ff161515610ed457600080fd5b600160a060020a0382166000908152600a602052604081205411610ef757600080fd5b50600160a060020a038082166000908152600a602052604080822080549290556003546002549293908116926323b872dd92911690859085905160e060020a63ffffffff8616028152600160a060020a0393841660048201529190921660248201526044810191909152606401602060405180830381600087803b1515610f7d57600080fd5b5af11515610f8a57600080fd5b50505060405180515050600160a060020a0382167fd4ff88d569a7ad2bfd6b17da9dbe82e2ccc016fd0051e08ff1ad1e6fe8fc9b558260405190815260200160405180910390a25050565b60005460a060020a900460ff1615610fec57600080fd5b600160a060020a0382166000908152600a6020526040902054151561107d57600d805463ffffffff9081166000908152600c60209081526040808320805473ffffffffffffffffffffffffffffffffffffffff1916600160a060020a038916908117909155855463ffffffff19811690861660010190951694909417909455918152600b90915220805460ff191690555b600160a060020a0382166000908152600a60205260409020546110a6908263ffffffff610e5d16565b600160a060020a039092166000908152600a602052604090209190915550565b6011546005541061087c57600f80546001019081905560105490106110f05760045460115561087c565b600f546000908152600e602052604090208054601155600101546007555600a165627a7a72305820f05a3a93e53fad0f81f7259094d84f917943ee37e3996e5790c00590abfb94fc0029`

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
	return address, tx, &ChangeableRateCrowdsale{ChangeableRateCrowdsaleCaller: ChangeableRateCrowdsaleCaller{contract: contract}, ChangeableRateCrowdsaleTransactor: ChangeableRateCrowdsaleTransactor{contract: contract}, ChangeableRateCrowdsaleFilterer: ChangeableRateCrowdsaleFilterer{contract: contract}}, nil
}

// ChangeableRateCrowdsale is an auto generated Go binding around an Ethereum contract.
type ChangeableRateCrowdsale struct {
	ChangeableRateCrowdsaleCaller     // Read-only binding to the contract
	ChangeableRateCrowdsaleTransactor // Write-only binding to the contract
	ChangeableRateCrowdsaleFilterer   // Log filterer for contract events
}

// ChangeableRateCrowdsaleCaller is an auto generated read-only Go binding around an Ethereum contract.
type ChangeableRateCrowdsaleCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ChangeableRateCrowdsaleTransactor is an auto generated write-only Go binding around an Ethereum contract.
type ChangeableRateCrowdsaleTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ChangeableRateCrowdsaleFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type ChangeableRateCrowdsaleFilterer struct {
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
	contract, err := bindChangeableRateCrowdsale(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &ChangeableRateCrowdsale{ChangeableRateCrowdsaleCaller: ChangeableRateCrowdsaleCaller{contract: contract}, ChangeableRateCrowdsaleTransactor: ChangeableRateCrowdsaleTransactor{contract: contract}, ChangeableRateCrowdsaleFilterer: ChangeableRateCrowdsaleFilterer{contract: contract}}, nil
}

// NewChangeableRateCrowdsaleCaller creates a new read-only instance of ChangeableRateCrowdsale, bound to a specific deployed contract.
func NewChangeableRateCrowdsaleCaller(address common.Address, caller bind.ContractCaller) (*ChangeableRateCrowdsaleCaller, error) {
	contract, err := bindChangeableRateCrowdsale(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &ChangeableRateCrowdsaleCaller{contract: contract}, nil
}

// NewChangeableRateCrowdsaleTransactor creates a new write-only instance of ChangeableRateCrowdsale, bound to a specific deployed contract.
func NewChangeableRateCrowdsaleTransactor(address common.Address, transactor bind.ContractTransactor) (*ChangeableRateCrowdsaleTransactor, error) {
	contract, err := bindChangeableRateCrowdsale(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &ChangeableRateCrowdsaleTransactor{contract: contract}, nil
}

// NewChangeableRateCrowdsaleFilterer creates a new log filterer instance of ChangeableRateCrowdsale, bound to a specific deployed contract.
func NewChangeableRateCrowdsaleFilterer(address common.Address, filterer bind.ContractFilterer) (*ChangeableRateCrowdsaleFilterer, error) {
	contract, err := bindChangeableRateCrowdsale(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &ChangeableRateCrowdsaleFilterer{contract: contract}, nil
}

// bindChangeableRateCrowdsale binds a generic wrapper to an already deployed contract.
func bindChangeableRateCrowdsale(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(ChangeableRateCrowdsaleABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
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

// ChangeableRateCrowdsalePauseIterator is returned from FilterPause and is used to iterate over the raw logs and unpacked data for Pause events raised by the ChangeableRateCrowdsale contract.
type ChangeableRateCrowdsalePauseIterator struct {
	Event *ChangeableRateCrowdsalePause // Event containing the contract specifics and raw log

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
func (it *ChangeableRateCrowdsalePauseIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ChangeableRateCrowdsalePause)
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
		it.Event = new(ChangeableRateCrowdsalePause)
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
func (it *ChangeableRateCrowdsalePauseIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ChangeableRateCrowdsalePauseIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ChangeableRateCrowdsalePause represents a Pause event raised by the ChangeableRateCrowdsale contract.
type ChangeableRateCrowdsalePause struct {
	Raw types.Log // Blockchain specific contextual infos
}

// FilterPause is a free log retrieval operation binding the contract event 0x6985a02210a168e66602d3235cb6db0e70f92b3ba4d376a33c0f3d9434bff625.
//
// Solidity: event Pause()
func (_ChangeableRateCrowdsale *ChangeableRateCrowdsaleFilterer) FilterPause(opts *bind.FilterOpts) (*ChangeableRateCrowdsalePauseIterator, error) {

	logs, sub, err := _ChangeableRateCrowdsale.contract.FilterLogs(opts, "Pause")
	if err != nil {
		return nil, err
	}
	return &ChangeableRateCrowdsalePauseIterator{contract: _ChangeableRateCrowdsale.contract, event: "Pause", logs: logs, sub: sub}, nil
}

// WatchPause is a free log subscription operation binding the contract event 0x6985a02210a168e66602d3235cb6db0e70f92b3ba4d376a33c0f3d9434bff625.
//
// Solidity: event Pause()
func (_ChangeableRateCrowdsale *ChangeableRateCrowdsaleFilterer) WatchPause(opts *bind.WatchOpts, sink chan<- *ChangeableRateCrowdsalePause) (event.Subscription, error) {

	logs, sub, err := _ChangeableRateCrowdsale.contract.WatchLogs(opts, "Pause")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ChangeableRateCrowdsalePause)
				if err := _ChangeableRateCrowdsale.contract.UnpackLog(event, "Pause", log); err != nil {
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

// ChangeableRateCrowdsaleTokenAddedIterator is returned from FilterTokenAdded and is used to iterate over the raw logs and unpacked data for TokenAdded events raised by the ChangeableRateCrowdsale contract.
type ChangeableRateCrowdsaleTokenAddedIterator struct {
	Event *ChangeableRateCrowdsaleTokenAdded // Event containing the contract specifics and raw log

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
func (it *ChangeableRateCrowdsaleTokenAddedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ChangeableRateCrowdsaleTokenAdded)
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
		it.Event = new(ChangeableRateCrowdsaleTokenAdded)
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
func (it *ChangeableRateCrowdsaleTokenAddedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ChangeableRateCrowdsaleTokenAddedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ChangeableRateCrowdsaleTokenAdded represents a TokenAdded event raised by the ChangeableRateCrowdsale contract.
type ChangeableRateCrowdsaleTokenAdded struct {
	Receiver         common.Address
	Tokens           *big.Int
	EquivalentAmount *big.Int
	Raw              types.Log // Blockchain specific contextual infos
}

// FilterTokenAdded is a free log retrieval operation binding the contract event 0xa818a22273fc309f0a3682b642c74c5b5c25c0615ff378d07767cd231e19fffc.
//
// Solidity: event TokenAdded(_receiver indexed address, _tokens uint256, _equivalentAmount uint256)
func (_ChangeableRateCrowdsale *ChangeableRateCrowdsaleFilterer) FilterTokenAdded(opts *bind.FilterOpts, _receiver []common.Address) (*ChangeableRateCrowdsaleTokenAddedIterator, error) {

	var _receiverRule []interface{}
	for _, _receiverItem := range _receiver {
		_receiverRule = append(_receiverRule, _receiverItem)
	}

	logs, sub, err := _ChangeableRateCrowdsale.contract.FilterLogs(opts, "TokenAdded", _receiverRule)
	if err != nil {
		return nil, err
	}
	return &ChangeableRateCrowdsaleTokenAddedIterator{contract: _ChangeableRateCrowdsale.contract, event: "TokenAdded", logs: logs, sub: sub}, nil
}

// WatchTokenAdded is a free log subscription operation binding the contract event 0xa818a22273fc309f0a3682b642c74c5b5c25c0615ff378d07767cd231e19fffc.
//
// Solidity: event TokenAdded(_receiver indexed address, _tokens uint256, _equivalentAmount uint256)
func (_ChangeableRateCrowdsale *ChangeableRateCrowdsaleFilterer) WatchTokenAdded(opts *bind.WatchOpts, sink chan<- *ChangeableRateCrowdsaleTokenAdded, _receiver []common.Address) (event.Subscription, error) {

	var _receiverRule []interface{}
	for _, _receiverItem := range _receiver {
		_receiverRule = append(_receiverRule, _receiverItem)
	}

	logs, sub, err := _ChangeableRateCrowdsale.contract.WatchLogs(opts, "TokenAdded", _receiverRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ChangeableRateCrowdsaleTokenAdded)
				if err := _ChangeableRateCrowdsale.contract.UnpackLog(event, "TokenAdded", log); err != nil {
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

// ChangeableRateCrowdsaleTokenBoughtIterator is returned from FilterTokenBought and is used to iterate over the raw logs and unpacked data for TokenBought events raised by the ChangeableRateCrowdsale contract.
type ChangeableRateCrowdsaleTokenBoughtIterator struct {
	Event *ChangeableRateCrowdsaleTokenBought // Event containing the contract specifics and raw log

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
func (it *ChangeableRateCrowdsaleTokenBoughtIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ChangeableRateCrowdsaleTokenBought)
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
		it.Event = new(ChangeableRateCrowdsaleTokenBought)
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
func (it *ChangeableRateCrowdsaleTokenBoughtIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ChangeableRateCrowdsaleTokenBoughtIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ChangeableRateCrowdsaleTokenBought represents a TokenBought event raised by the ChangeableRateCrowdsale contract.
type ChangeableRateCrowdsaleTokenBought struct {
	Buyer  common.Address
	Tokens *big.Int
	Amount *big.Int
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterTokenBought is a free log retrieval operation binding the contract event 0x28cab0d660ed8aedd61a8c9db00b97f6a2d67e07d87795994f440b18bc5f1aa3.
//
// Solidity: event TokenBought(_buyer indexed address, _tokens uint256, _amount uint256)
func (_ChangeableRateCrowdsale *ChangeableRateCrowdsaleFilterer) FilterTokenBought(opts *bind.FilterOpts, _buyer []common.Address) (*ChangeableRateCrowdsaleTokenBoughtIterator, error) {

	var _buyerRule []interface{}
	for _, _buyerItem := range _buyer {
		_buyerRule = append(_buyerRule, _buyerItem)
	}

	logs, sub, err := _ChangeableRateCrowdsale.contract.FilterLogs(opts, "TokenBought", _buyerRule)
	if err != nil {
		return nil, err
	}
	return &ChangeableRateCrowdsaleTokenBoughtIterator{contract: _ChangeableRateCrowdsale.contract, event: "TokenBought", logs: logs, sub: sub}, nil
}

// WatchTokenBought is a free log subscription operation binding the contract event 0x28cab0d660ed8aedd61a8c9db00b97f6a2d67e07d87795994f440b18bc5f1aa3.
//
// Solidity: event TokenBought(_buyer indexed address, _tokens uint256, _amount uint256)
func (_ChangeableRateCrowdsale *ChangeableRateCrowdsaleFilterer) WatchTokenBought(opts *bind.WatchOpts, sink chan<- *ChangeableRateCrowdsaleTokenBought, _buyer []common.Address) (event.Subscription, error) {

	var _buyerRule []interface{}
	for _, _buyerItem := range _buyer {
		_buyerRule = append(_buyerRule, _buyerItem)
	}

	logs, sub, err := _ChangeableRateCrowdsale.contract.WatchLogs(opts, "TokenBought", _buyerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ChangeableRateCrowdsaleTokenBought)
				if err := _ChangeableRateCrowdsale.contract.UnpackLog(event, "TokenBought", log); err != nil {
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

// ChangeableRateCrowdsaleTokenSentIterator is returned from FilterTokenSent and is used to iterate over the raw logs and unpacked data for TokenSent events raised by the ChangeableRateCrowdsale contract.
type ChangeableRateCrowdsaleTokenSentIterator struct {
	Event *ChangeableRateCrowdsaleTokenSent // Event containing the contract specifics and raw log

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
func (it *ChangeableRateCrowdsaleTokenSentIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ChangeableRateCrowdsaleTokenSent)
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
		it.Event = new(ChangeableRateCrowdsaleTokenSent)
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
func (it *ChangeableRateCrowdsaleTokenSentIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ChangeableRateCrowdsaleTokenSentIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ChangeableRateCrowdsaleTokenSent represents a TokenSent event raised by the ChangeableRateCrowdsale contract.
type ChangeableRateCrowdsaleTokenSent struct {
	Receiver common.Address
	Tokens   *big.Int
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterTokenSent is a free log retrieval operation binding the contract event 0xd4ff88d569a7ad2bfd6b17da9dbe82e2ccc016fd0051e08ff1ad1e6fe8fc9b55.
//
// Solidity: event TokenSent(_receiver indexed address, _tokens uint256)
func (_ChangeableRateCrowdsale *ChangeableRateCrowdsaleFilterer) FilterTokenSent(opts *bind.FilterOpts, _receiver []common.Address) (*ChangeableRateCrowdsaleTokenSentIterator, error) {

	var _receiverRule []interface{}
	for _, _receiverItem := range _receiver {
		_receiverRule = append(_receiverRule, _receiverItem)
	}

	logs, sub, err := _ChangeableRateCrowdsale.contract.FilterLogs(opts, "TokenSent", _receiverRule)
	if err != nil {
		return nil, err
	}
	return &ChangeableRateCrowdsaleTokenSentIterator{contract: _ChangeableRateCrowdsale.contract, event: "TokenSent", logs: logs, sub: sub}, nil
}

// WatchTokenSent is a free log subscription operation binding the contract event 0xd4ff88d569a7ad2bfd6b17da9dbe82e2ccc016fd0051e08ff1ad1e6fe8fc9b55.
//
// Solidity: event TokenSent(_receiver indexed address, _tokens uint256)
func (_ChangeableRateCrowdsale *ChangeableRateCrowdsaleFilterer) WatchTokenSent(opts *bind.WatchOpts, sink chan<- *ChangeableRateCrowdsaleTokenSent, _receiver []common.Address) (event.Subscription, error) {

	var _receiverRule []interface{}
	for _, _receiverItem := range _receiver {
		_receiverRule = append(_receiverRule, _receiverItem)
	}

	logs, sub, err := _ChangeableRateCrowdsale.contract.WatchLogs(opts, "TokenSent", _receiverRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ChangeableRateCrowdsaleTokenSent)
				if err := _ChangeableRateCrowdsale.contract.UnpackLog(event, "TokenSent", log); err != nil {
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

// ChangeableRateCrowdsaleTokenSubtractedIterator is returned from FilterTokenSubtracted and is used to iterate over the raw logs and unpacked data for TokenSubtracted events raised by the ChangeableRateCrowdsale contract.
type ChangeableRateCrowdsaleTokenSubtractedIterator struct {
	Event *ChangeableRateCrowdsaleTokenSubtracted // Event containing the contract specifics and raw log

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
func (it *ChangeableRateCrowdsaleTokenSubtractedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ChangeableRateCrowdsaleTokenSubtracted)
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
		it.Event = new(ChangeableRateCrowdsaleTokenSubtracted)
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
func (it *ChangeableRateCrowdsaleTokenSubtractedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ChangeableRateCrowdsaleTokenSubtractedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ChangeableRateCrowdsaleTokenSubtracted represents a TokenSubtracted event raised by the ChangeableRateCrowdsale contract.
type ChangeableRateCrowdsaleTokenSubtracted struct {
	Receiver         common.Address
	Tokens           *big.Int
	EquivalentAmount *big.Int
	Raw              types.Log // Blockchain specific contextual infos
}

// FilterTokenSubtracted is a free log retrieval operation binding the contract event 0x2da8b2becb27b189a54d21689aa110c877379a94c6f59f115cedc91d66239b33.
//
// Solidity: event TokenSubtracted(_receiver indexed address, _tokens uint256, _equivalentAmount uint256)
func (_ChangeableRateCrowdsale *ChangeableRateCrowdsaleFilterer) FilterTokenSubtracted(opts *bind.FilterOpts, _receiver []common.Address) (*ChangeableRateCrowdsaleTokenSubtractedIterator, error) {

	var _receiverRule []interface{}
	for _, _receiverItem := range _receiver {
		_receiverRule = append(_receiverRule, _receiverItem)
	}

	logs, sub, err := _ChangeableRateCrowdsale.contract.FilterLogs(opts, "TokenSubtracted", _receiverRule)
	if err != nil {
		return nil, err
	}
	return &ChangeableRateCrowdsaleTokenSubtractedIterator{contract: _ChangeableRateCrowdsale.contract, event: "TokenSubtracted", logs: logs, sub: sub}, nil
}

// WatchTokenSubtracted is a free log subscription operation binding the contract event 0x2da8b2becb27b189a54d21689aa110c877379a94c6f59f115cedc91d66239b33.
//
// Solidity: event TokenSubtracted(_receiver indexed address, _tokens uint256, _equivalentAmount uint256)
func (_ChangeableRateCrowdsale *ChangeableRateCrowdsaleFilterer) WatchTokenSubtracted(opts *bind.WatchOpts, sink chan<- *ChangeableRateCrowdsaleTokenSubtracted, _receiver []common.Address) (event.Subscription, error) {

	var _receiverRule []interface{}
	for _, _receiverItem := range _receiver {
		_receiverRule = append(_receiverRule, _receiverItem)
	}

	logs, sub, err := _ChangeableRateCrowdsale.contract.WatchLogs(opts, "TokenSubtracted", _receiverRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ChangeableRateCrowdsaleTokenSubtracted)
				if err := _ChangeableRateCrowdsale.contract.UnpackLog(event, "TokenSubtracted", log); err != nil {
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

// ChangeableRateCrowdsaleTokenToppedUpIterator is returned from FilterTokenToppedUp and is used to iterate over the raw logs and unpacked data for TokenToppedUp events raised by the ChangeableRateCrowdsale contract.
type ChangeableRateCrowdsaleTokenToppedUpIterator struct {
	Event *ChangeableRateCrowdsaleTokenToppedUp // Event containing the contract specifics and raw log

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
func (it *ChangeableRateCrowdsaleTokenToppedUpIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ChangeableRateCrowdsaleTokenToppedUp)
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
		it.Event = new(ChangeableRateCrowdsaleTokenToppedUp)
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
func (it *ChangeableRateCrowdsaleTokenToppedUpIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ChangeableRateCrowdsaleTokenToppedUpIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ChangeableRateCrowdsaleTokenToppedUp represents a TokenToppedUp event raised by the ChangeableRateCrowdsale contract.
type ChangeableRateCrowdsaleTokenToppedUp struct {
	Receiver         common.Address
	Tokens           *big.Int
	EquivalentAmount *big.Int
	Raw              types.Log // Blockchain specific contextual infos
}

// FilterTokenToppedUp is a free log retrieval operation binding the contract event 0xf10143e1b37ebaf3698e98976d8bddabcea045adc939c03b245da18e81ea9651.
//
// Solidity: event TokenToppedUp(_receiver indexed address, _tokens uint256, _equivalentAmount uint256)
func (_ChangeableRateCrowdsale *ChangeableRateCrowdsaleFilterer) FilterTokenToppedUp(opts *bind.FilterOpts, _receiver []common.Address) (*ChangeableRateCrowdsaleTokenToppedUpIterator, error) {

	var _receiverRule []interface{}
	for _, _receiverItem := range _receiver {
		_receiverRule = append(_receiverRule, _receiverItem)
	}

	logs, sub, err := _ChangeableRateCrowdsale.contract.FilterLogs(opts, "TokenToppedUp", _receiverRule)
	if err != nil {
		return nil, err
	}
	return &ChangeableRateCrowdsaleTokenToppedUpIterator{contract: _ChangeableRateCrowdsale.contract, event: "TokenToppedUp", logs: logs, sub: sub}, nil
}

// WatchTokenToppedUp is a free log subscription operation binding the contract event 0xf10143e1b37ebaf3698e98976d8bddabcea045adc939c03b245da18e81ea9651.
//
// Solidity: event TokenToppedUp(_receiver indexed address, _tokens uint256, _equivalentAmount uint256)
func (_ChangeableRateCrowdsale *ChangeableRateCrowdsaleFilterer) WatchTokenToppedUp(opts *bind.WatchOpts, sink chan<- *ChangeableRateCrowdsaleTokenToppedUp, _receiver []common.Address) (event.Subscription, error) {

	var _receiverRule []interface{}
	for _, _receiverItem := range _receiver {
		_receiverRule = append(_receiverRule, _receiverItem)
	}

	logs, sub, err := _ChangeableRateCrowdsale.contract.WatchLogs(opts, "TokenToppedUp", _receiverRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ChangeableRateCrowdsaleTokenToppedUp)
				if err := _ChangeableRateCrowdsale.contract.UnpackLog(event, "TokenToppedUp", log); err != nil {
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

// ChangeableRateCrowdsaleUnpauseIterator is returned from FilterUnpause and is used to iterate over the raw logs and unpacked data for Unpause events raised by the ChangeableRateCrowdsale contract.
type ChangeableRateCrowdsaleUnpauseIterator struct {
	Event *ChangeableRateCrowdsaleUnpause // Event containing the contract specifics and raw log

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
func (it *ChangeableRateCrowdsaleUnpauseIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ChangeableRateCrowdsaleUnpause)
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
		it.Event = new(ChangeableRateCrowdsaleUnpause)
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
func (it *ChangeableRateCrowdsaleUnpauseIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ChangeableRateCrowdsaleUnpauseIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ChangeableRateCrowdsaleUnpause represents a Unpause event raised by the ChangeableRateCrowdsale contract.
type ChangeableRateCrowdsaleUnpause struct {
	Raw types.Log // Blockchain specific contextual infos
}

// FilterUnpause is a free log retrieval operation binding the contract event 0x7805862f689e2f13df9f062ff482ad3ad112aca9e0847911ed832e158c525b33.
//
// Solidity: event Unpause()
func (_ChangeableRateCrowdsale *ChangeableRateCrowdsaleFilterer) FilterUnpause(opts *bind.FilterOpts) (*ChangeableRateCrowdsaleUnpauseIterator, error) {

	logs, sub, err := _ChangeableRateCrowdsale.contract.FilterLogs(opts, "Unpause")
	if err != nil {
		return nil, err
	}
	return &ChangeableRateCrowdsaleUnpauseIterator{contract: _ChangeableRateCrowdsale.contract, event: "Unpause", logs: logs, sub: sub}, nil
}

// WatchUnpause is a free log subscription operation binding the contract event 0x7805862f689e2f13df9f062ff482ad3ad112aca9e0847911ed832e158c525b33.
//
// Solidity: event Unpause()
func (_ChangeableRateCrowdsale *ChangeableRateCrowdsaleFilterer) WatchUnpause(opts *bind.WatchOpts, sink chan<- *ChangeableRateCrowdsaleUnpause) (event.Subscription, error) {

	logs, sub, err := _ChangeableRateCrowdsale.contract.WatchLogs(opts, "Unpause")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ChangeableRateCrowdsaleUnpause)
				if err := _ChangeableRateCrowdsale.contract.UnpackLog(event, "Unpause", log); err != nil {
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

// ClaimableCrowdsaleABI is the input ABI used to generate the binding from.
const ClaimableCrowdsaleABI = "[{\"constant\":false,\"inputs\":[{\"name\":\"_receiver\",\"type\":\"address\"}],\"name\":\"claimOne\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"maxTokensAmount\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"minBuyableAmount\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_receiver\",\"type\":\"address\"},{\"name\":\"_equivalentEthAmount\",\"type\":\"uint256\"}],\"name\":\"sub\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"tokenRate\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"i\",\"type\":\"uint32\"}],\"name\":\"getReceiver\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[],\"name\":\"unpause\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[],\"name\":\"claim\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"paused\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_tokenRate\",\"type\":\"uint256\"}],\"name\":\"setTokenRate\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"bankAddress\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"isFinished\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[],\"name\":\"pause\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[],\"name\":\"buy\",\"outputs\":[],\"payable\":true,\"stateMutability\":\"payable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"getReceiversCount\",\"outputs\":[{\"name\":\"\",\"type\":\"uint32\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[],\"name\":\"claimAll\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[],\"name\":\"finish\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_receiver\",\"type\":\"address\"},{\"name\":\"_equivalentEthAmount\",\"type\":\"uint256\"}],\"name\":\"topUp\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"\",\"type\":\"address\"}],\"name\":\"approved\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_receiver\",\"type\":\"address\"}],\"name\":\"approve\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"\",\"type\":\"address\"}],\"name\":\"tokens\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"issuedTokensAmount\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"transferOwnership\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_receiver\",\"type\":\"address\"},{\"name\":\"_equivalentEthAmount\",\"type\":\"uint256\"}],\"name\":\"add\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"token\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"name\":\"_tokenAddress\",\"type\":\"address\"},{\"name\":\"_bankAddress\",\"type\":\"address\"},{\"name\":\"_beneficiaryAddress\",\"type\":\"address\"},{\"name\":\"_tokenRate\",\"type\":\"uint256\"},{\"name\":\"_minBuyableAmount\",\"type\":\"uint256\"},{\"name\":\"_maxTokensAmount\",\"type\":\"uint256\"},{\"name\":\"_endDate\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"payable\":true,\"stateMutability\":\"payable\",\"type\":\"fallback\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"_buyer\",\"type\":\"address\"},{\"indexed\":false,\"name\":\"_tokens\",\"type\":\"uint256\"},{\"indexed\":false,\"name\":\"_amount\",\"type\":\"uint256\"}],\"name\":\"TokenBought\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"_receiver\",\"type\":\"address\"},{\"indexed\":false,\"name\":\"_tokens\",\"type\":\"uint256\"},{\"indexed\":false,\"name\":\"_equivalentAmount\",\"type\":\"uint256\"}],\"name\":\"TokenAdded\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"_receiver\",\"type\":\"address\"},{\"indexed\":false,\"name\":\"_tokens\",\"type\":\"uint256\"},{\"indexed\":false,\"name\":\"_equivalentAmount\",\"type\":\"uint256\"}],\"name\":\"TokenToppedUp\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"_receiver\",\"type\":\"address\"},{\"indexed\":false,\"name\":\"_tokens\",\"type\":\"uint256\"},{\"indexed\":false,\"name\":\"_equivalentAmount\",\"type\":\"uint256\"}],\"name\":\"TokenSubtracted\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"_receiver\",\"type\":\"address\"},{\"indexed\":false,\"name\":\"_tokens\",\"type\":\"uint256\"}],\"name\":\"TokenSent\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[],\"name\":\"Pause\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[],\"name\":\"Unpause\",\"type\":\"event\"}]"

// ClaimableCrowdsaleBin is the compiled bytecode used for deploying new contracts.
const ClaimableCrowdsaleBin = `0x60606040526000805460a060020a60ff02191681556005556009805460ff19169055341561002c57600080fd5b60405160e0806110e88339810160405280805191906020018051919060200180519190602001805191906020018051919060200180519190602001805160008054600160a060020a03338116600160a060020a031992831617909255600380549b83169b82169b909b17909a5560028054998216998b169990991790985560018054979098169690981695909517909555600792909255600655600455505060085561100b806100dd6000396000f30060606040526004361061013a5763ffffffff60e060020a6000350416630262dc1381146101445780630a384665146101635780632511b1821461018857806326ffee081461019b57806331711884146101bd5780633c78fe07146101d05780633f4ba83a146102085780634e71d92d1461022f5780635c975abb1461024257806361241c28146102555780637822ed491461026b5780637b3529621461027e5780638456cb59146102915780638da5cb5b146102a4578063a6f2ae3a1461013a578063a72a05f7146102b7578063d1058e59146102e3578063d56b2889146102f6578063d6f7ddf914610309578063d8b964e61461032b578063daea85c51461034a578063e486033914610369578063e69b414b14610388578063f2fde38b1461039b578063f5d82b6b146103ba578063fc0c546a146103dc575b6101426103ef565b005b341561014f57600080fd5b610142600160a060020a036004351661058e565b341561016e57600080fd5b6101766105cc565b60405190815260200160405180910390f35b341561019357600080fd5b6101766105d2565b34156101a657600080fd5b610142600160a060020a03600435166024356105d8565b34156101c857600080fd5b6101766106e8565b34156101db57600080fd5b6101ec63ffffffff600435166106ee565b604051600160a060020a03909116815260200160405180910390f35b341561021357600080fd5b61021b61072c565b604051901515815260200160405180910390f35b341561023a57600080fd5b6101426107b0565b341561024d57600080fd5b61021b6107d2565b341561026057600080fd5b6101426004356107e2565b341561027657600080fd5b6101ec61080f565b341561028957600080fd5b61021b61081e565b341561029c57600080fd5b61021b610827565b34156102af57600080fd5b6101ec6108b0565b34156102c257600080fd5b6102ca6108bf565b60405163ffffffff909116815260200160405180910390f35b34156102ee57600080fd5b6101426108e8565b341561030157600080fd5b6101426109a7565b341561031457600080fd5b610142600160a060020a0360043516602435610acf565b341561033657600080fd5b61021b600160a060020a0360043516610b81565b341561035557600080fd5b610142600160a060020a0360043516610b96565b341561037457600080fd5b610176600160a060020a0360043516610bec565b341561039357600080fd5b610176610bfe565b34156103a657600080fd5b610142600160a060020a0360043516610c04565b34156103c557600080fd5b610142600160a060020a0360043516602435610c5b565b34156103e757600080fd5b6101ec610d3c565b6009546000908190819060ff161561040657600080fd5b6004546005541061041657600080fd5b60085442111561042557600080fd5b60005460a060020a900460ff161561043c57600080fd5b60075434935060009250610456908463ffffffff610d4b16565b9050600454816005540111156104a75760055460045461047b9163ffffffff610d7616565b905061049260075482610d8890919063ffffffff16565b92506104a4348463ffffffff610d7616565b91505b6005546104ba908263ffffffff610d9f16565b60058190556004549011156104ce57600080fd5b6104d83382610dae565b33600160a060020a03167f28cab0d660ed8aedd61a8c9db00b97f6a2d67e07d87795994f440b18bc5f1aa3828560405191825260208201526040908101905180910390a2600154600160a060020a031683156108fc0284604051600060405180830381858888f19350505050151561054f57600080fd5b600082111561058957600160a060020a03331682156108fc0283604051600060405180830381858888f19350505050151561058957600080fd5b505050565b60005433600160a060020a039081169116146105a957600080fd5b60005460a060020a900460ff16156105c057600080fd5b6105c981610e9f565b50565b60045481565b60065481565b6000805433600160a060020a039081169116146105f457600080fd5b60005460a060020a900460ff161561060b57600080fd5b60075461061e908363ffffffff610d4b16565b600160a060020a0384166000908152600a60205260409020549091508190101561064757600080fd5b600160a060020a0383166000908152600a6020526040902054610670908263ffffffff610d7616565b600160a060020a0384166000908152600a602052604090205560055461069c908263ffffffff610d7616565b600555600160a060020a0383167f2da8b2becb27b189a54d21689aa110c877379a94c6f59f115cedc91d66239b33828460405191825260208201526040908101905180910390a2505050565b60075481565b6000805433600160a060020a0390811691161461070a57600080fd5b5063ffffffff166000908152600c6020526040902054600160a060020a031690565b6000805433600160a060020a0390811691161461074857600080fd5b60005460a060020a900460ff16151561076057600080fd5b6000805474ff0000000000000000000000000000000000000000191690557f7805862f689e2f13df9f062ff482ad3ad112aca9e0847911ed832e158c525b3360405160405180910390a150600190565b60005460a060020a900460ff16156107c757600080fd5b6107d033610e9f565b565b60005460a060020a900460ff1681565b60005433600160a060020a039081169116146107fd57600080fd5b6000811161080a57600080fd5b600755565b600254600160a060020a031681565b60095460ff1681565b6000805433600160a060020a0390811691161461084357600080fd5b60005460a060020a900460ff161561085a57600080fd5b6000805474ff0000000000000000000000000000000000000000191660a060020a1790557f6985a02210a168e66602d3235cb6db0e70f92b3ba4d376a33c0f3d9434bff62560405160405180910390a150600190565b600054600160a060020a031681565b6000805433600160a060020a039081169116146108db57600080fd5b50600d5463ffffffff1690565b60008054819033600160a060020a0390811691161461090657600080fd5b60005460a060020a900460ff161561091d57600080fd5b600091505b600d5463ffffffff90811690831610156109a3575063ffffffff81166000908152600c6020908152604080832054600160a060020a0316808452600b9092529091205460ff16801561098a5750600160a060020a0381166000908152600a6020526040812054115b156109985761099881610e9f565b600190910190610922565b5050565b60005433600160a060020a039081169116146109c257600080fd5b6004546005541015806109d6575060085442115b15156109e157600080fd5b60095460ff16156109f157600080fd5b6009805460ff19166001179055600354600254600160a060020a039182169163a9059cbb9116826370a082313060405160e060020a63ffffffff8416028152600160a060020a039091166004820152602401602060405180830381600087803b1515610a5c57600080fd5b5af11515610a6957600080fd5b5050506040518051905060405160e060020a63ffffffff8516028152600160a060020a0390921660048301526024820152604401602060405180830381600087803b1515610ab657600080fd5b5af11515610ac357600080fd5b50505060405180515050565b6000805433600160a060020a03908116911614610aeb57600080fd5b60005460a060020a900460ff1615610b0257600080fd5b600754610b15908363ffffffff610d4b16565b600554909150610b2b908263ffffffff610d9f16565b600555610b388382610dae565b82600160a060020a03167ff10143e1b37ebaf3698e98976d8bddabcea045adc939c03b245da18e81ea9651828460405191825260208201526040908101905180910390a2505050565b600b6020526000908152604090205460ff1681565b60005433600160a060020a03908116911614610bb157600080fd5b60005460a060020a900460ff1615610bc857600080fd5b600160a060020a03166000908152600b60205260409020805460ff19166001179055565b600a6020526000908152604090205481565b60055481565b60005433600160a060020a03908116911614610c1f57600080fd5b600160a060020a038116156105c95760008054600160a060020a03831673ffffffffffffffffffffffffffffffffffffffff1990911617905550565b6000805433600160a060020a03908116911614610c7757600080fd5b60095460ff1615610c8757600080fd5b60045460055410610c9757600080fd5b600854421115610ca657600080fd5b60005460a060020a900460ff1615610cbd57600080fd5b600754610cd0908363ffffffff610d4b16565b600554909150610ce6908263ffffffff610d9f16565b600555610cf38382610dae565b82600160a060020a03167fa818a22273fc309f0a3682b642c74c5b5c25c0615ff378d07767cd231e19fffc828460405191825260208201526040908101905180910390a2505050565b600354600160a060020a031681565b6000828202831580610d675750828482811515610d6457fe5b04145b1515610d6f57fe5b9392505050565b600082821115610d8257fe5b50900390565b6000808284811515610d9657fe5b04949350505050565b600082820183811015610d6f57fe5b60005460a060020a900460ff1615610dc557600080fd5b600160a060020a0382166000908152600a60205260409020541515610e5657600d805463ffffffff9081166000908152600c60209081526040808320805473ffffffffffffffffffffffffffffffffffffffff1916600160a060020a038916908117909155855463ffffffff19811690861660010190951694909417909455918152600b90915220805460ff191690555b600160a060020a0382166000908152600a6020526040902054610e7f908263ffffffff610d9f16565b600160a060020a039092166000908152600a602052604090209190915550565b6000805460a060020a900460ff1615610eb757600080fd5b600160a060020a0382166000908152600b602052604090205460ff161515610ede57600080fd5b600160a060020a0382166000908152600a602052604081205411610f0157600080fd5b50600160a060020a038082166000908152600a602052604080822080549290556003546002549293908116926323b872dd92911690859085905160e060020a63ffffffff8616028152600160a060020a0393841660048201529190921660248201526044810191909152606401602060405180830381600087803b1515610f8757600080fd5b5af11515610f9457600080fd5b50505060405180515050600160a060020a0382167fd4ff88d569a7ad2bfd6b17da9dbe82e2ccc016fd0051e08ff1ad1e6fe8fc9b558260405190815260200160405180910390a250505600a165627a7a72305820704156da49f9e97a9fa34828b5c3747a0f8c6596eff56ddcb9625a7b9a7225cf0029`

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
	return address, tx, &ClaimableCrowdsale{ClaimableCrowdsaleCaller: ClaimableCrowdsaleCaller{contract: contract}, ClaimableCrowdsaleTransactor: ClaimableCrowdsaleTransactor{contract: contract}, ClaimableCrowdsaleFilterer: ClaimableCrowdsaleFilterer{contract: contract}}, nil
}

// ClaimableCrowdsale is an auto generated Go binding around an Ethereum contract.
type ClaimableCrowdsale struct {
	ClaimableCrowdsaleCaller     // Read-only binding to the contract
	ClaimableCrowdsaleTransactor // Write-only binding to the contract
	ClaimableCrowdsaleFilterer   // Log filterer for contract events
}

// ClaimableCrowdsaleCaller is an auto generated read-only Go binding around an Ethereum contract.
type ClaimableCrowdsaleCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ClaimableCrowdsaleTransactor is an auto generated write-only Go binding around an Ethereum contract.
type ClaimableCrowdsaleTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ClaimableCrowdsaleFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type ClaimableCrowdsaleFilterer struct {
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
	contract, err := bindClaimableCrowdsale(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &ClaimableCrowdsale{ClaimableCrowdsaleCaller: ClaimableCrowdsaleCaller{contract: contract}, ClaimableCrowdsaleTransactor: ClaimableCrowdsaleTransactor{contract: contract}, ClaimableCrowdsaleFilterer: ClaimableCrowdsaleFilterer{contract: contract}}, nil
}

// NewClaimableCrowdsaleCaller creates a new read-only instance of ClaimableCrowdsale, bound to a specific deployed contract.
func NewClaimableCrowdsaleCaller(address common.Address, caller bind.ContractCaller) (*ClaimableCrowdsaleCaller, error) {
	contract, err := bindClaimableCrowdsale(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &ClaimableCrowdsaleCaller{contract: contract}, nil
}

// NewClaimableCrowdsaleTransactor creates a new write-only instance of ClaimableCrowdsale, bound to a specific deployed contract.
func NewClaimableCrowdsaleTransactor(address common.Address, transactor bind.ContractTransactor) (*ClaimableCrowdsaleTransactor, error) {
	contract, err := bindClaimableCrowdsale(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &ClaimableCrowdsaleTransactor{contract: contract}, nil
}

// NewClaimableCrowdsaleFilterer creates a new log filterer instance of ClaimableCrowdsale, bound to a specific deployed contract.
func NewClaimableCrowdsaleFilterer(address common.Address, filterer bind.ContractFilterer) (*ClaimableCrowdsaleFilterer, error) {
	contract, err := bindClaimableCrowdsale(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &ClaimableCrowdsaleFilterer{contract: contract}, nil
}

// bindClaimableCrowdsale binds a generic wrapper to an already deployed contract.
func bindClaimableCrowdsale(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(ClaimableCrowdsaleABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
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

// ClaimableCrowdsalePauseIterator is returned from FilterPause and is used to iterate over the raw logs and unpacked data for Pause events raised by the ClaimableCrowdsale contract.
type ClaimableCrowdsalePauseIterator struct {
	Event *ClaimableCrowdsalePause // Event containing the contract specifics and raw log

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
func (it *ClaimableCrowdsalePauseIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ClaimableCrowdsalePause)
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
		it.Event = new(ClaimableCrowdsalePause)
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
func (it *ClaimableCrowdsalePauseIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ClaimableCrowdsalePauseIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ClaimableCrowdsalePause represents a Pause event raised by the ClaimableCrowdsale contract.
type ClaimableCrowdsalePause struct {
	Raw types.Log // Blockchain specific contextual infos
}

// FilterPause is a free log retrieval operation binding the contract event 0x6985a02210a168e66602d3235cb6db0e70f92b3ba4d376a33c0f3d9434bff625.
//
// Solidity: event Pause()
func (_ClaimableCrowdsale *ClaimableCrowdsaleFilterer) FilterPause(opts *bind.FilterOpts) (*ClaimableCrowdsalePauseIterator, error) {

	logs, sub, err := _ClaimableCrowdsale.contract.FilterLogs(opts, "Pause")
	if err != nil {
		return nil, err
	}
	return &ClaimableCrowdsalePauseIterator{contract: _ClaimableCrowdsale.contract, event: "Pause", logs: logs, sub: sub}, nil
}

// WatchPause is a free log subscription operation binding the contract event 0x6985a02210a168e66602d3235cb6db0e70f92b3ba4d376a33c0f3d9434bff625.
//
// Solidity: event Pause()
func (_ClaimableCrowdsale *ClaimableCrowdsaleFilterer) WatchPause(opts *bind.WatchOpts, sink chan<- *ClaimableCrowdsalePause) (event.Subscription, error) {

	logs, sub, err := _ClaimableCrowdsale.contract.WatchLogs(opts, "Pause")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ClaimableCrowdsalePause)
				if err := _ClaimableCrowdsale.contract.UnpackLog(event, "Pause", log); err != nil {
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

// ClaimableCrowdsaleTokenAddedIterator is returned from FilterTokenAdded and is used to iterate over the raw logs and unpacked data for TokenAdded events raised by the ClaimableCrowdsale contract.
type ClaimableCrowdsaleTokenAddedIterator struct {
	Event *ClaimableCrowdsaleTokenAdded // Event containing the contract specifics and raw log

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
func (it *ClaimableCrowdsaleTokenAddedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ClaimableCrowdsaleTokenAdded)
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
		it.Event = new(ClaimableCrowdsaleTokenAdded)
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
func (it *ClaimableCrowdsaleTokenAddedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ClaimableCrowdsaleTokenAddedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ClaimableCrowdsaleTokenAdded represents a TokenAdded event raised by the ClaimableCrowdsale contract.
type ClaimableCrowdsaleTokenAdded struct {
	Receiver         common.Address
	Tokens           *big.Int
	EquivalentAmount *big.Int
	Raw              types.Log // Blockchain specific contextual infos
}

// FilterTokenAdded is a free log retrieval operation binding the contract event 0xa818a22273fc309f0a3682b642c74c5b5c25c0615ff378d07767cd231e19fffc.
//
// Solidity: event TokenAdded(_receiver indexed address, _tokens uint256, _equivalentAmount uint256)
func (_ClaimableCrowdsale *ClaimableCrowdsaleFilterer) FilterTokenAdded(opts *bind.FilterOpts, _receiver []common.Address) (*ClaimableCrowdsaleTokenAddedIterator, error) {

	var _receiverRule []interface{}
	for _, _receiverItem := range _receiver {
		_receiverRule = append(_receiverRule, _receiverItem)
	}

	logs, sub, err := _ClaimableCrowdsale.contract.FilterLogs(opts, "TokenAdded", _receiverRule)
	if err != nil {
		return nil, err
	}
	return &ClaimableCrowdsaleTokenAddedIterator{contract: _ClaimableCrowdsale.contract, event: "TokenAdded", logs: logs, sub: sub}, nil
}

// WatchTokenAdded is a free log subscription operation binding the contract event 0xa818a22273fc309f0a3682b642c74c5b5c25c0615ff378d07767cd231e19fffc.
//
// Solidity: event TokenAdded(_receiver indexed address, _tokens uint256, _equivalentAmount uint256)
func (_ClaimableCrowdsale *ClaimableCrowdsaleFilterer) WatchTokenAdded(opts *bind.WatchOpts, sink chan<- *ClaimableCrowdsaleTokenAdded, _receiver []common.Address) (event.Subscription, error) {

	var _receiverRule []interface{}
	for _, _receiverItem := range _receiver {
		_receiverRule = append(_receiverRule, _receiverItem)
	}

	logs, sub, err := _ClaimableCrowdsale.contract.WatchLogs(opts, "TokenAdded", _receiverRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ClaimableCrowdsaleTokenAdded)
				if err := _ClaimableCrowdsale.contract.UnpackLog(event, "TokenAdded", log); err != nil {
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

// ClaimableCrowdsaleTokenBoughtIterator is returned from FilterTokenBought and is used to iterate over the raw logs and unpacked data for TokenBought events raised by the ClaimableCrowdsale contract.
type ClaimableCrowdsaleTokenBoughtIterator struct {
	Event *ClaimableCrowdsaleTokenBought // Event containing the contract specifics and raw log

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
func (it *ClaimableCrowdsaleTokenBoughtIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ClaimableCrowdsaleTokenBought)
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
		it.Event = new(ClaimableCrowdsaleTokenBought)
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
func (it *ClaimableCrowdsaleTokenBoughtIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ClaimableCrowdsaleTokenBoughtIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ClaimableCrowdsaleTokenBought represents a TokenBought event raised by the ClaimableCrowdsale contract.
type ClaimableCrowdsaleTokenBought struct {
	Buyer  common.Address
	Tokens *big.Int
	Amount *big.Int
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterTokenBought is a free log retrieval operation binding the contract event 0x28cab0d660ed8aedd61a8c9db00b97f6a2d67e07d87795994f440b18bc5f1aa3.
//
// Solidity: event TokenBought(_buyer indexed address, _tokens uint256, _amount uint256)
func (_ClaimableCrowdsale *ClaimableCrowdsaleFilterer) FilterTokenBought(opts *bind.FilterOpts, _buyer []common.Address) (*ClaimableCrowdsaleTokenBoughtIterator, error) {

	var _buyerRule []interface{}
	for _, _buyerItem := range _buyer {
		_buyerRule = append(_buyerRule, _buyerItem)
	}

	logs, sub, err := _ClaimableCrowdsale.contract.FilterLogs(opts, "TokenBought", _buyerRule)
	if err != nil {
		return nil, err
	}
	return &ClaimableCrowdsaleTokenBoughtIterator{contract: _ClaimableCrowdsale.contract, event: "TokenBought", logs: logs, sub: sub}, nil
}

// WatchTokenBought is a free log subscription operation binding the contract event 0x28cab0d660ed8aedd61a8c9db00b97f6a2d67e07d87795994f440b18bc5f1aa3.
//
// Solidity: event TokenBought(_buyer indexed address, _tokens uint256, _amount uint256)
func (_ClaimableCrowdsale *ClaimableCrowdsaleFilterer) WatchTokenBought(opts *bind.WatchOpts, sink chan<- *ClaimableCrowdsaleTokenBought, _buyer []common.Address) (event.Subscription, error) {

	var _buyerRule []interface{}
	for _, _buyerItem := range _buyer {
		_buyerRule = append(_buyerRule, _buyerItem)
	}

	logs, sub, err := _ClaimableCrowdsale.contract.WatchLogs(opts, "TokenBought", _buyerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ClaimableCrowdsaleTokenBought)
				if err := _ClaimableCrowdsale.contract.UnpackLog(event, "TokenBought", log); err != nil {
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

// ClaimableCrowdsaleTokenSentIterator is returned from FilterTokenSent and is used to iterate over the raw logs and unpacked data for TokenSent events raised by the ClaimableCrowdsale contract.
type ClaimableCrowdsaleTokenSentIterator struct {
	Event *ClaimableCrowdsaleTokenSent // Event containing the contract specifics and raw log

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
func (it *ClaimableCrowdsaleTokenSentIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ClaimableCrowdsaleTokenSent)
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
		it.Event = new(ClaimableCrowdsaleTokenSent)
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
func (it *ClaimableCrowdsaleTokenSentIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ClaimableCrowdsaleTokenSentIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ClaimableCrowdsaleTokenSent represents a TokenSent event raised by the ClaimableCrowdsale contract.
type ClaimableCrowdsaleTokenSent struct {
	Receiver common.Address
	Tokens   *big.Int
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterTokenSent is a free log retrieval operation binding the contract event 0xd4ff88d569a7ad2bfd6b17da9dbe82e2ccc016fd0051e08ff1ad1e6fe8fc9b55.
//
// Solidity: event TokenSent(_receiver indexed address, _tokens uint256)
func (_ClaimableCrowdsale *ClaimableCrowdsaleFilterer) FilterTokenSent(opts *bind.FilterOpts, _receiver []common.Address) (*ClaimableCrowdsaleTokenSentIterator, error) {

	var _receiverRule []interface{}
	for _, _receiverItem := range _receiver {
		_receiverRule = append(_receiverRule, _receiverItem)
	}

	logs, sub, err := _ClaimableCrowdsale.contract.FilterLogs(opts, "TokenSent", _receiverRule)
	if err != nil {
		return nil, err
	}
	return &ClaimableCrowdsaleTokenSentIterator{contract: _ClaimableCrowdsale.contract, event: "TokenSent", logs: logs, sub: sub}, nil
}

// WatchTokenSent is a free log subscription operation binding the contract event 0xd4ff88d569a7ad2bfd6b17da9dbe82e2ccc016fd0051e08ff1ad1e6fe8fc9b55.
//
// Solidity: event TokenSent(_receiver indexed address, _tokens uint256)
func (_ClaimableCrowdsale *ClaimableCrowdsaleFilterer) WatchTokenSent(opts *bind.WatchOpts, sink chan<- *ClaimableCrowdsaleTokenSent, _receiver []common.Address) (event.Subscription, error) {

	var _receiverRule []interface{}
	for _, _receiverItem := range _receiver {
		_receiverRule = append(_receiverRule, _receiverItem)
	}

	logs, sub, err := _ClaimableCrowdsale.contract.WatchLogs(opts, "TokenSent", _receiverRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ClaimableCrowdsaleTokenSent)
				if err := _ClaimableCrowdsale.contract.UnpackLog(event, "TokenSent", log); err != nil {
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

// ClaimableCrowdsaleTokenSubtractedIterator is returned from FilterTokenSubtracted and is used to iterate over the raw logs and unpacked data for TokenSubtracted events raised by the ClaimableCrowdsale contract.
type ClaimableCrowdsaleTokenSubtractedIterator struct {
	Event *ClaimableCrowdsaleTokenSubtracted // Event containing the contract specifics and raw log

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
func (it *ClaimableCrowdsaleTokenSubtractedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ClaimableCrowdsaleTokenSubtracted)
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
		it.Event = new(ClaimableCrowdsaleTokenSubtracted)
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
func (it *ClaimableCrowdsaleTokenSubtractedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ClaimableCrowdsaleTokenSubtractedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ClaimableCrowdsaleTokenSubtracted represents a TokenSubtracted event raised by the ClaimableCrowdsale contract.
type ClaimableCrowdsaleTokenSubtracted struct {
	Receiver         common.Address
	Tokens           *big.Int
	EquivalentAmount *big.Int
	Raw              types.Log // Blockchain specific contextual infos
}

// FilterTokenSubtracted is a free log retrieval operation binding the contract event 0x2da8b2becb27b189a54d21689aa110c877379a94c6f59f115cedc91d66239b33.
//
// Solidity: event TokenSubtracted(_receiver indexed address, _tokens uint256, _equivalentAmount uint256)
func (_ClaimableCrowdsale *ClaimableCrowdsaleFilterer) FilterTokenSubtracted(opts *bind.FilterOpts, _receiver []common.Address) (*ClaimableCrowdsaleTokenSubtractedIterator, error) {

	var _receiverRule []interface{}
	for _, _receiverItem := range _receiver {
		_receiverRule = append(_receiverRule, _receiverItem)
	}

	logs, sub, err := _ClaimableCrowdsale.contract.FilterLogs(opts, "TokenSubtracted", _receiverRule)
	if err != nil {
		return nil, err
	}
	return &ClaimableCrowdsaleTokenSubtractedIterator{contract: _ClaimableCrowdsale.contract, event: "TokenSubtracted", logs: logs, sub: sub}, nil
}

// WatchTokenSubtracted is a free log subscription operation binding the contract event 0x2da8b2becb27b189a54d21689aa110c877379a94c6f59f115cedc91d66239b33.
//
// Solidity: event TokenSubtracted(_receiver indexed address, _tokens uint256, _equivalentAmount uint256)
func (_ClaimableCrowdsale *ClaimableCrowdsaleFilterer) WatchTokenSubtracted(opts *bind.WatchOpts, sink chan<- *ClaimableCrowdsaleTokenSubtracted, _receiver []common.Address) (event.Subscription, error) {

	var _receiverRule []interface{}
	for _, _receiverItem := range _receiver {
		_receiverRule = append(_receiverRule, _receiverItem)
	}

	logs, sub, err := _ClaimableCrowdsale.contract.WatchLogs(opts, "TokenSubtracted", _receiverRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ClaimableCrowdsaleTokenSubtracted)
				if err := _ClaimableCrowdsale.contract.UnpackLog(event, "TokenSubtracted", log); err != nil {
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

// ClaimableCrowdsaleTokenToppedUpIterator is returned from FilterTokenToppedUp and is used to iterate over the raw logs and unpacked data for TokenToppedUp events raised by the ClaimableCrowdsale contract.
type ClaimableCrowdsaleTokenToppedUpIterator struct {
	Event *ClaimableCrowdsaleTokenToppedUp // Event containing the contract specifics and raw log

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
func (it *ClaimableCrowdsaleTokenToppedUpIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ClaimableCrowdsaleTokenToppedUp)
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
		it.Event = new(ClaimableCrowdsaleTokenToppedUp)
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
func (it *ClaimableCrowdsaleTokenToppedUpIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ClaimableCrowdsaleTokenToppedUpIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ClaimableCrowdsaleTokenToppedUp represents a TokenToppedUp event raised by the ClaimableCrowdsale contract.
type ClaimableCrowdsaleTokenToppedUp struct {
	Receiver         common.Address
	Tokens           *big.Int
	EquivalentAmount *big.Int
	Raw              types.Log // Blockchain specific contextual infos
}

// FilterTokenToppedUp is a free log retrieval operation binding the contract event 0xf10143e1b37ebaf3698e98976d8bddabcea045adc939c03b245da18e81ea9651.
//
// Solidity: event TokenToppedUp(_receiver indexed address, _tokens uint256, _equivalentAmount uint256)
func (_ClaimableCrowdsale *ClaimableCrowdsaleFilterer) FilterTokenToppedUp(opts *bind.FilterOpts, _receiver []common.Address) (*ClaimableCrowdsaleTokenToppedUpIterator, error) {

	var _receiverRule []interface{}
	for _, _receiverItem := range _receiver {
		_receiverRule = append(_receiverRule, _receiverItem)
	}

	logs, sub, err := _ClaimableCrowdsale.contract.FilterLogs(opts, "TokenToppedUp", _receiverRule)
	if err != nil {
		return nil, err
	}
	return &ClaimableCrowdsaleTokenToppedUpIterator{contract: _ClaimableCrowdsale.contract, event: "TokenToppedUp", logs: logs, sub: sub}, nil
}

// WatchTokenToppedUp is a free log subscription operation binding the contract event 0xf10143e1b37ebaf3698e98976d8bddabcea045adc939c03b245da18e81ea9651.
//
// Solidity: event TokenToppedUp(_receiver indexed address, _tokens uint256, _equivalentAmount uint256)
func (_ClaimableCrowdsale *ClaimableCrowdsaleFilterer) WatchTokenToppedUp(opts *bind.WatchOpts, sink chan<- *ClaimableCrowdsaleTokenToppedUp, _receiver []common.Address) (event.Subscription, error) {

	var _receiverRule []interface{}
	for _, _receiverItem := range _receiver {
		_receiverRule = append(_receiverRule, _receiverItem)
	}

	logs, sub, err := _ClaimableCrowdsale.contract.WatchLogs(opts, "TokenToppedUp", _receiverRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ClaimableCrowdsaleTokenToppedUp)
				if err := _ClaimableCrowdsale.contract.UnpackLog(event, "TokenToppedUp", log); err != nil {
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

// ClaimableCrowdsaleUnpauseIterator is returned from FilterUnpause and is used to iterate over the raw logs and unpacked data for Unpause events raised by the ClaimableCrowdsale contract.
type ClaimableCrowdsaleUnpauseIterator struct {
	Event *ClaimableCrowdsaleUnpause // Event containing the contract specifics and raw log

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
func (it *ClaimableCrowdsaleUnpauseIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ClaimableCrowdsaleUnpause)
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
		it.Event = new(ClaimableCrowdsaleUnpause)
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
func (it *ClaimableCrowdsaleUnpauseIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ClaimableCrowdsaleUnpauseIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ClaimableCrowdsaleUnpause represents a Unpause event raised by the ClaimableCrowdsale contract.
type ClaimableCrowdsaleUnpause struct {
	Raw types.Log // Blockchain specific contextual infos
}

// FilterUnpause is a free log retrieval operation binding the contract event 0x7805862f689e2f13df9f062ff482ad3ad112aca9e0847911ed832e158c525b33.
//
// Solidity: event Unpause()
func (_ClaimableCrowdsale *ClaimableCrowdsaleFilterer) FilterUnpause(opts *bind.FilterOpts) (*ClaimableCrowdsaleUnpauseIterator, error) {

	logs, sub, err := _ClaimableCrowdsale.contract.FilterLogs(opts, "Unpause")
	if err != nil {
		return nil, err
	}
	return &ClaimableCrowdsaleUnpauseIterator{contract: _ClaimableCrowdsale.contract, event: "Unpause", logs: logs, sub: sub}, nil
}

// WatchUnpause is a free log subscription operation binding the contract event 0x7805862f689e2f13df9f062ff482ad3ad112aca9e0847911ed832e158c525b33.
//
// Solidity: event Unpause()
func (_ClaimableCrowdsale *ClaimableCrowdsaleFilterer) WatchUnpause(opts *bind.WatchOpts, sink chan<- *ClaimableCrowdsaleUnpause) (event.Subscription, error) {

	logs, sub, err := _ClaimableCrowdsale.contract.WatchLogs(opts, "Unpause")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ClaimableCrowdsaleUnpause)
				if err := _ClaimableCrowdsale.contract.UnpackLog(event, "Unpause", log); err != nil {
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
