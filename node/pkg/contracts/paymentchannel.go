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

// PaymentChannelMetaData contains all meta data concerning the PaymentChannel contract.
var PaymentChannelMetaData = &bind.MetaData{
	ABI: "[{\"type\":\"function\",\"name\":\"channels\",\"inputs\":[{\"name\":\"\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"outputs\":[{\"name\":\"user\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"provider\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"token\",\"type\":\"address\",\"internalType\":\"contractIERC20\"},{\"name\":\"deposit\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"expiration\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"open\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"claimTimeout\",\"inputs\":[{\"name\":\"channelId\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"closeChannel\",\"inputs\":[{\"name\":\"channelId\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"amount\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"signature\",\"type\":\"bytes\",\"internalType\":\"bytes\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"openChannel\",\"inputs\":[{\"name\":\"provider\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"token\",\"type\":\"address\",\"internalType\":\"contractIERC20\"},{\"name\":\"deposit\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"duration\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[{\"name\":\"channelId\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"stateMutability\":\"nonpayable\"},{\"type\":\"event\",\"name\":\"ChannelClosed\",\"inputs\":[{\"name\":\"channelId\",\"type\":\"bytes32\",\"indexed\":true,\"internalType\":\"bytes32\"},{\"name\":\"amountReceived\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"ChannelExpired\",\"inputs\":[{\"name\":\"channelId\",\"type\":\"bytes32\",\"indexed\":true,\"internalType\":\"bytes32\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"ChannelOpened\",\"inputs\":[{\"name\":\"channelId\",\"type\":\"bytes32\",\"indexed\":true,\"internalType\":\"bytes32\"},{\"name\":\"user\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"provider\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"deposit\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"},{\"name\":\"expiration\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"}],\"anonymous\":false},{\"type\":\"error\",\"name\":\"ECDSAInvalidSignature\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"ECDSAInvalidSignatureLength\",\"inputs\":[{\"name\":\"length\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]},{\"type\":\"error\",\"name\":\"ECDSAInvalidSignatureS\",\"inputs\":[{\"name\":\"s\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}]},{\"type\":\"error\",\"name\":\"ReentrancyGuardReentrantCall\",\"inputs\":[]}]",
}

// PaymentChannelABI is the input ABI used to generate the binding from.
// Deprecated: Use PaymentChannelMetaData.ABI instead.
var PaymentChannelABI = PaymentChannelMetaData.ABI

// PaymentChannel is an auto generated Go binding around an Ethereum contract.
type PaymentChannel struct {
	PaymentChannelCaller     // Read-only binding to the contract
	PaymentChannelTransactor // Write-only binding to the contract
	PaymentChannelFilterer   // Log filterer for contract events
}

// PaymentChannelCaller is an auto generated read-only Go binding around an Ethereum contract.
type PaymentChannelCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// PaymentChannelTransactor is an auto generated write-only Go binding around an Ethereum contract.
type PaymentChannelTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// PaymentChannelFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type PaymentChannelFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// PaymentChannelSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type PaymentChannelSession struct {
	Contract     *PaymentChannel   // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// PaymentChannelCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type PaymentChannelCallerSession struct {
	Contract *PaymentChannelCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts         // Call options to use throughout this session
}

// PaymentChannelTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type PaymentChannelTransactorSession struct {
	Contract     *PaymentChannelTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts         // Transaction auth options to use throughout this session
}

// PaymentChannelRaw is an auto generated low-level Go binding around an Ethereum contract.
type PaymentChannelRaw struct {
	Contract *PaymentChannel // Generic contract binding to access the raw methods on
}

// PaymentChannelCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type PaymentChannelCallerRaw struct {
	Contract *PaymentChannelCaller // Generic read-only contract binding to access the raw methods on
}

// PaymentChannelTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type PaymentChannelTransactorRaw struct {
	Contract *PaymentChannelTransactor // Generic write-only contract binding to access the raw methods on
}

