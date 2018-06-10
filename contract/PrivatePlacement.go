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

// PrivatePlacementABI is the input ABI used to generate the binding from.
const PrivatePlacementABI = "[{\"constant\":true,\"inputs\":[],\"name\":\"maxTokensAmount\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"supportAddress\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[],\"name\":\"allocateInternalWallets\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"totalSupply\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"minBuyableAmount\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[],\"name\":\"pauseToken\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"tokenRate\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[],\"name\":\"unpause\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"initialFoundersAmount\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[],\"name\":\"unpauseToken\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"paused\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_tokenRate\",\"type\":\"uint256\"}],\"name\":\"setTokenRate\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"bankAddress\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"isFinished\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[],\"name\":\"pause\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"initialSupportAmount\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"bountyAddress\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[],\"name\":\"buyTokens\",\"outputs\":[],\"payable\":true,\"stateMutability\":\"payable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_minBuyableAmount\",\"type\":\"uint256\"}],\"name\":\"setMinBuyableAmount\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[],\"name\":\"finish\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"issuedTokensAmount\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"foundersAddress\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"transferOwnership\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"initialBountyAmount\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"token\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"name\":\"_bankAddress\",\"type\":\"address\"},{\"name\":\"_foundersAddress\",\"type\":\"address\"},{\"name\":\"_supportAddress\",\"type\":\"address\"},{\"name\":\"_bountyAddress\",\"type\":\"address\"},{\"name\":\"_beneficiaryAddress\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"payable\":true,\"stateMutability\":\"payable\",\"type\":\"fallback\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"buyer\",\"type\":\"address\"},{\"indexed\":false,\"name\":\"tokens\",\"type\":\"uint256\"},{\"indexed\":false,\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"TokenBought\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[],\"name\":\"Pause\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[],\"name\":\"Unpause\",\"type\":\"event\"}]"

