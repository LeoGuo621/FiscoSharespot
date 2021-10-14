// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package user_management

import (
	"math/big"
	"strings"

	"github.com/FISCO-BCOS/go-sdk/abi"
	"github.com/FISCO-BCOS/go-sdk/abi/bind"
	"github.com/FISCO-BCOS/go-sdk/core/types"
	"github.com/FISCO-BCOS/go-sdk/event"
	ethereum "github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/common"
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

// Struct0 is an auto generated low-level Go binding around an user-defined struct.
type Struct0 struct {
	Id              string
	Addr            common.Address
	Name            string
	PasswordEncoded [32]byte
}

// UserManagementABI is the input ABI used to generate the binding from.
const UserManagementABI = "[{\"constant\":true,\"inputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"users\",\"outputs\":[{\"name\":\"id\",\"type\":\"string\"},{\"name\":\"addr\",\"type\":\"address\"},{\"name\":\"name\",\"type\":\"string\"},{\"name\":\"passwordEncoded\",\"type\":\"bytes32\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"id\",\"type\":\"string\"},{\"name\":\"addr\",\"type\":\"address\"},{\"name\":\"name\",\"type\":\"string\"},{\"name\":\"passwordEncoded\",\"type\":\"bytes32\"}],\"name\":\"addUser\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"getAllUsers\",\"outputs\":[{\"components\":[{\"name\":\"id\",\"type\":\"string\"},{\"name\":\"addr\",\"type\":\"address\"},{\"name\":\"name\",\"type\":\"string\"},{\"name\":\"passwordEncoded\",\"type\":\"bytes32\"}],\"name\":\"\",\"type\":\"tuple[]\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"}]"

