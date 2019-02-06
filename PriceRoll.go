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
const BufferBin = `0x604c602c600b82828239805160001a60731460008114601c57601e565bfe5b5030600052607381538281f3fe73000000000000000000000000000000000000000030146080604052600080fdfea165627a7a72305820be8923b417d8a52a35781409cdf848f8a4b9df196b6033ef8fc9ad0c5e91e7d10029`

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
const CBORBin = `0x604c602c600b82828239805160001a60731460008114601c57601e565bfe5b5030600052607381538281f3fe73000000000000000000000000000000000000000030146080604052600080fdfea165627a7a72305820e2f44c37d2b6a22aa02c9b8c23c1c179db67608d18fba374eb9857c0599c8c3d0029`

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
const PriceRollABI = "[{\"constant\":true,\"inputs\":[],\"name\":\"latest_roll\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"current_coin\",\"outputs\":[{\"name\":\"\",\"type\":\"uint8\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"query_stringETH\",\"outputs\":[{\"name\":\"\",\"type\":\"string\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"new_bonus\",\"type\":\"uint256\"}],\"name\":\"setBonus\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"new_gaslimit\",\"type\":\"uint256\"}],\"name\":\"setRandomGasLimit\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"pool\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"withdrawHouse\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_myid\",\"type\":\"bytes32\"},{\"name\":\"_result\",\"type\":\"string\"}],\"name\":\"__callback\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"expected_value\",\"type\":\"uint8\"},{\"name\":\"is_up\",\"type\":\"bool\"}],\"name\":\"placeBet\",\"outputs\":[],\"payable\":true,\"stateMutability\":\"payable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"round\",\"type\":\"uint256\"}],\"name\":\"claim\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_queryId\",\"type\":\"bytes32\"},{\"name\":\"_result\",\"type\":\"string\"},{\"name\":\"_proof\",\"type\":\"bytes\"}],\"name\":\"__callback\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[],\"name\":\"unpause\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"config_bonus_mult\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"account\",\"type\":\"address\"}],\"name\":\"isPauser\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"current_roll\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"new_cooldown\",\"type\":\"uint256\"}],\"name\":\"setCooldown\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"config_refund_delay\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"config_random_gas_limit\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"config_pricecheck_delay\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"paused\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"config_house_edge\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"rolls\",\"outputs\":[{\"name\":\"result_rngseed\",\"type\":\"string\"},{\"name\":\"query_rng\",\"type\":\"bytes32\"},{\"name\":\"query_price1\",\"type\":\"bytes32\"},{\"name\":\"query_price2\",\"type\":\"bytes32\"},{\"name\":\"result_price1\",\"type\":\"uint256\"},{\"name\":\"result_price2\",\"type\":\"uint256\"},{\"name\":\"timestamp\",\"type\":\"uint256\"},{\"name\":\"pool\",\"type\":\"uint256\"},{\"name\":\"state\",\"type\":\"uint8\"},{\"name\":\"coin\",\"type\":\"uint8\"},{\"name\":\"is_up\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"new_edge\",\"type\":\"uint256\"}],\"name\":\"setHouseEdge\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[],\"name\":\"renouncePauser\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[],\"name\":\"newRoll\",\"outputs\":[],\"payable\":true,\"stateMutability\":\"payable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"\",\"type\":\"address\"}],\"name\":\"balanceOf\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[],\"name\":\"renounceOwnership\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"amount\",\"type\":\"uint256\"},{\"name\":\"expected_value\",\"type\":\"uint8\"},{\"name\":\"is_up\",\"type\":\"bool\"}],\"name\":\"betFromInternalWallet\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"account\",\"type\":\"address\"}],\"name\":\"addPauser\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[],\"name\":\"destroy\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[],\"name\":\"pause\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"new_maxbet\",\"type\":\"uint256\"}],\"name\":\"setMaxBet\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"new_minbet\",\"type\":\"uint256\"}],\"name\":\"setMinBet\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"config_max_bet\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"isOwner\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"query_stringLTC\",\"outputs\":[{\"name\":\"\",\"type\":\"string\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"config_roll_cooldown\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"new_desination\",\"type\":\"address\"}],\"name\":\"setCutDestination\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[],\"name\":\"creditWallet\",\"outputs\":[],\"payable\":true,\"stateMutability\":\"payable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"withdrawWallet\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"userBalance\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"config_gas_limit\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"query_stringBTC\",\"outputs\":[{\"name\":\"\",\"type\":\"string\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"config_min_bet\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"new_gaslimit\",\"type\":\"uint256\"}],\"name\":\"setGasLimit\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"transferOwnership\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"new_delay\",\"type\":\"uint256\"}],\"name\":\"setPriceCheckDelay\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"house\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"payable\":true,\"stateMutability\":\"payable\",\"type\":\"fallback\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"round\",\"type\":\"uint256\"}],\"name\":\"Rolling\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"round\",\"type\":\"uint256\"}],\"name\":\"NewRoll\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"round\",\"type\":\"uint256\"},{\"indexed\":false,\"name\":\"seed\",\"type\":\"string\"},{\"indexed\":false,\"name\":\"start_price\",\"type\":\"uint256\"},{\"indexed\":false,\"name\":\"end_price\",\"type\":\"uint256\"}],\"name\":\"RollEnded\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"round\",\"type\":\"uint256\"}],\"name\":\"RollRefunded\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"round\",\"type\":\"uint256\"},{\"indexed\":true,\"name\":\"player\",\"type\":\"address\"},{\"indexed\":false,\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"RollClaimed\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"round\",\"type\":\"uint256\"},{\"indexed\":false,\"name\":\"amount\",\"type\":\"uint256\"},{\"indexed\":true,\"name\":\"player\",\"type\":\"address\"},{\"indexed\":false,\"name\":\"expected_value\",\"type\":\"uint8\"},{\"indexed\":false,\"name\":\"is_up\",\"type\":\"uint8\"}],\"name\":\"BetPlaced\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"OraclizeError\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"previousOwner\",\"type\":\"address\"},{\"indexed\":true,\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"OwnershipTransferred\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"name\":\"account\",\"type\":\"address\"}],\"name\":\"Paused\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"name\":\"account\",\"type\":\"address\"}],\"name\":\"Unpaused\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"account\",\"type\":\"address\"}],\"name\":\"PauserAdded\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"account\",\"type\":\"address\"}],\"name\":\"PauserRemoved\",\"type\":\"event\"}]"

