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

// BufferABI is the input ABI used to generate the binding from.
const BufferABI = "[]"

// BufferBin is the compiled bytecode used for deploying new contracts.
const BufferBin = `0x604c602c600b82828239805160001a60731460008114601c57601e565bfe5b5030600052607381538281f3fe73000000000000000000000000000000000000000030146080604052600080fdfea165627a7a72305820d033f9c2b84a09f09de7c26e7d33b6829c655b7880e5e0e2153d8f71f35a29be0029`

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
const CBORBin = `0x604c602c600b82828239805160001a60731460008114601c57601e565bfe5b5030600052607381538281f3fe73000000000000000000000000000000000000000030146080604052600080fdfea165627a7a723058201c4f9c1be2989806a59240d31968fb620a6764d41ea3ebb5ce2e742f2c85ce350029`

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
// Solidity: function getAddress() returns(_address address)
func (_OraclizeAddrResolverI *OraclizeAddrResolverITransactor) GetAddress(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _OraclizeAddrResolverI.contract.Transact(opts, "getAddress")
}

// GetAddress is a paid mutator transaction binding the contract method 0x38cc4831.
//
// Solidity: function getAddress() returns(_address address)
func (_OraclizeAddrResolverI *OraclizeAddrResolverISession) GetAddress() (*types.Transaction, error) {
	return _OraclizeAddrResolverI.Contract.GetAddress(&_OraclizeAddrResolverI.TransactOpts)
}

// GetAddress is a paid mutator transaction binding the contract method 0x38cc4831.
//
// Solidity: function getAddress() returns(_address address)
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
// Solidity: function randomDS_getSessionPubKeyHash() constant returns(_sessionKeyHash bytes32)
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
// Solidity: function randomDS_getSessionPubKeyHash() constant returns(_sessionKeyHash bytes32)
func (_OraclizeI *OraclizeISession) RandomDSGetSessionPubKeyHash() ([32]byte, error) {
	return _OraclizeI.Contract.RandomDSGetSessionPubKeyHash(&_OraclizeI.CallOpts)
}

// RandomDSGetSessionPubKeyHash is a free data retrieval call binding the contract method 0xabaa5f3e.
//
// Solidity: function randomDS_getSessionPubKeyHash() constant returns(_sessionKeyHash bytes32)
func (_OraclizeI *OraclizeICallerSession) RandomDSGetSessionPubKeyHash() ([32]byte, error) {
	return _OraclizeI.Contract.RandomDSGetSessionPubKeyHash(&_OraclizeI.CallOpts)
}

// GetPrice is a paid mutator transaction binding the contract method 0x524f3889.
//
// Solidity: function getPrice(_datasource string) returns(_dsprice uint256)
func (_OraclizeI *OraclizeITransactor) GetPrice(opts *bind.TransactOpts, _datasource string) (*types.Transaction, error) {
	return _OraclizeI.contract.Transact(opts, "getPrice", _datasource)
}

// GetPrice is a paid mutator transaction binding the contract method 0x524f3889.
//
// Solidity: function getPrice(_datasource string) returns(_dsprice uint256)
func (_OraclizeI *OraclizeISession) GetPrice(_datasource string) (*types.Transaction, error) {
	return _OraclizeI.Contract.GetPrice(&_OraclizeI.TransactOpts, _datasource)
}

// GetPrice is a paid mutator transaction binding the contract method 0x524f3889.
//
// Solidity: function getPrice(_datasource string) returns(_dsprice uint256)
func (_OraclizeI *OraclizeITransactorSession) GetPrice(_datasource string) (*types.Transaction, error) {
	return _OraclizeI.Contract.GetPrice(&_OraclizeI.TransactOpts, _datasource)
}

// Query is a paid mutator transaction binding the contract method 0xadf59f99.
//
// Solidity: function query(_timestamp uint256, _datasource string, _arg string) returns(_id bytes32)
func (_OraclizeI *OraclizeITransactor) Query(opts *bind.TransactOpts, _timestamp *big.Int, _datasource string, _arg string) (*types.Transaction, error) {
	return _OraclizeI.contract.Transact(opts, "query", _timestamp, _datasource, _arg)
}

// Query is a paid mutator transaction binding the contract method 0xadf59f99.
//
// Solidity: function query(_timestamp uint256, _datasource string, _arg string) returns(_id bytes32)
func (_OraclizeI *OraclizeISession) Query(_timestamp *big.Int, _datasource string, _arg string) (*types.Transaction, error) {
	return _OraclizeI.Contract.Query(&_OraclizeI.TransactOpts, _timestamp, _datasource, _arg)
}

// Query is a paid mutator transaction binding the contract method 0xadf59f99.
//
// Solidity: function query(_timestamp uint256, _datasource string, _arg string) returns(_id bytes32)
func (_OraclizeI *OraclizeITransactorSession) Query(_timestamp *big.Int, _datasource string, _arg string) (*types.Transaction, error) {
	return _OraclizeI.Contract.Query(&_OraclizeI.TransactOpts, _timestamp, _datasource, _arg)
}

// Query2 is a paid mutator transaction binding the contract method 0x77228659.
//
// Solidity: function query2(_timestamp uint256, _datasource string, _arg1 string, _arg2 string) returns(_id bytes32)
func (_OraclizeI *OraclizeITransactor) Query2(opts *bind.TransactOpts, _timestamp *big.Int, _datasource string, _arg1 string, _arg2 string) (*types.Transaction, error) {
	return _OraclizeI.contract.Transact(opts, "query2", _timestamp, _datasource, _arg1, _arg2)
}

// Query2 is a paid mutator transaction binding the contract method 0x77228659.
//
// Solidity: function query2(_timestamp uint256, _datasource string, _arg1 string, _arg2 string) returns(_id bytes32)
func (_OraclizeI *OraclizeISession) Query2(_timestamp *big.Int, _datasource string, _arg1 string, _arg2 string) (*types.Transaction, error) {
	return _OraclizeI.Contract.Query2(&_OraclizeI.TransactOpts, _timestamp, _datasource, _arg1, _arg2)
}

// Query2 is a paid mutator transaction binding the contract method 0x77228659.
//
// Solidity: function query2(_timestamp uint256, _datasource string, _arg1 string, _arg2 string) returns(_id bytes32)
func (_OraclizeI *OraclizeITransactorSession) Query2(_timestamp *big.Int, _datasource string, _arg1 string, _arg2 string) (*types.Transaction, error) {
	return _OraclizeI.Contract.Query2(&_OraclizeI.TransactOpts, _timestamp, _datasource, _arg1, _arg2)
}

// Query2WithGasLimit is a paid mutator transaction binding the contract method 0x85dee34c.
//
// Solidity: function query2_withGasLimit(_timestamp uint256, _datasource string, _arg1 string, _arg2 string, _gasLimit uint256) returns(_id bytes32)
func (_OraclizeI *OraclizeITransactor) Query2WithGasLimit(opts *bind.TransactOpts, _timestamp *big.Int, _datasource string, _arg1 string, _arg2 string, _gasLimit *big.Int) (*types.Transaction, error) {
	return _OraclizeI.contract.Transact(opts, "query2_withGasLimit", _timestamp, _datasource, _arg1, _arg2, _gasLimit)
}

// Query2WithGasLimit is a paid mutator transaction binding the contract method 0x85dee34c.
//
// Solidity: function query2_withGasLimit(_timestamp uint256, _datasource string, _arg1 string, _arg2 string, _gasLimit uint256) returns(_id bytes32)
func (_OraclizeI *OraclizeISession) Query2WithGasLimit(_timestamp *big.Int, _datasource string, _arg1 string, _arg2 string, _gasLimit *big.Int) (*types.Transaction, error) {
	return _OraclizeI.Contract.Query2WithGasLimit(&_OraclizeI.TransactOpts, _timestamp, _datasource, _arg1, _arg2, _gasLimit)
}

// Query2WithGasLimit is a paid mutator transaction binding the contract method 0x85dee34c.
//
// Solidity: function query2_withGasLimit(_timestamp uint256, _datasource string, _arg1 string, _arg2 string, _gasLimit uint256) returns(_id bytes32)
func (_OraclizeI *OraclizeITransactorSession) Query2WithGasLimit(_timestamp *big.Int, _datasource string, _arg1 string, _arg2 string, _gasLimit *big.Int) (*types.Transaction, error) {
	return _OraclizeI.Contract.Query2WithGasLimit(&_OraclizeI.TransactOpts, _timestamp, _datasource, _arg1, _arg2, _gasLimit)
}

// QueryN is a paid mutator transaction binding the contract method 0x83eed3d5.
//
// Solidity: function queryN(_timestamp uint256, _datasource string, _argN bytes) returns(_id bytes32)
func (_OraclizeI *OraclizeITransactor) QueryN(opts *bind.TransactOpts, _timestamp *big.Int, _datasource string, _argN []byte) (*types.Transaction, error) {
	return _OraclizeI.contract.Transact(opts, "queryN", _timestamp, _datasource, _argN)
}

// QueryN is a paid mutator transaction binding the contract method 0x83eed3d5.
//
// Solidity: function queryN(_timestamp uint256, _datasource string, _argN bytes) returns(_id bytes32)
func (_OraclizeI *OraclizeISession) QueryN(_timestamp *big.Int, _datasource string, _argN []byte) (*types.Transaction, error) {
	return _OraclizeI.Contract.QueryN(&_OraclizeI.TransactOpts, _timestamp, _datasource, _argN)
}

// QueryN is a paid mutator transaction binding the contract method 0x83eed3d5.
//
// Solidity: function queryN(_timestamp uint256, _datasource string, _argN bytes) returns(_id bytes32)
func (_OraclizeI *OraclizeITransactorSession) QueryN(_timestamp *big.Int, _datasource string, _argN []byte) (*types.Transaction, error) {
	return _OraclizeI.Contract.QueryN(&_OraclizeI.TransactOpts, _timestamp, _datasource, _argN)
}

// QueryNWithGasLimit is a paid mutator transaction binding the contract method 0xc55c1cb6.
//
// Solidity: function queryN_withGasLimit(_timestamp uint256, _datasource string, _argN bytes, _gasLimit uint256) returns(_id bytes32)
func (_OraclizeI *OraclizeITransactor) QueryNWithGasLimit(opts *bind.TransactOpts, _timestamp *big.Int, _datasource string, _argN []byte, _gasLimit *big.Int) (*types.Transaction, error) {
	return _OraclizeI.contract.Transact(opts, "queryN_withGasLimit", _timestamp, _datasource, _argN, _gasLimit)
}

// QueryNWithGasLimit is a paid mutator transaction binding the contract method 0xc55c1cb6.
//
// Solidity: function queryN_withGasLimit(_timestamp uint256, _datasource string, _argN bytes, _gasLimit uint256) returns(_id bytes32)
func (_OraclizeI *OraclizeISession) QueryNWithGasLimit(_timestamp *big.Int, _datasource string, _argN []byte, _gasLimit *big.Int) (*types.Transaction, error) {
	return _OraclizeI.Contract.QueryNWithGasLimit(&_OraclizeI.TransactOpts, _timestamp, _datasource, _argN, _gasLimit)
}

// QueryNWithGasLimit is a paid mutator transaction binding the contract method 0xc55c1cb6.
//
// Solidity: function queryN_withGasLimit(_timestamp uint256, _datasource string, _argN bytes, _gasLimit uint256) returns(_id bytes32)
func (_OraclizeI *OraclizeITransactorSession) QueryNWithGasLimit(_timestamp *big.Int, _datasource string, _argN []byte, _gasLimit *big.Int) (*types.Transaction, error) {
	return _OraclizeI.Contract.QueryNWithGasLimit(&_OraclizeI.TransactOpts, _timestamp, _datasource, _argN, _gasLimit)
}

// QueryWithGasLimit is a paid mutator transaction binding the contract method 0xc51be90f.
//
// Solidity: function query_withGasLimit(_timestamp uint256, _datasource string, _arg string, _gasLimit uint256) returns(_id bytes32)
func (_OraclizeI *OraclizeITransactor) QueryWithGasLimit(opts *bind.TransactOpts, _timestamp *big.Int, _datasource string, _arg string, _gasLimit *big.Int) (*types.Transaction, error) {
	return _OraclizeI.contract.Transact(opts, "query_withGasLimit", _timestamp, _datasource, _arg, _gasLimit)
}

// QueryWithGasLimit is a paid mutator transaction binding the contract method 0xc51be90f.
//
// Solidity: function query_withGasLimit(_timestamp uint256, _datasource string, _arg string, _gasLimit uint256) returns(_id bytes32)
func (_OraclizeI *OraclizeISession) QueryWithGasLimit(_timestamp *big.Int, _datasource string, _arg string, _gasLimit *big.Int) (*types.Transaction, error) {
	return _OraclizeI.Contract.QueryWithGasLimit(&_OraclizeI.TransactOpts, _timestamp, _datasource, _arg, _gasLimit)
}

// QueryWithGasLimit is a paid mutator transaction binding the contract method 0xc51be90f.
//
// Solidity: function query_withGasLimit(_timestamp uint256, _datasource string, _arg string, _gasLimit uint256) returns(_id bytes32)
func (_OraclizeI *OraclizeITransactorSession) QueryWithGasLimit(_timestamp *big.Int, _datasource string, _arg string, _gasLimit *big.Int) (*types.Transaction, error) {
	return _OraclizeI.Contract.QueryWithGasLimit(&_OraclizeI.TransactOpts, _timestamp, _datasource, _arg, _gasLimit)
}

// SetCustomGasPrice is a paid mutator transaction binding the contract method 0xca6ad1e4.
//
// Solidity: function setCustomGasPrice(_gasPrice uint256) returns()
func (_OraclizeI *OraclizeITransactor) SetCustomGasPrice(opts *bind.TransactOpts, _gasPrice *big.Int) (*types.Transaction, error) {
	return _OraclizeI.contract.Transact(opts, "setCustomGasPrice", _gasPrice)
}

// SetCustomGasPrice is a paid mutator transaction binding the contract method 0xca6ad1e4.
//
// Solidity: function setCustomGasPrice(_gasPrice uint256) returns()
func (_OraclizeI *OraclizeISession) SetCustomGasPrice(_gasPrice *big.Int) (*types.Transaction, error) {
	return _OraclizeI.Contract.SetCustomGasPrice(&_OraclizeI.TransactOpts, _gasPrice)
}

// SetCustomGasPrice is a paid mutator transaction binding the contract method 0xca6ad1e4.
//
// Solidity: function setCustomGasPrice(_gasPrice uint256) returns()
func (_OraclizeI *OraclizeITransactorSession) SetCustomGasPrice(_gasPrice *big.Int) (*types.Transaction, error) {
	return _OraclizeI.Contract.SetCustomGasPrice(&_OraclizeI.TransactOpts, _gasPrice)
}

// SetProofType is a paid mutator transaction binding the contract method 0x688dcfd7.
//
// Solidity: function setProofType(_proofType bytes1) returns()
func (_OraclizeI *OraclizeITransactor) SetProofType(opts *bind.TransactOpts, _proofType [1]byte) (*types.Transaction, error) {
	return _OraclizeI.contract.Transact(opts, "setProofType", _proofType)
}

// SetProofType is a paid mutator transaction binding the contract method 0x688dcfd7.
//
// Solidity: function setProofType(_proofType bytes1) returns()
func (_OraclizeI *OraclizeISession) SetProofType(_proofType [1]byte) (*types.Transaction, error) {
	return _OraclizeI.Contract.SetProofType(&_OraclizeI.TransactOpts, _proofType)
}

// SetProofType is a paid mutator transaction binding the contract method 0x688dcfd7.
//
// Solidity: function setProofType(_proofType bytes1) returns()
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
// Solidity: e OwnershipTransferred(previousOwner indexed address, newOwner indexed address)
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
// Solidity: e OwnershipTransferred(previousOwner indexed address, newOwner indexed address)
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
// Solidity: function isPauser(account address) constant returns(bool)
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
// Solidity: function isPauser(account address) constant returns(bool)
func (_Pausable *PausableSession) IsPauser(account common.Address) (bool, error) {
	return _Pausable.Contract.IsPauser(&_Pausable.CallOpts, account)
}

// IsPauser is a free data retrieval call binding the contract method 0x46fbf68e.
//
// Solidity: function isPauser(account address) constant returns(bool)
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
// Solidity: function addPauser(account address) returns()
func (_Pausable *PausableTransactor) AddPauser(opts *bind.TransactOpts, account common.Address) (*types.Transaction, error) {
	return _Pausable.contract.Transact(opts, "addPauser", account)
}

// AddPauser is a paid mutator transaction binding the contract method 0x82dc1ec4.
//
// Solidity: function addPauser(account address) returns()
func (_Pausable *PausableSession) AddPauser(account common.Address) (*types.Transaction, error) {
	return _Pausable.Contract.AddPauser(&_Pausable.TransactOpts, account)
}

// AddPauser is a paid mutator transaction binding the contract method 0x82dc1ec4.
//
// Solidity: function addPauser(account address) returns()
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
// Solidity: e Paused(account address)
func (_Pausable *PausableFilterer) FilterPaused(opts *bind.FilterOpts) (*PausablePausedIterator, error) {

	logs, sub, err := _Pausable.contract.FilterLogs(opts, "Paused")
	if err != nil {
		return nil, err
	}
	return &PausablePausedIterator{contract: _Pausable.contract, event: "Paused", logs: logs, sub: sub}, nil
}

// WatchPaused is a free log subscription operation binding the contract event 0x62e78cea01bee320cd4e420270b5ea74000d11b0c9f74754ebdbfc544b05a258.
//
// Solidity: e Paused(account address)
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
// Solidity: e PauserAdded(account indexed address)
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
// Solidity: e PauserAdded(account indexed address)
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
// Solidity: e PauserRemoved(account indexed address)
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
// Solidity: e PauserRemoved(account indexed address)
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
// Solidity: e Unpaused(account address)
func (_Pausable *PausableFilterer) FilterUnpaused(opts *bind.FilterOpts) (*PausableUnpausedIterator, error) {

	logs, sub, err := _Pausable.contract.FilterLogs(opts, "Unpaused")
	if err != nil {
		return nil, err
	}
	return &PausableUnpausedIterator{contract: _Pausable.contract, event: "Unpaused", logs: logs, sub: sub}, nil
}

// WatchUnpaused is a free log subscription operation binding the contract event 0x5db9ee0a495bf2e6ff9c91a7834c1ba4fdd244a5e8aa4e537bd38aeae4b073aa.
//
// Solidity: e Unpaused(account address)
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
// Solidity: function isPauser(account address) constant returns(bool)
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
// Solidity: function isPauser(account address) constant returns(bool)
func (_PauserRole *PauserRoleSession) IsPauser(account common.Address) (bool, error) {
	return _PauserRole.Contract.IsPauser(&_PauserRole.CallOpts, account)
}

// IsPauser is a free data retrieval call binding the contract method 0x46fbf68e.
//
// Solidity: function isPauser(account address) constant returns(bool)
func (_PauserRole *PauserRoleCallerSession) IsPauser(account common.Address) (bool, error) {
	return _PauserRole.Contract.IsPauser(&_PauserRole.CallOpts, account)
}

// AddPauser is a paid mutator transaction binding the contract method 0x82dc1ec4.
//
// Solidity: function addPauser(account address) returns()
func (_PauserRole *PauserRoleTransactor) AddPauser(opts *bind.TransactOpts, account common.Address) (*types.Transaction, error) {
	return _PauserRole.contract.Transact(opts, "addPauser", account)
}

// AddPauser is a paid mutator transaction binding the contract method 0x82dc1ec4.
//
// Solidity: function addPauser(account address) returns()
func (_PauserRole *PauserRoleSession) AddPauser(account common.Address) (*types.Transaction, error) {
	return _PauserRole.Contract.AddPauser(&_PauserRole.TransactOpts, account)
}

