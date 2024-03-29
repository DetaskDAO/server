// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package abi

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
)

// DeStageStage is an auto generated low-level Go binding around an user-defined struct.
type DeStageStage struct {
	Amount *big.Int
	Period uint32
	Status uint8
}

// IPermit2PermitTransferFrom is an auto generated low-level Go binding around an user-defined struct.
type IPermit2PermitTransferFrom struct {
	Permitted IPermit2TokenPermissions
	Nonce     *big.Int
	Deadline  *big.Int
}

// IPermit2TokenPermissions is an auto generated low-level Go binding around an user-defined struct.
type IPermit2TokenPermissions struct {
	Token  common.Address
	Amount *big.Int
}

// Order is an auto generated low-level Go binding around an user-defined struct.
type Order struct {
	TaskId    *big.Int
	Issuer    common.Address
	Amount    *big.Int
	Worker    common.Address
	Payed     *big.Int
	Token     common.Address
	Progress  uint8
	PayType   uint8
	StartDate uint32
}

// DeOrderMetaData contains all meta data concerning the DeOrder contract.
var DeOrderMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_weth\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_permit2\",\"type\":\"address\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"reason\",\"type\":\"uint256\"}],\"name\":\"AmountError\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"Expired\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidCaller\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"NonceError\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"ParamError\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"PermissionsError\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"ProgressError\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"season\",\"type\":\"uint256\"}],\"name\":\"RecoverError\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"StatusError\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"UnSupportToken\",\"type\":\"error\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"orderId\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"period\",\"type\":\"uint256\"}],\"name\":\"AppendStage\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"orderId\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"string\",\"name\":\"attachment\",\"type\":\"string\"}],\"name\":\"AttachmentUpdated\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"orderId\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"stageIndex\",\"type\":\"uint256\"}],\"name\":\"ConfirmOrderStage\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"fee\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"feeTo\",\"type\":\"address\"}],\"name\":\"FeeUpdated\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"orderId\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"who\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"stageIndex\",\"type\":\"uint256\"}],\"name\":\"OrderAbort\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"taskId\",\"type\":\"uint256\"},{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"orderId\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"issuer\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"worker\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"OrderCreated\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"orderId\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"OrderModified\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"orderId\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"who\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"payType\",\"type\":\"uint256\"}],\"name\":\"OrderStarted\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"previousOwner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"OwnershipTransferred\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"orderId\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"stageIndex\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"appendPeriod\",\"type\":\"uint256\"}],\"name\":\"ProlongStage\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"orderId\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256[]\",\"name\":\"amounts\",\"type\":\"uint256[]\"},{\"indexed\":false,\"internalType\":\"uint256[]\",\"name\":\"periods\",\"type\":\"uint256[]\"}],\"name\":\"SetStage\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"bool\",\"name\":\"enabled\",\"type\":\"bool\"}],\"name\":\"SupportToken\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"orderId\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"stageIndex\",\"type\":\"uint256\"}],\"name\":\"Withdraw\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"DOMAIN_SEPARATOR\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"PERMIT2\",\"outputs\":[{\"internalType\":\"contractIPermit2\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"PERMITAPPENDSTAGE_TYPEHASH\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"PERMITPROSTAGE_TYPEHASH\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"PERMITSTAGE_TYPEHASH\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"WETH\",\"outputs\":[{\"internalType\":\"contractIWETH9\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_orderId\",\"type\":\"uint256\"}],\"name\":\"abortOrder\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_orderId\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"period\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"nonce\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"deadline\",\"type\":\"uint256\"},{\"internalType\":\"uint8\",\"name\":\"v\",\"type\":\"uint8\"},{\"internalType\":\"bytes32\",\"name\":\"r\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"s\",\"type\":\"bytes32\"}],\"name\":\"appendStage\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"builderSBT\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_orderId\",\"type\":\"uint256\"},{\"internalType\":\"uint256[]\",\"name\":\"_stageIndexs\",\"type\":\"uint256[]\"}],\"name\":\"confirmDelivery\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_taskId\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"_issuer\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_worker\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_token\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_amount\",\"type\":\"uint256\"}],\"name\":\"createOrder\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"currOrderId\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"fee\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"feeTo\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"orderId\",\"type\":\"uint256\"}],\"name\":\"getOrder\",\"outputs\":[{\"components\":[{\"internalType\":\"uint256\",\"name\":\"taskId\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"issuer\",\"type\":\"address\"},{\"internalType\":\"uint96\",\"name\":\"amount\",\"type\":\"uint96\"},{\"internalType\":\"address\",\"name\":\"worker\",\"type\":\"address\"},{\"internalType\":\"uint96\",\"name\":\"payed\",\"type\":\"uint96\"},{\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"},{\"internalType\":\"enumOrderProgess\",\"name\":\"progress\",\"type\":\"uint8\"},{\"internalType\":\"enumPaymentType\",\"name\":\"payType\",\"type\":\"uint8\"},{\"internalType\":\"uint32\",\"name\":\"startDate\",\"type\":\"uint32\"}],\"internalType\":\"structOrder\",\"name\":\"\",\"type\":\"tuple\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_orderId\",\"type\":\"uint256\"}],\"name\":\"getStages\",\"outputs\":[{\"components\":[{\"internalType\":\"uint96\",\"name\":\"amount\",\"type\":\"uint96\"},{\"internalType\":\"uint32\",\"name\":\"period\",\"type\":\"uint32\"},{\"internalType\":\"enumDeStage.StageStatus\",\"name\":\"status\",\"type\":\"uint8\"}],\"internalType\":\"structDeStage.Stage[]\",\"name\":\"stages\",\"type\":\"tuple[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"issuerSBT\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"orderId\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"modifyOrder\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes[]\",\"name\":\"data\",\"type\":\"bytes[]\"}],\"name\":\"multicall\",\"outputs\":[{\"internalType\":\"bytes[]\",\"name\":\"results\",\"type\":\"bytes[]\"}],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"nonces\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_orderId\",\"type\":\"uint256\"}],\"name\":\"ongoingStage\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"stageIndex\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"stageStartDate\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"orderId\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"payOrder\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"orderId\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"deadline\",\"type\":\"uint256\"},{\"internalType\":\"uint8\",\"name\":\"v\",\"type\":\"uint8\"},{\"internalType\":\"bytes32\",\"name\":\"r\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"s\",\"type\":\"bytes32\"}],\"name\":\"payOrderWithPermit\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"orderId\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"components\":[{\"components\":[{\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"internalType\":\"structIPermit2.TokenPermissions\",\"name\":\"permitted\",\"type\":\"tuple\"},{\"internalType\":\"uint256\",\"name\":\"nonce\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"deadline\",\"type\":\"uint256\"}],\"internalType\":\"structIPermit2.PermitTransferFrom\",\"name\":\"permit\",\"type\":\"tuple\"},{\"internalType\":\"bytes\",\"name\":\"signature\",\"type\":\"bytes\"}],\"name\":\"payOrderWithPermit2\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_orderId\",\"type\":\"uint256\"}],\"name\":\"pendingWithdraw\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"pending\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"nextStage\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_orderId\",\"type\":\"uint256\"},{\"internalType\":\"uint256[]\",\"name\":\"_amounts\",\"type\":\"uint256[]\"},{\"internalType\":\"uint256[]\",\"name\":\"_periods\",\"type\":\"uint256[]\"},{\"internalType\":\"enumPaymentType\",\"name\":\"payType\",\"type\":\"uint8\"},{\"internalType\":\"uint256\",\"name\":\"nonce\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"deadline\",\"type\":\"uint256\"},{\"internalType\":\"uint8\",\"name\":\"v\",\"type\":\"uint8\"},{\"internalType\":\"bytes32\",\"name\":\"r\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"s\",\"type\":\"bytes32\"}],\"name\":\"permitStage\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_orderId\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_stageIndex\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_appendPeriod\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"nonce\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"deadline\",\"type\":\"uint256\"},{\"internalType\":\"uint8\",\"name\":\"v\",\"type\":\"uint8\"},{\"internalType\":\"bytes32\",\"name\":\"r\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"s\",\"type\":\"bytes32\"}],\"name\":\"prolongStage\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_orderId\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"_to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_amount\",\"type\":\"uint256\"}],\"name\":\"refund\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"renounceOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_fee\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"_feeTo\",\"type\":\"address\"}],\"name\":\"setFeeTo\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_builder\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_issuer\",\"type\":\"address\"}],\"name\":\"setSBT\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_token\",\"type\":\"address\"},{\"internalType\":\"bool\",\"name\":\"enable\",\"type\":\"bool\"}],\"name\":\"setSupportToken\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"orderId\",\"type\":\"uint256\"}],\"name\":\"stageTotalAmount\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"total\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"orderId\",\"type\":\"uint256\"}],\"name\":\"stagesLength\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"len\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_orderId\",\"type\":\"uint256\"}],\"name\":\"startOrder\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"supportTokens\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"orderId\",\"type\":\"uint256\"}],\"name\":\"totalStagePeriod\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"total\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"transferOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_orderId\",\"type\":\"uint256\"},{\"internalType\":\"string\",\"name\":\"_attachment\",\"type\":\"string\"}],\"name\":\"updateAttachment\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_orderId\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"}],\"name\":\"withdraw\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"stateMutability\":\"payable\",\"type\":\"receive\"}]",
}

// DeOrderABI is the input ABI used to generate the binding from.
// Deprecated: Use DeOrderMetaData.ABI instead.
var DeOrderABI = DeOrderMetaData.ABI

// DeOrder is an auto generated Go binding around an Ethereum contract.
type DeOrder struct {
	DeOrderCaller     // Read-only binding to the contract
	DeOrderTransactor // Write-only binding to the contract
	DeOrderFilterer   // Log filterer for contract events
}

// DeOrderCaller is an auto generated read-only Go binding around an Ethereum contract.
type DeOrderCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// DeOrderTransactor is an auto generated write-only Go binding around an Ethereum contract.
type DeOrderTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// DeOrderFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type DeOrderFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// DeOrderSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type DeOrderSession struct {
	Contract     *DeOrder          // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// DeOrderCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type DeOrderCallerSession struct {
	Contract *DeOrderCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts  // Call options to use throughout this session
}

// DeOrderTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type DeOrderTransactorSession struct {
	Contract     *DeOrderTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts  // Transaction auth options to use throughout this session
}

// DeOrderRaw is an auto generated low-level Go binding around an Ethereum contract.
type DeOrderRaw struct {
	Contract *DeOrder // Generic contract binding to access the raw methods on
}

// DeOrderCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type DeOrderCallerRaw struct {
	Contract *DeOrderCaller // Generic read-only contract binding to access the raw methods on
}

// DeOrderTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type DeOrderTransactorRaw struct {
	Contract *DeOrderTransactor // Generic write-only contract binding to access the raw methods on
}