// PriceRollBin is the compiled bytecode used for deploying new contracts.
const PriceRollBin = `0x6080604052603c6007819055610bb860085562030d406009819055600a5566470de4df820000600b5568056bc75e2d63100000600c556014600d55604b600e55600f5560108054600160a060020a03191673a54741f7fe21689b59bd7eacbf3a2947cd3f3bd4179055600060118190556012556013805460ff191690553480156200008957600080fd5b506200009e336401000000006200010c810204565b60068054600160a860020a0319166101003381029190911791829055604051600160a060020a039190920416906000907f8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0908290a3620001066401000000006200015e810204565b620002a9565b6200012760058264010000000062004437620001fc82021704565b604051600160a060020a038216907f6719d08c1888103bea251a4ed56406bd0c3e69723c8a1686e017e7bbe159b6f890600090a250565b6011546200017c90600164010000000062000c0b6200025782021704565b60115560135460039060ff1660028111156200019457fe5b600101811515620001a157fe5b066002811115620001ae57fe5b6013805460ff19166001836002811115620001c557fe5b0217905550426012556011546040517fb4267207f21a1858b0724dfba74816419152c88804d2b6276e5c32f6cd47c74a90600090a2565b600160a060020a03811615156200021257600080fd5b62000227828264010000000062000271810204565b156200023257600080fd5b600160a060020a0316600090815260209190915260409020805460ff19166001179055565b6000828201838110156200026a57600080fd5b9392505050565b6000600160a060020a03821615156200028957600080fd5b50600160a060020a03166000908152602091909152604090205460ff1690565b615abc80620002b96000396000f3fe6080604052600436106102b25760003560e060020a900480636fabd2651161017957806390e2c9dd116100e0578063d316401311610099578063ee7d72b411610073578063ee7d72b414610b6f578063f2fde38b14610b99578063fb8512f714610bcc578063ff9b3acf14610bf6576102b2565b8063d316401314610b30578063d8a3f00f14610b45578063deedc5c514610b5a576102b2565b806390e2c9dd14610a8c578063987ae15414610aa1578063aad7472214610ab6578063b8b069fc14610ae9578063be87ab3f14610af1578063bf15276514610b1b576102b2565b80638456cb59116101325780638456cb59146109c8578063881eff1e146109dd57806388ea41b914610a075780638b1dc09e14610a315780638da5cb5b14610a465780638f32d59b14610a77576102b2565b80636fabd265146108f557806370a08231146108fd578063715018a6146109305780637de36c971461094557806382dc1ec41461098057806383197ef0146109b3576102b2565b8063402684121161021d5780635ad51f68116101d65780635ad51f681461076f5780635c975abb146107845780635d10c029146107995780635d69f16c146107ae5780636cd0f102146108b65780636ef8d66d146108e0576102b2565b806340268412146106aa57806346fbf68e146106bf5780634e0c12f0146107065780634fc3f41a1461071b578063550098991461074557806357b6702f1461075a576102b2565b8063232184111161026f578063232184111461042157806327dc297e1461044b5780632ed1317c14610504578063379607f51461052c57806338bbfa50146105565780633f4ba83a14610695576102b2565b8063022b02bc146102ca57806306792c07146102f1578063091085701461032a5780630b98f975146103b65780631361b3b2146103e257806316f0115b1461040c575b6015546102c5903463ffffffff610c0b16565b601555005b3480156102d657600080fd5b506102df610c26565b60405190815260200160405180910390f35b3480156102fd57600080fd5b50610306610c2c565b6040518082600281111561031657fe5b60ff16815260200191505060405180910390f35b34801561033657600080fd5b5061033f610c35565b60405160208082528190810183818151815260200191508051906020019080838360005b8381101561037b578082015183820152602001610363565b50505050905090810190601f1680156103a85780820380516001836020036101000a031916815260200191505b509250505060405180910390f35b3480156103c257600080fd5b506103e0600480360360208110156103d957600080fd5b5035610c52565b005b3480156103ee57600080fd5b506103e06004803603602081101561040557600080fd5b5035610cc3565b34801561041857600080fd5b506102df610cdb565b34801561042d57600080fd5b506103e06004803603602081101561044457600080fd5b5035610ce1565b34801561045757600080fd5b506103e06004803603604081101561046e57600080fd5b8135919081019060408101602082013564010000000081111561049057600080fd5b8201836020820111156104a257600080fd5b803590602001918460018302840111640100000000831117156104c457600080fd5b91908080601f0160208091040260200160405190810160405281815292919060208401838380828437600092019190915250929550610da0945050505050565b6103e06004803603604081101561051a57600080fd5b5060ff81351690602001351515610dd7565b34801561053857600080fd5b506103e06004803603602081101561054f57600080fd5b5035610e14565b34801561056257600080fd5b506103e06004803603606081101561057957600080fd5b8135919081019060408101602082013564010000000081111561059b57600080fd5b8201836020820111156105ad57600080fd5b803590602001918460018302840111640100000000831117156105cf57600080fd5b91908080601f0160208091040260200160405190810160405281815292919060208401838380828437600092019190915250929594936020810193503591505064010000000081111561062157600080fd5b82018360208201111561063357600080fd5b8035906020019184600183028401116401000000008311171561065557600080fd5b91908080601f0160208091040260200160405190810160405281815292919060208401838380828437600092019190915250929550611296945050505050565b3480156106a157600080fd5b506103e061165a565b3480156106b657600080fd5b506102df6116c8565b3480156106cb57600080fd5b506106f2600480360360208110156106e257600080fd5b5035600160a060020a03166116ce565b604051901515815260200160405180910390f35b34801561071257600080fd5b506102df6116e9565b34801561072757600080fd5b506103e06004803603602081101561073e57600080fd5b50356116ef565b34801561075157600080fd5b506102df61175f565b34801561076657600080fd5b506102df611765565b34801561077b57600080fd5b506102df61176b565b34801561079057600080fd5b506106f2611771565b3480156107a557600080fd5b506102df61177b565b3480156107ba57600080fd5b506107d8600480360360208110156107d157600080fd5b5035611781565b60405180806020018c81526020018b81526020018a815260200189815260200188815260200187815260200186815260200185600581111561081657fe5b60ff16815260200184600281111561082a57fe5b60ff1681526020018315151515815260200182810382528d818151815260200191508051906020019080838360005b83811015610871578082015183820152602001610859565b50505050905090810190601f16801561089e5780820380516001836020036101000a031916815260200191505b509c5050505050505050505050505060405180910390f35b3480156108c257600080fd5b506103e0600480360360208110156108d957600080fd5b503561187c565b3480156108ec57600080fd5b506103e06118ec565b6103e06118f7565b34801561090957600080fd5b506102df6004803603602081101561092057600080fd5b5035600160a060020a0316611c92565b34801561093c57600080fd5b506103e0611ca6565b34801561095157600080fd5b506103e06004803603606081101561096857600080fd5b5080359060ff60208201351690604001351515611d19565b34801561098c57600080fd5b506103e0600480360360208110156109a357600080fd5b5035600160a060020a031661209e565b3480156109bf57600080fd5b506103e06120bb565b3480156109d457600080fd5b506103e06120d1565b3480156109e957600080fd5b506103e060048036036020811015610a0057600080fd5b5035612141565b348015610a1357600080fd5b506103e060048036036020811015610a2a57600080fd5b50356121b0565b348015610a3d57600080fd5b506102df61221f565b348015610a5257600080fd5b50610a5b612225565b604051600160a060020a03909116815260200160405180910390f35b348015610a8357600080fd5b506106f2612239565b348015610a9857600080fd5b5061033f61224f565b348015610aad57600080fd5b506102df61226c565b348015610ac257600080fd5b506103e060048036036020811015610ad957600080fd5b5035600160a060020a0316612272565b6103e06122a7565b348015610afd57600080fd5b506103e060048036036020811015610b1457600080fd5b50356122db565b348015610b2757600080fd5b506102df6123a2565b348015610b3c57600080fd5b506102df6123b7565b348015610b5157600080fd5b5061033f6123bd565b348015610b6657600080fd5b506102df6123da565b348015610b7b57600080fd5b506103e060048036036020811015610b9257600080fd5b50356123e0565b348015610ba557600080fd5b506103e060048036036020811015610bbc57600080fd5b5035600160a060020a03166123f8565b348015610bd857600080fd5b506103e060048036036020811015610bef57600080fd5b5035612414565b348015610c0257600080fd5b506102df612470565b600082820183811015610c1d57600080fd5b90505b92915050565b60125481565b60135460ff1681565b610180604051908101604052610149808252615778602083013981565b610c5a612239565b1515610c6557600080fd5b6101f4811115610cbe5760405160e560020a62461bcd02815260206004820152600760248201527f6d61782035302500000000000000000000000000000000000000000000000000604482015260640160405180910390fd5b600e55565b610ccb612239565b1515610cd657600080fd5b600a55565b60155481565b610ce9612239565b1515610cf457600080fd5b601454811115610d4d5760405160e560020a62461bcd02815260206004820152601760248201527f43616e742077697468647261772074686174206d756368000000000000000000604482015260640160405180910390fd5b601454610d60908263ffffffff61247616565b601455601054600160a060020a031681156108fc0282604051600060405180830381858888f19350505050158015610d9c573d6000803e3d6000fd5b5050565b610d9c82826000604051818152601f19601f8301168101602001604052908015610dd1576020820181803883390190505b50611296565b3360009081526019602052610df990349060409020549063ffffffff610c0b16565b33600090815260196020526040902055610d9c348383611d19565b600081815260166020526040812033600090815260098201602052909150604081206001810154909150600060ff90911611610e995760405160e560020a62461bcd02815260206004820152600c60248201527f6e6f74206120626574746f720000000000000000000000000000000000000000604482015260640160405180910390fd5b6004600883015460ff166005811115610eae57fe5b14158015610ecf57506005600883015460ff166005811115610ecc57fe5b14155b8015610eee57506000600883015460ff166005811115610eeb57fe5b14155b15610f17576008546006830154429101108015610f155760088301805460ff191660051790555b505b6005600883015460ff166005811115610f2c57fe5b1415610fe3578054600010610f8a5760405160e560020a62461bcd02815260206004820152601060248201527f416c726561647920726566756e64656400000000000000000000000000000000604482015260640160405180910390fd5b80543360009081526019602052610fad919060409020549063ffffffff610c0b16565b33600090815260196020526040902055336000908152600983016020526040902060008155600101805461ffff19169055611291565b60006064833360405160200180838054600181600116156101000203166002900480156110475780601f10611025576101008083540402835291820191611047565b820191906000526020600020905b815481529060010190602001808311611033575b505082600160a060020a0316600160a060020a03166c01000000000000000000000000028152601401925050506040516020818303038152906040528051906020012081151561109357fe5b60018481015460088701549390920601925060ff808216841092620100009004811615156101009092041615151481806110ca5750805b151561111f5760405160e560020a62461bcd02815260206004820152601460248201527f4e6f2077696e6e696e677320746f20636c61696d000000000000000000000000604482015260640160405180910390fd5b60008061115c86606060405190810160409081528254825260019092015460ff80821660208401526101009091041615159181019190915261248b565b905083156111c1576111be6111b1826111a589606060405190810160409081528254825260019092015460ff8082166020840152610100909104161515918101919091526124ad565b9063ffffffff61247616565b839063ffffffff610c0b16565b91505b82156111f8576111f56111b16103e86111e9600e548a600001546124da90919063ffffffff16565b9063ffffffff61250516565b91505b336000908152601960205261121a90839060409020549063ffffffff610c0b16565b3360009081526019602052604090205560155461123d908363ffffffff61247616565b601555601454611253908263ffffffff610c0b16565b60145533887f170c93f55dab07b2a0c856acad14fbb479fc97d6c3f489faf09b5f01c1575d3e8460405190815260200160405180910390a350505050505b505050565b61129e612529565b600160a060020a031633146112fc5760405160e560020a62461bcd02815260206004820152600b60248201527f61757468206661696c6564000000000000000000000000000000000000000000604482015260640160405180910390fd5b60008381526018602052604090205460ff161561134d5760405160e560020a62461bcd0281526004018080602001828103825260218152602001806158c16021913960400191505060405180910390fd5b6000838152601860205260019060409020805460ff19169115159190911790556000838152601760205260408120549050600081116113d55760405160e560020a62461bcd02815260206004820152601060248201527f496e76616c6964205f7175657279496400000000000000000000000000000000604482015260640160405180910390fd5b60008181526016602052604081209050806001015485141561149c576113fc8585856126fb565b60ff16156114455760088101805460ff19166005179055817f796cfdf016e01ffdff99dbe8ce27df476b06d0fae2ac92ec3889b6008219ed5260405160405180910390a2611497565b80848051611457929160200190615508565b50600881015460ff16600581111561146b57fe5b600101600581111561147957fe5b60088201805460ff1916600183600581111561149157fe5b02179055505b611536565b80600201548514156114c9576114b184612812565b6004820155600881015460ff16600581111561146b57fe5b80600301548514156114f9576114de84612812565b600580830191909155600882015460ff169081111561146b57fe5b60088101805460ff19166005179055817f796cfdf016e01ffdff99dbe8ce27df476b06d0fae2ac92ec3889b6008219ed5260405160405180910390a25b6004600882015460ff16600581111561154b57fe5b1415611653576005810154600482015460088301805462ff00001916929091106201000002919091179055600781015460155461158791610c0b565b6015556004810154600582015483917f8cb1b96dbae53a90404e7c18d5482ddc5a99bb8ca665731bc035a9054f098faa918491906040516020810183905260408101829052606080825284546002600019610100600184161502019091160490820181905281906080820190869080156116425780601f1061161757610100808354040283529160200191611642565b820191906000526020600020905b81548152906001019060200180831161162557829003601f168201915b505094505050505060405180910390a25b5050505050565b611663336116ce565b151561166e57600080fd5b60065460ff16151561167f57600080fd5b6006805460ff191690557f5db9ee0a495bf2e6ff9c91a7834c1ba4fdd244a5e8aa4e537bd38aeae4b073aa33604051600160a060020a03909116815260200160405180910390a1565b600e5481565b60006116e160058363ffffffff6128f916565b90505b919050565b60115481565b6116f7612239565b151561170257600080fd5b603c81101561175a5760405160e560020a62461bcd02815260206004820152601360248201527f4d696e696d756d2069732031206d696e75746500000000000000000000000000604482015260640160405180910390fd5b600755565b60085481565b600a5481565b600f5481565b60065460ff165b90565b600d5481565b601660205280600052604060002080549091508190600260001961010060018416150201909116046020601f820181900481020160405190810160405280929190818152602001828054600181600116156101000203166002900480156118295780601f106117fe57610100808354040283529160200191611829565b820191906000526020600020905b81548152906001019060200180831161180c57829003601f168201915b50505060018401546002850154600386015460048701546005880154600689015460078a01546008909a0154989995989497509295509093909260ff80821691610100810482169162010000909104168b565b611884612239565b151561188f57600080fd5b60328111156118e75760405160e560020a62461bcd02815260206004820152600660248201527f6d61782035250000000000000000000000000000000000000000000000000000604482015260640160405180910390fd5b600d55565b6118f533612933565b565b60065460ff161561190757600080fd5b60075460125442910111156119655760405160e560020a62461bcd02815260206004820152601460248201527f726f6c6c20697320636f6f6c696e6720646f776e000000000000000000000000604482015260640160405180910390fd5b600061196f61297d565b9050348111156119eb577f0510c169028b6c6e93e31f93b1b642b756cea80cc1966887b5c8f0b93d3d5b688160405190815260200160405180910390a160003411156119e657333480156108fc0290604051600060405180830381858888f193505050501580156119e4573d6000803e3d6000fd5b505b611c8f565b601154600090815260166020526040812090506060600060135460ff166002811115611a1357fe5b1415611a3a5761018060405190810160405261014980825261577860208301399050611a91565b600160135460ff166002811115611a4d57fe5b1415611a74576101806040519081016040526101498082526155e060208301399050611a91565b610180604051908101604052610149808252615948602083013990505b611aba7f1100000000000000000000000000000000000000000000000000000000000000612a6c565b611afd600060408051908101604052600681527f6e6573746564000000000000000000000000000000000000000000000000000060208201526009548490612c43565b6002830155600f54611b479060408051908101604052600681527f6e6573746564000000000000000000000000000000000000000000000000000060208201526009548490612c43565b6003830155611b757f3000000000000000000000000000000000000000000000000000000000000000612a6c565b611b8460006001600a5461300f565b60018381019190915542600684015560088301805460ff191690911780825560135460ff16919061ff001916610100836002811115611bbf57fe5b02179055506011546001830154600090815260176020526040902055601154600283015460009081526017602052604090205560115460038301546000908152601760205260409020556011547fe86c6fb9a82a2c1fd4c6c98a2fe03844f019514b0cfe8dfbdf076c61f124c0d460405160405180910390a2611c406134eb565b6000611c52348563ffffffff61247616565b90506000811115611c8b573381156108fc0282604051600060405180830381858888f19350505050158015611653573d6000803e3d6000fd5b5050505b50565b601960205280600052604060002054905081565b611cae612239565b1515611cb957600080fd5b6006546000906101009004600160a060020a03167f8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e060405160405180910390a36006805474ffffffffffffffffffffffffffffffffffffffff0019169055565b60065460ff1615611d2957600080fd5b3360009081526019602052839060409020541015611d7b5760405160e560020a62461bcd0281526004018080602001828103825260268152602001806158e26026913960400191505060405180910390fd5b60018260ff16118015611d91575060648260ff16105b1515611dd15760405160e560020a62461bcd02815260040180806020018281038252602e815260200180615729602e913960400191505060405180910390fd5b600c54831115611e2a5760405160e560020a62461bcd02815260206004820152600c60248201527f42657420746f6f20686967680000000000000000000000000000000000000000604482015260640160405180910390fd5b600b54831015611e835760405160e560020a62461bcd02815260206004820152600b60248201527f42657420746f6f206c6f77000000000000000000000000000000000000000000604482015260640160405180910390fd5b60115460009081526016602052604081203360009081526009820160205290915060408120805490915015611f015760405160e560020a62461bcd02815260206004820152601460248201527f416c726561647920706c61636564206120626574000000000000000000000000604482015260640160405180910390fd5b84815560018101805460ff191660ff86161761ff001916610100851515021790556000611f5e82606060405190810160409081528254825260019092015460ff80821660208401526101009091041615159181019190915261357e565b90506000611f816103e86111e9600e5486600001546124da90919063ffffffff16565b9050611f93828263ffffffff610c0b16565b6015541015611feb5760405160e560020a62461bcd02815260206004820152601f60248201527f4e6f7420656e6f75676820696e20706f6f6c20666f7220746869732062657400604482015260640160405180910390fd5b336000908152601960205261200d90889060409020549063ffffffff61247616565b33600090815260196020526040902055825460078501546120339163ffffffff610c0b16565b600785015560115433907fa03a715f8e6fbcac20e92a38c1169f81b49b105441690f1e78ed31b17aafba9f89898961206c57600061206f565b60015b60405192835260ff9182166020840152166040808301919091526060909101905180910390a350505050505050565b6120a7336116ce565b15156120b257600080fd5b611c8f816135c3565b6120c3612239565b15156120ce57600080fd5b33ff5b6120da336116ce565b15156120e557600080fd5b60065460ff16156120f557600080fd5b6006805460ff191660011790557f62e78cea01bee320cd4e420270b5ea74000d11b0c9f74754ebdbfc544b05a25833604051600160a060020a03909116815260200160405180910390a1565b612149612239565b151561215457600080fd5b600081116121ab5760405160e560020a62461bcd02815260206004820152601960248201527f4d7573742062652067726561746572207468616e207a65726f00000000000000604482015260640160405180910390fd5b600c55565b6121b8612239565b15156121c357600080fd5b6000811161221a5760405160e560020a62461bcd02815260206004820152601960248201527f4d7573742062652067726561746572207468616e207a65726f00000000000000604482015260640160405180910390fd5b600b55565b600c5481565b6006546101009004600160a060020a031690565b6006546101009004600160a060020a0316331490565b610180604051908101604052610149808252615948602083013981565b60075481565b61227a612239565b151561228557600080fd5b60108054600160a060020a031916600160a060020a0392909216919091179055565b33600090815260196020526122c990349060409020549063ffffffff610c0b16565b33600090815260196020526040902055565b33600090815260196020528190604090205410156123425760405160e560020a62461bcd02815260206004820152601060248201527f4e6f7420656e6f7567682066756e647300000000000000000000000000000000604482015260640160405180910390fd5b336000908152601960205261236490829060409020549063ffffffff61247616565b336000908152601960205260409020553381156108fc0282604051600060405180830381858888f19350505050158015610d9c573d6000803e3d6000fd5b33600090815260196020526040812054905090565b60095481565b6101806040519081016040526101498082526155e0602083013981565b600b5481565b6123e8612239565b15156123f357600080fd5b600955565b612400612239565b151561240b57600080fd5b611c8f8161360d565b61241c612239565b151561242757600080fd5b624f1a00811061246b5760405160e560020a62461bcd0281526004018080602001828103825260218152602001806157576021913960400191505060405180910390fd5b600f55565b60145481565b60008282111561248557600080fd5b50900390565b6000600d546103e861249c846124ad565b8115156124a557fe5b040292915050565b6000806001836020015160ff1603905082518160648190038551028115156124d157fe5b04019392505050565b60008215156124eb57506000610c20565b8282028284828115156124fa57fe5b0414610c1d57600080fd5b600080821161251357600080fd5b6000828481151561252057fe5b04949350505050565b600154600090600160a060020a03161580612556575060015461255490600160a060020a0316613698565b155b1561256757612565600061369c565b505b600154600160a060020a03166338cc48316040518163ffffffff1660e060020a028152600401602060405180830381600087803b1580156125a757600080fd5b505af11580156125bb573d6000803e3d6000fd5b505050506040513d60208110156125d157600080fd5b810190808051600054600160a060020a03908116911614925061268591505057600154600160a060020a03166338cc48316040518163ffffffff1660e060020a028152600401602060405180830381600087803b15801561263157600080fd5b505af1158015612645573d6000803e3d6000fd5b505050506040513d602081101561265b57600080fd5b81019080805160008054600160a060020a031916600160a060020a03929092169190911790555050505b600054600160a060020a031663c281d19e6040518163ffffffff1660e060020a02815260040160206040518083038186803b1580156126c357600080fd5b505afa1580156126d7573d6000803e3d6000fd5b505050506040513d60208110156126ed57600080fd5b810190808051935050505090565b60008160008151811061270a57fe5b016020015160f860020a900460f860020a02600160f860020a0319167f4c0000000000000000000000000000000000000000000000000000000000000014158061279c57508160018151811061275c57fe5b016020015160f860020a900460f860020a02600160f860020a0319167f500000000000000000000000000000000000000000000000000000000000000014155b806127d057506001826002815181106127b157fe5b016020015160f860020a900460f860020a0260f860020a900460ff1614155b156127dd5750600161280b565b60006127f28386866127ed6136a6565b613748565b905080151561280557600291505061280b565b60009150505b9392505050565b600060028183815b81518110156128d8578215612830576001840393505b81818151811061283c57fe5b016020015160f860020a900460f860020a0260f860020a900460ff16602e141561286557600192505b600082828151811061287357fe5b016020015160f860020a900460f860020a0260f860020a900460ff169050603081101580156128a3575060398111155b156128b4576030810386600a020195505b8380156128bf575084155b156128cf57506116e49350505050565b5060010161281a565b5b83156128f05784600a0294506001840393506128d9565b50505050919050565b6000600160a060020a038216151561291057600080fd5b600160a060020a03821660009081526020849052604090205460ff169392505050565b61294460058263ffffffff613d8216565b80600160a060020a03167fcd265ebaf09df2871cc7bd4133404a235ba12eff2041bb89d9c714a2621c7c7e60405160405180910390a250565b60006129a87f1100000000000000000000000000000000000000000000000000000000000000612a6c565b60006129e960408051908101604052600381527f55524c00000000000000000000000000000000000000000000000000000000006020820152600954613dd7565b6002029050612a177f3000000000000000000000000000000000000000000000000000000000000000612a6c565b612a66612a5960408051908101604052600681527f52616e646f6d00000000000000000000000000000000000000000000000000006020820152600a54613dd7565b829063ffffffff610c0b16565b91505090565b600154600160a060020a03161580612a965750600154612a9490600160a060020a0316613698565b155b15612aa757612aa5600061369c565b505b600154600160a060020a03166338cc48316040518163ffffffff1660e060020a028152600401602060405180830381600087803b158015612ae757600080fd5b505af1158015612afb573d6000803e3d6000fd5b505050506040513d6020811015612b1157600080fd5b810190808051600054600160a060020a039081169116149250612bc591505057600154600160a060020a03166338cc48316040518163ffffffff1660e060020a028152600401602060405180830381600087803b158015612b7157600080fd5b505af1158015612b85573d6000803e3d6000fd5b505050506040513d6020811015612b9b57600080fd5b81019080805160008054600160a060020a031916600160a060020a03929092169190911790555050505b600054600160a060020a031663688dcfd78260405160e060020a63ffffffff84160281527fff000000000000000000000000000000000000000000000000000000000000009091166004820152602401600060405180830381600087803b158015612c2f57600080fd5b505af1158015611653573d6000803e3d6000fd5b600154600090600160a060020a03161580612c705750600154612c6e90600160a060020a0316613698565b155b15612c8157612c7f600061369c565b505b600154600160a060020a03166338cc48316040518163ffffffff1660e060020a028152600401602060405180830381600087803b158015612cc157600080fd5b505af1158015612cd5573d6000803e3d6000fd5b505050506040513d6020811015612ceb57600080fd5b810190808051600054600160a060020a039081169116149250612d9f91505057600154600160a060020a03166338cc48316040518163ffffffff1660e060020a028152600401602060405180830381600087803b158015612d4b57600080fd5b505af1158015612d5f573d6000803e3d6000fd5b505050506040513d6020811015612d7557600080fd5b81019080805160008054600160a060020a031916600160a060020a03929092169190911790555050505b60008054600160a060020a0316632ef3accc86856040518363ffffffff1660e060020a0281526004018080602001838152602001828103825284818151815260200191508051906020019080838360005b83811015612e08578082015183820152602001612df0565b50505050905090810190601f168015612e355780820380516001836020036101000a031916815260200191505b509350505050602060405180830381600087803b158015612e5557600080fd5b505af1158015612e69573d6000803e3d6000fd5b505050506040513d6020811015612e7f57600080fd5b8101908080519350505050670de0b6b3a76400003a840201811115612ea8575060009050613007565b600054600160a060020a031663c51be90f82888888886040518663ffffffff1660e060020a028152600401808581526020018060200180602001848152602001838103835286818151815260200191508051906020019080838360005b83811015612f1d578082015183820152602001612f05565b50505050905090810190601f168015612f4a5780820380516001836020036101000a031916815260200191505b50838103825285818151815260200191508051906020019080838360005b83811015612f80578082015183820152602001612f68565b50505050905090810190601f168015612fad5780820380516001836020036101000a031916815260200191505b5096505050505050506020604051808303818588803b158015612fcf57600080fd5b505af1158015612fe3573d6000803e3d6000fd5b50505050506040513d6020811015612ffa57600080fd5b8101908080519450505050505b949350505050565b60008083118015613021575060208311155b151561302c57600080fd5b600a8402935060606001604051818152601f19601f8301168101602001604052908015613060576020820181803883390190505b5090508360f860020a028160008151811061307757fe5b906020010190600160f860020a031916908160001a90535060606020604051818152601f19601f83011681016020016040529080156130bd576020820181803883390190505b50905060606020604051818152601f19601f83011681016020016040529080156130ee576020820181803883390190505b50905060006130fb614022565b90506020835242411860014303401860208401526020825280602083015260606020604051818152601f19601f8301168101602001604052908015613147576020820181803883390190505b50905088602082015260606008604051818152601f19601f830116810160200160405290801561317e576020820181803883390190505b50905061319182601860088460006141bc565b5061319a615586565b60806040519081016040528087815260200188815260200186815260200184815250905060006131fe60408051908101604052600681527f72616e646f6d00000000000000000000000000000000000000000000000000006020820152838c614206565b905060606008604051818152601f19601f830116810160200160405290801561322e576020820181803883390190505b509050602084015160f860020a810460278301537e01000000000000000000000000000000000000000000000000000000000000810460268301537d0100000000000000000000000000000000000000000000000000000000008104602583015360e060020a810460248301537b01000000000000000000000000000000000000000000000000000000810460238301537a0100000000000000000000000000000000000000000000000000008104602283015379010000000000000000000000000000000000000000000000000081046021830153780100000000000000000000000000000000000000000000000081046020830153506134db82826020860151600287516040518082805190602001908083835b602083106133635780518252601f199092019160209182019101613344565b6001836020036101000a038019825116818451168082178552505050505050905001915050602060405180830381855afa1580156133a5573d6000803e3d6000fd5b5050506040513d60208110156133ba57600080fd5b8101908080519250505060408801516040516020018085805190602001908083835b602083106133fb5780518252601f1990920191602091820191016133dc565b6001836020036101000a038019825116818451161790925250505091909101905084805190602001908083835b602083106134475780518252601f199092019160209182019101613428565b6001836020036101000a038019825116818451161790925250505091909101848152602001905082805190602001908083835b602083106134995780518252601f19909201916020918201910161347a565b6001836020036101000a038019825116818451161790925250505091909101955060409450505050505160208183030381529060405280519060200120614422565b509b9a5050505050505050505050565b6011546134ff90600163ffffffff610c0b16565b60115560135460039060ff16600281111561351657fe5b60010181151561352257fe5b06600281111561352e57fe5b6013805460ff1916600183600281111561354457fe5b0217905550426012556011547fb4267207f21a1858b0724dfba74816419152c88804d2b6276e5c32f6cd47c74a60405160405180910390a2565b6000806001836020015160ff16039050600d546103e8036103e884600001518360648190038751028115156135af57fe5b04018115156135ba57fe5b04029392505050565b6135d460058263ffffffff61443716565b80600160a060020a03167f6719d08c1888103bea251a4ed56406bd0c3e69723c8a1686e017e7bbe159b6f860405160405180910390a250565b600160a060020a038116151561362257600080fd5b600654600160a060020a03808316916101009004167f8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e060405160405180910390a360068054600160a060020a039092166101000274ffffffffffffffffffffffffffffffffffffffff0019909216919091179055565b3b90565b60006116e161448e565b606060028054600181600116156101000203166002900480601f01602080910402602001604051908101604052809291908181526020018280546001816001161561010002031660029004801561373e5780601f106137135761010080835404028352916020019161373e565b820191906000526020600020905b81548152906001019060200180831161372157829003601f168201915b5050505050905090565b6000808560458151811061375857fe5b016020015160f860020a900460f860020a0260f860020a900460ff16600201604401602001905060606020604051818152601f19601f83011681016020016040529080156137ad576020820181803883390190505b5090506137bf878360208460006141bc565b50600284876040516020018083805190602001908083835b602083106137f65780518252601f1990920191602091820191016137d7565b6001836020036101000a038019825116818451161790925250505091909101928352505060200190506040516020818303038152906040526040518082805190602001908083835b6020831061385d5780518252601f19909201916020918201910161383e565b6001836020036101000a038019825116818451168082178552505050505050905001915050602060405180830381855afa15801561389f573d6000803e3d6000fd5b5050506040513d60208110156138b457600080fd5b810190808051925060409150505160200180828152602001915050604051602081830303815290604052805190602001208180519060200120146138fd57600092505050613007565b606087604a84018151811061390e57fe5b016020015160f860020a900460f860020a0260f860020a900460ff16600201604051818152601f19601f8301168101602001604052908015613957576020820181803883390190505b50905061396c888460490183518460006141bc565b50613a366002826040518082805190602001908083835b602083106139a25780518252601f199092019160209182019101613983565b6001836020036101000a038019825116818451168082178552505050505050905001915050602060405180830381855afa1580156139e4573d6000803e3d6000fd5b5050506040513d60208110156139f957600080fd5b81019080805192508991508b90506028870181518110613a1557fe5b016020015160f860020a900460f860020a0260f860020a900460ff166147ac565b1515613a485760009350505050613007565b60606029604051818152601f19601f8301168101602001604052908015613a76576020820181803883390190505b509050613a8b898560200160298460006141bc565b50606060408051818152601f19601f8301168101602001604052908015613ab9576020820181803883390190505b5090506000835186602001602901016041019050613adf8b6040830360408560006141bc565b5060006002836040518082805190602001908083835b60208310613b145780518252601f199092019160209182019101613af5565b6001836020036101000a038019825116818451168082178552505050505050905001915050602060405180830381855afa158015613b56573d6000803e3d6000fd5b5050506040513d6020811015613b6b57600080fd5b810190808051935086925083915060409050516020018083805190602001908083835b60208310613bad5780518252601f199092019160209182019101613b8e565b6001836020036101000a038019825116818451161790925250505091909101928352505060200190506040516020818303038152906040528051906020012060008c8152600360205260409020541415613c185760008b815260036020526040902060009055613c28565b6000975050505050505050613007565b60606049604051818152601f19601f8301168101602001604052908015613c56576020820181803883390190505b509050613c688d8960498460006141bc565b50613d086002826040518082805190602001908083835b60208310613c9e5780518252601f199092019160209182019101613c7f565b6001836020036101000a038019825116818451168082178552505050505050905001915050602060405180830381855afa158015613ce0573d6000803e3d6000fd5b5050506040513d6020811015613cf557600080fd5b810190808051925089915087905061482c565b1515613d1f57600098505050505050505050613007565b60008281526004602052604090205460ff161515613d5f57613d418d846149ae565b6000838152600460205260409020805460ff19169115159190911790555b60008281526004602052604090205460ff169d9c50505050505050505050505050565b600160a060020a0381161515613d9757600080fd5b613da182826128f9565b1515613dac57600080fd5b600160a060020a0381166000908152602083905260408120805460ff19169115159190911790555050565b600154600090600160a060020a03161580613e045750600154613e0290600160a060020a0316613698565b155b15613e1557613e13600061369c565b505b600154600160a060020a03166338cc48316040518163ffffffff1660e060020a028152600401602060405180830381600087803b158015613e5557600080fd5b505af1158015613e69573d6000803e3d6000fd5b505050506040513d6020811015613e7f57600080fd5b810190808051600054600160a060020a039081169116149250613f3391505057600154600160a060020a03166338cc48316040518163ffffffff1660e060020a028152600401602060405180830381600087803b158015613edf57600080fd5b505af1158015613ef3573d6000803e3d6000fd5b505050506040513d6020811015613f0957600080fd5b81019080805160008054600160a060020a031916600160a060020a03929092169190911790555050505b600054600160a060020a0316632ef3accc84846040518363ffffffff1660e060020a0281526004018080602001838152602001828103825284818151815260200191508051906020019080838360005b83811015613f9b578082015183820152602001613f83565b50505050905090810190601f168015613fc85780820380516001836020036101000a031916815260200191505b509350505050602060405180830381600087803b158015613fe857600080fd5b505af1158015613ffc573d6000803e3d6000fd5b505050506040513d602081101561401257600080fd5b8101908080519695505050505050565b600154600090600160a060020a0316158061404f575060015461404d90600160a060020a0316613698565b155b156140605761405e600061369c565b505b600154600160a060020a03166338cc48316040518163ffffffff1660e060020a028152600401602060405180830381600087803b1580156140a057600080fd5b505af11580156140b4573d6000803e3d6000fd5b505050506040513d60208110156140ca57600080fd5b810190808051600054600160a060020a03908116911614925061417e91505057600154600160a060020a03166338cc48316040518163ffffffff1660e060020a028152600401602060405180830381600087803b15801561412a57600080fd5b505af115801561413e573d6000803e3d6000fd5b505050506040513d602081101561415457600080fd5b81019080805160008054600160a060020a031916600160a060020a03929092169190911790555050505b600054600160a060020a031663abaa5f3e6040518163ffffffff1660e060020a02815260040160206040518083038186803b1580156126c357600080fd5b606083820180845110156141cf57600080fd5b60208087019084015b8688602001018210156141f9578189015186820152602091820191016141d8565b5093979650505050505050565b600154600090600160a060020a03161580614233575060015461423190600160a060020a0316613698565b155b1561424457614242600061369c565b505b600154600160a060020a03166338cc48316040518163ffffffff1660e060020a028152600401602060405180830381600087803b15801561428457600080fd5b505af1158015614298573d6000803e3d6000fd5b505050506040513d60208110156142ae57600080fd5b810190808051600054600160a060020a03908116911614925061436291505057600154600160a060020a03166338cc48316040518163ffffffff1660e060020a028152600401602060405180830381600087803b15801561430e57600080fd5b505af1158015614322573d6000803e3d6000fd5b505050506040513d602081101561433857600080fd5b81019080805160008054600160a060020a031916600160a060020a03929092169190911790555050505b6060600460405190808252806020026020018201604052801561439957816020015b60608152602001906001900390816143845790505b5090508351816000815181106143ab57fe5b602090810290910101528360016020020151816001815181106143ca57fe5b60209081029190910101526040840151816002815181106143e757fe5b602090810291909101015260608401518160038151811061440457fe5b60209081029091010152614419858285614d8d565b95945050505050565b60008281526003602052819060409020555050565b600160a060020a038116151561444c57600080fd5b61445682826128f9565b1561446057600080fd5b600160a060020a0381166000908152602083905260019060409020805460ff19169115159190911790555050565b6000806144ae731d3b2638a7cc9f2cb3d298a3da7a90b67e5506ed613698565b111561451e5760018054600160a060020a031916731d3b2638a7cc9f2cb3d298a3da7a90b67e5506ed17905561451660408051908101604052600b81527f6574685f6d61696e6e6574000000000000000000000000000000000000000000602082015261516a565b506001611778565b600061453d73c03a2615d5efaf5f49f60b7bb6583eaec212fdf1613698565b11156145a55760018054600160a060020a03191673c03a2615d5efaf5f49f60b7bb6583eaec212fdf117905561451660408051908101604052600c81527f6574685f726f707374656e330000000000000000000000000000000000000000602082015261516a565b60006145c473b7a07bcf2ba2f2703b24c0691b5278999c59ac7e613698565b111561462c5760018054600160a060020a03191673b7a07bcf2ba2f2703b24c0691b5278999c59ac7e17905561451660408051908101604052600981527f6574685f6b6f76616e0000000000000000000000000000000000000000000000602082015261516a565b600061464b73146500cfd35b22e4a392fe0adc06de1a1368ed48613698565b11156146b35760018054600160a060020a03191673146500cfd35b22e4a392fe0adc06de1a1368ed4817905561451660408051908101604052600b81527f6574685f72696e6b656279000000000000000000000000000000000000000000602082015261516a565b60006146d2736f485c8bf6fc43ea212e93bbf8ce046c7f1cb475613698565b1115614704575060018054600160a060020a031916736f485c8bf6fc43ea212e93bbf8ce046c7f1cb475178155611778565b60006147237320e12a1f859b3feae5fb2a0a32c18f5a65555bbf613698565b1115614755575060018054600160a060020a0319167320e12a1f859b3feae5fb2a0a32c18f5a65555bbf178155611778565b60006147747351efaf4c8b3c9afbd5ab9f4bbc82784ab6ef8faa613698565b11156147a6575060018054600160a060020a0319167351efaf4c8b3c9afbd5ab9f4bbc82784ab6ef8faa178155611778565b50600090565b60006001828451146147bd57600080fd5b60005b83811015614823578481815181106147d457fe5b016020015160f860020a900460f860020a02600160f860020a03191686826020811015156147fe57fe5b1a60f860020a02600160f860020a03191614151561481b57600091505b6001016147c0565b50949350505050565b600080600080600060606020604051818152601f19601f8301168101602001604052908015614862576020820181803883390190505b509050600060208960038151811061487657fe5b016020015160f860020a900460f860020a0260f860020a900460ff160360040190506148a7898260208560006141bc565b915060606020604051818152601f19601f83011681016020016040529080156148d7576020820181803883390190505b50905060228201915061491e8a60208c60018603815181106148f557fe5b016020015160f860020a900460f860020a0260f860020a900460ff1603840160208460006141bc565b9050602083015194506020810151935061493b8b601b878761517d565b9097509550600160a060020a0386168980519060200120600160a060020a0316141561497157600197505050505050505061280b565b61497e8b601c878761517d565b9097509550600160a060020a0386168980519060200120600160a060020a031614975061280b9650505050505050565b60008060608484600101815181106149c257fe5b016020015160f860020a900460f860020a0260f860020a900460ff16600201604051818152601f19601f8301168101602001604052908015614a0b576020820181803883390190505b509050614a1d858583518460006141bc565b50606060408051818152601f19601f8301168101602001604052908015614a4b576020820181803883390190505b509050614a5e86600460408460006141bc565b5060606062604051818152601f19601f8301168101602001604052908015614a8d576020820181803883390190505b50905060f860020a81600081518110614aa257fe5b906020010190600160f860020a031916908160001a905350614acc876041880360418460016141bc565b5060606040805190810160405280602081526020017ffd94fa71bc0ba10d39d464d0d8f465efeef0a2764e3887fcc9df41ded20f505c8152509050614b1781600060208560426141bc565b50614bb76002836040518082805190602001908083835b60208310614b4d5780518252601f199092019160209182019101614b2e565b6001836020036101000a038019825116818451168082178552505050505050905001915050602060405180830381855afa158015614b8f573d6000803e3d6000fd5b5050506040513d6020811015614ba457600080fd5b810190808051925087915086905061482c565b9450841515614bce57600095505050505050610c20565b606080604051908101604052806040815260200161590860409139905060606042604051818152601f19601f8301168101602001604052908015614c19576020820181803883390190505b5090507ffe0000000000000000000000000000000000000000000000000000000000000081600081518110614c4a57fe5b906020010190600160f860020a031916908160001a905350614c728a600360418460016141bc565b5060608a604581518110614c8257fe5b016020015160f860020a900460f860020a0260f860020a900460ff16600201604051818152601f19601f8301168101602001604052908015614ccb576020820181803883390190505b509050614cde8b604483518460006141bc565b50614d7e6002836040518082805190602001908083835b60208310614d145780518252601f199092019160209182019101614cf5565b6001836020036101000a038019825116818451168082178552505050505050905001915050602060405180830381855afa158015614d56573d6000803e3d6000fd5b5050506040513d6020811015614d6b57600080fd5b810190808051925084915086905061482c565b9b9a5050505050505050505050565b600154600090600160a060020a03161580614dba5750600154614db890600160a060020a0316613698565b155b15614dcb57614dc9600061369c565b505b600154600160a060020a03166338cc48316040518163ffffffff1660e060020a028152600401602060405180830381600087803b158015614e0b57600080fd5b505af1158015614e1f573d6000803e3d6000fd5b505050506040513d6020811015614e3557600080fd5b810190808051600054600160a060020a039081169116149250614ee991505057600154600160a060020a03166338cc48316040518163ffffffff1660e060020a028152600401602060405180830381600087803b158015614e9557600080fd5b505af1158015614ea9573d6000803e3d6000fd5b505050506040513d6020811015614ebf57600080fd5b81019080805160008054600160a060020a031916600160a060020a03929092169190911790555050505b60008054600160a060020a0316632ef3accc86856040518363ffffffff1660e060020a0281526004018080602001838152602001828103825284818151815260200191508051906020019080838360005b83811015614f52578082015183820152602001614f3a565b50505050905090810190601f168015614f7f5780820380516001836020036101000a031916815260200191505b509350505050602060405180830381600087803b158015614f9f57600080fd5b505af1158015614fb3573d6000803e3d6000fd5b505050506040513d6020811015614fc957600080fd5b8101908080519350505050670de0b6b3a76400003a840201811115614ff257506000905061280b565b6060614ffd856151bd565b60008054919250600160a060020a039091169063c55c1cb69084908985896040518663ffffffff1660e060020a028152600401808581526020018060200180602001848152602001838103835286818151815260200191508051906020019080838360005b8381101561507a578082015183820152602001615062565b50505050905090810190601f1680156150a75780820380516001836020036101000a031916815260200191505b50838103825285818151815260200191508051906020019080838360005b838110156150dd5780820151838201526020016150c5565b50505050905090810190601f16801561510a5780820380516001836020036101000a031916815260200191505b5096505050505050506020604051808303818588803b15801561512c57600080fd5b505af1158015615140573d6000803e3d6000fd5b50505050506040513d602081101561515757600080fd5b8101908080519998505050505050505050565b6002818051610d9c929160200190615508565b60008060008060405188815287602082015286604082015285606082015260208160808360006001610bb8f1925080519299929850919650505050505050565b60606151c7615231565b6151cf6155ad565b6151db8161040061523d565b6151e48161526b565b60005b835181101561521e576152168482815181106151ff57fe5b90602001906020020151839063ffffffff61527616565b6001016151e7565b5061522881615293565b80519392505050565b60405180590338823950565b8060208106156152505760208106602003015b60208301819052604051928390526000835290910160405250565b611c8f81600461529a565b61528382600283516152b3565b611291828263ffffffff6153b116565b611c8f8160075b610d9c82601f602060ff8516021763ffffffff61544a16565b601781116152d4576152cf8360ff84811660200216831761544a565b611291565b60ff811161530d576152f5836018602060ff8616021763ffffffff61544a16565b6153078382600163ffffffff61548316565b50611291565b61ffff81116153415761532f836019602060ff8616021763ffffffff61544a16565b6153078382600263ffffffff61548316565b63ffffffff81116153775761536583601a602060ff8616021763ffffffff61544a16565b6153078382600463ffffffff61548316565b67ffffffffffffffff81116112915761539f83601b602060ff8616021763ffffffff61544a16565b611c8b8382600863ffffffff61548316565b6153b96155ad565b826020015183515183510111156153e5576153e5836153dd856020015185516154d7565b6002026154ee565b60008060008451905085518051602081830101945086510190526020850191505b6020811061542657815183526020928301929190910190601f1901615406565b60001960208290036101000a01801983511681855116179093525093949350505050565b8160200151825151600101111561546c5761546c8283602001516002026154ee565b815180516020818301018381535060010190525050565b61548b6155ad565b8360200151845151830111156154ad576154ad846153dd8660200151856154d7565b60001961010083900a01845180518481830101868419825116179052909301909252509192915050565b6000818311156154e8575081610c20565b50919050565b6060825190506154fe838361523d565b611c8b83826153b1565b828054600181600116156101000203166002900490600052602060002090601f016020900481019282601f1061554957805160ff1916838001178555615576565b82800160010185558215615576579182015b8281111561557657825182559160200191906001019061555b565b506155829291506155c5565b5090565b60806040519081016040526004815b60608152602001906001900390816155955790505090565b60408051908101604052606081526000602082015290565b61177891905b8082111561558257600081556001016155cb56fe5b55524c5d206a736f6e2868747470733a2f2f6d696e2d6170692e63727970746f636f6d706172652e636f6d2f646174612f70726963653f6673796d3d425443267473796d733d555344266578747261506172616d733d5072696365526f6c6c267369676e3d74727565266170695f6b65793d247b5b646563727970745d20424a45576f35613533415042724e346659707a35784a61447a506d434c4e6a4b64552b794d654433703656734d4c6b4652466671497652612b64342f71756b54426273465a716b7673744d4d63716f4c5a6153686f683448664839585155784c376341744b777541693847436b467073306b63466d4e4233454151517667474d5834466561616f683430594370357142644b6758714c68582b4256753478397030754b53395858422b43633271496c7661676b4737792b546f3162567270315867673d3d7d292e55534445787065637465642076616c7565206d75737420626520696e207468652072616e6765206f66203220746f2039394f7261636c697a652064656c6179206973206d6178696d756d20363020646179735b55524c5d206a736f6e2868747470733a2f2f6d696e2d6170692e63727970746f636f6d706172652e636f6d2f646174612f70726963653f6673796d3d455448267473796d733d555344266578747261506172616d733d5072696365526f6c6c267369676e3d74727565266170695f6b65793d247b5b646563727970745d20424a45576f35613533415042724e346659707a35784a61447a506d434c4e6a4b64552b794d654433703656734d4c6b4652466671497652612b64342f71756b54426273465a716b7673744d4d63716f4c5a6153686f683448664839585155784c376341744b777541693847436b467073306b63466d4e4233454151517667474d5834466561616f683430594370357142644b6758714c68582b4256753478397030754b53395858422b43633271496c7661676b4737792b546f3162567270315867673d3d7d292e55534451756572792068617320616c7265616479206265656e2070726f636573736564214e6f7420656e6f75676820746f20626574207468652073706563696669656420616d6f756e747fb956469c5c9b89840d55b43537e66a98dd4811ea0a27224272c2e5622911e8537a2f8e86a46baec82864e98dd01e9ccc2f8bc5dfc9cbe5a91a290498dd96e45b55524c5d206a736f6e2868747470733a2f2f6d696e2d6170692e63727970746f636f6d706172652e636f6d2f646174612f70726963653f6673796d3d4c5443267473796d733d555344266578747261506172616d733d5072696365526f6c6c267369676e3d74727565266170695f6b65793d247b5b646563727970745d20424a45576f35613533415042724e346659707a35784a61447a506d434c4e6a4b64552b794d654433703656734d4c6b4652466671497652612b64342f71756b54426273465a716b7673744d4d63716f4c5a6153686f683448664839585155784c376341744b777541693847436b467073306b63466d4e4233454151517667474d5834466561616f683430594370357142644b6758714c68582b4256753478397030754b53395858422b43633271496c7661676b4737792b546f3162567270315867673d3d7d292e555344a165627a7a72305820b9ed77e4c5d331839496d9393b50fa9539181df70e2c1832cf1e9a4a0821237a0029`

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
// Solidity: function rolls( uint256) constant returns(result_rngseed string, query_rng bytes32, query_price1 bytes32, query_price2 bytes32, result_price1 uint256, result_price2 uint256, timestamp uint256, pool uint256, state uint8, coin uint8, is_up bool)
func (_PriceRoll *PriceRollCaller) Rolls(opts *bind.CallOpts, arg0 *big.Int) (struct {
	ResultRngseed string
	QueryRng      [32]byte
	QueryPrice1   [32]byte
	QueryPrice2   [32]byte
	ResultPrice1  *big.Int
	ResultPrice2  *big.Int
	Timestamp     *big.Int
	Pool          *big.Int
	State         uint8
	Coin          uint8
	IsUp          bool
}, error) {
	ret := new(struct {
		ResultRngseed string
		QueryRng      [32]byte
		QueryPrice1   [32]byte
		QueryPrice2   [32]byte
		ResultPrice1  *big.Int
		ResultPrice2  *big.Int
		Timestamp     *big.Int
		Pool          *big.Int
		State         uint8
		Coin          uint8
		IsUp          bool
	})
	out := ret
	err := _PriceRoll.contract.Call(opts, out, "rolls", arg0)
	return *ret, err
}