// AddPauser is a paid mutator transaction binding the contract method 0x82dc1ec4.
//
// Solidity: function addPauser(account address) returns()
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
// Solidity: e PauserAdded(account indexed address)
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
// Solidity: e PauserAdded(account indexed address)
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
// Solidity: e PauserRemoved(account indexed address)
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
// Solidity: e PauserRemoved(account indexed address)
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

// PriceRollABI is the input ABI used to generate the binding from.
const PriceRollABI = "[{\"constant\":true,\"inputs\":[],\"name\":\"latest_roll\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"current_coin\",\"outputs\":[{\"name\":\"\",\"type\":\"uint8\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"query_stringETH\",\"outputs\":[{\"name\":\"\",\"type\":\"string\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"new_bonus\",\"type\":\"uint256\"}],\"name\":\"setBonus\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"new_gaslimit\",\"type\":\"uint256\"}],\"name\":\"setRandomGasLimit\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"pool\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"withdrawHouse\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_myid\",\"type\":\"bytes32\"},{\"name\":\"_result\",\"type\":\"string\"}],\"name\":\"__callback\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"expected_value\",\"type\":\"uint8\"},{\"name\":\"is_up\",\"type\":\"bool\"}],\"name\":\"placeBet\",\"outputs\":[],\"payable\":true,\"stateMutability\":\"payable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"round\",\"type\":\"uint256\"}],\"name\":\"claim\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_queryId\",\"type\":\"bytes32\"},{\"name\":\"_result\",\"type\":\"string\"},{\"name\":\"_proof\",\"type\":\"bytes\"}],\"name\":\"__callback\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[],\"name\":\"unpause\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"config_bonus_mult\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"account\",\"type\":\"address\"}],\"name\":\"isPauser\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"current_roll\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"new_cooldown\",\"type\":\"uint256\"}],\"name\":\"setCooldown\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"config_refund_delay\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"config_random_gas_limit\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"config_pricecheck_delay\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"paused\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"config_house_edge\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"rolls\",\"outputs\":[{\"name\":\"query_rng\",\"type\":\"bytes32\"},{\"name\":\"query_price1\",\"type\":\"bytes32\"},{\"name\":\"query_price2\",\"type\":\"bytes32\"},{\"name\":\"result_price1\",\"type\":\"uint256\"},{\"name\":\"result_price2\",\"type\":\"uint256\"},{\"name\":\"timestamp\",\"type\":\"uint256\"},{\"name\":\"pool\",\"type\":\"uint256\"},{\"name\":\"state\",\"type\":\"uint8\"},{\"name\":\"coin\",\"type\":\"uint8\"},{\"name\":\"result_rng\",\"type\":\"uint8\"},{\"name\":\"is_up\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"new_edge\",\"type\":\"uint256\"}],\"name\":\"setHouseEdge\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[],\"name\":\"renouncePauser\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[],\"name\":\"newRoll\",\"outputs\":[],\"payable\":true,\"stateMutability\":\"payable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"\",\"type\":\"address\"}],\"name\":\"balanceOf\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[],\"name\":\"renounceOwnership\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"amount\",\"type\":\"uint256\"},{\"name\":\"expected_value\",\"type\":\"uint8\"},{\"name\":\"is_up\",\"type\":\"bool\"}],\"name\":\"betFromInternalWallet\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"account\",\"type\":\"address\"}],\"name\":\"addPauser\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[],\"name\":\"destroy\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[],\"name\":\"pause\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"new_maxbet\",\"type\":\"uint256\"}],\"name\":\"setMaxBet\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"new_minbet\",\"type\":\"uint256\"}],\"name\":\"setMinBet\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"config_max_bet\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"isOwner\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"query_stringLTC\",\"outputs\":[{\"name\":\"\",\"type\":\"string\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"config_roll_cooldown\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"new_desination\",\"type\":\"address\"}],\"name\":\"setCutDestination\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[],\"name\":\"creditWallet\",\"outputs\":[],\"payable\":true,\"stateMutability\":\"payable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"withdrawWallet\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"userBalance\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"config_gas_limit\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"query_stringBTC\",\"outputs\":[{\"name\":\"\",\"type\":\"string\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"config_min_bet\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"new_gaslimit\",\"type\":\"uint256\"}],\"name\":\"setGasLimit\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"transferOwnership\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"new_delay\",\"type\":\"uint256\"}],\"name\":\"setPriceCheckDelay\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"house\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"payable\":true,\"stateMutability\":\"payable\",\"type\":\"fallback\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"name\":\"round\",\"type\":\"uint256\"}],\"name\":\"Rolling\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"name\":\"round\",\"type\":\"uint256\"}],\"name\":\"NewRoll\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"name\":\"round\",\"type\":\"uint256\"},{\"indexed\":false,\"name\":\"value\",\"type\":\"uint8\"},{\"indexed\":false,\"name\":\"start_price\",\"type\":\"uint256\"},{\"indexed\":false,\"name\":\"end_price\",\"type\":\"uint256\"}],\"name\":\"RollEnded\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"name\":\"round\",\"type\":\"uint256\"}],\"name\":\"RollRefunded\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"name\":\"round\",\"type\":\"uint256\"},{\"indexed\":true,\"name\":\"player\",\"type\":\"address\"},{\"indexed\":false,\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"RollClaimed\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"name\":\"round\",\"type\":\"uint256\"},{\"indexed\":true,\"name\":\"player\",\"type\":\"address\"},{\"indexed\":false,\"name\":\"amount\",\"type\":\"uint256\"},{\"indexed\":false,\"name\":\"expected_value\",\"type\":\"uint8\"},{\"indexed\":false,\"name\":\"is_up\",\"type\":\"bool\"}],\"name\":\"BetPlaced\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"OraclizeError\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"previousOwner\",\"type\":\"address\"},{\"indexed\":true,\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"OwnershipTransferred\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"name\":\"account\",\"type\":\"address\"}],\"name\":\"Paused\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"name\":\"account\",\"type\":\"address\"}],\"name\":\"Unpaused\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"account\",\"type\":\"address\"}],\"name\":\"PauserAdded\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"account\",\"type\":\"address\"}],\"name\":\"PauserRemoved\",\"type\":\"event\"}]"