// UserManagementBin is the compiled bytecode used for deploying new contracts.
var UserManagementBin = "0x608060405234801561001057600080fd5b50610e6a806100206000396000f300608060405260043610610057576000357c0100000000000000000000000000000000000000000000000000000000900463ffffffff168063365b98b21461005c578063431a93f71461009c578063e2842d79146100d9575b600080fd5b34801561006857600080fd5b50610083600480360361007e9190810190610b5e565b610104565b6040516100939493929190610cff565b60405180910390f35b3480156100a857600080fd5b506100c360048036036100be9190810190610ac3565b610293565b6040516100d09190610ce4565b60405180910390f35b3480156100e557600080fd5b506100ee61047c565b6040516100fb9190610cc2565b60405180910390f35b60008181548110151561011357fe5b9060005260206000209060040201600091509050806000018054600181600116156101000203166002900480601f0160208091040260200160405190810160405280929190818152602001828054600181600116156101000203166002900480156101bf5780601f10610194576101008083540402835291602001916101bf565b820191906000526020600020905b8154815290600101906020018083116101a257829003601f168201915b5050505050908060010160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1690806002018054600181600116156101000203166002900480601f0160208091040260200160405190810160405280929190818152602001828054600181600116156101000203166002900480156102835780601f1061025857610100808354040283529160200191610283565b820191906000526020600020905b81548152906001019060200180831161026657829003601f168201915b5050505050908060030154905084565b600061029d610954565b7fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff6102f989898080601f016020809104026020016040519081016040528093929190818152602001838380828437820191505050505050610688565b10156103085760009150610471565b60806040519081016040528089898080601f01602080910402602001604051908101604052809392919081815260200183838082843782019150505050505081526020018773ffffffffffffffffffffffffffffffffffffffff16815260200186868080601f0160208091040260200160405190810160405280939291908181526020018383808284378201915050505050508152602001846000191681525090506000819080600181540180825580915050906001820390600052602060002090600402016000909192909190915060008201518160000190805190602001906103f4929190610996565b5060208201518160010160006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff1602179055506040820151816002019080519060200190610458929190610996565b5060608201518160030190600019169055505050600191505b509695505050505050565b60606000805480602002602001604051908101604052809291908181526020016000905b8282101561067f578382906000526020600020906004020160806040519081016040529081600082018054600181600116156101000203166002900480601f01602080910402602001604051908101604052809291908181526020018280546001816001161561010002031660029004801561055d5780601f106105325761010080835404028352916020019161055d565b820191906000526020600020905b81548152906001019060200180831161054057829003601f168201915b505050505081526020016001820160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff168152602001600282018054600181600116156101000203166002900480601f0160208091040260200160405190810160405280929190818152602001828054600181600116156101000203166002900480156106555780601f1061062a57610100808354040283529160200191610655565b820191906000526020600020905b81548152906001019060200180831161063857829003601f168201915b505050505081526020016003820154600019166000191681525050815260200190600101906104a0565b50505050905090565b600080600090505b6000805490508110156107795761075f6000828154811015156106af57fe5b90600052602060002090600402016000018054600181600116156101000203166002900480601f0160208091040260200160405190810160405280929190818152602001828054600181600116156101000203166002900480156107545780601f1061072957610100808354040283529160200191610754565b820191906000526020600020905b81548152906001019060200180831161073757829003601f168201915b5050505050846107a3565b1561076c5780915061079d565b8080600101915050610690565b7fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff91505b50919050565b6000816040516020018082805190602001908083835b6020831015156107de57805182526020820191506020810190506020830392506107b9565b6001836020036101000a0380198251168184511680821785525050505050509050019150506040516020818303038152906040526040518082805190602001908083835b6020831015156108475780518252602082019150602081019050602083039250610822565b6001836020036101000a038019825116818451168082178552505050505050905001915050604051809103902060001916836040516020018082805190602001908083835b6020831015156108b1578051825260208201915060208101905060208303925061088c565b6001836020036101000a0380198251168184511680821785525050505050509050019150506040516020818303038152906040526040518082805190602001908083835b60208310151561091a57805182526020820191506020810190506020830392506108f5565b6001836020036101000a03801982511681845116808217855250505050505090500191505060405180910390206000191614905092915050565b60806040519081016040528060608152602001600073ffffffffffffffffffffffffffffffffffffffff16815260200160608152602001600080191681525090565b828054600181600116156101000203166002900490600052602060002090601f016020900481019282601f106109d757805160ff1916838001178555610a05565b82800160010185558215610a05579182015b82811115610a045782518255916020019190600101906109e9565b5b509050610a129190610a16565b5090565b610a3891905b80821115610a34576000816000905550600101610a1c565b5090565b90565b6000610a478235610db8565b905092915050565b6000610a5b8235610dd8565b905092915050565b60008083601f8401121515610a7757600080fd5b8235905067ffffffffffffffff811115610a9057600080fd5b602083019150836001820283011115610aa857600080fd5b9250929050565b6000610abb8235610de2565b905092915050565b60008060008060008060808789031215610adc57600080fd5b600087013567ffffffffffffffff811115610af657600080fd5b610b0289828a01610a63565b96509650506020610b1589828a01610a3b565b945050604087013567ffffffffffffffff811115610b3257600080fd5b610b3e89828a01610a63565b93509350506060610b5189828a01610a4f565b9150509295509295509295565b600060208284031215610b7057600080fd5b6000610b7e84828501610aaf565b91505092915050565b610b9081610d82565b82525050565b6000610ba182610d5f565b80845260208401935083602082028501610bba85610d52565b60005b84811015610bf3578383038852610bd5838351610c58565b9250610be082610d75565b9150602088019750600181019050610bbd565b508196508694505050505092915050565b610c0d81610da2565b82525050565b610c1c81610dae565b82525050565b6000610c2d82610d6a565b808452610c41816020860160208601610dec565b610c4a81610e1f565b602085010191505092915050565b60006080830160008301518482036000860152610c758282610c22565b9150506020830151610c8a6020860182610b87565b5060408301518482036040860152610ca28282610c22565b9150506060830151610cb76060860182610c13565b508091505092915050565b60006020820190508181036000830152610cdc8184610b96565b905092915050565b6000602082019050610cf96000830184610c04565b92915050565b60006080820190508181036000830152610d198187610c22565b9050610d286020830186610b87565b8181036040830152610d3a8185610c22565b9050610d496060830184610c13565b95945050505050565b6000602082019050919050565b600081519050919050565b600081519050919050565b6000602082019050919050565b600073ffffffffffffffffffffffffffffffffffffffff82169050919050565b60008115159050919050565b6000819050919050565b600073ffffffffffffffffffffffffffffffffffffffff82169050919050565b6000819050919050565b6000819050919050565b60005b83811015610e0a578082015181840152602081019050610def565b83811115610e19576000848401525b50505050565b6000601f19601f83011690509190505600a265627a7a7230582057688b0a02030647d88e8c24acbb2a8a6c875841233e78c71f21169da85e13426c6578706572696d656e74616cf50037"

