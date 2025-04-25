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

// ProviderRegistryMetaData contains all meta data concerning the ProviderRegistry contract.
var ProviderRegistryMetaData = &bind.MetaData{
	ABI: "[{\"type\":\"constructor\",\"inputs\":[{\"name\":\"_stakeToken\",\"type\":\"address\",\"internalType\":\"contractIERC20\"},{\"name\":\"_minStake\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"depositStake\",\"inputs\":[{\"name\":\"amount\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"deregister\",\"inputs\":[],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"minStake\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"owner\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"providers\",\"inputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[{\"name\":\"stake\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"qosScore\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"registered\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"register\",\"inputs\":[],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"renounceOwnership\",\"inputs\":[],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"setMinStake\",\"inputs\":[{\"name\":\"_minStake\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"slashProvider\",\"inputs\":[{\"name\":\"provider_\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"amount\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"stakeToken\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"contractIERC20\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"transferOwnership\",\"inputs\":[{\"name\":\"newOwner\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"updateQoS\",\"inputs\":[{\"name\":\"provider_\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"score\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"withdrawStake\",\"inputs\":[{\"name\":\"amount\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"event\",\"name\":\"OwnershipTransferred\",\"inputs\":[{\"name\":\"previousOwner\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"newOwner\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"ProviderDeregistered\",\"inputs\":[{\"name\":\"provider\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"ProviderRegistered\",\"inputs\":[{\"name\":\"provider\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"stake\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"ProviderSlashed\",\"inputs\":[{\"name\":\"provider\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"amount\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"QoSUpdated\",\"inputs\":[{\"name\":\"provider\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"newScore\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"StakeDeposited\",\"inputs\":[{\"name\":\"provider\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"amount\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"StakeWithdrawn\",\"inputs\":[{\"name\":\"provider\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"amount\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"}],\"anonymous\":false},{\"type\":\"error\",\"name\":\"OwnableInvalidOwner\",\"inputs\":[{\"name\":\"owner\",\"type\":\"address\",\"internalType\":\"address\"}]},{\"type\":\"error\",\"name\":\"OwnableUnauthorizedAccount\",\"inputs\":[{\"name\":\"account\",\"type\":\"address\",\"internalType\":\"address\"}]},{\"type\":\"error\",\"name\":\"ReentrancyGuardReentrantCall\",\"inputs\":[]}]",
}

// ProviderRegistryABI is the input ABI used to generate the binding from.
// Deprecated: Use ProviderRegistryMetaData.ABI instead.
var ProviderRegistryABI = ProviderRegistryMetaData.ABI

// ProviderRegistry is an auto generated Go binding around an Ethereum contract.
type ProviderRegistry struct {
	ProviderRegistryCaller     // Read-only binding to the contract
	ProviderRegistryTransactor // Write-only binding to the contract
	ProviderRegistryFilterer   // Log filterer for contract events
}

// ProviderRegistryCaller is an auto generated read-only Go binding around an Ethereum contract.
type ProviderRegistryCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ProviderRegistryTransactor is an auto generated write-only Go binding around an Ethereum contract.
type ProviderRegistryTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ProviderRegistryFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type ProviderRegistryFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ProviderRegistrySession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type ProviderRegistrySession struct {
	Contract     *ProviderRegistry // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// ProviderRegistryCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type ProviderRegistryCallerSession struct {
	Contract *ProviderRegistryCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts           // Call options to use throughout this session
}

// ProviderRegistryTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type ProviderRegistryTransactorSession struct {
	Contract     *ProviderRegistryTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts           // Transaction auth options to use throughout this session
}

// ProviderRegistryRaw is an auto generated low-level Go binding around an Ethereum contract.
type ProviderRegistryRaw struct {
	Contract *ProviderRegistry // Generic contract binding to access the raw methods on
}

// ProviderRegistryCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type ProviderRegistryCallerRaw struct {
	Contract *ProviderRegistryCaller // Generic read-only contract binding to access the raw methods on
}

// ProviderRegistryTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type ProviderRegistryTransactorRaw struct {
	Contract *ProviderRegistryTransactor // Generic write-only contract binding to access the raw methods on
}

