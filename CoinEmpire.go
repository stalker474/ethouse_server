// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package main

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

// Reference imports to suppress errors if they are not otherwise used.
var (
	_ = big.NewInt
	_ = strings.NewReader
	_ = ethereum.NotFound
	_ = abi.U256
	_ = bind.Bind
	_ = common.Big1
	_ = types.BloomLookup
	_ = event.NewSubscription
)

// BufferABI is the input ABI used to generate the binding from.
const BufferABI = "[]"

// BufferBin is the compiled bytecode used for deploying new contracts.
const BufferBin = `0x604c602c600b82828239805160001a60731460008114601c57601e565bfe5b5030600052607381538281f3fe73000000000000000000000000000000000000000030146080604052600080fdfea165627a7a7230582078280ca2228dd0fb84231a4ca11ec56816a1152a4f5e2913545876b991bf2fca0029`

// DeployBuffer deploys a new Ethereum contract, binding an instance of Buffer to it.
func DeployBuffer(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *Buffer, error) {
	parsed, err := abi.JSON(strings.NewReader(BufferABI))
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	address, tx, contract, err := bind.DeployContract(auth, parsed, common.FromHex(BufferBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &Buffer{BufferCaller: BufferCaller{contract: contract}, BufferTransactor: BufferTransactor{contract: contract}, BufferFilterer: BufferFilterer{contract: contract}}, nil
}

// Buffer is an auto generated Go binding around an Ethereum contract.
type Buffer struct {
	BufferCaller     // Read-only binding to the contract
	BufferTransactor // Write-only binding to the contract
	BufferFilterer   // Log filterer for contract events
}

// BufferCaller is an auto generated read-only Go binding around an Ethereum contract.
type BufferCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// BufferTransactor is an auto generated write-only Go binding around an Ethereum contract.
type BufferTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// BufferFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type BufferFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// BufferSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type BufferSession struct {
	Contract     *Buffer           // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// BufferCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type BufferCallerSession struct {
	Contract *BufferCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts // Call options to use throughout this session
}

// BufferTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type BufferTransactorSession struct {
	Contract     *BufferTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// BufferRaw is an auto generated low-level Go binding around an Ethereum contract.
type BufferRaw struct {
	Contract *Buffer // Generic contract binding to access the raw methods on
}

// BufferCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type BufferCallerRaw struct {
	Contract *BufferCaller // Generic read-only contract binding to access the raw methods on
}

// BufferTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type BufferTransactorRaw struct {
	Contract *BufferTransactor // Generic write-only contract binding to access the raw methods on
}