// Rolls is a free data retrieval call binding the contract method 0x5d69f16c.
//
// Solidity: function rolls( uint256) constant returns(result_rngseed string, query_rng bytes32, query_price1 bytes32, query_price2 bytes32, result_price1 uint256, result_price2 uint256, timestamp uint256, pool uint256, state uint8, coin uint8, is_up bool)
func (_PriceRoll *PriceRollSession) Rolls(arg0 *big.Int) (struct {
	ResultRngseed string
	QueryRng      [32]byte
	QueryPrice1   [32]byte
	QueryPrice2   [32]byte
	ResultPrice1  *big.Int
	ResultPrice2  *big.Int
	Timestamp     *big.Int
	Pool          *big.Int
	State         uint8
	Coin          uint8
	IsUp          bool
}, error) {
	return _PriceRoll.Contract.Rolls(&_PriceRoll.CallOpts, arg0)
}

// Rolls is a free data retrieval call binding the contract method 0x5d69f16c.
//
// Solidity: function rolls( uint256) constant returns(result_rngseed string, query_rng bytes32, query_price1 bytes32, query_price2 bytes32, result_price1 uint256, result_price2 uint256, timestamp uint256, pool uint256, state uint8, coin uint8, is_up bool)
func (_PriceRoll *PriceRollCallerSession) Rolls(arg0 *big.Int) (struct {
	ResultRngseed string
	QueryRng      [32]byte
	QueryPrice1   [32]byte
	QueryPrice2   [32]byte
	ResultPrice1  *big.Int
	ResultPrice2  *big.Int
	Timestamp     *big.Int
	Pool          *big.Int
	State         uint8
	Coin          uint8
	IsUp          bool
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
	Amount        *big.Int
	Player        common.Address
	ExpectedValue uint8
	IsUp          uint8
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterBetPlaced is a free log retrieval operation binding the contract event 0xa03a715f8e6fbcac20e92a38c1169f81b49b105441690f1e78ed31b17aafba9f.
//
// Solidity: e BetPlaced(round indexed uint256, amount uint256, player indexed address, expected_value uint8, is_up uint8)
func (_PriceRoll *PriceRollFilterer) FilterBetPlaced(opts *bind.FilterOpts, round []*big.Int, player []common.Address) (*PriceRollBetPlacedIterator, error) {

	var roundRule []interface{}
	for _, roundItem := range round {
		roundRule = append(roundRule, roundItem)
	}

	var playerRule []interface{}
	for _, playerItem := range player {
		playerRule = append(playerRule, playerItem)
	}

	logs, sub, err := _PriceRoll.contract.FilterLogs(opts, "BetPlaced", roundRule, playerRule)
	if err != nil {
		return nil, err
	}
	return &PriceRollBetPlacedIterator{contract: _PriceRoll.contract, event: "BetPlaced", logs: logs, sub: sub}, nil
}

// WatchBetPlaced is a free log subscription operation binding the contract event 0xa03a715f8e6fbcac20e92a38c1169f81b49b105441690f1e78ed31b17aafba9f.
//
// Solidity: e BetPlaced(round indexed uint256, amount uint256, player indexed address, expected_value uint8, is_up uint8)
func (_PriceRoll *PriceRollFilterer) WatchBetPlaced(opts *bind.WatchOpts, sink chan<- *PriceRollBetPlaced, round []*big.Int, player []common.Address) (event.Subscription, error) {

	var roundRule []interface{}
	for _, roundItem := range round {
		roundRule = append(roundRule, roundItem)
	}

	var playerRule []interface{}
	for _, playerItem := range player {
		playerRule = append(playerRule, playerItem)
	}

	logs, sub, err := _PriceRoll.contract.WatchLogs(opts, "BetPlaced", roundRule, playerRule)
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
// Solidity: e NewRoll(round indexed uint256)
func (_PriceRoll *PriceRollFilterer) FilterNewRoll(opts *bind.FilterOpts, round []*big.Int) (*PriceRollNewRollIterator, error) {

	var roundRule []interface{}
	for _, roundItem := range round {
		roundRule = append(roundRule, roundItem)
	}

	logs, sub, err := _PriceRoll.contract.FilterLogs(opts, "NewRoll", roundRule)
	if err != nil {
		return nil, err
	}
	return &PriceRollNewRollIterator{contract: _PriceRoll.contract, event: "NewRoll", logs: logs, sub: sub}, nil
}

// WatchNewRoll is a free log subscription operation binding the contract event 0xb4267207f21a1858b0724dfba74816419152c88804d2b6276e5c32f6cd47c74a.
//
// Solidity: e NewRoll(round indexed uint256)
func (_PriceRoll *PriceRollFilterer) WatchNewRoll(opts *bind.WatchOpts, sink chan<- *PriceRollNewRoll, round []*big.Int) (event.Subscription, error) {

	var roundRule []interface{}
	for _, roundItem := range round {
		roundRule = append(roundRule, roundItem)
	}

	logs, sub, err := _PriceRoll.contract.WatchLogs(opts, "NewRoll", roundRule)
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
// Solidity: e RollClaimed(round indexed uint256, player indexed address, amount uint256)
func (_PriceRoll *PriceRollFilterer) FilterRollClaimed(opts *bind.FilterOpts, round []*big.Int, player []common.Address) (*PriceRollRollClaimedIterator, error) {

	var roundRule []interface{}
	for _, roundItem := range round {
		roundRule = append(roundRule, roundItem)
	}
	var playerRule []interface{}
	for _, playerItem := range player {
		playerRule = append(playerRule, playerItem)
	}

	logs, sub, err := _PriceRoll.contract.FilterLogs(opts, "RollClaimed", roundRule, playerRule)
	if err != nil {
		return nil, err
	}
	return &PriceRollRollClaimedIterator{contract: _PriceRoll.contract, event: "RollClaimed", logs: logs, sub: sub}, nil
}

// WatchRollClaimed is a free log subscription operation binding the contract event 0x170c93f55dab07b2a0c856acad14fbb479fc97d6c3f489faf09b5f01c1575d3e.
//
// Solidity: e RollClaimed(round indexed uint256, player indexed address, amount uint256)
func (_PriceRoll *PriceRollFilterer) WatchRollClaimed(opts *bind.WatchOpts, sink chan<- *PriceRollRollClaimed, round []*big.Int, player []common.Address) (event.Subscription, error) {

	var roundRule []interface{}
	for _, roundItem := range round {
		roundRule = append(roundRule, roundItem)
	}
	var playerRule []interface{}
	for _, playerItem := range player {
		playerRule = append(playerRule, playerItem)
	}

	logs, sub, err := _PriceRoll.contract.WatchLogs(opts, "RollClaimed", roundRule, playerRule)
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
	Seed       string
	StartPrice *big.Int
	EndPrice   *big.Int
	Raw        types.Log // Blockchain specific contextual infos
}

// FilterRollEnded is a free log retrieval operation binding the contract event 0x8cb1b96dbae53a90404e7c18d5482ddc5a99bb8ca665731bc035a9054f098faa.
//
// Solidity: e RollEnded(round indexed uint256, seed string, start_price uint256, end_price uint256)
func (_PriceRoll *PriceRollFilterer) FilterRollEnded(opts *bind.FilterOpts, round []*big.Int) (*PriceRollRollEndedIterator, error) {

	var roundRule []interface{}
	for _, roundItem := range round {
		roundRule = append(roundRule, roundItem)
	}

	logs, sub, err := _PriceRoll.contract.FilterLogs(opts, "RollEnded", roundRule)
	if err != nil {
		return nil, err
	}
	return &PriceRollRollEndedIterator{contract: _PriceRoll.contract, event: "RollEnded", logs: logs, sub: sub}, nil
}

// WatchRollEnded is a free log subscription operation binding the contract event 0x8cb1b96dbae53a90404e7c18d5482ddc5a99bb8ca665731bc035a9054f098faa.
//
// Solidity: e RollEnded(round indexed uint256, seed string, start_price uint256, end_price uint256)
func (_PriceRoll *PriceRollFilterer) WatchRollEnded(opts *bind.WatchOpts, sink chan<- *PriceRollRollEnded, round []*big.Int) (event.Subscription, error) {

	var roundRule []interface{}
	for _, roundItem := range round {
		roundRule = append(roundRule, roundItem)
	}

	logs, sub, err := _PriceRoll.contract.WatchLogs(opts, "RollEnded", roundRule)
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
// Solidity: e RollRefunded(round indexed uint256)
func (_PriceRoll *PriceRollFilterer) FilterRollRefunded(opts *bind.FilterOpts, round []*big.Int) (*PriceRollRollRefundedIterator, error) {

	var roundRule []interface{}
	for _, roundItem := range round {
		roundRule = append(roundRule, roundItem)
	}

	logs, sub, err := _PriceRoll.contract.FilterLogs(opts, "RollRefunded", roundRule)
	if err != nil {
		return nil, err
	}
	return &PriceRollRollRefundedIterator{contract: _PriceRoll.contract, event: "RollRefunded", logs: logs, sub: sub}, nil
}

// WatchRollRefunded is a free log subscription operation binding the contract event 0x796cfdf016e01ffdff99dbe8ce27df476b06d0fae2ac92ec3889b6008219ed52.
//
// Solidity: e RollRefunded(round indexed uint256)
func (_PriceRoll *PriceRollFilterer) WatchRollRefunded(opts *bind.WatchOpts, sink chan<- *PriceRollRollRefunded, round []*big.Int) (event.Subscription, error) {

	var roundRule []interface{}
	for _, roundItem := range round {
		roundRule = append(roundRule, roundItem)
	}

	logs, sub, err := _PriceRoll.contract.WatchLogs(opts, "RollRefunded", roundRule)
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
// Solidity: e Rolling(round indexed uint256)
func (_PriceRoll *PriceRollFilterer) FilterRolling(opts *bind.FilterOpts, round []*big.Int) (*PriceRollRollingIterator, error) {

	var roundRule []interface{}
	for _, roundItem := range round {
		roundRule = append(roundRule, roundItem)
	}

	logs, sub, err := _PriceRoll.contract.FilterLogs(opts, "Rolling", roundRule)
	if err != nil {
		return nil, err
	}
	return &PriceRollRollingIterator{contract: _PriceRoll.contract, event: "Rolling", logs: logs, sub: sub}, nil
}

// WatchRolling is a free log subscription operation binding the contract event 0xe86c6fb9a82a2c1fd4c6c98a2fe03844f019514b0cfe8dfbdf076c61f124c0d4.
//
// Solidity: e Rolling(round indexed uint256)
func (_PriceRoll *PriceRollFilterer) WatchRolling(opts *bind.WatchOpts, sink chan<- *PriceRollRolling, round []*big.Int) (event.Subscription, error) {

	var roundRule []interface{}
	for _, roundItem := range round {
		roundRule = append(roundRule, roundItem)
	}

	logs, sub, err := _PriceRoll.contract.WatchLogs(opts, "Rolling", roundRule)
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
const RolesBin = `0x604c602c600b82828239805160001a60731460008114601c57601e565bfe5b5030600052607381538281f3fe73000000000000000000000000000000000000000030146080604052600080fdfea165627a7a723058206d2956c348f6b79507fb40dee8276dd379d9203de1dca72c8a64745abee2fb680029`

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
const SafeMathBin = `0x604c602c600b82828239805160001a60731460008114601c57601e565bfe5b5030600052607381538281f3fe73000000000000000000000000000000000000000030146080604052600080fdfea165627a7a72305820b3d7438a0b7d29d3b3ba133c366d3ffe147fadefe99935d13fca9d3e0573e5880029`

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
const UsingOraclizeBin = `0x608060405234801561001057600080fd5b5061028b806100206000396000f3fe608060405234801561001057600080fd5b5060043610610052577c0100000000000000000000000000000000000000000000000000000000600035046327dc297e811461005757806338bbfa5014610106575b600080fd5b6101046004803603604081101561006d57600080fd5b8135919081019060408101602082013564010000000081111561008f57600080fd5b8201836020820111156100a157600080fd5b803590602001918460018302840111640100000000831117156100c357600080fd5b91908080601f01602080910402602001604051908101604052809392919081815260200183838082843760009201919091525092955061023a945050505050565b005b6101046004803603606081101561011c57600080fd5b8135919081019060408101602082013564010000000081111561013e57600080fd5b82018360208201111561015057600080fd5b8035906020019184600183028401116401000000008311171561017257600080fd5b91908080601f01602080910402602001604051908101604052809392919081815260200183838082843760009201919091525092959493602081019350359150506401000000008111156101c557600080fd5b8201836020820111156101d757600080fd5b803590602001918460018302840111640100000000831117156101f957600080fd5b91908080601f01602080910402602001604051908101604052809392919081815260200183838082843760009201919091525092955061025a945050505050565b604080516000815260208101909152610256908390839061025a565b5050565b50505056fea165627a7a72305820d4a07dbbdf0cc7640b89c67799a175e641b98c6d7d574751c20e10f94c697c6a0029`

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