// PriceRollBin is the compiled bytecode used for deploying new contracts.
const PriceRollBin = `0x6080604052603c6007819055610bb8600855620493e06009819055600a5566470de4df820000600b5568056bc75e2d63100000600c556014600d55604b600e55600f5560108054600160a060020a03191673a54741f7fe21689b59bd7eacbf3a2947cd3f3bd4179055600060118190556012556013805460ff191690553480156200008957600080fd5b506200009e336401000000006200010c810204565b60068054600160a860020a0319166101003381029190911791829055604051600160a060020a039190920416906000907f8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0908290a3620001066401000000006200015e810204565b620002b2565b620001276005826401000000006200429d6200020582021704565b604051600160a060020a038216907f6719d08c1888103bea251a4ed56406bd0c3e69723c8a1686e017e7bbe159b6f890600090a250565b6011546200017c90600164010000000062000baf6200026082021704565b60115560135460039060ff1660028111156200019457fe5b600101811515620001a157fe5b066002811115620001ae57fe5b6013805460ff19166001836002811115620001c557fe5b02179055504260125560115460408051918252517fb4267207f21a1858b0724dfba74816419152c88804d2b6276e5c32f6cd47c74a9181900360200190a1565b600160a060020a03811615156200021b57600080fd5b6200023082826401000000006200027a810204565b156200023b57600080fd5b600160a060020a0316600090815260209190915260409020805460ff19166001179055565b6000828201838110156200027357600080fd5b9392505050565b6000600160a060020a03821615156200029257600080fd5b50600160a060020a03166000908152602091909152604090205460ff1690565b61592280620002c26000396000f3fe6080604052600436106102b25760003560e060020a900480636fabd2651161017957806390e2c9dd116100e0578063d316401311610099578063ee7d72b411610073578063ee7d72b414610b13578063f2fde38b14610b3d578063fb8512f714610b70578063ff9b3acf14610b9a576102b2565b8063d316401314610ad4578063d8a3f00f14610ae9578063deedc5c514610afe576102b2565b806390e2c9dd14610a30578063987ae15414610a45578063aad7472214610a5a578063b8b069fc14610a8d578063be87ab3f14610a95578063bf15276514610abf576102b2565b80638456cb59116101325780638456cb591461096c578063881eff1e1461098157806388ea41b9146109ab5780638b1dc09e146109d55780638da5cb5b146109ea5780638f32d59b14610a1b576102b2565b80636fabd2651461089957806370a08231146108a1578063715018a6146108d45780637de36c97146108e957806382dc1ec41461092457806383197ef014610957576102b2565b8063402684121161021d5780635ad51f68116101d65780635ad51f681461076f5780635c975abb146107845780635d10c029146107995780635d69f16c146107ae5780636cd0f1021461085a5780636ef8d66d14610884576102b2565b806340268412146106aa57806346fbf68e146106bf5780634e0c12f0146107065780634fc3f41a1461071b578063550098991461074557806357b6702f1461075a576102b2565b8063232184111161026f578063232184111461042157806327dc297e1461044b5780632ed1317c14610504578063379607f51461052c57806338bbfa50146105565780633f4ba83a14610695576102b2565b8063022b02bc146102ca57806306792c07146102f1578063091085701461032a5780630b98f975146103b65780631361b3b2146103e257806316f0115b1461040c575b6015546102c5903463ffffffff610baf16565b601555005b3480156102d657600080fd5b506102df610bca565b60405190815260200160405180910390f35b3480156102fd57600080fd5b50610306610bd0565b6040518082600281111561031657fe5b60ff16815260200191505060405180910390f35b34801561033657600080fd5b5061033f610bd9565b60405160208082528190810183818151815260200191508051906020019080838360005b8381101561037b578082015183820152602001610363565b50505050905090810190601f1680156103a85780820380516001836020036101000a031916815260200191505b509250505060405180910390f35b3480156103c257600080fd5b506103e0600480360360208110156103d957600080fd5b5035610bf6565b005b3480156103ee57600080fd5b506103e06004803603602081101561040557600080fd5b5035610c67565b34801561041857600080fd5b506102df610c7f565b34801561042d57600080fd5b506103e06004803603602081101561044457600080fd5b5035610c85565b34801561045757600080fd5b506103e06004803603604081101561046e57600080fd5b8135919081019060408101602082013564010000000081111561049057600080fd5b8201836020820111156104a257600080fd5b803590602001918460018302840111640100000000831117156104c457600080fd5b91908080601f0160208091040260200160405190810160405281815292919060208401838380828437600092019190915250929550610d44945050505050565b6103e06004803603604081101561051a57600080fd5b5060ff81351690602001351515610d7b565b34801561053857600080fd5b506103e06004803603602081101561054f57600080fd5b5035610db8565b34801561056257600080fd5b506103e06004803603606081101561057957600080fd5b8135919081019060408101602082013564010000000081111561059b57600080fd5b8201836020820111156105ad57600080fd5b803590602001918460018302840111640100000000831117156105cf57600080fd5b91908080601f0160208091040260200160405190810160405281815292919060208401838380828437600092019190915250929594936020810193503591505064010000000081111561062157600080fd5b82018360208201111561063357600080fd5b8035906020019184600183028401116401000000008311171561065557600080fd5b91908080601f016020809104026020016040519081016040528181529291906020840183838082843760009201919091525092955061118f945050505050565b3480156106a157600080fd5b506103e0611559565b3480156106b657600080fd5b506102df6115c7565b3480156106cb57600080fd5b506106f2600480360360208110156106e257600080fd5b5035600160a060020a03166115cd565b604051901515815260200160405180910390f35b34801561071257600080fd5b506102df6115e8565b34801561072757600080fd5b506103e06004803603602081101561073e57600080fd5b50356115ee565b34801561075157600080fd5b506102df61165e565b34801561076657600080fd5b506102df611664565b34801561077b57600080fd5b506102df61166a565b34801561079057600080fd5b506106f2611670565b3480156107a557600080fd5b506102df61167a565b3480156107ba57600080fd5b506107d8600480360360208110156107d157600080fd5b5035611680565b604051808c81526020018b81526020018a815260200189815260200188815260200187815260200186815260200185600581111561081257fe5b60ff16815260200184600281111561082657fe5b60ff1681526020018360ff1660ff168152602001821515151581526020019b50505050505050505050505060405180910390f35b34801561086657600080fd5b506103e06004803603602081101561087d57600080fd5b50356116e3565b34801561089057600080fd5b506103e0611753565b6103e061175e565b3480156108ad57600080fd5b506102df600480360360208110156108c457600080fd5b5035600160a060020a0316611af6565b3480156108e057600080fd5b506103e0611b0a565b3480156108f557600080fd5b506103e06004803603606081101561090c57600080fd5b5080359060ff60208201351690604001351515611b7d565b34801561093057600080fd5b506103e06004803603602081101561094757600080fd5b5035600160a060020a0316611efd565b34801561096357600080fd5b506103e0611f1a565b34801561097857600080fd5b506103e0611f30565b34801561098d57600080fd5b506103e0600480360360208110156109a457600080fd5b5035611fa0565b3480156109b757600080fd5b506103e0600480360360208110156109ce57600080fd5b503561200f565b3480156109e157600080fd5b506102df61207e565b3480156109f657600080fd5b506109ff612084565b604051600160a060020a03909116815260200160405180910390f35b348015610a2757600080fd5b506106f2612098565b348015610a3c57600080fd5b5061033f6120ae565b348015610a5157600080fd5b506102df6120cb565b348015610a6657600080fd5b506103e060048036036020811015610a7d57600080fd5b5035600160a060020a03166120d1565b6103e0612106565b348015610aa157600080fd5b506103e060048036036020811015610ab857600080fd5b503561213a565b348015610acb57600080fd5b506102df612201565b348015610ae057600080fd5b506102df612216565b348015610af557600080fd5b5061033f61221c565b348015610b0a57600080fd5b506102df612239565b348015610b1f57600080fd5b506103e060048036036020811015610b3657600080fd5b503561223f565b348015610b4957600080fd5b506103e060048036036020811015610b6057600080fd5b5035600160a060020a0316612257565b348015610b7c57600080fd5b506103e060048036036020811015610b9357600080fd5b5035612273565b348015610ba657600080fd5b506102df6122cf565b600082820183811015610bc157600080fd5b90505b92915050565b60125481565b60135460ff1681565b6101806040519081016040526101498082526155de602083013981565b610bfe612098565b1515610c0957600080fd5b6101f4811115610c625760405160e560020a62461bcd02815260206004820152600760248201527f6d61782035302500000000000000000000000000000000000000000000000000604482015260640160405180910390fd5b600e55565b610c6f612098565b1515610c7a57600080fd5b600a55565b60155481565b610c8d612098565b1515610c9857600080fd5b601454811115610cf15760405160e560020a62461bcd02815260206004820152601760248201527f43616e742077697468647261772074686174206d756368000000000000000000604482015260640160405180910390fd5b601454610d04908263ffffffff6122d516565b601455601054600160a060020a031681156108fc0282604051600060405180830381858888f19350505050158015610d40573d6000803e3d6000fd5b5050565b610d4082826000604051818152601f19601f8301168101602001604052908015610d75576020820181803883390190505b5061118f565b3360009081526019602052610d9d90349060409020549063ffffffff610baf16565b33600090815260196020526040902055610d40348383611b7d565b600081815260166020526040812033600090815260088201602052909150604081206001810154909150600060ff90911611610e3d5760405160e560020a62461bcd02815260206004820152600c60248201527f6e6f74206120626574746f720000000000000000000000000000000000000000604482015260640160405180910390fd5b6004600783015460ff166005811115610e5257fe5b14158015610e7357506005600783015460ff166005811115610e7057fe5b14155b8015610e9257506000600783015460ff166005811115610e8f57fe5b14155b15610ebb576008546005830154429101108015610eb95760078301805460ff191660051790555b505b6005600783015460ff166005811115610ed057fe5b1415610f87578054600010610f2e5760405160e560020a62461bcd02815260206004820152601060248201527f416c726561647920726566756e64656400000000000000000000000000000000604482015260640160405180910390fd5b80543360009081526019602052610f51919060409020549063ffffffff610baf16565b33600090815260196020526040902055336000908152600883016020526040902060008155600101805461ffff1916905561118a565b6001810154600783015460ff808316620100008304821610926101009004811615156301000000909204161515148180610fbe5750805b15156110135760405160e560020a62461bcd02815260206004820152601460248201527f4e6f2077696e6e696e677320746f20636c61696d000000000000000000000000604482015260640160405180910390fd5b60008061105085606060405190810160409081528254825260019092015460ff8082166020840152610100909104161515918101919091526122ea565b905083156110b5576110b26110a58261109988606060405190810160409081528254825260019092015460ff80821660208401526101009091041615159181019190915261230c565b9063ffffffff6122d516565b839063ffffffff610baf16565b91505b82156110ec576110e96110a56103e86110dd600e54896000015461233990919063ffffffff16565b9063ffffffff61236416565b91505b336000908152601960205261110e90839060409020549063ffffffff610baf16565b33600090815260196020526040902055601554611131908363ffffffff6122d516565b601555601454611147908263ffffffff610baf16565b601455337f170c93f55dab07b2a0c856acad14fbb479fc97d6c3f489faf09b5f01c1575d3e888460405191825260208201526040908101905180910390a2505050505b505050565b611197612388565b600160a060020a031633146111f55760405160e560020a62461bcd02815260206004820152600b60248201527f61757468206661696c6564000000000000000000000000000000000000000000604482015260640160405180910390fd5b60008381526018602052604090205460ff16156112465760405160e560020a62461bcd0281526004018080602001828103825260218152602001806157276021913960400191505060405180910390fd5b6000838152601860205260019060409020805460ff19169115159190911790556000838152601760205260408120549050600081116112ce5760405160e560020a62461bcd02815260206004820152601060248201527f496e76616c6964205f7175657279496400000000000000000000000000000000604482015260640160405180910390fd5b600081815260166020526040812080549091508514156113c0576112f385858561255a565b60ff16156113425760078101805460ff191660051790557f796cfdf016e01ffdff99dbe8ce27df476b06d0fae2ac92ec3889b6008219ed528260405190815260200160405180910390a16113bb565b60006064858051906020012081151561135757fe5b60078401805462ff00001916620100009390920660010160ff81811694909402929092179081905590925016600581111561138e57fe5b600101600581111561139c57fe5b60078301805460ff191660018360058111156113b457fe5b0217905550505b61148d565b806001015485141561141d576113d584612671565b6003820155600781015460ff1660058111156113ed57fe5b60010160058111156113fb57fe5b60078201805460ff1916600183600581111561141357fe5b021790555061148d565b806002015485141561144a5761143284612671565b6004820155600781015460ff1660058111156113ed57fe5b60078101805460ff191660051790557f796cfdf016e01ffdff99dbe8ce27df476b06d0fae2ac92ec3889b6008219ed528260405190815260200160405180910390a15b6004600782015460ff1660058111156114a257fe5b1415611552576004810154600382015460078301805463ff00000019169290911063010000000291909117905560068101546015546114e091610baf565b6015819055507ffa1b3a1a2a7d28c6303b4f4d3f3dc998928c92532bf19e80f8bfa461528e4569828260070160029054906101000a900460ff168360030154846004015460405193845260ff909216602084015260408084019190915260608301919091526080909101905180910390a15b5050505050565b611562336115cd565b151561156d57600080fd5b60065460ff16151561157e57600080fd5b6006805460ff191690557f5db9ee0a495bf2e6ff9c91a7834c1ba4fdd244a5e8aa4e537bd38aeae4b073aa33604051600160a060020a03909116815260200160405180910390a1565b600e5481565b60006115e060058363ffffffff61275816565b90505b919050565b60115481565b6115f6612098565b151561160157600080fd5b603c8110156116595760405160e560020a62461bcd02815260206004820152601360248201527f4d696e696d756d2069732031206d696e75746500000000000000000000000000604482015260640160405180910390fd5b600755565b60085481565b600a5481565b600f5481565b60065460ff165b90565b600d5481565b601660205280600052604060002080546001820154600283015460038401546004850154600586015460068701546007909701549597509395929491939092909160ff80821691610100810482169162010000820481169163010000009004168b565b6116eb612098565b15156116f657600080fd5b603281111561174e5760405160e560020a62461bcd02815260206004820152600660248201527f6d61782035250000000000000000000000000000000000000000000000000000604482015260640160405180910390fd5b600d55565b61175c33612792565b565b60065460ff161561176e57600080fd5b60075460125442910111156117cc5760405160e560020a62461bcd02815260206004820152601460248201527f726f6c6c20697320636f6f6c696e6720646f776e000000000000000000000000604482015260640160405180910390fd5b60006117d66127dc565b905034811115611852577f0510c169028b6c6e93e31f93b1b642b756cea80cc1966887b5c8f0b93d3d5b688160405190815260200160405180910390a1600034111561184d57333480156108fc0290604051600060405180830381858888f1935050505015801561184b573d6000803e3d6000fd5b505b611af3565b601154600090815260166020526040812090506060600060135460ff16600281111561187a57fe5b14156118a1576101806040519081016040526101498082526155de602083013990506118f8565b600160135460ff1660028111156118b457fe5b14156118db57610180604051908101604052610149808252615446602083013990506118f8565b6101806040519081016040526101498082526157ae602083013990505b6119217f11000000000000000000000000000000000000000000000000000000000000006128cb565b611964600060408051908101604052600681527f6e6573746564000000000000000000000000000000000000000000000000000060208201526009548490612aa2565b6001830155600f546119ae9060408051908101604052600681527f6e6573746564000000000000000000000000000000000000000000000000000060208201526009548490612aa2565b60028301556119dc7f30000000000000000000000000000000000000000000000000000000000000006128cb565b6119eb60006001600a54612e6e565b825542600583015560078201805460ff191660011780825560135460ff16919061ff001916610100836002811115611a1f57fe5b02179055506011548254600090815260176020526040902055601154600183015460009081526017602052604090205560115460028301546000908152601760205260409020556011547fe86c6fb9a82a2c1fd4c6c98a2fe03844f019514b0cfe8dfbdf076c61f124c0d49060405190815260200160405180910390a1611aa461334a565b6000611ab6348563ffffffff6122d516565b90506000811115611aef573381156108fc0282604051600060405180830381858888f19350505050158015611552573d6000803e3d6000fd5b5050505b50565b601960205280600052604060002054905081565b611b12612098565b1515611b1d57600080fd5b6006546000906101009004600160a060020a03167f8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e060405160405180910390a36006805474ffffffffffffffffffffffffffffffffffffffff0019169055565b60065460ff1615611b8d57600080fd5b3360009081526019602052839060409020541015611bdf5760405160e560020a62461bcd0281526004018080602001828103825260268152602001806157486026913960400191505060405180910390fd5b60018260ff16118015611bf5575060648260ff16105b1515611c355760405160e560020a62461bcd02815260040180806020018281038252602e81526020018061558f602e913960400191505060405180910390fd5b600c54831115611c8e5760405160e560020a62461bcd02815260206004820152600c60248201527f42657420746f6f20686967680000000000000000000000000000000000000000604482015260640160405180910390fd5b600b54831015611ce75760405160e560020a62461bcd02815260206004820152600b60248201527f42657420746f6f206c6f77000000000000000000000000000000000000000000604482015260640160405180910390fd5b60115460009081526016602052604081203360009081526008820160205290915060408120805490915015611d655760405160e560020a62461bcd02815260206004820152601460248201527f416c726561647920706c61636564206120626574000000000000000000000000604482015260640160405180910390fd5b84815560018101805460ff191660ff86161761ff001916610100851515021790556000611dc282606060405190810160409081528254825260019092015460ff8082166020840152610100909104161515918101919091526133e4565b90506000611de56103e86110dd600e54866000015461233990919063ffffffff16565b9050611df7828263ffffffff610baf16565b6015541015611e4f5760405160e560020a62461bcd02815260206004820152601f60248201527f4e6f7420656e6f75676820696e20706f6f6c20666f7220746869732062657400604482015260640160405180910390fd5b3360009081526019602052611e7190889060409020549063ffffffff6122d516565b3360009081526019602052604090205582546006850154611e979163ffffffff610baf16565b600685015560115433907f917ae17adb72fe5dfbe4db3784fc9c4489530a50dbd693383d1502cfe49e68fd90898989604051938452602084019290925260ff1660408084019190915290151560608301526080909101905180910390a250505050505050565b611f06336115cd565b1515611f1157600080fd5b611af381613429565b611f22612098565b1515611f2d57600080fd5b33ff5b611f39336115cd565b1515611f4457600080fd5b60065460ff1615611f5457600080fd5b6006805460ff191660011790557f62e78cea01bee320cd4e420270b5ea74000d11b0c9f74754ebdbfc544b05a25833604051600160a060020a03909116815260200160405180910390a1565b611fa8612098565b1515611fb357600080fd5b6000811161200a5760405160e560020a62461bcd02815260206004820152601960248201527f4d7573742062652067726561746572207468616e207a65726f00000000000000604482015260640160405180910390fd5b600c55565b612017612098565b151561202257600080fd5b600081116120795760405160e560020a62461bcd02815260206004820152601960248201527f4d7573742062652067726561746572207468616e207a65726f00000000000000604482015260640160405180910390fd5b600b55565b600c5481565b6006546101009004600160a060020a031690565b6006546101009004600160a060020a0316331490565b6101806040519081016040526101498082526157ae602083013981565b60075481565b6120d9612098565b15156120e457600080fd5b60108054600160a060020a031916600160a060020a0392909216919091179055565b336000908152601960205261212890349060409020549063ffffffff610baf16565b33600090815260196020526040902055565b33600090815260196020528190604090205410156121a15760405160e560020a62461bcd02815260206004820152601060248201527f4e6f7420656e6f7567682066756e647300000000000000000000000000000000604482015260640160405180910390fd5b33600090815260196020526121c390829060409020549063ffffffff6122d516565b336000908152601960205260409020553381156108fc0282604051600060405180830381858888f19350505050158015610d40573d6000803e3d6000fd5b33600090815260196020526040812054905090565b60095481565b610180604051908101604052610149808252615446602083013981565b600b5481565b612247612098565b151561225257600080fd5b600955565b61225f612098565b151561226a57600080fd5b611af381613473565b61227b612098565b151561228657600080fd5b624f1a0081106122ca5760405160e560020a62461bcd0281526004018080602001828103825260218152602001806155bd6021913960400191505060405180910390fd5b600f55565b60145481565b6000828211156122e457600080fd5b50900390565b6000600d546103e86122fb8461230c565b81151561230457fe5b040292915050565b6000806001836020015160ff16039050825181606481900385510281151561233057fe5b04019392505050565b600082151561234a57506000610bc4565b82820282848281151561235957fe5b0414610bc157600080fd5b600080821161237257600080fd5b6000828481151561237f57fe5b04949350505050565b600154600090600160a060020a031615806123b557506001546123b390600160a060020a03166134fe565b155b156123c6576123c46000613502565b505b600154600160a060020a03166338cc48316040518163ffffffff1660e060020a028152600401602060405180830381600087803b15801561240657600080fd5b505af115801561241a573d6000803e3d6000fd5b505050506040513d602081101561243057600080fd5b810190808051600054600160a060020a0390811691161492506124e491505057600154600160a060020a03166338cc48316040518163ffffffff1660e060020a028152600401602060405180830381600087803b15801561249057600080fd5b505af11580156124a4573d6000803e3d6000fd5b505050506040513d60208110156124ba57600080fd5b81019080805160008054600160a060020a031916600160a060020a03929092169190911790555050505b600054600160a060020a031663c281d19e6040518163ffffffff1660e060020a02815260040160206040518083038186803b15801561252257600080fd5b505afa158015612536573d6000803e3d6000fd5b505050506040513d602081101561254c57600080fd5b810190808051935050505090565b60008160008151811061256957fe5b016020015160f860020a900460f860020a02600160f860020a0319167f4c000000000000000000000000000000000000000000000000000000000000001415806125fb5750816001815181106125bb57fe5b016020015160f860020a900460f860020a02600160f860020a0319167f500000000000000000000000000000000000000000000000000000000000000014155b8061262f575060018260028151811061261057fe5b016020015160f860020a900460f860020a0260f860020a900460ff1614155b1561263c5750600161266a565b600061265183868661264c61350c565b6135ae565b905080151561266457600291505061266a565b60009150505b9392505050565b600060028183815b815181101561273757821561268f576001840393505b81818151811061269b57fe5b016020015160f860020a900460f860020a0260f860020a900460ff16602e14156126c457600192505b60008282815181106126d257fe5b016020015160f860020a900460f860020a0260f860020a900460ff16905060308110158015612702575060398111155b15612713576030810386600a020195505b83801561271e575084155b1561272e57506115e39350505050565b50600101612679565b5b831561274f5784600a029450600184039350612738565b50505050919050565b6000600160a060020a038216151561276f57600080fd5b600160a060020a03821660009081526020849052604090205460ff169392505050565b6127a360058263ffffffff613be816565b80600160a060020a03167fcd265ebaf09df2871cc7bd4133404a235ba12eff2041bb89d9c714a2621c7c7e60405160405180910390a250565b60006128077f11000000000000000000000000000000000000000000000000000000000000006128cb565b600061284860408051908101604052600381527f55524c00000000000000000000000000000000000000000000000000000000006020820152600954613c3d565b60020290506128767f30000000000000000000000000000000000000000000000000000000000000006128cb565b6128c56128b860408051908101604052600681527f52616e646f6d00000000000000000000000000000000000000000000000000006020820152600a54613c3d565b829063ffffffff610baf16565b91505090565b600154600160a060020a031615806128f557506001546128f390600160a060020a03166134fe565b155b15612906576129046000613502565b505b600154600160a060020a03166338cc48316040518163ffffffff1660e060020a028152600401602060405180830381600087803b15801561294657600080fd5b505af115801561295a573d6000803e3d6000fd5b505050506040513d602081101561297057600080fd5b810190808051600054600160a060020a039081169116149250612a2491505057600154600160a060020a03166338cc48316040518163ffffffff1660e060020a028152600401602060405180830381600087803b1580156129d057600080fd5b505af11580156129e4573d6000803e3d6000fd5b505050506040513d60208110156129fa57600080fd5b81019080805160008054600160a060020a031916600160a060020a03929092169190911790555050505b600054600160a060020a031663688dcfd78260405160e060020a63ffffffff84160281527fff000000000000000000000000000000000000000000000000000000000000009091166004820152602401600060405180830381600087803b158015612a8e57600080fd5b505af1158015611552573d6000803e3d6000fd5b600154600090600160a060020a03161580612acf5750600154612acd90600160a060020a03166134fe565b155b15612ae057612ade6000613502565b505b600154600160a060020a03166338cc48316040518163ffffffff1660e060020a028152600401602060405180830381600087803b158015612b2057600080fd5b505af1158015612b34573d6000803e3d6000fd5b505050506040513d6020811015612b4a57600080fd5b810190808051600054600160a060020a039081169116149250612bfe91505057600154600160a060020a03166338cc48316040518163ffffffff1660e060020a028152600401602060405180830381600087803b158015612baa57600080fd5b505af1158015612bbe573d6000803e3d6000fd5b505050506040513d6020811015612bd457600080fd5b81019080805160008054600160a060020a031916600160a060020a03929092169190911790555050505b60008054600160a060020a0316632ef3accc86856040518363ffffffff1660e060020a0281526004018080602001838152602001828103825284818151815260200191508051906020019080838360005b83811015612c67578082015183820152602001612c4f565b50505050905090810190601f168015612c945780820380516001836020036101000a031916815260200191505b509350505050602060405180830381600087803b158015612cb457600080fd5b505af1158015612cc8573d6000803e3d6000fd5b505050506040513d6020811015612cde57600080fd5b8101908080519350505050670de0b6b3a76400003a840201811115612d07575060009050612e66565b600054600160a060020a031663c51be90f82888888886040518663ffffffff1660e060020a028152600401808581526020018060200180602001848152602001838103835286818151815260200191508051906020019080838360005b83811015612d7c578082015183820152602001612d64565b50505050905090810190601f168015612da95780820380516001836020036101000a031916815260200191505b50838103825285818151815260200191508051906020019080838360005b83811015612ddf578082015183820152602001612dc7565b50505050905090810190601f168015612e0c5780820380516001836020036101000a031916815260200191505b5096505050505050506020604051808303818588803b158015612e2e57600080fd5b505af1158015612e42573d6000803e3d6000fd5b50505050506040513d6020811015612e5957600080fd5b8101908080519450505050505b949350505050565b60008083118015612e80575060208311155b1515612e8b57600080fd5b600a8402935060606001604051818152601f19601f8301168101602001604052908015612ebf576020820181803883390190505b5090508360f860020a0281600081518110612ed657fe5b906020010190600160f860020a031916908160001a90535060606020604051818152601f19601f8301168101602001604052908015612f1c576020820181803883390190505b50905060606020604051818152601f19601f8301168101602001604052908015612f4d576020820181803883390190505b5090506000612f5a613e88565b90506020835242411860014303401860208401526020825280602083015260606020604051818152601f19601f8301168101602001604052908015612fa6576020820181803883390190505b50905088602082015260606008604051818152601f19601f8301168101602001604052908015612fdd576020820181803883390190505b509050612ff08260186008846000614022565b50612ff961536e565b608060405190810160405280878152602001888152602001868152602001848152509050600061305d60408051908101604052600681527f72616e646f6d00000000000000000000000000000000000000000000000000006020820152838c61406c565b905060606008604051818152601f19601f830116810160200160405290801561308d576020820181803883390190505b509050602084015160f860020a810460278301537e01000000000000000000000000000000000000000000000000000000000000810460268301537d0100000000000000000000000000000000000000000000000000000000008104602583015360e060020a810460248301537b01000000000000000000000000000000000000000000000000000000810460238301537a01000000000000000000000000000000000000000000000000000081046022830153790100000000000000000000000000000000000000000000000000810460218301537801000000000000000000000000000000000000000000000000810460208301535061333a82826020860151600287516040518082805190602001908083835b602083106131c25780518252601f1990920191602091820191016131a3565b6001836020036101000a038019825116818451168082178552505050505050905001915050602060405180830381855afa158015613204573d6000803e3d6000fd5b5050506040513d602081101561321957600080fd5b8101908080519250505060408801516040516020018085805190602001908083835b6020831061325a5780518252601f19909201916020918201910161323b565b6001836020036101000a038019825116818451161790925250505091909101905084805190602001908083835b602083106132a65780518252601f199092019160209182019101613287565b6001836020036101000a038019825116818451161790925250505091909101848152602001905082805190602001908083835b602083106132f85780518252601f1990920191602091820191016132d9565b6001836020036101000a038019825116818451161790925250505091909101955060409450505050505160208183030381529060405280519060200120614288565b509b9a5050505050505050505050565b60115461335e90600163ffffffff610baf16565b60115560135460039060ff16600281111561337557fe5b60010181151561338157fe5b06600281111561338d57fe5b6013805460ff191660018360028111156133a357fe5b0217905550426012556011547fb4267207f21a1858b0724dfba74816419152c88804d2b6276e5c32f6cd47c74a9060405190815260200160405180910390a1565b6000806001836020015160ff16039050600d546103e8036103e8846000015183606481900387510281151561341557fe5b040181151561342057fe5b04029392505050565b61343a60058263ffffffff61429d16565b80600160a060020a03167f6719d08c1888103bea251a4ed56406bd0c3e69723c8a1686e017e7bbe159b6f860405160405180910390a250565b600160a060020a038116151561348857600080fd5b600654600160a060020a03808316916101009004167f8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e060405160405180910390a360068054600160a060020a039092166101000274ffffffffffffffffffffffffffffffffffffffff0019909216919091179055565b3b90565b60006115e06142f4565b606060028054600181600116156101000203166002900480601f0160208091040260200160405190810160405280929190818152602001828054600181600116156101000203166002900480156135a45780601f10613579576101008083540402835291602001916135a4565b820191906000526020600020905b81548152906001019060200180831161358757829003601f168201915b5050505050905090565b600080856045815181106135be57fe5b016020015160f860020a900460f860020a0260f860020a900460ff16600201604401602001905060606020604051818152601f19601f8301168101602001604052908015613613576020820181803883390190505b50905061362587836020846000614022565b50600284876040516020018083805190602001908083835b6020831061365c5780518252601f19909201916020918201910161363d565b6001836020036101000a038019825116818451161790925250505091909101928352505060200190506040516020818303038152906040526040518082805190602001908083835b602083106136c35780518252601f1990920191602091820191016136a4565b6001836020036101000a038019825116818451168082178552505050505050905001915050602060405180830381855afa158015613705573d6000803e3d6000fd5b5050506040513d602081101561371a57600080fd5b8101908080519250604091505051602001808281526020019150506040516020818303038152906040528051906020012081805190602001201461376357600092505050612e66565b606087604a84018151811061377457fe5b016020015160f860020a900460f860020a0260f860020a900460ff16600201604051818152601f19601f83011681016020016040529080156137bd576020820181803883390190505b5090506137d288846049018351846000614022565b5061389c6002826040518082805190602001908083835b602083106138085780518252601f1990920191602091820191016137e9565b6001836020036101000a038019825116818451168082178552505050505050905001915050602060405180830381855afa15801561384a573d6000803e3d6000fd5b5050506040513d602081101561385f57600080fd5b81019080805192508991508b9050602887018151811061387b57fe5b016020015160f860020a900460f860020a0260f860020a900460ff16614612565b15156138ae5760009350505050612e66565b60606029604051818152601f19601f83011681016020016040529080156138dc576020820181803883390190505b5090506138f189856020016029846000614022565b50606060408051818152601f19601f830116810160200160405290801561391f576020820181803883390190505b50905060008351866020016029010160410190506139458b604083036040856000614022565b5060006002836040518082805190602001908083835b6020831061397a5780518252601f19909201916020918201910161395b565b6001836020036101000a038019825116818451168082178552505050505050905001915050602060405180830381855afa1580156139bc573d6000803e3d6000fd5b5050506040513d60208110156139d157600080fd5b810190808051935086925083915060409050516020018083805190602001908083835b60208310613a135780518252601f1990920191602091820191016139f4565b6001836020036101000a038019825116818451161790925250505091909101928352505060200190506040516020818303038152906040528051906020012060008c8152600360205260409020541415613a7e5760008b815260036020526040902060009055613a8e565b6000975050505050505050612e66565b60606049604051818152601f19601f8301168101602001604052908015613abc576020820181803883390190505b509050613ace8d896049846000614022565b50613b6e6002826040518082805190602001908083835b60208310613b045780518252601f199092019160209182019101613ae5565b6001836020036101000a038019825116818451168082178552505050505050905001915050602060405180830381855afa158015613b46573d6000803e3d6000fd5b5050506040513d6020811015613b5b57600080fd5b8101908080519250899150879050614692565b1515613b8557600098505050505050505050612e66565b60008281526004602052604090205460ff161515613bc557613ba78d84614814565b6000838152600460205260409020805460ff19169115159190911790555b60008281526004602052604090205460ff169d9c50505050505050505050505050565b600160a060020a0381161515613bfd57600080fd5b613c078282612758565b1515613c1257600080fd5b600160a060020a0381166000908152602083905260408120805460ff19169115159190911790555050565b600154600090600160a060020a03161580613c6a5750600154613c6890600160a060020a03166134fe565b155b15613c7b57613c796000613502565b505b600154600160a060020a03166338cc48316040518163ffffffff1660e060020a028152600401602060405180830381600087803b158015613cbb57600080fd5b505af1158015613ccf573d6000803e3d6000fd5b505050506040513d6020811015613ce557600080fd5b810190808051600054600160a060020a039081169116149250613d9991505057600154600160a060020a03166338cc48316040518163ffffffff1660e060020a028152600401602060405180830381600087803b158015613d4557600080fd5b505af1158015613d59573d6000803e3d6000fd5b505050506040513d6020811015613d6f57600080fd5b81019080805160008054600160a060020a031916600160a060020a03929092169190911790555050505b600054600160a060020a0316632ef3accc84846040518363ffffffff1660e060020a0281526004018080602001838152602001828103825284818151815260200191508051906020019080838360005b83811015613e01578082015183820152602001613de9565b50505050905090810190601f168015613e2e5780820380516001836020036101000a031916815260200191505b509350505050602060405180830381600087803b158015613e4e57600080fd5b505af1158015613e62573d6000803e3d6000fd5b505050506040513d6020811015613e7857600080fd5b8101908080519695505050505050565b600154600090600160a060020a03161580613eb55750600154613eb390600160a060020a03166134fe565b155b15613ec657613ec46000613502565b505b600154600160a060020a03166338cc48316040518163ffffffff1660e060020a028152600401602060405180830381600087803b158015613f0657600080fd5b505af1158015613f1a573d6000803e3d6000fd5b505050506040513d6020811015613f3057600080fd5b810190808051600054600160a060020a039081169116149250613fe491505057600154600160a060020a03166338cc48316040518163ffffffff1660e060020a028152600401602060405180830381600087803b158015613f9057600080fd5b505af1158015613fa4573d6000803e3d6000fd5b505050506040513d6020811015613fba57600080fd5b81019080805160008054600160a060020a031916600160a060020a03929092169190911790555050505b600054600160a060020a031663abaa5f3e6040518163ffffffff1660e060020a02815260040160206040518083038186803b15801561252257600080fd5b6060838201808451101561403557600080fd5b60208087019084015b86886020010182101561405f5781890151868201526020918201910161403e565b5093979650505050505050565b600154600090600160a060020a03161580614099575060015461409790600160a060020a03166134fe565b155b156140aa576140a86000613502565b505b600154600160a060020a03166338cc48316040518163ffffffff1660e060020a028152600401602060405180830381600087803b1580156140ea57600080fd5b505af11580156140fe573d6000803e3d6000fd5b505050506040513d602081101561411457600080fd5b810190808051600054600160a060020a0390811691161492506141c891505057600154600160a060020a03166338cc48316040518163ffffffff1660e060020a028152600401602060405180830381600087803b15801561417457600080fd5b505af1158015614188573d6000803e3d6000fd5b505050506040513d602081101561419e57600080fd5b81019080805160008054600160a060020a031916600160a060020a03929092169190911790555050505b606060046040519080825280602002602001820160405280156141ff57816020015b60608152602001906001900390816141ea5790505b50905083518160008151811061421157fe5b6020908102909101015283600160200201518160018151811061423057fe5b602090810291909101015260408401518160028151811061424d57fe5b602090810291909101015260608401518160038151811061426a57fe5b6020908102909101015261427f858285614bf3565b95945050505050565b60008281526003602052819060409020555050565b600160a060020a03811615156142b257600080fd5b6142bc8282612758565b156142c657600080fd5b600160a060020a0381166000908152602083905260019060409020805460ff19169115159190911790555050565b600080614314731d3b2638a7cc9f2cb3d298a3da7a90b67e5506ed6134fe565b11156143845760018054600160a060020a031916731d3b2638a7cc9f2cb3d298a3da7a90b67e5506ed17905561437c60408051908101604052600b81527f6574685f6d61696e6e65740000000000000000000000000000000000000000006020820152614fd0565b506001611677565b60006143a373c03a2615d5efaf5f49f60b7bb6583eaec212fdf16134fe565b111561440b5760018054600160a060020a03191673c03a2615d5efaf5f49f60b7bb6583eaec212fdf117905561437c60408051908101604052600c81527f6574685f726f707374656e3300000000000000000000000000000000000000006020820152614fd0565b600061442a73b7a07bcf2ba2f2703b24c0691b5278999c59ac7e6134fe565b11156144925760018054600160a060020a03191673b7a07bcf2ba2f2703b24c0691b5278999c59ac7e17905561437c60408051908101604052600981527f6574685f6b6f76616e00000000000000000000000000000000000000000000006020820152614fd0565b60006144b173146500cfd35b22e4a392fe0adc06de1a1368ed486134fe565b11156145195760018054600160a060020a03191673146500cfd35b22e4a392fe0adc06de1a1368ed4817905561437c60408051908101604052600b81527f6574685f72696e6b6562790000000000000000000000000000000000000000006020820152614fd0565b6000614538736f485c8bf6fc43ea212e93bbf8ce046c7f1cb4756134fe565b111561456a575060018054600160a060020a031916736f485c8bf6fc43ea212e93bbf8ce046c7f1cb475178155611677565b60006145897320e12a1f859b3feae5fb2a0a32c18f5a65555bbf6134fe565b11156145bb575060018054600160a060020a0319167320e12a1f859b3feae5fb2a0a32c18f5a65555bbf178155611677565b60006145da7351efaf4c8b3c9afbd5ab9f4bbc82784ab6ef8faa6134fe565b111561460c575060018054600160a060020a0319167351efaf4c8b3c9afbd5ab9f4bbc82784ab6ef8faa178155611677565b50600090565b600060018284511461462357600080fd5b60005b838110156146895784818151811061463a57fe5b016020015160f860020a900460f860020a02600160f860020a031916868260208110151561466457fe5b1a60f860020a02600160f860020a03191614151561468157600091505b600101614626565b50949350505050565b600080600080600060606020604051818152601f19601f83011681016020016040529080156146c8576020820181803883390190505b50905060006020896003815181106146dc57fe5b016020015160f860020a900460f860020a0260f860020a900460ff1603600401905061470d89826020856000614022565b915060606020604051818152601f19601f830116810160200160405290801561473d576020820181803883390190505b5090506022820191506147848a60208c600186038151811061475b57fe5b016020015160f860020a900460f860020a0260f860020a900460ff160384016020846000614022565b905060208301519450602081015193506147a18b601b8787614fe3565b9097509550600160a060020a0386168980519060200120600160a060020a031614156147d757600197505050505050505061266a565b6147e48b601c8787614fe3565b9097509550600160a060020a0386168980519060200120600160a060020a031614975061266a9650505050505050565b600080606084846001018151811061482857fe5b016020015160f860020a900460f860020a0260f860020a900460ff16600201604051818152601f19601f8301168101602001604052908015614871576020820181803883390190505b50905061488385858351846000614022565b50606060408051818152601f19601f83011681016020016040529080156148b1576020820181803883390190505b5090506148c48660046040846000614022565b5060606062604051818152601f19601f83011681016020016040529080156148f3576020820181803883390190505b50905060f860020a8160008151811061490857fe5b906020010190600160f860020a031916908160001a90535061493287604188036041846001614022565b5060606040805190810160405280602081526020017ffd94fa71bc0ba10d39d464d0d8f465efeef0a2764e3887fcc9df41ded20f505c815250905061497d8160006020856042614022565b50614a1d6002836040518082805190602001908083835b602083106149b35780518252601f199092019160209182019101614994565b6001836020036101000a038019825116818451168082178552505050505050905001915050602060405180830381855afa1580156149f5573d6000803e3d6000fd5b5050506040513d6020811015614a0a57600080fd5b8101908080519250879150869050614692565b9450841515614a3457600095505050505050610bc4565b606080604051908101604052806040815260200161576e60409139905060606042604051818152601f19601f8301168101602001604052908015614a7f576020820181803883390190505b5090507ffe0000000000000000000000000000000000000000000000000000000000000081600081518110614ab057fe5b906020010190600160f860020a031916908160001a905350614ad88a60036041846001614022565b5060608a604581518110614ae857fe5b016020015160f860020a900460f860020a0260f860020a900460ff16600201604051818152601f19601f8301168101602001604052908015614b31576020820181803883390190505b509050614b448b60448351846000614022565b50614be46002836040518082805190602001908083835b60208310614b7a5780518252601f199092019160209182019101614b5b565b6001836020036101000a038019825116818451168082178552505050505050905001915050602060405180830381855afa158015614bbc573d6000803e3d6000fd5b5050506040513d6020811015614bd157600080fd5b8101908080519250849150869050614692565b9b9a5050505050505050505050565b600154600090600160a060020a03161580614c205750600154614c1e90600160a060020a03166134fe565b155b15614c3157614c2f6000613502565b505b600154600160a060020a03166338cc48316040518163ffffffff1660e060020a028152600401602060405180830381600087803b158015614c7157600080fd5b505af1158015614c85573d6000803e3d6000fd5b505050506040513d6020811015614c9b57600080fd5b810190808051600054600160a060020a039081169116149250614d4f91505057600154600160a060020a03166338cc48316040518163ffffffff1660e060020a028152600401602060405180830381600087803b158015614cfb57600080fd5b505af1158015614d0f573d6000803e3d6000fd5b505050506040513d6020811015614d2557600080fd5b81019080805160008054600160a060020a031916600160a060020a03929092169190911790555050505b60008054600160a060020a0316632ef3accc86856040518363ffffffff1660e060020a0281526004018080602001838152602001828103825284818151815260200191508051906020019080838360005b83811015614db8578082015183820152602001614da0565b50505050905090810190601f168015614de55780820380516001836020036101000a031916815260200191505b509350505050602060405180830381600087803b158015614e0557600080fd5b505af1158015614e19573d6000803e3d6000fd5b505050506040513d6020811015614e2f57600080fd5b8101908080519350505050670de0b6b3a76400003a840201811115614e5857506000905061266a565b6060614e6385615023565b60008054919250600160a060020a039091169063c55c1cb69084908985896040518663ffffffff1660e060020a028152600401808581526020018060200180602001848152602001838103835286818151815260200191508051906020019080838360005b83811015614ee0578082015183820152602001614ec8565b50505050905090810190601f168015614f0d5780820380516001836020036101000a031916815260200191505b50838103825285818151815260200191508051906020019080838360005b83811015614f43578082015183820152602001614f2b565b50505050905090810190601f168015614f705780820380516001836020036101000a031916815260200191505b5096505050505050506020604051808303818588803b158015614f9257600080fd5b505af1158015614fa6573d6000803e3d6000fd5b50505050506040513d6020811015614fbd57600080fd5b8101908080519998505050505050505050565b6002818051610d40929160200190615395565b60008060008060405188815287602082015286604082015285606082015260208160808360006001610bb8f1925080519299929850919650505050505050565b606061502d615097565b615035615413565b615041816104006150a3565b61504a816150d1565b60005b83518110156150845761507c84828151811061506557fe5b90602001906020020151839063ffffffff6150dc16565b60010161504d565b5061508e816150f9565b80519392505050565b60405180590338823950565b8060208106156150b65760208106602003015b60208301819052604051928390526000835290910160405250565b611af3816004615100565b6150e98260028351615119565b61118a828263ffffffff61521716565b611af38160075b610d4082601f602060ff8516021763ffffffff6152b016565b6017811161513a576151358360ff8481166020021683176152b0565b61118a565b60ff81116151735761515b836018602060ff8616021763ffffffff6152b016565b61516d8382600163ffffffff6152e916565b5061118a565b61ffff81116151a757615195836019602060ff8616021763ffffffff6152b016565b61516d8382600263ffffffff6152e916565b63ffffffff81116151dd576151cb83601a602060ff8616021763ffffffff6152b016565b61516d8382600463ffffffff6152e916565b67ffffffffffffffff811161118a5761520583601b602060ff8616021763ffffffff6152b016565b611aef8382600863ffffffff6152e916565b61521f615413565b8260200151835151835101111561524b5761524b836152438560200151855161533d565b600202615354565b60008060008451905085518051602081830101945086510190526020850191505b6020811061528c57815183526020928301929190910190601f190161526c565b60001960208290036101000a01801983511681855116179093525093949350505050565b816020015182515160010111156152d2576152d2828360200151600202615354565b815180516020818301018381535060010190525050565b6152f1615413565b836020015184515183011115615313576153138461524386602001518561533d565b60001961010083900a01845180518481830101868419825116179052909301909252509192915050565b60008183111561534e575081610bc4565b50919050565b60608251905061536483836150a3565b611aef8382615217565b60806040519081016040526004815b606081526020019060019003908161537d5790505090565b828054600181600116156101000203166002900490600052602060002090601f016020900481019282601f106153d657805160ff1916838001178555615403565b82800160010185558215615403579182015b828111156154035782518255916020019190600101906153e8565b5061540f92915061542b565b5090565b60408051908101604052606081526000602082015290565b61167791905b8082111561540f576000815560010161543156fe5b55524c5d206a736f6e2868747470733a2f2f6d696e2d6170692e63727970746f636f6d706172652e636f6d2f646174612f70726963653f6673796d3d425443267473796d733d555344266578747261506172616d733d5072696365526f6c6c267369676e3d74727565266170695f6b65793d247b5b646563727970745d20424a45576f35613533415042724e346659707a35784a61447a506d434c4e6a4b64552b794d654433703656734d4c6b4652466671497652612b64342f71756b54426273465a716b7673744d4d63716f4c5a6153686f683448664839585155784c376341744b777541693847436b467073306b63466d4e4233454151517667474d5834466561616f683430594370357142644b6758714c68582b4256753478397030754b53395858422b43633271496c7661676b4737792b546f3162567270315867673d3d7d292e55534445787065637465642076616c7565206d75737420626520696e207468652072616e6765206f66203220746f2039394f7261636c697a652064656c6179206973206d6178696d756d20363020646179735b55524c5d206a736f6e2868747470733a2f2f6d696e2d6170692e63727970746f636f6d706172652e636f6d2f646174612f70726963653f6673796d3d455448267473796d733d555344266578747261506172616d733d5072696365526f6c6c267369676e3d74727565266170695f6b65793d247b5b646563727970745d20424a45576f35613533415042724e346659707a35784a61447a506d434c4e6a4b64552b794d654433703656734d4c6b4652466671497652612b64342f71756b54426273465a716b7673744d4d63716f4c5a6153686f683448664839585155784c376341744b777541693847436b467073306b63466d4e4233454151517667474d5834466561616f683430594370357142644b6758714c68582b4256753478397030754b53395858422b43633271496c7661676b4737792b546f3162567270315867673d3d7d292e55534451756572792068617320616c7265616479206265656e2070726f636573736564214e6f7420656e6f75676820746f20626574207468652073706563696669656420616d6f756e747fb956469c5c9b89840d55b43537e66a98dd4811ea0a27224272c2e5622911e8537a2f8e86a46baec82864e98dd01e9ccc2f8bc5dfc9cbe5a91a290498dd96e45b55524c5d206a736f6e2868747470733a2f2f6d696e2d6170692e63727970746f636f6d706172652e636f6d2f646174612f70726963653f6673796d3d4c5443267473796d733d555344266578747261506172616d733d5072696365526f6c6c267369676e3d74727565266170695f6b65793d247b5b646563727970745d20424a45576f35613533415042724e346659707a35784a61447a506d434c4e6a4b64552b794d654433703656734d4c6b4652466671497652612b64342f71756b54426273465a716b7673744d4d63716f4c5a6153686f683448664839585155784c376341744b777541693847436b467073306b63466d4e4233454151517667474d5834466561616f683430594370357142644b6758714c68582b4256753478397030754b53395858422b43633271496c7661676b4737792b546f3162567270315867673d3d7d292e555344a165627a7a723058202a6a70cd0d5073016f2a154edf2d62c044fce368cb497cc284231c6c77fdcfbe0029`