// NewBuffer creates a new instance of Buffer, bound to a specific deployed contract.
func NewBuffer(address common.Address, backend bind.ContractBackend) (*Buffer, error) {
	contract, err := bindBuffer(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Buffer{BufferCaller: BufferCaller{contract: contract}, BufferTransactor: BufferTransactor{contract: contract}, BufferFilterer: BufferFilterer{contract: contract}}, nil
}

// NewBufferCaller creates a new read-only instance of Buffer, bound to a specific deployed contract.
func NewBufferCaller(address common.Address, caller bind.ContractCaller) (*BufferCaller, error) {
	contract, err := bindBuffer(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &BufferCaller{contract: contract}, nil
}

// NewBufferTransactor creates a new write-only instance of Buffer, bound to a specific deployed contract.
func NewBufferTransactor(address common.Address, transactor bind.ContractTransactor) (*BufferTransactor, error) {
	contract, err := bindBuffer(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &BufferTransactor{contract: contract}, nil
}

// NewBufferFilterer creates a new log filterer instance of Buffer, bound to a specific deployed contract.
func NewBufferFilterer(address common.Address, filterer bind.ContractFilterer) (*BufferFilterer, error) {
	contract, err := bindBuffer(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &BufferFilterer{contract: contract}, nil
}

// bindBuffer binds a generic wrapper to an already deployed contract.
func bindBuffer(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(BufferABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Buffer *BufferRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _Buffer.Contract.BufferCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Buffer *BufferRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Buffer.Contract.BufferTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Buffer *BufferRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Buffer.Contract.BufferTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Buffer *BufferCallerRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _Buffer.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Buffer *BufferTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Buffer.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Buffer *BufferTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Buffer.Contract.contract.Transact(opts, method, params...)
}

// CBORABI is the input ABI used to generate the binding from.
const CBORABI = "[]"

// CBORBin is the compiled bytecode used for deploying new contracts.
const CBORBin = `0x604c602c600b82828239805160001a60731460008114601c57601e565bfe5b5030600052607381538281f3fe73000000000000000000000000000000000000000030146080604052600080fdfea165627a7a72305820bba8c4f72177663ad5f6b30d4d388c5b46fa9c1518b8c2dca36e9b5e9d49c6a20029`

// DeployCBOR deploys a new Ethereum contract, binding an instance of CBOR to it.
func DeployCBOR(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *CBOR, error) {
	parsed, err := abi.JSON(strings.NewReader(CBORABI))
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	address, tx, contract, err := bind.DeployContract(auth, parsed, common.FromHex(CBORBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &CBOR{CBORCaller: CBORCaller{contract: contract}, CBORTransactor: CBORTransactor{contract: contract}, CBORFilterer: CBORFilterer{contract: contract}}, nil
}

// CBOR is an auto generated Go binding around an Ethereum contract.
type CBOR struct {
	CBORCaller     // Read-only binding to the contract
	CBORTransactor // Write-only binding to the contract
	CBORFilterer   // Log filterer for contract events
}

// CBORCaller is an auto generated read-only Go binding around an Ethereum contract.
type CBORCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// CBORTransactor is an auto generated write-only Go binding around an Ethereum contract.
type CBORTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// CBORFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type CBORFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// CBORSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type CBORSession struct {
	Contract     *CBOR             // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// CBORCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type CBORCallerSession struct {
	Contract *CBORCaller   // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts // Call options to use throughout this session
}

// CBORTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type CBORTransactorSession struct {
	Contract     *CBORTransactor   // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// CBORRaw is an auto generated low-level Go binding around an Ethereum contract.
type CBORRaw struct {
	Contract *CBOR // Generic contract binding to access the raw methods on
}

// CBORCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type CBORCallerRaw struct {
	Contract *CBORCaller // Generic read-only contract binding to access the raw methods on
}

// CBORTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type CBORTransactorRaw struct {
	Contract *CBORTransactor // Generic write-only contract binding to access the raw methods on
}

// NewCBOR creates a new instance of CBOR, bound to a specific deployed contract.
func NewCBOR(address common.Address, backend bind.ContractBackend) (*CBOR, error) {
	contract, err := bindCBOR(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &CBOR{CBORCaller: CBORCaller{contract: contract}, CBORTransactor: CBORTransactor{contract: contract}, CBORFilterer: CBORFilterer{contract: contract}}, nil
}

// NewCBORCaller creates a new read-only instance of CBOR, bound to a specific deployed contract.
func NewCBORCaller(address common.Address, caller bind.ContractCaller) (*CBORCaller, error) {
	contract, err := bindCBOR(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &CBORCaller{contract: contract}, nil
}

// NewCBORTransactor creates a new write-only instance of CBOR, bound to a specific deployed contract.
func NewCBORTransactor(address common.Address, transactor bind.ContractTransactor) (*CBORTransactor, error) {
	contract, err := bindCBOR(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &CBORTransactor{contract: contract}, nil
}

// NewCBORFilterer creates a new log filterer instance of CBOR, bound to a specific deployed contract.
func NewCBORFilterer(address common.Address, filterer bind.ContractFilterer) (*CBORFilterer, error) {
	contract, err := bindCBOR(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &CBORFilterer{contract: contract}, nil
}

// bindCBOR binds a generic wrapper to an already deployed contract.
func bindCBOR(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(CBORABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_CBOR *CBORRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _CBOR.Contract.CBORCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_CBOR *CBORRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _CBOR.Contract.CBORTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_CBOR *CBORRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _CBOR.Contract.CBORTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_CBOR *CBORCallerRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _CBOR.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_CBOR *CBORTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _CBOR.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_CBOR *CBORTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _CBOR.Contract.contract.Transact(opts, method, params...)
}

// CoinEmpireABI is the input ABI used to generate the binding from.
const CoinEmpireABI = "[{\"constant\":true,\"inputs\":[{\"name\":\"price\",\"type\":\"uint256\"},{\"name\":\"tier\",\"type\":\"uint8\"}],\"name\":\"getSlotId\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"pure\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"pool\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"slot_to_owner\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[],\"name\":\"newSample\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"withdrawHouse\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_myid\",\"type\":\"bytes32\"},{\"name\":\"_result\",\"type\":\"string\"}],\"name\":\"__callback\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"price\",\"type\":\"uint256\"},{\"name\":\"tier\",\"type\":\"uint8\"}],\"name\":\"getTierData\",\"outputs\":[{\"name\":\"_index_data\",\"type\":\"uint256[10]\"},{\"name\":\"_earnings_data\",\"type\":\"uint256[10]\"},{\"name\":\"_price_data\",\"type\":\"uint256[10]\"},{\"name\":\"_owner_data\",\"type\":\"address[10]\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"\",\"type\":\"address\"}],\"name\":\"resell_tickets\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"prices\",\"type\":\"uint256[]\"},{\"name\":\"tiers\",\"type\":\"uint8[]\"}],\"name\":\"sellSlots\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_queryId\",\"type\":\"bytes32\"},{\"name\":\"_result\",\"type\":\"string\"},{\"name\":\"_proof\",\"type\":\"bytes\"}],\"name\":\"__callback\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[],\"name\":\"unpause\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"account\",\"type\":\"address\"}],\"name\":\"isPauser\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"prices\",\"type\":\"uint256[]\"},{\"name\":\"tiers\",\"type\":\"uint8[]\"}],\"name\":\"buySlots\",\"outputs\":[],\"payable\":true,\"stateMutability\":\"payable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"\",\"type\":\"address\"}],\"name\":\"capitalOf\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"\",\"type\":\"address\"}],\"name\":\"profitOf\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"paused\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[],\"name\":\"renouncePauser\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"\",\"type\":\"address\"}],\"name\":\"balanceOf\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[],\"name\":\"renounceOwnership\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"slot_to_price\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"account\",\"type\":\"address\"}],\"name\":\"addPauser\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[],\"name\":\"destroy\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[],\"name\":\"pause\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"query_string\",\"outputs\":[{\"name\":\"\",\"type\":\"string\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"isOwner\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[],\"name\":\"checkPrice\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"slot_to_earnings\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"getMinMax\",\"outputs\":[{\"name\":\"_min\",\"type\":\"uint256\"},{\"name\":\"_max\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"new_destination\",\"type\":\"address\"}],\"name\":\"setCutDestination\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"current_price\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"name\",\"type\":\"bytes32\"},{\"name\":\"new_value\",\"type\":\"uint256\"}],\"name\":\"setConfigValue\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"getTier1Data\",\"outputs\":[{\"name\":\"_index_data\",\"type\":\"uint256[60]\"},{\"name\":\"_earnings_data\",\"type\":\"uint256[60]\"},{\"name\":\"_price_data\",\"type\":\"uint256[60]\"},{\"name\":\"_owner_data\",\"type\":\"address[60]\"},{\"name\":\"_length\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"withdrawWallet\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"latest_sample\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"\",\"type\":\"bytes32\"}],\"name\":\"config\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"price\",\"type\":\"uint256\"}],\"name\":\"getHotnessModifier\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"transferOwnership\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"house\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"payable\":true,\"stateMutability\":\"payable\",\"type\":\"constructor\"},{\"payable\":true,\"stateMutability\":\"payable\",\"type\":\"fallback\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"name\":\"query_id\",\"type\":\"bytes32\"}],\"name\":\"SamplingPriceStarted\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"name\":\"price\",\"type\":\"uint256\"},{\"indexed\":false,\"name\":\"result\",\"type\":\"string\"},{\"indexed\":false,\"name\":\"proof\",\"type\":\"bytes\"}],\"name\":\"SamplingPriceEnded\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"name\":\"price\",\"type\":\"uint256\"},{\"indexed\":false,\"name\":\"new_value\",\"type\":\"uint256\"},{\"indexed\":false,\"name\":\"tier\",\"type\":\"uint8\"},{\"indexed\":false,\"name\":\"from\",\"type\":\"address\"},{\"indexed\":false,\"name\":\"by\",\"type\":\"address\"}],\"name\":\"SlotPurchased\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"name\":\"price\",\"type\":\"uint256\"},{\"indexed\":false,\"name\":\"tier\",\"type\":\"uint8\"},{\"indexed\":false,\"name\":\"by\",\"type\":\"address\"}],\"name\":\"SlotAbandoned\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"name\":\"name\",\"type\":\"bytes32\"},{\"indexed\":false,\"name\":\"new_value\",\"type\":\"uint256\"}],\"name\":\"ConfigChanged\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"previousOwner\",\"type\":\"address\"},{\"indexed\":true,\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"OwnershipTransferred\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"name\":\"account\",\"type\":\"address\"}],\"name\":\"Paused\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"name\":\"account\",\"type\":\"address\"}],\"name\":\"Unpaused\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"account\",\"type\":\"address\"}],\"name\":\"PauserAdded\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"account\",\"type\":\"address\"}],\"name\":\"PauserRemoved\",\"type\":\"event\"}]"

// CoinEmpireBin is the compiled bytecode used for deploying new contracts.
const CoinEmpireBin = `0x608060405260088054600160a060020a03191673a54741f7fe21689b59bd7eacbf3a2947cd3f3bd417905560006009819055600a5562000048336401000000006200036b810204565b60068054600160a860020a0319166101003381029190911791829055604051600160a060020a039190920416906000907f8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0908290a360076020526203827060008051602062005554833981519152556404a817c8007f1f8f10ace8067951adf4dd30e7f9fb4ea4d21b6af6300a20ec0d74b91f84667a819055610384600080516020620056bf8339815191525560327f7729a41062816f77329ca956209d1a1aa7b7e1abd27f215c57ad3a356065d005556101f47f1934936e913495172a68f489eba8f04b1b844c8d38a3035a45b08189f9af57cd556113887facc1313bfd621232066a9c54ce789299c5df44b1ef9af83db841e307b10705f4556706f05b59d3b200007f94d03aac1c5adffb43bea72d0d9572fc98b730fe4f355bdda67a1c5de5d6c6775567011c37937e0800007f161aa644b3561c0a2ed9ab0eb5ebcb59d229bba8d49af22be9d0922f604d1f2955662386f26fc100007fae6e034939092767869a45875d60d5a3f30e8ca6e83ecd66c948bce2cdf3c70e556216e3607f03565cf03431b9675ce6f670301da7655569d596d084d861c90455dcee47bff5556207a1207fc74b37efb558a657abc1a4c0bb36f4e7aa6d90bc016115585b63ee7bdb5965b155620186a07f93c763557f8d2895a7b0d4126cd48a84b0acc0bbab8ef8a3d65ce52fe773f03055620f42407fea096be5dc82e6eb331c8393f48594ecb147108039ee44932379ca150d6fe15355620493e07f1893b97643a98089770f715d0305eea450b78da77cb2504a1c69ed2b77b740ad5561c3507f9428f94ef6f4ec3c650bfb4a2a71424d241e94b5c309003957bb775f67e0c432557f4741535f50524943450000000000000000000000000000000000000000000000600052620002fe90640100000000620003bd810204565b620003327f1100000000000000000000000000000000000000000000000000000000000000640100000000620005fa810204565b600c546200034f903464010000000062000ce46200083c82021704565b600c556200036564010000000062000856810204565b620013cb565b6200038660058264010000000062003b1062000a2d82021704565b604051600160a060020a038216907f6719d08c1888103bea251a4ed56406bd0c3e69723c8a1686e017e7bbe159b6f890600090a250565b600154600160a060020a03161580620003f35750600154620003f190600160a060020a031664010000000062000a88810204565b155b1562000410576200040e600064010000000062000a8c810204565b505b600160009054906101000a9004600160a060020a0316600160a060020a03166338cc48316040518163ffffffff167c0100000000000000000000000000000000000000000000000000000000028152600401602060405180830381600087803b1580156200047d57600080fd5b505af115801562000492573d6000803e3d6000fd5b505050506040513d6020811015620004a957600080fd5b5051600054600160a060020a039081169116146200057c57600160009054906101000a9004600160a060020a0316600160a060020a03166338cc48316040518163ffffffff167c0100000000000000000000000000000000000000000000000000000000028152600401602060405180830381600087803b1580156200052e57600080fd5b505af115801562000543573d6000803e3d6000fd5b505050506040513d60208110156200055a57600080fd5b505160008054600160a060020a031916600160a060020a039092169190911790555b60008054604080517fca6ad1e4000000000000000000000000000000000000000000000000000000008152600481018590529051600160a060020a039092169263ca6ad1e49260248084019382900301818387803b158015620005de57600080fd5b505af1158015620005f3573d6000803e3d6000fd5b5050505050565b600154600160a060020a031615806200063057506001546200062e90600160a060020a031664010000000062000a88810204565b155b156200064d576200064b600064010000000062000a8c810204565b505b600160009054906101000a9004600160a060020a0316600160a060020a03166338cc48316040518163ffffffff167c0100000000000000000000000000000000000000000000000000000000028152600401602060405180830381600087803b158015620006ba57600080fd5b505af1158015620006cf573d6000803e3d6000fd5b505050506040513d6020811015620006e657600080fd5b5051600054600160a060020a03908116911614620007b957600160009054906101000a9004600160a060020a0316600160a060020a03166338cc48316040518163ffffffff167c0100000000000000000000000000000000000000000000000000000000028152600401602060405180830381600087803b1580156200076b57600080fd5b505af115801562000780573d6000803e3d6000fd5b505050506040513d60208110156200079757600080fd5b505160008054600160a060020a031916600160a060020a039092169190911790555b60008054604080517f688dcfd70000000000000000000000000000000000000000000000000000000081527fff00000000000000000000000000000000000000000000000000000000000000851660048201529051600160a060020a039092169263688dcfd79260248084019382900301818387803b158015620005de57600080fd5b6000828201838110156200084f57600080fd5b9392505050565b60065460ff16156200086757600080fd5b7f5052494345434845434b5f44454c4159000000000000000000000000000000006000526007602052600080516020620056bf8339815191525460095442910111156200091557604080517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152600c60248201527f636f6f6c696e6720646f776e0000000000000000000000000000000000000000604482015290519081900360640190fd5b7f5052494345434845434b5f44454c41590000000000000000000000000000000060005260076020908152600080516020620056bf83398151915254604080518082018252600681527f6e65737465640000000000000000000000000000000000000000000000000000818501528151610180810190925261014b808352303194620009ef9493919062005574908301397f4741535f4c494d495400000000000000000000000000000000000000000000006000526007602052600080516020620055548339815191525464010000000062000aa7810204565b5062000a2762000a0f8230316401000000006200295262000f1b82021704565b600c54906401000000006200295262000f1b82021704565b600c5550565b600160a060020a038116151562000a4357600080fd5b62000a58828264010000000062000f31810204565b1562000a6357600080fd5b600160a060020a0316600090815260209190915260409020805460ff19166001179055565b3b90565b600062000aa164010000000062000f69810204565b92915050565b600154600090600160a060020a0316158062000ae0575060015462000ade90600160a060020a031664010000000062000a88810204565b155b1562000afd5762000afb600064010000000062000a8c810204565b505b600160009054906101000a9004600160a060020a0316600160a060020a03166338cc48316040518163ffffffff167c0100000000000000000000000000000000000000000000000000000000028152600401602060405180830381600087803b15801562000b6a57600080fd5b505af115801562000b7f573d6000803e3d6000fd5b505050506040513d602081101562000b9657600080fd5b5051600054600160a060020a0390811691161462000c6957600160009054906101000a9004600160a060020a0316600160a060020a03166338cc48316040518163ffffffff167c0100000000000000000000000000000000000000000000000000000000028152600401602060405180830381600087803b15801562000c1b57600080fd5b505af115801562000c30573d6000803e3d6000fd5b505050506040513d602081101562000c4757600080fd5b505160008054600160a060020a031916600160a060020a039092169190911790555b60008054604080517f2ef3accc0000000000000000000000000000000000000000000000000000000081526024810186905260048101918252875160448201528751600160a060020a0390931692632ef3accc928992889282916064019060208601908083838c5b8381101562000ceb57818101518382015260200162000cd1565b50505050905090810190601f16801562000d195780820380516001836020036101000a031916815260200191505b509350505050602060405180830381600087803b15801562000d3a57600080fd5b505af115801562000d4f573d6000803e3d6000fd5b505050506040513d602081101562000d6657600080fd5b50519050670de0b6b3a76400003a84020181111562000d8a57506000905062000f13565b6000809054906101000a9004600160a060020a0316600160a060020a031663c51be90f82888888886040518663ffffffff167c0100000000000000000000000000000000000000000000000000000000028152600401808581526020018060200180602001848152602001838103835286818151815260200191508051906020019080838360005b8381101562000e2c57818101518382015260200162000e12565b50505050905090810190601f16801562000e5a5780820380516001836020036101000a031916815260200191505b50838103825285518152855160209182019187019080838360005b8381101562000e8f57818101518382015260200162000e75565b50505050905090810190601f16801562000ebd5780820380516001836020036101000a031916815260200191505b5096505050505050506020604051808303818588803b15801562000ee057600080fd5b505af115801562000ef5573d6000803e3d6000fd5b50505050506040513d602081101562000f0d57600080fd5b50519150505b949350505050565b60008282111562000f2b57600080fd5b50900390565b6000600160a060020a038216151562000f4957600080fd5b50600160a060020a03166000908152602091909152604090205460ff1690565b60008062000f94731d3b2638a7cc9f2cb3d298a3da7a90b67e5506ed64010000000062000a88810204565b1115620010125760018054600160a060020a031916731d3b2638a7cc9f2cb3d298a3da7a90b67e5506ed17905560408051808201909152600b81527f6574685f6d61696e6e65740000000000000000000000000000000000000000006020820152620010099064010000000062001310810204565b5060016200130d565b60006200103c73c03a2615d5efaf5f49f60b7bb6583eaec212fdf164010000000062000a88810204565b1115620010b15760018054600160a060020a03191673c03a2615d5efaf5f49f60b7bb6583eaec212fdf117905560408051808201909152600c81527f6574685f726f707374656e3300000000000000000000000000000000000000006020820152620010099064010000000062001310810204565b6000620010db73b7a07bcf2ba2f2703b24c0691b5278999c59ac7e64010000000062000a88810204565b1115620011505760018054600160a060020a03191673b7a07bcf2ba2f2703b24c0691b5278999c59ac7e17905560408051808201909152600981527f6574685f6b6f76616e00000000000000000000000000000000000000000000006020820152620010099064010000000062001310810204565b60006200117a73146500cfd35b22e4a392fe0adc06de1a1368ed4864010000000062000a88810204565b1115620011ef5760018054600160a060020a03191673146500cfd35b22e4a392fe0adc06de1a1368ed4817905560408051808201909152600b81527f6574685f72696e6b6562790000000000000000000000000000000000000000006020820152620010099064010000000062001310810204565b600062001219736f485c8bf6fc43ea212e93bbf8ce046c7f1cb47564010000000062000a88810204565b11156200124d575060018054600160a060020a031916736f485c8bf6fc43ea212e93bbf8ce046c7f1cb4751781556200130d565b6000620012777320e12a1f859b3feae5fb2a0a32c18f5a65555bbf64010000000062000a88810204565b1115620012ab575060018054600160a060020a0319167320e12a1f859b3feae5fb2a0a32c18f5a65555bbf1781556200130d565b6000620012d57351efaf4c8b3c9afbd5ab9f4bbc82784ab6ef8faa64010000000062000a88810204565b111562001309575060018054600160a060020a0319167351efaf4c8b3c9afbd5ab9f4bbc82784ab6ef8faa1781556200130d565b5060005b90565b80516200132590600290602084019062001329565b5050565b828054600181600116156101000203166002900490600052602060002090601f016020900481019282601f106200136c57805160ff19168380011785556200139c565b828001600101855582156200139c579182015b828111156200139c5782518255916020019190600101906200137f565b50620013aa929150620013ae565b5090565b6200130d91905b80821115620013aa5760008155600101620013b5565b61417980620013db6000396000f3fe6080604052600436106102345760003560e060020a900480637c9d452711610132578063aad74722116100af578063c595d6b911610073578063c595d6b914610c33578063cc718f7614610c48578063ddbb73b714610c72578063f2fde38b14610c9c578063ff9b3acf14610ccf57610234565b8063aad7472214610ab8578063ab2b313714610aeb578063b7671acd14610b00578063b93298f214610b30578063be87ab3f14610c0957610234565b80638da5cb5b116100f65780638da5cb5b14610a215780638f32d59b14610a365780638fc3047d14610a4b578063a4ac25bf14610a60578063a538d28714610a8a57610234565b80637c9d45271461091057806382dc1ec41461093a57806383197ef01461096d5780638456cb59146109825780638d1f5f171461099757610234565b806338bbfa50116101c057806354198ce91161018457806354198ce91461086b5780635c975abb1461089e5780636ef8d66d146108b357806370a08231146108c8578063715018a6146108fb57610234565b806338bbfa50146105d95780633f4ba83a1461071a57806346fbf68e1461072f57806347513621146107765780634b0c80141461083857610234565b80632321841111610207578063232184111461030357806327dc297e1461032d5780632932c76d146103e7578063330a6b33146104d7578063344b78871461050a57610234565b806310026c631461024c57806316f0115b14610291578063179a8d3a146102a657806322d61869146102ec575b600c54610247903463ffffffff610ce416565b600c55005b34801561025857600080fd5b5061027f6004803603604081101561026f57600080fd5b508035906020013560ff16610cff565b60408051918252519081900360200190f35b34801561029d57600080fd5b5061027f610d8e565b3480156102b257600080fd5b506102d0600480360360208110156102c957600080fd5b5035610d94565b60408051600160a060020a039092168252519081900360200190f35b3480156102f857600080fd5b50610301610daf565b005b34801561030f57600080fd5b506103016004803603602081101561032657600080fd5b5035610f80565b34801561033957600080fd5b506103016004803603604081101561035057600080fd5b8135919081019060408101602082013564010000000081111561037257600080fd5b82018360208201111561038457600080fd5b803590602001918460018302840111640100000000831117156103a657600080fd5b91908080601f016020809104026020016040519081016040528093929190818152602001838380828437600092019190915250929550611041945050505050565b3480156103f357600080fd5b5061041a6004803603604081101561040a57600080fd5b508035906020013560ff1661105d565b604051808561014080838360005b83811015610440578181015183820152602001610428565b5050505090500184600a60200280838360005b8381101561046b578181015183820152602001610453565b5050505090500183600a60200280838360005b8381101561049657818101518382015260200161047e565b5050505090500182600a60200280838360005b838110156104c15781810151838201526020016104a9565b5050505090500194505050505060405180910390f35b3480156104e357600080fd5b5061027f600480360360208110156104fa57600080fd5b5035600160a060020a0316611222565b34801561051657600080fd5b506103016004803603604081101561052d57600080fd5b81019060208101813564010000000081111561054857600080fd5b82018360208201111561055a57600080fd5b8035906020019184602083028401116401000000008311171561057c57600080fd5b91939092909160208101903564010000000081111561059a57600080fd5b8201836020820111156105ac57600080fd5b803590602001918460208302840111640100000000831117156105ce57600080fd5b509092509050611234565b3480156105e557600080fd5b50610301600480360360608110156105fc57600080fd5b8135919081019060408101602082013564010000000081111561061e57600080fd5b82018360208201111561063057600080fd5b8035906020019184600183028401116401000000008311171561065257600080fd5b91908080601f01602080910402602001604051908101604052809392919081815260200183838082843760009201919091525092959493602081019350359150506401000000008111156106a557600080fd5b8201836020820111156106b757600080fd5b803590602001918460018302840111640100000000831117156106d957600080fd5b91908080601f016020809104026020016040519081016040528093929190818152602001838380828437600092019190915250929550611675945050505050565b34801561072657600080fd5b50610301611a12565b34801561073b57600080fd5b506107626004803603602081101561075257600080fd5b5035600160a060020a0316611a76565b604080519115158252519081900360200190f35b6103016004803603604081101561078c57600080fd5b8101906020810181356401000000008111156107a757600080fd5b8201836020820111156107b957600080fd5b803590602001918460208302840111640100000000831117156107db57600080fd5b9193909290916020810190356401000000008111156107f957600080fd5b82018360208201111561080b57600080fd5b8035906020019184602083028401116401000000008311171561082d57600080fd5b509092509050611a91565b34801561084457600080fd5b5061027f6004803603602081101561085b57600080fd5b5035600160a060020a0316611cfd565b34801561087757600080fd5b5061027f6004803603602081101561088e57600080fd5b5035600160a060020a0316611d0f565b3480156108aa57600080fd5b50610762611d21565b3480156108bf57600080fd5b50610301611d2b565b3480156108d457600080fd5b5061027f600480360360208110156108eb57600080fd5b5035600160a060020a0316611d36565b34801561090757600080fd5b50610301611d48565b34801561091c57600080fd5b5061027f6004803603602081101561093357600080fd5b5035611db8565b34801561094657600080fd5b506103016004803603602081101561095d57600080fd5b5035600160a060020a0316611dca565b34801561097957600080fd5b50610301611dea565b34801561098e57600080fd5b50610301611e00565b3480156109a357600080fd5b506109ac611e66565b6040805160208082528351818301528351919283929083019185019080838360005b838110156109e65781810151838201526020016109ce565b50505050905090810190601f168015610a135780820380516001836020036101000a031916815260200191505b509250505060405180910390f35b348015610a2d57600080fd5b506102d0611e86565b348015610a4257600080fd5b50610762611e9a565b348015610a5757600080fd5b5061027f611eb0565b348015610a6c57600080fd5b5061027f60048036036020811015610a8357600080fd5b5035611ffd565b348015610a9657600080fd5b50610a9f61200f565b6040805192835260208301919091528051918290030190f35b348015610ac457600080fd5b5061030160048036036020811015610adb57600080fd5b5035600160a060020a03166120af565b348015610af757600080fd5b5061027f6120e4565b348015610b0c57600080fd5b5061030160048036036040811015610b2357600080fd5b50803590602001356120ea565b348015610b3c57600080fd5b50610b4561214e565b604051808661078080838360005b83811015610b6b578181015183820152602001610b53565b5050505090500185603c60200280838360005b83811015610b96578181015183820152602001610b7e565b5050505090500184603c60200280838360005b83811015610bc1578181015183820152602001610ba9565b5050505090500183603c60200280838360005b83811015610bec578181015183820152602001610bd4565b505050509050018281526020019550505050505060405180910390f35b348015610c1557600080fd5b5061030160048036036020811015610c2c57600080fd5b50356122a8565b348015610c3f57600080fd5b5061027f61236d565b348015610c5457600080fd5b5061027f60048036036020811015610c6b57600080fd5b5035612373565b348015610c7e57600080fd5b5061027f60048036036020811015610c9557600080fd5b5035612385565b348015610ca857600080fd5b5061030160048036036020811015610cbf57600080fd5b5035600160a060020a03166124e8565b348015610cdb57600080fd5b5061027f612504565b600082820183811015610cf657600080fd5b90505b92915050565b600060ff82161515610d3357610d2c6064610d20858263ffffffff61250a16565b9063ffffffff61252e16565b9250610d57565b8160ff1660011415610d5757610d54600a610d20858263ffffffff61250a16565b92505b506040805160208082019490945260ff9290921660f860020a02828201528051808303602101815260419092019052805191012090565b600c5481565b601160205260009081526040902054600160a060020a031681565b60065460ff1615610dbf57600080fd5b7f5052494345434845434b5f44454c41590000000000000000000000000000000060005260076020527faac44287c560009bee22f8b7ce1a7c08ba31eb4ba447eaa1305aad77210bd576546009544291011115610e66576040805160e560020a62461bcd02815260206004820152600c60248201527f636f6f6c696e6720646f776e0000000000000000000000000000000000000000604482015290519081900360640190fd5b7f5052494345434845434b5f44454c415900000000000000000000000000000000600052600760209081527faac44287c560009bee22f8b7ce1a7c08ba31eb4ba447eaa1305aad77210bd57654604080518082018252600681527f6e65737465640000000000000000000000000000000000000000000000000000818501528151610180810190925261014b808352303194610f5694939190613fe2908301397f4741535f4c494d4954000000000000000000000000000000000000000000000060005260076020527f3be7d4c26a22519a2102254bad47ede729c28e16609d56101f2de180abec19eb54612559565b50610f7a610f6b82303163ffffffff61295216565b600c549063ffffffff61295216565b600c5550565b610f88611e9a565b1515610f9357600080fd5b600b54811115610fed576040805160e560020a62461bcd02815260206004820152601760248201527f43616e742077697468647261772074686174206d756368000000000000000000604482015290519081900360640190fd5b600b54611000908263ffffffff61295216565b600b55600854604051600160a060020a039091169082156108fc029083906000818181858888f1935050505015801561103d573d6000803e3d6000fd5b5050565b60408051600081526020810190915261103d9083908390611675565b611065613e93565b61106d613e93565b611075613e93565b61107d613e93565b60008560ff16118015611094575060028560ff1611155b15156110ea576040805160e560020a62461bcd02815260206004820152601b60248201527f466f722074696572312063616c6c206765745469657231446174610000000000604482015290519081900360640190fd5b60005b600a811015611218578560ff16600114156111405761112a600a820261111e6064610d208b8263ffffffff61250a16565b9063ffffffff610ce416565b8582600a811061113657fe5b6020020152611169565b80611156600a610d208a8263ffffffff61250a16565b018582600a811061116357fe5b60200201525b60006111858683600a811061117a57fe5b602002015188610cff565b6000818152601360205260409020549091508583600a81106111a357fe5b602090810291909101919091526000828152601290915260409020548483600a81106111cb57fe5b60209081029190910191909152600082815260119091526040902054600160a060020a03168383600a81106111fc57fe5b600160a060020a039092166020929092020152506001016110ed565b5092959194509250565b60106020526000908152604090205481565b60065460ff161561124457600080fd5b82811461129b576040805160e560020a62461bcd02815260206004820152600d60248201527f496e76616c696420696e70757400000000000000000000000000000000000000604482015290519081900360640190fd5b600a8311156112f4576040805160e560020a62461bcd02815260206004820152601a60248201527f4d6178696d756d20313020736c6f747320617420612074696d65000000000000604482015290519081900360640190fd5b60008080805b60ff8116871115611504576000611340898960ff851681811061131957fe5b9050602002013588888560ff16818110151561133157fe5b9050602002013560ff16610cff565b600081815260116020526040902054909150600160a060020a031633146113b1576040805160e560020a62461bcd02815260206004820152601760248201527f4e6f7420616c6c20736c6f74732061726520796f757273000000000000000000604482015290519081900360640190fd5b6000818152601260205260409020546113d190869063ffffffff610ce416565b7f93c763557f8d2895a7b0d4126cd48a84b0acc0bbab8ef8a3d65ce52fe773f0305460008381526012602052604090205491965061142a91620f42409161141e919063ffffffff61252e16565b9063ffffffff61250a16565b935061145d611450888860ff861681811061144157fe5b9050602002013560ff16612967565b849063ffffffff610ce416565b6000828152601260209081526040808320839055601190915290208054600160a060020a031916905592507f3a58ca2fc87c3d2045f75bd819cb1aea9244cdfadebc1be13fded769abf7d66c898960ff85168181106114b857fe5b9050602002013588888560ff1681811015156114d057fe5b6040805194855260ff602092830294909401359390931690840152503382820152519081900360600190a1506001016112fa565b50600c54611518908463ffffffff61295216565b600c55336000908152600f602052604090205461153b908463ffffffff61295216565b336000908152600f602052604090205561155c81600263ffffffff61252e16565b3360009081526010602052604090205410156115c2576040805160e560020a62461bcd02815260206004820152601960248201527f4e6f7420656e6f75676820726573656c6c207469636b65747300000000000000604482015290519081900360640190fd5b6115f26115d682600263ffffffff61252e16565b336000908152601060205260409020549063ffffffff61295216565b33600090815260106020526040812091909155611615848463ffffffff61295216565b604051909150339082156108fc029083906000818181858888f19350505050151561166b57336000908152600d602052604090205461165a908263ffffffff610ce416565b336000908152600d60205260409020555b5050505050505050565b61167d6129df565b600160a060020a031633146116dc576040805160e560020a62461bcd02815260206004820152600b60248201527f61757468206661696c6564000000000000000000000000000000000000000000604482015290519081900360640190fd5b60008381526014602052604090205460ff161561172d5760405160e560020a62461bcd02815260040180806020018281038252602181526020018061412d6021913960400191505060405180910390fd5b6000838152601460205260409020805460ff1916600117905561174f82612bd1565b600a5560005b60038160ff1610156118e557600061176f600a5483610cff565b90506000611795620f424061141e61178686612cbe565b600c549063ffffffff61252e16565b6000838152601360205260409020549091506117b7908263ffffffff610ce416565b600083815260136020908152604080832093909355601190522054600160a060020a0316156118db57600c546117f3908263ffffffff61295216565b600c55600082815260116020908152604080832054600160a060020a03168352600e90915290205461182b908263ffffffff610ce416565b60008381526011602081815260408084208054600160a060020a039081168652600e845282862096909655878552929091529054905192169183156108fc0291849190818181858888f1935050505015156118db57600082815260116020908152604080832054600160a060020a03168352600d9091529020546118b5908263ffffffff610ce416565b600083815260116020908152604080832054600160a060020a03168352600d9091529020555b5050600101611755565b506118ee611d21565b15156118fc576118fc610daf565b426009819055507f768a664c22763f8327912aa5d327af4d164f2704893bba4178a753a13537f48e600a548383604051808481526020018060200180602001838103835285818151815260200191508051906020019080838360005b83811015611970578181015183820152602001611958565b50505050905090810190601f16801561199d5780820380516001836020036101000a031916815260200191505b50838103825284518152845160209182019186019080838360005b838110156119d05781810151838201526020016119b8565b50505050905090810190601f1680156119fd5780820380516001836020036101000a031916815260200191505b509550505050505060405180910390a1505050565b611a1b33611a76565b1515611a2657600080fd5b60065460ff161515611a3757600080fd5b6006805460ff191690556040805133815290517f5db9ee0a495bf2e6ff9c91a7834c1ba4fdd244a5e8aa4e537bd38aeae4b073aa9181900360200190a1565b6000611a8960058363ffffffff612e1116565b90505b919050565b60065460ff1615611aa157600080fd5b7f5052494345434845434b5f44454c41590000000000000000000000000000000060005260076020527faac44287c560009bee22f8b7ce1a7c08ba31eb4ba447eaa1305aad77210bd57654600954611afe9163ffffffff610ce416565b421115611b3f5760405160e560020a62461bcd02815260040180806020018281038252602b815260200180613f8e602b913960400191505060405180910390fd5b828114611b96576040805160e560020a62461bcd02815260206004820152600d60248201527f496e76616c696420696e70757400000000000000000000000000000000000000604482015290519081900360640190fd5b600a831115611bef576040805160e560020a62461bcd02815260206004820152601a60248201527f4d6178696d756d20313020736c6f747320617420612074696d65000000000000604482015290519081900360640190fd5b6000805b60ff8116851115611c5057611c47611c3a878760ff8516818110611c1357fe5b9050602002013586868560ff168181101515611c2b57fe5b9050602002013560ff16612e48565b839063ffffffff610ce416565b50600101611bf3565b5034811115611ca9576040805160e560020a62461bcd02815260206004820152601060248201527f4e6f7420656e6f7567682066756e647300000000000000000000000000000000604482015290519081900360640190fd5b6000611cbb348363ffffffff61295216565b90506000811115611cf557604051339082156108fc029083906000818181858888f19350505050158015611cf3573d6000803e3d6000fd5b505b505050505050565b600f6020526000908152604090205481565b600e6020526000908152604090205481565b60065460ff165b90565b611d34336133c8565b565b600d6020526000908152604090205481565b611d50611e9a565b1515611d5b57600080fd5b6006546040516000916101009004600160a060020a0316907f8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0908390a36006805474ffffffffffffffffffffffffffffffffffffffff0019169055565b60126020526000908152604090205481565b611dd333611a76565b1515611dde57600080fd5b611de781613410565b50565b611df2611e9a565b1515611dfd57600080fd5b33ff5b611e0933611a76565b1515611e1457600080fd5b60065460ff1615611e2457600080fd5b6006805460ff191660011790556040805133815290517f62e78cea01bee320cd4e420270b5ea74000d11b0c9f74754ebdbfc544b05a2589181900360200190a1565b6101806040519081016040528061014b8152602001613fe261014b913981565b6006546101009004600160a060020a031690565b6006546101009004600160a060020a0316331490565b6000611edb7f1100000000000000000000000000000000000000000000000000000000000000613458565b6000611f4f6040805190810160405280600381526020017f55524c0000000000000000000000000000000000000000000000000000000000815250600760007f4741535f4c494d49540000000000000000000000000000000000000000000000815260200190815260200160002054613663565b9050611f5b6000613458565b60408051808201909152600381527f55524c00000000000000000000000000000000000000000000000000000000006020808301919091527f5550444154455f4741535f4c494d495400000000000000000000000000000000600052600790527fb8dbbb2e1e4b9d9097b19c757457fd364ab8ad9119d26d09faf1137b9a66672f54611ff791611fea91613663565b829063ffffffff610ce416565b91505090565b60136020526000908152604090205481565b7f5350524541440000000000000000000000000000000000000000000000000000600090815260076020527f1893b97643a98089770f715d0305eea450b78da77cb2504a1c69ed2b77b740ad5481906002900481612086620f424061141e6120778286610ce4565b600a549063ffffffff61252e16565b905060006120a4620f424061141e612077828763ffffffff61295216565b945090925050509091565b6120b7611e9a565b15156120c257600080fd5b60088054600160a060020a031916600160a060020a0392909216919091179055565b600a5481565b6120f2611e9a565b15156120fd57600080fd5b600082815260076020908152604091829020839055815184815290810183905281517f867716e28221fc5767f7ac9e2623573aa2b682e2b925334db66a0753751708d6929181900390910190a15050565b612156613eb3565b61215e613eb3565b612166613eb3565b61216e613eb3565b600080600061217b61200f565b6064918290049350049050818103600101603c8111156121cf5760405160e560020a62461bcd028152600401808060200182810382526029815260200180613fb96029913960400191505060405180910390fd5b825b82811161229c5760006121e8826064026000610cff565b9050818a868203603c81106121f957fe5b6020908102919091019190915260008281526013909152604090205489868403603c811061222357fe5b6020908102919091019190915260008281526012909152604090205488868403603c811061224d57fe5b60209081029190910191909152600082815260119091526040902054600160a060020a031687868403603c811061228057fe5b600160a060020a039092166020929092020152506001016121d1565b50925050509091929394565b336000908152600d602052604090205481111561230f576040805160e560020a62461bcd02815260206004820152601060248201527f4e6f7420656e6f7567682066756e647300000000000000000000000000000000604482015290519081900360640190fd5b336000908152600d602052604090205461232f908263ffffffff61295216565b336000818152600d6020526040808220939093559151909183156108fc02918491818181858888f1935050505015801561103d573d6000803e3d6000fd5b60095481565b60076020526000908152604090205481565b600080600061239261200f565b909250905060028282030482850381900360008082126123b257816123b7565b816000035b7f484f544e4553535f4d4f44000000000000000000000000000000000000000000600090815260076020527fea096be5dc82e6eb331c8393f48594ecb147108039ee44932379ca150d6fe153549192509061242790620f42409061141e90610d208883888663ffffffff61252e16565b7f484f544e4553535f4d4f4400000000000000000000000000000000000000000060005260076020527fea096be5dc82e6eb331c8393f48594ecb147108039ee44932379ca150d6fe153549091508111156124c957507f484f544e4553535f4d4f4400000000000000000000000000000000000000000060005260076020527fea096be5dc82e6eb331c8393f48594ecb147108039ee44932379ca150d6fe153545b6124dc620f42408263ffffffff61295216565b98975050505050505050565b6124f0611e9a565b15156124fb57600080fd5b611de7816138da565b600b5481565b600080821161251857600080fd5b6000828481151561252557fe5b04949350505050565b600082151561253f57506000610cf9565b82820282848281151561254e57fe5b0414610cf657600080fd5b600154600090600160a060020a03161580612586575060015461258490600160a060020a0316613963565b155b15612597576125956000613967565b505b600160009054906101000a9004600160a060020a0316600160a060020a03166338cc48316040518163ffffffff1660e060020a028152600401602060405180830381600087803b1580156125ea57600080fd5b505af11580156125fe573d6000803e3d6000fd5b505050506040513d602081101561261457600080fd5b5051600054600160a060020a039081169116146126ca57600160009054906101000a9004600160a060020a0316600160a060020a03166338cc48316040518163ffffffff1660e060020a028152600401602060405180830381600087803b15801561267e57600080fd5b505af1158015612692573d6000803e3d6000fd5b505050506040513d60208110156126a857600080fd5b505160008054600160a060020a031916600160a060020a039092169190911790555b60008054604080517f2ef3accc0000000000000000000000000000000000000000000000000000000081526024810186905260048101918252875160448201528751600160a060020a0390931692632ef3accc928992889282916064019060208601908083838c5b8381101561274a578181015183820152602001612732565b50505050905090810190601f1680156127775780820380516001836020036101000a031916815260200191505b509350505050602060405180830381600087803b15801561279757600080fd5b505af11580156127ab573d6000803e3d6000fd5b505050506040513d60208110156127c157600080fd5b50519050670de0b6b3a76400003a8402018111156127e357506000905061294a565b6000809054906101000a9004600160a060020a0316600160a060020a031663c51be90f82888888886040518663ffffffff1660e060020a028152600401808581526020018060200180602001848152602001838103835286818151815260200191508051906020019080838360005b8381101561286a578181015183820152602001612852565b50505050905090810190601f1680156128975780820380516001836020036101000a031916815260200191505b50838103825285518152855160209182019187019080838360005b838110156128ca5781810151838201526020016128b2565b50505050905090810190601f1680156128f75780820380516001836020036101000a031916815260200191505b5096505050505050506020604051808303818588803b15801561291957600080fd5b505af115801561292d573d6000803e3d6000fd5b50505050506040513d602081101561294457600080fd5b50519150505b949350505050565b60008282111561296157600080fd5b50900390565b6000600260ff831611156129af5760405160e560020a62461bcd028152600401808060200182810382526022815260200180613f6c6022913960400191505060405180910390fd5b8160ff16600214156129c357506001611a8c565b8160ff16600114156129d75750600a611a8c565b506064611a8c565b600154600090600160a060020a03161580612a0c5750600154612a0a90600160a060020a0316613963565b155b15612a1d57612a1b6000613967565b505b600160009054906101000a9004600160a060020a0316600160a060020a03166338cc48316040518163ffffffff1660e060020a028152600401602060405180830381600087803b158015612a7057600080fd5b505af1158015612a84573d6000803e3d6000fd5b505050506040513d6020811015612a9a57600080fd5b5051600054600160a060020a03908116911614612b5057600160009054906101000a9004600160a060020a0316600160a060020a03166338cc48316040518163ffffffff1660e060020a028152600401602060405180830381600087803b158015612b0457600080fd5b505af1158015612b18573d6000803e3d6000fd5b505050506040513d6020811015612b2e57600080fd5b505160008054600160a060020a031916600160a060020a039092169190911790555b6000809054906101000a9004600160a060020a0316600160a060020a031663c281d19e6040518163ffffffff1660e060020a02815260040160206040518083038186803b158015612ba057600080fd5b505afa158015612bb4573d6000803e3d6000fd5b505050506040513d6020811015612bca57600080fd5b5051905090565b600060028183815b8151811015612c9d578215612bef576001840393505b8181815181101515612bfd57fe5b90602001015160f860020a900460f860020a0260f860020a900460ff16602e1415612c2757600192505b60008282815181101515612c3757fe5b90602001015160f860020a900460f860020a0260f860020a900460ff16905060308110158015612c68575060398111155b15612c79576030810386600a020195505b838015612c84575084155b15612c945750611a8c9350505050565b50600101612bd9565b5b8315612cb55784600a029450600184039350612c9e565b50505050919050565b6000600260ff83161115612d065760405160e560020a62461bcd028152600401808060200182810382526022815260200180613f6c6022913960400191505060405180910390fd5b8160ff1660021415612d6357507f54494552335f5041594f5554000000000000000000000000000000000000000060005260076020527facc1313bfd621232066a9c54ce789299c5df44b1ef9af83db841e307b10705f454611a8c565b8160ff1660011415612dc057507f54494552325f5041594f5554000000000000000000000000000000000000000060005260076020527f1934936e913495172a68f489eba8f04b1b844c8d38a3035a45b08189f9af57cd54611a8c565b507f54494552315f5041594f5554000000000000000000000000000000000000000060005260076020527f7729a41062816f77329ca956209d1a1aa7b7e1abd27f215c57ad3a356065d00554611a8c565b6000600160a060020a0382161515612e2857600080fd5b50600160a060020a03166000908152602091909152604090205460ff1690565b60065460009060ff1615612e5b57600080fd5b600260ff83161115612ea15760405160e560020a62461bcd028152600401808060200182810382526022815260200180613f6c6022913960400191505060405180910390fd5b600a541515612efa576040805160e560020a62461bcd02815260206004820152601660248201527f47616d65206861736e7420737461727465642079657400000000000000000000604482015290519081900360640190fd5b6000612f068484610cff565b9050600080612f1485612967565b905060008080612f2261200f565b9092509050818910801590612f375750808911155b1515612f8d576040805160e560020a62461bcd02815260206004820152601a60248201527f43616e742062757920746869732070726f706572747920796574000000000000604482015290519081900360640190fd5b6000868152601260205260409020541515612fdf576000612fad89613971565b90506000612fba8b612385565b9050612fd6611c3a620f424061141e858563ffffffff61252e16565b9650505061322f565b6000868152601260205260409020547f03565cf03431b9675ce6f670301da7655569d596d084d861c90455dcee47bff55461302490620f42409061141e90849061252e565b7f52454255595f4645450000000000000000000000000000000000000000000000600090815260076020527fc74b37efb558a657abc1a4c0bb36f4e7aa6d90bc016115585b63ee7bdb5965b1549197509061309290620f42409061141e90610d208b8763ffffffff61295216565b600089815260116020908152604080832054600160a060020a03168084526010909252909120549096509091508590871161310e57600160a060020a0381166000908152601060205260409020546130f0908863ffffffff61295216565b600160a060020a038216600090815260106020526040902055613147565b600160a060020a038116600090815260106020526040812054111561314757600160a060020a0381166000908152601060205260408120555b600160a060020a0381166000908152600e6020526040902054613170908363ffffffff610ce416565b600160a060020a0382166000908152600e602052604090205561319c610f6b848463ffffffff610ce416565b600c55600160a060020a0381166108fc6131bc858563ffffffff610ce416565b6040518115909202916000818181858888f19350505050151561322b576132116131ec848463ffffffff610ce416565b600160a060020a0383166000908152600d60205260409020549063ffffffff610ce416565b600160a060020a0382166000908152600d60205260409020555b5050505b7f484f5553455f4355540000000000000000000000000000000000000000000000600090815260076020527f9428f94ef6f4ec3c650bfb4a2a71424d241e94b5c309003957bb775f67e0c4325461329690620f42409061141e90899063ffffffff61252e16565b905060006132aa878363ffffffff61295216565b600c549091506132c0908263ffffffff610ce416565b600c55600b546132d6908363ffffffff610ce416565b600b556000888152601260209081526040808320849055601182528083208054600160a060020a031916339081179091558352600f909152902054613321908263ffffffff610ce416565b336000908152600f602090815260408083209390935560109052205461334d908763ffffffff610ce416565b336000818152601060209081526040918290209390935580518e815292830184905260ff8d1683820152600160a060020a03881660608401526080830191909152517f34176d7031c41eaafbc895af3f7b5b4e04f92dbb724a34fbb917b1ac56f7ae539181900360a00190a150949998505050505050505050565b6133d960058263ffffffff613ac416565b604051600160a060020a038216907fcd265ebaf09df2871cc7bd4133404a235ba12eff2041bb89d9c714a2621c7c7e90600090a250565b61342160058263ffffffff613b1016565b604051600160a060020a038216907f6719d08c1888103bea251a4ed56406bd0c3e69723c8a1686e017e7bbe159b6f890600090a250565b600154600160a060020a03161580613482575060015461348090600160a060020a0316613963565b155b15613493576134916000613967565b505b600160009054906101000a9004600160a060020a0316600160a060020a03166338cc48316040518163ffffffff1660e060020a028152600401602060405180830381600087803b1580156134e657600080fd5b505af11580156134fa573d6000803e3d6000fd5b505050506040513d602081101561351057600080fd5b5051600054600160a060020a039081169116146135c657600160009054906101000a9004600160a060020a0316600160a060020a03166338cc48316040518163ffffffff1660e060020a028152600401602060405180830381600087803b15801561357a57600080fd5b505af115801561358e573d6000803e3d6000fd5b505050506040513d60208110156135a457600080fd5b505160008054600160a060020a031916600160a060020a039092169190911790555b60008054604080517f688dcfd70000000000000000000000000000000000000000000000000000000081527fff00000000000000000000000000000000000000000000000000000000000000851660048201529051600160a060020a039092169263688dcfd79260248084019382900301818387803b15801561364857600080fd5b505af115801561365c573d6000803e3d6000fd5b5050505050565b600154600090600160a060020a03161580613690575060015461368e90600160a060020a0316613963565b155b156136a15761369f6000613967565b505b600160009054906101000a9004600160a060020a0316600160a060020a03166338cc48316040518163ffffffff1660e060020a028152600401602060405180830381600087803b1580156136f457600080fd5b505af1158015613708573d6000803e3d6000fd5b505050506040513d602081101561371e57600080fd5b5051600054600160a060020a039081169116146137d457600160009054906101000a9004600160a060020a0316600160a060020a03166338cc48316040518163ffffffff1660e060020a028152600401602060405180830381600087803b15801561378857600080fd5b505af115801561379c573d6000803e3d6000fd5b505050506040513d60208110156137b257600080fd5b505160008054600160a060020a031916600160a060020a039092169190911790555b60008054604080517f2ef3accc0000000000000000000000000000000000000000000000000000000081526024810186905260048101918252865160448201528651600160a060020a0390931693632ef3accc93889388939092839260649092019160208701918190849084905b8381101561385a578181015183820152602001613842565b50505050905090810190601f1680156138875780820380516001836020036101000a031916815260200191505b509350505050602060405180830381600087803b1580156138a757600080fd5b505af11580156138bb573d6000803e3d6000fd5b505050506040513d60208110156138d157600080fd5b50519392505050565b600160a060020a03811615156138ef57600080fd5b600654604051600160a060020a0380841692610100900416907f8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e090600090a360068054600160a060020a039092166101000274ffffffffffffffffffffffffffffffffffffffff0019909216919091179055565b3b90565b6000611a89613b5e565b6000600260ff831611156139b95760405160e560020a62461bcd028152600401808060200182810382526022815260200180613f6c6022913960400191505060405180910390fd5b8160ff1660021415613a1657507f54494552335f505249434500000000000000000000000000000000000000000060005260076020527fae6e034939092767869a45875d60d5a3f30e8ca6e83ecd66c948bce2cdf3c70e54611a8c565b8160ff1660011415613a7357507f54494552325f505249434500000000000000000000000000000000000000000060005260076020527f161aa644b3561c0a2ed9ab0eb5ebcb59d229bba8d49af22be9d0922f604d1f2954611a8c565b507f54494552315f505249434500000000000000000000000000000000000000000060005260076020527f94d03aac1c5adffb43bea72d0d9572fc98b730fe4f355bdda67a1c5de5d6c67754611a8c565b600160a060020a0381161515613ad957600080fd5b613ae38282612e11565b1515613aee57600080fd5b600160a060020a0316600090815260209190915260409020805460ff19169055565b600160a060020a0381161515613b2557600080fd5b613b2f8282612e11565b15613b3957600080fd5b600160a060020a0316600090815260209190915260409020805460ff19166001179055565b600080613b7e731d3b2638a7cc9f2cb3d298a3da7a90b67e5506ed613963565b1115613bef5760018054600160a060020a031916731d3b2638a7cc9f2cb3d298a3da7a90b67e5506ed17905560408051808201909152600b81527f6574685f6d61696e6e65740000000000000000000000000000000000000000006020820152613be790613e80565b506001611d28565b6000613c0e73c03a2615d5efaf5f49f60b7bb6583eaec212fdf1613963565b1115613c775760018054600160a060020a03191673c03a2615d5efaf5f49f60b7bb6583eaec212fdf117905560408051808201909152600c81527f6574685f726f707374656e3300000000000000000000000000000000000000006020820152613be790613e80565b6000613c9673b7a07bcf2ba2f2703b24c0691b5278999c59ac7e613963565b1115613cff5760018054600160a060020a03191673b7a07bcf2ba2f2703b24c0691b5278999c59ac7e17905560408051808201909152600981527f6574685f6b6f76616e00000000000000000000000000000000000000000000006020820152613be790613e80565b6000613d1e73146500cfd35b22e4a392fe0adc06de1a1368ed48613963565b1115613d875760018054600160a060020a03191673146500cfd35b22e4a392fe0adc06de1a1368ed4817905560408051808201909152600b81527f6574685f72696e6b6562790000000000000000000000000000000000000000006020820152613be790613e80565b6000613da6736f485c8bf6fc43ea212e93bbf8ce046c7f1cb475613963565b1115613dd8575060018054600160a060020a031916736f485c8bf6fc43ea212e93bbf8ce046c7f1cb475178155611d28565b6000613df77320e12a1f859b3feae5fb2a0a32c18f5a65555bbf613963565b1115613e29575060018054600160a060020a0319167320e12a1f859b3feae5fb2a0a32c18f5a65555bbf178155611d28565b6000613e487351efaf4c8b3c9afbd5ab9f4bbc82784ab6ef8faa613963565b1115613e7a575060018054600160a060020a0319167351efaf4c8b3c9afbd5ab9f4bbc82784ab6ef8faa178155611d28565b50600090565b805161103d906002906020840190613ed3565b61014060405190810160405280600a906020820280388339509192915050565b61078060405190810160405280603c906020820280388339509192915050565b828054600181600116156101000203166002900490600052602060002090601f016020900481019282601f10613f1457805160ff1916838001178555613f41565b82800160010185558215613f41579182015b82811115613f41578251825591602001919060010190613f26565b50613f4d929150613f51565b5090565b611d2891905b80821115613f4d5760008155600101613f5756fe6d6178696d756d20322064696769747320707265636973696f6e20616c6c6f77656443616e277420707572636861736520647572696e672070726963652077616974696e67206c6f636b6f7574546869732066756e6374696f6e2063616e742068616e646c6520612073707265616420736f206269675b55524c5d206a736f6e2868747470733a2f2f6d696e2d6170692e63727970746f636f6d706172652e636f6d2f646174612f70726963653f6673796d3d4c5443267473796d733d555344266578747261506172616d733d5072696365456d70697265267369676e3d74727565266170695f6b65793d247b5b646563727970745d204250526a742b4e6c6356337839366d7034726665675a4544535a415548654d34717279414e6565347177386a6362454c6d324e6f78544255674e655547375833465a396e6e31302b564c742f3271797365396c34426979506471664e4534474a76512f4d7130716633626455724b586e50665842644b445336656a595963395438374e766a6a645544697942336e6e59444937586978556b7865685135795847635a473663717a6c466a4345327630735578743864737633646c74553174302f57413d3d7d292e55534451756572792068617320616c7265616479206265656e2070726f63657373656421a165627a7a72305820a29f3ccc0f8b9edf8f854f918d5e9e133bb7e8de61e823e7ca8eda3ce234aa4100293be7d4c26a22519a2102254bad47ede729c28e16609d56101f2de180abec19eb5b55524c5d206a736f6e2868747470733a2f2f6d696e2d6170692e63727970746f636f6d706172652e636f6d2f646174612f70726963653f6673796d3d4c5443267473796d733d555344266578747261506172616d733d5072696365456d70697265267369676e3d74727565266170695f6b65793d247b5b646563727970745d204250526a742b4e6c6356337839366d7034726665675a4544535a415548654d34717279414e6565347177386a6362454c6d324e6f78544255674e655547375833465a396e6e31302b564c742f3271797365396c34426979506471664e4534474a76512f4d7130716633626455724b586e50665842644b445336656a595963395438374e766a6a645544697942336e6e59444937586978556b7865685135795847635a473663717a6c466a4345327630735578743864737633646c74553174302f57413d3d7d292e555344aac44287c560009bee22f8b7ce1a7c08ba31eb4ba447eaa1305aad77210bd576`

// DeployCoinEmpire deploys a new Ethereum contract, binding an instance of CoinEmpire to it.
func DeployCoinEmpire(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *CoinEmpire, error) {
	parsed, err := abi.JSON(strings.NewReader(CoinEmpireABI))
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	address, tx, contract, err := bind.DeployContract(auth, parsed, common.FromHex(CoinEmpireBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &CoinEmpire{CoinEmpireCaller: CoinEmpireCaller{contract: contract}, CoinEmpireTransactor: CoinEmpireTransactor{contract: contract}, CoinEmpireFilterer: CoinEmpireFilterer{contract: contract}}, nil
}

// CoinEmpire is an auto generated Go binding around an Ethereum contract.
type CoinEmpire struct {
	CoinEmpireCaller     // Read-only binding to the contract
	CoinEmpireTransactor // Write-only binding to the contract
	CoinEmpireFilterer   // Log filterer for contract events
}

// CoinEmpireCaller is an auto generated read-only Go binding around an Ethereum contract.
type CoinEmpireCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// CoinEmpireTransactor is an auto generated write-only Go binding around an Ethereum contract.
type CoinEmpireTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// CoinEmpireFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type CoinEmpireFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// CoinEmpireSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type CoinEmpireSession struct {
	Contract     *CoinEmpire       // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// CoinEmpireCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type CoinEmpireCallerSession struct {
	Contract *CoinEmpireCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts     // Call options to use throughout this session
}

// CoinEmpireTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type CoinEmpireTransactorSession struct {
	Contract     *CoinEmpireTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts     // Transaction auth options to use throughout this session
}

// CoinEmpireRaw is an auto generated low-level Go binding around an Ethereum contract.
type CoinEmpireRaw struct {
	Contract *CoinEmpire // Generic contract binding to access the raw methods on
}

// CoinEmpireCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type CoinEmpireCallerRaw struct {
	Contract *CoinEmpireCaller // Generic read-only contract binding to access the raw methods on
}

// CoinEmpireTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type CoinEmpireTransactorRaw struct {
	Contract *CoinEmpireTransactor // Generic write-only contract binding to access the raw methods on
}

// NewCoinEmpire creates a new instance of CoinEmpire, bound to a specific deployed contract.
func NewCoinEmpire(address common.Address, backend bind.ContractBackend) (*CoinEmpire, error) {
	contract, err := bindCoinEmpire(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &CoinEmpire{CoinEmpireCaller: CoinEmpireCaller{contract: contract}, CoinEmpireTransactor: CoinEmpireTransactor{contract: contract}, CoinEmpireFilterer: CoinEmpireFilterer{contract: contract}}, nil
}

// NewCoinEmpireCaller creates a new read-only instance of CoinEmpire, bound to a specific deployed contract.
func NewCoinEmpireCaller(address common.Address, caller bind.ContractCaller) (*CoinEmpireCaller, error) {
	contract, err := bindCoinEmpire(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &CoinEmpireCaller{contract: contract}, nil
}

// NewCoinEmpireTransactor creates a new write-only instance of CoinEmpire, bound to a specific deployed contract.
func NewCoinEmpireTransactor(address common.Address, transactor bind.ContractTransactor) (*CoinEmpireTransactor, error) {
	contract, err := bindCoinEmpire(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &CoinEmpireTransactor{contract: contract}, nil
}

// NewCoinEmpireFilterer creates a new log filterer instance of CoinEmpire, bound to a specific deployed contract.
func NewCoinEmpireFilterer(address common.Address, filterer bind.ContractFilterer) (*CoinEmpireFilterer, error) {
	contract, err := bindCoinEmpire(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &CoinEmpireFilterer{contract: contract}, nil
}

// bindCoinEmpire binds a generic wrapper to an already deployed contract.
func bindCoinEmpire(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(CoinEmpireABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_CoinEmpire *CoinEmpireRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _CoinEmpire.Contract.CoinEmpireCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_CoinEmpire *CoinEmpireRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _CoinEmpire.Contract.CoinEmpireTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_CoinEmpire *CoinEmpireRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _CoinEmpire.Contract.CoinEmpireTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_CoinEmpire *CoinEmpireCallerRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _CoinEmpire.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_CoinEmpire *CoinEmpireTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _CoinEmpire.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_CoinEmpire *CoinEmpireTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _CoinEmpire.Contract.contract.Transact(opts, method, params...)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address ) constant returns(uint256)
func (_CoinEmpire *CoinEmpireCaller) BalanceOf(opts *bind.CallOpts, arg0 common.Address) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _CoinEmpire.contract.Call(opts, out, "balanceOf", arg0)
	return *ret0, err
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address ) constant returns(uint256)
func (_CoinEmpire *CoinEmpireSession) BalanceOf(arg0 common.Address) (*big.Int, error) {
	return _CoinEmpire.Contract.BalanceOf(&_CoinEmpire.CallOpts, arg0)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address ) constant returns(uint256)
func (_CoinEmpire *CoinEmpireCallerSession) BalanceOf(arg0 common.Address) (*big.Int, error) {
	return _CoinEmpire.Contract.BalanceOf(&_CoinEmpire.CallOpts, arg0)
}

// CapitalOf is a free data retrieval call binding the contract method 0x4b0c8014.
//
// Solidity: function capitalOf(address ) constant returns(uint256)
func (_CoinEmpire *CoinEmpireCaller) CapitalOf(opts *bind.CallOpts, arg0 common.Address) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _CoinEmpire.contract.Call(opts, out, "capitalOf", arg0)
	return *ret0, err
}

// CapitalOf is a free data retrieval call binding the contract method 0x4b0c8014.
//
// Solidity: function capitalOf(address ) constant returns(uint256)
func (_CoinEmpire *CoinEmpireSession) CapitalOf(arg0 common.Address) (*big.Int, error) {
	return _CoinEmpire.Contract.CapitalOf(&_CoinEmpire.CallOpts, arg0)
}

// CapitalOf is a free data retrieval call binding the contract method 0x4b0c8014.
//
// Solidity: function capitalOf(address ) constant returns(uint256)
func (_CoinEmpire *CoinEmpireCallerSession) CapitalOf(arg0 common.Address) (*big.Int, error) {
	return _CoinEmpire.Contract.CapitalOf(&_CoinEmpire.CallOpts, arg0)
}

// Config is a free data retrieval call binding the contract method 0xcc718f76.
//
// Solidity: function config(bytes32 ) constant returns(uint256)
func (_CoinEmpire *CoinEmpireCaller) Config(opts *bind.CallOpts, arg0 [32]byte) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _CoinEmpire.contract.Call(opts, out, "config", arg0)
	return *ret0, err
}

// Config is a free data retrieval call binding the contract method 0xcc718f76.
//
// Solidity: function config(bytes32 ) constant returns(uint256)
func (_CoinEmpire *CoinEmpireSession) Config(arg0 [32]byte) (*big.Int, error) {
	return _CoinEmpire.Contract.Config(&_CoinEmpire.CallOpts, arg0)
}

// Config is a free data retrieval call binding the contract method 0xcc718f76.
//
// Solidity: function config(bytes32 ) constant returns(uint256)
func (_CoinEmpire *CoinEmpireCallerSession) Config(arg0 [32]byte) (*big.Int, error) {
	return _CoinEmpire.Contract.Config(&_CoinEmpire.CallOpts, arg0)
}

// CurrentPrice is a free data retrieval call binding the contract method 0xab2b3137.
//
// Solidity: function current_price() constant returns(uint256)
func (_CoinEmpire *CoinEmpireCaller) CurrentPrice(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _CoinEmpire.contract.Call(opts, out, "current_price")
	return *ret0, err
}

// CurrentPrice is a free data retrieval call binding the contract method 0xab2b3137.
//
// Solidity: function current_price() constant returns(uint256)
func (_CoinEmpire *CoinEmpireSession) CurrentPrice() (*big.Int, error) {
	return _CoinEmpire.Contract.CurrentPrice(&_CoinEmpire.CallOpts)
}

// CurrentPrice is a free data retrieval call binding the contract method 0xab2b3137.
//
// Solidity: function current_price() constant returns(uint256)
func (_CoinEmpire *CoinEmpireCallerSession) CurrentPrice() (*big.Int, error) {
	return _CoinEmpire.Contract.CurrentPrice(&_CoinEmpire.CallOpts)
}

// GetHotnessModifier is a free data retrieval call binding the contract method 0xddbb73b7.
//
// Solidity: function getHotnessModifier(uint256 price) constant returns(uint256)
func (_CoinEmpire *CoinEmpireCaller) GetHotnessModifier(opts *bind.CallOpts, price *big.Int) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _CoinEmpire.contract.Call(opts, out, "getHotnessModifier", price)
	return *ret0, err
}

// GetHotnessModifier is a free data retrieval call binding the contract method 0xddbb73b7.
//
// Solidity: function getHotnessModifier(uint256 price) constant returns(uint256)
func (_CoinEmpire *CoinEmpireSession) GetHotnessModifier(price *big.Int) (*big.Int, error) {
	return _CoinEmpire.Contract.GetHotnessModifier(&_CoinEmpire.CallOpts, price)
}

// GetHotnessModifier is a free data retrieval call binding the contract method 0xddbb73b7.
//
// Solidity: function getHotnessModifier(uint256 price) constant returns(uint256)
func (_CoinEmpire *CoinEmpireCallerSession) GetHotnessModifier(price *big.Int) (*big.Int, error) {
	return _CoinEmpire.Contract.GetHotnessModifier(&_CoinEmpire.CallOpts, price)
}

// GetMinMax is a free data retrieval call binding the contract method 0xa538d287.
//
// Solidity: function getMinMax() constant returns(uint256 _min, uint256 _max)
func (_CoinEmpire *CoinEmpireCaller) GetMinMax(opts *bind.CallOpts) (struct {
	Min *big.Int
	Max *big.Int
}, error) {
	ret := new(struct {
		Min *big.Int
		Max *big.Int
	})
	out := ret
	err := _CoinEmpire.contract.Call(opts, out, "getMinMax")
	return *ret, err
}

// GetMinMax is a free data retrieval call binding the contract method 0xa538d287.
//
// Solidity: function getMinMax() constant returns(uint256 _min, uint256 _max)
func (_CoinEmpire *CoinEmpireSession) GetMinMax() (struct {
	Min *big.Int
	Max *big.Int
}, error) {
	return _CoinEmpire.Contract.GetMinMax(&_CoinEmpire.CallOpts)
}

// GetMinMax is a free data retrieval call binding the contract method 0xa538d287.
//
// Solidity: function getMinMax() constant returns(uint256 _min, uint256 _max)
func (_CoinEmpire *CoinEmpireCallerSession) GetMinMax() (struct {
	Min *big.Int
	Max *big.Int
}, error) {
	return _CoinEmpire.Contract.GetMinMax(&_CoinEmpire.CallOpts)
}

// GetSlotId is a free data retrieval call binding the contract method 0x10026c63.
//
// Solidity: function getSlotId(uint256 price, uint8 tier) constant returns(uint256)
func (_CoinEmpire *CoinEmpireCaller) GetSlotId(opts *bind.CallOpts, price *big.Int, tier uint8) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _CoinEmpire.contract.Call(opts, out, "getSlotId", price, tier)
	return *ret0, err
}

// GetSlotId is a free data retrieval call binding the contract method 0x10026c63.
//
// Solidity: function getSlotId(uint256 price, uint8 tier) constant returns(uint256)
func (_CoinEmpire *CoinEmpireSession) GetSlotId(price *big.Int, tier uint8) (*big.Int, error) {
	return _CoinEmpire.Contract.GetSlotId(&_CoinEmpire.CallOpts, price, tier)
}

// GetSlotId is a free data retrieval call binding the contract method 0x10026c63.
//
// Solidity: function getSlotId(uint256 price, uint8 tier) constant returns(uint256)
func (_CoinEmpire *CoinEmpireCallerSession) GetSlotId(price *big.Int, tier uint8) (*big.Int, error) {
	return _CoinEmpire.Contract.GetSlotId(&_CoinEmpire.CallOpts, price, tier)
}

// GetTier1Data is a free data retrieval call binding the contract method 0xb93298f2.
//
// Solidity: function getTier1Data() constant returns(uint256[60] _index_data, uint256[60] _earnings_data, uint256[60] _price_data, address[60] _owner_data, uint256 _length)
func (_CoinEmpire *CoinEmpireCaller) GetTier1Data(opts *bind.CallOpts) (struct {
	IndexData    [60]*big.Int
	EarningsData [60]*big.Int
	PriceData    [60]*big.Int
	OwnerData    [60]common.Address
	Length       *big.Int
}, error) {
	ret := new(struct {
		IndexData    [60]*big.Int
		EarningsData [60]*big.Int
		PriceData    [60]*big.Int
		OwnerData    [60]common.Address
		Length       *big.Int
	})
	out := ret
	err := _CoinEmpire.contract.Call(opts, out, "getTier1Data")
	return *ret, err
}

// GetTier1Data is a free data retrieval call binding the contract method 0xb93298f2.
//
// Solidity: function getTier1Data() constant returns(uint256[60] _index_data, uint256[60] _earnings_data, uint256[60] _price_data, address[60] _owner_data, uint256 _length)
func (_CoinEmpire *CoinEmpireSession) GetTier1Data() (struct {
	IndexData    [60]*big.Int
	EarningsData [60]*big.Int
	PriceData    [60]*big.Int
	OwnerData    [60]common.Address
	Length       *big.Int
}, error) {
	return _CoinEmpire.Contract.GetTier1Data(&_CoinEmpire.CallOpts)
}

// GetTier1Data is a free data retrieval call binding the contract method 0xb93298f2.
//
// Solidity: function getTier1Data() constant returns(uint256[60] _index_data, uint256[60] _earnings_data, uint256[60] _price_data, address[60] _owner_data, uint256 _length)
func (_CoinEmpire *CoinEmpireCallerSession) GetTier1Data() (struct {
	IndexData    [60]*big.Int
	EarningsData [60]*big.Int
	PriceData    [60]*big.Int
	OwnerData    [60]common.Address
	Length       *big.Int
}, error) {
	return _CoinEmpire.Contract.GetTier1Data(&_CoinEmpire.CallOpts)
}

// GetTierData is a free data retrieval call binding the contract method 0x2932c76d.
//
// Solidity: function getTierData(uint256 price, uint8 tier) constant returns(uint256[10] _index_data, uint256[10] _earnings_data, uint256[10] _price_data, address[10] _owner_data)
func (_CoinEmpire *CoinEmpireCaller) GetTierData(opts *bind.CallOpts, price *big.Int, tier uint8) (struct {
	IndexData    [10]*big.Int
	EarningsData [10]*big.Int
	PriceData    [10]*big.Int
	OwnerData    [10]common.Address
}, error) {
	ret := new(struct {
		IndexData    [10]*big.Int
		EarningsData [10]*big.Int
		PriceData    [10]*big.Int
		OwnerData    [10]common.Address
	})
	out := ret
	err := _CoinEmpire.contract.Call(opts, out, "getTierData", price, tier)
	return *ret, err
}

// GetTierData is a free data retrieval call binding the contract method 0x2932c76d.
//
// Solidity: function getTierData(uint256 price, uint8 tier) constant returns(uint256[10] _index_data, uint256[10] _earnings_data, uint256[10] _price_data, address[10] _owner_data)
func (_CoinEmpire *CoinEmpireSession) GetTierData(price *big.Int, tier uint8) (struct {
	IndexData    [10]*big.Int
	EarningsData [10]*big.Int
	PriceData    [10]*big.Int
	OwnerData    [10]common.Address
}, error) {
	return _CoinEmpire.Contract.GetTierData(&_CoinEmpire.CallOpts, price, tier)
}

// GetTierData is a free data retrieval call binding the contract method 0x2932c76d.
//
// Solidity: function getTierData(uint256 price, uint8 tier) constant returns(uint256[10] _index_data, uint256[10] _earnings_data, uint256[10] _price_data, address[10] _owner_data)
func (_CoinEmpire *CoinEmpireCallerSession) GetTierData(price *big.Int, tier uint8) (struct {
	IndexData    [10]*big.Int
	EarningsData [10]*big.Int
	PriceData    [10]*big.Int
	OwnerData    [10]common.Address
}, error) {
	return _CoinEmpire.Contract.GetTierData(&_CoinEmpire.CallOpts, price, tier)
}

// House is a free data retrieval call binding the contract method 0xff9b3acf.
//
// Solidity: function house() constant returns(uint256)
func (_CoinEmpire *CoinEmpireCaller) House(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _CoinEmpire.contract.Call(opts, out, "house")
	return *ret0, err
}

// House is a free data retrieval call binding the contract method 0xff9b3acf.
//
// Solidity: function house() constant returns(uint256)
func (_CoinEmpire *CoinEmpireSession) House() (*big.Int, error) {
	return _CoinEmpire.Contract.House(&_CoinEmpire.CallOpts)
}

// House is a free data retrieval call binding the contract method 0xff9b3acf.
//
// Solidity: function house() constant returns(uint256)
func (_CoinEmpire *CoinEmpireCallerSession) House() (*big.Int, error) {
	return _CoinEmpire.Contract.House(&_CoinEmpire.CallOpts)
}

// IsOwner is a free data retrieval call binding the contract method 0x8f32d59b.
//
// Solidity: function isOwner() constant returns(bool)
func (_CoinEmpire *CoinEmpireCaller) IsOwner(opts *bind.CallOpts) (bool, error) {
	var (
		ret0 = new(bool)
	)
	out := ret0
	err := _CoinEmpire.contract.Call(opts, out, "isOwner")
	return *ret0, err
}

// IsOwner is a free data retrieval call binding the contract method 0x8f32d59b.
//
// Solidity: function isOwner() constant returns(bool)
func (_CoinEmpire *CoinEmpireSession) IsOwner() (bool, error) {
	return _CoinEmpire.Contract.IsOwner(&_CoinEmpire.CallOpts)
}

// IsOwner is a free data retrieval call binding the contract method 0x8f32d59b.
//
// Solidity: function isOwner() constant returns(bool)
func (_CoinEmpire *CoinEmpireCallerSession) IsOwner() (bool, error) {
	return _CoinEmpire.Contract.IsOwner(&_CoinEmpire.CallOpts)
}

// IsPauser is a free data retrieval call binding the contract method 0x46fbf68e.
//
// Solidity: function isPauser(address account) constant returns(bool)
func (_CoinEmpire *CoinEmpireCaller) IsPauser(opts *bind.CallOpts, account common.Address) (bool, error) {
	var (
		ret0 = new(bool)
	)
	out := ret0
	err := _CoinEmpire.contract.Call(opts, out, "isPauser", account)
	return *ret0, err
}

// IsPauser is a free data retrieval call binding the contract method 0x46fbf68e.
//
// Solidity: function isPauser(address account) constant returns(bool)
func (_CoinEmpire *CoinEmpireSession) IsPauser(account common.Address) (bool, error) {
	return _CoinEmpire.Contract.IsPauser(&_CoinEmpire.CallOpts, account)
}

// IsPauser is a free data retrieval call binding the contract method 0x46fbf68e.
//
// Solidity: function isPauser(address account) constant returns(bool)
func (_CoinEmpire *CoinEmpireCallerSession) IsPauser(account common.Address) (bool, error) {
	return _CoinEmpire.Contract.IsPauser(&_CoinEmpire.CallOpts, account)
}

// LatestSample is a free data retrieval call binding the contract method 0xc595d6b9.
//
// Solidity: function latest_sample() constant returns(uint256)
func (_CoinEmpire *CoinEmpireCaller) LatestSample(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _CoinEmpire.contract.Call(opts, out, "latest_sample")
	return *ret0, err
}

// LatestSample is a free data retrieval call binding the contract method 0xc595d6b9.
//
// Solidity: function latest_sample() constant returns(uint256)
func (_CoinEmpire *CoinEmpireSession) LatestSample() (*big.Int, error) {
	return _CoinEmpire.Contract.LatestSample(&_CoinEmpire.CallOpts)
}

// LatestSample is a free data retrieval call binding the contract method 0xc595d6b9.
//
// Solidity: function latest_sample() constant returns(uint256)
func (_CoinEmpire *CoinEmpireCallerSession) LatestSample() (*big.Int, error) {
	return _CoinEmpire.Contract.LatestSample(&_CoinEmpire.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() constant returns(address)
func (_CoinEmpire *CoinEmpireCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _CoinEmpire.contract.Call(opts, out, "owner")
	return *ret0, err
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() constant returns(address)
func (_CoinEmpire *CoinEmpireSession) Owner() (common.Address, error) {
	return _CoinEmpire.Contract.Owner(&_CoinEmpire.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() constant returns(address)
func (_CoinEmpire *CoinEmpireCallerSession) Owner() (common.Address, error) {
	return _CoinEmpire.Contract.Owner(&_CoinEmpire.CallOpts)
}

// Paused is a free data retrieval call binding the contract method 0x5c975abb.
//
// Solidity: function paused() constant returns(bool)
func (_CoinEmpire *CoinEmpireCaller) Paused(opts *bind.CallOpts) (bool, error) {
	var (
		ret0 = new(bool)
	)
	out := ret0
	err := _CoinEmpire.contract.Call(opts, out, "paused")
	return *ret0, err
}

// Paused is a free data retrieval call binding the contract method 0x5c975abb.
//
// Solidity: function paused() constant returns(bool)
func (_CoinEmpire *CoinEmpireSession) Paused() (bool, error) {
	return _CoinEmpire.Contract.Paused(&_CoinEmpire.CallOpts)
}

// Paused is a free data retrieval call binding the contract method 0x5c975abb.
//
// Solidity: function paused() constant returns(bool)
func (_CoinEmpire *CoinEmpireCallerSession) Paused() (bool, error) {
	return _CoinEmpire.Contract.Paused(&_CoinEmpire.CallOpts)
}

// Pool is a free data retrieval call binding the contract method 0x16f0115b.
//
// Solidity: function pool() constant returns(uint256)
func (_CoinEmpire *CoinEmpireCaller) Pool(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _CoinEmpire.contract.Call(opts, out, "pool")
	return *ret0, err
}

// Pool is a free data retrieval call binding the contract method 0x16f0115b.
//
// Solidity: function pool() constant returns(uint256)
func (_CoinEmpire *CoinEmpireSession) Pool() (*big.Int, error) {
	return _CoinEmpire.Contract.Pool(&_CoinEmpire.CallOpts)
}

// Pool is a free data retrieval call binding the contract method 0x16f0115b.
//
// Solidity: function pool() constant returns(uint256)
func (_CoinEmpire *CoinEmpireCallerSession) Pool() (*big.Int, error) {
	return _CoinEmpire.Contract.Pool(&_CoinEmpire.CallOpts)
}

// ProfitOf is a free data retrieval call binding the contract method 0x54198ce9.
//
// Solidity: function profitOf(address ) constant returns(uint256)
func (_CoinEmpire *CoinEmpireCaller) ProfitOf(opts *bind.CallOpts, arg0 common.Address) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _CoinEmpire.contract.Call(opts, out, "profitOf", arg0)
	return *ret0, err
}

// ProfitOf is a free data retrieval call binding the contract method 0x54198ce9.
//
// Solidity: function profitOf(address ) constant returns(uint256)
func (_CoinEmpire *CoinEmpireSession) ProfitOf(arg0 common.Address) (*big.Int, error) {
	return _CoinEmpire.Contract.ProfitOf(&_CoinEmpire.CallOpts, arg0)
}

// ProfitOf is a free data retrieval call binding the contract method 0x54198ce9.
//
// Solidity: function profitOf(address ) constant returns(uint256)
func (_CoinEmpire *CoinEmpireCallerSession) ProfitOf(arg0 common.Address) (*big.Int, error) {
	return _CoinEmpire.Contract.ProfitOf(&_CoinEmpire.CallOpts, arg0)
}

// QueryString is a free data retrieval call binding the contract method 0x8d1f5f17.
//
// Solidity: function query_string() constant returns(string)
func (_CoinEmpire *CoinEmpireCaller) QueryString(opts *bind.CallOpts) (string, error) {
	var (
		ret0 = new(string)
	)
	out := ret0
	err := _CoinEmpire.contract.Call(opts, out, "query_string")
	return *ret0, err
}

// QueryString is a free data retrieval call binding the contract method 0x8d1f5f17.
//
// Solidity: function query_string() constant returns(string)
func (_CoinEmpire *CoinEmpireSession) QueryString() (string, error) {
	return _CoinEmpire.Contract.QueryString(&_CoinEmpire.CallOpts)
}

// QueryString is a free data retrieval call binding the contract method 0x8d1f5f17.
//
// Solidity: function query_string() constant returns(string)
func (_CoinEmpire *CoinEmpireCallerSession) QueryString() (string, error) {
	return _CoinEmpire.Contract.QueryString(&_CoinEmpire.CallOpts)
}

// ResellTickets is a free data retrieval call binding the contract method 0x330a6b33.
//
// Solidity: function resell_tickets(address ) constant returns(uint256)
func (_CoinEmpire *CoinEmpireCaller) ResellTickets(opts *bind.CallOpts, arg0 common.Address) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _CoinEmpire.contract.Call(opts, out, "resell_tickets", arg0)
	return *ret0, err
}

// ResellTickets is a free data retrieval call binding the contract method 0x330a6b33.
//
// Solidity: function resell_tickets(address ) constant returns(uint256)
func (_CoinEmpire *CoinEmpireSession) ResellTickets(arg0 common.Address) (*big.Int, error) {
	return _CoinEmpire.Contract.ResellTickets(&_CoinEmpire.CallOpts, arg0)
}

// ResellTickets is a free data retrieval call binding the contract method 0x330a6b33.
//
// Solidity: function resell_tickets(address ) constant returns(uint256)
func (_CoinEmpire *CoinEmpireCallerSession) ResellTickets(arg0 common.Address) (*big.Int, error) {
	return _CoinEmpire.Contract.ResellTickets(&_CoinEmpire.CallOpts, arg0)
}

// SlotToEarnings is a free data retrieval call binding the contract method 0xa4ac25bf.
//
// Solidity: function slot_to_earnings(uint256 ) constant returns(uint256)
func (_CoinEmpire *CoinEmpireCaller) SlotToEarnings(opts *bind.CallOpts, arg0 *big.Int) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _CoinEmpire.contract.Call(opts, out, "slot_to_earnings", arg0)
	return *ret0, err
}

// SlotToEarnings is a free data retrieval call binding the contract method 0xa4ac25bf.
//
// Solidity: function slot_to_earnings(uint256 ) constant returns(uint256)
func (_CoinEmpire *CoinEmpireSession) SlotToEarnings(arg0 *big.Int) (*big.Int, error) {
	return _CoinEmpire.Contract.SlotToEarnings(&_CoinEmpire.CallOpts, arg0)
}

// SlotToEarnings is a free data retrieval call binding the contract method 0xa4ac25bf.
//
// Solidity: function slot_to_earnings(uint256 ) constant returns(uint256)
func (_CoinEmpire *CoinEmpireCallerSession) SlotToEarnings(arg0 *big.Int) (*big.Int, error) {
	return _CoinEmpire.Contract.SlotToEarnings(&_CoinEmpire.CallOpts, arg0)
}

// SlotToOwner is a free data retrieval call binding the contract method 0x179a8d3a.
//
// Solidity: function slot_to_owner(uint256 ) constant returns(address)
func (_CoinEmpire *CoinEmpireCaller) SlotToOwner(opts *bind.CallOpts, arg0 *big.Int) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _CoinEmpire.contract.Call(opts, out, "slot_to_owner", arg0)
	return *ret0, err
}

// SlotToOwner is a free data retrieval call binding the contract method 0x179a8d3a.
//
// Solidity: function slot_to_owner(uint256 ) constant returns(address)
func (_CoinEmpire *CoinEmpireSession) SlotToOwner(arg0 *big.Int) (common.Address, error) {
	return _CoinEmpire.Contract.SlotToOwner(&_CoinEmpire.CallOpts, arg0)
}

// SlotToOwner is a free data retrieval call binding the contract method 0x179a8d3a.
//
// Solidity: function slot_to_owner(uint256 ) constant returns(address)
func (_CoinEmpire *CoinEmpireCallerSession) SlotToOwner(arg0 *big.Int) (common.Address, error) {
	return _CoinEmpire.Contract.SlotToOwner(&_CoinEmpire.CallOpts, arg0)
}

// SlotToPrice is a free data retrieval call binding the contract method 0x7c9d4527.
//
// Solidity: function slot_to_price(uint256 ) constant returns(uint256)
func (_CoinEmpire *CoinEmpireCaller) SlotToPrice(opts *bind.CallOpts, arg0 *big.Int) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _CoinEmpire.contract.Call(opts, out, "slot_to_price", arg0)
	return *ret0, err
}

// SlotToPrice is a free data retrieval call binding the contract method 0x7c9d4527.
//
// Solidity: function slot_to_price(uint256 ) constant returns(uint256)
func (_CoinEmpire *CoinEmpireSession) SlotToPrice(arg0 *big.Int) (*big.Int, error) {
	return _CoinEmpire.Contract.SlotToPrice(&_CoinEmpire.CallOpts, arg0)
}

// SlotToPrice is a free data retrieval call binding the contract method 0x7c9d4527.
//
// Solidity: function slot_to_price(uint256 ) constant returns(uint256)
func (_CoinEmpire *CoinEmpireCallerSession) SlotToPrice(arg0 *big.Int) (*big.Int, error) {
	return _CoinEmpire.Contract.SlotToPrice(&_CoinEmpire.CallOpts, arg0)
}

// Callback is a paid mutator transaction binding the contract method 0x38bbfa50.
//
// Solidity: function __callback(bytes32 _queryId, string _result, bytes _proof) returns()
func (_CoinEmpire *CoinEmpireTransactor) Callback(opts *bind.TransactOpts, _queryId [32]byte, _result string, _proof []byte) (*types.Transaction, error) {
	return _CoinEmpire.contract.Transact(opts, "__callback", _queryId, _result, _proof)
}

// Callback is a paid mutator transaction binding the contract method 0x38bbfa50.
//
// Solidity: function __callback(bytes32 _queryId, string _result, bytes _proof) returns()
func (_CoinEmpire *CoinEmpireSession) Callback(_queryId [32]byte, _result string, _proof []byte) (*types.Transaction, error) {
	return _CoinEmpire.Contract.Callback(&_CoinEmpire.TransactOpts, _queryId, _result, _proof)
}

// Callback is a paid mutator transaction binding the contract method 0x38bbfa50.
//
// Solidity: function __callback(bytes32 _queryId, string _result, bytes _proof) returns()
func (_CoinEmpire *CoinEmpireTransactorSession) Callback(_queryId [32]byte, _result string, _proof []byte) (*types.Transaction, error) {
	return _CoinEmpire.Contract.Callback(&_CoinEmpire.TransactOpts, _queryId, _result, _proof)
}

// AddPauser is a paid mutator transaction binding the contract method 0x82dc1ec4.
//
// Solidity: function addPauser(address account) returns()
func (_CoinEmpire *CoinEmpireTransactor) AddPauser(opts *bind.TransactOpts, account common.Address) (*types.Transaction, error) {
	return _CoinEmpire.contract.Transact(opts, "addPauser", account)
}

// AddPauser is a paid mutator transaction binding the contract method 0x82dc1ec4.
//
// Solidity: function addPauser(address account) returns()
func (_CoinEmpire *CoinEmpireSession) AddPauser(account common.Address) (*types.Transaction, error) {
	return _CoinEmpire.Contract.AddPauser(&_CoinEmpire.TransactOpts, account)
}

// AddPauser is a paid mutator transaction binding the contract method 0x82dc1ec4.
//
// Solidity: function addPauser(address account) returns()
func (_CoinEmpire *CoinEmpireTransactorSession) AddPauser(account common.Address) (*types.Transaction, error) {
	return _CoinEmpire.Contract.AddPauser(&_CoinEmpire.TransactOpts, account)
}

// BuySlots is a paid mutator transaction binding the contract method 0x47513621.
//
// Solidity: function buySlots(uint256[] prices, uint8[] tiers) returns()
func (_CoinEmpire *CoinEmpireTransactor) BuySlots(opts *bind.TransactOpts, prices []*big.Int, tiers []uint8) (*types.Transaction, error) {
	return _CoinEmpire.contract.Transact(opts, "buySlots", prices, tiers)
}

// BuySlots is a paid mutator transaction binding the contract method 0x47513621.
//
// Solidity: function buySlots(uint256[] prices, uint8[] tiers) returns()
func (_CoinEmpire *CoinEmpireSession) BuySlots(prices []*big.Int, tiers []uint8) (*types.Transaction, error) {
	return _CoinEmpire.Contract.BuySlots(&_CoinEmpire.TransactOpts, prices, tiers)
}

// BuySlots is a paid mutator transaction binding the contract method 0x47513621.
//
// Solidity: function buySlots(uint256[] prices, uint8[] tiers) returns()
func (_CoinEmpire *CoinEmpireTransactorSession) BuySlots(prices []*big.Int, tiers []uint8) (*types.Transaction, error) {
	return _CoinEmpire.Contract.BuySlots(&_CoinEmpire.TransactOpts, prices, tiers)
}

// CheckPrice is a paid mutator transaction binding the contract method 0x8fc3047d.
//
// Solidity: function checkPrice() returns(uint256)
func (_CoinEmpire *CoinEmpireTransactor) CheckPrice(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _CoinEmpire.contract.Transact(opts, "checkPrice")
}

// CheckPrice is a paid mutator transaction binding the contract method 0x8fc3047d.
//
// Solidity: function checkPrice() returns(uint256)
func (_CoinEmpire *CoinEmpireSession) CheckPrice() (*types.Transaction, error) {
	return _CoinEmpire.Contract.CheckPrice(&_CoinEmpire.TransactOpts)
}

// CheckPrice is a paid mutator transaction binding the contract method 0x8fc3047d.
//
// Solidity: function checkPrice() returns(uint256)
func (_CoinEmpire *CoinEmpireTransactorSession) CheckPrice() (*types.Transaction, error) {
	return _CoinEmpire.Contract.CheckPrice(&_CoinEmpire.TransactOpts)
}

// Destroy is a paid mutator transaction binding the contract method 0x83197ef0.
//
// Solidity: function destroy() returns()
func (_CoinEmpire *CoinEmpireTransactor) Destroy(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _CoinEmpire.contract.Transact(opts, "destroy")
}

// Destroy is a paid mutator transaction binding the contract method 0x83197ef0.
//
// Solidity: function destroy() returns()
func (_CoinEmpire *CoinEmpireSession) Destroy() (*types.Transaction, error) {
	return _CoinEmpire.Contract.Destroy(&_CoinEmpire.TransactOpts)
}

// Destroy is a paid mutator transaction binding the contract method 0x83197ef0.
//
// Solidity: function destroy() returns()
func (_CoinEmpire *CoinEmpireTransactorSession) Destroy() (*types.Transaction, error) {
	return _CoinEmpire.Contract.Destroy(&_CoinEmpire.TransactOpts)
}

// NewSample is a paid mutator transaction binding the contract method 0x22d61869.
//
// Solidity: function newSample() returns()
func (_CoinEmpire *CoinEmpireTransactor) NewSample(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _CoinEmpire.contract.Transact(opts, "newSample")
}

// NewSample is a paid mutator transaction binding the contract method 0x22d61869.
//
// Solidity: function newSample() returns()
func (_CoinEmpire *CoinEmpireSession) NewSample() (*types.Transaction, error) {
	return _CoinEmpire.Contract.NewSample(&_CoinEmpire.TransactOpts)
}

// NewSample is a paid mutator transaction binding the contract method 0x22d61869.
//
// Solidity: function newSample() returns()
func (_CoinEmpire *CoinEmpireTransactorSession) NewSample() (*types.Transaction, error) {
	return _CoinEmpire.Contract.NewSample(&_CoinEmpire.TransactOpts)
}

// Pause is a paid mutator transaction binding the contract method 0x8456cb59.
//
// Solidity: function pause() returns()
func (_CoinEmpire *CoinEmpireTransactor) Pause(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _CoinEmpire.contract.Transact(opts, "pause")
}

// Pause is a paid mutator transaction binding the contract method 0x8456cb59.
//
// Solidity: function pause() returns()
func (_CoinEmpire *CoinEmpireSession) Pause() (*types.Transaction, error) {
	return _CoinEmpire.Contract.Pause(&_CoinEmpire.TransactOpts)
}

// Pause is a paid mutator transaction binding the contract method 0x8456cb59.
//
// Solidity: function pause() returns()
func (_CoinEmpire *CoinEmpireTransactorSession) Pause() (*types.Transaction, error) {
	return _CoinEmpire.Contract.Pause(&_CoinEmpire.TransactOpts)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_CoinEmpire *CoinEmpireTransactor) RenounceOwnership(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _CoinEmpire.contract.Transact(opts, "renounceOwnership")
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_CoinEmpire *CoinEmpireSession) RenounceOwnership() (*types.Transaction, error) {
	return _CoinEmpire.Contract.RenounceOwnership(&_CoinEmpire.TransactOpts)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_CoinEmpire *CoinEmpireTransactorSession) RenounceOwnership() (*types.Transaction, error) {
	return _CoinEmpire.Contract.RenounceOwnership(&_CoinEmpire.TransactOpts)
}

// RenouncePauser is a paid mutator transaction binding the contract method 0x6ef8d66d.
//
// Solidity: function renouncePauser() returns()
func (_CoinEmpire *CoinEmpireTransactor) RenouncePauser(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _CoinEmpire.contract.Transact(opts, "renouncePauser")
}

// RenouncePauser is a paid mutator transaction binding the contract method 0x6ef8d66d.
//
// Solidity: function renouncePauser() returns()
func (_CoinEmpire *CoinEmpireSession) RenouncePauser() (*types.Transaction, error) {
	return _CoinEmpire.Contract.RenouncePauser(&_CoinEmpire.TransactOpts)
}

// RenouncePauser is a paid mutator transaction binding the contract method 0x6ef8d66d.
//
// Solidity: function renouncePauser() returns()
func (_CoinEmpire *CoinEmpireTransactorSession) RenouncePauser() (*types.Transaction, error) {
	return _CoinEmpire.Contract.RenouncePauser(&_CoinEmpire.TransactOpts)
}

// SellSlots is a paid mutator transaction binding the contract method 0x344b7887.
//
// Solidity: function sellSlots(uint256[] prices, uint8[] tiers) returns()
func (_CoinEmpire *CoinEmpireTransactor) SellSlots(opts *bind.TransactOpts, prices []*big.Int, tiers []uint8) (*types.Transaction, error) {
	return _CoinEmpire.contract.Transact(opts, "sellSlots", prices, tiers)
}

// SellSlots is a paid mutator transaction binding the contract method 0x344b7887.
//
// Solidity: function sellSlots(uint256[] prices, uint8[] tiers) returns()
func (_CoinEmpire *CoinEmpireSession) SellSlots(prices []*big.Int, tiers []uint8) (*types.Transaction, error) {
	return _CoinEmpire.Contract.SellSlots(&_CoinEmpire.TransactOpts, prices, tiers)
}

// SellSlots is a paid mutator transaction binding the contract method 0x344b7887.
//
// Solidity: function sellSlots(uint256[] prices, uint8[] tiers) returns()
func (_CoinEmpire *CoinEmpireTransactorSession) SellSlots(prices []*big.Int, tiers []uint8) (*types.Transaction, error) {
	return _CoinEmpire.Contract.SellSlots(&_CoinEmpire.TransactOpts, prices, tiers)
}

// SetConfigValue is a paid mutator transaction binding the contract method 0xb7671acd.
//
// Solidity: function setConfigValue(bytes32 name, uint256 new_value) returns()
func (_CoinEmpire *CoinEmpireTransactor) SetConfigValue(opts *bind.TransactOpts, name [32]byte, new_value *big.Int) (*types.Transaction, error) {
	return _CoinEmpire.contract.Transact(opts, "setConfigValue", name, new_value)
}

// SetConfigValue is a paid mutator transaction binding the contract method 0xb7671acd.
//
// Solidity: function setConfigValue(bytes32 name, uint256 new_value) returns()
func (_CoinEmpire *CoinEmpireSession) SetConfigValue(name [32]byte, new_value *big.Int) (*types.Transaction, error) {
	return _CoinEmpire.Contract.SetConfigValue(&_CoinEmpire.TransactOpts, name, new_value)
}

// SetConfigValue is a paid mutator transaction binding the contract method 0xb7671acd.
//
// Solidity: function setConfigValue(bytes32 name, uint256 new_value) returns()
func (_CoinEmpire *CoinEmpireTransactorSession) SetConfigValue(name [32]byte, new_value *big.Int) (*types.Transaction, error) {
	return _CoinEmpire.Contract.SetConfigValue(&_CoinEmpire.TransactOpts, name, new_value)
}

// SetCutDestination is a paid mutator transaction binding the contract method 0xaad74722.
//
// Solidity: function setCutDestination(address new_destination) returns()
func (_CoinEmpire *CoinEmpireTransactor) SetCutDestination(opts *bind.TransactOpts, new_destination common.Address) (*types.Transaction, error) {
	return _CoinEmpire.contract.Transact(opts, "setCutDestination", new_destination)
}

// SetCutDestination is a paid mutator transaction binding the contract method 0xaad74722.
//
// Solidity: function setCutDestination(address new_destination) returns()
func (_CoinEmpire *CoinEmpireSession) SetCutDestination(new_destination common.Address) (*types.Transaction, error) {
	return _CoinEmpire.Contract.SetCutDestination(&_CoinEmpire.TransactOpts, new_destination)
}

// SetCutDestination is a paid mutator transaction binding the contract method 0xaad74722.
//
// Solidity: function setCutDestination(address new_destination) returns()
func (_CoinEmpire *CoinEmpireTransactorSession) SetCutDestination(new_destination common.Address) (*types.Transaction, error) {
	return _CoinEmpire.Contract.SetCutDestination(&_CoinEmpire.TransactOpts, new_destination)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_CoinEmpire *CoinEmpireTransactor) TransferOwnership(opts *bind.TransactOpts, newOwner common.Address) (*types.Transaction, error) {
	return _CoinEmpire.contract.Transact(opts, "transferOwnership", newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_CoinEmpire *CoinEmpireSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _CoinEmpire.Contract.TransferOwnership(&_CoinEmpire.TransactOpts, newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_CoinEmpire *CoinEmpireTransactorSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _CoinEmpire.Contract.TransferOwnership(&_CoinEmpire.TransactOpts, newOwner)
}

// Unpause is a paid mutator transaction binding the contract method 0x3f4ba83a.
//
// Solidity: function unpause() returns()
func (_CoinEmpire *CoinEmpireTransactor) Unpause(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _CoinEmpire.contract.Transact(opts, "unpause")
}

// Unpause is a paid mutator transaction binding the contract method 0x3f4ba83a.
//
// Solidity: function unpause() returns()
func (_CoinEmpire *CoinEmpireSession) Unpause() (*types.Transaction, error) {
	return _CoinEmpire.Contract.Unpause(&_CoinEmpire.TransactOpts)
}

// Unpause is a paid mutator transaction binding the contract method 0x3f4ba83a.
//
// Solidity: function unpause() returns()
func (_CoinEmpire *CoinEmpireTransactorSession) Unpause() (*types.Transaction, error) {
	return _CoinEmpire.Contract.Unpause(&_CoinEmpire.TransactOpts)
}

// WithdrawHouse is a paid mutator transaction binding the contract method 0x23218411.
//
// Solidity: function withdrawHouse(uint256 amount) returns()
func (_CoinEmpire *CoinEmpireTransactor) WithdrawHouse(opts *bind.TransactOpts, amount *big.Int) (*types.Transaction, error) {
	return _CoinEmpire.contract.Transact(opts, "withdrawHouse", amount)
}

// WithdrawHouse is a paid mutator transaction binding the contract method 0x23218411.
//
// Solidity: function withdrawHouse(uint256 amount) returns()
func (_CoinEmpire *CoinEmpireSession) WithdrawHouse(amount *big.Int) (*types.Transaction, error) {
	return _CoinEmpire.Contract.WithdrawHouse(&_CoinEmpire.TransactOpts, amount)
}

// WithdrawHouse is a paid mutator transaction binding the contract method 0x23218411.
//
// Solidity: function withdrawHouse(uint256 amount) returns()
func (_CoinEmpire *CoinEmpireTransactorSession) WithdrawHouse(amount *big.Int) (*types.Transaction, error) {
	return _CoinEmpire.Contract.WithdrawHouse(&_CoinEmpire.TransactOpts, amount)
}

// WithdrawWallet is a paid mutator transaction binding the contract method 0xbe87ab3f.
//
// Solidity: function withdrawWallet(uint256 amount) returns()
func (_CoinEmpire *CoinEmpireTransactor) WithdrawWallet(opts *bind.TransactOpts, amount *big.Int) (*types.Transaction, error) {
	return _CoinEmpire.contract.Transact(opts, "withdrawWallet", amount)
}

// WithdrawWallet is a paid mutator transaction binding the contract method 0xbe87ab3f.
//
// Solidity: function withdrawWallet(uint256 amount) returns()
func (_CoinEmpire *CoinEmpireSession) WithdrawWallet(amount *big.Int) (*types.Transaction, error) {
	return _CoinEmpire.Contract.WithdrawWallet(&_CoinEmpire.TransactOpts, amount)
}

// WithdrawWallet is a paid mutator transaction binding the contract method 0xbe87ab3f.
//
// Solidity: function withdrawWallet(uint256 amount) returns()
func (_CoinEmpire *CoinEmpireTransactorSession) WithdrawWallet(amount *big.Int) (*types.Transaction, error) {
	return _CoinEmpire.Contract.WithdrawWallet(&_CoinEmpire.TransactOpts, amount)
}

// CoinEmpireConfigChangedIterator is returned from FilterConfigChanged and is used to iterate over the raw logs and unpacked data for ConfigChanged events raised by the CoinEmpire contract.
type CoinEmpireConfigChangedIterator struct {
	Event *CoinEmpireConfigChanged // Event containing the contract specifics and raw log

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
func (it *CoinEmpireConfigChangedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(CoinEmpireConfigChanged)
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
		it.Event = new(CoinEmpireConfigChanged)
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
func (it *CoinEmpireConfigChangedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *CoinEmpireConfigChangedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// CoinEmpireConfigChanged represents a ConfigChanged event raised by the CoinEmpire contract.
type CoinEmpireConfigChanged struct {
	Name     [32]byte
	NewValue *big.Int
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterConfigChanged is a free log retrieval operation binding the contract event 0x867716e28221fc5767f7ac9e2623573aa2b682e2b925334db66a0753751708d6.
//
// Solidity: event ConfigChanged(bytes32 name, uint256 new_value)
func (_CoinEmpire *CoinEmpireFilterer) FilterConfigChanged(opts *bind.FilterOpts) (*CoinEmpireConfigChangedIterator, error) {

	logs, sub, err := _CoinEmpire.contract.FilterLogs(opts, "ConfigChanged")
	if err != nil {
		return nil, err
	}
	return &CoinEmpireConfigChangedIterator{contract: _CoinEmpire.contract, event: "ConfigChanged", logs: logs, sub: sub}, nil
}

// WatchConfigChanged is a free log subscription operation binding the contract event 0x867716e28221fc5767f7ac9e2623573aa2b682e2b925334db66a0753751708d6.
//
// Solidity: event ConfigChanged(bytes32 name, uint256 new_value)
func (_CoinEmpire *CoinEmpireFilterer) WatchConfigChanged(opts *bind.WatchOpts, sink chan<- *CoinEmpireConfigChanged) (event.Subscription, error) {

	logs, sub, err := _CoinEmpire.contract.WatchLogs(opts, "ConfigChanged")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(CoinEmpireConfigChanged)
				if err := _CoinEmpire.contract.UnpackLog(event, "ConfigChanged", log); err != nil {
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

// CoinEmpireOwnershipTransferredIterator is returned from FilterOwnershipTransferred and is used to iterate over the raw logs and unpacked data for OwnershipTransferred events raised by the CoinEmpire contract.
type CoinEmpireOwnershipTransferredIterator struct {
	Event *CoinEmpireOwnershipTransferred // Event containing the contract specifics and raw log

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
func (it *CoinEmpireOwnershipTransferredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(CoinEmpireOwnershipTransferred)
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
		it.Event = new(CoinEmpireOwnershipTransferred)
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
func (it *CoinEmpireOwnershipTransferredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *CoinEmpireOwnershipTransferredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// CoinEmpireOwnershipTransferred represents a OwnershipTransferred event raised by the CoinEmpire contract.
type CoinEmpireOwnershipTransferred struct {
	PreviousOwner common.Address
	NewOwner      common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterOwnershipTransferred is a free log retrieval operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_CoinEmpire *CoinEmpireFilterer) FilterOwnershipTransferred(opts *bind.FilterOpts, previousOwner []common.Address, newOwner []common.Address) (*CoinEmpireOwnershipTransferredIterator, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _CoinEmpire.contract.FilterLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return &CoinEmpireOwnershipTransferredIterator{contract: _CoinEmpire.contract, event: "OwnershipTransferred", logs: logs, sub: sub}, nil
}

// WatchOwnershipTransferred is a free log subscription operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_CoinEmpire *CoinEmpireFilterer) WatchOwnershipTransferred(opts *bind.WatchOpts, sink chan<- *CoinEmpireOwnershipTransferred, previousOwner []common.Address, newOwner []common.Address) (event.Subscription, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _CoinEmpire.contract.WatchLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(CoinEmpireOwnershipTransferred)
				if err := _CoinEmpire.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
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

// CoinEmpirePausedIterator is returned from FilterPaused and is used to iterate over the raw logs and unpacked data for Paused events raised by the CoinEmpire contract.
type CoinEmpirePausedIterator struct {
	Event *CoinEmpirePaused // Event containing the contract specifics and raw log

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
func (it *CoinEmpirePausedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(CoinEmpirePaused)
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
		it.Event = new(CoinEmpirePaused)
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
func (it *CoinEmpirePausedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *CoinEmpirePausedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// CoinEmpirePaused represents a Paused event raised by the CoinEmpire contract.
type CoinEmpirePaused struct {
	Account common.Address
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterPaused is a free log retrieval operation binding the contract event 0x62e78cea01bee320cd4e420270b5ea74000d11b0c9f74754ebdbfc544b05a258.
//
// Solidity: event Paused(address account)
func (_CoinEmpire *CoinEmpireFilterer) FilterPaused(opts *bind.FilterOpts) (*CoinEmpirePausedIterator, error) {

	logs, sub, err := _CoinEmpire.contract.FilterLogs(opts, "Paused")
	if err != nil {
		return nil, err
	}
	return &CoinEmpirePausedIterator{contract: _CoinEmpire.contract, event: "Paused", logs: logs, sub: sub}, nil
}

// WatchPaused is a free log subscription operation binding the contract event 0x62e78cea01bee320cd4e420270b5ea74000d11b0c9f74754ebdbfc544b05a258.
//
// Solidity: event Paused(address account)
func (_CoinEmpire *CoinEmpireFilterer) WatchPaused(opts *bind.WatchOpts, sink chan<- *CoinEmpirePaused) (event.Subscription, error) {

	logs, sub, err := _CoinEmpire.contract.WatchLogs(opts, "Paused")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(CoinEmpirePaused)
				if err := _CoinEmpire.contract.UnpackLog(event, "Paused", log); err != nil {
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

// CoinEmpirePauserAddedIterator is returned from FilterPauserAdded and is used to iterate over the raw logs and unpacked data for PauserAdded events raised by the CoinEmpire contract.
type CoinEmpirePauserAddedIterator struct {
	Event *CoinEmpirePauserAdded // Event containing the contract specifics and raw log

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
func (it *CoinEmpirePauserAddedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(CoinEmpirePauserAdded)
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
		it.Event = new(CoinEmpirePauserAdded)
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
func (it *CoinEmpirePauserAddedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *CoinEmpirePauserAddedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// CoinEmpirePauserAdded represents a PauserAdded event raised by the CoinEmpire contract.
type CoinEmpirePauserAdded struct {
	Account common.Address
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterPauserAdded is a free log retrieval operation binding the contract event 0x6719d08c1888103bea251a4ed56406bd0c3e69723c8a1686e017e7bbe159b6f8.
//
// Solidity: event PauserAdded(address indexed account)
func (_CoinEmpire *CoinEmpireFilterer) FilterPauserAdded(opts *bind.FilterOpts, account []common.Address) (*CoinEmpirePauserAddedIterator, error) {

	var accountRule []interface{}
	for _, accountItem := range account {
		accountRule = append(accountRule, accountItem)
	}

	logs, sub, err := _CoinEmpire.contract.FilterLogs(opts, "PauserAdded", accountRule)
	if err != nil {
		return nil, err
	}
	return &CoinEmpirePauserAddedIterator{contract: _CoinEmpire.contract, event: "PauserAdded", logs: logs, sub: sub}, nil
}

// WatchPauserAdded is a free log subscription operation binding the contract event 0x6719d08c1888103bea251a4ed56406bd0c3e69723c8a1686e017e7bbe159b6f8.
//
// Solidity: event PauserAdded(address indexed account)
func (_CoinEmpire *CoinEmpireFilterer) WatchPauserAdded(opts *bind.WatchOpts, sink chan<- *CoinEmpirePauserAdded, account []common.Address) (event.Subscription, error) {

	var accountRule []interface{}
	for _, accountItem := range account {
		accountRule = append(accountRule, accountItem)
	}

	logs, sub, err := _CoinEmpire.contract.WatchLogs(opts, "PauserAdded", accountRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(CoinEmpirePauserAdded)
				if err := _CoinEmpire.contract.UnpackLog(event, "PauserAdded", log); err != nil {
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

// CoinEmpirePauserRemovedIterator is returned from FilterPauserRemoved and is used to iterate over the raw logs and unpacked data for PauserRemoved events raised by the CoinEmpire contract.
type CoinEmpirePauserRemovedIterator struct {
	Event *CoinEmpirePauserRemoved // Event containing the contract specifics and raw log

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
func (it *CoinEmpirePauserRemovedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(CoinEmpirePauserRemoved)
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
		it.Event = new(CoinEmpirePauserRemoved)
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
func (it *CoinEmpirePauserRemovedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *CoinEmpirePauserRemovedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// CoinEmpirePauserRemoved represents a PauserRemoved event raised by the CoinEmpire contract.
type CoinEmpirePauserRemoved struct {
	Account common.Address
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterPauserRemoved is a free log retrieval operation binding the contract event 0xcd265ebaf09df2871cc7bd4133404a235ba12eff2041bb89d9c714a2621c7c7e.
//
// Solidity: event PauserRemoved(address indexed account)
func (_CoinEmpire *CoinEmpireFilterer) FilterPauserRemoved(opts *bind.FilterOpts, account []common.Address) (*CoinEmpirePauserRemovedIterator, error) {

	var accountRule []interface{}
	for _, accountItem := range account {
		accountRule = append(accountRule, accountItem)
	}

	logs, sub, err := _CoinEmpire.contract.FilterLogs(opts, "PauserRemoved", accountRule)
	if err != nil {
		return nil, err
	}
	return &CoinEmpirePauserRemovedIterator{contract: _CoinEmpire.contract, event: "PauserRemoved", logs: logs, sub: sub}, nil
}

// WatchPauserRemoved is a free log subscription operation binding the contract event 0xcd265ebaf09df2871cc7bd4133404a235ba12eff2041bb89d9c714a2621c7c7e.
//
// Solidity: event PauserRemoved(address indexed account)
func (_CoinEmpire *CoinEmpireFilterer) WatchPauserRemoved(opts *bind.WatchOpts, sink chan<- *CoinEmpirePauserRemoved, account []common.Address) (event.Subscription, error) {

	var accountRule []interface{}
	for _, accountItem := range account {
		accountRule = append(accountRule, accountItem)
	}

	logs, sub, err := _CoinEmpire.contract.WatchLogs(opts, "PauserRemoved", accountRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(CoinEmpirePauserRemoved)
				if err := _CoinEmpire.contract.UnpackLog(event, "PauserRemoved", log); err != nil {
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

// CoinEmpireSamplingPriceEndedIterator is returned from FilterSamplingPriceEnded and is used to iterate over the raw logs and unpacked data for SamplingPriceEnded events raised by the CoinEmpire contract.
type CoinEmpireSamplingPriceEndedIterator struct {
	Event *CoinEmpireSamplingPriceEnded // Event containing the contract specifics and raw log

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
func (it *CoinEmpireSamplingPriceEndedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(CoinEmpireSamplingPriceEnded)
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
		it.Event = new(CoinEmpireSamplingPriceEnded)
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
func (it *CoinEmpireSamplingPriceEndedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *CoinEmpireSamplingPriceEndedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// CoinEmpireSamplingPriceEnded represents a SamplingPriceEnded event raised by the CoinEmpire contract.
type CoinEmpireSamplingPriceEnded struct {
	Price  *big.Int
	Result string
	Proof  []byte
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterSamplingPriceEnded is a free log retrieval operation binding the contract event 0x768a664c22763f8327912aa5d327af4d164f2704893bba4178a753a13537f48e.
//
// Solidity: event SamplingPriceEnded(uint256 price, string result, bytes proof)
func (_CoinEmpire *CoinEmpireFilterer) FilterSamplingPriceEnded(opts *bind.FilterOpts) (*CoinEmpireSamplingPriceEndedIterator, error) {

	logs, sub, err := _CoinEmpire.contract.FilterLogs(opts, "SamplingPriceEnded")
	if err != nil {
		return nil, err
	}
	return &CoinEmpireSamplingPriceEndedIterator{contract: _CoinEmpire.contract, event: "SamplingPriceEnded", logs: logs, sub: sub}, nil
}

// WatchSamplingPriceEnded is a free log subscription operation binding the contract event 0x768a664c22763f8327912aa5d327af4d164f2704893bba4178a753a13537f48e.
//
// Solidity: event SamplingPriceEnded(uint256 price, string result, bytes proof)
func (_CoinEmpire *CoinEmpireFilterer) WatchSamplingPriceEnded(opts *bind.WatchOpts, sink chan<- *CoinEmpireSamplingPriceEnded) (event.Subscription, error) {

	logs, sub, err := _CoinEmpire.contract.WatchLogs(opts, "SamplingPriceEnded")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(CoinEmpireSamplingPriceEnded)
				if err := _CoinEmpire.contract.UnpackLog(event, "SamplingPriceEnded", log); err != nil {
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

// CoinEmpireSamplingPriceStartedIterator is returned from FilterSamplingPriceStarted and is used to iterate over the raw logs and unpacked data for SamplingPriceStarted events raised by the CoinEmpire contract.
type CoinEmpireSamplingPriceStartedIterator struct {
	Event *CoinEmpireSamplingPriceStarted // Event containing the contract specifics and raw log

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
func (it *CoinEmpireSamplingPriceStartedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(CoinEmpireSamplingPriceStarted)
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
		it.Event = new(CoinEmpireSamplingPriceStarted)
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
func (it *CoinEmpireSamplingPriceStartedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *CoinEmpireSamplingPriceStartedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// CoinEmpireSamplingPriceStarted represents a SamplingPriceStarted event raised by the CoinEmpire contract.
type CoinEmpireSamplingPriceStarted struct {
	QueryId [32]byte
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterSamplingPriceStarted is a free log retrieval operation binding the contract event 0x151bd9f313552302c8995e9797e0e3f49881e47d0d4271ef0d12e5925bbc32de.
//
// Solidity: event SamplingPriceStarted(bytes32 query_id)
func (_CoinEmpire *CoinEmpireFilterer) FilterSamplingPriceStarted(opts *bind.FilterOpts) (*CoinEmpireSamplingPriceStartedIterator, error) {

	logs, sub, err := _CoinEmpire.contract.FilterLogs(opts, "SamplingPriceStarted")
	if err != nil {
		return nil, err
	}
	return &CoinEmpireSamplingPriceStartedIterator{contract: _CoinEmpire.contract, event: "SamplingPriceStarted", logs: logs, sub: sub}, nil
}

// WatchSamplingPriceStarted is a free log subscription operation binding the contract event 0x151bd9f313552302c8995e9797e0e3f49881e47d0d4271ef0d12e5925bbc32de.
//
// Solidity: event SamplingPriceStarted(bytes32 query_id)
func (_CoinEmpire *CoinEmpireFilterer) WatchSamplingPriceStarted(opts *bind.WatchOpts, sink chan<- *CoinEmpireSamplingPriceStarted) (event.Subscription, error) {

	logs, sub, err := _CoinEmpire.contract.WatchLogs(opts, "SamplingPriceStarted")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(CoinEmpireSamplingPriceStarted)
				if err := _CoinEmpire.contract.UnpackLog(event, "SamplingPriceStarted", log); err != nil {
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

// CoinEmpireSlotAbandonedIterator is returned from FilterSlotAbandoned and is used to iterate over the raw logs and unpacked data for SlotAbandoned events raised by the CoinEmpire contract.
type CoinEmpireSlotAbandonedIterator struct {
	Event *CoinEmpireSlotAbandoned // Event containing the contract specifics and raw log

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
func (it *CoinEmpireSlotAbandonedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(CoinEmpireSlotAbandoned)
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
		it.Event = new(CoinEmpireSlotAbandoned)
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
func (it *CoinEmpireSlotAbandonedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *CoinEmpireSlotAbandonedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// CoinEmpireSlotAbandoned represents a SlotAbandoned event raised by the CoinEmpire contract.
type CoinEmpireSlotAbandoned struct {
	Price *big.Int
	Tier  uint8
	By    common.Address
	Raw   types.Log // Blockchain specific contextual infos
}

// FilterSlotAbandoned is a free log retrieval operation binding the contract event 0x3a58ca2fc87c3d2045f75bd819cb1aea9244cdfadebc1be13fded769abf7d66c.
//
// Solidity: event SlotAbandoned(uint256 price, uint8 tier, address by)
func (_CoinEmpire *CoinEmpireFilterer) FilterSlotAbandoned(opts *bind.FilterOpts) (*CoinEmpireSlotAbandonedIterator, error) {

	logs, sub, err := _CoinEmpire.contract.FilterLogs(opts, "SlotAbandoned")
	if err != nil {
		return nil, err
	}
	return &CoinEmpireSlotAbandonedIterator{contract: _CoinEmpire.contract, event: "SlotAbandoned", logs: logs, sub: sub}, nil
}

// WatchSlotAbandoned is a free log subscription operation binding the contract event 0x3a58ca2fc87c3d2045f75bd819cb1aea9244cdfadebc1be13fded769abf7d66c.
//
// Solidity: event SlotAbandoned(uint256 price, uint8 tier, address by)
func (_CoinEmpire *CoinEmpireFilterer) WatchSlotAbandoned(opts *bind.WatchOpts, sink chan<- *CoinEmpireSlotAbandoned) (event.Subscription, error) {

	logs, sub, err := _CoinEmpire.contract.WatchLogs(opts, "SlotAbandoned")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(CoinEmpireSlotAbandoned)
				if err := _CoinEmpire.contract.UnpackLog(event, "SlotAbandoned", log); err != nil {
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

// CoinEmpireSlotPurchasedIterator is returned from FilterSlotPurchased and is used to iterate over the raw logs and unpacked data for SlotPurchased events raised by the CoinEmpire contract.
type CoinEmpireSlotPurchasedIterator struct {
	Event *CoinEmpireSlotPurchased // Event containing the contract specifics and raw log

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
func (it *CoinEmpireSlotPurchasedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(CoinEmpireSlotPurchased)
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
		it.Event = new(CoinEmpireSlotPurchased)
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
func (it *CoinEmpireSlotPurchasedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *CoinEmpireSlotPurchasedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// CoinEmpireSlotPurchased represents a SlotPurchased event raised by the CoinEmpire contract.
type CoinEmpireSlotPurchased struct {
	Price    *big.Int
	NewValue *big.Int
	Tier     uint8
	From     common.Address
	By       common.Address
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterSlotPurchased is a free log retrieval operation binding the contract event 0x34176d7031c41eaafbc895af3f7b5b4e04f92dbb724a34fbb917b1ac56f7ae53.
//
// Solidity: event SlotPurchased(uint256 price, uint256 new_value, uint8 tier, address from, address by)
func (_CoinEmpire *CoinEmpireFilterer) FilterSlotPurchased(opts *bind.FilterOpts) (*CoinEmpireSlotPurchasedIterator, error) {

	logs, sub, err := _CoinEmpire.contract.FilterLogs(opts, "SlotPurchased")
	if err != nil {
		return nil, err
	}
	return &CoinEmpireSlotPurchasedIterator{contract: _CoinEmpire.contract, event: "SlotPurchased", logs: logs, sub: sub}, nil
}

// WatchSlotPurchased is a free log subscription operation binding the contract event 0x34176d7031c41eaafbc895af3f7b5b4e04f92dbb724a34fbb917b1ac56f7ae53.
//
// Solidity: event SlotPurchased(uint256 price, uint256 new_value, uint8 tier, address from, address by)
func (_CoinEmpire *CoinEmpireFilterer) WatchSlotPurchased(opts *bind.WatchOpts, sink chan<- *CoinEmpireSlotPurchased) (event.Subscription, error) {

	logs, sub, err := _CoinEmpire.contract.WatchLogs(opts, "SlotPurchased")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(CoinEmpireSlotPurchased)
				if err := _CoinEmpire.contract.UnpackLog(event, "SlotPurchased", log); err != nil {
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

// CoinEmpireUnpausedIterator is returned from FilterUnpaused and is used to iterate over the raw logs and unpacked data for Unpaused events raised by the CoinEmpire contract.
type CoinEmpireUnpausedIterator struct {
	Event *CoinEmpireUnpaused // Event containing the contract specifics and raw log

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
func (it *CoinEmpireUnpausedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(CoinEmpireUnpaused)
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
		it.Event = new(CoinEmpireUnpaused)
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
func (it *CoinEmpireUnpausedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *CoinEmpireUnpausedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// CoinEmpireUnpaused represents a Unpaused event raised by the CoinEmpire contract.
type CoinEmpireUnpaused struct {
	Account common.Address
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterUnpaused is a free log retrieval operation binding the contract event 0x5db9ee0a495bf2e6ff9c91a7834c1ba4fdd244a5e8aa4e537bd38aeae4b073aa.
//
// Solidity: event Unpaused(address account)
func (_CoinEmpire *CoinEmpireFilterer) FilterUnpaused(opts *bind.FilterOpts) (*CoinEmpireUnpausedIterator, error) {

	logs, sub, err := _CoinEmpire.contract.FilterLogs(opts, "Unpaused")
	if err != nil {
		return nil, err
	}
	return &CoinEmpireUnpausedIterator{contract: _CoinEmpire.contract, event: "Unpaused", logs: logs, sub: sub}, nil
}

// WatchUnpaused is a free log subscription operation binding the contract event 0x5db9ee0a495bf2e6ff9c91a7834c1ba4fdd244a5e8aa4e537bd38aeae4b073aa.
//
// Solidity: event Unpaused(address account)
func (_CoinEmpire *CoinEmpireFilterer) WatchUnpaused(opts *bind.WatchOpts, sink chan<- *CoinEmpireUnpaused) (event.Subscription, error) {

	logs, sub, err := _CoinEmpire.contract.WatchLogs(opts, "Unpaused")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(CoinEmpireUnpaused)
				if err := _CoinEmpire.contract.UnpackLog(event, "Unpaused", log); err != nil {
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

// OraclizeAddrResolverIABI is the input ABI used to generate the binding from.
const OraclizeAddrResolverIABI = "[{\"constant\":false,\"inputs\":[],\"name\":\"getAddress\",\"outputs\":[{\"name\":\"_address\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]"

// OraclizeAddrResolverIBin is the compiled bytecode used for deploying new contracts.
const OraclizeAddrResolverIBin = `0x`

// DeployOraclizeAddrResolverI deploys a new Ethereum contract, binding an instance of OraclizeAddrResolverI to it.
func DeployOraclizeAddrResolverI(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *OraclizeAddrResolverI, error) {
	parsed, err := abi.JSON(strings.NewReader(OraclizeAddrResolverIABI))
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	address, tx, contract, err := bind.DeployContract(auth, parsed, common.FromHex(OraclizeAddrResolverIBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &OraclizeAddrResolverI{OraclizeAddrResolverICaller: OraclizeAddrResolverICaller{contract: contract}, OraclizeAddrResolverITransactor: OraclizeAddrResolverITransactor{contract: contract}, OraclizeAddrResolverIFilterer: OraclizeAddrResolverIFilterer{contract: contract}}, nil
}

// OraclizeAddrResolverI is an auto generated Go binding around an Ethereum contract.
type OraclizeAddrResolverI struct {
	OraclizeAddrResolverICaller     // Read-only binding to the contract
	OraclizeAddrResolverITransactor // Write-only binding to the contract
	OraclizeAddrResolverIFilterer   // Log filterer for contract events
}

// OraclizeAddrResolverICaller is an auto generated read-only Go binding around an Ethereum contract.
type OraclizeAddrResolverICaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// OraclizeAddrResolverITransactor is an auto generated write-only Go binding around an Ethereum contract.
type OraclizeAddrResolverITransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// OraclizeAddrResolverIFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type OraclizeAddrResolverIFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// OraclizeAddrResolverISession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type OraclizeAddrResolverISession struct {
	Contract     *OraclizeAddrResolverI // Generic contract binding to set the session for
	CallOpts     bind.CallOpts          // Call options to use throughout this session
	TransactOpts bind.TransactOpts      // Transaction auth options to use throughout this session
}

// OraclizeAddrResolverICallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type OraclizeAddrResolverICallerSession struct {
	Contract *OraclizeAddrResolverICaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts                // Call options to use throughout this session
}

// OraclizeAddrResolverITransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type OraclizeAddrResolverITransactorSession struct {
	Contract     *OraclizeAddrResolverITransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts                // Transaction auth options to use throughout this session
}

// OraclizeAddrResolverIRaw is an auto generated low-level Go binding around an Ethereum contract.
type OraclizeAddrResolverIRaw struct {
	Contract *OraclizeAddrResolverI // Generic contract binding to access the raw methods on
}

// OraclizeAddrResolverICallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type OraclizeAddrResolverICallerRaw struct {
	Contract *OraclizeAddrResolverICaller // Generic read-only contract binding to access the raw methods on
}

// OraclizeAddrResolverITransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type OraclizeAddrResolverITransactorRaw struct {
	Contract *OraclizeAddrResolverITransactor // Generic write-only contract binding to access the raw methods on
}

// NewOraclizeAddrResolverI creates a new instance of OraclizeAddrResolverI, bound to a specific deployed contract.
func NewOraclizeAddrResolverI(address common.Address, backend bind.ContractBackend) (*OraclizeAddrResolverI, error) {
	contract, err := bindOraclizeAddrResolverI(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &OraclizeAddrResolverI{OraclizeAddrResolverICaller: OraclizeAddrResolverICaller{contract: contract}, OraclizeAddrResolverITransactor: OraclizeAddrResolverITransactor{contract: contract}, OraclizeAddrResolverIFilterer: OraclizeAddrResolverIFilterer{contract: contract}}, nil
}

// NewOraclizeAddrResolverICaller creates a new read-only instance of OraclizeAddrResolverI, bound to a specific deployed contract.
func NewOraclizeAddrResolverICaller(address common.Address, caller bind.ContractCaller) (*OraclizeAddrResolverICaller, error) {
	contract, err := bindOraclizeAddrResolverI(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &OraclizeAddrResolverICaller{contract: contract}, nil
}

// NewOraclizeAddrResolverITransactor creates a new write-only instance of OraclizeAddrResolverI, bound to a specific deployed contract.
func NewOraclizeAddrResolverITransactor(address common.Address, transactor bind.ContractTransactor) (*OraclizeAddrResolverITransactor, error) {
	contract, err := bindOraclizeAddrResolverI(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &OraclizeAddrResolverITransactor{contract: contract}, nil
}

// NewOraclizeAddrResolverIFilterer creates a new log filterer instance of OraclizeAddrResolverI, bound to a specific deployed contract.
func NewOraclizeAddrResolverIFilterer(address common.Address, filterer bind.ContractFilterer) (*OraclizeAddrResolverIFilterer, error) {
	contract, err := bindOraclizeAddrResolverI(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &OraclizeAddrResolverIFilterer{contract: contract}, nil
}

// bindOraclizeAddrResolverI binds a generic wrapper to an already deployed contract.
func bindOraclizeAddrResolverI(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(OraclizeAddrResolverIABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_OraclizeAddrResolverI *OraclizeAddrResolverIRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _OraclizeAddrResolverI.Contract.OraclizeAddrResolverICaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_OraclizeAddrResolverI *OraclizeAddrResolverIRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _OraclizeAddrResolverI.Contract.OraclizeAddrResolverITransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_OraclizeAddrResolverI *OraclizeAddrResolverIRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _OraclizeAddrResolverI.Contract.OraclizeAddrResolverITransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_OraclizeAddrResolverI *OraclizeAddrResolverICallerRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _OraclizeAddrResolverI.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_OraclizeAddrResolverI *OraclizeAddrResolverITransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _OraclizeAddrResolverI.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_OraclizeAddrResolverI *OraclizeAddrResolverITransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _OraclizeAddrResolverI.Contract.contract.Transact(opts, method, params...)
}

// GetAddress is a paid mutator transaction binding the contract method 0x38cc4831.
//
// Solidity: function getAddress() returns(address _address)
func (_OraclizeAddrResolverI *OraclizeAddrResolverITransactor) GetAddress(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _OraclizeAddrResolverI.contract.Transact(opts, "getAddress")
}

// GetAddress is a paid mutator transaction binding the contract method 0x38cc4831.
//
// Solidity: function getAddress() returns(address _address)
func (_OraclizeAddrResolverI *OraclizeAddrResolverISession) GetAddress() (*types.Transaction, error) {
	return _OraclizeAddrResolverI.Contract.GetAddress(&_OraclizeAddrResolverI.TransactOpts)
}

// GetAddress is a paid mutator transaction binding the contract method 0x38cc4831.
//
// Solidity: function getAddress() returns(address _address)
func (_OraclizeAddrResolverI *OraclizeAddrResolverITransactorSession) GetAddress() (*types.Transaction, error) {
	return _OraclizeAddrResolverI.Contract.GetAddress(&_OraclizeAddrResolverI.TransactOpts)
}

// OraclizeIABI is the input ABI used to generate the binding from.
const OraclizeIABI = "[{\"constant\":false,\"inputs\":[{\"name\":\"_datasource\",\"type\":\"string\"},{\"name\":\"_gasLimit\",\"type\":\"uint256\"}],\"name\":\"getPrice\",\"outputs\":[{\"name\":\"_dsprice\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_datasource\",\"type\":\"string\"}],\"name\":\"getPrice\",\"outputs\":[{\"name\":\"_dsprice\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_proofType\",\"type\":\"bytes1\"}],\"name\":\"setProofType\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_timestamp\",\"type\":\"uint256\"},{\"name\":\"_datasource\",\"type\":\"string\"},{\"name\":\"_arg1\",\"type\":\"string\"},{\"name\":\"_arg2\",\"type\":\"string\"}],\"name\":\"query2\",\"outputs\":[{\"name\":\"_id\",\"type\":\"bytes32\"}],\"payable\":true,\"stateMutability\":\"payable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_timestamp\",\"type\":\"uint256\"},{\"name\":\"_datasource\",\"type\":\"string\"},{\"name\":\"_argN\",\"type\":\"bytes\"}],\"name\":\"queryN\",\"outputs\":[{\"name\":\"_id\",\"type\":\"bytes32\"}],\"payable\":true,\"stateMutability\":\"payable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_timestamp\",\"type\":\"uint256\"},{\"name\":\"_datasource\",\"type\":\"string\"},{\"name\":\"_arg1\",\"type\":\"string\"},{\"name\":\"_arg2\",\"type\":\"string\"},{\"name\":\"_gasLimit\",\"type\":\"uint256\"}],\"name\":\"query2_withGasLimit\",\"outputs\":[{\"name\":\"_id\",\"type\":\"bytes32\"}],\"payable\":true,\"stateMutability\":\"payable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"randomDS_getSessionPubKeyHash\",\"outputs\":[{\"name\":\"_sessionKeyHash\",\"type\":\"bytes32\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_timestamp\",\"type\":\"uint256\"},{\"name\":\"_datasource\",\"type\":\"string\"},{\"name\":\"_arg\",\"type\":\"string\"}],\"name\":\"query\",\"outputs\":[{\"name\":\"_id\",\"type\":\"bytes32\"}],\"payable\":true,\"stateMutability\":\"payable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"cbAddress\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_timestamp\",\"type\":\"uint256\"},{\"name\":\"_datasource\",\"type\":\"string\"},{\"name\":\"_arg\",\"type\":\"string\"},{\"name\":\"_gasLimit\",\"type\":\"uint256\"}],\"name\":\"query_withGasLimit\",\"outputs\":[{\"name\":\"_id\",\"type\":\"bytes32\"}],\"payable\":true,\"stateMutability\":\"payable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_timestamp\",\"type\":\"uint256\"},{\"name\":\"_datasource\",\"type\":\"string\"},{\"name\":\"_argN\",\"type\":\"bytes\"},{\"name\":\"_gasLimit\",\"type\":\"uint256\"}],\"name\":\"queryN_withGasLimit\",\"outputs\":[{\"name\":\"_id\",\"type\":\"bytes32\"}],\"payable\":true,\"stateMutability\":\"payable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_gasPrice\",\"type\":\"uint256\"}],\"name\":\"setCustomGasPrice\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]"

// OraclizeIBin is the compiled bytecode used for deploying new contracts.
const OraclizeIBin = `0x`

// DeployOraclizeI deploys a new Ethereum contract, binding an instance of OraclizeI to it.
func DeployOraclizeI(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *OraclizeI, error) {
	parsed, err := abi.JSON(strings.NewReader(OraclizeIABI))
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	address, tx, contract, err := bind.DeployContract(auth, parsed, common.FromHex(OraclizeIBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &OraclizeI{OraclizeICaller: OraclizeICaller{contract: contract}, OraclizeITransactor: OraclizeITransactor{contract: contract}, OraclizeIFilterer: OraclizeIFilterer{contract: contract}}, nil
}

// OraclizeI is an auto generated Go binding around an Ethereum contract.
type OraclizeI struct {
	OraclizeICaller     // Read-only binding to the contract
	OraclizeITransactor // Write-only binding to the contract
	OraclizeIFilterer   // Log filterer for contract events
}

// OraclizeICaller is an auto generated read-only Go binding around an Ethereum contract.
type OraclizeICaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// OraclizeITransactor is an auto generated write-only Go binding around an Ethereum contract.
type OraclizeITransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// OraclizeIFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type OraclizeIFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// OraclizeISession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type OraclizeISession struct {
	Contract     *OraclizeI        // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// OraclizeICallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type OraclizeICallerSession struct {
	Contract *OraclizeICaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts    // Call options to use throughout this session
}

// OraclizeITransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type OraclizeITransactorSession struct {
	Contract     *OraclizeITransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts    // Transaction auth options to use throughout this session
}

// OraclizeIRaw is an auto generated low-level Go binding around an Ethereum contract.
type OraclizeIRaw struct {
	Contract *OraclizeI // Generic contract binding to access the raw methods on
}

// OraclizeICallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type OraclizeICallerRaw struct {
	Contract *OraclizeICaller // Generic read-only contract binding to access the raw methods on
}

// OraclizeITransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type OraclizeITransactorRaw struct {
	Contract *OraclizeITransactor // Generic write-only contract binding to access the raw methods on
}

// NewOraclizeI creates a new instance of OraclizeI, bound to a specific deployed contract.
func NewOraclizeI(address common.Address, backend bind.ContractBackend) (*OraclizeI, error) {
	contract, err := bindOraclizeI(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &OraclizeI{OraclizeICaller: OraclizeICaller{contract: contract}, OraclizeITransactor: OraclizeITransactor{contract: contract}, OraclizeIFilterer: OraclizeIFilterer{contract: contract}}, nil
}

// NewOraclizeICaller creates a new read-only instance of OraclizeI, bound to a specific deployed contract.
func NewOraclizeICaller(address common.Address, caller bind.ContractCaller) (*OraclizeICaller, error) {
	contract, err := bindOraclizeI(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &OraclizeICaller{contract: contract}, nil
}

// NewOraclizeITransactor creates a new write-only instance of OraclizeI, bound to a specific deployed contract.
func NewOraclizeITransactor(address common.Address, transactor bind.ContractTransactor) (*OraclizeITransactor, error) {
	contract, err := bindOraclizeI(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &OraclizeITransactor{contract: contract}, nil
}

// NewOraclizeIFilterer creates a new log filterer instance of OraclizeI, bound to a specific deployed contract.
func NewOraclizeIFilterer(address common.Address, filterer bind.ContractFilterer) (*OraclizeIFilterer, error) {
	contract, err := bindOraclizeI(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &OraclizeIFilterer{contract: contract}, nil
}

// bindOraclizeI binds a generic wrapper to an already deployed contract.
func bindOraclizeI(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(OraclizeIABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_OraclizeI *OraclizeIRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _OraclizeI.Contract.OraclizeICaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_OraclizeI *OraclizeIRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _OraclizeI.Contract.OraclizeITransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_OraclizeI *OraclizeIRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _OraclizeI.Contract.OraclizeITransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_OraclizeI *OraclizeICallerRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _OraclizeI.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_OraclizeI *OraclizeITransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _OraclizeI.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_OraclizeI *OraclizeITransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _OraclizeI.Contract.contract.Transact(opts, method, params...)
}

// CbAddress is a free data retrieval call binding the contract method 0xc281d19e.
//
// Solidity: function cbAddress() constant returns(address)
func (_OraclizeI *OraclizeICaller) CbAddress(opts *bind.CallOpts) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _OraclizeI.contract.Call(opts, out, "cbAddress")
	return *ret0, err
}

// CbAddress is a free data retrieval call binding the contract method 0xc281d19e.
//
// Solidity: function cbAddress() constant returns(address)
func (_OraclizeI *OraclizeISession) CbAddress() (common.Address, error) {
	return _OraclizeI.Contract.CbAddress(&_OraclizeI.CallOpts)
}

// CbAddress is a free data retrieval call binding the contract method 0xc281d19e.
//
// Solidity: function cbAddress() constant returns(address)
func (_OraclizeI *OraclizeICallerSession) CbAddress() (common.Address, error) {
	return _OraclizeI.Contract.CbAddress(&_OraclizeI.CallOpts)
}

// RandomDSGetSessionPubKeyHash is a free data retrieval call binding the contract method 0xabaa5f3e.
//
// Solidity: function randomDS_getSessionPubKeyHash() constant returns(bytes32 _sessionKeyHash)
func (_OraclizeI *OraclizeICaller) RandomDSGetSessionPubKeyHash(opts *bind.CallOpts) ([32]byte, error) {
	var (
		ret0 = new([32]byte)
	)
	out := ret0
	err := _OraclizeI.contract.Call(opts, out, "randomDS_getSessionPubKeyHash")
	return *ret0, err
}

// RandomDSGetSessionPubKeyHash is a free data retrieval call binding the contract method 0xabaa5f3e.
//
// Solidity: function randomDS_getSessionPubKeyHash() constant returns(bytes32 _sessionKeyHash)
func (_OraclizeI *OraclizeISession) RandomDSGetSessionPubKeyHash() ([32]byte, error) {
	return _OraclizeI.Contract.RandomDSGetSessionPubKeyHash(&_OraclizeI.CallOpts)
}

// RandomDSGetSessionPubKeyHash is a free data retrieval call binding the contract method 0xabaa5f3e.
//
// Solidity: function randomDS_getSessionPubKeyHash() constant returns(bytes32 _sessionKeyHash)
func (_OraclizeI *OraclizeICallerSession) RandomDSGetSessionPubKeyHash() ([32]byte, error) {
	return _OraclizeI.Contract.RandomDSGetSessionPubKeyHash(&_OraclizeI.CallOpts)
}

// GetPrice is a paid mutator transaction binding the contract method 0x524f3889.
//
// Solidity: function getPrice(string _datasource) returns(uint256 _dsprice)
func (_OraclizeI *OraclizeITransactor) GetPrice(opts *bind.TransactOpts, _datasource string) (*types.Transaction, error) {
	return _OraclizeI.contract.Transact(opts, "getPrice", _datasource)
}

// GetPrice is a paid mutator transaction binding the contract method 0x524f3889.
//
// Solidity: function getPrice(string _datasource) returns(uint256 _dsprice)
func (_OraclizeI *OraclizeISession) GetPrice(_datasource string) (*types.Transaction, error) {
	return _OraclizeI.Contract.GetPrice(&_OraclizeI.TransactOpts, _datasource)
}

// GetPrice is a paid mutator transaction binding the contract method 0x524f3889.
//
// Solidity: function getPrice(string _datasource) returns(uint256 _dsprice)
func (_OraclizeI *OraclizeITransactorSession) GetPrice(_datasource string) (*types.Transaction, error) {
	return _OraclizeI.Contract.GetPrice(&_OraclizeI.TransactOpts, _datasource)
}

// Query is a paid mutator transaction binding the contract method 0xadf59f99.
//
// Solidity: function query(uint256 _timestamp, string _datasource, string _arg) returns(bytes32 _id)
func (_OraclizeI *OraclizeITransactor) Query(opts *bind.TransactOpts, _timestamp *big.Int, _datasource string, _arg string) (*types.Transaction, error) {
	return _OraclizeI.contract.Transact(opts, "query", _timestamp, _datasource, _arg)
}

// Query is a paid mutator transaction binding the contract method 0xadf59f99.
//
// Solidity: function query(uint256 _timestamp, string _datasource, string _arg) returns(bytes32 _id)
func (_OraclizeI *OraclizeISession) Query(_timestamp *big.Int, _datasource string, _arg string) (*types.Transaction, error) {
	return _OraclizeI.Contract.Query(&_OraclizeI.TransactOpts, _timestamp, _datasource, _arg)
}

// Query is a paid mutator transaction binding the contract method 0xadf59f99.
//
// Solidity: function query(uint256 _timestamp, string _datasource, string _arg) returns(bytes32 _id)
func (_OraclizeI *OraclizeITransactorSession) Query(_timestamp *big.Int, _datasource string, _arg string) (*types.Transaction, error) {
	return _OraclizeI.Contract.Query(&_OraclizeI.TransactOpts, _timestamp, _datasource, _arg)
}

// Query2 is a paid mutator transaction binding the contract method 0x77228659.
//
// Solidity: function query2(uint256 _timestamp, string _datasource, string _arg1, string _arg2) returns(bytes32 _id)
func (_OraclizeI *OraclizeITransactor) Query2(opts *bind.TransactOpts, _timestamp *big.Int, _datasource string, _arg1 string, _arg2 string) (*types.Transaction, error) {
	return _OraclizeI.contract.Transact(opts, "query2", _timestamp, _datasource, _arg1, _arg2)
}

// Query2 is a paid mutator transaction binding the contract method 0x77228659.
//
// Solidity: function query2(uint256 _timestamp, string _datasource, string _arg1, string _arg2) returns(bytes32 _id)
func (_OraclizeI *OraclizeISession) Query2(_timestamp *big.Int, _datasource string, _arg1 string, _arg2 string) (*types.Transaction, error) {
	return _OraclizeI.Contract.Query2(&_OraclizeI.TransactOpts, _timestamp, _datasource, _arg1, _arg2)
}

// Query2 is a paid mutator transaction binding the contract method 0x77228659.
//
// Solidity: function query2(uint256 _timestamp, string _datasource, string _arg1, string _arg2) returns(bytes32 _id)
func (_OraclizeI *OraclizeITransactorSession) Query2(_timestamp *big.Int, _datasource string, _arg1 string, _arg2 string) (*types.Transaction, error) {
	return _OraclizeI.Contract.Query2(&_OraclizeI.TransactOpts, _timestamp, _datasource, _arg1, _arg2)
}

// Query2WithGasLimit is a paid mutator transaction binding the contract method 0x85dee34c.
//
// Solidity: function query2_withGasLimit(uint256 _timestamp, string _datasource, string _arg1, string _arg2, uint256 _gasLimit) returns(bytes32 _id)
func (_OraclizeI *OraclizeITransactor) Query2WithGasLimit(opts *bind.TransactOpts, _timestamp *big.Int, _datasource string, _arg1 string, _arg2 string, _gasLimit *big.Int) (*types.Transaction, error) {
	return _OraclizeI.contract.Transact(opts, "query2_withGasLimit", _timestamp, _datasource, _arg1, _arg2, _gasLimit)
}

// Query2WithGasLimit is a paid mutator transaction binding the contract method 0x85dee34c.
//
// Solidity: function query2_withGasLimit(uint256 _timestamp, string _datasource, string _arg1, string _arg2, uint256 _gasLimit) returns(bytes32 _id)
func (_OraclizeI *OraclizeISession) Query2WithGasLimit(_timestamp *big.Int, _datasource string, _arg1 string, _arg2 string, _gasLimit *big.Int) (*types.Transaction, error) {
	return _OraclizeI.Contract.Query2WithGasLimit(&_OraclizeI.TransactOpts, _timestamp, _datasource, _arg1, _arg2, _gasLimit)
}

// Query2WithGasLimit is a paid mutator transaction binding the contract method 0x85dee34c.
//
// Solidity: function query2_withGasLimit(uint256 _timestamp, string _datasource, string _arg1, string _arg2, uint256 _gasLimit) returns(bytes32 _id)
func (_OraclizeI *OraclizeITransactorSession) Query2WithGasLimit(_timestamp *big.Int, _datasource string, _arg1 string, _arg2 string, _gasLimit *big.Int) (*types.Transaction, error) {
	return _OraclizeI.Contract.Query2WithGasLimit(&_OraclizeI.TransactOpts, _timestamp, _datasource, _arg1, _arg2, _gasLimit)
}

// QueryN is a paid mutator transaction binding the contract method 0x83eed3d5.
//
// Solidity: function queryN(uint256 _timestamp, string _datasource, bytes _argN) returns(bytes32 _id)
func (_OraclizeI *OraclizeITransactor) QueryN(opts *bind.TransactOpts, _timestamp *big.Int, _datasource string, _argN []byte) (*types.Transaction, error) {
	return _OraclizeI.contract.Transact(opts, "queryN", _timestamp, _datasource, _argN)
}

// QueryN is a paid mutator transaction binding the contract method 0x83eed3d5.
//
// Solidity: function queryN(uint256 _timestamp, string _datasource, bytes _argN) returns(bytes32 _id)
func (_OraclizeI *OraclizeISession) QueryN(_timestamp *big.Int, _datasource string, _argN []byte) (*types.Transaction, error) {
	return _OraclizeI.Contract.QueryN(&_OraclizeI.TransactOpts, _timestamp, _datasource, _argN)
}

// QueryN is a paid mutator transaction binding the contract method 0x83eed3d5.
//
// Solidity: function queryN(uint256 _timestamp, string _datasource, bytes _argN) returns(bytes32 _id)
func (_OraclizeI *OraclizeITransactorSession) QueryN(_timestamp *big.Int, _datasource string, _argN []byte) (*types.Transaction, error) {
	return _OraclizeI.Contract.QueryN(&_OraclizeI.TransactOpts, _timestamp, _datasource, _argN)
}

// QueryNWithGasLimit is a paid mutator transaction binding the contract method 0xc55c1cb6.
//
// Solidity: function queryN_withGasLimit(uint256 _timestamp, string _datasource, bytes _argN, uint256 _gasLimit) returns(bytes32 _id)
func (_OraclizeI *OraclizeITransactor) QueryNWithGasLimit(opts *bind.TransactOpts, _timestamp *big.Int, _datasource string, _argN []byte, _gasLimit *big.Int) (*types.Transaction, error) {
	return _OraclizeI.contract.Transact(opts, "queryN_withGasLimit", _timestamp, _datasource, _argN, _gasLimit)
}

// QueryNWithGasLimit is a paid mutator transaction binding the contract method 0xc55c1cb6.
//
// Solidity: function queryN_withGasLimit(uint256 _timestamp, string _datasource, bytes _argN, uint256 _gasLimit) returns(bytes32 _id)
func (_OraclizeI *OraclizeISession) QueryNWithGasLimit(_timestamp *big.Int, _datasource string, _argN []byte, _gasLimit *big.Int) (*types.Transaction, error) {
	return _OraclizeI.Contract.QueryNWithGasLimit(&_OraclizeI.TransactOpts, _timestamp, _datasource, _argN, _gasLimit)
}

// QueryNWithGasLimit is a paid mutator transaction binding the contract method 0xc55c1cb6.
//
// Solidity: function queryN_withGasLimit(uint256 _timestamp, string _datasource, bytes _argN, uint256 _gasLimit) returns(bytes32 _id)
func (_OraclizeI *OraclizeITransactorSession) QueryNWithGasLimit(_timestamp *big.Int, _datasource string, _argN []byte, _gasLimit *big.Int) (*types.Transaction, error) {
	return _OraclizeI.Contract.QueryNWithGasLimit(&_OraclizeI.TransactOpts, _timestamp, _datasource, _argN, _gasLimit)
}

// QueryWithGasLimit is a paid mutator transaction binding the contract method 0xc51be90f.
//
// Solidity: function query_withGasLimit(uint256 _timestamp, string _datasource, string _arg, uint256 _gasLimit) returns(bytes32 _id)
func (_OraclizeI *OraclizeITransactor) QueryWithGasLimit(opts *bind.TransactOpts, _timestamp *big.Int, _datasource string, _arg string, _gasLimit *big.Int) (*types.Transaction, error) {
	return _OraclizeI.contract.Transact(opts, "query_withGasLimit", _timestamp, _datasource, _arg, _gasLimit)
}

// QueryWithGasLimit is a paid mutator transaction binding the contract method 0xc51be90f.
//
// Solidity: function query_withGasLimit(uint256 _timestamp, string _datasource, string _arg, uint256 _gasLimit) returns(bytes32 _id)
func (_OraclizeI *OraclizeISession) QueryWithGasLimit(_timestamp *big.Int, _datasource string, _arg string, _gasLimit *big.Int) (*types.Transaction, error) {
	return _OraclizeI.Contract.QueryWithGasLimit(&_OraclizeI.TransactOpts, _timestamp, _datasource, _arg, _gasLimit)
}

// QueryWithGasLimit is a paid mutator transaction binding the contract method 0xc51be90f.
//
// Solidity: function query_withGasLimit(uint256 _timestamp, string _datasource, string _arg, uint256 _gasLimit) returns(bytes32 _id)
func (_OraclizeI *OraclizeITransactorSession) QueryWithGasLimit(_timestamp *big.Int, _datasource string, _arg string, _gasLimit *big.Int) (*types.Transaction, error) {
	return _OraclizeI.Contract.QueryWithGasLimit(&_OraclizeI.TransactOpts, _timestamp, _datasource, _arg, _gasLimit)
}

// SetCustomGasPrice is a paid mutator transaction binding the contract method 0xca6ad1e4.
//
// Solidity: function setCustomGasPrice(uint256 _gasPrice) returns()
func (_OraclizeI *OraclizeITransactor) SetCustomGasPrice(opts *bind.TransactOpts, _gasPrice *big.Int) (*types.Transaction, error) {
	return _OraclizeI.contract.Transact(opts, "setCustomGasPrice", _gasPrice)
}

// SetCustomGasPrice is a paid mutator transaction binding the contract method 0xca6ad1e4.
//
// Solidity: function setCustomGasPrice(uint256 _gasPrice) returns()
func (_OraclizeI *OraclizeISession) SetCustomGasPrice(_gasPrice *big.Int) (*types.Transaction, error) {
	return _OraclizeI.Contract.SetCustomGasPrice(&_OraclizeI.TransactOpts, _gasPrice)
}

// SetCustomGasPrice is a paid mutator transaction binding the contract method 0xca6ad1e4.
//
// Solidity: function setCustomGasPrice(uint256 _gasPrice) returns()
func (_OraclizeI *OraclizeITransactorSession) SetCustomGasPrice(_gasPrice *big.Int) (*types.Transaction, error) {
	return _OraclizeI.Contract.SetCustomGasPrice(&_OraclizeI.TransactOpts, _gasPrice)
}

// SetProofType is a paid mutator transaction binding the contract method 0x688dcfd7.
//
// Solidity: function setProofType(bytes1 _proofType) returns()
func (_OraclizeI *OraclizeITransactor) SetProofType(opts *bind.TransactOpts, _proofType [1]byte) (*types.Transaction, error) {
	return _OraclizeI.contract.Transact(opts, "setProofType", _proofType)
}

// SetProofType is a paid mutator transaction binding the contract method 0x688dcfd7.
//
// Solidity: function setProofType(bytes1 _proofType) returns()
func (_OraclizeI *OraclizeISession) SetProofType(_proofType [1]byte) (*types.Transaction, error) {
	return _OraclizeI.Contract.SetProofType(&_OraclizeI.TransactOpts, _proofType)
}

// SetProofType is a paid mutator transaction binding the contract method 0x688dcfd7.
//
// Solidity: function setProofType(bytes1 _proofType) returns()
func (_OraclizeI *OraclizeITransactorSession) SetProofType(_proofType [1]byte) (*types.Transaction, error) {
	return _OraclizeI.Contract.SetProofType(&_OraclizeI.TransactOpts, _proofType)
}

// OwnableABI is the input ABI used to generate the binding from.
const OwnableABI = "[{\"constant\":false,\"inputs\":[],\"name\":\"renounceOwnership\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"isOwner\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"transferOwnership\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"previousOwner\",\"type\":\"address\"},{\"indexed\":true,\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"OwnershipTransferred\",\"type\":\"event\"}]"

// OwnableBin is the compiled bytecode used for deploying new contracts.
const OwnableBin = `0x`

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

// IsOwner is a free data retrieval call binding the contract method 0x8f32d59b.
//
// Solidity: function isOwner() constant returns(bool)
func (_Ownable *OwnableCaller) IsOwner(opts *bind.CallOpts) (bool, error) {
	var (
		ret0 = new(bool)
	)
	out := ret0
	err := _Ownable.contract.Call(opts, out, "isOwner")
	return *ret0, err
}

// IsOwner is a free data retrieval call binding the contract method 0x8f32d59b.
//
// Solidity: function isOwner() constant returns(bool)
func (_Ownable *OwnableSession) IsOwner() (bool, error) {
	return _Ownable.Contract.IsOwner(&_Ownable.CallOpts)
}

// IsOwner is a free data retrieval call binding the contract method 0x8f32d59b.
//
// Solidity: function isOwner() constant returns(bool)
func (_Ownable *OwnableCallerSession) IsOwner() (bool, error) {
	return _Ownable.Contract.IsOwner(&_Ownable.CallOpts)
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
// Solidity: function transferOwnership(address newOwner) returns()
func (_Ownable *OwnableTransactor) TransferOwnership(opts *bind.TransactOpts, newOwner common.Address) (*types.Transaction, error) {
	return _Ownable.contract.Transact(opts, "transferOwnership", newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_Ownable *OwnableSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _Ownable.Contract.TransferOwnership(&_Ownable.TransactOpts, newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_Ownable *OwnableTransactorSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _Ownable.Contract.TransferOwnership(&_Ownable.TransactOpts, newOwner)
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
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
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
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
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

// PausableABI is the input ABI used to generate the binding from.
const PausableABI = "[{\"constant\":false,\"inputs\":[],\"name\":\"unpause\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"account\",\"type\":\"address\"}],\"name\":\"isPauser\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"paused\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[],\"name\":\"renouncePauser\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"account\",\"type\":\"address\"}],\"name\":\"addPauser\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[],\"name\":\"pause\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"name\":\"account\",\"type\":\"address\"}],\"name\":\"Paused\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"name\":\"account\",\"type\":\"address\"}],\"name\":\"Unpaused\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"account\",\"type\":\"address\"}],\"name\":\"PauserAdded\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"account\",\"type\":\"address\"}],\"name\":\"PauserRemoved\",\"type\":\"event\"}]"

// PausableBin is the compiled bytecode used for deploying new contracts.
const PausableBin = `0x`

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
	return address, tx, &Pausable{PausableCaller: PausableCaller{contract: contract}, PausableTransactor: PausableTransactor{contract: contract}, PausableFilterer: PausableFilterer{contract: contract}}, nil
}

// Pausable is an auto generated Go binding around an Ethereum contract.
type Pausable struct {
	PausableCaller     // Read-only binding to the contract
	PausableTransactor // Write-only binding to the contract
	PausableFilterer   // Log filterer for contract events
}

// PausableCaller is an auto generated read-only Go binding around an Ethereum contract.
type PausableCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// PausableTransactor is an auto generated write-only Go binding around an Ethereum contract.
type PausableTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// PausableFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type PausableFilterer struct {
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
	contract, err := bindPausable(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Pausable{PausableCaller: PausableCaller{contract: contract}, PausableTransactor: PausableTransactor{contract: contract}, PausableFilterer: PausableFilterer{contract: contract}}, nil
}

// NewPausableCaller creates a new read-only instance of Pausable, bound to a specific deployed contract.
func NewPausableCaller(address common.Address, caller bind.ContractCaller) (*PausableCaller, error) {
	contract, err := bindPausable(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &PausableCaller{contract: contract}, nil
}

// NewPausableTransactor creates a new write-only instance of Pausable, bound to a specific deployed contract.
func NewPausableTransactor(address common.Address, transactor bind.ContractTransactor) (*PausableTransactor, error) {
	contract, err := bindPausable(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &PausableTransactor{contract: contract}, nil
}

// NewPausableFilterer creates a new log filterer instance of Pausable, bound to a specific deployed contract.
func NewPausableFilterer(address common.Address, filterer bind.ContractFilterer) (*PausableFilterer, error) {
	contract, err := bindPausable(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &PausableFilterer{contract: contract}, nil
}

// bindPausable binds a generic wrapper to an already deployed contract.
func bindPausable(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(PausableABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
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

// IsPauser is a free data retrieval call binding the contract method 0x46fbf68e.
//
// Solidity: function isPauser(address account) constant returns(bool)
func (_Pausable *PausableCaller) IsPauser(opts *bind.CallOpts, account common.Address) (bool, error) {
	var (
		ret0 = new(bool)
	)
	out := ret0
	err := _Pausable.contract.Call(opts, out, "isPauser", account)
	return *ret0, err
}

// IsPauser is a free data retrieval call binding the contract method 0x46fbf68e.
//
// Solidity: function isPauser(address account) constant returns(bool)
func (_Pausable *PausableSession) IsPauser(account common.Address) (bool, error) {
	return _Pausable.Contract.IsPauser(&_Pausable.CallOpts, account)
}

// IsPauser is a free data retrieval call binding the contract method 0x46fbf68e.
//
// Solidity: function isPauser(address account) constant returns(bool)
func (_Pausable *PausableCallerSession) IsPauser(account common.Address) (bool, error) {
	return _Pausable.Contract.IsPauser(&_Pausable.CallOpts, account)
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

// AddPauser is a paid mutator transaction binding the contract method 0x82dc1ec4.
//
// Solidity: function addPauser(address account) returns()
func (_Pausable *PausableTransactor) AddPauser(opts *bind.TransactOpts, account common.Address) (*types.Transaction, error) {
	return _Pausable.contract.Transact(opts, "addPauser", account)
}

// AddPauser is a paid mutator transaction binding the contract method 0x82dc1ec4.
//
// Solidity: function addPauser(address account) returns()
func (_Pausable *PausableSession) AddPauser(account common.Address) (*types.Transaction, error) {
	return _Pausable.Contract.AddPauser(&_Pausable.TransactOpts, account)
}

// AddPauser is a paid mutator transaction binding the contract method 0x82dc1ec4.
//
// Solidity: function addPauser(address account) returns()
func (_Pausable *PausableTransactorSession) AddPauser(account common.Address) (*types.Transaction, error) {
	return _Pausable.Contract.AddPauser(&_Pausable.TransactOpts, account)
}

// Pause is a paid mutator transaction binding the contract method 0x8456cb59.
//
// Solidity: function pause() returns()
func (_Pausable *PausableTransactor) Pause(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Pausable.contract.Transact(opts, "pause")
}

// Pause is a paid mutator transaction binding the contract method 0x8456cb59.
//
// Solidity: function pause() returns()
func (_Pausable *PausableSession) Pause() (*types.Transaction, error) {
	return _Pausable.Contract.Pause(&_Pausable.TransactOpts)
}

// Pause is a paid mutator transaction binding the contract method 0x8456cb59.
//
// Solidity: function pause() returns()
func (_Pausable *PausableTransactorSession) Pause() (*types.Transaction, error) {
	return _Pausable.Contract.Pause(&_Pausable.TransactOpts)
}

// RenouncePauser is a paid mutator transaction binding the contract method 0x6ef8d66d.
//
// Solidity: function renouncePauser() returns()
func (_Pausable *PausableTransactor) RenouncePauser(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Pausable.contract.Transact(opts, "renouncePauser")
}

// RenouncePauser is a paid mutator transaction binding the contract method 0x6ef8d66d.
//
// Solidity: function renouncePauser() returns()
func (_Pausable *PausableSession) RenouncePauser() (*types.Transaction, error) {
	return _Pausable.Contract.RenouncePauser(&_Pausable.TransactOpts)
}

// RenouncePauser is a paid mutator transaction binding the contract method 0x6ef8d66d.
//
// Solidity: function renouncePauser() returns()
func (_Pausable *PausableTransactorSession) RenouncePauser() (*types.Transaction, error) {
	return _Pausable.Contract.RenouncePauser(&_Pausable.TransactOpts)
}

// Unpause is a paid mutator transaction binding the contract method 0x3f4ba83a.
//
// Solidity: function unpause() returns()
func (_Pausable *PausableTransactor) Unpause(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Pausable.contract.Transact(opts, "unpause")
}

// Unpause is a paid mutator transaction binding the contract method 0x3f4ba83a.
//
// Solidity: function unpause() returns()
func (_Pausable *PausableSession) Unpause() (*types.Transaction, error) {
	return _Pausable.Contract.Unpause(&_Pausable.TransactOpts)
}

// Unpause is a paid mutator transaction binding the contract method 0x3f4ba83a.
//
// Solidity: function unpause() returns()
func (_Pausable *PausableTransactorSession) Unpause() (*types.Transaction, error) {
	return _Pausable.Contract.Unpause(&_Pausable.TransactOpts)
}

// PausablePausedIterator is returned from FilterPaused and is used to iterate over the raw logs and unpacked data for Paused events raised by the Pausable contract.
type PausablePausedIterator struct {
	Event *PausablePaused // Event containing the contract specifics and raw log

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
func (it *PausablePausedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(PausablePaused)
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
		it.Event = new(PausablePaused)
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
func (it *PausablePausedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *PausablePausedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// PausablePaused represents a Paused event raised by the Pausable contract.
type PausablePaused struct {
	Account common.Address
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterPaused is a free log retrieval operation binding the contract event 0x62e78cea01bee320cd4e420270b5ea74000d11b0c9f74754ebdbfc544b05a258.
//
// Solidity: event Paused(address account)
func (_Pausable *PausableFilterer) FilterPaused(opts *bind.FilterOpts) (*PausablePausedIterator, error) {

	logs, sub, err := _Pausable.contract.FilterLogs(opts, "Paused")
	if err != nil {
		return nil, err
	}
	return &PausablePausedIterator{contract: _Pausable.contract, event: "Paused", logs: logs, sub: sub}, nil
}

// WatchPaused is a free log subscription operation binding the contract event 0x62e78cea01bee320cd4e420270b5ea74000d11b0c9f74754ebdbfc544b05a258.
//
// Solidity: event Paused(address account)
func (_Pausable *PausableFilterer) WatchPaused(opts *bind.WatchOpts, sink chan<- *PausablePaused) (event.Subscription, error) {

	logs, sub, err := _Pausable.contract.WatchLogs(opts, "Paused")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(PausablePaused)
				if err := _Pausable.contract.UnpackLog(event, "Paused", log); err != nil {
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

// PausablePauserAddedIterator is returned from FilterPauserAdded and is used to iterate over the raw logs and unpacked data for PauserAdded events raised by the Pausable contract.
type PausablePauserAddedIterator struct {
	Event *PausablePauserAdded // Event containing the contract specifics and raw log

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
func (it *PausablePauserAddedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(PausablePauserAdded)
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
		it.Event = new(PausablePauserAdded)
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
func (it *PausablePauserAddedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *PausablePauserAddedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// PausablePauserAdded represents a PauserAdded event raised by the Pausable contract.
type PausablePauserAdded struct {
	Account common.Address
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterPauserAdded is a free log retrieval operation binding the contract event 0x6719d08c1888103bea251a4ed56406bd0c3e69723c8a1686e017e7bbe159b6f8.
//
// Solidity: event PauserAdded(address indexed account)
func (_Pausable *PausableFilterer) FilterPauserAdded(opts *bind.FilterOpts, account []common.Address) (*PausablePauserAddedIterator, error) {

	var accountRule []interface{}
	for _, accountItem := range account {
		accountRule = append(accountRule, accountItem)
	}

	logs, sub, err := _Pausable.contract.FilterLogs(opts, "PauserAdded", accountRule)
	if err != nil {
		return nil, err
	}
	return &PausablePauserAddedIterator{contract: _Pausable.contract, event: "PauserAdded", logs: logs, sub: sub}, nil
}

// WatchPauserAdded is a free log subscription operation binding the contract event 0x6719d08c1888103bea251a4ed56406bd0c3e69723c8a1686e017e7bbe159b6f8.
//
// Solidity: event PauserAdded(address indexed account)
func (_Pausable *PausableFilterer) WatchPauserAdded(opts *bind.WatchOpts, sink chan<- *PausablePauserAdded, account []common.Address) (event.Subscription, error) {

	var accountRule []interface{}
	for _, accountItem := range account {
		accountRule = append(accountRule, accountItem)
	}

	logs, sub, err := _Pausable.contract.WatchLogs(opts, "PauserAdded", accountRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(PausablePauserAdded)
				if err := _Pausable.contract.UnpackLog(event, "PauserAdded", log); err != nil {
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

// PausablePauserRemovedIterator is returned from FilterPauserRemoved and is used to iterate over the raw logs and unpacked data for PauserRemoved events raised by the Pausable contract.
type PausablePauserRemovedIterator struct {
	Event *PausablePauserRemoved // Event containing the contract specifics and raw log

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
func (it *PausablePauserRemovedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(PausablePauserRemoved)
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
		it.Event = new(PausablePauserRemoved)
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
func (it *PausablePauserRemovedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *PausablePauserRemovedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// PausablePauserRemoved represents a PauserRemoved event raised by the Pausable contract.
type PausablePauserRemoved struct {
	Account common.Address
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterPauserRemoved is a free log retrieval operation binding the contract event 0xcd265ebaf09df2871cc7bd4133404a235ba12eff2041bb89d9c714a2621c7c7e.
//
// Solidity: event PauserRemoved(address indexed account)
func (_Pausable *PausableFilterer) FilterPauserRemoved(opts *bind.FilterOpts, account []common.Address) (*PausablePauserRemovedIterator, error) {

	var accountRule []interface{}
	for _, accountItem := range account {
		accountRule = append(accountRule, accountItem)
	}

	logs, sub, err := _Pausable.contract.FilterLogs(opts, "PauserRemoved", accountRule)
	if err != nil {
		return nil, err
	}
	return &PausablePauserRemovedIterator{contract: _Pausable.contract, event: "PauserRemoved", logs: logs, sub: sub}, nil
}

// WatchPauserRemoved is a free log subscription operation binding the contract event 0xcd265ebaf09df2871cc7bd4133404a235ba12eff2041bb89d9c714a2621c7c7e.
//
// Solidity: event PauserRemoved(address indexed account)
func (_Pausable *PausableFilterer) WatchPauserRemoved(opts *bind.WatchOpts, sink chan<- *PausablePauserRemoved, account []common.Address) (event.Subscription, error) {

	var accountRule []interface{}
	for _, accountItem := range account {
		accountRule = append(accountRule, accountItem)
	}

	logs, sub, err := _Pausable.contract.WatchLogs(opts, "PauserRemoved", accountRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(PausablePauserRemoved)
				if err := _Pausable.contract.UnpackLog(event, "PauserRemoved", log); err != nil {
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

// PausableUnpausedIterator is returned from FilterUnpaused and is used to iterate over the raw logs and unpacked data for Unpaused events raised by the Pausable contract.
type PausableUnpausedIterator struct {
	Event *PausableUnpaused // Event containing the contract specifics and raw log

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
func (it *PausableUnpausedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(PausableUnpaused)
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
		it.Event = new(PausableUnpaused)
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
func (it *PausableUnpausedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *PausableUnpausedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// PausableUnpaused represents a Unpaused event raised by the Pausable contract.
type PausableUnpaused struct {
	Account common.Address
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterUnpaused is a free log retrieval operation binding the contract event 0x5db9ee0a495bf2e6ff9c91a7834c1ba4fdd244a5e8aa4e537bd38aeae4b073aa.
//
// Solidity: event Unpaused(address account)
func (_Pausable *PausableFilterer) FilterUnpaused(opts *bind.FilterOpts) (*PausableUnpausedIterator, error) {

	logs, sub, err := _Pausable.contract.FilterLogs(opts, "Unpaused")
	if err != nil {
		return nil, err
	}
	return &PausableUnpausedIterator{contract: _Pausable.contract, event: "Unpaused", logs: logs, sub: sub}, nil
}

// WatchUnpaused is a free log subscription operation binding the contract event 0x5db9ee0a495bf2e6ff9c91a7834c1ba4fdd244a5e8aa4e537bd38aeae4b073aa.
//
// Solidity: event Unpaused(address account)
func (_Pausable *PausableFilterer) WatchUnpaused(opts *bind.WatchOpts, sink chan<- *PausableUnpaused) (event.Subscription, error) {

	logs, sub, err := _Pausable.contract.WatchLogs(opts, "Unpaused")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(PausableUnpaused)
				if err := _Pausable.contract.UnpackLog(event, "Unpaused", log); err != nil {
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

// PauserRoleABI is the input ABI used to generate the binding from.
const PauserRoleABI = "[{\"constant\":true,\"inputs\":[{\"name\":\"account\",\"type\":\"address\"}],\"name\":\"isPauser\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[],\"name\":\"renouncePauser\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"account\",\"type\":\"address\"}],\"name\":\"addPauser\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"account\",\"type\":\"address\"}],\"name\":\"PauserAdded\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"account\",\"type\":\"address\"}],\"name\":\"PauserRemoved\",\"type\":\"event\"}]"

// PauserRoleBin is the compiled bytecode used for deploying new contracts.
const PauserRoleBin = `0x`

// DeployPauserRole deploys a new Ethereum contract, binding an instance of PauserRole to it.
func DeployPauserRole(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *PauserRole, error) {
	parsed, err := abi.JSON(strings.NewReader(PauserRoleABI))
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	address, tx, contract, err := bind.DeployContract(auth, parsed, common.FromHex(PauserRoleBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &PauserRole{PauserRoleCaller: PauserRoleCaller{contract: contract}, PauserRoleTransactor: PauserRoleTransactor{contract: contract}, PauserRoleFilterer: PauserRoleFilterer{contract: contract}}, nil
}

// PauserRole is an auto generated Go binding around an Ethereum contract.
type PauserRole struct {
	PauserRoleCaller     // Read-only binding to the contract
	PauserRoleTransactor // Write-only binding to the contract
	PauserRoleFilterer   // Log filterer for contract events
}

// PauserRoleCaller is an auto generated read-only Go binding around an Ethereum contract.
type PauserRoleCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// PauserRoleTransactor is an auto generated write-only Go binding around an Ethereum contract.
type PauserRoleTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// PauserRoleFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type PauserRoleFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// PauserRoleSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type PauserRoleSession struct {
	Contract     *PauserRole       // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// PauserRoleCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type PauserRoleCallerSession struct {
	Contract *PauserRoleCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts     // Call options to use throughout this session
}

// PauserRoleTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type PauserRoleTransactorSession struct {
	Contract     *PauserRoleTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts     // Transaction auth options to use throughout this session
}

// PauserRoleRaw is an auto generated low-level Go binding around an Ethereum contract.
type PauserRoleRaw struct {
	Contract *PauserRole // Generic contract binding to access the raw methods on
}

// PauserRoleCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type PauserRoleCallerRaw struct {
	Contract *PauserRoleCaller // Generic read-only contract binding to access the raw methods on
}

// PauserRoleTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type PauserRoleTransactorRaw struct {
	Contract *PauserRoleTransactor // Generic write-only contract binding to access the raw methods on
}

// NewPauserRole creates a new instance of PauserRole, bound to a specific deployed contract.
func NewPauserRole(address common.Address, backend bind.ContractBackend) (*PauserRole, error) {
	contract, err := bindPauserRole(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &PauserRole{PauserRoleCaller: PauserRoleCaller{contract: contract}, PauserRoleTransactor: PauserRoleTransactor{contract: contract}, PauserRoleFilterer: PauserRoleFilterer{contract: contract}}, nil
}

// NewPauserRoleCaller creates a new read-only instance of PauserRole, bound to a specific deployed contract.
func NewPauserRoleCaller(address common.Address, caller bind.ContractCaller) (*PauserRoleCaller, error) {
	contract, err := bindPauserRole(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &PauserRoleCaller{contract: contract}, nil
}

// NewPauserRoleTransactor creates a new write-only instance of PauserRole, bound to a specific deployed contract.
func NewPauserRoleTransactor(address common.Address, transactor bind.ContractTransactor) (*PauserRoleTransactor, error) {
	contract, err := bindPauserRole(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &PauserRoleTransactor{contract: contract}, nil
}

// NewPauserRoleFilterer creates a new log filterer instance of PauserRole, bound to a specific deployed contract.
func NewPauserRoleFilterer(address common.Address, filterer bind.ContractFilterer) (*PauserRoleFilterer, error) {
	contract, err := bindPauserRole(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &PauserRoleFilterer{contract: contract}, nil
}

// bindPauserRole binds a generic wrapper to an already deployed contract.
func bindPauserRole(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(PauserRoleABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_PauserRole *PauserRoleRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _PauserRole.Contract.PauserRoleCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_PauserRole *PauserRoleRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _PauserRole.Contract.PauserRoleTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_PauserRole *PauserRoleRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _PauserRole.Contract.PauserRoleTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_PauserRole *PauserRoleCallerRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _PauserRole.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_PauserRole *PauserRoleTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _PauserRole.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_PauserRole *PauserRoleTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _PauserRole.Contract.contract.Transact(opts, method, params...)
}

// IsPauser is a free data retrieval call binding the contract method 0x46fbf68e.
//
// Solidity: function isPauser(address account) constant returns(bool)
func (_PauserRole *PauserRoleCaller) IsPauser(opts *bind.CallOpts, account common.Address) (bool, error) {
	var (
		ret0 = new(bool)
	)
	out := ret0
	err := _PauserRole.contract.Call(opts, out, "isPauser", account)
	return *ret0, err
}

// IsPauser is a free data retrieval call binding the contract method 0x46fbf68e.
//
// Solidity: function isPauser(address account) constant returns(bool)
func (_PauserRole *PauserRoleSession) IsPauser(account common.Address) (bool, error) {
	return _PauserRole.Contract.IsPauser(&_PauserRole.CallOpts, account)
}

// IsPauser is a free data retrieval call binding the contract method 0x46fbf68e.
//
// Solidity: function isPauser(address account) constant returns(bool)
func (_PauserRole *PauserRoleCallerSession) IsPauser(account common.Address) (bool, error) {
	return _PauserRole.Contract.IsPauser(&_PauserRole.CallOpts, account)
}

// AddPauser is a paid mutator transaction binding the contract method 0x82dc1ec4.
//
// Solidity: function addPauser(address account) returns()
func (_PauserRole *PauserRoleTransactor) AddPauser(opts *bind.TransactOpts, account common.Address) (*types.Transaction, error) {
	return _PauserRole.contract.Transact(opts, "addPauser", account)
}

// AddPauser is a paid mutator transaction binding the contract method 0x82dc1ec4.
//
// Solidity: function addPauser(address account) returns()
func (_PauserRole *PauserRoleSession) AddPauser(account common.Address) (*types.Transaction, error) {
	return _PauserRole.Contract.AddPauser(&_PauserRole.TransactOpts, account)
}

// AddPauser is a paid mutator transaction binding the contract method 0x82dc1ec4.
//
// Solidity: function addPauser(address account) returns()
func (_PauserRole *PauserRoleTransactorSession) AddPauser(account common.Address) (*types.Transaction, error) {
	return _PauserRole.Contract.AddPauser(&_PauserRole.TransactOpts, account)
}

// RenouncePauser is a paid mutator transaction binding the contract method 0x6ef8d66d.
//
// Solidity: function renouncePauser() returns()
func (_PauserRole *PauserRoleTransactor) RenouncePauser(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _PauserRole.contract.Transact(opts, "renouncePauser")
}

// RenouncePauser is a paid mutator transaction binding the contract method 0x6ef8d66d.
//
// Solidity: function renouncePauser() returns()
func (_PauserRole *PauserRoleSession) RenouncePauser() (*types.Transaction, error) {
	return _PauserRole.Contract.RenouncePauser(&_PauserRole.TransactOpts)
}

// RenouncePauser is a paid mutator transaction binding the contract method 0x6ef8d66d.
//
// Solidity: function renouncePauser() returns()
func (_PauserRole *PauserRoleTransactorSession) RenouncePauser() (*types.Transaction, error) {
	return _PauserRole.Contract.RenouncePauser(&_PauserRole.TransactOpts)
}

// PauserRolePauserAddedIterator is returned from FilterPauserAdded and is used to iterate over the raw logs and unpacked data for PauserAdded events raised by the PauserRole contract.
type PauserRolePauserAddedIterator struct {
	Event *PauserRolePauserAdded // Event containing the contract specifics and raw log

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
func (it *PauserRolePauserAddedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(PauserRolePauserAdded)
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
		it.Event = new(PauserRolePauserAdded)
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
func (it *PauserRolePauserAddedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *PauserRolePauserAddedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// PauserRolePauserAdded represents a PauserAdded event raised by the PauserRole contract.
type PauserRolePauserAdded struct {
	Account common.Address
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterPauserAdded is a free log retrieval operation binding the contract event 0x6719d08c1888103bea251a4ed56406bd0c3e69723c8a1686e017e7bbe159b6f8.
//
// Solidity: event PauserAdded(address indexed account)
func (_PauserRole *PauserRoleFilterer) FilterPauserAdded(opts *bind.FilterOpts, account []common.Address) (*PauserRolePauserAddedIterator, error) {

	var accountRule []interface{}
	for _, accountItem := range account {
		accountRule = append(accountRule, accountItem)
	}

	logs, sub, err := _PauserRole.contract.FilterLogs(opts, "PauserAdded", accountRule)
	if err != nil {
		return nil, err
	}
	return &PauserRolePauserAddedIterator{contract: _PauserRole.contract, event: "PauserAdded", logs: logs, sub: sub}, nil
}

// WatchPauserAdded is a free log subscription operation binding the contract event 0x6719d08c1888103bea251a4ed56406bd0c3e69723c8a1686e017e7bbe159b6f8.
//
// Solidity: event PauserAdded(address indexed account)
func (_PauserRole *PauserRoleFilterer) WatchPauserAdded(opts *bind.WatchOpts, sink chan<- *PauserRolePauserAdded, account []common.Address) (event.Subscription, error) {

	var accountRule []interface{}
	for _, accountItem := range account {
		accountRule = append(accountRule, accountItem)
	}

	logs, sub, err := _PauserRole.contract.WatchLogs(opts, "PauserAdded", accountRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(PauserRolePauserAdded)
				if err := _PauserRole.contract.UnpackLog(event, "PauserAdded", log); err != nil {
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

// PauserRolePauserRemovedIterator is returned from FilterPauserRemoved and is used to iterate over the raw logs and unpacked data for PauserRemoved events raised by the PauserRole contract.
type PauserRolePauserRemovedIterator struct {
	Event *PauserRolePauserRemoved // Event containing the contract specifics and raw log

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
func (it *PauserRolePauserRemovedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(PauserRolePauserRemoved)
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
		it.Event = new(PauserRolePauserRemoved)
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
func (it *PauserRolePauserRemovedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *PauserRolePauserRemovedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// PauserRolePauserRemoved represents a PauserRemoved event raised by the PauserRole contract.
type PauserRolePauserRemoved struct {
	Account common.Address
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterPauserRemoved is a free log retrieval operation binding the contract event 0xcd265ebaf09df2871cc7bd4133404a235ba12eff2041bb89d9c714a2621c7c7e.
//
// Solidity: event PauserRemoved(address indexed account)
func (_PauserRole *PauserRoleFilterer) FilterPauserRemoved(opts *bind.FilterOpts, account []common.Address) (*PauserRolePauserRemovedIterator, error) {

	var accountRule []interface{}
	for _, accountItem := range account {
		accountRule = append(accountRule, accountItem)
	}

	logs, sub, err := _PauserRole.contract.FilterLogs(opts, "PauserRemoved", accountRule)
	if err != nil {
		return nil, err
	}
	return &PauserRolePauserRemovedIterator{contract: _PauserRole.contract, event: "PauserRemoved", logs: logs, sub: sub}, nil
}

// WatchPauserRemoved is a free log subscription operation binding the contract event 0xcd265ebaf09df2871cc7bd4133404a235ba12eff2041bb89d9c714a2621c7c7e.
//
// Solidity: event PauserRemoved(address indexed account)
func (_PauserRole *PauserRoleFilterer) WatchPauserRemoved(opts *bind.WatchOpts, sink chan<- *PauserRolePauserRemoved, account []common.Address) (event.Subscription, error) {

	var accountRule []interface{}
	for _, accountItem := range account {
		accountRule = append(accountRule, accountItem)
	}

	logs, sub, err := _PauserRole.contract.WatchLogs(opts, "PauserRemoved", accountRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(PauserRolePauserRemoved)
				if err := _PauserRole.contract.UnpackLog(event, "PauserRemoved", log); err != nil {
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

// RolesABI is the input ABI used to generate the binding from.
const RolesABI = "[]"

// RolesBin is the compiled bytecode used for deploying new contracts.
const RolesBin = `0x604c602c600b82828239805160001a60731460008114601c57601e565bfe5b5030600052607381538281f3fe73000000000000000000000000000000000000000030146080604052600080fdfea165627a7a72305820a8f0dad597a68be48b2ecd3447c4123fba454f485c47871a2de548906bf797440029`

// DeployRoles deploys a new Ethereum contract, binding an instance of Roles to it.
func DeployRoles(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *Roles, error) {
	parsed, err := abi.JSON(strings.NewReader(RolesABI))
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	address, tx, contract, err := bind.DeployContract(auth, parsed, common.FromHex(RolesBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &Roles{RolesCaller: RolesCaller{contract: contract}, RolesTransactor: RolesTransactor{contract: contract}, RolesFilterer: RolesFilterer{contract: contract}}, nil
}

// Roles is an auto generated Go binding around an Ethereum contract.
type Roles struct {
	RolesCaller     // Read-only binding to the contract
	RolesTransactor // Write-only binding to the contract
	RolesFilterer   // Log filterer for contract events
}

// RolesCaller is an auto generated read-only Go binding around an Ethereum contract.
type RolesCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// RolesTransactor is an auto generated write-only Go binding around an Ethereum contract.
type RolesTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// RolesFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type RolesFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// RolesSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type RolesSession struct {
	Contract     *Roles            // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// RolesCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type RolesCallerSession struct {
	Contract *RolesCaller  // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts // Call options to use throughout this session
}

// RolesTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type RolesTransactorSession struct {
	Contract     *RolesTransactor  // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// RolesRaw is an auto generated low-level Go binding around an Ethereum contract.
type RolesRaw struct {
	Contract *Roles // Generic contract binding to access the raw methods on
}

// RolesCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type RolesCallerRaw struct {
	Contract *RolesCaller // Generic read-only contract binding to access the raw methods on
}

// RolesTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type RolesTransactorRaw struct {
	Contract *RolesTransactor // Generic write-only contract binding to access the raw methods on
}

// NewRoles creates a new instance of Roles, bound to a specific deployed contract.
func NewRoles(address common.Address, backend bind.ContractBackend) (*Roles, error) {
	contract, err := bindRoles(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Roles{RolesCaller: RolesCaller{contract: contract}, RolesTransactor: RolesTransactor{contract: contract}, RolesFilterer: RolesFilterer{contract: contract}}, nil
}

// NewRolesCaller creates a new read-only instance of Roles, bound to a specific deployed contract.
func NewRolesCaller(address common.Address, caller bind.ContractCaller) (*RolesCaller, error) {
	contract, err := bindRoles(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &RolesCaller{contract: contract}, nil
}

// NewRolesTransactor creates a new write-only instance of Roles, bound to a specific deployed contract.
func NewRolesTransactor(address common.Address, transactor bind.ContractTransactor) (*RolesTransactor, error) {
	contract, err := bindRoles(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &RolesTransactor{contract: contract}, nil
}

// NewRolesFilterer creates a new log filterer instance of Roles, bound to a specific deployed contract.
func NewRolesFilterer(address common.Address, filterer bind.ContractFilterer) (*RolesFilterer, error) {
	contract, err := bindRoles(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &RolesFilterer{contract: contract}, nil
}

// bindRoles binds a generic wrapper to an already deployed contract.
func bindRoles(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(RolesABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Roles *RolesRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _Roles.Contract.RolesCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Roles *RolesRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Roles.Contract.RolesTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Roles *RolesRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Roles.Contract.RolesTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Roles *RolesCallerRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _Roles.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Roles *RolesTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Roles.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Roles *RolesTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Roles.Contract.contract.Transact(opts, method, params...)
}

// SafeMathABI is the input ABI used to generate the binding from.
const SafeMathABI = "[]"

// SafeMathBin is the compiled bytecode used for deploying new contracts.
const SafeMathBin = `0x604c602c600b82828239805160001a60731460008114601c57601e565bfe5b5030600052607381538281f3fe73000000000000000000000000000000000000000030146080604052600080fdfea165627a7a72305820492a7cd55fb605fd7d3e564e8600da7ee8e18c25ef071b826e8f806b1755d82f0029`

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

// SolcCheckerABI is the input ABI used to generate the binding from.
const SolcCheckerABI = "[{\"constant\":false,\"inputs\":[{\"name\":\"x\",\"type\":\"bytes\"}],\"name\":\"f\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]"

// SolcCheckerBin is the compiled bytecode used for deploying new contracts.
const SolcCheckerBin = `0x`

// DeploySolcChecker deploys a new Ethereum contract, binding an instance of SolcChecker to it.
func DeploySolcChecker(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *SolcChecker, error) {
	parsed, err := abi.JSON(strings.NewReader(SolcCheckerABI))
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	address, tx, contract, err := bind.DeployContract(auth, parsed, common.FromHex(SolcCheckerBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &SolcChecker{SolcCheckerCaller: SolcCheckerCaller{contract: contract}, SolcCheckerTransactor: SolcCheckerTransactor{contract: contract}, SolcCheckerFilterer: SolcCheckerFilterer{contract: contract}}, nil
}

// SolcChecker is an auto generated Go binding around an Ethereum contract.
type SolcChecker struct {
	SolcCheckerCaller     // Read-only binding to the contract
	SolcCheckerTransactor // Write-only binding to the contract
	SolcCheckerFilterer   // Log filterer for contract events
}

// SolcCheckerCaller is an auto generated read-only Go binding around an Ethereum contract.
type SolcCheckerCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// SolcCheckerTransactor is an auto generated write-only Go binding around an Ethereum contract.
type SolcCheckerTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// SolcCheckerFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type SolcCheckerFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// SolcCheckerSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type SolcCheckerSession struct {
	Contract     *SolcChecker      // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// SolcCheckerCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type SolcCheckerCallerSession struct {
	Contract *SolcCheckerCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts      // Call options to use throughout this session
}

// SolcCheckerTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type SolcCheckerTransactorSession struct {
	Contract     *SolcCheckerTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts      // Transaction auth options to use throughout this session
}

// SolcCheckerRaw is an auto generated low-level Go binding around an Ethereum contract.
type SolcCheckerRaw struct {
	Contract *SolcChecker // Generic contract binding to access the raw methods on
}

// SolcCheckerCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type SolcCheckerCallerRaw struct {
	Contract *SolcCheckerCaller // Generic read-only contract binding to access the raw methods on
}

// SolcCheckerTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type SolcCheckerTransactorRaw struct {
	Contract *SolcCheckerTransactor // Generic write-only contract binding to access the raw methods on
}

// NewSolcChecker creates a new instance of SolcChecker, bound to a specific deployed contract.
func NewSolcChecker(address common.Address, backend bind.ContractBackend) (*SolcChecker, error) {
	contract, err := bindSolcChecker(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &SolcChecker{SolcCheckerCaller: SolcCheckerCaller{contract: contract}, SolcCheckerTransactor: SolcCheckerTransactor{contract: contract}, SolcCheckerFilterer: SolcCheckerFilterer{contract: contract}}, nil
}

// NewSolcCheckerCaller creates a new read-only instance of SolcChecker, bound to a specific deployed contract.
func NewSolcCheckerCaller(address common.Address, caller bind.ContractCaller) (*SolcCheckerCaller, error) {
	contract, err := bindSolcChecker(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &SolcCheckerCaller{contract: contract}, nil
}

// NewSolcCheckerTransactor creates a new write-only instance of SolcChecker, bound to a specific deployed contract.
func NewSolcCheckerTransactor(address common.Address, transactor bind.ContractTransactor) (*SolcCheckerTransactor, error) {
	contract, err := bindSolcChecker(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &SolcCheckerTransactor{contract: contract}, nil
}

// NewSolcCheckerFilterer creates a new log filterer instance of SolcChecker, bound to a specific deployed contract.
func NewSolcCheckerFilterer(address common.Address, filterer bind.ContractFilterer) (*SolcCheckerFilterer, error) {
	contract, err := bindSolcChecker(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &SolcCheckerFilterer{contract: contract}, nil
}

// bindSolcChecker binds a generic wrapper to an already deployed contract.
func bindSolcChecker(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(SolcCheckerABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_SolcChecker *SolcCheckerRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _SolcChecker.Contract.SolcCheckerCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_SolcChecker *SolcCheckerRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _SolcChecker.Contract.SolcCheckerTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_SolcChecker *SolcCheckerRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _SolcChecker.Contract.SolcCheckerTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_SolcChecker *SolcCheckerCallerRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _SolcChecker.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_SolcChecker *SolcCheckerTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _SolcChecker.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_SolcChecker *SolcCheckerTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _SolcChecker.Contract.contract.Transact(opts, method, params...)
}

// F is a paid mutator transaction binding the contract method 0xd45754f8.
//
// Solidity: function f(bytes x) returns()
func (_SolcChecker *SolcCheckerTransactor) F(opts *bind.TransactOpts, x []byte) (*types.Transaction, error) {
	return _SolcChecker.contract.Transact(opts, "f", x)
}

// F is a paid mutator transaction binding the contract method 0xd45754f8.
//
// Solidity: function f(bytes x) returns()
func (_SolcChecker *SolcCheckerSession) F(x []byte) (*types.Transaction, error) {
	return _SolcChecker.Contract.F(&_SolcChecker.TransactOpts, x)
}

// F is a paid mutator transaction binding the contract method 0xd45754f8.
//
// Solidity: function f(bytes x) returns()
func (_SolcChecker *SolcCheckerTransactorSession) F(x []byte) (*types.Transaction, error) {
	return _SolcChecker.Contract.F(&_SolcChecker.TransactOpts, x)
}

// UsingOraclizeABI is the input ABI used to generate the binding from.
const UsingOraclizeABI = "[{\"constant\":false,\"inputs\":[{\"name\":\"_myid\",\"type\":\"bytes32\"},{\"name\":\"_result\",\"type\":\"string\"}],\"name\":\"__callback\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_myid\",\"type\":\"bytes32\"},{\"name\":\"_result\",\"type\":\"string\"},{\"name\":\"_proof\",\"type\":\"bytes\"}],\"name\":\"__callback\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]"

// UsingOraclizeBin is the compiled bytecode used for deploying new contracts.
const UsingOraclizeBin = `0x608060405234801561001057600080fd5b5061028b806100206000396000f3fe608060405234801561001057600080fd5b5060043610610052577c0100000000000000000000000000000000000000000000000000000000600035046327dc297e811461005757806338bbfa5014610106575b600080fd5b6101046004803603604081101561006d57600080fd5b8135919081019060408101602082013564010000000081111561008f57600080fd5b8201836020820111156100a157600080fd5b803590602001918460018302840111640100000000831117156100c357600080fd5b91908080601f01602080910402602001604051908101604052809392919081815260200183838082843760009201919091525092955061023a945050505050565b005b6101046004803603606081101561011c57600080fd5b8135919081019060408101602082013564010000000081111561013e57600080fd5b82018360208201111561015057600080fd5b8035906020019184600183028401116401000000008311171561017257600080fd5b91908080601f01602080910402602001604051908101604052809392919081815260200183838082843760009201919091525092959493602081019350359150506401000000008111156101c557600080fd5b8201836020820111156101d757600080fd5b803590602001918460018302840111640100000000831117156101f957600080fd5b91908080601f01602080910402602001604051908101604052809392919081815260200183838082843760009201919091525092955061025a945050505050565b604080516000815260208101909152610256908390839061025a565b5050565b50505056fea165627a7a72305820fa9b8a7b9a4c293579ff18b9fb631a3dd22e718dc17d0101b8e48504cc8b4b220029`

// DeployUsingOraclize deploys a new Ethereum contract, binding an instance of UsingOraclize to it.
func DeployUsingOraclize(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *UsingOraclize, error) {
	parsed, err := abi.JSON(strings.NewReader(UsingOraclizeABI))
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	address, tx, contract, err := bind.DeployContract(auth, parsed, common.FromHex(UsingOraclizeBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &UsingOraclize{UsingOraclizeCaller: UsingOraclizeCaller{contract: contract}, UsingOraclizeTransactor: UsingOraclizeTransactor{contract: contract}, UsingOraclizeFilterer: UsingOraclizeFilterer{contract: contract}}, nil
}

// UsingOraclize is an auto generated Go binding around an Ethereum contract.
type UsingOraclize struct {
	UsingOraclizeCaller     // Read-only binding to the contract
	UsingOraclizeTransactor // Write-only binding to the contract
	UsingOraclizeFilterer   // Log filterer for contract events
}

// UsingOraclizeCaller is an auto generated read-only Go binding around an Ethereum contract.
type UsingOraclizeCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// UsingOraclizeTransactor is an auto generated write-only Go binding around an Ethereum contract.
type UsingOraclizeTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// UsingOraclizeFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type UsingOraclizeFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// UsingOraclizeSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type UsingOraclizeSession struct {
	Contract     *UsingOraclize    // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// UsingOraclizeCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type UsingOraclizeCallerSession struct {
	Contract *UsingOraclizeCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts        // Call options to use throughout this session
}

// UsingOraclizeTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type UsingOraclizeTransactorSession struct {
	Contract     *UsingOraclizeTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts        // Transaction auth options to use throughout this session
}

// UsingOraclizeRaw is an auto generated low-level Go binding around an Ethereum contract.
type UsingOraclizeRaw struct {
	Contract *UsingOraclize // Generic contract binding to access the raw methods on
}

// UsingOraclizeCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type UsingOraclizeCallerRaw struct {
	Contract *UsingOraclizeCaller // Generic read-only contract binding to access the raw methods on
}

// UsingOraclizeTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type UsingOraclizeTransactorRaw struct {
	Contract *UsingOraclizeTransactor // Generic write-only contract binding to access the raw methods on
}

// NewUsingOraclize creates a new instance of UsingOraclize, bound to a specific deployed contract.
func NewUsingOraclize(address common.Address, backend bind.ContractBackend) (*UsingOraclize, error) {
	contract, err := bindUsingOraclize(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &UsingOraclize{UsingOraclizeCaller: UsingOraclizeCaller{contract: contract}, UsingOraclizeTransactor: UsingOraclizeTransactor{contract: contract}, UsingOraclizeFilterer: UsingOraclizeFilterer{contract: contract}}, nil
}

// NewUsingOraclizeCaller creates a new read-only instance of UsingOraclize, bound to a specific deployed contract.
func NewUsingOraclizeCaller(address common.Address, caller bind.ContractCaller) (*UsingOraclizeCaller, error) {
	contract, err := bindUsingOraclize(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &UsingOraclizeCaller{contract: contract}, nil
}

// NewUsingOraclizeTransactor creates a new write-only instance of UsingOraclize, bound to a specific deployed contract.
func NewUsingOraclizeTransactor(address common.Address, transactor bind.ContractTransactor) (*UsingOraclizeTransactor, error) {
	contract, err := bindUsingOraclize(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &UsingOraclizeTransactor{contract: contract}, nil
}

// NewUsingOraclizeFilterer creates a new log filterer instance of UsingOraclize, bound to a specific deployed contract.
func NewUsingOraclizeFilterer(address common.Address, filterer bind.ContractFilterer) (*UsingOraclizeFilterer, error) {
	contract, err := bindUsingOraclize(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &UsingOraclizeFilterer{contract: contract}, nil
}

// bindUsingOraclize binds a generic wrapper to an already deployed contract.
func bindUsingOraclize(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(UsingOraclizeABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_UsingOraclize *UsingOraclizeRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _UsingOraclize.Contract.UsingOraclizeCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_UsingOraclize *UsingOraclizeRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _UsingOraclize.Contract.UsingOraclizeTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_UsingOraclize *UsingOraclizeRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _UsingOraclize.Contract.UsingOraclizeTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_UsingOraclize *UsingOraclizeCallerRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _UsingOraclize.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_UsingOraclize *UsingOraclizeTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _UsingOraclize.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_UsingOraclize *UsingOraclizeTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _UsingOraclize.Contract.contract.Transact(opts, method, params...)
}

// Callback is a paid mutator transaction binding the contract method 0x38bbfa50.
//
// Solidity: function __callback(bytes32 _myid, string _result, bytes _proof) returns()
func (_UsingOraclize *UsingOraclizeTransactor) Callback(opts *bind.TransactOpts, _myid [32]byte, _result string, _proof []byte) (*types.Transaction, error) {
	return _UsingOraclize.contract.Transact(opts, "__callback", _myid, _result, _proof)
}

// Callback is a paid mutator transaction binding the contract method 0x38bbfa50.
//
// Solidity: function __callback(bytes32 _myid, string _result, bytes _proof) returns()
func (_UsingOraclize *UsingOraclizeSession) Callback(_myid [32]byte, _result string, _proof []byte) (*types.Transaction, error) {
	return _UsingOraclize.Contract.Callback(&_UsingOraclize.TransactOpts, _myid, _result, _proof)
}

// Callback is a paid mutator transaction binding the contract method 0x38bbfa50.
//
// Solidity: function __callback(bytes32 _myid, string _result, bytes _proof) returns()
func (_UsingOraclize *UsingOraclizeTransactorSession) Callback(_myid [32]byte, _result string, _proof []byte) (*types.Transaction, error) {
	return _UsingOraclize.Contract.Callback(&_UsingOraclize.TransactOpts, _myid, _result, _proof)
}