// NewDeOrder creates a new instance of DeOrder, bound to a specific deployed contract.
func NewDeOrder(address common.Address, backend bind.ContractBackend) (*DeOrder, error) {
	contract, err := bindDeOrder(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &DeOrder{DeOrderCaller: DeOrderCaller{contract: contract}, DeOrderTransactor: DeOrderTransactor{contract: contract}, DeOrderFilterer: DeOrderFilterer{contract: contract}}, nil
}

// NewDeOrderCaller creates a new read-only instance of DeOrder, bound to a specific deployed contract.
func NewDeOrderCaller(address common.Address, caller bind.ContractCaller) (*DeOrderCaller, error) {
	contract, err := bindDeOrder(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &DeOrderCaller{contract: contract}, nil
}

// NewDeOrderTransactor creates a new write-only instance of DeOrder, bound to a specific deployed contract.
func NewDeOrderTransactor(address common.Address, transactor bind.ContractTransactor) (*DeOrderTransactor, error) {
	contract, err := bindDeOrder(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &DeOrderTransactor{contract: contract}, nil
}

// NewDeOrderFilterer creates a new log filterer instance of DeOrder, bound to a specific deployed contract.
func NewDeOrderFilterer(address common.Address, filterer bind.ContractFilterer) (*DeOrderFilterer, error) {
	contract, err := bindDeOrder(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &DeOrderFilterer{contract: contract}, nil
}

// bindDeOrder binds a generic wrapper to an already deployed contract.
func bindDeOrder(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(DeOrderABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_DeOrder *DeOrderRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _DeOrder.Contract.DeOrderCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_DeOrder *DeOrderRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _DeOrder.Contract.DeOrderTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_DeOrder *DeOrderRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _DeOrder.Contract.DeOrderTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_DeOrder *DeOrderCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _DeOrder.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_DeOrder *DeOrderTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _DeOrder.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_DeOrder *DeOrderTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _DeOrder.Contract.contract.Transact(opts, method, params...)
}

// DOMAINSEPARATOR is a free data retrieval call binding the contract method 0x3644e515.
//
// Solidity: function DOMAIN_SEPARATOR() view returns(bytes32)
func (_DeOrder *DeOrderCaller) DOMAINSEPARATOR(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _DeOrder.contract.Call(opts, &out, "DOMAIN_SEPARATOR")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// DOMAINSEPARATOR is a free data retrieval call binding the contract method 0x3644e515.
//
// Solidity: function DOMAIN_SEPARATOR() view returns(bytes32)
func (_DeOrder *DeOrderSession) DOMAINSEPARATOR() ([32]byte, error) {
	return _DeOrder.Contract.DOMAINSEPARATOR(&_DeOrder.CallOpts)
}

// DOMAINSEPARATOR is a free data retrieval call binding the contract method 0x3644e515.
//
// Solidity: function DOMAIN_SEPARATOR() view returns(bytes32)
func (_DeOrder *DeOrderCallerSession) DOMAINSEPARATOR() ([32]byte, error) {
	return _DeOrder.Contract.DOMAINSEPARATOR(&_DeOrder.CallOpts)
}

// PERMIT2 is a free data retrieval call binding the contract method 0x6afdd850.
//
// Solidity: function PERMIT2() view returns(address)
func (_DeOrder *DeOrderCaller) PERMIT2(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _DeOrder.contract.Call(opts, &out, "PERMIT2")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// PERMIT2 is a free data retrieval call binding the contract method 0x6afdd850.
//
// Solidity: function PERMIT2() view returns(address)
func (_DeOrder *DeOrderSession) PERMIT2() (common.Address, error) {
	return _DeOrder.Contract.PERMIT2(&_DeOrder.CallOpts)
}

// PERMIT2 is a free data retrieval call binding the contract method 0x6afdd850.
//
// Solidity: function PERMIT2() view returns(address)
func (_DeOrder *DeOrderCallerSession) PERMIT2() (common.Address, error) {
	return _DeOrder.Contract.PERMIT2(&_DeOrder.CallOpts)
}

// PERMITAPPENDSTAGETYPEHASH is a free data retrieval call binding the contract method 0x235bb0ca.
//
// Solidity: function PERMITAPPENDSTAGE_TYPEHASH() view returns(bytes32)
func (_DeOrder *DeOrderCaller) PERMITAPPENDSTAGETYPEHASH(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _DeOrder.contract.Call(opts, &out, "PERMITAPPENDSTAGE_TYPEHASH")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// PERMITAPPENDSTAGETYPEHASH is a free data retrieval call binding the contract method 0x235bb0ca.
//
// Solidity: function PERMITAPPENDSTAGE_TYPEHASH() view returns(bytes32)
func (_DeOrder *DeOrderSession) PERMITAPPENDSTAGETYPEHASH() ([32]byte, error) {
	return _DeOrder.Contract.PERMITAPPENDSTAGETYPEHASH(&_DeOrder.CallOpts)
}

// PERMITAPPENDSTAGETYPEHASH is a free data retrieval call binding the contract method 0x235bb0ca.
//
// Solidity: function PERMITAPPENDSTAGE_TYPEHASH() view returns(bytes32)
func (_DeOrder *DeOrderCallerSession) PERMITAPPENDSTAGETYPEHASH() ([32]byte, error) {
	return _DeOrder.Contract.PERMITAPPENDSTAGETYPEHASH(&_DeOrder.CallOpts)
}

// PERMITPROSTAGETYPEHASH is a free data retrieval call binding the contract method 0x31f3a62c.
//
// Solidity: function PERMITPROSTAGE_TYPEHASH() view returns(bytes32)
func (_DeOrder *DeOrderCaller) PERMITPROSTAGETYPEHASH(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _DeOrder.contract.Call(opts, &out, "PERMITPROSTAGE_TYPEHASH")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// PERMITPROSTAGETYPEHASH is a free data retrieval call binding the contract method 0x31f3a62c.
//
// Solidity: function PERMITPROSTAGE_TYPEHASH() view returns(bytes32)
func (_DeOrder *DeOrderSession) PERMITPROSTAGETYPEHASH() ([32]byte, error) {
	return _DeOrder.Contract.PERMITPROSTAGETYPEHASH(&_DeOrder.CallOpts)
}

// PERMITPROSTAGETYPEHASH is a free data retrieval call binding the contract method 0x31f3a62c.
//
// Solidity: function PERMITPROSTAGE_TYPEHASH() view returns(bytes32)
func (_DeOrder *DeOrderCallerSession) PERMITPROSTAGETYPEHASH() ([32]byte, error) {
	return _DeOrder.Contract.PERMITPROSTAGETYPEHASH(&_DeOrder.CallOpts)
}

// PERMITSTAGETYPEHASH is a free data retrieval call binding the contract method 0xbc4491c2.
//
// Solidity: function PERMITSTAGE_TYPEHASH() view returns(bytes32)
func (_DeOrder *DeOrderCaller) PERMITSTAGETYPEHASH(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _DeOrder.contract.Call(opts, &out, "PERMITSTAGE_TYPEHASH")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// PERMITSTAGETYPEHASH is a free data retrieval call binding the contract method 0xbc4491c2.
//
// Solidity: function PERMITSTAGE_TYPEHASH() view returns(bytes32)
func (_DeOrder *DeOrderSession) PERMITSTAGETYPEHASH() ([32]byte, error) {
	return _DeOrder.Contract.PERMITSTAGETYPEHASH(&_DeOrder.CallOpts)
}

// PERMITSTAGETYPEHASH is a free data retrieval call binding the contract method 0xbc4491c2.
//
// Solidity: function PERMITSTAGE_TYPEHASH() view returns(bytes32)
func (_DeOrder *DeOrderCallerSession) PERMITSTAGETYPEHASH() ([32]byte, error) {
	return _DeOrder.Contract.PERMITSTAGETYPEHASH(&_DeOrder.CallOpts)
}

// WETH is a free data retrieval call binding the contract method 0xad5c4648.
//
// Solidity: function WETH() view returns(address)
func (_DeOrder *DeOrderCaller) WETH(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _DeOrder.contract.Call(opts, &out, "WETH")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// WETH is a free data retrieval call binding the contract method 0xad5c4648.
//
// Solidity: function WETH() view returns(address)
func (_DeOrder *DeOrderSession) WETH() (common.Address, error) {
	return _DeOrder.Contract.WETH(&_DeOrder.CallOpts)
}

// WETH is a free data retrieval call binding the contract method 0xad5c4648.
//
// Solidity: function WETH() view returns(address)
func (_DeOrder *DeOrderCallerSession) WETH() (common.Address, error) {
	return _DeOrder.Contract.WETH(&_DeOrder.CallOpts)
}

// BuilderSBT is a free data retrieval call binding the contract method 0x7c1b112c.
//
// Solidity: function builderSBT() view returns(address)
func (_DeOrder *DeOrderCaller) BuilderSBT(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _DeOrder.contract.Call(opts, &out, "builderSBT")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// BuilderSBT is a free data retrieval call binding the contract method 0x7c1b112c.
//
// Solidity: function builderSBT() view returns(address)
func (_DeOrder *DeOrderSession) BuilderSBT() (common.Address, error) {
	return _DeOrder.Contract.BuilderSBT(&_DeOrder.CallOpts)
}

// BuilderSBT is a free data retrieval call binding the contract method 0x7c1b112c.
//
// Solidity: function builderSBT() view returns(address)
func (_DeOrder *DeOrderCallerSession) BuilderSBT() (common.Address, error) {
	return _DeOrder.Contract.BuilderSBT(&_DeOrder.CallOpts)
}

// CurrOrderId is a free data retrieval call binding the contract method 0x7d96a9b9.
//
// Solidity: function currOrderId() view returns(uint256)
func (_DeOrder *DeOrderCaller) CurrOrderId(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _DeOrder.contract.Call(opts, &out, "currOrderId")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// CurrOrderId is a free data retrieval call binding the contract method 0x7d96a9b9.
//
// Solidity: function currOrderId() view returns(uint256)
func (_DeOrder *DeOrderSession) CurrOrderId() (*big.Int, error) {
	return _DeOrder.Contract.CurrOrderId(&_DeOrder.CallOpts)
}

// CurrOrderId is a free data retrieval call binding the contract method 0x7d96a9b9.
//
// Solidity: function currOrderId() view returns(uint256)
func (_DeOrder *DeOrderCallerSession) CurrOrderId() (*big.Int, error) {
	return _DeOrder.Contract.CurrOrderId(&_DeOrder.CallOpts)
}

// Fee is a free data retrieval call binding the contract method 0xddca3f43.
//
// Solidity: function fee() view returns(uint256)
func (_DeOrder *DeOrderCaller) Fee(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _DeOrder.contract.Call(opts, &out, "fee")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Fee is a free data retrieval call binding the contract method 0xddca3f43.
//
// Solidity: function fee() view returns(uint256)
func (_DeOrder *DeOrderSession) Fee() (*big.Int, error) {
	return _DeOrder.Contract.Fee(&_DeOrder.CallOpts)
}

// Fee is a free data retrieval call binding the contract method 0xddca3f43.
//
// Solidity: function fee() view returns(uint256)
func (_DeOrder *DeOrderCallerSession) Fee() (*big.Int, error) {
	return _DeOrder.Contract.Fee(&_DeOrder.CallOpts)
}

// FeeTo is a free data retrieval call binding the contract method 0x017e7e58.
//
// Solidity: function feeTo() view returns(address)
func (_DeOrder *DeOrderCaller) FeeTo(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _DeOrder.contract.Call(opts, &out, "feeTo")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// FeeTo is a free data retrieval call binding the contract method 0x017e7e58.
//
// Solidity: function feeTo() view returns(address)
func (_DeOrder *DeOrderSession) FeeTo() (common.Address, error) {
	return _DeOrder.Contract.FeeTo(&_DeOrder.CallOpts)
}

// FeeTo is a free data retrieval call binding the contract method 0x017e7e58.
//
// Solidity: function feeTo() view returns(address)
func (_DeOrder *DeOrderCallerSession) FeeTo() (common.Address, error) {
	return _DeOrder.Contract.FeeTo(&_DeOrder.CallOpts)
}

// GetOrder is a free data retrieval call binding the contract method 0xd09ef241.
//
// Solidity: function getOrder(uint256 orderId) view returns((uint256,address,uint96,address,uint96,address,uint8,uint8,uint32))
func (_DeOrder *DeOrderCaller) GetOrder(opts *bind.CallOpts, orderId *big.Int) (Order, error) {
	var out []interface{}
	err := _DeOrder.contract.Call(opts, &out, "getOrder", orderId)

	if err != nil {
		return *new(Order), err
	}

	out0 := *abi.ConvertType(out[0], new(Order)).(*Order)

	return out0, err

}

// GetOrder is a free data retrieval call binding the contract method 0xd09ef241.
//
// Solidity: function getOrder(uint256 orderId) view returns((uint256,address,uint96,address,uint96,address,uint8,uint8,uint32))
func (_DeOrder *DeOrderSession) GetOrder(orderId *big.Int) (Order, error) {
	return _DeOrder.Contract.GetOrder(&_DeOrder.CallOpts, orderId)
}

// GetOrder is a free data retrieval call binding the contract method 0xd09ef241.
//
// Solidity: function getOrder(uint256 orderId) view returns((uint256,address,uint96,address,uint96,address,uint8,uint8,uint32))
func (_DeOrder *DeOrderCallerSession) GetOrder(orderId *big.Int) (Order, error) {
	return _DeOrder.Contract.GetOrder(&_DeOrder.CallOpts, orderId)
}

// GetStages is a free data retrieval call binding the contract method 0xef8cccb0.
//
// Solidity: function getStages(uint256 _orderId) view returns((uint96,uint32,uint8)[] stages)
func (_DeOrder *DeOrderCaller) GetStages(opts *bind.CallOpts, _orderId *big.Int) ([]DeStageStage, error) {
	var out []interface{}
	err := _DeOrder.contract.Call(opts, &out, "getStages", _orderId)

	if err != nil {
		return *new([]DeStageStage), err
	}

	out0 := *abi.ConvertType(out[0], new([]DeStageStage)).(*[]DeStageStage)

	return out0, err

}

// GetStages is a free data retrieval call binding the contract method 0xef8cccb0.
//
// Solidity: function getStages(uint256 _orderId) view returns((uint96,uint32,uint8)[] stages)
func (_DeOrder *DeOrderSession) GetStages(_orderId *big.Int) ([]DeStageStage, error) {
	return _DeOrder.Contract.GetStages(&_DeOrder.CallOpts, _orderId)
}

// GetStages is a free data retrieval call binding the contract method 0xef8cccb0.
//
// Solidity: function getStages(uint256 _orderId) view returns((uint96,uint32,uint8)[] stages)
func (_DeOrder *DeOrderCallerSession) GetStages(_orderId *big.Int) ([]DeStageStage, error) {
	return _DeOrder.Contract.GetStages(&_DeOrder.CallOpts, _orderId)
}

// IssuerSBT is a free data retrieval call binding the contract method 0xe23590ef.
//
// Solidity: function issuerSBT() view returns(address)
func (_DeOrder *DeOrderCaller) IssuerSBT(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _DeOrder.contract.Call(opts, &out, "issuerSBT")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// IssuerSBT is a free data retrieval call binding the contract method 0xe23590ef.
//
// Solidity: function issuerSBT() view returns(address)
func (_DeOrder *DeOrderSession) IssuerSBT() (common.Address, error) {
	return _DeOrder.Contract.IssuerSBT(&_DeOrder.CallOpts)
}

// IssuerSBT is a free data retrieval call binding the contract method 0xe23590ef.
//
// Solidity: function issuerSBT() view returns(address)
func (_DeOrder *DeOrderCallerSession) IssuerSBT() (common.Address, error) {
	return _DeOrder.Contract.IssuerSBT(&_DeOrder.CallOpts)
}

// Nonces is a free data retrieval call binding the contract method 0x502e1a16.
//
// Solidity: function nonces(address , uint256 ) view returns(uint256)
func (_DeOrder *DeOrderCaller) Nonces(opts *bind.CallOpts, arg0 common.Address, arg1 *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _DeOrder.contract.Call(opts, &out, "nonces", arg0, arg1)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Nonces is a free data retrieval call binding the contract method 0x502e1a16.
//
// Solidity: function nonces(address , uint256 ) view returns(uint256)
func (_DeOrder *DeOrderSession) Nonces(arg0 common.Address, arg1 *big.Int) (*big.Int, error) {
	return _DeOrder.Contract.Nonces(&_DeOrder.CallOpts, arg0, arg1)
}

// Nonces is a free data retrieval call binding the contract method 0x502e1a16.
//
// Solidity: function nonces(address , uint256 ) view returns(uint256)
func (_DeOrder *DeOrderCallerSession) Nonces(arg0 common.Address, arg1 *big.Int) (*big.Int, error) {
	return _DeOrder.Contract.Nonces(&_DeOrder.CallOpts, arg0, arg1)
}

// OngoingStage is a free data retrieval call binding the contract method 0x59ac59eb.
//
// Solidity: function ongoingStage(uint256 _orderId) view returns(uint256 stageIndex, uint256 stageStartDate)
func (_DeOrder *DeOrderCaller) OngoingStage(opts *bind.CallOpts, _orderId *big.Int) (struct {
	StageIndex     *big.Int
	StageStartDate *big.Int
}, error) {
	var out []interface{}
	err := _DeOrder.contract.Call(opts, &out, "ongoingStage", _orderId)

	outstruct := new(struct {
		StageIndex     *big.Int
		StageStartDate *big.Int
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.StageIndex = *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)
	outstruct.StageStartDate = *abi.ConvertType(out[1], new(*big.Int)).(**big.Int)

	return *outstruct, err

}

// OngoingStage is a free data retrieval call binding the contract method 0x59ac59eb.
//
// Solidity: function ongoingStage(uint256 _orderId) view returns(uint256 stageIndex, uint256 stageStartDate)
func (_DeOrder *DeOrderSession) OngoingStage(_orderId *big.Int) (struct {
	StageIndex     *big.Int
	StageStartDate *big.Int
}, error) {
	return _DeOrder.Contract.OngoingStage(&_DeOrder.CallOpts, _orderId)
}

// OngoingStage is a free data retrieval call binding the contract method 0x59ac59eb.
//
// Solidity: function ongoingStage(uint256 _orderId) view returns(uint256 stageIndex, uint256 stageStartDate)
func (_DeOrder *DeOrderCallerSession) OngoingStage(_orderId *big.Int) (struct {
	StageIndex     *big.Int
	StageStartDate *big.Int
}, error) {
	return _DeOrder.Contract.OngoingStage(&_DeOrder.CallOpts, _orderId)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_DeOrder *DeOrderCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _DeOrder.contract.Call(opts, &out, "owner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_DeOrder *DeOrderSession) Owner() (common.Address, error) {
	return _DeOrder.Contract.Owner(&_DeOrder.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_DeOrder *DeOrderCallerSession) Owner() (common.Address, error) {
	return _DeOrder.Contract.Owner(&_DeOrder.CallOpts)
}

// PendingWithdraw is a free data retrieval call binding the contract method 0x2d72ac53.
//
// Solidity: function pendingWithdraw(uint256 _orderId) view returns(uint256 pending, uint256 nextStage)
func (_DeOrder *DeOrderCaller) PendingWithdraw(opts *bind.CallOpts, _orderId *big.Int) (struct {
	Pending   *big.Int
	NextStage *big.Int
}, error) {
	var out []interface{}
	err := _DeOrder.contract.Call(opts, &out, "pendingWithdraw", _orderId)

	outstruct := new(struct {
		Pending   *big.Int
		NextStage *big.Int
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.Pending = *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)
	outstruct.NextStage = *abi.ConvertType(out[1], new(*big.Int)).(**big.Int)

	return *outstruct, err

}

// PendingWithdraw is a free data retrieval call binding the contract method 0x2d72ac53.
//
// Solidity: function pendingWithdraw(uint256 _orderId) view returns(uint256 pending, uint256 nextStage)
func (_DeOrder *DeOrderSession) PendingWithdraw(_orderId *big.Int) (struct {
	Pending   *big.Int
	NextStage *big.Int
}, error) {
	return _DeOrder.Contract.PendingWithdraw(&_DeOrder.CallOpts, _orderId)
}

// PendingWithdraw is a free data retrieval call binding the contract method 0x2d72ac53.
//
// Solidity: function pendingWithdraw(uint256 _orderId) view returns(uint256 pending, uint256 nextStage)
func (_DeOrder *DeOrderCallerSession) PendingWithdraw(_orderId *big.Int) (struct {
	Pending   *big.Int
	NextStage *big.Int
}, error) {
	return _DeOrder.Contract.PendingWithdraw(&_DeOrder.CallOpts, _orderId)
}

// StageTotalAmount is a free data retrieval call binding the contract method 0x16c14a78.
//
// Solidity: function stageTotalAmount(uint256 orderId) view returns(uint256 total)
func (_DeOrder *DeOrderCaller) StageTotalAmount(opts *bind.CallOpts, orderId *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _DeOrder.contract.Call(opts, &out, "stageTotalAmount", orderId)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// StageTotalAmount is a free data retrieval call binding the contract method 0x16c14a78.
//
// Solidity: function stageTotalAmount(uint256 orderId) view returns(uint256 total)
func (_DeOrder *DeOrderSession) StageTotalAmount(orderId *big.Int) (*big.Int, error) {
	return _DeOrder.Contract.StageTotalAmount(&_DeOrder.CallOpts, orderId)
}

// StageTotalAmount is a free data retrieval call binding the contract method 0x16c14a78.
//
// Solidity: function stageTotalAmount(uint256 orderId) view returns(uint256 total)
func (_DeOrder *DeOrderCallerSession) StageTotalAmount(orderId *big.Int) (*big.Int, error) {
	return _DeOrder.Contract.StageTotalAmount(&_DeOrder.CallOpts, orderId)
}

// StagesLength is a free data retrieval call binding the contract method 0xed1f751e.
//
// Solidity: function stagesLength(uint256 orderId) view returns(uint256 len)
func (_DeOrder *DeOrderCaller) StagesLength(opts *bind.CallOpts, orderId *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _DeOrder.contract.Call(opts, &out, "stagesLength", orderId)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// StagesLength is a free data retrieval call binding the contract method 0xed1f751e.
//
// Solidity: function stagesLength(uint256 orderId) view returns(uint256 len)
func (_DeOrder *DeOrderSession) StagesLength(orderId *big.Int) (*big.Int, error) {
	return _DeOrder.Contract.StagesLength(&_DeOrder.CallOpts, orderId)
}

// StagesLength is a free data retrieval call binding the contract method 0xed1f751e.
//
// Solidity: function stagesLength(uint256 orderId) view returns(uint256 len)
func (_DeOrder *DeOrderCallerSession) StagesLength(orderId *big.Int) (*big.Int, error) {
	return _DeOrder.Contract.StagesLength(&_DeOrder.CallOpts, orderId)
}

// SupportTokens is a free data retrieval call binding the contract method 0x8a0737a0.
//
// Solidity: function supportTokens(address ) view returns(bool)
func (_DeOrder *DeOrderCaller) SupportTokens(opts *bind.CallOpts, arg0 common.Address) (bool, error) {
	var out []interface{}
	err := _DeOrder.contract.Call(opts, &out, "supportTokens", arg0)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// SupportTokens is a free data retrieval call binding the contract method 0x8a0737a0.
//
// Solidity: function supportTokens(address ) view returns(bool)
func (_DeOrder *DeOrderSession) SupportTokens(arg0 common.Address) (bool, error) {
	return _DeOrder.Contract.SupportTokens(&_DeOrder.CallOpts, arg0)
}

// SupportTokens is a free data retrieval call binding the contract method 0x8a0737a0.
//
// Solidity: function supportTokens(address ) view returns(bool)
func (_DeOrder *DeOrderCallerSession) SupportTokens(arg0 common.Address) (bool, error) {
	return _DeOrder.Contract.SupportTokens(&_DeOrder.CallOpts, arg0)
}

// TotalStagePeriod is a free data retrieval call binding the contract method 0xb7014608.
//
// Solidity: function totalStagePeriod(uint256 orderId) view returns(uint256 total)
func (_DeOrder *DeOrderCaller) TotalStagePeriod(opts *bind.CallOpts, orderId *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _DeOrder.contract.Call(opts, &out, "totalStagePeriod", orderId)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// TotalStagePeriod is a free data retrieval call binding the contract method 0xb7014608.
//
// Solidity: function totalStagePeriod(uint256 orderId) view returns(uint256 total)
func (_DeOrder *DeOrderSession) TotalStagePeriod(orderId *big.Int) (*big.Int, error) {
	return _DeOrder.Contract.TotalStagePeriod(&_DeOrder.CallOpts, orderId)
}

// TotalStagePeriod is a free data retrieval call binding the contract method 0xb7014608.
//
// Solidity: function totalStagePeriod(uint256 orderId) view returns(uint256 total)
func (_DeOrder *DeOrderCallerSession) TotalStagePeriod(orderId *big.Int) (*big.Int, error) {
	return _DeOrder.Contract.TotalStagePeriod(&_DeOrder.CallOpts, orderId)
}

// AbortOrder is a paid mutator transaction binding the contract method 0x4122a29e.
//
// Solidity: function abortOrder(uint256 _orderId) returns()
func (_DeOrder *DeOrderTransactor) AbortOrder(opts *bind.TransactOpts, _orderId *big.Int) (*types.Transaction, error) {
	return _DeOrder.contract.Transact(opts, "abortOrder", _orderId)
}

// AbortOrder is a paid mutator transaction binding the contract method 0x4122a29e.
//
// Solidity: function abortOrder(uint256 _orderId) returns()
func (_DeOrder *DeOrderSession) AbortOrder(_orderId *big.Int) (*types.Transaction, error) {
	return _DeOrder.Contract.AbortOrder(&_DeOrder.TransactOpts, _orderId)
}

// AbortOrder is a paid mutator transaction binding the contract method 0x4122a29e.
//
// Solidity: function abortOrder(uint256 _orderId) returns()
func (_DeOrder *DeOrderTransactorSession) AbortOrder(_orderId *big.Int) (*types.Transaction, error) {
	return _DeOrder.Contract.AbortOrder(&_DeOrder.TransactOpts, _orderId)
}

// AppendStage is a paid mutator transaction binding the contract method 0xba546ccf.
//
// Solidity: function appendStage(uint256 _orderId, uint256 amount, uint256 period, uint256 nonce, uint256 deadline, uint8 v, bytes32 r, bytes32 s) payable returns()
func (_DeOrder *DeOrderTransactor) AppendStage(opts *bind.TransactOpts, _orderId *big.Int, amount *big.Int, period *big.Int, nonce *big.Int, deadline *big.Int, v uint8, r [32]byte, s [32]byte) (*types.Transaction, error) {
	return _DeOrder.contract.Transact(opts, "appendStage", _orderId, amount, period, nonce, deadline, v, r, s)
}

// AppendStage is a paid mutator transaction binding the contract method 0xba546ccf.
//
// Solidity: function appendStage(uint256 _orderId, uint256 amount, uint256 period, uint256 nonce, uint256 deadline, uint8 v, bytes32 r, bytes32 s) payable returns()
func (_DeOrder *DeOrderSession) AppendStage(_orderId *big.Int, amount *big.Int, period *big.Int, nonce *big.Int, deadline *big.Int, v uint8, r [32]byte, s [32]byte) (*types.Transaction, error) {
	return _DeOrder.Contract.AppendStage(&_DeOrder.TransactOpts, _orderId, amount, period, nonce, deadline, v, r, s)
}

// AppendStage is a paid mutator transaction binding the contract method 0xba546ccf.
//
// Solidity: function appendStage(uint256 _orderId, uint256 amount, uint256 period, uint256 nonce, uint256 deadline, uint8 v, bytes32 r, bytes32 s) payable returns()
func (_DeOrder *DeOrderTransactorSession) AppendStage(_orderId *big.Int, amount *big.Int, period *big.Int, nonce *big.Int, deadline *big.Int, v uint8, r [32]byte, s [32]byte) (*types.Transaction, error) {
	return _DeOrder.Contract.AppendStage(&_DeOrder.TransactOpts, _orderId, amount, period, nonce, deadline, v, r, s)
}

// ConfirmDelivery is a paid mutator transaction binding the contract method 0x386f586d.
//
// Solidity: function confirmDelivery(uint256 _orderId, uint256[] _stageIndexs) returns()
func (_DeOrder *DeOrderTransactor) ConfirmDelivery(opts *bind.TransactOpts, _orderId *big.Int, _stageIndexs []*big.Int) (*types.Transaction, error) {
	return _DeOrder.contract.Transact(opts, "confirmDelivery", _orderId, _stageIndexs)
}

// ConfirmDelivery is a paid mutator transaction binding the contract method 0x386f586d.
//
// Solidity: function confirmDelivery(uint256 _orderId, uint256[] _stageIndexs) returns()
func (_DeOrder *DeOrderSession) ConfirmDelivery(_orderId *big.Int, _stageIndexs []*big.Int) (*types.Transaction, error) {
	return _DeOrder.Contract.ConfirmDelivery(&_DeOrder.TransactOpts, _orderId, _stageIndexs)
}

// ConfirmDelivery is a paid mutator transaction binding the contract method 0x386f586d.
//
// Solidity: function confirmDelivery(uint256 _orderId, uint256[] _stageIndexs) returns()
func (_DeOrder *DeOrderTransactorSession) ConfirmDelivery(_orderId *big.Int, _stageIndexs []*big.Int) (*types.Transaction, error) {
	return _DeOrder.Contract.ConfirmDelivery(&_DeOrder.TransactOpts, _orderId, _stageIndexs)
}

// CreateOrder is a paid mutator transaction binding the contract method 0x4a25fb12.
//
// Solidity: function createOrder(uint256 _taskId, address _issuer, address _worker, address _token, uint256 _amount) payable returns()
func (_DeOrder *DeOrderTransactor) CreateOrder(opts *bind.TransactOpts, _taskId *big.Int, _issuer common.Address, _worker common.Address, _token common.Address, _amount *big.Int) (*types.Transaction, error) {
	return _DeOrder.contract.Transact(opts, "createOrder", _taskId, _issuer, _worker, _token, _amount)
}

// CreateOrder is a paid mutator transaction binding the contract method 0x4a25fb12.
//
// Solidity: function createOrder(uint256 _taskId, address _issuer, address _worker, address _token, uint256 _amount) payable returns()
func (_DeOrder *DeOrderSession) CreateOrder(_taskId *big.Int, _issuer common.Address, _worker common.Address, _token common.Address, _amount *big.Int) (*types.Transaction, error) {
	return _DeOrder.Contract.CreateOrder(&_DeOrder.TransactOpts, _taskId, _issuer, _worker, _token, _amount)
}

// CreateOrder is a paid mutator transaction binding the contract method 0x4a25fb12.
//
// Solidity: function createOrder(uint256 _taskId, address _issuer, address _worker, address _token, uint256 _amount) payable returns()
func (_DeOrder *DeOrderTransactorSession) CreateOrder(_taskId *big.Int, _issuer common.Address, _worker common.Address, _token common.Address, _amount *big.Int) (*types.Transaction, error) {
	return _DeOrder.Contract.CreateOrder(&_DeOrder.TransactOpts, _taskId, _issuer, _worker, _token, _amount)
}

// ModifyOrder is a paid mutator transaction binding the contract method 0x59bfb274.
//
// Solidity: function modifyOrder(uint256 orderId, address token, uint256 amount) payable returns()
func (_DeOrder *DeOrderTransactor) ModifyOrder(opts *bind.TransactOpts, orderId *big.Int, token common.Address, amount *big.Int) (*types.Transaction, error) {
	return _DeOrder.contract.Transact(opts, "modifyOrder", orderId, token, amount)
}

// ModifyOrder is a paid mutator transaction binding the contract method 0x59bfb274.
//
// Solidity: function modifyOrder(uint256 orderId, address token, uint256 amount) payable returns()
func (_DeOrder *DeOrderSession) ModifyOrder(orderId *big.Int, token common.Address, amount *big.Int) (*types.Transaction, error) {
	return _DeOrder.Contract.ModifyOrder(&_DeOrder.TransactOpts, orderId, token, amount)
}

// ModifyOrder is a paid mutator transaction binding the contract method 0x59bfb274.
//
// Solidity: function modifyOrder(uint256 orderId, address token, uint256 amount) payable returns()
func (_DeOrder *DeOrderTransactorSession) ModifyOrder(orderId *big.Int, token common.Address, amount *big.Int) (*types.Transaction, error) {
	return _DeOrder.Contract.ModifyOrder(&_DeOrder.TransactOpts, orderId, token, amount)
}

// Multicall is a paid mutator transaction binding the contract method 0xac9650d8.
//
// Solidity: function multicall(bytes[] data) payable returns(bytes[] results)
func (_DeOrder *DeOrderTransactor) Multicall(opts *bind.TransactOpts, data [][]byte) (*types.Transaction, error) {
	return _DeOrder.contract.Transact(opts, "multicall", data)
}

// Multicall is a paid mutator transaction binding the contract method 0xac9650d8.
//
// Solidity: function multicall(bytes[] data) payable returns(bytes[] results)
func (_DeOrder *DeOrderSession) Multicall(data [][]byte) (*types.Transaction, error) {
	return _DeOrder.Contract.Multicall(&_DeOrder.TransactOpts, data)
}

// Multicall is a paid mutator transaction binding the contract method 0xac9650d8.
//
// Solidity: function multicall(bytes[] data) payable returns(bytes[] results)
func (_DeOrder *DeOrderTransactorSession) Multicall(data [][]byte) (*types.Transaction, error) {
	return _DeOrder.Contract.Multicall(&_DeOrder.TransactOpts, data)
}

// PayOrder is a paid mutator transaction binding the contract method 0xd2fbd0ed.
//
// Solidity: function payOrder(uint256 orderId, uint256 amount) payable returns()
func (_DeOrder *DeOrderTransactor) PayOrder(opts *bind.TransactOpts, orderId *big.Int, amount *big.Int) (*types.Transaction, error) {
	return _DeOrder.contract.Transact(opts, "payOrder", orderId, amount)
}

// PayOrder is a paid mutator transaction binding the contract method 0xd2fbd0ed.
//
// Solidity: function payOrder(uint256 orderId, uint256 amount) payable returns()
func (_DeOrder *DeOrderSession) PayOrder(orderId *big.Int, amount *big.Int) (*types.Transaction, error) {
	return _DeOrder.Contract.PayOrder(&_DeOrder.TransactOpts, orderId, amount)
}

// PayOrder is a paid mutator transaction binding the contract method 0xd2fbd0ed.
//
// Solidity: function payOrder(uint256 orderId, uint256 amount) payable returns()
func (_DeOrder *DeOrderTransactorSession) PayOrder(orderId *big.Int, amount *big.Int) (*types.Transaction, error) {
	return _DeOrder.Contract.PayOrder(&_DeOrder.TransactOpts, orderId, amount)
}

// PayOrderWithPermit is a paid mutator transaction binding the contract method 0x1ee7d8d4.
//
// Solidity: function payOrderWithPermit(uint256 orderId, uint256 amount, uint256 deadline, uint8 v, bytes32 r, bytes32 s) returns()
func (_DeOrder *DeOrderTransactor) PayOrderWithPermit(opts *bind.TransactOpts, orderId *big.Int, amount *big.Int, deadline *big.Int, v uint8, r [32]byte, s [32]byte) (*types.Transaction, error) {
	return _DeOrder.contract.Transact(opts, "payOrderWithPermit", orderId, amount, deadline, v, r, s)
}

// PayOrderWithPermit is a paid mutator transaction binding the contract method 0x1ee7d8d4.
//
// Solidity: function payOrderWithPermit(uint256 orderId, uint256 amount, uint256 deadline, uint8 v, bytes32 r, bytes32 s) returns()
func (_DeOrder *DeOrderSession) PayOrderWithPermit(orderId *big.Int, amount *big.Int, deadline *big.Int, v uint8, r [32]byte, s [32]byte) (*types.Transaction, error) {
	return _DeOrder.Contract.PayOrderWithPermit(&_DeOrder.TransactOpts, orderId, amount, deadline, v, r, s)
}

// PayOrderWithPermit is a paid mutator transaction binding the contract method 0x1ee7d8d4.
//
// Solidity: function payOrderWithPermit(uint256 orderId, uint256 amount, uint256 deadline, uint8 v, bytes32 r, bytes32 s) returns()
func (_DeOrder *DeOrderTransactorSession) PayOrderWithPermit(orderId *big.Int, amount *big.Int, deadline *big.Int, v uint8, r [32]byte, s [32]byte) (*types.Transaction, error) {
	return _DeOrder.Contract.PayOrderWithPermit(&_DeOrder.TransactOpts, orderId, amount, deadline, v, r, s)
}

// PayOrderWithPermit2 is a paid mutator transaction binding the contract method 0xedb86978.
//
// Solidity: function payOrderWithPermit2(uint256 orderId, uint256 amount, ((address,uint256),uint256,uint256) permit, bytes signature) returns()
func (_DeOrder *DeOrderTransactor) PayOrderWithPermit2(opts *bind.TransactOpts, orderId *big.Int, amount *big.Int, permit IPermit2PermitTransferFrom, signature []byte) (*types.Transaction, error) {
	return _DeOrder.contract.Transact(opts, "payOrderWithPermit2", orderId, amount, permit, signature)
}

// PayOrderWithPermit2 is a paid mutator transaction binding the contract method 0xedb86978.
//
// Solidity: function payOrderWithPermit2(uint256 orderId, uint256 amount, ((address,uint256),uint256,uint256) permit, bytes signature) returns()
func (_DeOrder *DeOrderSession) PayOrderWithPermit2(orderId *big.Int, amount *big.Int, permit IPermit2PermitTransferFrom, signature []byte) (*types.Transaction, error) {
	return _DeOrder.Contract.PayOrderWithPermit2(&_DeOrder.TransactOpts, orderId, amount, permit, signature)
}

// PayOrderWithPermit2 is a paid mutator transaction binding the contract method 0xedb86978.
//
// Solidity: function payOrderWithPermit2(uint256 orderId, uint256 amount, ((address,uint256),uint256,uint256) permit, bytes signature) returns()
func (_DeOrder *DeOrderTransactorSession) PayOrderWithPermit2(orderId *big.Int, amount *big.Int, permit IPermit2PermitTransferFrom, signature []byte) (*types.Transaction, error) {
	return _DeOrder.Contract.PayOrderWithPermit2(&_DeOrder.TransactOpts, orderId, amount, permit, signature)
}

// PermitStage is a paid mutator transaction binding the contract method 0xe626cf89.
//
// Solidity: function permitStage(uint256 _orderId, uint256[] _amounts, uint256[] _periods, uint8 payType, uint256 nonce, uint256 deadline, uint8 v, bytes32 r, bytes32 s) payable returns()
func (_DeOrder *DeOrderTransactor) PermitStage(opts *bind.TransactOpts, _orderId *big.Int, _amounts []*big.Int, _periods []*big.Int, payType uint8, nonce *big.Int, deadline *big.Int, v uint8, r [32]byte, s [32]byte) (*types.Transaction, error) {
	return _DeOrder.contract.Transact(opts, "permitStage", _orderId, _amounts, _periods, payType, nonce, deadline, v, r, s)
}

// PermitStage is a paid mutator transaction binding the contract method 0xe626cf89.
//
// Solidity: function permitStage(uint256 _orderId, uint256[] _amounts, uint256[] _periods, uint8 payType, uint256 nonce, uint256 deadline, uint8 v, bytes32 r, bytes32 s) payable returns()
func (_DeOrder *DeOrderSession) PermitStage(_orderId *big.Int, _amounts []*big.Int, _periods []*big.Int, payType uint8, nonce *big.Int, deadline *big.Int, v uint8, r [32]byte, s [32]byte) (*types.Transaction, error) {
	return _DeOrder.Contract.PermitStage(&_DeOrder.TransactOpts, _orderId, _amounts, _periods, payType, nonce, deadline, v, r, s)
}

// PermitStage is a paid mutator transaction binding the contract method 0xe626cf89.
//
// Solidity: function permitStage(uint256 _orderId, uint256[] _amounts, uint256[] _periods, uint8 payType, uint256 nonce, uint256 deadline, uint8 v, bytes32 r, bytes32 s) payable returns()
func (_DeOrder *DeOrderTransactorSession) PermitStage(_orderId *big.Int, _amounts []*big.Int, _periods []*big.Int, payType uint8, nonce *big.Int, deadline *big.Int, v uint8, r [32]byte, s [32]byte) (*types.Transaction, error) {
	return _DeOrder.Contract.PermitStage(&_DeOrder.TransactOpts, _orderId, _amounts, _periods, payType, nonce, deadline, v, r, s)
}

// ProlongStage is a paid mutator transaction binding the contract method 0xfccfac9d.
//
// Solidity: function prolongStage(uint256 _orderId, uint256 _stageIndex, uint256 _appendPeriod, uint256 nonce, uint256 deadline, uint8 v, bytes32 r, bytes32 s) returns()
func (_DeOrder *DeOrderTransactor) ProlongStage(opts *bind.TransactOpts, _orderId *big.Int, _stageIndex *big.Int, _appendPeriod *big.Int, nonce *big.Int, deadline *big.Int, v uint8, r [32]byte, s [32]byte) (*types.Transaction, error) {
	return _DeOrder.contract.Transact(opts, "prolongStage", _orderId, _stageIndex, _appendPeriod, nonce, deadline, v, r, s)
}

// ProlongStage is a paid mutator transaction binding the contract method 0xfccfac9d.
//
// Solidity: function prolongStage(uint256 _orderId, uint256 _stageIndex, uint256 _appendPeriod, uint256 nonce, uint256 deadline, uint8 v, bytes32 r, bytes32 s) returns()
func (_DeOrder *DeOrderSession) ProlongStage(_orderId *big.Int, _stageIndex *big.Int, _appendPeriod *big.Int, nonce *big.Int, deadline *big.Int, v uint8, r [32]byte, s [32]byte) (*types.Transaction, error) {
	return _DeOrder.Contract.ProlongStage(&_DeOrder.TransactOpts, _orderId, _stageIndex, _appendPeriod, nonce, deadline, v, r, s)
}

// ProlongStage is a paid mutator transaction binding the contract method 0xfccfac9d.
//
// Solidity: function prolongStage(uint256 _orderId, uint256 _stageIndex, uint256 _appendPeriod, uint256 nonce, uint256 deadline, uint8 v, bytes32 r, bytes32 s) returns()
func (_DeOrder *DeOrderTransactorSession) ProlongStage(_orderId *big.Int, _stageIndex *big.Int, _appendPeriod *big.Int, nonce *big.Int, deadline *big.Int, v uint8, r [32]byte, s [32]byte) (*types.Transaction, error) {
	return _DeOrder.Contract.ProlongStage(&_DeOrder.TransactOpts, _orderId, _stageIndex, _appendPeriod, nonce, deadline, v, r, s)
}

// Refund is a paid mutator transaction binding the contract method 0xe796a6eb.
//
// Solidity: function refund(uint256 _orderId, address _to, uint256 _amount) payable returns()
func (_DeOrder *DeOrderTransactor) Refund(opts *bind.TransactOpts, _orderId *big.Int, _to common.Address, _amount *big.Int) (*types.Transaction, error) {
	return _DeOrder.contract.Transact(opts, "refund", _orderId, _to, _amount)
}

// Refund is a paid mutator transaction binding the contract method 0xe796a6eb.
//
// Solidity: function refund(uint256 _orderId, address _to, uint256 _amount) payable returns()
func (_DeOrder *DeOrderSession) Refund(_orderId *big.Int, _to common.Address, _amount *big.Int) (*types.Transaction, error) {
	return _DeOrder.Contract.Refund(&_DeOrder.TransactOpts, _orderId, _to, _amount)
}

// Refund is a paid mutator transaction binding the contract method 0xe796a6eb.
//
// Solidity: function refund(uint256 _orderId, address _to, uint256 _amount) payable returns()
func (_DeOrder *DeOrderTransactorSession) Refund(_orderId *big.Int, _to common.Address, _amount *big.Int) (*types.Transaction, error) {
	return _DeOrder.Contract.Refund(&_DeOrder.TransactOpts, _orderId, _to, _amount)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_DeOrder *DeOrderTransactor) RenounceOwnership(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _DeOrder.contract.Transact(opts, "renounceOwnership")
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_DeOrder *DeOrderSession) RenounceOwnership() (*types.Transaction, error) {
	return _DeOrder.Contract.RenounceOwnership(&_DeOrder.TransactOpts)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_DeOrder *DeOrderTransactorSession) RenounceOwnership() (*types.Transaction, error) {
	return _DeOrder.Contract.RenounceOwnership(&_DeOrder.TransactOpts)
}

// SetFeeTo is a paid mutator transaction binding the contract method 0x384bb4a3.
//
// Solidity: function setFeeTo(uint256 _fee, address _feeTo) returns()
func (_DeOrder *DeOrderTransactor) SetFeeTo(opts *bind.TransactOpts, _fee *big.Int, _feeTo common.Address) (*types.Transaction, error) {
	return _DeOrder.contract.Transact(opts, "setFeeTo", _fee, _feeTo)
}

// SetFeeTo is a paid mutator transaction binding the contract method 0x384bb4a3.
//
// Solidity: function setFeeTo(uint256 _fee, address _feeTo) returns()
func (_DeOrder *DeOrderSession) SetFeeTo(_fee *big.Int, _feeTo common.Address) (*types.Transaction, error) {
	return _DeOrder.Contract.SetFeeTo(&_DeOrder.TransactOpts, _fee, _feeTo)
}

// SetFeeTo is a paid mutator transaction binding the contract method 0x384bb4a3.
//
// Solidity: function setFeeTo(uint256 _fee, address _feeTo) returns()
func (_DeOrder *DeOrderTransactorSession) SetFeeTo(_fee *big.Int, _feeTo common.Address) (*types.Transaction, error) {
	return _DeOrder.Contract.SetFeeTo(&_DeOrder.TransactOpts, _fee, _feeTo)
}

// SetSBT is a paid mutator transaction binding the contract method 0x6c406a9b.
//
// Solidity: function setSBT(address _builder, address _issuer) returns()
func (_DeOrder *DeOrderTransactor) SetSBT(opts *bind.TransactOpts, _builder common.Address, _issuer common.Address) (*types.Transaction, error) {
	return _DeOrder.contract.Transact(opts, "setSBT", _builder, _issuer)
}

// SetSBT is a paid mutator transaction binding the contract method 0x6c406a9b.
//
// Solidity: function setSBT(address _builder, address _issuer) returns()
func (_DeOrder *DeOrderSession) SetSBT(_builder common.Address, _issuer common.Address) (*types.Transaction, error) {
	return _DeOrder.Contract.SetSBT(&_DeOrder.TransactOpts, _builder, _issuer)
}

// SetSBT is a paid mutator transaction binding the contract method 0x6c406a9b.
//
// Solidity: function setSBT(address _builder, address _issuer) returns()
func (_DeOrder *DeOrderTransactorSession) SetSBT(_builder common.Address, _issuer common.Address) (*types.Transaction, error) {
	return _DeOrder.Contract.SetSBT(&_DeOrder.TransactOpts, _builder, _issuer)
}

// SetSupportToken is a paid mutator transaction binding the contract method 0x204daa93.
//
// Solidity: function setSupportToken(address _token, bool enable) returns()
func (_DeOrder *DeOrderTransactor) SetSupportToken(opts *bind.TransactOpts, _token common.Address, enable bool) (*types.Transaction, error) {
	return _DeOrder.contract.Transact(opts, "setSupportToken", _token, enable)
}

// SetSupportToken is a paid mutator transaction binding the contract method 0x204daa93.
//
// Solidity: function setSupportToken(address _token, bool enable) returns()
func (_DeOrder *DeOrderSession) SetSupportToken(_token common.Address, enable bool) (*types.Transaction, error) {
	return _DeOrder.Contract.SetSupportToken(&_DeOrder.TransactOpts, _token, enable)
}

// SetSupportToken is a paid mutator transaction binding the contract method 0x204daa93.
//
// Solidity: function setSupportToken(address _token, bool enable) returns()
func (_DeOrder *DeOrderTransactorSession) SetSupportToken(_token common.Address, enable bool) (*types.Transaction, error) {
	return _DeOrder.Contract.SetSupportToken(&_DeOrder.TransactOpts, _token, enable)
}

// StartOrder is a paid mutator transaction binding the contract method 0x7d41fdd9.
//
// Solidity: function startOrder(uint256 _orderId) payable returns()
func (_DeOrder *DeOrderTransactor) StartOrder(opts *bind.TransactOpts, _orderId *big.Int) (*types.Transaction, error) {
	return _DeOrder.contract.Transact(opts, "startOrder", _orderId)
}

// StartOrder is a paid mutator transaction binding the contract method 0x7d41fdd9.
//
// Solidity: function startOrder(uint256 _orderId) payable returns()
func (_DeOrder *DeOrderSession) StartOrder(_orderId *big.Int) (*types.Transaction, error) {
	return _DeOrder.Contract.StartOrder(&_DeOrder.TransactOpts, _orderId)
}

// StartOrder is a paid mutator transaction binding the contract method 0x7d41fdd9.
//
// Solidity: function startOrder(uint256 _orderId) payable returns()
func (_DeOrder *DeOrderTransactorSession) StartOrder(_orderId *big.Int) (*types.Transaction, error) {
	return _DeOrder.Contract.StartOrder(&_DeOrder.TransactOpts, _orderId)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_DeOrder *DeOrderTransactor) TransferOwnership(opts *bind.TransactOpts, newOwner common.Address) (*types.Transaction, error) {
	return _DeOrder.contract.Transact(opts, "transferOwnership", newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_DeOrder *DeOrderSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _DeOrder.Contract.TransferOwnership(&_DeOrder.TransactOpts, newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_DeOrder *DeOrderTransactorSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _DeOrder.Contract.TransferOwnership(&_DeOrder.TransactOpts, newOwner)
}

// UpdateAttachment is a paid mutator transaction binding the contract method 0x8e4736b2.
//
// Solidity: function updateAttachment(uint256 _orderId, string _attachment) returns()
func (_DeOrder *DeOrderTransactor) UpdateAttachment(opts *bind.TransactOpts, _orderId *big.Int, _attachment string) (*types.Transaction, error) {
	return _DeOrder.contract.Transact(opts, "updateAttachment", _orderId, _attachment)
}

// UpdateAttachment is a paid mutator transaction binding the contract method 0x8e4736b2.
//
// Solidity: function updateAttachment(uint256 _orderId, string _attachment) returns()
func (_DeOrder *DeOrderSession) UpdateAttachment(_orderId *big.Int, _attachment string) (*types.Transaction, error) {
	return _DeOrder.Contract.UpdateAttachment(&_DeOrder.TransactOpts, _orderId, _attachment)
}

// UpdateAttachment is a paid mutator transaction binding the contract method 0x8e4736b2.
//
// Solidity: function updateAttachment(uint256 _orderId, string _attachment) returns()
func (_DeOrder *DeOrderTransactorSession) UpdateAttachment(_orderId *big.Int, _attachment string) (*types.Transaction, error) {
	return _DeOrder.Contract.UpdateAttachment(&_DeOrder.TransactOpts, _orderId, _attachment)
}

// Withdraw is a paid mutator transaction binding the contract method 0x00f714ce.
//
// Solidity: function withdraw(uint256 _orderId, address to) returns()
func (_DeOrder *DeOrderTransactor) Withdraw(opts *bind.TransactOpts, _orderId *big.Int, to common.Address) (*types.Transaction, error) {
	return _DeOrder.contract.Transact(opts, "withdraw", _orderId, to)
}

// Withdraw is a paid mutator transaction binding the contract method 0x00f714ce.
//
// Solidity: function withdraw(uint256 _orderId, address to) returns()
func (_DeOrder *DeOrderSession) Withdraw(_orderId *big.Int, to common.Address) (*types.Transaction, error) {
	return _DeOrder.Contract.Withdraw(&_DeOrder.TransactOpts, _orderId, to)
}

// Withdraw is a paid mutator transaction binding the contract method 0x00f714ce.
//
// Solidity: function withdraw(uint256 _orderId, address to) returns()
func (_DeOrder *DeOrderTransactorSession) Withdraw(_orderId *big.Int, to common.Address) (*types.Transaction, error) {
	return _DeOrder.Contract.Withdraw(&_DeOrder.TransactOpts, _orderId, to)
}

// Receive is a paid mutator transaction binding the contract receive function.
//
// Solidity: receive() payable returns()
func (_DeOrder *DeOrderTransactor) Receive(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _DeOrder.contract.RawTransact(opts, nil) // calldata is disallowed for receive function
}

// Receive is a paid mutator transaction binding the contract receive function.
//
// Solidity: receive() payable returns()
func (_DeOrder *DeOrderSession) Receive() (*types.Transaction, error) {
	return _DeOrder.Contract.Receive(&_DeOrder.TransactOpts)
}

// Receive is a paid mutator transaction binding the contract receive function.
//
// Solidity: receive() payable returns()
func (_DeOrder *DeOrderTransactorSession) Receive() (*types.Transaction, error) {
	return _DeOrder.Contract.Receive(&_DeOrder.TransactOpts)
}

// DeOrderAppendStageIterator is returned from FilterAppendStage and is used to iterate over the raw logs and unpacked data for AppendStage events raised by the DeOrder contract.
type DeOrderAppendStageIterator struct {
	Event *DeOrderAppendStage // Event containing the contract specifics and raw log

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
func (it *DeOrderAppendStageIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(DeOrderAppendStage)
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
		it.Event = new(DeOrderAppendStage)
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
func (it *DeOrderAppendStageIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *DeOrderAppendStageIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// DeOrderAppendStage represents a AppendStage event raised by the DeOrder contract.
type DeOrderAppendStage struct {
	OrderId *big.Int
	Amount  *big.Int
	Period  *big.Int
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterAppendStage is a free log retrieval operation binding the contract event 0xfe21c1565116bcc6dd38b86e8ad093ad769866a2bda7ec261b0af259e5b980aa.
//
// Solidity: event AppendStage(uint256 indexed orderId, uint256 amount, uint256 period)
func (_DeOrder *DeOrderFilterer) FilterAppendStage(opts *bind.FilterOpts, orderId []*big.Int) (*DeOrderAppendStageIterator, error) {

	var orderIdRule []interface{}
	for _, orderIdItem := range orderId {
		orderIdRule = append(orderIdRule, orderIdItem)
	}

	logs, sub, err := _DeOrder.contract.FilterLogs(opts, "AppendStage", orderIdRule)
	if err != nil {
		return nil, err
	}
	return &DeOrderAppendStageIterator{contract: _DeOrder.contract, event: "AppendStage", logs: logs, sub: sub}, nil
}

// WatchAppendStage is a free log subscription operation binding the contract event 0xfe21c1565116bcc6dd38b86e8ad093ad769866a2bda7ec261b0af259e5b980aa.
//
// Solidity: event AppendStage(uint256 indexed orderId, uint256 amount, uint256 period)
func (_DeOrder *DeOrderFilterer) WatchAppendStage(opts *bind.WatchOpts, sink chan<- *DeOrderAppendStage, orderId []*big.Int) (event.Subscription, error) {

	var orderIdRule []interface{}
	for _, orderIdItem := range orderId {
		orderIdRule = append(orderIdRule, orderIdItem)
	}

	logs, sub, err := _DeOrder.contract.WatchLogs(opts, "AppendStage", orderIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(DeOrderAppendStage)
				if err := _DeOrder.contract.UnpackLog(event, "AppendStage", log); err != nil {
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

// ParseAppendStage is a log parse operation binding the contract event 0xfe21c1565116bcc6dd38b86e8ad093ad769866a2bda7ec261b0af259e5b980aa.
//
// Solidity: event AppendStage(uint256 indexed orderId, uint256 amount, uint256 period)
func (_DeOrder *DeOrderFilterer) ParseAppendStage(log types.Log) (*DeOrderAppendStage, error) {
	event := new(DeOrderAppendStage)
	if err := _DeOrder.contract.UnpackLog(event, "AppendStage", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// DeOrderAttachmentUpdatedIterator is returned from FilterAttachmentUpdated and is used to iterate over the raw logs and unpacked data for AttachmentUpdated events raised by the DeOrder contract.
type DeOrderAttachmentUpdatedIterator struct {
	Event *DeOrderAttachmentUpdated // Event containing the contract specifics and raw log

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
func (it *DeOrderAttachmentUpdatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(DeOrderAttachmentUpdated)
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
		it.Event = new(DeOrderAttachmentUpdated)
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
func (it *DeOrderAttachmentUpdatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *DeOrderAttachmentUpdatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// DeOrderAttachmentUpdated represents a AttachmentUpdated event raised by the DeOrder contract.
type DeOrderAttachmentUpdated struct {
	OrderId    *big.Int
	Attachment string
	Raw        types.Log // Blockchain specific contextual infos
}

// FilterAttachmentUpdated is a free log retrieval operation binding the contract event 0x9866b9bcafba6f147d29d348f70082ef6ebae7756608c1c6f2d093606b29a05e.
//
// Solidity: event AttachmentUpdated(uint256 indexed orderId, string attachment)
func (_DeOrder *DeOrderFilterer) FilterAttachmentUpdated(opts *bind.FilterOpts, orderId []*big.Int) (*DeOrderAttachmentUpdatedIterator, error) {

	var orderIdRule []interface{}
	for _, orderIdItem := range orderId {
		orderIdRule = append(orderIdRule, orderIdItem)
	}

	logs, sub, err := _DeOrder.contract.FilterLogs(opts, "AttachmentUpdated", orderIdRule)
	if err != nil {
		return nil, err
	}
	return &DeOrderAttachmentUpdatedIterator{contract: _DeOrder.contract, event: "AttachmentUpdated", logs: logs, sub: sub}, nil
}

// WatchAttachmentUpdated is a free log subscription operation binding the contract event 0x9866b9bcafba6f147d29d348f70082ef6ebae7756608c1c6f2d093606b29a05e.
//
// Solidity: event AttachmentUpdated(uint256 indexed orderId, string attachment)
func (_DeOrder *DeOrderFilterer) WatchAttachmentUpdated(opts *bind.WatchOpts, sink chan<- *DeOrderAttachmentUpdated, orderId []*big.Int) (event.Subscription, error) {

	var orderIdRule []interface{}
	for _, orderIdItem := range orderId {
		orderIdRule = append(orderIdRule, orderIdItem)
	}

	logs, sub, err := _DeOrder.contract.WatchLogs(opts, "AttachmentUpdated", orderIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(DeOrderAttachmentUpdated)
				if err := _DeOrder.contract.UnpackLog(event, "AttachmentUpdated", log); err != nil {
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

// ParseAttachmentUpdated is a log parse operation binding the contract event 0x9866b9bcafba6f147d29d348f70082ef6ebae7756608c1c6f2d093606b29a05e.
//
// Solidity: event AttachmentUpdated(uint256 indexed orderId, string attachment)
func (_DeOrder *DeOrderFilterer) ParseAttachmentUpdated(log types.Log) (*DeOrderAttachmentUpdated, error) {
	event := new(DeOrderAttachmentUpdated)
	if err := _DeOrder.contract.UnpackLog(event, "AttachmentUpdated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// DeOrderConfirmOrderStageIterator is returned from FilterConfirmOrderStage and is used to iterate over the raw logs and unpacked data for ConfirmOrderStage events raised by the DeOrder contract.
type DeOrderConfirmOrderStageIterator struct {
	Event *DeOrderConfirmOrderStage // Event containing the contract specifics and raw log

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
func (it *DeOrderConfirmOrderStageIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(DeOrderConfirmOrderStage)
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
		it.Event = new(DeOrderConfirmOrderStage)
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
func (it *DeOrderConfirmOrderStageIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *DeOrderConfirmOrderStageIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// DeOrderConfirmOrderStage represents a ConfirmOrderStage event raised by the DeOrder contract.
type DeOrderConfirmOrderStage struct {
	OrderId    *big.Int
	StageIndex *big.Int
	Raw        types.Log // Blockchain specific contextual infos
}

// FilterConfirmOrderStage is a free log retrieval operation binding the contract event 0x53182a8b7a3c3c3cab7d343a3256f3b24f11a2c9637aef95199bd47159cb783d.
//
// Solidity: event ConfirmOrderStage(uint256 indexed orderId, uint256 stageIndex)
func (_DeOrder *DeOrderFilterer) FilterConfirmOrderStage(opts *bind.FilterOpts, orderId []*big.Int) (*DeOrderConfirmOrderStageIterator, error) {

	var orderIdRule []interface{}
	for _, orderIdItem := range orderId {
		orderIdRule = append(orderIdRule, orderIdItem)
	}

	logs, sub, err := _DeOrder.contract.FilterLogs(opts, "ConfirmOrderStage", orderIdRule)
	if err != nil {
		return nil, err
	}
	return &DeOrderConfirmOrderStageIterator{contract: _DeOrder.contract, event: "ConfirmOrderStage", logs: logs, sub: sub}, nil
}

// WatchConfirmOrderStage is a free log subscription operation binding the contract event 0x53182a8b7a3c3c3cab7d343a3256f3b24f11a2c9637aef95199bd47159cb783d.
//
// Solidity: event ConfirmOrderStage(uint256 indexed orderId, uint256 stageIndex)
func (_DeOrder *DeOrderFilterer) WatchConfirmOrderStage(opts *bind.WatchOpts, sink chan<- *DeOrderConfirmOrderStage, orderId []*big.Int) (event.Subscription, error) {

	var orderIdRule []interface{}
	for _, orderIdItem := range orderId {
		orderIdRule = append(orderIdRule, orderIdItem)
	}

	logs, sub, err := _DeOrder.contract.WatchLogs(opts, "ConfirmOrderStage", orderIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(DeOrderConfirmOrderStage)
				if err := _DeOrder.contract.UnpackLog(event, "ConfirmOrderStage", log); err != nil {
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

// ParseConfirmOrderStage is a log parse operation binding the contract event 0x53182a8b7a3c3c3cab7d343a3256f3b24f11a2c9637aef95199bd47159cb783d.
//
// Solidity: event ConfirmOrderStage(uint256 indexed orderId, uint256 stageIndex)
func (_DeOrder *DeOrderFilterer) ParseConfirmOrderStage(log types.Log) (*DeOrderConfirmOrderStage, error) {
	event := new(DeOrderConfirmOrderStage)
	if err := _DeOrder.contract.UnpackLog(event, "ConfirmOrderStage", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// DeOrderFeeUpdatedIterator is returned from FilterFeeUpdated and is used to iterate over the raw logs and unpacked data for FeeUpdated events raised by the DeOrder contract.
type DeOrderFeeUpdatedIterator struct {
	Event *DeOrderFeeUpdated // Event containing the contract specifics and raw log

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
func (it *DeOrderFeeUpdatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(DeOrderFeeUpdated)
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
		it.Event = new(DeOrderFeeUpdated)
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
func (it *DeOrderFeeUpdatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *DeOrderFeeUpdatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// DeOrderFeeUpdated represents a FeeUpdated event raised by the DeOrder contract.
type DeOrderFeeUpdated struct {
	Fee   *big.Int
	FeeTo common.Address
	Raw   types.Log // Blockchain specific contextual infos
}

// FilterFeeUpdated is a free log retrieval operation binding the contract event 0x7cfad8b150be9751a5386cc4e0f549618032ff63d14fab4f77cd4b0aaaedc242.
//
// Solidity: event FeeUpdated(uint256 fee, address feeTo)
func (_DeOrder *DeOrderFilterer) FilterFeeUpdated(opts *bind.FilterOpts) (*DeOrderFeeUpdatedIterator, error) {

	logs, sub, err := _DeOrder.contract.FilterLogs(opts, "FeeUpdated")
	if err != nil {
		return nil, err
	}
	return &DeOrderFeeUpdatedIterator{contract: _DeOrder.contract, event: "FeeUpdated", logs: logs, sub: sub}, nil
}

// WatchFeeUpdated is a free log subscription operation binding the contract event 0x7cfad8b150be9751a5386cc4e0f549618032ff63d14fab4f77cd4b0aaaedc242.
//
// Solidity: event FeeUpdated(uint256 fee, address feeTo)
func (_DeOrder *DeOrderFilterer) WatchFeeUpdated(opts *bind.WatchOpts, sink chan<- *DeOrderFeeUpdated) (event.Subscription, error) {

	logs, sub, err := _DeOrder.contract.WatchLogs(opts, "FeeUpdated")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(DeOrderFeeUpdated)
				if err := _DeOrder.contract.UnpackLog(event, "FeeUpdated", log); err != nil {
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

// ParseFeeUpdated is a log parse operation binding the contract event 0x7cfad8b150be9751a5386cc4e0f549618032ff63d14fab4f77cd4b0aaaedc242.
//
// Solidity: event FeeUpdated(uint256 fee, address feeTo)
func (_DeOrder *DeOrderFilterer) ParseFeeUpdated(log types.Log) (*DeOrderFeeUpdated, error) {
	event := new(DeOrderFeeUpdated)
	if err := _DeOrder.contract.UnpackLog(event, "FeeUpdated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// DeOrderOrderAbortIterator is returned from FilterOrderAbort and is used to iterate over the raw logs and unpacked data for OrderAbort events raised by the DeOrder contract.
type DeOrderOrderAbortIterator struct {
	Event *DeOrderOrderAbort // Event containing the contract specifics and raw log

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
func (it *DeOrderOrderAbortIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(DeOrderOrderAbort)
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
		it.Event = new(DeOrderOrderAbort)
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
func (it *DeOrderOrderAbortIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *DeOrderOrderAbortIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// DeOrderOrderAbort represents a OrderAbort event raised by the DeOrder contract.
type DeOrderOrderAbort struct {
	OrderId    *big.Int
	Who        common.Address
	StageIndex *big.Int
	Raw        types.Log // Blockchain specific contextual infos
}

// FilterOrderAbort is a free log retrieval operation binding the contract event 0xad749093d9bc75cae179ecf479b40b860f762b1a9fff1b935a6ab014a2089fb2.
//
// Solidity: event OrderAbort(uint256 indexed orderId, address who, uint256 stageIndex)
func (_DeOrder *DeOrderFilterer) FilterOrderAbort(opts *bind.FilterOpts, orderId []*big.Int) (*DeOrderOrderAbortIterator, error) {

	var orderIdRule []interface{}
	for _, orderIdItem := range orderId {
		orderIdRule = append(orderIdRule, orderIdItem)
	}

	logs, sub, err := _DeOrder.contract.FilterLogs(opts, "OrderAbort", orderIdRule)
	if err != nil {
		return nil, err
	}
	return &DeOrderOrderAbortIterator{contract: _DeOrder.contract, event: "OrderAbort", logs: logs, sub: sub}, nil
}

// WatchOrderAbort is a free log subscription operation binding the contract event 0xad749093d9bc75cae179ecf479b40b860f762b1a9fff1b935a6ab014a2089fb2.
//
// Solidity: event OrderAbort(uint256 indexed orderId, address who, uint256 stageIndex)
func (_DeOrder *DeOrderFilterer) WatchOrderAbort(opts *bind.WatchOpts, sink chan<- *DeOrderOrderAbort, orderId []*big.Int) (event.Subscription, error) {

	var orderIdRule []interface{}
	for _, orderIdItem := range orderId {
		orderIdRule = append(orderIdRule, orderIdItem)
	}

	logs, sub, err := _DeOrder.contract.WatchLogs(opts, "OrderAbort", orderIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(DeOrderOrderAbort)
				if err := _DeOrder.contract.UnpackLog(event, "OrderAbort", log); err != nil {
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

// ParseOrderAbort is a log parse operation binding the contract event 0xad749093d9bc75cae179ecf479b40b860f762b1a9fff1b935a6ab014a2089fb2.
//
// Solidity: event OrderAbort(uint256 indexed orderId, address who, uint256 stageIndex)
func (_DeOrder *DeOrderFilterer) ParseOrderAbort(log types.Log) (*DeOrderOrderAbort, error) {
	event := new(DeOrderOrderAbort)
	if err := _DeOrder.contract.UnpackLog(event, "OrderAbort", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// DeOrderOrderCreatedIterator is returned from FilterOrderCreated and is used to iterate over the raw logs and unpacked data for OrderCreated events raised by the DeOrder contract.
type DeOrderOrderCreatedIterator struct {
	Event *DeOrderOrderCreated // Event containing the contract specifics and raw log

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
func (it *DeOrderOrderCreatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(DeOrderOrderCreated)
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
		it.Event = new(DeOrderOrderCreated)
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
func (it *DeOrderOrderCreatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *DeOrderOrderCreatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// DeOrderOrderCreated represents a OrderCreated event raised by the DeOrder contract.
type DeOrderOrderCreated struct {
	TaskId  *big.Int
	OrderId *big.Int
	Issuer  common.Address
	Worker  common.Address
	Token   common.Address
	Amount  *big.Int
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterOrderCreated is a free log retrieval operation binding the contract event 0x783ccb0007f88ccc789789cdb8b560e6519914e3b4dcba0c266d3bb52345a05c.
//
// Solidity: event OrderCreated(uint256 indexed taskId, uint256 indexed orderId, address issuer, address worker, address token, uint256 amount)
func (_DeOrder *DeOrderFilterer) FilterOrderCreated(opts *bind.FilterOpts, taskId []*big.Int, orderId []*big.Int) (*DeOrderOrderCreatedIterator, error) {

	var taskIdRule []interface{}
	for _, taskIdItem := range taskId {
		taskIdRule = append(taskIdRule, taskIdItem)
	}
	var orderIdRule []interface{}
	for _, orderIdItem := range orderId {
		orderIdRule = append(orderIdRule, orderIdItem)
	}

	logs, sub, err := _DeOrder.contract.FilterLogs(opts, "OrderCreated", taskIdRule, orderIdRule)
	if err != nil {
		return nil, err
	}
	return &DeOrderOrderCreatedIterator{contract: _DeOrder.contract, event: "OrderCreated", logs: logs, sub: sub}, nil
}

// WatchOrderCreated is a free log subscription operation binding the contract event 0x783ccb0007f88ccc789789cdb8b560e6519914e3b4dcba0c266d3bb52345a05c.
//
// Solidity: event OrderCreated(uint256 indexed taskId, uint256 indexed orderId, address issuer, address worker, address token, uint256 amount)
func (_DeOrder *DeOrderFilterer) WatchOrderCreated(opts *bind.WatchOpts, sink chan<- *DeOrderOrderCreated, taskId []*big.Int, orderId []*big.Int) (event.Subscription, error) {

	var taskIdRule []interface{}
	for _, taskIdItem := range taskId {
		taskIdRule = append(taskIdRule, taskIdItem)
	}
	var orderIdRule []interface{}
	for _, orderIdItem := range orderId {
		orderIdRule = append(orderIdRule, orderIdItem)
	}

	logs, sub, err := _DeOrder.contract.WatchLogs(opts, "OrderCreated", taskIdRule, orderIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(DeOrderOrderCreated)
				if err := _DeOrder.contract.UnpackLog(event, "OrderCreated", log); err != nil {
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

// ParseOrderCreated is a log parse operation binding the contract event 0x783ccb0007f88ccc789789cdb8b560e6519914e3b4dcba0c266d3bb52345a05c.
//
// Solidity: event OrderCreated(uint256 indexed taskId, uint256 indexed orderId, address issuer, address worker, address token, uint256 amount)
func (_DeOrder *DeOrderFilterer) ParseOrderCreated(log types.Log) (*DeOrderOrderCreated, error) {
	event := new(DeOrderOrderCreated)
	if err := _DeOrder.contract.UnpackLog(event, "OrderCreated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// DeOrderOrderModifiedIterator is returned from FilterOrderModified and is used to iterate over the raw logs and unpacked data for OrderModified events raised by the DeOrder contract.
type DeOrderOrderModifiedIterator struct {
	Event *DeOrderOrderModified // Event containing the contract specifics and raw log

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
func (it *DeOrderOrderModifiedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(DeOrderOrderModified)
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
		it.Event = new(DeOrderOrderModified)
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
func (it *DeOrderOrderModifiedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *DeOrderOrderModifiedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// DeOrderOrderModified represents a OrderModified event raised by the DeOrder contract.
type DeOrderOrderModified struct {
	OrderId *big.Int
	Token   common.Address
	Amount  *big.Int
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterOrderModified is a free log retrieval operation binding the contract event 0x4add6162d869cec70eb9deab7fffe7d235bdc6e05208868eb1fbd8c581c40385.
//
// Solidity: event OrderModified(uint256 indexed orderId, address token, uint256 amount)
func (_DeOrder *DeOrderFilterer) FilterOrderModified(opts *bind.FilterOpts, orderId []*big.Int) (*DeOrderOrderModifiedIterator, error) {

	var orderIdRule []interface{}
	for _, orderIdItem := range orderId {
		orderIdRule = append(orderIdRule, orderIdItem)
	}

	logs, sub, err := _DeOrder.contract.FilterLogs(opts, "OrderModified", orderIdRule)
	if err != nil {
		return nil, err
	}
	return &DeOrderOrderModifiedIterator{contract: _DeOrder.contract, event: "OrderModified", logs: logs, sub: sub}, nil
}

// WatchOrderModified is a free log subscription operation binding the contract event 0x4add6162d869cec70eb9deab7fffe7d235bdc6e05208868eb1fbd8c581c40385.
//
// Solidity: event OrderModified(uint256 indexed orderId, address token, uint256 amount)
func (_DeOrder *DeOrderFilterer) WatchOrderModified(opts *bind.WatchOpts, sink chan<- *DeOrderOrderModified, orderId []*big.Int) (event.Subscription, error) {

	var orderIdRule []interface{}
	for _, orderIdItem := range orderId {
		orderIdRule = append(orderIdRule, orderIdItem)
	}

	logs, sub, err := _DeOrder.contract.WatchLogs(opts, "OrderModified", orderIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(DeOrderOrderModified)
				if err := _DeOrder.contract.UnpackLog(event, "OrderModified", log); err != nil {
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

// ParseOrderModified is a log parse operation binding the contract event 0x4add6162d869cec70eb9deab7fffe7d235bdc6e05208868eb1fbd8c581c40385.
//
// Solidity: event OrderModified(uint256 indexed orderId, address token, uint256 amount)
func (_DeOrder *DeOrderFilterer) ParseOrderModified(log types.Log) (*DeOrderOrderModified, error) {
	event := new(DeOrderOrderModified)
	if err := _DeOrder.contract.UnpackLog(event, "OrderModified", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// DeOrderOrderStartedIterator is returned from FilterOrderStarted and is used to iterate over the raw logs and unpacked data for OrderStarted events raised by the DeOrder contract.
type DeOrderOrderStartedIterator struct {
	Event *DeOrderOrderStarted // Event containing the contract specifics and raw log

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
func (it *DeOrderOrderStartedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(DeOrderOrderStarted)
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
		it.Event = new(DeOrderOrderStarted)
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
func (it *DeOrderOrderStartedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *DeOrderOrderStartedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// DeOrderOrderStarted represents a OrderStarted event raised by the DeOrder contract.
type DeOrderOrderStarted struct {
	OrderId *big.Int
	Who     common.Address
	PayType *big.Int
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterOrderStarted is a free log retrieval operation binding the contract event 0xbbef050dc707f179bc1395f0bcb8a732c4f3ab402205b507f03e0ad8d3935236.
//
// Solidity: event OrderStarted(uint256 indexed orderId, address who, uint256 payType)
func (_DeOrder *DeOrderFilterer) FilterOrderStarted(opts *bind.FilterOpts, orderId []*big.Int) (*DeOrderOrderStartedIterator, error) {

	var orderIdRule []interface{}
	for _, orderIdItem := range orderId {
		orderIdRule = append(orderIdRule, orderIdItem)
	}

	logs, sub, err := _DeOrder.contract.FilterLogs(opts, "OrderStarted", orderIdRule)
	if err != nil {
		return nil, err
	}
	return &DeOrderOrderStartedIterator{contract: _DeOrder.contract, event: "OrderStarted", logs: logs, sub: sub}, nil
}

// WatchOrderStarted is a free log subscription operation binding the contract event 0xbbef050dc707f179bc1395f0bcb8a732c4f3ab402205b507f03e0ad8d3935236.
//
// Solidity: event OrderStarted(uint256 indexed orderId, address who, uint256 payType)
func (_DeOrder *DeOrderFilterer) WatchOrderStarted(opts *bind.WatchOpts, sink chan<- *DeOrderOrderStarted, orderId []*big.Int) (event.Subscription, error) {

	var orderIdRule []interface{}
	for _, orderIdItem := range orderId {
		orderIdRule = append(orderIdRule, orderIdItem)
	}

	logs, sub, err := _DeOrder.contract.WatchLogs(opts, "OrderStarted", orderIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(DeOrderOrderStarted)
				if err := _DeOrder.contract.UnpackLog(event, "OrderStarted", log); err != nil {
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

// ParseOrderStarted is a log parse operation binding the contract event 0xbbef050dc707f179bc1395f0bcb8a732c4f3ab402205b507f03e0ad8d3935236.
//
// Solidity: event OrderStarted(uint256 indexed orderId, address who, uint256 payType)
func (_DeOrder *DeOrderFilterer) ParseOrderStarted(log types.Log) (*DeOrderOrderStarted, error) {
	event := new(DeOrderOrderStarted)
	if err := _DeOrder.contract.UnpackLog(event, "OrderStarted", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// DeOrderOwnershipTransferredIterator is returned from FilterOwnershipTransferred and is used to iterate over the raw logs and unpacked data for OwnershipTransferred events raised by the DeOrder contract.
type DeOrderOwnershipTransferredIterator struct {
	Event *DeOrderOwnershipTransferred // Event containing the contract specifics and raw log

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
func (it *DeOrderOwnershipTransferredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(DeOrderOwnershipTransferred)
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
		it.Event = new(DeOrderOwnershipTransferred)
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
func (it *DeOrderOwnershipTransferredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *DeOrderOwnershipTransferredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// DeOrderOwnershipTransferred represents a OwnershipTransferred event raised by the DeOrder contract.
type DeOrderOwnershipTransferred struct {
	PreviousOwner common.Address
	NewOwner      common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterOwnershipTransferred is a free log retrieval operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_DeOrder *DeOrderFilterer) FilterOwnershipTransferred(opts *bind.FilterOpts, previousOwner []common.Address, newOwner []common.Address) (*DeOrderOwnershipTransferredIterator, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _DeOrder.contract.FilterLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return &DeOrderOwnershipTransferredIterator{contract: _DeOrder.contract, event: "OwnershipTransferred", logs: logs, sub: sub}, nil
}

// WatchOwnershipTransferred is a free log subscription operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_DeOrder *DeOrderFilterer) WatchOwnershipTransferred(opts *bind.WatchOpts, sink chan<- *DeOrderOwnershipTransferred, previousOwner []common.Address, newOwner []common.Address) (event.Subscription, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _DeOrder.contract.WatchLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(DeOrderOwnershipTransferred)
				if err := _DeOrder.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
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
func (_DeOrder *DeOrderFilterer) ParseOwnershipTransferred(log types.Log) (*DeOrderOwnershipTransferred, error) {
	event := new(DeOrderOwnershipTransferred)
	if err := _DeOrder.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// DeOrderProlongStageIterator is returned from FilterProlongStage and is used to iterate over the raw logs and unpacked data for ProlongStage events raised by the DeOrder contract.
type DeOrderProlongStageIterator struct {
	Event *DeOrderProlongStage // Event containing the contract specifics and raw log

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
func (it *DeOrderProlongStageIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(DeOrderProlongStage)
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
		it.Event = new(DeOrderProlongStage)
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
func (it *DeOrderProlongStageIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *DeOrderProlongStageIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// DeOrderProlongStage represents a ProlongStage event raised by the DeOrder contract.
type DeOrderProlongStage struct {
	OrderId      *big.Int
	StageIndex   *big.Int
	AppendPeriod *big.Int
	Raw          types.Log // Blockchain specific contextual infos
}

// FilterProlongStage is a free log retrieval operation binding the contract event 0x5a4b481a886d8dbf31ab860a2b1045da1c4df5585248f8012f0161b78fc21ff0.
//
// Solidity: event ProlongStage(uint256 indexed orderId, uint256 stageIndex, uint256 appendPeriod)
func (_DeOrder *DeOrderFilterer) FilterProlongStage(opts *bind.FilterOpts, orderId []*big.Int) (*DeOrderProlongStageIterator, error) {

	var orderIdRule []interface{}
	for _, orderIdItem := range orderId {
		orderIdRule = append(orderIdRule, orderIdItem)
	}

	logs, sub, err := _DeOrder.contract.FilterLogs(opts, "ProlongStage", orderIdRule)
	if err != nil {
		return nil, err
	}
	return &DeOrderProlongStageIterator{contract: _DeOrder.contract, event: "ProlongStage", logs: logs, sub: sub}, nil
}

// WatchProlongStage is a free log subscription operation binding the contract event 0x5a4b481a886d8dbf31ab860a2b1045da1c4df5585248f8012f0161b78fc21ff0.
//
// Solidity: event ProlongStage(uint256 indexed orderId, uint256 stageIndex, uint256 appendPeriod)
func (_DeOrder *DeOrderFilterer) WatchProlongStage(opts *bind.WatchOpts, sink chan<- *DeOrderProlongStage, orderId []*big.Int) (event.Subscription, error) {

	var orderIdRule []interface{}
	for _, orderIdItem := range orderId {
		orderIdRule = append(orderIdRule, orderIdItem)
	}

	logs, sub, err := _DeOrder.contract.WatchLogs(opts, "ProlongStage", orderIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(DeOrderProlongStage)
				if err := _DeOrder.contract.UnpackLog(event, "ProlongStage", log); err != nil {
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

// ParseProlongStage is a log parse operation binding the contract event 0x5a4b481a886d8dbf31ab860a2b1045da1c4df5585248f8012f0161b78fc21ff0.
//
// Solidity: event ProlongStage(uint256 indexed orderId, uint256 stageIndex, uint256 appendPeriod)
func (_DeOrder *DeOrderFilterer) ParseProlongStage(log types.Log) (*DeOrderProlongStage, error) {
	event := new(DeOrderProlongStage)
	if err := _DeOrder.contract.UnpackLog(event, "ProlongStage", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// DeOrderSetStageIterator is returned from FilterSetStage and is used to iterate over the raw logs and unpacked data for SetStage events raised by the DeOrder contract.
type DeOrderSetStageIterator struct {
	Event *DeOrderSetStage // Event containing the contract specifics and raw log

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
func (it *DeOrderSetStageIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(DeOrderSetStage)
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
		it.Event = new(DeOrderSetStage)
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
func (it *DeOrderSetStageIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *DeOrderSetStageIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// DeOrderSetStage represents a SetStage event raised by the DeOrder contract.
type DeOrderSetStage struct {
	OrderId *big.Int
	Amounts []*big.Int
	Periods []*big.Int
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterSetStage is a free log retrieval operation binding the contract event 0x2afaf2443808078691312232696466d71cd881f3b4529e39d9d72fd15c15eb12.
//
// Solidity: event SetStage(uint256 indexed orderId, uint256[] amounts, uint256[] periods)
func (_DeOrder *DeOrderFilterer) FilterSetStage(opts *bind.FilterOpts, orderId []*big.Int) (*DeOrderSetStageIterator, error) {

	var orderIdRule []interface{}
	for _, orderIdItem := range orderId {
		orderIdRule = append(orderIdRule, orderIdItem)
	}

	logs, sub, err := _DeOrder.contract.FilterLogs(opts, "SetStage", orderIdRule)
	if err != nil {
		return nil, err
	}
	return &DeOrderSetStageIterator{contract: _DeOrder.contract, event: "SetStage", logs: logs, sub: sub}, nil
}

// WatchSetStage is a free log subscription operation binding the contract event 0x2afaf2443808078691312232696466d71cd881f3b4529e39d9d72fd15c15eb12.
//
// Solidity: event SetStage(uint256 indexed orderId, uint256[] amounts, uint256[] periods)
func (_DeOrder *DeOrderFilterer) WatchSetStage(opts *bind.WatchOpts, sink chan<- *DeOrderSetStage, orderId []*big.Int) (event.Subscription, error) {

	var orderIdRule []interface{}
	for _, orderIdItem := range orderId {
		orderIdRule = append(orderIdRule, orderIdItem)
	}

	logs, sub, err := _DeOrder.contract.WatchLogs(opts, "SetStage", orderIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(DeOrderSetStage)
				if err := _DeOrder.contract.UnpackLog(event, "SetStage", log); err != nil {
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

// ParseSetStage is a log parse operation binding the contract event 0x2afaf2443808078691312232696466d71cd881f3b4529e39d9d72fd15c15eb12.
//
// Solidity: event SetStage(uint256 indexed orderId, uint256[] amounts, uint256[] periods)
func (_DeOrder *DeOrderFilterer) ParseSetStage(log types.Log) (*DeOrderSetStage, error) {
	event := new(DeOrderSetStage)
	if err := _DeOrder.contract.UnpackLog(event, "SetStage", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// DeOrderSupportTokenIterator is returned from FilterSupportToken and is used to iterate over the raw logs and unpacked data for SupportToken events raised by the DeOrder contract.
type DeOrderSupportTokenIterator struct {
	Event *DeOrderSupportToken // Event containing the contract specifics and raw log

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
func (it *DeOrderSupportTokenIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(DeOrderSupportToken)
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
		it.Event = new(DeOrderSupportToken)
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
func (it *DeOrderSupportTokenIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *DeOrderSupportTokenIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// DeOrderSupportToken represents a SupportToken event raised by the DeOrder contract.
type DeOrderSupportToken struct {
	Token   common.Address
	Enabled bool
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterSupportToken is a free log retrieval operation binding the contract event 0x5e2b4e0b770b15ed507e19f9f3b164f07087641b80ca21ef0052f6e8611daaa7.
//
// Solidity: event SupportToken(address token, bool enabled)
func (_DeOrder *DeOrderFilterer) FilterSupportToken(opts *bind.FilterOpts) (*DeOrderSupportTokenIterator, error) {

	logs, sub, err := _DeOrder.contract.FilterLogs(opts, "SupportToken")
	if err != nil {
		return nil, err
	}
	return &DeOrderSupportTokenIterator{contract: _DeOrder.contract, event: "SupportToken", logs: logs, sub: sub}, nil
}

// WatchSupportToken is a free log subscription operation binding the contract event 0x5e2b4e0b770b15ed507e19f9f3b164f07087641b80ca21ef0052f6e8611daaa7.
//
// Solidity: event SupportToken(address token, bool enabled)
func (_DeOrder *DeOrderFilterer) WatchSupportToken(opts *bind.WatchOpts, sink chan<- *DeOrderSupportToken) (event.Subscription, error) {

	logs, sub, err := _DeOrder.contract.WatchLogs(opts, "SupportToken")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(DeOrderSupportToken)
				if err := _DeOrder.contract.UnpackLog(event, "SupportToken", log); err != nil {
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

// ParseSupportToken is a log parse operation binding the contract event 0x5e2b4e0b770b15ed507e19f9f3b164f07087641b80ca21ef0052f6e8611daaa7.
//
// Solidity: event SupportToken(address token, bool enabled)
func (_DeOrder *DeOrderFilterer) ParseSupportToken(log types.Log) (*DeOrderSupportToken, error) {
	event := new(DeOrderSupportToken)
	if err := _DeOrder.contract.UnpackLog(event, "SupportToken", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// DeOrderWithdrawIterator is returned from FilterWithdraw and is used to iterate over the raw logs and unpacked data for Withdraw events raised by the DeOrder contract.
type DeOrderWithdrawIterator struct {
	Event *DeOrderWithdraw // Event containing the contract specifics and raw log

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
func (it *DeOrderWithdrawIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(DeOrderWithdraw)
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
		it.Event = new(DeOrderWithdraw)
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
func (it *DeOrderWithdrawIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *DeOrderWithdrawIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// DeOrderWithdraw represents a Withdraw event raised by the DeOrder contract.
type DeOrderWithdraw struct {
	OrderId    *big.Int
	Amount     *big.Int
	StageIndex *big.Int
	Raw        types.Log // Blockchain specific contextual infos
}

// FilterWithdraw is a free log retrieval operation binding the contract event 0xa01a72713bf837059e3a668d28f0de277fb7f24f2a4e95bf926703c95b5f12b2.
//
// Solidity: event Withdraw(uint256 indexed orderId, uint256 amount, uint256 stageIndex)
func (_DeOrder *DeOrderFilterer) FilterWithdraw(opts *bind.FilterOpts, orderId []*big.Int) (*DeOrderWithdrawIterator, error) {

	var orderIdRule []interface{}
	for _, orderIdItem := range orderId {
		orderIdRule = append(orderIdRule, orderIdItem)
	}

	logs, sub, err := _DeOrder.contract.FilterLogs(opts, "Withdraw", orderIdRule)
	if err != nil {
		return nil, err
	}
	return &DeOrderWithdrawIterator{contract: _DeOrder.contract, event: "Withdraw", logs: logs, sub: sub}, nil
}

// WatchWithdraw is a free log subscription operation binding the contract event 0xa01a72713bf837059e3a668d28f0de277fb7f24f2a4e95bf926703c95b5f12b2.
//
// Solidity: event Withdraw(uint256 indexed orderId, uint256 amount, uint256 stageIndex)
func (_DeOrder *DeOrderFilterer) WatchWithdraw(opts *bind.WatchOpts, sink chan<- *DeOrderWithdraw, orderId []*big.Int) (event.Subscription, error) {

	var orderIdRule []interface{}
	for _, orderIdItem := range orderId {
		orderIdRule = append(orderIdRule, orderIdItem)
	}

	logs, sub, err := _DeOrder.contract.WatchLogs(opts, "Withdraw", orderIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(DeOrderWithdraw)
				if err := _DeOrder.contract.UnpackLog(event, "Withdraw", log); err != nil {
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

// ParseWithdraw is a log parse operation binding the contract event 0xa01a72713bf837059e3a668d28f0de277fb7f24f2a4e95bf926703c95b5f12b2.
//
// Solidity: event Withdraw(uint256 indexed orderId, uint256 amount, uint256 stageIndex)
func (_DeOrder *DeOrderFilterer) ParseWithdraw(log types.Log) (*DeOrderWithdraw, error) {
	event := new(DeOrderWithdraw)
	if err := _DeOrder.contract.UnpackLog(event, "Withdraw", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
