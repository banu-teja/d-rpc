// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package contracts

import (
	"errors"
	"math/big"
	"strings"

	ethereum "github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/event"
)

// Reference imports to suppress errors if they are not otherwise used.
var (
	_ = errors.New
	_ = big.NewInt
	_ = strings.NewReader
	_ = ethereum.NotFound
	_ = bind.Bind
	_ = common.Big1
	_ = types.BloomLookup
	_ = event.NewSubscription
	_ = abi.ConvertType
)

// ContractsMetaData contains all meta data concerning the Contracts contract.
var ContractsMetaData = &bind.MetaData{
	ABI: "[{\"type\":\"constructor\",\"inputs\":[{\"name\":\"_stakeToken\",\"type\":\"address\",\"internalType\":\"contractIERC20\"},{\"name\":\"_minStake\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"DEREGISTRATION_COOLDOWN\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"depositStake\",\"inputs\":[{\"name\":\"amount\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"deregister\",\"inputs\":[],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"minStake\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"owner\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"providers\",\"inputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[{\"name\":\"stake\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"qosScore\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"registered\",\"type\":\"bool\",\"internalType\":\"bool\"},{\"name\":\"registrationTime\",\"type\":\"uint40\",\"internalType\":\"uint40\"},{\"name\":\"deregistrationTime\",\"type\":\"uint40\",\"internalType\":\"uint40\"},{\"name\":\"endpointURL\",\"type\":\"string\",\"internalType\":\"string\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"register\",\"inputs\":[],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"renounceOwnership\",\"inputs\":[],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"setEndpointURL\",\"inputs\":[{\"name\":\"_url\",\"type\":\"string\",\"internalType\":\"string\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"setMinStake\",\"inputs\":[{\"name\":\"_minStake\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"slashProvider\",\"inputs\":[{\"name\":\"provider_\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"amount\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"stakeToken\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"contractIERC20\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"transferOwnership\",\"inputs\":[{\"name\":\"newOwner\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"updateQoS\",\"inputs\":[{\"name\":\"provider_\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"score\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"withdrawStake\",\"inputs\":[{\"name\":\"amount\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"event\",\"name\":\"OwnershipTransferred\",\"inputs\":[{\"name\":\"previousOwner\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"newOwner\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"ProviderDeregistered\",\"inputs\":[{\"name\":\"provider\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"ProviderRegistered\",\"inputs\":[{\"name\":\"provider\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"stake\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"},{\"name\":\"endpointURL\",\"type\":\"string\",\"indexed\":false,\"internalType\":\"string\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"ProviderSlashed\",\"inputs\":[{\"name\":\"provider\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"amount\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"ProviderURLUpdated\",\"inputs\":[{\"name\":\"provider\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"newURL\",\"type\":\"string\",\"indexed\":false,\"internalType\":\"string\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"QoSUpdated\",\"inputs\":[{\"name\":\"provider\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"newScore\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"StakeDeposited\",\"inputs\":[{\"name\":\"provider\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"amount\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"StakeWithdrawn\",\"inputs\":[{\"name\":\"provider\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"amount\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"}],\"anonymous\":false},{\"type\":\"error\",\"name\":\"OwnableInvalidOwner\",\"inputs\":[{\"name\":\"owner\",\"type\":\"address\",\"internalType\":\"address\"}]},{\"type\":\"error\",\"name\":\"OwnableUnauthorizedAccount\",\"inputs\":[{\"name\":\"account\",\"type\":\"address\",\"internalType\":\"address\"}]},{\"type\":\"error\",\"name\":\"ReentrancyGuardReentrantCall\",\"inputs\":[]}]",
}

// ContractsABI is the input ABI used to generate the binding from.
// Deprecated: Use ContractsMetaData.ABI instead.
var ContractsABI = ContractsMetaData.ABI

// Contracts is an auto generated Go binding around an Ethereum contract.
type Contracts struct {
	ContractsCaller     // Read-only binding to the contract
	ContractsTransactor // Write-only binding to the contract
	ContractsFilterer   // Log filterer for contract events
}

// ContractsCaller is an auto generated read-only Go binding around an Ethereum contract.
type ContractsCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ContractsTransactor is an auto generated write-only Go binding around an Ethereum contract.
type ContractsTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ContractsFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type ContractsFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ContractsSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type ContractsSession struct {
	Contract     *Contracts        // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// ContractsCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type ContractsCallerSession struct {
	Contract *ContractsCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts    // Call options to use throughout this session
}

// ContractsTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type ContractsTransactorSession struct {
	Contract     *ContractsTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts    // Transaction auth options to use throughout this session
}

// ContractsRaw is an auto generated low-level Go binding around an Ethereum contract.
type ContractsRaw struct {
	Contract *Contracts // Generic contract binding to access the raw methods on
}

// ContractsCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type ContractsCallerRaw struct {
	Contract *ContractsCaller // Generic read-only contract binding to access the raw methods on
}

// ContractsTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type ContractsTransactorRaw struct {
	Contract *ContractsTransactor // Generic write-only contract binding to access the raw methods on
}