// NewProviderRegistry creates a new instance of ProviderRegistry, bound to a specific deployed contract.
func NewProviderRegistry(address common.Address, backend bind.ContractBackend) (*ProviderRegistry, error) {
	contract, err := bindProviderRegistry(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &ProviderRegistry{ProviderRegistryCaller: ProviderRegistryCaller{contract: contract}, ProviderRegistryTransactor: ProviderRegistryTransactor{contract: contract}, ProviderRegistryFilterer: ProviderRegistryFilterer{contract: contract}}, nil
}

// NewProviderRegistryCaller creates a new read-only instance of ProviderRegistry, bound to a specific deployed contract.
func NewProviderRegistryCaller(address common.Address, caller bind.ContractCaller) (*ProviderRegistryCaller, error) {
	contract, err := bindProviderRegistry(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &ProviderRegistryCaller{contract: contract}, nil
}

// NewProviderRegistryTransactor creates a new write-only instance of ProviderRegistry, bound to a specific deployed contract.
func NewProviderRegistryTransactor(address common.Address, transactor bind.ContractTransactor) (*ProviderRegistryTransactor, error) {
	contract, err := bindProviderRegistry(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &ProviderRegistryTransactor{contract: contract}, nil
}

// NewProviderRegistryFilterer creates a new log filterer instance of ProviderRegistry, bound to a specific deployed contract.
func NewProviderRegistryFilterer(address common.Address, filterer bind.ContractFilterer) (*ProviderRegistryFilterer, error) {
	contract, err := bindProviderRegistry(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &ProviderRegistryFilterer{contract: contract}, nil
}

// bindProviderRegistry binds a generic wrapper to an already deployed contract.
func bindProviderRegistry(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := ProviderRegistryMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_ProviderRegistry *ProviderRegistryRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _ProviderRegistry.Contract.ProviderRegistryCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_ProviderRegistry *ProviderRegistryRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ProviderRegistry.Contract.ProviderRegistryTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_ProviderRegistry *ProviderRegistryRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _ProviderRegistry.Contract.ProviderRegistryTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_ProviderRegistry *ProviderRegistryCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _ProviderRegistry.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_ProviderRegistry *ProviderRegistryTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ProviderRegistry.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_ProviderRegistry *ProviderRegistryTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _ProviderRegistry.Contract.contract.Transact(opts, method, params...)
}

// MinStake is a free data retrieval call binding the contract method 0x375b3c0a.
//
// Solidity: function minStake() view returns(uint256)
func (_ProviderRegistry *ProviderRegistryCaller) MinStake(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _ProviderRegistry.contract.Call(opts, &out, "minStake")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// MinStake is a free data retrieval call binding the contract method 0x375b3c0a.
//
// Solidity: function minStake() view returns(uint256)
func (_ProviderRegistry *ProviderRegistrySession) MinStake() (*big.Int, error) {
	return _ProviderRegistry.Contract.MinStake(&_ProviderRegistry.CallOpts)
}

// MinStake is a free data retrieval call binding the contract method 0x375b3c0a.
//
// Solidity: function minStake() view returns(uint256)
func (_ProviderRegistry *ProviderRegistryCallerSession) MinStake() (*big.Int, error) {
	return _ProviderRegistry.Contract.MinStake(&_ProviderRegistry.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_ProviderRegistry *ProviderRegistryCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _ProviderRegistry.contract.Call(opts, &out, "owner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_ProviderRegistry *ProviderRegistrySession) Owner() (common.Address, error) {
	return _ProviderRegistry.Contract.Owner(&_ProviderRegistry.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_ProviderRegistry *ProviderRegistryCallerSession) Owner() (common.Address, error) {
	return _ProviderRegistry.Contract.Owner(&_ProviderRegistry.CallOpts)
}

// Providers is a free data retrieval call binding the contract method 0x0787bc27.
//
// Solidity: function providers(address ) view returns(uint256 stake, uint256 qosScore, bool registered)
func (_ProviderRegistry *ProviderRegistryCaller) Providers(opts *bind.CallOpts, arg0 common.Address) (struct {
	Stake      *big.Int
	QosScore   *big.Int
	Registered bool
}, error) {
	var out []interface{}
	err := _ProviderRegistry.contract.Call(opts, &out, "providers", arg0)

	outstruct := new(struct {
		Stake      *big.Int
		QosScore   *big.Int
		Registered bool
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.Stake = *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)
	outstruct.QosScore = *abi.ConvertType(out[1], new(*big.Int)).(**big.Int)
	outstruct.Registered = *abi.ConvertType(out[2], new(bool)).(*bool)

	return *outstruct, err

}

// Providers is a free data retrieval call binding the contract method 0x0787bc27.
//
// Solidity: function providers(address ) view returns(uint256 stake, uint256 qosScore, bool registered)
func (_ProviderRegistry *ProviderRegistrySession) Providers(arg0 common.Address) (struct {
	Stake      *big.Int
	QosScore   *big.Int
	Registered bool
}, error) {
	return _ProviderRegistry.Contract.Providers(&_ProviderRegistry.CallOpts, arg0)
}

// Providers is a free data retrieval call binding the contract method 0x0787bc27.
//
// Solidity: function providers(address ) view returns(uint256 stake, uint256 qosScore, bool registered)
func (_ProviderRegistry *ProviderRegistryCallerSession) Providers(arg0 common.Address) (struct {
	Stake      *big.Int
	QosScore   *big.Int
	Registered bool
}, error) {
	return _ProviderRegistry.Contract.Providers(&_ProviderRegistry.CallOpts, arg0)
}

// StakeToken is a free data retrieval call binding the contract method 0x51ed6a30.
//
// Solidity: function stakeToken() view returns(address)
func (_ProviderRegistry *ProviderRegistryCaller) StakeToken(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _ProviderRegistry.contract.Call(opts, &out, "stakeToken")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// StakeToken is a free data retrieval call binding the contract method 0x51ed6a30.
//
// Solidity: function stakeToken() view returns(address)
func (_ProviderRegistry *ProviderRegistrySession) StakeToken() (common.Address, error) {
	return _ProviderRegistry.Contract.StakeToken(&_ProviderRegistry.CallOpts)
}

// StakeToken is a free data retrieval call binding the contract method 0x51ed6a30.
//
// Solidity: function stakeToken() view returns(address)
func (_ProviderRegistry *ProviderRegistryCallerSession) StakeToken() (common.Address, error) {
	return _ProviderRegistry.Contract.StakeToken(&_ProviderRegistry.CallOpts)
}

// DepositStake is a paid mutator transaction binding the contract method 0xcb82cc8f.
//
// Solidity: function depositStake(uint256 amount) returns()
func (_ProviderRegistry *ProviderRegistryTransactor) DepositStake(opts *bind.TransactOpts, amount *big.Int) (*types.Transaction, error) {
	return _ProviderRegistry.contract.Transact(opts, "depositStake", amount)
}

// DepositStake is a paid mutator transaction binding the contract method 0xcb82cc8f.
//
// Solidity: function depositStake(uint256 amount) returns()
func (_ProviderRegistry *ProviderRegistrySession) DepositStake(amount *big.Int) (*types.Transaction, error) {
	return _ProviderRegistry.Contract.DepositStake(&_ProviderRegistry.TransactOpts, amount)
}

// DepositStake is a paid mutator transaction binding the contract method 0xcb82cc8f.
//
// Solidity: function depositStake(uint256 amount) returns()
func (_ProviderRegistry *ProviderRegistryTransactorSession) DepositStake(amount *big.Int) (*types.Transaction, error) {
	return _ProviderRegistry.Contract.DepositStake(&_ProviderRegistry.TransactOpts, amount)
}

// Deregister is a paid mutator transaction binding the contract method 0xaff5edb1.
//
// Solidity: function deregister() returns()
func (_ProviderRegistry *ProviderRegistryTransactor) Deregister(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ProviderRegistry.contract.Transact(opts, "deregister")
}

// Deregister is a paid mutator transaction binding the contract method 0xaff5edb1.
//
// Solidity: function deregister() returns()
func (_ProviderRegistry *ProviderRegistrySession) Deregister() (*types.Transaction, error) {
	return _ProviderRegistry.Contract.Deregister(&_ProviderRegistry.TransactOpts)
}

// Deregister is a paid mutator transaction binding the contract method 0xaff5edb1.
//
// Solidity: function deregister() returns()
func (_ProviderRegistry *ProviderRegistryTransactorSession) Deregister() (*types.Transaction, error) {
	return _ProviderRegistry.Contract.Deregister(&_ProviderRegistry.TransactOpts)
}

// Register is a paid mutator transaction binding the contract method 0x1aa3a008.
//
// Solidity: function register() returns()
func (_ProviderRegistry *ProviderRegistryTransactor) Register(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ProviderRegistry.contract.Transact(opts, "register")
}

// Register is a paid mutator transaction binding the contract method 0x1aa3a008.
//
// Solidity: function register() returns()
func (_ProviderRegistry *ProviderRegistrySession) Register() (*types.Transaction, error) {
	return _ProviderRegistry.Contract.Register(&_ProviderRegistry.TransactOpts)
}

// Register is a paid mutator transaction binding the contract method 0x1aa3a008.
//
// Solidity: function register() returns()
func (_ProviderRegistry *ProviderRegistryTransactorSession) Register() (*types.Transaction, error) {
	return _ProviderRegistry.Contract.Register(&_ProviderRegistry.TransactOpts)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_ProviderRegistry *ProviderRegistryTransactor) RenounceOwnership(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ProviderRegistry.contract.Transact(opts, "renounceOwnership")
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_ProviderRegistry *ProviderRegistrySession) RenounceOwnership() (*types.Transaction, error) {
	return _ProviderRegistry.Contract.RenounceOwnership(&_ProviderRegistry.TransactOpts)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_ProviderRegistry *ProviderRegistryTransactorSession) RenounceOwnership() (*types.Transaction, error) {
	return _ProviderRegistry.Contract.RenounceOwnership(&_ProviderRegistry.TransactOpts)
}

// SetMinStake is a paid mutator transaction binding the contract method 0x8c80fd90.
//
// Solidity: function setMinStake(uint256 _minStake) returns()
func (_ProviderRegistry *ProviderRegistryTransactor) SetMinStake(opts *bind.TransactOpts, _minStake *big.Int) (*types.Transaction, error) {
	return _ProviderRegistry.contract.Transact(opts, "setMinStake", _minStake)
}

// SetMinStake is a paid mutator transaction binding the contract method 0x8c80fd90.
//
// Solidity: function setMinStake(uint256 _minStake) returns()
func (_ProviderRegistry *ProviderRegistrySession) SetMinStake(_minStake *big.Int) (*types.Transaction, error) {
	return _ProviderRegistry.Contract.SetMinStake(&_ProviderRegistry.TransactOpts, _minStake)
}

// SetMinStake is a paid mutator transaction binding the contract method 0x8c80fd90.
//
// Solidity: function setMinStake(uint256 _minStake) returns()
func (_ProviderRegistry *ProviderRegistryTransactorSession) SetMinStake(_minStake *big.Int) (*types.Transaction, error) {
	return _ProviderRegistry.Contract.SetMinStake(&_ProviderRegistry.TransactOpts, _minStake)
}

// SlashProvider is a paid mutator transaction binding the contract method 0x3febfc0f.
//
// Solidity: function slashProvider(address provider_, uint256 amount) returns()
func (_ProviderRegistry *ProviderRegistryTransactor) SlashProvider(opts *bind.TransactOpts, provider_ common.Address, amount *big.Int) (*types.Transaction, error) {
	return _ProviderRegistry.contract.Transact(opts, "slashProvider", provider_, amount)
}

// SlashProvider is a paid mutator transaction binding the contract method 0x3febfc0f.
//
// Solidity: function slashProvider(address provider_, uint256 amount) returns()
func (_ProviderRegistry *ProviderRegistrySession) SlashProvider(provider_ common.Address, amount *big.Int) (*types.Transaction, error) {
	return _ProviderRegistry.Contract.SlashProvider(&_ProviderRegistry.TransactOpts, provider_, amount)
}

// SlashProvider is a paid mutator transaction binding the contract method 0x3febfc0f.
//
// Solidity: function slashProvider(address provider_, uint256 amount) returns()
func (_ProviderRegistry *ProviderRegistryTransactorSession) SlashProvider(provider_ common.Address, amount *big.Int) (*types.Transaction, error) {
	return _ProviderRegistry.Contract.SlashProvider(&_ProviderRegistry.TransactOpts, provider_, amount)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_ProviderRegistry *ProviderRegistryTransactor) TransferOwnership(opts *bind.TransactOpts, newOwner common.Address) (*types.Transaction, error) {
	return _ProviderRegistry.contract.Transact(opts, "transferOwnership", newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_ProviderRegistry *ProviderRegistrySession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _ProviderRegistry.Contract.TransferOwnership(&_ProviderRegistry.TransactOpts, newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_ProviderRegistry *ProviderRegistryTransactorSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _ProviderRegistry.Contract.TransferOwnership(&_ProviderRegistry.TransactOpts, newOwner)
}

// UpdateQoS is a paid mutator transaction binding the contract method 0x117835a4.
//
// Solidity: function updateQoS(address provider_, uint256 score) returns()
func (_ProviderRegistry *ProviderRegistryTransactor) UpdateQoS(opts *bind.TransactOpts, provider_ common.Address, score *big.Int) (*types.Transaction, error) {
	return _ProviderRegistry.contract.Transact(opts, "updateQoS", provider_, score)
}

// UpdateQoS is a paid mutator transaction binding the contract method 0x117835a4.
//
// Solidity: function updateQoS(address provider_, uint256 score) returns()
func (_ProviderRegistry *ProviderRegistrySession) UpdateQoS(provider_ common.Address, score *big.Int) (*types.Transaction, error) {
	return _ProviderRegistry.Contract.UpdateQoS(&_ProviderRegistry.TransactOpts, provider_, score)
}

// UpdateQoS is a paid mutator transaction binding the contract method 0x117835a4.
//
// Solidity: function updateQoS(address provider_, uint256 score) returns()
func (_ProviderRegistry *ProviderRegistryTransactorSession) UpdateQoS(provider_ common.Address, score *big.Int) (*types.Transaction, error) {
	return _ProviderRegistry.Contract.UpdateQoS(&_ProviderRegistry.TransactOpts, provider_, score)
}

// WithdrawStake is a paid mutator transaction binding the contract method 0x25d5971f.
//
// Solidity: function withdrawStake(uint256 amount) returns()
func (_ProviderRegistry *ProviderRegistryTransactor) WithdrawStake(opts *bind.TransactOpts, amount *big.Int) (*types.Transaction, error) {
	return _ProviderRegistry.contract.Transact(opts, "withdrawStake", amount)
}

// WithdrawStake is a paid mutator transaction binding the contract method 0x25d5971f.
//
// Solidity: function withdrawStake(uint256 amount) returns()
func (_ProviderRegistry *ProviderRegistrySession) WithdrawStake(amount *big.Int) (*types.Transaction, error) {
	return _ProviderRegistry.Contract.WithdrawStake(&_ProviderRegistry.TransactOpts, amount)
}

// WithdrawStake is a paid mutator transaction binding the contract method 0x25d5971f.
//
// Solidity: function withdrawStake(uint256 amount) returns()
func (_ProviderRegistry *ProviderRegistryTransactorSession) WithdrawStake(amount *big.Int) (*types.Transaction, error) {
	return _ProviderRegistry.Contract.WithdrawStake(&_ProviderRegistry.TransactOpts, amount)
}

// ProviderRegistryOwnershipTransferredIterator is returned from FilterOwnershipTransferred and is used to iterate over the raw logs and unpacked data for OwnershipTransferred events raised by the ProviderRegistry contract.
type ProviderRegistryOwnershipTransferredIterator struct {
	Event *ProviderRegistryOwnershipTransferred // Event containing the contract specifics and raw log

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
func (it *ProviderRegistryOwnershipTransferredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ProviderRegistryOwnershipTransferred)
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
		it.Event = new(ProviderRegistryOwnershipTransferred)
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
func (it *ProviderRegistryOwnershipTransferredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ProviderRegistryOwnershipTransferredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ProviderRegistryOwnershipTransferred represents a OwnershipTransferred event raised by the ProviderRegistry contract.
type ProviderRegistryOwnershipTransferred struct {
	PreviousOwner common.Address
	NewOwner      common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterOwnershipTransferred is a free log retrieval operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_ProviderRegistry *ProviderRegistryFilterer) FilterOwnershipTransferred(opts *bind.FilterOpts, previousOwner []common.Address, newOwner []common.Address) (*ProviderRegistryOwnershipTransferredIterator, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _ProviderRegistry.contract.FilterLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return &ProviderRegistryOwnershipTransferredIterator{contract: _ProviderRegistry.contract, event: "OwnershipTransferred", logs: logs, sub: sub}, nil
}

// WatchOwnershipTransferred is a free log subscription operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_ProviderRegistry *ProviderRegistryFilterer) WatchOwnershipTransferred(opts *bind.WatchOpts, sink chan<- *ProviderRegistryOwnershipTransferred, previousOwner []common.Address, newOwner []common.Address) (event.Subscription, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _ProviderRegistry.contract.WatchLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ProviderRegistryOwnershipTransferred)
				if err := _ProviderRegistry.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
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
func (_ProviderRegistry *ProviderRegistryFilterer) ParseOwnershipTransferred(log types.Log) (*ProviderRegistryOwnershipTransferred, error) {
	event := new(ProviderRegistryOwnershipTransferred)
	if err := _ProviderRegistry.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ProviderRegistryProviderDeregisteredIterator is returned from FilterProviderDeregistered and is used to iterate over the raw logs and unpacked data for ProviderDeregistered events raised by the ProviderRegistry contract.
type ProviderRegistryProviderDeregisteredIterator struct {
	Event *ProviderRegistryProviderDeregistered // Event containing the contract specifics and raw log

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
func (it *ProviderRegistryProviderDeregisteredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ProviderRegistryProviderDeregistered)
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
		it.Event = new(ProviderRegistryProviderDeregistered)
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
func (it *ProviderRegistryProviderDeregisteredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ProviderRegistryProviderDeregisteredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ProviderRegistryProviderDeregistered represents a ProviderDeregistered event raised by the ProviderRegistry contract.
type ProviderRegistryProviderDeregistered struct {
	Provider common.Address
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterProviderDeregistered is a free log retrieval operation binding the contract event 0xf04091b4a187e321a42001e46961e45b6a75b203fc6fb766b7e05505f6080abb.
//
// Solidity: event ProviderDeregistered(address indexed provider)
func (_ProviderRegistry *ProviderRegistryFilterer) FilterProviderDeregistered(opts *bind.FilterOpts, provider []common.Address) (*ProviderRegistryProviderDeregisteredIterator, error) {

	var providerRule []interface{}
	for _, providerItem := range provider {
		providerRule = append(providerRule, providerItem)
	}

	logs, sub, err := _ProviderRegistry.contract.FilterLogs(opts, "ProviderDeregistered", providerRule)
	if err != nil {
		return nil, err
	}
	return &ProviderRegistryProviderDeregisteredIterator{contract: _ProviderRegistry.contract, event: "ProviderDeregistered", logs: logs, sub: sub}, nil
}

// WatchProviderDeregistered is a free log subscription operation binding the contract event 0xf04091b4a187e321a42001e46961e45b6a75b203fc6fb766b7e05505f6080abb.
//
// Solidity: event ProviderDeregistered(address indexed provider)
func (_ProviderRegistry *ProviderRegistryFilterer) WatchProviderDeregistered(opts *bind.WatchOpts, sink chan<- *ProviderRegistryProviderDeregistered, provider []common.Address) (event.Subscription, error) {

	var providerRule []interface{}
	for _, providerItem := range provider {
		providerRule = append(providerRule, providerItem)
	}

	logs, sub, err := _ProviderRegistry.contract.WatchLogs(opts, "ProviderDeregistered", providerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ProviderRegistryProviderDeregistered)
				if err := _ProviderRegistry.contract.UnpackLog(event, "ProviderDeregistered", log); err != nil {
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
func (_ProviderRegistry *ProviderRegistryFilterer) ParseProviderDeregistered(log types.Log) (*ProviderRegistryProviderDeregistered, error) {
	event := new(ProviderRegistryProviderDeregistered)
	if err := _ProviderRegistry.contract.UnpackLog(event, "ProviderDeregistered", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ProviderRegistryProviderRegisteredIterator is returned from FilterProviderRegistered and is used to iterate over the raw logs and unpacked data for ProviderRegistered events raised by the ProviderRegistry contract.
type ProviderRegistryProviderRegisteredIterator struct {
	Event *ProviderRegistryProviderRegistered // Event containing the contract specifics and raw log

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
func (it *ProviderRegistryProviderRegisteredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ProviderRegistryProviderRegistered)
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
		it.Event = new(ProviderRegistryProviderRegistered)
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
func (it *ProviderRegistryProviderRegisteredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ProviderRegistryProviderRegisteredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ProviderRegistryProviderRegistered represents a ProviderRegistered event raised by the ProviderRegistry contract.
type ProviderRegistryProviderRegistered struct {
	Provider common.Address
	Stake    *big.Int
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterProviderRegistered is a free log retrieval operation binding the contract event 0x90c9734131c1e4fb36cde2d71e6feb93fb258f71be8a85411c173d25e1516e80.
//
// Solidity: event ProviderRegistered(address indexed provider, uint256 stake)
func (_ProviderRegistry *ProviderRegistryFilterer) FilterProviderRegistered(opts *bind.FilterOpts, provider []common.Address) (*ProviderRegistryProviderRegisteredIterator, error) {

	var providerRule []interface{}
	for _, providerItem := range provider {
		providerRule = append(providerRule, providerItem)
	}

	logs, sub, err := _ProviderRegistry.contract.FilterLogs(opts, "ProviderRegistered", providerRule)
	if err != nil {
		return nil, err
	}
	return &ProviderRegistryProviderRegisteredIterator{contract: _ProviderRegistry.contract, event: "ProviderRegistered", logs: logs, sub: sub}, nil
}

// WatchProviderRegistered is a free log subscription operation binding the contract event 0x90c9734131c1e4fb36cde2d71e6feb93fb258f71be8a85411c173d25e1516e80.
//
// Solidity: event ProviderRegistered(address indexed provider, uint256 stake)
func (_ProviderRegistry *ProviderRegistryFilterer) WatchProviderRegistered(opts *bind.WatchOpts, sink chan<- *ProviderRegistryProviderRegistered, provider []common.Address) (event.Subscription, error) {

	var providerRule []interface{}
	for _, providerItem := range provider {
		providerRule = append(providerRule, providerItem)
	}

	logs, sub, err := _ProviderRegistry.contract.WatchLogs(opts, "ProviderRegistered", providerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ProviderRegistryProviderRegistered)
				if err := _ProviderRegistry.contract.UnpackLog(event, "ProviderRegistered", log); err != nil {
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

// ParseProviderRegistered is a log parse operation binding the contract event 0x90c9734131c1e4fb36cde2d71e6feb93fb258f71be8a85411c173d25e1516e80.
//
// Solidity: event ProviderRegistered(address indexed provider, uint256 stake)
func (_ProviderRegistry *ProviderRegistryFilterer) ParseProviderRegistered(log types.Log) (*ProviderRegistryProviderRegistered, error) {
	event := new(ProviderRegistryProviderRegistered)
	if err := _ProviderRegistry.contract.UnpackLog(event, "ProviderRegistered", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ProviderRegistryProviderSlashedIterator is returned from FilterProviderSlashed and is used to iterate over the raw logs and unpacked data for ProviderSlashed events raised by the ProviderRegistry contract.
type ProviderRegistryProviderSlashedIterator struct {
	Event *ProviderRegistryProviderSlashed // Event containing the contract specifics and raw log

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
func (it *ProviderRegistryProviderSlashedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ProviderRegistryProviderSlashed)
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
		it.Event = new(ProviderRegistryProviderSlashed)
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
func (it *ProviderRegistryProviderSlashedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ProviderRegistryProviderSlashedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ProviderRegistryProviderSlashed represents a ProviderSlashed event raised by the ProviderRegistry contract.
type ProviderRegistryProviderSlashed struct {
	Provider common.Address
	Amount   *big.Int
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterProviderSlashed is a free log retrieval operation binding the contract event 0x1f23fbfb3eedfa7285fbf3046ae9bdcfdfc0cbc70611752972939d76849ed5c5.
//
// Solidity: event ProviderSlashed(address indexed provider, uint256 amount)
func (_ProviderRegistry *ProviderRegistryFilterer) FilterProviderSlashed(opts *bind.FilterOpts, provider []common.Address) (*ProviderRegistryProviderSlashedIterator, error) {

	var providerRule []interface{}
	for _, providerItem := range provider {
		providerRule = append(providerRule, providerItem)
	}

	logs, sub, err := _ProviderRegistry.contract.FilterLogs(opts, "ProviderSlashed", providerRule)
	if err != nil {
		return nil, err
	}
	return &ProviderRegistryProviderSlashedIterator{contract: _ProviderRegistry.contract, event: "ProviderSlashed", logs: logs, sub: sub}, nil
}

// WatchProviderSlashed is a free log subscription operation binding the contract event 0x1f23fbfb3eedfa7285fbf3046ae9bdcfdfc0cbc70611752972939d76849ed5c5.
//
// Solidity: event ProviderSlashed(address indexed provider, uint256 amount)
func (_ProviderRegistry *ProviderRegistryFilterer) WatchProviderSlashed(opts *bind.WatchOpts, sink chan<- *ProviderRegistryProviderSlashed, provider []common.Address) (event.Subscription, error) {

	var providerRule []interface{}
	for _, providerItem := range provider {
		providerRule = append(providerRule, providerItem)
	}

	logs, sub, err := _ProviderRegistry.contract.WatchLogs(opts, "ProviderSlashed", providerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ProviderRegistryProviderSlashed)
				if err := _ProviderRegistry.contract.UnpackLog(event, "ProviderSlashed", log); err != nil {
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
func (_ProviderRegistry *ProviderRegistryFilterer) ParseProviderSlashed(log types.Log) (*ProviderRegistryProviderSlashed, error) {
	event := new(ProviderRegistryProviderSlashed)
	if err := _ProviderRegistry.contract.UnpackLog(event, "ProviderSlashed", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ProviderRegistryQoSUpdatedIterator is returned from FilterQoSUpdated and is used to iterate over the raw logs and unpacked data for QoSUpdated events raised by the ProviderRegistry contract.
type ProviderRegistryQoSUpdatedIterator struct {
	Event *ProviderRegistryQoSUpdated // Event containing the contract specifics and raw log

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
func (it *ProviderRegistryQoSUpdatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ProviderRegistryQoSUpdated)
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
		it.Event = new(ProviderRegistryQoSUpdated)
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
func (it *ProviderRegistryQoSUpdatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ProviderRegistryQoSUpdatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ProviderRegistryQoSUpdated represents a QoSUpdated event raised by the ProviderRegistry contract.
type ProviderRegistryQoSUpdated struct {
	Provider common.Address
	NewScore *big.Int
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterQoSUpdated is a free log retrieval operation binding the contract event 0xaeef8199fb5a7c06c4f963578d283254c5b64a4c2cf1dd22ea4191858e41c80a.
//
// Solidity: event QoSUpdated(address indexed provider, uint256 newScore)
func (_ProviderRegistry *ProviderRegistryFilterer) FilterQoSUpdated(opts *bind.FilterOpts, provider []common.Address) (*ProviderRegistryQoSUpdatedIterator, error) {

	var providerRule []interface{}
	for _, providerItem := range provider {
		providerRule = append(providerRule, providerItem)
	}

	logs, sub, err := _ProviderRegistry.contract.FilterLogs(opts, "QoSUpdated", providerRule)
	if err != nil {
		return nil, err
	}
	return &ProviderRegistryQoSUpdatedIterator{contract: _ProviderRegistry.contract, event: "QoSUpdated", logs: logs, sub: sub}, nil
}

// WatchQoSUpdated is a free log subscription operation binding the contract event 0xaeef8199fb5a7c06c4f963578d283254c5b64a4c2cf1dd22ea4191858e41c80a.
//
// Solidity: event QoSUpdated(address indexed provider, uint256 newScore)
func (_ProviderRegistry *ProviderRegistryFilterer) WatchQoSUpdated(opts *bind.WatchOpts, sink chan<- *ProviderRegistryQoSUpdated, provider []common.Address) (event.Subscription, error) {

	var providerRule []interface{}
	for _, providerItem := range provider {
		providerRule = append(providerRule, providerItem)
	}

	logs, sub, err := _ProviderRegistry.contract.WatchLogs(opts, "QoSUpdated", providerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ProviderRegistryQoSUpdated)
				if err := _ProviderRegistry.contract.UnpackLog(event, "QoSUpdated", log); err != nil {
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
func (_ProviderRegistry *ProviderRegistryFilterer) ParseQoSUpdated(log types.Log) (*ProviderRegistryQoSUpdated, error) {
	event := new(ProviderRegistryQoSUpdated)
	if err := _ProviderRegistry.contract.UnpackLog(event, "QoSUpdated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ProviderRegistryStakeDepositedIterator is returned from FilterStakeDeposited and is used to iterate over the raw logs and unpacked data for StakeDeposited events raised by the ProviderRegistry contract.
type ProviderRegistryStakeDepositedIterator struct {
	Event *ProviderRegistryStakeDeposited // Event containing the contract specifics and raw log

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
func (it *ProviderRegistryStakeDepositedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ProviderRegistryStakeDeposited)
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
		it.Event = new(ProviderRegistryStakeDeposited)
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
func (it *ProviderRegistryStakeDepositedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ProviderRegistryStakeDepositedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ProviderRegistryStakeDeposited represents a StakeDeposited event raised by the ProviderRegistry contract.
type ProviderRegistryStakeDeposited struct {
	Provider common.Address
	Amount   *big.Int
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterStakeDeposited is a free log retrieval operation binding the contract event 0x0a7bb2e28cc4698aac06db79cf9163bfcc20719286cf59fa7d492ceda1b8edc2.
//
// Solidity: event StakeDeposited(address indexed provider, uint256 amount)
func (_ProviderRegistry *ProviderRegistryFilterer) FilterStakeDeposited(opts *bind.FilterOpts, provider []common.Address) (*ProviderRegistryStakeDepositedIterator, error) {

	var providerRule []interface{}
	for _, providerItem := range provider {
		providerRule = append(providerRule, providerItem)
	}

	logs, sub, err := _ProviderRegistry.contract.FilterLogs(opts, "StakeDeposited", providerRule)
	if err != nil {
		return nil, err
	}
	return &ProviderRegistryStakeDepositedIterator{contract: _ProviderRegistry.contract, event: "StakeDeposited", logs: logs, sub: sub}, nil
}

// WatchStakeDeposited is a free log subscription operation binding the contract event 0x0a7bb2e28cc4698aac06db79cf9163bfcc20719286cf59fa7d492ceda1b8edc2.
//
// Solidity: event StakeDeposited(address indexed provider, uint256 amount)
func (_ProviderRegistry *ProviderRegistryFilterer) WatchStakeDeposited(opts *bind.WatchOpts, sink chan<- *ProviderRegistryStakeDeposited, provider []common.Address) (event.Subscription, error) {

	var providerRule []interface{}
	for _, providerItem := range provider {
		providerRule = append(providerRule, providerItem)
	}

	logs, sub, err := _ProviderRegistry.contract.WatchLogs(opts, "StakeDeposited", providerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ProviderRegistryStakeDeposited)
				if err := _ProviderRegistry.contract.UnpackLog(event, "StakeDeposited", log); err != nil {
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
func (_ProviderRegistry *ProviderRegistryFilterer) ParseStakeDeposited(log types.Log) (*ProviderRegistryStakeDeposited, error) {
	event := new(ProviderRegistryStakeDeposited)
	if err := _ProviderRegistry.contract.UnpackLog(event, "StakeDeposited", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ProviderRegistryStakeWithdrawnIterator is returned from FilterStakeWithdrawn and is used to iterate over the raw logs and unpacked data for StakeWithdrawn events raised by the ProviderRegistry contract.
type ProviderRegistryStakeWithdrawnIterator struct {
	Event *ProviderRegistryStakeWithdrawn // Event containing the contract specifics and raw log

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
func (it *ProviderRegistryStakeWithdrawnIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ProviderRegistryStakeWithdrawn)
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
		it.Event = new(ProviderRegistryStakeWithdrawn)
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
func (it *ProviderRegistryStakeWithdrawnIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ProviderRegistryStakeWithdrawnIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ProviderRegistryStakeWithdrawn represents a StakeWithdrawn event raised by the ProviderRegistry contract.
type ProviderRegistryStakeWithdrawn struct {
	Provider common.Address
	Amount   *big.Int
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterStakeWithdrawn is a free log retrieval operation binding the contract event 0x8108595eb6bad3acefa9da467d90cc2217686d5c5ac85460f8b7849c840645fc.
//
// Solidity: event StakeWithdrawn(address indexed provider, uint256 amount)
func (_ProviderRegistry *ProviderRegistryFilterer) FilterStakeWithdrawn(opts *bind.FilterOpts, provider []common.Address) (*ProviderRegistryStakeWithdrawnIterator, error) {

	var providerRule []interface{}
	for _, providerItem := range provider {
		providerRule = append(providerRule, providerItem)
	}

	logs, sub, err := _ProviderRegistry.contract.FilterLogs(opts, "StakeWithdrawn", providerRule)
	if err != nil {
		return nil, err
	}
	return &ProviderRegistryStakeWithdrawnIterator{contract: _ProviderRegistry.contract, event: "StakeWithdrawn", logs: logs, sub: sub}, nil
}

// WatchStakeWithdrawn is a free log subscription operation binding the contract event 0x8108595eb6bad3acefa9da467d90cc2217686d5c5ac85460f8b7849c840645fc.
//
// Solidity: event StakeWithdrawn(address indexed provider, uint256 amount)
func (_ProviderRegistry *ProviderRegistryFilterer) WatchStakeWithdrawn(opts *bind.WatchOpts, sink chan<- *ProviderRegistryStakeWithdrawn, provider []common.Address) (event.Subscription, error) {

	var providerRule []interface{}
	for _, providerItem := range provider {
		providerRule = append(providerRule, providerItem)
	}

	logs, sub, err := _ProviderRegistry.contract.WatchLogs(opts, "StakeWithdrawn", providerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ProviderRegistryStakeWithdrawn)
				if err := _ProviderRegistry.contract.UnpackLog(event, "StakeWithdrawn", log); err != nil {
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
func (_ProviderRegistry *ProviderRegistryFilterer) ParseStakeWithdrawn(log types.Log) (*ProviderRegistryStakeWithdrawn, error) {
	event := new(ProviderRegistryStakeWithdrawn)
	if err := _ProviderRegistry.contract.UnpackLog(event, "StakeWithdrawn", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