// DeployPriceRoll deploys a new Ethereum contract, binding an instance of PriceRoll to it.
func DeployPriceRoll(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *PriceRoll, error) {
	parsed, err := abi.JSON(strings.NewReader(PriceRollABI))
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	address, tx, contract, err := bind.DeployContract(auth, parsed, common.FromHex(PriceRollBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &PriceRoll{PriceRollCaller: PriceRollCaller{contract: contract}, PriceRollTransactor: PriceRollTransactor{contract: contract}, PriceRollFilterer: PriceRollFilterer{contract: contract}}, nil
}

// PriceRoll is an auto generated Go binding around an Ethereum contract.
type PriceRoll struct {
	PriceRollCaller     // Read-only binding to the contract
	PriceRollTransactor // Write-only binding to the contract
	PriceRollFilterer   // Log filterer for contract events
}

// PriceRollCaller is an auto generated read-only Go binding around an Ethereum contract.
type PriceRollCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// PriceRollTransactor is an auto generated write-only Go binding around an Ethereum contract.
type PriceRollTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// PriceRollFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type PriceRollFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// PriceRollSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type PriceRollSession struct {
	Contract     *PriceRoll        // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// PriceRollCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type PriceRollCallerSession struct {
	Contract *PriceRollCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts    // Call options to use throughout this session
}

// PriceRollTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type PriceRollTransactorSession struct {
	Contract     *PriceRollTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts    // Transaction auth options to use throughout this session
}

// PriceRollRaw is an auto generated low-level Go binding around an Ethereum contract.
type PriceRollRaw struct {
	Contract *PriceRoll // Generic contract binding to access the raw methods on
}

// PriceRollCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type PriceRollCallerRaw struct {
	Contract *PriceRollCaller // Generic read-only contract binding to access the raw methods on
}

// PriceRollTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type PriceRollTransactorRaw struct {
	Contract *PriceRollTransactor // Generic write-only contract binding to access the raw methods on
}