// NewContracts creates a new instance of Contracts, bound to a specific deployed contract.
func NewContracts(address common.Address, backend bind.ContractBackend) (*Contracts, error) {
	contract, err := bindContracts(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Contracts{ContractsCaller: ContractsCaller{contract: contract}, ContractsTransactor: ContractsTransactor{contract: contract}, ContractsFilterer: ContractsFilterer{contract: contract}}, nil
}

// NewContractsCaller creates a new read-only instance of Contracts, bound to a specific deployed contract.
func NewContractsCaller(address common.Address, caller bind.ContractCaller) (*ContractsCaller, error) {
	contract, err := bindContracts(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &ContractsCaller{contract: contract}, nil
}

// NewContractsTransactor creates a new write-only instance of Contracts, bound to a specific deployed contract.
func NewContractsTransactor(address common.Address, transactor bind.ContractTransactor) (*ContractsTransactor, error) {
	contract, err := bindContracts(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &ContractsTransactor{contract: contract}, nil
}

// NewContractsFilterer creates a new log filterer instance of Contracts, bound to a specific deployed contract.
func NewContractsFilterer(address common.Address, filterer bind.ContractFilterer) (*ContractsFilterer, error) {
	contract, err := bindContracts(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &ContractsFilterer{contract: contract}, nil
}

// bindContracts binds a generic wrapper to an already deployed contract.
func bindContracts(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := ContractsMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Contracts *ContractsRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Contracts.Contract.ContractsCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Contracts *ContractsRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Contracts.Contract.ContractsTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Contracts *ContractsRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Contracts.Contract.ContractsTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Contracts *ContractsCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Contracts.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Contracts *ContractsTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Contracts.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Contracts *ContractsTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Contracts.Contract.contract.Transact(opts, method, params...)
}

// DEREGISTRATIONCOOLDOWN is a free data retrieval call binding the contract method 0x2de6211b.
//
// Solidity: function DEREGISTRATION_COOLDOWN() view returns(uint256)
func (_Contracts *ContractsCaller) DEREGISTRATIONCOOLDOWN(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Contracts.contract.Call(opts, &out, "DEREGISTRATION_COOLDOWN")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// DEREGISTRATIONCOOLDOWN is a free data retrieval call binding the contract method 0x2de6211b.
//
// Solidity: function DEREGISTRATION_COOLDOWN() view returns(uint256)
func (_Contracts *ContractsSession) DEREGISTRATIONCOOLDOWN() (*big.Int, error) {
	return _Contracts.Contract.DEREGISTRATIONCOOLDOWN(&_Contracts.CallOpts)
}

// DEREGISTRATIONCOOLDOWN is a free data retrieval call binding the contract method 0x2de6211b.
//
// Solidity: function DEREGISTRATION_COOLDOWN() view returns(uint256)
func (_Contracts *ContractsCallerSession) DEREGISTRATIONCOOLDOWN() (*big.Int, error) {
	return _Contracts.Contract.DEREGISTRATIONCOOLDOWN(&_Contracts.CallOpts)
}

// MinStake is a free data retrieval call binding the contract method 0x375b3c0a.
//
// Solidity: function minStake() view returns(uint256)
func (_Contracts *ContractsCaller) MinStake(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Contracts.contract.Call(opts, &out, "minStake")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// MinStake is a free data retrieval call binding the contract method 0x375b3c0a.
//
// Solidity: function minStake() view returns(uint256)
func (_Contracts *ContractsSession) MinStake() (*big.Int, error) {
	return _Contracts.Contract.MinStake(&_Contracts.CallOpts)
}

// MinStake is a free data retrieval call binding the contract method 0x375b3c0a.
//
// Solidity: function minStake() view returns(uint256)
func (_Contracts *ContractsCallerSession) MinStake() (*big.Int, error) {
	return _Contracts.Contract.MinStake(&_Contracts.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_Contracts *ContractsCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Contracts.contract.Call(opts, &out, "owner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_Contracts *ContractsSession) Owner() (common.Address, error) {
	return _Contracts.Contract.Owner(&_Contracts.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_Contracts *ContractsCallerSession) Owner() (common.Address, error) {
	return _Contracts.Contract.Owner(&_Contracts.CallOpts)
}

// Providers is a free data retrieval call binding the contract method 0x0787bc27.
//
// Solidity: function providers(address ) view returns(uint256 stake, uint256 qosScore, bool registered, uint40 registrationTime, uint40 deregistrationTime, string endpointURL)
func (_Contracts *ContractsCaller) Providers(opts *bind.CallOpts, arg0 common.Address) (struct {
	Stake              *big.Int
	QosScore           *big.Int
	Registered         bool
	RegistrationTime   *big.Int
	DeregistrationTime *big.Int
	EndpointURL        string
}, error) {
	var out []interface{}
	err := _Contracts.contract.Call(opts, &out, "providers", arg0)

	outstruct := new(struct {
		Stake              *big.Int
		QosScore           *big.Int
		Registered         bool
		RegistrationTime   *big.Int
		DeregistrationTime *big.Int
		EndpointURL        string
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.Stake = *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)
	outstruct.QosScore = *abi.ConvertType(out[1], new(*big.Int)).(**big.Int)
	outstruct.Registered = *abi.ConvertType(out[2], new(bool)).(*bool)
	outstruct.RegistrationTime = *abi.ConvertType(out[3], new(*big.Int)).(**big.Int)
	outstruct.DeregistrationTime = *abi.ConvertType(out[4], new(*big.Int)).(**big.Int)
	outstruct.EndpointURL = *abi.ConvertType(out[5], new(string)).(*string)

	return *outstruct, err

}

// Providers is a free data retrieval call binding the contract method 0x0787bc27.
//
// Solidity: function providers(address ) view returns(uint256 stake, uint256 qosScore, bool registered, uint40 registrationTime, uint40 deregistrationTime, string endpointURL)
func (_Contracts *ContractsSession) Providers(arg0 common.Address) (struct {
	Stake              *big.Int
	QosScore           *big.Int
	Registered         bool
	RegistrationTime   *big.Int
	DeregistrationTime *big.Int
	EndpointURL        string
}, error) {
	return _Contracts.Contract.Providers(&_Contracts.CallOpts, arg0)
}

// Providers is a free data retrieval call binding the contract method 0x0787bc27.
//
// Solidity: function providers(address ) view returns(uint256 stake, uint256 qosScore, bool registered, uint40 registrationTime, uint40 deregistrationTime, string endpointURL)
func (_Contracts *ContractsCallerSession) Providers(arg0 common.Address) (struct {
	Stake              *big.Int
	QosScore           *big.Int
	Registered         bool
	RegistrationTime   *big.Int
	DeregistrationTime *big.Int
	EndpointURL        string
}, error) {
	return _Contracts.Contract.Providers(&_Contracts.CallOpts, arg0)
}

// StakeToken is a free data retrieval call binding the contract method 0x51ed6a30.
//
// Solidity: function stakeToken() view returns(address)
func (_Contracts *ContractsCaller) StakeToken(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Contracts.contract.Call(opts, &out, "stakeToken")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// StakeToken is a free data retrieval call binding the contract method 0x51ed6a30.
//
// Solidity: function stakeToken() view returns(address)
func (_Contracts *ContractsSession) StakeToken() (common.Address, error) {
	return _Contracts.Contract.StakeToken(&_Contracts.CallOpts)
}

// StakeToken is a free data retrieval call binding the contract method 0x51ed6a30.
//
// Solidity: function stakeToken() view returns(address)
func (_Contracts *ContractsCallerSession) StakeToken() (common.Address, error) {
	return _Contracts.Contract.StakeToken(&_Contracts.CallOpts)
}

// DepositStake is a paid mutator transaction binding the contract method 0xcb82cc8f.
//
// Solidity: function depositStake(uint256 amount) returns()
func (_Contracts *ContractsTransactor) DepositStake(opts *bind.TransactOpts, amount *big.Int) (*types.Transaction, error) {
	return _Contracts.contract.Transact(opts, "depositStake", amount)
}

// DepositStake is a paid mutator transaction binding the contract method 0xcb82cc8f.
//
// Solidity: function depositStake(uint256 amount) returns()
func (_Contracts *ContractsSession) DepositStake(amount *big.Int) (*types.Transaction, error) {
	return _Contracts.Contract.DepositStake(&_Contracts.TransactOpts, amount)
}

// DepositStake is a paid mutator transaction binding the contract method 0xcb82cc8f.
//
// Solidity: function depositStake(uint256 amount) returns()
func (_Contracts *ContractsTransactorSession) DepositStake(amount *big.Int) (*types.Transaction, error) {
	return _Contracts.Contract.DepositStake(&_Contracts.TransactOpts, amount)
}

// Deregister is a paid mutator transaction binding the contract method 0xaff5edb1.
//
// Solidity: function deregister() returns()
func (_Contracts *ContractsTransactor) Deregister(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Contracts.contract.Transact(opts, "deregister")
}

// Deregister is a paid mutator transaction binding the contract method 0xaff5edb1.
//
// Solidity: function deregister() returns()
func (_Contracts *ContractsSession) Deregister() (*types.Transaction, error) {
	return _Contracts.Contract.Deregister(&_Contracts.TransactOpts)
}

// Deregister is a paid mutator transaction binding the contract method 0xaff5edb1.
//
// Solidity: function deregister() returns()
func (_Contracts *ContractsTransactorSession) Deregister() (*types.Transaction, error) {
	return _Contracts.Contract.Deregister(&_Contracts.TransactOpts)
}

// Register is a paid mutator transaction binding the contract method 0x1aa3a008.
//
// Solidity: function register() returns()
func (_Contracts *ContractsTransactor) Register(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Contracts.contract.Transact(opts, "register")
}

// Register is a paid mutator transaction binding the contract method 0x1aa3a008.
//
// Solidity: function register() returns()
func (_Contracts *ContractsSession) Register() (*types.Transaction, error) {
	return _Contracts.Contract.Register(&_Contracts.TransactOpts)
}

// Register is a paid mutator transaction binding the contract method 0x1aa3a008.
//
// Solidity: function register() returns()
func (_Contracts *ContractsTransactorSession) Register() (*types.Transaction, error) {
	return _Contracts.Contract.Register(&_Contracts.TransactOpts)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_Contracts *ContractsTransactor) RenounceOwnership(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Contracts.contract.Transact(opts, "renounceOwnership")
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_Contracts *ContractsSession) RenounceOwnership() (*types.Transaction, error) {
	return _Contracts.Contract.RenounceOwnership(&_Contracts.TransactOpts)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_Contracts *ContractsTransactorSession) RenounceOwnership() (*types.Transaction, error) {
	return _Contracts.Contract.RenounceOwnership(&_Contracts.TransactOpts)
}

// SetEndpointURL is a paid mutator transaction binding the contract method 0xf6c9b2d4.
//
// Solidity: function setEndpointURL(string _url) returns()
func (_Contracts *ContractsTransactor) SetEndpointURL(opts *bind.TransactOpts, _url string) (*types.Transaction, error) {
	return _Contracts.contract.Transact(opts, "setEndpointURL", _url)
}

// SetEndpointURL is a paid mutator transaction binding the contract method 0xf6c9b2d4.
//
// Solidity: function setEndpointURL(string _url) returns()
func (_Contracts *ContractsSession) SetEndpointURL(_url string) (*types.Transaction, error) {
	return _Contracts.Contract.SetEndpointURL(&_Contracts.TransactOpts, _url)
}

// SetEndpointURL is a paid mutator transaction binding the contract method 0xf6c9b2d4.
//
// Solidity: function setEndpointURL(string _url) returns()
func (_Contracts *ContractsTransactorSession) SetEndpointURL(_url string) (*types.Transaction, error) {
	return _Contracts.Contract.SetEndpointURL(&_Contracts.TransactOpts, _url)
}

// SetMinStake is a paid mutator transaction binding the contract method 0x8c80fd90.
//
// Solidity: function setMinStake(uint256 _minStake) returns()
func (_Contracts *ContractsTransactor) SetMinStake(opts *bind.TransactOpts, _minStake *big.Int) (*types.Transaction, error) {
	return _Contracts.contract.Transact(opts, "setMinStake", _minStake)
}

// SetMinStake is a paid mutator transaction binding the contract method 0x8c80fd90.
//
// Solidity: function setMinStake(uint256 _minStake) returns()
func (_Contracts *ContractsSession) SetMinStake(_minStake *big.Int) (*types.Transaction, error) {
	return _Contracts.Contract.SetMinStake(&_Contracts.TransactOpts, _minStake)
}

// SetMinStake is a paid mutator transaction binding the contract method 0x8c80fd90.
//
// Solidity: function setMinStake(uint256 _minStake) returns()
func (_Contracts *ContractsTransactorSession) SetMinStake(_minStake *big.Int) (*types.Transaction, error) {
	return _Contracts.Contract.SetMinStake(&_Contracts.TransactOpts, _minStake)
}

// SlashProvider is a paid mutator transaction binding the contract method 0x3febfc0f.
//
// Solidity: function slashProvider(address provider_, uint256 amount) returns()
func (_Contracts *ContractsTransactor) SlashProvider(opts *bind.TransactOpts, provider_ common.Address, amount *big.Int) (*types.Transaction, error) {
	return _Contracts.contract.Transact(opts, "slashProvider", provider_, amount)
}

// SlashProvider is a paid mutator transaction binding the contract method 0x3febfc0f.
//
// Solidity: function slashProvider(address provider_, uint256 amount) returns()
func (_Contracts *ContractsSession) SlashProvider(provider_ common.Address, amount *big.Int) (*types.Transaction, error) {
	return _Contracts.Contract.SlashProvider(&_Contracts.TransactOpts, provider_, amount)
}

// SlashProvider is a paid mutator transaction binding the contract method 0x3febfc0f.
//
// Solidity: function slashProvider(address provider_, uint256 amount) returns()
func (_Contracts *ContractsTransactorSession) SlashProvider(provider_ common.Address, amount *big.Int) (*types.Transaction, error) {
	return _Contracts.Contract.SlashProvider(&_Contracts.TransactOpts, provider_, amount)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_Contracts *ContractsTransactor) TransferOwnership(opts *bind.TransactOpts, newOwner common.Address) (*types.Transaction, error) {
	return _Contracts.contract.Transact(opts, "transferOwnership", newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_Contracts *ContractsSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _Contracts.Contract.TransferOwnership(&_Contracts.TransactOpts, newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_Contracts *ContractsTransactorSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _Contracts.Contract.TransferOwnership(&_Contracts.TransactOpts, newOwner)
}

// UpdateQoS is a paid mutator transaction binding the contract method 0x117835a4.
//
// Solidity: function updateQoS(address provider_, uint256 score) returns()
func (_Contracts *ContractsTransactor) UpdateQoS(opts *bind.TransactOpts, provider_ common.Address, score *big.Int) (*types.Transaction, error) {
	return _Contracts.contract.Transact(opts, "updateQoS", provider_, score)
}

// UpdateQoS is a paid mutator transaction binding the contract method 0x117835a4.
//
// Solidity: function updateQoS(address provider_, uint256 score) returns()
func (_Contracts *ContractsSession) UpdateQoS(provider_ common.Address, score *big.Int) (*types.Transaction, error) {
	return _Contracts.Contract.UpdateQoS(&_Contracts.TransactOpts, provider_, score)
}

// UpdateQoS is a paid mutator transaction binding the contract method 0x117835a4.
//
// Solidity: function updateQoS(address provider_, uint256 score) returns()
func (_Contracts *ContractsTransactorSession) UpdateQoS(provider_ common.Address, score *big.Int) (*types.Transaction, error) {
	return _Contracts.Contract.UpdateQoS(&_Contracts.TransactOpts, provider_, score)
}

// WithdrawStake is a paid mutator transaction binding the contract method 0x25d5971f.
//
// Solidity: function withdrawStake(uint256 amount) returns()
func (_Contracts *ContractsTransactor) WithdrawStake(opts *bind.TransactOpts, amount *big.Int) (*types.Transaction, error) {
	return _Contracts.contract.Transact(opts, "withdrawStake", amount)
}

// WithdrawStake is a paid mutator transaction binding the contract method 0x25d5971f.
//
// Solidity: function withdrawStake(uint256 amount) returns()
func (_Contracts *ContractsSession) WithdrawStake(amount *big.Int) (*types.Transaction, error) {
	return _Contracts.Contract.WithdrawStake(&_Contracts.TransactOpts, amount)
}

// WithdrawStake is a paid mutator transaction binding the contract method 0x25d5971f.
//
// Solidity: function withdrawStake(uint256 amount) returns()
func (_Contracts *ContractsTransactorSession) WithdrawStake(amount *big.Int) (*types.Transaction, error) {
	return _Contracts.Contract.WithdrawStake(&_Contracts.TransactOpts, amount)
}

// ContractsOwnershipTransferredIterator is returned from FilterOwnershipTransferred and is used to iterate over the raw logs and unpacked data for OwnershipTransferred events raised by the Contracts contract.
type ContractsOwnershipTransferredIterator struct {
	Event *ContractsOwnershipTransferred // Event containing the contract specifics and raw log

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
func (it *ContractsOwnershipTransferredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ContractsOwnershipTransferred)
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
		it.Event = new(ContractsOwnershipTransferred)
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
func (it *ContractsOwnershipTransferredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ContractsOwnershipTransferredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ContractsOwnershipTransferred represents a OwnershipTransferred event raised by the Contracts contract.
type ContractsOwnershipTransferred struct {
	PreviousOwner common.Address
	NewOwner      common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterOwnershipTransferred is a free log retrieval operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_Contracts *ContractsFilterer) FilterOwnershipTransferred(opts *bind.FilterOpts, previousOwner []common.Address, newOwner []common.Address) (*ContractsOwnershipTransferredIterator, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _Contracts.contract.FilterLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return &ContractsOwnershipTransferredIterator{contract: _Contracts.contract, event: "OwnershipTransferred", logs: logs, sub: sub}, nil
}

// WatchOwnershipTransferred is a free log subscription operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_Contracts *ContractsFilterer) WatchOwnershipTransferred(opts *bind.WatchOpts, sink chan<- *ContractsOwnershipTransferred, previousOwner []common.Address, newOwner []common.Address) (event.Subscription, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _Contracts.contract.WatchLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ContractsOwnershipTransferred)
				if err := _Contracts.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
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

// ParseOwnershipTransferred is a log parse operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_Contracts *ContractsFilterer) ParseOwnershipTransferred(log types.Log) (*ContractsOwnershipTransferred, error) {
	event := new(ContractsOwnershipTransferred)
	if err := _Contracts.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ContractsProviderDeregisteredIterator is returned from FilterProviderDeregistered and is used to iterate over the raw logs and unpacked data for ProviderDeregistered events raised by the Contracts contract.
type ContractsProviderDeregisteredIterator struct {
	Event *ContractsProviderDeregistered // Event containing the contract specifics and raw log

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
func (it *ContractsProviderDeregisteredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ContractsProviderDeregistered)
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
		it.Event = new(ContractsProviderDeregistered)
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
func (it *ContractsProviderDeregisteredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ContractsProviderDeregisteredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ContractsProviderDeregistered represents a ProviderDeregistered event raised by the Contracts contract.
type ContractsProviderDeregistered struct {
	Provider common.Address
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterProviderDeregistered is a free log retrieval operation binding the contract event 0xf04091b4a187e321a42001e46961e45b6a75b203fc6fb766b7e05505f6080abb.
//
// Solidity: event ProviderDeregistered(address indexed provider)
func (_Contracts *ContractsFilterer) FilterProviderDeregistered(opts *bind.FilterOpts, provider []common.Address) (*ContractsProviderDeregisteredIterator, error) {

	var providerRule []interface{}
	for _, providerItem := range provider {
		providerRule = append(providerRule, providerItem)
	}

	logs, sub, err := _Contracts.contract.FilterLogs(opts, "ProviderDeregistered", providerRule)
	if err != nil {
		return nil, err
	}
	return &ContractsProviderDeregisteredIterator{contract: _Contracts.contract, event: "ProviderDeregistered", logs: logs, sub: sub}, nil
}

// WatchProviderDeregistered is a free log subscription operation binding the contract event 0xf04091b4a187e321a42001e46961e45b6a75b203fc6fb766b7e05505f6080abb.
//
// Solidity: event ProviderDeregistered(address indexed provider)
func (_Contracts *ContractsFilterer) WatchProviderDeregistered(opts *bind.WatchOpts, sink chan<- *ContractsProviderDeregistered, provider []common.Address) (event.Subscription, error) {

	var providerRule []interface{}
	for _, providerItem := range provider {
		providerRule = append(providerRule, providerItem)
	}

	logs, sub, err := _Contracts.contract.WatchLogs(opts, "ProviderDeregistered", providerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ContractsProviderDeregistered)
				if err := _Contracts.contract.UnpackLog(event, "ProviderDeregistered", log); err != nil {
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

// ParseProviderDeregistered is a log parse operation binding the contract event 0xf04091b4a187e321a42001e46961e45b6a75b203fc6fb766b7e05505f6080abb.
//
// Solidity: event ProviderDeregistered(address indexed provider)
func (_Contracts *ContractsFilterer) ParseProviderDeregistered(log types.Log) (*ContractsProviderDeregistered, error) {
	event := new(ContractsProviderDeregistered)
	if err := _Contracts.contract.UnpackLog(event, "ProviderDeregistered", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ContractsProviderRegisteredIterator is returned from FilterProviderRegistered and is used to iterate over the raw logs and unpacked data for ProviderRegistered events raised by the Contracts contract.
type ContractsProviderRegisteredIterator struct {
	Event *ContractsProviderRegistered // Event containing the contract specifics and raw log

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
func (it *ContractsProviderRegisteredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ContractsProviderRegistered)
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
		it.Event = new(ContractsProviderRegistered)
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
func (it *ContractsProviderRegisteredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ContractsProviderRegisteredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ContractsProviderRegistered represents a ProviderRegistered event raised by the Contracts contract.
type ContractsProviderRegistered struct {
	Provider    common.Address
	Stake       *big.Int
	EndpointURL string
	Raw         types.Log // Blockchain specific contextual infos
}

// FilterProviderRegistered is a free log retrieval operation binding the contract event 0xf2499bc9b6818e3b377491d8b0fbef0d628a495d3ddd4b7cc5ce77bf9e228fa2.
//
// Solidity: event ProviderRegistered(address indexed provider, uint256 stake, string endpointURL)
func (_Contracts *ContractsFilterer) FilterProviderRegistered(opts *bind.FilterOpts, provider []common.Address) (*ContractsProviderRegisteredIterator, error) {

	var providerRule []interface{}
	for _, providerItem := range provider {
		providerRule = append(providerRule, providerItem)
	}

	logs, sub, err := _Contracts.contract.FilterLogs(opts, "ProviderRegistered", providerRule)
	if err != nil {
		return nil, err
	}
	return &ContractsProviderRegisteredIterator{contract: _Contracts.contract, event: "ProviderRegistered", logs: logs, sub: sub}, nil
}

// WatchProviderRegistered is a free log subscription operation binding the contract event 0xf2499bc9b6818e3b377491d8b0fbef0d628a495d3ddd4b7cc5ce77bf9e228fa2.
//
// Solidity: event ProviderRegistered(address indexed provider, uint256 stake, string endpointURL)
func (_Contracts *ContractsFilterer) WatchProviderRegistered(opts *bind.WatchOpts, sink chan<- *ContractsProviderRegistered, provider []common.Address) (event.Subscription, error) {

	var providerRule []interface{}
	for _, providerItem := range provider {
		providerRule = append(providerRule, providerItem)
	}

	logs, sub, err := _Contracts.contract.WatchLogs(opts, "ProviderRegistered", providerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ContractsProviderRegistered)
				if err := _Contracts.contract.UnpackLog(event, "ProviderRegistered", log); err != nil {
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

// ParseProviderRegistered is a log parse operation binding the contract event 0xf2499bc9b6818e3b377491d8b0fbef0d628a495d3ddd4b7cc5ce77bf9e228fa2.
//
// Solidity: event ProviderRegistered(address indexed provider, uint256 stake, string endpointURL)
func (_Contracts *ContractsFilterer) ParseProviderRegistered(log types.Log) (*ContractsProviderRegistered, error) {
	event := new(ContractsProviderRegistered)
	if err := _Contracts.contract.UnpackLog(event, "ProviderRegistered", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ContractsProviderSlashedIterator is returned from FilterProviderSlashed and is used to iterate over the raw logs and unpacked data for ProviderSlashed events raised by the Contracts contract.
type ContractsProviderSlashedIterator struct {
	Event *ContractsProviderSlashed // Event containing the contract specifics and raw log

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
func (it *ContractsProviderSlashedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ContractsProviderSlashed)
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
		it.Event = new(ContractsProviderSlashed)
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
func (it *ContractsProviderSlashedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ContractsProviderSlashedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ContractsProviderSlashed represents a ProviderSlashed event raised by the Contracts contract.
type ContractsProviderSlashed struct {
	Provider common.Address
	Amount   *big.Int
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterProviderSlashed is a free log retrieval operation binding the contract event 0x1f23fbfb3eedfa7285fbf3046ae9bdcfdfc0cbc70611752972939d76849ed5c5.
//
// Solidity: event ProviderSlashed(address indexed provider, uint256 amount)
func (_Contracts *ContractsFilterer) FilterProviderSlashed(opts *bind.FilterOpts, provider []common.Address) (*ContractsProviderSlashedIterator, error) {

	var providerRule []interface{}
	for _, providerItem := range provider {
		providerRule = append(providerRule, providerItem)
	}

	logs, sub, err := _Contracts.contract.FilterLogs(opts, "ProviderSlashed", providerRule)
	if err != nil {
		return nil, err
	}
	return &ContractsProviderSlashedIterator{contract: _Contracts.contract, event: "ProviderSlashed", logs: logs, sub: sub}, nil
}

// WatchProviderSlashed is a free log subscription operation binding the contract event 0x1f23fbfb3eedfa7285fbf3046ae9bdcfdfc0cbc70611752972939d76849ed5c5.
//
// Solidity: event ProviderSlashed(address indexed provider, uint256 amount)
func (_Contracts *ContractsFilterer) WatchProviderSlashed(opts *bind.WatchOpts, sink chan<- *ContractsProviderSlashed, provider []common.Address) (event.Subscription, error) {

	var providerRule []interface{}
	for _, providerItem := range provider {
		providerRule = append(providerRule, providerItem)
	}

	logs, sub, err := _Contracts.contract.WatchLogs(opts, "ProviderSlashed", providerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ContractsProviderSlashed)
				if err := _Contracts.contract.UnpackLog(event, "ProviderSlashed", log); err != nil {
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

// ParseProviderSlashed is a log parse operation binding the contract event 0x1f23fbfb3eedfa7285fbf3046ae9bdcfdfc0cbc70611752972939d76849ed5c5.
//
// Solidity: event ProviderSlashed(address indexed provider, uint256 amount)
func (_Contracts *ContractsFilterer) ParseProviderSlashed(log types.Log) (*ContractsProviderSlashed, error) {
	event := new(ContractsProviderSlashed)
	if err := _Contracts.contract.UnpackLog(event, "ProviderSlashed", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ContractsProviderURLUpdatedIterator is returned from FilterProviderURLUpdated and is used to iterate over the raw logs and unpacked data for ProviderURLUpdated events raised by the Contracts contract.
type ContractsProviderURLUpdatedIterator struct {
	Event *ContractsProviderURLUpdated // Event containing the contract specifics and raw log

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
func (it *ContractsProviderURLUpdatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ContractsProviderURLUpdated)
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
		it.Event = new(ContractsProviderURLUpdated)
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
func (it *ContractsProviderURLUpdatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ContractsProviderURLUpdatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ContractsProviderURLUpdated represents a ProviderURLUpdated event raised by the Contracts contract.
type ContractsProviderURLUpdated struct {
	Provider common.Address
	NewURL   string
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterProviderURLUpdated is a free log retrieval operation binding the contract event 0xf3f3d311c4f10915cec36706dee2e7d87ff331c7b97db73cbbc7b88a9407a34a.
//
// Solidity: event ProviderURLUpdated(address indexed provider, string newURL)
func (_Contracts *ContractsFilterer) FilterProviderURLUpdated(opts *bind.FilterOpts, provider []common.Address) (*ContractsProviderURLUpdatedIterator, error) {

	var providerRule []interface{}
	for _, providerItem := range provider {
		providerRule = append(providerRule, providerItem)
	}

	logs, sub, err := _Contracts.contract.FilterLogs(opts, "ProviderURLUpdated", providerRule)
	if err != nil {
		return nil, err
	}
	return &ContractsProviderURLUpdatedIterator{contract: _Contracts.contract, event: "ProviderURLUpdated", logs: logs, sub: sub}, nil
}

// WatchProviderURLUpdated is a free log subscription operation binding the contract event 0xf3f3d311c4f10915cec36706dee2e7d87ff331c7b97db73cbbc7b88a9407a34a.
//
// Solidity: event ProviderURLUpdated(address indexed provider, string newURL)
func (_Contracts *ContractsFilterer) WatchProviderURLUpdated(opts *bind.WatchOpts, sink chan<- *ContractsProviderURLUpdated, provider []common.Address) (event.Subscription, error) {

	var providerRule []interface{}
	for _, providerItem := range provider {
		providerRule = append(providerRule, providerItem)
	}

	logs, sub, err := _Contracts.contract.WatchLogs(opts, "ProviderURLUpdated", providerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ContractsProviderURLUpdated)
				if err := _Contracts.contract.UnpackLog(event, "ProviderURLUpdated", log); err != nil {
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

// ParseProviderURLUpdated is a log parse operation binding the contract event 0xf3f3d311c4f10915cec36706dee2e7d87ff331c7b97db73cbbc7b88a9407a34a.
//
// Solidity: event ProviderURLUpdated(address indexed provider, string newURL)
func (_Contracts *ContractsFilterer) ParseProviderURLUpdated(log types.Log) (*ContractsProviderURLUpdated, error) {
	event := new(ContractsProviderURLUpdated)
	if err := _Contracts.contract.UnpackLog(event, "ProviderURLUpdated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ContractsQoSUpdatedIterator is returned from FilterQoSUpdated and is used to iterate over the raw logs and unpacked data for QoSUpdated events raised by the Contracts contract.
type ContractsQoSUpdatedIterator struct {
	Event *ContractsQoSUpdated // Event containing the contract specifics and raw log

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
func (it *ContractsQoSUpdatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ContractsQoSUpdated)
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
		it.Event = new(ContractsQoSUpdated)
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
func (it *ContractsQoSUpdatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ContractsQoSUpdatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ContractsQoSUpdated represents a QoSUpdated event raised by the Contracts contract.
type ContractsQoSUpdated struct {
	Provider common.Address
	NewScore *big.Int
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterQoSUpdated is a free log retrieval operation binding the contract event 0xaeef8199fb5a7c06c4f963578d283254c5b64a4c2cf1dd22ea4191858e41c80a.
//
// Solidity: event QoSUpdated(address indexed provider, uint256 newScore)
func (_Contracts *ContractsFilterer) FilterQoSUpdated(opts *bind.FilterOpts, provider []common.Address) (*ContractsQoSUpdatedIterator, error) {

	var providerRule []interface{}
	for _, providerItem := range provider {
		providerRule = append(providerRule, providerItem)
	}

	logs, sub, err := _Contracts.contract.FilterLogs(opts, "QoSUpdated", providerRule)
	if err != nil {
		return nil, err
	}
	return &ContractsQoSUpdatedIterator{contract: _Contracts.contract, event: "QoSUpdated", logs: logs, sub: sub}, nil
}

// WatchQoSUpdated is a free log subscription operation binding the contract event 0xaeef8199fb5a7c06c4f963578d283254c5b64a4c2cf1dd22ea4191858e41c80a.
//
// Solidity: event QoSUpdated(address indexed provider, uint256 newScore)
func (_Contracts *ContractsFilterer) WatchQoSUpdated(opts *bind.WatchOpts, sink chan<- *ContractsQoSUpdated, provider []common.Address) (event.Subscription, error) {

	var providerRule []interface{}
	for _, providerItem := range provider {
		providerRule = append(providerRule, providerItem)
	}

	logs, sub, err := _Contracts.contract.WatchLogs(opts, "QoSUpdated", providerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ContractsQoSUpdated)
				if err := _Contracts.contract.UnpackLog(event, "QoSUpdated", log); err != nil {
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

// ParseQoSUpdated is a log parse operation binding the contract event 0xaeef8199fb5a7c06c4f963578d283254c5b64a4c2cf1dd22ea4191858e41c80a.
//
// Solidity: event QoSUpdated(address indexed provider, uint256 newScore)
func (_Contracts *ContractsFilterer) ParseQoSUpdated(log types.Log) (*ContractsQoSUpdated, error) {
	event := new(ContractsQoSUpdated)
	if err := _Contracts.contract.UnpackLog(event, "QoSUpdated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ContractsStakeDepositedIterator is returned from FilterStakeDeposited and is used to iterate over the raw logs and unpacked data for StakeDeposited events raised by the Contracts contract.
type ContractsStakeDepositedIterator struct {
	Event *ContractsStakeDeposited // Event containing the contract specifics and raw log

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
func (it *ContractsStakeDepositedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ContractsStakeDeposited)
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
		it.Event = new(ContractsStakeDeposited)
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
func (it *ContractsStakeDepositedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ContractsStakeDepositedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ContractsStakeDeposited represents a StakeDeposited event raised by the Contracts contract.
type ContractsStakeDeposited struct {
	Provider common.Address
	Amount   *big.Int
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterStakeDeposited is a free log retrieval operation binding the contract event 0x0a7bb2e28cc4698aac06db79cf9163bfcc20719286cf59fa7d492ceda1b8edc2.
//
// Solidity: event StakeDeposited(address indexed provider, uint256 amount)
func (_Contracts *ContractsFilterer) FilterStakeDeposited(opts *bind.FilterOpts, provider []common.Address) (*ContractsStakeDepositedIterator, error) {

	var providerRule []interface{}
	for _, providerItem := range provider {
		providerRule = append(providerRule, providerItem)
	}

	logs, sub, err := _Contracts.contract.FilterLogs(opts, "StakeDeposited", providerRule)
	if err != nil {
		return nil, err
	}
	return &ContractsStakeDepositedIterator{contract: _Contracts.contract, event: "StakeDeposited", logs: logs, sub: sub}, nil
}

// WatchStakeDeposited is a free log subscription operation binding the contract event 0x0a7bb2e28cc4698aac06db79cf9163bfcc20719286cf59fa7d492ceda1b8edc2.
//
// Solidity: event StakeDeposited(address indexed provider, uint256 amount)
func (_Contracts *ContractsFilterer) WatchStakeDeposited(opts *bind.WatchOpts, sink chan<- *ContractsStakeDeposited, provider []common.Address) (event.Subscription, error) {

	var providerRule []interface{}
	for _, providerItem := range provider {
		providerRule = append(providerRule, providerItem)
	}

	logs, sub, err := _Contracts.contract.WatchLogs(opts, "StakeDeposited", providerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ContractsStakeDeposited)
				if err := _Contracts.contract.UnpackLog(event, "StakeDeposited", log); err != nil {
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

// ParseStakeDeposited is a log parse operation binding the contract event 0x0a7bb2e28cc4698aac06db79cf9163bfcc20719286cf59fa7d492ceda1b8edc2.
//
// Solidity: event StakeDeposited(address indexed provider, uint256 amount)
func (_Contracts *ContractsFilterer) ParseStakeDeposited(log types.Log) (*ContractsStakeDeposited, error) {
	event := new(ContractsStakeDeposited)
	if err := _Contracts.contract.UnpackLog(event, "StakeDeposited", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ContractsStakeWithdrawnIterator is returned from FilterStakeWithdrawn and is used to iterate over the raw logs and unpacked data for StakeWithdrawn events raised by the Contracts contract.
type ContractsStakeWithdrawnIterator struct {
	Event *ContractsStakeWithdrawn // Event containing the contract specifics and raw log

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
func (it *ContractsStakeWithdrawnIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ContractsStakeWithdrawn)
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
		it.Event = new(ContractsStakeWithdrawn)
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
func (it *ContractsStakeWithdrawnIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ContractsStakeWithdrawnIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ContractsStakeWithdrawn represents a StakeWithdrawn event raised by the Contracts contract.
type ContractsStakeWithdrawn struct {
	Provider common.Address
	Amount   *big.Int
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterStakeWithdrawn is a free log retrieval operation binding the contract event 0x8108595eb6bad3acefa9da467d90cc2217686d5c5ac85460f8b7849c840645fc.
//
// Solidity: event StakeWithdrawn(address indexed provider, uint256 amount)
func (_Contracts *ContractsFilterer) FilterStakeWithdrawn(opts *bind.FilterOpts, provider []common.Address) (*ContractsStakeWithdrawnIterator, error) {

	var providerRule []interface{}
	for _, providerItem := range provider {
		providerRule = append(providerRule, providerItem)
	}

	logs, sub, err := _Contracts.contract.FilterLogs(opts, "StakeWithdrawn", providerRule)
	if err != nil {
		return nil, err
	}
	return &ContractsStakeWithdrawnIterator{contract: _Contracts.contract, event: "StakeWithdrawn", logs: logs, sub: sub}, nil
}

// WatchStakeWithdrawn is a free log subscription operation binding the contract event 0x8108595eb6bad3acefa9da467d90cc2217686d5c5ac85460f8b7849c840645fc.
//
// Solidity: event StakeWithdrawn(address indexed provider, uint256 amount)
func (_Contracts *ContractsFilterer) WatchStakeWithdrawn(opts *bind.WatchOpts, sink chan<- *ContractsStakeWithdrawn, provider []common.Address) (event.Subscription, error) {

	var providerRule []interface{}
	for _, providerItem := range provider {
		providerRule = append(providerRule, providerItem)
	}

	logs, sub, err := _Contracts.contract.WatchLogs(opts, "StakeWithdrawn", providerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ContractsStakeWithdrawn)
				if err := _Contracts.contract.UnpackLog(event, "StakeWithdrawn", log); err != nil {
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

// ParseStakeWithdrawn is a log parse operation binding the contract event 0x8108595eb6bad3acefa9da467d90cc2217686d5c5ac85460f8b7849c840645fc.
//
// Solidity: event StakeWithdrawn(address indexed provider, uint256 amount)
func (_Contracts *ContractsFilterer) ParseStakeWithdrawn(log types.Log) (*ContractsStakeWithdrawn, error) {
	event := new(ContractsStakeWithdrawn)
	if err := _Contracts.contract.UnpackLog(event, "StakeWithdrawn", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
