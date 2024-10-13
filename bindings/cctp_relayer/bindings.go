// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package bindings

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

// ICCTPRelayerReceiveCall is an auto generated low-level Go binding around an user-defined struct.
type ICCTPRelayerReceiveCall struct {
	Message     []byte
	Attestation []byte
}

// BindingsMetaData contains all meta data concerning the Bindings contract.
var BindingsMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"target\",\"type\":\"address\"}],\"name\":\"AddressEmptyCode\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"AddressInsufficientBalance\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"implementation\",\"type\":\"address\"}],\"name\":\"ERC1967InvalidImplementation\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"ERC1967NonPayable\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"ETHSendFailed\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"FailedInnerCall\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InsufficientNativeToken\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InsufficientSwapOutput\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidInitialization\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"MissingBalance\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"NotInitializing\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"}],\"name\":\"OwnableInvalidOwner\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"OwnableUnauthorizedAccount\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"PaymentCannotBeZero\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"Reentrancy\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"}],\"name\":\"SafeERC20FailedOperation\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"SwapFailed\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"TransferFailed\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"UUPSUnauthorizedCallContext\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"slot\",\"type\":\"bytes32\"}],\"name\":\"UUPSUnsupportedProxiableUUID\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"ZeroAddress\",\"type\":\"error\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"bytes\",\"name\":\"message\",\"type\":\"bytes\"},{\"indexed\":false,\"internalType\":\"bytes\",\"name\":\"attestation\",\"type\":\"bytes\"}],\"name\":\"FailedReceiveMessage\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint64\",\"name\":\"version\",\"type\":\"uint64\"}],\"name\":\"Initialized\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"previousOwner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"OwnershipTransferStarted\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"previousOwner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"OwnershipTransferred\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint64\",\"name\":\"nonce\",\"type\":\"uint64\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"paymentAmount\",\"type\":\"uint256\"}],\"name\":\"PaymentForRelay\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"implementation\",\"type\":\"address\"}],\"name\":\"Upgraded\",\"type\":\"event\"},{\"stateMutability\":\"payable\",\"type\":\"fallback\"},{\"inputs\":[],\"name\":\"UPGRADE_INTERFACE_VERSION\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"acceptOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"bytes\",\"name\":\"message\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"attestation\",\"type\":\"bytes\"}],\"internalType\":\"structICCTPRelayer.ReceiveCall[]\",\"name\":\"receiveCalls\",\"type\":\"tuple[]\"}],\"name\":\"batchReceiveMessage\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"usdc_\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"messenger_\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"transmitter_\",\"type\":\"address\"}],\"name\":\"initialize\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint64\",\"name\":\"nonce\",\"type\":\"uint64\"},{\"internalType\":\"uint256\",\"name\":\"paymentAmount\",\"type\":\"uint256\"}],\"name\":\"makePaymentForRelay\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"messenger\",\"outputs\":[{\"internalType\":\"contractITokenMessenger\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"pendingOwner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"proxiableUUID\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"renounceOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"transferAmount\",\"type\":\"uint256\"},{\"internalType\":\"uint32\",\"name\":\"destinationDomain\",\"type\":\"uint32\"},{\"internalType\":\"bytes32\",\"name\":\"mintRecipient\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"burnToken\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"feeAmount\",\"type\":\"uint256\"}],\"name\":\"requestCCTPTransfer\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"transferAmount\",\"type\":\"uint256\"},{\"internalType\":\"uint32\",\"name\":\"destinationDomain\",\"type\":\"uint32\"},{\"internalType\":\"bytes32\",\"name\":\"mintRecipient\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"burnToken\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"feeAmount\",\"type\":\"uint256\"},{\"internalType\":\"bytes32\",\"name\":\"destinationCaller\",\"type\":\"bytes32\"}],\"name\":\"requestCCTPTransferWithCaller\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_swapRouter\",\"type\":\"address\"}],\"name\":\"setSwapRouter\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"inputToken\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"inputAmount\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"swapCalldata\",\"type\":\"bytes\"},{\"internalType\":\"uint32\",\"name\":\"destinationDomain\",\"type\":\"uint32\"},{\"internalType\":\"bytes32\",\"name\":\"mintRecipient\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"burnToken\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"feeAmount\",\"type\":\"uint256\"}],\"name\":\"swapAndRequestCCTPTransfer\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"inputToken\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"inputAmount\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"swapCalldata\",\"type\":\"bytes\"},{\"internalType\":\"uint32\",\"name\":\"destinationDomain\",\"type\":\"uint32\"},{\"internalType\":\"bytes32\",\"name\":\"mintRecipient\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"burnToken\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"feeAmount\",\"type\":\"uint256\"},{\"internalType\":\"bytes32\",\"name\":\"destinationCaller\",\"type\":\"bytes32\"}],\"name\":\"swapAndRequestCCTPTransferWithCaller\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"swapRouter\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"transferOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"transmitter\",\"outputs\":[{\"internalType\":\"contractIMessageTransmitter\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newImplementation\",\"type\":\"address\"},{\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"}],\"name\":\"upgradeToAndCall\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"usdc\",\"outputs\":[{\"internalType\":\"contractIERC20\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"receiver\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"withdraw\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"stateMutability\":\"payable\",\"type\":\"receive\"}]",
}

// BindingsABI is the input ABI used to generate the binding from.
// Deprecated: Use BindingsMetaData.ABI instead.
var BindingsABI = BindingsMetaData.ABI

// Bindings is an auto generated Go binding around an Ethereum contract.
type Bindings struct {
	BindingsCaller     // Read-only binding to the contract
	BindingsTransactor // Write-only binding to the contract
	BindingsFilterer   // Log filterer for contract events
}

// BindingsCaller is an auto generated read-only Go binding around an Ethereum contract.
type BindingsCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// BindingsTransactor is an auto generated write-only Go binding around an Ethereum contract.
type BindingsTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// BindingsFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type BindingsFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// BindingsSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type BindingsSession struct {
	Contract     *Bindings         // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// BindingsCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type BindingsCallerSession struct {
	Contract *BindingsCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts   // Call options to use throughout this session
}

// BindingsTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type BindingsTransactorSession struct {
	Contract     *BindingsTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts   // Transaction auth options to use throughout this session
}

// BindingsRaw is an auto generated low-level Go binding around an Ethereum contract.
type BindingsRaw struct {
	Contract *Bindings // Generic contract binding to access the raw methods on
}

// BindingsCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type BindingsCallerRaw struct {
	Contract *BindingsCaller // Generic read-only contract binding to access the raw methods on
}

// BindingsTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type BindingsTransactorRaw struct {
	Contract *BindingsTransactor // Generic write-only contract binding to access the raw methods on
}

// NewBindings creates a new instance of Bindings, bound to a specific deployed contract.
func NewBindings(address common.Address, backend bind.ContractBackend) (*Bindings, error) {
	contract, err := bindBindings(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Bindings{BindingsCaller: BindingsCaller{contract: contract}, BindingsTransactor: BindingsTransactor{contract: contract}, BindingsFilterer: BindingsFilterer{contract: contract}}, nil
}

// NewBindingsCaller creates a new read-only instance of Bindings, bound to a specific deployed contract.
func NewBindingsCaller(address common.Address, caller bind.ContractCaller) (*BindingsCaller, error) {
	contract, err := bindBindings(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &BindingsCaller{contract: contract}, nil
}

// NewBindingsTransactor creates a new write-only instance of Bindings, bound to a specific deployed contract.
func NewBindingsTransactor(address common.Address, transactor bind.ContractTransactor) (*BindingsTransactor, error) {
	contract, err := bindBindings(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &BindingsTransactor{contract: contract}, nil
}

// NewBindingsFilterer creates a new log filterer instance of Bindings, bound to a specific deployed contract.
func NewBindingsFilterer(address common.Address, filterer bind.ContractFilterer) (*BindingsFilterer, error) {
	contract, err := bindBindings(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &BindingsFilterer{contract: contract}, nil
}

// bindBindings binds a generic wrapper to an already deployed contract.
func bindBindings(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := BindingsMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Bindings *BindingsRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Bindings.Contract.BindingsCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Bindings *BindingsRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Bindings.Contract.BindingsTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Bindings *BindingsRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Bindings.Contract.BindingsTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Bindings *BindingsCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Bindings.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Bindings *BindingsTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Bindings.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Bindings *BindingsTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Bindings.Contract.contract.Transact(opts, method, params...)
}

// UPGRADEINTERFACEVERSION is a free data retrieval call binding the contract method 0xad3cb1cc.
//
// Solidity: function UPGRADE_INTERFACE_VERSION() view returns(string)
func (_Bindings *BindingsCaller) UPGRADEINTERFACEVERSION(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _Bindings.contract.Call(opts, &out, "UPGRADE_INTERFACE_VERSION")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// UPGRADEINTERFACEVERSION is a free data retrieval call binding the contract method 0xad3cb1cc.
//
// Solidity: function UPGRADE_INTERFACE_VERSION() view returns(string)
func (_Bindings *BindingsSession) UPGRADEINTERFACEVERSION() (string, error) {
	return _Bindings.Contract.UPGRADEINTERFACEVERSION(&_Bindings.CallOpts)
}

// UPGRADEINTERFACEVERSION is a free data retrieval call binding the contract method 0xad3cb1cc.
//
// Solidity: function UPGRADE_INTERFACE_VERSION() view returns(string)
func (_Bindings *BindingsCallerSession) UPGRADEINTERFACEVERSION() (string, error) {
	return _Bindings.Contract.UPGRADEINTERFACEVERSION(&_Bindings.CallOpts)
}

// Messenger is a free data retrieval call binding the contract method 0x3cb747bf.
//
// Solidity: function messenger() view returns(address)
func (_Bindings *BindingsCaller) Messenger(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Bindings.contract.Call(opts, &out, "messenger")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Messenger is a free data retrieval call binding the contract method 0x3cb747bf.
//
// Solidity: function messenger() view returns(address)
func (_Bindings *BindingsSession) Messenger() (common.Address, error) {
	return _Bindings.Contract.Messenger(&_Bindings.CallOpts)
}

// Messenger is a free data retrieval call binding the contract method 0x3cb747bf.
//
// Solidity: function messenger() view returns(address)
func (_Bindings *BindingsCallerSession) Messenger() (common.Address, error) {
	return _Bindings.Contract.Messenger(&_Bindings.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_Bindings *BindingsCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Bindings.contract.Call(opts, &out, "owner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_Bindings *BindingsSession) Owner() (common.Address, error) {
	return _Bindings.Contract.Owner(&_Bindings.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_Bindings *BindingsCallerSession) Owner() (common.Address, error) {
	return _Bindings.Contract.Owner(&_Bindings.CallOpts)
}

// PendingOwner is a free data retrieval call binding the contract method 0xe30c3978.
//
// Solidity: function pendingOwner() view returns(address)
func (_Bindings *BindingsCaller) PendingOwner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Bindings.contract.Call(opts, &out, "pendingOwner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// PendingOwner is a free data retrieval call binding the contract method 0xe30c3978.
//
// Solidity: function pendingOwner() view returns(address)
func (_Bindings *BindingsSession) PendingOwner() (common.Address, error) {
	return _Bindings.Contract.PendingOwner(&_Bindings.CallOpts)
}

// PendingOwner is a free data retrieval call binding the contract method 0xe30c3978.
//
// Solidity: function pendingOwner() view returns(address)
func (_Bindings *BindingsCallerSession) PendingOwner() (common.Address, error) {
	return _Bindings.Contract.PendingOwner(&_Bindings.CallOpts)
}

// ProxiableUUID is a free data retrieval call binding the contract method 0x52d1902d.
//
// Solidity: function proxiableUUID() view returns(bytes32)
func (_Bindings *BindingsCaller) ProxiableUUID(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _Bindings.contract.Call(opts, &out, "proxiableUUID")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// ProxiableUUID is a free data retrieval call binding the contract method 0x52d1902d.
//
// Solidity: function proxiableUUID() view returns(bytes32)
func (_Bindings *BindingsSession) ProxiableUUID() ([32]byte, error) {
	return _Bindings.Contract.ProxiableUUID(&_Bindings.CallOpts)
}

// ProxiableUUID is a free data retrieval call binding the contract method 0x52d1902d.
//
// Solidity: function proxiableUUID() view returns(bytes32)
func (_Bindings *BindingsCallerSession) ProxiableUUID() ([32]byte, error) {
	return _Bindings.Contract.ProxiableUUID(&_Bindings.CallOpts)
}

// SwapRouter is a free data retrieval call binding the contract method 0xc31c9c07.
//
// Solidity: function swapRouter() view returns(address)
func (_Bindings *BindingsCaller) SwapRouter(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Bindings.contract.Call(opts, &out, "swapRouter")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// SwapRouter is a free data retrieval call binding the contract method 0xc31c9c07.
//
// Solidity: function swapRouter() view returns(address)
func (_Bindings *BindingsSession) SwapRouter() (common.Address, error) {
	return _Bindings.Contract.SwapRouter(&_Bindings.CallOpts)
}

// SwapRouter is a free data retrieval call binding the contract method 0xc31c9c07.
//
// Solidity: function swapRouter() view returns(address)
func (_Bindings *BindingsCallerSession) SwapRouter() (common.Address, error) {
	return _Bindings.Contract.SwapRouter(&_Bindings.CallOpts)
}

// Transmitter is a free data retrieval call binding the contract method 0xcec46f6c.
//
// Solidity: function transmitter() view returns(address)
func (_Bindings *BindingsCaller) Transmitter(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Bindings.contract.Call(opts, &out, "transmitter")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Transmitter is a free data retrieval call binding the contract method 0xcec46f6c.
//
// Solidity: function transmitter() view returns(address)
func (_Bindings *BindingsSession) Transmitter() (common.Address, error) {
	return _Bindings.Contract.Transmitter(&_Bindings.CallOpts)
}

// Transmitter is a free data retrieval call binding the contract method 0xcec46f6c.
//
// Solidity: function transmitter() view returns(address)
func (_Bindings *BindingsCallerSession) Transmitter() (common.Address, error) {
	return _Bindings.Contract.Transmitter(&_Bindings.CallOpts)
}

// Usdc is a free data retrieval call binding the contract method 0x3e413bee.
//
// Solidity: function usdc() view returns(address)
func (_Bindings *BindingsCaller) Usdc(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Bindings.contract.Call(opts, &out, "usdc")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Usdc is a free data retrieval call binding the contract method 0x3e413bee.
//
// Solidity: function usdc() view returns(address)
func (_Bindings *BindingsSession) Usdc() (common.Address, error) {
	return _Bindings.Contract.Usdc(&_Bindings.CallOpts)
}

// Usdc is a free data retrieval call binding the contract method 0x3e413bee.
//
// Solidity: function usdc() view returns(address)
func (_Bindings *BindingsCallerSession) Usdc() (common.Address, error) {
	return _Bindings.Contract.Usdc(&_Bindings.CallOpts)
}

// AcceptOwnership is a paid mutator transaction binding the contract method 0x79ba5097.
//
// Solidity: function acceptOwnership() returns()
func (_Bindings *BindingsTransactor) AcceptOwnership(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Bindings.contract.Transact(opts, "acceptOwnership")
}

// AcceptOwnership is a paid mutator transaction binding the contract method 0x79ba5097.
//
// Solidity: function acceptOwnership() returns()
func (_Bindings *BindingsSession) AcceptOwnership() (*types.Transaction, error) {
	return _Bindings.Contract.AcceptOwnership(&_Bindings.TransactOpts)
}

// AcceptOwnership is a paid mutator transaction binding the contract method 0x79ba5097.
//
// Solidity: function acceptOwnership() returns()
func (_Bindings *BindingsTransactorSession) AcceptOwnership() (*types.Transaction, error) {
	return _Bindings.Contract.AcceptOwnership(&_Bindings.TransactOpts)
}

// BatchReceiveMessage is a paid mutator transaction binding the contract method 0x4b54d875.
//
// Solidity: function batchReceiveMessage((bytes,bytes)[] receiveCalls) returns()
func (_Bindings *BindingsTransactor) BatchReceiveMessage(opts *bind.TransactOpts, receiveCalls []ICCTPRelayerReceiveCall) (*types.Transaction, error) {
	return _Bindings.contract.Transact(opts, "batchReceiveMessage", receiveCalls)
}

// BatchReceiveMessage is a paid mutator transaction binding the contract method 0x4b54d875.
//
// Solidity: function batchReceiveMessage((bytes,bytes)[] receiveCalls) returns()
func (_Bindings *BindingsSession) BatchReceiveMessage(receiveCalls []ICCTPRelayerReceiveCall) (*types.Transaction, error) {
	return _Bindings.Contract.BatchReceiveMessage(&_Bindings.TransactOpts, receiveCalls)
}

// BatchReceiveMessage is a paid mutator transaction binding the contract method 0x4b54d875.
//
// Solidity: function batchReceiveMessage((bytes,bytes)[] receiveCalls) returns()
func (_Bindings *BindingsTransactorSession) BatchReceiveMessage(receiveCalls []ICCTPRelayerReceiveCall) (*types.Transaction, error) {
	return _Bindings.Contract.BatchReceiveMessage(&_Bindings.TransactOpts, receiveCalls)
}

// Initialize is a paid mutator transaction binding the contract method 0xc0c53b8b.
//
// Solidity: function initialize(address usdc_, address messenger_, address transmitter_) returns()
func (_Bindings *BindingsTransactor) Initialize(opts *bind.TransactOpts, usdc_ common.Address, messenger_ common.Address, transmitter_ common.Address) (*types.Transaction, error) {
	return _Bindings.contract.Transact(opts, "initialize", usdc_, messenger_, transmitter_)
}

// Initialize is a paid mutator transaction binding the contract method 0xc0c53b8b.
//
// Solidity: function initialize(address usdc_, address messenger_, address transmitter_) returns()
func (_Bindings *BindingsSession) Initialize(usdc_ common.Address, messenger_ common.Address, transmitter_ common.Address) (*types.Transaction, error) {
	return _Bindings.Contract.Initialize(&_Bindings.TransactOpts, usdc_, messenger_, transmitter_)
}

// Initialize is a paid mutator transaction binding the contract method 0xc0c53b8b.
//
// Solidity: function initialize(address usdc_, address messenger_, address transmitter_) returns()
func (_Bindings *BindingsTransactorSession) Initialize(usdc_ common.Address, messenger_ common.Address, transmitter_ common.Address) (*types.Transaction, error) {
	return _Bindings.Contract.Initialize(&_Bindings.TransactOpts, usdc_, messenger_, transmitter_)
}

// MakePaymentForRelay is a paid mutator transaction binding the contract method 0xdc432692.
//
// Solidity: function makePaymentForRelay(uint64 nonce, uint256 paymentAmount) returns()
func (_Bindings *BindingsTransactor) MakePaymentForRelay(opts *bind.TransactOpts, nonce uint64, paymentAmount *big.Int) (*types.Transaction, error) {
	return _Bindings.contract.Transact(opts, "makePaymentForRelay", nonce, paymentAmount)
}

// MakePaymentForRelay is a paid mutator transaction binding the contract method 0xdc432692.
//
// Solidity: function makePaymentForRelay(uint64 nonce, uint256 paymentAmount) returns()
func (_Bindings *BindingsSession) MakePaymentForRelay(nonce uint64, paymentAmount *big.Int) (*types.Transaction, error) {
	return _Bindings.Contract.MakePaymentForRelay(&_Bindings.TransactOpts, nonce, paymentAmount)
}

// MakePaymentForRelay is a paid mutator transaction binding the contract method 0xdc432692.
//
// Solidity: function makePaymentForRelay(uint64 nonce, uint256 paymentAmount) returns()
func (_Bindings *BindingsTransactorSession) MakePaymentForRelay(nonce uint64, paymentAmount *big.Int) (*types.Transaction, error) {
	return _Bindings.Contract.MakePaymentForRelay(&_Bindings.TransactOpts, nonce, paymentAmount)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_Bindings *BindingsTransactor) RenounceOwnership(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Bindings.contract.Transact(opts, "renounceOwnership")
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_Bindings *BindingsSession) RenounceOwnership() (*types.Transaction, error) {
	return _Bindings.Contract.RenounceOwnership(&_Bindings.TransactOpts)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_Bindings *BindingsTransactorSession) RenounceOwnership() (*types.Transaction, error) {
	return _Bindings.Contract.RenounceOwnership(&_Bindings.TransactOpts)
}

// RequestCCTPTransfer is a paid mutator transaction binding the contract method 0x78c940c0.
//
// Solidity: function requestCCTPTransfer(uint256 transferAmount, uint32 destinationDomain, bytes32 mintRecipient, address burnToken, uint256 feeAmount) returns()
func (_Bindings *BindingsTransactor) RequestCCTPTransfer(opts *bind.TransactOpts, transferAmount *big.Int, destinationDomain uint32, mintRecipient [32]byte, burnToken common.Address, feeAmount *big.Int) (*types.Transaction, error) {
	return _Bindings.contract.Transact(opts, "requestCCTPTransfer", transferAmount, destinationDomain, mintRecipient, burnToken, feeAmount)
}

// RequestCCTPTransfer is a paid mutator transaction binding the contract method 0x78c940c0.
//
// Solidity: function requestCCTPTransfer(uint256 transferAmount, uint32 destinationDomain, bytes32 mintRecipient, address burnToken, uint256 feeAmount) returns()
func (_Bindings *BindingsSession) RequestCCTPTransfer(transferAmount *big.Int, destinationDomain uint32, mintRecipient [32]byte, burnToken common.Address, feeAmount *big.Int) (*types.Transaction, error) {
	return _Bindings.Contract.RequestCCTPTransfer(&_Bindings.TransactOpts, transferAmount, destinationDomain, mintRecipient, burnToken, feeAmount)
}

// RequestCCTPTransfer is a paid mutator transaction binding the contract method 0x78c940c0.
//
// Solidity: function requestCCTPTransfer(uint256 transferAmount, uint32 destinationDomain, bytes32 mintRecipient, address burnToken, uint256 feeAmount) returns()
func (_Bindings *BindingsTransactorSession) RequestCCTPTransfer(transferAmount *big.Int, destinationDomain uint32, mintRecipient [32]byte, burnToken common.Address, feeAmount *big.Int) (*types.Transaction, error) {
	return _Bindings.Contract.RequestCCTPTransfer(&_Bindings.TransactOpts, transferAmount, destinationDomain, mintRecipient, burnToken, feeAmount)
}

// RequestCCTPTransferWithCaller is a paid mutator transaction binding the contract method 0xd77d6ec0.
//
// Solidity: function requestCCTPTransferWithCaller(uint256 transferAmount, uint32 destinationDomain, bytes32 mintRecipient, address burnToken, uint256 feeAmount, bytes32 destinationCaller) returns()
func (_Bindings *BindingsTransactor) RequestCCTPTransferWithCaller(opts *bind.TransactOpts, transferAmount *big.Int, destinationDomain uint32, mintRecipient [32]byte, burnToken common.Address, feeAmount *big.Int, destinationCaller [32]byte) (*types.Transaction, error) {
	return _Bindings.contract.Transact(opts, "requestCCTPTransferWithCaller", transferAmount, destinationDomain, mintRecipient, burnToken, feeAmount, destinationCaller)
}

// RequestCCTPTransferWithCaller is a paid mutator transaction binding the contract method 0xd77d6ec0.
//
// Solidity: function requestCCTPTransferWithCaller(uint256 transferAmount, uint32 destinationDomain, bytes32 mintRecipient, address burnToken, uint256 feeAmount, bytes32 destinationCaller) returns()
func (_Bindings *BindingsSession) RequestCCTPTransferWithCaller(transferAmount *big.Int, destinationDomain uint32, mintRecipient [32]byte, burnToken common.Address, feeAmount *big.Int, destinationCaller [32]byte) (*types.Transaction, error) {
	return _Bindings.Contract.RequestCCTPTransferWithCaller(&_Bindings.TransactOpts, transferAmount, destinationDomain, mintRecipient, burnToken, feeAmount, destinationCaller)
}

// RequestCCTPTransferWithCaller is a paid mutator transaction binding the contract method 0xd77d6ec0.
//
// Solidity: function requestCCTPTransferWithCaller(uint256 transferAmount, uint32 destinationDomain, bytes32 mintRecipient, address burnToken, uint256 feeAmount, bytes32 destinationCaller) returns()
func (_Bindings *BindingsTransactorSession) RequestCCTPTransferWithCaller(transferAmount *big.Int, destinationDomain uint32, mintRecipient [32]byte, burnToken common.Address, feeAmount *big.Int, destinationCaller [32]byte) (*types.Transaction, error) {
	return _Bindings.Contract.RequestCCTPTransferWithCaller(&_Bindings.TransactOpts, transferAmount, destinationDomain, mintRecipient, burnToken, feeAmount, destinationCaller)
}

// SetSwapRouter is a paid mutator transaction binding the contract method 0x41273657.
//
// Solidity: function setSwapRouter(address _swapRouter) returns()
func (_Bindings *BindingsTransactor) SetSwapRouter(opts *bind.TransactOpts, _swapRouter common.Address) (*types.Transaction, error) {
	return _Bindings.contract.Transact(opts, "setSwapRouter", _swapRouter)
}

// SetSwapRouter is a paid mutator transaction binding the contract method 0x41273657.
//
// Solidity: function setSwapRouter(address _swapRouter) returns()
func (_Bindings *BindingsSession) SetSwapRouter(_swapRouter common.Address) (*types.Transaction, error) {
	return _Bindings.Contract.SetSwapRouter(&_Bindings.TransactOpts, _swapRouter)
}

// SetSwapRouter is a paid mutator transaction binding the contract method 0x41273657.
//
// Solidity: function setSwapRouter(address _swapRouter) returns()
func (_Bindings *BindingsTransactorSession) SetSwapRouter(_swapRouter common.Address) (*types.Transaction, error) {
	return _Bindings.Contract.SetSwapRouter(&_Bindings.TransactOpts, _swapRouter)
}

// SwapAndRequestCCTPTransfer is a paid mutator transaction binding the contract method 0xed58841d.
//
// Solidity: function swapAndRequestCCTPTransfer(address inputToken, uint256 inputAmount, bytes swapCalldata, uint32 destinationDomain, bytes32 mintRecipient, address burnToken, uint256 feeAmount) payable returns()
func (_Bindings *BindingsTransactor) SwapAndRequestCCTPTransfer(opts *bind.TransactOpts, inputToken common.Address, inputAmount *big.Int, swapCalldata []byte, destinationDomain uint32, mintRecipient [32]byte, burnToken common.Address, feeAmount *big.Int) (*types.Transaction, error) {
	return _Bindings.contract.Transact(opts, "swapAndRequestCCTPTransfer", inputToken, inputAmount, swapCalldata, destinationDomain, mintRecipient, burnToken, feeAmount)
}

// SwapAndRequestCCTPTransfer is a paid mutator transaction binding the contract method 0xed58841d.
//
// Solidity: function swapAndRequestCCTPTransfer(address inputToken, uint256 inputAmount, bytes swapCalldata, uint32 destinationDomain, bytes32 mintRecipient, address burnToken, uint256 feeAmount) payable returns()
func (_Bindings *BindingsSession) SwapAndRequestCCTPTransfer(inputToken common.Address, inputAmount *big.Int, swapCalldata []byte, destinationDomain uint32, mintRecipient [32]byte, burnToken common.Address, feeAmount *big.Int) (*types.Transaction, error) {
	return _Bindings.Contract.SwapAndRequestCCTPTransfer(&_Bindings.TransactOpts, inputToken, inputAmount, swapCalldata, destinationDomain, mintRecipient, burnToken, feeAmount)
}

// SwapAndRequestCCTPTransfer is a paid mutator transaction binding the contract method 0xed58841d.
//
// Solidity: function swapAndRequestCCTPTransfer(address inputToken, uint256 inputAmount, bytes swapCalldata, uint32 destinationDomain, bytes32 mintRecipient, address burnToken, uint256 feeAmount) payable returns()
func (_Bindings *BindingsTransactorSession) SwapAndRequestCCTPTransfer(inputToken common.Address, inputAmount *big.Int, swapCalldata []byte, destinationDomain uint32, mintRecipient [32]byte, burnToken common.Address, feeAmount *big.Int) (*types.Transaction, error) {
	return _Bindings.Contract.SwapAndRequestCCTPTransfer(&_Bindings.TransactOpts, inputToken, inputAmount, swapCalldata, destinationDomain, mintRecipient, burnToken, feeAmount)
}

// SwapAndRequestCCTPTransferWithCaller is a paid mutator transaction binding the contract method 0xef63a915.
//
// Solidity: function swapAndRequestCCTPTransferWithCaller(address inputToken, uint256 inputAmount, bytes swapCalldata, uint32 destinationDomain, bytes32 mintRecipient, address burnToken, uint256 feeAmount, bytes32 destinationCaller) payable returns()
func (_Bindings *BindingsTransactor) SwapAndRequestCCTPTransferWithCaller(opts *bind.TransactOpts, inputToken common.Address, inputAmount *big.Int, swapCalldata []byte, destinationDomain uint32, mintRecipient [32]byte, burnToken common.Address, feeAmount *big.Int, destinationCaller [32]byte) (*types.Transaction, error) {
	return _Bindings.contract.Transact(opts, "swapAndRequestCCTPTransferWithCaller", inputToken, inputAmount, swapCalldata, destinationDomain, mintRecipient, burnToken, feeAmount, destinationCaller)
}

// SwapAndRequestCCTPTransferWithCaller is a paid mutator transaction binding the contract method 0xef63a915.
//
// Solidity: function swapAndRequestCCTPTransferWithCaller(address inputToken, uint256 inputAmount, bytes swapCalldata, uint32 destinationDomain, bytes32 mintRecipient, address burnToken, uint256 feeAmount, bytes32 destinationCaller) payable returns()
func (_Bindings *BindingsSession) SwapAndRequestCCTPTransferWithCaller(inputToken common.Address, inputAmount *big.Int, swapCalldata []byte, destinationDomain uint32, mintRecipient [32]byte, burnToken common.Address, feeAmount *big.Int, destinationCaller [32]byte) (*types.Transaction, error) {
	return _Bindings.Contract.SwapAndRequestCCTPTransferWithCaller(&_Bindings.TransactOpts, inputToken, inputAmount, swapCalldata, destinationDomain, mintRecipient, burnToken, feeAmount, destinationCaller)
}

// SwapAndRequestCCTPTransferWithCaller is a paid mutator transaction binding the contract method 0xef63a915.
//
// Solidity: function swapAndRequestCCTPTransferWithCaller(address inputToken, uint256 inputAmount, bytes swapCalldata, uint32 destinationDomain, bytes32 mintRecipient, address burnToken, uint256 feeAmount, bytes32 destinationCaller) payable returns()
func (_Bindings *BindingsTransactorSession) SwapAndRequestCCTPTransferWithCaller(inputToken common.Address, inputAmount *big.Int, swapCalldata []byte, destinationDomain uint32, mintRecipient [32]byte, burnToken common.Address, feeAmount *big.Int, destinationCaller [32]byte) (*types.Transaction, error) {
	return _Bindings.Contract.SwapAndRequestCCTPTransferWithCaller(&_Bindings.TransactOpts, inputToken, inputAmount, swapCalldata, destinationDomain, mintRecipient, burnToken, feeAmount, destinationCaller)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_Bindings *BindingsTransactor) TransferOwnership(opts *bind.TransactOpts, newOwner common.Address) (*types.Transaction, error) {
	return _Bindings.contract.Transact(opts, "transferOwnership", newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_Bindings *BindingsSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _Bindings.Contract.TransferOwnership(&_Bindings.TransactOpts, newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_Bindings *BindingsTransactorSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _Bindings.Contract.TransferOwnership(&_Bindings.TransactOpts, newOwner)
}

// UpgradeToAndCall is a paid mutator transaction binding the contract method 0x4f1ef286.
//
// Solidity: function upgradeToAndCall(address newImplementation, bytes data) payable returns()
func (_Bindings *BindingsTransactor) UpgradeToAndCall(opts *bind.TransactOpts, newImplementation common.Address, data []byte) (*types.Transaction, error) {
	return _Bindings.contract.Transact(opts, "upgradeToAndCall", newImplementation, data)
}

// UpgradeToAndCall is a paid mutator transaction binding the contract method 0x4f1ef286.
//
// Solidity: function upgradeToAndCall(address newImplementation, bytes data) payable returns()
func (_Bindings *BindingsSession) UpgradeToAndCall(newImplementation common.Address, data []byte) (*types.Transaction, error) {
	return _Bindings.Contract.UpgradeToAndCall(&_Bindings.TransactOpts, newImplementation, data)
}

// UpgradeToAndCall is a paid mutator transaction binding the contract method 0x4f1ef286.
//
// Solidity: function upgradeToAndCall(address newImplementation, bytes data) payable returns()
func (_Bindings *BindingsTransactorSession) UpgradeToAndCall(newImplementation common.Address, data []byte) (*types.Transaction, error) {
	return _Bindings.Contract.UpgradeToAndCall(&_Bindings.TransactOpts, newImplementation, data)
}

// Withdraw is a paid mutator transaction binding the contract method 0xf3fef3a3.
//
// Solidity: function withdraw(address receiver, uint256 amount) returns()
func (_Bindings *BindingsTransactor) Withdraw(opts *bind.TransactOpts, receiver common.Address, amount *big.Int) (*types.Transaction, error) {
	return _Bindings.contract.Transact(opts, "withdraw", receiver, amount)
}

// Withdraw is a paid mutator transaction binding the contract method 0xf3fef3a3.
//
// Solidity: function withdraw(address receiver, uint256 amount) returns()
func (_Bindings *BindingsSession) Withdraw(receiver common.Address, amount *big.Int) (*types.Transaction, error) {
	return _Bindings.Contract.Withdraw(&_Bindings.TransactOpts, receiver, amount)
}

// Withdraw is a paid mutator transaction binding the contract method 0xf3fef3a3.
//
// Solidity: function withdraw(address receiver, uint256 amount) returns()
func (_Bindings *BindingsTransactorSession) Withdraw(receiver common.Address, amount *big.Int) (*types.Transaction, error) {
	return _Bindings.Contract.Withdraw(&_Bindings.TransactOpts, receiver, amount)
}

// Fallback is a paid mutator transaction binding the contract fallback function.
//
// Solidity: fallback() payable returns()
func (_Bindings *BindingsTransactor) Fallback(opts *bind.TransactOpts, calldata []byte) (*types.Transaction, error) {
	return _Bindings.contract.RawTransact(opts, calldata)
}

// Fallback is a paid mutator transaction binding the contract fallback function.
//
// Solidity: fallback() payable returns()
func (_Bindings *BindingsSession) Fallback(calldata []byte) (*types.Transaction, error) {
	return _Bindings.Contract.Fallback(&_Bindings.TransactOpts, calldata)
}

// Fallback is a paid mutator transaction binding the contract fallback function.
//
// Solidity: fallback() payable returns()
func (_Bindings *BindingsTransactorSession) Fallback(calldata []byte) (*types.Transaction, error) {
	return _Bindings.Contract.Fallback(&_Bindings.TransactOpts, calldata)
}

// Receive is a paid mutator transaction binding the contract receive function.
//
// Solidity: receive() payable returns()
func (_Bindings *BindingsTransactor) Receive(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Bindings.contract.RawTransact(opts, nil) // calldata is disallowed for receive function
}

// Receive is a paid mutator transaction binding the contract receive function.
//
// Solidity: receive() payable returns()
func (_Bindings *BindingsSession) Receive() (*types.Transaction, error) {
	return _Bindings.Contract.Receive(&_Bindings.TransactOpts)
}

// Receive is a paid mutator transaction binding the contract receive function.
//
// Solidity: receive() payable returns()
func (_Bindings *BindingsTransactorSession) Receive() (*types.Transaction, error) {
	return _Bindings.Contract.Receive(&_Bindings.TransactOpts)
}

// BindingsFailedReceiveMessageIterator is returned from FilterFailedReceiveMessage and is used to iterate over the raw logs and unpacked data for FailedReceiveMessage events raised by the Bindings contract.
type BindingsFailedReceiveMessageIterator struct {
	Event *BindingsFailedReceiveMessage // Event containing the contract specifics and raw log

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
func (it *BindingsFailedReceiveMessageIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(BindingsFailedReceiveMessage)
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
		it.Event = new(BindingsFailedReceiveMessage)
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
func (it *BindingsFailedReceiveMessageIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *BindingsFailedReceiveMessageIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// BindingsFailedReceiveMessage represents a FailedReceiveMessage event raised by the Bindings contract.
type BindingsFailedReceiveMessage struct {
	Message     []byte
	Attestation []byte
	Raw         types.Log // Blockchain specific contextual infos
}

// FilterFailedReceiveMessage is a free log retrieval operation binding the contract event 0x2a30267cdd951d6585d7afbe0e939c70408957ca95494eb906e4fa22943ad6de.
//
// Solidity: event FailedReceiveMessage(bytes message, bytes attestation)
func (_Bindings *BindingsFilterer) FilterFailedReceiveMessage(opts *bind.FilterOpts) (*BindingsFailedReceiveMessageIterator, error) {

	logs, sub, err := _Bindings.contract.FilterLogs(opts, "FailedReceiveMessage")
	if err != nil {
		return nil, err
	}
	return &BindingsFailedReceiveMessageIterator{contract: _Bindings.contract, event: "FailedReceiveMessage", logs: logs, sub: sub}, nil
}

// WatchFailedReceiveMessage is a free log subscription operation binding the contract event 0x2a30267cdd951d6585d7afbe0e939c70408957ca95494eb906e4fa22943ad6de.
//
// Solidity: event FailedReceiveMessage(bytes message, bytes attestation)
func (_Bindings *BindingsFilterer) WatchFailedReceiveMessage(opts *bind.WatchOpts, sink chan<- *BindingsFailedReceiveMessage) (event.Subscription, error) {

	logs, sub, err := _Bindings.contract.WatchLogs(opts, "FailedReceiveMessage")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(BindingsFailedReceiveMessage)
				if err := _Bindings.contract.UnpackLog(event, "FailedReceiveMessage", log); err != nil {
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

// ParseFailedReceiveMessage is a log parse operation binding the contract event 0x2a30267cdd951d6585d7afbe0e939c70408957ca95494eb906e4fa22943ad6de.
//
// Solidity: event FailedReceiveMessage(bytes message, bytes attestation)
func (_Bindings *BindingsFilterer) ParseFailedReceiveMessage(log types.Log) (*BindingsFailedReceiveMessage, error) {
	event := new(BindingsFailedReceiveMessage)
	if err := _Bindings.contract.UnpackLog(event, "FailedReceiveMessage", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// BindingsInitializedIterator is returned from FilterInitialized and is used to iterate over the raw logs and unpacked data for Initialized events raised by the Bindings contract.
type BindingsInitializedIterator struct {
	Event *BindingsInitialized // Event containing the contract specifics and raw log

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
func (it *BindingsInitializedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(BindingsInitialized)
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
		it.Event = new(BindingsInitialized)
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
func (it *BindingsInitializedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *BindingsInitializedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// BindingsInitialized represents a Initialized event raised by the Bindings contract.
type BindingsInitialized struct {
	Version uint64
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterInitialized is a free log retrieval operation binding the contract event 0xc7f505b2f371ae2175ee4913f4499e1f2633a7b5936321eed1cdaeb6115181d2.
//
// Solidity: event Initialized(uint64 version)
func (_Bindings *BindingsFilterer) FilterInitialized(opts *bind.FilterOpts) (*BindingsInitializedIterator, error) {

	logs, sub, err := _Bindings.contract.FilterLogs(opts, "Initialized")
	if err != nil {
		return nil, err
	}
	return &BindingsInitializedIterator{contract: _Bindings.contract, event: "Initialized", logs: logs, sub: sub}, nil
}

// WatchInitialized is a free log subscription operation binding the contract event 0xc7f505b2f371ae2175ee4913f4499e1f2633a7b5936321eed1cdaeb6115181d2.
//
// Solidity: event Initialized(uint64 version)
func (_Bindings *BindingsFilterer) WatchInitialized(opts *bind.WatchOpts, sink chan<- *BindingsInitialized) (event.Subscription, error) {

	logs, sub, err := _Bindings.contract.WatchLogs(opts, "Initialized")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(BindingsInitialized)
				if err := _Bindings.contract.UnpackLog(event, "Initialized", log); err != nil {
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

// ParseInitialized is a log parse operation binding the contract event 0xc7f505b2f371ae2175ee4913f4499e1f2633a7b5936321eed1cdaeb6115181d2.
//
// Solidity: event Initialized(uint64 version)
func (_Bindings *BindingsFilterer) ParseInitialized(log types.Log) (*BindingsInitialized, error) {
	event := new(BindingsInitialized)
	if err := _Bindings.contract.UnpackLog(event, "Initialized", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// BindingsOwnershipTransferStartedIterator is returned from FilterOwnershipTransferStarted and is used to iterate over the raw logs and unpacked data for OwnershipTransferStarted events raised by the Bindings contract.
type BindingsOwnershipTransferStartedIterator struct {
	Event *BindingsOwnershipTransferStarted // Event containing the contract specifics and raw log

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
func (it *BindingsOwnershipTransferStartedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(BindingsOwnershipTransferStarted)
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
		it.Event = new(BindingsOwnershipTransferStarted)
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
func (it *BindingsOwnershipTransferStartedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *BindingsOwnershipTransferStartedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// BindingsOwnershipTransferStarted represents a OwnershipTransferStarted event raised by the Bindings contract.
type BindingsOwnershipTransferStarted struct {
	PreviousOwner common.Address
	NewOwner      common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterOwnershipTransferStarted is a free log retrieval operation binding the contract event 0x38d16b8cac22d99fc7c124b9cd0de2d3fa1faef420bfe791d8c362d765e22700.
//
// Solidity: event OwnershipTransferStarted(address indexed previousOwner, address indexed newOwner)
func (_Bindings *BindingsFilterer) FilterOwnershipTransferStarted(opts *bind.FilterOpts, previousOwner []common.Address, newOwner []common.Address) (*BindingsOwnershipTransferStartedIterator, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _Bindings.contract.FilterLogs(opts, "OwnershipTransferStarted", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return &BindingsOwnershipTransferStartedIterator{contract: _Bindings.contract, event: "OwnershipTransferStarted", logs: logs, sub: sub}, nil
}

// WatchOwnershipTransferStarted is a free log subscription operation binding the contract event 0x38d16b8cac22d99fc7c124b9cd0de2d3fa1faef420bfe791d8c362d765e22700.
//
// Solidity: event OwnershipTransferStarted(address indexed previousOwner, address indexed newOwner)
func (_Bindings *BindingsFilterer) WatchOwnershipTransferStarted(opts *bind.WatchOpts, sink chan<- *BindingsOwnershipTransferStarted, previousOwner []common.Address, newOwner []common.Address) (event.Subscription, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _Bindings.contract.WatchLogs(opts, "OwnershipTransferStarted", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(BindingsOwnershipTransferStarted)
				if err := _Bindings.contract.UnpackLog(event, "OwnershipTransferStarted", log); err != nil {
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

// ParseOwnershipTransferStarted is a log parse operation binding the contract event 0x38d16b8cac22d99fc7c124b9cd0de2d3fa1faef420bfe791d8c362d765e22700.
//
// Solidity: event OwnershipTransferStarted(address indexed previousOwner, address indexed newOwner)
func (_Bindings *BindingsFilterer) ParseOwnershipTransferStarted(log types.Log) (*BindingsOwnershipTransferStarted, error) {
	event := new(BindingsOwnershipTransferStarted)
	if err := _Bindings.contract.UnpackLog(event, "OwnershipTransferStarted", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// BindingsOwnershipTransferredIterator is returned from FilterOwnershipTransferred and is used to iterate over the raw logs and unpacked data for OwnershipTransferred events raised by the Bindings contract.
type BindingsOwnershipTransferredIterator struct {
	Event *BindingsOwnershipTransferred // Event containing the contract specifics and raw log

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
func (it *BindingsOwnershipTransferredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(BindingsOwnershipTransferred)
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
		it.Event = new(BindingsOwnershipTransferred)
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
func (it *BindingsOwnershipTransferredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *BindingsOwnershipTransferredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// BindingsOwnershipTransferred represents a OwnershipTransferred event raised by the Bindings contract.
type BindingsOwnershipTransferred struct {
	PreviousOwner common.Address
	NewOwner      common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterOwnershipTransferred is a free log retrieval operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_Bindings *BindingsFilterer) FilterOwnershipTransferred(opts *bind.FilterOpts, previousOwner []common.Address, newOwner []common.Address) (*BindingsOwnershipTransferredIterator, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _Bindings.contract.FilterLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return &BindingsOwnershipTransferredIterator{contract: _Bindings.contract, event: "OwnershipTransferred", logs: logs, sub: sub}, nil
}

// WatchOwnershipTransferred is a free log subscription operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_Bindings *BindingsFilterer) WatchOwnershipTransferred(opts *bind.WatchOpts, sink chan<- *BindingsOwnershipTransferred, previousOwner []common.Address, newOwner []common.Address) (event.Subscription, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _Bindings.contract.WatchLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(BindingsOwnershipTransferred)
				if err := _Bindings.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
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
func (_Bindings *BindingsFilterer) ParseOwnershipTransferred(log types.Log) (*BindingsOwnershipTransferred, error) {
	event := new(BindingsOwnershipTransferred)
	if err := _Bindings.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// BindingsPaymentForRelayIterator is returned from FilterPaymentForRelay and is used to iterate over the raw logs and unpacked data for PaymentForRelay events raised by the Bindings contract.
type BindingsPaymentForRelayIterator struct {
	Event *BindingsPaymentForRelay // Event containing the contract specifics and raw log

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
func (it *BindingsPaymentForRelayIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(BindingsPaymentForRelay)
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
		it.Event = new(BindingsPaymentForRelay)
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
func (it *BindingsPaymentForRelayIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *BindingsPaymentForRelayIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// BindingsPaymentForRelay represents a PaymentForRelay event raised by the Bindings contract.
type BindingsPaymentForRelay struct {
	Nonce         uint64
	PaymentAmount *big.Int
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterPaymentForRelay is a free log retrieval operation binding the contract event 0x454970efb901c4c9b7539333d09e15311313f3cb4fbbe20bdb7667b7c2e0baff.
//
// Solidity: event PaymentForRelay(uint64 nonce, uint256 paymentAmount)
func (_Bindings *BindingsFilterer) FilterPaymentForRelay(opts *bind.FilterOpts) (*BindingsPaymentForRelayIterator, error) {

	logs, sub, err := _Bindings.contract.FilterLogs(opts, "PaymentForRelay")
	if err != nil {
		return nil, err
	}
	return &BindingsPaymentForRelayIterator{contract: _Bindings.contract, event: "PaymentForRelay", logs: logs, sub: sub}, nil
}

// WatchPaymentForRelay is a free log subscription operation binding the contract event 0x454970efb901c4c9b7539333d09e15311313f3cb4fbbe20bdb7667b7c2e0baff.
//
// Solidity: event PaymentForRelay(uint64 nonce, uint256 paymentAmount)
func (_Bindings *BindingsFilterer) WatchPaymentForRelay(opts *bind.WatchOpts, sink chan<- *BindingsPaymentForRelay) (event.Subscription, error) {

	logs, sub, err := _Bindings.contract.WatchLogs(opts, "PaymentForRelay")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(BindingsPaymentForRelay)
				if err := _Bindings.contract.UnpackLog(event, "PaymentForRelay", log); err != nil {
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

// ParsePaymentForRelay is a log parse operation binding the contract event 0x454970efb901c4c9b7539333d09e15311313f3cb4fbbe20bdb7667b7c2e0baff.
//
// Solidity: event PaymentForRelay(uint64 nonce, uint256 paymentAmount)
func (_Bindings *BindingsFilterer) ParsePaymentForRelay(log types.Log) (*BindingsPaymentForRelay, error) {
	event := new(BindingsPaymentForRelay)
	if err := _Bindings.contract.UnpackLog(event, "PaymentForRelay", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// BindingsUpgradedIterator is returned from FilterUpgraded and is used to iterate over the raw logs and unpacked data for Upgraded events raised by the Bindings contract.
type BindingsUpgradedIterator struct {
	Event *BindingsUpgraded // Event containing the contract specifics and raw log

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
func (it *BindingsUpgradedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(BindingsUpgraded)
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
		it.Event = new(BindingsUpgraded)
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
func (it *BindingsUpgradedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *BindingsUpgradedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// BindingsUpgraded represents a Upgraded event raised by the Bindings contract.
type BindingsUpgraded struct {
	Implementation common.Address
	Raw            types.Log // Blockchain specific contextual infos
}

// FilterUpgraded is a free log retrieval operation binding the contract event 0xbc7cd75a20ee27fd9adebab32041f755214dbc6bffa90cc0225b39da2e5c2d3b.
//
// Solidity: event Upgraded(address indexed implementation)
func (_Bindings *BindingsFilterer) FilterUpgraded(opts *bind.FilterOpts, implementation []common.Address) (*BindingsUpgradedIterator, error) {

	var implementationRule []interface{}
	for _, implementationItem := range implementation {
		implementationRule = append(implementationRule, implementationItem)
	}

	logs, sub, err := _Bindings.contract.FilterLogs(opts, "Upgraded", implementationRule)
	if err != nil {
		return nil, err
	}
	return &BindingsUpgradedIterator{contract: _Bindings.contract, event: "Upgraded", logs: logs, sub: sub}, nil
}

// WatchUpgraded is a free log subscription operation binding the contract event 0xbc7cd75a20ee27fd9adebab32041f755214dbc6bffa90cc0225b39da2e5c2d3b.
//
// Solidity: event Upgraded(address indexed implementation)
func (_Bindings *BindingsFilterer) WatchUpgraded(opts *bind.WatchOpts, sink chan<- *BindingsUpgraded, implementation []common.Address) (event.Subscription, error) {

	var implementationRule []interface{}
	for _, implementationItem := range implementation {
		implementationRule = append(implementationRule, implementationItem)
	}

	logs, sub, err := _Bindings.contract.WatchLogs(opts, "Upgraded", implementationRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(BindingsUpgraded)
				if err := _Bindings.contract.UnpackLog(event, "Upgraded", log); err != nil {
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

// ParseUpgraded is a log parse operation binding the contract event 0xbc7cd75a20ee27fd9adebab32041f755214dbc6bffa90cc0225b39da2e5c2d3b.
//
// Solidity: event Upgraded(address indexed implementation)
func (_Bindings *BindingsFilterer) ParseUpgraded(log types.Log) (*BindingsUpgraded, error) {
	event := new(BindingsUpgraded)
	if err := _Bindings.contract.UnpackLog(event, "Upgraded", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
