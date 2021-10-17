// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package testcontract

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

// TestcontractMetaData contains all meta data concerning the Testcontract contract.
var TestcontractMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"string\",\"name\":\"key\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"value\",\"type\":\"string\"}],\"name\":\"add\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"key\",\"type\":\"string\"}],\"name\":\"get\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"}]",
	Bin: "0x608060405234801561001057600080fd5b506105c0806100206000396000f3fe608060405234801561001057600080fd5b50600436106100365760003560e01c8063693ec85e1461003b578063ebdf86ca1461006b575b600080fd5b6100556004803603810190610050919061036b565b610087565b604051610062919061043c565b60405180910390f35b6100856004803603810190610080919061045e565b610137565b005b60606001826040516100999190610512565b908152602001604051809103902080546100b290610558565b80601f01602080910402602001604051908101604052809291908181526020018280546100de90610558565b801561012b5780601f106101005761010080835404028352916020019161012b565b820191906000526020600020905b81548152906001019060200180831161010e57829003601f168201915b50505050509050919050565b806001836040516101489190610512565b9081526020016040518091039020908051906020019061016992919061016e565b505050565b82805461017a90610558565b90600052602060002090601f01602090048101928261019c57600085556101e3565b82601f106101b557805160ff19168380011785556101e3565b828001600101855582156101e3579182015b828111156101e25782518255916020019190600101906101c7565b5b5090506101f091906101f4565b5090565b5b8082111561020d5760008160009055506001016101f5565b5090565b6000604051905090565b600080fd5b600080fd5b600080fd5b600080fd5b6000601f19601f8301169050919050565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052604160045260246000fd5b6102788261022f565b810181811067ffffffffffffffff8211171561029757610296610240565b5b80604052505050565b60006102aa610211565b90506102b6828261026f565b919050565b600067ffffffffffffffff8211156102d6576102d5610240565b5b6102df8261022f565b9050602081019050919050565b82818337600083830152505050565b600061030e610309846102bb565b6102a0565b90508281526020810184848401111561032a5761032961022a565b5b6103358482856102ec565b509392505050565b600082601f83011261035257610351610225565b5b81356103628482602086016102fb565b91505092915050565b6000602082840312156103815761038061021b565b5b600082013567ffffffffffffffff81111561039f5761039e610220565b5b6103ab8482850161033d565b91505092915050565b600081519050919050565b600082825260208201905092915050565b60005b838110156103ee5780820151818401526020810190506103d3565b838111156103fd576000848401525b50505050565b600061040e826103b4565b61041881856103bf565b93506104288185602086016103d0565b6104318161022f565b840191505092915050565b600060208201905081810360008301526104568184610403565b905092915050565b600080604083850312156104755761047461021b565b5b600083013567ffffffffffffffff81111561049357610492610220565b5b61049f8582860161033d565b925050602083013567ffffffffffffffff8111156104c0576104bf610220565b5b6104cc8582860161033d565b9150509250929050565b600081905092915050565b60006104ec826103b4565b6104f681856104d6565b93506105068185602086016103d0565b80840191505092915050565b600061051e82846104e1565b915081905092915050565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052602260045260246000fd5b6000600282049050600182168061057057607f821691505b6020821081141561058457610583610529565b5b5091905056fea2646970667358221220ea2bfa716fad240abe5a741a49d75473b58febd2d58bd8064e8f810a0b17bd4464736f6c63430008090033",
}

// TestcontractABI is the input ABI used to generate the binding from.
// Deprecated: Use TestcontractMetaData.ABI instead.
var TestcontractABI = TestcontractMetaData.ABI

// TestcontractBin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use TestcontractMetaData.Bin instead.
var TestcontractBin = TestcontractMetaData.Bin

// DeployTestcontract deploys a new Ethereum contract, binding an instance of Testcontract to it.
func DeployTestcontract(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *Testcontract, error) {
	parsed, err := TestcontractMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(TestcontractBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &Testcontract{TestcontractCaller: TestcontractCaller{contract: contract}, TestcontractTransactor: TestcontractTransactor{contract: contract}, TestcontractFilterer: TestcontractFilterer{contract: contract}}, nil
}

// Testcontract is an auto generated Go binding around an Ethereum contract.
type Testcontract struct {
	TestcontractCaller     // Read-only binding to the contract
	TestcontractTransactor // Write-only binding to the contract
	TestcontractFilterer   // Log filterer for contract events
}

// TestcontractCaller is an auto generated read-only Go binding around an Ethereum contract.
type TestcontractCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// TestcontractTransactor is an auto generated write-only Go binding around an Ethereum contract.
type TestcontractTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// TestcontractFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type TestcontractFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// TestcontractSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type TestcontractSession struct {
	Contract     *Testcontract     // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// TestcontractCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type TestcontractCallerSession struct {
	Contract *TestcontractCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts       // Call options to use throughout this session
}

// TestcontractTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type TestcontractTransactorSession struct {
	Contract     *TestcontractTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts       // Transaction auth options to use throughout this session
}

// TestcontractRaw is an auto generated low-level Go binding around an Ethereum contract.
type TestcontractRaw struct {
	Contract *Testcontract // Generic contract binding to access the raw methods on
}

// TestcontractCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type TestcontractCallerRaw struct {
	Contract *TestcontractCaller // Generic read-only contract binding to access the raw methods on
}

// TestcontractTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type TestcontractTransactorRaw struct {
	Contract *TestcontractTransactor // Generic write-only contract binding to access the raw methods on
}

// NewTestcontract creates a new instance of Testcontract, bound to a specific deployed contract.
func NewTestcontract(address common.Address, backend bind.ContractBackend) (*Testcontract, error) {
	contract, err := bindTestcontract(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Testcontract{TestcontractCaller: TestcontractCaller{contract: contract}, TestcontractTransactor: TestcontractTransactor{contract: contract}, TestcontractFilterer: TestcontractFilterer{contract: contract}}, nil
}

// NewTestcontractCaller creates a new read-only instance of Testcontract, bound to a specific deployed contract.
func NewTestcontractCaller(address common.Address, caller bind.ContractCaller) (*TestcontractCaller, error) {
	contract, err := bindTestcontract(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &TestcontractCaller{contract: contract}, nil
}

// NewTestcontractTransactor creates a new write-only instance of Testcontract, bound to a specific deployed contract.
func NewTestcontractTransactor(address common.Address, transactor bind.ContractTransactor) (*TestcontractTransactor, error) {
	contract, err := bindTestcontract(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &TestcontractTransactor{contract: contract}, nil
}

// NewTestcontractFilterer creates a new log filterer instance of Testcontract, bound to a specific deployed contract.
func NewTestcontractFilterer(address common.Address, filterer bind.ContractFilterer) (*TestcontractFilterer, error) {
	contract, err := bindTestcontract(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &TestcontractFilterer{contract: contract}, nil
}

// bindTestcontract binds a generic wrapper to an already deployed contract.
func bindTestcontract(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(TestcontractABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Testcontract *TestcontractRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Testcontract.Contract.TestcontractCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Testcontract *TestcontractRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Testcontract.Contract.TestcontractTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Testcontract *TestcontractRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Testcontract.Contract.TestcontractTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Testcontract *TestcontractCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Testcontract.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Testcontract *TestcontractTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Testcontract.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Testcontract *TestcontractTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Testcontract.Contract.contract.Transact(opts, method, params...)
}

// Get is a free data retrieval call binding the contract method 0x693ec85e.
//
// Solidity: function get(string key) view returns(string)
func (_Testcontract *TestcontractCaller) Get(opts *bind.CallOpts, key string) (string, error) {
	var out []interface{}
	err := _Testcontract.contract.Call(opts, &out, "get", key)

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// Get is a free data retrieval call binding the contract method 0x693ec85e.
//
// Solidity: function get(string key) view returns(string)
func (_Testcontract *TestcontractSession) Get(key string) (string, error) {
	return _Testcontract.Contract.Get(&_Testcontract.CallOpts, key)
}

// Get is a free data retrieval call binding the contract method 0x693ec85e.
//
// Solidity: function get(string key) view returns(string)
func (_Testcontract *TestcontractCallerSession) Get(key string) (string, error) {
	return _Testcontract.Contract.Get(&_Testcontract.CallOpts, key)
}

// Add is a paid mutator transaction binding the contract method 0xebdf86ca.
//
// Solidity: function add(string key, string value) returns()
func (_Testcontract *TestcontractTransactor) Add(opts *bind.TransactOpts, key string, value string) (*types.Transaction, error) {
	return _Testcontract.contract.Transact(opts, "add", key, value)
}

// Add is a paid mutator transaction binding the contract method 0xebdf86ca.
//
// Solidity: function add(string key, string value) returns()
func (_Testcontract *TestcontractSession) Add(key string, value string) (*types.Transaction, error) {
	return _Testcontract.Contract.Add(&_Testcontract.TransactOpts, key, value)
}

// Add is a paid mutator transaction binding the contract method 0xebdf86ca.
//
// Solidity: function add(string key, string value) returns()
func (_Testcontract *TestcontractTransactorSession) Add(key string, value string) (*types.Transaction, error) {
	return _Testcontract.Contract.Add(&_Testcontract.TransactOpts, key, value)
}