// DeployUserManagement deploys a new contract, binding an instance of UserManagement to it.
func DeployUserManagement(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *UserManagement, error) {
	parsed, err := abi.JSON(strings.NewReader(UserManagementABI))
	if err != nil {
		return common.Address{}, nil, nil, err
	}

	address, tx, contract, err := bind.DeployContract(auth, parsed, common.FromHex(UserManagementBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &UserManagement{UserManagementCaller: UserManagementCaller{contract: contract}, UserManagementTransactor: UserManagementTransactor{contract: contract}, UserManagementFilterer: UserManagementFilterer{contract: contract}}, nil
}

func AsyncDeployUserManagement(auth *bind.TransactOpts, handler func(*types.Receipt, error), backend bind.ContractBackend) (*types.Transaction, error) {
	parsed, err := abi.JSON(strings.NewReader(UserManagementABI))
	if err != nil {
		return nil, err
	}

	tx, err := bind.AsyncDeployContract(auth, handler, parsed, common.FromHex(UserManagementBin), backend)
	if err != nil {
		return nil, err
	}
	return tx, nil
}

// UserManagement is an auto generated Go binding around a Solidity contract.
type UserManagement struct {
	UserManagementCaller     // Read-only binding to the contract
	UserManagementTransactor // Write-only binding to the contract
	UserManagementFilterer   // Log filterer for contract events
}

// UserManagementCaller is an auto generated read-only Go binding around a Solidity contract.
type UserManagementCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// UserManagementTransactor is an auto generated write-only Go binding around a Solidity contract.
type UserManagementTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// UserManagementFilterer is an auto generated log filtering Go binding around a Solidity contract events.
type UserManagementFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// UserManagementSession is an auto generated Go binding around a Solidity contract,
// with pre-set call and transact options.
type UserManagementSession struct {
	Contract     *UserManagement   // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// UserManagementCallerSession is an auto generated read-only Go binding around a Solidity contract,
// with pre-set call options.
type UserManagementCallerSession struct {
	Contract *UserManagementCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts         // Call options to use throughout this session
}

// UserManagementTransactorSession is an auto generated write-only Go binding around a Solidity contract,
// with pre-set transact options.
type UserManagementTransactorSession struct {
	Contract     *UserManagementTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts         // Transaction auth options to use throughout this session
}

// UserManagementRaw is an auto generated low-level Go binding around a Solidity contract.
type UserManagementRaw struct {
	Contract *UserManagement // Generic contract binding to access the raw methods on
}

// UserManagementCallerRaw is an auto generated low-level read-only Go binding around a Solidity contract.
type UserManagementCallerRaw struct {
	Contract *UserManagementCaller // Generic read-only contract binding to access the raw methods on
}

// UserManagementTransactorRaw is an auto generated low-level write-only Go binding around a Solidity contract.
type UserManagementTransactorRaw struct {
	Contract *UserManagementTransactor // Generic write-only contract binding to access the raw methods on
}

// NewUserManagement creates a new instance of UserManagement, bound to a specific deployed contract.
func NewUserManagement(address common.Address, backend bind.ContractBackend) (*UserManagement, error) {
	contract, err := bindUserManagement(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &UserManagement{UserManagementCaller: UserManagementCaller{contract: contract}, UserManagementTransactor: UserManagementTransactor{contract: contract}, UserManagementFilterer: UserManagementFilterer{contract: contract}}, nil
}

// NewUserManagementCaller creates a new read-only instance of UserManagement, bound to a specific deployed contract.
func NewUserManagementCaller(address common.Address, caller bind.ContractCaller) (*UserManagementCaller, error) {
	contract, err := bindUserManagement(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &UserManagementCaller{contract: contract}, nil
}

// NewUserManagementTransactor creates a new write-only instance of UserManagement, bound to a specific deployed contract.
func NewUserManagementTransactor(address common.Address, transactor bind.ContractTransactor) (*UserManagementTransactor, error) {
	contract, err := bindUserManagement(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &UserManagementTransactor{contract: contract}, nil
}

// NewUserManagementFilterer creates a new log filterer instance of UserManagement, bound to a specific deployed contract.
func NewUserManagementFilterer(address common.Address, filterer bind.ContractFilterer) (*UserManagementFilterer, error) {
	contract, err := bindUserManagement(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &UserManagementFilterer{contract: contract}, nil
}

// bindUserManagement binds a generic wrapper to an already deployed contract.
func bindUserManagement(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(UserManagementABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_UserManagement *UserManagementRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _UserManagement.Contract.UserManagementCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_UserManagement *UserManagementRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, *types.Receipt, error) {
	return _UserManagement.Contract.UserManagementTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_UserManagement *UserManagementRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, *types.Receipt, error) {
	return _UserManagement.Contract.UserManagementTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_UserManagement *UserManagementCallerRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _UserManagement.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_UserManagement *UserManagementTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, *types.Receipt, error) {
	return _UserManagement.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_UserManagement *UserManagementTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, *types.Receipt, error) {
	return _UserManagement.Contract.contract.Transact(opts, method, params...)
}

// GetAllUsers is a free data retrieval call binding the contract method 0xe2842d79.
//
// Solidity: function getAllUsers() constant returns([]Struct0)
func (_UserManagement *UserManagementCaller) GetAllUsers(opts *bind.CallOpts) ([]Struct0, error) {
	var (
		ret0 = new([]Struct0)
	)
	out := ret0
	err := _UserManagement.contract.Call(opts, out, "getAllUsers")
	return *ret0, err
}

// GetAllUsers is a free data retrieval call binding the contract method 0xe2842d79.
//
// Solidity: function getAllUsers() constant returns([]Struct0)
func (_UserManagement *UserManagementSession) GetAllUsers() ([]Struct0, error) {
	return _UserManagement.Contract.GetAllUsers(&_UserManagement.CallOpts)
}

// GetAllUsers is a free data retrieval call binding the contract method 0xe2842d79.
//
// Solidity: function getAllUsers() constant returns([]Struct0)
func (_UserManagement *UserManagementCallerSession) GetAllUsers() ([]Struct0, error) {
	return _UserManagement.Contract.GetAllUsers(&_UserManagement.CallOpts)
}

// Users is a free data retrieval call binding the contract method 0x365b98b2.
//
// Solidity: function users(uint256 ) constant returns(string id, address addr, string name, bytes32 passwordEncoded)
func (_UserManagement *UserManagementCaller) Users(opts *bind.CallOpts, arg0 *big.Int) (struct {
	Id              string
	Addr            common.Address
	Name            string
	PasswordEncoded [32]byte
}, error) {
	ret := new(struct {
		Id              string
		Addr            common.Address
		Name            string
		PasswordEncoded [32]byte
	})
	out := ret
	err := _UserManagement.contract.Call(opts, out, "users", arg0)
	return *ret, err
}

// Users is a free data retrieval call binding the contract method 0x365b98b2.
//
// Solidity: function users(uint256 ) constant returns(string id, address addr, string name, bytes32 passwordEncoded)
func (_UserManagement *UserManagementSession) Users(arg0 *big.Int) (struct {
	Id              string
	Addr            common.Address
	Name            string
	PasswordEncoded [32]byte
}, error) {
	return _UserManagement.Contract.Users(&_UserManagement.CallOpts, arg0)
}

// Users is a free data retrieval call binding the contract method 0x365b98b2.
//
// Solidity: function users(uint256 ) constant returns(string id, address addr, string name, bytes32 passwordEncoded)
func (_UserManagement *UserManagementCallerSession) Users(arg0 *big.Int) (struct {
	Id              string
	Addr            common.Address
	Name            string
	PasswordEncoded [32]byte
}, error) {
	return _UserManagement.Contract.Users(&_UserManagement.CallOpts, arg0)
}

// AddUser is a paid mutator transaction binding the contract method 0x431a93f7.
//
// Solidity: function addUser(string id, address addr, string name, bytes32 passwordEncoded) returns(bool)
func (_UserManagement *UserManagementTransactor) AddUser(opts *bind.TransactOpts, id string, addr common.Address, name string, passwordEncoded [32]byte) (*types.Transaction, *types.Receipt, error) {
	return _UserManagement.contract.Transact(opts, "addUser", id, addr, name, passwordEncoded)
}

func (_UserManagement *UserManagementTransactor) AsyncAddUser(handler func(*types.Receipt, error), opts *bind.TransactOpts, id string, addr common.Address, name string, passwordEncoded [32]byte) (*types.Transaction, error) {
	return _UserManagement.contract.AsyncTransact(opts, handler, "addUser", id, addr, name, passwordEncoded)
}

// AddUser is a paid mutator transaction binding the contract method 0x431a93f7.
//
// Solidity: function addUser(string id, address addr, string name, bytes32 passwordEncoded) returns(bool)
func (_UserManagement *UserManagementSession) AddUser(id string, addr common.Address, name string, passwordEncoded [32]byte) (*types.Transaction, *types.Receipt, error) {
	return _UserManagement.Contract.AddUser(&_UserManagement.TransactOpts, id, addr, name, passwordEncoded)
}

func (_UserManagement *UserManagementSession) AsyncAddUser(handler func(*types.Receipt, error), id string, addr common.Address, name string, passwordEncoded [32]byte) (*types.Transaction, error) {
	return _UserManagement.Contract.AsyncAddUser(handler, &_UserManagement.TransactOpts, id, addr, name, passwordEncoded)
}

// AddUser is a paid mutator transaction binding the contract method 0x431a93f7.
//
// Solidity: function addUser(string id, address addr, string name, bytes32 passwordEncoded) returns(bool)
func (_UserManagement *UserManagementTransactorSession) AddUser(id string, addr common.Address, name string, passwordEncoded [32]byte) (*types.Transaction, *types.Receipt, error) {
	return _UserManagement.Contract.AddUser(&_UserManagement.TransactOpts, id, addr, name, passwordEncoded)
}

func (_UserManagement *UserManagementTransactorSession) AsyncAddUser(handler func(*types.Receipt, error), id string, addr common.Address, name string, passwordEncoded [32]byte) (*types.Transaction, error) {
	return _UserManagement.Contract.AsyncAddUser(handler, &_UserManagement.TransactOpts, id, addr, name, passwordEncoded)
}