// NewPaymentChannel creates a new instance of PaymentChannel, bound to a specific deployed contract.
func NewPaymentChannel(address common.Address, backend bind.ContractBackend) (*PaymentChannel, error) {
	contract, err := bindPaymentChannel(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &PaymentChannel{PaymentChannelCaller: PaymentChannelCaller{contract: contract}, PaymentChannelTransactor: PaymentChannelTransactor{contract: contract}, PaymentChannelFilterer: PaymentChannelFilterer{contract: contract}}, nil
}

// NewPaymentChannelCaller creates a new read-only instance of PaymentChannel, bound to a specific deployed contract.
func NewPaymentChannelCaller(address common.Address, caller bind.ContractCaller) (*PaymentChannelCaller, error) {
	contract, err := bindPaymentChannel(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &PaymentChannelCaller{contract: contract}, nil
}

// NewPaymentChannelTransactor creates a new write-only instance of PaymentChannel, bound to a specific deployed contract.
func NewPaymentChannelTransactor(address common.Address, transactor bind.ContractTransactor) (*PaymentChannelTransactor, error) {
	contract, err := bindPaymentChannel(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &PaymentChannelTransactor{contract: contract}, nil
}

// NewPaymentChannelFilterer creates a new log filterer instance of PaymentChannel, bound to a specific deployed contract.
func NewPaymentChannelFilterer(address common.Address, filterer bind.ContractFilterer) (*PaymentChannelFilterer, error) {
	contract, err := bindPaymentChannel(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &PaymentChannelFilterer{contract: contract}, nil
}

// bindPaymentChannel binds a generic wrapper to an already deployed contract.
func bindPaymentChannel(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := PaymentChannelMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_PaymentChannel *PaymentChannelRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _PaymentChannel.Contract.PaymentChannelCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_PaymentChannel *PaymentChannelRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _PaymentChannel.Contract.PaymentChannelTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_PaymentChannel *PaymentChannelRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _PaymentChannel.Contract.PaymentChannelTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_PaymentChannel *PaymentChannelCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _PaymentChannel.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_PaymentChannel *PaymentChannelTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _PaymentChannel.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_PaymentChannel *PaymentChannelTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _PaymentChannel.Contract.contract.Transact(opts, method, params...)
}

// Channels is a free data retrieval call binding the contract method 0x7a7ebd7b.
//
// Solidity: function channels(bytes32 ) view returns(address user, address provider, address token, uint256 deposit, uint256 expiration, bool open)
func (_PaymentChannel *PaymentChannelCaller) Channels(opts *bind.CallOpts, arg0 [32]byte) (struct {
	User       common.Address
	Provider   common.Address
	Token      common.Address
	Deposit    *big.Int
	Expiration *big.Int
	Open       bool
}, error) {
	var out []interface{}
	err := _PaymentChannel.contract.Call(opts, &out, "channels", arg0)

	outstruct := new(struct {
		User       common.Address
		Provider   common.Address
		Token      common.Address
		Deposit    *big.Int
		Expiration *big.Int
		Open       bool
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.User = *abi.ConvertType(out[0], new(common.Address)).(*common.Address)
	outstruct.Provider = *abi.ConvertType(out[1], new(common.Address)).(*common.Address)
	outstruct.Token = *abi.ConvertType(out[2], new(common.Address)).(*common.Address)
	outstruct.Deposit = *abi.ConvertType(out[3], new(*big.Int)).(**big.Int)
	outstruct.Expiration = *abi.ConvertType(out[4], new(*big.Int)).(**big.Int)
	outstruct.Open = *abi.ConvertType(out[5], new(bool)).(*bool)

	return *outstruct, err

}

// Channels is a free data retrieval call binding the contract method 0x7a7ebd7b.
//
// Solidity: function channels(bytes32 ) view returns(address user, address provider, address token, uint256 deposit, uint256 expiration, bool open)
func (_PaymentChannel *PaymentChannelSession) Channels(arg0 [32]byte) (struct {
	User       common.Address
	Provider   common.Address
	Token      common.Address
	Deposit    *big.Int
	Expiration *big.Int
	Open       bool
}, error) {
	return _PaymentChannel.Contract.Channels(&_PaymentChannel.CallOpts, arg0)
}

// Channels is a free data retrieval call binding the contract method 0x7a7ebd7b.
//
// Solidity: function channels(bytes32 ) view returns(address user, address provider, address token, uint256 deposit, uint256 expiration, bool open)
func (_PaymentChannel *PaymentChannelCallerSession) Channels(arg0 [32]byte) (struct {
	User       common.Address
	Provider   common.Address
	Token      common.Address
	Deposit    *big.Int
	Expiration *big.Int
	Open       bool
}, error) {
	return _PaymentChannel.Contract.Channels(&_PaymentChannel.CallOpts, arg0)
}

// ClaimTimeout is a paid mutator transaction binding the contract method 0xaaa5a02a.
//
// Solidity: function claimTimeout(bytes32 channelId) returns()
func (_PaymentChannel *PaymentChannelTransactor) ClaimTimeout(opts *bind.TransactOpts, channelId [32]byte) (*types.Transaction, error) {
	return _PaymentChannel.contract.Transact(opts, "claimTimeout", channelId)
}

// ClaimTimeout is a paid mutator transaction binding the contract method 0xaaa5a02a.
//
// Solidity: function claimTimeout(bytes32 channelId) returns()
func (_PaymentChannel *PaymentChannelSession) ClaimTimeout(channelId [32]byte) (*types.Transaction, error) {
	return _PaymentChannel.Contract.ClaimTimeout(&_PaymentChannel.TransactOpts, channelId)
}

// ClaimTimeout is a paid mutator transaction binding the contract method 0xaaa5a02a.
//
// Solidity: function claimTimeout(bytes32 channelId) returns()
func (_PaymentChannel *PaymentChannelTransactorSession) ClaimTimeout(channelId [32]byte) (*types.Transaction, error) {
	return _PaymentChannel.Contract.ClaimTimeout(&_PaymentChannel.TransactOpts, channelId)
}

// CloseChannel is a paid mutator transaction binding the contract method 0xbac068ce.
//
// Solidity: function closeChannel(bytes32 channelId, uint256 amount, bytes signature) returns()
func (_PaymentChannel *PaymentChannelTransactor) CloseChannel(opts *bind.TransactOpts, channelId [32]byte, amount *big.Int, signature []byte) (*types.Transaction, error) {
	return _PaymentChannel.contract.Transact(opts, "closeChannel", channelId, amount, signature)
}

// CloseChannel is a paid mutator transaction binding the contract method 0xbac068ce.
//
// Solidity: function closeChannel(bytes32 channelId, uint256 amount, bytes signature) returns()
func (_PaymentChannel *PaymentChannelSession) CloseChannel(channelId [32]byte, amount *big.Int, signature []byte) (*types.Transaction, error) {
	return _PaymentChannel.Contract.CloseChannel(&_PaymentChannel.TransactOpts, channelId, amount, signature)
}

// CloseChannel is a paid mutator transaction binding the contract method 0xbac068ce.
//
// Solidity: function closeChannel(bytes32 channelId, uint256 amount, bytes signature) returns()
func (_PaymentChannel *PaymentChannelTransactorSession) CloseChannel(channelId [32]byte, amount *big.Int, signature []byte) (*types.Transaction, error) {
	return _PaymentChannel.Contract.CloseChannel(&_PaymentChannel.TransactOpts, channelId, amount, signature)
}

// OpenChannel is a paid mutator transaction binding the contract method 0xd8d9965a.
//
// Solidity: function openChannel(address provider, address token, uint256 deposit, uint256 duration) returns(bytes32 channelId)
func (_PaymentChannel *PaymentChannelTransactor) OpenChannel(opts *bind.TransactOpts, provider common.Address, token common.Address, deposit *big.Int, duration *big.Int) (*types.Transaction, error) {
	return _PaymentChannel.contract.Transact(opts, "openChannel", provider, token, deposit, duration)
}

// OpenChannel is a paid mutator transaction binding the contract method 0xd8d9965a.
//
// Solidity: function openChannel(address provider, address token, uint256 deposit, uint256 duration) returns(bytes32 channelId)
func (_PaymentChannel *PaymentChannelSession) OpenChannel(provider common.Address, token common.Address, deposit *big.Int, duration *big.Int) (*types.Transaction, error) {
	return _PaymentChannel.Contract.OpenChannel(&_PaymentChannel.TransactOpts, provider, token, deposit, duration)
}

// OpenChannel is a paid mutator transaction binding the contract method 0xd8d9965a.
//
// Solidity: function openChannel(address provider, address token, uint256 deposit, uint256 duration) returns(bytes32 channelId)
func (_PaymentChannel *PaymentChannelTransactorSession) OpenChannel(provider common.Address, token common.Address, deposit *big.Int, duration *big.Int) (*types.Transaction, error) {
	return _PaymentChannel.Contract.OpenChannel(&_PaymentChannel.TransactOpts, provider, token, deposit, duration)
}

// PaymentChannelChannelClosedIterator is returned from FilterChannelClosed and is used to iterate over the raw logs and unpacked data for ChannelClosed events raised by the PaymentChannel contract.
type PaymentChannelChannelClosedIterator struct {
	Event *PaymentChannelChannelClosed // Event containing the contract specifics and raw log

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
func (it *PaymentChannelChannelClosedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(PaymentChannelChannelClosed)
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
		it.Event = new(PaymentChannelChannelClosed)
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
func (it *PaymentChannelChannelClosedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *PaymentChannelChannelClosedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// PaymentChannelChannelClosed represents a ChannelClosed event raised by the PaymentChannel contract.
type PaymentChannelChannelClosed struct {
	ChannelId      [32]byte
	AmountReceived *big.Int
	Raw            types.Log // Blockchain specific contextual infos
}

// FilterChannelClosed is a free log retrieval operation binding the contract event 0x74e9aa18d6bb2c4887e76896296ce0a296a2e8315bb319b08b7607ff92fbef79.
//
// Solidity: event ChannelClosed(bytes32 indexed channelId, uint256 amountReceived)
func (_PaymentChannel *PaymentChannelFilterer) FilterChannelClosed(opts *bind.FilterOpts, channelId [][32]byte) (*PaymentChannelChannelClosedIterator, error) {

	var channelIdRule []interface{}
	for _, channelIdItem := range channelId {
		channelIdRule = append(channelIdRule, channelIdItem)
	}

	logs, sub, err := _PaymentChannel.contract.FilterLogs(opts, "ChannelClosed", channelIdRule)
	if err != nil {
		return nil, err
	}
	return &PaymentChannelChannelClosedIterator{contract: _PaymentChannel.contract, event: "ChannelClosed", logs: logs, sub: sub}, nil
}

// WatchChannelClosed is a free log subscription operation binding the contract event 0x74e9aa18d6bb2c4887e76896296ce0a296a2e8315bb319b08b7607ff92fbef79.
//
// Solidity: event ChannelClosed(bytes32 indexed channelId, uint256 amountReceived)
func (_PaymentChannel *PaymentChannelFilterer) WatchChannelClosed(opts *bind.WatchOpts, sink chan<- *PaymentChannelChannelClosed, channelId [][32]byte) (event.Subscription, error) {

	var channelIdRule []interface{}
	for _, channelIdItem := range channelId {
		channelIdRule = append(channelIdRule, channelIdItem)
	}

	logs, sub, err := _PaymentChannel.contract.WatchLogs(opts, "ChannelClosed", channelIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(PaymentChannelChannelClosed)
				if err := _PaymentChannel.contract.UnpackLog(event, "ChannelClosed", log); err != nil {
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

// ParseChannelClosed is a log parse operation binding the contract event 0x74e9aa18d6bb2c4887e76896296ce0a296a2e8315bb319b08b7607ff92fbef79.
//
// Solidity: event ChannelClosed(bytes32 indexed channelId, uint256 amountReceived)
func (_PaymentChannel *PaymentChannelFilterer) ParseChannelClosed(log types.Log) (*PaymentChannelChannelClosed, error) {
	event := new(PaymentChannelChannelClosed)
	if err := _PaymentChannel.contract.UnpackLog(event, "ChannelClosed", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// PaymentChannelChannelExpiredIterator is returned from FilterChannelExpired and is used to iterate over the raw logs and unpacked data for ChannelExpired events raised by the PaymentChannel contract.
type PaymentChannelChannelExpiredIterator struct {
	Event *PaymentChannelChannelExpired // Event containing the contract specifics and raw log

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
func (it *PaymentChannelChannelExpiredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(PaymentChannelChannelExpired)
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
		it.Event = new(PaymentChannelChannelExpired)
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
func (it *PaymentChannelChannelExpiredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *PaymentChannelChannelExpiredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// PaymentChannelChannelExpired represents a ChannelExpired event raised by the PaymentChannel contract.
type PaymentChannelChannelExpired struct {
	ChannelId [32]byte
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterChannelExpired is a free log retrieval operation binding the contract event 0xb0b1a2376e938b887ac88a6049e44d46d0042dd3d17f70089e61339792cb2fbe.
//
// Solidity: event ChannelExpired(bytes32 indexed channelId)
func (_PaymentChannel *PaymentChannelFilterer) FilterChannelExpired(opts *bind.FilterOpts, channelId [][32]byte) (*PaymentChannelChannelExpiredIterator, error) {

	var channelIdRule []interface{}
	for _, channelIdItem := range channelId {
		channelIdRule = append(channelIdRule, channelIdItem)
	}

	logs, sub, err := _PaymentChannel.contract.FilterLogs(opts, "ChannelExpired", channelIdRule)
	if err != nil {
		return nil, err
	}
	return &PaymentChannelChannelExpiredIterator{contract: _PaymentChannel.contract, event: "ChannelExpired", logs: logs, sub: sub}, nil
}

// WatchChannelExpired is a free log subscription operation binding the contract event 0xb0b1a2376e938b887ac88a6049e44d46d0042dd3d17f70089e61339792cb2fbe.
//
// Solidity: event ChannelExpired(bytes32 indexed channelId)
func (_PaymentChannel *PaymentChannelFilterer) WatchChannelExpired(opts *bind.WatchOpts, sink chan<- *PaymentChannelChannelExpired, channelId [][32]byte) (event.Subscription, error) {

	var channelIdRule []interface{}
	for _, channelIdItem := range channelId {
		channelIdRule = append(channelIdRule, channelIdItem)
	}

	logs, sub, err := _PaymentChannel.contract.WatchLogs(opts, "ChannelExpired", channelIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(PaymentChannelChannelExpired)
				if err := _PaymentChannel.contract.UnpackLog(event, "ChannelExpired", log); err != nil {
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

// ParseChannelExpired is a log parse operation binding the contract event 0xb0b1a2376e938b887ac88a6049e44d46d0042dd3d17f70089e61339792cb2fbe.
//
// Solidity: event ChannelExpired(bytes32 indexed channelId)
func (_PaymentChannel *PaymentChannelFilterer) ParseChannelExpired(log types.Log) (*PaymentChannelChannelExpired, error) {
	event := new(PaymentChannelChannelExpired)
	if err := _PaymentChannel.contract.UnpackLog(event, "ChannelExpired", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// PaymentChannelChannelOpenedIterator is returned from FilterChannelOpened and is used to iterate over the raw logs and unpacked data for ChannelOpened events raised by the PaymentChannel contract.
type PaymentChannelChannelOpenedIterator struct {
	Event *PaymentChannelChannelOpened // Event containing the contract specifics and raw log

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
func (it *PaymentChannelChannelOpenedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(PaymentChannelChannelOpened)
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
		it.Event = new(PaymentChannelChannelOpened)
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
func (it *PaymentChannelChannelOpenedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *PaymentChannelChannelOpenedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// PaymentChannelChannelOpened represents a ChannelOpened event raised by the PaymentChannel contract.
type PaymentChannelChannelOpened struct {
	ChannelId  [32]byte
	User       common.Address
	Provider   common.Address
	Deposit    *big.Int
	Expiration *big.Int
	Raw        types.Log // Blockchain specific contextual infos
}

// FilterChannelOpened is a free log retrieval operation binding the contract event 0x506f81b7a67b45bfbc6167fd087b3dd9b65b4531a2380ec406aab5b57ac62152.
//
// Solidity: event ChannelOpened(bytes32 indexed channelId, address indexed user, address indexed provider, uint256 deposit, uint256 expiration)
func (_PaymentChannel *PaymentChannelFilterer) FilterChannelOpened(opts *bind.FilterOpts, channelId [][32]byte, user []common.Address, provider []common.Address) (*PaymentChannelChannelOpenedIterator, error) {

	var channelIdRule []interface{}
	for _, channelIdItem := range channelId {
		channelIdRule = append(channelIdRule, channelIdItem)
	}
	var userRule []interface{}
	for _, userItem := range user {
		userRule = append(userRule, userItem)
	}
	var providerRule []interface{}
	for _, providerItem := range provider {
		providerRule = append(providerRule, providerItem)
	}

	logs, sub, err := _PaymentChannel.contract.FilterLogs(opts, "ChannelOpened", channelIdRule, userRule, providerRule)
	if err != nil {
		return nil, err
	}
	return &PaymentChannelChannelOpenedIterator{contract: _PaymentChannel.contract, event: "ChannelOpened", logs: logs, sub: sub}, nil
}

// WatchChannelOpened is a free log subscription operation binding the contract event 0x506f81b7a67b45bfbc6167fd087b3dd9b65b4531a2380ec406aab5b57ac62152.
//
// Solidity: event ChannelOpened(bytes32 indexed channelId, address indexed user, address indexed provider, uint256 deposit, uint256 expiration)
func (_PaymentChannel *PaymentChannelFilterer) WatchChannelOpened(opts *bind.WatchOpts, sink chan<- *PaymentChannelChannelOpened, channelId [][32]byte, user []common.Address, provider []common.Address) (event.Subscription, error) {

	var channelIdRule []interface{}
	for _, channelIdItem := range channelId {
		channelIdRule = append(channelIdRule, channelIdItem)
	}
	var userRule []interface{}
	for _, userItem := range user {
		userRule = append(userRule, userItem)
	}
	var providerRule []interface{}
	for _, providerItem := range provider {
		providerRule = append(providerRule, providerItem)
	}

	logs, sub, err := _PaymentChannel.contract.WatchLogs(opts, "ChannelOpened", channelIdRule, userRule, providerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(PaymentChannelChannelOpened)
				if err := _PaymentChannel.contract.UnpackLog(event, "ChannelOpened", log); err != nil {
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

// ParseChannelOpened is a log parse operation binding the contract event 0x506f81b7a67b45bfbc6167fd087b3dd9b65b4531a2380ec406aab5b57ac62152.
//
// Solidity: event ChannelOpened(bytes32 indexed channelId, address indexed user, address indexed provider, uint256 deposit, uint256 expiration)
func (_PaymentChannel *PaymentChannelFilterer) ParseChannelOpened(log types.Log) (*PaymentChannelChannelOpened, error) {
	event := new(PaymentChannelChannelOpened)
	if err := _PaymentChannel.contract.UnpackLog(event, "ChannelOpened", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
