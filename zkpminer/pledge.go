// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package zkpminer

import (
	"math/big"
	"strings"

	ethereum "github.com/Evanesco-Labs/go-evanesco"
	"github.com/Evanesco-Labs/go-evanesco/accounts/abi"
	"github.com/Evanesco-Labs/go-evanesco/accounts/abi/bind"
	"github.com/Evanesco-Labs/go-evanesco/common"
	"github.com/Evanesco-Labs/go-evanesco/core/types"
	"github.com/Evanesco-Labs/go-evanesco/event"
)

// Reference imports to suppress errors if they are not otherwise used.
var (
	_ = big.NewInt
	_ = strings.NewReader
	_ = ethereum.NotFound
	_ = bind.Bind
	_ = common.Big1
	_ = types.BloomLookup
	_ = event.NewSubscription
)

// PledgeUserInfo is an auto generated low-level Go binding around an user-defined struct.
type PledgeUserInfo struct {
	Index     *big.Int
	User      common.Address
	DepositAt uint32
	ExpireAt  uint32
	Nodes     uint32
	Amount    *big.Int
	Settled   bool
}

// PledgeWithdrawInfo is an auto generated low-level Go binding around an user-defined struct.
type PledgeWithdrawInfo struct {
	Index      *big.Int
	WithdrawAt *big.Int
	UIndex     *big.Int
	WAmount    *big.Int
}

// PledgeABI is the input ABI used to generate the binding from.
const PledgeABI = "[{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"index\",\"type\":\"uint256\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"user\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"Deposit\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"previousOwner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"OwnershipTransferred\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"index\",\"type\":\"uint256\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"user\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"Withdraw\",\"type\":\"event\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"borrowsRecard\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"index\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"user\",\"type\":\"address\"},{\"internalType\":\"uint32\",\"name\":\"depositAt\",\"type\":\"uint32\"},{\"internalType\":\"uint32\",\"name\":\"expireAt\",\"type\":\"uint32\"},{\"internalType\":\"uint32\",\"name\":\"nodes\",\"type\":\"uint32\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"internalType\":\"bool\",\"name\":\"settled\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\",\"constant\":true},{\"inputs\":[],\"name\":\"minPledgeAmount\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\",\"constant\":true},{\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\",\"constant\":true},{\"inputs\":[],\"name\":\"renounceOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"totalBorrows\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\",\"constant\":true},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"transferOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"userDepositsLending\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\",\"constant\":true},{\"stateMutability\":\"payable\",\"type\":\"receive\",\"payable\":true},{\"inputs\":[],\"name\":\"initialize\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_nodes\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_amount\",\"type\":\"uint256\"},{\"internalType\":\"address[]\",\"name\":\"_addrs\",\"type\":\"address[]\"},{\"internalType\":\"uint256\",\"name\":\"_expiredAt\",\"type\":\"uint256\"}],\"name\":\"deposit\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\",\"payable\":true},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_index\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"_receiver\",\"type\":\"address\"}],\"name\":\"withdraw\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_index\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_curTime\",\"type\":\"uint256\"}],\"name\":\"allowWithdraw\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\",\"constant\":true},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"user\",\"type\":\"address\"}],\"name\":\"iseffective\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\",\"constant\":true},{\"inputs\":[{\"internalType\":\"contractIQLF\",\"name\":\"_qlf\",\"type\":\"address\"},{\"internalType\":\"bool\",\"name\":\"_isWhiteList\",\"type\":\"bool\"}],\"name\":\"setAddressQLF\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_time\",\"type\":\"uint256\"}],\"name\":\"setMinPledgeAmount\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_lockTime\",\"type\":\"uint256\"}],\"name\":\"setLockTime\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bool\",\"name\":\"_isDepositPaused\",\"type\":\"bool\"},{\"internalType\":\"bool\",\"name\":\"_isWithdrawPaused\",\"type\":\"bool\"}],\"name\":\"setPaused\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"user\",\"type\":\"address\"}],\"name\":\"isValidateWhite\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\",\"constant\":true},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"user\",\"type\":\"address\"}],\"name\":\"getUserDeposits\",\"outputs\":[{\"components\":[{\"internalType\":\"uint256\",\"name\":\"index\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"user\",\"type\":\"address\"},{\"internalType\":\"uint32\",\"name\":\"depositAt\",\"type\":\"uint32\"},{\"internalType\":\"uint32\",\"name\":\"expireAt\",\"type\":\"uint32\"},{\"internalType\":\"uint32\",\"name\":\"nodes\",\"type\":\"uint32\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"internalType\":\"bool\",\"name\":\"settled\",\"type\":\"bool\"}],\"internalType\":\"structPledge.UserInfo[]\",\"name\":\"\",\"type\":\"tuple[]\"}],\"stateMutability\":\"view\",\"type\":\"function\",\"constant\":true},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"user\",\"type\":\"address\"}],\"name\":\"getUserWithdraws\",\"outputs\":[{\"components\":[{\"internalType\":\"uint256\",\"name\":\"index\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"withdrawAt\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"uIndex\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"wAmount\",\"type\":\"uint256\"}],\"internalType\":\"structPledge.WithdrawInfo[]\",\"name\":\"\",\"type\":\"tuple[]\"}],\"stateMutability\":\"view\",\"type\":\"function\",\"constant\":true},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"user\",\"type\":\"address\"}],\"name\":\"getUserDepositsMininAddr\",\"outputs\":[{\"internalType\":\"address[]\",\"name\":\"\",\"type\":\"address[]\"}],\"stateMutability\":\"view\",\"type\":\"function\",\"constant\":true},{\"inputs\":[],\"name\":\"getAddressQLF\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\",\"constant\":true},{\"inputs\":[],\"name\":\"getisWhiteList\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\",\"constant\":true},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_projectStartTime\",\"type\":\"uint256\"}],\"name\":\"setProjectStartTime\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getBlockTime\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"time\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\",\"constant\":true},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_address\",\"type\":\"address\"},{\"internalType\":\"bool\",\"name\":\"_flag\",\"type\":\"bool\"}],\"name\":\"addMininAddr\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"user\",\"type\":\"address\"}],\"name\":\"isExtisAddress\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\",\"constant\":true},{\"inputs\":[{\"internalType\":\"address[]\",\"name\":\"users\",\"type\":\"address[]\"}],\"name\":\"isExtisBathAddress\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\",\"constant\":true},{\"inputs\":[],\"name\":\"getUserDepositsMininAddrLength\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\",\"constant\":true},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_user\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_oldAddress\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_index\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"_replace\",\"type\":\"address\"}],\"name\":\"updateDepositsMininAddress\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_replace\",\"type\":\"address\"}],\"name\":\"updateDepositsCoinbaseAddr\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_sender\",\"type\":\"address\"}],\"name\":\"getDepositsCoinbaseAddr\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\",\"constant\":true},{\"inputs\":[],\"name\":\"getBorrowsRecard\",\"outputs\":[{\"components\":[{\"internalType\":\"uint256\",\"name\":\"index\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"user\",\"type\":\"address\"},{\"internalType\":\"uint32\",\"name\":\"depositAt\",\"type\":\"uint32\"},{\"internalType\":\"uint32\",\"name\":\"expireAt\",\"type\":\"uint32\"},{\"internalType\":\"uint32\",\"name\":\"nodes\",\"type\":\"uint32\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"internalType\":\"bool\",\"name\":\"settled\",\"type\":\"bool\"}],\"internalType\":\"structPledge.UserInfo[]\",\"name\":\"\",\"type\":\"tuple[]\"}],\"stateMutability\":\"view\",\"type\":\"function\",\"constant\":true},{\"inputs\":[],\"name\":\"getUserDepositsLending\",\"outputs\":[{\"internalType\":\"uint256[]\",\"name\":\"\",\"type\":\"uint256[]\"}],\"stateMutability\":\"view\",\"type\":\"function\",\"constant\":true},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_sender\",\"type\":\"address\"}],\"name\":\"getUserLendingDeposits\",\"outputs\":[{\"components\":[{\"internalType\":\"uint256\",\"name\":\"index\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"user\",\"type\":\"address\"},{\"internalType\":\"uint32\",\"name\":\"depositAt\",\"type\":\"uint32\"},{\"internalType\":\"uint32\",\"name\":\"expireAt\",\"type\":\"uint32\"},{\"internalType\":\"uint32\",\"name\":\"nodes\",\"type\":\"uint32\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"internalType\":\"bool\",\"name\":\"settled\",\"type\":\"bool\"}],\"internalType\":\"structPledge.UserInfo[]\",\"name\":\"\",\"type\":\"tuple[]\"}],\"stateMutability\":\"view\",\"type\":\"function\",\"constant\":true},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_sender\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_nodes\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_amount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_expiredAt\",\"type\":\"uint256\"},{\"internalType\":\"address[]\",\"name\":\"_addrs\",\"type\":\"address[]\"}],\"name\":\"depositLending\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\",\"payable\":true},{\"inputs\":[],\"name\":\"getLendingAccount\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"_totalBorrow\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_notExpireBorrow\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_expireBorrow\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\",\"constant\":true},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_address\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_expireAt\",\"type\":\"uint256\"}],\"name\":\"setdataMigration\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"user\",\"type\":\"address\"}],\"name\":\"iseffectiveNew\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"},{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\",\"constant\":true}]"

// Pledge is an auto generated Go binding around an Ethereum contract.
type Pledge struct {
	PledgeCaller     // Read-only binding to the contract
	PledgeTransactor // Write-only binding to the contract
	PledgeFilterer   // Log filterer for contract events
}

// PledgeCaller is an auto generated read-only Go binding around an Ethereum contract.
type PledgeCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// PledgeTransactor is an auto generated write-only Go binding around an Ethereum contract.
type PledgeTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// PledgeFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type PledgeFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// PledgeSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type PledgeSession struct {
	Contract     *Pledge           // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// PledgeCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type PledgeCallerSession struct {
	Contract *PledgeCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts // Call options to use throughout this session
}

// PledgeTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type PledgeTransactorSession struct {
	Contract     *PledgeTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// PledgeRaw is an auto generated low-level Go binding around an Ethereum contract.
type PledgeRaw struct {
	Contract *Pledge // Generic contract binding to access the raw methods on
}

// PledgeCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type PledgeCallerRaw struct {
	Contract *PledgeCaller // Generic read-only contract binding to access the raw methods on
}

// PledgeTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type PledgeTransactorRaw struct {
	Contract *PledgeTransactor // Generic write-only contract binding to access the raw methods on
}

// NewPledge creates a new instance of Pledge, bound to a specific deployed contract.
func NewPledge(address common.Address, backend bind.ContractBackend) (*Pledge, error) {
	contract, err := bindPledge(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Pledge{PledgeCaller: PledgeCaller{contract: contract}, PledgeTransactor: PledgeTransactor{contract: contract}, PledgeFilterer: PledgeFilterer{contract: contract}}, nil
}

// NewPledgeCaller creates a new read-only instance of Pledge, bound to a specific deployed contract.
func NewPledgeCaller(address common.Address, caller bind.ContractCaller) (*PledgeCaller, error) {
	contract, err := bindPledge(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &PledgeCaller{contract: contract}, nil
}

// NewPledgeTransactor creates a new write-only instance of Pledge, bound to a specific deployed contract.
func NewPledgeTransactor(address common.Address, transactor bind.ContractTransactor) (*PledgeTransactor, error) {
	contract, err := bindPledge(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &PledgeTransactor{contract: contract}, nil
}

// NewPledgeFilterer creates a new log filterer instance of Pledge, bound to a specific deployed contract.
func NewPledgeFilterer(address common.Address, filterer bind.ContractFilterer) (*PledgeFilterer, error) {
	contract, err := bindPledge(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &PledgeFilterer{contract: contract}, nil
}

// bindPledge binds a generic wrapper to an already deployed contract.
func bindPledge(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(PledgeABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Pledge *PledgeRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Pledge.Contract.PledgeCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Pledge *PledgeRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Pledge.Contract.PledgeTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Pledge *PledgeRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Pledge.Contract.PledgeTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Pledge *PledgeCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Pledge.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Pledge *PledgeTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Pledge.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Pledge *PledgeTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Pledge.Contract.contract.Transact(opts, method, params...)
}

// AllowWithdraw is a free data retrieval call binding the contract method 0xa3bd4585.
//
// Solidity: function allowWithdraw(uint256 _index, uint256 _curTime) view returns(bool)
func (_Pledge *PledgeCaller) AllowWithdraw(opts *bind.CallOpts, _index *big.Int, _curTime *big.Int) (bool, error) {
	var out []interface{}
	err := _Pledge.contract.Call(opts, &out, "allowWithdraw", _index, _curTime)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// AllowWithdraw is a free data retrieval call binding the contract method 0xa3bd4585.
//
// Solidity: function allowWithdraw(uint256 _index, uint256 _curTime) view returns(bool)
func (_Pledge *PledgeSession) AllowWithdraw(_index *big.Int, _curTime *big.Int) (bool, error) {
	return _Pledge.Contract.AllowWithdraw(&_Pledge.CallOpts, _index, _curTime)
}

// AllowWithdraw is a free data retrieval call binding the contract method 0xa3bd4585.
//
// Solidity: function allowWithdraw(uint256 _index, uint256 _curTime) view returns(bool)
func (_Pledge *PledgeCallerSession) AllowWithdraw(_index *big.Int, _curTime *big.Int) (bool, error) {
	return _Pledge.Contract.AllowWithdraw(&_Pledge.CallOpts, _index, _curTime)
}

// BorrowsRecard is a free data retrieval call binding the contract method 0xe9d1fd0d.
//
// Solidity: function borrowsRecard(uint256 ) view returns(uint256 index, address user, uint32 depositAt, uint32 expireAt, uint32 nodes, uint256 amount, bool settled)
func (_Pledge *PledgeCaller) BorrowsRecard(opts *bind.CallOpts, arg0 *big.Int) (struct {
	Index     *big.Int
	User      common.Address
	DepositAt uint32
	ExpireAt  uint32
	Nodes     uint32
	Amount    *big.Int
	Settled   bool
}, error) {
	var out []interface{}
	err := _Pledge.contract.Call(opts, &out, "borrowsRecard", arg0)

	outstruct := new(struct {
		Index     *big.Int
		User      common.Address
		DepositAt uint32
		ExpireAt  uint32
		Nodes     uint32
		Amount    *big.Int
		Settled   bool
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.Index = *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)
	outstruct.User = *abi.ConvertType(out[1], new(common.Address)).(*common.Address)
	outstruct.DepositAt = *abi.ConvertType(out[2], new(uint32)).(*uint32)
	outstruct.ExpireAt = *abi.ConvertType(out[3], new(uint32)).(*uint32)
	outstruct.Nodes = *abi.ConvertType(out[4], new(uint32)).(*uint32)
	outstruct.Amount = *abi.ConvertType(out[5], new(*big.Int)).(**big.Int)
	outstruct.Settled = *abi.ConvertType(out[6], new(bool)).(*bool)

	return *outstruct, err

}

// BorrowsRecard is a free data retrieval call binding the contract method 0xe9d1fd0d.
//
// Solidity: function borrowsRecard(uint256 ) view returns(uint256 index, address user, uint32 depositAt, uint32 expireAt, uint32 nodes, uint256 amount, bool settled)
func (_Pledge *PledgeSession) BorrowsRecard(arg0 *big.Int) (struct {
	Index     *big.Int
	User      common.Address
	DepositAt uint32
	ExpireAt  uint32
	Nodes     uint32
	Amount    *big.Int
	Settled   bool
}, error) {
	return _Pledge.Contract.BorrowsRecard(&_Pledge.CallOpts, arg0)
}

// BorrowsRecard is a free data retrieval call binding the contract method 0xe9d1fd0d.
//
// Solidity: function borrowsRecard(uint256 ) view returns(uint256 index, address user, uint32 depositAt, uint32 expireAt, uint32 nodes, uint256 amount, bool settled)
func (_Pledge *PledgeCallerSession) BorrowsRecard(arg0 *big.Int) (struct {
	Index     *big.Int
	User      common.Address
	DepositAt uint32
	ExpireAt  uint32
	Nodes     uint32
	Amount    *big.Int
	Settled   bool
}, error) {
	return _Pledge.Contract.BorrowsRecard(&_Pledge.CallOpts, arg0)
}

// GetAddressQLF is a free data retrieval call binding the contract method 0xa2de3904.
//
// Solidity: function getAddressQLF() view returns(address)
func (_Pledge *PledgeCaller) GetAddressQLF(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Pledge.contract.Call(opts, &out, "getAddressQLF")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// GetAddressQLF is a free data retrieval call binding the contract method 0xa2de3904.
//
// Solidity: function getAddressQLF() view returns(address)
func (_Pledge *PledgeSession) GetAddressQLF() (common.Address, error) {
	return _Pledge.Contract.GetAddressQLF(&_Pledge.CallOpts)
}

// GetAddressQLF is a free data retrieval call binding the contract method 0xa2de3904.
//
// Solidity: function getAddressQLF() view returns(address)
func (_Pledge *PledgeCallerSession) GetAddressQLF() (common.Address, error) {
	return _Pledge.Contract.GetAddressQLF(&_Pledge.CallOpts)
}

// GetBlockTime is a free data retrieval call binding the contract method 0x87ceff09.
//
// Solidity: function getBlockTime() view returns(uint256 time)
func (_Pledge *PledgeCaller) GetBlockTime(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Pledge.contract.Call(opts, &out, "getBlockTime")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetBlockTime is a free data retrieval call binding the contract method 0x87ceff09.
//
// Solidity: function getBlockTime() view returns(uint256 time)
func (_Pledge *PledgeSession) GetBlockTime() (*big.Int, error) {
	return _Pledge.Contract.GetBlockTime(&_Pledge.CallOpts)
}

// GetBlockTime is a free data retrieval call binding the contract method 0x87ceff09.
//
// Solidity: function getBlockTime() view returns(uint256 time)
func (_Pledge *PledgeCallerSession) GetBlockTime() (*big.Int, error) {
	return _Pledge.Contract.GetBlockTime(&_Pledge.CallOpts)
}

// GetBorrowsRecard is a free data retrieval call binding the contract method 0xecf57d4e.
//
// Solidity: function getBorrowsRecard() view returns((uint256,address,uint32,uint32,uint32,uint256,bool)[])
func (_Pledge *PledgeCaller) GetBorrowsRecard(opts *bind.CallOpts) ([]PledgeUserInfo, error) {
	var out []interface{}
	err := _Pledge.contract.Call(opts, &out, "getBorrowsRecard")

	if err != nil {
		return *new([]PledgeUserInfo), err
	}

	out0 := *abi.ConvertType(out[0], new([]PledgeUserInfo)).(*[]PledgeUserInfo)

	return out0, err

}

// GetBorrowsRecard is a free data retrieval call binding the contract method 0xecf57d4e.
//
// Solidity: function getBorrowsRecard() view returns((uint256,address,uint32,uint32,uint32,uint256,bool)[])
func (_Pledge *PledgeSession) GetBorrowsRecard() ([]PledgeUserInfo, error) {
	return _Pledge.Contract.GetBorrowsRecard(&_Pledge.CallOpts)
}

// GetBorrowsRecard is a free data retrieval call binding the contract method 0xecf57d4e.
//
// Solidity: function getBorrowsRecard() view returns((uint256,address,uint32,uint32,uint32,uint256,bool)[])
func (_Pledge *PledgeCallerSession) GetBorrowsRecard() ([]PledgeUserInfo, error) {
	return _Pledge.Contract.GetBorrowsRecard(&_Pledge.CallOpts)
}

// GetDepositsCoinbaseAddr is a free data retrieval call binding the contract method 0xaaaca3cc.
//
// Solidity: function getDepositsCoinbaseAddr(address _sender) view returns(address)
func (_Pledge *PledgeCaller) GetDepositsCoinbaseAddr(opts *bind.CallOpts, _sender common.Address) (common.Address, error) {
	var out []interface{}
	err := _Pledge.contract.Call(opts, &out, "getDepositsCoinbaseAddr", _sender)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// GetDepositsCoinbaseAddr is a free data retrieval call binding the contract method 0xaaaca3cc.
//
// Solidity: function getDepositsCoinbaseAddr(address _sender) view returns(address)
func (_Pledge *PledgeSession) GetDepositsCoinbaseAddr(_sender common.Address) (common.Address, error) {
	return _Pledge.Contract.GetDepositsCoinbaseAddr(&_Pledge.CallOpts, _sender)
}

// GetDepositsCoinbaseAddr is a free data retrieval call binding the contract method 0xaaaca3cc.
//
// Solidity: function getDepositsCoinbaseAddr(address _sender) view returns(address)
func (_Pledge *PledgeCallerSession) GetDepositsCoinbaseAddr(_sender common.Address) (common.Address, error) {
	return _Pledge.Contract.GetDepositsCoinbaseAddr(&_Pledge.CallOpts, _sender)
}

// GetLendingAccount is a free data retrieval call binding the contract method 0xc86f1b71.
//
// Solidity: function getLendingAccount() view returns(uint256 _totalBorrow, uint256 _notExpireBorrow, uint256 _expireBorrow)
func (_Pledge *PledgeCaller) GetLendingAccount(opts *bind.CallOpts) (struct {
	TotalBorrow     *big.Int
	NotExpireBorrow *big.Int
	ExpireBorrow    *big.Int
}, error) {
	var out []interface{}
	err := _Pledge.contract.Call(opts, &out, "getLendingAccount")

	outstruct := new(struct {
		TotalBorrow     *big.Int
		NotExpireBorrow *big.Int
		ExpireBorrow    *big.Int
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.TotalBorrow = *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)
	outstruct.NotExpireBorrow = *abi.ConvertType(out[1], new(*big.Int)).(**big.Int)
	outstruct.ExpireBorrow = *abi.ConvertType(out[2], new(*big.Int)).(**big.Int)

	return *outstruct, err

}

// GetLendingAccount is a free data retrieval call binding the contract method 0xc86f1b71.
//
// Solidity: function getLendingAccount() view returns(uint256 _totalBorrow, uint256 _notExpireBorrow, uint256 _expireBorrow)
func (_Pledge *PledgeSession) GetLendingAccount() (struct {
	TotalBorrow     *big.Int
	NotExpireBorrow *big.Int
	ExpireBorrow    *big.Int
}, error) {
	return _Pledge.Contract.GetLendingAccount(&_Pledge.CallOpts)
}

// GetLendingAccount is a free data retrieval call binding the contract method 0xc86f1b71.
//
// Solidity: function getLendingAccount() view returns(uint256 _totalBorrow, uint256 _notExpireBorrow, uint256 _expireBorrow)
func (_Pledge *PledgeCallerSession) GetLendingAccount() (struct {
	TotalBorrow     *big.Int
	NotExpireBorrow *big.Int
	ExpireBorrow    *big.Int
}, error) {
	return _Pledge.Contract.GetLendingAccount(&_Pledge.CallOpts)
}

// GetUserDeposits is a free data retrieval call binding the contract method 0x2a5bf6d2.
//
// Solidity: function getUserDeposits(address user) view returns((uint256,address,uint32,uint32,uint32,uint256,bool)[])
func (_Pledge *PledgeCaller) GetUserDeposits(opts *bind.CallOpts, user common.Address) ([]PledgeUserInfo, error) {
	var out []interface{}
	err := _Pledge.contract.Call(opts, &out, "getUserDeposits", user)

	if err != nil {
		return *new([]PledgeUserInfo), err
	}

	out0 := *abi.ConvertType(out[0], new([]PledgeUserInfo)).(*[]PledgeUserInfo)

	return out0, err

}

// GetUserDeposits is a free data retrieval call binding the contract method 0x2a5bf6d2.
//
// Solidity: function getUserDeposits(address user) view returns((uint256,address,uint32,uint32,uint32,uint256,bool)[])
func (_Pledge *PledgeSession) GetUserDeposits(user common.Address) ([]PledgeUserInfo, error) {
	return _Pledge.Contract.GetUserDeposits(&_Pledge.CallOpts, user)
}

// GetUserDeposits is a free data retrieval call binding the contract method 0x2a5bf6d2.
//
// Solidity: function getUserDeposits(address user) view returns((uint256,address,uint32,uint32,uint32,uint256,bool)[])
func (_Pledge *PledgeCallerSession) GetUserDeposits(user common.Address) ([]PledgeUserInfo, error) {
	return _Pledge.Contract.GetUserDeposits(&_Pledge.CallOpts, user)
}

// GetUserDepositsLending is a free data retrieval call binding the contract method 0xbebdca09.
//
// Solidity: function getUserDepositsLending() view returns(uint256[])
func (_Pledge *PledgeCaller) GetUserDepositsLending(opts *bind.CallOpts) ([]*big.Int, error) {
	var out []interface{}
	err := _Pledge.contract.Call(opts, &out, "getUserDepositsLending")

	if err != nil {
		return *new([]*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new([]*big.Int)).(*[]*big.Int)

	return out0, err

}

// GetUserDepositsLending is a free data retrieval call binding the contract method 0xbebdca09.
//
// Solidity: function getUserDepositsLending() view returns(uint256[])
func (_Pledge *PledgeSession) GetUserDepositsLending() ([]*big.Int, error) {
	return _Pledge.Contract.GetUserDepositsLending(&_Pledge.CallOpts)
}

// GetUserDepositsLending is a free data retrieval call binding the contract method 0xbebdca09.
//
// Solidity: function getUserDepositsLending() view returns(uint256[])
func (_Pledge *PledgeCallerSession) GetUserDepositsLending() ([]*big.Int, error) {
	return _Pledge.Contract.GetUserDepositsLending(&_Pledge.CallOpts)
}

// GetUserDepositsMininAddr is a free data retrieval call binding the contract method 0xee4f542c.
//
// Solidity: function getUserDepositsMininAddr(address user) view returns(address[])
func (_Pledge *PledgeCaller) GetUserDepositsMininAddr(opts *bind.CallOpts, user common.Address) ([]common.Address, error) {
	var out []interface{}
	err := _Pledge.contract.Call(opts, &out, "getUserDepositsMininAddr", user)

	if err != nil {
		return *new([]common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new([]common.Address)).(*[]common.Address)

	return out0, err

}

// GetUserDepositsMininAddr is a free data retrieval call binding the contract method 0xee4f542c.
//
// Solidity: function getUserDepositsMininAddr(address user) view returns(address[])
func (_Pledge *PledgeSession) GetUserDepositsMininAddr(user common.Address) ([]common.Address, error) {
	return _Pledge.Contract.GetUserDepositsMininAddr(&_Pledge.CallOpts, user)
}

// GetUserDepositsMininAddr is a free data retrieval call binding the contract method 0xee4f542c.
//
// Solidity: function getUserDepositsMininAddr(address user) view returns(address[])
func (_Pledge *PledgeCallerSession) GetUserDepositsMininAddr(user common.Address) ([]common.Address, error) {
	return _Pledge.Contract.GetUserDepositsMininAddr(&_Pledge.CallOpts, user)
}

// GetUserDepositsMininAddrLength is a free data retrieval call binding the contract method 0x9cf205bd.
//
// Solidity: function getUserDepositsMininAddrLength() view returns(uint256)
func (_Pledge *PledgeCaller) GetUserDepositsMininAddrLength(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Pledge.contract.Call(opts, &out, "getUserDepositsMininAddrLength")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetUserDepositsMininAddrLength is a free data retrieval call binding the contract method 0x9cf205bd.
//
// Solidity: function getUserDepositsMininAddrLength() view returns(uint256)
func (_Pledge *PledgeSession) GetUserDepositsMininAddrLength() (*big.Int, error) {
	return _Pledge.Contract.GetUserDepositsMininAddrLength(&_Pledge.CallOpts)
}

// GetUserDepositsMininAddrLength is a free data retrieval call binding the contract method 0x9cf205bd.
//
// Solidity: function getUserDepositsMininAddrLength() view returns(uint256)
func (_Pledge *PledgeCallerSession) GetUserDepositsMininAddrLength() (*big.Int, error) {
	return _Pledge.Contract.GetUserDepositsMininAddrLength(&_Pledge.CallOpts)
}

// GetUserLendingDeposits is a free data retrieval call binding the contract method 0x013f9435.
//
// Solidity: function getUserLendingDeposits(address _sender) view returns((uint256,address,uint32,uint32,uint32,uint256,bool)[])
func (_Pledge *PledgeCaller) GetUserLendingDeposits(opts *bind.CallOpts, _sender common.Address) ([]PledgeUserInfo, error) {
	var out []interface{}
	err := _Pledge.contract.Call(opts, &out, "getUserLendingDeposits", _sender)

	if err != nil {
		return *new([]PledgeUserInfo), err
	}

	out0 := *abi.ConvertType(out[0], new([]PledgeUserInfo)).(*[]PledgeUserInfo)

	return out0, err

}

// GetUserLendingDeposits is a free data retrieval call binding the contract method 0x013f9435.
//
// Solidity: function getUserLendingDeposits(address _sender) view returns((uint256,address,uint32,uint32,uint32,uint256,bool)[])
func (_Pledge *PledgeSession) GetUserLendingDeposits(_sender common.Address) ([]PledgeUserInfo, error) {
	return _Pledge.Contract.GetUserLendingDeposits(&_Pledge.CallOpts, _sender)
}

// GetUserLendingDeposits is a free data retrieval call binding the contract method 0x013f9435.
//
// Solidity: function getUserLendingDeposits(address _sender) view returns((uint256,address,uint32,uint32,uint32,uint256,bool)[])
func (_Pledge *PledgeCallerSession) GetUserLendingDeposits(_sender common.Address) ([]PledgeUserInfo, error) {
	return _Pledge.Contract.GetUserLendingDeposits(&_Pledge.CallOpts, _sender)
}

// GetUserWithdraws is a free data retrieval call binding the contract method 0x515d223a.
//
// Solidity: function getUserWithdraws(address user) view returns((uint256,uint256,uint256,uint256)[])
func (_Pledge *PledgeCaller) GetUserWithdraws(opts *bind.CallOpts, user common.Address) ([]PledgeWithdrawInfo, error) {
	var out []interface{}
	err := _Pledge.contract.Call(opts, &out, "getUserWithdraws", user)

	if err != nil {
		return *new([]PledgeWithdrawInfo), err
	}

	out0 := *abi.ConvertType(out[0], new([]PledgeWithdrawInfo)).(*[]PledgeWithdrawInfo)

	return out0, err

}

// GetUserWithdraws is a free data retrieval call binding the contract method 0x515d223a.
//
// Solidity: function getUserWithdraws(address user) view returns((uint256,uint256,uint256,uint256)[])
func (_Pledge *PledgeSession) GetUserWithdraws(user common.Address) ([]PledgeWithdrawInfo, error) {
	return _Pledge.Contract.GetUserWithdraws(&_Pledge.CallOpts, user)
}

// GetUserWithdraws is a free data retrieval call binding the contract method 0x515d223a.
//
// Solidity: function getUserWithdraws(address user) view returns((uint256,uint256,uint256,uint256)[])
func (_Pledge *PledgeCallerSession) GetUserWithdraws(user common.Address) ([]PledgeWithdrawInfo, error) {
	return _Pledge.Contract.GetUserWithdraws(&_Pledge.CallOpts, user)
}

// GetisWhiteList is a free data retrieval call binding the contract method 0x2154d560.
//
// Solidity: function getisWhiteList() view returns(bool)
func (_Pledge *PledgeCaller) GetisWhiteList(opts *bind.CallOpts) (bool, error) {
	var out []interface{}
	err := _Pledge.contract.Call(opts, &out, "getisWhiteList")

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// GetisWhiteList is a free data retrieval call binding the contract method 0x2154d560.
//
// Solidity: function getisWhiteList() view returns(bool)
func (_Pledge *PledgeSession) GetisWhiteList() (bool, error) {
	return _Pledge.Contract.GetisWhiteList(&_Pledge.CallOpts)
}

// GetisWhiteList is a free data retrieval call binding the contract method 0x2154d560.
//
// Solidity: function getisWhiteList() view returns(bool)
func (_Pledge *PledgeCallerSession) GetisWhiteList() (bool, error) {
	return _Pledge.Contract.GetisWhiteList(&_Pledge.CallOpts)
}

// IsExtisAddress is a free data retrieval call binding the contract method 0x2480f68c.
//
// Solidity: function isExtisAddress(address user) view returns(bool)
func (_Pledge *PledgeCaller) IsExtisAddress(opts *bind.CallOpts, user common.Address) (bool, error) {
	var out []interface{}
	err := _Pledge.contract.Call(opts, &out, "isExtisAddress", user)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsExtisAddress is a free data retrieval call binding the contract method 0x2480f68c.
//
// Solidity: function isExtisAddress(address user) view returns(bool)
func (_Pledge *PledgeSession) IsExtisAddress(user common.Address) (bool, error) {
	return _Pledge.Contract.IsExtisAddress(&_Pledge.CallOpts, user)
}

// IsExtisAddress is a free data retrieval call binding the contract method 0x2480f68c.
//
// Solidity: function isExtisAddress(address user) view returns(bool)
func (_Pledge *PledgeCallerSession) IsExtisAddress(user common.Address) (bool, error) {
	return _Pledge.Contract.IsExtisAddress(&_Pledge.CallOpts, user)
}

// IsExtisBathAddress is a free data retrieval call binding the contract method 0xc3d30934.
//
// Solidity: function isExtisBathAddress(address[] users) view returns(bool)
func (_Pledge *PledgeCaller) IsExtisBathAddress(opts *bind.CallOpts, users []common.Address) (bool, error) {
	var out []interface{}
	err := _Pledge.contract.Call(opts, &out, "isExtisBathAddress", users)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsExtisBathAddress is a free data retrieval call binding the contract method 0xc3d30934.
//
// Solidity: function isExtisBathAddress(address[] users) view returns(bool)
func (_Pledge *PledgeSession) IsExtisBathAddress(users []common.Address) (bool, error) {
	return _Pledge.Contract.IsExtisBathAddress(&_Pledge.CallOpts, users)
}

// IsExtisBathAddress is a free data retrieval call binding the contract method 0xc3d30934.
//
// Solidity: function isExtisBathAddress(address[] users) view returns(bool)
func (_Pledge *PledgeCallerSession) IsExtisBathAddress(users []common.Address) (bool, error) {
	return _Pledge.Contract.IsExtisBathAddress(&_Pledge.CallOpts, users)
}

// IsValidateWhite is a free data retrieval call binding the contract method 0xac574946.
//
// Solidity: function isValidateWhite(address user) view returns(bool)
func (_Pledge *PledgeCaller) IsValidateWhite(opts *bind.CallOpts, user common.Address) (bool, error) {
	var out []interface{}
	err := _Pledge.contract.Call(opts, &out, "isValidateWhite", user)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsValidateWhite is a free data retrieval call binding the contract method 0xac574946.
//
// Solidity: function isValidateWhite(address user) view returns(bool)
func (_Pledge *PledgeSession) IsValidateWhite(user common.Address) (bool, error) {
	return _Pledge.Contract.IsValidateWhite(&_Pledge.CallOpts, user)
}

// IsValidateWhite is a free data retrieval call binding the contract method 0xac574946.
//
// Solidity: function isValidateWhite(address user) view returns(bool)
func (_Pledge *PledgeCallerSession) IsValidateWhite(user common.Address) (bool, error) {
	return _Pledge.Contract.IsValidateWhite(&_Pledge.CallOpts, user)
}

// Iseffective is a free data retrieval call binding the contract method 0x13e286c4.
//
// Solidity: function iseffective(address user) view returns(bool)
func (_Pledge *PledgeCaller) Iseffective(opts *bind.CallOpts, user common.Address) (bool, error) {
	var out []interface{}
	err := _Pledge.contract.Call(opts, &out, "iseffective", user)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// Iseffective is a free data retrieval call binding the contract method 0x13e286c4.
//
// Solidity: function iseffective(address user) view returns(bool)
func (_Pledge *PledgeSession) Iseffective(user common.Address) (bool, error) {
	return _Pledge.Contract.Iseffective(&_Pledge.CallOpts, user)
}

// Iseffective is a free data retrieval call binding the contract method 0x13e286c4.
//
// Solidity: function iseffective(address user) view returns(bool)
func (_Pledge *PledgeCallerSession) Iseffective(user common.Address) (bool, error) {
	return _Pledge.Contract.Iseffective(&_Pledge.CallOpts, user)
}

// IseffectiveNew is a free data retrieval call binding the contract method 0xb39a8baa.
//
// Solidity: function iseffectiveNew(address user) view returns(bool, address)
func (_Pledge *PledgeCaller) IseffectiveNew(opts *bind.CallOpts, user common.Address) (bool, common.Address, error) {
	var out []interface{}
	err := _Pledge.contract.Call(opts, &out, "iseffectiveNew", user)

	if err != nil {
		return *new(bool), *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)
	out1 := *abi.ConvertType(out[1], new(common.Address)).(*common.Address)

	return out0, out1, err

}

// IseffectiveNew is a free data retrieval call binding the contract method 0xb39a8baa.
//
// Solidity: function iseffectiveNew(address user) view returns(bool, address)
func (_Pledge *PledgeSession) IseffectiveNew(user common.Address) (bool, common.Address, error) {
	return _Pledge.Contract.IseffectiveNew(&_Pledge.CallOpts, user)
}

// IseffectiveNew is a free data retrieval call binding the contract method 0xb39a8baa.
//
// Solidity: function iseffectiveNew(address user) view returns(bool, address)
func (_Pledge *PledgeCallerSession) IseffectiveNew(user common.Address) (bool, common.Address, error) {
	return _Pledge.Contract.IseffectiveNew(&_Pledge.CallOpts, user)
}

// MinPledgeAmount is a free data retrieval call binding the contract method 0xec555e91.
//
// Solidity: function minPledgeAmount() view returns(uint256)
func (_Pledge *PledgeCaller) MinPledgeAmount(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Pledge.contract.Call(opts, &out, "minPledgeAmount")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// MinPledgeAmount is a free data retrieval call binding the contract method 0xec555e91.
//
// Solidity: function minPledgeAmount() view returns(uint256)
func (_Pledge *PledgeSession) MinPledgeAmount() (*big.Int, error) {
	return _Pledge.Contract.MinPledgeAmount(&_Pledge.CallOpts)
}

// MinPledgeAmount is a free data retrieval call binding the contract method 0xec555e91.
//
// Solidity: function minPledgeAmount() view returns(uint256)
func (_Pledge *PledgeCallerSession) MinPledgeAmount() (*big.Int, error) {
	return _Pledge.Contract.MinPledgeAmount(&_Pledge.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_Pledge *PledgeCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Pledge.contract.Call(opts, &out, "owner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_Pledge *PledgeSession) Owner() (common.Address, error) {
	return _Pledge.Contract.Owner(&_Pledge.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_Pledge *PledgeCallerSession) Owner() (common.Address, error) {
	return _Pledge.Contract.Owner(&_Pledge.CallOpts)
}

// TotalBorrows is a free data retrieval call binding the contract method 0x47bd3718.
//
// Solidity: function totalBorrows() view returns(uint256)
func (_Pledge *PledgeCaller) TotalBorrows(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Pledge.contract.Call(opts, &out, "totalBorrows")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// TotalBorrows is a free data retrieval call binding the contract method 0x47bd3718.
//
// Solidity: function totalBorrows() view returns(uint256)
func (_Pledge *PledgeSession) TotalBorrows() (*big.Int, error) {
	return _Pledge.Contract.TotalBorrows(&_Pledge.CallOpts)
}

// TotalBorrows is a free data retrieval call binding the contract method 0x47bd3718.
//
// Solidity: function totalBorrows() view returns(uint256)
func (_Pledge *PledgeCallerSession) TotalBorrows() (*big.Int, error) {
	return _Pledge.Contract.TotalBorrows(&_Pledge.CallOpts)
}

// UserDepositsLending is a free data retrieval call binding the contract method 0xb787774a.
//
// Solidity: function userDepositsLending(address , uint256 ) view returns(uint256)
func (_Pledge *PledgeCaller) UserDepositsLending(opts *bind.CallOpts, arg0 common.Address, arg1 *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _Pledge.contract.Call(opts, &out, "userDepositsLending", arg0, arg1)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// UserDepositsLending is a free data retrieval call binding the contract method 0xb787774a.
//
// Solidity: function userDepositsLending(address , uint256 ) view returns(uint256)
func (_Pledge *PledgeSession) UserDepositsLending(arg0 common.Address, arg1 *big.Int) (*big.Int, error) {
	return _Pledge.Contract.UserDepositsLending(&_Pledge.CallOpts, arg0, arg1)
}

// UserDepositsLending is a free data retrieval call binding the contract method 0xb787774a.
//
// Solidity: function userDepositsLending(address , uint256 ) view returns(uint256)
func (_Pledge *PledgeCallerSession) UserDepositsLending(arg0 common.Address, arg1 *big.Int) (*big.Int, error) {
	return _Pledge.Contract.UserDepositsLending(&_Pledge.CallOpts, arg0, arg1)
}

// AddMininAddr is a paid mutator transaction binding the contract method 0xf0cbcb68.
//
// Solidity: function addMininAddr(address _address, bool _flag) returns()
func (_Pledge *PledgeTransactor) AddMininAddr(opts *bind.TransactOpts, _address common.Address, _flag bool) (*types.Transaction, error) {
	return _Pledge.contract.Transact(opts, "addMininAddr", _address, _flag)
}

// AddMininAddr is a paid mutator transaction binding the contract method 0xf0cbcb68.
//
// Solidity: function addMininAddr(address _address, bool _flag) returns()
func (_Pledge *PledgeSession) AddMininAddr(_address common.Address, _flag bool) (*types.Transaction, error) {
	return _Pledge.Contract.AddMininAddr(&_Pledge.TransactOpts, _address, _flag)
}

// AddMininAddr is a paid mutator transaction binding the contract method 0xf0cbcb68.
//
// Solidity: function addMininAddr(address _address, bool _flag) returns()
func (_Pledge *PledgeTransactorSession) AddMininAddr(_address common.Address, _flag bool) (*types.Transaction, error) {
	return _Pledge.Contract.AddMininAddr(&_Pledge.TransactOpts, _address, _flag)
}

// Deposit is a paid mutator transaction binding the contract method 0xf13e0b13.
//
// Solidity: function deposit(uint256 _nodes, uint256 _amount, address[] _addrs, uint256 _expiredAt) payable returns()
func (_Pledge *PledgeTransactor) Deposit(opts *bind.TransactOpts, _nodes *big.Int, _amount *big.Int, _addrs []common.Address, _expiredAt *big.Int) (*types.Transaction, error) {
	return _Pledge.contract.Transact(opts, "deposit", _nodes, _amount, _addrs, _expiredAt)
}

// Deposit is a paid mutator transaction binding the contract method 0xf13e0b13.
//
// Solidity: function deposit(uint256 _nodes, uint256 _amount, address[] _addrs, uint256 _expiredAt) payable returns()
func (_Pledge *PledgeSession) Deposit(_nodes *big.Int, _amount *big.Int, _addrs []common.Address, _expiredAt *big.Int) (*types.Transaction, error) {
	return _Pledge.Contract.Deposit(&_Pledge.TransactOpts, _nodes, _amount, _addrs, _expiredAt)
}

// Deposit is a paid mutator transaction binding the contract method 0xf13e0b13.
//
// Solidity: function deposit(uint256 _nodes, uint256 _amount, address[] _addrs, uint256 _expiredAt) payable returns()
func (_Pledge *PledgeTransactorSession) Deposit(_nodes *big.Int, _amount *big.Int, _addrs []common.Address, _expiredAt *big.Int) (*types.Transaction, error) {
	return _Pledge.Contract.Deposit(&_Pledge.TransactOpts, _nodes, _amount, _addrs, _expiredAt)
}

// DepositLending is a paid mutator transaction binding the contract method 0xa0fb9f0d.
//
// Solidity: function depositLending(address _sender, uint256 _nodes, uint256 _amount, uint256 _expiredAt, address[] _addrs) payable returns()
func (_Pledge *PledgeTransactor) DepositLending(opts *bind.TransactOpts, _sender common.Address, _nodes *big.Int, _amount *big.Int, _expiredAt *big.Int, _addrs []common.Address) (*types.Transaction, error) {
	return _Pledge.contract.Transact(opts, "depositLending", _sender, _nodes, _amount, _expiredAt, _addrs)
}

// DepositLending is a paid mutator transaction binding the contract method 0xa0fb9f0d.
//
// Solidity: function depositLending(address _sender, uint256 _nodes, uint256 _amount, uint256 _expiredAt, address[] _addrs) payable returns()
func (_Pledge *PledgeSession) DepositLending(_sender common.Address, _nodes *big.Int, _amount *big.Int, _expiredAt *big.Int, _addrs []common.Address) (*types.Transaction, error) {
	return _Pledge.Contract.DepositLending(&_Pledge.TransactOpts, _sender, _nodes, _amount, _expiredAt, _addrs)
}

// DepositLending is a paid mutator transaction binding the contract method 0xa0fb9f0d.
//
// Solidity: function depositLending(address _sender, uint256 _nodes, uint256 _amount, uint256 _expiredAt, address[] _addrs) payable returns()
func (_Pledge *PledgeTransactorSession) DepositLending(_sender common.Address, _nodes *big.Int, _amount *big.Int, _expiredAt *big.Int, _addrs []common.Address) (*types.Transaction, error) {
	return _Pledge.Contract.DepositLending(&_Pledge.TransactOpts, _sender, _nodes, _amount, _expiredAt, _addrs)
}

// Initialize is a paid mutator transaction binding the contract method 0x8129fc1c.
//
// Solidity: function initialize() returns()
func (_Pledge *PledgeTransactor) Initialize(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Pledge.contract.Transact(opts, "initialize")
}

// Initialize is a paid mutator transaction binding the contract method 0x8129fc1c.
//
// Solidity: function initialize() returns()
func (_Pledge *PledgeSession) Initialize() (*types.Transaction, error) {
	return _Pledge.Contract.Initialize(&_Pledge.TransactOpts)
}

// Initialize is a paid mutator transaction binding the contract method 0x8129fc1c.
//
// Solidity: function initialize() returns()
func (_Pledge *PledgeTransactorSession) Initialize() (*types.Transaction, error) {
	return _Pledge.Contract.Initialize(&_Pledge.TransactOpts)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_Pledge *PledgeTransactor) RenounceOwnership(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Pledge.contract.Transact(opts, "renounceOwnership")
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_Pledge *PledgeSession) RenounceOwnership() (*types.Transaction, error) {
	return _Pledge.Contract.RenounceOwnership(&_Pledge.TransactOpts)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_Pledge *PledgeTransactorSession) RenounceOwnership() (*types.Transaction, error) {
	return _Pledge.Contract.RenounceOwnership(&_Pledge.TransactOpts)
}

// SetAddressQLF is a paid mutator transaction binding the contract method 0x13f3c4b9.
//
// Solidity: function setAddressQLF(address _qlf, bool _isWhiteList) returns()
func (_Pledge *PledgeTransactor) SetAddressQLF(opts *bind.TransactOpts, _qlf common.Address, _isWhiteList bool) (*types.Transaction, error) {
	return _Pledge.contract.Transact(opts, "setAddressQLF", _qlf, _isWhiteList)
}

// SetAddressQLF is a paid mutator transaction binding the contract method 0x13f3c4b9.
//
// Solidity: function setAddressQLF(address _qlf, bool _isWhiteList) returns()
func (_Pledge *PledgeSession) SetAddressQLF(_qlf common.Address, _isWhiteList bool) (*types.Transaction, error) {
	return _Pledge.Contract.SetAddressQLF(&_Pledge.TransactOpts, _qlf, _isWhiteList)
}

// SetAddressQLF is a paid mutator transaction binding the contract method 0x13f3c4b9.
//
// Solidity: function setAddressQLF(address _qlf, bool _isWhiteList) returns()
func (_Pledge *PledgeTransactorSession) SetAddressQLF(_qlf common.Address, _isWhiteList bool) (*types.Transaction, error) {
	return _Pledge.Contract.SetAddressQLF(&_Pledge.TransactOpts, _qlf, _isWhiteList)
}

// SetLockTime is a paid mutator transaction binding the contract method 0xae04d45d.
//
// Solidity: function setLockTime(uint256 _lockTime) returns()
func (_Pledge *PledgeTransactor) SetLockTime(opts *bind.TransactOpts, _lockTime *big.Int) (*types.Transaction, error) {
	return _Pledge.contract.Transact(opts, "setLockTime", _lockTime)
}

// SetLockTime is a paid mutator transaction binding the contract method 0xae04d45d.
//
// Solidity: function setLockTime(uint256 _lockTime) returns()
func (_Pledge *PledgeSession) SetLockTime(_lockTime *big.Int) (*types.Transaction, error) {
	return _Pledge.Contract.SetLockTime(&_Pledge.TransactOpts, _lockTime)
}

// SetLockTime is a paid mutator transaction binding the contract method 0xae04d45d.
//
// Solidity: function setLockTime(uint256 _lockTime) returns()
func (_Pledge *PledgeTransactorSession) SetLockTime(_lockTime *big.Int) (*types.Transaction, error) {
	return _Pledge.Contract.SetLockTime(&_Pledge.TransactOpts, _lockTime)
}

// SetMinPledgeAmount is a paid mutator transaction binding the contract method 0x573fe31f.
//
// Solidity: function setMinPledgeAmount(uint256 _time) returns()
func (_Pledge *PledgeTransactor) SetMinPledgeAmount(opts *bind.TransactOpts, _time *big.Int) (*types.Transaction, error) {
	return _Pledge.contract.Transact(opts, "setMinPledgeAmount", _time)
}

// SetMinPledgeAmount is a paid mutator transaction binding the contract method 0x573fe31f.
//
// Solidity: function setMinPledgeAmount(uint256 _time) returns()
func (_Pledge *PledgeSession) SetMinPledgeAmount(_time *big.Int) (*types.Transaction, error) {
	return _Pledge.Contract.SetMinPledgeAmount(&_Pledge.TransactOpts, _time)
}

// SetMinPledgeAmount is a paid mutator transaction binding the contract method 0x573fe31f.
//
// Solidity: function setMinPledgeAmount(uint256 _time) returns()
func (_Pledge *PledgeTransactorSession) SetMinPledgeAmount(_time *big.Int) (*types.Transaction, error) {
	return _Pledge.Contract.SetMinPledgeAmount(&_Pledge.TransactOpts, _time)
}

// SetPaused is a paid mutator transaction binding the contract method 0x6426be48.
//
// Solidity: function setPaused(bool _isDepositPaused, bool _isWithdrawPaused) returns()
func (_Pledge *PledgeTransactor) SetPaused(opts *bind.TransactOpts, _isDepositPaused bool, _isWithdrawPaused bool) (*types.Transaction, error) {
	return _Pledge.contract.Transact(opts, "setPaused", _isDepositPaused, _isWithdrawPaused)
}

// SetPaused is a paid mutator transaction binding the contract method 0x6426be48.
//
// Solidity: function setPaused(bool _isDepositPaused, bool _isWithdrawPaused) returns()
func (_Pledge *PledgeSession) SetPaused(_isDepositPaused bool, _isWithdrawPaused bool) (*types.Transaction, error) {
	return _Pledge.Contract.SetPaused(&_Pledge.TransactOpts, _isDepositPaused, _isWithdrawPaused)
}

// SetPaused is a paid mutator transaction binding the contract method 0x6426be48.
//
// Solidity: function setPaused(bool _isDepositPaused, bool _isWithdrawPaused) returns()
func (_Pledge *PledgeTransactorSession) SetPaused(_isDepositPaused bool, _isWithdrawPaused bool) (*types.Transaction, error) {
	return _Pledge.Contract.SetPaused(&_Pledge.TransactOpts, _isDepositPaused, _isWithdrawPaused)
}

// SetProjectStartTime is a paid mutator transaction binding the contract method 0xc6d1f78d.
//
// Solidity: function setProjectStartTime(uint256 _projectStartTime) returns()
func (_Pledge *PledgeTransactor) SetProjectStartTime(opts *bind.TransactOpts, _projectStartTime *big.Int) (*types.Transaction, error) {
	return _Pledge.contract.Transact(opts, "setProjectStartTime", _projectStartTime)
}

// SetProjectStartTime is a paid mutator transaction binding the contract method 0xc6d1f78d.
//
// Solidity: function setProjectStartTime(uint256 _projectStartTime) returns()
func (_Pledge *PledgeSession) SetProjectStartTime(_projectStartTime *big.Int) (*types.Transaction, error) {
	return _Pledge.Contract.SetProjectStartTime(&_Pledge.TransactOpts, _projectStartTime)
}

// SetProjectStartTime is a paid mutator transaction binding the contract method 0xc6d1f78d.
//
// Solidity: function setProjectStartTime(uint256 _projectStartTime) returns()
func (_Pledge *PledgeTransactorSession) SetProjectStartTime(_projectStartTime *big.Int) (*types.Transaction, error) {
	return _Pledge.Contract.SetProjectStartTime(&_Pledge.TransactOpts, _projectStartTime)
}

// SetdataMigration is a paid mutator transaction binding the contract method 0x6e2c9c91.
//
// Solidity: function setdataMigration(address _address, uint256 _expireAt) returns()
func (_Pledge *PledgeTransactor) SetdataMigration(opts *bind.TransactOpts, _address common.Address, _expireAt *big.Int) (*types.Transaction, error) {
	return _Pledge.contract.Transact(opts, "setdataMigration", _address, _expireAt)
}

// SetdataMigration is a paid mutator transaction binding the contract method 0x6e2c9c91.
//
// Solidity: function setdataMigration(address _address, uint256 _expireAt) returns()
func (_Pledge *PledgeSession) SetdataMigration(_address common.Address, _expireAt *big.Int) (*types.Transaction, error) {
	return _Pledge.Contract.SetdataMigration(&_Pledge.TransactOpts, _address, _expireAt)
}

// SetdataMigration is a paid mutator transaction binding the contract method 0x6e2c9c91.
//
// Solidity: function setdataMigration(address _address, uint256 _expireAt) returns()
func (_Pledge *PledgeTransactorSession) SetdataMigration(_address common.Address, _expireAt *big.Int) (*types.Transaction, error) {
	return _Pledge.Contract.SetdataMigration(&_Pledge.TransactOpts, _address, _expireAt)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_Pledge *PledgeTransactor) TransferOwnership(opts *bind.TransactOpts, newOwner common.Address) (*types.Transaction, error) {
	return _Pledge.contract.Transact(opts, "transferOwnership", newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_Pledge *PledgeSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _Pledge.Contract.TransferOwnership(&_Pledge.TransactOpts, newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_Pledge *PledgeTransactorSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _Pledge.Contract.TransferOwnership(&_Pledge.TransactOpts, newOwner)
}

// UpdateDepositsCoinbaseAddr is a paid mutator transaction binding the contract method 0x01d3a7de.
//
// Solidity: function updateDepositsCoinbaseAddr(address _replace) returns()
func (_Pledge *PledgeTransactor) UpdateDepositsCoinbaseAddr(opts *bind.TransactOpts, _replace common.Address) (*types.Transaction, error) {
	return _Pledge.contract.Transact(opts, "updateDepositsCoinbaseAddr", _replace)
}

// UpdateDepositsCoinbaseAddr is a paid mutator transaction binding the contract method 0x01d3a7de.
//
// Solidity: function updateDepositsCoinbaseAddr(address _replace) returns()
func (_Pledge *PledgeSession) UpdateDepositsCoinbaseAddr(_replace common.Address) (*types.Transaction, error) {
	return _Pledge.Contract.UpdateDepositsCoinbaseAddr(&_Pledge.TransactOpts, _replace)
}

// UpdateDepositsCoinbaseAddr is a paid mutator transaction binding the contract method 0x01d3a7de.
//
// Solidity: function updateDepositsCoinbaseAddr(address _replace) returns()
func (_Pledge *PledgeTransactorSession) UpdateDepositsCoinbaseAddr(_replace common.Address) (*types.Transaction, error) {
	return _Pledge.Contract.UpdateDepositsCoinbaseAddr(&_Pledge.TransactOpts, _replace)
}

// UpdateDepositsMininAddress is a paid mutator transaction binding the contract method 0xc91048c7.
//
// Solidity: function updateDepositsMininAddress(address _user, address _oldAddress, uint256 _index, address _replace) returns()
func (_Pledge *PledgeTransactor) UpdateDepositsMininAddress(opts *bind.TransactOpts, _user common.Address, _oldAddress common.Address, _index *big.Int, _replace common.Address) (*types.Transaction, error) {
	return _Pledge.contract.Transact(opts, "updateDepositsMininAddress", _user, _oldAddress, _index, _replace)
}

// UpdateDepositsMininAddress is a paid mutator transaction binding the contract method 0xc91048c7.
//
// Solidity: function updateDepositsMininAddress(address _user, address _oldAddress, uint256 _index, address _replace) returns()
func (_Pledge *PledgeSession) UpdateDepositsMininAddress(_user common.Address, _oldAddress common.Address, _index *big.Int, _replace common.Address) (*types.Transaction, error) {
	return _Pledge.Contract.UpdateDepositsMininAddress(&_Pledge.TransactOpts, _user, _oldAddress, _index, _replace)
}

// UpdateDepositsMininAddress is a paid mutator transaction binding the contract method 0xc91048c7.
//
// Solidity: function updateDepositsMininAddress(address _user, address _oldAddress, uint256 _index, address _replace) returns()
func (_Pledge *PledgeTransactorSession) UpdateDepositsMininAddress(_user common.Address, _oldAddress common.Address, _index *big.Int, _replace common.Address) (*types.Transaction, error) {
	return _Pledge.Contract.UpdateDepositsMininAddress(&_Pledge.TransactOpts, _user, _oldAddress, _index, _replace)
}

// Withdraw is a paid mutator transaction binding the contract method 0x00f714ce.
//
// Solidity: function withdraw(uint256 _index, address _receiver) returns()
func (_Pledge *PledgeTransactor) Withdraw(opts *bind.TransactOpts, _index *big.Int, _receiver common.Address) (*types.Transaction, error) {
	return _Pledge.contract.Transact(opts, "withdraw", _index, _receiver)
}

// Withdraw is a paid mutator transaction binding the contract method 0x00f714ce.
//
// Solidity: function withdraw(uint256 _index, address _receiver) returns()
func (_Pledge *PledgeSession) Withdraw(_index *big.Int, _receiver common.Address) (*types.Transaction, error) {
	return _Pledge.Contract.Withdraw(&_Pledge.TransactOpts, _index, _receiver)
}

// Withdraw is a paid mutator transaction binding the contract method 0x00f714ce.
//
// Solidity: function withdraw(uint256 _index, address _receiver) returns()
func (_Pledge *PledgeTransactorSession) Withdraw(_index *big.Int, _receiver common.Address) (*types.Transaction, error) {
	return _Pledge.Contract.Withdraw(&_Pledge.TransactOpts, _index, _receiver)
}

// Receive is a paid mutator transaction binding the contract receive function.
//
// Solidity: receive() payable returns()
func (_Pledge *PledgeTransactor) Receive(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Pledge.contract.RawTransact(opts, nil) // calldata is disallowed for receive function
}

// Receive is a paid mutator transaction binding the contract receive function.
//
// Solidity: receive() payable returns()
func (_Pledge *PledgeSession) Receive() (*types.Transaction, error) {
	return _Pledge.Contract.Receive(&_Pledge.TransactOpts)
}

// Receive is a paid mutator transaction binding the contract receive function.
//
// Solidity: receive() payable returns()
func (_Pledge *PledgeTransactorSession) Receive() (*types.Transaction, error) {
	return _Pledge.Contract.Receive(&_Pledge.TransactOpts)
}

// PledgeDepositIterator is returned from FilterDeposit and is used to iterate over the raw logs and unpacked data for Deposit events raised by the Pledge contract.
type PledgeDepositIterator struct {
	Event *PledgeDeposit // Event containing the contract specifics and raw log

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
func (it *PledgeDepositIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(PledgeDeposit)
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
		it.Event = new(PledgeDeposit)
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
func (it *PledgeDepositIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *PledgeDepositIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// PledgeDeposit represents a Deposit event raised by the Pledge contract.
type PledgeDeposit struct {
	Index  *big.Int
	User   common.Address
	Amount *big.Int
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterDeposit is a free log retrieval operation binding the contract event 0xeaa18152488ce5959073c9c79c88ca90b3d96c00de1f118cfaad664c3dab06b9.
//
// Solidity: event Deposit(uint256 index, address indexed user, uint256 amount)
func (_Pledge *PledgeFilterer) FilterDeposit(opts *bind.FilterOpts, user []common.Address) (*PledgeDepositIterator, error) {

	var userRule []interface{}
	for _, userItem := range user {
		userRule = append(userRule, userItem)
	}

	logs, sub, err := _Pledge.contract.FilterLogs(opts, "Deposit", userRule)
	if err != nil {
		return nil, err
	}
	return &PledgeDepositIterator{contract: _Pledge.contract, event: "Deposit", logs: logs, sub: sub}, nil
}

// WatchDeposit is a free log subscription operation binding the contract event 0xeaa18152488ce5959073c9c79c88ca90b3d96c00de1f118cfaad664c3dab06b9.
//
// Solidity: event Deposit(uint256 index, address indexed user, uint256 amount)
func (_Pledge *PledgeFilterer) WatchDeposit(opts *bind.WatchOpts, sink chan<- *PledgeDeposit, user []common.Address) (event.Subscription, error) {

	var userRule []interface{}
	for _, userItem := range user {
		userRule = append(userRule, userItem)
	}

	logs, sub, err := _Pledge.contract.WatchLogs(opts, "Deposit", userRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(PledgeDeposit)
				if err := _Pledge.contract.UnpackLog(event, "Deposit", log); err != nil {
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

// ParseDeposit is a log parse operation binding the contract event 0xeaa18152488ce5959073c9c79c88ca90b3d96c00de1f118cfaad664c3dab06b9.
//
// Solidity: event Deposit(uint256 index, address indexed user, uint256 amount)
func (_Pledge *PledgeFilterer) ParseDeposit(log types.Log) (*PledgeDeposit, error) {
	event := new(PledgeDeposit)
	if err := _Pledge.contract.UnpackLog(event, "Deposit", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// PledgeOwnershipTransferredIterator is returned from FilterOwnershipTransferred and is used to iterate over the raw logs and unpacked data for OwnershipTransferred events raised by the Pledge contract.
type PledgeOwnershipTransferredIterator struct {
	Event *PledgeOwnershipTransferred // Event containing the contract specifics and raw log

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
func (it *PledgeOwnershipTransferredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(PledgeOwnershipTransferred)
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
		it.Event = new(PledgeOwnershipTransferred)
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
func (it *PledgeOwnershipTransferredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *PledgeOwnershipTransferredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// PledgeOwnershipTransferred represents a OwnershipTransferred event raised by the Pledge contract.
type PledgeOwnershipTransferred struct {
	PreviousOwner common.Address
	NewOwner      common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterOwnershipTransferred is a free log retrieval operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_Pledge *PledgeFilterer) FilterOwnershipTransferred(opts *bind.FilterOpts, previousOwner []common.Address, newOwner []common.Address) (*PledgeOwnershipTransferredIterator, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _Pledge.contract.FilterLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return &PledgeOwnershipTransferredIterator{contract: _Pledge.contract, event: "OwnershipTransferred", logs: logs, sub: sub}, nil
}

// WatchOwnershipTransferred is a free log subscription operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_Pledge *PledgeFilterer) WatchOwnershipTransferred(opts *bind.WatchOpts, sink chan<- *PledgeOwnershipTransferred, previousOwner []common.Address, newOwner []common.Address) (event.Subscription, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _Pledge.contract.WatchLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(PledgeOwnershipTransferred)
				if err := _Pledge.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
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
func (_Pledge *PledgeFilterer) ParseOwnershipTransferred(log types.Log) (*PledgeOwnershipTransferred, error) {
	event := new(PledgeOwnershipTransferred)
	if err := _Pledge.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// PledgeWithdrawIterator is returned from FilterWithdraw and is used to iterate over the raw logs and unpacked data for Withdraw events raised by the Pledge contract.
type PledgeWithdrawIterator struct {
	Event *PledgeWithdraw // Event containing the contract specifics and raw log

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
func (it *PledgeWithdrawIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(PledgeWithdraw)
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
		it.Event = new(PledgeWithdraw)
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
func (it *PledgeWithdrawIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *PledgeWithdrawIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// PledgeWithdraw represents a Withdraw event raised by the Pledge contract.
type PledgeWithdraw struct {
	Index  *big.Int
	User   common.Address
	Amount *big.Int
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterWithdraw is a free log retrieval operation binding the contract event 0x9da6493a92039daf47d1f2d7a782299c5994c6323eb1e972f69c432089ec52bf.
//
// Solidity: event Withdraw(uint256 index, address indexed user, uint256 amount)
func (_Pledge *PledgeFilterer) FilterWithdraw(opts *bind.FilterOpts, user []common.Address) (*PledgeWithdrawIterator, error) {

	var userRule []interface{}
	for _, userItem := range user {
		userRule = append(userRule, userItem)
	}

	logs, sub, err := _Pledge.contract.FilterLogs(opts, "Withdraw", userRule)
	if err != nil {
		return nil, err
	}
	return &PledgeWithdrawIterator{contract: _Pledge.contract, event: "Withdraw", logs: logs, sub: sub}, nil
}

// WatchWithdraw is a free log subscription operation binding the contract event 0x9da6493a92039daf47d1f2d7a782299c5994c6323eb1e972f69c432089ec52bf.
//
// Solidity: event Withdraw(uint256 index, address indexed user, uint256 amount)
func (_Pledge *PledgeFilterer) WatchWithdraw(opts *bind.WatchOpts, sink chan<- *PledgeWithdraw, user []common.Address) (event.Subscription, error) {

	var userRule []interface{}
	for _, userItem := range user {
		userRule = append(userRule, userItem)
	}

	logs, sub, err := _Pledge.contract.WatchLogs(opts, "Withdraw", userRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(PledgeWithdraw)
				if err := _Pledge.contract.UnpackLog(event, "Withdraw", log); err != nil {
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

// ParseWithdraw is a log parse operation binding the contract event 0x9da6493a92039daf47d1f2d7a782299c5994c6323eb1e972f69c432089ec52bf.
//
// Solidity: event Withdraw(uint256 index, address indexed user, uint256 amount)
func (_Pledge *PledgeFilterer) ParseWithdraw(log types.Log) (*PledgeWithdraw, error) {
	event := new(PledgeWithdraw)
	if err := _Pledge.contract.UnpackLog(event, "Withdraw", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