// PrivatePlacementBin is the compiled bytecode used for deploying new contracts.
const PrivatePlacementBin = `0x60606040526000805460a060020a60ff021990811682556005919091556009805460ff19169055600b8054909116905534156200003b57600080fd5b60405160a080620021da833981016040528080519190602001805191906020018051919060200180519190602001805191506200009490506b02df455b08e99ea2b300000064010000000062000d666200019a82021704565b60008054600160a060020a03338116600160a060020a0319928316179092556003805484841690831617905560028054838a1690831617905560018054928516929091169190911790556127106007819055869083906064630160a5b06359d030006200011883670de0b6b3a764000064010000000062000d03620001ca82021704565b6006556200013d82670de0b6b3a764000064010000000062000d03620001ca82021704565b60045560085550506009805461010060a860020a031916610100600160a060020a039a8b16021790555050600a8054600160a060020a03199081169688169690961790555050600b80549093169190931617905550620002099050565b600081620001a7620001f8565b908152602001604051809103906000f0801515620001c457600080fd5b92915050565b6000828202831580620001e85750828482811515620001e557fe5b04145b1515620001f157fe5b9392505050565b6040516108f980620018e183390190565b6116c880620002196000396000f30060606040526004361061013a5763ffffffff60e060020a6000350416630a38466581146101445780630f6d37d6146101695780631135b3ac1461019857806318160ddd146101ab5780632511b182146101be5780632c349627146101d157806331711884146101f85780633f4ba83a1461020b5780634d39ed061461021e57806350669a03146102315780635c975abb1461024457806361241c28146102575780637822ed491461026d5780637b352962146102805780638456cb59146102935780638da5cb5b146102a6578063ab28c704146102b9578063c516358f146102cc578063d0febe4c1461013a578063d561ef2b146102df578063d56b2889146102f5578063e69b414b14610308578063ebfdc6571461031b578063f2fde38b1461032e578063f338c9841461034d578063fc0c546a14610360575b610142610373565b005b341561014f57600080fd5b610157610584565b60405190815260200160405180910390f35b341561017457600080fd5b61017c61058a565b604051600160a060020a03909116815260200160405180910390f35b34156101a357600080fd5b610142610599565b34156101b657600080fd5b610157610773565b34156101c957600080fd5b610157610783565b34156101dc57600080fd5b6101e4610789565b604051901515815260200160405180910390f35b341561020357600080fd5b610157610862565b341561021657600080fd5b6101e4610868565b341561022957600080fd5b6101576108ec565b341561023c57600080fd5b6101e46108fb565b341561024f57600080fd5b6101e46109b7565b341561026257600080fd5b6101426004356109c7565b341561027857600080fd5b61017c610a23565b341561028b57600080fd5b6101e4610a32565b341561029e57600080fd5b6101e4610a3b565b34156102b157600080fd5b61017c610ac4565b34156102c457600080fd5b610157610ad3565b34156102d757600080fd5b61017c610ae2565b34156102ea57600080fd5b610142600435610af1565b341561030057600080fd5b610142610b66565b341561031357600080fd5b610157610c75565b341561032657600080fd5b61017c610c7b565b341561033957600080fd5b610142600160a060020a0360043516610c8f565b341561035857600080fd5b610157610ce5565b341561036b57600080fd5b61017c610cf4565b6009546000908190819060ff161561038a57600080fd5b6004546005541061039a57600080fd5b6008544211156103a957600080fd5b60005460a060020a900460ff16156103c057600080fd5b6006543410156103cf57600080fd5b600754349350600092506103e9908463ffffffff610d0316565b90506004548160055401111561043a5760055460045461040e9163ffffffff610d2e16565b905061042560075482610d4090919063ffffffff16565b9250610437348463ffffffff610d2e16565b91505b60055461044d908263ffffffff610d5716565b600581905560045490111561046157600080fd5b600354600160a060020a031663a9059cbb338360405160e060020a63ffffffff8516028152600160a060020a0390921660048301526024820152604401602060405180830381600087803b15156104b757600080fd5b5af115156104c457600080fd5b50505060405180515050600160a060020a0333167f28cab0d660ed8aedd61a8c9db00b97f6a2d67e07d87795994f440b18bc5f1aa3828560405191825260208201526040908101905180910390a2600154600160a060020a031683156108fc0284604051600060405180830381858888f19350505050151561054557600080fd5b600082111561057f57600160a060020a03331682156108fc0283604051600060405180830381858888f19350505050151561057f57600080fd5b505050565b60045481565b600a54600160a060020a031681565b60005433600160a060020a039081169116146105b457600080fd5b600b5460a060020a900460ff16156105cb57600080fd5b600b805474ff0000000000000000000000000000000000000000191660a060020a179055600354600954600160a060020a039182169163a9059cbb916101009004166adc94ce82ac7c640280000060405160e060020a63ffffffff8516028152600160a060020a0390921660048301526024820152604401602060405180830381600087803b151561065c57600080fd5b5af1151561066957600080fd5b50505060405180515050600354600a54600160a060020a039182169163a9059cbb91166a075a4b267d3758aac0000060405160e060020a63ffffffff8516028152600160a060020a0390921660048301526024820152604401602060405180830381600087803b15156106db57600080fd5b5af115156106e857600080fd5b50505060405180515050600354600b54600160a060020a039182169163a9059cbb91166a1d692c99f4dd62ab00000060405160e060020a63ffffffff8516028152600160a060020a0390921660048301526024820152604401602060405180830381600087803b151561075a57600080fd5b5af1151561076757600080fd5b50505060405180515050565b6b02df455b08e99ea2b300000081565b60065481565b6000805433600160a060020a039081169116146107a557600080fd5b600354600160a060020a0316635c975abb6040518163ffffffff1660e060020a028152600401602060405180830381600087803b15156107e457600080fd5b5af115156107f157600080fd5b505050604051805115905061080557600080fd5b600354600160a060020a0316638456cb596040518163ffffffff1660e060020a028152600401602060405180830381600087803b151561084457600080fd5b5af1151561085157600080fd5b505050604051805150600191505090565b60075481565b6000805433600160a060020a0390811691161461088457600080fd5b60005460a060020a900460ff16151561089c57600080fd5b6000805474ff0000000000000000000000000000000000000000191690557f7805862f689e2f13df9f062ff482ad3ad112aca9e0847911ed832e158c525b3360405160405180910390a150600190565b6adc94ce82ac7c640280000081565b6000805433600160a060020a0390811691161461091757600080fd5b600354600160a060020a0316635c975abb6040518163ffffffff1660e060020a028152600401602060405180830381600087803b151561095657600080fd5b5af1151561096357600080fd5b50505060405180519050151561097857600080fd5b600354600160a060020a0316633f4ba83a6040518163ffffffff1660e060020a028152600401602060405180830381600087803b151561084457600080fd5b60005460a060020a900460ff1681565b60005433600160a060020a039081169116146109e257600080fd5b60095460ff16156109f257600080fd5b60045460055410610a0257600080fd5b600854421115610a1157600080fd5b60008111610a1e57600080fd5b600755565b600254600160a060020a031681565b60095460ff1681565b6000805433600160a060020a03908116911614610a5757600080fd5b60005460a060020a900460ff1615610a6e57600080fd5b6000805474ff0000000000000000000000000000000000000000191660a060020a1790557f6985a02210a168e66602d3235cb6db0e70f92b3ba4d376a33c0f3d9434bff62560405160405180910390a150600190565b600054600160a060020a031681565b6a075a4b267d3758aac0000081565b600b54600160a060020a031681565b60005433600160a060020a03908116911614610b0c57600080fd5b60095460ff1615610b1c57600080fd5b60045460055410610b2c57600080fd5b600854421115610b3b57600080fd5b60008111610b4857600080fd5b610b6081670de0b6b3a764000063ffffffff610d0316565b60065550565b60005433600160a060020a03908116911614610b8157600080fd5b600454600554101580610b95575060085442115b1515610ba057600080fd5b60095460ff1615610bb057600080fd5b6009805460ff19166001179055600354600254600160a060020a039182169163a9059cbb9116826370a082313060405160e060020a63ffffffff8416028152600160a060020a039091166004820152602401602060405180830381600087803b1515610c1b57600080fd5b5af11515610c2857600080fd5b5050506040518051905060405160e060020a63ffffffff8516028152600160a060020a0390921660048301526024820152604401602060405180830381600087803b151561075a57600080fd5b60055481565b6009546101009004600160a060020a031681565b60005433600160a060020a03908116911614610caa57600080fd5b600160a060020a03811615610ce2576000805473ffffffffffffffffffffffffffffffffffffffff1916600160a060020a0383161790555b50565b6a1d692c99f4dd62ab00000081565b600354600160a060020a031681565b6000828202831580610d1f5750828482811515610d1c57fe5b04145b1515610d2757fe5b9392505050565b600082821115610d3a57fe5b50900390565b6000808284811515610d4e57fe5b04949350505050565b600082820183811015610d2757fe5b600081610d71610d93565b908152602001604051809103906000f0801515610d8d57600080fd5b92915050565b6040516108f980610da483390190560060606040526003805460a060020a60ff0219169055341561001f57600080fd5b6040516020806108f98339810160405280805160038054600160a060020a03191633600160a060020a03161790559150506000811161005d57600080fd5b600160a060020a033316600090815260016020526040812082905555610871806100886000396000f3006060604052600436106100cf5763ffffffff7c010000000000000000000000000000000000000000000000000000000060003504166306fdde0381146100d4578063095ea7b31461015e57806318160ddd1461019457806323b872dd146101b9578063313ce567146101e15780633f4ba83a1461020d5780635c975abb1461022057806370a08231146102335780638456cb59146102525780638da5cb5b1461026557806395d89b4114610294578063a9059cbb146102a7578063dd62ed3e146102c9578063f2fde38b146102ee575b600080fd5b34156100df57600080fd5b6100e761030f565b60405160208082528190810183818151815260200191508051906020019080838360005b8381101561012357808201518382015260200161010b565b50505050905090810190601f1680156101505780820380516001836020036101000a031916815260200191505b509250505060405180910390f35b341561016957600080fd5b610180600160a060020a0360043516602435610346565b604051901515815260200160405180910390f35b341561019f57600080fd5b6101a76103ec565b60405190815260200160405180910390f35b34156101c457600080fd5b610180600160a060020a03600435811690602435166044356103f2565b34156101ec57600080fd5b6101f461041f565b60405163ffffffff909116815260200160405180910390f35b341561021857600080fd5b610180610424565b341561022b57600080fd5b6101806104aa565b341561023e57600080fd5b6101a7600160a060020a03600435166104ba565b341561025d57600080fd5b6101806104d5565b341561027057600080fd5b610278610560565b604051600160a060020a03909116815260200160405180910390f35b341561029f57600080fd5b6100e761056f565b34156102b257600080fd5b610180600160a060020a03600435166024356105a6565b34156102d457600080fd5b6101a7600160a060020a03600435811690602435166105d1565b34156102f957600080fd5b61030d600160a060020a03600435166105fc565b005b60408051908101604052600a81527f484f515520546f6b656e00000000000000000000000000000000000000000000602082015281565b60008115806103785750600160a060020a03338116600090815260026020908152604080832093871683529290522054155b151561038357600080fd5b600160a060020a03338116600081815260026020908152604080832094881680845294909152908190208590557f8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b9259085905190815260200160405180910390a350600192915050565b60005481565b60035460009060a060020a900460ff161561040c57600080fd5b610417848484610652565b949350505050565b601281565b60035460009033600160a060020a0390811691161461044257600080fd5b60035460a060020a900460ff16151561045a57600080fd5b6003805474ff0000000000000000000000000000000000000000191690557f7805862f689e2f13df9f062ff482ad3ad112aca9e0847911ed832e158c525b3360405160405180910390a150600190565b60035460a060020a900460ff1681565b600160a060020a031660009081526001602052604090205490565b60035460009033600160a060020a039081169116146104f357600080fd5b60035460a060020a900460ff161561050a57600080fd5b6003805474ff0000000000000000000000000000000000000000191660a060020a1790557f6985a02210a168e66602d3235cb6db0e70f92b3ba4d376a33c0f3d9434bff62560405160405180910390a150600190565b600354600160a060020a031681565b60408051908101604052600381527f4851580000000000000000000000000000000000000000000000000000000000602082015281565b60035460009060a060020a900460ff16156105c057600080fd5b6105ca8383610765565b9392505050565b600160a060020a03918216600090815260026020908152604080832093909416825291909152205490565b60035433600160a060020a0390811691161461061757600080fd5b600160a060020a0381161561064f576003805473ffffffffffffffffffffffffffffffffffffffff1916600160a060020a0383161790555b50565b600160a060020a038084166000908152600260209081526040808320338516845282528083205493861683526001909152812054909190610699908463ffffffff61082416565b600160a060020a0380861660009081526001602052604080822093909355908716815220546106ce908463ffffffff61083316565b600160a060020a0386166000908152600160205260409020556106f7818463ffffffff61083316565b600160a060020a03808716600081815260026020908152604080832033861684529091529081902093909355908616917fddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef9086905190815260200160405180910390a3506001949350505050565b600160a060020a03331660009081526001602052604081205461078e908363ffffffff61083316565b600160a060020a0333811660009081526001602052604080822093909355908516815220546107c3908363ffffffff61082416565b600160a060020a0380851660008181526001602052604090819020939093559133909116907fddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef9085905190815260200160405180910390a350600192915050565b6000828201838110156105ca57fe5b60008282111561083f57fe5b509003905600a165627a7a723058209c386ceb1c89159603b5853635f586d3292243cdc50012c560431fec96dea9d30029a165627a7a7230582036f5924fd0ca859ae209913c054008f3f553d65b4f37f836de2a919a51d67c93002960606040526003805460a060020a60ff0219169055341561001f57600080fd5b6040516020806108f98339810160405280805160038054600160a060020a03191633600160a060020a03161790559150506000811161005d57600080fd5b600160a060020a033316600090815260016020526040812082905555610871806100886000396000f3006060604052600436106100cf5763ffffffff7c010000000000000000000000000000000000000000000000000000000060003504166306fdde0381146100d4578063095ea7b31461015e57806318160ddd1461019457806323b872dd146101b9578063313ce567146101e15780633f4ba83a1461020d5780635c975abb1461022057806370a08231146102335780638456cb59146102525780638da5cb5b1461026557806395d89b4114610294578063a9059cbb146102a7578063dd62ed3e146102c9578063f2fde38b146102ee575b600080fd5b34156100df57600080fd5b6100e761030f565b60405160208082528190810183818151815260200191508051906020019080838360005b8381101561012357808201518382015260200161010b565b50505050905090810190601f1680156101505780820380516001836020036101000a031916815260200191505b509250505060405180910390f35b341561016957600080fd5b610180600160a060020a0360043516602435610346565b604051901515815260200160405180910390f35b341561019f57600080fd5b6101a76103ec565b60405190815260200160405180910390f35b34156101c457600080fd5b610180600160a060020a03600435811690602435166044356103f2565b34156101ec57600080fd5b6101f461041f565b60405163ffffffff909116815260200160405180910390f35b341561021857600080fd5b610180610424565b341561022b57600080fd5b6101806104aa565b341561023e57600080fd5b6101a7600160a060020a03600435166104ba565b341561025d57600080fd5b6101806104d5565b341561027057600080fd5b610278610560565b604051600160a060020a03909116815260200160405180910390f35b341561029f57600080fd5b6100e761056f565b34156102b257600080fd5b610180600160a060020a03600435166024356105a6565b34156102d457600080fd5b6101a7600160a060020a03600435811690602435166105d1565b34156102f957600080fd5b61030d600160a060020a03600435166105fc565b005b60408051908101604052600a81527f484f515520546f6b656e00000000000000000000000000000000000000000000602082015281565b60008115806103785750600160a060020a03338116600090815260026020908152604080832093871683529290522054155b151561038357600080fd5b600160a060020a03338116600081815260026020908152604080832094881680845294909152908190208590557f8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b9259085905190815260200160405180910390a350600192915050565b60005481565b60035460009060a060020a900460ff161561040c57600080fd5b610417848484610652565b949350505050565b601281565b60035460009033600160a060020a0390811691161461044257600080fd5b60035460a060020a900460ff16151561045a57600080fd5b6003805474ff0000000000000000000000000000000000000000191690557f7805862f689e2f13df9f062ff482ad3ad112aca9e0847911ed832e158c525b3360405160405180910390a150600190565b60035460a060020a900460ff1681565b600160a060020a031660009081526001602052604090205490565b60035460009033600160a060020a039081169116146104f357600080fd5b60035460a060020a900460ff161561050a57600080fd5b6003805474ff0000000000000000000000000000000000000000191660a060020a1790557f6985a02210a168e66602d3235cb6db0e70f92b3ba4d376a33c0f3d9434bff62560405160405180910390a150600190565b600354600160a060020a031681565b60408051908101604052600381527f4851580000000000000000000000000000000000000000000000000000000000602082015281565b60035460009060a060020a900460ff16156105c057600080fd5b6105ca8383610765565b9392505050565b600160a060020a03918216600090815260026020908152604080832093909416825291909152205490565b60035433600160a060020a0390811691161461061757600080fd5b600160a060020a0381161561064f576003805473ffffffffffffffffffffffffffffffffffffffff1916600160a060020a0383161790555b50565b600160a060020a038084166000908152600260209081526040808320338516845282528083205493861683526001909152812054909190610699908463ffffffff61082416565b600160a060020a0380861660009081526001602052604080822093909355908716815220546106ce908463ffffffff61083316565b600160a060020a0386166000908152600160205260409020556106f7818463ffffffff61083316565b600160a060020a03808716600081815260026020908152604080832033861684529091529081902093909355908616917fddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef9086905190815260200160405180910390a3506001949350505050565b600160a060020a03331660009081526001602052604081205461078e908363ffffffff61083316565b600160a060020a0333811660009081526001602052604080822093909355908516815220546107c3908363ffffffff61082416565b600160a060020a0380851660008181526001602052604090819020939093559133909116907fddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef9085905190815260200160405180910390a350600192915050565b6000828201838110156105ca57fe5b60008282111561083f57fe5b509003905600a165627a7a723058209c386ceb1c89159603b5853635f586d3292243cdc50012c560431fec96dea9d30029`

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
	return address, tx, &PrivatePlacement{PrivatePlacementCaller: PrivatePlacementCaller{contract: contract}, PrivatePlacementTransactor: PrivatePlacementTransactor{contract: contract}, PrivatePlacementFilterer: PrivatePlacementFilterer{contract: contract}}, nil
}