// NewPriceRoll creates a new instance of PriceRoll, bound to a specific deployed contract.
func NewPriceRoll(address common.Address, backend bind.ContractBackend) (*PriceRoll, error) {
	contract, err := bindPriceRoll(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &PriceRoll{PriceRollCaller: PriceRollCaller{contract: contract}, PriceRollTransactor: PriceRollTransactor{contract: contract}, PriceRollFilterer: PriceRollFilterer{contract: contract}}, nil
}

// NewPriceRollCaller creates a new read-only instance of PriceRoll, bound to a specific deployed contract.
func NewPriceRollCaller(address common.Address, caller bind.ContractCaller) (*PriceRollCaller, error) {
	contract, err := bindPriceRoll(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &PriceRollCaller{contract: contract}, nil
}

// NewPriceRollTransactor creates a new write-only instance of PriceRoll, bound to a specific deployed contract.
func NewPriceRollTransactor(address common.Address, transactor bind.ContractTransactor) (*PriceRollTransactor, error) {
	contract, err := bindPriceRoll(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &PriceRollTransactor{contract: contract}, nil
}

// NewPriceRollFilterer creates a new log filterer instance of PriceRoll, bound to a specific deployed contract.
func NewPriceRollFilterer(address common.Address, filterer bind.ContractFilterer) (*PriceRollFilterer, error) {
	contract, err := bindPriceRoll(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &PriceRollFilterer{contract: contract}, nil
}

// bindPriceRoll binds a generic wrapper to an already deployed contract.
func bindPriceRoll(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(PriceRollABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_PriceRoll *PriceRollRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _PriceRoll.Contract.PriceRollCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_PriceRoll *PriceRollRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _PriceRoll.Contract.PriceRollTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_PriceRoll *PriceRollRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _PriceRoll.Contract.PriceRollTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_PriceRoll *PriceRollCallerRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _PriceRoll.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_PriceRoll *PriceRollTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _PriceRoll.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_PriceRoll *PriceRollTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _PriceRoll.Contract.contract.Transact(opts, method, params...)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf( address) constant returns(uint256)
func (_PriceRoll *PriceRollCaller) BalanceOf(opts *bind.CallOpts, arg0 common.Address) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _PriceRoll.contract.Call(opts, out, "balanceOf", arg0)
	return *ret0, err
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf( address) constant returns(uint256)
func (_PriceRoll *PriceRollSession) BalanceOf(arg0 common.Address) (*big.Int, error) {
	return _PriceRoll.Contract.BalanceOf(&_PriceRoll.CallOpts, arg0)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf( address) constant returns(uint256)
func (_PriceRoll *PriceRollCallerSession) BalanceOf(arg0 common.Address) (*big.Int, error) {
	return _PriceRoll.Contract.BalanceOf(&_PriceRoll.CallOpts, arg0)
}

// ConfigBonusMult is a free data retrieval call binding the contract method 0x40268412.
//
// Solidity: function config_bonus_mult() constant returns(uint256)
func (_PriceRoll *PriceRollCaller) ConfigBonusMult(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _PriceRoll.contract.Call(opts, out, "config_bonus_mult")
	return *ret0, err
}

// ConfigBonusMult is a free data retrieval call binding the contract method 0x40268412.
//
// Solidity: function config_bonus_mult() constant returns(uint256)
func (_PriceRoll *PriceRollSession) ConfigBonusMult() (*big.Int, error) {
	return _PriceRoll.Contract.ConfigBonusMult(&_PriceRoll.CallOpts)
}

// ConfigBonusMult is a free data retrieval call binding the contract method 0x40268412.
//
// Solidity: function config_bonus_mult() constant returns(uint256)
func (_PriceRoll *PriceRollCallerSession) ConfigBonusMult() (*big.Int, error) {
	return _PriceRoll.Contract.ConfigBonusMult(&_PriceRoll.CallOpts)
}

// ConfigGasLimit is a free data retrieval call binding the contract method 0xd3164013.
//
// Solidity: function config_gas_limit() constant returns(uint256)
func (_PriceRoll *PriceRollCaller) ConfigGasLimit(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _PriceRoll.contract.Call(opts, out, "config_gas_limit")
	return *ret0, err
}

// ConfigGasLimit is a free data retrieval call binding the contract method 0xd3164013.
//
// Solidity: function config_gas_limit() constant returns(uint256)
func (_PriceRoll *PriceRollSession) ConfigGasLimit() (*big.Int, error) {
	return _PriceRoll.Contract.ConfigGasLimit(&_PriceRoll.CallOpts)
}

// ConfigGasLimit is a free data retrieval call binding the contract method 0xd3164013.
//
// Solidity: function config_gas_limit() constant returns(uint256)
func (_PriceRoll *PriceRollCallerSession) ConfigGasLimit() (*big.Int, error) {
	return _PriceRoll.Contract.ConfigGasLimit(&_PriceRoll.CallOpts)
}

// ConfigHouseEdge is a free data retrieval call binding the contract method 0x5d10c029.
//
// Solidity: function config_house_edge() constant returns(uint256)
func (_PriceRoll *PriceRollCaller) ConfigHouseEdge(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _PriceRoll.contract.Call(opts, out, "config_house_edge")
	return *ret0, err
}

// ConfigHouseEdge is a free data retrieval call binding the contract method 0x5d10c029.
//
// Solidity: function config_house_edge() constant returns(uint256)
func (_PriceRoll *PriceRollSession) ConfigHouseEdge() (*big.Int, error) {
	return _PriceRoll.Contract.ConfigHouseEdge(&_PriceRoll.CallOpts)
}

// ConfigHouseEdge is a free data retrieval call binding the contract method 0x5d10c029.
//
// Solidity: function config_house_edge() constant returns(uint256)
func (_PriceRoll *PriceRollCallerSession) ConfigHouseEdge() (*big.Int, error) {
	return _PriceRoll.Contract.ConfigHouseEdge(&_PriceRoll.CallOpts)
}

// ConfigMaxBet is a free data retrieval call binding the contract method 0x8b1dc09e.
//
// Solidity: function config_max_bet() constant returns(uint256)
func (_PriceRoll *PriceRollCaller) ConfigMaxBet(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _PriceRoll.contract.Call(opts, out, "config_max_bet")
	return *ret0, err
}

// ConfigMaxBet is a free data retrieval call binding the contract method 0x8b1dc09e.
//
// Solidity: function config_max_bet() constant returns(uint256)
func (_PriceRoll *PriceRollSession) ConfigMaxBet() (*big.Int, error) {
	return _PriceRoll.Contract.ConfigMaxBet(&_PriceRoll.CallOpts)
}

// ConfigMaxBet is a free data retrieval call binding the contract method 0x8b1dc09e.
//
// Solidity: function config_max_bet() constant returns(uint256)
func (_PriceRoll *PriceRollCallerSession) ConfigMaxBet() (*big.Int, error) {
	return _PriceRoll.Contract.ConfigMaxBet(&_PriceRoll.CallOpts)
}

// ConfigMinBet is a free data retrieval call binding the contract method 0xdeedc5c5.
//
// Solidity: function config_min_bet() constant returns(uint256)
func (_PriceRoll *PriceRollCaller) ConfigMinBet(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _PriceRoll.contract.Call(opts, out, "config_min_bet")
	return *ret0, err
}

// ConfigMinBet is a free data retrieval call binding the contract method 0xdeedc5c5.
//
// Solidity: function config_min_bet() constant returns(uint256)
func (_PriceRoll *PriceRollSession) ConfigMinBet() (*big.Int, error) {
	return _PriceRoll.Contract.ConfigMinBet(&_PriceRoll.CallOpts)
}

// ConfigMinBet is a free data retrieval call binding the contract method 0xdeedc5c5.
//
// Solidity: function config_min_bet() constant returns(uint256)
func (_PriceRoll *PriceRollCallerSession) ConfigMinBet() (*big.Int, error) {
	return _PriceRoll.Contract.ConfigMinBet(&_PriceRoll.CallOpts)
}

// ConfigPricecheckDelay is a free data retrieval call binding the contract method 0x5ad51f68.
//
// Solidity: function config_pricecheck_delay() constant returns(uint256)
func (_PriceRoll *PriceRollCaller) ConfigPricecheckDelay(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _PriceRoll.contract.Call(opts, out, "config_pricecheck_delay")
	return *ret0, err
}

// ConfigPricecheckDelay is a free data retrieval call binding the contract method 0x5ad51f68.
//
// Solidity: function config_pricecheck_delay() constant returns(uint256)
func (_PriceRoll *PriceRollSession) ConfigPricecheckDelay() (*big.Int, error) {
	return _PriceRoll.Contract.ConfigPricecheckDelay(&_PriceRoll.CallOpts)
}

// ConfigPricecheckDelay is a free data retrieval call binding the contract method 0x5ad51f68.
//
// Solidity: function config_pricecheck_delay() constant returns(uint256)
func (_PriceRoll *PriceRollCallerSession) ConfigPricecheckDelay() (*big.Int, error) {
	return _PriceRoll.Contract.ConfigPricecheckDelay(&_PriceRoll.CallOpts)
}

// ConfigRandomGasLimit is a free data retrieval call binding the contract method 0x57b6702f.
//
// Solidity: function config_random_gas_limit() constant returns(uint256)
func (_PriceRoll *PriceRollCaller) ConfigRandomGasLimit(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _PriceRoll.contract.Call(opts, out, "config_random_gas_limit")
	return *ret0, err
}

// ConfigRandomGasLimit is a free data retrieval call binding the contract method 0x57b6702f.
//
// Solidity: function config_random_gas_limit() constant returns(uint256)
func (_PriceRoll *PriceRollSession) ConfigRandomGasLimit() (*big.Int, error) {
	return _PriceRoll.Contract.ConfigRandomGasLimit(&_PriceRoll.CallOpts)
}

// ConfigRandomGasLimit is a free data retrieval call binding the contract method 0x57b6702f.
//
// Solidity: function config_random_gas_limit() constant returns(uint256)
func (_PriceRoll *PriceRollCallerSession) ConfigRandomGasLimit() (*big.Int, error) {
	return _PriceRoll.Contract.ConfigRandomGasLimit(&_PriceRoll.CallOpts)
}

// ConfigRefundDelay is a free data retrieval call binding the contract method 0x55009899.
//
// Solidity: function config_refund_delay() constant returns(uint256)
func (_PriceRoll *PriceRollCaller) ConfigRefundDelay(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _PriceRoll.contract.Call(opts, out, "config_refund_delay")
	return *ret0, err
}

// ConfigRefundDelay is a free data retrieval call binding the contract method 0x55009899.
//
// Solidity: function config_refund_delay() constant returns(uint256)
func (_PriceRoll *PriceRollSession) ConfigRefundDelay() (*big.Int, error) {
	return _PriceRoll.Contract.ConfigRefundDelay(&_PriceRoll.CallOpts)
}

// ConfigRefundDelay is a free data retrieval call binding the contract method 0x55009899.
//
// Solidity: function config_refund_delay() constant returns(uint256)
func (_PriceRoll *PriceRollCallerSession) ConfigRefundDelay() (*big.Int, error) {
	return _PriceRoll.Contract.ConfigRefundDelay(&_PriceRoll.CallOpts)
}

// ConfigRollCooldown is a free data retrieval call binding the contract method 0x987ae154.
//
// Solidity: function config_roll_cooldown() constant returns(uint256)
func (_PriceRoll *PriceRollCaller) ConfigRollCooldown(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _PriceRoll.contract.Call(opts, out, "config_roll_cooldown")
	return *ret0, err
}

// ConfigRollCooldown is a free data retrieval call binding the contract method 0x987ae154.
//
// Solidity: function config_roll_cooldown() constant returns(uint256)
func (_PriceRoll *PriceRollSession) ConfigRollCooldown() (*big.Int, error) {
	return _PriceRoll.Contract.ConfigRollCooldown(&_PriceRoll.CallOpts)
}

// ConfigRollCooldown is a free data retrieval call binding the contract method 0x987ae154.
//
// Solidity: function config_roll_cooldown() constant returns(uint256)
func (_PriceRoll *PriceRollCallerSession) ConfigRollCooldown() (*big.Int, error) {
	return _PriceRoll.Contract.ConfigRollCooldown(&_PriceRoll.CallOpts)
}

// CurrentCoin is a free data retrieval call binding the contract method 0x06792c07.
//
// Solidity: function current_coin() constant returns(uint8)
func (_PriceRoll *PriceRollCaller) CurrentCoin(opts *bind.CallOpts) (uint8, error) {
	var (
		ret0 = new(uint8)
	)
	out := ret0
	err := _PriceRoll.contract.Call(opts, out, "current_coin")
	return *ret0, err
}

// CurrentCoin is a free data retrieval call binding the contract method 0x06792c07.
//
// Solidity: function current_coin() constant returns(uint8)
func (_PriceRoll *PriceRollSession) CurrentCoin() (uint8, error) {
	return _PriceRoll.Contract.CurrentCoin(&_PriceRoll.CallOpts)
}

// CurrentCoin is a free data retrieval call binding the contract method 0x06792c07.
//
// Solidity: function current_coin() constant returns(uint8)
func (_PriceRoll *PriceRollCallerSession) CurrentCoin() (uint8, error) {
	return _PriceRoll.Contract.CurrentCoin(&_PriceRoll.CallOpts)
}

// CurrentRoll is a free data retrieval call binding the contract method 0x4e0c12f0.
//
// Solidity: function current_roll() constant returns(uint256)
func (_PriceRoll *PriceRollCaller) CurrentRoll(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _PriceRoll.contract.Call(opts, out, "current_roll")
	return *ret0, err
}

// CurrentRoll is a free data retrieval call binding the contract method 0x4e0c12f0.
//
// Solidity: function current_roll() constant returns(uint256)
func (_PriceRoll *PriceRollSession) CurrentRoll() (*big.Int, error) {
	return _PriceRoll.Contract.CurrentRoll(&_PriceRoll.CallOpts)
}

// CurrentRoll is a free data retrieval call binding the contract method 0x4e0c12f0.
//
// Solidity: function current_roll() constant returns(uint256)
func (_PriceRoll *PriceRollCallerSession) CurrentRoll() (*big.Int, error) {
	return _PriceRoll.Contract.CurrentRoll(&_PriceRoll.CallOpts)
}

// House is a free data retrieval call binding the contract method 0xff9b3acf.
//
// Solidity: function house() constant returns(uint256)
func (_PriceRoll *PriceRollCaller) House(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _PriceRoll.contract.Call(opts, out, "house")
	return *ret0, err
}

// House is a free data retrieval call binding the contract method 0xff9b3acf.
//
// Solidity: function house() constant returns(uint256)
func (_PriceRoll *PriceRollSession) House() (*big.Int, error) {
	return _PriceRoll.Contract.House(&_PriceRoll.CallOpts)
}

// House is a free data retrieval call binding the contract method 0xff9b3acf.
//
// Solidity: function house() constant returns(uint256)
func (_PriceRoll *PriceRollCallerSession) House() (*big.Int, error) {
	return _PriceRoll.Contract.House(&_PriceRoll.CallOpts)
}

// IsOwner is a free data retrieval call binding the contract method 0x8f32d59b.
//
// Solidity: function isOwner() constant returns(bool)
func (_PriceRoll *PriceRollCaller) IsOwner(opts *bind.CallOpts) (bool, error) {
	var (
		ret0 = new(bool)
	)
	out := ret0
	err := _PriceRoll.contract.Call(opts, out, "isOwner")
	return *ret0, err
}

// IsOwner is a free data retrieval call binding the contract method 0x8f32d59b.
//
// Solidity: function isOwner() constant returns(bool)
func (_PriceRoll *PriceRollSession) IsOwner() (bool, error) {
	return _PriceRoll.Contract.IsOwner(&_PriceRoll.CallOpts)
}

// IsOwner is a free data retrieval call binding the contract method 0x8f32d59b.
//
// Solidity: function isOwner() constant returns(bool)
func (_PriceRoll *PriceRollCallerSession) IsOwner() (bool, error) {
	return _PriceRoll.Contract.IsOwner(&_PriceRoll.CallOpts)
}

// IsPauser is a free data retrieval call binding the contract method 0x46fbf68e.
//
// Solidity: function isPauser(account address) constant returns(bool)
func (_PriceRoll *PriceRollCaller) IsPauser(opts *bind.CallOpts, account common.Address) (bool, error) {
	var (
		ret0 = new(bool)
	)
	out := ret0
	err := _PriceRoll.contract.Call(opts, out, "isPauser", account)
	return *ret0, err
}

// IsPauser is a free data retrieval call binding the contract method 0x46fbf68e.
//
// Solidity: function isPauser(account address) constant returns(bool)
func (_PriceRoll *PriceRollSession) IsPauser(account common.Address) (bool, error) {
	return _PriceRoll.Contract.IsPauser(&_PriceRoll.CallOpts, account)
}

// IsPauser is a free data retrieval call binding the contract method 0x46fbf68e.
//
// Solidity: function isPauser(account address) constant returns(bool)
func (_PriceRoll *PriceRollCallerSession) IsPauser(account common.Address) (bool, error) {
	return _PriceRoll.Contract.IsPauser(&_PriceRoll.CallOpts, account)
}

// LatestRoll is a free data retrieval call binding the contract method 0x022b02bc.
//
// Solidity: function latest_roll() constant returns(uint256)
func (_PriceRoll *PriceRollCaller) LatestRoll(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _PriceRoll.contract.Call(opts, out, "latest_roll")
	return *ret0, err
}

// LatestRoll is a free data retrieval call binding the contract method 0x022b02bc.
//
// Solidity: function latest_roll() constant returns(uint256)
func (_PriceRoll *PriceRollSession) LatestRoll() (*big.Int, error) {
	return _PriceRoll.Contract.LatestRoll(&_PriceRoll.CallOpts)
}

// LatestRoll is a free data retrieval call binding the contract method 0x022b02bc.
//
// Solidity: function latest_roll() constant returns(uint256)
func (_PriceRoll *PriceRollCallerSession) LatestRoll() (*big.Int, error) {
	return _PriceRoll.Contract.LatestRoll(&_PriceRoll.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() constant returns(address)
func (_PriceRoll *PriceRollCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _PriceRoll.contract.Call(opts, out, "owner")
	return *ret0, err
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() constant returns(address)
func (_PriceRoll *PriceRollSession) Owner() (common.Address, error) {
	return _PriceRoll.Contract.Owner(&_PriceRoll.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() constant returns(address)
func (_PriceRoll *PriceRollCallerSession) Owner() (common.Address, error) {
	return _PriceRoll.Contract.Owner(&_PriceRoll.CallOpts)
}

// Paused is a free data retrieval call binding the contract method 0x5c975abb.
//
// Solidity: function paused() constant returns(bool)
func (_PriceRoll *PriceRollCaller) Paused(opts *bind.CallOpts) (bool, error) {
	var (
		ret0 = new(bool)
	)
	out := ret0
	err := _PriceRoll.contract.Call(opts, out, "paused")
	return *ret0, err
}

// Paused is a free data retrieval call binding the contract method 0x5c975abb.
//
// Solidity: function paused() constant returns(bool)
func (_PriceRoll *PriceRollSession) Paused() (bool, error) {
	return _PriceRoll.Contract.Paused(&_PriceRoll.CallOpts)
}

// Paused is a free data retrieval call binding the contract method 0x5c975abb.
//
// Solidity: function paused() constant returns(bool)
func (_PriceRoll *PriceRollCallerSession) Paused() (bool, error) {
	return _PriceRoll.Contract.Paused(&_PriceRoll.CallOpts)
}

// Pool is a free data retrieval call binding the contract method 0x16f0115b.
//
// Solidity: function pool() constant returns(uint256)
func (_PriceRoll *PriceRollCaller) Pool(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _PriceRoll.contract.Call(opts, out, "pool")
	return *ret0, err
}

// Pool is a free data retrieval call binding the contract method 0x16f0115b.
//
// Solidity: function pool() constant returns(uint256)
func (_PriceRoll *PriceRollSession) Pool() (*big.Int, error) {
	return _PriceRoll.Contract.Pool(&_PriceRoll.CallOpts)
}

// Pool is a free data retrieval call binding the contract method 0x16f0115b.
//
// Solidity: function pool() constant returns(uint256)
func (_PriceRoll *PriceRollCallerSession) Pool() (*big.Int, error) {
	return _PriceRoll.Contract.Pool(&_PriceRoll.CallOpts)
}

// QueryStringBTC is a free data retrieval call binding the contract method 0xd8a3f00f.
//
// Solidity: function query_stringBTC() constant returns(string)
func (_PriceRoll *PriceRollCaller) QueryStringBTC(opts *bind.CallOpts) (string, error) {
	var (
		ret0 = new(string)
	)
	out := ret0
	err := _PriceRoll.contract.Call(opts, out, "query_stringBTC")
	return *ret0, err
}

// QueryStringBTC is a free data retrieval call binding the contract method 0xd8a3f00f.
//
// Solidity: function query_stringBTC() constant returns(string)
func (_PriceRoll *PriceRollSession) QueryStringBTC() (string, error) {
	return _PriceRoll.Contract.QueryStringBTC(&_PriceRoll.CallOpts)
}

// QueryStringBTC is a free data retrieval call binding the contract method 0xd8a3f00f.
//
// Solidity: function query_stringBTC() constant returns(string)
func (_PriceRoll *PriceRollCallerSession) QueryStringBTC() (string, error) {
	return _PriceRoll.Contract.QueryStringBTC(&_PriceRoll.CallOpts)
}

// QueryStringETH is a free data retrieval call binding the contract method 0x09108570.
//
// Solidity: function query_stringETH() constant returns(string)
func (_PriceRoll *PriceRollCaller) QueryStringETH(opts *bind.CallOpts) (string, error) {
	var (
		ret0 = new(string)
	)
	out := ret0
	err := _PriceRoll.contract.Call(opts, out, "query_stringETH")
	return *ret0, err
}

// QueryStringETH is a free data retrieval call binding the contract method 0x09108570.
//
// Solidity: function query_stringETH() constant returns(string)
func (_PriceRoll *PriceRollSession) QueryStringETH() (string, error) {
	return _PriceRoll.Contract.QueryStringETH(&_PriceRoll.CallOpts)
}

// QueryStringETH is a free data retrieval call binding the contract method 0x09108570.
//
// Solidity: function query_stringETH() constant returns(string)
func (_PriceRoll *PriceRollCallerSession) QueryStringETH() (string, error) {
	return _PriceRoll.Contract.QueryStringETH(&_PriceRoll.CallOpts)
}

// QueryStringLTC is a free data retrieval call binding the contract method 0x90e2c9dd.
//
// Solidity: function query_stringLTC() constant returns(string)
func (_PriceRoll *PriceRollCaller) QueryStringLTC(opts *bind.CallOpts) (string, error) {
	var (
		ret0 = new(string)
	)
	out := ret0
	err := _PriceRoll.contract.Call(opts, out, "query_stringLTC")
	return *ret0, err
}

// QueryStringLTC is a free data retrieval call binding the contract method 0x90e2c9dd.
//
// Solidity: function query_stringLTC() constant returns(string)
func (_PriceRoll *PriceRollSession) QueryStringLTC() (string, error) {
	return _PriceRoll.Contract.QueryStringLTC(&_PriceRoll.CallOpts)
}

// QueryStringLTC is a free data retrieval call binding the contract method 0x90e2c9dd.
//
// Solidity: function query_stringLTC() constant returns(string)
func (_PriceRoll *PriceRollCallerSession) QueryStringLTC() (string, error) {
	return _PriceRoll.Contract.QueryStringLTC(&_PriceRoll.CallOpts)
}

// Rolls is a free data retrieval call binding the contract method 0x5d69f16c.
//
// Solidity: function rolls( uint256) constant returns(query_rng bytes32, query_price1 bytes32, query_price2 bytes32, result_price1 uint256, result_price2 uint256, timestamp uint256, pool uint256, state uint8, coin uint8, result_rng uint8, is_up bool)
func (_PriceRoll *PriceRollCaller) Rolls(opts *bind.CallOpts, arg0 *big.Int) (struct {
	QueryRng     [32]byte
	QueryPrice1  [32]byte
	QueryPrice2  [32]byte
	ResultPrice1 *big.Int
	ResultPrice2 *big.Int
	Timestamp    *big.Int
	Pool         *big.Int
	State        uint8
	Coin         uint8
	ResultRng    uint8
	IsUp         bool
}, error) {
	ret := new(struct {
		QueryRng     [32]byte
		QueryPrice1  [32]byte
		QueryPrice2  [32]byte
		ResultPrice1 *big.Int
		ResultPrice2 *big.Int
		Timestamp    *big.Int
		Pool         *big.Int
		State        uint8
		Coin         uint8
		ResultRng    uint8
		IsUp         bool
	})
	out := ret
	err := _PriceRoll.contract.Call(opts, out, "rolls", arg0)
	return *ret, err
}

// Rolls is a free data retrieval call binding the contract method 0x5d69f16c.
//
// Solidity: function rolls( uint256) constant returns(query_rng bytes32, query_price1 bytes32, query_price2 bytes32, result_price1 uint256, result_price2 uint256, timestamp uint256, pool uint256, state uint8, coin uint8, result_rng uint8, is_up bool)
func (_PriceRoll *PriceRollSession) Rolls(arg0 *big.Int) (struct {
	QueryRng     [32]byte
	QueryPrice1  [32]byte
	QueryPrice2  [32]byte
	ResultPrice1 *big.Int
	ResultPrice2 *big.Int
	Timestamp    *big.Int
	Pool         *big.Int
	State        uint8
	Coin         uint8
	ResultRng    uint8
	IsUp         bool
}, error) {
	return _PriceRoll.Contract.Rolls(&_PriceRoll.CallOpts, arg0)
}

// Rolls is a free data retrieval call binding the contract method 0x5d69f16c.
//
// Solidity: function rolls( uint256) constant returns(query_rng bytes32, query_price1 bytes32, query_price2 bytes32, result_price1 uint256, result_price2 uint256, timestamp uint256, pool uint256, state uint8, coin uint8, result_rng uint8, is_up bool)
func (_PriceRoll *PriceRollCallerSession) Rolls(arg0 *big.Int) (struct {
	QueryRng     [32]byte
	QueryPrice1  [32]byte
	QueryPrice2  [32]byte
	ResultPrice1 *big.Int
	ResultPrice2 *big.Int
	Timestamp    *big.Int
	Pool         *big.Int
	State        uint8
	Coin         uint8
	ResultRng    uint8
	IsUp         bool
}, error) {
	return _PriceRoll.Contract.Rolls(&_PriceRoll.CallOpts, arg0)
}

// UserBalance is a free data retrieval call binding the contract method 0xbf152765.
//
// Solidity: function userBalance() constant returns(uint256)
func (_PriceRoll *PriceRollCaller) UserBalance(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _PriceRoll.contract.Call(opts, out, "userBalance")
	return *ret0, err
}

// UserBalance is a free data retrieval call binding the contract method 0xbf152765.
//
// Solidity: function userBalance() constant returns(uint256)
func (_PriceRoll *PriceRollSession) UserBalance() (*big.Int, error) {
	return _PriceRoll.Contract.UserBalance(&_PriceRoll.CallOpts)
}

// UserBalance is a free data retrieval call binding the contract method 0xbf152765.
//
// Solidity: function userBalance() constant returns(uint256)
func (_PriceRoll *PriceRollCallerSession) UserBalance() (*big.Int, error) {
	return _PriceRoll.Contract.UserBalance(&_PriceRoll.CallOpts)
}

// Callback is a paid mutator transaction binding the contract method 0x38bbfa50.
//
// Solidity: function __callback(_queryId bytes32, _result string, _proof bytes) returns()
func (_PriceRoll *PriceRollTransactor) Callback(opts *bind.TransactOpts, _queryId [32]byte, _result string, _proof []byte) (*types.Transaction, error) {
	return _PriceRoll.contract.Transact(opts, "__callback", _queryId, _result, _proof)
}

// Callback is a paid mutator transaction binding the contract method 0x38bbfa50.
//
// Solidity: function __callback(_queryId bytes32, _result string, _proof bytes) returns()
func (_PriceRoll *PriceRollSession) Callback(_queryId [32]byte, _result string, _proof []byte) (*types.Transaction, error) {
	return _PriceRoll.Contract.Callback(&_PriceRoll.TransactOpts, _queryId, _result, _proof)
}

// Callback is a paid mutator transaction binding the contract method 0x38bbfa50.
//
// Solidity: function __callback(_queryId bytes32, _result string, _proof bytes) returns()
func (_PriceRoll *PriceRollTransactorSession) Callback(_queryId [32]byte, _result string, _proof []byte) (*types.Transaction, error) {
	return _PriceRoll.Contract.Callback(&_PriceRoll.TransactOpts, _queryId, _result, _proof)
}

// AddPauser is a paid mutator transaction binding the contract method 0x82dc1ec4.
//
// Solidity: function addPauser(account address) returns()
func (_PriceRoll *PriceRollTransactor) AddPauser(opts *bind.TransactOpts, account common.Address) (*types.Transaction, error) {
	return _PriceRoll.contract.Transact(opts, "addPauser", account)
}

// AddPauser is a paid mutator transaction binding the contract method 0x82dc1ec4.
//
// Solidity: function addPauser(account address) returns()
func (_PriceRoll *PriceRollSession) AddPauser(account common.Address) (*types.Transaction, error) {
	return _PriceRoll.Contract.AddPauser(&_PriceRoll.TransactOpts, account)
}

// AddPauser is a paid mutator transaction binding the contract method 0x82dc1ec4.
//
// Solidity: function addPauser(account address) returns()
func (_PriceRoll *PriceRollTransactorSession) AddPauser(account common.Address) (*types.Transaction, error) {
	return _PriceRoll.Contract.AddPauser(&_PriceRoll.TransactOpts, account)
}

// BetFromInternalWallet is a paid mutator transaction binding the contract method 0x7de36c97.
//
// Solidity: function betFromInternalWallet(amount uint256, expected_value uint8, is_up bool) returns()
func (_PriceRoll *PriceRollTransactor) BetFromInternalWallet(opts *bind.TransactOpts, amount *big.Int, expected_value uint8, is_up bool) (*types.Transaction, error) {
	return _PriceRoll.contract.Transact(opts, "betFromInternalWallet", amount, expected_value, is_up)
}

// BetFromInternalWallet is a paid mutator transaction binding the contract method 0x7de36c97.
//
// Solidity: function betFromInternalWallet(amount uint256, expected_value uint8, is_up bool) returns()
func (_PriceRoll *PriceRollSession) BetFromInternalWallet(amount *big.Int, expected_value uint8, is_up bool) (*types.Transaction, error) {
	return _PriceRoll.Contract.BetFromInternalWallet(&_PriceRoll.TransactOpts, amount, expected_value, is_up)
}

// BetFromInternalWallet is a paid mutator transaction binding the contract method 0x7de36c97.
//
// Solidity: function betFromInternalWallet(amount uint256, expected_value uint8, is_up bool) returns()
func (_PriceRoll *PriceRollTransactorSession) BetFromInternalWallet(amount *big.Int, expected_value uint8, is_up bool) (*types.Transaction, error) {
	return _PriceRoll.Contract.BetFromInternalWallet(&_PriceRoll.TransactOpts, amount, expected_value, is_up)
}

// Claim is a paid mutator transaction binding the contract method 0x379607f5.
//
// Solidity: function claim(round uint256) returns()
func (_PriceRoll *PriceRollTransactor) Claim(opts *bind.TransactOpts, round *big.Int) (*types.Transaction, error) {
	return _PriceRoll.contract.Transact(opts, "claim", round)
}

// Claim is a paid mutator transaction binding the contract method 0x379607f5.
//
// Solidity: function claim(round uint256) returns()
func (_PriceRoll *PriceRollSession) Claim(round *big.Int) (*types.Transaction, error) {
	return _PriceRoll.Contract.Claim(&_PriceRoll.TransactOpts, round)
}

// Claim is a paid mutator transaction binding the contract method 0x379607f5.
//
// Solidity: function claim(round uint256) returns()
func (_PriceRoll *PriceRollTransactorSession) Claim(round *big.Int) (*types.Transaction, error) {
	return _PriceRoll.Contract.Claim(&_PriceRoll.TransactOpts, round)
}

// CreditWallet is a paid mutator transaction binding the contract method 0xb8b069fc.
//
// Solidity: function creditWallet() returns()
func (_PriceRoll *PriceRollTransactor) CreditWallet(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _PriceRoll.contract.Transact(opts, "creditWallet")
}

// CreditWallet is a paid mutator transaction binding the contract method 0xb8b069fc.
//
// Solidity: function creditWallet() returns()
func (_PriceRoll *PriceRollSession) CreditWallet() (*types.Transaction, error) {
	return _PriceRoll.Contract.CreditWallet(&_PriceRoll.TransactOpts)
}

// CreditWallet is a paid mutator transaction binding the contract method 0xb8b069fc.
//
// Solidity: function creditWallet() returns()
func (_PriceRoll *PriceRollTransactorSession) CreditWallet() (*types.Transaction, error) {
	return _PriceRoll.Contract.CreditWallet(&_PriceRoll.TransactOpts)
}

// Destroy is a paid mutator transaction binding the contract method 0x83197ef0.
//
// Solidity: function destroy() returns()
func (_PriceRoll *PriceRollTransactor) Destroy(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _PriceRoll.contract.Transact(opts, "destroy")
}

// Destroy is a paid mutator transaction binding the contract method 0x83197ef0.
//
// Solidity: function destroy() returns()
func (_PriceRoll *PriceRollSession) Destroy() (*types.Transaction, error) {
	return _PriceRoll.Contract.Destroy(&_PriceRoll.TransactOpts)
}

// Destroy is a paid mutator transaction binding the contract method 0x83197ef0.
//
// Solidity: function destroy() returns()
func (_PriceRoll *PriceRollTransactorSession) Destroy() (*types.Transaction, error) {
	return _PriceRoll.Contract.Destroy(&_PriceRoll.TransactOpts)
}

// NewRoll is a paid mutator transaction binding the contract method 0x6fabd265.
//
// Solidity: function newRoll() returns()
func (_PriceRoll *PriceRollTransactor) NewRoll(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _PriceRoll.contract.Transact(opts, "newRoll")
}

// NewRoll is a paid mutator transaction binding the contract method 0x6fabd265.
//
// Solidity: function newRoll() returns()
func (_PriceRoll *PriceRollSession) NewRoll() (*types.Transaction, error) {
	return _PriceRoll.Contract.NewRoll(&_PriceRoll.TransactOpts)
}

// NewRoll is a paid mutator transaction binding the contract method 0x6fabd265.
//
// Solidity: function newRoll() returns()
func (_PriceRoll *PriceRollTransactorSession) NewRoll() (*types.Transaction, error) {
	return _PriceRoll.Contract.NewRoll(&_PriceRoll.TransactOpts)
}

// Pause is a paid mutator transaction binding the contract method 0x8456cb59.
//
// Solidity: function pause() returns()
func (_PriceRoll *PriceRollTransactor) Pause(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _PriceRoll.contract.Transact(opts, "pause")
}

// Pause is a paid mutator transaction binding the contract method 0x8456cb59.
//
// Solidity: function pause() returns()
func (_PriceRoll *PriceRollSession) Pause() (*types.Transaction, error) {
	return _PriceRoll.Contract.Pause(&_PriceRoll.TransactOpts)
}

// Pause is a paid mutator transaction binding the contract method 0x8456cb59.
//
// Solidity: function pause() returns()
func (_PriceRoll *PriceRollTransactorSession) Pause() (*types.Transaction, error) {
	return _PriceRoll.Contract.Pause(&_PriceRoll.TransactOpts)
}

// PlaceBet is a paid mutator transaction binding the contract method 0x2ed1317c.
//
// Solidity: function placeBet(expected_value uint8, is_up bool) returns()
func (_PriceRoll *PriceRollTransactor) PlaceBet(opts *bind.TransactOpts, expected_value uint8, is_up bool) (*types.Transaction, error) {
	return _PriceRoll.contract.Transact(opts, "placeBet", expected_value, is_up)
}

// PlaceBet is a paid mutator transaction binding the contract method 0x2ed1317c.
//
// Solidity: function placeBet(expected_value uint8, is_up bool) returns()
func (_PriceRoll *PriceRollSession) PlaceBet(expected_value uint8, is_up bool) (*types.Transaction, error) {
	return _PriceRoll.Contract.PlaceBet(&_PriceRoll.TransactOpts, expected_value, is_up)
}

// PlaceBet is a paid mutator transaction binding the contract method 0x2ed1317c.
//
// Solidity: function placeBet(expected_value uint8, is_up bool) returns()
func (_PriceRoll *PriceRollTransactorSession) PlaceBet(expected_value uint8, is_up bool) (*types.Transaction, error) {
	return _PriceRoll.Contract.PlaceBet(&_PriceRoll.TransactOpts, expected_value, is_up)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_PriceRoll *PriceRollTransactor) RenounceOwnership(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _PriceRoll.contract.Transact(opts, "renounceOwnership")
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_PriceRoll *PriceRollSession) RenounceOwnership() (*types.Transaction, error) {
	return _PriceRoll.Contract.RenounceOwnership(&_PriceRoll.TransactOpts)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_PriceRoll *PriceRollTransactorSession) RenounceOwnership() (*types.Transaction, error) {
	return _PriceRoll.Contract.RenounceOwnership(&_PriceRoll.TransactOpts)
}

// RenouncePauser is a paid mutator transaction binding the contract method 0x6ef8d66d.
//
// Solidity: function renouncePauser() returns()
func (_PriceRoll *PriceRollTransactor) RenouncePauser(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _PriceRoll.contract.Transact(opts, "renouncePauser")
}

// RenouncePauser is a paid mutator transaction binding the contract method 0x6ef8d66d.
//
// Solidity: function renouncePauser() returns()
func (_PriceRoll *PriceRollSession) RenouncePauser() (*types.Transaction, error) {
	return _PriceRoll.Contract.RenouncePauser(&_PriceRoll.TransactOpts)
}

// RenouncePauser is a paid mutator transaction binding the contract method 0x6ef8d66d.
//
// Solidity: function renouncePauser() returns()
func (_PriceRoll *PriceRollTransactorSession) RenouncePauser() (*types.Transaction, error) {
	return _PriceRoll.Contract.RenouncePauser(&_PriceRoll.TransactOpts)
}

// SetBonus is a paid mutator transaction binding the contract method 0x0b98f975.
//
// Solidity: function setBonus(new_bonus uint256) returns()
func (_PriceRoll *PriceRollTransactor) SetBonus(opts *bind.TransactOpts, new_bonus *big.Int) (*types.Transaction, error) {
	return _PriceRoll.contract.Transact(opts, "setBonus", new_bonus)
}

// SetBonus is a paid mutator transaction binding the contract method 0x0b98f975.
//
// Solidity: function setBonus(new_bonus uint256) returns()
func (_PriceRoll *PriceRollSession) SetBonus(new_bonus *big.Int) (*types.Transaction, error) {
	return _PriceRoll.Contract.SetBonus(&_PriceRoll.TransactOpts, new_bonus)
}

// SetBonus is a paid mutator transaction binding the contract method 0x0b98f975.
//
// Solidity: function setBonus(new_bonus uint256) returns()
func (_PriceRoll *PriceRollTransactorSession) SetBonus(new_bonus *big.Int) (*types.Transaction, error) {
	return _PriceRoll.Contract.SetBonus(&_PriceRoll.TransactOpts, new_bonus)
}

// SetCooldown is a paid mutator transaction binding the contract method 0x4fc3f41a.
//
// Solidity: function setCooldown(new_cooldown uint256) returns()
func (_PriceRoll *PriceRollTransactor) SetCooldown(opts *bind.TransactOpts, new_cooldown *big.Int) (*types.Transaction, error) {
	return _PriceRoll.contract.Transact(opts, "setCooldown", new_cooldown)
}

// SetCooldown is a paid mutator transaction binding the contract method 0x4fc3f41a.
//
// Solidity: function setCooldown(new_cooldown uint256) returns()
func (_PriceRoll *PriceRollSession) SetCooldown(new_cooldown *big.Int) (*types.Transaction, error) {
	return _PriceRoll.Contract.SetCooldown(&_PriceRoll.TransactOpts, new_cooldown)
}

// SetCooldown is a paid mutator transaction binding the contract method 0x4fc3f41a.
//
// Solidity: function setCooldown(new_cooldown uint256) returns()
func (_PriceRoll *PriceRollTransactorSession) SetCooldown(new_cooldown *big.Int) (*types.Transaction, error) {
	return _PriceRoll.Contract.SetCooldown(&_PriceRoll.TransactOpts, new_cooldown)
}

// SetCutDestination is a paid mutator transaction binding the contract method 0xaad74722.
//
// Solidity: function setCutDestination(new_desination address) returns()
func (_PriceRoll *PriceRollTransactor) SetCutDestination(opts *bind.TransactOpts, new_desination common.Address) (*types.Transaction, error) {
	return _PriceRoll.contract.Transact(opts, "setCutDestination", new_desination)
}

// SetCutDestination is a paid mutator transaction binding the contract method 0xaad74722.
//
// Solidity: function setCutDestination(new_desination address) returns()
func (_PriceRoll *PriceRollSession) SetCutDestination(new_desination common.Address) (*types.Transaction, error) {
	return _PriceRoll.Contract.SetCutDestination(&_PriceRoll.TransactOpts, new_desination)
}

// SetCutDestination is a paid mutator transaction binding the contract method 0xaad74722.
//
// Solidity: function setCutDestination(new_desination address) returns()
func (_PriceRoll *PriceRollTransactorSession) SetCutDestination(new_desination common.Address) (*types.Transaction, error) {
	return _PriceRoll.Contract.SetCutDestination(&_PriceRoll.TransactOpts, new_desination)
}

// SetGasLimit is a paid mutator transaction binding the contract method 0xee7d72b4.
//
// Solidity: function setGasLimit(new_gaslimit uint256) returns()
func (_PriceRoll *PriceRollTransactor) SetGasLimit(opts *bind.TransactOpts, new_gaslimit *big.Int) (*types.Transaction, error) {
	return _PriceRoll.contract.Transact(opts, "setGasLimit", new_gaslimit)
}

// SetGasLimit is a paid mutator transaction binding the contract method 0xee7d72b4.
//
// Solidity: function setGasLimit(new_gaslimit uint256) returns()
func (_PriceRoll *PriceRollSession) SetGasLimit(new_gaslimit *big.Int) (*types.Transaction, error) {
	return _PriceRoll.Contract.SetGasLimit(&_PriceRoll.TransactOpts, new_gaslimit)
}

// SetGasLimit is a paid mutator transaction binding the contract method 0xee7d72b4.
//
// Solidity: function setGasLimit(new_gaslimit uint256) returns()
func (_PriceRoll *PriceRollTransactorSession) SetGasLimit(new_gaslimit *big.Int) (*types.Transaction, error) {
	return _PriceRoll.Contract.SetGasLimit(&_PriceRoll.TransactOpts, new_gaslimit)
}

// SetHouseEdge is a paid mutator transaction binding the contract method 0x6cd0f102.
//
// Solidity: function setHouseEdge(new_edge uint256) returns()
func (_PriceRoll *PriceRollTransactor) SetHouseEdge(opts *bind.TransactOpts, new_edge *big.Int) (*types.Transaction, error) {
	return _PriceRoll.contract.Transact(opts, "setHouseEdge", new_edge)
}

// SetHouseEdge is a paid mutator transaction binding the contract method 0x6cd0f102.
//
// Solidity: function setHouseEdge(new_edge uint256) returns()
func (_PriceRoll *PriceRollSession) SetHouseEdge(new_edge *big.Int) (*types.Transaction, error) {
	return _PriceRoll.Contract.SetHouseEdge(&_PriceRoll.TransactOpts, new_edge)
}

// SetHouseEdge is a paid mutator transaction binding the contract method 0x6cd0f102.
//
// Solidity: function setHouseEdge(new_edge uint256) returns()
func (_PriceRoll *PriceRollTransactorSession) SetHouseEdge(new_edge *big.Int) (*types.Transaction, error) {
	return _PriceRoll.Contract.SetHouseEdge(&_PriceRoll.TransactOpts, new_edge)
}

// SetMaxBet is a paid mutator transaction binding the contract method 0x881eff1e.
//
// Solidity: function setMaxBet(new_maxbet uint256) returns()
func (_PriceRoll *PriceRollTransactor) SetMaxBet(opts *bind.TransactOpts, new_maxbet *big.Int) (*types.Transaction, error) {
	return _PriceRoll.contract.Transact(opts, "setMaxBet", new_maxbet)
}

// SetMaxBet is a paid mutator transaction binding the contract method 0x881eff1e.
//
// Solidity: function setMaxBet(new_maxbet uint256) returns()
func (_PriceRoll *PriceRollSession) SetMaxBet(new_maxbet *big.Int) (*types.Transaction, error) {
	return _PriceRoll.Contract.SetMaxBet(&_PriceRoll.TransactOpts, new_maxbet)
}

// SetMaxBet is a paid mutator transaction binding the contract method 0x881eff1e.
//
// Solidity: function setMaxBet(new_maxbet uint256) returns()
func (_PriceRoll *PriceRollTransactorSession) SetMaxBet(new_maxbet *big.Int) (*types.Transaction, error) {
	return _PriceRoll.Contract.SetMaxBet(&_PriceRoll.TransactOpts, new_maxbet)
}

// SetMinBet is a paid mutator transaction binding the contract method 0x88ea41b9.
//
// Solidity: function setMinBet(new_minbet uint256) returns()
func (_PriceRoll *PriceRollTransactor) SetMinBet(opts *bind.TransactOpts, new_minbet *big.Int) (*types.Transaction, error) {
	return _PriceRoll.contract.Transact(opts, "setMinBet", new_minbet)
}

// SetMinBet is a paid mutator transaction binding the contract method 0x88ea41b9.
//
// Solidity: function setMinBet(new_minbet uint256) returns()
func (_PriceRoll *PriceRollSession) SetMinBet(new_minbet *big.Int) (*types.Transaction, error) {
	return _PriceRoll.Contract.SetMinBet(&_PriceRoll.TransactOpts, new_minbet)
}

// SetMinBet is a paid mutator transaction binding the contract method 0x88ea41b9.
//
// Solidity: function setMinBet(new_minbet uint256) returns()
func (_PriceRoll *PriceRollTransactorSession) SetMinBet(new_minbet *big.Int) (*types.Transaction, error) {
	return _PriceRoll.Contract.SetMinBet(&_PriceRoll.TransactOpts, new_minbet)
}

// SetPriceCheckDelay is a paid mutator transaction binding the contract method 0xfb8512f7.
//
// Solidity: function setPriceCheckDelay(new_delay uint256) returns()
func (_PriceRoll *PriceRollTransactor) SetPriceCheckDelay(opts *bind.TransactOpts, new_delay *big.Int) (*types.Transaction, error) {
	return _PriceRoll.contract.Transact(opts, "setPriceCheckDelay", new_delay)
}

// SetPriceCheckDelay is a paid mutator transaction binding the contract method 0xfb8512f7.
//
// Solidity: function setPriceCheckDelay(new_delay uint256) returns()
func (_PriceRoll *PriceRollSession) SetPriceCheckDelay(new_delay *big.Int) (*types.Transaction, error) {
	return _PriceRoll.Contract.SetPriceCheckDelay(&_PriceRoll.TransactOpts, new_delay)
}

// SetPriceCheckDelay is a paid mutator transaction binding the contract method 0xfb8512f7.
//
// Solidity: function setPriceCheckDelay(new_delay uint256) returns()
func (_PriceRoll *PriceRollTransactorSession) SetPriceCheckDelay(new_delay *big.Int) (*types.Transaction, error) {
	return _PriceRoll.Contract.SetPriceCheckDelay(&_PriceRoll.TransactOpts, new_delay)
}

// SetRandomGasLimit is a paid mutator transaction binding the contract method 0x1361b3b2.
//
// Solidity: function setRandomGasLimit(new_gaslimit uint256) returns()
func (_PriceRoll *PriceRollTransactor) SetRandomGasLimit(opts *bind.TransactOpts, new_gaslimit *big.Int) (*types.Transaction, error) {
	return _PriceRoll.contract.Transact(opts, "setRandomGasLimit", new_gaslimit)
}

// SetRandomGasLimit is a paid mutator transaction binding the contract method 0x1361b3b2.
//
// Solidity: function setRandomGasLimit(new_gaslimit uint256) returns()
func (_PriceRoll *PriceRollSession) SetRandomGasLimit(new_gaslimit *big.Int) (*types.Transaction, error) {
	return _PriceRoll.Contract.SetRandomGasLimit(&_PriceRoll.TransactOpts, new_gaslimit)
}

// SetRandomGasLimit is a paid mutator transaction binding the contract method 0x1361b3b2.
//
// Solidity: function setRandomGasLimit(new_gaslimit uint256) returns()
func (_PriceRoll *PriceRollTransactorSession) SetRandomGasLimit(new_gaslimit *big.Int) (*types.Transaction, error) {
	return _PriceRoll.Contract.SetRandomGasLimit(&_PriceRoll.TransactOpts, new_gaslimit)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(newOwner address) returns()
func (_PriceRoll *PriceRollTransactor) TransferOwnership(opts *bind.TransactOpts, newOwner common.Address) (*types.Transaction, error) {
	return _PriceRoll.contract.Transact(opts, "transferOwnership", newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(newOwner address) returns()
func (_PriceRoll *PriceRollSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _PriceRoll.Contract.TransferOwnership(&_PriceRoll.TransactOpts, newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(newOwner address) returns()
func (_PriceRoll *PriceRollTransactorSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _PriceRoll.Contract.TransferOwnership(&_PriceRoll.TransactOpts, newOwner)
}

// Unpause is a paid mutator transaction binding the contract method 0x3f4ba83a.
//
// Solidity: function unpause() returns()
func (_PriceRoll *PriceRollTransactor) Unpause(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _PriceRoll.contract.Transact(opts, "unpause")
}

// Unpause is a paid mutator transaction binding the contract method 0x3f4ba83a.
//
// Solidity: function unpause() returns()
func (_PriceRoll *PriceRollSession) Unpause() (*types.Transaction, error) {
	return _PriceRoll.Contract.Unpause(&_PriceRoll.TransactOpts)
}

// Unpause is a paid mutator transaction binding the contract method 0x3f4ba83a.
//
// Solidity: function unpause() returns()
func (_PriceRoll *PriceRollTransactorSession) Unpause() (*types.Transaction, error) {
	return _PriceRoll.Contract.Unpause(&_PriceRoll.TransactOpts)
}

// WithdrawHouse is a paid mutator transaction binding the contract method 0x23218411.
//
// Solidity: function withdrawHouse(amount uint256) returns()
func (_PriceRoll *PriceRollTransactor) WithdrawHouse(opts *bind.TransactOpts, amount *big.Int) (*types.Transaction, error) {
	return _PriceRoll.contract.Transact(opts, "withdrawHouse", amount)
}

// WithdrawHouse is a paid mutator transaction binding the contract method 0x23218411.
//
// Solidity: function withdrawHouse(amount uint256) returns()
func (_PriceRoll *PriceRollSession) WithdrawHouse(amount *big.Int) (*types.Transaction, error) {
	return _PriceRoll.Contract.WithdrawHouse(&_PriceRoll.TransactOpts, amount)
}

// WithdrawHouse is a paid mutator transaction binding the contract method 0x23218411.
//
// Solidity: function withdrawHouse(amount uint256) returns()
func (_PriceRoll *PriceRollTransactorSession) WithdrawHouse(amount *big.Int) (*types.Transaction, error) {
	return _PriceRoll.Contract.WithdrawHouse(&_PriceRoll.TransactOpts, amount)
}

// WithdrawWallet is a paid mutator transaction binding the contract method 0xbe87ab3f.
//
// Solidity: function withdrawWallet(amount uint256) returns()
func (_PriceRoll *PriceRollTransactor) WithdrawWallet(opts *bind.TransactOpts, amount *big.Int) (*types.Transaction, error) {
	return _PriceRoll.contract.Transact(opts, "withdrawWallet", amount)
}

// WithdrawWallet is a paid mutator transaction binding the contract method 0xbe87ab3f.
//
// Solidity: function withdrawWallet(amount uint256) returns()
func (_PriceRoll *PriceRollSession) WithdrawWallet(amount *big.Int) (*types.Transaction, error) {
	return _PriceRoll.Contract.WithdrawWallet(&_PriceRoll.TransactOpts, amount)
}

// WithdrawWallet is a paid mutator transaction binding the contract method 0xbe87ab3f.
//
// Solidity: function withdrawWallet(amount uint256) returns()
func (_PriceRoll *PriceRollTransactorSession) WithdrawWallet(amount *big.Int) (*types.Transaction, error) {
	return _PriceRoll.Contract.WithdrawWallet(&_PriceRoll.TransactOpts, amount)
}

// PriceRollBetPlacedIterator is returned from FilterBetPlaced and is used to iterate over the raw logs and unpacked data for BetPlaced events raised by the PriceRoll contract.
type PriceRollBetPlacedIterator struct {
	Event *PriceRollBetPlaced // Event containing the contract specifics and raw log

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
func (it *PriceRollBetPlacedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(PriceRollBetPlaced)
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
		it.Event = new(PriceRollBetPlaced)
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
func (it *PriceRollBetPlacedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *PriceRollBetPlacedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// PriceRollBetPlaced represents a BetPlaced event raised by the PriceRoll contract.
type PriceRollBetPlaced struct {
	Round         *big.Int
	Player        common.Address
	Amount        *big.Int
	ExpectedValue uint8
	IsUp          bool
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterBetPlaced is a free log retrieval operation binding the contract event 0x917ae17adb72fe5dfbe4db3784fc9c4489530a50dbd693383d1502cfe49e68fd.
//
// Solidity: e BetPlaced(round uint256, player indexed address, amount uint256, expected_value uint8, is_up bool)
func (_PriceRoll *PriceRollFilterer) FilterBetPlaced(opts *bind.FilterOpts, player []common.Address) (*PriceRollBetPlacedIterator, error) {

	var playerRule []interface{}
	for _, playerItem := range player {
		playerRule = append(playerRule, playerItem)
	}

	logs, sub, err := _PriceRoll.contract.FilterLogs(opts, "BetPlaced", playerRule)
	if err != nil {
		return nil, err
	}
	return &PriceRollBetPlacedIterator{contract: _PriceRoll.contract, event: "BetPlaced", logs: logs, sub: sub}, nil
}

// WatchBetPlaced is a free log subscription operation binding the contract event 0x917ae17adb72fe5dfbe4db3784fc9c4489530a50dbd693383d1502cfe49e68fd.
//
// Solidity: e BetPlaced(round uint256, player indexed address, amount uint256, expected_value uint8, is_up bool)
func (_PriceRoll *PriceRollFilterer) WatchBetPlaced(opts *bind.WatchOpts, sink chan<- *PriceRollBetPlaced, player []common.Address) (event.Subscription, error) {

	var playerRule []interface{}
	for _, playerItem := range player {
		playerRule = append(playerRule, playerItem)
	}

	logs, sub, err := _PriceRoll.contract.WatchLogs(opts, "BetPlaced", playerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(PriceRollBetPlaced)
				if err := _PriceRoll.contract.UnpackLog(event, "BetPlaced", log); err != nil {
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

// PriceRollNewRollIterator is returned from FilterNewRoll and is used to iterate over the raw logs and unpacked data for NewRoll events raised by the PriceRoll contract.
type PriceRollNewRollIterator struct {
	Event *PriceRollNewRoll // Event containing the contract specifics and raw log

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
func (it *PriceRollNewRollIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(PriceRollNewRoll)
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
		it.Event = new(PriceRollNewRoll)
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
func (it *PriceRollNewRollIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *PriceRollNewRollIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// PriceRollNewRoll represents a NewRoll event raised by the PriceRoll contract.
type PriceRollNewRoll struct {
	Round *big.Int
	Raw   types.Log // Blockchain specific contextual infos
}

// FilterNewRoll is a free log retrieval operation binding the contract event 0xb4267207f21a1858b0724dfba74816419152c88804d2b6276e5c32f6cd47c74a.
//
// Solidity: e NewRoll(round uint256)
func (_PriceRoll *PriceRollFilterer) FilterNewRoll(opts *bind.FilterOpts) (*PriceRollNewRollIterator, error) {

	logs, sub, err := _PriceRoll.contract.FilterLogs(opts, "NewRoll")
	if err != nil {
		return nil, err
	}
	return &PriceRollNewRollIterator{contract: _PriceRoll.contract, event: "NewRoll", logs: logs, sub: sub}, nil
}

// WatchNewRoll is a free log subscription operation binding the contract event 0xb4267207f21a1858b0724dfba74816419152c88804d2b6276e5c32f6cd47c74a.
//
// Solidity: e NewRoll(round uint256)
func (_PriceRoll *PriceRollFilterer) WatchNewRoll(opts *bind.WatchOpts, sink chan<- *PriceRollNewRoll) (event.Subscription, error) {

	logs, sub, err := _PriceRoll.contract.WatchLogs(opts, "NewRoll")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(PriceRollNewRoll)
				if err := _PriceRoll.contract.UnpackLog(event, "NewRoll", log); err != nil {
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

// PriceRollOraclizeErrorIterator is returned from FilterOraclizeError and is used to iterate over the raw logs and unpacked data for OraclizeError events raised by the PriceRoll contract.
type PriceRollOraclizeErrorIterator struct {
	Event *PriceRollOraclizeError // Event containing the contract specifics and raw log

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
func (it *PriceRollOraclizeErrorIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(PriceRollOraclizeError)
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
		it.Event = new(PriceRollOraclizeError)
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
func (it *PriceRollOraclizeErrorIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *PriceRollOraclizeErrorIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// PriceRollOraclizeError represents a OraclizeError event raised by the PriceRoll contract.
type PriceRollOraclizeError struct {
	Value *big.Int
	Raw   types.Log // Blockchain specific contextual infos
}

// FilterOraclizeError is a free log retrieval operation binding the contract event 0x0510c169028b6c6e93e31f93b1b642b756cea80cc1966887b5c8f0b93d3d5b68.
//
// Solidity: e OraclizeError(value uint256)
func (_PriceRoll *PriceRollFilterer) FilterOraclizeError(opts *bind.FilterOpts) (*PriceRollOraclizeErrorIterator, error) {

	logs, sub, err := _PriceRoll.contract.FilterLogs(opts, "OraclizeError")
	if err != nil {
		return nil, err
	}
	return &PriceRollOraclizeErrorIterator{contract: _PriceRoll.contract, event: "OraclizeError", logs: logs, sub: sub}, nil
}

// WatchOraclizeError is a free log subscription operation binding the contract event 0x0510c169028b6c6e93e31f93b1b642b756cea80cc1966887b5c8f0b93d3d5b68.
//
// Solidity: e OraclizeError(value uint256)
func (_PriceRoll *PriceRollFilterer) WatchOraclizeError(opts *bind.WatchOpts, sink chan<- *PriceRollOraclizeError) (event.Subscription, error) {

	logs, sub, err := _PriceRoll.contract.WatchLogs(opts, "OraclizeError")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(PriceRollOraclizeError)
				if err := _PriceRoll.contract.UnpackLog(event, "OraclizeError", log); err != nil {
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

// PriceRollOwnershipTransferredIterator is returned from FilterOwnershipTransferred and is used to iterate over the raw logs and unpacked data for OwnershipTransferred events raised by the PriceRoll contract.
type PriceRollOwnershipTransferredIterator struct {
	Event *PriceRollOwnershipTransferred // Event containing the contract specifics and raw log

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
func (it *PriceRollOwnershipTransferredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(PriceRollOwnershipTransferred)
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
		it.Event = new(PriceRollOwnershipTransferred)
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
func (it *PriceRollOwnershipTransferredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *PriceRollOwnershipTransferredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// PriceRollOwnershipTransferred represents a OwnershipTransferred event raised by the PriceRoll contract.
type PriceRollOwnershipTransferred struct {
	PreviousOwner common.Address
	NewOwner      common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterOwnershipTransferred is a free log retrieval operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: e OwnershipTransferred(previousOwner indexed address, newOwner indexed address)
func (_PriceRoll *PriceRollFilterer) FilterOwnershipTransferred(opts *bind.FilterOpts, previousOwner []common.Address, newOwner []common.Address) (*PriceRollOwnershipTransferredIterator, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _PriceRoll.contract.FilterLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return &PriceRollOwnershipTransferredIterator{contract: _PriceRoll.contract, event: "OwnershipTransferred", logs: logs, sub: sub}, nil
}

// WatchOwnershipTransferred is a free log subscription operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: e OwnershipTransferred(previousOwner indexed address, newOwner indexed address)
func (_PriceRoll *PriceRollFilterer) WatchOwnershipTransferred(opts *bind.WatchOpts, sink chan<- *PriceRollOwnershipTransferred, previousOwner []common.Address, newOwner []common.Address) (event.Subscription, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _PriceRoll.contract.WatchLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(PriceRollOwnershipTransferred)
				if err := _PriceRoll.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
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

// PriceRollPausedIterator is returned from FilterPaused and is used to iterate over the raw logs and unpacked data for Paused events raised by the PriceRoll contract.
type PriceRollPausedIterator struct {
	Event *PriceRollPaused // Event containing the contract specifics and raw log

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
func (it *PriceRollPausedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(PriceRollPaused)
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
		it.Event = new(PriceRollPaused)
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
func (it *PriceRollPausedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *PriceRollPausedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// PriceRollPaused represents a Paused event raised by the PriceRoll contract.
type PriceRollPaused struct {
	Account common.Address
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterPaused is a free log retrieval operation binding the contract event 0x62e78cea01bee320cd4e420270b5ea74000d11b0c9f74754ebdbfc544b05a258.
//
// Solidity: e Paused(account address)
func (_PriceRoll *PriceRollFilterer) FilterPaused(opts *bind.FilterOpts) (*PriceRollPausedIterator, error) {

	logs, sub, err := _PriceRoll.contract.FilterLogs(opts, "Paused")
	if err != nil {
		return nil, err
	}
	return &PriceRollPausedIterator{contract: _PriceRoll.contract, event: "Paused", logs: logs, sub: sub}, nil
}

// WatchPaused is a free log subscription operation binding the contract event 0x62e78cea01bee320cd4e420270b5ea74000d11b0c9f74754ebdbfc544b05a258.
//
// Solidity: e Paused(account address)
func (_PriceRoll *PriceRollFilterer) WatchPaused(opts *bind.WatchOpts, sink chan<- *PriceRollPaused) (event.Subscription, error) {

	logs, sub, err := _PriceRoll.contract.WatchLogs(opts, "Paused")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(PriceRollPaused)
				if err := _PriceRoll.contract.UnpackLog(event, "Paused", log); err != nil {
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

// PriceRollPauserAddedIterator is returned from FilterPauserAdded and is used to iterate over the raw logs and unpacked data for PauserAdded events raised by the PriceRoll contract.
type PriceRollPauserAddedIterator struct {
	Event *PriceRollPauserAdded // Event containing the contract specifics and raw log

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
func (it *PriceRollPauserAddedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(PriceRollPauserAdded)
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
		it.Event = new(PriceRollPauserAdded)
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
func (it *PriceRollPauserAddedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *PriceRollPauserAddedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// PriceRollPauserAdded represents a PauserAdded event raised by the PriceRoll contract.
type PriceRollPauserAdded struct {
	Account common.Address
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterPauserAdded is a free log retrieval operation binding the contract event 0x6719d08c1888103bea251a4ed56406bd0c3e69723c8a1686e017e7bbe159b6f8.
//
// Solidity: e PauserAdded(account indexed address)
func (_PriceRoll *PriceRollFilterer) FilterPauserAdded(opts *bind.FilterOpts, account []common.Address) (*PriceRollPauserAddedIterator, error) {

	var accountRule []interface{}
	for _, accountItem := range account {
		accountRule = append(accountRule, accountItem)
	}

	logs, sub, err := _PriceRoll.contract.FilterLogs(opts, "PauserAdded", accountRule)
	if err != nil {
		return nil, err
	}
	return &PriceRollPauserAddedIterator{contract: _PriceRoll.contract, event: "PauserAdded", logs: logs, sub: sub}, nil
}

// WatchPauserAdded is a free log subscription operation binding the contract event 0x6719d08c1888103bea251a4ed56406bd0c3e69723c8a1686e017e7bbe159b6f8.
//
// Solidity: e PauserAdded(account indexed address)
func (_PriceRoll *PriceRollFilterer) WatchPauserAdded(opts *bind.WatchOpts, sink chan<- *PriceRollPauserAdded, account []common.Address) (event.Subscription, error) {

	var accountRule []interface{}
	for _, accountItem := range account {
		accountRule = append(accountRule, accountItem)
	}

	logs, sub, err := _PriceRoll.contract.WatchLogs(opts, "PauserAdded", accountRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(PriceRollPauserAdded)
				if err := _PriceRoll.contract.UnpackLog(event, "PauserAdded", log); err != nil {
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

// PriceRollPauserRemovedIterator is returned from FilterPauserRemoved and is used to iterate over the raw logs and unpacked data for PauserRemoved events raised by the PriceRoll contract.
type PriceRollPauserRemovedIterator struct {
	Event *PriceRollPauserRemoved // Event containing the contract specifics and raw log

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
func (it *PriceRollPauserRemovedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(PriceRollPauserRemoved)
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
		it.Event = new(PriceRollPauserRemoved)
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
func (it *PriceRollPauserRemovedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *PriceRollPauserRemovedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// PriceRollPauserRemoved represents a PauserRemoved event raised by the PriceRoll contract.
type PriceRollPauserRemoved struct {
	Account common.Address
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterPauserRemoved is a free log retrieval operation binding the contract event 0xcd265ebaf09df2871cc7bd4133404a235ba12eff2041bb89d9c714a2621c7c7e.
//
// Solidity: e PauserRemoved(account indexed address)
func (_PriceRoll *PriceRollFilterer) FilterPauserRemoved(opts *bind.FilterOpts, account []common.Address) (*PriceRollPauserRemovedIterator, error) {

	var accountRule []interface{}
	for _, accountItem := range account {
		accountRule = append(accountRule, accountItem)
	}

	logs, sub, err := _PriceRoll.contract.FilterLogs(opts, "PauserRemoved", accountRule)
	if err != nil {
		return nil, err
	}
	return &PriceRollPauserRemovedIterator{contract: _PriceRoll.contract, event: "PauserRemoved", logs: logs, sub: sub}, nil
}

// WatchPauserRemoved is a free log subscription operation binding the contract event 0xcd265ebaf09df2871cc7bd4133404a235ba12eff2041bb89d9c714a2621c7c7e.
//
// Solidity: e PauserRemoved(account indexed address)
func (_PriceRoll *PriceRollFilterer) WatchPauserRemoved(opts *bind.WatchOpts, sink chan<- *PriceRollPauserRemoved, account []common.Address) (event.Subscription, error) {

	var accountRule []interface{}
	for _, accountItem := range account {
		accountRule = append(accountRule, accountItem)
	}

	logs, sub, err := _PriceRoll.contract.WatchLogs(opts, "PauserRemoved", accountRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(PriceRollPauserRemoved)
				if err := _PriceRoll.contract.UnpackLog(event, "PauserRemoved", log); err != nil {
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

// PriceRollRollClaimedIterator is returned from FilterRollClaimed and is used to iterate over the raw logs and unpacked data for RollClaimed events raised by the PriceRoll contract.
type PriceRollRollClaimedIterator struct {
	Event *PriceRollRollClaimed // Event containing the contract specifics and raw log

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
func (it *PriceRollRollClaimedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(PriceRollRollClaimed)
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
		it.Event = new(PriceRollRollClaimed)
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
func (it *PriceRollRollClaimedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *PriceRollRollClaimedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// PriceRollRollClaimed represents a RollClaimed event raised by the PriceRoll contract.
type PriceRollRollClaimed struct {
	Round  *big.Int
	Player common.Address
	Amount *big.Int
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterRollClaimed is a free log retrieval operation binding the contract event 0x170c93f55dab07b2a0c856acad14fbb479fc97d6c3f489faf09b5f01c1575d3e.
//
// Solidity: e RollClaimed(round uint256, player indexed address, amount uint256)
func (_PriceRoll *PriceRollFilterer) FilterRollClaimed(opts *bind.FilterOpts, player []common.Address) (*PriceRollRollClaimedIterator, error) {

	var playerRule []interface{}
	for _, playerItem := range player {
		playerRule = append(playerRule, playerItem)
	}

	logs, sub, err := _PriceRoll.contract.FilterLogs(opts, "RollClaimed", playerRule)
	if err != nil {
		return nil, err
	}
	return &PriceRollRollClaimedIterator{contract: _PriceRoll.contract, event: "RollClaimed", logs: logs, sub: sub}, nil
}

// WatchRollClaimed is a free log subscription operation binding the contract event 0x170c93f55dab07b2a0c856acad14fbb479fc97d6c3f489faf09b5f01c1575d3e.
//
// Solidity: e RollClaimed(round uint256, player indexed address, amount uint256)
func (_PriceRoll *PriceRollFilterer) WatchRollClaimed(opts *bind.WatchOpts, sink chan<- *PriceRollRollClaimed, player []common.Address) (event.Subscription, error) {

	var playerRule []interface{}
	for _, playerItem := range player {
		playerRule = append(playerRule, playerItem)
	}

	logs, sub, err := _PriceRoll.contract.WatchLogs(opts, "RollClaimed", playerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(PriceRollRollClaimed)
				if err := _PriceRoll.contract.UnpackLog(event, "RollClaimed", log); err != nil {
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

// PriceRollRollEndedIterator is returned from FilterRollEnded and is used to iterate over the raw logs and unpacked data for RollEnded events raised by the PriceRoll contract.
type PriceRollRollEndedIterator struct {
	Event *PriceRollRollEnded // Event containing the contract specifics and raw log

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
func (it *PriceRollRollEndedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(PriceRollRollEnded)
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
		it.Event = new(PriceRollRollEnded)
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
func (it *PriceRollRollEndedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *PriceRollRollEndedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// PriceRollRollEnded represents a RollEnded event raised by the PriceRoll contract.
type PriceRollRollEnded struct {
	Round      *big.Int
	Value      uint8
	StartPrice *big.Int
	EndPrice   *big.Int
	Raw        types.Log // Blockchain specific contextual infos
}

// FilterRollEnded is a free log retrieval operation binding the contract event 0xfa1b3a1a2a7d28c6303b4f4d3f3dc998928c92532bf19e80f8bfa461528e4569.
//
// Solidity: e RollEnded(round uint256, value uint8, start_price uint256, end_price uint256)
func (_PriceRoll *PriceRollFilterer) FilterRollEnded(opts *bind.FilterOpts) (*PriceRollRollEndedIterator, error) {

	logs, sub, err := _PriceRoll.contract.FilterLogs(opts, "RollEnded")
	if err != nil {
		return nil, err
	}
	return &PriceRollRollEndedIterator{contract: _PriceRoll.contract, event: "RollEnded", logs: logs, sub: sub}, nil
}

// WatchRollEnded is a free log subscription operation binding the contract event 0xfa1b3a1a2a7d28c6303b4f4d3f3dc998928c92532bf19e80f8bfa461528e4569.
//
// Solidity: e RollEnded(round uint256, value uint8, start_price uint256, end_price uint256)
func (_PriceRoll *PriceRollFilterer) WatchRollEnded(opts *bind.WatchOpts, sink chan<- *PriceRollRollEnded) (event.Subscription, error) {

	logs, sub, err := _PriceRoll.contract.WatchLogs(opts, "RollEnded")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(PriceRollRollEnded)
				if err := _PriceRoll.contract.UnpackLog(event, "RollEnded", log); err != nil {
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

// PriceRollRollRefundedIterator is returned from FilterRollRefunded and is used to iterate over the raw logs and unpacked data for RollRefunded events raised by the PriceRoll contract.
type PriceRollRollRefundedIterator struct {
	Event *PriceRollRollRefunded // Event containing the contract specifics and raw log

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
func (it *PriceRollRollRefundedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(PriceRollRollRefunded)
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
		it.Event = new(PriceRollRollRefunded)
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
func (it *PriceRollRollRefundedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *PriceRollRollRefundedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// PriceRollRollRefunded represents a RollRefunded event raised by the PriceRoll contract.
type PriceRollRollRefunded struct {
	Round *big.Int
	Raw   types.Log // Blockchain specific contextual infos
}

// FilterRollRefunded is a free log retrieval operation binding the contract event 0x796cfdf016e01ffdff99dbe8ce27df476b06d0fae2ac92ec3889b6008219ed52.
//
// Solidity: e RollRefunded(round uint256)
func (_PriceRoll *PriceRollFilterer) FilterRollRefunded(opts *bind.FilterOpts) (*PriceRollRollRefundedIterator, error) {

	logs, sub, err := _PriceRoll.contract.FilterLogs(opts, "RollRefunded")
	if err != nil {
		return nil, err
	}
	return &PriceRollRollRefundedIterator{contract: _PriceRoll.contract, event: "RollRefunded", logs: logs, sub: sub}, nil
}

// WatchRollRefunded is a free log subscription operation binding the contract event 0x796cfdf016e01ffdff99dbe8ce27df476b06d0fae2ac92ec3889b6008219ed52.
//
// Solidity: e RollRefunded(round uint256)
func (_PriceRoll *PriceRollFilterer) WatchRollRefunded(opts *bind.WatchOpts, sink chan<- *PriceRollRollRefunded) (event.Subscription, error) {

	logs, sub, err := _PriceRoll.contract.WatchLogs(opts, "RollRefunded")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(PriceRollRollRefunded)
				if err := _PriceRoll.contract.UnpackLog(event, "RollRefunded", log); err != nil {
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

// PriceRollRollingIterator is returned from FilterRolling and is used to iterate over the raw logs and unpacked data for Rolling events raised by the PriceRoll contract.
type PriceRollRollingIterator struct {
	Event *PriceRollRolling // Event containing the contract specifics and raw log

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
func (it *PriceRollRollingIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(PriceRollRolling)
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
		it.Event = new(PriceRollRolling)
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
func (it *PriceRollRollingIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *PriceRollRollingIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// PriceRollRolling represents a Rolling event raised by the PriceRoll contract.
type PriceRollRolling struct {
	Round *big.Int
	Raw   types.Log // Blockchain specific contextual infos
}

// FilterRolling is a free log retrieval operation binding the contract event 0xe86c6fb9a82a2c1fd4c6c98a2fe03844f019514b0cfe8dfbdf076c61f124c0d4.
//
// Solidity: e Rolling(round uint256)
func (_PriceRoll *PriceRollFilterer) FilterRolling(opts *bind.FilterOpts) (*PriceRollRollingIterator, error) {

	logs, sub, err := _PriceRoll.contract.FilterLogs(opts, "Rolling")
	if err != nil {
		return nil, err
	}
	return &PriceRollRollingIterator{contract: _PriceRoll.contract, event: "Rolling", logs: logs, sub: sub}, nil
}

// WatchRolling is a free log subscription operation binding the contract event 0xe86c6fb9a82a2c1fd4c6c98a2fe03844f019514b0cfe8dfbdf076c61f124c0d4.
//
// Solidity: e Rolling(round uint256)
func (_PriceRoll *PriceRollFilterer) WatchRolling(opts *bind.WatchOpts, sink chan<- *PriceRollRolling) (event.Subscription, error) {

	logs, sub, err := _PriceRoll.contract.WatchLogs(opts, "Rolling")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(PriceRollRolling)
				if err := _PriceRoll.contract.UnpackLog(event, "Rolling", log); err != nil {
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

// PriceRollUnpausedIterator is returned from FilterUnpaused and is used to iterate over the raw logs and unpacked data for Unpaused events raised by the PriceRoll contract.
type PriceRollUnpausedIterator struct {
	Event *PriceRollUnpaused // Event containing the contract specifics and raw log

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
func (it *PriceRollUnpausedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(PriceRollUnpaused)
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
		it.Event = new(PriceRollUnpaused)
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
func (it *PriceRollUnpausedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *PriceRollUnpausedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// PriceRollUnpaused represents a Unpaused event raised by the PriceRoll contract.
type PriceRollUnpaused struct {
	Account common.Address
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterUnpaused is a free log retrieval operation binding the contract event 0x5db9ee0a495bf2e6ff9c91a7834c1ba4fdd244a5e8aa4e537bd38aeae4b073aa.
//
// Solidity: e Unpaused(account address)
func (_PriceRoll *PriceRollFilterer) FilterUnpaused(opts *bind.FilterOpts) (*PriceRollUnpausedIterator, error) {

	logs, sub, err := _PriceRoll.contract.FilterLogs(opts, "Unpaused")
	if err != nil {
		return nil, err
	}
	return &PriceRollUnpausedIterator{contract: _PriceRoll.contract, event: "Unpaused", logs: logs, sub: sub}, nil
}

// WatchUnpaused is a free log subscription operation binding the contract event 0x5db9ee0a495bf2e6ff9c91a7834c1ba4fdd244a5e8aa4e537bd38aeae4b073aa.
//
// Solidity: e Unpaused(account address)
func (_PriceRoll *PriceRollFilterer) WatchUnpaused(opts *bind.WatchOpts, sink chan<- *PriceRollUnpaused) (event.Subscription, error) {

	logs, sub, err := _PriceRoll.contract.WatchLogs(opts, "Unpaused")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(PriceRollUnpaused)
				if err := _PriceRoll.contract.UnpackLog(event, "Unpaused", log); err != nil {
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
const RolesBin = `0x604c602c600b82828239805160001a60731460008114601c57601e565bfe5b5030600052607381538281f3fe73000000000000000000000000000000000000000030146080604052600080fdfea165627a7a72305820c415b73fe772266f27c46d5b208da184b13577e8d5144d7a59a8cdd3fc187db60029`

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
const SafeMathBin = `0x604c602c600b82828239805160001a60731460008114601c57601e565bfe5b5030600052607381538281f3fe73000000000000000000000000000000000000000030146080604052600080fdfea165627a7a72305820684c4fbaa4f43f7da2adfe164a6f11773ce4dd67222413f0d1bd548be754ee280029`

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
// Solidity: function f(x bytes) returns()
func (_SolcChecker *SolcCheckerTransactor) F(opts *bind.TransactOpts, x []byte) (*types.Transaction, error) {
	return _SolcChecker.contract.Transact(opts, "f", x)
}

// F is a paid mutator transaction binding the contract method 0xd45754f8.
//
// Solidity: function f(x bytes) returns()
func (_SolcChecker *SolcCheckerSession) F(x []byte) (*types.Transaction, error) {
	return _SolcChecker.Contract.F(&_SolcChecker.TransactOpts, x)
}

// F is a paid mutator transaction binding the contract method 0xd45754f8.
//
// Solidity: function f(x bytes) returns()
func (_SolcChecker *SolcCheckerTransactorSession) F(x []byte) (*types.Transaction, error) {
	return _SolcChecker.Contract.F(&_SolcChecker.TransactOpts, x)
}

// UsingOraclizeABI is the input ABI used to generate the binding from.
const UsingOraclizeABI = "[{\"constant\":false,\"inputs\":[{\"name\":\"_myid\",\"type\":\"bytes32\"},{\"name\":\"_result\",\"type\":\"string\"}],\"name\":\"__callback\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_myid\",\"type\":\"bytes32\"},{\"name\":\"_result\",\"type\":\"string\"},{\"name\":\"_proof\",\"type\":\"bytes\"}],\"name\":\"__callback\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]"

// UsingOraclizeBin is the compiled bytecode used for deploying new contracts.
const UsingOraclizeBin = `0x608060405234801561001057600080fd5b5061028b806100206000396000f3fe608060405234801561001057600080fd5b5060043610610052577c0100000000000000000000000000000000000000000000000000000000600035046327dc297e811461005757806338bbfa5014610106575b600080fd5b6101046004803603604081101561006d57600080fd5b8135919081019060408101602082013564010000000081111561008f57600080fd5b8201836020820111156100a157600080fd5b803590602001918460018302840111640100000000831117156100c357600080fd5b91908080601f01602080910402602001604051908101604052809392919081815260200183838082843760009201919091525092955061023a945050505050565b005b6101046004803603606081101561011c57600080fd5b8135919081019060408101602082013564010000000081111561013e57600080fd5b82018360208201111561015057600080fd5b8035906020019184600183028401116401000000008311171561017257600080fd5b91908080601f01602080910402602001604051908101604052809392919081815260200183838082843760009201919091525092959493602081019350359150506401000000008111156101c557600080fd5b8201836020820111156101d757600080fd5b803590602001918460018302840111640100000000831117156101f957600080fd5b91908080601f01602080910402602001604051908101604052809392919081815260200183838082843760009201919091525092955061025a945050505050565b604080516000815260208101909152610256908390839061025a565b5050565b50505056fea165627a7a723058201ac246271fb01a5987b667b34445619d9449ae7432f2375145de8ca98dc062b20029`

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
// Solidity: function __callback(_myid bytes32, _result string, _proof bytes) returns()
func (_UsingOraclize *UsingOraclizeTransactor) Callback(opts *bind.TransactOpts, _myid [32]byte, _result string, _proof []byte) (*types.Transaction, error) {
	return _UsingOraclize.contract.Transact(opts, "__callback", _myid, _result, _proof)
}

// Callback is a paid mutator transaction binding the contract method 0x38bbfa50.
//
// Solidity: function __callback(_myid bytes32, _result string, _proof bytes) returns()
func (_UsingOraclize *UsingOraclizeSession) Callback(_myid [32]byte, _result string, _proof []byte) (*types.Transaction, error) {
	return _UsingOraclize.Contract.Callback(&_UsingOraclize.TransactOpts, _myid, _result, _proof)
}

// Callback is a paid mutator transaction binding the contract method 0x38bbfa50.
//
// Solidity: function __callback(_myid bytes32, _result string, _proof bytes) returns()
func (_UsingOraclize *UsingOraclizeTransactorSession) Callback(_myid [32]byte, _result string, _proof []byte) (*types.Transaction, error) {
	return _UsingOraclize.Contract.Callback(&_UsingOraclize.TransactOpts, _myid, _result, _proof)
}