// PrivatePlacement is an auto generated Go binding around an Ethereum contract.
type PrivatePlacement struct {
	PrivatePlacementCaller     // Read-only binding to the contract
	PrivatePlacementTransactor // Write-only binding to the contract
	PrivatePlacementFilterer   // Log filterer for contract events
}

// PrivatePlacementCaller is an auto generated read-only Go binding around an Ethereum contract.
type PrivatePlacementCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// PrivatePlacementTransactor is an auto generated write-only Go binding around an Ethereum contract.
type PrivatePlacementTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// PrivatePlacementFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type PrivatePlacementFilterer struct {
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
	contract, err := bindPrivatePlacement(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &PrivatePlacement{PrivatePlacementCaller: PrivatePlacementCaller{contract: contract}, PrivatePlacementTransactor: PrivatePlacementTransactor{contract: contract}, PrivatePlacementFilterer: PrivatePlacementFilterer{contract: contract}}, nil
}

// NewPrivatePlacementCaller creates a new read-only instance of PrivatePlacement, bound to a specific deployed contract.
func NewPrivatePlacementCaller(address common.Address, caller bind.ContractCaller) (*PrivatePlacementCaller, error) {
	contract, err := bindPrivatePlacement(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &PrivatePlacementCaller{contract: contract}, nil
}

// NewPrivatePlacementTransactor creates a new write-only instance of PrivatePlacement, bound to a specific deployed contract.
func NewPrivatePlacementTransactor(address common.Address, transactor bind.ContractTransactor) (*PrivatePlacementTransactor, error) {
	contract, err := bindPrivatePlacement(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &PrivatePlacementTransactor{contract: contract}, nil
}

// NewPrivatePlacementFilterer creates a new log filterer instance of PrivatePlacement, bound to a specific deployed contract.
func NewPrivatePlacementFilterer(address common.Address, filterer bind.ContractFilterer) (*PrivatePlacementFilterer, error) {
	contract, err := bindPrivatePlacement(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &PrivatePlacementFilterer{contract: contract}, nil
}

// bindPrivatePlacement binds a generic wrapper to an already deployed contract.
func bindPrivatePlacement(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(PrivatePlacementABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
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
// Solidity: function pause() returns(bool)
func (_PrivatePlacement *PrivatePlacementTransactor) Pause(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _PrivatePlacement.contract.Transact(opts, "pause")
}

// Pause is a paid mutator transaction binding the contract method 0x8456cb59.
//
// Solidity: function pause() returns(bool)
func (_PrivatePlacement *PrivatePlacementSession) Pause() (*types.Transaction, error) {
	return _PrivatePlacement.Contract.Pause(&_PrivatePlacement.TransactOpts)
}

// Pause is a paid mutator transaction binding the contract method 0x8456cb59.
//
// Solidity: function pause() returns(bool)
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
// Solidity: function unpause() returns(bool)
func (_PrivatePlacement *PrivatePlacementTransactor) Unpause(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _PrivatePlacement.contract.Transact(opts, "unpause")
}

// Unpause is a paid mutator transaction binding the contract method 0x3f4ba83a.
//
// Solidity: function unpause() returns(bool)
func (_PrivatePlacement *PrivatePlacementSession) Unpause() (*types.Transaction, error) {
	return _PrivatePlacement.Contract.Unpause(&_PrivatePlacement.TransactOpts)
}

// Unpause is a paid mutator transaction binding the contract method 0x3f4ba83a.
//
// Solidity: function unpause() returns(bool)
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

// PrivatePlacementPauseIterator is returned from FilterPause and is used to iterate over the raw logs and unpacked data for Pause events raised by the PrivatePlacement contract.
type PrivatePlacementPauseIterator struct {
	Event *PrivatePlacementPause // Event containing the contract specifics and raw log

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
func (it *PrivatePlacementPauseIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(PrivatePlacementPause)
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
		it.Event = new(PrivatePlacementPause)
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
func (it *PrivatePlacementPauseIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *PrivatePlacementPauseIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// PrivatePlacementPause represents a Pause event raised by the PrivatePlacement contract.
type PrivatePlacementPause struct {
	Raw types.Log // Blockchain specific contextual infos
}

// FilterPause is a free log retrieval operation binding the contract event 0x6985a02210a168e66602d3235cb6db0e70f92b3ba4d376a33c0f3d9434bff625.
//
// Solidity: event Pause()
func (_PrivatePlacement *PrivatePlacementFilterer) FilterPause(opts *bind.FilterOpts) (*PrivatePlacementPauseIterator, error) {

	logs, sub, err := _PrivatePlacement.contract.FilterLogs(opts, "Pause")
	if err != nil {
		return nil, err
	}
	return &PrivatePlacementPauseIterator{contract: _PrivatePlacement.contract, event: "Pause", logs: logs, sub: sub}, nil
}

// WatchPause is a free log subscription operation binding the contract event 0x6985a02210a168e66602d3235cb6db0e70f92b3ba4d376a33c0f3d9434bff625.
//
// Solidity: event Pause()
func (_PrivatePlacement *PrivatePlacementFilterer) WatchPause(opts *bind.WatchOpts, sink chan<- *PrivatePlacementPause) (event.Subscription, error) {

	logs, sub, err := _PrivatePlacement.contract.WatchLogs(opts, "Pause")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(PrivatePlacementPause)
				if err := _PrivatePlacement.contract.UnpackLog(event, "Pause", log); err != nil {
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

// PrivatePlacementTokenBoughtIterator is returned from FilterTokenBought and is used to iterate over the raw logs and unpacked data for TokenBought events raised by the PrivatePlacement contract.
type PrivatePlacementTokenBoughtIterator struct {
	Event *PrivatePlacementTokenBought // Event containing the contract specifics and raw log

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
func (it *PrivatePlacementTokenBoughtIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(PrivatePlacementTokenBought)
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
		it.Event = new(PrivatePlacementTokenBought)
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
func (it *PrivatePlacementTokenBoughtIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *PrivatePlacementTokenBoughtIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// PrivatePlacementTokenBought represents a TokenBought event raised by the PrivatePlacement contract.
type PrivatePlacementTokenBought struct {
	Buyer  common.Address
	Tokens *big.Int
	Amount *big.Int
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterTokenBought is a free log retrieval operation binding the contract event 0x28cab0d660ed8aedd61a8c9db00b97f6a2d67e07d87795994f440b18bc5f1aa3.
//
// Solidity: event TokenBought(buyer indexed address, tokens uint256, amount uint256)
func (_PrivatePlacement *PrivatePlacementFilterer) FilterTokenBought(opts *bind.FilterOpts, buyer []common.Address) (*PrivatePlacementTokenBoughtIterator, error) {

	var buyerRule []interface{}
	for _, buyerItem := range buyer {
		buyerRule = append(buyerRule, buyerItem)
	}

	logs, sub, err := _PrivatePlacement.contract.FilterLogs(opts, "TokenBought", buyerRule)
	if err != nil {
		return nil, err
	}
	return &PrivatePlacementTokenBoughtIterator{contract: _PrivatePlacement.contract, event: "TokenBought", logs: logs, sub: sub}, nil
}

// WatchTokenBought is a free log subscription operation binding the contract event 0x28cab0d660ed8aedd61a8c9db00b97f6a2d67e07d87795994f440b18bc5f1aa3.
//
// Solidity: event TokenBought(buyer indexed address, tokens uint256, amount uint256)
func (_PrivatePlacement *PrivatePlacementFilterer) WatchTokenBought(opts *bind.WatchOpts, sink chan<- *PrivatePlacementTokenBought, buyer []common.Address) (event.Subscription, error) {

	var buyerRule []interface{}
	for _, buyerItem := range buyer {
		buyerRule = append(buyerRule, buyerItem)
	}

	logs, sub, err := _PrivatePlacement.contract.WatchLogs(opts, "TokenBought", buyerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(PrivatePlacementTokenBought)
				if err := _PrivatePlacement.contract.UnpackLog(event, "TokenBought", log); err != nil {
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

// PrivatePlacementUnpauseIterator is returned from FilterUnpause and is used to iterate over the raw logs and unpacked data for Unpause events raised by the PrivatePlacement contract.
type PrivatePlacementUnpauseIterator struct {
	Event *PrivatePlacementUnpause // Event containing the contract specifics and raw log

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
func (it *PrivatePlacementUnpauseIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(PrivatePlacementUnpause)
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
		it.Event = new(PrivatePlacementUnpause)
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
func (it *PrivatePlacementUnpauseIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *PrivatePlacementUnpauseIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// PrivatePlacementUnpause represents a Unpause event raised by the PrivatePlacement contract.
type PrivatePlacementUnpause struct {
	Raw types.Log // Blockchain specific contextual infos
}

// FilterUnpause is a free log retrieval operation binding the contract event 0x7805862f689e2f13df9f062ff482ad3ad112aca9e0847911ed832e158c525b33.
//
// Solidity: event Unpause()
func (_PrivatePlacement *PrivatePlacementFilterer) FilterUnpause(opts *bind.FilterOpts) (*PrivatePlacementUnpauseIterator, error) {

	logs, sub, err := _PrivatePlacement.contract.FilterLogs(opts, "Unpause")
	if err != nil {
		return nil, err
	}
	return &PrivatePlacementUnpauseIterator{contract: _PrivatePlacement.contract, event: "Unpause", logs: logs, sub: sub}, nil
}

// WatchUnpause is a free log subscription operation binding the contract event 0x7805862f689e2f13df9f062ff482ad3ad112aca9e0847911ed832e158c525b33.
//
// Solidity: event Unpause()
func (_PrivatePlacement *PrivatePlacementFilterer) WatchUnpause(opts *bind.WatchOpts, sink chan<- *PrivatePlacementUnpause) (event.Subscription, error) {

	logs, sub, err := _PrivatePlacement.contract.WatchLogs(opts, "Unpause")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(PrivatePlacementUnpause)
				if err := _PrivatePlacement.contract.UnpackLog(event, "Unpause", log); err != nil {
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
